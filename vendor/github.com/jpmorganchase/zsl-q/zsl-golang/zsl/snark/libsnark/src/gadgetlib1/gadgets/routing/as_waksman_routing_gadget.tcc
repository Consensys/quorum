/** @file
 *****************************************************************************

 Implementation of interfaces for the AS-Waksman routing gadget.

 See as_waksman_routing_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef AS_WAKSMAN_ROUTING_GADGET_TCC_
#define AS_WAKSMAN_ROUTING_GADGET_TCC_

#include <algorithm>

#include "common/routing_algorithms/as_waksman_routing_algorithm.hpp"
#include "common/profiling.hpp"

namespace libsnark {

template<typename FieldT>
as_waksman_routing_gadget<FieldT>::as_waksman_routing_gadget(protoboard<FieldT> &pb,
                                                             const size_t num_packets,
                                                             const std::vector<pb_variable_array<FieldT> > &routing_input_bits,
                                                             const std::vector<pb_variable_array<FieldT> > &routing_output_bits,
                                                             const std::string& annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    num_packets(num_packets),
    num_columns(as_waksman_num_columns(num_packets)),
    routing_input_bits(routing_input_bits),
    routing_output_bits(routing_output_bits),
    packet_size(routing_input_bits[0].size()),
    num_subpackets(div_ceil(packet_size, FieldT::capacity()))
{
    neighbors = generate_as_waksman_topology(num_packets);
    routed_packets.resize(num_columns+1);

    /* Two pass allocation. First allocate LHS packets, then for every
       switch either copy over the variables from previously allocated
       to allocate target packets */
    routed_packets[0].resize(num_packets);
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        routed_packets[0][packet_idx].allocate(pb, num_subpackets, FMT(annotation_prefix, " routed_packets_0_%zu", packet_idx));
    }

    for (size_t column_idx = 0; column_idx < num_columns; ++column_idx)
    {
        routed_packets[column_idx+1].resize(num_packets);

        for (size_t row_idx = 0; row_idx < num_packets; ++row_idx)
        {
            if (neighbors[column_idx][row_idx].first == neighbors[column_idx][row_idx].second)
            {
                /* This is a straight edge, so just copy over the previously allocated subpackets */
                routed_packets[column_idx+1][neighbors[column_idx][row_idx].first] = routed_packets[column_idx][row_idx];
            }
            else
            {
                const size_t straight_edge = neighbors[column_idx][row_idx].first;
                const size_t cross_edge = neighbors[column_idx][row_idx].second;
                routed_packets[column_idx+1][straight_edge].allocate(pb, num_subpackets, FMT(annotation_prefix, " routed_packets_%zu_%zu", column_idx+1, straight_edge));
                routed_packets[column_idx+1][cross_edge].allocate(pb, num_subpackets, FMT(annotation_prefix, " routed_packets_%zu_%zu", column_idx+1, cross_edge));
                ++row_idx; /* skip the next idx, as it to refers to the same packets */
            }
        }
    }

    /* create packing/unpacking gadgets */
    pack_inputs.reserve(num_packets); unpack_outputs.reserve(num_packets);
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        pack_inputs.emplace_back(
            multipacking_gadget<FieldT>(pb,
                                        pb_variable_array<FieldT>(routing_input_bits[packet_idx].begin(), routing_input_bits[packet_idx].end()),
                                        routed_packets[0][packet_idx],
                                        FieldT::capacity(),
                                        FMT(this->annotation_prefix, " pack_inputs_%zu", packet_idx)));
        unpack_outputs.emplace_back(
            multipacking_gadget<FieldT>(pb,
                                        pb_variable_array<FieldT>(routing_output_bits[packet_idx].begin(), routing_output_bits[packet_idx].end()),
                                        routed_packets[num_columns][packet_idx],
                                        FieldT::capacity(),
                                        FMT(this->annotation_prefix, " unpack_outputs_%zu", packet_idx)));
    }

    /* allocate switch bits */
    if (num_subpackets > 1)
    {
        asw_switch_bits.resize(num_columns);

        for (size_t column_idx = 0; column_idx < num_columns; ++column_idx)
        {
            for (size_t row_idx = 0; row_idx < num_packets; ++row_idx)
            {
                if (neighbors[column_idx][row_idx].first != neighbors[column_idx][row_idx].second)
                {
                    asw_switch_bits[column_idx][row_idx].allocate(pb, FMT(annotation_prefix, " asw_switch_bits_%zu_%zu", column_idx, row_idx));
                    ++row_idx; /* next row_idx corresponds to the same switch, so skip it */
                }
            }
        }
    }
}

template<typename FieldT>
void as_waksman_routing_gadget<FieldT>::generate_r1cs_constraints()
{
    /* packing/unpacking */
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        pack_inputs[packet_idx].generate_r1cs_constraints(false);
        unpack_outputs[packet_idx].generate_r1cs_constraints(true);
    }

    /* actual routing constraints */
    for (size_t column_idx = 0; column_idx < num_columns; ++column_idx)
    {
        for (size_t row_idx = 0; row_idx < num_packets; ++row_idx)
        {
            if (neighbors[column_idx][row_idx].first == neighbors[column_idx][row_idx].second)
            {
                /* if there is no switch at this position, then just continue with next row_idx */
                continue;
            }

            if (num_subpackets == 1)
            {
                /* easy case: require that
                   (cur-straight_edge)*(cur-cross_edge) = 0 for both
                   switch inputs */
                for (size_t switch_input : { row_idx, row_idx+1 })
                {
                    const size_t straight_edge = neighbors[column_idx][switch_input].first;
                    const size_t cross_edge = neighbors[column_idx][switch_input].second;

                    this->pb.add_r1cs_constraint(
                        r1cs_constraint<FieldT>(routed_packets[column_idx][switch_input][0] - routed_packets[column_idx+1][straight_edge][0],
                                                routed_packets[column_idx][switch_input][0] - routed_packets[column_idx+1][cross_edge][0],
                                                0),
                        FMT(this->annotation_prefix, " easy_route_%zu_%zu", column_idx, switch_input));
                }
            }
            else
            {
                /* require switching bit to be boolean */
                generate_boolean_r1cs_constraint<FieldT>(this->pb, asw_switch_bits[column_idx][row_idx],
                                                         FMT(this->annotation_prefix, " asw_switch_bits_%zu_%zu", column_idx, row_idx));

                /* route forward according to the switch bit */
                for (size_t subpacket_idx = 0; subpacket_idx < num_subpackets; ++subpacket_idx)
                {
                    /*
                      (1-switch_bit) * (cur-straight_edge) + switch_bit * (cur-cross_edge) = 0
                      switch_bit * (cross_edge-straight_edge) = cur-straight_edge
                     */
                    for (size_t switch_input : { row_idx, row_idx+1 })
                    {
                        const size_t straight_edge = neighbors[column_idx][switch_input].first;
                        const size_t cross_edge = neighbors[column_idx][switch_input].second;

                        this->pb.add_r1cs_constraint(
                            r1cs_constraint<FieldT>(
                                asw_switch_bits[column_idx][row_idx],
                                routed_packets[column_idx+1][cross_edge][subpacket_idx] - routed_packets[column_idx+1][straight_edge][subpacket_idx],
                                routed_packets[column_idx][switch_input][subpacket_idx] - routed_packets[column_idx+1][straight_edge][subpacket_idx]),
                            FMT(this->annotation_prefix, " route_forward_%zu_%zu_%zu", column_idx, switch_input, subpacket_idx));
                    }
                }
            }

            /* we processed both switch inputs at once, so skip the next iteration */
            ++row_idx;
        }
    }
}

template<typename FieldT>
void as_waksman_routing_gadget<FieldT>::generate_r1cs_witness(const integer_permutation& permutation)
{
    /* pack inputs */
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        pack_inputs[packet_idx].generate_r1cs_witness_from_bits();
    }

    /* do the routing */
    as_waksman_routing routing = get_as_waksman_routing(permutation);

    for (size_t column_idx = 0; column_idx < num_columns; ++column_idx)
    {
        for (size_t row_idx = 0; row_idx < num_packets; ++row_idx)
        {
            if (neighbors[column_idx][row_idx].first == neighbors[column_idx][row_idx].second)
            {
                /* this is a straight edge, so just pass the values forward */
                const size_t next = neighbors[column_idx][row_idx].first;

                for (size_t subpacket_idx = 0; subpacket_idx < num_subpackets; ++subpacket_idx)
                {
                    this->pb.val(routed_packets[column_idx+1][next][subpacket_idx]) = this->pb.val(routed_packets[column_idx][row_idx][subpacket_idx]);
                }
            }
            else
            {
                if (num_subpackets > 1)
                {
                    /* update the switch bit */
                    this->pb.val(asw_switch_bits[column_idx][row_idx]) = FieldT(routing[column_idx][row_idx] ? 1 : 0);
                }

                /* route according to the switch bit */
                const bool switch_val = routing[column_idx][row_idx];

                for (size_t switch_input : { row_idx, row_idx+1 })
                {
                    const size_t straight_edge = neighbors[column_idx][switch_input].first;
                    const size_t cross_edge = neighbors[column_idx][switch_input].second;

                    const size_t switched_edge = (switch_val ? cross_edge : straight_edge);

                    for (size_t subpacket_idx = 0; subpacket_idx < num_subpackets; ++subpacket_idx)
                    {
                        this->pb.val(routed_packets[column_idx+1][switched_edge][subpacket_idx]) = this->pb.val(routed_packets[column_idx][switch_input][subpacket_idx]);
                    }
                }

                /* we processed both switch inputs at once, so skip the next iteration */
                ++row_idx;
            }
        }
    }

    /* unpack outputs */
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        unpack_outputs[packet_idx].generate_r1cs_witness_from_packed();
    }
}

template<typename FieldT>
void test_as_waksman_routing_gadget(const size_t num_packets, const size_t packet_size)
{
    printf("testing as_waksman_routing_gadget by routing %zu element vector of %zu bits (Fp fits all %zu bit integers)\n", num_packets, packet_size, FieldT::capacity());
    protoboard<FieldT> pb;
    integer_permutation permutation(num_packets);
    permutation.random_shuffle();
    print_time("generated permutation");

    std::vector<pb_variable_array<FieldT> > randbits(num_packets), outbits(num_packets);
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        randbits[packet_idx].allocate(pb, packet_size, FMT("", "randbits_%zu", packet_idx));
        outbits[packet_idx].allocate(pb, packet_size, FMT("", "outbits_%zu", packet_idx));

        for (size_t bit_idx = 0; bit_idx < packet_size; ++bit_idx)
        {
            pb.val(randbits[packet_idx][bit_idx]) = (rand() % 2) ? FieldT::one() : FieldT::zero();
        }
    }
    print_time("generated bits to be routed");

    as_waksman_routing_gadget<FieldT> r(pb, num_packets, randbits, outbits, "main_routing_gadget");
    r.generate_r1cs_constraints();
    print_time("generated routing constraints");

    r.generate_r1cs_witness(permutation);
    print_time("generated routing assignment");

    printf("positive test\n");
    assert(pb.is_satisfied());
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        for (size_t bit_idx = 0; bit_idx < packet_size; ++bit_idx)
        {
            assert(pb.val(outbits[permutation.get(packet_idx)][bit_idx]) == pb.val(randbits[packet_idx][bit_idx]));
        }
    }

    printf("negative test\n");
    pb.val(pb_variable<FieldT>(10)) = FieldT(12345);
    assert(!pb.is_satisfied());

    printf("num_constraints = %zu, num_variables = %zu\n",
           pb.num_constraints(),
           pb.constraint_system.num_variables);
}

} // libsnark

#endif // AS_WAKSMAN_ROUTING_GADGET_TCC_

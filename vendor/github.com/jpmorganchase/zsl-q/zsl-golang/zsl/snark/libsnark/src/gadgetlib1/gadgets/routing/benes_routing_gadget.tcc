/** @file
 *****************************************************************************

 Implementation of interfaces for the Benes routing gadget.

 See benes_routing_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BENES_ROUTING_GADGET_TCC_
#define BENES_ROUTING_GADGET_TCC_

#include "common/profiling.hpp"

#include <algorithm>

namespace libsnark {

template<typename FieldT>
benes_routing_gadget<FieldT>::benes_routing_gadget(protoboard<FieldT> &pb,
                                                   const size_t num_packets,
                                                   const std::vector<pb_variable_array<FieldT> > &routing_input_bits,
                                                   const std::vector<pb_variable_array<FieldT> > &routing_output_bits,
                                                   const size_t lines_to_unpack,
                                                   const std::string& annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    num_packets(num_packets),
    num_columns(benes_num_columns(num_packets)),
    routing_input_bits(routing_input_bits),
    routing_output_bits(routing_output_bits),
    lines_to_unpack(lines_to_unpack),
    packet_size(routing_input_bits[0].size()),
    num_subpackets(div_ceil(packet_size, FieldT::capacity()))
{
    assert(lines_to_unpack <= routing_input_bits.size());
    assert(num_packets == 1ul<<log2(num_packets));
    assert(routing_input_bits.size() == num_packets);

    neighbors = generate_benes_topology(num_packets);

    routed_packets.resize(num_columns+1);
    for (size_t column_idx = 0; column_idx <= num_columns; ++column_idx)
    {
        routed_packets[column_idx].resize(num_packets);
        for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
        {
            routed_packets[column_idx][packet_idx].allocate(pb, num_subpackets, FMT(annotation_prefix, " routed_packets_%zu_%zu", column_idx, packet_idx));
        }
    }

    pack_inputs.reserve(num_packets);
    unpack_outputs.reserve(num_packets);

    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        pack_inputs.emplace_back(
            multipacking_gadget<FieldT>(pb,
                                        pb_variable_array<FieldT>(routing_input_bits[packet_idx].begin(), routing_input_bits[packet_idx].end()),
                                        routed_packets[0][packet_idx],
                                        FieldT::capacity(),
                                        FMT(this->annotation_prefix, " pack_inputs_%zu", packet_idx)));
        if (packet_idx < lines_to_unpack)
        {
            unpack_outputs.emplace_back(
                multipacking_gadget<FieldT>(pb,
                                            pb_variable_array<FieldT>(routing_output_bits[packet_idx].begin(), routing_output_bits[packet_idx].end()),
                                            routed_packets[num_columns][packet_idx],
                                            FieldT::capacity(),
                                            FMT(this->annotation_prefix, " unpack_outputs_%zu", packet_idx)));
        }
    }

    if (num_subpackets > 1)
    {
        benes_switch_bits.resize(num_columns);
        for (size_t column_idx = 0; column_idx < num_columns; ++column_idx)
        {
            benes_switch_bits[column_idx].allocate(pb, num_packets, FMT(this->annotation_prefix, " benes_switch_bits_%zu", column_idx));
        }
    }
}

template<typename FieldT>
void benes_routing_gadget<FieldT>::generate_r1cs_constraints()
{
    /* packing/unpacking */
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        pack_inputs[packet_idx].generate_r1cs_constraints(false);
        if (packet_idx < lines_to_unpack)
        {
            unpack_outputs[packet_idx].generate_r1cs_constraints(true);
        }
        else
        {
            for (size_t subpacket_idx = 0; subpacket_idx < num_subpackets; ++subpacket_idx)
            {
                this->pb.add_r1cs_constraint(
                    r1cs_constraint<FieldT>(1, routed_packets[0][packet_idx][subpacket_idx], routed_packets[num_columns][packet_idx][subpacket_idx]),
                    FMT(this->annotation_prefix, " fix_line_%zu_subpacket_%zu", packet_idx, subpacket_idx));
            }
        }
    }

    /* actual routing constraints */
    for (size_t column_idx = 0; column_idx < num_columns; ++column_idx)
    {
        for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
        {
            const size_t straight_edge = neighbors[column_idx][packet_idx].first;
            const size_t cross_edge = neighbors[column_idx][packet_idx].second;

            if (num_subpackets == 1)
            {
                /* easy case: (cur-next)*(cur-cross) = 0 */
                this->pb.add_r1cs_constraint(
                    r1cs_constraint<FieldT>(
                        routed_packets[column_idx][packet_idx][0] - routed_packets[column_idx+1][straight_edge][0],
                        routed_packets[column_idx][packet_idx][0] - routed_packets[column_idx+1][cross_edge][0],
                        0),
                    FMT(this->annotation_prefix, " easy_route_%zu_%zu", column_idx, packet_idx));
            }
            else
            {
                /* routing bit must be boolean */
                generate_boolean_r1cs_constraint<FieldT>(this->pb, benes_switch_bits[column_idx][packet_idx],
                                                         FMT(this->annotation_prefix, " routing_bit_%zu_%zu", column_idx, packet_idx));

                /* route forward according to routing bits */
                for (size_t subpacket_idx = 0; subpacket_idx < num_subpackets; ++subpacket_idx)
                {
                    /*
                      (1-switch_bit) * (cur-straight_edge) + switch_bit * (cur-cross_edge) = 0
                      switch_bit * (cross_edge-straight_edge) = cur-straight_edge
                    */
                    this->pb.add_r1cs_constraint(
                        r1cs_constraint<FieldT>(
                            benes_switch_bits[column_idx][packet_idx],
                            routed_packets[column_idx+1][cross_edge][subpacket_idx] - routed_packets[column_idx+1][straight_edge][subpacket_idx],
                            routed_packets[column_idx][packet_idx][subpacket_idx] - routed_packets[column_idx+1][straight_edge][subpacket_idx]),
                        FMT(this->annotation_prefix, " route_forward_%zu_%zu_%zu", column_idx, packet_idx, subpacket_idx));
                }
            }
        }
    }
}

template<typename FieldT>
void benes_routing_gadget<FieldT>::generate_r1cs_witness(const integer_permutation& permutation)
{
    /* pack inputs */
    for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
    {
        pack_inputs[packet_idx].generate_r1cs_witness_from_bits();
    }

    /* do the routing */
    const benes_routing routing = get_benes_routing(permutation);

    for (size_t column_idx = 0; column_idx < num_columns; ++column_idx)
    {
        for (size_t packet_idx = 0; packet_idx < num_packets; ++packet_idx)
        {
            const size_t straight_edge = neighbors[column_idx][packet_idx].first;
            const size_t cross_edge = neighbors[column_idx][packet_idx].second;

            if (num_subpackets > 1)
            {
                this->pb.val(benes_switch_bits[column_idx][packet_idx]) = FieldT(routing[column_idx][packet_idx] ? 1 : 0);
            }

            for (size_t subpacket_idx = 0; subpacket_idx < num_subpackets; ++subpacket_idx)
            {
                this->pb.val(routing[column_idx][packet_idx] ?
                             routed_packets[column_idx+1][cross_edge][subpacket_idx] :
                             routed_packets[column_idx+1][straight_edge][subpacket_idx]) =
                    this->pb.val(routed_packets[column_idx][packet_idx][subpacket_idx]);
            }
        }
    }

    /* unpack outputs */
    for (size_t packet_idx = 0; packet_idx < lines_to_unpack; ++packet_idx)
    {
        unpack_outputs[packet_idx].generate_r1cs_witness_from_packed();
    }
}

template<typename FieldT>
void test_benes_routing_gadget(const size_t num_packets, const size_t packet_size)
{
    const size_t dimension = log2(num_packets);
    assert(num_packets == 1ul<<dimension);

    printf("testing benes_routing_gadget by routing 2^%zu-entry vector of %zu bits (Fp fits all %zu bit integers)\n", dimension, packet_size, FieldT::capacity());

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

    benes_routing_gadget<FieldT> r(pb, num_packets, randbits, outbits, num_packets, "main_routing_gadget");
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

#endif // BENES_ROUTING_GADGET_TCC_

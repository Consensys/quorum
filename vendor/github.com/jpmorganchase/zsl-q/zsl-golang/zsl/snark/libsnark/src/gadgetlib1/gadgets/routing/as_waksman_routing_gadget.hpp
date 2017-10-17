/** @file
 *****************************************************************************

 Declaration of interfaces for the AS-Waksman routing gadget.

 The gadget verifies that the outputs are a permutation of the inputs,
 by use of an AS-Waksman network.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef AS_WAKSMAN_ROUTING_GADGET_HPP_
#define AS_WAKSMAN_ROUTING_GADGET_HPP_

#include "gadgetlib1/protoboard.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "common/data_structures/integer_permutation.hpp"
#include "common/routing_algorithms/as_waksman_routing_algorithm.hpp"

namespace libsnark {

template<typename FieldT>
class as_waksman_routing_gadget : public gadget<FieldT> {
private:
    /*
      Indexing conventions:

      routed_packets[column_idx][packet_idx][subpacket_idx]
      pack_inputs/unpack_outputs[packet_idx]
      asw_switch_bits[column_idx][row_idx]

      Where column_idx ranges is in range 0 .. width and packet_idx is
      in range 0 .. num_packets-1.

      Note that unlike in Bene\v{s} routing networks row_idx are
      *not* necessarily consecutive; similarly for straight edges
      routed_packets[column_idx][packet_idx] will *reuse* previously
      allocated variables.

    */
    std::vector<std::vector<pb_variable_array<FieldT> > > routed_packets;
    std::vector<multipacking_gadget<FieldT> > pack_inputs, unpack_outputs;

    /*
      If #packets = 1 then we can route without explicit switch bits
      (and save half the constraints); in this case asw_switch_bits will
      be unused.

      For asw_switch_bits 0 corresponds to switch off (straight
      connection), and 1 corresponds to switch on (crossed
      connection).
    */
    std::vector<std::map<size_t, pb_variable<FieldT> > > asw_switch_bits;
    as_waksman_topology neighbors;
public:
    const size_t num_packets;
    const size_t num_columns;
    const std::vector<pb_variable_array<FieldT>> routing_input_bits;
    const std::vector<pb_variable_array<FieldT>> routing_output_bits;

    const size_t packet_size, num_subpackets;

    as_waksman_routing_gadget(protoboard<FieldT> &pb,
                              const size_t num_packets,
                              const std::vector<pb_variable_array<FieldT>> &routing_input_bits,
                              const std::vector<pb_variable_array<FieldT>> &routing_output_bits,
                              const std::string& annotation_prefix="");
    void generate_r1cs_constraints();
    void generate_r1cs_witness(const integer_permutation& permutation);
};

template<typename FieldT>
void test_as_waksman_routing_gadget(const size_t num_packets, const size_t packet_size);

} // libsnark

#include "gadgetlib1/gadgets/routing/as_waksman_routing_gadget.tcc"

#endif // AS_WAKSMAN_ROUTING_GADGET_HPP_

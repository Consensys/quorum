/** @file
 *****************************************************************************

 Declaration of interfaces for the Benes routing gadget.

 The gadget verifies that the outputs are a permutation of the inputs,
 by use of a Benes network.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BENES_ROUTING_GADGET_HPP_
#define BENES_ROUTING_GADGET_HPP_

#include "common/data_structures/integer_permutation.hpp"
#include "common/routing_algorithms/benes_routing_algorithm.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "gadgetlib1/protoboard.hpp"

namespace libsnark {

template<typename FieldT>
class benes_routing_gadget : public gadget<FieldT> {
private:
    /*
      Indexing conventions:

      routed_packets[column_idx][packet_idx][subpacket_idx]
      pack_inputs/unpack_outputs[packet_idx]
      benes_switch_bits[column_idx][row_idx]

      Where column_idx ranges is in range 0 .. 2*dimension
      (2*dimension-1 for switch bits/topology) and packet_idx is in
      range 0 .. num_packets-1.
    */
    std::vector<std::vector<pb_variable_array<FieldT> > > routed_packets;
    std::vector<multipacking_gadget<FieldT> > pack_inputs, unpack_outputs;

    /*
      If #packets = 1 then we can route without explicit routing bits
      (and save half the constraints); in this case benes_switch_bits will
      be unused.

      For benes_switch_bits 0 corresponds to straight edge and 1
      corresponds to cross edge.
    */
    std::vector<pb_variable_array<FieldT>> benes_switch_bits;
    benes_topology neighbors;
public:
    const size_t num_packets;
    const size_t num_columns;

    const std::vector<pb_variable_array<FieldT> > routing_input_bits;
    const std::vector<pb_variable_array<FieldT> > routing_output_bits;
    size_t lines_to_unpack;

    const size_t packet_size, num_subpackets;

    benes_routing_gadget(protoboard<FieldT> &pb,
                         const size_t num_packets,
                         const std::vector<pb_variable_array<FieldT>> &routing_input_bits,
                         const std::vector<pb_variable_array<FieldT>> &routing_output_bits,
                         const size_t lines_to_unpack,
                         const std::string& annotation_prefix="");

    void generate_r1cs_constraints();

    void generate_r1cs_witness(const integer_permutation &permutation);
};

template<typename FieldT>
void test_benes_routing_gadget(const size_t num_packets, const size_t packet_size);

} // libsnark

#include "gadgetlib1/gadgets/routing/benes_routing_gadget.tcc"

#endif // BENES_ROUTING_GADGET_HPP_

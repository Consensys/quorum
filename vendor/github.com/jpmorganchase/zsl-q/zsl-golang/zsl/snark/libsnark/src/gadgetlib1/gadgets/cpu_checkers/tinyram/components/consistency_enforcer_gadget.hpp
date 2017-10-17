/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM consistency enforcer gadget.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef CONSISTENCY_ENFORCER_GADGET_HPP_
#define CONSISTENCY_ENFORCER_GADGET_HPP_

namespace libsnark {

template<typename FieldT>
class consistency_enforcer_gadget : public tinyram_standard_gadget<FieldT>  {
private:
    pb_variable<FieldT> is_register_instruction;
    pb_variable<FieldT> is_control_flow_instruction;
    pb_variable<FieldT> is_stall_instruction;

    pb_variable<FieldT> packed_desidx;
    std::shared_ptr<packing_gadget<FieldT> > pack_desidx;

    pb_variable<FieldT> computed_result;
    pb_variable<FieldT> computed_flag;
    std::shared_ptr<inner_product_gadget<FieldT> > compute_computed_result;
    std::shared_ptr<inner_product_gadget<FieldT> > compute_computed_flag;

    pb_variable<FieldT> pc_from_cf_or_zero;

    std::shared_ptr<loose_multiplexing_gadget<FieldT> > demux_packed_outgoing_desval;
public:
    pb_variable_array<FieldT> opcode_indicators;
    pb_variable_array<FieldT> instruction_results;
    pb_variable_array<FieldT> instruction_flags;
    pb_variable_array<FieldT> desidx;
    pb_variable<FieldT> packed_incoming_pc;
    pb_variable_array<FieldT> packed_incoming_registers;
    pb_variable<FieldT> packed_incoming_desval;
    pb_variable<FieldT> incoming_flag;
    pb_variable<FieldT> packed_outgoing_pc;
    pb_variable_array<FieldT> packed_outgoing_registers;
    pb_variable<FieldT> outgoing_flag;
    pb_variable<FieldT> packed_outgoing_desval;

    consistency_enforcer_gadget(tinyram_protoboard<FieldT> &pb,
                                const pb_variable_array<FieldT> &opcode_indicators,
                                const pb_variable_array<FieldT> &instruction_results,
                                const pb_variable_array<FieldT> &instruction_flags,
                                const pb_variable_array<FieldT> &desidx,
                                const pb_variable<FieldT> &packed_incoming_pc,
                                const pb_variable_array<FieldT> &packed_incoming_registers,
                                const pb_variable<FieldT> &packed_incoming_desval,
                                const pb_variable<FieldT> &incoming_flag,
                                const pb_variable<FieldT> &packed_outgoing_pc,
                                const pb_variable_array<FieldT> &packed_outgoing_registers,
                                const pb_variable<FieldT> &outgoing_flag,
                                const std::string &annotation_prefix="");

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/consistency_enforcer_gadget.tcc"

#endif // CONSISTENCY_ENFORCER_GADGET_HPP_

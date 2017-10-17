/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM instruction packing gadget.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef INSTRUCTION_PACKING_GADGET_HPP_
#define INSTRUCTION_PACKING_GADGET_HPP_

namespace libsnark {

template<typename FieldT>
class tinyram_instruction_packing_gadget : public tinyram_gadget<FieldT> {
private:
    pb_variable_array<FieldT> all_bits;

    std::shared_ptr<packing_gadget<FieldT> > pack_instruction;
public:
    pb_variable_array<FieldT> opcode;
    pb_variable<FieldT> arg2_is_imm;
    pb_variable_array<FieldT> desidx;
    pb_variable_array<FieldT> arg1idx;
    pb_variable_array<FieldT> arg2idx;
    pb_variable<FieldT> packed_instruction;

    pb_variable_array<FieldT> dummy;

    tinyram_instruction_packing_gadget(tinyram_protoboard<FieldT> &pb,
                                       const pb_variable_array<FieldT> &opcode,
                                       const pb_variable<FieldT> &arg2_is_imm,
                                       const pb_variable_array<FieldT> &desidx,
                                       const pb_variable_array<FieldT> &arg1idx,
                                       const pb_variable_array<FieldT> &arg2idx,
                                       const pb_variable<FieldT> &packed_instruction,
                                       const std::string &annotation_prefix="");

    void generate_r1cs_constraints(const bool enforce_bitness);
    void generate_r1cs_witness_from_packed();
    void generate_r1cs_witness_from_bits();
};

template<typename FieldT>
FieldT pack_instruction(const tinyram_architecture_params &ap,
                        const tinyram_instruction &instr);

template<typename FieldT>
void test_instruction_packing();

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/instruction_packing_gadget.tcc"

#endif // INSTRUCTION_PACKING_GADGET_HPP_

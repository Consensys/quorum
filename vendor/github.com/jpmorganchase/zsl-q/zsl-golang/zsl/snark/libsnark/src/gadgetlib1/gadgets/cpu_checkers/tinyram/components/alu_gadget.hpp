/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM ALU gadget.

 The gadget checks the correct execution of a given TinyRAM instruction.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef ALU_GADGET_HPP_
#define ALU_GADGET_HPP_

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/alu_arithmetic.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/alu_control_flow.hpp"

namespace libsnark {

template<typename FieldT>
class ALU_gadget : public tinyram_standard_gadget<FieldT> {
private:
    std::vector<std::shared_ptr<tinyram_standard_gadget<FieldT> > > components;
public:
    pb_variable_array<FieldT> opcode_indicators;
    word_variable_gadget<FieldT> pc;
    word_variable_gadget<FieldT> desval;
    word_variable_gadget<FieldT> arg1val;
    word_variable_gadget<FieldT> arg2val;
    pb_variable<FieldT> flag;
    pb_variable_array<FieldT> instruction_results;
    pb_variable_array<FieldT> instruction_flags;

    ALU_gadget<FieldT>(tinyram_protoboard<FieldT> &pb,
                       const pb_variable_array<FieldT> &opcode_indicators,
                       const word_variable_gadget<FieldT> &pc,
                       const word_variable_gadget<FieldT> &desval,
                       const word_variable_gadget<FieldT> &arg1val,
                       const word_variable_gadget<FieldT> &arg2val,
                       const pb_variable<FieldT> &flag,
                       const pb_variable_array<FieldT> &instruction_results,
                       const pb_variable_array<FieldT> &instruction_flags,
                       const std::string &annotation_prefix="") :
        tinyram_standard_gadget<FieldT>(pb, annotation_prefix),
        opcode_indicators(opcode_indicators),
        pc(pc),
        desval(desval),
        arg1val(arg1val),
        arg2val(arg2val),
        flag(flag),
        instruction_results(instruction_results),
        instruction_flags(instruction_flags)
    {
        components.resize(1ul<<pb.ap.opcode_width());

        /* arithmetic */
        components[tinyram_opcode_AND].reset(
            new ALU_and_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                       instruction_results[tinyram_opcode_AND],
                                       instruction_flags[tinyram_opcode_AND],
                                       FMT(this->annotation_prefix, " AND")));

        components[tinyram_opcode_OR].reset(
            new ALU_or_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                      instruction_results[tinyram_opcode_OR],
                                      instruction_flags[tinyram_opcode_OR],
                                      FMT(this->annotation_prefix, " OR")));

        components[tinyram_opcode_XOR].reset(
            new ALU_xor_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                       instruction_results[tinyram_opcode_XOR],
                                       instruction_flags[tinyram_opcode_XOR],
                                       FMT(this->annotation_prefix, " XOR")));

        components[tinyram_opcode_NOT].reset(
            new ALU_not_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                       instruction_results[tinyram_opcode_NOT],
                                       instruction_flags[tinyram_opcode_NOT],
                                       FMT(this->annotation_prefix, " NOT")));

        components[tinyram_opcode_ADD].reset(
            new ALU_add_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                       instruction_results[tinyram_opcode_ADD],
                                       instruction_flags[tinyram_opcode_ADD],
                                       FMT(this->annotation_prefix, " ADD")));

        components[tinyram_opcode_SUB].reset(
            new ALU_sub_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                       instruction_results[tinyram_opcode_SUB],
                                       instruction_flags[tinyram_opcode_SUB],
                                       FMT(this->annotation_prefix, " SUB")));

        components[tinyram_opcode_MOV].reset(
            new ALU_mov_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                       instruction_results[tinyram_opcode_MOV],
                                       instruction_flags[tinyram_opcode_MOV],
                                       FMT(this->annotation_prefix, " MOV")));

        components[tinyram_opcode_CMOV].reset(
            new ALU_cmov_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                        instruction_results[tinyram_opcode_CMOV],
                                        instruction_flags[tinyram_opcode_CMOV],
                                        FMT(this->annotation_prefix, " CMOV")));

        components[tinyram_opcode_CMPA].reset(
            new ALU_cmp_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                       instruction_results[tinyram_opcode_CMPE],
                                       instruction_flags[tinyram_opcode_CMPE],
                                       instruction_results[tinyram_opcode_CMPA],
                                       instruction_flags[tinyram_opcode_CMPA],
                                       instruction_results[tinyram_opcode_CMPAE],
                                       instruction_flags[tinyram_opcode_CMPAE],
                                       FMT(this->annotation_prefix, " CMP_unsigned")));

        components[tinyram_opcode_CMPG].reset(
            new ALU_cmps_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                        instruction_results[tinyram_opcode_CMPG],
                                        instruction_flags[tinyram_opcode_CMPG],
                                        instruction_results[tinyram_opcode_CMPGE],
                                        instruction_flags[tinyram_opcode_CMPGE],
                                        FMT(this->annotation_prefix, " CMP_signed")));

        components[tinyram_opcode_UMULH].reset(
            new ALU_umul_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                        instruction_results[tinyram_opcode_MULL],
                                        instruction_flags[tinyram_opcode_MULL],
                                        instruction_results[tinyram_opcode_UMULH],
                                        instruction_flags[tinyram_opcode_UMULH],
                                        FMT(this->annotation_prefix, " MUL_unsigned")));

        components[tinyram_opcode_SMULH].reset(
            new ALU_smul_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                        instruction_results[tinyram_opcode_SMULH],
                                        instruction_flags[tinyram_opcode_SMULH],
                                        FMT(this->annotation_prefix, " MUL_signed")));


        components[tinyram_opcode_UDIV].reset(
            new ALU_divmod_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                          instruction_results[tinyram_opcode_UDIV],
                                          instruction_flags[tinyram_opcode_UDIV],
                                          instruction_results[tinyram_opcode_UMOD],
                                          instruction_flags[tinyram_opcode_UMOD],
                                          FMT(this->annotation_prefix, " DIV")));

        components[tinyram_opcode_SHR].reset(
            new ALU_shr_shl_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                           instruction_results[tinyram_opcode_SHR],
                                           instruction_flags[tinyram_opcode_SHR],
                                           instruction_results[tinyram_opcode_SHL],
                                           instruction_flags[tinyram_opcode_SHL],
                                           FMT(this->annotation_prefix, " SHR_SHL")));

        /* control flow */
        components[tinyram_opcode_JMP].reset(
            new ALU_jmp_gadget<FieldT>(pb, pc, arg2val, flag,
                                       instruction_results[tinyram_opcode_JMP],
                                       FMT(this->annotation_prefix, " JMP")));

        components[tinyram_opcode_CJMP].reset(
            new ALU_cjmp_gadget<FieldT>(pb, pc, arg2val, flag,
                                        instruction_results[tinyram_opcode_CJMP],
                                        FMT(this->annotation_prefix, " CJMP")));

        components[tinyram_opcode_CNJMP].reset(
            new ALU_cnjmp_gadget<FieldT>(pb, pc, arg2val, flag,
                                         instruction_results[tinyram_opcode_CNJMP],
                                         FMT(this->annotation_prefix, " CNJMP")));
    }

    void generate_r1cs_constraints();

    void generate_r1cs_witness();

};

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/alu_gadget.tcc"

#endif // ALU_GADGET_HPP_

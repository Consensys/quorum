/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM ALU control-flow gadgets.

 These gadget check the correct execution of control-flow TinyRAM instructions.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef ALU_CONTROL_FLOW_HPP_
#define ALU_CONTROL_FLOW_HPP_

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/tinyram_protoboard.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"

namespace libsnark {

/* control flow gadgets */
template<typename FieldT>
class ALU_control_flow_gadget : public tinyram_standard_gadget<FieldT> {
public:
    const word_variable_gadget<FieldT> pc;
    const word_variable_gadget<FieldT> argval2;
    const pb_variable<FieldT> flag;
    const pb_variable<FieldT> result;

    ALU_control_flow_gadget(tinyram_protoboard<FieldT> &pb,
                            const word_variable_gadget<FieldT> &pc,
                            const word_variable_gadget<FieldT> &argval2,
                            const pb_variable<FieldT> &flag,
                            const pb_variable<FieldT> &result,
                            const std::string &annotation_prefix="") :
        tinyram_standard_gadget<FieldT>(pb, annotation_prefix),
        pc(pc),
        argval2(argval2),
        flag(flag),
        result(result) {};
};

template<typename FieldT>
class ALU_jmp_gadget : public ALU_control_flow_gadget<FieldT> {
public:
    ALU_jmp_gadget(tinyram_protoboard<FieldT> &pb,
                   const word_variable_gadget<FieldT> &pc,
                   const word_variable_gadget<FieldT> &argval2,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &result,
                   const std::string &annotation_prefix="") :
        ALU_control_flow_gadget<FieldT>(pb, pc, argval2, flag, result, annotation_prefix) {}

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_jmp_gadget();

template<typename FieldT>
class ALU_cjmp_gadget : public ALU_control_flow_gadget<FieldT> {
public:
    ALU_cjmp_gadget(tinyram_protoboard<FieldT> &pb,
                    const word_variable_gadget<FieldT> &pc,
                    const word_variable_gadget<FieldT> &argval2,
                    const pb_variable<FieldT> &flag,
                    const pb_variable<FieldT> &result,
                    const std::string &annotation_prefix="") :
        ALU_control_flow_gadget<FieldT>(pb, pc, argval2, flag, result, annotation_prefix) {}

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_cjmp_gadget();

template<typename FieldT>
class ALU_cnjmp_gadget : public ALU_control_flow_gadget<FieldT> {
public:
    ALU_cnjmp_gadget(tinyram_protoboard<FieldT> &pb,
                     const word_variable_gadget<FieldT> &pc,
                     const word_variable_gadget<FieldT> &argval2,
                     const pb_variable<FieldT> &flag,
                     const pb_variable<FieldT> &result,
                     const std::string &annotation_prefix="") :
        ALU_control_flow_gadget<FieldT>(pb, pc, argval2, flag, result, annotation_prefix) {}

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_cnjmp_gadget();

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/alu_control_flow.tcc"

#endif // ALU_CONTROL_FLOW_HPP_

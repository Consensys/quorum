/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM CPU checker gadget.

 The gadget checks the correct operation for the CPU of the TinyRAM architecture.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TINYRAM_CPU_CHECKER_HPP_
#define TINYRAM_CPU_CHECKER_HPP_

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/tinyram_protoboard.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/word_variable_gadget.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/alu_gadget.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/argument_decoder_gadget.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/consistency_enforcer_gadget.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/memory_masking_gadget.hpp"

namespace libsnark {

template<typename FieldT>
class tinyram_cpu_checker : public tinyram_standard_gadget<FieldT> {
private:
    pb_variable_array<FieldT> opcode;
    pb_variable<FieldT> arg2_is_imm;
    pb_variable_array<FieldT> desidx;
    pb_variable_array<FieldT> arg1idx;
    pb_variable_array<FieldT> arg2idx;

    std::vector<word_variable_gadget<FieldT> > prev_registers;
    std::vector<word_variable_gadget<FieldT> > next_registers;
    pb_variable<FieldT> prev_flag;
    pb_variable<FieldT> next_flag;
    pb_variable<FieldT> prev_tape1_exhausted;
    pb_variable<FieldT> next_tape1_exhausted;

    std::shared_ptr<word_variable_gadget<FieldT> > prev_pc_addr_as_word_variable;
    std::shared_ptr<word_variable_gadget<FieldT> > desval;
    std::shared_ptr<word_variable_gadget<FieldT> > arg1val;
    std::shared_ptr<word_variable_gadget<FieldT> > arg2val;

    std::shared_ptr<argument_decoder_gadget<FieldT> > decode_arguments;
    pb_variable_array<FieldT> opcode_indicators;
    std::shared_ptr<ALU_gadget<FieldT> > ALU;

    std::shared_ptr<doubleword_variable_gadget<FieldT> > ls_prev_val_as_doubleword_variable;
    std::shared_ptr<doubleword_variable_gadget<FieldT> > ls_next_val_as_doubleword_variable;
    std::shared_ptr<dual_variable_gadget<FieldT> > memory_subaddress;
    pb_variable<FieldT> memory_subcontents;
    pb_linear_combination<FieldT> memory_access_is_word;
    pb_linear_combination<FieldT> memory_access_is_byte;
    std::shared_ptr<memory_masking_gadget<FieldT> > check_memory;

    std::shared_ptr<word_variable_gadget<FieldT> > next_pc_addr_as_word_variable;
    std::shared_ptr<consistency_enforcer_gadget<FieldT> > consistency_enforcer;

    pb_variable_array<FieldT> instruction_results;
    pb_variable_array<FieldT> instruction_flags;

    pb_variable<FieldT> read_not1;
public:
    pb_variable_array<FieldT> prev_pc_addr;
    pb_variable_array<FieldT> prev_pc_val;
    pb_variable_array<FieldT> prev_state;
    pb_variable_array<FieldT> ls_addr;
    pb_variable_array<FieldT> ls_prev_val;
    pb_variable_array<FieldT> ls_next_val;
    pb_variable_array<FieldT> next_state;
    pb_variable_array<FieldT> next_pc_addr;
    pb_variable<FieldT> next_has_accepted;

    tinyram_cpu_checker(tinyram_protoboard<FieldT> &pb,
                        pb_variable_array<FieldT> &prev_pc_addr,
                        pb_variable_array<FieldT> &prev_pc_val,
                        pb_variable_array<FieldT> &prev_state,
                        pb_variable_array<FieldT> &ls_addr,
                        pb_variable_array<FieldT> &ls_prev_val,
                        pb_variable_array<FieldT> &ls_next_val,
                        pb_variable_array<FieldT> &next_state,
                        pb_variable_array<FieldT> &next_pc_addr,
                        pb_variable<FieldT> &next_has_accepted,
                        const std::string &annotation_prefix);

    void generate_r1cs_constraints();
    void generate_r1cs_witness() { assert(0); }
    void generate_r1cs_witness_address();
    void generate_r1cs_witness_other(tinyram_input_tape_iterator &aux_it,
                                     const tinyram_input_tape_iterator &aux_end);
    void dump() const;
};

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/tinyram_cpu_checker.tcc"

#endif // TINYRAM_CPU_CHECKER_HPP_

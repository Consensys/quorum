/** @file
 *****************************************************************************

 Implementation of interfaces for the TinyRAM CPU checker gadget.

 See tinyram_cpu_checker.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TINYRAM_CPU_CHECKER_TCC_
#define TINYRAM_CPU_CHECKER_TCC_

#include "algebra/fields/field_utils.hpp"

namespace libsnark {

template<typename FieldT>
tinyram_cpu_checker<FieldT>::tinyram_cpu_checker(tinyram_protoboard<FieldT> &pb,
                                                 pb_variable_array<FieldT> &prev_pc_addr,
                                                 pb_variable_array<FieldT> &prev_pc_val,
                                                 pb_variable_array<FieldT> &prev_state,
                                                 pb_variable_array<FieldT> &ls_addr,
                                                 pb_variable_array<FieldT> &ls_prev_val,
                                                 pb_variable_array<FieldT> &ls_next_val,
                                                 pb_variable_array<FieldT> &next_state,
                                                 pb_variable_array<FieldT> &next_pc_addr,
                                                 pb_variable<FieldT> &next_has_accepted,
                                                 const std::string &annotation_prefix) :
tinyram_standard_gadget<FieldT>(pb, annotation_prefix), prev_pc_addr(prev_pc_addr), prev_pc_val(prev_pc_val),
    prev_state(prev_state), ls_addr(ls_addr), ls_prev_val(ls_prev_val), ls_next_val(ls_next_val),
    next_state(next_state), next_pc_addr(next_pc_addr), next_has_accepted(next_has_accepted)
{
    /* parse previous PC value as an instruction (note that we start
       parsing from LSB of the instruction doubleword and go to the
       MSB) */
    auto pc_val_it = prev_pc_val.begin();

    arg2idx = pb_variable_array<FieldT>(pc_val_it, pc_val_it + pb.ap.reg_arg_or_imm_width()); std::advance(pc_val_it, pb.ap.reg_arg_or_imm_width());
    std::advance(pc_val_it, pb.ap.instruction_padding_width());
    arg1idx = pb_variable_array<FieldT>(pc_val_it, pc_val_it + pb.ap.reg_arg_width()); std::advance(pc_val_it, pb.ap.reg_arg_width());
    desidx = pb_variable_array<FieldT>(pc_val_it, pc_val_it + pb.ap.reg_arg_width()); std::advance(pc_val_it, pb.ap.reg_arg_width());
    arg2_is_imm = *pc_val_it; std::advance(pc_val_it, 1);
    opcode = pb_variable_array<FieldT>(pc_val_it, pc_val_it + pb.ap.opcode_width()); std::advance(pc_val_it, pb.ap.opcode_width());

    assert(pc_val_it == prev_pc_val.end());

    /* parse state as registers + flags */
    pb_variable_array<FieldT> packed_prev_registers, packed_next_registers;
    for (size_t i = 0; i < pb.ap.k; ++i)
    {
        prev_registers.emplace_back(word_variable_gadget<FieldT>(pb, pb_variable_array<FieldT>(prev_state.begin() + i * pb.ap.w, prev_state.begin() + (i + 1) * pb.ap.w), FMT(annotation_prefix, " prev_registers_%zu", i)));
        next_registers.emplace_back(word_variable_gadget<FieldT>(pb, pb_variable_array<FieldT>(next_state.begin() + i * pb.ap.w, next_state.begin() + (i + 1) * pb.ap.w), FMT(annotation_prefix, " next_registers_%zu", i)));

        packed_prev_registers.emplace_back(prev_registers[i].packed);
        packed_next_registers.emplace_back(next_registers[i].packed);
    }
    prev_flag = *(++prev_state.rbegin());
    next_flag = *(++next_state.rbegin());
    prev_tape1_exhausted = *(prev_state.rbegin());
    next_tape1_exhausted = *(next_state.rbegin());

    /* decode arguments */
    prev_pc_addr_as_word_variable.reset(new word_variable_gadget<FieldT>(pb, prev_pc_addr, FMT(annotation_prefix, " prev_pc_addr_as_word_variable")));
    desval.reset(new word_variable_gadget<FieldT>(pb, FMT(annotation_prefix, " desval")));
    arg1val.reset(new word_variable_gadget<FieldT>(pb, FMT(annotation_prefix, " arg1val")));
    arg2val.reset(new word_variable_gadget<FieldT>(pb, FMT(annotation_prefix, " arg2val")));

    decode_arguments.reset(new argument_decoder_gadget<FieldT>(pb, arg2_is_imm, desidx, arg1idx, arg2idx, packed_prev_registers,
                                                               desval->packed, arg1val->packed, arg2val->packed,
                                                               FMT(annotation_prefix, " decode_arguments")));

    /* create indicator variables for opcodes */
    opcode_indicators.allocate(pb, 1ul<<pb.ap.opcode_width(), FMT(annotation_prefix, " opcode_indicators"));

    /* perform the ALU operations */
    instruction_results.allocate(pb, 1ul<<pb.ap.opcode_width(), FMT(annotation_prefix, " instruction_results"));
    instruction_flags.allocate(pb, 1ul<<pb.ap.opcode_width(), FMT(annotation_prefix, " instruction_flags"));

    ALU.reset(new ALU_gadget<FieldT>(pb, opcode_indicators, *prev_pc_addr_as_word_variable, *desval, *arg1val, *arg2val, prev_flag, instruction_results, instruction_flags,
                                     FMT(annotation_prefix, " ALU")));

    /* check correctness of memory operations */
    ls_prev_val_as_doubleword_variable.reset(new doubleword_variable_gadget<FieldT>(pb, ls_prev_val, FMT(annotation_prefix, " ls_prev_val_as_doubleword_variable")))
;
    ls_next_val_as_doubleword_variable.reset(new doubleword_variable_gadget<FieldT>(pb, ls_next_val, FMT(annotation_prefix, " ls_next_val_as_doubleword_variable")));
    memory_subaddress.reset(new dual_variable_gadget<FieldT>(pb, pb_variable_array<FieldT>(arg2val->bits.begin(), arg2val->bits.begin() + pb.ap.subaddr_len()),
                                                             FMT(annotation_prefix, " memory_subaddress")));

    memory_subcontents.allocate(pb, FMT(annotation_prefix, " memory_subcontents"));
    memory_access_is_word.assign(pb, 1 - (opcode_indicators[tinyram_opcode_LOADB] + opcode_indicators[tinyram_opcode_STOREB]));
    memory_access_is_byte.assign(pb, opcode_indicators[tinyram_opcode_LOADB] + opcode_indicators[tinyram_opcode_STOREB]);

    check_memory.reset(new memory_masking_gadget<FieldT>(pb,
                                                         *ls_prev_val_as_doubleword_variable,
                                                         *memory_subaddress,
                                                         memory_subcontents,
                                                         memory_access_is_word,
                                                         memory_access_is_byte,
                                                         *ls_next_val_as_doubleword_variable,
                                                         FMT(annotation_prefix, " check_memory")));

    /* handle reads */
    read_not1.allocate(pb, FMT(annotation_prefix, " read_not1"));

    /* check consistency of the states according to the ALU results */
    next_pc_addr_as_word_variable.reset(new word_variable_gadget<FieldT>(pb, next_pc_addr, FMT(annotation_prefix, " next_pc_addr_as_word_variable")));

    consistency_enforcer.reset(new consistency_enforcer_gadget<FieldT>(pb, opcode_indicators, instruction_results, instruction_flags,
                                                                       desidx, prev_pc_addr_as_word_variable->packed,
                                                                       packed_prev_registers,
                                                                       desval->packed,
                                                                       prev_flag,
                                                                       next_pc_addr_as_word_variable->packed,
                                                                       packed_next_registers,
                                                                       next_flag,
                                                                       FMT(annotation_prefix, " consistency_enforcer")));
}

template<typename FieldT>
void tinyram_cpu_checker<FieldT>::generate_r1cs_constraints()
{
    decode_arguments->generate_r1cs_constraints();

    /* generate indicator variables for opcode */
    for (size_t i = 0; i < 1ul<<this->pb.ap.opcode_width(); ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(opcode_indicators[i], pb_packing_sum<FieldT>(opcode) - i, 0),
                                     FMT(this->annotation_prefix, " opcode_indicators_%zu", i));
    }
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, pb_sum<FieldT>(opcode_indicators), 1),
                                 FMT(this->annotation_prefix, " opcode_indicators_sum_to_1"));

    /* consistency checks for repacked variables */
    for (size_t i = 0; i < this->pb.ap.k; ++i)
    {
        prev_registers[i].generate_r1cs_constraints(true);
        next_registers[i].generate_r1cs_constraints(true);
    }
    prev_pc_addr_as_word_variable->generate_r1cs_constraints(true);
    next_pc_addr_as_word_variable->generate_r1cs_constraints(true);
    ls_prev_val_as_doubleword_variable->generate_r1cs_constraints(true);
    ls_next_val_as_doubleword_variable->generate_r1cs_constraints(true);

    /* main consistency checks */
    decode_arguments->generate_r1cs_constraints();
    ALU->generate_r1cs_constraints();
    consistency_enforcer->generate_r1cs_constraints();

    /* check correct access to memory */
    ls_prev_val_as_doubleword_variable->generate_r1cs_constraints(false);
    ls_next_val_as_doubleword_variable->generate_r1cs_constraints(false);
    memory_subaddress->generate_r1cs_constraints(false);
    check_memory->generate_r1cs_constraints();

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1,
                                                         pb_packing_sum<FieldT>(
                                                             pb_variable_array<FieldT>(arg2val->bits.begin() + this->pb.ap.subaddr_len(),
                                                                                       arg2val->bits.end())),
                                                         pb_packing_sum<FieldT>(ls_addr)),
                                 FMT(this->annotation_prefix, " ls_addr_is_arg2val_minus_subaddress"));

    /* We require that if opcode is one of load.{b,w}, then
       subcontents is appropriately stored in instruction_results. If
       opcode is store.b we only take the necessary portion of arg1val
       (i.e. last byte), and take entire arg1val for store.w.

       Note that ls_addr is *always* going to be arg2val. If the
       instruction is a non-memory instruction, we will treat it as a
       load from that memory location. */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(opcode_indicators[tinyram_opcode_LOADB],
                                                         memory_subcontents - instruction_results[tinyram_opcode_LOADB],
                                                         0),
                                 FMT(this->annotation_prefix, " handle_loadb"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(opcode_indicators[tinyram_opcode_LOADW],
                                                         memory_subcontents - instruction_results[tinyram_opcode_LOADW],
                                                         0),
                                 FMT(this->annotation_prefix, " handle_loadw"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(opcode_indicators[tinyram_opcode_STOREB],
                                                         memory_subcontents - pb_packing_sum<FieldT>(
                                                             pb_variable_array<FieldT>(desval->bits.begin(),
                                                                                       desval->bits.begin() + 8)),
                                                         0),
                                 FMT(this->annotation_prefix, " handle_storeb"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(opcode_indicators[tinyram_opcode_STOREW],
                                                         memory_subcontents - desval->packed,
                                                         0),
                                 FMT(this->annotation_prefix, " handle_storew"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1 - (opcode_indicators[tinyram_opcode_STOREB] + opcode_indicators[tinyram_opcode_STOREW]),
                                                         ls_prev_val_as_doubleword_variable->packed - ls_next_val_as_doubleword_variable->packed,
                                                         0),
                                 FMT(this->annotation_prefix, " non_store_instructions_dont_change_memory"));

    /* specify that accepting state implies opcode = answer && arg2val == 0 */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(next_has_accepted,
                                                         1 - opcode_indicators[tinyram_opcode_ANSWER],
                                                         0),
                                 FMT(this->annotation_prefix, " accepting_requires_answer"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(next_has_accepted,
                                                         arg2val->packed,
                                                         0),
                                 FMT(this->annotation_prefix, " accepting_requires_arg2val_equal_zero"));

    /*
       handle tapes:

       we require that:
       prev_tape1_exhausted implies next_tape1_exhausted,
       prev_tape1_exhausted implies flag to be set
       reads other than from tape 1 imply flag to be set
       flag implies result to be 0
    */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(prev_tape1_exhausted,
                                                         1 - next_tape1_exhausted,
                                                         0),
                                 FMT(this->annotation_prefix, " prev_tape1_exhausted_implies_next_tape1_exhausted"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(prev_tape1_exhausted,
                                                         1 - instruction_flags[tinyram_opcode_READ],
                                                         0),
                                 FMT(this->annotation_prefix, " prev_tape1_exhausted_implies_flag"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(opcode_indicators[tinyram_opcode_READ],
                                                         1 - arg2val->packed,
                                                         read_not1),
                                 FMT(this->annotation_prefix, " read_not1")); /* will be nonzero for read X for X != 1 */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(read_not1,
                                                         1 - instruction_flags[tinyram_opcode_READ],
                                                         0),
                                 FMT(this->annotation_prefix, " other_reads_imply_flag"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(instruction_flags[tinyram_opcode_READ],
                                                         instruction_results[tinyram_opcode_READ],
                                                         0),
                                 FMT(this->annotation_prefix, " read_flag_implies_result_0"));
}

template<typename FieldT>
void tinyram_cpu_checker<FieldT>::generate_r1cs_witness_address()
{
    /* decode instruction and arguments */
    prev_pc_addr_as_word_variable->generate_r1cs_witness_from_bits();
    for (size_t i = 0; i < this->pb.ap.k; ++i)
    {
        prev_registers[i].generate_r1cs_witness_from_bits();
    }

    decode_arguments->generate_r1cs_witness();

    desval->generate_r1cs_witness_from_packed();
    arg1val->generate_r1cs_witness_from_packed();
    arg2val->generate_r1cs_witness_from_packed();

    /* clear out ls_addr and fill with everything of arg2val except the subaddress */
    ls_addr.fill_with_bits_of_field_element(this->pb, this->pb.val(arg2val->packed).as_ulong() >> this->pb.ap.subaddr_len());
}

template<typename FieldT>
void tinyram_cpu_checker<FieldT>::generate_r1cs_witness_other(tinyram_input_tape_iterator &aux_it,
                                                              const tinyram_input_tape_iterator &aux_end)
{
    /* now ls_prev_val is filled with memory contents at ls_addr. we
       now ensure consistency with its doubleword representation */
    ls_prev_val_as_doubleword_variable->generate_r1cs_witness_from_bits();

    /* fill in the opcode indicators */
    const size_t opcode_val = opcode.get_field_element_from_bits(this->pb).as_ulong();
    for (size_t i = 0; i < 1ul<<this->pb.ap.opcode_width(); ++i)
    {
        this->pb.val(opcode_indicators[i]) = (i == opcode_val ? FieldT::one() : FieldT::zero());
    }

    /* execute the ALU */
    ALU->generate_r1cs_witness();

    /* fill memory_subaddress */
    memory_subaddress->bits.fill_with_bits(this->pb, pb_variable_array<FieldT>(arg2val->bits.begin(),
                                                                               arg2val->bits.begin() +  + this->pb.ap.subaddr_len()).get_bits(this->pb));
    memory_subaddress->generate_r1cs_witness_from_bits();

    /* we distinguish four cases for memory handling:
       a) load.b
       b) store.b
       c) store.w
       d) load.w or any non-memory instruction */
    const size_t prev_doubleword = this->pb.val(ls_prev_val_as_doubleword_variable->packed).as_ulong();
    const size_t subaddress = this->pb.val(memory_subaddress->packed).as_ulong();

    if (this->pb.val(opcode_indicators[tinyram_opcode_LOADB]) == FieldT::one())
    {
        const size_t loaded_byte = (prev_doubleword >> (8 * subaddress)) & 0xFF;
        this->pb.val(instruction_results[tinyram_opcode_LOADB]) = FieldT(loaded_byte);
        this->pb.val(memory_subcontents) = FieldT(loaded_byte);
    }
    else if (this->pb.val(opcode_indicators[tinyram_opcode_STOREB]) == FieldT::one())
    {
        const size_t stored_byte = (this->pb.val(desval->packed).as_ulong()) & 0xFF;
        this->pb.val(memory_subcontents) = FieldT(stored_byte);
    }
    else if (this->pb.val(opcode_indicators[tinyram_opcode_STOREW]) == FieldT::one())
    {
        const size_t stored_word = (this->pb.val(desval->packed).as_ulong());
        this->pb.val(memory_subcontents) = FieldT(stored_word);
    }
    else
    {
        const bool access_is_word0 = (this->pb.val(*memory_subaddress->bits.rbegin()) == FieldT::zero());
        const size_t loaded_word = (prev_doubleword >> (access_is_word0 ? 0 : this->pb.ap.w)) & ((1ul << this->pb.ap.w) - 1);
        this->pb.val(instruction_results[tinyram_opcode_LOADW]) = FieldT(loaded_word); /* does not hurt even for non-memory instructions */
        this->pb.val(memory_subcontents) = FieldT(loaded_word);
    }

    memory_access_is_word.evaluate(this->pb);
    memory_access_is_byte.evaluate(this->pb);

    check_memory->generate_r1cs_witness();

    /* handle reads */
    if (this->pb.val(prev_tape1_exhausted) == FieldT::one())
    {
        /* if tape was exhausted before, it will always be
           exhausted. we also need to only handle reads from tape 1,
           so we can safely set flag here */
        this->pb.val(next_tape1_exhausted) = FieldT::one();
        this->pb.val(instruction_flags[tinyram_opcode_READ]) = FieldT::one();
    }

    this->pb.val(read_not1) = this->pb.val(opcode_indicators[tinyram_opcode_READ]) * (FieldT::one() - this->pb.val(arg2val->packed));
    if (this->pb.val(read_not1) != FieldT::one())
    {
        /* reading from tape other than 0 raises the flag */
        this->pb.val(instruction_flags[tinyram_opcode_READ]) = FieldT::one();
    }
    else
    {
        /* otherwise perform the actual read */
        if (aux_it != aux_end)
        {
            this->pb.val(instruction_results[tinyram_opcode_READ]) = FieldT(*aux_it);
            if (++aux_it == aux_end)
            {
                /* tape has ended! */
                this->pb.val(next_tape1_exhausted) = FieldT::one();
            }
        }
        else
        {
            /* handled above, so nothing to do here */
        }
    }

    /* flag implies result zero */
    if (this->pb.val(instruction_flags[tinyram_opcode_READ]) == FieldT::one())
    {
        this->pb.val(instruction_results[tinyram_opcode_READ]) = FieldT::zero();
    }

    /* execute consistency enforcer */
    consistency_enforcer->generate_r1cs_witness();
    next_pc_addr_as_word_variable->generate_r1cs_witness_from_packed();

    for (size_t i = 0; i < this->pb.ap.k; ++i)
    {
        next_registers[i].generate_r1cs_witness_from_packed();
    }

    /* finally set has_accepted to 1 if both the opcode is ANSWER and arg2val is 0 */
    this->pb.val(next_has_accepted) = (this->pb.val(opcode_indicators[tinyram_opcode_ANSWER]) == FieldT::one() &&
                                       this->pb.val(arg2val->packed) == FieldT::zero()) ? FieldT::one() : FieldT::zero();
}

template<typename FieldT>
void tinyram_cpu_checker<FieldT>::dump() const
{
    printf("   pc = %lu, flag = %lu\n",
           this->pb.val(prev_pc_addr_as_word_variable->packed).as_ulong(),
           this->pb.val(prev_flag).as_ulong());
    printf("   ");

    for (size_t j = 0; j < this->pb.ap.k; ++j)
    {
        printf("r%zu = %2lu ", j, this->pb.val(prev_registers[j].packed).as_ulong());
    }
    printf("\n");

    const size_t opcode_val = opcode.get_field_element_from_bits(this->pb).as_ulong();
    printf("   %s r%lu, r%lu, %s%lu\n",
           tinyram_opcode_names[static_cast<tinyram_opcode>(opcode_val)].c_str(),
           desidx.get_field_element_from_bits(this->pb).as_ulong(),
           arg1idx.get_field_element_from_bits(this->pb).as_ulong(),
           (this->pb.val(arg2_is_imm) == FieldT::one() ? "" : "r"),
           arg2idx.get_field_element_from_bits(this->pb).as_ulong());
}

} // libsnark

#endif // TINYRAM_CPU_CHECKER_TCC_

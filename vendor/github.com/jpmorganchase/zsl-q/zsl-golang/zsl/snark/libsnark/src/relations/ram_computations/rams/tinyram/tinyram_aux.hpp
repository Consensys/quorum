/** @file
 *****************************************************************************

 Declaration of auxiliary functions for TinyRAM.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TINYRAM_AUX_HPP_
#define TINYRAM_AUX_HPP_

#include <cassert>
#include <iostream>
#include <map>

#include "common/utils.hpp"
#include "relations/constraint_satisfaction_problems/r1cs/r1cs.hpp"
#include "relations/ram_computations/memory/memory_interface.hpp"
#include "relations/ram_computations/rams/ram_params.hpp"

namespace libsnark {

enum tinyram_opcode {
    tinyram_opcode_AND    = 0b00000,
    tinyram_opcode_OR     = 0b00001,
    tinyram_opcode_XOR    = 0b00010,
    tinyram_opcode_NOT    = 0b00011,
    tinyram_opcode_ADD    = 0b00100,
    tinyram_opcode_SUB    = 0b00101,
    tinyram_opcode_MULL   = 0b00110,
    tinyram_opcode_UMULH  = 0b00111,
    tinyram_opcode_SMULH  = 0b01000,
    tinyram_opcode_UDIV   = 0b01001,
    tinyram_opcode_UMOD   = 0b01010,
    tinyram_opcode_SHL    = 0b01011,
    tinyram_opcode_SHR    = 0b01100,

    tinyram_opcode_CMPE   = 0b01101,
    tinyram_opcode_CMPA   = 0b01110,
    tinyram_opcode_CMPAE  = 0b01111,
    tinyram_opcode_CMPG   = 0b10000,
    tinyram_opcode_CMPGE  = 0b10001,

    tinyram_opcode_MOV    = 0b10010,
    tinyram_opcode_CMOV   = 0b10011,

    tinyram_opcode_JMP    = 0b10100,
    tinyram_opcode_CJMP   = 0b10101,
    tinyram_opcode_CNJMP  = 0b10110,

    tinyram_opcode_10111  = 0b10111,
    tinyram_opcode_11000  = 0b11000,
    tinyram_opcode_11001  = 0b11001,

    tinyram_opcode_STOREB = 0b11010,
    tinyram_opcode_LOADB  = 0b11011,
    tinyram_opcode_STOREW = 0b11100,
    tinyram_opcode_LOADW  = 0b11101,
    tinyram_opcode_READ   = 0b11110,
    tinyram_opcode_ANSWER = 0b11111
};

enum tinyram_opcode_args {
    tinyram_opcode_args_des_arg1_arg2 = 1,
    tinyram_opcode_args_des_arg2 = 2,
    tinyram_opcode_args_arg1_arg2 = 3,
    tinyram_opcode_args_arg2 = 4,
    tinyram_opcode_args_none = 5,
    tinyram_opcode_args_arg2_des = 6
};

/**
 * Instructions that may change a register or the flag.
 * All other instructions leave all registers and the flag intact.
 */
const static int tinyram_opcodes_register[] = {
    tinyram_opcode_AND,
    tinyram_opcode_OR,
    tinyram_opcode_XOR,
    tinyram_opcode_NOT,
    tinyram_opcode_ADD,
    tinyram_opcode_SUB,
    tinyram_opcode_MULL,
    tinyram_opcode_UMULH,
    tinyram_opcode_SMULH,
    tinyram_opcode_UDIV,
    tinyram_opcode_UMOD,
    tinyram_opcode_SHL,
    tinyram_opcode_SHR,

    tinyram_opcode_CMPE,
    tinyram_opcode_CMPA,
    tinyram_opcode_CMPAE,
    tinyram_opcode_CMPG,
    tinyram_opcode_CMPGE,

    tinyram_opcode_MOV,
    tinyram_opcode_CMOV,

    tinyram_opcode_LOADB,
    tinyram_opcode_LOADW,
    tinyram_opcode_READ
};

/**
 * Instructions that modify the program counter.
 * All other instructions either advance it (+1) or stall (see below).
 */
const static int tinyram_opcodes_control_flow[] = {
    tinyram_opcode_JMP,
    tinyram_opcode_CJMP,
    tinyram_opcode_CNJMP
};

/**
 * Instructions that make the program counter stall;
 * these are "answer" plus all the undefined opcodes.
 */
const static int tinyram_opcodes_stall[] = {
    tinyram_opcode_10111,
    tinyram_opcode_11000,
    tinyram_opcode_11001,

    tinyram_opcode_ANSWER
};

typedef size_t reg_count_t; // type for the number of registers
typedef size_t reg_width_t; // type for the width of a register

extern std::map<tinyram_opcode, std::string> tinyram_opcode_names;

extern std::map<std::string, tinyram_opcode> opcode_values;

extern std::map<tinyram_opcode, tinyram_opcode_args> opcode_args;

void ensure_tinyram_opcode_value_map();

class tinyram_program;
typedef std::vector<size_t> tinyram_input_tape;
typedef typename tinyram_input_tape::const_iterator tinyram_input_tape_iterator;

class tinyram_architecture_params {
public:
    reg_width_t w; /* width of a register */
    reg_count_t k; /* number of registers */

    tinyram_architecture_params() {};
    tinyram_architecture_params(const reg_width_t w, const reg_count_t k) : w(w), k(k) { assert(w == 1ul << log2(w)); };

    size_t address_size() const;
    size_t value_size() const;
    size_t cpu_state_size() const;
    size_t initial_pc_addr() const;

    bit_vector initial_cpu_state() const;
    memory_contents initial_memory_contents(const tinyram_program &program,
                                            const tinyram_input_tape &primary_input) const;

    size_t opcode_width() const;
    size_t reg_arg_width() const;
    size_t instruction_padding_width() const;
    size_t reg_arg_or_imm_width() const;

    size_t dwaddr_len() const;
    size_t subaddr_len() const;

    size_t bytes_in_word() const;

    size_t instr_size() const;

    bool operator==(const tinyram_architecture_params &other) const;

    friend std::ostream& operator<<(std::ostream &out, const tinyram_architecture_params &ap);
    friend std::istream& operator>>(std::istream &in, tinyram_architecture_params &ap);

    void print() const;
};

/* order everywhere is reversed (i.e. MSB comes first),
   corresponding to the order in memory */

class tinyram_instruction {
public:
    tinyram_opcode opcode;
    bool arg2_is_imm;
    size_t desidx;
    size_t arg1idx;
    size_t arg2idx_or_imm;

    tinyram_instruction(const tinyram_opcode &opcode,
                        const bool arg2_is_imm,
                        const size_t &desidx,
                        const size_t &arg1idx,
                        const size_t &arg2idx_or_imm);

    size_t as_dword(const tinyram_architecture_params &ap) const;
};

tinyram_instruction random_tinyram_instruction(const tinyram_architecture_params &ap);

std::vector<tinyram_instruction> generate_tinyram_prelude(const tinyram_architecture_params &ap);
extern tinyram_instruction tinyram_default_instruction;

class tinyram_program {
public:
    std::vector<tinyram_instruction> instructions;
    size_t size() const { return instructions.size(); }
    void add_instruction(const tinyram_instruction &instr);
};

tinyram_program load_preprocessed_program(const tinyram_architecture_params &ap,
                                          std::istream &preprocessed);

memory_store_trace tinyram_boot_trace_from_program_and_input(const tinyram_architecture_params &ap,
                                                             const size_t boot_trace_size_bound,
                                                             const tinyram_program &program,
                                                             const tinyram_input_tape &primary_input);

tinyram_input_tape load_tape(std::istream &tape);

} // libsnark

#endif // TINYRAM_AUX_HPP_

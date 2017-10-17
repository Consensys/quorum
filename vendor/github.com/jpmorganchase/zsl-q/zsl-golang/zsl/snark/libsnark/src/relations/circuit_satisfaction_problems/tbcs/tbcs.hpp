/** @file
 *****************************************************************************

 Declaration of interfaces for:
 - a TBCS gate,
 - a TBCS variable assignment, and
 - a TBCS circuit.

 Above, TBCS stands for "Two-input Boolean Circuit Satisfiability".

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TBCS_HPP_
#define TBCS_HPP_

#include "common/profiling.hpp"
#include "relations/variable.hpp"

namespace libsnark {

/*********************** BACS variable assignment ****************************/

/**
 * A TBCS variable assignment is a vector of bools.
 */
typedef std::vector<bool> tbcs_variable_assignment;


/**************************** TBCS gate **************************************/

typedef size_t tbcs_wire_t;

/**
 * Types of TBCS gates (2-input boolean gates).
 *
 * The order and names used below is taken from page 4 of [1].
 *
 * Note that each gate's truth table is encoded in its 4-bit opcode. Namely,
 * if g(X,Y) denotes the output of gate g with inputs X and Y, then
 *            OPCODE(g) = (g(0,0),g(0,1),g(1,0),g(1,1))
 * For example, if g is of type IF_X_THEN_Y, which has opcode 13, then the
 * truth table of g is 1101 (13 in binary).
 *
 * (Note that MSB above is g(0,0) and LSB is g(1,1))
 *
 * References:
 *
 * [1] = https://mitpress.mit.edu/sites/default/files/titles/content/9780262640688_sch_0001.pdf
 */
enum tbcs_gate_type {
    TBCS_GATE_CONSTANT_0 = 0,
    TBCS_GATE_AND = 1,
    TBCS_GATE_X_AND_NOT_Y = 2,
    TBCS_GATE_X = 3,
    TBCS_GATE_NOT_X_AND_Y = 4,
    TBCS_GATE_Y = 5,
    TBCS_GATE_XOR = 6,
    TBCS_GATE_OR = 7,
    TBCS_GATE_NOR = 8,
    TBCS_GATE_EQUIVALENCE = 9,
    TBCS_GATE_NOT_Y = 10,
    TBCS_GATE_IF_Y_THEN_X = 11,
    TBCS_GATE_NOT_X = 12,
    TBCS_GATE_IF_X_THEN_Y = 13,
    TBCS_GATE_NAND = 14,
    TBCS_GATE_CONSTANT_1 = 15
};

static const int num_tbcs_gate_types = 16;

/**
 * A TBCS gate is a formal expression of the form
 *
 *                g(left_wire,right_wire) = output ,
 *
 * where 'left_wire' and 'right_wire' are the two input wires, and 'output' is
 * the output wire. In other words, a TBCS gate is a 2-input boolean gate;
 * there are 16 possible such gates (see tbcs_gate_type above).
 *
 * A TBCS gate is used to construct a TBCS circuit (see below).
 */
class tbcs_gate {
public:

    tbcs_wire_t left_wire;
    tbcs_wire_t right_wire;

    tbcs_gate_type type;

    tbcs_wire_t output;

    bool is_circuit_output;

    bool evaluate(const tbcs_variable_assignment &input) const;
    void print(const std::map<size_t, std::string> &variable_annotations = std::map<size_t, std::string>()) const;
    bool operator==(const tbcs_gate &other) const;

    friend std::ostream& operator<<(std::ostream &out, const tbcs_gate &g);
    friend std::istream& operator>>(std::istream &in, tbcs_gate &g);
};


/****************************** TBCS inputs **********************************/

/**
 * A TBCS primary input is a TBCS variable assignment.
 */
typedef tbcs_variable_assignment tbcs_primary_input;

/**
 * A TBCS auxiliary input is a TBCS variable assigment.
 */
typedef tbcs_variable_assignment tbcs_auxiliary_input;


/************************** TBCS circuit *************************************/

/**
 * A TBCS circuit is a boolean circuit in which every gate has 2 inputs.
 *
 * A TBCS circuit is satisfied by a TBCS variable assignment if every output
 * evaluates to zero.
 *
 * NOTE:
 * The 0-th variable (i.e., "x_{0}") always represents the constant 1.
 * Thus, the 0-th variable is not included in num_variables.
 */
class tbcs_circuit {
public:
    size_t primary_input_size;
    size_t auxiliary_input_size;
    std::vector<tbcs_gate> gates;

    tbcs_circuit() : primary_input_size(0), auxiliary_input_size(0) {}

    size_t num_inputs() const;
    size_t num_gates() const;
    size_t num_wires() const;

    std::vector<size_t> wire_depths() const;
    size_t depth() const;

#ifdef DEBUG
    std::map<size_t, std::string> gate_annotations;
    std::map<size_t, std::string> variable_annotations;
#endif

    bool is_valid() const;
    bool is_satisfied(const tbcs_primary_input &primary_input,
                      const tbcs_auxiliary_input &auxiliary_input) const;

    tbcs_variable_assignment get_all_wires(const tbcs_primary_input &primary_input,
                                           const tbcs_auxiliary_input &auxiliary_input) const;
    tbcs_variable_assignment get_all_outputs(const tbcs_primary_input &primary_input,
                                             const tbcs_auxiliary_input &auxiliary_input) const;

    void add_gate(const tbcs_gate &g);
    void add_gate(const tbcs_gate &g, const std::string &annotation);

    bool operator==(const tbcs_circuit &other) const;

    void print() const;
    void print_info() const;

    friend std::ostream& operator<<(std::ostream &out, const tbcs_circuit &circuit);
    friend std::istream& operator>>(std::istream &in, tbcs_circuit &circuit);
};

} // libsnark

#endif // TBCS_HPP_

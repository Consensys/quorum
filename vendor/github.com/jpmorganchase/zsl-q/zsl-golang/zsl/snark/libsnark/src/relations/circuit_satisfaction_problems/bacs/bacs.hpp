/** @file
 *****************************************************************************

 Declaration of interfaces for:
 - a BACS variable assigment,
 - a BACS gate,
 - a BACS primary input,
 - a BACS auxiliary input,
 - a BACS circuit.

 Above, BACS stands for "Bilinear Arithmetic Circuit Satisfiability".

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_HPP_
#define BACS_HPP_

#include <vector>

#include "relations/variable.hpp"

namespace libsnark {

/*********************** BACS variable assignment ****************************/

/**
 * A BACS variable assignment is a vector of field elements.
 */
template<typename FieldT>
using bacs_variable_assignment = std::vector<FieldT>;


/**************************** BACS gate **************************************/

template<typename FieldT>
struct bacs_gate;

template<typename FieldT>
std::ostream& operator<<(std::ostream &out, const bacs_gate<FieldT> &g);

template<typename FieldT>
std::istream& operator>>(std::istream &in, bacs_gate<FieldT> &g);

/**
 * A BACS gate is a formal expression of the form lhs * rhs = output ,
 * where lhs and rhs are linear combinations (of variables) and output is a variable.
 *
 * In other words, a BACS gate is an arithmetic gate that is bilinear.
 */
template<typename FieldT>
struct bacs_gate {

    linear_combination<FieldT> lhs;
    linear_combination<FieldT> rhs;

    variable<FieldT> output;
    bool is_circuit_output;

    FieldT evaluate(const bacs_variable_assignment<FieldT> &input) const;
    void print(const std::map<size_t, std::string> &variable_annotations = std::map<size_t, std::string>()) const;

    bool operator==(const bacs_gate<FieldT> &other) const;

    friend std::ostream& operator<< <FieldT>(std::ostream &out, const bacs_gate<FieldT> &g);
    friend std::istream& operator>> <FieldT>(std::istream &in, bacs_gate<FieldT> &g);
};


/****************************** BACS inputs **********************************/

/**
 * A BACS primary input is a BACS variable assignment.
 */
template<typename FieldT>
using bacs_primary_input = bacs_variable_assignment<FieldT>;

/**
 * A BACS auxiliary input is a BACS variable assigment.
 */
template<typename FieldT>
using bacs_auxiliary_input = bacs_variable_assignment<FieldT>;


/************************** BACS circuit *************************************/

template<typename FieldT>
class bacs_circuit;

template<typename FieldT>
std::ostream& operator<<(std::ostream &out, const bacs_circuit<FieldT> &circuit);

template<typename FieldT>
std::istream& operator>>(std::istream &in, bacs_circuit<FieldT> &circuit);

/**
 * A BACS circuit is an arithmetic circuit in which every gate is a BACS gate.
 *
 * Given a BACS primary input and a BACS auxiliary input, the circuit can be evaluated.
 * If every output evaluates to zero, then the circuit is satisfied.
 *
 * NOTE:
 * The 0-th variable (i.e., "x_{0}") always represents the constant 1.
 * Thus, the 0-th variable is not included in num_variables.
 */
template<typename FieldT>
class bacs_circuit {
public:
    size_t primary_input_size;
    size_t auxiliary_input_size;
    std::vector<bacs_gate<FieldT> > gates;

    bacs_circuit() : primary_input_size(0), auxiliary_input_size(0) {}

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
    bool is_satisfied(const bacs_primary_input<FieldT> &primary_input,
                      const bacs_auxiliary_input<FieldT> &auxiliary_input) const;

    bacs_variable_assignment<FieldT> get_all_outputs(const bacs_primary_input<FieldT> &primary_input,
                                                     const bacs_auxiliary_input<FieldT> &auxiliary_input) const;
    bacs_variable_assignment<FieldT> get_all_wires(const bacs_primary_input<FieldT> &primary_input,
                                                   const bacs_auxiliary_input<FieldT> &auxiliary_input) const;

    void add_gate(const bacs_gate<FieldT> &g);
    void add_gate(const bacs_gate<FieldT> &g, const std::string &annotation);

    bool operator==(const bacs_circuit<FieldT> &other) const;

    void print() const;
    void print_info() const;

    friend std::ostream& operator<< <FieldT>(std::ostream &out, const bacs_circuit<FieldT> &circuit);
    friend std::istream& operator>> <FieldT>(std::istream &in, bacs_circuit<FieldT> &circuit);
};

} // libsnark

#include "relations/circuit_satisfaction_problems/bacs/bacs.tcc"

#endif // BACS_HPP_

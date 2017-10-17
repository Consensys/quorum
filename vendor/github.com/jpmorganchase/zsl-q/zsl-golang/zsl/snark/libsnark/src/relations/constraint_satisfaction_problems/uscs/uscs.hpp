/** @file
 *****************************************************************************

 Declaration of interfaces for:
 - a USCS constraint,
 - a USCS variable assignment, and
 - a USCS constraint system.

 Above, USCS stands for "Unitary-Square Constraint System".

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef USCS_HPP_
#define USCS_HPP_

#include <cstdlib>
#include <iostream>
#include <map>
#include <string>
#include <vector>

#include "relations/variable.hpp"

namespace libsnark {

/************************* USCS constraint ***********************************/

/**
 * A USCS constraint is a formal expression of the form
 *
 *                \sum_{i=1}^{m} a_i * x_{i} ,
 *
 * where each a_i is in <FieldT> and each x_{i} is a formal variable.
 *
 * A USCS constraint is used to construct a USCS constraint system (see below).
 */
template<typename FieldT>
using uscs_constraint = linear_combination<FieldT>;


/************************* USCS variable assignment **************************/

/**
 * A USCS variable assignment is a vector of <FieldT> elements that represents
 * a candidate solution to a USCS constraint system (see below).
 */
template<typename FieldT>
using uscs_primary_input = std::vector<FieldT>;

template<typename FieldT>
using uscs_auxiliary_input = std::vector<FieldT>;

template<typename FieldT>
using uscs_variable_assignment = std::vector<FieldT>;



/************************* USCS constraint system ****************************/

template<typename FieldT>
class uscs_constraint_system;

template<typename FieldT>
std::ostream& operator<<(std::ostream &out, const uscs_constraint_system<FieldT> &cs);

template<typename FieldT>
std::istream& operator>>(std::istream &in, uscs_constraint_system<FieldT> &cs);

/**
 * A system of USCS constraints looks like
 *
 *     { ( \sum_{i=1}^{m_k} a_{k,i} * x_{k,i} )^2 = 1 }_{k=1}^{n}  .
 *
 * In other words, the system is satisfied if and only if there exist a
 * USCS variable assignment for which each USCS constraint evaluates to -1 or 1.
 *
 * NOTE:
 * The 0-th variable (i.e., "x_{0}") always represents the constant 1.
 * Thus, the 0-th variable is not included in num_variables.
 */
template<typename FieldT>
class uscs_constraint_system {
public:
    size_t primary_input_size;
    size_t auxiliary_input_size;

    std::vector<uscs_constraint<FieldT> > constraints;

    uscs_constraint_system() : primary_input_size(0), auxiliary_input_size(0) {};

    size_t num_inputs() const;
    size_t num_variables() const;
    size_t num_constraints() const;

#ifdef DEBUG
    std::map<size_t, std::string> constraint_annotations;
    std::map<size_t, std::string> variable_annotations;
#endif

    bool is_valid() const;
    bool is_satisfied(const uscs_primary_input<FieldT> &primary_input,
                      const uscs_auxiliary_input<FieldT> &auxiliary_input) const;

    void add_constraint(const uscs_constraint<FieldT> &constraint);
    void add_constraint(const uscs_constraint<FieldT> &constraint, const std::string &annotation);

    bool operator==(const uscs_constraint_system<FieldT> &other) const;

    friend std::ostream& operator<< <FieldT>(std::ostream &out, const uscs_constraint_system<FieldT> &cs);
    friend std::istream& operator>> <FieldT>(std::istream &in, uscs_constraint_system<FieldT> &cs);

    void report_linear_constraint_statistics() const;
};


} // libsnark

#include "relations/constraint_satisfaction_problems/uscs/uscs.tcc"

#endif // USCS_HPP_

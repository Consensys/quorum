/** @file
 *****************************************************************************

 Implementation of functions to sample USCS examples with prescribed parameters
 (according to some distribution).

 See uscs_examples.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef USCS_EXAMPLES_TCC_
#define USCS_EXAMPLES_TCC_

#include <cassert>

#include "common/utils.hpp"

namespace libsnark {

template<typename FieldT>
uscs_example<FieldT> generate_uscs_example_with_field_input(const size_t num_constraints,
                                                            const size_t num_inputs)
{
    enter_block("Call to generate_uscs_example_with_field_input");

    assert(num_inputs >= 1);
    assert(num_constraints >= num_inputs);

    uscs_constraint_system<FieldT> cs;
    cs.primary_input_size = num_inputs;
    cs.auxiliary_input_size = num_constraints - num_inputs;

    uscs_variable_assignment<FieldT> full_variable_assignment;
    for (size_t i = 0; i < num_constraints; ++i)
    {
        full_variable_assignment.emplace_back(FieldT(std::rand()));
    }

    for (size_t i = 0; i < num_constraints; ++i)
    {
        size_t x, y, z;

        do
        {
            x = std::rand() % num_constraints;
            y = std::rand() % num_constraints;
            z = std::rand() % num_constraints;
        } while (x == z || y == z);

        const FieldT x_coeff = FieldT(std::rand());
        const FieldT y_coeff = FieldT(std::rand());
        const FieldT val = (std::rand() % 2 == 0 ? FieldT::one() : -FieldT::one());
        const FieldT z_coeff = (val - x_coeff * full_variable_assignment[x] - y_coeff * full_variable_assignment[y]) * full_variable_assignment[z].inverse();

        uscs_constraint<FieldT> constr;
        constr.add_term(x+1, x_coeff);
        constr.add_term(y+1, y_coeff);
        constr.add_term(z+1, z_coeff);

        cs.add_constraint(constr);
    }

    /* split variable assignment */
    uscs_primary_input<FieldT> primary_input(full_variable_assignment.begin(), full_variable_assignment.begin() + num_inputs);
    uscs_primary_input<FieldT> auxiliary_input(full_variable_assignment.begin() + num_inputs, full_variable_assignment.end());

    /* sanity checks */
    assert(cs.num_variables() == full_variable_assignment.size());
    assert(cs.num_variables() >= num_inputs);
    assert(cs.num_inputs() == num_inputs);
    assert(cs.num_constraints() == num_constraints);
    assert(cs.is_satisfied(primary_input, auxiliary_input));

    leave_block("Call to generate_uscs_example_with_field_input");

    return uscs_example<FieldT>(std::move(cs), std::move(primary_input), std::move(auxiliary_input));
}

template<typename FieldT>
uscs_example<FieldT> generate_uscs_example_with_binary_input(const size_t num_constraints,
                                                             const size_t num_inputs)
{
    enter_block("Call to generate_uscs_example_with_binary_input");

    assert(num_inputs >= 1);

    uscs_constraint_system<FieldT> cs;
    cs.primary_input_size = num_inputs;
    cs.auxiliary_input_size = num_constraints;

    uscs_variable_assignment<FieldT> full_variable_assignment;
    for (size_t i = 0; i < num_inputs; ++i)
    {
        full_variable_assignment.push_back(FieldT(std::rand() % 2));
    }

    size_t lastvar = num_inputs-1;
    for (size_t i = 0; i < num_constraints; ++i)
    {
        ++lastvar;

        /* chose two random bits and XOR them together */
        const size_t u = (i == 0 ? std::rand() % num_inputs : std::rand() % i);
        const size_t v = (i == 0 ? std::rand() % num_inputs : std::rand() % i);

        uscs_constraint<FieldT> constr;
        constr.add_term(u+1, 1);
        constr.add_term(v+1, 1);
        constr.add_term(lastvar+1, 1);
        constr.add_term(0,-FieldT::one()); // shift constant term (which is 0) by 1

        cs.add_constraint(constr);
        full_variable_assignment.push_back(full_variable_assignment[u] + full_variable_assignment[v] - full_variable_assignment[u] * full_variable_assignment[v] - full_variable_assignment[u] * full_variable_assignment[v]);
    }

    /* split variable assignment */
    uscs_primary_input<FieldT> primary_input(full_variable_assignment.begin(), full_variable_assignment.begin() + num_inputs);
    uscs_primary_input<FieldT> auxiliary_input(full_variable_assignment.begin() + num_inputs, full_variable_assignment.end());

    /* sanity checks */
    assert(cs.num_variables() == full_variable_assignment.size());
    assert(cs.num_variables() >= num_inputs);
    assert(cs.num_inputs() == num_inputs);
    assert(cs.num_constraints() == num_constraints);
    assert(cs.is_satisfied(primary_input, auxiliary_input));

    leave_block("Call to generate_uscs_example_with_binary_input");

    return uscs_example<FieldT>(std::move(cs), std::move(primary_input), std::move(auxiliary_input));
}

} // libsnark
#endif // USCS_EXAMPLES_TCC

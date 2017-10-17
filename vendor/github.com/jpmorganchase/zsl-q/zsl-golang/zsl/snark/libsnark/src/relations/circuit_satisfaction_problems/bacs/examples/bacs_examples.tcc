/** @file
 *****************************************************************************

 Implementation of functions to sample BACS examples with prescribed parameters
 (according to some distribution).

 See bacs_examples.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_EXAMPLES_TCC_
#define BACS_EXAMPLES_TCC_

#include <cassert>

#include "common/utils.hpp"

namespace libsnark {

template<typename FieldT>
linear_combination<FieldT> random_linear_combination(const size_t num_variables)
{
    const size_t terms = 1 + (std::rand() % 3);
    linear_combination<FieldT> result;

    for (size_t i = 0; i < terms; ++i)
    {
        const FieldT coeff = FieldT(std::rand()); // TODO: replace with FieldT::random_element(), when it becomes faster...
        result = result + coeff * variable<FieldT>(std::rand() % (num_variables + 1));
    }

    return result;
}

template<typename FieldT>
bacs_example<FieldT> generate_bacs_example(const size_t primary_input_size,
                                           const size_t auxiliary_input_size,
                                           const size_t num_gates,
                                           const size_t num_outputs)
{
    bacs_example<FieldT> example;
    for (size_t i = 0; i < primary_input_size; ++i)
    {
        example.primary_input.emplace_back(FieldT::random_element());
    }

    for (size_t i = 0; i < auxiliary_input_size; ++i)
    {
        example.auxiliary_input.emplace_back(FieldT::random_element());
    }

    example.circuit.primary_input_size = primary_input_size;
    example.circuit.auxiliary_input_size = auxiliary_input_size;

    bacs_variable_assignment<FieldT> all_vals;
    all_vals.insert(all_vals.end(), example.primary_input.begin(), example.primary_input.end());
    all_vals.insert(all_vals.end(), example.auxiliary_input.begin(), example.auxiliary_input.end());

    for (size_t i = 0; i < num_gates; ++i)
    {
        const size_t num_variables = primary_input_size + auxiliary_input_size + i;
        bacs_gate<FieldT> gate;
        gate.lhs = random_linear_combination<FieldT>(num_variables);
        gate.rhs = random_linear_combination<FieldT>(num_variables);
        gate.output = variable<FieldT>(num_variables+1);

        if (i >= num_gates - num_outputs)
        {
            /* make gate a circuit output and fix */
            gate.is_circuit_output = true;
            const var_index_t var_idx = std::rand() % (1 + primary_input_size + std::min(num_gates-num_outputs, i));
            const FieldT var_val = (var_idx == 0 ? FieldT::one() : all_vals[var_idx-1]);

            if (std::rand() % 2 == 0)
            {
                const FieldT lhs_val = gate.lhs.evaluate(all_vals);
                const FieldT coeff = -(lhs_val * var_val.inverse());
                gate.lhs = gate.lhs + coeff * variable<FieldT>(var_idx);
            }
            else
            {
                const FieldT rhs_val = gate.rhs.evaluate(all_vals);
                const FieldT coeff = -(rhs_val * var_val.inverse());
                gate.rhs = gate.rhs + coeff * variable<FieldT>(var_idx);
            }

            assert(gate.evaluate(all_vals).is_zero());
        }
        else
        {
            gate.is_circuit_output = false;
        }

        example.circuit.add_gate(gate);
        all_vals.emplace_back(gate.evaluate(all_vals));
    }

    assert(example.circuit.is_satisfied(example.primary_input, example.auxiliary_input));

    return example;
}

} // libsnark

#endif // BACS_EXAMPLES_TCC

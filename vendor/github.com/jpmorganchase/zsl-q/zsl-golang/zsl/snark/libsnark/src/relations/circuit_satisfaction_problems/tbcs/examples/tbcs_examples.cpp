/** @file
 *****************************************************************************

 Implementation of functions to sample TBCS examples with prescribed parameters
 (according to some distribution).

 See tbcs_examples.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "relations/circuit_satisfaction_problems/tbcs/examples/tbcs_examples.hpp"

#include <cassert>

#include "common/utils.hpp"

namespace libsnark {

tbcs_example generate_tbcs_example(const size_t primary_input_size,
                                   const size_t auxiliary_input_size,
                                   const size_t num_gates,
                                   const size_t num_outputs)
{
    tbcs_example example;
    for (size_t i = 0; i < primary_input_size; ++i)
    {
        example.primary_input.push_back(std::rand() % 2 == 0 ? false : true);
    }

    for (size_t i = 0; i < auxiliary_input_size; ++i)
    {
        example.auxiliary_input.push_back(std::rand() % 2 == 0 ? false : true);
    }

    example.circuit.primary_input_size = primary_input_size;
    example.circuit.auxiliary_input_size = auxiliary_input_size;

    tbcs_variable_assignment all_vals;
    all_vals.insert(all_vals.end(), example.primary_input.begin(), example.primary_input.end());
    all_vals.insert(all_vals.end(), example.auxiliary_input.begin(), example.auxiliary_input.end());

    for (size_t i = 0; i < num_gates; ++i)
    {
        const size_t num_variables = primary_input_size + auxiliary_input_size + i;
        tbcs_gate gate;
        gate.left_wire = std::rand() % (num_variables+1);
        gate.right_wire = std::rand() % (num_variables+1);
        gate.output = num_variables+1;

        if (i >= num_gates - num_outputs)
        {
            /* make gate a circuit output and fix */
            do
            {
                gate.type = (tbcs_gate_type)(std::rand() % num_tbcs_gate_types);
            }
            while (gate.evaluate(all_vals));

            gate.is_circuit_output = true;
        }
        else
        {
            gate.type = (tbcs_gate_type)(std::rand() % num_tbcs_gate_types);
            gate.is_circuit_output = false;
        }

        example.circuit.add_gate(gate);
        all_vals.push_back(gate.evaluate(all_vals));
    }

    assert(example.circuit.is_satisfied(example.primary_input, example.auxiliary_input));

    return example;
}

} // libsnark

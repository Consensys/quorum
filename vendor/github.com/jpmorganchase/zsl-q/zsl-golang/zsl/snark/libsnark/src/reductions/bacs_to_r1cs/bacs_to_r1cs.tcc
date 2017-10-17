/** @file
*****************************************************************************

Implementation of interfaces for a BACS-to-R1CS reduction.

See bacs_to_r1cs.hpp .

*****************************************************************************
* @author     This file is part of libsnark, developed by SCIPR Lab
*             and contributors (see AUTHORS).
* @copyright  MIT license (see LICENSE file)
*****************************************************************************/

#ifndef BACS_TO_R1CS_TCC_
#define BACS_TO_R1CS_TCC_

#include "relations/circuit_satisfaction_problems/bacs/bacs.hpp"
#include "relations/constraint_satisfaction_problems/r1cs/r1cs.hpp"

namespace libsnark {

template<typename FieldT>
r1cs_constraint_system<FieldT> bacs_to_r1cs_instance_map(const bacs_circuit<FieldT> &circuit)
{
    enter_block("Call to bacs_to_r1cs_instance_map");
    assert(circuit.is_valid());
    r1cs_constraint_system<FieldT> result;

#ifdef DEBUG
    result.variable_annotations = circuit.variable_annotations;
#endif

    result.primary_input_size = circuit.primary_input_size;
    result.auxiliary_input_size = circuit.auxiliary_input_size + circuit.gates.size();

    for (auto &g : circuit.gates)
    {
        result.constraints.emplace_back(r1cs_constraint<FieldT>(g.lhs, g.rhs, g.output));
#ifdef DEBUG
        auto it = circuit.gate_annotations.find(g.output.index);
        if (it != circuit.gate_annotations.end())
        {
            result.constraint_annotations[result.constraints.size()-1] = it->second;
        }
#endif
    }

    for (auto &g : circuit.gates)
    {
        if (g.is_circuit_output)
        {
            result.constraints.emplace_back(r1cs_constraint<FieldT>(1, g.output, 0));

#ifdef DEBUG
            result.constraint_annotations[result.constraints.size()-1] = FMT("", "output_%zu_is_circuit_output", g.output.index);
#endif
        }
    }

    leave_block("Call to bacs_to_r1cs_instance_map");

    return result;
}

template<typename FieldT>
r1cs_variable_assignment<FieldT> bacs_to_r1cs_witness_map(const bacs_circuit<FieldT> &circuit,
                                                               const bacs_primary_input<FieldT> &primary_input,
                                                               const bacs_auxiliary_input<FieldT> &auxiliary_input)
{
    enter_block("Call to bacs_to_r1cs_witness_map");
    const r1cs_variable_assignment<FieldT> result = circuit.get_all_wires(primary_input, auxiliary_input);
    leave_block("Call to bacs_to_r1cs_witness_map");

    return result;
}

} // libsnark

#endif // BACS_TO_R1CS_TCC_

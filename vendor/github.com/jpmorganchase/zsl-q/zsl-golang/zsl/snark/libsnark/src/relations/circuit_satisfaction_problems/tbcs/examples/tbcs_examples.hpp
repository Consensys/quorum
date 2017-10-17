/** @file
 *****************************************************************************

 Declaration of interfaces for a TBCS example, as well as functions to sample
 TBCS examples with prescribed parameters (according to some distribution).

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TBCS_EXAMPLES_HPP_
#define TBCS_EXAMPLES_HPP_

#include "relations/circuit_satisfaction_problems/tbcs/tbcs.hpp"

namespace libsnark {

/**
 * A TBCS example comprises a TBCS circuit, TBCS primary input, and TBCS auxiliary input.
 */
struct tbcs_example {

    tbcs_circuit circuit;
    tbcs_primary_input primary_input;
    tbcs_auxiliary_input auxiliary_input;

    tbcs_example() = default;
    tbcs_example(const tbcs_example &other) = default;
    tbcs_example(const tbcs_circuit &circuit,
                 const tbcs_primary_input &primary_input,
                 const tbcs_auxiliary_input &auxiliary_input) :
        circuit(circuit),
        primary_input(primary_input),
        auxiliary_input(auxiliary_input)
    {}

    tbcs_example(tbcs_circuit &&circuit,
                 tbcs_primary_input &&primary_input,
                 tbcs_auxiliary_input &&auxiliary_input) :
        circuit(std::move(circuit)),
        primary_input(std::move(primary_input)),
        auxiliary_input(std::move(auxiliary_input))
    {}
};

/**
 * Generate a TBCS example such that:
 * - the primary input has size primary_input_size;
 * - the auxiliary input has size auxiliary_input_size;
 * - the circuit has num_gates gates;
 * - the circuit has num_outputs (<= num_gates) output gates.
 *
 * This is done by first selecting primary and auxiliary inputs uniformly at random, and then for each gate:
 * - selecting random left and right wires from primary inputs, auxiliary inputs, and outputs of previous gates,
 * - selecting a gate type at random (subject to the constraint "output = 0" if this is an output gate).
 */
tbcs_example generate_tbcs_example(const size_t primary_input_size,
                                   const size_t auxiliary_input_size,
                                   const size_t num_gates,
                                   const size_t num_outputs);

} // libsnark

#endif // TBCS_EXAMPLES_HPP_

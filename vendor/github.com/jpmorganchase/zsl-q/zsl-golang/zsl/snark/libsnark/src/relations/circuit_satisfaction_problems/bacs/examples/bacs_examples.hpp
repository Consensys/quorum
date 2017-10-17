/** @file
 *****************************************************************************

 Declaration of interfaces for a BACS example, as well as functions to sample
 BACS examples with prescribed parameters (according to some distribution).

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_EXAMPLES_HPP_
#define BACS_EXAMPLES_HPP_

#include "relations/circuit_satisfaction_problems/bacs/bacs.hpp"

namespace libsnark {

/**
 * A BACS example comprises a BACS circuit, BACS primary input, and BACS auxiliary input.
 */
template<typename FieldT>
struct bacs_example {

    bacs_circuit<FieldT> circuit;
    bacs_primary_input<FieldT> primary_input;
    bacs_auxiliary_input<FieldT> auxiliary_input;

    bacs_example<FieldT>() = default;
    bacs_example<FieldT>(const bacs_example<FieldT> &other) = default;
    bacs_example<FieldT>(const bacs_circuit<FieldT> &circuit,
                         const bacs_primary_input<FieldT> &primary_input,
                         const bacs_auxiliary_input<FieldT> &auxiliary_input) :
        circuit(circuit),
        primary_input(primary_input),
        auxiliary_input(auxiliary_input)
    {}

    bacs_example<FieldT>(bacs_circuit<FieldT> &&circuit,
                         bacs_primary_input<FieldT> &&primary_input,
                         bacs_auxiliary_input<FieldT> &&auxiliary_input) :
        circuit(std::move(circuit)),
        primary_input(std::move(primary_input)),
        auxiliary_input(std::move(auxiliary_input))
    {}
};

/**
 * Generate a BACS example such that:
 * - the primary input has size primary_input_size;
 * - the auxiliary input has size auxiliary_input_size;
 * - the circuit has num_gates gates;
 * - the circuit has num_outputs (<= num_gates) output gates.
 *
 * This is done by first selecting primary and auxiliary inputs uniformly at random, and then for each gate:
 * - selecting random left and right wires from primary inputs, auxiliary inputs, and outputs of previous gates,
 * - selecting random linear combinations for left and right wires, consisting of 1, 2, 3 or 4 terms each, with random coefficients,
 * - if the gate is an output gate, then adding a random non-output wire to either left or right linear combination, with appropriate coefficient, so that the linear combination evaluates to 0.
 */
template<typename FieldT>
bacs_example<FieldT> generate_bacs_example(const size_t primary_input_size,
                                           const size_t auxiliary_input_size,
                                           const size_t num_gates,
                                           const size_t num_outputs);

} // libsnark

#include "relations/circuit_satisfaction_problems/bacs/examples/bacs_examples.tcc"

#endif // BACS_EXAMPLES_HPP_

/** @file
 *****************************************************************************

 Declaration of interfaces for a BACS-to-R1CS reduction, that is, constructing
 a R1CS ("Rank-1 Constraint System") from a BACS ("Bilinear Arithmetic Circuit Satisfiability").

 The reduction is straightforward: each bilinear gate gives rises to a
 corresponding R1CS constraint that enforces correct computation of the gate;
 also, each output gives rise to a corresponding R1CS constraint that enforces
 that the output is zero.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_TO_R1CS_HPP_
#define BACS_TO_R1CS_HPP_

#include "relations/circuit_satisfaction_problems/bacs/bacs.hpp"
#include "relations/constraint_satisfaction_problems/r1cs/r1cs.hpp"

namespace libsnark {

/**
 * Instance map for the BACS-to-R1CS reduction.
 */
template<typename FieldT>
r1cs_constraint_system<FieldT> bacs_to_r1cs_instance_map(const bacs_circuit<FieldT> &circuit);

/**
 * Witness map for the BACS-to-R1CS reduction.
 */
template<typename FieldT>
r1cs_variable_assignment<FieldT> bacs_to_r1cs_witness_map(const bacs_circuit<FieldT> &circuit,
                                                               const bacs_primary_input<FieldT> &primary_input,
                                                               const bacs_auxiliary_input<FieldT> &auxiliary_input);

} // libsnark

#include "reductions/bacs_to_r1cs/bacs_to_r1cs.tcc"

#endif // BACS_TO_R1CS_HPP_

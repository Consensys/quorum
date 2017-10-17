/** @file
 *****************************************************************************

 Declaration of interfaces for a TBCS-to-USCS reduction, that is, constructing
 a USCS ("Unitary-Square Constraint System") from a TBCS ("Two-input Boolean Circuit Satisfiability").

 The reduction is straightforward: each non-output wire is mapped to a
 corresponding USCS constraint that enforces the wire to carry a boolean value;
 each 2-input boolean gate is mapped to a corresponding USCS constraint that
 enforces correct computation of the gate; each output wire is mapped to a
 corresponding USCS constraint that enforces that the output is zero.

 The mapping of a gate to a USCS constraint is due to \[GOS12].

 References:

 \[GOS12]:
 "New techniques for noninteractive zero-knowledge",
 Jens Groth, Rafail Ostrovsky, Amit Sahai
 JACM 2012,
 <http://www0.cs.ucl.ac.uk/staff/J.Groth/NIZKJournal.pdf>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TBCS_TO_USCS_HPP_
#define TBCS_TO_USCS_HPP_

#include "relations/constraint_satisfaction_problems/uscs/uscs.hpp"
#include "relations/circuit_satisfaction_problems/tbcs/tbcs.hpp"

namespace libsnark {

/**
 * Instance map for the TBCS-to-USCS reduction.
 */
template<typename FieldT>
uscs_constraint_system<FieldT> tbcs_to_uscs_instance_map(const tbcs_circuit &circuit);

/**
 * Witness map for the TBCS-to-USCS reduction.
 */
template<typename FieldT>
uscs_variable_assignment<FieldT> tbcs_to_uscs_witness_map(const tbcs_circuit &circuit,
                                                               const tbcs_primary_input &primary_input,
                                                               const tbcs_auxiliary_input &auxiliary_input);

} // libsnark

#include "reductions/tbcs_to_uscs/tbcs_to_uscs.tcc"

#endif // TBCS_TO_USCS_HPP_

/** @file
 *****************************************************************************

 Declaration of functionality that runs the USCS ppzkSNARK for
 a given USCS example.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_USCS_PPZKSNARK_HPP_
#define RUN_USCS_PPZKSNARK_HPP_

#include "relations/constraint_satisfaction_problems/uscs/examples/uscs_examples.hpp"

namespace libsnark {

/**
 * Runs the ppzkSNARK (generator, prover, and verifier) for a given
 * USCS example (specified by a constraint system, input, and witness).
 *
 * Optionally, also test the serialization routines for keys and proofs.
 * (This takes additional time.)
 */
template<typename ppT>
bool run_uscs_ppzksnark(const uscs_example<Fr<ppT> > &example,
                        const bool test_serialization);

} // libsnark

#include "zk_proof_systems/ppzksnark/uscs_ppzksnark/examples/run_uscs_ppzksnark.tcc"

#endif // RUN_USCS_PPZKSNARK_HPP_

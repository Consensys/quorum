/** @file
 *****************************************************************************

 Declaration of functionality that runs the TBCS ppzkSNARK for
 a given TBCS example.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_TBCS_PPZKSNARK_HPP_
#define RUN_TBCS_PPZKSNARK_HPP_

#include "relations/circuit_satisfaction_problems/tbcs/examples/tbcs_examples.hpp"

namespace libsnark {

/**
 * Runs the ppzkSNARK (generator, prover, and verifier) for a given
 * TBCS example (specified by a circuit, primary input, and auxiliary input).
 *
 * Optionally, also test the serialization routines for keys and proofs.
 * (This takes additional time.)
 */
template<typename ppT>
bool run_tbcs_ppzksnark(const tbcs_example &example,
                        const bool test_serialization);

} // libsnark

#include "zk_proof_systems/ppzksnark/tbcs_ppzksnark/examples/run_tbcs_ppzksnark.tcc"

#endif // RUN_TBCS_PPZKSNARK_HPP_

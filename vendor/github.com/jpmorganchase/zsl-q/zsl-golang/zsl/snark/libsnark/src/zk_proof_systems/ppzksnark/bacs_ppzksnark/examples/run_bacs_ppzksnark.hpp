/** @file
 *****************************************************************************

 Declaration of functionality that runs the BACS ppzkSNARK for
 a given BACS example.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_BACS_PPZKSNARK_HPP_
#define RUN_BACS_PPZKSNARK_HPP_

#include "relations/circuit_satisfaction_problems/bacs/examples/bacs_examples.hpp"

namespace libsnark {

/**
 * Runs the ppzkSNARK (generator, prover, and verifier) for a given
 * BACS example (specified by a circuit, primary input, and auxiliary input).
 *
 * Optionally, also test the serialization routines for keys and proofs.
 * (This takes additional time.)
 */
template<typename ppT>
bool run_bacs_ppzksnark(const bacs_example<Fr<ppT> > &example,
                        const bool test_serialization);

} // libsnark

#include "zk_proof_systems/ppzksnark/bacs_ppzksnark/examples/run_bacs_ppzksnark.tcc"

#endif // RUN_BACS_PPZKSNARK_HPP_

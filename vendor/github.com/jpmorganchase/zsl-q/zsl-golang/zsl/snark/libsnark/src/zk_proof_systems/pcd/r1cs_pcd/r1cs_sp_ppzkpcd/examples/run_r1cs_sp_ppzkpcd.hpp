/** @file
 *****************************************************************************

 Declaration of functionality that runs the R1CS single-predicate ppzkPCD
 for a compliance predicate example.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_R1CS_SP_PPZKPCD_HPP_
#define RUN_R1CS_SP_PPZKPCD_HPP_

namespace libsnark {

/**
 * Runs the single-predicate ppzkPCD (generator, prover, and verifier) for the
 * "tally compliance predicate", of a given wordsize, arity, and depth.
 *
 * Optionally, also test the serialization routines for keys and proofs.
 * (This takes additional time.)
 */
template<typename PCD_ppT>
bool run_r1cs_sp_ppzkpcd_tally_example(const size_t wordsize,
                                       const size_t arity,
                                       const size_t depth,
                                       const bool test_serialization);

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/examples/run_r1cs_sp_ppzkpcd.tcc"

#endif // RUN_R1CS_SP_PPZKPCD_HPP_

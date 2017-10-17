/** @file
 *****************************************************************************

 Declaration of functionality that runs the R1CS multi-predicate ppzkPCD
 for a compliance predicate example.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_R1CS_MP_PPZKPCD_HPP_
#define RUN_R1CS_MP_PPZKPCD_HPP_

namespace libsnark {

/**
 * Runs the multi-predicate ppzkPCD (generator, prover, and verifier) for the
 * "tally compliance predicate", of a given wordsize, arity, and depth.
 *
 * Optionally, also test the serialization routines for keys and proofs.
 * (This takes additional time.)
 *
 * Optionally, also test the case of compliance predicates with different types.
 */
template<typename PCD_ppT>
bool run_r1cs_mp_ppzkpcd_tally_example(const size_t wordsize,
                                       const size_t max_arity,
                                       const size_t depth,
                                       const bool test_serialization,
                                       const bool test_multi_type,
                                       const bool test_same_type_optimization);

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_mp_ppzkpcd/examples/run_r1cs_mp_ppzkpcd.tcc"

#endif // RUN_R1CS_MP_PPZKPCD_HPP_

/** @file
 *****************************************************************************

 Declaration of functionality that runs the RAM zkSNARK for
 a given RAM example.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_RAM_ZKSNARK_HPP_
#define RUN_RAM_ZKSNARK_HPP_

#include "relations/ram_computations/rams/examples/ram_examples.hpp"
#include "zk_proof_systems/zksnark/ram_zksnark/ram_zksnark_params.hpp"

namespace libsnark {

/**
 * Runs the zkSNARK (generator, prover, and verifier) for a given
 * RAM example (specified by an architecture, boot trace, auxiliary input, and time bound).
 *
 * Optionally, also test the serialization routines for keys and proofs.
 * (This takes additional time.)
 */
template<typename ram_zksnark_ppT>
bool run_ram_zksnark(const ram_example<ram_zksnark_machine_pp<ram_zksnark_ppT> > &example,
                     const bool test_serialization);

} // libsnark

#include "zk_proof_systems/zksnark/ram_zksnark/examples/run_ram_zksnark.tcc"

#endif // RUN_RAM_ZKSNARK_HPP_

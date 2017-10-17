/** @file
 *****************************************************************************

 Declaration of functionality that runs the RAM ppzkSNARK for
 a given RAM example.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_RAM_PPZKSNARK_HPP_
#define RUN_RAM_PPZKSNARK_HPP_

#include "relations/ram_computations/rams/examples/ram_examples.hpp"
#include "zk_proof_systems/ppzksnark/ram_ppzksnark/ram_ppzksnark_params.hpp"

namespace libsnark {

/**
 * Runs the ppzkSNARK (generator, prover, and verifier) for a given
 * RAM example (specified by an architecture, boot trace, auxiliary input, and time bound).
 *
 * Optionally, also test the serialization routines for keys and proofs.
 * (This takes additional time.)
 */
template<typename ram_ppzksnark_ppT>
bool run_ram_ppzksnark(const ram_example<ram_ppzksnark_machine_pp<ram_ppzksnark_ppT> > &example,
                       const bool test_serialization);

} // libsnark

#include "zk_proof_systems/ppzksnark/ram_ppzksnark/examples/run_ram_ppzksnark.tcc"

#endif // RUN_RAM_PPZKSNARK_HPP_

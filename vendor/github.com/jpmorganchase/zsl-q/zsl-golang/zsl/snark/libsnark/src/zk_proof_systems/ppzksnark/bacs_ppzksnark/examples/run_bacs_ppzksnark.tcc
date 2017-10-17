/** @file
 *****************************************************************************

 Implementation of functionality that runs the BACS ppzkSNARK for
 a given BACS example.

 See run_bacs_ppzksnark.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_BACS_PPZKSNARK_TCC_
#define RUN_BACS_PPZKSNARK_TCC_

#include "zk_proof_systems/ppzksnark/bacs_ppzksnark/bacs_ppzksnark.hpp"

#include <sstream>

#include "common/profiling.hpp"

namespace libsnark {

/**
 * The code below provides an example of all stages of running a BACS ppzkSNARK.
 *
 * Of course, in a real-life scenario, we would have three distinct entities,
 * mangled into one in the demonstration below. The three entities are as follows.
 * (1) The "generator", which runs the ppzkSNARK generator on input a given
 *     circuit C to create a proving and a verification key for C.
 * (2) The "prover", which runs the ppzkSNARK prover on input the proving key,
 *     a primary input for C, and an auxiliary input for C.
 * (3) The "verifier", which runs the ppzkSNARK verifier on input the verification key,
 *     a primary input for C, and a proof.
 */
template<typename ppT>
bool run_bacs_ppzksnark(const bacs_example<Fr<ppT> > &example,
                        const bool test_serialization)
{
    enter_block("Call to run_bacs_ppzksnark");

    print_header("BACS ppzkSNARK Generator");
    bacs_ppzksnark_keypair<ppT> keypair = bacs_ppzksnark_generator<ppT>(example.circuit);
    printf("\n"); print_indent(); print_mem("after generator");

    print_header("Preprocess verification key");
    bacs_ppzksnark_processed_verification_key<ppT> pvk = bacs_ppzksnark_verifier_process_vk<ppT>(keypair.vk);

    if (test_serialization)
    {
        enter_block("Test serialization of keys");
        keypair.pk = reserialize<bacs_ppzksnark_proving_key<ppT> >(keypair.pk);
        keypair.vk = reserialize<bacs_ppzksnark_verification_key<ppT> >(keypair.vk);
        pvk = reserialize<bacs_ppzksnark_processed_verification_key<ppT> >(pvk);
        leave_block("Test serialization of keys");
    }

    print_header("BACS ppzkSNARK Prover");
    bacs_ppzksnark_proof<ppT> proof = bacs_ppzksnark_prover<ppT>(keypair.pk, example.primary_input, example.auxiliary_input);
    printf("\n"); print_indent(); print_mem("after prover");

    if (test_serialization)
    {
        enter_block("Test serialization of proof");
        proof = reserialize<bacs_ppzksnark_proof<ppT> >(proof);
        leave_block("Test serialization of proof");
    }

    print_header("BACS ppzkSNARK Verifier");
    bool ans = bacs_ppzksnark_verifier_strong_IC<ppT>(keypair.vk, example.primary_input, proof);
    printf("\n"); print_indent(); print_mem("after verifier");
    printf("* The verification result is: %s\n", (ans ? "PASS" : "FAIL"));

    print_header("BACS ppzkSNARK Online Verifier");
    bool ans2 = bacs_ppzksnark_online_verifier_strong_IC<ppT>(pvk, example.primary_input, proof);
    assert(ans == ans2);

    leave_block("Call to run_bacs_ppzksnark");

    return ans;
}

} // libsnark

#endif // RUN_BACS_PPZKSNARK_TCC_

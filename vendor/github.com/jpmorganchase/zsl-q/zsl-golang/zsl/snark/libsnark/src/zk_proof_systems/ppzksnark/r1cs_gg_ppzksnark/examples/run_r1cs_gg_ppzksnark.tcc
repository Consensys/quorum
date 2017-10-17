/** @file
 *****************************************************************************

 Implementation of functionality that runs the R1CS GG-ppzkSNARK for
 a given R1CS example.

 See run_r1cs_gg_ppzksnark.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_R1CS_GG_PPZKSNARK_TCC_
#define RUN_R1CS_GG_PPZKSNARK_TCC_

#include "zk_proof_systems/ppzksnark/r1cs_gg_ppzksnark/r1cs_gg_ppzksnark.hpp"

#include <sstream>
#include <type_traits>

#include "common/profiling.hpp"

namespace libsnark {

template<typename ppT>
typename std::enable_if<ppT::has_affine_pairing, void>::type
test_affine_verifier(const r1cs_gg_ppzksnark_verification_key<ppT> &vk,
                     const r1cs_gg_ppzksnark_primary_input<ppT> &primary_input,
                     const r1cs_gg_ppzksnark_proof<ppT> &proof,
                     const bool expected_answer)
{
    print_header("R1CS GG-ppzkSNARK Affine Verifier");
    const bool answer = r1cs_gg_ppzksnark_affine_verifier_weak_IC<ppT>(vk, primary_input, proof);
    assert(answer == expected_answer);
}

template<typename ppT>
typename std::enable_if<!ppT::has_affine_pairing, void>::type
test_affine_verifier(const r1cs_gg_ppzksnark_verification_key<ppT> &vk,
                     const r1cs_gg_ppzksnark_primary_input<ppT> &primary_input,
                     const r1cs_gg_ppzksnark_proof<ppT> &proof,
                     const bool expected_answer)
{
    print_header("R1CS GG-ppzkSNARK Affine Verifier");
    UNUSED(vk, primary_input, proof, expected_answer);
    printf("Affine verifier is not supported; not testing anything.\n");
}

/**
 * The code below provides an example of all stages of running a R1CS GG-ppzkSNARK.
 *
 * Of course, in a real-life scenario, we would have three distinct entities,
 * mangled into one in the demonstration below. The three entities are as follows.
 * (1) The "generator", which runs the ppzkSNARK generator on input a given
 *     constraint system CS to create a proving and a verification key for CS.
 * (2) The "prover", which runs the ppzkSNARK prover on input the proving key,
 *     a primary input for CS, and an auxiliary input for CS.
 * (3) The "verifier", which runs the ppzkSNARK verifier on input the verification key,
 *     a primary input for CS, and a proof.
 */
template<typename ppT>
bool run_r1cs_gg_ppzksnark(const r1cs_example<Fr<ppT> > &example,
                        const bool test_serialization)
{
    enter_block("Call to run_r1cs_gg_ppzksnark");

    print_header("R1CS GG-ppzkSNARK Generator");
    r1cs_gg_ppzksnark_keypair<ppT> keypair = r1cs_gg_ppzksnark_generator<ppT>(example.constraint_system);
    printf("\n"); print_indent(); print_mem("after generator");

    print_header("Preprocess verification key");
    r1cs_gg_ppzksnark_processed_verification_key<ppT> pvk = r1cs_gg_ppzksnark_verifier_process_vk<ppT>(keypair.vk);

    if (test_serialization)
    {
        enter_block("Test serialization of keys");
        keypair.pk = reserialize<r1cs_gg_ppzksnark_proving_key<ppT> >(keypair.pk);
        keypair.vk = reserialize<r1cs_gg_ppzksnark_verification_key<ppT> >(keypair.vk);
        pvk = reserialize<r1cs_gg_ppzksnark_processed_verification_key<ppT> >(pvk);
        leave_block("Test serialization of keys");
    }

    print_header("R1CS GG-ppzkSNARK Prover");
    r1cs_gg_ppzksnark_proof<ppT> proof = r1cs_gg_ppzksnark_prover<ppT>(keypair.pk, example.primary_input, example.auxiliary_input);
    printf("\n"); print_indent(); print_mem("after prover");

    if (test_serialization)
    {
        enter_block("Test serialization of proof");
        proof = reserialize<r1cs_gg_ppzksnark_proof<ppT> >(proof);
        leave_block("Test serialization of proof");
    }

    print_header("R1CS GG-ppzkSNARK Verifier");
    const bool ans = r1cs_gg_ppzksnark_verifier_strong_IC<ppT>(keypair.vk, example.primary_input, proof);
    printf("\n"); print_indent(); print_mem("after verifier");
    printf("* The verification result is: %s\n", (ans ? "PASS" : "FAIL"));

    print_header("R1CS GG-ppzkSNARK Online Verifier");
    const bool ans2 = r1cs_gg_ppzksnark_online_verifier_strong_IC<ppT>(pvk, example.primary_input, proof);
    assert(ans == ans2);

    test_affine_verifier<ppT>(keypair.vk, example.primary_input, proof, ans);

    leave_block("Call to run_r1cs_gg_ppzksnark");

    return ans;
}

} // libsnark

#endif // RUN_R1CS_GG_PPZKSNARK_TCC_

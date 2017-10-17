/** @file
 *****************************************************************************

 Implementation of functionality that runs the R1CS ppzkADSNARK for
 a given R1CS example.

 See run_r1cs_ppzkadsnark.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_R1CS_PPZKADSNARK_TCC_
#define RUN_R1CS_PPZKADSNARK_TCC_

#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark.hpp"
#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/examples/prf/aes_ctr_prf.tcc"
#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/examples/signature/ed25519_signature.tcc"

#include <sstream>
#include <type_traits>

#include "common/profiling.hpp"

namespace libsnark {

/**
 * The code below provides an example of all stages of running a R1CS ppzkADSNARK.
 *
 * Of course, in a real-life scenario, we would have three distinct entities,
 * mangled into one in the demonstration below. The three entities are as follows.
 * (1) The "generator", which runs the ppzkADSNARK generator on input a given
 *     constraint system CS to create a proving and a verification key for CS.
 * (2) The "prover", which runs the ppzkADSNARK prover on input the proving key,
 *     a primary input for CS, and an auxiliary input for CS.
 * (3) The "verifier", which runs the ppzkADSNARK verifier on input the verification key,
 *     a primary input for CS, and a proof.
 */
template<typename ppT>
bool run_r1cs_ppzkadsnark(const r1cs_example<Fr<snark_pp<ppT>> > &example,
                          const bool test_serialization)
{
    enter_block("Call to run_r1cs_ppzkadsnark");

    r1cs_ppzkadsnark_auth_keys<ppT> auth_keys = r1cs_ppzkadsnark_auth_generator<ppT>();

    print_header("R1CS ppzkADSNARK Generator");
    r1cs_ppzkadsnark_keypair<ppT> keypair = r1cs_ppzkadsnark_generator<ppT>(example.constraint_system,auth_keys.pap);
    printf("\n"); print_indent(); print_mem("after generator");

    print_header("Preprocess verification key");
    r1cs_ppzkadsnark_processed_verification_key<ppT> pvk = r1cs_ppzkadsnark_verifier_process_vk<ppT>(keypair.vk);

    if (test_serialization)
    {
        enter_block("Test serialization of keys");
        keypair.pk = reserialize<r1cs_ppzkadsnark_proving_key<ppT> >(keypair.pk);
        keypair.vk = reserialize<r1cs_ppzkadsnark_verification_key<ppT> >(keypair.vk);
        pvk = reserialize<r1cs_ppzkadsnark_processed_verification_key<ppT> >(pvk);
        leave_block("Test serialization of keys");
    }

    print_header("R1CS ppzkADSNARK Authenticate");
    std::vector<Fr<snark_pp<ppT>>> data;
    data.reserve(example.constraint_system.num_inputs());
    std::vector<labelT> labels;
    labels.reserve(example.constraint_system.num_inputs());
    for (size_t i = 0; i < example.constraint_system.num_inputs(); i++) {
        labels.emplace_back(labelT());
        data.emplace_back(example.primary_input[i]);
    }
    std::vector<r1cs_ppzkadsnark_auth_data<ppT>> auth_data =
        r1cs_ppzkadsnark_auth_sign<ppT>(data,auth_keys.sak,labels);

    print_header("R1CS ppzkADSNARK Verify Symmetric");
    bool auth_res =
        r1cs_ppzkadsnark_auth_verify<ppT>(data,auth_data,auth_keys.sak,labels);
    printf("* The verification result is: %s\n", (auth_res ? "PASS" : "FAIL"));

    print_header("R1CS ppzkADSNARK Verify Public");
    bool auth_resp =
        r1cs_ppzkadsnark_auth_verify<ppT>(data,auth_data,auth_keys.pak,labels);
    assert (auth_res == auth_resp);

    print_header("R1CS ppzkADSNARK Prover");
    r1cs_ppzkadsnark_proof<ppT> proof = r1cs_ppzkadsnark_prover<ppT>(keypair.pk, example.primary_input, example.auxiliary_input,auth_data);
    printf("\n"); print_indent(); print_mem("after prover");

    if (test_serialization)
    {
        enter_block("Test serialization of proof");
        proof = reserialize<r1cs_ppzkadsnark_proof<ppT> >(proof);
        leave_block("Test serialization of proof");
    }

    print_header("R1CS ppzkADSNARK Symmetric Verifier");
    bool ans = r1cs_ppzkadsnark_verifier<ppT>(keypair.vk, proof,auth_keys.sak,labels);
    printf("\n"); print_indent(); print_mem("after verifier");
    printf("* The verification result is: %s\n", (ans ? "PASS" : "FAIL"));

    print_header("R1CS ppzkADSNARK Symmetric Online Verifier");
    bool ans2 = r1cs_ppzkadsnark_online_verifier<ppT>(pvk, proof,auth_keys.sak,labels);
    assert(ans == ans2);

    print_header("R1CS ppzkADSNARK Public Verifier");
    ans = r1cs_ppzkadsnark_verifier<ppT>(keypair.vk, auth_data, proof,auth_keys.pak,labels);
    printf("\n"); print_indent(); print_mem("after verifier");
    printf("* The verification result is: %s\n", (ans ? "PASS" : "FAIL"));

    print_header("R1CS ppzkADSNARK Public Online Verifier");
    ans2 = r1cs_ppzkadsnark_online_verifier<ppT>(pvk, auth_data, proof,auth_keys.pak,labels);
    assert(ans == ans2);

    leave_block("Call to run_r1cs_ppzkadsnark");

    return ans;
}

} // libsnark

#endif // RUN_R1CS_PPZKADSNARK_TCC_

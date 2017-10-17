/** @file
 *****************************************************************************

 Declaration of interfaces for the the R1CS ppzkSNARK verifier gadget.

 The gadget r1cs_ppzksnark_verifier_gadget verifiers correct computation of r1cs_ppzksnark_verifier_strong_IC.
 The gadget is built from two main sub-gadgets:
 - r1cs_ppzksnark_verifier_process_vk_gadget, which verifies correct computation of r1cs_ppzksnark_verifier_process_vk, and
 - r1cs_ppzksnark_online_verifier_gadget, which verifies correct computation of r1cs_ppzksnark_online_verifier_strong_IC.
 See r1cs_ppzksnark.hpp for description of the aforementioned functions.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_PPZKSNARK_VERIFIER_GADGET_HPP_
#define R1CS_PPZKSNARK_VERIFIER_GADGET_HPP_

#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "gadgetlib1/gadgets/curves/weierstrass_g1_gadget.hpp"
#include "gadgetlib1/gadgets/curves/weierstrass_g2_gadget.hpp"
#include "gadgetlib1/gadgets/pairing/pairing_checks.hpp"
#include "gadgetlib1/gadgets/pairing/pairing_params.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp"

namespace libsnark {

template<typename ppT>
class r1cs_ppzksnark_proof_variable : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    std::shared_ptr<G1_variable<ppT> > g_A_g;
    std::shared_ptr<G1_variable<ppT> > g_A_h;
    std::shared_ptr<G2_variable<ppT> > g_B_g;
    std::shared_ptr<G1_variable<ppT> > g_B_h;
    std::shared_ptr<G1_variable<ppT> > g_C_g;
    std::shared_ptr<G1_variable<ppT> > g_C_h;
    std::shared_ptr<G1_variable<ppT> > g_H;
    std::shared_ptr<G1_variable<ppT> > g_K;

    std::vector<std::shared_ptr<G1_variable<ppT> > > all_G1_vars;
    std::vector<std::shared_ptr<G2_variable<ppT> > > all_G2_vars;

    std::vector<std::shared_ptr<G1_checker_gadget<ppT> > > all_G1_checkers;
    std::shared_ptr<G2_checker_gadget<ppT> > G2_checker;

    pb_variable_array<FieldT> proof_contents;

    r1cs_ppzksnark_proof_variable(protoboard<FieldT> &pb,
                                  const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness(const r1cs_ppzksnark_proof<other_curve<ppT> > &proof);
    static size_t size();
};

template<typename ppT>
class r1cs_ppzksnark_verification_key_variable : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    std::shared_ptr<G2_variable<ppT> > alphaA_g2;
    std::shared_ptr<G1_variable<ppT> > alphaB_g1;
    std::shared_ptr<G2_variable<ppT> > alphaC_g2;
    std::shared_ptr<G2_variable<ppT> > gamma_g2;
    std::shared_ptr<G1_variable<ppT> > gamma_beta_g1;
    std::shared_ptr<G2_variable<ppT> > gamma_beta_g2;
    std::shared_ptr<G2_variable<ppT> > rC_Z_g2;
    std::shared_ptr<G1_variable<ppT> > encoded_IC_base;
    std::vector<std::shared_ptr<G1_variable<ppT> > > encoded_IC_query;

    pb_variable_array<FieldT> all_bits;
    pb_linear_combination_array<FieldT> all_vars;
    size_t input_size;

    std::vector<std::shared_ptr<G1_variable<ppT> > > all_G1_vars;
    std::vector<std::shared_ptr<G2_variable<ppT> > > all_G2_vars;

    std::shared_ptr<multipacking_gadget<FieldT> > packer;

    // Unfortunately, g++ 4.9 and g++ 5.0 have a bug related to
    // incorrect inlining of small functions:
    // https://gcc.gnu.org/bugzilla/show_bug.cgi?id=65307, which
    // produces wrong assembly even at -O1. The test case at the bug
    // report is directly derived from this code here. As a temporary
    // work-around we mark the key functions noinline to hint compiler
    // that inlining should not be performed.

    // TODO: remove later, when g++ developers fix the bug.

    __attribute__((noinline)) r1cs_ppzksnark_verification_key_variable(protoboard<FieldT> &pb,
                                                                       const pb_variable_array<FieldT> &all_bits,
                                                                       const size_t input_size,
                                                                       const std::string &annotation_prefix);
    void generate_r1cs_constraints(const bool enforce_bitness);
    void generate_r1cs_witness(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &vk);
    void generate_r1cs_witness(const bit_vector &vk_bits);
    bit_vector get_bits() const;
    static size_t __attribute__((noinline)) size_in_bits(const size_t input_size);
    static bit_vector get_verification_key_bits(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &r1cs_vk);
};

template<typename ppT>
class r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable {
public:
    typedef Fr<ppT> FieldT;

    std::shared_ptr<G1_variable<ppT> > encoded_IC_base;
    std::vector<std::shared_ptr<G1_variable<ppT> > > encoded_IC_query;

    std::shared_ptr<G1_precomputation<ppT> > vk_alphaB_g1_precomp;
    std::shared_ptr<G1_precomputation<ppT> > vk_gamma_beta_g1_precomp;

    std::shared_ptr<G2_precomputation<ppT> > pp_G2_one_precomp;
    std::shared_ptr<G2_precomputation<ppT> > vk_alphaA_g2_precomp;
    std::shared_ptr<G2_precomputation<ppT> > vk_alphaC_g2_precomp;
    std::shared_ptr<G2_precomputation<ppT> > vk_gamma_beta_g2_precomp;
    std::shared_ptr<G2_precomputation<ppT> > vk_gamma_g2_precomp;
    std::shared_ptr<G2_precomputation<ppT> > vk_rC_Z_g2_precomp;

    r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable();
    r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable(protoboard<FieldT> &pb,
                                                                         const r1cs_ppzksnark_verification_key<other_curve<ppT> > &r1cs_vk,
                                                                         const std::string &annotation_prefix);
};

template<typename ppT>
class r1cs_ppzksnark_verifier_process_vk_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    std::shared_ptr<precompute_G1_gadget<ppT> > compute_vk_alphaB_g1_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_vk_gamma_beta_g1_precomp;

    std::shared_ptr<precompute_G2_gadget<ppT> > compute_vk_alphaA_g2_precomp;
    std::shared_ptr<precompute_G2_gadget<ppT> > compute_vk_alphaC_g2_precomp;
    std::shared_ptr<precompute_G2_gadget<ppT> > compute_vk_gamma_beta_g2_precomp;
    std::shared_ptr<precompute_G2_gadget<ppT> > compute_vk_gamma_g2_precomp;
    std::shared_ptr<precompute_G2_gadget<ppT> > compute_vk_rC_Z_g2_precomp;

    r1cs_ppzksnark_verification_key_variable<ppT> vk;
    r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> &pvk; // important to have a reference here

    r1cs_ppzksnark_verifier_process_vk_gadget(protoboard<FieldT> &pb,
                                              const r1cs_ppzksnark_verification_key_variable<ppT> &vk,
                                              r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> &pvk,
                                              const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename ppT>
class r1cs_ppzksnark_online_verifier_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> pvk;

    pb_variable_array<FieldT> input;
    size_t elt_size;
    r1cs_ppzksnark_proof_variable<ppT> proof;
    pb_variable<FieldT> result;
    const size_t input_len;

    std::shared_ptr<G1_variable<ppT> > acc;
    std::shared_ptr<G1_multiscalar_mul_gadget<ppT> > accumulate_input;

    std::shared_ptr<G1_variable<ppT> > proof_g_A_g_acc;
    std::shared_ptr<G1_add_gadget<ppT> > compute_proof_g_A_g_acc;
    std::shared_ptr<G1_variable<ppT> > proof_g_A_g_acc_C;
    std::shared_ptr<G1_add_gadget<ppT> > compute_proof_g_A_g_acc_C;

    std::shared_ptr<G1_precomputation<ppT> > proof_g_A_h_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_A_g_acc_C_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_A_g_acc_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_A_g_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_B_h_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_C_h_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_C_g_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_K_precomp;
    std::shared_ptr<G1_precomputation<ppT> > proof_g_H_precomp;

    std::shared_ptr<G2_precomputation<ppT> > proof_g_B_g_precomp;

    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_A_h_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_A_g_acc_C_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_A_g_acc_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_A_g_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_B_h_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_C_h_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_C_g_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_K_precomp;
    std::shared_ptr<precompute_G1_gadget<ppT> > compute_proof_g_H_precomp;

    std::shared_ptr<precompute_G2_gadget<ppT> > compute_proof_g_B_g_precomp;

    std::shared_ptr<check_e_equals_e_gadget<ppT> > check_kc_A_valid;
    std::shared_ptr<check_e_equals_e_gadget<ppT> > check_kc_B_valid;
    std::shared_ptr<check_e_equals_e_gadget<ppT> > check_kc_C_valid;
    std::shared_ptr<check_e_equals_ee_gadget<ppT> > check_QAP_valid;
    std::shared_ptr<check_e_equals_ee_gadget<ppT> > check_CC_valid;

    pb_variable<FieldT> kc_A_valid;
    pb_variable<FieldT> kc_B_valid;
    pb_variable<FieldT> kc_C_valid;
    pb_variable<FieldT> QAP_valid;
    pb_variable<FieldT> CC_valid;

    pb_variable_array<FieldT> all_test_results;
    std::shared_ptr<conjunction_gadget<FieldT> > all_tests_pass;

    r1cs_ppzksnark_online_verifier_gadget(protoboard<FieldT> &pb,
                                          const r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> &pvk,
                                          const pb_variable_array<FieldT> &input,
                                          const size_t elt_size,
                                          const r1cs_ppzksnark_proof_variable<ppT> &proof,
                                          const pb_variable<FieldT> &result,
                                          const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename ppT>
class r1cs_ppzksnark_verifier_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    std::shared_ptr<r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> > pvk;
    std::shared_ptr<r1cs_ppzksnark_verifier_process_vk_gadget<ppT> > compute_pvk;
    std::shared_ptr<r1cs_ppzksnark_online_verifier_gadget<ppT> > online_verifier;

    r1cs_ppzksnark_verifier_gadget(protoboard<FieldT> &pb,
                                   const r1cs_ppzksnark_verification_key_variable<ppT> &vk,
                                   const pb_variable_array<FieldT> &input,
                                   const size_t elt_size,
                                   const r1cs_ppzksnark_proof_variable<ppT> &proof,
                                   const pb_variable<FieldT> &result,
                                   const std::string &annotation_prefix);

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/verifiers/r1cs_ppzksnark_verifier_gadget.tcc"

#endif // R1CS_PPZKSNARK_VERIFIER_GADGET_HPP_

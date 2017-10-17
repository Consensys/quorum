/** @file
 *****************************************************************************

 Implementation of interfaces for the the R1CS ppzkSNARK verifier gadget.

 See r1cs_ppzksnark_verifier_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_PPZKSNARK_VERIFIER_GADGET_TCC_
#define R1CS_PPZKSNARK_VERIFIER_GADGET_TCC_

#include "gadgetlib1/constraint_profiling.hpp"

namespace libsnark {

template<typename ppT>
r1cs_ppzksnark_proof_variable<ppT>::r1cs_ppzksnark_proof_variable(protoboard<FieldT> &pb,
                                                                  const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix)
{
    const size_t num_G1 = 7;
    const size_t num_G2 = 1;

    g_A_g.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " g_A_g")));
    g_A_h.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " g_A_h")));
    g_B_g.reset(new G2_variable<ppT>(pb, FMT(annotation_prefix, " g_B_g")));
    g_B_h.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " g_B_h")));
    g_C_g.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " g_C_g")));
    g_C_h.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " g_C_h")));
    g_H.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " g_H")));
    g_K.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " g_K")));

    all_G1_vars = { g_A_g, g_A_h, g_B_h, g_C_g, g_C_h, g_H,g_K };
    all_G2_vars = { g_B_g };

    all_G1_checkers.resize(all_G1_vars.size());

    for (size_t i = 0; i < all_G1_vars.size(); ++i)
    {
        all_G1_checkers[i].reset(new G1_checker_gadget<ppT>(pb, *all_G1_vars[i], FMT(annotation_prefix, " all_G1_checkers_%zu", i)));
    }
    G2_checker.reset(new G2_checker_gadget<ppT>(pb, *g_B_g, FMT(annotation_prefix, " G2_checker")));

    assert(all_G1_vars.size() == num_G1);
    assert(all_G2_vars.size() == num_G2);
}

template<typename ppT>
void r1cs_ppzksnark_proof_variable<ppT>::generate_r1cs_constraints()
{
    for (auto &G1_checker : all_G1_checkers)
    {
        G1_checker->generate_r1cs_constraints();
    }

    G2_checker->generate_r1cs_constraints();
}

template<typename ppT>
void r1cs_ppzksnark_proof_variable<ppT>::generate_r1cs_witness(const r1cs_ppzksnark_proof<other_curve<ppT> > &proof)
{
    std::vector<G1<other_curve<ppT> > > G1_elems;
    std::vector<G2<other_curve<ppT> > > G2_elems;

    G1_elems = { proof.g_A.g, proof.g_A.h, proof.g_B.h, proof.g_C.g, proof.g_C.h, proof.g_H, proof.g_K };
    G2_elems = { proof.g_B.g };

    assert(G1_elems.size() == all_G1_vars.size());
    assert(G2_elems.size() == all_G2_vars.size());

    for (size_t i = 0; i < G1_elems.size(); ++i)
    {
        all_G1_vars[i]->generate_r1cs_witness(G1_elems[i]);
    }

    for (size_t i = 0; i < G2_elems.size(); ++i)
    {
        all_G2_vars[i]->generate_r1cs_witness(G2_elems[i]);
    }

    for (auto &G1_checker : all_G1_checkers)
    {
        G1_checker->generate_r1cs_witness();
    }

    G2_checker->generate_r1cs_witness();
}

template<typename ppT>
size_t r1cs_ppzksnark_proof_variable<ppT>::size()
{
    const size_t num_G1 = 7;
    const size_t num_G2 = 1;
    return (num_G1 * G1_variable<ppT>::num_field_elems + num_G2 * G2_variable<ppT>::num_field_elems);
}

template<typename ppT>
r1cs_ppzksnark_verification_key_variable<ppT>::r1cs_ppzksnark_verification_key_variable(protoboard<FieldT> &pb,
                                                                                        const pb_variable_array<FieldT> &all_bits,
                                                                                        const size_t input_size,
                                                                                        const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    all_bits(all_bits),
    input_size(input_size)
{
    const size_t num_G1 = 2 + (input_size + 1);
    const size_t num_G2 = 5;

    assert(all_bits.size() == (G1_variable<ppT>::size_in_bits() * num_G1 + G2_variable<ppT>::size_in_bits() * num_G2));

    this->alphaA_g2.reset(new G2_variable<ppT>(pb, FMT(annotation_prefix, " alphaA_g2")));
    this->alphaB_g1.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " alphaB_g1")));
    this->alphaC_g2.reset(new G2_variable<ppT>(pb, FMT(annotation_prefix, " alphaC_g2")));
    this->gamma_g2.reset(new G2_variable<ppT>(pb, FMT(annotation_prefix, " gamma_g2")));
    this->gamma_beta_g1.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " gamma_beta_g1")));
    this->gamma_beta_g2.reset(new G2_variable<ppT>(pb, FMT(annotation_prefix, " gamma_beta_g2")));
    this->rC_Z_g2.reset(new G2_variable<ppT>(pb, FMT(annotation_prefix, " rC_Z_g2")));

    all_G1_vars = { this->alphaB_g1, this->gamma_beta_g1 };
    all_G2_vars = { this->alphaA_g2, this->alphaC_g2, this->gamma_g2, this->gamma_beta_g2, this->rC_Z_g2 };

    this->encoded_IC_query.resize(input_size);
    this->encoded_IC_base.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " encoded_IC_base")));
    this->all_G1_vars.emplace_back(this->encoded_IC_base);

    for (size_t i = 0; i < input_size; ++i)
    {
        this->encoded_IC_query[i].reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " encoded_IC_query_%zu", i)));
        all_G1_vars.emplace_back(this->encoded_IC_query[i]);
    }

    for (auto &G1_var : all_G1_vars)
    {
        all_vars.insert(all_vars.end(), G1_var->all_vars.begin(), G1_var->all_vars.end());
    }

    for (auto &G2_var : all_G2_vars)
    {
        all_vars.insert(all_vars.end(), G2_var->all_vars.begin(), G2_var->all_vars.end());
    }

    assert(all_G1_vars.size() == num_G1);
    assert(all_G2_vars.size() == num_G2);
    assert(all_vars.size() == (num_G1 * G1_variable<ppT>::num_variables() + num_G2 * G2_variable<ppT>::num_variables()));

    packer.reset(new multipacking_gadget<FieldT>(pb, all_bits, all_vars, FieldT::size_in_bits(), FMT(annotation_prefix, " packer")));
}

template<typename ppT>
void r1cs_ppzksnark_verification_key_variable<ppT>::generate_r1cs_constraints(const bool enforce_bitness)
{
    packer->generate_r1cs_constraints(enforce_bitness);
}

template<typename ppT>
void r1cs_ppzksnark_verification_key_variable<ppT>::generate_r1cs_witness(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &vk)
{
    std::vector<G1<other_curve<ppT> > > G1_elems;
    std::vector<G2<other_curve<ppT> > > G2_elems;

    G1_elems = { vk.alphaB_g1, vk.gamma_beta_g1 };
    G2_elems = { vk.alphaA_g2, vk.alphaC_g2, vk.gamma_g2, vk.gamma_beta_g2, vk.rC_Z_g2 };

    assert(vk.encoded_IC_query.rest.indices.size() == input_size);
    G1_elems.emplace_back(vk.encoded_IC_query.first);
    for (size_t i = 0; i < input_size; ++i)
    {
        assert(vk.encoded_IC_query.rest.indices[i] == i);
        G1_elems.emplace_back(vk.encoded_IC_query.rest.values[i]);
    }

    assert(G1_elems.size() == all_G1_vars.size());
    assert(G2_elems.size() == all_G2_vars.size());

    for (size_t i = 0; i < G1_elems.size(); ++i)
    {
        all_G1_vars[i]->generate_r1cs_witness(G1_elems[i]);
    }

    for (size_t i = 0; i < G2_elems.size(); ++i)
    {
        all_G2_vars[i]->generate_r1cs_witness(G2_elems[i]);
    }

    packer->generate_r1cs_witness_from_packed();
}

template<typename ppT>
void r1cs_ppzksnark_verification_key_variable<ppT>::generate_r1cs_witness(const bit_vector &vk_bits)
{
    all_bits.fill_with_bits(this->pb, vk_bits);
    packer->generate_r1cs_witness_from_bits();
}

template<typename ppT>
bit_vector r1cs_ppzksnark_verification_key_variable<ppT>::get_bits() const
{
    return all_bits.get_bits(this->pb);
}

template<typename ppT>
size_t r1cs_ppzksnark_verification_key_variable<ppT>::size_in_bits(const size_t input_size)
{
    const size_t num_G1 = 2 + (input_size + 1);
    const size_t num_G2 = 5;
    const size_t result = G1_variable<ppT>::size_in_bits() * num_G1 + G2_variable<ppT>::size_in_bits() * num_G2;
    printf("G1_size_in_bits = %zu, G2_size_in_bits = %zu\n", G1_variable<ppT>::size_in_bits(), G2_variable<ppT>::size_in_bits());
    printf("r1cs_ppzksnark_verification_key_variable<ppT>::size_in_bits(%zu) = %zu\n", input_size, result);
    return result;
}

template<typename ppT>
bit_vector r1cs_ppzksnark_verification_key_variable<ppT>::get_verification_key_bits(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &r1cs_vk)
{
    typedef Fr<ppT> FieldT;

    const size_t input_size_in_elts = r1cs_vk.encoded_IC_query.rest.indices.size(); // this might be approximate for bound verification keys, however they are not supported by r1cs_ppzksnark_verification_key_variable
    const size_t vk_size_in_bits = r1cs_ppzksnark_verification_key_variable<ppT>::size_in_bits(input_size_in_elts);

    protoboard<FieldT> pb;
    pb_variable_array<FieldT> vk_bits;
    vk_bits.allocate(pb, vk_size_in_bits, "vk_bits");
    r1cs_ppzksnark_verification_key_variable<ppT> vk(pb, vk_bits, input_size_in_elts, "translation_step_vk");
    vk.generate_r1cs_witness(r1cs_vk);

    return vk.get_bits();
}

template<typename ppT>
r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT>::r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable()
{
    // will be allocated outside
}

template<typename ppT>
r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT>::r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable(protoboard<FieldT> &pb,
                                                                                                                                                const r1cs_ppzksnark_verification_key<other_curve<ppT> > &r1cs_vk,
                                                                                                                                                const std::string &annotation_prefix)
{
    encoded_IC_base.reset(new G1_variable<ppT>(pb, r1cs_vk.encoded_IC_query.first, FMT(annotation_prefix, " encoded_IC_base")));
    encoded_IC_query.resize(r1cs_vk.encoded_IC_query.rest.indices.size());
    for (size_t i = 0; i < r1cs_vk.encoded_IC_query.rest.indices.size(); ++i)
    {
        assert(r1cs_vk.encoded_IC_query.rest.indices[i] == i);
        encoded_IC_query[i].reset(new G1_variable<ppT>(pb, r1cs_vk.encoded_IC_query.rest.values[i], FMT(annotation_prefix, " encoded_IC_query")));
    }

    vk_alphaB_g1_precomp.reset(new G1_precomputation<ppT>(pb, r1cs_vk.alphaB_g1, FMT(annotation_prefix, " vk_alphaB_g1_precomp")));
    vk_gamma_beta_g1_precomp.reset(new G1_precomputation<ppT>(pb, r1cs_vk.gamma_beta_g1, FMT(annotation_prefix, " vk_gamma_beta_g1_precomp")));

    pp_G2_one_precomp.reset(new G2_precomputation<ppT>(pb, G2<other_curve<ppT> >::one(), FMT(annotation_prefix, " pp_G2_one_precomp")));
    vk_alphaA_g2_precomp.reset(new G2_precomputation<ppT>(pb, r1cs_vk.alphaA_g2, FMT(annotation_prefix, " vk_alphaA_g2_precomp")));
    vk_alphaC_g2_precomp.reset(new G2_precomputation<ppT>(pb, r1cs_vk.alphaC_g2, FMT(annotation_prefix, " vk_alphaC_g2_precomp")));
    vk_gamma_beta_g2_precomp.reset(new G2_precomputation<ppT>(pb, r1cs_vk.gamma_beta_g2, FMT(annotation_prefix, " vk_gamma_beta_g2_precomp")));
    vk_gamma_g2_precomp.reset(new G2_precomputation<ppT>(pb, r1cs_vk.gamma_g2, FMT(annotation_prefix, " vk_gamma_g2_precomp")));
    vk_rC_Z_g2_precomp.reset(new G2_precomputation<ppT>(pb, r1cs_vk.rC_Z_g2, FMT(annotation_prefix, " vk_rC_Z_g2_precomp")));
}

template<typename ppT>
r1cs_ppzksnark_verifier_process_vk_gadget<ppT>::r1cs_ppzksnark_verifier_process_vk_gadget(protoboard<FieldT> &pb,
                                                                                          const r1cs_ppzksnark_verification_key_variable<ppT> &vk,
                                                                                          r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> &pvk,
                                                                                          const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    vk(vk),
    pvk(pvk)
{
    pvk.encoded_IC_base = vk.encoded_IC_base;
    pvk.encoded_IC_query = vk.encoded_IC_query;

    pvk.vk_alphaB_g1_precomp.reset(new G1_precomputation<ppT>());
    pvk.vk_gamma_beta_g1_precomp.reset(new G1_precomputation<ppT>());

    pvk.pp_G2_one_precomp.reset(new G2_precomputation<ppT>());
    pvk.vk_alphaA_g2_precomp.reset(new G2_precomputation<ppT>());
    pvk.vk_alphaC_g2_precomp.reset(new G2_precomputation<ppT>());
    pvk.vk_gamma_beta_g2_precomp.reset(new G2_precomputation<ppT>());
    pvk.vk_gamma_g2_precomp.reset(new G2_precomputation<ppT>());
    pvk.vk_rC_Z_g2_precomp.reset(new G2_precomputation<ppT>());

    compute_vk_alphaB_g1_precomp.reset(new precompute_G1_gadget<ppT>(pb, *vk.alphaB_g1, *pvk.vk_alphaB_g1_precomp, FMT(annotation_prefix, " compute_vk_alphaB_g1_precomp")));
    compute_vk_gamma_beta_g1_precomp.reset(new precompute_G1_gadget<ppT>(pb, *vk.gamma_beta_g1, *pvk.vk_gamma_beta_g1_precomp, FMT(annotation_prefix, " compute_vk_gamma_beta_g1_precomp")));

    pvk.pp_G2_one_precomp.reset(new G2_precomputation<ppT>(pb, G2<other_curve<ppT> >::one(), FMT(annotation_prefix, " pp_G2_one_precomp")));
    compute_vk_alphaA_g2_precomp.reset(new precompute_G2_gadget<ppT>(pb, *vk.alphaA_g2, *pvk.vk_alphaA_g2_precomp, FMT(annotation_prefix, " compute_vk_alphaA_g2_precomp")));
    compute_vk_alphaC_g2_precomp.reset(new precompute_G2_gadget<ppT>(pb, *vk.alphaC_g2, *pvk.vk_alphaC_g2_precomp, FMT(annotation_prefix, " compute_vk_alphaC_g2_precomp")));
    compute_vk_gamma_beta_g2_precomp.reset(new precompute_G2_gadget<ppT>(pb, *vk.gamma_beta_g2, *pvk.vk_gamma_beta_g2_precomp, FMT(annotation_prefix, " compute_vk_gamma_beta_g2_precomp")));
    compute_vk_gamma_g2_precomp.reset(new precompute_G2_gadget<ppT>(pb, *vk.gamma_g2, *pvk.vk_gamma_g2_precomp, FMT(annotation_prefix, " compute_vk_gamma_g2_precomp")));
    compute_vk_rC_Z_g2_precomp.reset(new precompute_G2_gadget<ppT>(pb, *vk.rC_Z_g2, *pvk.vk_rC_Z_g2_precomp, FMT(annotation_prefix, " compute_vk_rC_Z_g2_precomp")));
}

template<typename ppT>
void r1cs_ppzksnark_verifier_process_vk_gadget<ppT>::generate_r1cs_constraints()
{
    compute_vk_alphaB_g1_precomp->generate_r1cs_constraints();
    compute_vk_gamma_beta_g1_precomp->generate_r1cs_constraints();

    compute_vk_alphaA_g2_precomp->generate_r1cs_constraints();
    compute_vk_alphaC_g2_precomp->generate_r1cs_constraints();
    compute_vk_gamma_beta_g2_precomp->generate_r1cs_constraints();
    compute_vk_gamma_g2_precomp->generate_r1cs_constraints();
    compute_vk_rC_Z_g2_precomp->generate_r1cs_constraints();
}

template<typename ppT>
void r1cs_ppzksnark_verifier_process_vk_gadget<ppT>::generate_r1cs_witness()
{
    compute_vk_alphaB_g1_precomp->generate_r1cs_witness();
    compute_vk_gamma_beta_g1_precomp->generate_r1cs_witness();

    compute_vk_alphaA_g2_precomp->generate_r1cs_witness();
    compute_vk_alphaC_g2_precomp->generate_r1cs_witness();
    compute_vk_gamma_beta_g2_precomp->generate_r1cs_witness();
    compute_vk_gamma_g2_precomp->generate_r1cs_witness();
    compute_vk_rC_Z_g2_precomp->generate_r1cs_witness();
}

template<typename ppT>
r1cs_ppzksnark_online_verifier_gadget<ppT>::r1cs_ppzksnark_online_verifier_gadget(protoboard<FieldT> &pb,
                                                                                  const r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> &pvk,
                                                                                  const pb_variable_array<FieldT> &input,
                                                                                  const size_t elt_size,
                                                                                  const r1cs_ppzksnark_proof_variable<ppT> &proof,
                                                                                  const pb_variable<FieldT> &result,
                                                                                  const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    pvk(pvk),
    input(input),
    elt_size(elt_size),
    proof(proof),
    result(result),
    input_len(input.size())
{
    // accumulate input and store base in acc
    acc.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " acc")));
    std::vector<G1_variable<ppT> > IC_terms;
    for (size_t i = 0; i < pvk.encoded_IC_query.size(); ++i)
    {
        IC_terms.emplace_back(*(pvk.encoded_IC_query[i]));
    }
    accumulate_input.reset(new G1_multiscalar_mul_gadget<ppT>(pb, *(pvk.encoded_IC_base), input, elt_size, IC_terms, *acc, FMT(annotation_prefix, " accumulate_input")));

    // allocate results for precomputation
    proof_g_A_h_precomp.reset(new G1_precomputation<ppT>());
    proof_g_A_g_acc_C_precomp.reset(new G1_precomputation<ppT>());
    proof_g_A_g_acc_precomp.reset(new G1_precomputation<ppT>());
    proof_g_A_g_precomp.reset(new G1_precomputation<ppT>());
    proof_g_B_h_precomp.reset(new G1_precomputation<ppT>());
    proof_g_C_h_precomp.reset(new G1_precomputation<ppT>());
    proof_g_C_g_precomp.reset(new G1_precomputation<ppT>());
    proof_g_K_precomp.reset(new G1_precomputation<ppT>());
    proof_g_H_precomp.reset(new G1_precomputation<ppT>());

    proof_g_B_g_precomp.reset(new G2_precomputation<ppT>());

    // do the necessary precomputations
    // compute things not available in plain from proof/vk
    proof_g_A_g_acc.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " proof_g_A_g_acc")));
    compute_proof_g_A_g_acc.reset(new G1_add_gadget<ppT>(pb, *(proof.g_A_g), *acc , *proof_g_A_g_acc, FMT(annotation_prefix, " compute_proof_g_A_g_acc")));
    proof_g_A_g_acc_C.reset(new G1_variable<ppT>(pb, FMT(annotation_prefix, " proof_g_A_g_acc_C")));
    compute_proof_g_A_g_acc_C.reset(new G1_add_gadget<ppT>(pb, *proof_g_A_g_acc, *(proof.g_C_g) , *proof_g_A_g_acc_C, FMT(annotation_prefix, " compute_proof_g_A_g_acc_C")));

    compute_proof_g_A_g_acc_precomp.reset(new precompute_G1_gadget<ppT>(pb, *proof_g_A_g_acc, *proof_g_A_g_acc_precomp, FMT(annotation_prefix, " compute_proof_g_A_g_acc_precomp")));
    compute_proof_g_A_g_acc_C_precomp.reset(new precompute_G1_gadget<ppT>(pb, *proof_g_A_g_acc_C, *proof_g_A_g_acc_C_precomp, FMT(annotation_prefix, " compute_proof_g_A_g_acc_C_precomp")));

    // do other precomputations
    compute_proof_g_A_h_precomp.reset(new precompute_G1_gadget<ppT>(pb, *(proof.g_A_h), *proof_g_A_h_precomp, FMT(annotation_prefix, " compute_proof_g_A_h_precomp")));
    compute_proof_g_A_g_precomp.reset(new precompute_G1_gadget<ppT>(pb, *(proof.g_A_g), *proof_g_A_g_precomp, FMT(annotation_prefix, " compute_proof_g_A_g_precomp")));
    compute_proof_g_B_h_precomp.reset(new precompute_G1_gadget<ppT>(pb, *(proof.g_B_h), *proof_g_B_h_precomp, FMT(annotation_prefix, " compute_proof_g_B_h_precomp")));
    compute_proof_g_C_h_precomp.reset(new precompute_G1_gadget<ppT>(pb, *(proof.g_C_h), *proof_g_C_h_precomp, FMT(annotation_prefix, " compute_proof_g_C_h_precomp")));
    compute_proof_g_C_g_precomp.reset(new precompute_G1_gadget<ppT>(pb, *(proof.g_C_g), *proof_g_C_g_precomp, FMT(annotation_prefix, " compute_proof_g_C_g_precomp")));
    compute_proof_g_H_precomp.reset(new precompute_G1_gadget<ppT>(pb, *(proof.g_H), *proof_g_H_precomp, FMT(annotation_prefix, " compute_proof_g_H_precomp")));
    compute_proof_g_K_precomp.reset(new precompute_G1_gadget<ppT>(pb, *(proof.g_K), *proof_g_K_precomp, FMT(annotation_prefix, " compute_proof_g_K_precomp")));
    compute_proof_g_B_g_precomp.reset(new precompute_G2_gadget<ppT>(pb, *(proof.g_B_g), *proof_g_B_g_precomp, FMT(annotation_prefix, " compute_proof_g_B_g_precomp")));

    // check validity of A knowledge commitment
    kc_A_valid.allocate(pb, FMT(annotation_prefix, " kc_A_valid"));
    check_kc_A_valid.reset(new check_e_equals_e_gadget<ppT>(pb, *proof_g_A_g_precomp, *(pvk.vk_alphaA_g2_precomp), *proof_g_A_h_precomp, *(pvk.pp_G2_one_precomp), kc_A_valid, FMT(annotation_prefix, " check_kc_A_valid")));

    // check validity of B knowledge commitment
    kc_B_valid.allocate(pb, FMT(annotation_prefix, " kc_B_valid"));
    check_kc_B_valid.reset(new check_e_equals_e_gadget<ppT>(pb, *(pvk.vk_alphaB_g1_precomp), *proof_g_B_g_precomp, *proof_g_B_h_precomp, *(pvk.pp_G2_one_precomp), kc_B_valid, FMT(annotation_prefix, " check_kc_B_valid")));

    // check validity of C knowledge commitment
    kc_C_valid.allocate(pb, FMT(annotation_prefix, " kc_C_valid"));
    check_kc_C_valid.reset(new check_e_equals_e_gadget<ppT>(pb, *proof_g_C_g_precomp, *(pvk.vk_alphaC_g2_precomp), *proof_g_C_h_precomp, *(pvk.pp_G2_one_precomp), kc_C_valid, FMT(annotation_prefix, " check_kc_C_valid")));

    // check QAP divisibility
    QAP_valid.allocate(pb, FMT(annotation_prefix, " QAP_valid"));
    check_QAP_valid.reset(new check_e_equals_ee_gadget<ppT>(pb, *proof_g_A_g_acc_precomp, *proof_g_B_g_precomp, *proof_g_H_precomp, *(pvk.vk_rC_Z_g2_precomp), *proof_g_C_g_precomp, *(pvk.pp_G2_one_precomp), QAP_valid, FMT(annotation_prefix, " check_QAP_valid")));

    // check coefficients
    CC_valid.allocate(pb, FMT(annotation_prefix, " CC_valid"));
    check_CC_valid.reset(new check_e_equals_ee_gadget<ppT>(pb, *proof_g_K_precomp, *(pvk.vk_gamma_g2_precomp), *proof_g_A_g_acc_C_precomp, *(pvk.vk_gamma_beta_g2_precomp), *(pvk.vk_gamma_beta_g1_precomp), *proof_g_B_g_precomp, CC_valid, FMT(annotation_prefix, " check_CC_valid")));

    // final constraint
    all_test_results.emplace_back(kc_A_valid);
    all_test_results.emplace_back(kc_B_valid);
    all_test_results.emplace_back(kc_C_valid);
    all_test_results.emplace_back(QAP_valid);
    all_test_results.emplace_back(CC_valid);

    all_tests_pass.reset(new conjunction_gadget<FieldT>(pb, all_test_results, result, FMT(annotation_prefix, " all_tests_pass")));
}

template<typename ppT>
void r1cs_ppzksnark_online_verifier_gadget<ppT>::generate_r1cs_constraints()
{
    PROFILE_CONSTRAINTS(this->pb, "accumulate verifier input")
    {
        print_indent(); printf("* Number of bits as an input to verifier gadget: %zu\n", input.size());
        accumulate_input->generate_r1cs_constraints();
    }

    PROFILE_CONSTRAINTS(this->pb, "rest of the verifier")
    {
        compute_proof_g_A_g_acc->generate_r1cs_constraints();
        compute_proof_g_A_g_acc_C->generate_r1cs_constraints();

        compute_proof_g_A_g_acc_precomp->generate_r1cs_constraints();
        compute_proof_g_A_g_acc_C_precomp->generate_r1cs_constraints();

        compute_proof_g_A_h_precomp->generate_r1cs_constraints();
        compute_proof_g_A_g_precomp->generate_r1cs_constraints();
        compute_proof_g_B_h_precomp->generate_r1cs_constraints();
        compute_proof_g_C_h_precomp->generate_r1cs_constraints();
        compute_proof_g_C_g_precomp->generate_r1cs_constraints();
        compute_proof_g_H_precomp->generate_r1cs_constraints();
        compute_proof_g_K_precomp->generate_r1cs_constraints();
        compute_proof_g_B_g_precomp->generate_r1cs_constraints();

        check_kc_A_valid->generate_r1cs_constraints();
        check_kc_B_valid->generate_r1cs_constraints();
        check_kc_C_valid->generate_r1cs_constraints();
        check_QAP_valid->generate_r1cs_constraints();
        check_CC_valid->generate_r1cs_constraints();

        all_tests_pass->generate_r1cs_constraints();
    }
}

template<typename ppT>
void r1cs_ppzksnark_online_verifier_gadget<ppT>::generate_r1cs_witness()
{
    accumulate_input->generate_r1cs_witness();

    compute_proof_g_A_g_acc->generate_r1cs_witness();
    compute_proof_g_A_g_acc_C->generate_r1cs_witness();

    compute_proof_g_A_g_acc_precomp->generate_r1cs_witness();
    compute_proof_g_A_g_acc_C_precomp->generate_r1cs_witness();

    compute_proof_g_A_h_precomp->generate_r1cs_witness();
    compute_proof_g_A_g_precomp->generate_r1cs_witness();
    compute_proof_g_B_h_precomp->generate_r1cs_witness();
    compute_proof_g_C_h_precomp->generate_r1cs_witness();
    compute_proof_g_C_g_precomp->generate_r1cs_witness();
    compute_proof_g_H_precomp->generate_r1cs_witness();
    compute_proof_g_K_precomp->generate_r1cs_witness();
    compute_proof_g_B_g_precomp->generate_r1cs_witness();

    check_kc_A_valid->generate_r1cs_witness();
    check_kc_B_valid->generate_r1cs_witness();
    check_kc_C_valid->generate_r1cs_witness();
    check_QAP_valid->generate_r1cs_witness();
    check_CC_valid->generate_r1cs_witness();

    all_tests_pass->generate_r1cs_witness();
}

template<typename ppT>
r1cs_ppzksnark_verifier_gadget<ppT>::r1cs_ppzksnark_verifier_gadget(protoboard<FieldT> &pb,
                                                                    const r1cs_ppzksnark_verification_key_variable<ppT> &vk,
                                                                    const pb_variable_array<FieldT> &input,
                                                                    const size_t elt_size,
                                                                    const r1cs_ppzksnark_proof_variable<ppT> &proof,
                                                                    const pb_variable<FieldT> &result,
                                                                    const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix)
{
    pvk.reset(new r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT>());
    compute_pvk.reset(new r1cs_ppzksnark_verifier_process_vk_gadget<ppT>(pb, vk, *pvk, FMT(annotation_prefix, " compute_pvk")));
    online_verifier.reset(new r1cs_ppzksnark_online_verifier_gadget<ppT>(pb, *pvk, input, elt_size, proof, result, FMT(annotation_prefix, " online_verifier")));
}

template<typename ppT>
void r1cs_ppzksnark_verifier_gadget<ppT>::generate_r1cs_constraints()
{
    PROFILE_CONSTRAINTS(this->pb, "precompute pvk")
    {
        compute_pvk->generate_r1cs_constraints();
    }

    PROFILE_CONSTRAINTS(this->pb, "online verifier")
    {
        online_verifier->generate_r1cs_constraints();
    }
}

template<typename ppT>
void r1cs_ppzksnark_verifier_gadget<ppT>::generate_r1cs_witness()
{
    compute_pvk->generate_r1cs_witness();
    online_verifier->generate_r1cs_witness();
}

} // libsnark

#endif // R1CS_PPZKSNARK_VERIFIER_GADGET_TCC_

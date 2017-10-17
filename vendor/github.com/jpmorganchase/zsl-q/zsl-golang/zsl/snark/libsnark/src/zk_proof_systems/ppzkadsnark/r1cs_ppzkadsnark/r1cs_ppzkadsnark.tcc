/** @file
*****************************************************************************

Implementation of interfaces for a ppzkADSNARK for R1CS.

See r1cs_ppzkadsnark.hpp .

*****************************************************************************
* @author     This file is part of libsnark, developed by SCIPR Lab
*             and contributors (see AUTHORS).
* @copyright  MIT license (see LICENSE file)
*****************************************************************************/

#ifndef R1CS_PPZKADSNARK_TCC_
#define R1CS_PPZKADSNARK_TCC_

#include <algorithm>
#include <cassert>
#include <functional>
#include <iostream>
#include <sstream>

#include "common/profiling.hpp"
#include "common/utils.hpp"
#include "algebra/scalar_multiplication/multiexp.hpp"
#include "algebra/scalar_multiplication/kc_multiexp.hpp"
#include "reductions/r1cs_to_qap/r1cs_to_qap.hpp"

namespace libsnark {


template<typename ppT>
bool r1cs_ppzkadsnark_pub_auth_prms<ppT>::operator==(const r1cs_ppzkadsnark_pub_auth_prms<ppT> &other) const
{
    return (this->I1 == other.I1);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_pub_auth_prms<ppT> &pap)
{
    out << pap.I1;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_pub_auth_prms<ppT> &pap)
{
    in >> pap.I1;

    return in;
}

template<typename ppT>
bool r1cs_ppzkadsnark_sec_auth_key<ppT>::operator==(const r1cs_ppzkadsnark_sec_auth_key<ppT> &other) const
{
    return (this->i == other.i) &&
        (this->skp == other.skp) &&
        (this->S == other.S);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_sec_auth_key<ppT> &key)
{
    out << key.i;
    out << key.skp;
    out << key.S;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_sec_auth_key<ppT> &key)
{
    in >> key.i;
    in >> key.skp;
    in >> key.S;

    return in;
}

template<typename ppT>
bool r1cs_ppzkadsnark_pub_auth_key<ppT>::operator==(const r1cs_ppzkadsnark_pub_auth_key<ppT> &other) const
{
    return (this->minusI2 == other.minusI2) &&
        (this->vkp == other.vkp);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_pub_auth_key<ppT> &key)
{
    out << key.minusI2;
    out << key.vkp;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_pub_auth_key<ppT> &key)
{
    in >> key.minusI2;
    in >> key.vkp;

    return in;
}

template<typename ppT>
bool r1cs_ppzkadsnark_auth_data<ppT>::operator==(const r1cs_ppzkadsnark_auth_data<ppT> &other) const
{
    return (this->mu == other.mu) &&
        (this->Lambda == other.Lambda) &&
        (this->sigma == other.sigma);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_auth_data<ppT> &data)
{
    out << data.mu;
    out << data.Lambda;
    out << data.sigma;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_auth_data<ppT> &data)
{
    in >> data.mu;
    in >> data.Lambda;
    data.sigma;

    return in;
}

template<typename ppT>
bool r1cs_ppzkadsnark_proving_key<ppT>::operator==(const r1cs_ppzkadsnark_proving_key<ppT> &other) const
{
    return (this->A_query == other.A_query &&
            this->B_query == other.B_query &&
            this->C_query == other.C_query &&
            this->H_query == other.H_query &&
            this->K_query == other.K_query &&
            this->rA_i_Z_g1 == other.rA_i_Z_g1 &&
            this->constraint_system == other.constraint_system);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_proving_key<ppT> &pk)
{
    out << pk.A_query;
    out << pk.B_query;
    out << pk.C_query;
    out << pk.H_query;
    out << pk.K_query;
    out << pk.rA_i_Z_g1;
    out << pk.constraint_system;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_proving_key<ppT> &pk)
{
    in >> pk.A_query;
    in >> pk.B_query;
    in >> pk.C_query;
    in >> pk.H_query;
    in >> pk.K_query;
    in >> pk.rA_i_Z_g1;
    in >> pk.constraint_system;

    return in;
}

template<typename ppT>
bool r1cs_ppzkadsnark_verification_key<ppT>::operator==(const r1cs_ppzkadsnark_verification_key<ppT> &other) const
{
    return (this->alphaA_g2 == other.alphaA_g2 &&
            this->alphaB_g1 == other.alphaB_g1 &&
            this->alphaC_g2 == other.alphaC_g2 &&
            this->gamma_g2 == other.gamma_g2 &&
            this->gamma_beta_g1 == other.gamma_beta_g1 &&
            this->gamma_beta_g2 == other.gamma_beta_g2 &&
            this->rC_Z_g2 == other.rC_Z_g2 &&
            this->A0 == other.A0 &&
            this->Ain == other.Ain);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_verification_key<ppT> &vk)
{
    out << vk.alphaA_g2 << OUTPUT_NEWLINE;
    out << vk.alphaB_g1 << OUTPUT_NEWLINE;
    out << vk.alphaC_g2 << OUTPUT_NEWLINE;
    out << vk.gamma_g2 << OUTPUT_NEWLINE;
    out << vk.gamma_beta_g1 << OUTPUT_NEWLINE;
    out << vk.gamma_beta_g2 << OUTPUT_NEWLINE;
    out << vk.rC_Z_g2 << OUTPUT_NEWLINE;
    out << vk.A0 << OUTPUT_NEWLINE;
    out << vk.Ain << OUTPUT_NEWLINE;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_verification_key<ppT> &vk)
{
    in >> vk.alphaA_g2;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.alphaB_g1;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.alphaC_g2;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.gamma_g2;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.gamma_beta_g1;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.gamma_beta_g2;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.rC_Z_g2;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.A0;
    consume_OUTPUT_NEWLINE(in);
    in >> vk.Ain;
    consume_OUTPUT_NEWLINE(in);

    return in;
}

template<typename ppT>
bool r1cs_ppzkadsnark_processed_verification_key<ppT>::operator==(
    const r1cs_ppzkadsnark_processed_verification_key<ppT> &other) const
{
    bool result = (this->pp_G2_one_precomp == other.pp_G2_one_precomp &&
                   this->vk_alphaA_g2_precomp == other.vk_alphaA_g2_precomp &&
                   this->vk_alphaB_g1_precomp == other.vk_alphaB_g1_precomp &&
                   this->vk_alphaC_g2_precomp == other.vk_alphaC_g2_precomp &&
                   this->vk_rC_Z_g2_precomp == other.vk_rC_Z_g2_precomp &&
                   this->vk_gamma_g2_precomp == other.vk_gamma_g2_precomp &&
                   this->vk_gamma_beta_g1_precomp == other.vk_gamma_beta_g1_precomp &&
                   this->vk_gamma_beta_g2_precomp == other.vk_gamma_beta_g2_precomp &&
                   this->vk_rC_i_g2_precomp == other.vk_rC_i_g2_precomp &&
                   this->A0 == other.A0 &&
                   this->Ain == other.Ain &&
                   this->proof_g_vki_precomp.size() == other.proof_g_vki_precomp.size());
    if (result) {
        for(size_t i=0;i<this->proof_g_vki_precomp.size();i++)
            result &= this->proof_g_vki_precomp[i] == other.proof_g_vki_precomp[i];
    }
    return result;
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk)
{
    out << pvk.pp_G2_one_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_alphaA_g2_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_alphaB_g1_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_alphaC_g2_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_rC_Z_g2_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_gamma_g2_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_gamma_beta_g1_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_gamma_beta_g2_precomp << OUTPUT_NEWLINE;
    out << pvk.vk_rC_i_g2_precomp << OUTPUT_NEWLINE;
    out << pvk.A0 << OUTPUT_NEWLINE;
    out << pvk.Ain << OUTPUT_NEWLINE;
    out << pvk.proof_g_vki_precomp  << OUTPUT_NEWLINE;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk)
{
    in >> pvk.pp_G2_one_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_alphaA_g2_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_alphaB_g1_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_alphaC_g2_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_rC_Z_g2_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_gamma_g2_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_gamma_beta_g1_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_gamma_beta_g2_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.vk_rC_i_g2_precomp;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.A0;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.Ain;
    consume_OUTPUT_NEWLINE(in);
    in >> pvk.proof_g_vki_precomp;
    consume_OUTPUT_NEWLINE(in);

    return in;
}

template<typename ppT>
bool r1cs_ppzkadsnark_proof<ppT>::operator==(const r1cs_ppzkadsnark_proof<ppT> &other) const
{
    return (this->g_A == other.g_A &&
            this->g_B == other.g_B &&
            this->g_C == other.g_C &&
            this->g_H == other.g_H &&
            this->g_K == other.g_K &&
            this->g_Aau == other.g_Aau &&
            this->muA == other.muA);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_proof<ppT> &proof)
{
    out << proof.g_A << OUTPUT_NEWLINE;
    out << proof.g_B << OUTPUT_NEWLINE;
    out << proof.g_C << OUTPUT_NEWLINE;
    out << proof.g_H << OUTPUT_NEWLINE;
    out << proof.g_K << OUTPUT_NEWLINE;
    out << proof.g_Aau << OUTPUT_NEWLINE;
    out << proof.muA << OUTPUT_NEWLINE;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_proof<ppT> &proof)
{
    in >> proof.g_A;
    consume_OUTPUT_NEWLINE(in);
    in >> proof.g_B;
    consume_OUTPUT_NEWLINE(in);
    in >> proof.g_C;
    consume_OUTPUT_NEWLINE(in);
    in >> proof.g_H;
    consume_OUTPUT_NEWLINE(in);
    in >> proof.g_K;
    consume_OUTPUT_NEWLINE(in);
    in >> proof.g_Aau;
    consume_OUTPUT_NEWLINE(in);
    in >> proof.muA;
    consume_OUTPUT_NEWLINE(in);

    return in;
}

template<typename ppT>
r1cs_ppzkadsnark_verification_key<ppT> r1cs_ppzkadsnark_verification_key<ppT>::dummy_verification_key(const size_t input_size)
{
    r1cs_ppzkadsnark_verification_key<ppT> result;
    result.alphaA_g2 = Fr<snark_pp<ppT>>::random_element() * G2<snark_pp<ppT>>::one();
    result.alphaB_g1 = Fr<snark_pp<ppT>>::random_element() * G1<snark_pp<ppT>>::one();
    result.alphaC_g2 = Fr<snark_pp<ppT>>::random_element() * G2<snark_pp<ppT>>::one();
    result.gamma_g2 = Fr<snark_pp<ppT>>::random_element() * G2<snark_pp<ppT>>::one();
    result.gamma_beta_g1 = Fr<snark_pp<ppT>>::random_element() * G1<snark_pp<ppT>>::one();
    result.gamma_beta_g2 = Fr<snark_pp<ppT>>::random_element() * G2<snark_pp<ppT>>::one();
    result.rC_Z_g2 = Fr<snark_pp<ppT>>::random_element() * G2<snark_pp<ppT>>::one();

    result.A0 = Fr<snark_pp<ppT>>::random_element() * G1<snark_pp<ppT>>::one();
    for (size_t i = 0; i < input_size; ++i)
    {
        result.Ain.emplace_back(Fr<snark_pp<ppT>>::random_element() *
                                G1<snark_pp<ppT>>::one());
    }

    return result;
}

template<typename ppT>
r1cs_ppzkadsnark_auth_keys<ppT> r1cs_ppzkadsnark_auth_generator(void) {
    kpT<ppT> sigkp = sigGen<ppT>();
    r1cs_ppzkadsnark_prfKeyT<ppT>prfseed = prfGen<ppT>();
    Fr<snark_pp<ppT>> i = Fr<snark_pp<ppT>>::random_element();
    G1<snark_pp<ppT>> I1 = i * G1<snark_pp<ppT>>::one();
    G2<snark_pp<ppT>> minusI2 = G2<snark_pp<ppT>>::zero() -
        i * G2<snark_pp<ppT>>::one();
    return r1cs_ppzkadsnark_auth_keys<ppT>(
        r1cs_ppzkadsnark_pub_auth_prms<ppT>(std::move(I1)),
        r1cs_ppzkadsnark_pub_auth_key<ppT>(std::move(minusI2),std::move(sigkp.vk)),
        r1cs_ppzkadsnark_sec_auth_key<ppT>(std::move(i),std::move(sigkp.sk),std::move(prfseed)));
}

template<typename ppT>
std::vector<r1cs_ppzkadsnark_auth_data<ppT>> r1cs_ppzkadsnark_auth_sign(
    const std::vector<Fr<snark_pp<ppT>>> &ins,
    const r1cs_ppzkadsnark_sec_auth_key<ppT> &sk,
    const std::vector<labelT> labels) {
    enter_block("Call to r1cs_ppzkadsnark_auth_sign");
    assert (labels.size()==ins.size());
    std::vector<r1cs_ppzkadsnark_auth_data<ppT>> res;
    res.reserve(ins.size());
    for (size_t i = 0; i < ins.size();i++) {
        Fr<snark_pp<ppT>> lambda = prfCompute<ppT>(sk.S,labels[i]);
        G2<snark_pp<ppT>> Lambda = lambda * G2<snark_pp<ppT>>::one();
        r1cs_ppzkadsnark_sigT<ppT>sig = sigSign<ppT>(sk.skp,labels[i],Lambda);
        r1cs_ppzkadsnark_auth_data<ppT> val(std::move(lambda + sk.i * ins[i]),
                                            std::move(Lambda),
                                            std::move(sig));
        res.emplace_back(val);
    }
    leave_block("Call to r1cs_ppzkadsnark_auth_sign");
    return std::move(res);
}

// symmetric
template<typename ppT>
bool r1cs_ppzkadsnark_auth_verify(const std::vector<Fr<snark_pp<ppT>>> &data,
                                  const std::vector<r1cs_ppzkadsnark_auth_data<ppT>> & auth_data,
                                  const r1cs_ppzkadsnark_sec_auth_key<ppT> &sak,
                                  const std::vector<labelT> &labels) {
    enter_block("Call to r1cs_ppzkadsnark_auth_verify");
    assert ((data.size()==labels.size()) && (auth_data.size()==labels.size()));
    bool res = true;
    for (size_t i = 0; i < data.size();i++) {
        Fr<snark_pp<ppT>> lambda = prfCompute<ppT>(sak.S,labels[i]);
        Fr<snark_pp<ppT>> mup = lambda + sak.i * data[i];
        res = res && (auth_data[i].mu == mup);
    }
    leave_block("Call to r1cs_ppzkadsnark_auth_verify");
    return res;
}

// public
template<typename ppT>
bool r1cs_ppzkadsnark_auth_verify(const std::vector<Fr<snark_pp<ppT>>> &data,
                                  const std::vector<r1cs_ppzkadsnark_auth_data<ppT>> & auth_data,
                                  const r1cs_ppzkadsnark_pub_auth_key<ppT> &pak,
                                  const std::vector<labelT> &labels) {
    enter_block("Call to r1cs_ppzkadsnark_auth_verify");
    assert ((data.size()==labels.size()) && (data.size()==auth_data.size()));
    bool res = true;
    for (size_t i = 0; i < auth_data.size();i++) {
        G2<snark_pp<ppT>> Mup = auth_data[i].Lambda - data[i] * pak.minusI2;
        res = res && (auth_data[i].mu * G2<snark_pp<ppT>>::one() == Mup);
        res = res && sigVerif<ppT>(pak.vkp,labels[i],auth_data[i].Lambda,auth_data[i].sigma);
    }
    leave_block("Call to r1cs_ppzkadsnark_auth_verify");
    return res;
}

template <typename ppT>
r1cs_ppzkadsnark_keypair<ppT> r1cs_ppzkadsnark_generator(const r1cs_ppzkadsnark_constraint_system<ppT> &cs,
                                                         const r1cs_ppzkadsnark_pub_auth_prms<ppT> &prms)
{
    enter_block("Call to r1cs_ppzkadsnark_generator");

    /* make the B_query "lighter" if possible */
    r1cs_ppzkadsnark_constraint_system<ppT> cs_copy(cs);
    cs_copy.swap_AB_if_beneficial();

    /* draw random element at which the QAP is evaluated */
    const  Fr<snark_pp<ppT>> t = Fr<snark_pp<ppT>>::random_element();

    qap_instance_evaluation<Fr<snark_pp<ppT>> > qap_inst =
        r1cs_to_qap_instance_map_with_evaluation(cs_copy, t);

    print_indent(); printf("* QAP number of variables: %zu\n", qap_inst.num_variables());
    print_indent(); printf("* QAP pre degree: %zu\n", cs_copy.constraints.size());
    print_indent(); printf("* QAP degree: %zu\n", qap_inst.degree());
    print_indent(); printf("* QAP number of input variables: %zu\n", qap_inst.num_inputs());

    enter_block("Compute query densities");
    size_t non_zero_At = 0, non_zero_Bt = 0, non_zero_Ct = 0, non_zero_Ht = 0;
    for (size_t i = 0; i < qap_inst.num_variables()+1; ++i)
    {
        if (!qap_inst.At[i].is_zero())
        {
            ++non_zero_At;
        }
        if (!qap_inst.Bt[i].is_zero())
        {
            ++non_zero_Bt;
        }
        if (!qap_inst.Ct[i].is_zero())
        {
            ++non_zero_Ct;
        }
    }
    for (size_t i = 0; i < qap_inst.degree()+1; ++i)
    {
        if (!qap_inst.Ht[i].is_zero())
        {
            ++non_zero_Ht;
        }
    }
    leave_block("Compute query densities");

    Fr_vector<snark_pp<ppT>> At = std::move(qap_inst.At); // qap_inst.At is now in unspecified state, but we do not use it later
    Fr_vector<snark_pp<ppT>> Bt = std::move(qap_inst.Bt); // qap_inst.Bt is now in unspecified state, but we do not use it later
    Fr_vector<snark_pp<ppT>> Ct = std::move(qap_inst.Ct); // qap_inst.Ct is now in unspecified state, but we do not use it later
    Fr_vector<snark_pp<ppT>> Ht = std::move(qap_inst.Ht); // qap_inst.Ht is now in unspecified state, but we do not use it later

    /* append Zt to At,Bt,Ct with */
    At.emplace_back(qap_inst.Zt);
    Bt.emplace_back(qap_inst.Zt);
    Ct.emplace_back(qap_inst.Zt);

    const  Fr<snark_pp<ppT>> alphaA = Fr<snark_pp<ppT>>::random_element(),
        alphaB = Fr<snark_pp<ppT>>::random_element(),
        alphaC = Fr<snark_pp<ppT>>::random_element(),
        rA = Fr<snark_pp<ppT>>::random_element(),
        rB = Fr<snark_pp<ppT>>::random_element(),
        beta = Fr<snark_pp<ppT>>::random_element(),
        gamma = Fr<snark_pp<ppT>>::random_element();
    const Fr<snark_pp<ppT>>      rC = rA * rB;

    // consrtuct the same-coefficient-check query (must happen before zeroing out the prefix of At)
    Fr_vector<snark_pp<ppT>> Kt;
    Kt.reserve(qap_inst.num_variables()+4);
    for (size_t i = 0; i < qap_inst.num_variables()+1; ++i)
    {
        Kt.emplace_back( beta * (rA * At[i] + rB * Bt[i] + rC * Ct[i] ) );
    }
    Kt.emplace_back(beta * rA * qap_inst.Zt);
    Kt.emplace_back(beta * rB * qap_inst.Zt);
    Kt.emplace_back(beta * rC * qap_inst.Zt);

    const size_t g1_exp_count = 2*(non_zero_At - qap_inst.num_inputs() + non_zero_Ct) + non_zero_Bt + non_zero_Ht + Kt.size();
    const size_t g2_exp_count = non_zero_Bt;

    size_t g1_window = get_exp_window_size<G1<snark_pp<ppT>> >(g1_exp_count);
    size_t g2_window = get_exp_window_size<G2<snark_pp<ppT>> >(g2_exp_count);
    print_indent(); printf("* G1 window: %zu\n", g1_window);
    print_indent(); printf("* G2 window: %zu\n", g2_window);

#ifdef MULTICORE
    const size_t chunks = omp_get_max_threads(); // to override, set OMP_NUM_THREADS env var or call omp_set_num_threads()
#else
    const size_t chunks = 1;
#endif

    enter_block("Generating G1 multiexp table");
    window_table<G1<snark_pp<ppT>> > g1_table =
        get_window_table(Fr<snark_pp<ppT>>::size_in_bits(), g1_window,
                         G1<snark_pp<ppT>>::one());
    leave_block("Generating G1 multiexp table");

    enter_block("Generating G2 multiexp table");
    window_table<G2<snark_pp<ppT>> > g2_table =
        get_window_table(Fr<snark_pp<ppT>>::size_in_bits(),
                         g2_window, G2<snark_pp<ppT>>::one());
    leave_block("Generating G2 multiexp table");

    enter_block("Generate R1CS proving key");

    enter_block("Generate knowledge commitments");
    enter_block("Compute the A-query", false);
    knowledge_commitment_vector<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > A_query =
        kc_batch_exp(Fr<snark_pp<ppT>>::size_in_bits(), g1_window, g1_window, g1_table,
                     g1_table, rA, rA*alphaA, At, chunks);
    leave_block("Compute the A-query", false);

    enter_block("Compute the B-query", false);
    knowledge_commitment_vector<G2<snark_pp<ppT>>, G1<snark_pp<ppT>> > B_query =
        kc_batch_exp(Fr<snark_pp<ppT>>::size_in_bits(), g2_window, g1_window, g2_table,
                     g1_table, rB, rB*alphaB, Bt, chunks);
    leave_block("Compute the B-query", false);

    enter_block("Compute the C-query", false);
    knowledge_commitment_vector<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > C_query =
        kc_batch_exp(Fr<snark_pp<ppT>>::size_in_bits(), g1_window, g1_window, g1_table,
                     g1_table, rC, rC*alphaC, Ct, chunks);
    leave_block("Compute the C-query", false);

    enter_block("Compute the H-query", false);
    G1_vector<snark_pp<ppT>> H_query = batch_exp(Fr<snark_pp<ppT>>::size_in_bits(), g1_window, g1_table, Ht);
    leave_block("Compute the H-query", false);

    enter_block("Compute the K-query", false);
    G1_vector<snark_pp<ppT>> K_query = batch_exp(Fr<snark_pp<ppT>>::size_in_bits(), g1_window, g1_table, Kt);
#ifdef USE_MIXED_ADDITION
    batch_to_special<G1<snark_pp<ppT>> >(K_query);
#endif
    leave_block("Compute the K-query", false);

    leave_block("Generate knowledge commitments");

    leave_block("Generate R1CS proving key");

    enter_block("Generate R1CS verification key");
    G2<snark_pp<ppT>> alphaA_g2 = alphaA * G2<snark_pp<ppT>>::one();
    G1<snark_pp<ppT>> alphaB_g1 = alphaB * G1<snark_pp<ppT>>::one();
    G2<snark_pp<ppT>> alphaC_g2 = alphaC * G2<snark_pp<ppT>>::one();
    G2<snark_pp<ppT>> gamma_g2 = gamma * G2<snark_pp<ppT>>::one();
    G1<snark_pp<ppT>> gamma_beta_g1 = (gamma * beta) * G1<snark_pp<ppT>>::one();
    G2<snark_pp<ppT>> gamma_beta_g2 = (gamma * beta) * G2<snark_pp<ppT>>::one();
    G2<snark_pp<ppT>> rC_Z_g2 = (rC * qap_inst.Zt) * G2<snark_pp<ppT>>::one();

    enter_block("Generate extra authentication elements");
    G1<snark_pp<ppT>> rA_i_Z_g1 = (rA * qap_inst.Zt) * prms.I1;
    leave_block("Generate extra authentication elements");

    enter_block("Copy encoded input coefficients for R1CS verification key");
    G1<snark_pp<ppT>> A0 = A_query[0].g;
    G1_vector<snark_pp<ppT>> Ain;
    Ain.reserve(qap_inst.num_inputs());
    for (size_t i = 0; i < qap_inst.num_inputs(); ++i)
    {
        Ain.emplace_back(A_query[1+i].g);
    }

    leave_block("Copy encoded input coefficients for R1CS verification key");

    leave_block("Generate R1CS verification key");

    leave_block("Call to r1cs_ppzkadsnark_generator");

    r1cs_ppzkadsnark_verification_key<ppT> vk = r1cs_ppzkadsnark_verification_key<ppT>(alphaA_g2,
                                                                                       alphaB_g1,
                                                                                       alphaC_g2,
                                                                                       gamma_g2,
                                                                                       gamma_beta_g1,
                                                                                       gamma_beta_g2,
                                                                                       rC_Z_g2,
                                                                                       A0,
                                                                                       Ain);
    r1cs_ppzkadsnark_proving_key<ppT> pk = r1cs_ppzkadsnark_proving_key<ppT>(std::move(A_query),
                                                                             std::move(B_query),
                                                                             std::move(C_query),
                                                                             std::move(H_query),
                                                                             std::move(K_query),
                                                                             std::move(rA_i_Z_g1),
                                                                             std::move(cs_copy));

    pk.print_size();
    vk.print_size();

    return r1cs_ppzkadsnark_keypair<ppT>(std::move(pk), std::move(vk));
}

template <typename ppT>
r1cs_ppzkadsnark_proof<ppT> r1cs_ppzkadsnark_prover(const r1cs_ppzkadsnark_proving_key<ppT> &pk,
                                                    const r1cs_ppzkadsnark_primary_input<ppT> &primary_input,
                                                    const r1cs_ppzkadsnark_auxiliary_input<ppT> &auxiliary_input,
                                                    const std::vector<r1cs_ppzkadsnark_auth_data<ppT>> &auth_data)
{
    enter_block("Call to r1cs_ppzkadsnark_prover");

#ifdef DEBUG
    assert(pk.constraint_system.is_satisfied(primary_input, auxiliary_input));
#endif

    const Fr<snark_pp<ppT>> d1 = Fr<snark_pp<ppT>>::random_element(),
        d2 = Fr<snark_pp<ppT>>::random_element(),
        d3 = Fr<snark_pp<ppT>>::random_element(),
        dauth = Fr<snark_pp<ppT>>::random_element();

    enter_block("Compute the polynomial H");
    const qap_witness<Fr<snark_pp<ppT>> > qap_wit = r1cs_to_qap_witness_map(pk.constraint_system, primary_input,
                                                                            auxiliary_input, d1 + dauth, d2, d3);
    leave_block("Compute the polynomial H");

#ifdef DEBUG
    const Fr<snark_pp<ppT>> t = Fr<snark_pp<ppT>>::random_element();
    qap_instance_evaluation<Fr<snark_pp<ppT>> > qap_inst = r1cs_to_qap_instance_map_with_evaluation(pk.constraint_system, t);
    assert(qap_inst.is_satisfied(qap_wit));
#endif

    knowledge_commitment<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_A =
        /* pk.A_query[0] + */ d1*pk.A_query[qap_wit.num_variables()+1];
    knowledge_commitment<G2<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_B =
        pk.B_query[0] + qap_wit.d2*pk.B_query[qap_wit.num_variables()+1];
    knowledge_commitment<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_C =
        pk.C_query[0] + qap_wit.d3*pk.C_query[qap_wit.num_variables()+1];

    knowledge_commitment<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_Ain = dauth*pk.A_query[qap_wit.num_variables()+1];

    G1<snark_pp<ppT>> g_H = G1<snark_pp<ppT>>::zero();
    G1<snark_pp<ppT>> g_K = (pk.K_query[0] +
                             qap_wit.d1*pk.K_query[qap_wit.num_variables()+1] +
                             qap_wit.d2*pk.K_query[qap_wit.num_variables()+2] +
                             qap_wit.d3*pk.K_query[qap_wit.num_variables()+3]);

#ifdef DEBUG
    for (size_t i = 0; i < qap_wit.num_inputs() + 1; ++i)
    {
        assert(pk.A_query[i].g == G1<snark_pp<ppT>>::zero());
    }
    assert(pk.A_query.domain_size() == qap_wit.num_variables()+2);
    assert(pk.B_query.domain_size() == qap_wit.num_variables()+2);
    assert(pk.C_query.domain_size() == qap_wit.num_variables()+2);
    assert(pk.H_query.size() == qap_wit.degree()+1);
    assert(pk.K_query.size() == qap_wit.num_variables()+4);
#endif

#ifdef MULTICORE
    const size_t chunks = omp_get_max_threads(); // to override, set OMP_NUM_THREADS env var or call omp_set_num_threads()
#else
    const size_t chunks = 1;
#endif

    enter_block("Compute the proof");

    enter_block("Compute answer to A-query", false);
    g_A = g_A + kc_multi_exp_with_mixed_addition<G1<snark_pp<ppT>>, G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(pk.A_query,
                                                                                                           1+qap_wit.num_inputs(), 1+qap_wit.num_variables(),
                                                                                                           qap_wit.coefficients_for_ABCs.begin()+qap_wit.num_inputs(),
                                                                                                           qap_wit.coefficients_for_ABCs.begin()+qap_wit.num_variables(),
                                                                                                           chunks, true);
    leave_block("Compute answer to A-query", false);

    enter_block("Compute answer to Ain-query", false);
    g_Ain = g_Ain + kc_multi_exp_with_mixed_addition<G1<snark_pp<ppT>>, G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(pk.A_query,
                                                                                                               1, 1+qap_wit.num_inputs(),
                                                                                                               qap_wit.coefficients_for_ABCs.begin(),
                                                                                                               qap_wit.coefficients_for_ABCs.begin()+qap_wit.num_inputs(),
                                                                                                               chunks, true);
    //std :: cout << "The input proof term: " << g_Ain << "\n";
    leave_block("Compute answer to Ain-query", false);

    enter_block("Compute answer to B-query", false);
    g_B = g_B + kc_multi_exp_with_mixed_addition<G2<snark_pp<ppT>>, G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(pk.B_query,
                                                                                                           1, 1+qap_wit.num_variables(),
                                                                                                           qap_wit.coefficients_for_ABCs.begin(),
                                                                                                           qap_wit.coefficients_for_ABCs.begin()+qap_wit.num_variables(),
                                                                                                           chunks, true);
    leave_block("Compute answer to B-query", false);

    enter_block("Compute answer to C-query", false);
    g_C = g_C + kc_multi_exp_with_mixed_addition<G1<snark_pp<ppT>>, G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(pk.C_query,
                                                                                                           1, 1+qap_wit.num_variables(),
                                                                                                           qap_wit.coefficients_for_ABCs.begin(),
                                                                                                           qap_wit.coefficients_for_ABCs.begin()+qap_wit.num_variables(),
                                                                                                           chunks, true);
    leave_block("Compute answer to C-query", false);

    enter_block("Compute answer to H-query", false);
    g_H = g_H + multi_exp<G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(pk.H_query.begin(),
                                                                 pk.H_query.begin()+qap_wit.degree()+1,
                                                                 qap_wit.coefficients_for_H.begin(),
                                                                 qap_wit.coefficients_for_H.begin()+qap_wit.degree()+1,
                                                                 chunks, true);
    leave_block("Compute answer to H-query", false);

    enter_block("Compute answer to K-query", false);
    g_K = g_K + multi_exp_with_mixed_addition<G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(pk.K_query.begin()+1,
                                                                                     pk.K_query.begin()+1+qap_wit.num_variables(),
                                                                                     qap_wit.coefficients_for_ABCs.begin(),
                                                                                     qap_wit.coefficients_for_ABCs.begin()+qap_wit.num_variables(),
                                                                                     chunks, true);
    leave_block("Compute answer to K-query", false);

    enter_block("Compute extra auth terms", false);
    std::vector<Fr<snark_pp<ppT>>> mus;
    std::vector<G1<snark_pp<ppT>>> Ains;
    mus.reserve(qap_wit.num_inputs());
    Ains.reserve(qap_wit.num_inputs());
    for (size_t i=0;i<qap_wit.num_inputs();i++) {
        mus.emplace_back(auth_data[i].mu);
        Ains.emplace_back(pk.A_query[i+1].g);
    }
    G1<snark_pp<ppT>> muA = dauth * pk.rA_i_Z_g1;
    muA = muA + multi_exp<G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(Ains.begin(), Ains.begin()+qap_wit.num_inputs(),
                                                                 mus.begin(), mus.begin()+qap_wit.num_inputs(),
                                                                 chunks, true);

    // To Do: Decide whether to include relevant parts of auth_data in proof
    leave_block("Compute extra auth terms", false);

    leave_block("Compute the proof");

    leave_block("Call to r1cs_ppzkadsnark_prover");

    r1cs_ppzkadsnark_proof<ppT> proof = r1cs_ppzkadsnark_proof<ppT>(std::move(g_A),
                                                                    std::move(g_B),
                                                                    std::move(g_C),
                                                                    std::move(g_H),
                                                                    std::move(g_K),
                                                                    std::move(g_Ain),
                                                                    std::move(muA));
    proof.print_size();

    return proof;
}

template <typename ppT>
r1cs_ppzkadsnark_processed_verification_key<ppT> r1cs_ppzkadsnark_verifier_process_vk(
    const r1cs_ppzkadsnark_verification_key<ppT> &vk)
{
    enter_block("Call to r1cs_ppzkadsnark_verifier_process_vk");

    r1cs_ppzkadsnark_processed_verification_key<ppT> pvk;
    pvk.pp_G2_one_precomp        = snark_pp<ppT>::precompute_G2(G2<snark_pp<ppT>>::one());
    pvk.vk_alphaA_g2_precomp     = snark_pp<ppT>::precompute_G2(vk.alphaA_g2);
    pvk.vk_alphaB_g1_precomp     = snark_pp<ppT>::precompute_G1(vk.alphaB_g1);
    pvk.vk_alphaC_g2_precomp     = snark_pp<ppT>::precompute_G2(vk.alphaC_g2);
    pvk.vk_rC_Z_g2_precomp       = snark_pp<ppT>::precompute_G2(vk.rC_Z_g2);
    pvk.vk_gamma_g2_precomp      = snark_pp<ppT>::precompute_G2(vk.gamma_g2);
    pvk.vk_gamma_beta_g1_precomp = snark_pp<ppT>::precompute_G1(vk.gamma_beta_g1);
    pvk.vk_gamma_beta_g2_precomp = snark_pp<ppT>::precompute_G2(vk.gamma_beta_g2);

    enter_block("Pre-processing for additional auth elements");
    G2_precomp<snark_pp<ppT>> vk_rC_z_g2_precomp = snark_pp<ppT>::precompute_G2(vk.rC_Z_g2);

    pvk.A0 = G1<snark_pp<ppT>>(vk.A0);
    pvk.Ain = G1_vector<snark_pp<ppT>>(vk.Ain);

    pvk.proof_g_vki_precomp.reserve(pvk.Ain.size());
    for(size_t i = 0; i < pvk.Ain.size();i++) {
        pvk.proof_g_vki_precomp.emplace_back(snark_pp<ppT>::precompute_G1(pvk.Ain[i]));
    }

    leave_block("Pre-processing for additional auth elements");

    leave_block("Call to r1cs_ppzkadsnark_verifier_process_vk");

    return pvk;
}

// symmetric
template<typename ppT>
bool r1cs_ppzkadsnark_online_verifier(const r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk,
                                      const r1cs_ppzkadsnark_proof<ppT> &proof,
                                      const r1cs_ppzkadsnark_sec_auth_key<ppT> & sak,
                                      const std::vector<labelT> &labels)
{
    bool result = true;
    enter_block("Call to r1cs_ppzkadsnark_online_verifier");

    enter_block("Check if the proof is well-formed");
    if (!proof.is_well_formed())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("At least one of the proof elements does not lie on the curve.\n");
        }
        result = false;
    }
    leave_block("Check if the proof is well-formed");

    enter_block("Checking auth-specific elements");

    enter_block("Checking A1");

    enter_block("Compute PRFs");
    std::vector<Fr<snark_pp<ppT>>>lambdas;
    lambdas.reserve(labels.size());
    for (size_t i = 0; i < labels.size();i++) {
        lambdas.emplace_back(prfCompute<ppT>(sak.S,labels[i]));
    }
    leave_block("Compute PRFs");
    G1<snark_pp<ppT>> prodA = sak.i * proof.g_Aau.g;
    prodA = prodA + multi_exp<G1<snark_pp<ppT>>, Fr<snark_pp<ppT>> >(pvk.Ain.begin(),
                                                                     pvk.Ain.begin() + labels.size(),
                                                                     lambdas.begin(),
                                                                     lambdas.begin() + labels.size(), 1, true);

    bool result_auth = true;

    if (!(prodA == proof.muA)) {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Authentication check failed.\n");
        }
        result_auth = false;
    }

    leave_block("Checking A1");

    enter_block("Checking A2");
    G1_precomp<snark_pp<ppT>> proof_g_Aau_g_precomp      = snark_pp<ppT>::precompute_G1(proof.g_Aau.g);
    G1_precomp<snark_pp<ppT>> proof_g_Aau_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_Aau.h);
    Fqk<snark_pp<ppT>> kc_Aau_1 = snark_pp<ppT>::miller_loop(proof_g_Aau_g_precomp, pvk.vk_alphaA_g2_precomp);
    Fqk<snark_pp<ppT>> kc_Aau_2 = snark_pp<ppT>::miller_loop(proof_g_Aau_h_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_Aau = snark_pp<ppT>::final_exponentiation(kc_Aau_1 * kc_Aau_2.unitary_inverse());
    if (kc_Aau != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for Aau query incorrect.\n");
        }
        result_auth = false;
    }
    leave_block("Checking A2");

    leave_block("Checking auth-specific elements");

    result &= result_auth;

    enter_block("Online pairing computations");
    enter_block("Check knowledge commitment for A is valid");
    G1_precomp<snark_pp<ppT>> proof_g_A_g_precomp      = snark_pp<ppT>::precompute_G1(proof.g_A.g);
    G1_precomp<snark_pp<ppT>> proof_g_A_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_A.h);
    Fqk<snark_pp<ppT>> kc_A_1 = snark_pp<ppT>::miller_loop(proof_g_A_g_precomp,      pvk.vk_alphaA_g2_precomp);
    Fqk<snark_pp<ppT>> kc_A_2 = snark_pp<ppT>::miller_loop(proof_g_A_h_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_A = snark_pp<ppT>::final_exponentiation(kc_A_1 * kc_A_2.unitary_inverse());
    if (kc_A != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for A query incorrect.\n");
        }
        result = false;
    }
    leave_block("Check knowledge commitment for A is valid");

    enter_block("Check knowledge commitment for B is valid");
    G2_precomp<snark_pp<ppT>> proof_g_B_g_precomp      = snark_pp<ppT>::precompute_G2(proof.g_B.g);
    G1_precomp<snark_pp<ppT>> proof_g_B_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_B.h);
    Fqk<snark_pp<ppT>> kc_B_1 = snark_pp<ppT>::miller_loop(pvk.vk_alphaB_g1_precomp, proof_g_B_g_precomp);
    Fqk<snark_pp<ppT>> kc_B_2 = snark_pp<ppT>::miller_loop(proof_g_B_h_precomp,    pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_B = snark_pp<ppT>::final_exponentiation(kc_B_1 * kc_B_2.unitary_inverse());
    if (kc_B != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for B query incorrect.\n");
        }
        result = false;
    }
    leave_block("Check knowledge commitment for B is valid");

    enter_block("Check knowledge commitment for C is valid");
    G1_precomp<snark_pp<ppT>> proof_g_C_g_precomp      = snark_pp<ppT>::precompute_G1(proof.g_C.g);
    G1_precomp<snark_pp<ppT>> proof_g_C_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_C.h);
    Fqk<snark_pp<ppT>> kc_C_1 = snark_pp<ppT>::miller_loop(proof_g_C_g_precomp,      pvk.vk_alphaC_g2_precomp);
    Fqk<snark_pp<ppT>> kc_C_2 = snark_pp<ppT>::miller_loop(proof_g_C_h_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_C = snark_pp<ppT>::final_exponentiation(kc_C_1 * kc_C_2.unitary_inverse());
    if (kc_C != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for C query incorrect.\n");
        }
        result = false;
    }
    leave_block("Check knowledge commitment for C is valid");

    G1<snark_pp<ppT>> Aacc = pvk.A0 + proof.g_Aau.g + proof.g_A.g;

    enter_block("Check QAP divisibility");
    G1_precomp<snark_pp<ppT>> proof_g_Aacc_precomp = snark_pp<ppT>::precompute_G1(Aacc);
    G1_precomp<snark_pp<ppT>> proof_g_H_precomp = snark_pp<ppT>::precompute_G1(proof.g_H);
    Fqk<snark_pp<ppT>> QAP_1  = snark_pp<ppT>::miller_loop(proof_g_Aacc_precomp,  proof_g_B_g_precomp);
    Fqk<snark_pp<ppT>> QAP_23  = snark_pp<ppT>::double_miller_loop(proof_g_H_precomp, pvk.vk_rC_Z_g2_precomp,
                                                                   proof_g_C_g_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> QAP = snark_pp<ppT>::final_exponentiation(QAP_1 * QAP_23.unitary_inverse());
    if (QAP != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("QAP divisibility check failed.\n");
        }
        result = false;
    }
    leave_block("Check QAP divisibility");

    enter_block("Check same coefficients were used");
    G1_precomp<snark_pp<ppT>> proof_g_K_precomp = snark_pp<ppT>::precompute_G1(proof.g_K);
    G1_precomp<snark_pp<ppT>> proof_g_Aacc_C_precomp = snark_pp<ppT>::precompute_G1(Aacc + proof.g_C.g);
    Fqk<snark_pp<ppT>> K_1 = snark_pp<ppT>::miller_loop(proof_g_K_precomp, pvk.vk_gamma_g2_precomp);
    Fqk<snark_pp<ppT>> K_23 = snark_pp<ppT>::double_miller_loop(proof_g_Aacc_C_precomp, pvk.vk_gamma_beta_g2_precomp,
                                                                pvk.vk_gamma_beta_g1_precomp, proof_g_B_g_precomp);
    GT<snark_pp<ppT>> K = snark_pp<ppT>::final_exponentiation(K_1 * K_23.unitary_inverse());
    if (K != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Same-coefficient check failed.\n");
        }
        result = false;
    }
    leave_block("Check same coefficients were used");
    leave_block("Online pairing computations");
    leave_block("Call to r1cs_ppzkadsnark_online_verifier");

    return result;
}

template<typename ppT>
bool r1cs_ppzkadsnark_verifier(const r1cs_ppzkadsnark_verification_key<ppT> &vk,
                               const r1cs_ppzkadsnark_proof<ppT> &proof,
                               const r1cs_ppzkadsnark_sec_auth_key<ppT> &sak,
                               const std::vector<labelT> &labels)
{
    enter_block("Call to r1cs_ppzkadsnark_verifier");
    r1cs_ppzkadsnark_processed_verification_key<ppT> pvk = r1cs_ppzkadsnark_verifier_process_vk<ppT>(vk);
    bool result = r1cs_ppzkadsnark_online_verifier<ppT>(pvk, proof, sak, labels);
    leave_block("Call to r1cs_ppzkadsnark_verifier");
    return result;
}


// public
template<typename ppT>
bool r1cs_ppzkadsnark_online_verifier(const r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk,
                                      const std::vector<r1cs_ppzkadsnark_auth_data<ppT>>  &auth_data,
                                      const r1cs_ppzkadsnark_proof<ppT> &proof,
                                      const r1cs_ppzkadsnark_pub_auth_key<ppT> & pak,
                                      const std::vector<labelT> &labels)
{
    bool result = true;
    enter_block("Call to r1cs_ppzkadsnark_online_verifier");

    enter_block("Check if the proof is well-formed");
    if (!proof.is_well_formed())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("At least one of the proof elements does not lie on the curve.\n");
        }
        result = false;
    }
    leave_block("Check if the proof is well-formed");

    enter_block("Checking auth-specific elements");
    assert (labels.size()==auth_data.size());

    enter_block("Checking A1");

    enter_block("Checking signatures");
    std::vector<G2<snark_pp<ppT>>> Lambdas;
    std::vector<r1cs_ppzkadsnark_sigT<ppT>> sigs;
    Lambdas.reserve(labels.size());
    sigs.reserve(labels.size());
    for (size_t i = 0; i < labels.size();i++) {
        Lambdas.emplace_back(auth_data[i].Lambda);
        sigs.emplace_back(auth_data[i].sigma);
    }
    bool result_auth = sigBatchVerif<ppT>(pak.vkp,labels,Lambdas,sigs);
    if (! result_auth)
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Auth sig check failed.\n");
        }
    }

    leave_block("Checking signatures");

    enter_block("Checking pairings");
    // To Do: Decide whether to move pak and lambda preprocessing to offline
    std::vector<G2_precomp<snark_pp<ppT>>> g_Lambdas_precomp;
    g_Lambdas_precomp.reserve(auth_data.size());
    for(size_t i=0; i < auth_data.size(); i++)
        g_Lambdas_precomp.emplace_back(snark_pp<ppT>::precompute_G2(auth_data[i].Lambda));
    G2_precomp<snark_pp<ppT>> g_minusi_precomp = snark_pp<ppT>::precompute_G2(pak.minusI2);

    enter_block("Computation");
    Fqk<snark_pp<ppT>> accum;
    if(auth_data.size() % 2 == 1) {
        accum = snark_pp<ppT>::miller_loop(pvk.proof_g_vki_precomp[0]  , g_Lambdas_precomp[0]);
    }
    else {
        accum = Fqk<snark_pp<ppT>>::one();
    }
    for(size_t i = auth_data.size() % 2; i < labels.size();i=i+2) {
        accum = accum * snark_pp<ppT>::double_miller_loop(pvk.proof_g_vki_precomp[i]  , g_Lambdas_precomp[i],
                                                          pvk.proof_g_vki_precomp[i+1], g_Lambdas_precomp[i+1]);
    }
    G1_precomp<snark_pp<ppT>> proof_g_muA_precomp = snark_pp<ppT>::precompute_G1(proof.muA);
    G1_precomp<snark_pp<ppT>> proof_g_Aau_precomp = snark_pp<ppT>::precompute_G1(proof.g_Aau.g);
    Fqk<snark_pp<ppT>> accum2 = snark_pp<ppT>::double_miller_loop(proof_g_muA_precomp, pvk.pp_G2_one_precomp,
                                                                  proof_g_Aau_precomp, g_minusi_precomp);
    GT<snark_pp<ppT>> authPair = snark_pp<ppT>::final_exponentiation(accum * accum2.unitary_inverse());
    if (authPair != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Auth pairing check failed.\n");
        }
        result_auth = false;
    }
    leave_block("Computation");
    leave_block("Checking pairings");


    if (!(result_auth)) {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Authentication check failed.\n");
        }
    }

    leave_block("Checking A1");

    enter_block("Checking A2");
    G1_precomp<snark_pp<ppT>> proof_g_Aau_g_precomp = snark_pp<ppT>::precompute_G1(proof.g_Aau.g);
    G1_precomp<snark_pp<ppT>> proof_g_Aau_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_Aau.h);
    Fqk<snark_pp<ppT>> kc_Aau_1 = snark_pp<ppT>::miller_loop(proof_g_Aau_g_precomp, pvk.vk_alphaA_g2_precomp);
    Fqk<snark_pp<ppT>> kc_Aau_2 = snark_pp<ppT>::miller_loop(proof_g_Aau_h_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_Aau = snark_pp<ppT>::final_exponentiation(kc_Aau_1 * kc_Aau_2.unitary_inverse());
    if (kc_Aau != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for Aau query incorrect.\n");
        }
        result_auth = false;
    }
    leave_block("Checking A2");

    leave_block("Checking auth-specific elements");

    result &= result_auth;

    enter_block("Online pairing computations");
    enter_block("Check knowledge commitment for A is valid");
    G1_precomp<snark_pp<ppT>> proof_g_A_g_precomp      = snark_pp<ppT>::precompute_G1(proof.g_A.g);
    G1_precomp<snark_pp<ppT>> proof_g_A_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_A.h);
    Fqk<snark_pp<ppT>> kc_A_1 = snark_pp<ppT>::miller_loop(proof_g_A_g_precomp,      pvk.vk_alphaA_g2_precomp);
    Fqk<snark_pp<ppT>> kc_A_2 = snark_pp<ppT>::miller_loop(proof_g_A_h_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_A = snark_pp<ppT>::final_exponentiation(kc_A_1 * kc_A_2.unitary_inverse());
    if (kc_A != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for A query incorrect.\n");
        }
        result = false;
    }
    leave_block("Check knowledge commitment for A is valid");

    enter_block("Check knowledge commitment for B is valid");
    G2_precomp<snark_pp<ppT>> proof_g_B_g_precomp      = snark_pp<ppT>::precompute_G2(proof.g_B.g);
    G1_precomp<snark_pp<ppT>> proof_g_B_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_B.h);
    Fqk<snark_pp<ppT>> kc_B_1 = snark_pp<ppT>::miller_loop(pvk.vk_alphaB_g1_precomp, proof_g_B_g_precomp);
    Fqk<snark_pp<ppT>> kc_B_2 = snark_pp<ppT>::miller_loop(proof_g_B_h_precomp,    pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_B = snark_pp<ppT>::final_exponentiation(kc_B_1 * kc_B_2.unitary_inverse());
    if (kc_B != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for B query incorrect.\n");
        }
        result = false;
    }
    leave_block("Check knowledge commitment for B is valid");

    enter_block("Check knowledge commitment for C is valid");
    G1_precomp<snark_pp<ppT>> proof_g_C_g_precomp      = snark_pp<ppT>::precompute_G1(proof.g_C.g);
    G1_precomp<snark_pp<ppT>> proof_g_C_h_precomp = snark_pp<ppT>::precompute_G1(proof.g_C.h);
    Fqk<snark_pp<ppT>> kc_C_1 = snark_pp<ppT>::miller_loop(proof_g_C_g_precomp,      pvk.vk_alphaC_g2_precomp);
    Fqk<snark_pp<ppT>> kc_C_2 = snark_pp<ppT>::miller_loop(proof_g_C_h_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> kc_C = snark_pp<ppT>::final_exponentiation(kc_C_1 * kc_C_2.unitary_inverse());
    if (kc_C != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Knowledge commitment for C query incorrect.\n");
        }
        result = false;
    }
    leave_block("Check knowledge commitment for C is valid");

    G1<snark_pp<ppT>> Aacc = pvk.A0 + proof.g_Aau.g + proof.g_A.g;

    enter_block("Check QAP divisibility");
    G1_precomp<snark_pp<ppT>> proof_g_Aacc_precomp = snark_pp<ppT>::precompute_G1(Aacc);
    G1_precomp<snark_pp<ppT>> proof_g_H_precomp = snark_pp<ppT>::precompute_G1(proof.g_H);
    Fqk<snark_pp<ppT>> QAP_1  = snark_pp<ppT>::miller_loop(proof_g_Aacc_precomp,  proof_g_B_g_precomp);
    Fqk<snark_pp<ppT>> QAP_23  = snark_pp<ppT>::double_miller_loop(proof_g_H_precomp, pvk.vk_rC_Z_g2_precomp,
                                                                   proof_g_C_g_precomp, pvk.pp_G2_one_precomp);
    GT<snark_pp<ppT>> QAP = snark_pp<ppT>::final_exponentiation(QAP_1 * QAP_23.unitary_inverse());
    if (QAP != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("QAP divisibility check failed.\n");
        }
        result = false;
    }
    leave_block("Check QAP divisibility");

    enter_block("Check same coefficients were used");
    G1_precomp<snark_pp<ppT>> proof_g_K_precomp = snark_pp<ppT>::precompute_G1(proof.g_K);
    G1_precomp<snark_pp<ppT>> proof_g_Aacc_C_precomp = snark_pp<ppT>::precompute_G1(Aacc + proof.g_C.g);
    Fqk<snark_pp<ppT>> K_1 = snark_pp<ppT>::miller_loop(proof_g_K_precomp, pvk.vk_gamma_g2_precomp);
    Fqk<snark_pp<ppT>> K_23 = snark_pp<ppT>::double_miller_loop(proof_g_Aacc_C_precomp, pvk.vk_gamma_beta_g2_precomp,
                                                                pvk.vk_gamma_beta_g1_precomp, proof_g_B_g_precomp);
    GT<snark_pp<ppT>> K = snark_pp<ppT>::final_exponentiation(K_1 * K_23.unitary_inverse());
    if (K != GT<snark_pp<ppT>>::one())
    {
        if (!inhibit_profiling_info)
        {
            print_indent(); printf("Same-coefficient check failed.\n");
        }
        result = false;
    }
    leave_block("Check same coefficients were used");
    leave_block("Online pairing computations");
    leave_block("Call to r1cs_ppzkadsnark_online_verifier");

    return result;
}

// public
template<typename ppT>
bool r1cs_ppzkadsnark_verifier(const r1cs_ppzkadsnark_verification_key<ppT> &vk,
                               const std::vector<r1cs_ppzkadsnark_auth_data<ppT>> &auth_data,
                               const r1cs_ppzkadsnark_proof<ppT> &proof,
                               const r1cs_ppzkadsnark_pub_auth_key<ppT> &pak,
                               const std::vector<labelT> &labels)
{
    assert(labels.size() == auth_data.size());
    enter_block("Call to r1cs_ppzkadsnark_verifier");
    r1cs_ppzkadsnark_processed_verification_key<ppT> pvk = r1cs_ppzkadsnark_verifier_process_vk<ppT>(vk);
    bool result = r1cs_ppzkadsnark_online_verifier<ppT>(pvk, auth_data, proof, pak,labels);
    leave_block("Call to r1cs_ppzkadsnark_verifier");
    return result;
}

} // libsnark
#endif // R1CS_PPZKADSNARK_TCC_

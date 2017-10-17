/** @file
 *****************************************************************************

 Declaration of interfaces for a *single-predicate* ppzkPCD for R1CS.

 This includes:
 - class for proving key
 - class for verification key
 - class for processed verification key
 - class for key pair (proving key & verification key)
 - class for proof
 - generator algorithm
 - prover algorithm
 - verifier algorithm
 - online verifier algorithm

 The implementation follows, extends, and optimizes the approach described
 in \[BCTV14]. Thus, PCD is constructed from two "matched" ppzkSNARKs for R1CS.

 Acronyms:

 "R1CS" = "Rank-1 Constraint Systems"
 "ppzkSNARK" = "PreProcessing Zero-Knowledge Succinct Non-interactive ARgument of Knowledge"
 "ppzkPCD" = "Pre-Processing Zero-Knowledge Proof-Carrying Data"

 References:

 \[BCTV14]:
 "Scalable Zero Knowledge via Cycles of Elliptic Curves",
 Eli Ben-Sasson, Alessandro Chiesa, Eran Tromer, Madars Virza,
 CRYPTO 2014,
 <http://eprint.iacr.org/2014/595>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_SP_PPZKPCD_HPP_
#define R1CS_SP_PPZKPCD_HPP_

#include <memory>

#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/r1cs_sp_ppzkpcd_params.hpp"

namespace libsnark {

/******************************** Proving key ********************************/

template<typename PCD_ppT>
class r1cs_sp_ppzkpcd_proving_key;

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk);

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk);

/**
 * A proving key for the R1CS (single-predicate) ppzkPCD.
 */
template<typename PCD_ppT>
class r1cs_sp_ppzkpcd_proving_key {
public:
    typedef typename PCD_ppT::curve_A_pp A_pp;
    typedef typename PCD_ppT::curve_B_pp B_pp;

    r1cs_sp_ppzkpcd_compliance_predicate<PCD_ppT> compliance_predicate;

    r1cs_ppzksnark_proving_key<A_pp> compliance_step_r1cs_pk;
    r1cs_ppzksnark_proving_key<B_pp> translation_step_r1cs_pk;

    r1cs_ppzksnark_verification_key<A_pp> compliance_step_r1cs_vk;
    r1cs_ppzksnark_verification_key<B_pp> translation_step_r1cs_vk;

    r1cs_sp_ppzkpcd_proving_key() {};
    r1cs_sp_ppzkpcd_proving_key(const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &other) = default;
    r1cs_sp_ppzkpcd_proving_key(r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &&other) = default;
    r1cs_sp_ppzkpcd_proving_key(const r1cs_sp_ppzkpcd_compliance_predicate<PCD_ppT> &compliance_predicate,
                                r1cs_ppzksnark_proving_key<A_pp> &&compliance_step_r1cs_pk,
                                r1cs_ppzksnark_proving_key<B_pp> &&translation_step_r1cs_pk,
                                const r1cs_ppzksnark_verification_key<A_pp> &compliance_step_r1cs_vk,
                                const r1cs_ppzksnark_verification_key<B_pp> &translation_step_r1cs_vk) :
        compliance_predicate(compliance_predicate),
        compliance_step_r1cs_pk(std::move(compliance_step_r1cs_pk)),
        translation_step_r1cs_pk(std::move(translation_step_r1cs_pk)),
        compliance_step_r1cs_vk(std::move(compliance_step_r1cs_vk)),
        translation_step_r1cs_vk(std::move(translation_step_r1cs_vk))
    {};

    r1cs_sp_ppzkpcd_proving_key<PCD_ppT>& operator=(const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &other) = default;

    size_t size_in_bits() const
    {
        return (compliance_step_r1cs_pk.size_in_bits()
                + translation_step_r1cs_pk.size_in_bits()
                + compliance_step_r1cs_vk.size_in_bits()
                + translation_step_r1cs_vk.size_in_bits());
    }

    bool operator==(const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &other) const;
    friend std::ostream& operator<< <PCD_ppT>(std::ostream &out, const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk);
    friend std::istream& operator>> <PCD_ppT>(std::istream &in, r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk);
};


/******************************* Verification key ****************************/

template<typename PCD_ppT>
class r1cs_sp_ppzkpcd_verification_key;

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk);

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk);

/**
 * A verification key for the R1CS (single-predicate) ppzkPCD.
 */
template<typename PCD_ppT>
class r1cs_sp_ppzkpcd_verification_key {
public:
    typedef typename PCD_ppT::curve_A_pp A_pp;
    typedef typename PCD_ppT::curve_B_pp B_pp;

    r1cs_ppzksnark_verification_key<A_pp> compliance_step_r1cs_vk;
    r1cs_ppzksnark_verification_key<B_pp> translation_step_r1cs_vk;

    r1cs_sp_ppzkpcd_verification_key() = default;
    r1cs_sp_ppzkpcd_verification_key(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &other) = default;
    r1cs_sp_ppzkpcd_verification_key(r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &&other) = default;
    r1cs_sp_ppzkpcd_verification_key(const r1cs_ppzksnark_verification_key<A_pp> &compliance_step_r1cs_vk,
                                     const r1cs_ppzksnark_verification_key<B_pp> &translation_step_r1cs_vk) :
        compliance_step_r1cs_vk(std::move(compliance_step_r1cs_vk)),
        translation_step_r1cs_vk(std::move(translation_step_r1cs_vk))
    {};

    r1cs_sp_ppzkpcd_verification_key<PCD_ppT>& operator=(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &other) = default;

    size_t size_in_bits() const
    {
        return (compliance_step_r1cs_vk.size_in_bits()
                + translation_step_r1cs_vk.size_in_bits());
    }

    bool operator==(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &other) const;
    friend std::ostream& operator<< <PCD_ppT>(std::ostream &out, const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk);
    friend std::istream& operator>> <PCD_ppT>(std::istream &in, r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk);

    static r1cs_sp_ppzkpcd_verification_key<PCD_ppT> dummy_verification_key();
};


/************************ Processed verification key *************************/

template<typename PCD_ppT>
class r1cs_sp_ppzkpcd_processed_verification_key;

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);

/**
 * A processed verification key for the R1CS (single-predicate) ppzkPCD.
 *
 * Compared to a (non-processed) verification key, a processed verification key
 * contains a small constant amount of additional pre-computed information that
 * enables a faster verification time.
 */
template<typename PCD_ppT>
class r1cs_sp_ppzkpcd_processed_verification_key {
public:
    typedef typename PCD_ppT::curve_A_pp A_pp;
    typedef typename PCD_ppT::curve_B_pp B_pp;

    r1cs_ppzksnark_processed_verification_key<A_pp> compliance_step_r1cs_pvk;
    r1cs_ppzksnark_processed_verification_key<B_pp> translation_step_r1cs_pvk;
    bit_vector translation_step_r1cs_vk_bits;

    r1cs_sp_ppzkpcd_processed_verification_key() {};
    r1cs_sp_ppzkpcd_processed_verification_key(const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &other) = default;
    r1cs_sp_ppzkpcd_processed_verification_key(r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &&other) = default;
    r1cs_sp_ppzkpcd_processed_verification_key(r1cs_ppzksnark_processed_verification_key<A_pp> &&compliance_step_r1cs_pvk,
                                               r1cs_ppzksnark_processed_verification_key<B_pp> &&translation_step_r1cs_pvk,
                                               const bit_vector &translation_step_r1cs_vk_bits) :
        compliance_step_r1cs_pvk(std::move(compliance_step_r1cs_pvk)),
        translation_step_r1cs_pvk(std::move(translation_step_r1cs_pvk)),
        translation_step_r1cs_vk_bits(std::move(translation_step_r1cs_vk_bits))
    {};

    r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT>& operator=(const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &other) = default;

    size_t size_in_bits() const
    {
        return (compliance_step_r1cs_pvk.size_in_bits() +
                translation_step_r1cs_pvk.size_in_bits() +
                translation_step_r1cs_vk_bits.size());
    }

    bool operator==(const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &other) const;
    friend std::ostream& operator<< <PCD_ppT>(std::ostream &out, const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);
    friend std::istream& operator>> <PCD_ppT>(std::istream &in, r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);
};


/********************************* Key pair **********************************/

/**
 * A key pair for the R1CS (single-predicate) ppzkPC, which consists of a proving key and a verification key.
 */
template<typename PCD_ppT>
class r1cs_sp_ppzkpcd_keypair {
public:
    typedef typename PCD_ppT::curve_A_pp A_pp;
    typedef typename PCD_ppT::curve_B_pp B_pp;

    r1cs_sp_ppzkpcd_proving_key<PCD_ppT> pk;
    r1cs_sp_ppzkpcd_verification_key<PCD_ppT> vk;

    r1cs_sp_ppzkpcd_keypair() {};
    r1cs_sp_ppzkpcd_keypair(r1cs_sp_ppzkpcd_keypair<PCD_ppT> &&other) = default;
    r1cs_sp_ppzkpcd_keypair(r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &&pk,
                            r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &&vk) :
        pk(std::move(pk)),
        vk(std::move(vk))
    {};
    r1cs_sp_ppzkpcd_keypair(r1cs_ppzksnark_keypair<A_pp> &&kp_A,
                            r1cs_ppzksnark_keypair<B_pp> &&kp_B) :
        pk(std::move(kp_A.pk),std::move(kp_B.pk)),
        vk(std::move(kp_A.vk),std::move(kp_B.vk))
    {};
};


/*********************************** Proof ***********************************/

/**
 * A proof for the R1CS (single-predicate) ppzkPCD.
 */
template<typename PCD_ppT>
using r1cs_sp_ppzkpcd_proof = r1cs_ppzksnark_proof<typename PCD_ppT::curve_B_pp>;


/***************************** Main algorithms *******************************/

/**
 * A generator algorithm for the R1CS (single-predicate) ppzkPCD.
 *
 * Given a compliance predicate, this algorithm produces proving and verification keys for the predicate.
 */
template<typename PCD_ppT>
r1cs_sp_ppzkpcd_keypair<PCD_ppT> r1cs_sp_ppzkpcd_generator(const r1cs_sp_ppzkpcd_compliance_predicate<PCD_ppT> &compliance_predicate);

/**
 * A prover algorithm for the R1CS (single-predicate) ppzkPCD.
 *
 * Given a proving key, inputs for the compliance predicate, and proofs for
 * the predicate's input messages, this algorithm produces a proof (of knowledge)
 * that attests to the compliance of the output message.
 */
template <typename PCD_ppT>
r1cs_sp_ppzkpcd_proof<PCD_ppT> r1cs_sp_ppzkpcd_prover(const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk,
                                                      const r1cs_sp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                                                      const r1cs_sp_ppzkpcd_auxiliary_input<PCD_ppT> &auxiliary_input,
                                                      const std::vector<r1cs_sp_ppzkpcd_proof<PCD_ppT> > &incoming_proofs);

/*
 Below are two variants of verifier algorithm for the R1CS (single-predicate) ppzkPCD.

 These are the two cases that arise from whether the verifier accepts a
 (non-processed) verification key or, instead, a processed verification key.
 In the latter case, we call the algorithm an "online verifier".
 */

/**
 * A verifier algorithm for the R1CS (single-predicate) ppzkPCD that
 * accepts a non-processed verification key.
 */
template<typename PCD_ppT>
bool r1cs_sp_ppzkpcd_verifier(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk,
                              const r1cs_sp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                              const r1cs_sp_ppzkpcd_proof<PCD_ppT> &proof);

/**
 * Convert a (non-processed) verification key into a processed verification key.
 */
template<typename PCD_ppT>
r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> r1cs_sp_ppzkpcd_process_vk(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk);

/**
 * A verifier algorithm for the R1CS (single-predicate) ppzkPCD that
 * accepts a processed verification key.
 */
template<typename PCD_ppT>
bool r1cs_sp_ppzkpcd_online_verifier(const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk,
                                     const r1cs_sp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                                     const r1cs_sp_ppzkpcd_proof<PCD_ppT> &proof);

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/r1cs_sp_ppzkpcd.tcc"

#endif // R1CS_SP_PPZKPCD_HPP_

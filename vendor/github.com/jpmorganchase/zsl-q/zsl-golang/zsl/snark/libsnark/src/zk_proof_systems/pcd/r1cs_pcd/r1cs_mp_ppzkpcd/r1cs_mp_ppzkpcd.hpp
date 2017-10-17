/** @file
 *****************************************************************************

 Declaration of interfaces for a *multi-predicate* ppzkPCD for R1CS.

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
 in \[CTV15]. Thus, PCD is constructed from two "matched" ppzkSNARKs for R1CS.

 Acronyms:

 "R1CS" = "Rank-1 Constraint Systems"
 "ppzkSNARK" = "PreProcessing Zero-Knowledge Succinct Non-interactive ARgument of Knowledge"
 "ppzkPCD" = "Pre-Processing Zero-Knowledge Proof-Carrying Data"

 References:

 \[CTV15]:
 "Cluster Computing in Zero Knowledge",
 Alessandro Chiesa, Eran Tromer, Madars Virza,

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_MP_PPZKPCD_HPP_
#define R1CS_MP_PPZKPCD_HPP_

#include <memory>
#include <vector>

#include "common/data_structures/set_commitment.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/ppzkpcd_compliance_predicate.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_mp_ppzkpcd/r1cs_mp_ppzkpcd_params.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp"

namespace libsnark {

/******************************** Proving key ********************************/

template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_proving_key;

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &pk);

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &pk);

/**
 * A proving key for the R1CS (multi-predicate) ppzkPCD.
 */
template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_proving_key {
public:
    typedef typename PCD_ppT::curve_A_pp A_pp;
    typedef typename PCD_ppT::curve_B_pp B_pp;

    std::vector<r1cs_mp_ppzkpcd_compliance_predicate<PCD_ppT> > compliance_predicates;

    std::vector<r1cs_ppzksnark_proving_key<A_pp> > compliance_step_r1cs_pks;
    std::vector<r1cs_ppzksnark_proving_key<B_pp> > translation_step_r1cs_pks;

    std::vector<r1cs_ppzksnark_verification_key<A_pp> > compliance_step_r1cs_vks;
    std::vector<r1cs_ppzksnark_verification_key<B_pp> > translation_step_r1cs_vks;

    set_commitment commitment_to_translation_step_r1cs_vks;
    std::vector<set_membership_proof> compliance_step_r1cs_vk_membership_proofs;

    std::map<size_t, size_t> compliance_predicate_name_to_idx;

    r1cs_mp_ppzkpcd_proving_key() {};
    r1cs_mp_ppzkpcd_proving_key(const r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &other) = default;
    r1cs_mp_ppzkpcd_proving_key(r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &&other) = default;
    r1cs_mp_ppzkpcd_proving_key(const std::vector<r1cs_mp_ppzkpcd_compliance_predicate<PCD_ppT> > &compliance_predicates,
                                const std::vector<r1cs_ppzksnark_proving_key<A_pp> > &compliance_step_r1cs_pk,
                                const std::vector<r1cs_ppzksnark_proving_key<B_pp> > &translation_step_r1cs_pk,
                                const std::vector<r1cs_ppzksnark_verification_key<A_pp> > &compliance_step_r1cs_vk,
                                const std::vector<r1cs_ppzksnark_verification_key<B_pp> > &translation_step_r1cs_vk,
                                const set_commitment &commitment_to_translation_step_r1cs_vks,
                                const std::vector<set_membership_proof> &compliance_step_r1cs_vk_membership_proofs,
                                const std::map<size_t, size_t> &compliance_predicate_name_to_idx) :
    compliance_predicates(compliance_predicates),
        compliance_step_r1cs_pks(compliance_step_r1cs_pks),
        translation_step_r1cs_pks(translation_step_r1cs_pks),
        compliance_step_r1cs_vks(compliance_step_r1cs_vks),
        translation_step_r1cs_vks(translation_step_r1cs_vks),
        commitment_to_translation_step_r1cs_vks(commitment_to_translation_step_r1cs_vks),
        compliance_step_r1cs_vk_membership_proofs(compliance_step_r1cs_vk_membership_proofs),
        compliance_predicate_name_to_idx(compliance_predicate_name_to_idx)
    {}

    r1cs_mp_ppzkpcd_proving_key<PCD_ppT>& operator=(const r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &other) = default;

    size_t size_in_bits() const;

    bool is_well_formed() const;

    bool operator==(const r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &other) const;
    friend std::ostream& operator<< <PCD_ppT>(std::ostream &out, const r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &pk);
    friend std::istream& operator>> <PCD_ppT>(std::istream &in, r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &pk);
};


/******************************* Verification key ****************************/

template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_verification_key;

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &vk);

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &vk);

/**
 * A verification key for the R1CS (multi-predicate) ppzkPCD.
 */
template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_verification_key {
public:
    typedef typename PCD_ppT::curve_A_pp A_pp;
    typedef typename PCD_ppT::curve_B_pp B_pp;

    std::vector<r1cs_ppzksnark_verification_key<A_pp> > compliance_step_r1cs_vks;
    std::vector<r1cs_ppzksnark_verification_key<B_pp> > translation_step_r1cs_vks;
    set_commitment commitment_to_translation_step_r1cs_vks;

    r1cs_mp_ppzkpcd_verification_key() = default;
    r1cs_mp_ppzkpcd_verification_key(const r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &other) = default;
    r1cs_mp_ppzkpcd_verification_key(r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &&other) = default;
    r1cs_mp_ppzkpcd_verification_key(const std::vector<r1cs_ppzksnark_verification_key<A_pp> > &compliance_step_r1cs_vks,
                                     const std::vector<r1cs_ppzksnark_verification_key<B_pp> > &translation_step_r1cs_vks,
                                     const set_commitment &commitment_to_translation_step_r1cs_vks) :
        compliance_step_r1cs_vks(compliance_step_r1cs_vks),
        translation_step_r1cs_vks(translation_step_r1cs_vks),
        commitment_to_translation_step_r1cs_vks(commitment_to_translation_step_r1cs_vks)
    {}

    r1cs_mp_ppzkpcd_verification_key<PCD_ppT>& operator=(const r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &other) = default;

    size_t size_in_bits() const;

    bool operator==(const r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &other) const;
    friend std::ostream& operator<< <PCD_ppT>(std::ostream &out, const r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &vk);
    friend std::istream& operator>> <PCD_ppT>(std::istream &in, r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &vk);
};


/************************* Processed verification key **************************/

template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_processed_verification_key;

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);

/**
 * A processed verification key for the R1CS (multi-predicate) ppzkPCD.
 *
 * Compared to a (non-processed) verification key, a processed verification key
 * contains a small constant amount of additional pre-computed information that
 * enables a faster verification time.
 */
template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_processed_verification_key {
public:
    typedef typename PCD_ppT::curve_A_pp A_pp;
    typedef typename PCD_ppT::curve_B_pp B_pp;

    std::vector<r1cs_ppzksnark_processed_verification_key<A_pp> > compliance_step_r1cs_pvks;
    std::vector<r1cs_ppzksnark_processed_verification_key<B_pp> > translation_step_r1cs_pvks;
    set_commitment commitment_to_translation_step_r1cs_vks;

    r1cs_mp_ppzkpcd_processed_verification_key() = default;
    r1cs_mp_ppzkpcd_processed_verification_key(const r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &other) = default;
    r1cs_mp_ppzkpcd_processed_verification_key(r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &&other) = default;
    r1cs_mp_ppzkpcd_processed_verification_key(std::vector<r1cs_ppzksnark_processed_verification_key<A_pp> > &&compliance_step_r1cs_pvks,
                                               std::vector<r1cs_ppzksnark_processed_verification_key<B_pp> > &&translation_step_r1cs_pvks,
                                               const set_commitment &commitment_to_translation_step_r1cs_vks) :
        compliance_step_r1cs_pvks(std::move(compliance_step_r1cs_pvks)),
        translation_step_r1cs_pvks(std::move(translation_step_r1cs_pvks)),
        commitment_to_translation_step_r1cs_vks(commitment_to_translation_step_r1cs_vks)
    {};

    r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT>& operator=(const r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &other) = default;

    size_t size_in_bits() const;

    bool operator==(const r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &other) const;
    friend std::ostream& operator<< <PCD_ppT>(std::ostream &out, const r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);
    friend std::istream& operator>> <PCD_ppT>(std::istream &in, r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk);
};


/********************************** Key pair *********************************/

/**
 * A key pair for the R1CS (multi-predicate) ppzkPC, which consists of a proving key and a verification key.
 */
template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_keypair {
public:
    r1cs_mp_ppzkpcd_proving_key<PCD_ppT> pk;
    r1cs_mp_ppzkpcd_verification_key<PCD_ppT> vk;

    r1cs_mp_ppzkpcd_keypair() = default;
    r1cs_mp_ppzkpcd_keypair(r1cs_mp_ppzkpcd_keypair<PCD_ppT> &&other) = default;
    r1cs_mp_ppzkpcd_keypair(r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &&pk,
                            r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &&vk) :
        pk(std::move(pk)),
        vk(std::move(vk))
    {};
};


/*********************************** Proof ***********************************/

template<typename ppT>
class r1cs_mp_ppzkpcd_proof;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_mp_ppzkpcd_proof<ppT> &proof);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_mp_ppzkpcd_proof<ppT> &proof);

/**
 * A proof for the R1CS (multi-predicate) ppzkPCD.
 */
template<typename PCD_ppT>
class r1cs_mp_ppzkpcd_proof {
public:
    size_t compliance_predicate_idx;
    r1cs_ppzksnark_proof<typename PCD_ppT::curve_B_pp> r1cs_proof;

    r1cs_mp_ppzkpcd_proof() = default;
    r1cs_mp_ppzkpcd_proof(const size_t compliance_predicate_idx,
                          const r1cs_ppzksnark_proof<typename PCD_ppT::curve_B_pp> &r1cs_proof) :
        compliance_predicate_idx(compliance_predicate_idx),
        r1cs_proof(r1cs_proof)
    {}

    size_t size_in_bits() const;

    bool operator==(const r1cs_mp_ppzkpcd_proof<PCD_ppT> &other) const;
    friend std::ostream& operator<< <PCD_ppT>(std::ostream &out, const r1cs_mp_ppzkpcd_proof<PCD_ppT> &proof);
    friend std::istream& operator>> <PCD_ppT>(std::istream &in, r1cs_mp_ppzkpcd_proof<PCD_ppT> &proof);
};


/***************************** Main algorithms *******************************/

/**
 * A generator algorithm for the R1CS (multi-predicate) ppzkPCD.
 *
 * Given a vector of compliance predicates, this algorithm produces proving and verification keys for the vector.
 */
template<typename PCD_ppT>
r1cs_mp_ppzkpcd_keypair<PCD_ppT> r1cs_mp_ppzkpcd_generator(const std::vector<r1cs_mp_ppzkpcd_compliance_predicate<PCD_ppT> > &compliance_predicates);

/**
 * A prover algorithm for the R1CS (multi-predicate) ppzkPCD.
 *
 * Given a proving key, name of chosen compliance predicate, inputs for the
 * compliance predicate, and proofs for the predicate's input messages, this
 * algorithm produces a proof (of knowledge) that attests to the compliance of
 * the output message.
 */
template <typename PCD_ppT>
r1cs_mp_ppzkpcd_proof<PCD_ppT> r1cs_mp_ppzkpcd_prover(const r1cs_mp_ppzkpcd_proving_key<PCD_ppT> &pk,
                                                      const size_t compliance_predicate_name,
                                                      const r1cs_mp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                                                      const r1cs_mp_ppzkpcd_auxiliary_input<PCD_ppT> &auxiliary_input,
                                                      const std::vector<r1cs_mp_ppzkpcd_proof<PCD_ppT> > &incoming_proofs);

/*
  Below are two variants of verifier algorithm for the R1CS (multi-predicate) ppzkPCD.

  These are the two cases that arise from whether the verifier accepts a
  (non-processed) verification key or, instead, a processed verification key.
  In the latter case, we call the algorithm an "online verifier".
*/

/**
 * A verifier algorithm for the R1CS (multi-predicate) ppzkPCD that
 * accepts a non-processed verification key.
 */
template<typename PCD_ppT>
bool r1cs_mp_ppzkpcd_verifier(const r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &vk,
                              const r1cs_mp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                              const r1cs_mp_ppzkpcd_proof<PCD_ppT> &proof);

/**
 * Convert a (non-processed) verification key into a processed verification key.
 */
template<typename PCD_ppT>
r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> r1cs_mp_ppzkpcd_process_vk(const r1cs_mp_ppzkpcd_verification_key<PCD_ppT> &vk);

/**
 * A verifier algorithm for the R1CS (multi-predicate) ppzkPCD that
 * accepts a processed verification key.
 */
template<typename PCD_ppT>
bool r1cs_mp_ppzkpcd_online_verifier(const r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk,
                                     const r1cs_mp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                                     const r1cs_mp_ppzkpcd_proof<PCD_ppT> &proof);

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_mp_ppzkpcd/r1cs_mp_ppzkpcd.tcc"

#endif // R1CS_MP_PPZKPCD_HPP_

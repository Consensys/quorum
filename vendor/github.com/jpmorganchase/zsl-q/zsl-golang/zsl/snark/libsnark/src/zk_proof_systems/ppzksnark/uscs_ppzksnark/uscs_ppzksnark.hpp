/** @file
 *****************************************************************************

 Declaration of interfaces for a ppzkSNARK for USCS.

 This includes:
 - class for proving key
 - class for verification key
 - class for processed verification key
 - class for key pair (proving key & verification key)
 - class for proof
 - generator algorithm
 - prover algorithm
 - verifier algorithm (with strong or weak input consistency)
 - online verifier algorithm (with strong or weak input consistency)

 The implementation instantiates the protocol of \[DFGK14], by following
 extending, and optimizing the approach described in \[BCTV14].


 Acronyms:

 - "ppzkSNARK" = "Pre-Processing Zero-Knowledge Succinct Non-interactive ARgument of Knowledge"
 - "USCS" = "Unitary-Square Constraint Systems"

 References:

 \[BCTV14]:
 "Succinct Non-Interactive Zero Knowledge for a von Neumann Architecture",
 Eli Ben-Sasson, Alessandro Chiesa, Eran Tromer, Madars Virza,
 USENIX Security 2014,
 <http://eprint.iacr.org/2013/879>

 \[DFGK14]:
 "Square Span Programs with Applications to Succinct NIZK Arguments"
 George Danezis, Cedric Fournet, Jens Groth, Markulf Kohlweiss,
 ASIACRYPT 2014,
 <http://eprint.iacr.org/2014/718>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef USCS_PPZKSNARK_HPP_
#define USCS_PPZKSNARK_HPP_

#include <memory>

#include "algebra/curves/public_params.hpp"
#include "common/data_structures/accumulation_vector.hpp"
#include "algebra/knowledge_commitment/knowledge_commitment.hpp"
#include "relations/constraint_satisfaction_problems/uscs/uscs.hpp"
#include "zk_proof_systems/ppzksnark/uscs_ppzksnark/uscs_ppzksnark_params.hpp"

namespace libsnark {

/******************************** Proving key ********************************/

template<typename ppT>
class uscs_ppzksnark_proving_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const uscs_ppzksnark_proving_key<ppT> &pk);

template<typename ppT>
std::istream& operator>>(std::istream &in, uscs_ppzksnark_proving_key<ppT> &pk);

/**
 * A proving key for the USCS ppzkSNARK.
 */
template<typename ppT>
class uscs_ppzksnark_proving_key {
public:
    G1_vector<ppT> V_g1_query;
    G1_vector<ppT> alpha_V_g1_query;
    G1_vector<ppT> H_g1_query;
    G2_vector<ppT> V_g2_query;

    uscs_ppzksnark_constraint_system<ppT> constraint_system;

    uscs_ppzksnark_proving_key() {};
    uscs_ppzksnark_proving_key<ppT>& operator=(const uscs_ppzksnark_proving_key<ppT> &other) = default;
    uscs_ppzksnark_proving_key(const uscs_ppzksnark_proving_key<ppT> &other) = default;
    uscs_ppzksnark_proving_key(uscs_ppzksnark_proving_key<ppT> &&other) = default;
    uscs_ppzksnark_proving_key(G1_vector<ppT> &&V_g1_query,
                               G1_vector<ppT> &&alpha_V_g1_query,
                               G1_vector<ppT> &&H_g1_query,
                               G2_vector<ppT> &&V_g2_query,
                               uscs_ppzksnark_constraint_system<ppT> &&constraint_system) :
        V_g1_query(std::move(V_g1_query)),
        alpha_V_g1_query(std::move(alpha_V_g1_query)),
        H_g1_query(std::move(H_g1_query)),
        V_g2_query(std::move(V_g2_query)),
        constraint_system(std::move(constraint_system))
    {};

    size_t G1_size() const
    {
        return V_g1_query.size() + alpha_V_g1_query.size() + H_g1_query.size();
    }

    size_t G2_size() const
    {
        return V_g2_query.size();
    }

    size_t G1_sparse_size() const
    {
        return G1_size();
    }

    size_t G2_sparse_size() const
    {
        return G2_size();
    }

    size_t size_in_bits() const
    {
        return G1<ppT>::size_in_bits() * G1_size() + G2<ppT>::size_in_bits() * G2_size();
    }

    void print_size() const
    {
        print_indent(); printf("* G1 elements in PK: %zu\n", this->G1_size());
        print_indent(); printf("* Non-zero G1 elements in PK: %zu\n", this->G1_sparse_size());
        print_indent(); printf("* G2 elements in PK: %zu\n", this->G2_size());
        print_indent(); printf("* Non-zero G2 elements in PK: %zu\n", this->G2_sparse_size());
        print_indent(); printf("* PK size in bits: %zu\n", this->size_in_bits());
    }

    bool operator==(const uscs_ppzksnark_proving_key<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const uscs_ppzksnark_proving_key<ppT> &pk);
    friend std::istream& operator>> <ppT>(std::istream &in, uscs_ppzksnark_proving_key<ppT> &pk);
};


/******************************* Verification key ****************************/

template<typename ppT>
class uscs_ppzksnark_verification_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const uscs_ppzksnark_verification_key<ppT> &vk);

template<typename ppT>
std::istream& operator>>(std::istream &in, uscs_ppzksnark_verification_key<ppT> &vk);

/**
 * A verification key for the USCS ppzkSNARK.
 */
template<typename ppT>
class uscs_ppzksnark_verification_key {
public:
    G2<ppT> tilde_g2;
    G2<ppT> alpha_tilde_g2;
    G2<ppT> Z_g2;

    accumulation_vector<G1<ppT> > encoded_IC_query;

    uscs_ppzksnark_verification_key() = default;
    uscs_ppzksnark_verification_key(const G2<ppT> &tilde_g2,
                                    const G2<ppT> &alpha_tilde_g2,
                                    const G2<ppT> &Z_g2,
                                    const accumulation_vector<G1<ppT> > &eIC) :
        tilde_g2(tilde_g2),
        alpha_tilde_g2(alpha_tilde_g2),
        Z_g2(Z_g2),
        encoded_IC_query(eIC)
    {};

    size_t G1_size() const
    {
        return encoded_IC_query.size();
    }

    size_t G2_size() const
    {
        return 3;
    }

    size_t size_in_bits() const
    {
        return encoded_IC_query.size_in_bits() + 3 * G2<ppT>::size_in_bits();
    }

    void print_size() const
    {
        print_indent(); printf("* G1 elements in VK: %zu\n", this->G1_size());
        print_indent(); printf("* G2 elements in VK: %zu\n", this->G2_size());
        print_indent(); printf("* VK size in bits: %zu\n", this->size_in_bits());
    }

    bool operator==(const uscs_ppzksnark_verification_key<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const uscs_ppzksnark_verification_key<ppT> &vk);
    friend std::istream& operator>> <ppT>(std::istream &in, uscs_ppzksnark_verification_key<ppT> &vk);

    static uscs_ppzksnark_verification_key<ppT> dummy_verification_key(const size_t input_size);
};


/************************ Processed verification key *************************/

template<typename ppT>
class uscs_ppzksnark_processed_verification_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const uscs_ppzksnark_processed_verification_key<ppT> &pvk);

template<typename ppT>
std::istream& operator>>(std::istream &in, uscs_ppzksnark_processed_verification_key<ppT> &pvk);

/**
 * A processed verification key for the USCS ppzkSNARK.
 *
 * Compared to a (non-processed) verification key, a processed verification key
 * contains a small constant amount of additional pre-computed information that
 * enables a faster verification time.
 */
template<typename ppT>
class uscs_ppzksnark_processed_verification_key {
public:
    G1_precomp<ppT> pp_G1_one_precomp;
    G2_precomp<ppT> pp_G2_one_precomp;
    G2_precomp<ppT> vk_tilde_g2_precomp;
    G2_precomp<ppT> vk_alpha_tilde_g2_precomp;
    G2_precomp<ppT> vk_Z_g2_precomp;
    GT<ppT> pairing_of_g1_and_g2;

    accumulation_vector<G1<ppT> > encoded_IC_query;

    bool operator==(const uscs_ppzksnark_processed_verification_key &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const uscs_ppzksnark_processed_verification_key<ppT> &pvk);
    friend std::istream& operator>> <ppT>(std::istream &in, uscs_ppzksnark_processed_verification_key<ppT> &pvk);
};


/********************************** Key pair *********************************/

/**
 * A key pair for the USCS ppzkSNARK, which consists of a proving key and a verification key.
 */
template<typename ppT>
class uscs_ppzksnark_keypair {
public:
    uscs_ppzksnark_proving_key<ppT> pk;
    uscs_ppzksnark_verification_key<ppT> vk;

    uscs_ppzksnark_keypair() {};
    uscs_ppzksnark_keypair(uscs_ppzksnark_proving_key<ppT> &&pk,
                           uscs_ppzksnark_verification_key<ppT> &&vk) :
        pk(std::move(pk)),
        vk(std::move(vk))
    {}

    uscs_ppzksnark_keypair(uscs_ppzksnark_keypair<ppT> &&other) = default;
};


/*********************************** Proof ***********************************/

template<typename ppT>
class uscs_ppzksnark_proof;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const uscs_ppzksnark_proof<ppT> &proof);

template<typename ppT>
std::istream& operator>>(std::istream &in, uscs_ppzksnark_proof<ppT> &proof);

/**
 * A proof for the USCS ppzkSNARK.
 *
 * While the proof has a structure, externally one merely opaquely produces,
 * seralizes/deserializes, and verifies proofs. We only expose some information
 * about the structure for statistics purposes.
 */
template<typename ppT>
class uscs_ppzksnark_proof {
public:
    G1<ppT> V_g1;
    G1<ppT> alpha_V_g1;
    G1<ppT> H_g1;
    G2<ppT> V_g2;

    uscs_ppzksnark_proof()
    {
        // invalid proof with valid curve points
        this->V_g1 = G1<ppT> ::one();
        this->alpha_V_g1 = G1<ppT> ::one();
        this->H_g1 = G1<ppT> ::one();
        this->V_g2 = G2<ppT> ::one();
    }
    uscs_ppzksnark_proof(G1<ppT> &&V_g1,
                         G1<ppT> &&alpha_V_g1,
                         G1<ppT> &&H_g1,
                         G2<ppT> &&V_g2) :
        V_g1(std::move(V_g1)),
        alpha_V_g1(std::move(alpha_V_g1)),
        H_g1(std::move(H_g1)),
        V_g2(std::move(V_g2))
    {};

    size_t G1_size() const
    {
        return 3;
    }

    size_t G2_size() const
    {
        return 1;
    }

    size_t size_in_bits() const
    {
        return G1_size() * G1<ppT>::size_in_bits() + G2_size() * G2<ppT>::size_in_bits();
    }

    void print_size() const
    {
        print_indent(); printf("* G1 elements in proof: %zu\n", this->G1_size());
        print_indent(); printf("* G2 elements in proof: %zu\n", this->G2_size());
        print_indent(); printf("* Proof size in bits: %zu\n", this->size_in_bits());
    }

    bool is_well_formed() const
    {
        return (V_g1.is_well_formed() &&
                alpha_V_g1.is_well_formed() &&
                H_g1.is_well_formed() &&
                V_g2.is_well_formed());
    }

    bool operator==(const uscs_ppzksnark_proof<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const uscs_ppzksnark_proof<ppT> &proof);
    friend std::istream& operator>> <ppT>(std::istream &in, uscs_ppzksnark_proof<ppT> &proof);
};


/***************************** Main algorithms *******************************/

/**
 * A generator algorithm for the USCS ppzkSNARK.
 *
 * Given a USCS constraint system CS, this algorithm produces proving and verification keys for CS.
 */
template<typename ppT>
uscs_ppzksnark_keypair<ppT> uscs_ppzksnark_generator(const uscs_ppzksnark_constraint_system<ppT> &cs);

/**
 * A prover algorithm for the USCS ppzkSNARK.
 *
 * Given a USCS primary input X and a USCS auxiliary input Y, this algorithm
 * produces a proof (of knowledge) that attests to the following statement:
 *               ``there exists Y such that CS(X,Y)=0''.
 * Above, CS is the USCS constraint system that was given as input to the generator algorithm.
 */
template<typename ppT>
uscs_ppzksnark_proof<ppT> uscs_ppzksnark_prover(const uscs_ppzksnark_proving_key<ppT> &pk,
                                                const uscs_ppzksnark_primary_input<ppT> &primary_input,
                                                const uscs_ppzksnark_auxiliary_input<ppT> &auxiliary_input);

/*
 Below are four variants of verifier algorithm for the USCS ppzkSNARK.

 These are the four cases that arise from the following two choices:

 (1) The verifier accepts a (non-processed) verification key or, instead, a processed verification key.
     In the latter case, we call the algorithm an "online verifier".

 (2) The verifier checks for "weak" input consistency or, instead, "strong" input consistency.
     Strong input consistency requires that |primary_input| = CS.num_inputs, whereas
     weak input consistency requires that |primary_input| <= CS.num_inputs (and
     the primary input is implicitly padded with zeros up to length CS.num_inputs).
 */

/**
 * A verifier algorithm for the USCS ppzkSNARK that:
 * (1) accepts a non-processed verification key, and
 * (2) has weak input consistency.
 */
template<typename ppT>
bool uscs_ppzksnark_verifier_weak_IC(const uscs_ppzksnark_verification_key<ppT> &vk,
                                     const uscs_ppzksnark_primary_input<ppT> &primary_input,
                                     const uscs_ppzksnark_proof<ppT> &proof);

/**
 * A verifier algorithm for the USCS ppzkSNARK that:
 * (1) accepts a non-processed verification key, and
 * (2) has strong input consistency.
 */
template<typename ppT>
bool uscs_ppzksnark_verifier_strong_IC(const uscs_ppzksnark_verification_key<ppT> &vk,
                                       const uscs_ppzksnark_primary_input<ppT> &primary_input,
                                       const uscs_ppzksnark_proof<ppT> &proof);

/**
 * Convert a (non-processed) verification key into a processed verification key.
 */
template<typename ppT>
uscs_ppzksnark_processed_verification_key<ppT> uscs_ppzksnark_verifier_process_vk(const uscs_ppzksnark_verification_key<ppT> &vk);

/**
 * A verifier algorithm for the USCS ppzkSNARK that:
 * (1) accepts a processed verification key, and
 * (2) has weak input consistency.
 */
template<typename ppT>
bool uscs_ppzksnark_online_verifier_weak_IC(const uscs_ppzksnark_processed_verification_key<ppT> &pvk,
                                            const uscs_ppzksnark_primary_input<ppT> &primary_input,
                                            const uscs_ppzksnark_proof<ppT> &proof);

/**
 * A verifier algorithm for the USCS ppzkSNARK that:
 * (1) accepts a processed verification key, and
 * (2) has strong input consistency.
 */
template<typename ppT>
bool uscs_ppzksnark_online_verifier_strong_IC(const uscs_ppzksnark_processed_verification_key<ppT> &pvk,
                                              const uscs_ppzksnark_primary_input<ppT> &primary_input,
                                              const uscs_ppzksnark_proof<ppT> &proof);

} // libsnark

#include "zk_proof_systems/ppzksnark/uscs_ppzksnark/uscs_ppzksnark.tcc"

#endif // USCS_PPZKSNARK_HPP_

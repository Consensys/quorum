/** @file
 *****************************************************************************

 Declaration of interfaces for a ppzkADSNARK for R1CS.

 This includes:
 - class for authentication key (public and symmetric)
 - class for authentication verification key (public and symmetric)
 - class for proving key
 - class for verification key
 - class for processed verification key
 - class for key tuple (authentication key & proving key & verification key)
 - class for authenticated data
 - class for proof
 - generator algorithm
 - authentication key generator algorithm
 - prover algorithm
 - verifier algorithm (public and symmetric)
 - online verifier algorithm (public and symmetric)

 The implementation instantiates the construction in \[BBFR15], which in turn
 is based on the r1cs_ppzkadsnark proof system.

 Acronyms:

 - R1CS = "Rank-1 Constraint Systems"
 - ppzkADSNARK = "PreProcessing Zero-Knowledge Succinct Non-interactive ARgument of Knowledge Over Authenticated Data"

 References:

\[BBFR15]
"ADSNARK: Nearly Practical and Privacy-Preserving Proofs on Authenticated Data",
Michael Backes, Manuel Barbosa, Dario Fiore, Raphael M. Reischuk,
IEEE Symposium on Security and Privacy 2015,
 <http://eprint.iacr.org/2014/617>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_PPZKADSNARK_HPP_
#define R1CS_PPZKADSNARK_HPP_

#include <memory>

#include "algebra/curves/public_params.hpp"
#include "common/data_structures/accumulation_vector.hpp"
#include "algebra/knowledge_commitment/knowledge_commitment.hpp"
#include "relations/constraint_satisfaction_problems/r1cs/r1cs.hpp"
#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark_params.hpp"
#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark_signature.hpp"
#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark_prf.hpp"

namespace libsnark {

/******************************** Public authentication parameters ********************************/

template<typename ppT>
class r1cs_ppzkadsnark_pub_auth_prms;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_pub_auth_prms<ppT> &pap);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_pub_auth_prms<ppT> &pap);

/**
 * Public authentication parameters for the R1CS ppzkADSNARK
 */
template<typename ppT>
class r1cs_ppzkadsnark_pub_auth_prms {
public:
    G1<snark_pp<ppT>> I1;

    r1cs_ppzkadsnark_pub_auth_prms() {};
    r1cs_ppzkadsnark_pub_auth_prms<ppT>& operator=(const r1cs_ppzkadsnark_pub_auth_prms<ppT> &other) = default;
    r1cs_ppzkadsnark_pub_auth_prms(const r1cs_ppzkadsnark_pub_auth_prms<ppT> &other) = default;
    r1cs_ppzkadsnark_pub_auth_prms(r1cs_ppzkadsnark_pub_auth_prms<ppT> &&other) = default;
    r1cs_ppzkadsnark_pub_auth_prms(G1<snark_pp<ppT>> &&I1) : I1(std::move(I1)) {};

    bool operator==(const r1cs_ppzkadsnark_pub_auth_prms<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_pub_auth_prms<ppT> &pap);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_pub_auth_prms<ppT> &pap);
};

/******************************** Secret authentication key ********************************/

template<typename ppT>
class r1cs_ppzkadsnark_sec_auth_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_sec_auth_key<ppT> &key);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_sec_auth_key<ppT> &key);

/**
 * Secret authentication key for the R1CS ppzkADSNARK
 */
template<typename ppT>
class r1cs_ppzkadsnark_sec_auth_key {
public:
    Fr<snark_pp<ppT>> i;
    r1cs_ppzkadsnark_skT<ppT>skp;
    r1cs_ppzkadsnark_prfKeyT<ppT>S;

    r1cs_ppzkadsnark_sec_auth_key() {};
    r1cs_ppzkadsnark_sec_auth_key<ppT>& operator=(const r1cs_ppzkadsnark_sec_auth_key<ppT> &other) = default;
    r1cs_ppzkadsnark_sec_auth_key(const r1cs_ppzkadsnark_sec_auth_key<ppT> &other) = default;
    r1cs_ppzkadsnark_sec_auth_key(r1cs_ppzkadsnark_sec_auth_key<ppT> &&other) = default;
    r1cs_ppzkadsnark_sec_auth_key(Fr<snark_pp<ppT>> &&i,
                                  r1cs_ppzkadsnark_skT<ppT>&&skp, r1cs_ppzkadsnark_prfKeyT<ppT>&&S) :
        i(std::move(i)),
        skp(std::move(skp)),
        S(std::move(S)) {};

    bool operator==(const r1cs_ppzkadsnark_sec_auth_key<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_sec_auth_key<ppT> &key);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_sec_auth_key<ppT> &key);
};

/******************************** Public authentication key ********************************/

template<typename ppT>
class r1cs_ppzkadsnark_pub_auth_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_pub_auth_key<ppT> &key);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_pub_auth_key<ppT> &key);

/**
 * Public authentication key for the R1CS ppzkADSNARK
 */
template<typename ppT>
class r1cs_ppzkadsnark_pub_auth_key {
public:
    G2<snark_pp<ppT>> minusI2;
    r1cs_ppzkadsnark_vkT<ppT>vkp;

    r1cs_ppzkadsnark_pub_auth_key() {};
    r1cs_ppzkadsnark_pub_auth_key<ppT>& operator=(const r1cs_ppzkadsnark_pub_auth_key<ppT> &other) = default;
    r1cs_ppzkadsnark_pub_auth_key(const r1cs_ppzkadsnark_pub_auth_key<ppT> &other) = default;
    r1cs_ppzkadsnark_pub_auth_key(r1cs_ppzkadsnark_pub_auth_key<ppT> &&other) = default;
    r1cs_ppzkadsnark_pub_auth_key(G2<snark_pp<ppT>> &&minusI2, r1cs_ppzkadsnark_vkT<ppT>&&vkp) :
        minusI2(std::move(minusI2)),
        vkp(std::move(vkp)) {};

    bool operator==(const r1cs_ppzkadsnark_pub_auth_key<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_pub_auth_key<ppT> &key);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_pub_auth_key<ppT> &key);
};

/******************************** Authentication key material ********************************/

template<typename ppT>
class r1cs_ppzkadsnark_auth_keys {
public:
    r1cs_ppzkadsnark_pub_auth_prms<ppT> pap;
    r1cs_ppzkadsnark_pub_auth_key<ppT> pak;
    r1cs_ppzkadsnark_sec_auth_key<ppT> sak;

    r1cs_ppzkadsnark_auth_keys() {};
    r1cs_ppzkadsnark_auth_keys(r1cs_ppzkadsnark_auth_keys<ppT> &&other) = default;
    r1cs_ppzkadsnark_auth_keys(r1cs_ppzkadsnark_pub_auth_prms<ppT> &&pap,
                               r1cs_ppzkadsnark_pub_auth_key<ppT> &&pak,
                               r1cs_ppzkadsnark_sec_auth_key<ppT> &&sak) :
        pap(std::move(pap)),
        pak(std::move(pak)),
        sak(std::move(sak))
    {}
};

/******************************** Authenticated data ********************************/

template<typename ppT>
class r1cs_ppzkadsnark_auth_data;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_auth_data<ppT> &data);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_auth_data<ppT> &data);

/**
 * Authenticated data for the R1CS ppzkADSNARK
 */
template<typename ppT>
class r1cs_ppzkadsnark_auth_data {
public:
    Fr<snark_pp<ppT>> mu;
    G2<snark_pp<ppT>> Lambda;
    r1cs_ppzkadsnark_sigT<ppT>sigma;

    r1cs_ppzkadsnark_auth_data() {};
    r1cs_ppzkadsnark_auth_data<ppT>& operator=(const r1cs_ppzkadsnark_auth_data<ppT> &other) = default;
    r1cs_ppzkadsnark_auth_data(const r1cs_ppzkadsnark_auth_data<ppT> &other) = default;
    r1cs_ppzkadsnark_auth_data(r1cs_ppzkadsnark_auth_data<ppT> &&other) = default;
    r1cs_ppzkadsnark_auth_data(Fr<snark_pp<ppT>> &&mu,
                               G2<snark_pp<ppT>> &&Lambda,
                               r1cs_ppzkadsnark_sigT<ppT>&&sigma) :
        mu(std::move(mu)),
        Lambda(std::move(Lambda)),
        sigma(std::move(sigma)) {};

    bool operator==(const r1cs_ppzkadsnark_auth_data<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_auth_data<ppT> &key);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_auth_data<ppT> &key);
};

/******************************** Proving key ********************************/

template<typename ppT>
class r1cs_ppzkadsnark_proving_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_proving_key<ppT> &pk);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_proving_key<ppT> &pk);

/**
 * A proving key for the R1CS ppzkADSNARK.
 */
template<typename ppT>
class r1cs_ppzkadsnark_proving_key {
public:
    knowledge_commitment_vector<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > A_query;
    knowledge_commitment_vector<G2<snark_pp<ppT>>, G1<snark_pp<ppT>> > B_query;
    knowledge_commitment_vector<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > C_query;
    G1_vector<snark_pp<ppT>> H_query; // t powers
    G1_vector<snark_pp<ppT>> K_query;
    /* Now come the additional elements for ad */
    G1<snark_pp<ppT>> rA_i_Z_g1;

    r1cs_ppzkadsnark_constraint_system<ppT> constraint_system;

    r1cs_ppzkadsnark_proving_key() {};
    r1cs_ppzkadsnark_proving_key<ppT>& operator=(const r1cs_ppzkadsnark_proving_key<ppT> &other) = default;
    r1cs_ppzkadsnark_proving_key(const r1cs_ppzkadsnark_proving_key<ppT> &other) = default;
    r1cs_ppzkadsnark_proving_key(r1cs_ppzkadsnark_proving_key<ppT> &&other) = default;
    r1cs_ppzkadsnark_proving_key(knowledge_commitment_vector<G1<snark_pp<ppT>>,
                                 G1<snark_pp<ppT>> > &&A_query,
                                 knowledge_commitment_vector<G2<snark_pp<ppT>>,
                                 G1<snark_pp<ppT>> > &&B_query,
                                 knowledge_commitment_vector<G1<snark_pp<ppT>>,
                                 G1<snark_pp<ppT>> > &&C_query,
                                 G1_vector<snark_pp<ppT>> &&H_query,
                                 G1_vector<snark_pp<ppT>> &&K_query,
                                 G1<snark_pp<ppT>> &&rA_i_Z_g1,
                                 r1cs_ppzkadsnark_constraint_system<ppT> &&constraint_system) :
        A_query(std::move(A_query)),
        B_query(std::move(B_query)),
        C_query(std::move(C_query)),
        H_query(std::move(H_query)),
        K_query(std::move(K_query)),
        rA_i_Z_g1(std::move(rA_i_Z_g1)),
        constraint_system(std::move(constraint_system))
    {};

    size_t G1_size() const
    {
        return 2*(A_query.domain_size() + C_query.domain_size()) + B_query.domain_size() + H_query.size() + K_query.size() + 1;
    }

    size_t G2_size() const
    {
        return B_query.domain_size();
    }

    size_t G1_sparse_size() const
    {
        return 2*(A_query.size() + C_query.size()) + B_query.size() + H_query.size() + K_query.size() + 1;
    }

    size_t G2_sparse_size() const
    {
        return B_query.size();
    }

    size_t size_in_bits() const
    {
        return A_query.size_in_bits() + B_query.size_in_bits() + C_query.size_in_bits() + libsnark::size_in_bits(H_query) + libsnark::size_in_bits(K_query) + G1<snark_pp<ppT>>::size_in_bits();
    }

    void print_size() const
    {
        print_indent(); printf("* G1 elements in PK: %zu\n", this->G1_size());
        print_indent(); printf("* Non-zero G1 elements in PK: %zu\n", this->G1_sparse_size());
        print_indent(); printf("* G2 elements in PK: %zu\n", this->G2_size());
        print_indent(); printf("* Non-zero G2 elements in PK: %zu\n", this->G2_sparse_size());
        print_indent(); printf("* PK size in bits: %zu\n", this->size_in_bits());
    }

    bool operator==(const r1cs_ppzkadsnark_proving_key<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_proving_key<ppT> &pk);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_proving_key<ppT> &pk);
};


/******************************* Verification key ****************************/

template<typename ppT>
class r1cs_ppzkadsnark_verification_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_verification_key<ppT> &vk);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_verification_key<ppT> &vk);

/**
 * A verification key for the R1CS ppzkADSNARK.
 */
template<typename ppT>
class r1cs_ppzkadsnark_verification_key {
public:
    G2<snark_pp<ppT>> alphaA_g2;
    G1<snark_pp<ppT>> alphaB_g1;
    G2<snark_pp<ppT>> alphaC_g2;
    G2<snark_pp<ppT>> gamma_g2;
    G1<snark_pp<ppT>> gamma_beta_g1;
    G2<snark_pp<ppT>> gamma_beta_g2;
    G2<snark_pp<ppT>> rC_Z_g2;

    G1<snark_pp<ppT>> A0;
    G1_vector<snark_pp<ppT>> Ain;

    r1cs_ppzkadsnark_verification_key() = default;
    r1cs_ppzkadsnark_verification_key(const G2<snark_pp<ppT>> &alphaA_g2,
                                      const G1<snark_pp<ppT>> &alphaB_g1,
                                      const G2<snark_pp<ppT>> &alphaC_g2,
                                      const G2<snark_pp<ppT>> &gamma_g2,
                                      const G1<snark_pp<ppT>> &gamma_beta_g1,
                                      const G2<snark_pp<ppT>> &gamma_beta_g2,
                                      const G2<snark_pp<ppT>> &rC_Z_g2,
                                      const G1<snark_pp<ppT>> A0,
                                      const G1_vector<snark_pp<ppT>> Ain) :
        alphaA_g2(alphaA_g2),
        alphaB_g1(alphaB_g1),
        alphaC_g2(alphaC_g2),
        gamma_g2(gamma_g2),
        gamma_beta_g1(gamma_beta_g1),
        gamma_beta_g2(gamma_beta_g2),
        rC_Z_g2(rC_Z_g2),
        A0(A0),
        Ain(Ain)
    {};

    size_t G1_size() const
    {
        return 3 + Ain.size();
    }

    size_t G2_size() const
    {
        return 5;
    }

    size_t size_in_bits() const
    {
        return G1_size() * G1<snark_pp<ppT>>::size_in_bits() + G2_size() * G2<snark_pp<ppT>>::size_in_bits(); // possible zksnark bug
    }

    void print_size() const
    {
        print_indent(); printf("* G1 elements in VK: %zu\n", this->G1_size());
        print_indent(); printf("* G2 elements in VK: %zu\n", this->G2_size());
        print_indent(); printf("* VK size in bits: %zu\n", this->size_in_bits());
    }

    bool operator==(const r1cs_ppzkadsnark_verification_key<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_verification_key<ppT> &vk);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_verification_key<ppT> &vk);

    static r1cs_ppzkadsnark_verification_key<ppT> dummy_verification_key(const size_t input_size);
};


/************************ Processed verification key *************************/

template<typename ppT>
class r1cs_ppzkadsnark_processed_verification_key;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk);

/**
 * A processed verification key for the R1CS ppzkADSNARK.
 *
 * Compared to a (non-processed) verification key, a processed verification key
 * contains a small constant amount of additional pre-computed information that
 * enables a faster verification time.
 */
template<typename ppT>
class r1cs_ppzkadsnark_processed_verification_key {
public:
    G2_precomp<snark_pp<ppT>> pp_G2_one_precomp;
    G2_precomp<snark_pp<ppT>> vk_alphaA_g2_precomp;
    G1_precomp<snark_pp<ppT>> vk_alphaB_g1_precomp;
    G2_precomp<snark_pp<ppT>> vk_alphaC_g2_precomp;
    G2_precomp<snark_pp<ppT>> vk_rC_Z_g2_precomp;
    G2_precomp<snark_pp<ppT>> vk_gamma_g2_precomp;
    G1_precomp<snark_pp<ppT>> vk_gamma_beta_g1_precomp;
    G2_precomp<snark_pp<ppT>> vk_gamma_beta_g2_precomp;
    G2_precomp<snark_pp<ppT>> vk_rC_i_g2_precomp;

    G1<snark_pp<ppT>> A0;
    G1_vector<snark_pp<ppT>> Ain;

    std::vector<G1_precomp<snark_pp<ppT>>> proof_g_vki_precomp;

    bool operator==(const r1cs_ppzkadsnark_processed_verification_key &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk);
};


/********************************** Key pair *********************************/

/**
 * A key pair for the R1CS ppzkADSNARK, which consists of a proving key and a verification key.
 */
template<typename ppT>
class r1cs_ppzkadsnark_keypair {
public:
    r1cs_ppzkadsnark_proving_key<ppT> pk;
    r1cs_ppzkadsnark_verification_key<ppT> vk;

    r1cs_ppzkadsnark_keypair() = default;
    r1cs_ppzkadsnark_keypair(const r1cs_ppzkadsnark_keypair<ppT> &other) = default;
    r1cs_ppzkadsnark_keypair(r1cs_ppzkadsnark_proving_key<ppT> &&pk,
                             r1cs_ppzkadsnark_verification_key<ppT> &&vk) :
        pk(std::move(pk)),
        vk(std::move(vk))
    {}

    r1cs_ppzkadsnark_keypair(r1cs_ppzkadsnark_keypair<ppT> &&other) = default;
};


/*********************************** Proof ***********************************/

template<typename ppT>
class r1cs_ppzkadsnark_proof;

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_ppzkadsnark_proof<ppT> &proof);

template<typename ppT>
std::istream& operator>>(std::istream &in, r1cs_ppzkadsnark_proof<ppT> &proof);

/**
 * A proof for the R1CS ppzkADSNARK.
 *
 * While the proof has a structure, externally one merely opaquely produces,
 * seralizes/deserializes, and verifies proofs. We only expose some information
 * about the structure for statistics purposes.
 */
template<typename ppT>
class r1cs_ppzkadsnark_proof {
public:
    knowledge_commitment<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_A;
    knowledge_commitment<G2<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_B;
    knowledge_commitment<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_C;
    G1<snark_pp<ppT>> g_H;
    G1<snark_pp<ppT>> g_K;
    knowledge_commitment<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> > g_Aau;
    G1<snark_pp<ppT>> muA;

    r1cs_ppzkadsnark_proof()
    {
        // invalid proof with valid curve points
        this->g_A.g = G1<snark_pp<ppT>> ::one();
        this->g_A.h = G1<snark_pp<ppT>>::one();
        this->g_B.g = G2<snark_pp<ppT>> ::one();
        this->g_B.h = G1<snark_pp<ppT>>::one();
        this->g_C.g = G1<snark_pp<ppT>> ::one();
        this->g_C.h = G1<snark_pp<ppT>>::one();
        this->g_H = G1<snark_pp<ppT>>::one();
        this->g_K = G1<snark_pp<ppT>>::one();
        g_Aau = knowledge_commitment<G1<snark_pp<ppT>>, G1<snark_pp<ppT>> >
            (G1<snark_pp<ppT>>::one(),G1<snark_pp<ppT>>::one());
        this->muA = G1<snark_pp<ppT>>::one();
    }
    r1cs_ppzkadsnark_proof(knowledge_commitment<G1<snark_pp<ppT>>,
                           G1<snark_pp<ppT>> > &&g_A,
                           knowledge_commitment<G2<snark_pp<ppT>>,
                           G1<snark_pp<ppT>> > &&g_B,
                           knowledge_commitment<G1<snark_pp<ppT>>,
                           G1<snark_pp<ppT>> > &&g_C,
                           G1<snark_pp<ppT>> &&g_H,
                           G1<snark_pp<ppT>> &&g_K,
                           knowledge_commitment<G1<snark_pp<ppT>>,
                           G1<snark_pp<ppT>> > &&g_Aau,
                           G1<snark_pp<ppT>> &&muA) :
        g_A(std::move(g_A)),
        g_B(std::move(g_B)),
        g_C(std::move(g_C)),
        g_H(std::move(g_H)),
        g_K(std::move(g_K)),
        g_Aau(std::move(g_Aau)),
        muA(std::move(muA))
    {};

    size_t G1_size() const
    {
        return 10;
    }

    size_t G2_size() const
    {
        return 1;
    }

    size_t size_in_bits() const
    {
        return G1_size() * G1<snark_pp<ppT>>::size_in_bits() + G2_size() * G2<snark_pp<ppT>>::size_in_bits();
    }

    void print_size() const
    {
        print_indent(); printf("* G1 elements in proof: %zu\n", this->G1_size());
        print_indent(); printf("* G2 elements in proof: %zu\n", this->G2_size());
        print_indent(); printf("* Proof size in bits: %zu\n", this->size_in_bits());
    }

    bool is_well_formed() const
    {
        return (g_A.g.is_well_formed() && g_A.h.is_well_formed() &&
                g_B.g.is_well_formed() && g_B.h.is_well_formed() &&
                g_C.g.is_well_formed() && g_C.h.is_well_formed() &&
                g_H.is_well_formed() &&
                g_K.is_well_formed() &&
                g_Aau.g.is_well_formed() && g_Aau.h.is_well_formed() &&
                muA.is_well_formed());
    }

    bool operator==(const r1cs_ppzkadsnark_proof<ppT> &other) const;
    friend std::ostream& operator<< <ppT>(std::ostream &out, const r1cs_ppzkadsnark_proof<ppT> &proof);
    friend std::istream& operator>> <ppT>(std::istream &in, r1cs_ppzkadsnark_proof<ppT> &proof);
};


/***************************** Main algorithms *******************************/

/**
 * R1CS ppZKADSNARK authentication parameters generator algorithm.
 */
template<typename ppT>
r1cs_ppzkadsnark_auth_keys<ppT> r1cs_ppzkadsnark_auth_generator(void);

/**
 * R1CS ppZKADSNARK authentication algorithm.
 */
template<typename ppT>
std::vector<r1cs_ppzkadsnark_auth_data<ppT>> r1cs_ppzkadsnark_auth_sign(
    const std::vector<Fr<snark_pp<ppT>>> &ins,
    const r1cs_ppzkadsnark_sec_auth_key<ppT> &sk,
    const std::vector<labelT> labels);

/**
 * R1CS ppZKADSNARK authentication verification algorithms.
 */
template<typename ppT>
bool r1cs_ppzkadsnark_auth_verify(const std::vector<Fr<snark_pp<ppT>>> &data,
                                  const std::vector<r1cs_ppzkadsnark_auth_data<ppT>> & auth_data,
                                  const r1cs_ppzkadsnark_sec_auth_key<ppT> &sak,
                                  const std::vector<labelT> &labels);

template<typename ppT>
bool r1cs_ppzkadsnark_auth_verify(const std::vector<Fr<snark_pp<ppT>>> &data,
                                  const std::vector<r1cs_ppzkadsnark_auth_data<ppT>> & auth_data,
                                  const r1cs_ppzkadsnark_pub_auth_key<ppT> &pak,
                                  const std::vector<labelT> &labels);

/**
 * A generator algorithm for the R1CS ppzkADSNARK.
 *
 * Given a R1CS constraint system CS, this algorithm produces proving and verification keys for CS.
 */
template<typename ppT>
r1cs_ppzkadsnark_keypair<ppT> r1cs_ppzkadsnark_generator(const r1cs_ppzkadsnark_constraint_system<ppT> &cs,
                                                         const r1cs_ppzkadsnark_pub_auth_prms<ppT> &prms);

/**
 * A prover algorithm for the R1CS ppzkADSNARK.
 *
 * Given a R1CS primary input X and a R1CS auxiliary input Y, this algorithm
 * produces a proof (of knowledge) that attests to the following statement:
 *               ``there exists Y such that CS(X,Y)=0''.
 * Above, CS is the R1CS constraint system that was given as input to the generator algorithm.
 */
template<typename ppT>
r1cs_ppzkadsnark_proof<ppT> r1cs_ppzkadsnark_prover(const r1cs_ppzkadsnark_proving_key<ppT> &pk,
                                                    const r1cs_ppzkadsnark_primary_input<ppT> &primary_input,
                                                    const r1cs_ppzkadsnark_auxiliary_input<ppT> &auxiliary_input,
                                                    const std::vector<r1cs_ppzkadsnark_auth_data<ppT>> &auth_data);

/*
 Below are two variants of verifier algorithm for the R1CS ppzkADSNARK.

 These are the four cases that arise from the following choices:

1) The verifier accepts a (non-processed) verification key or, instead, a processed verification key.
     In the latter case, we call the algorithm an "online verifier".

2) The verifier uses the symmetric key or the public verification key.
     In the former case we call the algorithm a "symmetric verifier".

*/

/**
 * Convert a (non-processed) verification key into a processed verification key.
 */
template<typename ppT>
r1cs_ppzkadsnark_processed_verification_key<ppT> r1cs_ppzkadsnark_verifier_process_vk(
    const r1cs_ppzkadsnark_verification_key<ppT> &vk);

/**
 * A symmetric verifier algorithm for the R1CS ppzkADSNARK that
 * accepts a non-processed verification key
 */
template<typename ppT>
bool r1cs_ppzkadsnark_verifier(const r1cs_ppzkadsnark_verification_key<ppT> &vk,
                               const r1cs_ppzkadsnark_proof<ppT> &proof,
                               const r1cs_ppzkadsnark_sec_auth_key<ppT> & sak,
                               const std::vector<labelT> &labels);

/**
 * A symmetric verifier algorithm for the R1CS ppzkADSNARK that
 * accepts a processed verification key.
 */
template<typename ppT>
bool r1cs_ppzkadsnark_online_verifier(const r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk,
                                      const r1cs_ppzkadsnark_proof<ppT> &proof,
                                      const r1cs_ppzkadsnark_sec_auth_key<ppT> & sak,
                                      const std::vector<labelT> &labels);


/**
 * A verifier algorithm for the R1CS ppzkADSNARK that
 * accepts a non-processed verification key
 */
template<typename ppT>
bool r1cs_ppzkadsnark_verifier(const r1cs_ppzkadsnark_verification_key<ppT> &vk,
                               const std::vector<r1cs_ppzkadsnark_auth_data<ppT>>  &auth_data,
                               const r1cs_ppzkadsnark_proof<ppT> &proof,
                               const r1cs_ppzkadsnark_pub_auth_key<ppT> & pak,
                               const std::vector<labelT> &labels);

/**
 * A verifier algorithm for the R1CS ppzkADSNARK that
 * accepts a processed verification key.
 */
template<typename ppT>
bool r1cs_ppzkadsnark_online_verifier(const r1cs_ppzkadsnark_processed_verification_key<ppT> &pvk,
                                      const std::vector<r1cs_ppzkadsnark_auth_data<ppT>>  &auth_data,
                                      const r1cs_ppzkadsnark_proof<ppT> &proof,
                                      const r1cs_ppzkadsnark_pub_auth_key<ppT> & pak,
                                      const std::vector<labelT> &labels);


} // libsnark

#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark.tcc"

#endif // R1CS_PPZKSNARK_HPP_

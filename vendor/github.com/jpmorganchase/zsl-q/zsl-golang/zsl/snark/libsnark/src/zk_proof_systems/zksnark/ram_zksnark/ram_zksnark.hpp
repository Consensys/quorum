/** @file
 *****************************************************************************

 Declaration of interfaces for a zkSNARK for RAM.

 This includes:
 - the class for a proving key;
 - the class for a verification key;
 - the class for a key pair (proving key & verification key);
 - the class for a proof;
 - the generator algorithm;
 - the prover algorithm;
 - the verifier algorithm.

 The implementation follows, extends, and optimizes the approach described
 in \[BCTV14]. Thus, the zkSNARK is constructed from a ppzkPCD for R1CS.


 Acronyms:

 "R1CS" = "Rank-1 Constraint Systems"
 "RAM" = "Random-Access Machines"
 "zkSNARK" = "Zero-Knowledge Succinct Non-interactive ARgument of Knowledge"
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

#ifndef RAM_ZKSNARK_HPP_
#define RAM_ZKSNARK_HPP_

#include <memory>

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/r1cs_sp_ppzkpcd.hpp"
#include "zk_proof_systems/zksnark/ram_zksnark/ram_compliance_predicate.hpp"
#include "zk_proof_systems/zksnark/ram_zksnark/ram_zksnark_params.hpp"

namespace libsnark {

/******************************** Proving key ********************************/

template<typename ram_zksnark_ppT>
class ram_zksnark_proving_key;

template<typename ram_zksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_zksnark_proving_key<ram_zksnark_ppT> &pk);

template<typename ram_zksnark_ppT>
std::istream& operator>>(std::istream &in, ram_zksnark_proving_key<ram_zksnark_ppT> &pk);

/**
 * A proving key for the RAM zkSNARK.
 */
template<typename ram_zksnark_ppT>
class ram_zksnark_proving_key {
public:
    ram_zksnark_architecture_params<ram_zksnark_ppT> ap;
    r1cs_sp_ppzkpcd_proving_key<ram_zksnark_PCD_pp<ram_zksnark_ppT> > pcd_pk;

    ram_zksnark_proving_key() {}
    ram_zksnark_proving_key(const ram_zksnark_proving_key<ram_zksnark_ppT> &other) = default;
    ram_zksnark_proving_key(ram_zksnark_proving_key<ram_zksnark_ppT> &&other) = default;
    ram_zksnark_proving_key(const ram_zksnark_architecture_params<ram_zksnark_ppT> &ap,
                            r1cs_sp_ppzkpcd_proving_key<ram_zksnark_PCD_pp<ram_zksnark_ppT> > &&pcd_pk) :
        ap(ap),
        pcd_pk(std::move(pcd_pk))
    {};

    ram_zksnark_proving_key<ram_zksnark_ppT>& operator=(const ram_zksnark_proving_key<ram_zksnark_ppT> &other) = default;

    bool operator==(const ram_zksnark_proving_key<ram_zksnark_ppT> &other) const;
    friend std::ostream& operator<< <ram_zksnark_ppT>(std::ostream &out, const ram_zksnark_proving_key<ram_zksnark_ppT> &pk);
    friend std::istream& operator>> <ram_zksnark_ppT>(std::istream &in, ram_zksnark_proving_key<ram_zksnark_ppT> &pk);
};


/******************************* Verification key ****************************/

template<typename ram_zksnark_ppT>
class ram_zksnark_verification_key;

template<typename ram_zksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_zksnark_verification_key<ram_zksnark_ppT> &vk);

template<typename ram_zksnark_ppT>
std::istream& operator>>(std::istream &in, ram_zksnark_verification_key<ram_zksnark_ppT> &vk);

/**
 * A verification key for the RAM zkSNARK.
 */
template<typename ram_zksnark_ppT>
class ram_zksnark_verification_key {
public:
    ram_zksnark_architecture_params<ram_zksnark_ppT> ap;
    r1cs_sp_ppzkpcd_verification_key<ram_zksnark_PCD_pp<ram_zksnark_ppT> > pcd_vk;

    ram_zksnark_verification_key() = default;
    ram_zksnark_verification_key(const ram_zksnark_verification_key<ram_zksnark_ppT> &other) = default;
    ram_zksnark_verification_key(ram_zksnark_verification_key<ram_zksnark_ppT> &&other) = default;
    ram_zksnark_verification_key(const ram_zksnark_architecture_params<ram_zksnark_ppT> &ap,
                                 r1cs_sp_ppzkpcd_verification_key<ram_zksnark_PCD_pp<ram_zksnark_ppT> > &&pcd_vk) :
        ap(ap),
        pcd_vk(std::move(pcd_vk))
    {};

    ram_zksnark_verification_key<ram_zksnark_ppT>& operator=(const ram_zksnark_verification_key<ram_zksnark_ppT> &other) = default;

    bool operator==(const ram_zksnark_verification_key<ram_zksnark_ppT> &other) const;
    friend std::ostream& operator<< <ram_zksnark_ppT>(std::ostream &out, const ram_zksnark_verification_key<ram_zksnark_ppT> &vk);
    friend std::istream& operator>> <ram_zksnark_ppT>(std::istream &in, ram_zksnark_verification_key<ram_zksnark_ppT> &vk);

    static ram_zksnark_verification_key<ram_zksnark_ppT> dummy_verification_key(const ram_zksnark_architecture_params<ram_zksnark_ppT> &ap);
};


/********************************** Key pair *********************************/

/**
 * A key pair for the RAM zkSNARK, which consists of a proving key and a verification key.
 */
template<typename ram_zksnark_ppT>
struct ram_zksnark_keypair {
public:
    ram_zksnark_proving_key<ram_zksnark_ppT> pk;
    ram_zksnark_verification_key<ram_zksnark_ppT> vk;

    ram_zksnark_keypair() {};
    ram_zksnark_keypair(ram_zksnark_keypair<ram_zksnark_ppT> &&other) = default;
    ram_zksnark_keypair(ram_zksnark_proving_key<ram_zksnark_ppT> &&pk,
                        ram_zksnark_verification_key<ram_zksnark_ppT> &&vk) :
        pk(std::move(pk)),
        vk(std::move(vk))
    {};
};


/*********************************** Proof ***********************************/

template<typename ram_zksnark_ppT>
class ram_zksnark_proof;

template<typename ram_zksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_zksnark_proof<ram_zksnark_ppT> &proof);

template<typename ram_zksnark_ppT>
std::istream& operator>>(std::istream &in, ram_zksnark_proof<ram_zksnark_ppT> &proof);

/**
 * A proof for the RAM zkSNARK.
 */
template<typename ram_zksnark_ppT>
class ram_zksnark_proof {
public:
    r1cs_sp_ppzkpcd_proof<ram_zksnark_PCD_pp<ram_zksnark_ppT> > PCD_proof;

    ram_zksnark_proof() = default;
    ram_zksnark_proof(r1cs_sp_ppzkpcd_proof<ram_zksnark_PCD_pp<ram_zksnark_ppT> > &&PCD_proof) :
        PCD_proof(std::move(PCD_proof)) {};
    ram_zksnark_proof(const r1cs_sp_ppzkpcd_proof<ram_zksnark_PCD_pp<ram_zksnark_ppT> > &PCD_proof) :
        PCD_proof(PCD_proof) {};

    size_t size_in_bits() const
    {
        return PCD_proof.size_in_bits();
    }

    bool operator==(const ram_zksnark_proof<ram_zksnark_ppT> &other) const;
    friend std::ostream& operator<< <ram_zksnark_ppT>(std::ostream &out, const ram_zksnark_proof<ram_zksnark_ppT> &proof);
    friend std::istream& operator>> <ram_zksnark_ppT>(std::istream &in, ram_zksnark_proof<ram_zksnark_ppT> &proof);
};


/***************************** Main algorithms *******************************/

/**
 * A generator algorithm for the RAM zkSNARK.
 *
 * Given a choice of architecture parameters, this algorithm produces proving
 * and verification keys for all computations that respect this choice.
 */
template<typename ram_zksnark_ppT>
ram_zksnark_keypair<ram_zksnark_ppT> ram_zksnark_generator(const ram_zksnark_architecture_params<ram_zksnark_ppT> &ap);

/**
 * A prover algorithm for the RAM zkSNARK.
 *
 * Given a proving key, primary input X, time bound T, and auxiliary input Y, this algorithm
 * produces a proof (of knowledge) that attests to the following statement:
 *               ``there exists Y such that X(Y) accepts within T steps''.
 */
template<typename ram_zksnark_ppT>
ram_zksnark_proof<ram_zksnark_ppT> ram_zksnark_prover(const ram_zksnark_proving_key<ram_zksnark_ppT> &pk,
                                                      const ram_zksnark_primary_input<ram_zksnark_ppT> &primary_input,
                                                      const size_t time_bound,
                                                      const ram_zksnark_auxiliary_input<ram_zksnark_ppT> &auxiliary_input);

/**
 * A verifier algorithm for the RAM zkSNARK.
 *
 * This algorithm is universal in the sense that the verification key
 * supports proof verification for *any* choice of primary input and time bound.
 */
template<typename ram_zksnark_ppT>
bool ram_zksnark_verifier(const ram_zksnark_verification_key<ram_zksnark_ppT> &vk,
                          const ram_zksnark_primary_input<ram_zksnark_ppT> &primary_input,
                          const size_t time_bound,
                          const ram_zksnark_proof<ram_zksnark_ppT> &proof);

} // libsnark

#include "zk_proof_systems/zksnark/ram_zksnark/ram_zksnark.tcc"

#endif // RAM_ZKSNARK_HPP_

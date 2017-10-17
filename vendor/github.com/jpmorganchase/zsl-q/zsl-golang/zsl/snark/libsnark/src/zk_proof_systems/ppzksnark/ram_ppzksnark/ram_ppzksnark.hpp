/** @file
 *****************************************************************************

 Declaration of interfaces for a ppzkSNARK for RAM.

 This includes:
 - the class for a proving key;
 - the class for a verification key;
 - the class for a key pair (proving key & verification key);
 - the class for a proof;
 - the generator algorithm;
 - the prover algorithm;
 - the verifier algorithm.

 The implementation follows, extends, and optimizes the approach described
 in \[BCTV14] (itself building on \[BCGTV13]). In particular, the ppzkSNARK
 for RAM is constructed from a ppzkSNARK for R1CS.


 Acronyms:

 "R1CS" = "Rank-1 Constraint Systems"
 "RAM" = "Random-Access Machines"
 "ppzkSNARK" = "Pre-Processing Zero-Knowledge Succinct Non-interactive ARgument of Knowledge"

 References:

 \[BCGTV13]:
 "SNARKs for C: verifying program executions succinctly and in zero knowledge",
 Eli Ben-Sasson, Alessandro Chiesa, Daniel Genkin, Eran Tromer, Madars Virza,
 CRYPTO 2014,
 <http://eprint.iacr.org/2013/507>

 \[BCTV14]:
 "Succinct Non-Interactive Zero Knowledge for a von Neumann Architecture",
 Eli Ben-Sasson, Alessandro Chiesa, Eran Tromer, Madars Virza,
 USENIX Security 2014,
 <http://eprint.iacr.org/2013/879>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_PPZKSNARK_HPP_
#define RAM_PPZKSNARK_HPP_

#include <memory>

#include "reductions/ram_to_r1cs/ram_to_r1cs.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp"
#include "zk_proof_systems/ppzksnark/ram_ppzksnark/ram_ppzksnark_params.hpp"

namespace libsnark {

/******************************** Proving key ********************************/

template<typename ram_ppzksnark_ppT>
class ram_ppzksnark_proving_key;

template<typename ram_ppzksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk);

template<typename ram_ppzksnark_ppT>
std::istream& operator>>(std::istream &in, ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk);

/**
 * A proving key for the RAM ppzkSNARK.
 */
template<typename ram_ppzksnark_ppT>
class ram_ppzksnark_proving_key {
public:
    typedef ram_ppzksnark_snark_pp<ram_ppzksnark_ppT> snark_ppT;

    r1cs_ppzksnark_proving_key<snark_ppT> r1cs_pk;
    ram_ppzksnark_architecture_params<ram_ppzksnark_ppT> ap;
    size_t primary_input_size_bound;
    size_t time_bound;

    ram_ppzksnark_proving_key() {}
    ram_ppzksnark_proving_key(const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &other) = default;
    ram_ppzksnark_proving_key(ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &&other) = default;
    ram_ppzksnark_proving_key(r1cs_ppzksnark_proving_key<snark_ppT> &&r1cs_pk,
                              const ram_ppzksnark_architecture_params<ram_ppzksnark_ppT> &ap,
                              const size_t primary_input_size_bound,
                              const size_t time_bound) :
        r1cs_pk(std::move(r1cs_pk)),
        ap(ap),
        primary_input_size_bound(primary_input_size_bound),
        time_bound(time_bound)
    {}

    ram_ppzksnark_proving_key<ram_ppzksnark_ppT>& operator=(const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &other) = default;

    bool operator==(const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &other) const;
    friend std::ostream& operator<< <ram_ppzksnark_ppT>(std::ostream &out, const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk);
    friend std::istream& operator>> <ram_ppzksnark_ppT>(std::istream &in, ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk);
};


/******************************* Verification key ****************************/

template<typename ram_ppzksnark_ppT>
class ram_ppzksnark_verification_key;

template<typename ram_ppzksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk);

template<typename ram_ppzksnark_ppT>
std::istream& operator>>(std::istream &in, ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk);

/**
 * A verification key for the RAM ppzkSNARK.
 */
template<typename ram_ppzksnark_ppT>
class ram_ppzksnark_verification_key {
public:
    typedef ram_ppzksnark_snark_pp<ram_ppzksnark_ppT> snark_ppT;

    r1cs_ppzksnark_verification_key<snark_ppT> r1cs_vk;
    ram_ppzksnark_architecture_params<ram_ppzksnark_ppT> ap;
    size_t primary_input_size_bound;
    size_t time_bound;

    std::set<size_t> bound_primary_input_locations;

    ram_ppzksnark_verification_key() = default;
    ram_ppzksnark_verification_key(const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &other) = default;
    ram_ppzksnark_verification_key(ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &&other) = default;
    ram_ppzksnark_verification_key(const r1cs_ppzksnark_verification_key<snark_ppT> &r1cs_vk,
                                   const ram_ppzksnark_architecture_params<ram_ppzksnark_ppT> &ap,
                                   const size_t primary_input_size_bound,
                                   const size_t time_bound) :
        r1cs_vk(r1cs_vk),
        ap(ap),
        primary_input_size_bound(primary_input_size_bound),
        time_bound(time_bound)
    {}

    ram_ppzksnark_verification_key<ram_ppzksnark_ppT>& operator=(const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &other) = default;

    ram_ppzksnark_verification_key<ram_ppzksnark_ppT> bind_primary_input(const ram_ppzksnark_primary_input<ram_ppzksnark_ppT> &primary_input) const;

    bool operator==(const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &other) const;
    friend std::ostream& operator<< <ram_ppzksnark_ppT>(std::ostream &out, const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk);
    friend std::istream& operator>> <ram_ppzksnark_ppT>(std::istream &in, ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk);
};


/********************************** Key pair *********************************/

/**
 * A key pair for the RAM ppzkSNARK, which consists of a proving key and a verification key.
 */
template<typename ram_ppzksnark_ppT>
struct ram_ppzksnark_keypair {
public:
    ram_ppzksnark_proving_key<ram_ppzksnark_ppT> pk;
    ram_ppzksnark_verification_key<ram_ppzksnark_ppT> vk;

    ram_ppzksnark_keypair() = default;
    ram_ppzksnark_keypair(ram_ppzksnark_keypair<ram_ppzksnark_ppT> &&other) = default;
    ram_ppzksnark_keypair(const ram_ppzksnark_keypair<ram_ppzksnark_ppT> &other) = default;
    ram_ppzksnark_keypair(ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &&pk,
                          ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &&vk) :
        pk(std::move(pk)),
        vk(std::move(vk))
    {}
};


/*********************************** Proof ***********************************/

/**
 * A proof for the RAM ppzkSNARK.
 */
template<typename ram_ppzksnark_ppT>
using ram_ppzksnark_proof = r1cs_ppzksnark_proof<ram_ppzksnark_snark_pp<ram_ppzksnark_ppT> >;


/***************************** Main algorithms *******************************/

/**
 * A generator algorithm for the RAM ppzkSNARK.
 *
 * Given a choice of architecture parameters and computation bounds, this algorithm
 * produces proving and verification keys for all computations that respect these choices.
 */
template<typename ram_ppzksnark_ppT>
ram_ppzksnark_keypair<ram_ppzksnark_ppT> ram_ppzksnark_generator(const ram_ppzksnark_architecture_params<ram_ppzksnark_ppT> &ap,
                                                                 const size_t primary_input_size_bound,
                                                                 const size_t time_bound);

/**
 * A prover algorithm for the RAM ppzkSNARK.
 *
 * Given a proving key, primary input X, and auxiliary input Y, this algorithm
 * produces a proof (of knowledge) that attests to the following statement:
 *               ``there exists Y such that X(Y) accepts''.
 *
 * Above, it has to be the case that the computation respects the bounds:
 * - the size of X is at most primary_input_size_bound, and
 * - the time to compute X(Y) is at most time_bound.
 */
template<typename ram_ppzksnark_ppT>
ram_ppzksnark_proof<ram_ppzksnark_ppT> ram_ppzksnark_prover(const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk,
                                                            const ram_ppzksnark_primary_input<ram_ppzksnark_ppT> &primary_input,
                                                            const ram_ppzksnark_auxiliary_input<ram_ppzksnark_ppT> &auxiliary_input);

/**
 * A verifier algorithm for the RAM ppzkSNARK.
 *
 * This algorithm is universal in the sense that the verification key
 * supports proof verification for any choice of primary input
 * provided that the computation respects the bounds.
 */
template<typename ram_ppzksnark_ppT>
bool ram_ppzksnark_verifier(const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk,
                            const ram_ppzksnark_primary_input<ram_ppzksnark_ppT> &primary_input,
                            const ram_ppzksnark_proof<ram_ppzksnark_ppT> &proof);

} // libsnark

#include "zk_proof_systems/ppzksnark/ram_ppzksnark/ram_ppzksnark.tcc"

#endif // RAM_PPZKSNARK_HPP_

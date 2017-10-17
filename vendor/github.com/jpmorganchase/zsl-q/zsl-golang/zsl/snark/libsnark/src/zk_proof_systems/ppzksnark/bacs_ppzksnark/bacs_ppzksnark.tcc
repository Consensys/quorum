/** @file
 *****************************************************************************

 Implementation of interfaces for a ppzkSNARK for BACS.

 See bacs_ppzksnark.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_PPZKSNARK_TCC_
#define BACS_PPZKSNARK_TCC_

#include "reductions/bacs_to_r1cs/bacs_to_r1cs.hpp"

namespace libsnark {


template<typename ppT>
bool bacs_ppzksnark_proving_key<ppT>::operator==(const bacs_ppzksnark_proving_key<ppT> &other) const
{
    return (this->circuit == other.circuit &&
            this->r1cs_pk == other.r1cs_pk);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const bacs_ppzksnark_proving_key<ppT> &pk)
{
    out << pk.circuit << OUTPUT_NEWLINE;
    out << pk.r1cs_pk << OUTPUT_NEWLINE;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, bacs_ppzksnark_proving_key<ppT> &pk)
{
    in >> pk.circuit;
    consume_OUTPUT_NEWLINE(in);
    in >> pk.r1cs_pk;
    consume_OUTPUT_NEWLINE(in);

    return in;
}


template<typename ppT>
bacs_ppzksnark_keypair<ppT> bacs_ppzksnark_generator(const bacs_ppzksnark_circuit<ppT> &circuit)
{
    typedef Fr<ppT> FieldT;

    enter_block("Call to bacs_ppzksnark_generator");
    const r1cs_constraint_system<FieldT> r1cs_cs = bacs_to_r1cs_instance_map<FieldT>(circuit);
    const r1cs_ppzksnark_keypair<ppT> r1cs_keypair = r1cs_ppzksnark_generator<ppT>(r1cs_cs);
    leave_block("Call to bacs_ppzksnark_generator");

    return bacs_ppzksnark_keypair<ppT>(bacs_ppzksnark_proving_key<ppT>(circuit, r1cs_keypair.pk),
                                       r1cs_keypair.vk);
}

template<typename ppT>
bacs_ppzksnark_proof<ppT> bacs_ppzksnark_prover(const bacs_ppzksnark_proving_key<ppT> &pk,
                                                const bacs_ppzksnark_primary_input<ppT> &primary_input,
                                                const bacs_ppzksnark_auxiliary_input<ppT> &auxiliary_input)
{
    typedef Fr<ppT> FieldT;

    enter_block("Call to bacs_ppzksnark_prover");
    const r1cs_variable_assignment<FieldT> r1cs_va = bacs_to_r1cs_witness_map<FieldT>(pk.circuit, primary_input, auxiliary_input);
    const r1cs_auxiliary_input<FieldT> r1cs_ai(r1cs_va.begin() + primary_input.size(), r1cs_va.end()); // TODO: faster to just change bacs_to_r1cs_witness_map into two :(
    const r1cs_ppzksnark_proof<ppT> r1cs_proof = r1cs_ppzksnark_prover<ppT>(pk.r1cs_pk, primary_input, r1cs_ai);
    leave_block("Call to bacs_ppzksnark_prover");

    return r1cs_proof;
}

template<typename ppT>
bacs_ppzksnark_processed_verification_key<ppT> bacs_ppzksnark_verifier_process_vk(const bacs_ppzksnark_verification_key<ppT> &vk)
{
    enter_block("Call to bacs_ppzksnark_verifier_process_vk");
    const bacs_ppzksnark_processed_verification_key<ppT> pvk = r1cs_ppzksnark_verifier_process_vk<ppT>(vk);
    leave_block("Call to bacs_ppzksnark_verifier_process_vk");

    return pvk;
}

template<typename ppT>
bool bacs_ppzksnark_verifier_weak_IC(const bacs_ppzksnark_verification_key<ppT> &vk,
                                     const bacs_ppzksnark_primary_input<ppT> &primary_input,
                                     const bacs_ppzksnark_proof<ppT> &proof)
{
    enter_block("Call to bacs_ppzksnark_verifier_weak_IC");
    const bacs_ppzksnark_processed_verification_key<ppT> pvk = bacs_ppzksnark_verifier_process_vk<ppT>(vk);
    const bool bit = r1cs_ppzksnark_online_verifier_weak_IC<ppT>(pvk, primary_input, proof);
    leave_block("Call to bacs_ppzksnark_verifier_weak_IC");

    return bit;
}

template<typename ppT>
bool bacs_ppzksnark_verifier_strong_IC(const bacs_ppzksnark_verification_key<ppT> &vk,
                                       const bacs_ppzksnark_primary_input<ppT> &primary_input,
                                       const bacs_ppzksnark_proof<ppT> &proof)
{
    enter_block("Call to bacs_ppzksnark_verifier_strong_IC");
    const bacs_ppzksnark_processed_verification_key<ppT> pvk = bacs_ppzksnark_verifier_process_vk<ppT>(vk);
    const bool bit = r1cs_ppzksnark_online_verifier_strong_IC<ppT>(pvk, primary_input, proof);
    leave_block("Call to bacs_ppzksnark_verifier_strong_IC");

    return bit;
}

template<typename ppT>
bool bacs_ppzksnark_online_verifier_weak_IC(const bacs_ppzksnark_processed_verification_key<ppT> &pvk,
                                            const bacs_ppzksnark_primary_input<ppT> &primary_input,
                                            const bacs_ppzksnark_proof<ppT> &proof)
{
    enter_block("Call to bacs_ppzksnark_online_verifier_weak_IC");
    const bool bit = r1cs_ppzksnark_online_verifier_weak_IC<ppT>(pvk, primary_input, proof);
    leave_block("Call to bacs_ppzksnark_online_verifier_weak_IC");

    return bit;
}

template<typename ppT>
bool bacs_ppzksnark_online_verifier_strong_IC(const bacs_ppzksnark_processed_verification_key<ppT> &pvk,
                                              const bacs_ppzksnark_primary_input<ppT> &primary_input,
                                              const bacs_ppzksnark_proof<ppT> &proof)
{
    enter_block("Call to bacs_ppzksnark_online_verifier_strong_IC");
    const bool bit = r1cs_ppzksnark_online_verifier_strong_IC<ppT>(pvk, primary_input, proof);
    leave_block("Call to bacs_ppzksnark_online_verifier_strong_IC");

    return bit;
}

} // libsnark

#endif // BACS_PPZKSNARK_TCC_

/** @file
 *****************************************************************************

 Implementation of interfaces for a ppzkSNARK for TBCS.

 See tbcs_ppzksnark.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TBCS_PPZKSNARK_TCC_
#define TBCS_PPZKSNARK_TCC_

#include "reductions/tbcs_to_uscs/tbcs_to_uscs.hpp"

namespace libsnark {


template<typename ppT>
bool tbcs_ppzksnark_proving_key<ppT>::operator==(const tbcs_ppzksnark_proving_key<ppT> &other) const
{
    return (this->circuit == other.circuit &&
            this->uscs_pk == other.uscs_pk);
}

template<typename ppT>
std::ostream& operator<<(std::ostream &out, const tbcs_ppzksnark_proving_key<ppT> &pk)
{
    out << pk.circuit << OUTPUT_NEWLINE;
    out << pk.uscs_pk << OUTPUT_NEWLINE;

    return out;
}

template<typename ppT>
std::istream& operator>>(std::istream &in, tbcs_ppzksnark_proving_key<ppT> &pk)
{
    in >> pk.circuit;
    consume_OUTPUT_NEWLINE(in);
    in >> pk.uscs_pk;
    consume_OUTPUT_NEWLINE(in);

    return in;
}


template<typename ppT>
tbcs_ppzksnark_keypair<ppT> tbcs_ppzksnark_generator(const tbcs_ppzksnark_circuit &circuit)
{
    typedef Fr<ppT> FieldT;

    enter_block("Call to tbcs_ppzksnark_generator");
    const uscs_constraint_system<FieldT> uscs_cs = tbcs_to_uscs_instance_map<FieldT>(circuit);
    const uscs_ppzksnark_keypair<ppT> uscs_keypair = uscs_ppzksnark_generator<ppT>(uscs_cs);
    leave_block("Call to tbcs_ppzksnark_generator");

    return tbcs_ppzksnark_keypair<ppT>(tbcs_ppzksnark_proving_key<ppT>(circuit, uscs_keypair.pk),
                                       uscs_keypair.vk);
}

template<typename ppT>
tbcs_ppzksnark_proof<ppT> tbcs_ppzksnark_prover(const tbcs_ppzksnark_proving_key<ppT> &pk,
                                                const tbcs_ppzksnark_primary_input &primary_input,
                                                const tbcs_ppzksnark_auxiliary_input &auxiliary_input)
{
    typedef Fr<ppT> FieldT;

    enter_block("Call to tbcs_ppzksnark_prover");
    const uscs_variable_assignment<FieldT> uscs_va = tbcs_to_uscs_witness_map<FieldT>(pk.circuit, primary_input, auxiliary_input);
    const uscs_primary_input<FieldT> uscs_pi = convert_bit_vector_to_field_element_vector<FieldT>(primary_input);
    const uscs_auxiliary_input<FieldT> uscs_ai(uscs_va.begin() + primary_input.size(), uscs_va.end()); // TODO: faster to just change bacs_to_r1cs_witness_map into two :(
    const uscs_ppzksnark_proof<ppT> uscs_proof = uscs_ppzksnark_prover<ppT>(pk.uscs_pk, uscs_pi, uscs_ai);
    leave_block("Call to tbcs_ppzksnark_prover");

    return uscs_proof;
}

template<typename ppT>
tbcs_ppzksnark_processed_verification_key<ppT> tbcs_ppzksnark_verifier_process_vk(const tbcs_ppzksnark_verification_key<ppT> &vk)
{
    enter_block("Call to tbcs_ppzksnark_verifier_process_vk");
    const tbcs_ppzksnark_processed_verification_key<ppT> pvk = uscs_ppzksnark_verifier_process_vk<ppT>(vk);
    leave_block("Call to tbcs_ppzksnark_verifier_process_vk");

    return pvk;
}

template<typename ppT>
bool tbcs_ppzksnark_verifier_weak_IC(const tbcs_ppzksnark_verification_key<ppT> &vk,
                                     const tbcs_ppzksnark_primary_input &primary_input,
                                     const tbcs_ppzksnark_proof<ppT> &proof)
{
    typedef Fr<ppT> FieldT;
    enter_block("Call to tbcs_ppzksnark_verifier_weak_IC");
    const uscs_primary_input<FieldT> uscs_input = convert_bit_vector_to_field_element_vector<FieldT>(primary_input);
    const tbcs_ppzksnark_processed_verification_key<ppT> pvk = tbcs_ppzksnark_verifier_process_vk<ppT>(vk);
    const bool bit = uscs_ppzksnark_online_verifier_weak_IC<ppT>(pvk, uscs_input, proof);
    leave_block("Call to tbcs_ppzksnark_verifier_weak_IC");

    return bit;
}

template<typename ppT>
bool tbcs_ppzksnark_verifier_strong_IC(const tbcs_ppzksnark_verification_key<ppT> &vk,
                                       const tbcs_ppzksnark_primary_input &primary_input,
                                       const tbcs_ppzksnark_proof<ppT> &proof)
{
    typedef Fr<ppT> FieldT;
    enter_block("Call to tbcs_ppzksnark_verifier_strong_IC");
    const tbcs_ppzksnark_processed_verification_key<ppT> pvk = tbcs_ppzksnark_verifier_process_vk<ppT>(vk);
    const uscs_primary_input<FieldT> uscs_input = convert_bit_vector_to_field_element_vector<FieldT>(primary_input);
    const bool bit = uscs_ppzksnark_online_verifier_strong_IC<ppT>(pvk, uscs_input, proof);
    leave_block("Call to tbcs_ppzksnark_verifier_strong_IC");

    return bit;
}

template<typename ppT>
bool tbcs_ppzksnark_online_verifier_weak_IC(const tbcs_ppzksnark_processed_verification_key<ppT> &pvk,
                                            const tbcs_ppzksnark_primary_input &primary_input,
                                            const tbcs_ppzksnark_proof<ppT> &proof)
{
    typedef Fr<ppT> FieldT;
    enter_block("Call to tbcs_ppzksnark_online_verifier_weak_IC");
    const uscs_primary_input<FieldT> uscs_input = convert_bit_vector_to_field_element_vector<FieldT>(primary_input);
    const bool bit = uscs_ppzksnark_online_verifier_weak_IC<ppT>(pvk, uscs_input, proof);
    leave_block("Call to tbcs_ppzksnark_online_verifier_weak_IC");

    return bit;
}

template<typename ppT>
bool tbcs_ppzksnark_online_verifier_strong_IC(const tbcs_ppzksnark_processed_verification_key<ppT> &pvk,
                                              const tbcs_ppzksnark_primary_input &primary_input,
                                              const tbcs_ppzksnark_proof<ppT> &proof)
{
    typedef Fr<ppT> FieldT;
    enter_block("Call to tbcs_ppzksnark_online_verifier_strong_IC");
    const uscs_primary_input<FieldT> uscs_input = convert_bit_vector_to_field_element_vector<FieldT>(primary_input);
    const bool bit = uscs_ppzksnark_online_verifier_strong_IC<ppT>(pvk, uscs_input, proof);
    leave_block("Call to tbcs_ppzksnark_online_verifier_strong_IC");

    return bit;
}

} // libsnark

#endif // TBCS_PPZKSNARK_TCC_

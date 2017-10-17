/** @file
 *****************************************************************************

 Implementation of interfaces for a *single-predicate* ppzkPCD for R1CS.

 See r1cs_sp_ppzkpcd.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_SP_PPZKPCD_TCC_
#define R1CS_SP_PPZKPCD_TCC_

#include <algorithm>
#include <cassert>
#include <iostream>

#include "common/profiling.hpp"
#include "common/utils.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/sp_pcd_circuits.hpp"

namespace libsnark {

template<typename PCD_ppT>
bool r1cs_sp_ppzkpcd_proving_key<PCD_ppT>::operator==(const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &other) const
{
    return (this->compliance_predicate == other.compliance_predicate &&
            this->compliance_step_r1cs_pk == other.compliance_step_r1cs_pk &&
            this->translation_step_r1cs_pk == other.translation_step_r1cs_pk &&
            this->compliance_step_r1cs_vk == other.compliance_step_r1cs_vk &&
            this->translation_step_r1cs_vk == other.translation_step_r1cs_vk);
}

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk)
{
    out << pk.compliance_predicate;
    out << pk.compliance_step_r1cs_pk;
    out << pk.translation_step_r1cs_pk;
    out << pk.compliance_step_r1cs_vk;
    out << pk.translation_step_r1cs_vk;

    return out;
}

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk)
{
    in >> pk.compliance_predicate;
    in >> pk.compliance_step_r1cs_pk;
    in >> pk.translation_step_r1cs_pk;
    in >> pk.compliance_step_r1cs_vk;
    in >> pk.translation_step_r1cs_vk;

    return in;
}

template<typename PCD_ppT>
bool r1cs_sp_ppzkpcd_verification_key<PCD_ppT>::operator==(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &other) const
{
    return (this->compliance_step_r1cs_vk == other.compliance_step_r1cs_vk &&
            this->translation_step_r1cs_vk == other.translation_step_r1cs_vk);
}

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk)
{
    out << vk.compliance_step_r1cs_vk;
    out << vk.translation_step_r1cs_vk;

    return out;
}

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk)
{
    in >> vk.compliance_step_r1cs_vk;
    in >> vk.translation_step_r1cs_vk;

    return in;
}

template<typename PCD_ppT>
r1cs_sp_ppzkpcd_verification_key<PCD_ppT> r1cs_sp_ppzkpcd_verification_key<PCD_ppT>::dummy_verification_key()
{
    typedef typename PCD_ppT::curve_A_pp curve_A_pp;
    typedef typename PCD_ppT::curve_B_pp curve_B_pp;

    r1cs_sp_ppzkpcd_verification_key<PCD_ppT> result;
    result.compliance_step_r1cs_vk = r1cs_ppzksnark_verification_key<typename PCD_ppT::curve_A_pp>::dummy_verification_key(sp_compliance_step_pcd_circuit_maker<curve_A_pp>::input_size_in_elts());
    result.translation_step_r1cs_vk = r1cs_ppzksnark_verification_key<typename PCD_ppT::curve_B_pp>::dummy_verification_key(sp_translation_step_pcd_circuit_maker<curve_B_pp>::input_size_in_elts());

    return result;
}

template<typename PCD_ppT>
bool r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT>::operator==(const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &other) const
{
    return (this->compliance_step_r1cs_pvk == other.compliance_step_r1cs_pvk &&
            this->translation_step_r1cs_pvk == other.translation_step_r1cs_pvk &&
            this->translation_step_r1cs_vk_bits == other.translation_step_r1cs_vk_bits);
}

template<typename PCD_ppT>
std::ostream& operator<<(std::ostream &out, const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk)
{
    out << pvk.compliance_step_r1cs_pvk;
    out << pvk.translation_step_r1cs_pvk;
    serialize_bit_vector(out, pvk.translation_step_r1cs_vk_bits);

    return out;
}

template<typename PCD_ppT>
std::istream& operator>>(std::istream &in, r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk)
{
    in >> pvk.compliance_step_r1cs_pvk;
    in >> pvk.translation_step_r1cs_pvk;
    deserialize_bit_vector(in, pvk.translation_step_r1cs_vk_bits);

    return in;
}

template<typename PCD_ppT>
r1cs_sp_ppzkpcd_keypair<PCD_ppT> r1cs_sp_ppzkpcd_generator(const r1cs_sp_ppzkpcd_compliance_predicate<PCD_ppT> &compliance_predicate)
{
    assert(Fr<typename PCD_ppT::curve_A_pp>::mod == Fq<typename PCD_ppT::curve_B_pp>::mod);
    assert(Fq<typename PCD_ppT::curve_A_pp>::mod == Fr<typename PCD_ppT::curve_B_pp>::mod);

    typedef Fr<typename PCD_ppT::curve_A_pp> FieldT_A;
    typedef Fr<typename PCD_ppT::curve_B_pp> FieldT_B;

    typedef typename PCD_ppT::curve_A_pp curve_A_pp;
    typedef typename PCD_ppT::curve_B_pp curve_B_pp;

    enter_block("Call to r1cs_sp_ppzkpcd_generator");

    assert(compliance_predicate.is_well_formed());

    enter_block("Construct compliance step PCD circuit");
    sp_compliance_step_pcd_circuit_maker<curve_A_pp> compliance_step_pcd_circuit(compliance_predicate);
    compliance_step_pcd_circuit.generate_r1cs_constraints();
    const r1cs_constraint_system<FieldT_A> compliance_step_pcd_circuit_cs = compliance_step_pcd_circuit.get_circuit();
    compliance_step_pcd_circuit_cs.report_linear_constraint_statistics();
    leave_block("Construct compliance step PCD circuit");

    enter_block("Generate key pair for compliance step PCD circuit");
    r1cs_ppzksnark_keypair<curve_A_pp> compliance_step_keypair = r1cs_ppzksnark_generator<curve_A_pp>(compliance_step_pcd_circuit_cs);
    leave_block("Generate key pair for compliance step PCD circuit");

    enter_block("Construct translation step PCD circuit");
    sp_translation_step_pcd_circuit_maker<curve_B_pp> translation_step_pcd_circuit(compliance_step_keypair.vk);
    translation_step_pcd_circuit.generate_r1cs_constraints();
    const r1cs_constraint_system<FieldT_B> translation_step_pcd_circuit_cs = translation_step_pcd_circuit.get_circuit();
    translation_step_pcd_circuit_cs.report_linear_constraint_statistics();
    leave_block("Construct translation step PCD circuit");

    enter_block("Generate key pair for translation step PCD circuit");
    r1cs_ppzksnark_keypair<curve_B_pp> translation_step_keypair = r1cs_ppzksnark_generator<curve_B_pp>(translation_step_pcd_circuit_cs);
    leave_block("Generate key pair for translation step PCD circuit");

    print_indent(); print_mem("in generator");
    leave_block("Call to r1cs_sp_ppzkpcd_generator");

    return r1cs_sp_ppzkpcd_keypair<PCD_ppT>(r1cs_sp_ppzkpcd_proving_key<PCD_ppT>(compliance_predicate,
                                                                                 std::move(compliance_step_keypair.pk),
                                                                                 std::move(translation_step_keypair.pk),
                                                                                 compliance_step_keypair.vk,
                                                                                 translation_step_keypair.vk),
                                            r1cs_sp_ppzkpcd_verification_key<PCD_ppT>(compliance_step_keypair.vk,
                                                                                      translation_step_keypair.vk));
}

template <typename PCD_ppT>
r1cs_sp_ppzkpcd_proof<PCD_ppT> r1cs_sp_ppzkpcd_prover(const r1cs_sp_ppzkpcd_proving_key<PCD_ppT> &pk,
                                                      const r1cs_sp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                                                      const r1cs_sp_ppzkpcd_auxiliary_input<PCD_ppT> &auxiliary_input,
                                                      const std::vector<r1cs_sp_ppzkpcd_proof<PCD_ppT> > &incoming_proofs)
{
    typedef Fr<typename PCD_ppT::curve_A_pp> FieldT_A;
    typedef Fr<typename PCD_ppT::curve_B_pp> FieldT_B;

    typedef typename PCD_ppT::curve_A_pp curve_A_pp;
    typedef typename PCD_ppT::curve_B_pp curve_B_pp;

    enter_block("Call to r1cs_sp_ppzkpcd_prover");

    const bit_vector translation_step_r1cs_vk_bits = r1cs_ppzksnark_verification_key_variable<curve_A_pp>::get_verification_key_bits(pk.translation_step_r1cs_vk);
#ifdef DEBUG
    printf("Outgoing message:\n");
    primary_input.outgoing_message->print();
#endif

    enter_block("Prove compliance step");
    sp_compliance_step_pcd_circuit_maker<curve_A_pp> compliance_step_pcd_circuit(pk.compliance_predicate);
    compliance_step_pcd_circuit.generate_r1cs_witness(pk.translation_step_r1cs_vk,
                                                      primary_input,
                                                      auxiliary_input,
                                                      incoming_proofs);

    const r1cs_primary_input<FieldT_A> compliance_step_primary_input = compliance_step_pcd_circuit.get_primary_input();
    const r1cs_auxiliary_input<FieldT_A> compliance_step_auxiliary_input = compliance_step_pcd_circuit.get_auxiliary_input();

    const r1cs_ppzksnark_proof<curve_A_pp> compliance_step_proof = r1cs_ppzksnark_prover<curve_A_pp>(pk.compliance_step_r1cs_pk, compliance_step_primary_input, compliance_step_auxiliary_input);
    leave_block("Prove compliance step");

#ifdef DEBUG
    const r1cs_primary_input<FieldT_A> compliance_step_input = get_sp_compliance_step_pcd_circuit_input<curve_A_pp>(translation_step_r1cs_vk_bits, primary_input);
    const bool compliance_step_ok = r1cs_ppzksnark_verifier_strong_IC<curve_A_pp>(pk.compliance_step_r1cs_vk, compliance_step_input, compliance_step_proof);
    assert(compliance_step_ok);
#endif

    enter_block("Prove translation step");
    sp_translation_step_pcd_circuit_maker<curve_B_pp> translation_step_pcd_circuit(pk.compliance_step_r1cs_vk);

    const r1cs_primary_input<FieldT_B> translation_step_primary_input = get_sp_translation_step_pcd_circuit_input<curve_B_pp>(translation_step_r1cs_vk_bits, primary_input);
    translation_step_pcd_circuit.generate_r1cs_witness(translation_step_primary_input, compliance_step_proof); // TODO: potential for better naming

    const r1cs_auxiliary_input<FieldT_B> translation_step_auxiliary_input = translation_step_pcd_circuit.get_auxiliary_input();
    const r1cs_ppzksnark_proof<curve_B_pp> translation_step_proof = r1cs_ppzksnark_prover<curve_B_pp>(pk.translation_step_r1cs_pk, translation_step_primary_input, translation_step_auxiliary_input);
    leave_block("Prove translation step");

#ifdef DEBUG
    const bool translation_step_ok = r1cs_ppzksnark_verifier_strong_IC<curve_B_pp>(pk.translation_step_r1cs_vk, translation_step_primary_input, translation_step_proof);
    assert(translation_step_ok);
#endif

    print_indent(); print_mem("in prover");
    leave_block("Call to r1cs_sp_ppzkpcd_prover");

    return translation_step_proof;
}

template<typename PCD_ppT>
bool r1cs_sp_ppzkpcd_online_verifier(const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> &pvk,
                                     const r1cs_sp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                                     const r1cs_sp_ppzkpcd_proof<PCD_ppT> &proof)

{
    typedef typename PCD_ppT::curve_B_pp curve_B_pp;

    enter_block("Call to r1cs_sp_ppzkpcd_online_verifier");
    const r1cs_primary_input<Fr<curve_B_pp> > r1cs_input = get_sp_translation_step_pcd_circuit_input<curve_B_pp>(pvk.translation_step_r1cs_vk_bits, primary_input);
    const bool result = r1cs_ppzksnark_online_verifier_strong_IC(pvk.translation_step_r1cs_pvk, r1cs_input, proof);
    print_indent(); print_mem("in online verifier");
    leave_block("Call to r1cs_sp_ppzkpcd_online_verifier");

    return result;
}

template<typename PCD_ppT>
r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> r1cs_sp_ppzkpcd_process_vk(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk)
{
    typedef typename PCD_ppT::curve_A_pp curve_A_pp;
    typedef typename PCD_ppT::curve_B_pp curve_B_pp;

    enter_block("Call to r1cs_sp_ppzkpcd_processed_verification_key");
    r1cs_ppzksnark_processed_verification_key<curve_A_pp> compliance_step_r1cs_pvk = r1cs_ppzksnark_verifier_process_vk<curve_A_pp>(vk.compliance_step_r1cs_vk);
    r1cs_ppzksnark_processed_verification_key<curve_B_pp> translation_step_r1cs_pvk = r1cs_ppzksnark_verifier_process_vk<curve_B_pp>(vk.translation_step_r1cs_vk);
    const bit_vector translation_step_r1cs_vk_bits = r1cs_ppzksnark_verification_key_variable<curve_A_pp>::get_verification_key_bits(vk.translation_step_r1cs_vk);
    leave_block("Call to r1cs_sp_ppzkpcd_processed_verification_key");

    return r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT>(std::move(compliance_step_r1cs_pvk),
                                                               std::move(translation_step_r1cs_pvk),
                                                               translation_step_r1cs_vk_bits);
}


template<typename PCD_ppT>
bool r1cs_sp_ppzkpcd_verifier(const r1cs_sp_ppzkpcd_verification_key<PCD_ppT> &vk,
                                     const r1cs_sp_ppzkpcd_primary_input<PCD_ppT> &primary_input,
                              const r1cs_sp_ppzkpcd_proof<PCD_ppT> &proof)
{
    enter_block("Call to r1cs_sp_ppzkpcd_verifier");
    const r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> pvk = r1cs_sp_ppzkpcd_process_vk(vk);
    const bool result = r1cs_sp_ppzkpcd_online_verifier(pvk, primary_input, proof);
    print_indent(); print_mem("in verifier");
    leave_block("Call to r1cs_sp_ppzkpcd_verifier");

    return result;
}


} // libsnark

#endif // R1CS_SP_PPZKPCD_TCC_

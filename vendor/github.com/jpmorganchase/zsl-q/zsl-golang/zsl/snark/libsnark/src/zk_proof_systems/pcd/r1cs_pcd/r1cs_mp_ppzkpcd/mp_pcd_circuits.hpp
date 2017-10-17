/** @file
 *****************************************************************************

 Declaration of functionality for creating and using the two PCD circuits in
 a multi-predicate PCD construction.

 The implementation follows, extends, and optimizes the approach described
 in \[CTV15]. At high level, there is a "compliance step" circuit and a
 "translation step" circuit, for each compliance predicate. For more details,
 see \[CTV15].


 References:

 \[CTV15]:
 "Cluster Computing in Zero Knowledge",
 Alessandro Chiesa, Eran Tromer, Madars Virza

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MP_PCD_CIRCUITS_HPP_
#define MP_PCD_CIRCUITS_HPP_

#include "gadgetlib1/gadget.hpp"
#include "gadgetlib1/gadgets/gadget_from_r1cs.hpp"
#include "gadgetlib1/gadgets/hashes/crh_gadget.hpp"
#include "gadgetlib1/gadgets/set_commitment/set_commitment_gadget.hpp"
#include "gadgetlib1/gadgets/verifiers/r1cs_ppzksnark_verifier_gadget.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/cp_handler.hpp"

namespace libsnark {

/**************************** Compliance step ********************************/

/**
 * A compliance-step PCD circuit.
 *
 * The circuit is an R1CS that checks compliance (for the given compliance predicate)
 * and validity of previous proofs.
 */
template<typename ppT>
class mp_compliance_step_pcd_circuit_maker {
public:
    typedef Fr<ppT> FieldT;

    r1cs_pcd_compliance_predicate<FieldT> compliance_predicate;

    protoboard<FieldT> pb;

    pb_variable<FieldT> zero;

    std::shared_ptr<block_variable<FieldT> > block_for_outgoing_message;
    std::shared_ptr<CRH_with_field_out_gadget<FieldT> > hash_outgoing_message;

    std::vector<block_variable<FieldT> > block_for_incoming_messages;
    std::vector<pb_variable_array<FieldT> > commitment_and_incoming_message_digests;
    std::vector<multipacking_gadget<FieldT> > unpack_commitment_and_incoming_message_digests;
    std::vector<pb_variable_array<FieldT> > commitment_and_incoming_messages_digest_bits;
    std::vector<CRH_with_field_out_gadget<FieldT> > hash_incoming_messages;

    std::vector<r1cs_ppzksnark_verification_key_variable<ppT> > translation_step_vks;
    std::vector<pb_variable_array<FieldT> > translation_step_vks_bits;

    pb_variable<FieldT> outgoing_message_type;
    pb_variable_array<FieldT> outgoing_message_payload;
    pb_variable_array<FieldT> outgoing_message_vars;

    pb_variable<FieldT> arity;
    std::vector<pb_variable<FieldT> > incoming_message_types;
    std::vector<pb_variable_array<FieldT> > incoming_message_payloads;
    std::vector<pb_variable_array<FieldT> > incoming_message_vars;

    pb_variable_array<FieldT> local_data;
    pb_variable_array<FieldT> cp_witness;
    std::shared_ptr<gadget_from_r1cs<FieldT> > compliance_predicate_as_gadget;

    pb_variable_array<FieldT> outgoing_message_bits;
    std::shared_ptr<multipacking_gadget<FieldT> > unpack_outgoing_message;

    std::vector<pb_variable_array<FieldT> > incoming_messages_bits;
    std::vector<multipacking_gadget<FieldT> > unpack_incoming_messages;

    pb_variable_array<FieldT> mp_compliance_step_pcd_circuit_input;
    pb_variable_array<FieldT> padded_translation_step_vk_and_outgoing_message_digest;
    std::vector<pb_variable_array<FieldT> > padded_commitment_and_incoming_messages_digest;

    std::shared_ptr<set_commitment_variable<FieldT, CRH_with_bit_out_gadget<FieldT> > > commitment;
    std::vector<set_membership_proof_variable<FieldT, CRH_with_bit_out_gadget<FieldT> > > membership_proofs;
    std::vector<set_commitment_gadget<FieldT, CRH_with_bit_out_gadget<FieldT> > > membership_checkers;
    pb_variable_array<FieldT> membership_check_results;
    pb_variable<FieldT> common_type;
    pb_variable_array<FieldT> common_type_check_aux;

    std::vector<pb_variable_array<FieldT> > verifier_input;
    std::vector<r1cs_ppzksnark_proof_variable<ppT> > proof;
    pb_variable_array<FieldT> verification_results;
    std::vector<r1cs_ppzksnark_verifier_gadget<ppT> > verifier;

    mp_compliance_step_pcd_circuit_maker(const r1cs_pcd_compliance_predicate<FieldT> &compliance_predicate,
                                         const size_t max_number_of_predicates);
    void generate_r1cs_constraints();
    r1cs_constraint_system<FieldT> get_circuit() const;

    void generate_r1cs_witness(const set_commitment &commitment_to_translation_step_r1cs_vks,
                               const std::vector<r1cs_ppzksnark_verification_key<other_curve<ppT> > > &mp_translation_step_pcd_circuit_vks,
                               const std::vector<set_membership_proof> &vk_membership_proofs,
                               const r1cs_pcd_compliance_predicate_primary_input<FieldT> &compliance_predicate_primary_input,
                               const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> &compliance_predicate_auxiliary_input,
                               const std::vector<r1cs_ppzksnark_proof<other_curve<ppT> > > &translation_step_proofs);
    r1cs_primary_input<FieldT> get_primary_input() const;
    r1cs_auxiliary_input<FieldT> get_auxiliary_input() const;

    static size_t field_logsize();
    static size_t field_capacity();
    static size_t input_size_in_elts();
    static size_t input_capacity_in_bits();
    static size_t input_size_in_bits();
};

/*************************** Translation step ********************************/

/**
 * A translation-step PCD circuit.
 *
 * The circuit is an R1CS that checks validity of previous proofs.
 */
template<typename ppT>
class mp_translation_step_pcd_circuit_maker {
public:
    typedef Fr<ppT> FieldT;

    protoboard<FieldT> pb;

    pb_variable_array<FieldT> mp_translation_step_pcd_circuit_input;
    pb_variable_array<FieldT> unpacked_mp_translation_step_pcd_circuit_input;
    pb_variable_array<FieldT> verifier_input;
    std::shared_ptr<multipacking_gadget<FieldT> > unpack_mp_translation_step_pcd_circuit_input;

    std::shared_ptr<r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> > hardcoded_compliance_step_vk;
    std::shared_ptr<r1cs_ppzksnark_proof_variable<ppT> > proof;
    std::shared_ptr<r1cs_ppzksnark_online_verifier_gadget<ppT> > online_verifier;

    mp_translation_step_pcd_circuit_maker(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &compliance_step_vk);
    void generate_r1cs_constraints();
    r1cs_constraint_system<FieldT> get_circuit() const;

    void generate_r1cs_witness(const r1cs_primary_input<Fr<ppT> > translation_step_input,
                               const r1cs_ppzksnark_proof<other_curve<ppT> > &prev_proof);
    r1cs_primary_input<FieldT> get_primary_input() const;
    r1cs_auxiliary_input<FieldT> get_auxiliary_input() const;

    static size_t field_logsize();
    static size_t field_capacity();
    static size_t input_size_in_elts();
    static size_t input_capacity_in_bits();
    static size_t input_size_in_bits();
};

/****************************** Input maps ***********************************/

/**
 * Obtain the primary input for a compliance-step PCD circuit.
 */
template<typename ppT>
r1cs_primary_input<Fr<ppT> > get_mp_compliance_step_pcd_circuit_input(const set_commitment &commitment_to_translation_step_r1cs_vks,
                                                                      const r1cs_pcd_compliance_predicate_primary_input<Fr<ppT> > &primary_input);

/**
 * Obtain the primary input for a translation-step PCD circuit.
 */
template<typename ppT>
r1cs_primary_input<Fr<ppT> > get_mp_translation_step_pcd_circuit_input(const set_commitment &commitment_to_translation_step_r1cs_vks,
                                                                       const r1cs_pcd_compliance_predicate_primary_input<Fr<other_curve<ppT> > > &primary_input);

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_mp_ppzkpcd/mp_pcd_circuits.tcc"

#endif // MP_PCD_CIRCUITS_HPP_

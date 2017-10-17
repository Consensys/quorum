/** @file
 *****************************************************************************

 Declaration of functionality for creating and using the two PCD circuits in
 a single-predicate PCD construction.

 The implementation follows, extends, and optimizes the approach described
 in \[BCTV14]. At high level, there is a "compliance step" circuit and a
 "translation step" circuit. For more details see Section 4 of \[BCTV14].


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

#ifndef SP_PCD_CIRCUITS_HPP_
#define SP_PCD_CIRCUITS_HPP_

#include "gadgetlib1/protoboard.hpp"
#include "gadgetlib1/gadgets/gadget_from_r1cs.hpp"
#include "gadgetlib1/gadgets/hashes/crh_gadget.hpp"
#include "gadgetlib1/gadgets/pairing/pairing_params.hpp"
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
class sp_compliance_step_pcd_circuit_maker {
public:
    typedef Fr<ppT> FieldT;

    r1cs_pcd_compliance_predicate<FieldT> compliance_predicate;

    protoboard<FieldT> pb;

    pb_variable<FieldT> zero;

    std::shared_ptr<block_variable<FieldT> > block_for_outgoing_message;
    std::shared_ptr<CRH_with_field_out_gadget<FieldT> > hash_outgoing_message;

    std::vector<block_variable<FieldT> > blocks_for_incoming_messages;
    std::vector<pb_variable_array<FieldT> > sp_translation_step_vk_and_incoming_message_payload_digests;
    std::vector<multipacking_gadget<FieldT> > unpack_sp_translation_step_vk_and_incoming_message_payload_digests;
    std::vector<pb_variable_array<FieldT> > sp_translation_step_vk_and_incoming_message_payload_digest_bits;
    std::vector<CRH_with_field_out_gadget<FieldT> > hash_incoming_messages;

    std::shared_ptr<r1cs_ppzksnark_verification_key_variable<ppT> > sp_translation_step_vk;
    pb_variable_array<FieldT> sp_translation_step_vk_bits;

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

    pb_variable_array<FieldT> sp_compliance_step_pcd_circuit_input;
    pb_variable_array<FieldT> padded_translation_step_vk_and_outgoing_message_digest;
    std::vector<pb_variable_array<FieldT> > padded_translation_step_vk_and_incoming_messages_digests;

    std::vector<pb_variable_array<FieldT> > verifier_input;
    std::vector<r1cs_ppzksnark_proof_variable<ppT> > proof;
    pb_variable<FieldT> verification_result;
    std::vector<r1cs_ppzksnark_verifier_gadget<ppT> > verifiers;

    sp_compliance_step_pcd_circuit_maker(const r1cs_pcd_compliance_predicate<FieldT> &compliance_predicate);
    void generate_r1cs_constraints();
    r1cs_constraint_system<FieldT> get_circuit() const;

    void generate_r1cs_witness(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &translation_step_pcd_circuit_vk,
                               const r1cs_pcd_compliance_predicate_primary_input<FieldT> &compliance_predicate_primary_input,
                               const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> &compliance_predicate_auxiliary_input,
                               const std::vector<r1cs_ppzksnark_proof<other_curve<ppT> > > &incoming_proofs);
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
class sp_translation_step_pcd_circuit_maker {
public:
    typedef Fr<ppT> FieldT;

    protoboard<FieldT> pb;

    pb_variable_array<FieldT> sp_translation_step_pcd_circuit_input;
    pb_variable_array<FieldT> unpacked_sp_translation_step_pcd_circuit_input;
    pb_variable_array<FieldT> verifier_input;
    std::shared_ptr<multipacking_gadget<FieldT> > unpack_sp_translation_step_pcd_circuit_input;

    std::shared_ptr<r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT> > hardcoded_sp_compliance_step_vk;
    std::shared_ptr<r1cs_ppzksnark_proof_variable<ppT> > proof;
    std::shared_ptr<r1cs_ppzksnark_online_verifier_gadget<ppT> > online_verifier;

    sp_translation_step_pcd_circuit_maker(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &compliance_step_vk);
    void generate_r1cs_constraints();
    r1cs_constraint_system<FieldT> get_circuit() const;

    void generate_r1cs_witness(const r1cs_primary_input<Fr<ppT> > translation_step_input,
                               const r1cs_ppzksnark_proof<other_curve<ppT> > &compliance_step_proof);
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
r1cs_primary_input<Fr<ppT> > get_sp_compliance_step_pcd_circuit_input(const bit_vector &sp_translation_step_vk_bits,
                                                                      const r1cs_pcd_compliance_predicate_primary_input<Fr<ppT> > &primary_input);

/**
 * Obtain the primary input for a translation-step PCD circuit.
 */
template<typename ppT>
r1cs_primary_input<Fr<ppT> > get_sp_translation_step_pcd_circuit_input(const bit_vector &sp_translation_step_vk_bits,
                                                                       const r1cs_pcd_compliance_predicate_primary_input<Fr<other_curve<ppT> > > &primary_input);

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/sp_pcd_circuits.tcc"

#endif // SP_PCD_CIRCUITS_HPP_

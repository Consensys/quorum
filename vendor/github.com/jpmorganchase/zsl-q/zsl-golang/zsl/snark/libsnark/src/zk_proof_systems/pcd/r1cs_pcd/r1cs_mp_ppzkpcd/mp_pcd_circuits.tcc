/** @file
 *****************************************************************************

 Implementation of functionality for creating and using the two PCD circuits in
 a multi-predicate PCD construction.

 See mp_pcd_circuits.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MP_PCD_CIRCUITS_TCC_
#define MP_PCD_CIRCUITS_TCC_

#include <algorithm>
#include "common/utils.hpp"
#include "gadgetlib1/constraint_profiling.hpp"

namespace libsnark {

template<typename ppT>
mp_compliance_step_pcd_circuit_maker<ppT>::mp_compliance_step_pcd_circuit_maker(const r1cs_pcd_compliance_predicate<FieldT> &compliance_predicate,
                                                                                const size_t max_number_of_predicates) :
    compliance_predicate(compliance_predicate)
{
    /* calculate some useful sizes */
    const size_t digest_size = CRH_with_field_out_gadget<FieldT>::get_digest_len();
    const size_t outgoing_msg_size_in_bits = field_logsize() * (1 + compliance_predicate.outgoing_message_payload_length);
    assert(compliance_predicate.has_equal_input_lengths());
    const size_t translation_step_vk_size_in_bits = r1cs_ppzksnark_verification_key_variable<ppT>::size_in_bits(mp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_elts());
    const size_t padded_verifier_input_size = mp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_capacity_in_bits();
    const size_t commitment_size = set_commitment_gadget<FieldT, CRH_with_bit_out_gadget<FieldT> >::root_size_in_bits();

    const size_t output_block_size = commitment_size + outgoing_msg_size_in_bits;
    const size_t max_incoming_payload_length = *std::max_element(compliance_predicate.incoming_message_payload_lengths.begin(), compliance_predicate.incoming_message_payload_lengths.end());
    const size_t max_input_block_size = commitment_size + field_logsize() * (1 + max_incoming_payload_length);

    CRH_with_bit_out_gadget<FieldT>::sample_randomness(std::max(output_block_size, max_input_block_size));

    /* allocate input of the compliance MP_PCD circuit */
    mp_compliance_step_pcd_circuit_input.allocate(pb, input_size_in_elts(), "mp_compliance_step_pcd_circuit_input");

    /* allocate inputs to the compliance predicate */
    outgoing_message_type.allocate(pb, "outgoing_message_type");
    outgoing_message_payload.allocate(pb, compliance_predicate.outgoing_message_payload_length, "outgoing_message_payload");

    outgoing_message_vars.insert(outgoing_message_vars.end(), outgoing_message_type);
    outgoing_message_vars.insert(outgoing_message_vars.end(), outgoing_message_payload.begin(), outgoing_message_payload.end());

    arity.allocate(pb, "arity");

    incoming_message_types.resize(compliance_predicate.max_arity);
    incoming_message_payloads.resize(compliance_predicate.max_arity);
    incoming_message_vars.resize(compliance_predicate.max_arity);
    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        incoming_message_types[i].allocate(pb, FMT("", "incoming_message_type_%zu", i));
        incoming_message_payloads[i].allocate(pb, compliance_predicate.incoming_message_payload_lengths[i], FMT("", "incoming_message_payloads_%zu", i));

        incoming_message_vars[i].insert(incoming_message_vars[i].end(), incoming_message_types[i]);
        incoming_message_vars[i].insert(incoming_message_vars[i].end(), incoming_message_payloads[i].begin(), incoming_message_payloads[i].end());
    }

    local_data.allocate(pb, compliance_predicate.local_data_length, "local_data");
    cp_witness.allocate(pb, compliance_predicate.witness_length, "cp_witness");

    /* convert compliance predicate from a constraint system into a gadget */
    pb_variable_array<FieldT> incoming_messages_concat;
    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        incoming_messages_concat.insert(incoming_messages_concat.end(), incoming_message_vars[i].begin(), incoming_message_vars[i].end());
    }

    compliance_predicate_as_gadget.reset(new gadget_from_r1cs<FieldT>(pb,
        { outgoing_message_vars,
          pb_variable_array<FieldT>(1, arity),
          incoming_messages_concat,
          local_data,
          cp_witness },
            compliance_predicate.constraint_system, "compliance_predicate_as_gadget"));

    /* unpack messages to bits */
    outgoing_message_bits.allocate(pb, outgoing_msg_size_in_bits, "outgoing_message_bits");
    unpack_outgoing_message.reset(new multipacking_gadget<FieldT>(pb, outgoing_message_bits, outgoing_message_vars, field_logsize(), "unpack_outgoing_message"));

    incoming_messages_bits.resize(compliance_predicate.max_arity);
    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        const size_t incoming_msg_size_in_bits = field_logsize() * (1 + compliance_predicate.incoming_message_payload_lengths[i]);

        incoming_messages_bits[i].allocate(pb, incoming_msg_size_in_bits, FMT("", "incoming_messages_bits_%zu", i));
        unpack_incoming_messages.emplace_back(multipacking_gadget<FieldT>(pb, incoming_messages_bits[i], incoming_message_vars[i], field_logsize(), FMT("", "unpack_incoming_messages_%zu", i)));
    }

    /* allocate digests */
    commitment_and_incoming_message_digests.resize(compliance_predicate.max_arity);
    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        commitment_and_incoming_message_digests[i].allocate(pb, digest_size, FMT("", "commitment_and_incoming_message_digests_%zu", i));
    }

    /* allocate commitment, verification key(s) and membership checker(s)/proof(s) */
    commitment.reset(new set_commitment_variable<FieldT, CRH_with_bit_out_gadget<FieldT> >(pb, commitment_size, "commitment"));

    print_indent(); printf("* %s perform same type optimization for compliance predicate with type %zu\n",
                           (compliance_predicate.relies_on_same_type_inputs ? "Will" : "Will NOT"),
                           compliance_predicate.type);
    if (compliance_predicate.relies_on_same_type_inputs)
    {
        /* only one set_commitment_gadget is needed */
        common_type.allocate(pb, "common_type");
        common_type_check_aux.allocate(pb, compliance_predicate.accepted_input_types.size(), "common_type_check_aux");

        translation_step_vks_bits.resize(1);
        translation_step_vks_bits[0].allocate(pb, translation_step_vk_size_in_bits, "translation_step_vk_bits");
        membership_check_results.allocate(pb, 1, "membership_check_results");

        membership_proofs.emplace_back(set_membership_proof_variable<FieldT, CRH_with_bit_out_gadget<FieldT>>(pb,
                                                                                                              max_number_of_predicates,
                                                                                                              "membership_proof"));
        membership_checkers.emplace_back(set_commitment_gadget<FieldT, CRH_with_bit_out_gadget<FieldT>>(pb,
                                                                                                        max_number_of_predicates,
                                                                                                        translation_step_vks_bits[0],
                                                                                                        *commitment,
                                                                                                        membership_proofs[0],
                                                                                                        membership_check_results[0], "membership_checker"));
    }
    else
    {
        /* check for max_arity possibly different VKs */
        translation_step_vks_bits.resize(compliance_predicate.max_arity);
        membership_check_results.allocate(pb, compliance_predicate.max_arity, "membership_check_results");

        for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
        {
            translation_step_vks_bits[i].allocate(pb, translation_step_vk_size_in_bits, FMT("", "translation_step_vks_bits_%zu", i));

            membership_proofs.emplace_back(set_membership_proof_variable<FieldT, CRH_with_bit_out_gadget<FieldT> >(pb,
                                                                                                                   max_number_of_predicates,
                                                                                                                   FMT("", "membership_proof_%zu", i)));
            membership_checkers.emplace_back(set_commitment_gadget<FieldT, CRH_with_bit_out_gadget<FieldT> >(pb,
                                                                                                             max_number_of_predicates,
                                                                                                             translation_step_vks_bits[i],
                                                                                                             *commitment,
                                                                                                             membership_proofs[i],
                                                                                                             membership_check_results[i],
                                                                                                             FMT("", "membership_checkers_%zu", i)));
        }
    }

    /* allocate blocks */
    block_for_outgoing_message.reset(new block_variable<FieldT>(pb, { commitment->bits, outgoing_message_bits }, "block_for_outgoing_message"));

    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        block_for_incoming_messages.emplace_back(block_variable<FieldT>(pb, { commitment->bits, incoming_messages_bits[i] }, FMT("", "block_for_incoming_messages_%zu", i)));
    }

    /* allocate hash checkers */
    hash_outgoing_message.reset(new CRH_with_field_out_gadget<FieldT>(pb, output_block_size, *block_for_outgoing_message, mp_compliance_step_pcd_circuit_input, "hash_outgoing_message"));

    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        const size_t input_block_size = commitment_size + incoming_messages_bits[i].size();
        hash_incoming_messages.emplace_back(CRH_with_field_out_gadget<FieldT>(pb, input_block_size, block_for_incoming_messages[i], commitment_and_incoming_message_digests[i], FMT("", "hash_incoming_messages_%zu", i)));
    }

    /* allocate useful zero variable */
    zero.allocate(pb, "zero");

    /* prepare arguments for the verifier */
    if (compliance_predicate.relies_on_same_type_inputs)
    {
        translation_step_vks.emplace_back(r1cs_ppzksnark_verification_key_variable<ppT>(pb, translation_step_vks_bits[0], mp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_elts(), "translation_step_vk"));
    }
    else
    {
        for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
        {
            translation_step_vks.emplace_back(r1cs_ppzksnark_verification_key_variable<ppT>(pb, translation_step_vks_bits[i], mp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_elts(), FMT("", "translation_step_vks_%zu", i)));
        }
    }

    verification_results.allocate(pb, compliance_predicate.max_arity, "verification_results");
    commitment_and_incoming_messages_digest_bits.resize(compliance_predicate.max_arity);

    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        commitment_and_incoming_messages_digest_bits[i].allocate(pb, digest_size * field_logsize(), FMT("", "commitment_and_incoming_messages_digest_bits_%zu", i));
        unpack_commitment_and_incoming_message_digests.emplace_back(multipacking_gadget<FieldT>(pb,
                                                                                                commitment_and_incoming_messages_digest_bits[i],
                                                                                                commitment_and_incoming_message_digests[i],
                                                                                                field_logsize(),
                                                                                                FMT("", "unpack_commitment_and_incoming_message_digests_%zu", i)));

        verifier_input.emplace_back(commitment_and_incoming_messages_digest_bits[i]);
        while (verifier_input[i].size() < padded_verifier_input_size)
        {
            verifier_input[i].emplace_back(zero);
        }

        proof.emplace_back(r1cs_ppzksnark_proof_variable<ppT>(pb, FMT("", "proof_%zu", i)));
        const r1cs_ppzksnark_verification_key_variable<ppT> &vk_to_be_used = (compliance_predicate.relies_on_same_type_inputs ? translation_step_vks[0] : translation_step_vks[i]);
        verifier.emplace_back(r1cs_ppzksnark_verifier_gadget<ppT>(pb,
                                                                  vk_to_be_used,
                                                                  verifier_input[i],
                                                                  mp_translation_step_pcd_circuit_maker<other_curve<ppT> >::field_capacity(),
                                                                  proof[i],
                                                                  verification_results[i],
                                                                  FMT("", "verifier_%zu", i)));
    }

    pb.set_input_sizes(input_size_in_elts());
}

template<typename ppT>
void mp_compliance_step_pcd_circuit_maker<ppT>::generate_r1cs_constraints()
{
    const size_t digest_size = CRH_with_bit_out_gadget<FieldT>::get_digest_len();
    const size_t dimension = knapsack_dimension<FieldT>::dimension;
    print_indent(); printf("* Knapsack dimension: %zu\n", dimension);

    print_indent(); printf("* Compliance predicate arity: %zu\n", compliance_predicate.max_arity);
    print_indent(); printf("* Compliance predicate outgoing payload length: %zu\n", compliance_predicate.outgoing_message_payload_length);
    print_indent(); printf("* Compliance predicate inncoming payload lengts:");
    for (auto l : compliance_predicate.incoming_message_payload_lengths)
    {
        printf(" %zu", l);
    }
    printf("\n");
    print_indent(); printf("* Compliance predicate local data length: %zu\n", compliance_predicate.local_data_length);
    print_indent(); printf("* Compliance predicate witness length: %zu\n", compliance_predicate.witness_length);

    PROFILE_CONSTRAINTS(pb, "booleanity")
    {
        PROFILE_CONSTRAINTS(pb, "booleanity: unpack outgoing_message")
        {
            unpack_outgoing_message->generate_r1cs_constraints(true);
        }

        PROFILE_CONSTRAINTS(pb, "booleanity: unpack s incoming_messages")
        {
            for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
            {
                unpack_incoming_messages[i].generate_r1cs_constraints(true);
            }
        }

        PROFILE_CONSTRAINTS(pb, "booleanity: unpack verification key")
        {
            for (size_t i = 0; i < translation_step_vks.size(); ++i)
            {
                translation_step_vks[i].generate_r1cs_constraints(true);
            }
        }
    }

    PROFILE_CONSTRAINTS(pb, "(1+s) copies of hash")
    {
        print_indent(); printf("* Digest-size: %zu\n", digest_size);
        hash_outgoing_message->generate_r1cs_constraints();

        for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
        {
            hash_incoming_messages[i].generate_r1cs_constraints();
        }
    }

    PROFILE_CONSTRAINTS(pb, "s copies of repacking circuit for verifier")
    {
        for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
        {
            unpack_commitment_and_incoming_message_digests[i].generate_r1cs_constraints(true);
        }
    }

    PROFILE_CONSTRAINTS(pb, "set membership check")
    {
        for (auto &membership_proof : membership_proofs)
        {
            membership_proof.generate_r1cs_constraints();
        }

        for (auto &membership_checker : membership_checkers)
        {
            membership_checker.generate_r1cs_constraints();
        }
    }

    PROFILE_CONSTRAINTS(pb, "compliance predicate")
    {
        compliance_predicate_as_gadget->generate_r1cs_constraints();
    }

    PROFILE_CONSTRAINTS(pb, "s copies of verifier for translated proofs")
    {
        PROFILE_CONSTRAINTS(pb, "check that s proofs lie on the curve")
        {
            for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
            {
                proof[i].generate_r1cs_constraints();
            }
        }

        for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
        {
            verifier[i].generate_r1cs_constraints();
        }
    }

    PROFILE_CONSTRAINTS(pb, "miscellaneous")
    {
        generate_r1cs_equals_const_constraint<FieldT>(pb, zero, FieldT::zero(), "zero");

        PROFILE_CONSTRAINTS(pb, "check that s proofs lie on the curve")
        {
            for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
            {
                generate_boolean_r1cs_constraint<FieldT>(pb, verification_results[i], FMT("", "verification_results_%zu", i));
            }
        }

        /* either type = 0 or proof verified w.r.t. a valid verification key */
        PROFILE_CONSTRAINTS(pb, "check that s messages have valid proofs (or are base case)")
        {
            for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
            {
                pb.add_r1cs_constraint(r1cs_constraint<FieldT>(incoming_message_types[i], 1 - verification_results[i], 0), FMT("", "not_base_case_implies_valid_proof_%zu", i));
            }
        }

        if (compliance_predicate.relies_on_same_type_inputs)
        {
            PROFILE_CONSTRAINTS(pb, "check that all non-base case messages are of same type and that VK is validly selected")
            {
                for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
                {
                    pb.add_r1cs_constraint(r1cs_constraint<FieldT>(incoming_message_types[i], incoming_message_types[i] - common_type, 0), FMT("", "non_base_types_equal_%zu", i));
                }

                pb.add_r1cs_constraint(r1cs_constraint<FieldT>(common_type, 1 - membership_check_results[0], 0), "valid_vk_for_the_common_type");

                auto it = compliance_predicate.accepted_input_types.begin();
                for (size_t i = 0; i < compliance_predicate.accepted_input_types.size(); ++i, ++it)
                {
                    pb.add_r1cs_constraint(r1cs_constraint<FieldT>((i == 0 ? common_type : common_type_check_aux[i-1]),
                                                                   common_type - FieldT(*it),
                                                                   (i == compliance_predicate.accepted_input_types.size() - 1 ? 0 * ONE : common_type_check_aux[i])),
                                           FMT("", "common_type_in_prescribed_set_%zu_must_equal_%zu", i, *it));
                }
            }
        }
        else
        {
            PROFILE_CONSTRAINTS(pb, "check that all s messages have validly selected VKs")
            {
                for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
                {
                    pb.add_r1cs_constraint(r1cs_constraint<FieldT>(incoming_message_types[i], 1 - membership_check_results[i], 0), FMT("", "not_base_case_implies_valid_vk_%zu", i));
                }
            }
        }
        pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, outgoing_message_type, FieldT(compliance_predicate.type)), "enforce_outgoing_type");
    }

    PRINT_CONSTRAINT_PROFILING();
    print_indent(); printf("* Number of constraints in mp_compliance_step_pcd_circuit: %zu\n", pb.num_constraints());
}

template<typename ppT>
r1cs_constraint_system<Fr<ppT> > mp_compliance_step_pcd_circuit_maker<ppT>::get_circuit() const
{
    return pb.get_constraint_system();
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > mp_compliance_step_pcd_circuit_maker<ppT>::get_primary_input() const
{
    return pb.primary_input();
}

template<typename ppT>
r1cs_auxiliary_input<Fr<ppT> > mp_compliance_step_pcd_circuit_maker<ppT>::get_auxiliary_input() const
{
    return pb.auxiliary_input();
}

template<typename ppT>
void mp_compliance_step_pcd_circuit_maker<ppT>::generate_r1cs_witness(const set_commitment &commitment_to_translation_step_r1cs_vks,
                                                                      const std::vector<r1cs_ppzksnark_verification_key<other_curve<ppT> > > &mp_translation_step_pcd_circuit_vks,
                                                                      const std::vector<set_membership_proof> &vk_membership_proofs,
                                                                      const r1cs_pcd_compliance_predicate_primary_input<FieldT> &compliance_predicate_primary_input,
                                                                      const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> &compliance_predicate_auxiliary_input,
                                                                      const std::vector<r1cs_ppzksnark_proof<other_curve<ppT> > > &translation_step_proofs)
{
    this->pb.clear_values();
    this->pb.val(zero) = FieldT::zero();

    compliance_predicate_as_gadget->generate_r1cs_witness(compliance_predicate_primary_input.as_r1cs_primary_input(),
                                                          compliance_predicate_auxiliary_input.as_r1cs_auxiliary_input(compliance_predicate.incoming_message_payload_lengths));

    unpack_outgoing_message->generate_r1cs_witness_from_packed();
    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        unpack_incoming_messages[i].generate_r1cs_witness_from_packed();
    }

    for (size_t i = 0; i < translation_step_vks.size(); ++i)
    {
        translation_step_vks[i].generate_r1cs_witness(mp_translation_step_pcd_circuit_vks[i]);
    }

    commitment->generate_r1cs_witness(commitment_to_translation_step_r1cs_vks);

    if (compliance_predicate.relies_on_same_type_inputs)
    {
        /* all messages (except base case) must be of the same type */
        this->pb.val(common_type) = FieldT::zero();
        size_t nonzero_type_idx = 0;
        for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
        {
            if (this->pb.val(incoming_message_types[i]) == 0)
            {
                continue;
            }

            if (this->pb.val(common_type).is_zero())
            {
                this->pb.val(common_type) = this->pb.val(incoming_message_types[i]);
                nonzero_type_idx = i;
            }
            else
            {
                assert(this->pb.val(common_type) == this->pb.val(incoming_message_types[i]));
            }
        }

        this->pb.val(membership_check_results[0]) = (this->pb.val(common_type).is_zero() ? FieldT::zero() : FieldT::one());
        membership_proofs[0].generate_r1cs_witness(vk_membership_proofs[nonzero_type_idx]);
        membership_checkers[0].generate_r1cs_witness();

        auto it = compliance_predicate.accepted_input_types.begin();
        for (size_t i = 0; i < compliance_predicate.accepted_input_types.size(); ++i, ++it)
        {
            pb.val(common_type_check_aux[i]) = ((i == 0 ? pb.val(common_type) : pb.val(common_type_check_aux[i-1])) *
                                                (pb.val(common_type) - FieldT(*it)));
        }
    }
    else
    {
        for (size_t i = 0; i < membership_checkers.size(); ++i)
        {
            this->pb.val(membership_check_results[i]) = (this->pb.val(incoming_message_types[i]).is_zero() ? FieldT::zero() : FieldT::one());
            membership_proofs[i].generate_r1cs_witness(vk_membership_proofs[i]);
            membership_checkers[i].generate_r1cs_witness();
        }
    }

    hash_outgoing_message->generate_r1cs_witness();
    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        hash_incoming_messages[i].generate_r1cs_witness();
        unpack_commitment_and_incoming_message_digests[i].generate_r1cs_witness_from_packed();
    }

    for (size_t i = 0; i < compliance_predicate.max_arity; ++i)
    {
        proof[i].generate_r1cs_witness(translation_step_proofs[i]);
        verifier[i].generate_r1cs_witness();
    }

#ifdef DEBUG
    get_circuit(); // force generating constraints
    assert(this->pb.is_satisfied());
#endif
}

template<typename ppT>
size_t mp_compliance_step_pcd_circuit_maker<ppT>::field_logsize()
{
    return Fr<ppT>::size_in_bits();
}

template<typename ppT>
size_t mp_compliance_step_pcd_circuit_maker<ppT>::field_capacity()
{
    return Fr<ppT>::capacity();
}

template<typename ppT>
size_t mp_compliance_step_pcd_circuit_maker<ppT>::input_size_in_elts()
{
    const size_t digest_size = CRH_with_field_out_gadget<FieldT>::get_digest_len();
    return digest_size;
}

template<typename ppT>
size_t mp_compliance_step_pcd_circuit_maker<ppT>::input_capacity_in_bits()
{
    return input_size_in_elts() * field_capacity();
}

template<typename ppT>
size_t mp_compliance_step_pcd_circuit_maker<ppT>::input_size_in_bits()
{
    return input_size_in_elts() * field_logsize();
}

template<typename ppT>
mp_translation_step_pcd_circuit_maker<ppT>::mp_translation_step_pcd_circuit_maker(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &compliance_step_vk)
{
    /* allocate input of the translation MP_PCD circuit */
    mp_translation_step_pcd_circuit_input.allocate(pb, input_size_in_elts(), "mp_translation_step_pcd_circuit_input");

    /* unpack translation step MP_PCD circuit input */
    unpacked_mp_translation_step_pcd_circuit_input.allocate(pb, mp_compliance_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_bits(), "unpacked_mp_translation_step_pcd_circuit_input");
    unpack_mp_translation_step_pcd_circuit_input.reset(new multipacking_gadget<FieldT>(pb, unpacked_mp_translation_step_pcd_circuit_input, mp_translation_step_pcd_circuit_input, field_capacity(), "unpack_mp_translation_step_pcd_circuit_input"));

    /* prepare arguments for the verifier */
    hardcoded_compliance_step_vk.reset(new r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT>(pb, compliance_step_vk, "hardcoded_compliance_step_vk"));
    proof.reset(new r1cs_ppzksnark_proof_variable<ppT>(pb, "proof"));

    /* verify previous proof */
    online_verifier.reset(new r1cs_ppzksnark_online_verifier_gadget<ppT>(pb,
                                                                         *hardcoded_compliance_step_vk,
                                                                         unpacked_mp_translation_step_pcd_circuit_input,
                                                                         mp_compliance_step_pcd_circuit_maker<other_curve<ppT> >::field_logsize(),
                                                                         *proof,
                                                                         ONE, // must always accept
                                                                         "verifier"));

    pb.set_input_sizes(input_size_in_elts());
}

template<typename ppT>
void mp_translation_step_pcd_circuit_maker<ppT>::generate_r1cs_constraints()
{
    PROFILE_CONSTRAINTS(pb, "repacking: unpack circuit input")
    {
        unpack_mp_translation_step_pcd_circuit_input->generate_r1cs_constraints(true);
    }

    PROFILE_CONSTRAINTS(pb, "verifier for compliance proofs")
    {
        PROFILE_CONSTRAINTS(pb, "check that proof lies on the curve")
        {
            proof->generate_r1cs_constraints();
        }

        online_verifier->generate_r1cs_constraints();
    }

    PRINT_CONSTRAINT_PROFILING();
    print_indent(); printf("* Number of constraints in mp_translation_step_pcd_circuit: %zu\n", pb.num_constraints());
}

template<typename ppT>
r1cs_constraint_system<Fr<ppT> > mp_translation_step_pcd_circuit_maker<ppT>::get_circuit() const
{
    return pb.get_constraint_system();
}

template<typename ppT>
void mp_translation_step_pcd_circuit_maker<ppT>::generate_r1cs_witness(const r1cs_primary_input<Fr<ppT> > translation_step_input,
                                                                       const r1cs_ppzksnark_proof<other_curve<ppT> > &prev_proof)
{
    this->pb.clear_values();
    mp_translation_step_pcd_circuit_input.fill_with_field_elements(pb, translation_step_input);
    unpack_mp_translation_step_pcd_circuit_input->generate_r1cs_witness_from_packed();

    proof->generate_r1cs_witness(prev_proof);
    online_verifier->generate_r1cs_witness();

#ifdef DEBUG
    get_circuit(); // force generating constraints
    assert(this->pb.is_satisfied());
#endif
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > mp_translation_step_pcd_circuit_maker<ppT>::get_primary_input() const
{
    return pb.primary_input();
}

template<typename ppT>
r1cs_auxiliary_input<Fr<ppT> > mp_translation_step_pcd_circuit_maker<ppT>::get_auxiliary_input() const
{
    return pb.auxiliary_input();
}

template<typename ppT>
size_t mp_translation_step_pcd_circuit_maker<ppT>::field_logsize()
{
    return Fr<ppT>::size_in_bits();
}

template<typename ppT>
size_t mp_translation_step_pcd_circuit_maker<ppT>::field_capacity()
{
    return Fr<ppT>::capacity();
}

template<typename ppT>
size_t mp_translation_step_pcd_circuit_maker<ppT>::input_size_in_elts()
{
    return div_ceil(mp_compliance_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_bits(), mp_translation_step_pcd_circuit_maker<ppT>::field_capacity());
}

template<typename ppT>
size_t mp_translation_step_pcd_circuit_maker<ppT>::input_capacity_in_bits()
{
    return input_size_in_elts() * field_capacity();
}

template<typename ppT>
size_t mp_translation_step_pcd_circuit_maker<ppT>::input_size_in_bits()
{
    return input_size_in_elts() * field_logsize();
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > get_mp_compliance_step_pcd_circuit_input(const set_commitment &commitment_to_translation_step_r1cs_vks,
                                                                      const r1cs_pcd_compliance_predicate_primary_input<Fr<ppT> > &primary_input)
{
    enter_block("Call to get_mp_compliance_step_pcd_circuit_input");
    typedef Fr<ppT> FieldT;

    const r1cs_variable_assignment<FieldT> outgoing_message_as_va = primary_input.outgoing_message->as_r1cs_variable_assignment();
    bit_vector msg_bits;
    for (const FieldT &elt : outgoing_message_as_va)
    {
        const bit_vector elt_bits = convert_field_element_to_bit_vector(elt);
        msg_bits.insert(msg_bits.end(), elt_bits.begin(), elt_bits.end());
    }

    bit_vector block;
    block.insert(block.end(), commitment_to_translation_step_r1cs_vks.begin(), commitment_to_translation_step_r1cs_vks.end());
    block.insert(block.end(), msg_bits.begin(), msg_bits.end());

    enter_block("Sample CRH randomness");
    CRH_with_field_out_gadget<FieldT>::sample_randomness(block.size());
    leave_block("Sample CRH randomness");

    const std::vector<FieldT> digest = CRH_with_field_out_gadget<FieldT>::get_hash(block);
    leave_block("Call to get_mp_compliance_step_pcd_circuit_input");

    return digest;
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > get_mp_translation_step_pcd_circuit_input(const set_commitment &commitment_to_translation_step_r1cs_vks,
                                                                       const r1cs_pcd_compliance_predicate_primary_input<Fr<other_curve<ppT> > > &primary_input)
{
    enter_block("Call to get_mp_translation_step_pcd_circuit_input");
    typedef Fr<ppT> FieldT;

    const std::vector<Fr<other_curve<ppT> > > mp_compliance_step_pcd_circuit_input = get_mp_compliance_step_pcd_circuit_input<other_curve<ppT> >(commitment_to_translation_step_r1cs_vks, primary_input);
    bit_vector mp_compliance_step_pcd_circuit_input_bits;
    for (const Fr<other_curve<ppT> > &elt : mp_compliance_step_pcd_circuit_input)
    {
        const bit_vector elt_bits = convert_field_element_to_bit_vector<Fr<other_curve<ppT> > >(elt);
        mp_compliance_step_pcd_circuit_input_bits.insert(mp_compliance_step_pcd_circuit_input_bits.end(), elt_bits.begin(), elt_bits.end());
    }

    mp_compliance_step_pcd_circuit_input_bits.resize(mp_translation_step_pcd_circuit_maker<ppT>::input_capacity_in_bits(), false);

    const r1cs_primary_input<FieldT> result = pack_bit_vector_into_field_element_vector<FieldT>(mp_compliance_step_pcd_circuit_input_bits, mp_translation_step_pcd_circuit_maker<ppT>::field_capacity());
    leave_block("Call to get_mp_translation_step_pcd_circuit_input");

    return result;
}

} // libsnark

#endif // MP_PCD_CIRCUITS_TCC_

/** @file
 *****************************************************************************

 Implementation of functionality for creating and using the two PCD circuits in
 a single-predicate PCD construction.

 See sp_pcd_circuits.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SP_PCD_CIRCUITS_TCC_
#define SP_PCD_CIRCUITS_TCC_

#include "common/utils.hpp"
#include "gadgetlib1/constraint_profiling.hpp"

namespace libsnark {

template<typename ppT>
sp_compliance_step_pcd_circuit_maker<ppT>::sp_compliance_step_pcd_circuit_maker(const r1cs_pcd_compliance_predicate<FieldT> &compliance_predicate) :
    compliance_predicate(compliance_predicate)
{
    /* calculate some useful sizes */
    assert(compliance_predicate.is_well_formed());
    assert(compliance_predicate.has_equal_input_and_output_lengths());

    const size_t compliance_predicate_arity = compliance_predicate.max_arity;
    const size_t digest_size = CRH_with_field_out_gadget<FieldT>::get_digest_len();
    const size_t msg_size_in_bits = field_logsize() * (1+compliance_predicate.outgoing_message_payload_length);
    const size_t sp_translation_step_vk_size_in_bits = r1cs_ppzksnark_verification_key_variable<ppT>::size_in_bits(sp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_elts());
    const size_t padded_verifier_input_size = sp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_capacity_in_bits();

    printf("other curve input size = %zu\n", sp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_elts());
    printf("translation_vk_bits = %zu\n", sp_translation_step_vk_size_in_bits);
    printf("padded verifier input size = %zu\n", padded_verifier_input_size);

    const size_t block_size = msg_size_in_bits + sp_translation_step_vk_size_in_bits;
    CRH_with_bit_out_gadget<FieldT>::sample_randomness(block_size);

    /* allocate input of the compliance PCD circuit */
    sp_compliance_step_pcd_circuit_input.allocate(pb, input_size_in_elts(), "sp_compliance_step_pcd_circuit_input");

    /* allocate inputs to the compliance predicate */
    outgoing_message_type.allocate(pb, "outgoing_message_type");
    outgoing_message_payload.allocate(pb, compliance_predicate.outgoing_message_payload_length, "outgoing_message_payload");

    outgoing_message_vars.insert(outgoing_message_vars.end(), outgoing_message_type);
    outgoing_message_vars.insert(outgoing_message_vars.end(), outgoing_message_payload.begin(), outgoing_message_payload.end());

    arity.allocate(pb, "arity");

    incoming_message_types.resize(compliance_predicate_arity);
    incoming_message_payloads.resize(compliance_predicate_arity);
    incoming_message_vars.resize(compliance_predicate_arity);
    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        incoming_message_types[i].allocate(pb, FMT("", "incoming_message_type_%zu", i));
        incoming_message_payloads[i].allocate(pb, compliance_predicate.outgoing_message_payload_length, FMT("", "incoming_message_payloads_%zu", i));

        incoming_message_vars[i].insert(incoming_message_vars[i].end(), incoming_message_types[i]);
        incoming_message_vars[i].insert(incoming_message_vars[i].end(), incoming_message_payloads[i].begin(), incoming_message_payloads[i].end());
    }

    local_data.allocate(pb, compliance_predicate.local_data_length, "local_data");
    cp_witness.allocate(pb, compliance_predicate.witness_length, "cp_witness");

    /* convert compliance predicate from a constraint system into a gadget */
    pb_variable_array<FieldT> incoming_messages_concat;
    for (size_t i = 0; i < compliance_predicate_arity; ++i)
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
    outgoing_message_bits.allocate(pb, msg_size_in_bits, "outgoing_message_bits");
    unpack_outgoing_message.reset(new multipacking_gadget<FieldT>(pb, outgoing_message_bits, outgoing_message_vars, field_logsize(), "unpack_outgoing_message"));

    incoming_messages_bits.resize(compliance_predicate_arity);
    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        incoming_messages_bits[i].allocate(pb, msg_size_in_bits, FMT("", "incoming_messages_bits_%zu", i));
        unpack_incoming_messages.emplace_back(multipacking_gadget<FieldT>(pb, incoming_messages_bits[i], incoming_message_vars[i], field_logsize(), FMT("", "unpack_incoming_messages_%zu", i)));
    }

    /* allocate digests */
    sp_translation_step_vk_and_incoming_message_payload_digests.resize(compliance_predicate_arity);
    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        sp_translation_step_vk_and_incoming_message_payload_digests[i].allocate(pb, digest_size, FMT("", "sp_translation_step_vk_and_incoming_message_payload_digests_%zu", i));
    }

    /* allocate blocks */
    sp_translation_step_vk_bits.allocate(pb, sp_translation_step_vk_size_in_bits, "sp_translation_step_vk_bits");

    block_for_outgoing_message.reset(new block_variable<FieldT>(pb, {
                sp_translation_step_vk_bits,
                outgoing_message_bits }, "block_for_outgoing_message"));

    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        blocks_for_incoming_messages.emplace_back(block_variable<FieldT>(pb, {
                    sp_translation_step_vk_bits,
                    incoming_messages_bits[i] }, FMT("", "blocks_for_incoming_messages_zu", i)));
    }

    /* allocate hash checkers */
    hash_outgoing_message.reset(new CRH_with_field_out_gadget<FieldT>(pb, block_size, *block_for_outgoing_message, sp_compliance_step_pcd_circuit_input, "hash_outgoing_message"));

    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        hash_incoming_messages.emplace_back(CRH_with_field_out_gadget<FieldT>(pb, block_size, blocks_for_incoming_messages[i], sp_translation_step_vk_and_incoming_message_payload_digests[i], FMT("", "hash_incoming_messages_%zu", i)));
    }

    /* allocate useful zero variable */
    zero.allocate(pb, "zero");

    /* prepare arguments for the verifier */
    sp_translation_step_vk.reset(new r1cs_ppzksnark_verification_key_variable<ppT>(pb, sp_translation_step_vk_bits, sp_translation_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_elts(), "sp_translation_step_vk"));

    verification_result.allocate(pb, "verification_result");
    sp_translation_step_vk_and_incoming_message_payload_digest_bits.resize(compliance_predicate_arity);

    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        sp_translation_step_vk_and_incoming_message_payload_digest_bits[i].allocate(pb, digest_size * field_logsize(), FMT("", "sp_translation_step_vk_and_incoming_message_payload_digest_bits_%zu", i));
        unpack_sp_translation_step_vk_and_incoming_message_payload_digests.emplace_back(multipacking_gadget<FieldT>(pb,
                                                                                                            sp_translation_step_vk_and_incoming_message_payload_digest_bits[i],
                                                                                                            sp_translation_step_vk_and_incoming_message_payload_digests[i],
                                                                                                            field_logsize(),
                                                                                                            FMT("", "unpack_sp_translation_step_vk_and_incoming_message_payload_digests_%zu", i)));

        verifier_input.emplace_back(sp_translation_step_vk_and_incoming_message_payload_digest_bits[i]);
        while (verifier_input[i].size() < padded_verifier_input_size)
        {
            verifier_input[i].emplace_back(zero);
        }

        proof.emplace_back(r1cs_ppzksnark_proof_variable<ppT>(pb, FMT("", "proof_%zu", i)));
        verifiers.emplace_back(r1cs_ppzksnark_verifier_gadget<ppT>(pb,
                                                    *sp_translation_step_vk,
                                                    verifier_input[i],
                                                    sp_translation_step_pcd_circuit_maker<other_curve<ppT> >::field_capacity(),
                                                    proof[i],
                                                    verification_result,
                                                    FMT("", "verifiers_%zu", i)));
    }

    pb.set_input_sizes(input_size_in_elts());
    printf("done compliance\n");
}

template<typename ppT>
void sp_compliance_step_pcd_circuit_maker<ppT>::generate_r1cs_constraints()
{
    const size_t digest_size = CRH_with_bit_out_gadget<FieldT>::get_digest_len();
    const size_t dimension = knapsack_dimension<FieldT>::dimension;
    print_indent(); printf("* Knapsack dimension: %zu\n", dimension);

    const size_t compliance_predicate_arity = compliance_predicate.max_arity;
    print_indent(); printf("* Compliance predicate arity: %zu\n", compliance_predicate_arity);
    print_indent(); printf("* Compliance predicate payload length: %zu\n", compliance_predicate.outgoing_message_payload_length);
    print_indent(); printf("* Compliance predicate local data length: %zu\n", compliance_predicate.local_data_length);
    print_indent(); printf("* Compliance predicate witness length: %zu\n", compliance_predicate.witness_length);

    PROFILE_CONSTRAINTS(pb, "booleanity")
    {
        PROFILE_CONSTRAINTS(pb, "booleanity: unpack outgoing_message")
        {
            unpack_outgoing_message->generate_r1cs_constraints(true);
        }

        PROFILE_CONSTRAINTS(pb, "booleanity: unpack s incoming_message")
        {
            for (size_t i = 0; i < compliance_predicate_arity; ++i)
            {
                unpack_incoming_messages[i].generate_r1cs_constraints(true);
            }
        }

        PROFILE_CONSTRAINTS(pb, "booleanity: unpack verification key")
        {
            sp_translation_step_vk->generate_r1cs_constraints(true);
        }
    }

    PROFILE_CONSTRAINTS(pb, "(1+s) copies of hash")
    {
        print_indent(); printf("* Digest-size: %zu\n", digest_size);
        hash_outgoing_message->generate_r1cs_constraints();

        for (size_t i = 0; i < compliance_predicate_arity; ++i)
        {
            hash_incoming_messages[i].generate_r1cs_constraints();
        }
    }

    PROFILE_CONSTRAINTS(pb, "s copies of repacking circuit")
    {
        for (size_t i = 0; i < compliance_predicate_arity; ++i)
        {
            unpack_sp_translation_step_vk_and_incoming_message_payload_digests[i].generate_r1cs_constraints(true);
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
            for (size_t i = 0; i < compliance_predicate_arity; ++i)
            {
                proof[i].generate_r1cs_constraints();
            }
        }

        for (size_t i = 0; i < compliance_predicate_arity; ++i)
        {
            verifiers[i].generate_r1cs_constraints();
        }
    }

    PROFILE_CONSTRAINTS(pb, "miscellaneous")
    {
        generate_r1cs_equals_const_constraint<FieldT>(pb, zero, FieldT::zero(), "zero");
        generate_boolean_r1cs_constraint<FieldT>(pb, verification_result, "verification_result");

        /* type * (1-verification_result) = 0 */
        pb.add_r1cs_constraint(r1cs_constraint<FieldT>(incoming_message_types[0], 1 - verification_result, 0), "not_base_case_implies_valid_proofs");

        /* all types equal */
        for (size_t i = 1; i < compliance_predicate.max_arity; ++i)
        {
            pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, incoming_message_types[0], incoming_message_types[i]),
                                   FMT("", "type_%zu_equal_to_type_0", i));
        }

        pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, arity, compliance_predicate_arity), "full_arity");
        pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, outgoing_message_type, FieldT(compliance_predicate.type)), "enforce_outgoing_type");
    }

    PRINT_CONSTRAINT_PROFILING();
    print_indent(); printf("* Number of constraints in sp_compliance_step_pcd_circuit: %zu\n", pb.num_constraints());
}

template<typename ppT>
r1cs_constraint_system<Fr<ppT> > sp_compliance_step_pcd_circuit_maker<ppT>::get_circuit() const
{
    return pb.get_constraint_system();
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > sp_compliance_step_pcd_circuit_maker<ppT>::get_primary_input() const
{
    return pb.primary_input();
}

template<typename ppT>
r1cs_auxiliary_input<Fr<ppT> > sp_compliance_step_pcd_circuit_maker<ppT>::get_auxiliary_input() const
{
    return pb.auxiliary_input();
}

template<typename ppT>
void sp_compliance_step_pcd_circuit_maker<ppT>::generate_r1cs_witness(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &sp_translation_step_pcd_circuit_vk,
                                                                      const r1cs_pcd_compliance_predicate_primary_input<FieldT> &compliance_predicate_primary_input,
                                                                      const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> &compliance_predicate_auxiliary_input,
                                                                      const std::vector<r1cs_ppzksnark_proof<other_curve<ppT> > > &incoming_proofs)
{
    const size_t compliance_predicate_arity = compliance_predicate.max_arity;
    this->pb.clear_values();
    this->pb.val(zero) = FieldT::zero();

    compliance_predicate_as_gadget->generate_r1cs_witness(compliance_predicate_primary_input.as_r1cs_primary_input(),
                                                          compliance_predicate_auxiliary_input.as_r1cs_auxiliary_input(compliance_predicate.incoming_message_payload_lengths));
    this->pb.val(arity) = FieldT(compliance_predicate_arity);
    unpack_outgoing_message->generate_r1cs_witness_from_packed();
    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        unpack_incoming_messages[i].generate_r1cs_witness_from_packed();
    }

    sp_translation_step_vk->generate_r1cs_witness(sp_translation_step_pcd_circuit_vk);
    hash_outgoing_message->generate_r1cs_witness();
    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        hash_incoming_messages[i].generate_r1cs_witness();
        unpack_sp_translation_step_vk_and_incoming_message_payload_digests[i].generate_r1cs_witness_from_packed();
    }

    for (size_t i = 0; i < compliance_predicate_arity; ++i)
    {
        proof[i].generate_r1cs_witness(incoming_proofs[i]);
        verifiers[i].generate_r1cs_witness();
    }

    if (this->pb.val(incoming_message_types[0]) != FieldT::zero())
    {
        this->pb.val(verification_result) = FieldT::one();
    }

#ifdef DEBUG
    generate_r1cs_constraints(); // force generating constraints
    assert(this->pb.is_satisfied());
#endif
}

template<typename ppT>
size_t sp_compliance_step_pcd_circuit_maker<ppT>::field_logsize()
{
    return Fr<ppT>::size_in_bits();
}

template<typename ppT>
size_t sp_compliance_step_pcd_circuit_maker<ppT>::field_capacity()
{
    return Fr<ppT>::capacity();
}

template<typename ppT>
size_t sp_compliance_step_pcd_circuit_maker<ppT>::input_size_in_elts()
{
    const size_t digest_size = CRH_with_field_out_gadget<FieldT>::get_digest_len();
    return digest_size;
}

template<typename ppT>
size_t sp_compliance_step_pcd_circuit_maker<ppT>::input_capacity_in_bits()
{
    return input_size_in_elts() * field_capacity();
}

template<typename ppT>
size_t sp_compliance_step_pcd_circuit_maker<ppT>::input_size_in_bits()
{
    return input_size_in_elts() * field_logsize();
}

template<typename ppT>
sp_translation_step_pcd_circuit_maker<ppT>::sp_translation_step_pcd_circuit_maker(const r1cs_ppzksnark_verification_key<other_curve<ppT> > &sp_compliance_step_vk)
{
    /* allocate input of the translation PCD circuit */
    sp_translation_step_pcd_circuit_input.allocate(pb, input_size_in_elts(), "sp_translation_step_pcd_circuit_input");

    /* unpack translation step PCD circuit input */
    unpacked_sp_translation_step_pcd_circuit_input.allocate(pb, sp_compliance_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_bits(), "unpacked_sp_translation_step_pcd_circuit_input");
    unpack_sp_translation_step_pcd_circuit_input.reset(new multipacking_gadget<FieldT>(pb, unpacked_sp_translation_step_pcd_circuit_input, sp_translation_step_pcd_circuit_input, field_capacity(), "unpack_sp_translation_step_pcd_circuit_input"));

    /* prepare arguments for the verifier */
    hardcoded_sp_compliance_step_vk.reset(new r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT>(pb, sp_compliance_step_vk, "hardcoded_sp_compliance_step_vk"));
    proof.reset(new r1cs_ppzksnark_proof_variable<ppT>(pb, "proof"));

    /* verify previous proof */
    online_verifier.reset(new r1cs_ppzksnark_online_verifier_gadget<ppT>(pb,
                                                          *hardcoded_sp_compliance_step_vk,
                                                          unpacked_sp_translation_step_pcd_circuit_input,
                                                          sp_compliance_step_pcd_circuit_maker<other_curve<ppT> >::field_logsize(),
                                                          *proof,
                                                          ONE, // must always accept
                                                          "verifier"));
    pb.set_input_sizes(input_size_in_elts());

    printf("done translation\n");
}

template<typename ppT>
void sp_translation_step_pcd_circuit_maker<ppT>::generate_r1cs_constraints()
{
    PROFILE_CONSTRAINTS(pb, "repacking: unpack circuit input")
    {
        unpack_sp_translation_step_pcd_circuit_input->generate_r1cs_constraints(true);
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
    print_indent(); printf("* Number of constraints in sp_translation_step_pcd_circuit: %zu\n", pb.num_constraints());
}

template<typename ppT>
r1cs_constraint_system<Fr<ppT> > sp_translation_step_pcd_circuit_maker<ppT>::get_circuit() const
{
    return pb.get_constraint_system();
}

template<typename ppT>
void sp_translation_step_pcd_circuit_maker<ppT>::generate_r1cs_witness(const r1cs_primary_input<Fr<ppT> > sp_translation_step_input,
                                                                       const r1cs_ppzksnark_proof<other_curve<ppT> > &compliance_step_proof)
{
    this->pb.clear_values();
    sp_translation_step_pcd_circuit_input.fill_with_field_elements(pb, sp_translation_step_input);
    unpack_sp_translation_step_pcd_circuit_input->generate_r1cs_witness_from_packed();

    proof->generate_r1cs_witness(compliance_step_proof);
    online_verifier->generate_r1cs_witness();

#ifdef DEBUG
    generate_r1cs_constraints(); // force generating constraints

    printf("Input to the translation circuit:\n");
    for (size_t i = 0; i < this->pb.num_inputs(); ++i)
    {
        this->pb.val(pb_variable<FieldT>(i+1)).print();
    }

    assert(this->pb.is_satisfied());
#endif
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > sp_translation_step_pcd_circuit_maker<ppT>::get_primary_input() const
{
    return pb.primary_input();
}

template<typename ppT>
r1cs_auxiliary_input<Fr<ppT> > sp_translation_step_pcd_circuit_maker<ppT>::get_auxiliary_input() const
{
    return pb.auxiliary_input();
}

template<typename ppT>
size_t sp_translation_step_pcd_circuit_maker<ppT>::field_logsize()
{
    return Fr<ppT>::size_in_bits();
}

template<typename ppT>
size_t sp_translation_step_pcd_circuit_maker<ppT>::field_capacity()
{
    return Fr<ppT>::capacity();
}

template<typename ppT>
size_t sp_translation_step_pcd_circuit_maker<ppT>::input_size_in_elts()
{
    return div_ceil(sp_compliance_step_pcd_circuit_maker<other_curve<ppT> >::input_size_in_bits(), sp_translation_step_pcd_circuit_maker<ppT>::field_capacity());
}

template<typename ppT>
size_t sp_translation_step_pcd_circuit_maker<ppT>::input_capacity_in_bits()
{
    return input_size_in_elts() * field_capacity();
}

template<typename ppT>
size_t sp_translation_step_pcd_circuit_maker<ppT>::input_size_in_bits()
{
    return input_size_in_elts() * field_logsize();
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > get_sp_compliance_step_pcd_circuit_input(const bit_vector &sp_translation_step_vk_bits,
                                                                      const r1cs_pcd_compliance_predicate_primary_input<Fr<ppT> > &primary_input)
{
    enter_block("Call to get_sp_compliance_step_pcd_circuit_input");
    typedef Fr<ppT> FieldT;

    const r1cs_variable_assignment<FieldT> outgoing_message_as_va = primary_input.outgoing_message->as_r1cs_variable_assignment();
    bit_vector msg_bits;
    for (const FieldT &elt : outgoing_message_as_va)
    {
        const bit_vector elt_bits = convert_field_element_to_bit_vector(elt);
        msg_bits.insert(msg_bits.end(), elt_bits.begin(), elt_bits.end());
    }

    bit_vector block;
    block.insert(block.end(), sp_translation_step_vk_bits.begin(), sp_translation_step_vk_bits.end());
    block.insert(block.end(), msg_bits.begin(), msg_bits.end());

    enter_block("Sample CRH randomness");
    CRH_with_field_out_gadget<FieldT>::sample_randomness(block.size());
    leave_block("Sample CRH randomness");

    const std::vector<FieldT> digest = CRH_with_field_out_gadget<FieldT>::get_hash(block);
    leave_block("Call to get_sp_compliance_step_pcd_circuit_input");

    return digest;
}

template<typename ppT>
r1cs_primary_input<Fr<ppT> > get_sp_translation_step_pcd_circuit_input(const bit_vector &sp_translation_step_vk_bits,
                                                                       const r1cs_pcd_compliance_predicate_primary_input<Fr<other_curve<ppT> > > &primary_input)
{
    enter_block("Call to get_sp_translation_step_pcd_circuit_input");
    typedef Fr<ppT> FieldT;

    const std::vector<Fr<other_curve<ppT> > > sp_compliance_step_pcd_circuit_input = get_sp_compliance_step_pcd_circuit_input<other_curve<ppT> >(sp_translation_step_vk_bits, primary_input);
    bit_vector sp_compliance_step_pcd_circuit_input_bits;
    for (const Fr<other_curve<ppT> > &elt : sp_compliance_step_pcd_circuit_input)
    {
        const bit_vector elt_bits = convert_field_element_to_bit_vector<Fr<other_curve<ppT> > >(elt);
        sp_compliance_step_pcd_circuit_input_bits.insert(sp_compliance_step_pcd_circuit_input_bits.end(), elt_bits.begin(), elt_bits.end());
    }

    sp_compliance_step_pcd_circuit_input_bits.resize(sp_translation_step_pcd_circuit_maker<ppT>::input_capacity_in_bits(), false);

    const r1cs_primary_input<FieldT> result = pack_bit_vector_into_field_element_vector<FieldT>(sp_compliance_step_pcd_circuit_input_bits, sp_translation_step_pcd_circuit_maker<ppT>::field_capacity());
    leave_block("Call to get_sp_translation_step_pcd_circuit_input");

    return result;
}

} // libsnark

#endif // SP_PCD_CIRCUITS_TCC_

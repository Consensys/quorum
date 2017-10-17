/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef R1CS_PCD_PARAMS_TCC_
#define R1CS_PCD_PARAMS_TCC_

namespace libsnark {

template<typename FieldT>
r1cs_primary_input<FieldT> r1cs_pcd_compliance_predicate_primary_input<FieldT>::as_r1cs_primary_input() const
{
    return outgoing_message->as_r1cs_variable_assignment();
}

template<typename FieldT>
r1cs_auxiliary_input<FieldT> r1cs_pcd_compliance_predicate_auxiliary_input<FieldT>::as_r1cs_auxiliary_input(const std::vector<size_t> &incoming_message_payload_lengths) const
{
    const size_t arity = incoming_messages.size();

    r1cs_auxiliary_input<FieldT> result;
    result.emplace_back(FieldT(arity));

    const size_t max_arity = incoming_message_payload_lengths.size();
    assert(arity <= max_arity);

    for (size_t i = 0; i < arity; ++i)
    {
        const r1cs_variable_assignment<FieldT> msg_as_r1cs_va = incoming_messages[i]->as_r1cs_variable_assignment();
        assert(msg_as_r1cs_va.size() == (1 + incoming_message_payload_lengths[i]));
        result.insert(result.end(), msg_as_r1cs_va.begin(), msg_as_r1cs_va.end());
    }

    /* pad with dummy messages of appropriate size */
    for (size_t i = arity; i < max_arity; ++i)
    {
        result.resize(result.size() + (1 + incoming_message_payload_lengths[i]), FieldT::zero());
    }

    const r1cs_variable_assignment<FieldT> local_data_as_r1cs_va = local_data->as_r1cs_variable_assignment();
    result.insert(result.end(), local_data_as_r1cs_va.begin(), local_data_as_r1cs_va.end());
    result.insert(result.end(), witness.begin(), witness.end());

    return result;
}

} // libsnark

#endif // R1CS_PCD_PARAMS_TCC_

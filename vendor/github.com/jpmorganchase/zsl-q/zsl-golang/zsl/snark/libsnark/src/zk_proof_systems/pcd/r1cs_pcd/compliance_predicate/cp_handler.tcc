/** @file
 *****************************************************************************

 Implementation of interfaces for a compliance predicate handler.

 See cp_handler.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef CP_HANDLER_TCC_
#define CP_HANDLER_TCC_

#include <algorithm>

namespace libsnark {

template<typename FieldT>
r1cs_pcd_message_variable<FieldT>::r1cs_pcd_message_variable(protoboard<FieldT> &pb,
                                                             const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix)
{
    type.allocate(pb, FMT(annotation_prefix, " type"));
    all_vars.emplace_back(type);

    num_vars_at_construction = pb.num_variables();
}

template<typename FieldT>
void r1cs_pcd_message_variable<FieldT>::update_all_vars()
{
    /* NOTE: this assumes that r1cs_pcd_message_variable has been the
     * only gadget allocating variables on the protoboard and needs to
     * be updated, e.g., in multicore variable allocation scenario. */

    for (size_t var_idx = num_vars_at_construction + 1; var_idx <= this->pb.num_variables(); ++var_idx)
    {
        all_vars.emplace_back(pb_variable<FieldT>(var_idx));
    }
}

template<typename FieldT>
void r1cs_pcd_message_variable<FieldT>::generate_r1cs_witness(const std::shared_ptr<r1cs_pcd_message<FieldT> > &message)
{
    all_vars.fill_with_field_elements(this->pb, message->as_r1cs_variable_assignment());
}

template<typename FieldT>
r1cs_pcd_local_data_variable<FieldT>::r1cs_pcd_local_data_variable(protoboard<FieldT> &pb,
                                                                   const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix)
{
    num_vars_at_construction = pb.num_variables();
}

template<typename FieldT>
void r1cs_pcd_local_data_variable<FieldT>::update_all_vars()
{
    /* (the same NOTE as for r1cs_message_variable applies) */

    for (size_t var_idx = num_vars_at_construction + 1; var_idx <= this->pb.num_variables(); ++var_idx)
    {
        all_vars.emplace_back(pb_variable<FieldT>(var_idx));
    }
}

template<typename FieldT>
void r1cs_pcd_local_data_variable<FieldT>::generate_r1cs_witness(const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data)
{
    all_vars.fill_with_field_elements(this->pb, local_data->as_r1cs_variable_assignment());
}

template<typename FieldT, typename protoboardT>
compliance_predicate_handler<FieldT, protoboardT>::compliance_predicate_handler(const protoboardT &pb,
                                                                                const size_t name,
                                                                                const size_t type,
                                                                                const size_t max_arity,
                                                                                const bool relies_on_same_type_inputs,
                                                                                const std::set<size_t> accepted_input_types) :
    pb(pb), name(name), type(type), max_arity(max_arity), relies_on_same_type_inputs(relies_on_same_type_inputs),
    accepted_input_types(accepted_input_types)
{
    incoming_messages.resize(max_arity);
}

template<typename FieldT, typename protoboardT>
void compliance_predicate_handler<FieldT, protoboardT>::generate_r1cs_witness(const std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > &incoming_message_values,
                                                                              const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data_value)
{
    pb.clear_values();
    pb.val(outgoing_message->type) = FieldT(type);
    pb.val(arity) = FieldT(incoming_message_values.size());

    for (size_t i = 0; i < incoming_message_values.size(); ++i)
    {
        incoming_messages[i]->generate_r1cs_witness(incoming_message_values[i]);
    }

    local_data->generate_r1cs_witness(local_data_value);
}


template<typename FieldT, typename protoboardT>
r1cs_pcd_compliance_predicate<FieldT> compliance_predicate_handler<FieldT, protoboardT>::get_compliance_predicate() const
{
    assert(incoming_messages.size() == max_arity);

    const size_t outgoing_message_payload_length = outgoing_message->all_vars.size() - 1;

    std::vector<size_t> incoming_message_payload_lengths(max_arity);
    std::transform(incoming_messages.begin(), incoming_messages.end(),
                   incoming_message_payload_lengths.begin(),
                   [] (const std::shared_ptr<r1cs_pcd_message_variable<FieldT> > &msg) { return msg->all_vars.size() - 1; });

    const size_t local_data_length = local_data->all_vars.size();

    const size_t all_but_witness_length = ((1 + outgoing_message_payload_length) + 1 +
                                           (max_arity + std::accumulate(incoming_message_payload_lengths.begin(),
                                                                        incoming_message_payload_lengths.end(), 0)) +
                                           local_data_length);
    const size_t witness_length = pb.num_variables() - all_but_witness_length;

    r1cs_constraint_system<FieldT> constraint_system = pb.get_constraint_system();
    constraint_system.primary_input_size = 1 + outgoing_message_payload_length;
    constraint_system.auxiliary_input_size = pb.num_variables() - constraint_system.primary_input_size;

    return r1cs_pcd_compliance_predicate<FieldT>(name,
                                                 type,
                                                 constraint_system,
                                                 outgoing_message_payload_length,
                                                 max_arity,
                                                 incoming_message_payload_lengths,
                                                 local_data_length,
                                                 witness_length,
                                                 relies_on_same_type_inputs,
                                                 accepted_input_types);
}

template<typename FieldT, typename protoboardT>
r1cs_variable_assignment<FieldT> compliance_predicate_handler<FieldT, protoboardT>::get_full_variable_assignment() const
{
    return pb.full_variable_assignment();
}

template<typename FieldT, typename protoboardT>
std::shared_ptr<r1cs_pcd_message<FieldT> > compliance_predicate_handler<FieldT, protoboardT>::get_outgoing_message() const
{
    return outgoing_message->get_message();
}

template<typename FieldT, typename protoboardT>
size_t compliance_predicate_handler<FieldT, protoboardT>::get_arity() const
{
    return pb.val(arity).as_ulong();
}

template<typename FieldT, typename protoboardT>
std::shared_ptr<r1cs_pcd_message<FieldT> > compliance_predicate_handler<FieldT, protoboardT>::get_incoming_message(const size_t message_idx) const
{
    assert(message_idx < max_arity);
    return incoming_messages[message_idx]->get_message();
}

template<typename FieldT, typename protoboardT>
std::shared_ptr<r1cs_pcd_local_data<FieldT> > compliance_predicate_handler<FieldT, protoboardT>::get_local_data() const
{
    return local_data->get_local_data();
}

template<typename FieldT, typename protoboardT>
r1cs_pcd_witness<FieldT> compliance_predicate_handler<FieldT, protoboardT>::get_witness() const
{
    const r1cs_variable_assignment<FieldT> va = pb.full_variable_assignment();
    // outgoing_message + arity + incoming_messages + local_data
    const size_t witness_pos = (outgoing_message->all_vars.size() + 1 +
                                std::accumulate(incoming_messages.begin(), incoming_messages.end(),
                                                0, [](size_t acc, const std::shared_ptr<r1cs_pcd_message_variable<FieldT> > &msg) {
                                                    return acc + msg->all_vars.size(); }) +
                                local_data->all_vars.size());

    return r1cs_variable_assignment<FieldT>(va.begin() + witness_pos, va.end());
}

} // libsnark

#endif // CP_HANDLER_TCC_

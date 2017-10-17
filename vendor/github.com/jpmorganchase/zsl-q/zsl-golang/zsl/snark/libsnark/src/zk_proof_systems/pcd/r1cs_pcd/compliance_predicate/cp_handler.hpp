/** @file
 *****************************************************************************

 Declaration of interfaces for a compliance predicate handler.

 A compliance predicate handler is a base class for creating compliance predicates.
 It relies on classes declared in gadgetlib1.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef CP_HANDLER_HPP_
#define CP_HANDLER_HPP_

#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/compliance_predicate.hpp"
#include "gadgetlib1/protoboard.hpp"

namespace libsnark {

/***************************** Message variable ******************************/

/**
 * A variable to represent an r1cs_pcd_message.
 */
template<typename FieldT>
class r1cs_pcd_message_variable : public gadget<FieldT> {
protected:
    size_t num_vars_at_construction;
public:

    pb_variable<FieldT> type;

    pb_variable_array<FieldT> all_vars;

    r1cs_pcd_message_variable(protoboard<FieldT> &pb,
                              const std::string &annotation_prefix);
    void update_all_vars();

    void generate_r1cs_witness(const std::shared_ptr<r1cs_pcd_message<FieldT> > &message);
    virtual std::shared_ptr<r1cs_pcd_message<FieldT> > get_message() const = 0;

    virtual ~r1cs_pcd_message_variable() = default;
};
/*************************** Local data variable *****************************/

/**
 * A variable to represent an r1cs_pcd_local_data.
 */
template<typename FieldT>
class r1cs_pcd_local_data_variable : public gadget<FieldT> {
protected:
    size_t num_vars_at_construction;
public:

    pb_variable_array<FieldT> all_vars;

    r1cs_pcd_local_data_variable(protoboard<FieldT> &pb,
                                 const std::string &annotation_prefix);
    void update_all_vars();

    void generate_r1cs_witness(const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data);

    virtual ~r1cs_pcd_local_data_variable() = default;
};

/*********************** Compliance predicate handler ************************/

/**
 * A base class for creating compliance predicates.
 */
template<typename FieldT, typename protoboardT>
class compliance_predicate_handler {
protected:
    protoboardT pb;

    std::shared_ptr<r1cs_pcd_message_variable<FieldT> > outgoing_message;
    pb_variable<FieldT> arity;
    std::vector<std::shared_ptr<r1cs_pcd_message_variable<FieldT> > > incoming_messages;
    std::shared_ptr<r1cs_pcd_local_data_variable<FieldT> > local_data;
public:
    const size_t name;
    const size_t type;
    const size_t max_arity;
    const bool relies_on_same_type_inputs;
    const std::set<size_t> accepted_input_types;

    compliance_predicate_handler(const protoboardT &pb,
                                 const size_t name,
                                 const size_t type,
                                 const size_t max_arity,
                                 const bool relies_on_same_type_inputs,
                                 const std::set<size_t> accepted_input_types = std::set<size_t>());
    virtual void generate_r1cs_constraints() = 0;
    virtual void generate_r1cs_witness(const std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > &incoming_message_values,
                                       const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data_value);

    r1cs_pcd_compliance_predicate<FieldT> get_compliance_predicate() const;
    r1cs_variable_assignment<FieldT> get_full_variable_assignment() const;

    std::shared_ptr<r1cs_pcd_message<FieldT> > get_outgoing_message() const;
    size_t get_arity() const;
    std::shared_ptr<r1cs_pcd_message<FieldT> > get_incoming_message(const size_t message_idx) const;
    std::shared_ptr<r1cs_pcd_local_data<FieldT> > get_local_data() const;
    r1cs_variable_assignment<FieldT> get_witness() const;
};

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/cp_handler.tcc"

#endif // CP_HANDLER_HPP_

/** @file
 *****************************************************************************

 Declaration of interfaces for a compliance predicate for R1CS PCD.

 A compliance predicate specifies a local invariant to be enforced, by PCD,
 throughout a dynamic distributed computation. A compliance predicate
 receives input messages, local data, and an output message (and perhaps some
 other auxiliary information), and then either accepts or rejects.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef COMPLIANCE_PREDICATE_HPP_
#define COMPLIANCE_PREDICATE_HPP_

#include "relations/constraint_satisfaction_problems/r1cs/r1cs.hpp"

namespace libsnark {

/********************************* Message ***********************************/

/**
 * A message for R1CS PCD.
 *
 * It is a pair, consisting of
 * - a type (a positive integer), and
 * - a payload (a vector of field elements).
 */
template<typename FieldT>
class r1cs_pcd_message {
public:
    size_t type;

    r1cs_pcd_message(const size_t type);
    virtual r1cs_variable_assignment<FieldT> payload_as_r1cs_variable_assignment() const = 0;
    r1cs_variable_assignment<FieldT> as_r1cs_variable_assignment() const;

    virtual void print() const;
    virtual ~r1cs_pcd_message() = default;
};

/******************************* Local data **********************************/

/**
 * A local data for R1CS PCD.
 */
template<typename FieldT>
class r1cs_pcd_local_data {
public:
    r1cs_pcd_local_data() = default;
    virtual r1cs_variable_assignment<FieldT> as_r1cs_variable_assignment() const = 0;
    virtual ~r1cs_pcd_local_data() = default;
};

/******************************** Witness ************************************/

template<typename FieldT>
using r1cs_pcd_witness = std::vector<FieldT>;

/*************************** Compliance predicate ****************************/

template<typename FieldT>
class r1cs_pcd_compliance_predicate;

template<typename FieldT>
std::ostream& operator<<(std::ostream &out, const r1cs_pcd_compliance_predicate<FieldT> &cp);

template<typename FieldT>
std::istream& operator>>(std::istream &in, r1cs_pcd_compliance_predicate<FieldT> &cp);

/**
 * A compliance predicate for R1CS PCD.
 *
 * It is a wrapper around R1CS that also specifies how to parse a
 * variable assignment as:
 * - output message (the input)
 * - some number of input messages (part of the witness)
 * - local data (also part of the witness)
 * - auxiliary information (the remaining variables of the witness)
 *
 * A compliance predicate also has a type, allegedly the same
 * as the type of the output message.
 *
 * The input wires of R1CS appear in the following order:
 * - (1 + outgoing_message_payload_length) wires for outgoing message
 * - 1 wire for arity (allegedly, 0 <= arity <= max_arity)
 * - for i = 0, ..., max_arity-1:
 * - (1 + incoming_message_payload_lengths[i]) wires for i-th message of
 *   the input (in the array that's padded to max_arity messages)
 * - local_data_length wires for local data
 *
 * The rest witness_length wires of the R1CS constitute the witness.
 *
 * To allow for optimizations, the compliance predicate also
 * specififies a flag, called relies_on_same_type_inputs, denoting
 * whether the predicate works under the assumption that all input
 * messages have the same type. In such case a member
 * accepted_input_types lists all types accepted by the predicate
 * (accepted_input_types has no meaning if
 * relies_on_same_type_inputs=false).
 */

template<typename FieldT>
class r1cs_pcd_compliance_predicate {
public:

    size_t name;
    size_t type;

    r1cs_constraint_system<FieldT> constraint_system;

    size_t outgoing_message_payload_length;
    size_t max_arity;
    std::vector<size_t> incoming_message_payload_lengths;
    size_t local_data_length;
    size_t witness_length;

    bool relies_on_same_type_inputs;
    std::set<size_t> accepted_input_types;

    r1cs_pcd_compliance_predicate() = default;
    r1cs_pcd_compliance_predicate(r1cs_pcd_compliance_predicate<FieldT> &&other) = default;
    r1cs_pcd_compliance_predicate(const r1cs_pcd_compliance_predicate<FieldT> &other) = default;
    r1cs_pcd_compliance_predicate(const size_t name,
                                  const size_t type,
                                  const r1cs_constraint_system<FieldT> &constraint_system,
                                  const size_t outgoing_message_payload_length,
                                  const size_t max_arity,
                                  const std::vector<size_t> &incoming_message_payload_lengths,
                                  const size_t local_data_length,
                                  const size_t witness_length,
                                  const bool relies_on_same_type_inputs,
                                  const std::set<size_t> accepted_input_types = std::set<size_t>());

    r1cs_pcd_compliance_predicate<FieldT> & operator=(const r1cs_pcd_compliance_predicate<FieldT> &other) = default;

    bool is_well_formed() const;
    bool has_equal_input_and_output_lengths() const;
    bool has_equal_input_lengths() const;

    bool is_satisfied(const std::shared_ptr<r1cs_pcd_message<FieldT> > &outgoing_message,
                      const std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > &incoming_messages,
                      const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data,
                      const r1cs_pcd_witness<FieldT> &witness) const;

    bool operator==(const r1cs_pcd_compliance_predicate<FieldT> &other) const;
    friend std::ostream& operator<< <FieldT>(std::ostream &out, const r1cs_pcd_compliance_predicate<FieldT> &cp);
    friend std::istream& operator>> <FieldT>(std::istream &in, r1cs_pcd_compliance_predicate<FieldT> &cp);
};


} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/compliance_predicate.tcc"

#endif // COMPLIANCE_PREDICATE_HPP_

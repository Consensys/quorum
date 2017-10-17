/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef R1CS_PCD_PARAMS_HPP_
#define R1CS_PCD_PARAMS_HPP_

#include <memory>
#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/compliance_predicate.hpp"

namespace libsnark {

template<typename FieldT>
class r1cs_pcd_compliance_predicate_primary_input {
public:
    std::shared_ptr<r1cs_pcd_message<FieldT> > outgoing_message;

    r1cs_pcd_compliance_predicate_primary_input(const std::shared_ptr<r1cs_pcd_message<FieldT> > &outgoing_message) : outgoing_message(outgoing_message) {}
    r1cs_primary_input<FieldT> as_r1cs_primary_input() const;
};

template<typename FieldT>
class r1cs_pcd_compliance_predicate_auxiliary_input {
public:
    std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > incoming_messages;
    std::shared_ptr<r1cs_pcd_local_data<FieldT> > local_data;
    r1cs_pcd_witness<FieldT> witness;

    r1cs_pcd_compliance_predicate_auxiliary_input(const std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > &incoming_messages,
                                                  const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data,
                                                  const r1cs_pcd_witness<FieldT> &witness) :
        incoming_messages(incoming_messages), local_data(local_data), witness(witness) {}

    r1cs_auxiliary_input<FieldT> as_r1cs_auxiliary_input(const std::vector<size_t> &incoming_message_payload_lengths) const;
};

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_pcd_params.tcc"

#endif // R1CS_PCD_PARAMS_HPP_

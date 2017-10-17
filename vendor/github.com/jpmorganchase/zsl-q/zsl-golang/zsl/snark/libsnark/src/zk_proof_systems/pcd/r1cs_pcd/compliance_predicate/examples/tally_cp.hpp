/** @file
 *****************************************************************************

 Declaration of interfaces for the tally compliance predicate.

 The tally compliance predicate has two purposes:
 (1) it exemplifies the use of interfaces declared in cp_handler.hpp, and
 (2) it enables us to test r1cs_pcd functionalities.

 See
 - src/zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/examples/run_r1cs_sp_ppzkpcd.hpp
 - src/zk_proof_systems/pcd/r1cs_pcd/r1cs_mp_ppzkpcd/examples/run_r1cs_mp_ppzkpcd.hpp
 for code that uses the tally compliance predicate.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TALLY_CP_HPP_
#define TALLY_CP_HPP_

#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/compliance_predicate.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/cp_handler.hpp"

namespace libsnark {

/**
 * Subclasses a R1CS PCD message to the tally compliance predicate.
 */
template<typename FieldT>
class tally_pcd_message : public r1cs_pcd_message<FieldT> {
public:
    size_t wordsize;

    size_t sum;
    size_t count;

    tally_pcd_message(const size_t type,
                      const size_t wordsize,
                      const size_t sum,
                      const size_t count);
    r1cs_variable_assignment<FieldT> payload_as_r1cs_variable_assignment() const;
    void print() const;

    ~tally_pcd_message() = default;
};

template<typename FieldT>
class tally_pcd_local_data : public r1cs_pcd_local_data<FieldT> {
public:
    size_t summand;

    tally_pcd_local_data(const size_t summand);
    r1cs_variable_assignment<FieldT> as_r1cs_variable_assignment() const;
    void print() const;

    ~tally_pcd_local_data() = default;
};

/**
 * Subclass a R1CS compliance predicate handler to the tally compliance predicate handler.
 */
template<typename FieldT>
class tally_cp_handler : public compliance_predicate_handler<FieldT, protoboard<FieldT> > {
public:
    typedef compliance_predicate_handler<FieldT, protoboard<FieldT> > base_handler;
    pb_variable_array<FieldT> incoming_types;

    pb_variable<FieldT> sum_out_packed;
    pb_variable<FieldT> count_out_packed;
    pb_variable_array<FieldT> sum_in_packed;
    pb_variable_array<FieldT> count_in_packed;

    pb_variable_array<FieldT> sum_in_packed_aux;
    pb_variable_array<FieldT> count_in_packed_aux;

    std::shared_ptr<packing_gadget<FieldT> > unpack_sum_out;
    std::shared_ptr<packing_gadget<FieldT> > unpack_count_out;
    std::vector<packing_gadget<FieldT> > pack_sum_in;
    std::vector<packing_gadget<FieldT> > pack_count_in;

    pb_variable<FieldT> type_val_inner_product;
    std::shared_ptr<inner_product_gadget<FieldT> > compute_type_val_inner_product;

    pb_variable_array<FieldT> arity_indicators;

    size_t wordsize;
    size_t message_length;

    tally_cp_handler(const size_t type,
                     const size_t max_arity,
                     const size_t wordsize,
                     const bool relies_on_same_type_inputs = false,
                     const std::set<size_t> accepted_input_types = std::set<size_t>());

    void generate_r1cs_constraints();
    void generate_r1cs_witness(const std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > &incoming_messages,
                               const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data);

    std::shared_ptr<r1cs_pcd_message<FieldT> > get_base_case_message() const;
};

} // libsnark

#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/examples/tally_cp.tcc"

#endif // TALLY_CP_HPP_

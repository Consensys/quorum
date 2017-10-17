/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef SET_COMMITMENT_GADGET_HPP_
#define SET_COMMITMENT_GADGET_HPP_

#include "gadgetlib1/gadget.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "gadgetlib1/gadgets/hashes/hash_io.hpp"
#include "gadgetlib1/gadgets/merkle_tree/merkle_tree_check_read_gadget.hpp"
#include "gadgetlib1/gadgets/set_commitment/set_membership_proof_variable.hpp"

namespace libsnark {

template<typename FieldT, typename HashT>
using set_commitment_variable = digest_variable<FieldT>;

template<typename FieldT, typename HashT>
class set_commitment_gadget : public gadget<FieldT> {
private:
    std::shared_ptr<block_variable<FieldT> > element_block;
    std::shared_ptr<digest_variable<FieldT> > element_digest;
    std::shared_ptr<HashT> hash_element;
    std::shared_ptr<merkle_tree_check_read_gadget<FieldT, HashT> > check_membership;

public:
    size_t tree_depth;
    pb_variable_array<FieldT> element_bits;
    set_commitment_variable<FieldT, HashT> root_digest;
    set_membership_proof_variable<FieldT, HashT> proof;
    pb_linear_combination<FieldT> check_successful;

    set_commitment_gadget(protoboard<FieldT> &pb,
                          const size_t max_entries,
                          const pb_variable_array<FieldT> &element_bits,
                          const set_commitment_variable<FieldT, HashT> &root_digest,
                          const set_membership_proof_variable<FieldT, HashT> &proof,
                          const pb_linear_combination<FieldT> &check_successful,
                          const std::string &annotation_prefix);

    void generate_r1cs_constraints();
    void generate_r1cs_witness();

    static size_t root_size_in_bits();
};

template<typename FieldT, typename HashT>
void test_set_commitment_gadget();

} // libsnark

#include "gadgetlib1/gadgets/set_commitment/set_commitment_gadget.tcc"

#endif // SET_COMMITMENT_GADGET_HPP_

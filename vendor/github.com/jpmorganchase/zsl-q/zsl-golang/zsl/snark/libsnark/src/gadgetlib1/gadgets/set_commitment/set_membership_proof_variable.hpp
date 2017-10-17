/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SET_MEMBERSHIP_PROOF_VARIABLE_HPP_
#define SET_MEMBERSHIP_PROOF_VARIABLE_HPP_

#include "common/data_structures/set_commitment.hpp"
#include "gadgetlib1/gadget.hpp"
#include "gadgetlib1/gadgets/hashes/hash_io.hpp"
#include "gadgetlib1/gadgets/merkle_tree/merkle_authentication_path_variable.hpp"

namespace libsnark {

template<typename FieldT, typename HashT>
class set_membership_proof_variable : public gadget<FieldT> {
public:
    pb_variable_array<FieldT> address_bits;
    std::shared_ptr<merkle_authentication_path_variable<FieldT, HashT> > merkle_path;

    const size_t max_entries;
    const size_t tree_depth;

    set_membership_proof_variable(protoboard<FieldT> &pb,
                                  const size_t max_entries,
                                  const std::string &annotation_prefix);

    void generate_r1cs_constraints();
    void generate_r1cs_witness(const set_membership_proof &proof);

    set_membership_proof get_membership_proof() const;

    static r1cs_variable_assignment<FieldT> as_r1cs_variable_assignment(const set_membership_proof &proof);
};

} // libsnark

#include "gadgetlib1/gadgets/set_commitment/set_membership_proof_variable.tcc"

#endif // SET_MEMBERSHIP_PROOF_VARIABLE_HPP

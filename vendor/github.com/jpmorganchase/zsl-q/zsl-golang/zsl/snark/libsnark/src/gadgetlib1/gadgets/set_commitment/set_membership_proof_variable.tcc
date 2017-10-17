/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SET_MEMBERSHIP_PROOF_VARIABLE_TCC_
#define SET_MEMBERSHIP_PROOF_VARIABLE_TCC_

namespace libsnark {

template<typename FieldT, typename HashT>
set_membership_proof_variable<FieldT, HashT>::set_membership_proof_variable(protoboard<FieldT> &pb,
                                                                            const size_t max_entries,
                                                                            const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    max_entries(max_entries),
    tree_depth(log2(max_entries))
{
    if (tree_depth > 0)
    {
        address_bits.allocate(pb, tree_depth, FMT(annotation_prefix, " address_bits"));
        merkle_path.reset(new merkle_authentication_path_variable<FieldT, HashT>(pb, tree_depth, FMT(annotation_prefix, " merkle_path")));
    }
}

template<typename FieldT, typename HashT>
void set_membership_proof_variable<FieldT, HashT>::generate_r1cs_constraints()
{
    if (tree_depth > 0)
    {
        for (size_t i = 0; i < tree_depth; ++i)
        {
            generate_boolean_r1cs_constraint<FieldT>(this->pb, address_bits[i], FMT(this->annotation_prefix, " address_bits"));
        }
        merkle_path->generate_r1cs_constraints();
    }
}

template<typename FieldT, typename HashT>
void set_membership_proof_variable<FieldT, HashT>::generate_r1cs_witness(const set_membership_proof &proof)
{
    if (tree_depth > 0)
    {
        address_bits.fill_with_bits_of_field_element(this->pb, FieldT(proof.address));
        merkle_path->generate_r1cs_witness(proof.address, proof.merkle_path);
    }
}

template<typename FieldT, typename HashT>
set_membership_proof set_membership_proof_variable<FieldT, HashT>::get_membership_proof() const
{
    set_membership_proof result;

    if (tree_depth == 0)
    {
        result.address = 0;
    }
    else
    {
        result.address = address_bits.get_field_element_from_bits(this->pb).as_ulong();
        result.merkle_path = merkle_path->get_authentication_path(result.address);
    }

    return result;
}

template<typename FieldT, typename HashT>
r1cs_variable_assignment<FieldT> set_membership_proof_variable<FieldT, HashT>::as_r1cs_variable_assignment(const set_membership_proof &proof)
{
    protoboard<FieldT> pb;
    const size_t max_entries = (1ul << (proof.merkle_path.size()));
    set_membership_proof_variable<FieldT, HashT> proof_variable(pb, max_entries, "proof_variable");
    proof_variable.generate_r1cs_witness(proof);

    return pb.full_variable_assignment();
}

} // libsnark

#endif // SET_MEMBERSHIP_PROOF_VARIABLE_TCC

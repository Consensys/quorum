/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef SET_COMMITMENT_GADGET_TCC_
#define SET_COMMITMENT_GADGET_TCC_

#include "common/data_structures/set_commitment.hpp"

namespace libsnark {

template<typename FieldT, typename HashT>
set_commitment_gadget<FieldT, HashT>::set_commitment_gadget(protoboard<FieldT> &pb,
                                                            const size_t max_entries,
                                                            const pb_variable_array<FieldT> &element_bits,
                                                            const set_commitment_variable<FieldT, HashT> &root_digest,
                                                            const set_membership_proof_variable<FieldT, HashT> &proof,
                                                            const pb_linear_combination<FieldT> &check_successful,
                                                            const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), tree_depth(log2(max_entries)), element_bits(element_bits),
    root_digest(root_digest), proof(proof), check_successful(check_successful)
{
    element_block.reset(new block_variable<FieldT>(pb, { element_bits }, FMT(annotation_prefix, " element_block")));

    if (tree_depth == 0)
    {
        hash_element.reset(new HashT(pb, element_bits.size(), *element_block, root_digest, FMT(annotation_prefix, " hash_element")));
    }
    else
    {
        element_digest.reset(new digest_variable<FieldT>(pb, HashT::get_digest_len(),
                                                         FMT(annotation_prefix, " element_digest")));
        hash_element.reset(new HashT(pb, element_bits.size(), *element_block, *element_digest, FMT(annotation_prefix, " hash_element")));
        check_membership.reset(new merkle_tree_check_read_gadget<FieldT, HashT>(pb,
                                                                                tree_depth,
                                                                                proof.address_bits,
                                                                                *element_digest,
                                                                                root_digest,
                                                                                *proof.merkle_path,
                                                                                check_successful,
                                                                                FMT(annotation_prefix, " check_membership")));
    }
}

template<typename FieldT, typename HashT>
void set_commitment_gadget<FieldT, HashT>::generate_r1cs_constraints()
{
    hash_element->generate_r1cs_constraints();

    if (tree_depth > 0)
    {
        check_membership->generate_r1cs_constraints();
    }
}

template<typename FieldT, typename HashT>
void set_commitment_gadget<FieldT, HashT>::generate_r1cs_witness()
{
    hash_element->generate_r1cs_witness();

    if (tree_depth > 0)
    {
        check_membership->generate_r1cs_witness();
    }
}

template<typename FieldT, typename HashT>
size_t set_commitment_gadget<FieldT, HashT>::root_size_in_bits()
{
    return merkle_tree_check_read_gadget<FieldT, HashT>::root_size_in_bits();
}

template<typename FieldT, typename HashT>
void test_set_commitment_gadget()
{
    const size_t digest_len = HashT::get_digest_len();
    const size_t max_set_size = 16;
    const size_t value_size = (HashT::get_block_len() > 0 ? HashT::get_block_len() : 10);

    set_commitment_accumulator<HashT> accumulator(max_set_size, value_size);

    std::vector<bit_vector> set_elems;
    for (size_t i = 0; i < max_set_size; ++i)
    {
        bit_vector elem(value_size);
        std::generate(elem.begin(), elem.end(), [&]() { return std::rand() % 2; });
        set_elems.emplace_back(elem);
        accumulator.add(elem);
        assert(accumulator.is_in_set(elem));
    }

    protoboard<FieldT> pb;
    pb_variable_array<FieldT> element_bits;
    element_bits.allocate(pb, value_size, "element_bits");
    set_commitment_variable<FieldT, HashT> root_digest(pb, digest_len, "root_digest");

    pb_variable<FieldT> check_succesful;
    check_succesful.allocate(pb, "check_succesful");

    set_membership_proof_variable<FieldT, HashT> proof(pb, max_set_size, "proof");

    set_commitment_gadget<FieldT, HashT> sc(pb, max_set_size, element_bits, root_digest, proof, check_succesful, "sc");
    sc.generate_r1cs_constraints();

    /* test all elements from set */
    for (size_t i = 0; i < max_set_size; ++i)
    {
        element_bits.fill_with_bits(pb, set_elems[i]);
        pb.val(check_succesful) = FieldT::one();
        proof.generate_r1cs_witness(accumulator.get_membership_proof(set_elems[i]));
        sc.generate_r1cs_witness();
        root_digest.generate_r1cs_witness(accumulator.get_commitment());
        assert(pb.is_satisfied());
    }
    printf("membership tests OK\n");

    /* test an element not in set */
    for (size_t i = 0; i < value_size; ++i)
    {
        pb.val(element_bits[i]) = FieldT(std::rand() % 2);
    }

    pb.val(check_succesful) = FieldT::zero(); /* do not require the check result to be successful */
    proof.generate_r1cs_witness(accumulator.get_membership_proof(set_elems[0])); /* try it with invalid proof */
    sc.generate_r1cs_witness();
    root_digest.generate_r1cs_witness(accumulator.get_commitment());
    assert(pb.is_satisfied());

    pb.val(check_succesful) = FieldT::one(); /* now require the check result to be succesful */
    proof.generate_r1cs_witness(accumulator.get_membership_proof(set_elems[0])); /* try it with invalid proof */
    sc.generate_r1cs_witness();
    root_digest.generate_r1cs_witness(accumulator.get_commitment());
    assert(!pb.is_satisfied()); /* the protoboard should be unsatisfied */
    printf("non-membership test OK\n");
}

} // libsnark

#endif // SET_COMMITMENT_GADGET_TCC_

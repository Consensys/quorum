/** @file
 *****************************************************************************

 Implementation of a Merkle tree based set commitment scheme.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SET_COMMITMENT_TCC_
#define SET_COMMITMENT_TCC_

namespace libsnark {

template<typename HashT>
set_commitment_accumulator<HashT>::set_commitment_accumulator(const size_t max_entries, const size_t value_size) :
    value_size(value_size)
{
    depth = log2(max_entries);
    digest_size = HashT::get_digest_len();

    tree.reset(new merkle_tree<HashT>(depth, digest_size));
}

template<typename HashT>
void set_commitment_accumulator<HashT>::add(const bit_vector &value)
{
    assert(value_size == 0 || value.size() == value_size);
    const bit_vector hash = HashT::get_hash(value);
    if (hash_to_pos.find(hash) == hash_to_pos.end())
    {
        const size_t pos = hash_to_pos.size();
        tree->set_value(pos, hash);
        hash_to_pos[hash] = pos;
    }
}

template<typename HashT>
bool set_commitment_accumulator<HashT>::is_in_set(const bit_vector &value) const
{
    assert(value_size == 0 || value.size() == value_size);
    const bit_vector hash = HashT::get_hash(value);
    return (hash_to_pos.find(hash) != hash_to_pos.end());
}

template<typename HashT>
set_commitment set_commitment_accumulator<HashT>::get_commitment() const
{
    return tree->get_root();
}

template<typename HashT>
set_membership_proof set_commitment_accumulator<HashT>::get_membership_proof(const bit_vector &value) const
{
    const bit_vector hash = HashT::get_hash(value);
    auto it = hash_to_pos.find(hash);
    assert(it != hash_to_pos.end());

    set_membership_proof proof;
    proof.address = it->second;
    proof.merkle_path = tree->get_path(it->second);

    return proof;
}

} // libsnark

#endif // SET_COMMITMENT_TCC_

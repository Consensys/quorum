/** @file
 *****************************************************************************

 Declaration of interfaces for a Merkle tree based set commitment scheme.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SET_COMMITMENT_HPP_
#define SET_COMMITMENT_HPP_

#include "common/utils.hpp"
#include "common/data_structures/merkle_tree.hpp"
#include "gadgetlib1/gadgets/hashes/hash_io.hpp" // TODO: the current structure is suboptimal

namespace libsnark {

typedef bit_vector set_commitment;

struct set_membership_proof {
    size_t address;
    merkle_authentication_path merkle_path;

    bool operator==(const set_membership_proof &other) const;
    size_t size_in_bits() const;
    friend std::ostream& operator<<(std::ostream &out, const set_membership_proof &other);
    friend std::istream& operator>>(std::istream &in, set_membership_proof &other);
};

template<typename HashT>
class set_commitment_accumulator {
private:
    std::shared_ptr<merkle_tree<HashT> > tree;
    std::map<bit_vector, size_t> hash_to_pos;
public:

    size_t depth;
    size_t digest_size;
    size_t value_size;

    set_commitment_accumulator(const size_t max_entries, const size_t value_size=0);

    void add(const bit_vector &value);
    bool is_in_set(const bit_vector &value) const;
    set_commitment get_commitment() const;

    set_membership_proof get_membership_proof(const bit_vector &value) const;
};

} // libsnark

/* note that set_commitment has both .cpp, for implementation of
   non-templatized code (methods of set_membership_proof) and .tcc
   (implementation of set_commitment_accumulator<HashT> */
#include "common/data_structures/set_commitment.tcc"

#endif // SET_COMMITMENT_HPP_

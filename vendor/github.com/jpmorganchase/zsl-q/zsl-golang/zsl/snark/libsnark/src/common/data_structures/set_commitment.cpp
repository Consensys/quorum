/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "common/data_structures/set_commitment.hpp"
#include "common/serialization.hpp"

namespace libsnark {

bool set_membership_proof::operator==(const set_membership_proof &other) const
{
    return (this->address == other.address &&
            this->merkle_path == other.merkle_path);
}

size_t set_membership_proof::size_in_bits() const
{
    if (merkle_path.empty())
    {
        return (8 * sizeof(address));
    }
    else
    {
        return (8 * sizeof(address) + merkle_path[0].size() * merkle_path.size());
    }
}

std::ostream& operator<<(std::ostream &out, const set_membership_proof &proof)
{
    out << proof.address << "\n";
    out << proof.merkle_path.size() << "\n";
    for (size_t i = 0; i < proof.merkle_path.size(); ++i)
    {
        output_bool_vector(out, proof.merkle_path[i]);
    }

    return out;
}

std::istream& operator>>(std::istream &in, set_membership_proof &proof)
{
    in >> proof.address;
    consume_newline(in);
    size_t tree_depth;
    in >> tree_depth;
    consume_newline(in);
    proof.merkle_path.resize(tree_depth);

    for (size_t i = 0; i < tree_depth; ++i)
    {
        input_bool_vector(in, proof.merkle_path[i]);
    }

    return in;
}

} // libsnark

/** @file
 *****************************************************************************

 Declaration of interfaces for a delegated random-access memory.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef DELEGATED_RA_MEMORY_HPP_
#define DELEGATED_RA_MEMORY_HPP_

#include <map>
#include <memory>
#include <vector>

#include "common/data_structures/merkle_tree.hpp"
#include "relations/ram_computations/memory/memory_interface.hpp"

namespace libsnark {

template<typename HashT>
class delegated_ra_memory : public memory_interface {
private:
    bit_vector int_to_tree_elem(const size_t i) const;
    size_t int_from_tree_elem(const bit_vector &v) const;

    std::unique_ptr<merkle_tree<HashT> > contents;

public:
    delegated_ra_memory(const size_t num_addresses, const size_t value_size);
    delegated_ra_memory(const size_t num_addresses, const size_t value_size, const std::vector<size_t> &contents_as_vector);
    delegated_ra_memory(const size_t num_addresses, const size_t value_size, const memory_contents &contents_as_map);

    size_t get_value(const size_t address) const;
    void set_value(const size_t address, const size_t value);

    typename HashT::hash_value_type get_root() const;
    typename HashT::merkle_authentication_path_type get_path(const size_t address) const;

    void dump() const;
};

} // libsnark

#include "relations/ram_computations/memory/delegated_ra_memory.tcc"

#endif // DELEGATED_RA_MEMORY_HPP_

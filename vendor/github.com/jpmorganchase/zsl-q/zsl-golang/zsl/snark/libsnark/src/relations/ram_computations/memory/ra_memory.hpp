/** @file
 *****************************************************************************

 Declaration of interfaces for a random-access memory.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RA_MEMORY_HPP_
#define RA_MEMORY_HPP_

#include "relations/ram_computations/memory/memory_interface.hpp"

namespace libsnark {

/**
 * A random-access memory maintains the memory's contents via a map (from addresses to values).
 */
class ra_memory : public memory_interface {
public:

    memory_contents contents;

    ra_memory(const size_t num_addresses, const size_t value_size);
    ra_memory(const size_t num_addresses, const size_t value_size, const std::vector<size_t> &contents_as_vector);
    ra_memory(const size_t num_addresses, const size_t value_size, const memory_contents &contents);

    size_t get_value(const size_t address) const;
    void set_value(const size_t address, const size_t value);

};

} // libsnark

#endif // RA_MEMORY_HPP_

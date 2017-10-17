/** @file
 *****************************************************************************

 Implementation of interfaces for a random-access memory.

 See ra_memory.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <cassert>

#include "relations/ram_computations/memory/ra_memory.hpp"

namespace libsnark {

ra_memory::ra_memory(const size_t num_addresses, const size_t value_size) :
    memory_interface(num_addresses, value_size)
{
}

ra_memory::ra_memory(const size_t num_addresses,
                     const size_t value_size,
                     const std::vector<size_t> &contents_as_vector) :
    memory_interface(num_addresses, value_size)
{
    /* copy std::vector into std::map */
    for (size_t i = 0; i < contents_as_vector.size(); ++i)
    {
        contents[i] = contents_as_vector[i];
    }
}


ra_memory::ra_memory(const size_t num_addresses,
                     const size_t value_size,
                     const memory_contents &contents) :
    memory_interface(num_addresses, value_size), contents(contents)
{
}

size_t ra_memory::get_value(const size_t address) const
{
    assert(address < num_addresses);
    auto it = contents.find(address);
    return (it == contents.end() ? 0 : it->second);
}

void ra_memory::set_value(const size_t address, const size_t value)
{
    contents[address] = value;
}

} // libsnark

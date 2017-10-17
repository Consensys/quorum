/** @file
 *****************************************************************************

 Implementation of interfaces for functions to sample examples of memory contents.

 See memory_contents_examples.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "relations/ram_computations/memory/examples/memory_contents_examples.hpp"

#include <cstdlib>
#include <map>
#include <set>

namespace libsnark {

memory_contents block_memory_contents(const size_t num_addresses,
                                      const size_t value_size,
                                      const size_t block1_size,
                                      const size_t block2_size)
{
    const size_t max_unit = 1ul<<value_size;

    memory_contents result;
    for (size_t i = 0; i < block1_size; ++i)
    {
        result[i] = std::rand() % max_unit;
    }

    for (size_t i = 0; i < block2_size; ++i)
    {
        result[num_addresses/2+i] = std::rand() % max_unit;
    }

    return result;
}

memory_contents random_memory_contents(const size_t num_addresses,
                                       const size_t value_size,
                                       const size_t num_filled)
{
    const size_t max_unit = 1ul<<value_size;

    std::set<size_t> unfilled;
    for (size_t i = 0; i < num_addresses; ++i)
    {
        unfilled.insert(i);
    }

    memory_contents result;
    for (size_t i = 0; i < num_filled; ++i)
    {
        auto it = unfilled.begin();
        std::advance(it, std::rand() % unfilled.size());
        result[*it] = std::rand() % max_unit;
        unfilled.erase(it);
    }

    return result;
}

} // libsnark

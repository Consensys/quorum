/** @file
 *****************************************************************************

 Declaration of interfaces for functions to sample examples of memory contents.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MEMORY_CONTENTS_EXAMPLES_HPP_
#define MEMORY_CONTENTS_EXAMPLES_HPP_

#include "relations/ram_computations/memory/memory_interface.hpp"

namespace libsnark {

/**
 * Sample memory contents consisting of two blocks of random values;
 * the first block is located at the beginning of memory, while
 * the second block is located half-way through memory.
 */
memory_contents block_memory_contents(const size_t num_addresses,
                                      const size_t value_size,
                                      const size_t block1_size,
                                      const size_t block2_size);

/**
 * Sample memory contents having a given number of non-zero entries;
 * each non-zero entry is a random value at a random address (approximately).
 */
memory_contents random_memory_contents(const size_t num_addresses,
                                       const size_t value_size,
                                       const size_t num_filled);

} // libsnark

#endif // MEMORY_CONTENTS_EXAMPLES_HPP_

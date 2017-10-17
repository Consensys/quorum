/** @file
 *****************************************************************************

 Declaration of interfaces for a memory store trace.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MEMORY_STORE_TRACE_HPP_
#define MEMORY_STORE_TRACE_HPP_

#include "relations/ram_computations/memory/memory_interface.hpp"

namespace libsnark {

/**
 * A pair consisting of an address and a value.
 * It represents a memory store.
 */
typedef std::pair<size_t, size_t> address_and_value;

/**
 * A list in which each component consists of a timestamp and a memory store.
 */
class memory_store_trace {
private:
    std::map<size_t, address_and_value> entries;

public:
    memory_store_trace();
    address_and_value get_trace_entry(const size_t timestamp) const;
    std::map<size_t, address_and_value> get_all_trace_entries() const;
    void set_trace_entry(const size_t timestamp, const address_and_value &av);

    memory_contents as_memory_contents() const;
};

} // libsnark

#endif // MEMORY_STORE_TRACE_HPP_

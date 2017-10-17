/** @file
 *****************************************************************************

 Implementation of interfaces for a memory store trace.

 See memory_store_trace.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "relations/ram_computations/memory/memory_store_trace.hpp"

namespace libsnark {

memory_store_trace::memory_store_trace()
{
}

address_and_value memory_store_trace::get_trace_entry(const size_t timestamp) const
{
    auto it = entries.find(timestamp);
    return (it != entries.end() ? it->second : std::make_pair<size_t, size_t>(0, 0));
}

std::map<size_t, address_and_value> memory_store_trace::get_all_trace_entries() const
{
    return entries;
}

void memory_store_trace::set_trace_entry(const size_t timestamp, const address_and_value &av)
{
    entries[timestamp] = av;
}

memory_contents memory_store_trace::as_memory_contents() const
{
    memory_contents result;

    for (auto &ts_and_addrval : entries)
    {
        result[ts_and_addrval.second.first] = ts_and_addrval.second.second;
    }

    return result;
}

} // libsnark

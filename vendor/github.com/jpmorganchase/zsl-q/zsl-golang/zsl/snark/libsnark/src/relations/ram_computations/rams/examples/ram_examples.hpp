/** @file
 *****************************************************************************

 Declaration of interfaces for a RAM example, as well as functions to sample
 RAM examples with prescribed parameters (according to some distribution).

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_EXAMPLES_HPP_
#define RAM_EXAMPLES_HPP_

#include "relations/ram_computations/rams/ram_params.hpp"

namespace libsnark {

template<typename ramT>
struct ram_example {
    ram_architecture_params<ramT> ap;
    size_t boot_trace_size_bound;
    size_t time_bound;
    ram_boot_trace<ramT> boot_trace;
    ram_input_tape<ramT> auxiliary_input;
};

/**
 * For now: only specialized to TinyRAM
 */
template<typename ramT>
ram_example<ramT> gen_ram_example_simple(const ram_architecture_params<ramT> &ap, const size_t boot_trace_size_bound, const size_t time_bound, const bool satisfiable=true);

/**
 * For now: only specialized to TinyRAM
 */
template<typename ramT>
ram_example<ramT> gen_ram_example_complex(const ram_architecture_params<ramT> &ap, const size_t boot_trace_size_bound, const size_t time_bound, const bool satisfiable=true);

} // libsnark

#include "relations/ram_computations/rams/examples/ram_examples.tcc"

#endif // RAM_EXAMPLES_HPP_

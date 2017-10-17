/** @file
 *****************************************************************************

 Declaration of public-parameter selector for RAMs.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_PARAMS_HPP_
#define RAM_PARAMS_HPP_

#include <vector>

#include "relations/ram_computations/memory/memory_store_trace.hpp"

namespace libsnark {

/*
  When declaring a new ramT one should do a make it a class that declares typedefs for:

  base_field_type
  ram_cpu_checker_type
  architecture_params_type

  For ram_to_r1cs reduction currently the following are also necessary:
  protoboard_type (e.g. tinyram_protoboard<FieldT>)
  gadget_base_type (e.g. tinyram_gadget<FieldT>)
  cpu_state_variable_type (must have pb_variable_array<FieldT> all_vars)

  The ramT class must also have a static size_t variable
  timestamp_length, which specifies the zk-SNARK reduction timestamp
  length.
*/

template<typename ramT>
using ram_base_field = typename ramT::base_field_type;

template<typename ramT>
using ram_cpu_state = bit_vector;

template<typename ramT>
using ram_boot_trace = memory_store_trace;

template<typename ramT>
using ram_protoboard = typename ramT::protoboard_type;

template<typename ramT>
using ram_gadget_base = typename ramT::gadget_base_type;

template<typename ramT>
using ram_cpu_checker = typename ramT::cpu_checker_type;

template<typename ramT>
using ram_architecture_params = typename ramT::architecture_params_type;

template<typename ramT>
using ram_input_tape = std::vector<size_t>;

/*
  One should also make the following methods for ram_architecture_params

  (We are not yet making a ram_architecture_params base class, as it
  would require base class for ram_program

  TODO: make this base class)

  size_t address_size();
  size_t value_size();
  size_t cpu_state_size();
  size_t initial_pc_addr();
  bit_vector initial_cpu_state();
*/

} // libsnark

#endif // RAM_PARAMS_HPP_

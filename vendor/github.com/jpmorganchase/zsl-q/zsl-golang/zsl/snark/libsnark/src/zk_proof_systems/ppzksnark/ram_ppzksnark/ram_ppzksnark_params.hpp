/** @file
 *****************************************************************************

 Declaration of public-parameter selector for the RAM ppzkSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_PPZKSNARK_PARAMS_HPP_
#define RAM_PPZKSNARK_PARAMS_HPP_

namespace libsnark {

/**
 * The interfaces of the RAM ppzkSNARK are templatized via the parameter
 * ram_ppzksnark_ppT. When used, the interfaces must be invoked with
 * a particular parameter choice; let 'my_ram_ppzksnark_pp' denote this choice.
 *
 * my_ram_ppzksnark_pp needs to contain typedefs for the typenames
 * - snark_pp, and
 * - machine_pp.
 * as well as a method with the following signature:
 * - static void init_public_params();
 *
 * For example, if you want to use the types my_snark_pp and my_machine_pp,
 * then you could declare my_ram_ppzksnark_pp as follows:
 *
 *   class my_ram_ppzksnark_pp {
 *   public:
 *       typedef my_snark_pp snark_pp;
 *       typedef my_machine_pp machine_pp;
 *       static void init_public params()
 *       {
 *           snark_pp::init_public_params(); // and additional initialization if needed
 *       }
 *   };
 *
 * Having done the above, my_ram_ppzksnark_pp can be used as a template parameter.
 *
 * Look for for default_tinyram_ppzksnark_pp in the file
 *
 *   common/default_types/ram_ppzksnark_pp.hpp
 *
 * for an example of the above steps for the case of "RAM=TinyRAM".
 *
 */

/**
 * Below are various template aliases (used for convenience).
 */

template<typename ram_ppzksnark_ppT>
using ram_ppzksnark_snark_pp = typename ram_ppzksnark_ppT::snark_pp;

template<typename ram_ppzksnark_ppT>
using ram_ppzksnark_machine_pp = typename ram_ppzksnark_ppT::machine_pp;

template<typename ram_ppzksnark_ppT>
using ram_ppzksnark_architecture_params = ram_architecture_params<ram_ppzksnark_machine_pp<ram_ppzksnark_ppT> >;

template<typename ram_ppzksnark_ppT>
using ram_ppzksnark_primary_input = ram_boot_trace<ram_ppzksnark_machine_pp<ram_ppzksnark_ppT> >;

template<typename ram_ppzksnark_ppT>
using ram_ppzksnark_auxiliary_input = ram_input_tape<ram_ppzksnark_machine_pp<ram_ppzksnark_ppT> >;

} // libsnark

#endif // RAM_PPZKSNARK_PARAMS_HPP_

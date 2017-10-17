/** @file
 *****************************************************************************

 Declaration of public-parameter selector for the BACS ppzkSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_PPZKSNARK_PARAMS_HPP_
#define BACS_PPZKSNARK_PARAMS_HPP_

#include "relations/circuit_satisfaction_problems/bacs/bacs.hpp"

namespace libsnark {

/**
 * Below are various template aliases (used for convenience).
 */

template<typename ppT>
using bacs_ppzksnark_circuit = bacs_circuit<Fr<ppT> >;

template<typename ppT>
using bacs_ppzksnark_primary_input = bacs_primary_input<Fr<ppT> >;

template<typename ppT>
using bacs_ppzksnark_auxiliary_input = bacs_auxiliary_input<Fr<ppT> >;

} // libsnark

#endif // BACS_PPZKSNARK_PARAMS_HPP_

/** @file
 *****************************************************************************

 Declaration of public-parameter selector for the USCS ppzkSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef USCS_PPZKSNARK_PARAMS_HPP_
#define USCS_PPZKSNARK_PARAMS_HPP_

#include "relations/constraint_satisfaction_problems/uscs/uscs.hpp"

namespace libsnark {

/**
 * Below are various template aliases (used for convenience).
 */

template<typename ppT>
using uscs_ppzksnark_constraint_system = uscs_constraint_system<Fr<ppT> >;

template<typename ppT>
using uscs_ppzksnark_primary_input = uscs_primary_input<Fr<ppT> >;

template<typename ppT>
using uscs_ppzksnark_auxiliary_input = uscs_auxiliary_input<Fr<ppT> >;

} // libsnark

#endif // USCS_PPZKSNARK_PARAMS_HPP_

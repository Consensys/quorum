/** @file
 *****************************************************************************

 Declaration of public-parameter selector for the TBCS ppzkSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TBCS_PPZKSNARK_PARAMS_HPP_
#define TBCS_PPZKSNARK_PARAMS_HPP_

#include "relations/circuit_satisfaction_problems/tbcs/tbcs.hpp"

namespace libsnark {

/**
 * Below are various typedefs aliases (used for uniformity with other proof systems).
 */

typedef tbcs_circuit tbcs_ppzksnark_circuit;

typedef tbcs_primary_input tbcs_ppzksnark_primary_input;

typedef tbcs_auxiliary_input tbcs_ppzksnark_auxiliary_input;

} // libsnark

#endif // TBCS_PPZKSNARK_PARAMS_HPP_

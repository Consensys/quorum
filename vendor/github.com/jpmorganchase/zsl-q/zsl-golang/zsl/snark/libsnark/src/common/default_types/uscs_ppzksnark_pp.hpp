
/** @file
 *****************************************************************************

 This file defines default_uscs_ppzksnark_pp based on the elliptic curve
 choice selected in ec_pp.hpp.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef USCS_PPZKSNARK_PP_HPP_
#define USCS_PPZKSNARK_PP_HPP_

#include "common/default_types/ec_pp.hpp"

namespace libsnark {
typedef default_ec_pp default_uscs_ppzksnark_pp;
} // libsnark

#endif // USCS_PPZKSNARK_PP_HPP_


/** @file
 *****************************************************************************

 This file defines default_bacs_ppzksnark_pp based on the elliptic curve
 choice selected in ec_pp.hpp.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_PPZKSNARK_PP_HPP_
#define BACS_PPZKSNARK_PP_HPP_

#include "common/default_types/ec_pp.hpp"

namespace libsnark {
typedef default_ec_pp default_bacs_ppzksnark_pp;
} // libsnark

#endif // BACS_PPZKSNARK_PP_HPP_

/** @file
 *****************************************************************************

 This file provides the initialization methods for the default PCD cycle.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "common/default_types/r1cs_ppzkpcd_pp.hpp"

namespace libsnark {

void default_r1cs_ppzkpcd_pp::init_public_params()
{
    curve_A_pp::init_public_params();
    curve_B_pp::init_public_params();
}

} // libsnark

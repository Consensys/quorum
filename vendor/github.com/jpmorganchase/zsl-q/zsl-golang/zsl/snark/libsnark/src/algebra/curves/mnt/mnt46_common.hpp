/** @file
 *****************************************************************************

 Declaration of functionality that is shared among MNT curves.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT46_COMMON_HPP_
#define MNT46_COMMON_HPP_

#include "algebra/fields/bigint.hpp"

namespace libsnark {

const mp_size_t mnt46_A_bitcount = 298;
const mp_size_t mnt46_B_bitcount = 298;

const mp_size_t mnt46_A_limbs = (mnt46_A_bitcount+GMP_NUMB_BITS-1)/GMP_NUMB_BITS;
const mp_size_t mnt46_B_limbs = (mnt46_B_bitcount+GMP_NUMB_BITS-1)/GMP_NUMB_BITS;

extern bigint<mnt46_A_limbs> mnt46_modulus_A;
extern bigint<mnt46_B_limbs> mnt46_modulus_B;

} // libsnark

#endif

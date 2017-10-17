/** @file
 *****************************************************************************

 Declaration of functions for generating randomness.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RNG_HPP_
#define RNG_HPP_

#include <cstdint>

namespace libsnark {

template<typename FieldT>
FieldT SHA512_rng(const uint64_t idx);

} // libsnark

#include "common/rng.tcc"

#endif // RNG_HPP_

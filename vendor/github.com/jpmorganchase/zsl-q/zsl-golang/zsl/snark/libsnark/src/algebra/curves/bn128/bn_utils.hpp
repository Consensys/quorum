/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BN_UTILS_HPP_
#define BN_UTILS_HPP_
#include <vector>
#include "bn.h"

namespace libsnark {

template<typename FieldT>
void bn_batch_invert(std::vector<FieldT> &vec);

} // libsnark

#include "algebra/curves/bn128/bn_utils.tcc"

#endif // BN_UTILS_HPP_

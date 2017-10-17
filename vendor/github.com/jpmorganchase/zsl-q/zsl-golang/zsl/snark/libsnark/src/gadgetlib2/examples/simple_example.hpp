/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SIMPLE_EXAMPLE_HPP_
#define SIMPLE_EXAMPLE_HPP_

#include "common/default_types/ec_pp.hpp"
#include "relations/constraint_satisfaction_problems/r1cs/examples/r1cs_examples.hpp"

namespace libsnark {

r1cs_example<Fr<default_ec_pp> > gen_r1cs_example_from_gadgetlib2_protoboard(const size_t size);

} // libsnark

#endif // SIMPLE_EXAMPLE_HPP_

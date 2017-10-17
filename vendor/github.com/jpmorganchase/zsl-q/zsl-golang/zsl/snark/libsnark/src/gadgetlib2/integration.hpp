/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef INTEGRATION_HPP_
#define INTEGRATION_HPP_

#include "common/default_types/ec_pp.hpp"
#include "relations/constraint_satisfaction_problems/r1cs/r1cs.hpp"
#include "gadgetlib2/protoboard.hpp"

namespace libsnark {

r1cs_constraint_system<Fr<default_ec_pp> > get_constraint_system_from_gadgetlib2(const gadgetlib2::Protoboard &pb);
r1cs_variable_assignment<Fr<default_ec_pp> > get_variable_assignment_from_gadgetlib2(const gadgetlib2::Protoboard &pb);

} // libsnark

#endif // INTEGRATION_HPP_

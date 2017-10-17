/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#ifndef CRH_GADGET_HPP_
#define CRH_GADGET_HPP_

#include "gadgetlib1/gadgets/hashes/knapsack/knapsack_gadget.hpp"

namespace libsnark {

// for now all CRH gadgets are knapsack CRH's; can be easily extended
// later to more expressive selector types.
template<typename FieldT>
using CRH_with_field_out_gadget = knapsack_CRH_with_field_out_gadget<FieldT>;

template<typename FieldT>
using CRH_with_bit_out_gadget = knapsack_CRH_with_bit_out_gadget<FieldT>;

} // libsnark

#endif // CRH_GADGET_HPP_

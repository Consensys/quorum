/** @file
 *****************************************************************************

 Declaration of interfaces for the memory load&store gadget.

 The gadget can be used to verify a memory load, followed by a store to the
 same address, from a "delegated memory".

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MEMORY_LOAD_STORE_GADGET_HPP_
#define MEMORY_LOAD_STORE_GADGET_HPP_

#include "gadgetlib1/gadgets/merkle_tree/merkle_tree_check_update_gadget.hpp"

namespace libsnark {

template<typename FieldT, typename HashT>
using memory_load_store_gadget = merkle_tree_check_update_gadget<FieldT, HashT>;

} // libsnark

#endif // MEMORY_LOAD_STORE_GADGET_HPP_

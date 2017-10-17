/** @file
 *****************************************************************************

 Implementation of interfaces for a protoboard for the FOORAM CPU.

 See fooram_protoboard.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FOORAM_PROTOBOARD_TCC_
#define FOORAM_PROTOBOARD_TCC_

namespace libsnark {

template<typename FieldT>
fooram_protoboard<FieldT>::fooram_protoboard(const fooram_architecture_params &ap) :
    protoboard<FieldT>(), ap(ap)
{
}

template<typename FieldT>
fooram_gadget<FieldT>::fooram_gadget(fooram_protoboard<FieldT> &pb, const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), pb(pb)
{
}

} // libsnark

#endif // FOORAM_PROTOBOARD_HPP_

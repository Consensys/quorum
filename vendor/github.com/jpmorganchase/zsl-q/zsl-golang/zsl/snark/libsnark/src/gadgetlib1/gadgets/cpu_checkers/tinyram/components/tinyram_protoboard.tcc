/** @file
 *****************************************************************************

 Implementation of interfaces for a protoboard for TinyRAM.

 See tinyram_protoboard.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TINYRAM_PROTOBOARD_TCC_
#define TINYRAM_PROTOBOARD_TCC_

namespace libsnark {

template<typename FieldT>
tinyram_protoboard<FieldT>::tinyram_protoboard(const tinyram_architecture_params &ap) :
    ap(ap)
{
}

template<typename FieldT>
tinyram_gadget<FieldT>::tinyram_gadget(tinyram_protoboard<FieldT> &pb, const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), pb(pb)
{
}

template<typename FieldT>
tinyram_standard_gadget<FieldT>::tinyram_standard_gadget(tinyram_protoboard<FieldT> &pb, const std::string &annotation_prefix) :
    tinyram_gadget<FieldT>(pb, annotation_prefix)
{
}

} // libsnark

#endif // TINYRAM_PROTOBOARD_TCC_

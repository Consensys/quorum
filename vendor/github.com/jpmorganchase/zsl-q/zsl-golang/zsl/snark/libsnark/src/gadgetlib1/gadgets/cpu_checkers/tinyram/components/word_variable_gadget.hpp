/** @file
 *****************************************************************************

 Declaration of interfaces for (single and double) word gadgets.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef WORD_VARIABLE_GADGET_HPP_
#define WORD_VARIABLE_GADGET_HPP_

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/tinyram_protoboard.hpp"

namespace libsnark {

/**
 * Holds both binary and field representaton of a word.
 */
template<typename FieldT>
class word_variable_gadget : public dual_variable_gadget<FieldT> {
public:
    word_variable_gadget(tinyram_protoboard<FieldT> &pb, const std::string &annotation_prefix="") :
        dual_variable_gadget<FieldT>(pb, pb.ap.w, annotation_prefix) {}
    word_variable_gadget(tinyram_protoboard<FieldT> &pb, const pb_variable_array<FieldT> &bits, const std::string &annotation_prefix="") :
        dual_variable_gadget<FieldT>(pb, bits, annotation_prefix) {}
    word_variable_gadget(tinyram_protoboard<FieldT> &pb, const pb_variable<FieldT> &packed, const std::string &annotation_prefix="") :
        dual_variable_gadget<FieldT>(pb, packed, pb.ap.w, annotation_prefix) {}
};

/**
 * Holds both binary and field representaton of a double word.
 */
template<typename FieldT>
class doubleword_variable_gadget : public dual_variable_gadget<FieldT> {
public:
    doubleword_variable_gadget(tinyram_protoboard<FieldT> &pb, const std::string &annotation_prefix="") :
        dual_variable_gadget<FieldT>(pb, 2*pb.ap.w, annotation_prefix) {}
    doubleword_variable_gadget(tinyram_protoboard<FieldT> &pb, const pb_variable_array<FieldT> &bits, const std::string &annotation_prefix="") :
        dual_variable_gadget<FieldT>(pb, bits, annotation_prefix) {}
    doubleword_variable_gadget(tinyram_protoboard<FieldT> &pb, const pb_variable<FieldT> &packed, const std::string &annotation_prefix="") :
        dual_variable_gadget<FieldT>(pb, packed, 2*pb.ap.w, annotation_prefix) {}
};

} // libsnark

#endif // WORD_VARIABLE_GADGET_HPP_

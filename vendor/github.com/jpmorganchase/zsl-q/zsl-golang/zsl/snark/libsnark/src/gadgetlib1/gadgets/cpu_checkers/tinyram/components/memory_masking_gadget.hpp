/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM memory masking gadget.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MEMORY_MASKING_GADGET_HPP_
#define MEMORY_MASKING_GADGET_HPP_

namespace libsnark {

/**
 * The memory masking gadget checks if a specified part of a double
 * word is correctly modified. In TinyRAM CPU checker we use this to
 * implement byte addressing and word addressing for the memory that
 * consists of double words.
 *
 * More precisely, memory masking gadgets takes the following
 * arguments:
 *
 * dw_contents_prev, dw_contents_next -- the contents of the memory
 *
 * double word before and after the access
 *
 * access_is_word -- a boolean indicating if access is word
 *
 * access_is_byte -- a boolean indicating if access is byte
 *
 * subaddress -- an integer specifying which byte (if access_is_byte=1)
 * or word (if access_is_byte=1) this access is operating on
 *
 * subcontents -- contents of the byte, resp., word to be operated on
 *
 * Memory masking gadget enforces that dw_contents_prev is equal to
 * dw_contents_next everywhere, except subaddres-th byte (if
 * access_is_byte = 1), or MSB(subaddress)-th word (if access_is_word =
 * 1). The corresponding byte, resp., word in dw_contents_next is
 * required to equal subcontents.
 *
 * Note that indexing MSB(subaddress)-th word is the same as indexing
 * the word specified by subaddress expressed in bytes and aligned to
 * the word boundary by rounding the subaddress down.
 *
 * Requirements: The caller is required to perform bounds checks on
 * subcontents. The caller is also required to ensure that exactly one
 * of access_is_word and access_is_byte is set to 1.
 */
template<typename FieldT>
class memory_masking_gadget : public tinyram_standard_gadget<FieldT> {
private:
    pb_linear_combination<FieldT> shift;
    pb_variable<FieldT> is_word0;
    pb_variable<FieldT> is_word1;
    pb_variable_array<FieldT> is_subaddress;
    pb_variable_array<FieldT> is_byte;

    pb_linear_combination<FieldT> masked_out_word0;
    pb_linear_combination<FieldT> masked_out_word1;
    pb_linear_combination_array<FieldT> masked_out_bytes;

    std::shared_ptr<inner_product_gadget<FieldT> > get_masked_out_dw_contents_prev;

    pb_variable<FieldT> masked_out_dw_contents_prev;
    pb_variable<FieldT> expected_dw_contents_next;
public:
    doubleword_variable_gadget<FieldT> dw_contents_prev;
    dual_variable_gadget<FieldT> subaddress;
    pb_linear_combination<FieldT> subcontents;
    pb_linear_combination<FieldT> access_is_word;
    pb_linear_combination<FieldT> access_is_byte;
    doubleword_variable_gadget<FieldT> dw_contents_next;

    memory_masking_gadget(tinyram_protoboard<FieldT> &pb,
                          const doubleword_variable_gadget<FieldT> &dw_contents_prev,
                          const dual_variable_gadget<FieldT> &subaddress,
                          const pb_linear_combination<FieldT> &subcontents,
                          const pb_linear_combination<FieldT> &access_is_word,
                          const pb_linear_combination<FieldT> &access_is_byte,
                          const doubleword_variable_gadget<FieldT> &dw_contents_next,
                          const std::string& annotation_prefix="");
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/memory_masking_gadget.tcc"

#endif // MEMORY_MASKING_GADGET_HPP_

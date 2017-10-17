/** @file
 *****************************************************************************

 Implementation of interfaces for the TinyRAM memory masking gadget.

 See memory_masking_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MEMORY_MASKING_GADGET_TCC_
#define MEMORY_MASKING_GADGET_TCC_

namespace libsnark {

template<typename FieldT>
memory_masking_gadget<FieldT>::memory_masking_gadget(tinyram_protoboard<FieldT> &pb,
                                                     const doubleword_variable_gadget<FieldT> &dw_contents_prev,
                                                     const dual_variable_gadget<FieldT> &subaddress,
                                                     const pb_linear_combination<FieldT> &subcontents,
                                                     const pb_linear_combination<FieldT> &access_is_word,
                                                     const pb_linear_combination<FieldT> &access_is_byte,
                                                     const doubleword_variable_gadget<FieldT> &dw_contents_next,
                                                     const std::string& annotation_prefix) :
    tinyram_standard_gadget<FieldT>(pb, annotation_prefix),
    dw_contents_prev(dw_contents_prev),
    subaddress(subaddress),
    subcontents(subcontents),
    access_is_word(access_is_word),
    access_is_byte(access_is_byte),
    dw_contents_next(dw_contents_next)
{
    /*
      Indicator variables for access being to word_0, word_1, and
      byte_0, byte_1, ...

      We use little-endian indexing here (least significant
      bit/byte/word has the smallest address).
    */
    is_word0.allocate(pb, FMT(this->annotation_prefix, " is_word0"));
    is_word1.allocate(pb, FMT(this->annotation_prefix, " is_word1"));
    is_subaddress.allocate(pb, 2 * pb.ap.bytes_in_word(), FMT(this->annotation_prefix, " is_sub_address"));
    is_byte.allocate(pb, 2 * pb.ap.bytes_in_word(), FMT(this->annotation_prefix, " is_byte"));

    /*
      Get value of the dw_contents_prev for which the specified entity
      is masked out to be zero. E.g. the value of masked_out_bytes[3]
      will be the same as the value of dw_contents_prev, when 3rd
      (0-indexed) byte is set to all zeros.
    */
    masked_out_word0.assign(pb, (FieldT(2)^pb.ap.w) * pb_packing_sum<FieldT>(
                                pb_variable_array<FieldT>(dw_contents_prev.bits.begin() + pb.ap.w,
                                                          dw_contents_prev.bits.begin() + 2 * pb.ap.w)));
    masked_out_word1.assign(pb, pb_packing_sum<FieldT>(
                                pb_variable_array<FieldT>(dw_contents_prev.bits.begin(),
                                                          dw_contents_prev.bits.begin() + pb.ap.w)));
    masked_out_bytes.resize(2 * pb.ap.bytes_in_word());

    for (size_t i = 0; i < 2 * pb.ap.bytes_in_word(); ++i)
    {
        /* just subtract out the byte to be masked */
        masked_out_bytes[i].assign(pb, (dw_contents_prev.packed -
                                        (FieldT(2)^(8*i)) * pb_packing_sum<FieldT>(
                                            pb_variable_array<FieldT>(dw_contents_prev.bits.begin() + 8*i,
                                                                      dw_contents_prev.bits.begin() + 8*(i+1)))));
    }

    /*
      Define masked_out_dw_contents_prev to be the correct masked out
      contents for the current access type.
    */

    pb_linear_combination_array<FieldT> masked_out_indicators;
    masked_out_indicators.emplace_back(is_word0);
    masked_out_indicators.emplace_back(is_word1);
    masked_out_indicators.insert(masked_out_indicators.end(), is_byte.begin(), is_byte.end());

    pb_linear_combination_array<FieldT> masked_out_results;
    masked_out_results.emplace_back(masked_out_word0);
    masked_out_results.emplace_back(masked_out_word1);
    masked_out_results.insert(masked_out_results.end(), masked_out_bytes.begin(), masked_out_bytes.end());

    masked_out_dw_contents_prev.allocate(pb, FMT(this->annotation_prefix, " masked_out_dw_contents_prev"));
    get_masked_out_dw_contents_prev.reset(new inner_product_gadget<FieldT>(pb, masked_out_indicators, masked_out_results, masked_out_dw_contents_prev,
                                                                           FMT(this->annotation_prefix, " get_masked_out_dw_contents_prev")));

    /*
      Define shift so that masked_out_dw_contents_prev + shift * subcontents = dw_contents_next
     */
    linear_combination<FieldT> shift_lc = is_word0 * 1 + is_word1 * (FieldT(2)^this->pb.ap.w);
    for (size_t i = 0; i < 2 * this->pb.ap.bytes_in_word(); ++i)
    {
        shift_lc = shift_lc + is_byte[i] * (FieldT(2)^(8*i));
    }
    shift.assign(pb, shift_lc);
}

template<typename FieldT>
void memory_masking_gadget<FieldT>::generate_r1cs_constraints()
{
    /* get indicator variables for is_subaddress[i] by adding constraints
       is_subaddress[i] * (subaddress - i) = 0 and \sum_i is_subaddress[i] = 1 */
    for (size_t i = 0; i < 2 * this->pb.ap.bytes_in_word(); ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(is_subaddress[i], subaddress.packed - i, 0),
                                     FMT(this->annotation_prefix, " is_subaddress_%zu", i));
    }
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, pb_sum<FieldT>(is_subaddress), 1), FMT(this->annotation_prefix, " is_subaddress"));

    /* get indicator variables is_byte_X */
    for (size_t i = 0; i < 2 * this->pb.ap.bytes_in_word(); ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(access_is_byte, is_subaddress[i], is_byte[i]),
                                     FMT(this->annotation_prefix, " is_byte_%zu", i));
    }

    /* get indicator variables is_word_0/is_word_1 */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(access_is_word, 1 - subaddress.bits[this->pb.ap.subaddr_len()-1], is_word0),
                                 FMT(this->annotation_prefix, " is_word_0"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(access_is_word, subaddress.bits[this->pb.ap.subaddr_len()-1], is_word1),
                                 FMT(this->annotation_prefix, " is_word_1"));

    /* compute masked_out_dw_contents_prev */
    get_masked_out_dw_contents_prev->generate_r1cs_constraints();

    /*
       masked_out_dw_contents_prev + shift * subcontents = dw_contents_next
     */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(shift, subcontents, dw_contents_next.packed - masked_out_dw_contents_prev),
                                 FMT(this->annotation_prefix, " mask_difference"));
}

template<typename FieldT>
void memory_masking_gadget<FieldT>::generate_r1cs_witness()
{
    /* get indicator variables is_subaddress */
    for (size_t i = 0; i < 2 * this->pb.ap.bytes_in_word(); ++i)
    {
        this->pb.val(is_subaddress[i]) = (this->pb.val(subaddress.packed) == FieldT(i)) ? FieldT::one() : FieldT::zero();
    }

    /* get indicator variables is_byte_X */
    for (size_t i = 0; i < 2 * this->pb.ap.bytes_in_word(); ++i)
    {
        this->pb.val(is_byte[i]) = this->pb.val(is_subaddress[i]) * this->pb.lc_val(access_is_byte);
    }

    /* get indicator variables is_word_0/is_word_1 */
    this->pb.val(is_word0) = (FieldT::one() - this->pb.val(subaddress.bits[this->pb.ap.subaddr_len()-1])) * this->pb.lc_val(access_is_word);
    this->pb.val(is_word1) = this->pb.val(subaddress.bits[this->pb.ap.subaddr_len()-1]) * this->pb.lc_val(access_is_word);

    /* calculate shift and masked out words/bytes */
    shift.evaluate(this->pb);
    masked_out_word0.evaluate(this->pb);
    masked_out_word1.evaluate(this->pb);
    masked_out_bytes.evaluate(this->pb);

    /* get masked_out dw/word0/word1/bytes */
    get_masked_out_dw_contents_prev->generate_r1cs_witness();

    /* compute dw_contents_next */
    this->pb.val(dw_contents_next.packed) = this->pb.val(masked_out_dw_contents_prev) + this->pb.lc_val(shift) * this->pb.lc_val(subcontents);
    dw_contents_next.generate_r1cs_witness_from_packed();
}

} // libsnark

#endif // MEMORY_MASKING_GADGET_TCC_

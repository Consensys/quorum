/** @file
 *****************************************************************************

 Implementation of interfaces for an auxiliarry gadget for the FOORAM CPU.

 See bar_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BAR_GADGET_TCC_
#define BAR_GADGET_TCC_

namespace libsnark {

template<typename FieldT>
bar_gadget<FieldT>::bar_gadget(protoboard<FieldT> &pb,
                               const pb_linear_combination_array<FieldT> &X,
                               const FieldT &a,
                               const pb_linear_combination_array<FieldT> &Y,
                               const FieldT &b,
                               const pb_linear_combination<FieldT> &Z_packed,
                               const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    X(X),
    a(a),
    Y(Y),
    b(b),
    Z_packed(Z_packed)
{
    assert(X.size() == Y.size());
    width = X.size();

    result.allocate(pb, FMT(annotation_prefix, " result"));
    Z_bits.allocate(pb, width, FMT(annotation_prefix, " Z_bits"));
    overflow.allocate(pb, 2*width, FMT(annotation_prefix, " overflow"));

    unpacked_result.insert(unpacked_result.end(), Z_bits.begin(), Z_bits.end());
    unpacked_result.insert(unpacked_result.end(), overflow.begin(), overflow.end());

    unpack_result.reset(new packing_gadget<FieldT>(pb, unpacked_result, result, FMT(annotation_prefix, " unpack_result")));
    pack_Z.reset(new packing_gadget<FieldT>(pb, Z_bits, Z_packed, FMT(annotation_prefix, " pack_Z")));
}

template<typename FieldT>
void bar_gadget<FieldT>::generate_r1cs_constraints()
{
    unpack_result->generate_r1cs_constraints(true);
    pack_Z->generate_r1cs_constraints(false);

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, a * pb_packing_sum<FieldT>(X) + b * pb_packing_sum<FieldT>(Y), result), FMT(this->annotation_prefix, " compute_result"));
}

template<typename FieldT>
void bar_gadget<FieldT>::generate_r1cs_witness()
{
    this->pb.val(result) = X.get_field_element_from_bits(this->pb) * a + Y.get_field_element_from_bits(this->pb) * b;
    unpack_result->generate_r1cs_witness_from_packed();

    pack_Z->generate_r1cs_witness_from_bits();
}

} // libsnark

#endif // BAR_GADGET_TCC_

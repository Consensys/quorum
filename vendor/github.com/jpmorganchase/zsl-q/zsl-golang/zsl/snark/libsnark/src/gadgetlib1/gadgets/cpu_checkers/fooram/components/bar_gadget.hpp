/** @file
 *****************************************************************************

 Declaration of interfaces for an auxiliarry gadget for the FOORAM CPU.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BAR_GADGET_HPP_
#define BAR_GADGET_HPP_

#include "gadgetlib1/gadget.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"

namespace libsnark {

/**
 * The bar gadget checks linear combination
 *                   Z = aX + bY (mod 2^w)
 * for a, b - const, X, Y - vectors of w bits,
 * where w is implicitly inferred, Z - a packed variable.
 *
 * This gadget is used four times in fooram:
 * - PC' = PC + 1
 * - load_addr = 2 * x + PC'
 * - store_addr = x + PC
 */
template<typename FieldT>
class bar_gadget : public gadget<FieldT> {
public:
    pb_linear_combination_array<FieldT> X;
    FieldT a;
    pb_linear_combination_array<FieldT> Y;
    FieldT b;
    pb_linear_combination<FieldT> Z_packed;
    pb_variable_array<FieldT> Z_bits;

    pb_variable<FieldT> result;
    pb_variable_array<FieldT> overflow;
    pb_variable_array<FieldT> unpacked_result;

    std::shared_ptr<packing_gadget<FieldT> > unpack_result;
    std::shared_ptr<packing_gadget<FieldT> > pack_Z;

    size_t width;
    bar_gadget(protoboard<FieldT> &pb,
               const pb_linear_combination_array<FieldT> &X,
               const FieldT &a,
               const pb_linear_combination_array<FieldT> &Y,
               const FieldT &b,
               const pb_linear_combination<FieldT> &Z_packed,
               const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/fooram/components/bar_gadget.tcc"

#endif // BAR_GADGET_HPP_

/** @file
 *****************************************************************************

 Declaration of interfaces for Fp2 gadgets.

 The gadgets verify field arithmetic in Fp2 = Fp[U]/(U^2-non_residue),
 where non_residue is in Fp.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FP2_GADGETS_HPP_
#define FP2_GADGETS_HPP_

#include "gadgetlib1/gadget.hpp"
#include <memory>

namespace libsnark {

/**
 * Gadget that represents an Fp2 variable.
 */
template<typename Fp2T>
class Fp2_variable : public gadget<typename Fp2T::my_Fp> {
public:
    typedef typename Fp2T::my_Fp FieldT;

    pb_linear_combination<FieldT> c0;
    pb_linear_combination<FieldT> c1;

    pb_linear_combination_array<FieldT> all_vars;

    Fp2_variable(protoboard<FieldT> &pb,
                 const std::string &annotation_prefix);
    Fp2_variable(protoboard<FieldT> &pb,
                 const Fp2T &el,
                 const std::string &annotation_prefix);
    Fp2_variable(protoboard<FieldT> &pb,
                 const Fp2T &el,
                 const pb_linear_combination<FieldT> &coeff,
                 const std::string &annotation_prefix);
    Fp2_variable(protoboard<FieldT> &pb,
                 const pb_linear_combination<FieldT> &c0,
                 const pb_linear_combination<FieldT> &c1,
                 const std::string &annotation_prefix);

    void generate_r1cs_equals_const_constraints(const Fp2T &el);
    void generate_r1cs_witness(const Fp2T &el);
    Fp2T get_element();

    Fp2_variable<Fp2T> operator*(const FieldT &coeff) const;
    Fp2_variable<Fp2T> operator+(const Fp2_variable<Fp2T> &other) const;
    Fp2_variable<Fp2T> operator+(const Fp2T &other) const;
    Fp2_variable<Fp2T> mul_by_X() const;
    void evaluate() const;
    bool is_constant() const;

    static size_t size_in_bits();
    static size_t num_variables();
};

/**
 * Gadget that creates constraints for Fp2 by Fp2 multiplication.
 */
template<typename Fp2T>
class Fp2_mul_gadget : public gadget<typename Fp2T::my_Fp> {
public:
    typedef typename Fp2T::my_Fp FieldT;

    Fp2_variable<Fp2T> A;
    Fp2_variable<Fp2T> B;
    Fp2_variable<Fp2T> result;

    pb_variable<FieldT> v1;

    Fp2_mul_gadget(protoboard<FieldT> &pb,
                   const Fp2_variable<Fp2T> &A,
                   const Fp2_variable<Fp2T> &B,
                   const Fp2_variable<Fp2T> &result,
                   const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp2 multiplication by a linear combination.
 */
template<typename Fp2T>
class Fp2_mul_by_lc_gadget : public gadget<typename Fp2T::my_Fp> {
public:
    typedef typename Fp2T::my_Fp FieldT;

    Fp2_variable<Fp2T> A;
    pb_linear_combination<FieldT> lc;
    Fp2_variable<Fp2T> result;

    Fp2_mul_by_lc_gadget(protoboard<FieldT> &pb,
                         const Fp2_variable<Fp2T> &A,
                         const pb_linear_combination<FieldT> &lc,
                         const Fp2_variable<Fp2T> &result,
                         const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp2 squaring.
 */
template<typename Fp2T>
class Fp2_sqr_gadget : public gadget<typename Fp2T::my_Fp> {
public:
    typedef typename Fp2T::my_Fp FieldT;

    Fp2_variable<Fp2T> A;
    Fp2_variable<Fp2T> result;

    Fp2_sqr_gadget(protoboard<FieldT> &pb,
                   const Fp2_variable<Fp2T> &A,
                   const Fp2_variable<Fp2T> &result,
                   const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/fields/fp2_gadgets.tcc"

#endif // FP2_GADGETS_HPP_

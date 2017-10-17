/** @file
 *****************************************************************************

 Declaration of interfaces for Fp3 gadgets.

 The gadgets verify field arithmetic in Fp3 = Fp[U]/(U^3-non_residue),
 where non_residue is in Fp.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FP3_GADGETS_HPP_
#define FP3_GADGETS_HPP_

namespace libsnark {

/**
 * Gadget that represents an Fp3 variable.
 */
template<typename Fp3T>
class Fp3_variable : public gadget<typename Fp3T::my_Fp> {
public:
    typedef typename Fp3T::my_Fp FieldT;

    pb_linear_combination<FieldT> c0;
    pb_linear_combination<FieldT> c1;
    pb_linear_combination<FieldT> c2;

    pb_linear_combination_array<FieldT> all_vars;

    Fp3_variable(protoboard<FieldT> &pb,
                 const std::string &annotation_prefix);
    Fp3_variable(protoboard<FieldT> &pb,
                 const Fp3T &el,
                 const std::string &annotation_prefix);
    Fp3_variable(protoboard<FieldT> &pb,
                 const Fp3T &el,
                 const pb_linear_combination<FieldT> &coeff,
                 const std::string &annotation_prefix);
    Fp3_variable(protoboard<FieldT> &pb,
                 const pb_linear_combination<FieldT> &c0,
                 const pb_linear_combination<FieldT> &c1,
                 const pb_linear_combination<FieldT> &c2,
                 const std::string &annotation_prefix);

    void generate_r1cs_equals_const_constraints(const Fp3T &el);
    void generate_r1cs_witness(const Fp3T &el);
    Fp3T get_element();

    Fp3_variable<Fp3T> operator*(const FieldT &coeff) const;
    Fp3_variable<Fp3T> operator+(const Fp3_variable<Fp3T> &other) const;
    Fp3_variable<Fp3T> operator+(const Fp3T &other) const;
    Fp3_variable<Fp3T> mul_by_X() const;
    void evaluate() const;
    bool is_constant() const;

    static size_t size_in_bits();
    static size_t num_variables();
};

/**
 * Gadget that creates constraints for Fp3 by Fp3 multiplication.
 */
template<typename Fp3T>
class Fp3_mul_gadget : public gadget<typename Fp3T::my_Fp> {
public:
    typedef typename Fp3T::my_Fp FieldT;

    Fp3_variable<Fp3T> A;
    Fp3_variable<Fp3T> B;
    Fp3_variable<Fp3T> result;

    pb_variable<FieldT> v0;
    pb_variable<FieldT> v4;

    Fp3_mul_gadget(protoboard<FieldT> &pb,
                   const Fp3_variable<Fp3T> &A,
                   const Fp3_variable<Fp3T> &B,
                   const Fp3_variable<Fp3T> &result,
                   const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp3 multiplication by a linear combination.
 */
template<typename Fp3T>
class Fp3_mul_by_lc_gadget : public gadget<typename Fp3T::my_Fp> {
public:
    typedef typename Fp3T::my_Fp FieldT;

    Fp3_variable<Fp3T> A;
    pb_linear_combination<FieldT> lc;
    Fp3_variable<Fp3T> result;

    Fp3_mul_by_lc_gadget(protoboard<FieldT> &pb,
                         const Fp3_variable<Fp3T> &A,
                         const pb_linear_combination<FieldT> &lc,
                         const Fp3_variable<Fp3T> &result,
                         const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp3 squaring.
 */
template<typename Fp3T>
class Fp3_sqr_gadget : public gadget<typename Fp3T::my_Fp> {
public:
    typedef typename Fp3T::my_Fp FieldT;

    Fp3_variable<Fp3T> A;
    Fp3_variable<Fp3T> result;

    std::shared_ptr<Fp3_mul_gadget<Fp3T> > mul;

    Fp3_sqr_gadget(protoboard<FieldT> &pb,
                   const Fp3_variable<Fp3T> &A,
                   const Fp3_variable<Fp3T> &result,
                   const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};


} // libsnark

#include "gadgetlib1/gadgets/fields/fp3_gadgets.tcc"

#endif // FP3_GADGETS_HPP_

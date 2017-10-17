/** @file
 *****************************************************************************

 Declaration of interfaces for Fp4 gadgets.

 The gadgets verify field arithmetic in Fp4 = Fp2[V]/(V^2-U) where
 Fp2 = Fp[U]/(U^2-non_residue) and non_residue is in Fp.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FP4_GADGETS_HPP_
#define FP4_GADGETS_HPP_

namespace libsnark {

/**
 * Gadget that represents an Fp4 variable.
 */
template<typename Fp4T>
class Fp4_variable : public gadget<typename Fp4T::my_Fp> {
public:
    typedef typename Fp4T::my_Fp FieldT;
    typedef typename Fp4T::my_Fpe Fp2T;

    Fp2_variable<Fp2T> c0;
    Fp2_variable<Fp2T> c1;

    Fp4_variable(protoboard<FieldT> &pb, const std::string &annotation_prefix);
    Fp4_variable(protoboard<FieldT> &pb, const Fp4T &el, const std::string &annotation_prefix);
    Fp4_variable(protoboard<FieldT> &pb, const Fp2_variable<Fp2T> &c0, const Fp2_variable<Fp2T> &c1, const std::string &annotation_prefix);
    void generate_r1cs_equals_const_constraints(const Fp4T &el);
    void generate_r1cs_witness(const Fp4T &el);
    Fp4T get_element();

    Fp4_variable<Fp4T> Frobenius_map(const size_t power) const;
    void evaluate() const;
};

/**
 * Gadget that creates constraints for Fp4 multiplication (towering formulas).
 */
template<typename Fp4T>
class Fp4_tower_mul_gadget : public gadget<typename Fp4T::my_Fp> {
public:
    typedef typename Fp4T::my_Fp FieldT;
    typedef typename Fp4T::my_Fpe Fp2T;

    Fp4_variable<Fp4T> A;
    Fp4_variable<Fp4T> B;
    Fp4_variable<Fp4T> result;

    pb_linear_combination<FieldT> v0_c0;
    pb_linear_combination<FieldT> v0_c1;

    pb_linear_combination<FieldT> Ac0_plus_Ac1_c0;
    pb_linear_combination<FieldT> Ac0_plus_Ac1_c1;
    std::shared_ptr<Fp2_variable<Fp2T> > Ac0_plus_Ac1;

    std::shared_ptr<Fp2_variable<Fp2T> > v0;
    std::shared_ptr<Fp2_variable<Fp2T> > v1;

    pb_linear_combination<FieldT> Bc0_plus_Bc1_c0;
    pb_linear_combination<FieldT> Bc0_plus_Bc1_c1;
    std::shared_ptr<Fp2_variable<Fp2T> > Bc0_plus_Bc1;

    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c0;
    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c1;

    std::shared_ptr<Fp2_variable<Fp2T> > result_c1_plus_v0_plus_v1;

    std::shared_ptr<Fp2_mul_gadget<Fp2T> > compute_v0;
    std::shared_ptr<Fp2_mul_gadget<Fp2T> > compute_v1;
    std::shared_ptr<Fp2_mul_gadget<Fp2T> > compute_result_c1;

    Fp4_tower_mul_gadget(protoboard<FieldT> &pb,
                       const Fp4_variable<Fp4T> &A,
                       const Fp4_variable<Fp4T> &B,
                       const Fp4_variable<Fp4T> &result,
                       const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp4 multiplication (direct formulas).
 */
template<typename Fp4T>
class Fp4_direct_mul_gadget : public gadget<typename Fp4T::my_Fp> {
public:
    typedef typename Fp4T::my_Fp FieldT;
    typedef typename Fp4T::my_Fpe Fp2T;

    Fp4_variable<Fp4T> A;
    Fp4_variable<Fp4T> B;
    Fp4_variable<Fp4T> result;

    pb_variable<FieldT> v1;
    pb_variable<FieldT> v2;
    pb_variable<FieldT> v6;

    Fp4_direct_mul_gadget(protoboard<FieldT> &pb,
                          const Fp4_variable<Fp4T> &A,
                          const Fp4_variable<Fp4T> &B,
                          const Fp4_variable<Fp4T> &result,
                          const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Alias default multiplication gadget
 */
template<typename Fp4T>
using Fp4_mul_gadget = Fp4_direct_mul_gadget<Fp4T>;

/**
 * Gadget that creates constraints for Fp4 squaring.
 */
template<typename Fp4T>
class Fp4_sqr_gadget : public gadget<typename Fp4T::my_Fp> {
public:
    typedef typename Fp4T::my_Fp FieldT;
    typedef typename Fp4T::my_Fpe Fp2T;

    Fp4_variable<Fp4T> A;
    Fp4_variable<Fp4T> result;

    std::shared_ptr<Fp2_variable<Fp2T> > v1;

    pb_linear_combination<FieldT> v0_c0;
    pb_linear_combination<FieldT> v0_c1;
    std::shared_ptr<Fp2_variable<Fp2T> > v0;

    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_v0;
    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_v1;

    pb_linear_combination<FieldT> Ac0_plus_Ac1_c0;
    pb_linear_combination<FieldT> Ac0_plus_Ac1_c1;
    std::shared_ptr<Fp2_variable<Fp2T> > Ac0_plus_Ac1;

    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c0;
    pb_linear_combination<FieldT >result_c1_plus_v0_plus_v1_c1;

    std::shared_ptr<Fp2_variable<Fp2T> > result_c1_plus_v0_plus_v1;

    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_result_c1;

    Fp4_sqr_gadget(protoboard<FieldT> &pb,
                   const Fp4_variable<Fp4T> &A,
                   const Fp4_variable<Fp4T> &result,
                   const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp4 cyclotomic squaring
 */
template<typename Fp4T>
class Fp4_cyclotomic_sqr_gadget : public gadget<typename Fp4T::my_Fp> {
public:
/*
*/
    typedef typename Fp4T::my_Fp FieldT;
    typedef typename Fp4T::my_Fpe Fp2T;

    Fp4_variable<Fp4T> A;
    Fp4_variable<Fp4T> result;

    pb_linear_combination<FieldT> c0_expr_c0;
    pb_linear_combination<FieldT> c0_expr_c1;
    std::shared_ptr<Fp2_variable<Fp2T> > c0_expr;
    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_c0_expr;

    pb_linear_combination<FieldT> A_c0_plus_A_c1_c0;
    pb_linear_combination<FieldT> A_c0_plus_A_c1_c1;
    std::shared_ptr<Fp2_variable<Fp2T> > A_c0_plus_A_c1;

    pb_linear_combination<FieldT> c1_expr_c0;
    pb_linear_combination<FieldT> c1_expr_c1;
    std::shared_ptr<Fp2_variable<Fp2T> > c1_expr;
    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_c1_expr;

    Fp4_cyclotomic_sqr_gadget(protoboard<FieldT> &pb,
                              const Fp4_variable<Fp4T> &A,
                              const Fp4_variable<Fp4T> &result,
                              const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/fields/fp4_gadgets.tcc"

#endif // FP4_GADGETS_HPP_

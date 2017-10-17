/** @file
 *****************************************************************************

 Declaration of interfaces for Fp6 gadgets.

 The gadgets verify field arithmetic in Fp6 = Fp3[Y]/(Y^2-X) where
 Fp3 = Fp[X]/(X^3-non_residue) and non_residue is in Fp.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FP6_GADGETS_HPP_
#define FP6_GADGETS_HPP_

#include "gadgetlib1/gadgets/fields/fp2_gadgets.hpp"

namespace libsnark {

/**
 * Gadget that represents an Fp6 variable.
 */
template<typename Fp6T>
class Fp6_variable : public gadget<typename Fp6T::my_Fp> {
public:
    typedef typename Fp6T::my_Fp FieldT;
    typedef typename Fp6T::my_Fpe Fp3T;

    Fp3_variable<Fp3T> c0;
    Fp3_variable<Fp3T> c1;

    Fp6_variable(protoboard<FieldT> &pb, const std::string &annotation_prefix);
    Fp6_variable(protoboard<FieldT> &pb, const Fp6T &el, const std::string &annotation_prefix);
    Fp6_variable(protoboard<FieldT> &pb, const Fp3_variable<Fp3T> &c0, const Fp3_variable<Fp3T> &c1, const std::string &annotation_prefix);
    void generate_r1cs_equals_const_constraints(const Fp6T &el);
    void generate_r1cs_witness(const Fp6T &el);
    Fp6T get_element();
    Fp6_variable<Fp6T> Frobenius_map(const size_t power) const;
    void evaluate() const;
};

/**
 * Gadget that creates constraints for Fp6 multiplication.
 */
template<typename Fp6T>
class Fp6_mul_gadget : public gadget<typename Fp6T::my_Fp> {
public:
    typedef typename Fp6T::my_Fp FieldT;
    typedef typename Fp6T::my_Fpe Fp3T;

    Fp6_variable<Fp6T> A;
    Fp6_variable<Fp6T> B;
    Fp6_variable<Fp6T> result;

    pb_linear_combination<FieldT> v0_c0;
    pb_linear_combination<FieldT> v0_c1;
    pb_linear_combination<FieldT> v0_c2;

    pb_linear_combination<FieldT> Ac0_plus_Ac1_c0;
    pb_linear_combination<FieldT> Ac0_plus_Ac1_c1;
    pb_linear_combination<FieldT> Ac0_plus_Ac1_c2;
    std::shared_ptr<Fp3_variable<Fp3T> > Ac0_plus_Ac1;

    std::shared_ptr<Fp3_variable<Fp3T> > v0;
    std::shared_ptr<Fp3_variable<Fp3T> > v1;

    pb_linear_combination<FieldT> Bc0_plus_Bc1_c0;
    pb_linear_combination<FieldT> Bc0_plus_Bc1_c1;
    pb_linear_combination<FieldT> Bc0_plus_Bc1_c2;
    std::shared_ptr<Fp3_variable<Fp3T> > Bc0_plus_Bc1;

    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c0;
    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c1;
    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c2;
    std::shared_ptr<Fp3_variable<Fp3T> > result_c1_plus_v0_plus_v1;

    std::shared_ptr<Fp3_mul_gadget<Fp3T> > compute_v0;
    std::shared_ptr<Fp3_mul_gadget<Fp3T> > compute_v1;
    std::shared_ptr<Fp3_mul_gadget<Fp3T> > compute_result_c1;

    Fp6_mul_gadget(protoboard<FieldT> &pb,
                   const Fp6_variable<Fp6T> &A,
                   const Fp6_variable<Fp6T> &B,
                   const Fp6_variable<Fp6T> &result,
                   const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp6 multiplication by a Fp6 element B for which B.c0.c0 = B.c0.c1 = 0.
 */
template<typename Fp6T>
class Fp6_mul_by_2345_gadget : public gadget<typename Fp6T::my_Fp> {
public:
    typedef typename Fp6T::my_Fp FieldT;
    typedef typename Fp6T::my_Fpe Fp3T;

    Fp6_variable<Fp6T> A;
    Fp6_variable<Fp6T> B;
    Fp6_variable<Fp6T> result;

    pb_linear_combination<FieldT> v0_c0;
    pb_linear_combination<FieldT> v0_c1;
    pb_linear_combination<FieldT> v0_c2;

    pb_linear_combination<FieldT> Ac0_plus_Ac1_c0;
    pb_linear_combination<FieldT> Ac0_plus_Ac1_c1;
    pb_linear_combination<FieldT> Ac0_plus_Ac1_c2;
    std::shared_ptr<Fp3_variable<Fp3T> > Ac0_plus_Ac1;

    std::shared_ptr<Fp3_variable<Fp3T> > v0;
    std::shared_ptr<Fp3_variable<Fp3T> > v1;

    pb_linear_combination<FieldT> Bc0_plus_Bc1_c0;
    pb_linear_combination<FieldT> Bc0_plus_Bc1_c1;
    pb_linear_combination<FieldT> Bc0_plus_Bc1_c2;
    std::shared_ptr<Fp3_variable<Fp3T> > Bc0_plus_Bc1;

    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c0;
    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c1;
    pb_linear_combination<FieldT> result_c1_plus_v0_plus_v1_c2;
    std::shared_ptr<Fp3_variable<Fp3T> > result_c1_plus_v0_plus_v1;

    std::shared_ptr<Fp3_mul_gadget<Fp3T> > compute_v1;
    std::shared_ptr<Fp3_mul_gadget<Fp3T> > compute_result_c1;

    Fp6_mul_by_2345_gadget(protoboard<FieldT> &pb,
                           const Fp6_variable<Fp6T> &A,
                           const Fp6_variable<Fp6T> &B,
                           const Fp6_variable<Fp6T> &result,
                           const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp6 squaring.
 */
template<typename Fp6T>
class Fp6_sqr_gadget : public gadget<typename Fp6T::my_Fp> {
public:
    typedef typename Fp6T::my_Fp FieldT;

    Fp6_variable<Fp6T> A;
    Fp6_variable<Fp6T> result;

    std::shared_ptr<Fp6_mul_gadget<Fp6T> > mul;

    Fp6_sqr_gadget(protoboard<FieldT> &pb,
                   const Fp6_variable<Fp6T> &A,
                   const Fp6_variable<Fp6T> &result,
                   const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for Fp6 cyclotomic squaring
 */
template<typename Fp6T>
class Fp6_cyclotomic_sqr_gadget : public gadget<typename Fp6T::my_Fp> {
public:
    typedef typename Fp6T::my_Fp FieldT;
    typedef typename Fp6T::my_Fp2 Fp2T;

    Fp6_variable<Fp6T> A;
    Fp6_variable<Fp6T> result;

    std::shared_ptr<Fp2_variable<Fp2T> > a;
    std::shared_ptr<Fp2_variable<Fp2T> > b;
    std::shared_ptr<Fp2_variable<Fp2T> > c;

    pb_linear_combination<FieldT> asq_c0;
    pb_linear_combination<FieldT> asq_c1;

    pb_linear_combination<FieldT> bsq_c0;
    pb_linear_combination<FieldT> bsq_c1;

    pb_linear_combination<FieldT> csq_c0;
    pb_linear_combination<FieldT> csq_c1;

    std::shared_ptr<Fp2_variable<Fp2T> > asq;
    std::shared_ptr<Fp2_variable<Fp2T> > bsq;
    std::shared_ptr<Fp2_variable<Fp2T> > csq;

    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_asq;
    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_bsq;
    std::shared_ptr<Fp2_sqr_gadget<Fp2T> > compute_csq;

    Fp6_cyclotomic_sqr_gadget(protoboard<FieldT> &pb,
                              const Fp6_variable<Fp6T> &A,
                              const Fp6_variable<Fp6T> &result,
                              const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/fields/fp6_gadgets.tcc"

#endif // FP6_GADGETS_HPP_

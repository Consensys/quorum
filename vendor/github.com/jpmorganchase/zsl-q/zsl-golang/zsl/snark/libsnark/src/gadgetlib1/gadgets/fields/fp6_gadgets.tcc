/** @file
 *****************************************************************************

 Implementation of interfaces for Fp6 gadgets.

 See fp6_gadgets.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FP6_GADGETS_TCC_
#define FP6_GADGETS_TCC_

namespace libsnark {

template<typename Fp6T>
Fp6_variable<Fp6T>::Fp6_variable(protoboard<FieldT> &pb, const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), c0(pb, FMT(annotation_prefix, " c0")), c1(pb, FMT(annotation_prefix, " c1"))
{
}

template<typename Fp6T>
Fp6_variable<Fp6T>::Fp6_variable(protoboard<FieldT> &pb,
                                 const Fp6T &el,
                                 const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), c0(pb, el.c0, FMT(annotation_prefix, " c0")), c1(pb, el.c1, FMT(annotation_prefix, " c1"))
{
}

template<typename Fp6T>
Fp6_variable<Fp6T>::Fp6_variable(protoboard<FieldT> &pb, const Fp3_variable<Fp3T> &c0, const Fp3_variable<Fp3T> &c1, const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), c0(c0), c1(c1)
{
}

template<typename Fp6T>
void Fp6_variable<Fp6T>::generate_r1cs_equals_const_constraints(const Fp6T &el)
{
    c0.generate_r1cs_equals_const_constraints(el.c0);
    c1.generate_r1cs_equals_const_constraints(el.c1);
}

template<typename Fp6T>
void Fp6_variable<Fp6T>::generate_r1cs_witness(const Fp6T &el)
{
    c0.generate_r1cs_witness(el.c0);
    c1.generate_r1cs_witness(el.c1);
}

template<typename Fp6T>
Fp6T Fp6_variable<Fp6T>::get_element()
{
    Fp6T el;
    el.c0 = c0.get_element();
    el.c1 = c1.get_element();
    return el;
}

template<typename Fp6T>
Fp6_variable<Fp6T> Fp6_variable<Fp6T>::Frobenius_map(const size_t power) const
{
    pb_linear_combination<FieldT> new_c0c0, new_c0c1, new_c0c2, new_c1c0, new_c1c1, new_c1c2;
    new_c0c0.assign(this->pb, c0.c0);
    new_c0c1.assign(this->pb, c0.c1 * Fp3T::Frobenius_coeffs_c1[power % 3]);
    new_c0c2.assign(this->pb, c0.c2 * Fp3T::Frobenius_coeffs_c2[power % 3]);
    new_c1c0.assign(this->pb, c1.c0 * Fp6T::Frobenius_coeffs_c1[power % 6]);
    new_c1c1.assign(this->pb, c1.c1 * (Fp6T::Frobenius_coeffs_c1[power % 6] * Fp3T::Frobenius_coeffs_c1[power % 3]));
    new_c1c2.assign(this->pb, c1.c2 * (Fp6T::Frobenius_coeffs_c1[power % 6] * Fp3T::Frobenius_coeffs_c2[power % 3]));

    return Fp6_variable<Fp6T>(this->pb,
                              Fp3_variable<Fp3T>(this->pb, new_c0c0, new_c0c1, new_c0c2, FMT(this->annotation_prefix, " Frobenius_map_c0")),
                              Fp3_variable<Fp3T>(this->pb, new_c1c0, new_c1c1, new_c1c2, FMT(this->annotation_prefix, " Frobenius_map_c1")),
                              FMT(this->annotation_prefix, " Frobenius_map"));
}

template<typename Fp6T>
void Fp6_variable<Fp6T>::evaluate() const
{
    c0.evaluate();
    c1.evaluate();
}

template<typename Fp6T>
Fp6_mul_gadget<Fp6T>::Fp6_mul_gadget(protoboard<FieldT> &pb,
                                     const Fp6_variable<Fp6T> &A,
                                     const Fp6_variable<Fp6T> &B,
                                     const Fp6_variable<Fp6T> &result,
                                     const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), A(A), B(B), result(result)
{
/*
    Karatsuba multiplication for Fp6 as a quadratic extension of Fp3:
        v0 = A.c0 * B.c0
        v1 = A.c1 * B.c1
        result.c0 = v0 + non_residue * v1
        result.c1 = (A.c0 + A.c1) * (B.c0 + B.c1) - v0 - v1
    where "non_residue * elem" := (non_residue * elem.c2, elem.c0, elem.c1)

    Enforced with 3 Fp3_mul_gadget's that ensure that:
        A.c1 * B.c1 = v1
        A.c0 * B.c0 = v0
        (A.c0+A.c1)*(B.c0+B.c1) = result.c1 + v0 + v1

    Reference:
        "Multiplication and Squaring on Pairing-Friendly Fields"
        Devegili, OhEigeartaigh, Scott, Dahab
*/
    v1.reset(new Fp3_variable<Fp3T>(pb, FMT(annotation_prefix, " v1")));

    compute_v1.reset(new Fp3_mul_gadget<Fp3T>(pb, A.c1, B.c1, *v1, FMT(annotation_prefix, " compute_v1")));

    v0_c0.assign(pb, result.c0.c0 - Fp6T::non_residue * v1->c2);
    v0_c1.assign(pb, result.c0.c1 - v1->c0);
    v0_c2.assign(pb, result.c0.c2 - v1->c1);
    v0.reset(new Fp3_variable<Fp3T>(pb, v0_c0, v0_c1, v0_c2, FMT(annotation_prefix, " v0")));

    compute_v0.reset(new Fp3_mul_gadget<Fp3T>(pb, A.c0, B.c0, *v0, FMT(annotation_prefix, " compute_v0")));

    Ac0_plus_Ac1_c0.assign(pb, A.c0.c0 + A.c1.c0);
    Ac0_plus_Ac1_c1.assign(pb, A.c0.c1 + A.c1.c1);
    Ac0_plus_Ac1_c2.assign(pb, A.c0.c2 + A.c1.c2);
    Ac0_plus_Ac1.reset(new Fp3_variable<Fp3T>(pb, Ac0_plus_Ac1_c0, Ac0_plus_Ac1_c1, Ac0_plus_Ac1_c2, FMT(annotation_prefix, " Ac0_plus_Ac1")));

    Bc0_plus_Bc1_c0.assign(pb, B.c0.c0 + B.c1.c0);
    Bc0_plus_Bc1_c1.assign(pb, B.c0.c1 + B.c1.c1);
    Bc0_plus_Bc1_c2.assign(pb, B.c0.c2 + B.c1.c2);
    Bc0_plus_Bc1.reset(new Fp3_variable<Fp3T>(pb, Bc0_plus_Bc1_c0, Bc0_plus_Bc1_c1, Bc0_plus_Bc1_c2, FMT(annotation_prefix, " Bc0_plus_Bc1")));

    result_c1_plus_v0_plus_v1_c0.assign(pb, result.c1.c0 + v0->c0 + v1->c0);
    result_c1_plus_v0_plus_v1_c1.assign(pb, result.c1.c1 + v0->c1 + v1->c1);
    result_c1_plus_v0_plus_v1_c2.assign(pb, result.c1.c2 + v0->c2 + v1->c2);
    result_c1_plus_v0_plus_v1.reset(new Fp3_variable<Fp3T>(pb,
                                                           result_c1_plus_v0_plus_v1_c0,
                                                           result_c1_plus_v0_plus_v1_c1,
                                                           result_c1_plus_v0_plus_v1_c2,
                                                           FMT(annotation_prefix, " result_c1_plus_v0_plus_v1")));

    compute_result_c1.reset(new Fp3_mul_gadget<Fp3T>(pb, *Ac0_plus_Ac1, *Bc0_plus_Bc1, *result_c1_plus_v0_plus_v1, FMT(annotation_prefix, " compute_result_c1")));
}

template<typename Fp6T>
void Fp6_mul_gadget<Fp6T>::generate_r1cs_constraints()
{
    compute_v0->generate_r1cs_constraints();
    compute_v1->generate_r1cs_constraints();
    compute_result_c1->generate_r1cs_constraints();
}

template<typename Fp6T>
void Fp6_mul_gadget<Fp6T>::generate_r1cs_witness()
{
    compute_v0->generate_r1cs_witness();
    compute_v1->generate_r1cs_witness();

    Ac0_plus_Ac1_c0.evaluate(this->pb);
    Ac0_plus_Ac1_c1.evaluate(this->pb);
    Ac0_plus_Ac1_c2.evaluate(this->pb);

    Bc0_plus_Bc1_c0.evaluate(this->pb);
    Bc0_plus_Bc1_c1.evaluate(this->pb);
    Bc0_plus_Bc1_c2.evaluate(this->pb);

    compute_result_c1->generate_r1cs_witness();

    const Fp6T Aval = A.get_element();
    const Fp6T Bval = B.get_element();
    const Fp6T Rval = Aval * Bval;

    result.generate_r1cs_witness(Rval);

    result_c1_plus_v0_plus_v1_c0.evaluate(this->pb);
    result_c1_plus_v0_plus_v1_c1.evaluate(this->pb);
    result_c1_plus_v0_plus_v1_c2.evaluate(this->pb);

    compute_result_c1->generate_r1cs_witness();
}

template<typename Fp6T>
Fp6_mul_by_2345_gadget<Fp6T>::Fp6_mul_by_2345_gadget(protoboard<FieldT> &pb,
                                                     const Fp6_variable<Fp6T> &A,
                                                     const Fp6_variable<Fp6T> &B,
                                                     const Fp6_variable<Fp6T> &result,
                                                     const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), A(A), B(B), result(result)
{
/*
    Karatsuba multiplication for Fp6 as a quadratic extension of Fp3:
        v0 = A.c0 * B.c0
        v1 = A.c1 * B.c1
        result.c0 = v0 + non_residue * v1
        result.c1 = (A.c0 + A.c1) * (B.c0 + B.c1) - v0 - v1
    where "non_residue * elem" := (non_residue * elem.c2, elem.c0, elem.c1)

    We know that B.c0.c0 = B.c0.c1 = 0

    Enforced with 2 Fp3_mul_gadget's that ensure that:
        A.c1 * B.c1 = v1
        (A.c0+A.c1)*(B.c0+B.c1) = result.c1 + v0 + v1

    And one multiplication (three direct constraints) that enforces A.c0 * B.c0
    = v0, where B.c0.c0 = B.c0.c1 = 0.

    Note that (u + v * X + t * X^2) * (0 + 0 * X + z * X^2) =
    (v * z * non_residue + t * z * non_residue * X + u * z * X^2)

    Reference:
        "Multiplication and Squaring on Pairing-Friendly Fields"
        Devegili, OhEigeartaigh, Scott, Dahab
*/
    v1.reset(new Fp3_variable<Fp3T>(pb, FMT(annotation_prefix, " v1")));
    compute_v1.reset(new Fp3_mul_gadget<Fp3T>(pb, A.c1, B.c1, *v1, FMT(annotation_prefix, " compute_v1")));

    /* we inline result.c0 in v0 as follows: v0 = (result.c0.c0 - Fp6T::non_residue * v1->c2, result.c0.c1 - v1->c0, result.c0.c2 - v1->c1) */
    v0.reset(new Fp3_variable<Fp3T>(pb, FMT(annotation_prefix, " v0")));

    Ac0_plus_Ac1_c0.assign(pb, A.c0.c0 + A.c1.c0);
    Ac0_plus_Ac1_c1.assign(pb, A.c0.c1 + A.c1.c1);
    Ac0_plus_Ac1_c2.assign(pb, A.c0.c2 + A.c1.c2);
    Ac0_plus_Ac1.reset(new Fp3_variable<Fp3T>(pb, Ac0_plus_Ac1_c0, Ac0_plus_Ac1_c1, Ac0_plus_Ac1_c2, FMT(annotation_prefix, " Ac0_plus_Ac1")));

    Bc0_plus_Bc1_c0.assign(pb, B.c0.c0 + B.c1.c0);
    Bc0_plus_Bc1_c1.assign(pb, B.c0.c1 + B.c1.c1);
    Bc0_plus_Bc1_c2.assign(pb, B.c0.c2 + B.c1.c2);
    Bc0_plus_Bc1.reset(new Fp3_variable<Fp3T>(pb, Bc0_plus_Bc1_c0, Bc0_plus_Bc1_c1, Bc0_plus_Bc1_c2, FMT(annotation_prefix, " Bc0_plus_Bc1")));

    result_c1_plus_v0_plus_v1_c0.assign(pb, result.c1.c0 + v0->c0 + v1->c0);
    result_c1_plus_v0_plus_v1_c1.assign(pb, result.c1.c1 + v0->c1 + v1->c1);
    result_c1_plus_v0_plus_v1_c2.assign(pb, result.c1.c2 + v0->c2 + v1->c2);
    result_c1_plus_v0_plus_v1.reset(new Fp3_variable<Fp3T>(pb,
                                                           result_c1_plus_v0_plus_v1_c0,
                                                           result_c1_plus_v0_plus_v1_c1,
                                                           result_c1_plus_v0_plus_v1_c2,
                                                           FMT(annotation_prefix, " result_c1_plus_v0_plus_v1")));

    compute_result_c1.reset(new Fp3_mul_gadget<Fp3T>(pb, *Ac0_plus_Ac1, *Bc0_plus_Bc1, *result_c1_plus_v0_plus_v1, FMT(annotation_prefix, " compute_result_c1")));
}

template<typename Fp6T>
void Fp6_mul_by_2345_gadget<Fp6T>::generate_r1cs_constraints()
{
    compute_v1->generate_r1cs_constraints();
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(A.c0.c1,
                                                         Fp3T::non_residue * B.c0.c2,
                                                         result.c0.c0 - Fp6T::non_residue * v1->c2),
                                 FMT(this->annotation_prefix, " v0.c0"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(A.c0.c2,
                                                         Fp3T::non_residue * B.c0.c2,
                                                         result.c0.c1 - v1->c0),
                                 FMT(this->annotation_prefix, " v0.c1"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(A.c0.c0,
                                                         B.c0.c2,
                                                         result.c0.c2 - v1->c1),
                                 FMT(this->annotation_prefix, " v0.c2"));
    compute_result_c1->generate_r1cs_constraints();
}

template<typename Fp6T>
void Fp6_mul_by_2345_gadget<Fp6T>::generate_r1cs_witness()
{
    compute_v1->generate_r1cs_witness();

    const Fp3T A_c0_val = A.c0.get_element();
    const Fp3T B_c0_val = B.c0.get_element();
    assert(B_c0_val.c0.is_zero());
    assert(B_c0_val.c1.is_zero());

    const Fp3T v0_val = A_c0_val * B_c0_val;
    v0->generate_r1cs_witness(v0_val);

    Ac0_plus_Ac1_c0.evaluate(this->pb);
    Ac0_plus_Ac1_c1.evaluate(this->pb);
    Ac0_plus_Ac1_c2.evaluate(this->pb);

    Bc0_plus_Bc1_c0.evaluate(this->pb);
    Bc0_plus_Bc1_c1.evaluate(this->pb);
    Bc0_plus_Bc1_c2.evaluate(this->pb);

    compute_result_c1->generate_r1cs_witness();

    const Fp6T Aval = A.get_element();
    const Fp6T Bval = B.get_element();
    const Fp6T Rval = Aval * Bval;

    result.generate_r1cs_witness(Rval);

    result_c1_plus_v0_plus_v1_c0.evaluate(this->pb);
    result_c1_plus_v0_plus_v1_c1.evaluate(this->pb);
    result_c1_plus_v0_plus_v1_c2.evaluate(this->pb);

    compute_result_c1->generate_r1cs_witness();
}

template<typename Fp6T>
Fp6_sqr_gadget<Fp6T>::Fp6_sqr_gadget(protoboard<FieldT> &pb,
                                     const Fp6_variable<Fp6T> &A,
                                     const Fp6_variable<Fp6T> &result,
                                     const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), A(A), result(result)
{
    // We can't do better than 3 Fp3_mul_gadget's for squaring, so we just use multiplication.
    mul.reset(new Fp6_mul_gadget<Fp6T>(pb, A, A, result, FMT(annotation_prefix, " mul")));
}

template<typename Fp6T>
void Fp6_sqr_gadget<Fp6T>::generate_r1cs_constraints()
{
    mul->generate_r1cs_constraints();
}

template<typename Fp6T>
void Fp6_sqr_gadget<Fp6T>::generate_r1cs_witness()
{
    mul->generate_r1cs_witness();
}

template<typename Fp6T>
Fp6_cyclotomic_sqr_gadget<Fp6T>::Fp6_cyclotomic_sqr_gadget(protoboard<FieldT> &pb,
                                                           const Fp6_variable<Fp6T> &A,
                                                           const Fp6_variable<Fp6T> &result,
                                                           const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), A(A), result(result)
{
/*
    my_Fp2 a = my_Fp2(c0.c0, c1.c1);
    my_Fp2 b = my_Fp2(c1.c0, c0.c2);
    my_Fp2 c = my_Fp2(c0.c1, c1.c2);

    my_Fp2 asq = a.squared();
    my_Fp2 bsq = b.squared();
    my_Fp2 csq = c.squared();

    result.c0.c0 = 3 * asq_a - 2 * a_a;
    result.c1.c1 = 3 * asq_b + 2 * a_b;

    result.c0.c1 = 3 * bsq_a - 2 * c_a;
    result.c1.c2 = 3 * bsq_b + 2 * c_b;

    result.c0.c2 = 3 * csq_a - 2 * b_b;
    result.c1.c0 = 3 * my_Fp3::non_residue * csq_b + 2 * b_a;

    return Fp6_2over3_model<n, mbodulus>(my_Fp3(A_a, C_a, B_b),
                                         my_Fp3(B_a, A_b, C_b))
*/
    a.reset(new Fp2_variable<Fp2T>(pb, A.c0.c0, A.c1.c1, FMT(annotation_prefix, " a")));
    b.reset(new Fp2_variable<Fp2T>(pb, A.c1.c0, A.c0.c2, FMT(annotation_prefix, " b")));
    c.reset(new Fp2_variable<Fp2T>(pb, A.c0.c1, A.c1.c2, FMT(annotation_prefix, " c")));

    asq_c0.assign(pb, (result.c0.c0 + 2 * a->c0) * FieldT(3).inverse());
    asq_c1.assign(pb, (result.c1.c1 - 2 * a->c1) * FieldT(3).inverse());

    bsq_c0.assign(pb, (result.c0.c1 + 2 * c->c0) * FieldT(3).inverse());
    bsq_c1.assign(pb, (result.c1.c2 - 2 * c->c1) * FieldT(3).inverse());

    csq_c0.assign(pb, (result.c0.c2 + 2 * b->c1) * FieldT(3).inverse());
    csq_c1.assign(pb, (result.c1.c0 - 2 * b->c0) * (FieldT(3) * Fp2T::non_residue).inverse());

    asq.reset(new Fp2_variable<Fp2T>(pb, asq_c0, asq_c1, FMT(annotation_prefix, " asq")));
    bsq.reset(new Fp2_variable<Fp2T>(pb, bsq_c0, bsq_c1, FMT(annotation_prefix, " bsq")));
    csq.reset(new Fp2_variable<Fp2T>(pb, csq_c0, csq_c1, FMT(annotation_prefix, " csq")));

    compute_asq.reset(new Fp2_sqr_gadget<Fp2T>(pb, *a, *asq, FMT(annotation_prefix, " compute_asq")));
    compute_bsq.reset(new Fp2_sqr_gadget<Fp2T>(pb, *b, *bsq, FMT(annotation_prefix, " compute_bsq")));
    compute_csq.reset(new Fp2_sqr_gadget<Fp2T>(pb, *c, *csq, FMT(annotation_prefix, " compute_csq")));
}

template<typename Fp6T>
void Fp6_cyclotomic_sqr_gadget<Fp6T>::generate_r1cs_constraints()
{
    compute_asq->generate_r1cs_constraints();
    compute_bsq->generate_r1cs_constraints();
    compute_csq->generate_r1cs_constraints();
}

template<typename Fp6T>
void Fp6_cyclotomic_sqr_gadget<Fp6T>::generate_r1cs_witness()
{
    const Fp6T Aval = A.get_element();
    const Fp6T Rval = Aval.cyclotomic_squared();

    result.generate_r1cs_witness(Rval);

    asq->evaluate();
    bsq->evaluate();
    csq->evaluate();

    compute_asq->generate_r1cs_witness();
    compute_bsq->generate_r1cs_witness();
    compute_csq->generate_r1cs_witness();
}

} // libsnark

#endif // FP6_GADGETS_TCC_

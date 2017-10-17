/** @file
 *****************************************************************************

 Implementation of interfaces for gadgets for Miller loops.

 See weierstrass_miller_loop.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef WEIERSTRASS_MILLER_LOOP_TCC_
#define WEIERSTRASS_MILLER_LOOP_TCC_

#include "algebra/scalar_multiplication/wnaf.hpp"
#include "gadgetlib1/constraint_profiling.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"

namespace libsnark {

/*
  performs

  mnt_Fqk g_RR_at_P = mnt_Fqk(prec_P.PY_twist_squared,
  -prec_P.PX * c.gamma_twist + c.gamma_X - c.old_RY);

  (later in Miller loop: f = f.squared() * g_RR_at_P)
*/

/* Note the slight interface change: this gadget will allocate g_RR_at_P inside itself (!) */
template<typename ppT>
mnt_miller_loop_dbl_line_eval<ppT>::mnt_miller_loop_dbl_line_eval(protoboard<FieldT> &pb,
                                                                  const G1_precomputation<ppT> &prec_P,
                                                                  const precompute_G2_gadget_coeffs<ppT> &c,
                                                                  std::shared_ptr<Fqk_variable<ppT> > &g_RR_at_P,
                                                                  const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), prec_P(prec_P), c(c), g_RR_at_P(g_RR_at_P)
{
    gamma_twist.reset(new Fqe_variable<ppT>(c.gamma->mul_by_X()));
    // prec_P.PX * c.gamma_twist = c.gamma_X - c.old_RY - g_RR_at_P_c1
    if (gamma_twist->is_constant())
    {
        gamma_twist->evaluate();
        const FqeT gamma_twist_const = gamma_twist->get_element();
        g_RR_at_P_c1.reset(new Fqe_variable<ppT>(Fqe_variable<ppT>(this->pb, -gamma_twist_const, prec_P.P->X, FMT(annotation_prefix, " tmp")) +
                                                 *(c.gamma_X) + *(c.RY) * (-FieldT::one())));
    }
    else if (prec_P.P->X.is_constant())
    {
        prec_P.P->X.evaluate(pb);
        const FieldT P_X_const = prec_P.P->X.constant_term();
        g_RR_at_P_c1.reset(new Fqe_variable<ppT>(*gamma_twist * (-P_X_const) + *(c.gamma_X) + *(c.RY) * (-FieldT::one())));
    }
    else
    {
        g_RR_at_P_c1.reset(new Fqe_variable<ppT>(pb, FMT(annotation_prefix, " g_RR_at_P_c1")));
        compute_g_RR_at_P_c1.reset(new Fqe_mul_by_lc_gadget<ppT>(pb, *gamma_twist, prec_P.P->X,
                                                                 *(c.gamma_X) + *(c.RY) * (-FieldT::one()) + (*g_RR_at_P_c1) * (-FieldT::one()),
                                                                 FMT(annotation_prefix, " compute_g_RR_at_P_c1")));
    }
    g_RR_at_P.reset(new Fqk_variable<ppT>(pb, *(prec_P.PY_twist_squared), *g_RR_at_P_c1, FMT(annotation_prefix, " g_RR_at_P")));
}

template<typename ppT>
void mnt_miller_loop_dbl_line_eval<ppT>::generate_r1cs_constraints()
{
    if (!gamma_twist->is_constant() && !prec_P.P->X.is_constant())
    {
        compute_g_RR_at_P_c1->generate_r1cs_constraints();
    }
}

template<typename ppT>
void mnt_miller_loop_dbl_line_eval<ppT>::generate_r1cs_witness()
{
    gamma_twist->evaluate();
    const FqeT gamma_twist_val = gamma_twist->get_element();
    const FieldT PX_val = this->pb.lc_val(prec_P.P->X);
    const FqeT gamma_X_val = c.gamma_X->get_element();
    const FqeT RY_val = c.RY->get_element();
    const FqeT g_RR_at_P_c1_val = -PX_val * gamma_twist_val + gamma_X_val - RY_val;
    g_RR_at_P_c1->generate_r1cs_witness(g_RR_at_P_c1_val);

    if (!gamma_twist->is_constant() && !prec_P.P->X.is_constant())
    {
        compute_g_RR_at_P_c1->generate_r1cs_witness();
    }
    g_RR_at_P->evaluate();
}

/*
  performs
  mnt_Fqk g_RQ_at_P = mnt_Fqk(prec_P.PY_twist_squared,
  -prec_P.PX * c.gamma_twist + c.gamma_X - prec_Q.QY);

  (later in Miller loop: f = f * g_RQ_at_P)

  If invert_Q is set to true: use -QY in place of QY everywhere above.
*/

/* Note the slight interface change: this gadget will allocate g_RQ_at_P inside itself (!) */
template<typename ppT>
mnt_miller_loop_add_line_eval<ppT>::mnt_miller_loop_add_line_eval(protoboard<FieldT> &pb,
                                                                  const bool invert_Q,
                                                                  const G1_precomputation<ppT> &prec_P,
                                                                  const precompute_G2_gadget_coeffs<ppT> &c,
                                                                  const G2_variable<ppT> &Q,
                                                                  std::shared_ptr<Fqk_variable<ppT> > &g_RQ_at_P,
                                                                  const std::string &annotation_prefix) :
gadget<FieldT>(pb, annotation_prefix), invert_Q(invert_Q), prec_P(prec_P), c(c), Q(Q), g_RQ_at_P(g_RQ_at_P)
{
    gamma_twist.reset(new Fqe_variable<ppT>(c.gamma->mul_by_X()));
    // prec_P.PX * c.gamma_twist = c.gamma_X - prec_Q.QY - g_RQ_at_P_c1
    if (gamma_twist->is_constant())
    {
        gamma_twist->evaluate();
        const FqeT gamma_twist_const = gamma_twist->get_element();
        g_RQ_at_P_c1.reset(new Fqe_variable<ppT>(Fqe_variable<ppT>(this->pb, -gamma_twist_const, prec_P.P->X, FMT(annotation_prefix, " tmp")) +
                                                 *(c.gamma_X) + *(Q.Y) * (!invert_Q ? -FieldT::one() : FieldT::one())));
    }
    else if (prec_P.P->X.is_constant())
    {
        prec_P.P->X.evaluate(pb);
        const FieldT P_X_const = prec_P.P->X.constant_term();
        g_RQ_at_P_c1.reset(new Fqe_variable<ppT>(*gamma_twist * (-P_X_const) + *(c.gamma_X) + *(Q.Y) * (!invert_Q ? -FieldT::one() : FieldT::one())));
    }
    else
    {
        g_RQ_at_P_c1.reset(new Fqe_variable<ppT>(pb, FMT(annotation_prefix, " g_RQ_at_Q_c1")));
        compute_g_RQ_at_P_c1.reset(new Fqe_mul_by_lc_gadget<ppT>(pb, *gamma_twist, prec_P.P->X,
                                                                 *(c.gamma_X) + *(Q.Y) * (!invert_Q ? -FieldT::one() : FieldT::one()) + (*g_RQ_at_P_c1) * (-FieldT::one()),
                                                                 FMT(annotation_prefix, " compute_g_RQ_at_P_c1")));
    }
    g_RQ_at_P.reset(new Fqk_variable<ppT>(pb, *(prec_P.PY_twist_squared), *g_RQ_at_P_c1, FMT(annotation_prefix, " g_RQ_at_P")));
}

template<typename ppT>
void mnt_miller_loop_add_line_eval<ppT>::generate_r1cs_constraints()
{
    if (!gamma_twist->is_constant() && !prec_P.P->X.is_constant())
    {
        compute_g_RQ_at_P_c1->generate_r1cs_constraints();
    }
}

template<typename ppT>
void mnt_miller_loop_add_line_eval<ppT>::generate_r1cs_witness()
{
    gamma_twist->evaluate();
    const FqeT gamma_twist_val = gamma_twist->get_element();
    const FieldT PX_val = this->pb.lc_val(prec_P.P->X);
    const FqeT gamma_X_val = c.gamma_X->get_element();
    const FqeT QY_val = Q.Y->get_element();
    const FqeT g_RQ_at_P_c1_val = -PX_val * gamma_twist_val + gamma_X_val + (!invert_Q ? -QY_val : QY_val);
    g_RQ_at_P_c1->generate_r1cs_witness(g_RQ_at_P_c1_val);

    if (!gamma_twist->is_constant() && !prec_P.P->X.is_constant())
    {
        compute_g_RQ_at_P_c1->generate_r1cs_witness();
    }
    g_RQ_at_P->evaluate();
}

template<typename ppT>
mnt_miller_loop_gadget<ppT>::mnt_miller_loop_gadget(protoboard<FieldT> &pb,
                                                    const G1_precomputation<ppT> &prec_P,
                                                    const G2_precomputation<ppT> &prec_Q,
                                                    const Fqk_variable<ppT> &result,
                                                    const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), prec_P(prec_P), prec_Q(prec_Q), result(result)
{
    const auto &loop_count = pairing_selector<ppT>::pairing_loop_count;

    f_count = add_count = dbl_count = 0;

    bool found_nonzero = false;
    std::vector<long> NAF = find_wnaf(1, loop_count);
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        ++dbl_count;
        f_count += 2;

        if (NAF[i] != 0)
        {
            ++add_count;
            f_count += 1;
        }
    }

    fs.resize(f_count);
    doubling_steps.resize(dbl_count);
    addition_steps.resize(add_count);
    g_RR_at_Ps.resize(dbl_count);
    g_RQ_at_Ps.resize(add_count);

    for (size_t i = 0; i < f_count; ++i)
    {
        fs[i].reset(new Fqk_variable<ppT>(pb, FMT(annotation_prefix, " fs_%zu", i)));
    }

    dbl_sqrs.resize(dbl_count);
    dbl_muls.resize(dbl_count);
    add_muls.resize(add_count);

    size_t add_id = 0;
    size_t dbl_id = 0;
    size_t f_id = 0;
    size_t prec_id = 0;

    found_nonzero = false;
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        doubling_steps[dbl_id].reset(new mnt_miller_loop_dbl_line_eval<ppT>(pb,
                                                                            prec_P, *prec_Q.coeffs[prec_id],
                                                                            g_RR_at_Ps[dbl_id],
                                                                            FMT(annotation_prefix, " doubling_steps_%zu", dbl_id)));
        ++prec_id;
        dbl_sqrs[dbl_id].reset(new Fqk_sqr_gadget<ppT>(pb, *fs[f_id], *fs[f_id+1], FMT(annotation_prefix, " dbl_sqrs_%zu", dbl_id)));
        ++f_id;
        dbl_muls[dbl_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RR_at_Ps[dbl_id], (f_id + 1 == f_count ? result : *fs[f_id+1]), FMT(annotation_prefix, " dbl_muls_%zu", dbl_id)));
        ++f_id;
        ++dbl_id;

        if (NAF[i] != 0)
        {
            addition_steps[add_id].reset(new mnt_miller_loop_add_line_eval<ppT>(pb,
                                                                                NAF[i] < 0,
                                                                                prec_P, *prec_Q.coeffs[prec_id], *prec_Q.Q,
                                                                                g_RQ_at_Ps[add_id],
                                                                                FMT(annotation_prefix, " addition_steps_%zu", add_id)));
            ++prec_id;
            add_muls[add_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RQ_at_Ps[add_id], (f_id + 1 == f_count ? result : *fs[f_id+1]), FMT(annotation_prefix, " add_muls_%zu", add_id)));
            ++f_id;
            ++add_id;
        }
    }
}

template<typename ppT>
void mnt_miller_loop_gadget<ppT>::generate_r1cs_constraints()
{
    fs[0]->generate_r1cs_equals_const_constraints(FqkT::one());

    for (size_t i = 0; i < dbl_count; ++i)
    {
        doubling_steps[i]->generate_r1cs_constraints();
        dbl_sqrs[i]->generate_r1cs_constraints();
        dbl_muls[i]->generate_r1cs_constraints();
    }

    for (size_t i = 0; i < add_count; ++i)
    {
        addition_steps[i]->generate_r1cs_constraints();
        add_muls[i]->generate_r1cs_constraints();
    }
}

template<typename ppT>
void mnt_miller_loop_gadget<ppT>::generate_r1cs_witness()
{
    fs[0]->generate_r1cs_witness(FqkT::one());

    size_t add_id = 0;
    size_t dbl_id = 0;

    const auto &loop_count = pairing_selector<ppT>::pairing_loop_count;

    bool found_nonzero = false;
    std::vector<long> NAF = find_wnaf(1, loop_count);
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        doubling_steps[dbl_id]->generate_r1cs_witness();
        dbl_sqrs[dbl_id]->generate_r1cs_witness();
        dbl_muls[dbl_id]->generate_r1cs_witness();
        ++dbl_id;

        if (NAF[i] != 0)
        {
            addition_steps[add_id]->generate_r1cs_witness();
            add_muls[add_id]->generate_r1cs_witness();
            ++add_id;
        }
    }
}

template<typename ppT>
void test_mnt_miller_loop(const std::string &annotation)
{
    protoboard<Fr<ppT> > pb;
    G1<other_curve<ppT> > P_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1_variable<ppT> P(pb, "P");
    G2_variable<ppT> Q(pb, "Q");

    G1_precomputation<ppT> prec_P;
    G2_precomputation<ppT> prec_Q;

    precompute_G1_gadget<ppT> compute_prec_P(pb, P, prec_P, "prec_P");
    precompute_G2_gadget<ppT> compute_prec_Q(pb, Q, prec_Q, "prec_Q");

    Fqk_variable<ppT> result(pb, "result");
    mnt_miller_loop_gadget<ppT> miller(pb, prec_P, prec_Q, result, "miller");

    PROFILE_CONSTRAINTS(pb, "precompute P")
    {
        compute_prec_P.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "precompute Q")
    {
        compute_prec_Q.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "Miller loop")
    {
        miller.generate_r1cs_constraints();
    }
    PRINT_CONSTRAINT_PROFILING();

    P.generate_r1cs_witness(P_val);
    compute_prec_P.generate_r1cs_witness();
    Q.generate_r1cs_witness(Q_val);
    compute_prec_Q.generate_r1cs_witness();
    miller.generate_r1cs_witness();
    assert(pb.is_satisfied());

    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P = other_curve<ppT>::affine_ate_precompute_G1(P_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q = other_curve<ppT>::affine_ate_precompute_G2(Q_val);
    Fqk<other_curve<ppT> > native_result = other_curve<ppT>::affine_ate_miller_loop(native_prec_P, native_prec_Q);

    assert(result.get_element() == native_result);
    printf("number of constraints for Miller loop (Fr is %s)  = %zu\n", annotation.c_str(), pb.num_constraints());
}

template<typename ppT>
mnt_e_over_e_miller_loop_gadget<ppT>::mnt_e_over_e_miller_loop_gadget(protoboard<FieldT> &pb,
                                                                      const G1_precomputation<ppT> &prec_P1,
                                                                      const G2_precomputation<ppT> &prec_Q1,
                                                                      const G1_precomputation<ppT> &prec_P2,
                                                                      const G2_precomputation<ppT> &prec_Q2,
                                                                      const Fqk_variable<ppT> &result,
                                                                      const std::string &annotation_prefix) :
gadget<FieldT>(pb, annotation_prefix), prec_P1(prec_P1), prec_Q1(prec_Q1), prec_P2(prec_P2), prec_Q2(prec_Q2), result(result)
{
    const auto &loop_count = pairing_selector<ppT>::pairing_loop_count;

    f_count = add_count = dbl_count = 0;

    bool found_nonzero = false;
    std::vector<long> NAF = find_wnaf(1, loop_count);
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        ++dbl_count;
        f_count += 3;

        if (NAF[i] != 0)
        {
            ++add_count;
            f_count += 2;
        }
    }

    fs.resize(f_count);
    doubling_steps1.resize(dbl_count);
    addition_steps1.resize(add_count);
    doubling_steps2.resize(dbl_count);
    addition_steps2.resize(add_count);
    g_RR_at_P1s.resize(dbl_count);
    g_RQ_at_P1s.resize(add_count);
    g_RR_at_P2s.resize(dbl_count);
    g_RQ_at_P2s.resize(add_count);

    for (size_t i = 0; i < f_count; ++i)
    {
        fs[i].reset(new Fqk_variable<ppT>(pb, FMT(annotation_prefix, " fs_%zu", i)));
    }

    dbl_sqrs.resize(dbl_count);
    dbl_muls1.resize(dbl_count);
    add_muls1.resize(add_count);
    dbl_muls2.resize(dbl_count);
    add_muls2.resize(add_count);

    size_t add_id = 0;
    size_t dbl_id = 0;
    size_t f_id = 0;
    size_t prec_id = 0;

    found_nonzero = false;
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        doubling_steps1[dbl_id].reset(new mnt_miller_loop_dbl_line_eval<ppT>(pb,
                                                                             prec_P1, *prec_Q1.coeffs[prec_id],
                                                                             g_RR_at_P1s[dbl_id],
                                                                             FMT(annotation_prefix, " doubling_steps1_%zu", dbl_id)));
        doubling_steps2[dbl_id].reset(new mnt_miller_loop_dbl_line_eval<ppT>(pb,
                                                                             prec_P2, *prec_Q2.coeffs[prec_id],
                                                                             g_RR_at_P2s[dbl_id],
                                                                             FMT(annotation_prefix, " doubling_steps2_%zu", dbl_id)));
        ++prec_id;

        dbl_sqrs[dbl_id].reset(new Fqk_sqr_gadget<ppT>(pb, *fs[f_id], *fs[f_id+1], FMT(annotation_prefix, " dbl_sqrs_%zu", dbl_id)));
        ++f_id;
        dbl_muls1[dbl_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RR_at_P1s[dbl_id], *fs[f_id+1], FMT(annotation_prefix, " dbl_mul1s_%zu", dbl_id)));
        ++f_id;
        dbl_muls2[dbl_id].reset(new Fqk_special_mul_gadget<ppT>(pb, (f_id + 1 == f_count ? result : *fs[f_id+1]), *g_RR_at_P2s[dbl_id], *fs[f_id], FMT(annotation_prefix, " dbl_mul2s_%zu", dbl_id)));
        ++f_id;
        ++dbl_id;

        if (NAF[i] != 0)
        {
            addition_steps1[add_id].reset(new mnt_miller_loop_add_line_eval<ppT>(pb,
                                                                                 NAF[i] < 0,
                                                                                 prec_P1, *prec_Q1.coeffs[prec_id], *prec_Q1.Q,
                                                                                 g_RQ_at_P1s[add_id],
                                                                                 FMT(annotation_prefix, " addition_steps1_%zu", add_id)));
            addition_steps2[add_id].reset(new mnt_miller_loop_add_line_eval<ppT>(pb,
                                                                                 NAF[i] < 0,
                                                                                 prec_P2, *prec_Q2.coeffs[prec_id], *prec_Q2.Q,
                                                                                 g_RQ_at_P2s[add_id],
                                                                                 FMT(annotation_prefix, " addition_steps2_%zu", add_id)));
            ++prec_id;
            add_muls1[add_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RQ_at_P1s[add_id], *fs[f_id+1], FMT(annotation_prefix, " add_mul1s_%zu", add_id)));
            ++f_id;
            add_muls2[add_id].reset(new Fqk_special_mul_gadget<ppT>(pb, (f_id + 1 == f_count ? result : *fs[f_id+1]), *g_RQ_at_P2s[add_id], *fs[f_id], FMT(annotation_prefix, " add_mul2s_%zu", add_id)));
            ++f_id;
            ++add_id;
        }
    }
}

template<typename ppT>
void mnt_e_over_e_miller_loop_gadget<ppT>::generate_r1cs_constraints()
{
    fs[0]->generate_r1cs_equals_const_constraints(FqkT::one());

    for (size_t i = 0; i < dbl_count; ++i)
    {
        doubling_steps1[i]->generate_r1cs_constraints();
        doubling_steps2[i]->generate_r1cs_constraints();
        dbl_sqrs[i]->generate_r1cs_constraints();
        dbl_muls1[i]->generate_r1cs_constraints();
        dbl_muls2[i]->generate_r1cs_constraints();
    }

    for (size_t i = 0; i < add_count; ++i)
    {
        addition_steps1[i]->generate_r1cs_constraints();
        addition_steps2[i]->generate_r1cs_constraints();
        add_muls1[i]->generate_r1cs_constraints();
        add_muls2[i]->generate_r1cs_constraints();
    }
}

template<typename ppT>
void mnt_e_over_e_miller_loop_gadget<ppT>::generate_r1cs_witness()
{
    fs[0]->generate_r1cs_witness(FqkT::one());

    size_t add_id = 0;
    size_t dbl_id = 0;
    size_t f_id = 0;

    const auto &loop_count = pairing_selector<ppT>::pairing_loop_count;

    bool found_nonzero = false;
    std::vector<long> NAF = find_wnaf(1, loop_count);
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        doubling_steps1[dbl_id]->generate_r1cs_witness();
        doubling_steps2[dbl_id]->generate_r1cs_witness();
        dbl_sqrs[dbl_id]->generate_r1cs_witness();
        ++f_id;
        dbl_muls1[dbl_id]->generate_r1cs_witness();
        ++f_id;
        (f_id+1 == f_count ? result : *fs[f_id+1]).generate_r1cs_witness(fs[f_id]->get_element() * g_RR_at_P2s[dbl_id]->get_element().inverse());
        dbl_muls2[dbl_id]->generate_r1cs_witness();
        ++f_id;
        ++dbl_id;

        if (NAF[i] != 0)
        {
            addition_steps1[add_id]->generate_r1cs_witness();
            addition_steps2[add_id]->generate_r1cs_witness();
            add_muls1[add_id]->generate_r1cs_witness();
            ++f_id;
            (f_id+1 == f_count ? result : *fs[f_id+1]).generate_r1cs_witness(fs[f_id]->get_element() * g_RQ_at_P2s[add_id]->get_element().inverse());
            add_muls2[add_id]->generate_r1cs_witness();
            ++f_id;
            ++add_id;
        }
    }
}

template<typename ppT>
void test_mnt_e_over_e_miller_loop(const std::string &annotation)
{
    protoboard<Fr<ppT> > pb;
    G1<other_curve<ppT> > P1_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q1_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1<other_curve<ppT> > P2_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q2_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1_variable<ppT> P1(pb, "P1");
    G2_variable<ppT> Q1(pb, "Q1");
    G1_variable<ppT> P2(pb, "P2");
    G2_variable<ppT> Q2(pb, "Q2");

    G1_precomputation<ppT> prec_P1;
    precompute_G1_gadget<ppT> compute_prec_P1(pb, P1, prec_P1, "compute_prec_P1");
    G1_precomputation<ppT> prec_P2;
    precompute_G1_gadget<ppT> compute_prec_P2(pb, P2, prec_P2, "compute_prec_P2");
    G2_precomputation<ppT> prec_Q1;
    precompute_G2_gadget<ppT> compute_prec_Q1(pb, Q1, prec_Q1, "compute_prec_Q1");
    G2_precomputation<ppT> prec_Q2;
    precompute_G2_gadget<ppT> compute_prec_Q2(pb, Q2, prec_Q2, "compute_prec_Q2");

    Fqk_variable<ppT> result(pb, "result");
    mnt_e_over_e_miller_loop_gadget<ppT> miller(pb, prec_P1, prec_Q1, prec_P2, prec_Q2, result, "miller");

    PROFILE_CONSTRAINTS(pb, "precompute P")
    {
        compute_prec_P1.generate_r1cs_constraints();
        compute_prec_P2.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "precompute Q")
    {
        compute_prec_Q1.generate_r1cs_constraints();
        compute_prec_Q2.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "Miller loop")
    {
        miller.generate_r1cs_constraints();
    }
    PRINT_CONSTRAINT_PROFILING();

    P1.generate_r1cs_witness(P1_val);
    compute_prec_P1.generate_r1cs_witness();
    Q1.generate_r1cs_witness(Q1_val);
    compute_prec_Q1.generate_r1cs_witness();
    P2.generate_r1cs_witness(P2_val);
    compute_prec_P2.generate_r1cs_witness();
    Q2.generate_r1cs_witness(Q2_val);
    compute_prec_Q2.generate_r1cs_witness();
    miller.generate_r1cs_witness();
    assert(pb.is_satisfied());

    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P1 = other_curve<ppT>::affine_ate_precompute_G1(P1_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q1 = other_curve<ppT>::affine_ate_precompute_G2(Q1_val);
    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P2 = other_curve<ppT>::affine_ate_precompute_G1(P2_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q2 = other_curve<ppT>::affine_ate_precompute_G2(Q2_val);
    Fqk<other_curve<ppT> > native_result = (other_curve<ppT>::affine_ate_miller_loop(native_prec_P1, native_prec_Q1) *
                                            other_curve<ppT>::affine_ate_miller_loop(native_prec_P2, native_prec_Q2).inverse());

    assert(result.get_element() == native_result);
    printf("number of constraints for e over e Miller loop (Fr is %s)  = %zu\n", annotation.c_str(), pb.num_constraints());
}

template<typename ppT>
mnt_e_times_e_over_e_miller_loop_gadget<ppT>::mnt_e_times_e_over_e_miller_loop_gadget(protoboard<FieldT> &pb,
                                                                                      const G1_precomputation<ppT> &prec_P1,
                                                                                      const G2_precomputation<ppT> &prec_Q1,
                                                                                      const G1_precomputation<ppT> &prec_P2,
                                                                                      const G2_precomputation<ppT> &prec_Q2,
                                                                                      const G1_precomputation<ppT> &prec_P3,
                                                                                      const G2_precomputation<ppT> &prec_Q3,
                                                                                      const Fqk_variable<ppT> &result,
                                                                                      const std::string &annotation_prefix) :
gadget<FieldT>(pb, annotation_prefix), prec_P1(prec_P1), prec_Q1(prec_Q1), prec_P2(prec_P2), prec_Q2(prec_Q2), prec_P3(prec_P3), prec_Q3(prec_Q3), result(result)
{
    const auto &loop_count = pairing_selector<ppT>::pairing_loop_count;

    f_count = add_count = dbl_count = 0;

    bool found_nonzero = false;
    std::vector<long> NAF = find_wnaf(1, loop_count);
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        ++dbl_count;
        f_count += 4;

        if (NAF[i] != 0)
        {
            ++add_count;
            f_count += 3;
        }
    }

    fs.resize(f_count);
    doubling_steps1.resize(dbl_count);
    addition_steps1.resize(add_count);
    doubling_steps2.resize(dbl_count);
    addition_steps2.resize(add_count);
    doubling_steps3.resize(dbl_count);
    addition_steps3.resize(add_count);
    g_RR_at_P1s.resize(dbl_count);
    g_RQ_at_P1s.resize(add_count);
    g_RR_at_P2s.resize(dbl_count);
    g_RQ_at_P2s.resize(add_count);
    g_RR_at_P3s.resize(dbl_count);
    g_RQ_at_P3s.resize(add_count);

    for (size_t i = 0; i < f_count; ++i)
    {
        fs[i].reset(new Fqk_variable<ppT>(pb, FMT(annotation_prefix, " fs_%zu", i)));
    }

    dbl_sqrs.resize(dbl_count);
    dbl_muls1.resize(dbl_count);
    add_muls1.resize(add_count);
    dbl_muls2.resize(dbl_count);
    add_muls2.resize(add_count);
    dbl_muls3.resize(dbl_count);
    add_muls3.resize(add_count);

    size_t add_id = 0;
    size_t dbl_id = 0;
    size_t f_id = 0;
    size_t prec_id = 0;

    found_nonzero = false;
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        doubling_steps1[dbl_id].reset(new mnt_miller_loop_dbl_line_eval<ppT>(pb,
                                                                             prec_P1, *prec_Q1.coeffs[prec_id],
                                                                             g_RR_at_P1s[dbl_id],
                                                                             FMT(annotation_prefix, " doubling_steps1_%zu", dbl_id)));
        doubling_steps2[dbl_id].reset(new mnt_miller_loop_dbl_line_eval<ppT>(pb,
                                                                             prec_P2, *prec_Q2.coeffs[prec_id],
                                                                             g_RR_at_P2s[dbl_id],
                                                                             FMT(annotation_prefix, " doubling_steps2_%zu", dbl_id)));
        doubling_steps3[dbl_id].reset(new mnt_miller_loop_dbl_line_eval<ppT>(pb,
                                                                             prec_P3, *prec_Q3.coeffs[prec_id],
                                                                             g_RR_at_P3s[dbl_id],
                                                                             FMT(annotation_prefix, " doubling_steps3_%zu", dbl_id)));
        ++prec_id;

        dbl_sqrs[dbl_id].reset(new Fqk_sqr_gadget<ppT>(pb, *fs[f_id], *fs[f_id+1], FMT(annotation_prefix, " dbl_sqrs_%zu", dbl_id)));
        ++f_id;
        dbl_muls1[dbl_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RR_at_P1s[dbl_id], *fs[f_id+1], FMT(annotation_prefix, " dbl_muls1_%zu", dbl_id)));
        ++f_id;
        dbl_muls2[dbl_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RR_at_P2s[dbl_id], *fs[f_id+1], FMT(annotation_prefix, " dbl_muls2_%zu", dbl_id)));
        ++f_id;
        dbl_muls3[dbl_id].reset(new Fqk_special_mul_gadget<ppT>(pb, (f_id + 1 == f_count ? result : *fs[f_id+1]), *g_RR_at_P3s[dbl_id], *fs[f_id], FMT(annotation_prefix, " dbl_muls3_%zu", dbl_id)));
        ++f_id;
        ++dbl_id;

        if (NAF[i] != 0)
        {
            addition_steps1[add_id].reset(new mnt_miller_loop_add_line_eval<ppT>(pb,
                                                                                 NAF[i] < 0,
                                                                                 prec_P1, *prec_Q1.coeffs[prec_id], *prec_Q1.Q,
                                                                                 g_RQ_at_P1s[add_id],
                                                                                 FMT(annotation_prefix, " addition_steps1_%zu", add_id)));
            addition_steps2[add_id].reset(new mnt_miller_loop_add_line_eval<ppT>(pb,
                                                                                 NAF[i] < 0,
                                                                                 prec_P2, *prec_Q2.coeffs[prec_id], *prec_Q2.Q,
                                                                                 g_RQ_at_P2s[add_id],
                                                                                 FMT(annotation_prefix, " addition_steps2_%zu", add_id)));
            addition_steps3[add_id].reset(new mnt_miller_loop_add_line_eval<ppT>(pb,
                                                                                 NAF[i] < 0,
                                                                                 prec_P3, *prec_Q3.coeffs[prec_id], *prec_Q3.Q,
                                                                                 g_RQ_at_P3s[add_id],
                                                                                 FMT(annotation_prefix, " addition_steps3_%zu", add_id)));
            ++prec_id;
            add_muls1[add_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RQ_at_P1s[add_id], *fs[f_id+1], FMT(annotation_prefix, " add_muls1_%zu", add_id)));
            ++f_id;
            add_muls2[add_id].reset(new Fqk_special_mul_gadget<ppT>(pb, *fs[f_id], *g_RQ_at_P2s[add_id], *fs[f_id+1], FMT(annotation_prefix, " add_muls2_%zu", add_id)));
            ++f_id;
            add_muls3[add_id].reset(new Fqk_special_mul_gadget<ppT>(pb, (f_id + 1 == f_count ? result : *fs[f_id+1]), *g_RQ_at_P3s[add_id], *fs[f_id], FMT(annotation_prefix, " add_muls3_%zu", add_id)));
            ++f_id;
            ++add_id;
        }
    }
}

template<typename ppT>
void mnt_e_times_e_over_e_miller_loop_gadget<ppT>::generate_r1cs_constraints()
{
    fs[0]->generate_r1cs_equals_const_constraints(FqkT::one());

    for (size_t i = 0; i < dbl_count; ++i)
    {
        doubling_steps1[i]->generate_r1cs_constraints();
        doubling_steps2[i]->generate_r1cs_constraints();
        doubling_steps3[i]->generate_r1cs_constraints();
        dbl_sqrs[i]->generate_r1cs_constraints();
        dbl_muls1[i]->generate_r1cs_constraints();
        dbl_muls2[i]->generate_r1cs_constraints();
        dbl_muls3[i]->generate_r1cs_constraints();
    }

    for (size_t i = 0; i < add_count; ++i)
    {
        addition_steps1[i]->generate_r1cs_constraints();
        addition_steps2[i]->generate_r1cs_constraints();
        addition_steps3[i]->generate_r1cs_constraints();
        add_muls1[i]->generate_r1cs_constraints();
        add_muls2[i]->generate_r1cs_constraints();
        add_muls3[i]->generate_r1cs_constraints();
    }
}

template<typename ppT>
void mnt_e_times_e_over_e_miller_loop_gadget<ppT>::generate_r1cs_witness()
{
    fs[0]->generate_r1cs_witness(FqkT::one());

    size_t add_id = 0;
    size_t dbl_id = 0;
    size_t f_id = 0;

    const auto &loop_count = pairing_selector<ppT>::pairing_loop_count;

    bool found_nonzero = false;
    std::vector<long> NAF = find_wnaf(1, loop_count);
    for (long i = NAF.size()-1; i >= 0; --i)
    {
        if (!found_nonzero)
        {
            /* this skips the MSB itself */
            found_nonzero |= (NAF[i] != 0);
            continue;
        }

        doubling_steps1[dbl_id]->generate_r1cs_witness();
        doubling_steps2[dbl_id]->generate_r1cs_witness();
        doubling_steps3[dbl_id]->generate_r1cs_witness();
        dbl_sqrs[dbl_id]->generate_r1cs_witness();
        ++f_id;
        dbl_muls1[dbl_id]->generate_r1cs_witness();
        ++f_id;
        dbl_muls2[dbl_id]->generate_r1cs_witness();
        ++f_id;
        (f_id+1 == f_count ? result : *fs[f_id+1]).generate_r1cs_witness(fs[f_id]->get_element() * g_RR_at_P3s[dbl_id]->get_element().inverse());
        dbl_muls3[dbl_id]->generate_r1cs_witness();
        ++f_id;
        ++dbl_id;

        if (NAF[i] != 0)
        {
            addition_steps1[add_id]->generate_r1cs_witness();
            addition_steps2[add_id]->generate_r1cs_witness();
            addition_steps3[add_id]->generate_r1cs_witness();
            add_muls1[add_id]->generate_r1cs_witness();
            ++f_id;
            add_muls2[add_id]->generate_r1cs_witness();
            ++f_id;
            (f_id+1 == f_count ? result : *fs[f_id+1]).generate_r1cs_witness(fs[f_id]->get_element() * g_RQ_at_P3s[add_id]->get_element().inverse());
            add_muls3[add_id]->generate_r1cs_witness();
            ++f_id;
            ++add_id;
        }
    }
}

template<typename ppT>
void test_mnt_e_times_e_over_e_miller_loop(const std::string &annotation)
{
    protoboard<Fr<ppT> > pb;
    G1<other_curve<ppT> > P1_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q1_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1<other_curve<ppT> > P2_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q2_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1<other_curve<ppT> > P3_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q3_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1_variable<ppT> P1(pb, "P1");
    G2_variable<ppT> Q1(pb, "Q1");
    G1_variable<ppT> P2(pb, "P2");
    G2_variable<ppT> Q2(pb, "Q2");
    G1_variable<ppT> P3(pb, "P3");
    G2_variable<ppT> Q3(pb, "Q3");

    G1_precomputation<ppT> prec_P1;
    precompute_G1_gadget<ppT> compute_prec_P1(pb, P1, prec_P1, "compute_prec_P1");
    G1_precomputation<ppT> prec_P2;
    precompute_G1_gadget<ppT> compute_prec_P2(pb, P2, prec_P2, "compute_prec_P2");
    G1_precomputation<ppT> prec_P3;
    precompute_G1_gadget<ppT> compute_prec_P3(pb, P3, prec_P3, "compute_prec_P3");
    G2_precomputation<ppT> prec_Q1;
    precompute_G2_gadget<ppT> compute_prec_Q1(pb, Q1, prec_Q1, "compute_prec_Q1");
    G2_precomputation<ppT> prec_Q2;
    precompute_G2_gadget<ppT> compute_prec_Q2(pb, Q2, prec_Q2, "compute_prec_Q2");
    G2_precomputation<ppT> prec_Q3;
    precompute_G2_gadget<ppT> compute_prec_Q3(pb, Q3, prec_Q3, "compute_prec_Q3");

    Fqk_variable<ppT> result(pb, "result");
    mnt_e_times_e_over_e_miller_loop_gadget<ppT> miller(pb, prec_P1, prec_Q1, prec_P2, prec_Q2, prec_P3, prec_Q3, result, "miller");

    PROFILE_CONSTRAINTS(pb, "precompute P")
    {
        compute_prec_P1.generate_r1cs_constraints();
        compute_prec_P2.generate_r1cs_constraints();
        compute_prec_P3.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "precompute Q")
    {
        compute_prec_Q1.generate_r1cs_constraints();
        compute_prec_Q2.generate_r1cs_constraints();
        compute_prec_Q3.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "Miller loop")
    {
        miller.generate_r1cs_constraints();
    }
    PRINT_CONSTRAINT_PROFILING();

    P1.generate_r1cs_witness(P1_val);
    compute_prec_P1.generate_r1cs_witness();
    Q1.generate_r1cs_witness(Q1_val);
    compute_prec_Q1.generate_r1cs_witness();
    P2.generate_r1cs_witness(P2_val);
    compute_prec_P2.generate_r1cs_witness();
    Q2.generate_r1cs_witness(Q2_val);
    compute_prec_Q2.generate_r1cs_witness();
    P3.generate_r1cs_witness(P3_val);
    compute_prec_P3.generate_r1cs_witness();
    Q3.generate_r1cs_witness(Q3_val);
    compute_prec_Q3.generate_r1cs_witness();
    miller.generate_r1cs_witness();
    assert(pb.is_satisfied());

    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P1 = other_curve<ppT>::affine_ate_precompute_G1(P1_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q1 = other_curve<ppT>::affine_ate_precompute_G2(Q1_val);
    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P2 = other_curve<ppT>::affine_ate_precompute_G1(P2_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q2 = other_curve<ppT>::affine_ate_precompute_G2(Q2_val);
    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P3 = other_curve<ppT>::affine_ate_precompute_G1(P3_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q3 = other_curve<ppT>::affine_ate_precompute_G2(Q3_val);
    Fqk<other_curve<ppT> > native_result = (other_curve<ppT>::affine_ate_miller_loop(native_prec_P1, native_prec_Q1) *
                                            other_curve<ppT>::affine_ate_miller_loop(native_prec_P2, native_prec_Q2) *
                                            other_curve<ppT>::affine_ate_miller_loop(native_prec_P3, native_prec_Q3).inverse());

    assert(result.get_element() == native_result);
    printf("number of constraints for e times e over e Miller loop (Fr is %s)  = %zu\n", annotation.c_str(), pb.num_constraints());
}

} // libsnark

#endif // WEIERSTRASS_MILLER_LOOP_TCC_

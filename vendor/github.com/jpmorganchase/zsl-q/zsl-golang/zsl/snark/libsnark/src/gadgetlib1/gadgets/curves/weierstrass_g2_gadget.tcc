/** @file
 *****************************************************************************

 Implementation of interfaces for G2 gadgets.

 See weierstrass_g2_gadgets.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef WEIERSTRASS_G2_GADGET_TCC_
#define WEIERSTRASS_G2_GADGET_TCC_

#include "algebra/scalar_multiplication/wnaf.hpp"

namespace libsnark {

template<typename ppT>
G2_variable<ppT>::G2_variable(protoboard<FieldT> &pb,
                              const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix)
{
    X.reset(new Fqe_variable<ppT>(pb, FMT(annotation_prefix, " X")));
    Y.reset(new Fqe_variable<ppT>(pb, FMT(annotation_prefix, " Y")));

    all_vars.insert(all_vars.end(), X->all_vars.begin(), X->all_vars.end());
    all_vars.insert(all_vars.end(), Y->all_vars.begin(), Y->all_vars.end());
}

template<typename ppT>
G2_variable<ppT>::G2_variable(protoboard<FieldT> &pb,
                              const G2<other_curve<ppT> > &Q,
                              const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix)
{
    G2<other_curve<ppT> > Q_copy = Q;
    Q_copy.to_affine_coordinates();

    X.reset(new Fqe_variable<ppT>(pb, Q_copy.X(), FMT(annotation_prefix, " X")));
    Y.reset(new Fqe_variable<ppT>(pb, Q_copy.Y(), FMT(annotation_prefix, " Y")));

    all_vars.insert(all_vars.end(), X->all_vars.begin(), X->all_vars.end());
    all_vars.insert(all_vars.end(), Y->all_vars.begin(), Y->all_vars.end());
}

template<typename ppT>
void G2_variable<ppT>::generate_r1cs_witness(const G2<other_curve<ppT> > &Q)
{
    G2<other_curve<ppT> > Qcopy = Q;
    Qcopy.to_affine_coordinates();

    X->generate_r1cs_witness(Qcopy.X());
    Y->generate_r1cs_witness(Qcopy.Y());
}

template<typename ppT>
size_t G2_variable<ppT>::size_in_bits()
{
    return 2 * Fqe_variable<ppT>::size_in_bits();
}

template<typename ppT>
size_t G2_variable<ppT>::num_variables()
{
    return 2 * Fqe_variable<ppT>::num_variables();
}

template<typename ppT>
G2_checker_gadget<ppT>::G2_checker_gadget(protoboard<FieldT> &pb,
                                          const G2_variable<ppT> &Q,
                                          const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    Q(Q)
{
    Xsquared.reset(new Fqe_variable<ppT>(pb, FMT(annotation_prefix, " Xsquared")));
    Ysquared.reset(new Fqe_variable<ppT>(pb, FMT(annotation_prefix, " Ysquared")));

    compute_Xsquared.reset(new Fqe_sqr_gadget<ppT>(pb, *(Q.X), *Xsquared, FMT(annotation_prefix, " compute_Xsquared")));
    compute_Ysquared.reset(new Fqe_sqr_gadget<ppT>(pb, *(Q.Y), *Ysquared, FMT(annotation_prefix, " compute_Ysquared")));

    Xsquared_plus_a.reset(new Fqe_variable<ppT>((*Xsquared) + G2<other_curve<ppT> >::coeff_a));
    Ysquared_minus_b.reset(new Fqe_variable<ppT>((*Ysquared) + (-G2<other_curve<ppT> >::coeff_b)));

    curve_equation.reset(new Fqe_mul_gadget<ppT>(pb, *(Q.X), *Xsquared_plus_a, *Ysquared_minus_b, FMT(annotation_prefix, " curve_equation")));
}

template<typename ppT>
void G2_checker_gadget<ppT>::generate_r1cs_constraints()
{
    compute_Xsquared->generate_r1cs_constraints();
    compute_Ysquared->generate_r1cs_constraints();
    curve_equation->generate_r1cs_constraints();
}

template<typename ppT>
void G2_checker_gadget<ppT>::generate_r1cs_witness()
{
    compute_Xsquared->generate_r1cs_witness();
    compute_Ysquared->generate_r1cs_witness();
    Xsquared_plus_a->evaluate();
    curve_equation->generate_r1cs_witness();
}

template<typename ppT>
void test_G2_checker_gadget(const std::string &annotation)
{
    protoboard<Fr<ppT> > pb;
    G2_variable<ppT> g(pb, "g");
    G2_checker_gadget<ppT> g_check(pb, g, "g_check");
    g_check.generate_r1cs_constraints();

    printf("positive test\n");
    g.generate_r1cs_witness(G2<other_curve<ppT> >::one());
    g_check.generate_r1cs_witness();
    assert(pb.is_satisfied());

    printf("negative test\n");
    g.generate_r1cs_witness(G2<other_curve<ppT> >::zero());
    g_check.generate_r1cs_witness();
    assert(!pb.is_satisfied());

    printf("number of constraints for G2 checker (Fr is %s)  = %zu\n", annotation.c_str(), pb.num_constraints());
}

} // libsnark

#endif // WEIERSTRASS_G2_GADGET_TCC_

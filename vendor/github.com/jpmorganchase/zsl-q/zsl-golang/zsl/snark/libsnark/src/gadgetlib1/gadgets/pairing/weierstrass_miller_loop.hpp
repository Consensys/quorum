/** @file
 *****************************************************************************

 Declaration of interfaces for gadgets for Miller loops.

 The gadgets verify computations of (single or multiple simultaneous) Miller loops.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef WEIERSTRASS_MILLER_LOOP_HPP_
#define WEIERSTRASS_MILLER_LOOP_HPP_

#include <memory>
#include "gadgetlib1/gadgets/pairing/pairing_params.hpp"
#include "gadgetlib1/gadgets/pairing/weierstrass_precomputation.hpp"

namespace libsnark {

/**
 * Gadget for doubling step in the Miller loop.
 *
 * Technical note:
 *
 * mnt_Fqk g_RR_at_P = mnt_Fqk(prec_P.PY_twist_squared,
 *                             -prec_P.PX * c.gamma_twist + c.gamma_X - c.old_RY);
 *
 *(later in Miller loop: f = f.squared() * g_RR_at_P)
 *
 * Note the slight interface change: this gadget allocates g_RR_at_P inside itself (!)
 */
template<typename ppT>
class mnt_miller_loop_dbl_line_eval : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;
    typedef Fqe<other_curve<ppT> > FqeT;
    typedef Fqk<other_curve<ppT> > FqkT;

    G1_precomputation<ppT> prec_P;
    precompute_G2_gadget_coeffs<ppT> c;
    std::shared_ptr<Fqk_variable<ppT> > &g_RR_at_P; // reference from outside

    std::shared_ptr<Fqe_variable<ppT> > gamma_twist;
    std::shared_ptr<Fqe_variable<ppT> > g_RR_at_P_c1;
    std::shared_ptr<Fqe_mul_by_lc_gadget<ppT> > compute_g_RR_at_P_c1;

    mnt_miller_loop_dbl_line_eval(protoboard<FieldT> &pb,
                                  const G1_precomputation<ppT> &prec_P,
                                  const precompute_G2_gadget_coeffs<ppT> &c,
                                  std::shared_ptr<Fqk_variable<ppT> > &g_RR_at_P,
                                  const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget for addition step in the Miller loop.
 *
 * Technical note:
 *
 * mnt_Fqk g_RQ_at_P = mnt_Fqk(prec_P.PY_twist_squared,
 *                            -prec_P.PX * c.gamma_twist + c.gamma_X - prec_Q.QY);
 *
 * (later in Miller loop: f = f * g_RQ_at_P)
 *
 * Note the slight interface change: this gadget will allocate g_RQ_at_P inside itself (!)
 */
template<typename ppT>
class mnt_miller_loop_add_line_eval : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;
    typedef Fqe<other_curve<ppT> > FqeT;
    typedef Fqk<other_curve<ppT> > FqkT;

    bool invert_Q;
    G1_precomputation<ppT> prec_P;
    precompute_G2_gadget_coeffs<ppT> c;
    G2_variable<ppT> Q;
    std::shared_ptr<Fqk_variable<ppT> > &g_RQ_at_P; // reference from outside

    std::shared_ptr<Fqe_variable<ppT> > gamma_twist;
    std::shared_ptr<Fqe_variable<ppT> > g_RQ_at_P_c1;
    std::shared_ptr<Fqe_mul_by_lc_gadget<ppT> > compute_g_RQ_at_P_c1;

    mnt_miller_loop_add_line_eval(protoboard<FieldT> &pb,
                                  const bool invert_Q,
                                  const G1_precomputation<ppT> &prec_P,
                                  const precompute_G2_gadget_coeffs<ppT> &c,
                                  const G2_variable<ppT> &Q,
                                  std::shared_ptr<Fqk_variable<ppT> > &g_RQ_at_P,
                                  const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget for verifying a single Miller loop.
 */
template<typename ppT>
class mnt_miller_loop_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;
    typedef Fqe<other_curve<ppT> > FqeT;
    typedef Fqk<other_curve<ppT> > FqkT;

    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RR_at_Ps;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RQ_at_Ps;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > fs;

    std::vector<std::shared_ptr<mnt_miller_loop_add_line_eval<ppT> > > addition_steps;
    std::vector<std::shared_ptr<mnt_miller_loop_dbl_line_eval<ppT> > > doubling_steps;

    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > dbl_muls;
    std::vector<std::shared_ptr<Fqk_sqr_gadget<ppT> > > dbl_sqrs;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > add_muls;

    size_t f_count;
    size_t add_count;
    size_t dbl_count;

    G1_precomputation<ppT> prec_P;
    G2_precomputation<ppT> prec_Q;
    Fqk_variable<ppT> result;

    mnt_miller_loop_gadget(protoboard<FieldT> &pb,
                           const G1_precomputation<ppT> &prec_P,
                           const G2_precomputation<ppT> &prec_Q,
                           const Fqk_variable<ppT> &result,
                           const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename ppT>
void test_mnt_miller_loop(const std::string &annotation);

/**
 * Gadget for verifying a double Miller loop (where the second is inverted).
 */
template<typename ppT>
class mnt_e_over_e_miller_loop_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;
    typedef Fqe<other_curve<ppT> > FqeT;
    typedef Fqk<other_curve<ppT> > FqkT;

    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RR_at_P1s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RQ_at_P1s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RR_at_P2s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RQ_at_P2s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > fs;

    std::vector<std::shared_ptr<mnt_miller_loop_add_line_eval<ppT> > > addition_steps1;
    std::vector<std::shared_ptr<mnt_miller_loop_dbl_line_eval<ppT> > > doubling_steps1;
    std::vector<std::shared_ptr<mnt_miller_loop_add_line_eval<ppT> > > addition_steps2;
    std::vector<std::shared_ptr<mnt_miller_loop_dbl_line_eval<ppT> > > doubling_steps2;

    std::vector<std::shared_ptr<Fqk_sqr_gadget<ppT> > > dbl_sqrs;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > dbl_muls1;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > add_muls1;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > dbl_muls2;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > add_muls2;

    size_t f_count;
    size_t add_count;
    size_t dbl_count;

    G1_precomputation<ppT> prec_P1;
    G2_precomputation<ppT> prec_Q1;
    G1_precomputation<ppT> prec_P2;
    G2_precomputation<ppT> prec_Q2;
    Fqk_variable<ppT> result;

    mnt_e_over_e_miller_loop_gadget(protoboard<FieldT> &pb,
                                    const G1_precomputation<ppT> &prec_P1,
                                    const G2_precomputation<ppT> &prec_Q1,
                                    const G1_precomputation<ppT> &prec_P2,
                                    const G2_precomputation<ppT> &prec_Q2,
                                    const Fqk_variable<ppT> &result,
                                    const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename ppT>
void test_mnt_e_over_e_miller_loop(const std::string &annotation);

/**
 * Gadget for verifying a triple Miller loop (where the third is inverted).
 */
template<typename ppT>
class mnt_e_times_e_over_e_miller_loop_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;
    typedef Fqe<other_curve<ppT> > FqeT;
    typedef Fqk<other_curve<ppT> > FqkT;

    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RR_at_P1s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RQ_at_P1s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RR_at_P2s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RQ_at_P2s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RR_at_P3s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > g_RQ_at_P3s;
    std::vector<std::shared_ptr<Fqk_variable<ppT> > > fs;

    std::vector<std::shared_ptr<mnt_miller_loop_add_line_eval<ppT> > > addition_steps1;
    std::vector<std::shared_ptr<mnt_miller_loop_dbl_line_eval<ppT> > > doubling_steps1;
    std::vector<std::shared_ptr<mnt_miller_loop_add_line_eval<ppT> > > addition_steps2;
    std::vector<std::shared_ptr<mnt_miller_loop_dbl_line_eval<ppT> > > doubling_steps2;
    std::vector<std::shared_ptr<mnt_miller_loop_add_line_eval<ppT> > > addition_steps3;
    std::vector<std::shared_ptr<mnt_miller_loop_dbl_line_eval<ppT> > > doubling_steps3;

    std::vector<std::shared_ptr<Fqk_sqr_gadget<ppT> > > dbl_sqrs;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > dbl_muls1;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > add_muls1;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > dbl_muls2;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > add_muls2;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > dbl_muls3;
    std::vector<std::shared_ptr<Fqk_special_mul_gadget<ppT> > > add_muls3;

    size_t f_count;
    size_t add_count;
    size_t dbl_count;

    G1_precomputation<ppT> prec_P1;
    G2_precomputation<ppT> prec_Q1;
    G1_precomputation<ppT> prec_P2;
    G2_precomputation<ppT> prec_Q2;
    G1_precomputation<ppT> prec_P3;
    G2_precomputation<ppT> prec_Q3;
    Fqk_variable<ppT> result;

    mnt_e_times_e_over_e_miller_loop_gadget(protoboard<FieldT> &pb,
                                            const G1_precomputation<ppT> &prec_P1,
                                            const G2_precomputation<ppT> &prec_Q1,
                                            const G1_precomputation<ppT> &prec_P2,
                                            const G2_precomputation<ppT> &prec_Q2,
                                            const G1_precomputation<ppT> &prec_P3,
                                            const G2_precomputation<ppT> &prec_Q3,
                                            const Fqk_variable<ppT> &result,
                                            const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename ppT>
void test_mnt_e_times_e_over_e_miller_loop(const std::string &annotation);

} // libsnark

#include "gadgetlib1/gadgets/pairing/weierstrass_miller_loop.tcc"

#endif // WEIERSTRASS_MILLER_LOOP_HPP_

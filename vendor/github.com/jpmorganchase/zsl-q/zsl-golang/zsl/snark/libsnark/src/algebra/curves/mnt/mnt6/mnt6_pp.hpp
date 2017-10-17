/** @file
 *****************************************************************************

 Declaration of interfaces for public parameters of MNT6.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT6_PP_HPP_
#define MNT6_PP_HPP_

#include "algebra/curves/public_params.hpp"
#include "algebra/curves/mnt/mnt6/mnt6_init.hpp"
#include "algebra/curves/mnt/mnt6/mnt6_g1.hpp"
#include "algebra/curves/mnt/mnt6/mnt6_g2.hpp"
#include "algebra/curves/mnt/mnt6/mnt6_pairing.hpp"

namespace libsnark {

class mnt6_pp {
public:
    typedef mnt6_Fr Fp_type;
    typedef mnt6_G1 G1_type;
    typedef mnt6_G2 G2_type;
    typedef mnt6_affine_ate_G1_precomputation affine_ate_G1_precomp_type;
    typedef mnt6_affine_ate_G2_precomputation affine_ate_G2_precomp_type;
    typedef mnt6_G1_precomp G1_precomp_type;
    typedef mnt6_G2_precomp G2_precomp_type;
    typedef mnt6_Fq Fq_type;
    typedef mnt6_Fq3 Fqe_type;
    typedef mnt6_Fq6 Fqk_type;
    typedef mnt6_GT GT_type;

    static const bool has_affine_pairing = true;

    static void init_public_params();
    static mnt6_GT final_exponentiation(const mnt6_Fq6 &elt);
    static mnt6_G1_precomp precompute_G1(const mnt6_G1 &P);
    static mnt6_G2_precomp precompute_G2(const mnt6_G2 &Q);
    static mnt6_Fq6 miller_loop(const mnt6_G1_precomp &prec_P,
                                const mnt6_G2_precomp &prec_Q);
    static mnt6_affine_ate_G1_precomputation affine_ate_precompute_G1(const mnt6_G1 &P);
    static mnt6_affine_ate_G2_precomputation affine_ate_precompute_G2(const mnt6_G2 &Q);
    static mnt6_Fq6 affine_ate_miller_loop(const mnt6_affine_ate_G1_precomputation &prec_P,
                                           const mnt6_affine_ate_G2_precomputation &prec_Q);
    static mnt6_Fq6 affine_ate_e_over_e_miller_loop(const mnt6_affine_ate_G1_precomputation &prec_P1,
                                                    const mnt6_affine_ate_G2_precomputation &prec_Q1,
                                                    const mnt6_affine_ate_G1_precomputation &prec_P2,
                                                    const mnt6_affine_ate_G2_precomputation &prec_Q2);
    static mnt6_Fq6 affine_ate_e_times_e_over_e_miller_loop(const mnt6_affine_ate_G1_precomputation &prec_P1,
                                                            const mnt6_affine_ate_G2_precomputation &prec_Q1,
                                                            const mnt6_affine_ate_G1_precomputation &prec_P2,
                                                            const mnt6_affine_ate_G2_precomputation &prec_Q2,
                                                            const mnt6_affine_ate_G1_precomputation &prec_P3,
                                                            const mnt6_affine_ate_G2_precomputation &prec_Q3);
    static mnt6_Fq6 double_miller_loop(const mnt6_G1_precomp &prec_P1,
                                       const mnt6_G2_precomp &prec_Q1,
                                       const mnt6_G1_precomp &prec_P2,
                                       const mnt6_G2_precomp &prec_Q2);

    /* the following are used in test files */
    static mnt6_Fq6 pairing(const mnt6_G1 &P,
                            const mnt6_G2 &Q);
    static mnt6_Fq6 reduced_pairing(const mnt6_G1 &P,
                                    const mnt6_G2 &Q);
    static mnt6_Fq6 affine_reduced_pairing(const mnt6_G1 &P,
                                           const mnt6_G2 &Q);
};

} // libsnark

#endif // MNT6_PP_HPP_

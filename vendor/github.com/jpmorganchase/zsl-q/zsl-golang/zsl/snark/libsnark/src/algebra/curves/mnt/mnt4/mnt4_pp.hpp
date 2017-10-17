/** @file
 *****************************************************************************

 Declaration of interfaces for public parameters of MNT4.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT4_PP_HPP_
#define MNT4_PP_HPP_

#include "algebra/curves/public_params.hpp"
#include "algebra/curves/mnt/mnt4/mnt4_init.hpp"
#include "algebra/curves/mnt/mnt4/mnt4_g1.hpp"
#include "algebra/curves/mnt/mnt4/mnt4_g2.hpp"
#include "algebra/curves/mnt/mnt4/mnt4_pairing.hpp"

namespace libsnark {

class mnt4_pp {
public:
    typedef mnt4_Fr Fp_type;
    typedef mnt4_G1 G1_type;
    typedef mnt4_G2 G2_type;
    typedef mnt4_G1_precomp G1_precomp_type;
    typedef mnt4_G2_precomp G2_precomp_type;
    typedef mnt4_affine_ate_G1_precomputation affine_ate_G1_precomp_type;
    typedef mnt4_affine_ate_G2_precomputation affine_ate_G2_precomp_type;
    typedef mnt4_Fq Fq_type;
    typedef mnt4_Fq2 Fqe_type;
    typedef mnt4_Fq4 Fqk_type;
    typedef mnt4_GT GT_type;

    static const bool has_affine_pairing = true;

    static void init_public_params();
    static mnt4_GT final_exponentiation(const mnt4_Fq4 &elt);

    static mnt4_G1_precomp precompute_G1(const mnt4_G1 &P);
    static mnt4_G2_precomp precompute_G2(const mnt4_G2 &Q);

    static mnt4_Fq4 miller_loop(const mnt4_G1_precomp &prec_P,
                                const mnt4_G2_precomp &prec_Q);

    static mnt4_affine_ate_G1_precomputation affine_ate_precompute_G1(const mnt4_G1 &P);
    static mnt4_affine_ate_G2_precomputation affine_ate_precompute_G2(const mnt4_G2 &Q);
    static mnt4_Fq4 affine_ate_miller_loop(const mnt4_affine_ate_G1_precomputation &prec_P,
                                           const mnt4_affine_ate_G2_precomputation &prec_Q);

    static mnt4_Fq4 affine_ate_e_over_e_miller_loop(const mnt4_affine_ate_G1_precomputation &prec_P1,
                                                    const mnt4_affine_ate_G2_precomputation &prec_Q1,
                                                    const mnt4_affine_ate_G1_precomputation &prec_P2,
                                                    const mnt4_affine_ate_G2_precomputation &prec_Q2);
    static mnt4_Fq4 affine_ate_e_times_e_over_e_miller_loop(const mnt4_affine_ate_G1_precomputation &prec_P1,
                                                            const mnt4_affine_ate_G2_precomputation &prec_Q1,
                                                            const mnt4_affine_ate_G1_precomputation &prec_P2,
                                                            const mnt4_affine_ate_G2_precomputation &prec_Q2,
                                                            const mnt4_affine_ate_G1_precomputation &prec_P3,
                                                            const mnt4_affine_ate_G2_precomputation &prec_Q3);

    static mnt4_Fq4 double_miller_loop(const mnt4_G1_precomp &prec_P1,
                                       const mnt4_G2_precomp &prec_Q1,
                                       const mnt4_G1_precomp &prec_P2,
                                       const mnt4_G2_precomp &prec_Q2);

    /* the following are used in test files */
    static mnt4_Fq4 pairing(const mnt4_G1 &P,
                            const mnt4_G2 &Q);
    static mnt4_Fq4 reduced_pairing(const mnt4_G1 &P,
                                    const mnt4_G2 &Q);
    static mnt4_Fq4 affine_reduced_pairing(const mnt4_G1 &P,
                                           const mnt4_G2 &Q);
};

} // libsnark

#endif // MNT4_PP_HPP_

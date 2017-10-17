/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BN128_PP_HPP_
#define BN128_PP_HPP_
#include "algebra/curves/public_params.hpp"
#include "algebra/curves/bn128/bn128_init.hpp"
#include "algebra/curves/bn128/bn128_g1.hpp"
#include "algebra/curves/bn128/bn128_g2.hpp"
#include "algebra/curves/bn128/bn128_gt.hpp"
#include "algebra/curves/bn128/bn128_pairing.hpp"

namespace libsnark {

class bn128_pp {
public:
    typedef bn128_Fr Fp_type;
    typedef bn128_G1 G1_type;
    typedef bn128_G2 G2_type;
    typedef bn128_ate_G1_precomp G1_precomp_type;
    typedef bn128_ate_G2_precomp G2_precomp_type;
    typedef bn128_Fq Fq_type;
    typedef bn128_Fq12 Fqk_type;
    typedef bn128_GT GT_type;

    static const bool has_affine_pairing = false;

    static void init_public_params();
    static bn128_GT final_exponentiation(const bn128_Fq12 &elt);
    static bn128_ate_G1_precomp precompute_G1(const bn128_G1 &P);
    static bn128_ate_G2_precomp precompute_G2(const bn128_G2 &Q);
    static bn128_Fq12 miller_loop(const bn128_ate_G1_precomp &prec_P,
                                  const bn128_ate_G2_precomp &prec_Q);
    static bn128_Fq12 double_miller_loop(const bn128_ate_G1_precomp &prec_P1,
                                         const bn128_ate_G2_precomp &prec_Q1,
                                         const bn128_ate_G1_precomp &prec_P2,
                                         const bn128_ate_G2_precomp &prec_Q2);

    /* the following are used in test files */
    static bn128_GT pairing(const bn128_G1 &P,
                            const bn128_G2 &Q);
    static bn128_GT reduced_pairing(const bn128_G1 &P,
                                    const bn128_G2 &Q);
};

} // libsnark
#endif // BN128_PP_HPP_

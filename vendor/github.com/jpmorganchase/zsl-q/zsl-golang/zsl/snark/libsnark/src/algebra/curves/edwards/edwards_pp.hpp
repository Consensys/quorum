/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef EDWARDS_PP_HPP_
#define EDWARDS_PP_HPP_
#include "algebra/curves/public_params.hpp"
#include "algebra/curves/edwards/edwards_init.hpp"
#include "algebra/curves/edwards/edwards_g1.hpp"
#include "algebra/curves/edwards/edwards_g2.hpp"
#include "algebra/curves/edwards/edwards_pairing.hpp"

namespace libsnark {

class edwards_pp {
public:
    typedef edwards_Fr Fp_type;
    typedef edwards_G1 G1_type;
    typedef edwards_G2 G2_type;
    typedef edwards_G1_precomp G1_precomp_type;
    typedef edwards_G2_precomp G2_precomp_type;
    typedef edwards_Fq Fq_type;
    typedef edwards_Fq3 Fqe_type;
    typedef edwards_Fq6 Fqk_type;
    typedef edwards_GT GT_type;

    static const bool has_affine_pairing = false;

    static void init_public_params();
    static edwards_GT final_exponentiation(const edwards_Fq6 &elt);
    static edwards_G1_precomp precompute_G1(const edwards_G1 &P);
    static edwards_G2_precomp precompute_G2(const edwards_G2 &Q);
    static edwards_Fq6 miller_loop(const edwards_G1_precomp &prec_P,
                                   const edwards_G2_precomp &prec_Q);
    static edwards_Fq6 double_miller_loop(const edwards_G1_precomp &prec_P1,
                                          const edwards_G2_precomp &prec_Q1,
                                          const edwards_G1_precomp &prec_P2,
                                          const edwards_G2_precomp &prec_Q2);
    /* the following are used in test files */
    static edwards_Fq6 pairing(const edwards_G1 &P,
                               const edwards_G2 &Q);
    static edwards_Fq6 reduced_pairing(const edwards_G1 &P,
                                       const edwards_G2 &Q);
};

} // libsnark
#endif // EDWARDS_PP_HPP_

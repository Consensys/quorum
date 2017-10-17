/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/bn128/bn128_pp.hpp"
#include "common/profiling.hpp"

namespace libsnark {

void bn128_pp::init_public_params()
{
    init_bn128_params();
}

bn128_GT bn128_pp::final_exponentiation(const bn128_GT &elt)
{
    return bn128_final_exponentiation(elt);
}

bn128_ate_G1_precomp bn128_pp::precompute_G1(const bn128_G1 &P)
{
    return bn128_ate_precompute_G1(P);
}

bn128_ate_G2_precomp bn128_pp::precompute_G2(const bn128_G2 &Q)
{
    return bn128_ate_precompute_G2(Q);
}

bn128_Fq12 bn128_pp::miller_loop(const bn128_ate_G1_precomp &prec_P,
                                 const bn128_ate_G2_precomp &prec_Q)
{
    enter_block("Call to miller_loop<bn128_pp>");
    bn128_Fq12 result = bn128_ate_miller_loop(prec_P, prec_Q);
    leave_block("Call to miller_loop<bn128_pp>");
    return result;
}

bn128_Fq12 bn128_pp::double_miller_loop(const bn128_ate_G1_precomp &prec_P1,
                                        const bn128_ate_G2_precomp &prec_Q1,
                                        const bn128_ate_G1_precomp &prec_P2,
                                        const bn128_ate_G2_precomp &prec_Q2)
{
    enter_block("Call to double_miller_loop<bn128_pp>");
    bn128_Fq12 result = bn128_double_ate_miller_loop(prec_P1, prec_Q1, prec_P2, prec_Q2);
    leave_block("Call to double_miller_loop<bn128_pp>");
    return result;
}

bn128_Fq12 bn128_pp::pairing(const bn128_G1 &P,
                             const bn128_G2 &Q)
{
    enter_block("Call to pairing<bn128_pp>");
    bn128_ate_G1_precomp prec_P = bn128_pp::precompute_G1(P);
    bn128_ate_G2_precomp prec_Q = bn128_pp::precompute_G2(Q);

    bn128_Fq12 result = bn128_pp::miller_loop(prec_P, prec_Q);
    leave_block("Call to pairing<bn128_pp>");
    return result;
}

bn128_GT bn128_pp::reduced_pairing(const bn128_G1 &P,
                                   const bn128_G2 &Q)
{
    enter_block("Call to reduced_pairing<bn128_pp>");
    const bn128_Fq12 f = bn128_pp::pairing(P, Q);
    const bn128_GT result = bn128_pp::final_exponentiation(f);
    leave_block("Call to reduced_pairing<bn128_pp>");
    return result;
}

} // libsnark

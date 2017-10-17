/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/edwards/edwards_pp.hpp"

namespace libsnark {

void edwards_pp::init_public_params()
{
    init_edwards_params();
}

edwards_GT edwards_pp::final_exponentiation(const edwards_Fq6 &elt)
{
    return edwards_final_exponentiation(elt);
}

edwards_G1_precomp edwards_pp::precompute_G1(const edwards_G1 &P)
{
    return edwards_precompute_G1(P);
}

edwards_G2_precomp edwards_pp::precompute_G2(const edwards_G2 &Q)
{
    return edwards_precompute_G2(Q);
}

edwards_Fq6 edwards_pp::miller_loop(const edwards_G1_precomp &prec_P,
                                    const edwards_G2_precomp &prec_Q)
{
    return edwards_miller_loop(prec_P, prec_Q);
}

edwards_Fq6 edwards_pp::double_miller_loop(const edwards_G1_precomp &prec_P1,
                                           const edwards_G2_precomp &prec_Q1,
                                           const edwards_G1_precomp &prec_P2,
                                           const edwards_G2_precomp &prec_Q2)
{
    return edwards_double_miller_loop(prec_P1, prec_Q1, prec_P2, prec_Q2);
}

edwards_Fq6 edwards_pp::pairing(const edwards_G1 &P,
                                const edwards_G2 &Q)
{
    return edwards_pairing(P, Q);
}

edwards_Fq6 edwards_pp::reduced_pairing(const edwards_G1 &P,
                                        const edwards_G2 &Q)
{
    return edwards_reduced_pairing(P, Q);
}

} // libsnark

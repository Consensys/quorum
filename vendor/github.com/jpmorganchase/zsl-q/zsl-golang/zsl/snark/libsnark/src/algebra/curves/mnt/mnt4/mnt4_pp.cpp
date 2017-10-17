/** @file
 *****************************************************************************

 Implementation of interfaces for public parameters of MNT4.

 See mnt4_pp.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/mnt/mnt4/mnt4_pp.hpp"

namespace libsnark {

void mnt4_pp::init_public_params()
{
    init_mnt4_params();
}

mnt4_GT mnt4_pp::final_exponentiation(const mnt4_Fq4 &elt)
{
    return mnt4_final_exponentiation(elt);
}

mnt4_G1_precomp mnt4_pp::precompute_G1(const mnt4_G1 &P)
{
    return mnt4_precompute_G1(P);
}

mnt4_G2_precomp mnt4_pp::precompute_G2(const mnt4_G2 &Q)
{
    return mnt4_precompute_G2(Q);
}

mnt4_Fq4 mnt4_pp::miller_loop(const mnt4_G1_precomp &prec_P,
                              const mnt4_G2_precomp &prec_Q)
{
    return mnt4_miller_loop(prec_P, prec_Q);
}

mnt4_affine_ate_G1_precomputation mnt4_pp::affine_ate_precompute_G1(const mnt4_G1 &P)
{
    return mnt4_affine_ate_precompute_G1(P);
}

mnt4_affine_ate_G2_precomputation mnt4_pp::affine_ate_precompute_G2(const mnt4_G2 &Q)
{
    return mnt4_affine_ate_precompute_G2(Q);
}

mnt4_Fq4 mnt4_pp::affine_ate_miller_loop(const mnt4_affine_ate_G1_precomputation &prec_P,
                                         const mnt4_affine_ate_G2_precomputation &prec_Q)
{
    return mnt4_affine_ate_miller_loop(prec_P, prec_Q);
}

mnt4_Fq4 mnt4_pp::affine_ate_e_over_e_miller_loop(const mnt4_affine_ate_G1_precomputation &prec_P1,
                                                  const mnt4_affine_ate_G2_precomputation &prec_Q1,
                                                  const mnt4_affine_ate_G1_precomputation &prec_P2,
                                                  const mnt4_affine_ate_G2_precomputation &prec_Q2)
{
    return mnt4_affine_ate_miller_loop(prec_P1, prec_Q1) * mnt4_affine_ate_miller_loop(prec_P2, prec_Q2).unitary_inverse();
}

mnt4_Fq4 mnt4_pp::affine_ate_e_times_e_over_e_miller_loop(const mnt4_affine_ate_G1_precomputation &prec_P1,
                                                          const mnt4_affine_ate_G2_precomputation &prec_Q1,
                                                          const mnt4_affine_ate_G1_precomputation &prec_P2,
                                                          const mnt4_affine_ate_G2_precomputation &prec_Q2,
                                                          const mnt4_affine_ate_G1_precomputation &prec_P3,
                                                          const mnt4_affine_ate_G2_precomputation &prec_Q3)
{
    return ((mnt4_affine_ate_miller_loop(prec_P1, prec_Q1) * mnt4_affine_ate_miller_loop(prec_P2, prec_Q2)) *
            mnt4_affine_ate_miller_loop(prec_P3, prec_Q3).unitary_inverse());
}

mnt4_Fq4 mnt4_pp::double_miller_loop(const mnt4_G1_precomp &prec_P1,
                                     const mnt4_G2_precomp &prec_Q1,
                                     const mnt4_G1_precomp &prec_P2,
                                     const mnt4_G2_precomp &prec_Q2)
{
    return mnt4_double_miller_loop(prec_P1, prec_Q1, prec_P2, prec_Q2);
}

mnt4_Fq4 mnt4_pp::pairing(const mnt4_G1 &P,
                          const mnt4_G2 &Q)
{
    return mnt4_pairing(P, Q);
}

mnt4_Fq4 mnt4_pp::reduced_pairing(const mnt4_G1 &P,
                                  const mnt4_G2 &Q)
{
    return mnt4_reduced_pairing(P, Q);
}

mnt4_Fq4 mnt4_pp::affine_reduced_pairing(const mnt4_G1 &P,
                                         const mnt4_G2 &Q)
{
    return mnt4_affine_reduced_pairing(P, Q);
}

} // libsnark

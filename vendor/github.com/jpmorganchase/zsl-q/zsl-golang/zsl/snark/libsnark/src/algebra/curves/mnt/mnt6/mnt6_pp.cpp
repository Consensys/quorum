/** @file
 *****************************************************************************

 Implementation of interfaces for public parameters of MNT6.

 See mnt6_pp.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/mnt/mnt6/mnt6_pp.hpp"

namespace libsnark {

void mnt6_pp::init_public_params()
{
    init_mnt6_params();
}

mnt6_GT mnt6_pp::final_exponentiation(const mnt6_Fq6 &elt)
{
    return mnt6_final_exponentiation(elt);
}

mnt6_G1_precomp mnt6_pp::precompute_G1(const mnt6_G1 &P)
{
    return mnt6_precompute_G1(P);
}

mnt6_G2_precomp mnt6_pp::precompute_G2(const mnt6_G2 &Q)
{
    return mnt6_precompute_G2(Q);
}


mnt6_Fq6 mnt6_pp::miller_loop(const mnt6_G1_precomp &prec_P,
                              const mnt6_G2_precomp &prec_Q)
{
    return mnt6_miller_loop(prec_P, prec_Q);
}

mnt6_affine_ate_G1_precomputation mnt6_pp::affine_ate_precompute_G1(const mnt6_G1 &P)
{
    return mnt6_affine_ate_precompute_G1(P);
}

mnt6_affine_ate_G2_precomputation mnt6_pp::affine_ate_precompute_G2(const mnt6_G2 &Q)
{
    return mnt6_affine_ate_precompute_G2(Q);
}

mnt6_Fq6 mnt6_pp::affine_ate_miller_loop(const mnt6_affine_ate_G1_precomputation &prec_P,
                                         const mnt6_affine_ate_G2_precomputation &prec_Q)
{
    return mnt6_affine_ate_miller_loop(prec_P, prec_Q);
}

mnt6_Fq6 mnt6_pp::double_miller_loop(const mnt6_G1_precomp &prec_P1,
                                     const mnt6_G2_precomp &prec_Q1,
                                     const mnt6_G1_precomp &prec_P2,
                                     const mnt6_G2_precomp &prec_Q2)
{
    return mnt6_double_miller_loop(prec_P1, prec_Q1, prec_P2, prec_Q2);
}

mnt6_Fq6 mnt6_pp::affine_ate_e_over_e_miller_loop(const mnt6_affine_ate_G1_precomputation &prec_P1,
                                                  const mnt6_affine_ate_G2_precomputation &prec_Q1,
                                                  const mnt6_affine_ate_G1_precomputation &prec_P2,
                                                  const mnt6_affine_ate_G2_precomputation &prec_Q2)
{
    return mnt6_affine_ate_miller_loop(prec_P1, prec_Q1) * mnt6_affine_ate_miller_loop(prec_P2, prec_Q2).unitary_inverse();
}

mnt6_Fq6 mnt6_pp::affine_ate_e_times_e_over_e_miller_loop(const mnt6_affine_ate_G1_precomputation &prec_P1,
                                                          const mnt6_affine_ate_G2_precomputation &prec_Q1,
                                                          const mnt6_affine_ate_G1_precomputation &prec_P2,
                                                          const mnt6_affine_ate_G2_precomputation &prec_Q2,
                                                          const mnt6_affine_ate_G1_precomputation &prec_P3,
                                                          const mnt6_affine_ate_G2_precomputation &prec_Q3)
{
    return ((mnt6_affine_ate_miller_loop(prec_P1, prec_Q1) * mnt6_affine_ate_miller_loop(prec_P2, prec_Q2)) *
            mnt6_affine_ate_miller_loop(prec_P3, prec_Q3).unitary_inverse());
}

mnt6_Fq6 mnt6_pp::pairing(const mnt6_G1 &P,
                          const mnt6_G2 &Q)
{
    return mnt6_pairing(P, Q);
}

mnt6_Fq6 mnt6_pp::reduced_pairing(const mnt6_G1 &P,
                                  const mnt6_G2 &Q)
{
    return mnt6_reduced_pairing(P, Q);
}

mnt6_Fq6 mnt6_pp::affine_reduced_pairing(const mnt6_G1 &P,
                                         const mnt6_G2 &Q)
{
    return mnt6_affine_reduced_pairing(P, Q);
}

} // libsnark

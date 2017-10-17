/** @file
 *****************************************************************************

 Declaration of interfaces for pairing operations on MNT6.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT6_PAIRING_HPP_
#define MNT6_PAIRING_HPP_

#include <vector>

#include "algebra/curves/mnt/mnt6/mnt6_init.hpp"

namespace libsnark {

/* final exponentiation */

mnt6_Fq6 mnt6_final_exponentiation_last_chunk(const mnt6_Fq6 &elt,
                                              const mnt6_Fq6 &elt_inv);
mnt6_Fq6 mnt6_final_exponentiation_first_chunk(const mnt6_Fq6 &elt,
                                               const mnt6_Fq6 &elt_inv);
mnt6_GT mnt6_final_exponentiation(const mnt6_Fq6 &elt);

/* affine ate miller loop */

struct mnt6_affine_ate_G1_precomputation {
    mnt6_Fq PX;
    mnt6_Fq PY;
    mnt6_Fq3 PY_twist_squared;
};

struct mnt6_affine_ate_coeffs {
    // TODO: trim (not all of them are needed)
    mnt6_Fq3 old_RX;
    mnt6_Fq3 old_RY;
    mnt6_Fq3 gamma;
    mnt6_Fq3 gamma_twist;
    mnt6_Fq3 gamma_X;
};

struct mnt6_affine_ate_G2_precomputation {
    mnt6_Fq3 QX;
    mnt6_Fq3 QY;
    std::vector<mnt6_affine_ate_coeffs> coeffs;
};

mnt6_affine_ate_G1_precomputation mnt6_affine_ate_precompute_G1(const mnt6_G1& P);
mnt6_affine_ate_G2_precomputation mnt6_affine_ate_precompute_G2(const mnt6_G2& Q);

mnt6_Fq6 mnt6_affine_ate_miller_loop(const mnt6_affine_ate_G1_precomputation &prec_P,
                                     const mnt6_affine_ate_G2_precomputation &prec_Q);

/* ate pairing */

struct mnt6_ate_G1_precomp {
    mnt6_Fq PX;
    mnt6_Fq PY;
    mnt6_Fq3 PX_twist;
    mnt6_Fq3 PY_twist;

    bool operator==(const mnt6_ate_G1_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt6_ate_G1_precomp &prec_P);
    friend std::istream& operator>>(std::istream &in, mnt6_ate_G1_precomp &prec_P);
};

struct mnt6_ate_dbl_coeffs {
    mnt6_Fq3 c_H;
    mnt6_Fq3 c_4C;
    mnt6_Fq3 c_J;
    mnt6_Fq3 c_L;

    bool operator==(const mnt6_ate_dbl_coeffs &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt6_ate_dbl_coeffs &dc);
    friend std::istream& operator>>(std::istream &in, mnt6_ate_dbl_coeffs &dc);
};

struct mnt6_ate_add_coeffs {
    mnt6_Fq3 c_L1;
    mnt6_Fq3 c_RZ;

    bool operator==(const mnt6_ate_add_coeffs &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt6_ate_add_coeffs &dc);
    friend std::istream& operator>>(std::istream &in, mnt6_ate_add_coeffs &dc);
};

struct mnt6_ate_G2_precomp {
    mnt6_Fq3 QX;
    mnt6_Fq3 QY;
    mnt6_Fq3 QY2;
    mnt6_Fq3 QX_over_twist;
    mnt6_Fq3 QY_over_twist;
    std::vector<mnt6_ate_dbl_coeffs> dbl_coeffs;
    std::vector<mnt6_ate_add_coeffs> add_coeffs;

    bool operator==(const mnt6_ate_G2_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt6_ate_G2_precomp &prec_Q);
    friend std::istream& operator>>(std::istream &in, mnt6_ate_G2_precomp &prec_Q);
};

mnt6_ate_G1_precomp mnt6_ate_precompute_G1(const mnt6_G1& P);
mnt6_ate_G2_precomp mnt6_ate_precompute_G2(const mnt6_G2& Q);

mnt6_Fq6 mnt6_ate_miller_loop(const mnt6_ate_G1_precomp &prec_P,
                              const mnt6_ate_G2_precomp &prec_Q);
mnt6_Fq6 mnt6_ate_double_miller_loop(const mnt6_ate_G1_precomp &prec_P1,
                                     const mnt6_ate_G2_precomp &prec_Q1,
                                     const mnt6_ate_G1_precomp &prec_P2,
                                     const mnt6_ate_G2_precomp &prec_Q2);

mnt6_Fq6 mnt6_ate_pairing(const mnt6_G1& P,
                          const mnt6_G2 &Q);
mnt6_GT mnt6_ate_reduced_pairing(const mnt6_G1 &P,
                                 const mnt6_G2 &Q);

/* choice of pairing */

typedef mnt6_ate_G1_precomp mnt6_G1_precomp;
typedef mnt6_ate_G2_precomp mnt6_G2_precomp;

mnt6_G1_precomp mnt6_precompute_G1(const mnt6_G1& P);

mnt6_G2_precomp mnt6_precompute_G2(const mnt6_G2& Q);

mnt6_Fq6 mnt6_miller_loop(const mnt6_G1_precomp &prec_P,
                          const mnt6_G2_precomp &prec_Q);

mnt6_Fq6 mnt6_double_miller_loop(const mnt6_G1_precomp &prec_P1,
                                 const mnt6_G2_precomp &prec_Q1,
                                 const mnt6_G1_precomp &prec_P2,
                                 const mnt6_G2_precomp &prec_Q2);

mnt6_Fq6 mnt6_pairing(const mnt6_G1& P,
                      const mnt6_G2 &Q);

mnt6_GT mnt6_reduced_pairing(const mnt6_G1 &P,
                             const mnt6_G2 &Q);

mnt6_GT mnt6_affine_reduced_pairing(const mnt6_G1 &P,
                                    const mnt6_G2 &Q);

} // libsnark

#endif // MNT6_PAIRING_HPP_

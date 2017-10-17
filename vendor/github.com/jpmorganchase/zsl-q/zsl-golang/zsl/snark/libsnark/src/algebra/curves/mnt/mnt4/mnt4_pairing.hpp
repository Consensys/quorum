/** @file
 *****************************************************************************

 Declaration of interfaces for pairing operations on MNT4.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT4_PAIRING_HPP_
#define MNT4_PAIRING_HPP_

#include <vector>

#include "algebra/curves/mnt/mnt4/mnt4_init.hpp"

namespace libsnark {

/* final exponentiation */

mnt4_Fq4 mnt4_final_exponentiation_last_chunk(const mnt4_Fq4 &elt,
                                              const mnt4_Fq4 &elt_inv);
mnt4_Fq4 mnt4_final_exponentiation_first_chunk(const mnt4_Fq4 &elt,
                                               const mnt4_Fq4 &elt_inv);
mnt4_GT mnt4_final_exponentiation(const mnt4_Fq4 &elt);

/* affine ate miller loop */

struct mnt4_affine_ate_G1_precomputation {
    mnt4_Fq PX;
    mnt4_Fq PY;
    mnt4_Fq2 PY_twist_squared;
};

struct mnt4_affine_ate_coeffs {
    // TODO: trim (not all of them are needed)
    mnt4_Fq2 old_RX;
    mnt4_Fq2 old_RY;
    mnt4_Fq2 gamma;
    mnt4_Fq2 gamma_twist;
    mnt4_Fq2 gamma_X;
};

struct mnt4_affine_ate_G2_precomputation {
    mnt4_Fq2 QX;
    mnt4_Fq2 QY;
    std::vector<mnt4_affine_ate_coeffs> coeffs;
};

mnt4_affine_ate_G1_precomputation mnt4_affine_ate_precompute_G1(const mnt4_G1& P);
mnt4_affine_ate_G2_precomputation mnt4_affine_ate_precompute_G2(const mnt4_G2& Q);

mnt4_Fq4 mnt4_affine_ate_miller_loop(const mnt4_affine_ate_G1_precomputation &prec_P,
                                     const mnt4_affine_ate_G2_precomputation &prec_Q);

/* ate pairing */

struct mnt4_ate_G1_precomp {
    mnt4_Fq PX;
    mnt4_Fq PY;
    mnt4_Fq2 PX_twist;
    mnt4_Fq2 PY_twist;

    bool operator==(const mnt4_ate_G1_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt4_ate_G1_precomp &prec_P);
    friend std::istream& operator>>(std::istream &in, mnt4_ate_G1_precomp &prec_P);
};

struct mnt4_ate_dbl_coeffs {
    mnt4_Fq2 c_H;
    mnt4_Fq2 c_4C;
    mnt4_Fq2 c_J;
    mnt4_Fq2 c_L;

    bool operator==(const mnt4_ate_dbl_coeffs &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt4_ate_dbl_coeffs &dc);
    friend std::istream& operator>>(std::istream &in, mnt4_ate_dbl_coeffs &dc);
};

struct mnt4_ate_add_coeffs {
    mnt4_Fq2 c_L1;
    mnt4_Fq2 c_RZ;

    bool operator==(const mnt4_ate_add_coeffs &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt4_ate_add_coeffs &dc);
    friend std::istream& operator>>(std::istream &in, mnt4_ate_add_coeffs &dc);
};

struct mnt4_ate_G2_precomp {
    mnt4_Fq2 QX;
    mnt4_Fq2 QY;
    mnt4_Fq2 QY2;
    mnt4_Fq2 QX_over_twist;
    mnt4_Fq2 QY_over_twist;
    std::vector<mnt4_ate_dbl_coeffs> dbl_coeffs;
    std::vector<mnt4_ate_add_coeffs> add_coeffs;

    bool operator==(const mnt4_ate_G2_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const mnt4_ate_G2_precomp &prec_Q);
    friend std::istream& operator>>(std::istream &in, mnt4_ate_G2_precomp &prec_Q);
};

mnt4_ate_G1_precomp mnt4_ate_precompute_G1(const mnt4_G1& P);
mnt4_ate_G2_precomp mnt4_ate_precompute_G2(const mnt4_G2& Q);

mnt4_Fq4 mnt4_ate_miller_loop(const mnt4_ate_G1_precomp &prec_P,
                                    const mnt4_ate_G2_precomp &prec_Q);
mnt4_Fq4 mnt4_ate_double_miller_loop(const mnt4_ate_G1_precomp &prec_P1,
                                           const mnt4_ate_G2_precomp &prec_Q1,
                                           const mnt4_ate_G1_precomp &prec_P2,
                                           const mnt4_ate_G2_precomp &prec_Q2);

mnt4_Fq4 mnt4_ate_pairing(const mnt4_G1& P,
                          const mnt4_G2 &Q);
mnt4_GT mnt4_ate_reduced_pairing(const mnt4_G1 &P,
                                 const mnt4_G2 &Q);

/* choice of pairing */

typedef mnt4_ate_G1_precomp mnt4_G1_precomp;
typedef mnt4_ate_G2_precomp mnt4_G2_precomp;

mnt4_G1_precomp mnt4_precompute_G1(const mnt4_G1& P);

mnt4_G2_precomp mnt4_precompute_G2(const mnt4_G2& Q);

mnt4_Fq4 mnt4_miller_loop(const mnt4_G1_precomp &prec_P,
                          const mnt4_G2_precomp &prec_Q);

mnt4_Fq4 mnt4_double_miller_loop(const mnt4_G1_precomp &prec_P1,
                                 const mnt4_G2_precomp &prec_Q1,
                                 const mnt4_G1_precomp &prec_P2,
                                 const mnt4_G2_precomp &prec_Q2);

mnt4_Fq4 mnt4_pairing(const mnt4_G1& P,
                      const mnt4_G2 &Q);

mnt4_GT mnt4_reduced_pairing(const mnt4_G1 &P,
                             const mnt4_G2 &Q);

mnt4_GT mnt4_affine_reduced_pairing(const mnt4_G1 &P,
                                    const mnt4_G2 &Q);

} // libsnark

#endif // MNT4_PAIRING_HPP_

/** @file
 ********************************************************************************
 Declares functions for computing Ate pairings over the bn128 curves, split into a
 offline and online stages.
 ********************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *******************************************************************************/

#ifndef BN128_PAIRING_HPP_
#define BN128_PAIRING_HPP_
#include "algebra/curves/bn128/bn128_g1.hpp"
#include "algebra/curves/bn128/bn128_g2.hpp"
#include "algebra/curves/bn128/bn128_gt.hpp"
#include "bn.h"

namespace libsnark {

struct bn128_ate_G1_precomp {
    bn::Fp P[3];

    bool operator==(const bn128_ate_G1_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const bn128_ate_G1_precomp &prec_P);
    friend std::istream& operator>>(std::istream &in, bn128_ate_G1_precomp &prec_P);
};

typedef bn::Fp6 bn128_ate_ell_coeffs;

struct bn128_ate_G2_precomp {
    bn::Fp2 Q[3];
    std::vector<bn128_ate_ell_coeffs> coeffs;

    bool operator==(const bn128_ate_G2_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const bn128_ate_G2_precomp &prec_Q);
    friend std::istream& operator>>(std::istream &in, bn128_ate_G2_precomp &prec_Q);
};

bn128_ate_G1_precomp bn128_ate_precompute_G1(const bn128_G1& P);
bn128_ate_G2_precomp bn128_ate_precompute_G2(const bn128_G2& Q);

bn128_Fq12 bn128_double_ate_miller_loop(const bn128_ate_G1_precomp &prec_P1,
                                        const bn128_ate_G2_precomp &prec_Q1,
                                        const bn128_ate_G1_precomp &prec_P2,
                                        const bn128_ate_G2_precomp &prec_Q2);
bn128_Fq12 bn128_ate_miller_loop(const bn128_ate_G1_precomp &prec_P,
                                 const bn128_ate_G2_precomp &prec_Q);

bn128_GT bn128_final_exponentiation(const bn128_Fq12 &elt);

} // libsnark
#endif // BN128_PAIRING_HPP_

/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef EDWARDS_PAIRING_HPP_
#define EDWARDS_PAIRING_HPP_
#include <vector>
#include "algebra/curves/edwards/edwards_init.hpp"

namespace libsnark {

/* final exponentiation */

edwards_Fq6 edwards_final_exponentiation_last_chunk(const edwards_Fq6 &elt,
                                                    const edwards_Fq6 &elt_inv);
edwards_Fq6 edwards_final_exponentiation_first_chunk(const edwards_Fq6 &elt,
                                                     const edwards_Fq6 &elt_inv);
edwards_GT edwards_final_exponentiation(const edwards_Fq6 &elt);

/* Tate pairing */

struct edwards_Fq_conic_coefficients {
    edwards_Fq c_ZZ;
    edwards_Fq c_XY;
    edwards_Fq c_XZ;

    bool operator==(const edwards_Fq_conic_coefficients &other) const;
    friend std::ostream& operator<<(std::ostream &out, const edwards_Fq_conic_coefficients &cc);
    friend std::istream& operator>>(std::istream &in, edwards_Fq_conic_coefficients &cc);
};
typedef std::vector<edwards_Fq_conic_coefficients> edwards_tate_G1_precomp;

std::ostream& operator<<(std::ostream& out, const edwards_tate_G1_precomp &prec_P);
std::istream& operator>>(std::istream& in, edwards_tate_G1_precomp &prec_P);

struct edwards_tate_G2_precomp {
    edwards_Fq3 y0, eta;

    bool operator==(const edwards_tate_G2_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const edwards_tate_G2_precomp &prec_Q);
    friend std::istream& operator>>(std::istream &in, edwards_tate_G2_precomp &prec_Q);
};

edwards_tate_G1_precomp edwards_tate_precompute_G1(const edwards_G1& P);
edwards_tate_G2_precomp edwards_tate_precompute_G2(const edwards_G2& Q);

edwards_Fq6 edwards_tate_miller_loop(const edwards_tate_G1_precomp &prec_P,
                                     const edwards_tate_G2_precomp &prec_Q);

edwards_Fq6 edwards_tate_pairing(const edwards_G1& P,
                                 const edwards_G2 &Q);
edwards_GT edwards_tate_reduced_pairing(const edwards_G1 &P,
                                        const edwards_G2 &Q);

/* ate pairing */

struct edwards_Fq3_conic_coefficients {
    edwards_Fq3 c_ZZ;
    edwards_Fq3 c_XY;
    edwards_Fq3 c_XZ;

    bool operator==(const edwards_Fq3_conic_coefficients &other) const;
    friend std::ostream& operator<<(std::ostream &out, const edwards_Fq3_conic_coefficients &cc);
    friend std::istream& operator>>(std::istream &in, edwards_Fq3_conic_coefficients &cc);
};
typedef std::vector<edwards_Fq3_conic_coefficients> edwards_ate_G2_precomp;

std::ostream& operator<<(std::ostream& out, const edwards_ate_G2_precomp &prec_Q);
std::istream& operator>>(std::istream& in, edwards_ate_G2_precomp &prec_Q);

struct edwards_ate_G1_precomp {
    edwards_Fq P_XY;
    edwards_Fq P_XZ;
    edwards_Fq P_ZZplusYZ;

    bool operator==(const edwards_ate_G1_precomp &other) const;
    friend std::ostream& operator<<(std::ostream &out, const edwards_ate_G1_precomp &prec_P);
    friend std::istream& operator>>(std::istream &in, edwards_ate_G1_precomp &prec_P);
};

edwards_ate_G1_precomp edwards_ate_precompute_G1(const edwards_G1& P);
edwards_ate_G2_precomp edwards_ate_precompute_G2(const edwards_G2& Q);

edwards_Fq6 edwards_ate_miller_loop(const edwards_ate_G1_precomp &prec_P,
                                    const edwards_ate_G2_precomp &prec_Q);
edwards_Fq6 edwards_ate_double_miller_loop(const edwards_ate_G1_precomp &prec_P1,
                                           const edwards_ate_G2_precomp &prec_Q1,
                                           const edwards_ate_G1_precomp &prec_P2,
                                           const edwards_ate_G2_precomp &prec_Q2);

edwards_Fq6 edwards_ate_pairing(const edwards_G1& P,
                                const edwards_G2 &Q);
edwards_GT edwards_ate_reduced_pairing(const edwards_G1 &P,
                                       const edwards_G2 &Q);

/* choice of pairing */

typedef edwards_ate_G1_precomp edwards_G1_precomp;
typedef edwards_ate_G2_precomp edwards_G2_precomp;

edwards_G1_precomp edwards_precompute_G1(const edwards_G1& P);
edwards_G2_precomp edwards_precompute_G2(const edwards_G2& Q);

edwards_Fq6 edwards_miller_loop(const edwards_G1_precomp &prec_P,
                                const edwards_G2_precomp &prec_Q);

edwards_Fq6 edwards_double_miller_loop(const edwards_G1_precomp &prec_P1,
                                       const edwards_G2_precomp &prec_Q1,
                                       const edwards_G1_precomp &prec_P2,
                                       const edwards_G2_precomp &prec_Q2);

edwards_Fq6 edwards_pairing(const edwards_G1& P,
                            const edwards_G2 &Q);

edwards_GT edwards_reduced_pairing(const edwards_G1 &P,
                                   const edwards_G2 &Q);

} // libsnark
#endif // EDWARDS_PAIRING_HPP_

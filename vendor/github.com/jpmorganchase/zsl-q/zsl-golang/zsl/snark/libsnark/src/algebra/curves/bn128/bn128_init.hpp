/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BN128_INIT_HPP_
#define BN128_INIT_HPP_
#include "algebra/curves/public_params.hpp"
#include "algebra/fields/fp.hpp"
#include "bn.h"   // If you're missing this file, run libsnark's ./prepare-depends.sh

namespace libsnark {

const mp_size_t bn128_r_bitcount = 254;
const mp_size_t bn128_q_bitcount = 254;

const mp_size_t bn128_r_limbs = (bn128_r_bitcount+GMP_NUMB_BITS-1)/GMP_NUMB_BITS;
const mp_size_t bn128_q_limbs = (bn128_q_bitcount+GMP_NUMB_BITS-1)/GMP_NUMB_BITS;

extern bigint<bn128_r_limbs> bn128_modulus_r;
extern bigint<bn128_q_limbs> bn128_modulus_q;

extern bn::Fp bn128_coeff_b;
extern size_t bn128_Fq_s;
extern bn::Fp bn128_Fq_nqr_to_t;
extern mie::Vuint bn128_Fq_t_minus_1_over_2;

extern bn::Fp2 bn128_twist_coeff_b;
extern size_t bn128_Fq2_s;
extern bn::Fp2 bn128_Fq2_nqr_to_t;
extern mie::Vuint bn128_Fq2_t_minus_1_over_2;

typedef Fp_model<bn128_r_limbs, bn128_modulus_r> bn128_Fr;
typedef Fp_model<bn128_q_limbs, bn128_modulus_q> bn128_Fq;

void init_bn128_params();

class bn128_G1;
class bn128_G2;
class bn128_GT;
typedef bn128_GT bn128_Fq12;

} // libsnark
#endif // BN128_INIT_HPP_

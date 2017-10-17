/** @file
 *****************************************************************************

 Declaration of interfaces for initializing MNT6.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT6_INIT_HPP_
#define MNT6_INIT_HPP_

#include "algebra/curves/public_params.hpp"
#include "algebra/curves/mnt/mnt46_common.hpp"
#include "algebra/fields/fp.hpp"
#include "algebra/fields/fp3.hpp"
#include "algebra/fields/fp6_2over3.hpp"

namespace libsnark {

#define mnt6_modulus_r mnt46_modulus_B
#define mnt6_modulus_q mnt46_modulus_A

const mp_size_t mnt6_r_bitcount = mnt46_B_bitcount;
const mp_size_t mnt6_q_bitcount = mnt46_A_bitcount;

const mp_size_t mnt6_r_limbs = mnt46_B_limbs;
const mp_size_t mnt6_q_limbs = mnt46_A_limbs;

extern bigint<mnt6_r_limbs> mnt6_modulus_r;
extern bigint<mnt6_q_limbs> mnt6_modulus_q;

typedef Fp_model<mnt6_r_limbs, mnt6_modulus_r> mnt6_Fr;
typedef Fp_model<mnt6_q_limbs, mnt6_modulus_q> mnt6_Fq;
typedef Fp3_model<mnt6_q_limbs, mnt6_modulus_q> mnt6_Fq3;
typedef Fp6_2over3_model<mnt6_q_limbs, mnt6_modulus_q> mnt6_Fq6;
typedef mnt6_Fq6 mnt6_GT;

// parameters for twisted short Weierstrass curve E'/Fq3 : y^2 = x^3 + (a * twist^2) * x + (b * twist^3)
extern mnt6_Fq3 mnt6_twist;
extern mnt6_Fq3 mnt6_twist_coeff_a;
extern mnt6_Fq3 mnt6_twist_coeff_b;
extern mnt6_Fq mnt6_twist_mul_by_a_c0;
extern mnt6_Fq mnt6_twist_mul_by_a_c1;
extern mnt6_Fq mnt6_twist_mul_by_a_c2;
extern mnt6_Fq mnt6_twist_mul_by_b_c0;
extern mnt6_Fq mnt6_twist_mul_by_b_c1;
extern mnt6_Fq mnt6_twist_mul_by_b_c2;
extern mnt6_Fq mnt6_twist_mul_by_q_X;
extern mnt6_Fq mnt6_twist_mul_by_q_Y;

// parameters for pairing
extern bigint<mnt6_q_limbs> mnt6_ate_loop_count;
extern bool mnt6_ate_is_loop_count_neg;
extern bigint<6*mnt6_q_limbs> mnt6_final_exponent;
extern bigint<mnt6_q_limbs> mnt6_final_exponent_last_chunk_abs_of_w0;
extern bool mnt6_final_exponent_last_chunk_is_w0_neg;
extern bigint<mnt6_q_limbs> mnt6_final_exponent_last_chunk_w1;

void init_mnt6_params();

class mnt6_G1;
class mnt6_G2;

} // libsnark

#endif // MNT6_INIT_HPP_

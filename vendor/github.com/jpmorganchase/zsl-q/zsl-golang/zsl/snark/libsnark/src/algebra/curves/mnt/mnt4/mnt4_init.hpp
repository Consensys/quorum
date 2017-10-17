/** @file
 *****************************************************************************

 Declaration of interfaces for initializing MNT4.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT4_INIT_HPP_
#define MNT4_INIT_HPP_

#include "algebra/curves/public_params.hpp"
#include "algebra/curves/mnt/mnt46_common.hpp"
#include "algebra/fields/fp.hpp"
#include "algebra/fields/fp2.hpp"
#include "algebra/fields/fp4.hpp"

namespace libsnark {

#define mnt4_modulus_r mnt46_modulus_A
#define mnt4_modulus_q mnt46_modulus_B

const mp_size_t mnt4_r_bitcount = mnt46_A_bitcount;
const mp_size_t mnt4_q_bitcount = mnt46_B_bitcount;

const mp_size_t mnt4_r_limbs = mnt46_A_limbs;
const mp_size_t mnt4_q_limbs = mnt46_B_limbs;

extern bigint<mnt4_r_limbs> mnt4_modulus_r;
extern bigint<mnt4_q_limbs> mnt4_modulus_q;

typedef Fp_model<mnt4_r_limbs, mnt4_modulus_r> mnt4_Fr;
typedef Fp_model<mnt4_q_limbs, mnt4_modulus_q> mnt4_Fq;
typedef Fp2_model<mnt4_q_limbs, mnt4_modulus_q> mnt4_Fq2;
typedef Fp4_model<mnt4_q_limbs, mnt4_modulus_q> mnt4_Fq4;
typedef mnt4_Fq4 mnt4_GT;

// parameters for twisted short Weierstrass curve E'/Fq2 : y^2 = x^3 + (a * twist^2) * x + (b * twist^3)
extern mnt4_Fq2 mnt4_twist;
extern mnt4_Fq2 mnt4_twist_coeff_a;
extern mnt4_Fq2 mnt4_twist_coeff_b;
extern mnt4_Fq mnt4_twist_mul_by_a_c0;
extern mnt4_Fq mnt4_twist_mul_by_a_c1;
extern mnt4_Fq mnt4_twist_mul_by_b_c0;
extern mnt4_Fq mnt4_twist_mul_by_b_c1;
extern mnt4_Fq mnt4_twist_mul_by_q_X;
extern mnt4_Fq mnt4_twist_mul_by_q_Y;

// parameters for pairing
extern bigint<mnt4_q_limbs> mnt4_ate_loop_count;
extern bool mnt4_ate_is_loop_count_neg;
extern bigint<4*mnt4_q_limbs> mnt4_final_exponent;
extern bigint<mnt4_q_limbs> mnt4_final_exponent_last_chunk_abs_of_w0;
extern bool mnt4_final_exponent_last_chunk_is_w0_neg;
extern bigint<mnt4_q_limbs> mnt4_final_exponent_last_chunk_w1;

void init_mnt4_params();

class mnt4_G1;
class mnt4_G2;

} // libsnark

#endif // MNT4_INIT_HPP_

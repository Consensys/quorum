/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BN128_G2_HPP_
#define BN128_G2_HPP_
#include <iostream>
#include <vector>
#include "algebra/curves/bn128/bn128_init.hpp"
#include "algebra/curves/curve_utils.hpp"
#include "bn.h"

namespace libsnark {

class bn128_G2;
std::ostream& operator<<(std::ostream &, const bn128_G2&);
std::istream& operator>>(std::istream &, bn128_G2&);

class bn128_G2 {
private:
    static bn::Fp2 sqrt(const bn::Fp2 &el);
public:
#ifdef PROFILE_OP_COUNTS
    static long long add_cnt;
    static long long dbl_cnt;
#endif
    static std::vector<size_t> wnaf_window_table;
    static std::vector<size_t> fixed_base_exp_window_table;
    static bn128_G2 G2_zero;
    static bn128_G2 G2_one;

    bn::Fp2 coord[3];
    bn128_G2();
    typedef bn128_Fq base_field;
    typedef bn128_Fr scalar_field;

    void print() const;
    void print_coordinates() const;

    void to_affine_coordinates();
    void to_special();
    bool is_special() const;

    bool is_zero() const;

    bool operator==(const bn128_G2 &other) const;
    bool operator!=(const bn128_G2 &other) const;

    bn128_G2 operator+(const bn128_G2 &other) const;
    bn128_G2 operator-() const;
    bn128_G2 operator-(const bn128_G2 &other) const;

    bn128_G2 add(const bn128_G2 &other) const;
    bn128_G2 mixed_add(const bn128_G2 &other) const;
    bn128_G2 dbl() const;

    bool is_well_formed() const;

    static bn128_G2 zero();
    static bn128_G2 one();
    static bn128_G2 random_element();

    static size_t size_in_bits() { return 2*base_field::size_in_bits() + 1; }
    static bigint<base_field::num_limbs> base_field_char() { return base_field::field_char(); }
    static bigint<scalar_field::num_limbs> order() { return scalar_field::field_char(); }

    friend std::ostream& operator<<(std::ostream &out, const bn128_G2 &g);
    friend std::istream& operator>>(std::istream &in, bn128_G2 &g);
};

template<mp_size_t m>
bn128_G2 operator*(const bigint<m> &lhs, const bn128_G2 &rhs)
{
    return scalar_mul<bn128_G2, m>(rhs, lhs);
}

template<mp_size_t m, const bigint<m>& modulus_p>
bn128_G2 operator*(const Fp_model<m, modulus_p> &lhs, const bn128_G2 &rhs)
{
    return scalar_mul<bn128_G2, m>(rhs, lhs.as_bigint());
}

template<typename T>
void batch_to_special_all_non_zeros(std::vector<T> &vec);
template<>
void batch_to_special_all_non_zeros<bn128_G2>(std::vector<bn128_G2> &vec);

} // libsnark
#endif // BN128_G2_HPP_

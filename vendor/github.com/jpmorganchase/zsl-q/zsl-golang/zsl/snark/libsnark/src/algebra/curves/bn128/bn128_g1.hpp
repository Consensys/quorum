/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BN128_G1_HPP_
#define BN128_G1_HPP_
#include <vector>
#include "algebra/curves/bn128/bn128_init.hpp"
#include "algebra/curves/curve_utils.hpp"
#include "bn.h"

namespace libsnark {

class bn128_G1;
std::ostream& operator<<(std::ostream &, const bn128_G1&);
std::istream& operator>>(std::istream &, bn128_G1&);

class bn128_G1 {
private:
    static bn::Fp sqrt(const bn::Fp &el);
public:
#ifdef PROFILE_OP_COUNTS
    static long long add_cnt;
    static long long dbl_cnt;
#endif
    static std::vector<size_t> wnaf_window_table;
    static std::vector<size_t> fixed_base_exp_window_table;
    static bn128_G1 G1_zero;
    static bn128_G1 G1_one;

    bn::Fp coord[3];
    bn128_G1();
    typedef bn128_Fq base_field;
    typedef bn128_Fr scalar_field;

    void print() const;
    void print_coordinates() const;

    void to_affine_coordinates();
    void to_special();
    bool is_special() const;

    bool is_zero() const;

    bool operator==(const bn128_G1 &other) const;
    bool operator!=(const bn128_G1 &other) const;

    bn128_G1 operator+(const bn128_G1 &other) const;
    bn128_G1 operator-() const;
    bn128_G1 operator-(const bn128_G1 &other) const;

    bn128_G1 add(const bn128_G1 &other) const;
    bn128_G1 mixed_add(const bn128_G1 &other) const;
    bn128_G1 dbl() const;

    bool is_well_formed() const;

    static bn128_G1 zero();
    static bn128_G1 one();
    static bn128_G1 random_element();

    static size_t size_in_bits() { return bn128_Fq::size_in_bits() + 1; }
    static bigint<base_field::num_limbs> base_field_char() { return base_field::field_char(); }
    static bigint<scalar_field::num_limbs> order() { return scalar_field::field_char(); }

    friend std::ostream& operator<<(std::ostream &out, const bn128_G1 &g);
    friend std::istream& operator>>(std::istream &in, bn128_G1 &g);
};

template<mp_size_t m>
bn128_G1 operator*(const bigint<m> &lhs, const bn128_G1 &rhs)
{
    return scalar_mul<bn128_G1, m>(rhs, lhs);
}

template<mp_size_t m, const bigint<m>& modulus_p>
bn128_G1 operator*(const Fp_model<m,modulus_p> &lhs, const bn128_G1 &rhs)
{
    return scalar_mul<bn128_G1, m>(rhs, lhs.as_bigint());
}

std::ostream& operator<<(std::ostream& out, const std::vector<bn128_G1> &v);
std::istream& operator>>(std::istream& in, std::vector<bn128_G1> &v);

template<typename T>
void batch_to_special_all_non_zeros(std::vector<T> &vec);
template<>
void batch_to_special_all_non_zeros<bn128_G1>(std::vector<bn128_G1> &vec);

} // libsnark
#endif // BN128_G1_HPP_

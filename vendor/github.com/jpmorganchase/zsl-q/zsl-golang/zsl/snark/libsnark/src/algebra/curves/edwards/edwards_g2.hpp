/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef EDWARDS_G2_HPP_
#define EDWARDS_G2_HPP_
#include <iostream>
#include <vector>
#include "algebra/curves/edwards/edwards_init.hpp"
#include "algebra/curves/curve_utils.hpp"

namespace libsnark {

class edwards_G2;
std::ostream& operator<<(std::ostream &, const edwards_G2&);
std::istream& operator>>(std::istream &, edwards_G2&);

class edwards_G2 {
public:
#ifdef PROFILE_OP_COUNTS
    static long long add_cnt;
    static long long dbl_cnt;
#endif
    static std::vector<size_t> wnaf_window_table;
    static std::vector<size_t> fixed_base_exp_window_table;

    static edwards_G2 G2_zero;
    static edwards_G2 G2_one;

    edwards_Fq3 X, Y, Z;
    edwards_G2();
private:
    edwards_G2(const edwards_Fq3& X, const edwards_Fq3& Y, const edwards_Fq3& Z) : X(X), Y(Y), Z(Z) {};
public:
    static edwards_Fq3 mul_by_a(const edwards_Fq3 &elt);
    static edwards_Fq3 mul_by_d(const edwards_Fq3 &elt);
    typedef edwards_Fq base_field;
    typedef edwards_Fq3 twist_field;
    typedef edwards_Fr scalar_field;

    // using inverted coordinates
    edwards_G2(const edwards_Fq3& X, const edwards_Fq3& Y) : X(Y), Y(X), Z(X*Y) {};

    void print() const;
    void print_coordinates() const;

    void to_affine_coordinates();
    void to_special();
    bool is_special() const;

    bool is_zero() const;

    bool operator==(const edwards_G2 &other) const;
    bool operator!=(const edwards_G2 &other) const;

    edwards_G2 operator+(const edwards_G2 &other) const;
    edwards_G2 operator-() const;
    edwards_G2 operator-(const edwards_G2 &other) const;

    edwards_G2 add(const edwards_G2 &other) const;
    edwards_G2 mixed_add(const edwards_G2 &other) const;
    edwards_G2 dbl() const;
    edwards_G2 mul_by_q() const;

    bool is_well_formed() const;

    static edwards_G2 zero();
    static edwards_G2 one();
    static edwards_G2 random_element();

    static size_t size_in_bits() { return twist_field::size_in_bits() + 1; }
    static bigint<base_field::num_limbs> base_field_char() { return base_field::field_char(); }
    static bigint<scalar_field::num_limbs> order() { return scalar_field::field_char(); }

    friend std::ostream& operator<<(std::ostream &out, const edwards_G2 &g);
    friend std::istream& operator>>(std::istream &in, edwards_G2 &g);
};

template<mp_size_t m>
edwards_G2 operator*(const bigint<m> &lhs, const edwards_G2 &rhs)
{
    return scalar_mul<edwards_G2, m>(rhs, lhs);
}

template<mp_size_t m, const bigint<m>& modulus_p>
edwards_G2 operator*(const Fp_model<m, modulus_p> &lhs, const edwards_G2 &rhs)
{
   return scalar_mul<edwards_G2, m>(rhs, lhs.as_bigint());
}

template<typename T>
void batch_to_special_all_non_zeros(std::vector<T> &vec);
template<>
void batch_to_special_all_non_zeros<edwards_G2>(std::vector<edwards_G2> &vec);

} // libsnark
#endif // EDWARDS_G2_HPP_

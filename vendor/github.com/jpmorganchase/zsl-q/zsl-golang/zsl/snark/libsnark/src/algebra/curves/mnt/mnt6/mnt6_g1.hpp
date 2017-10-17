/** @file
 *****************************************************************************

 Declaration of interfaces for the MNT6 G1 group.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT6_G1_HPP_
#define MNT6_G1_HPP_

#include <vector>

#include "algebra/curves/curve_utils.hpp"
#include "algebra/curves/mnt/mnt6/mnt6_init.hpp"

namespace libsnark {

class mnt6_G1;
std::ostream& operator<<(std::ostream &, const mnt6_G1&);
std::istream& operator>>(std::istream &, mnt6_G1&);

class mnt6_G1 {
private:
    mnt6_Fq X_, Y_, Z_;
public:
#ifdef PROFILE_OP_COUNTS
    static long long add_cnt;
    static long long dbl_cnt;
#endif
    static std::vector<size_t> wnaf_window_table;
    static std::vector<size_t> fixed_base_exp_window_table;
    static mnt6_G1 G1_zero;
    static mnt6_G1 G1_one;
    static mnt6_Fq coeff_a;
    static mnt6_Fq coeff_b;

    typedef mnt6_Fq base_field;
    typedef mnt6_Fr scalar_field;

    // using projective coordinates
    mnt6_G1();
    mnt6_G1(const mnt6_Fq& X, const mnt6_Fq& Y) : X_(X), Y_(Y), Z_(base_field::one()) {}
    mnt6_G1(const mnt6_Fq& X, const mnt6_Fq& Y, const mnt6_Fq& Z) : X_(X), Y_(Y), Z_(Z) {}

    mnt6_Fq X() const { return X_; }
    mnt6_Fq Y() const { return Y_; }
    mnt6_Fq Z() const { return Z_; }

    void print() const;
    void print_coordinates() const;

    void to_affine_coordinates();
    void to_special();
    bool is_special() const;

    bool is_zero() const;

    bool operator==(const mnt6_G1 &other) const;
    bool operator!=(const mnt6_G1 &other) const;

    mnt6_G1 operator+(const mnt6_G1 &other) const;
    mnt6_G1 operator-() const;
    mnt6_G1 operator-(const mnt6_G1 &other) const;

    mnt6_G1 add(const mnt6_G1 &other) const;
    mnt6_G1 mixed_add(const mnt6_G1 &other) const;
    mnt6_G1 dbl() const;

    bool is_well_formed() const;

    static mnt6_G1 zero();
    static mnt6_G1 one();
    static mnt6_G1 random_element();

    static size_t size_in_bits() { return base_field::size_in_bits() + 1; }
    static bigint<base_field::num_limbs> base_field_char() { return base_field::field_char(); }
    static bigint<scalar_field::num_limbs> order() { return scalar_field::field_char(); }

    friend std::ostream& operator<<(std::ostream &out, const mnt6_G1 &g);
    friend std::istream& operator>>(std::istream &in, mnt6_G1 &g);
};

template<mp_size_t m>
mnt6_G1 operator*(const bigint<m> &lhs, const mnt6_G1 &rhs)
{
    return scalar_mul<mnt6_G1, m>(rhs, lhs);
}

template<mp_size_t m, const bigint<m>& modulus_p>
mnt6_G1 operator*(const Fp_model<m,modulus_p> &lhs, const mnt6_G1 &rhs)
{
    return scalar_mul<mnt6_G1, m>(rhs, lhs.as_bigint());
}

std::ostream& operator<<(std::ostream& out, const std::vector<mnt6_G1> &v);
std::istream& operator>>(std::istream& in, std::vector<mnt6_G1> &v);

template<typename T>
void batch_to_special_all_non_zeros(std::vector<T> &vec);
template<>
void batch_to_special_all_non_zeros<mnt6_G1>(std::vector<mnt6_G1> &vec);

} // libsnark

#endif // MNT6_G1_HPP_

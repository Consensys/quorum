/** @file
 *****************************************************************************

 Declaration of interfaces for the MNT6 G2 group.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MNT6_G2_HPP_
#define MNT6_G2_HPP_

#include <vector>

#include "algebra/curves/curve_utils.hpp"
#include "algebra/curves/mnt/mnt6/mnt6_init.hpp"

namespace libsnark {

class mnt6_G2;
std::ostream& operator<<(std::ostream &, const mnt6_G2&);
std::istream& operator>>(std::istream &, mnt6_G2&);

class mnt6_G2 {
private:
    mnt6_Fq3 X_, Y_, Z_;
public:
#ifdef PROFILE_OP_COUNTS
    static long long add_cnt;
    static long long dbl_cnt;
#endif
    static std::vector<size_t> wnaf_window_table;
    static std::vector<size_t> fixed_base_exp_window_table;
    static mnt6_G2 G2_zero;
    static mnt6_G2 G2_one;
    static mnt6_Fq3 twist;
    static mnt6_Fq3 coeff_a;
    static mnt6_Fq3 coeff_b;

    typedef mnt6_Fq base_field;
    typedef mnt6_Fq3 twist_field;
    typedef mnt6_Fr scalar_field;

    // using projective coordinates
    mnt6_G2();
    mnt6_G2(const mnt6_Fq3& X, const mnt6_Fq3& Y, const mnt6_Fq3& Z) : X_(X), Y_(Y), Z_(Z) {}

    mnt6_Fq3 X() const { return X_; }
    mnt6_Fq3 Y() const { return Y_; }
    mnt6_Fq3 Z() const { return Z_; }

    static mnt6_Fq3 mul_by_a(const mnt6_Fq3 &elt);
    static mnt6_Fq3 mul_by_b(const mnt6_Fq3 &elt);

    void print() const;
    void print_coordinates() const;

    void to_affine_coordinates();
    void to_special();
    bool is_special() const;

    bool is_zero() const;

    bool operator==(const mnt6_G2 &other) const;
    bool operator!=(const mnt6_G2 &other) const;

    mnt6_G2 operator+(const mnt6_G2 &other) const;
    mnt6_G2 operator-() const;
    mnt6_G2 operator-(const mnt6_G2 &other) const;

    mnt6_G2 add(const mnt6_G2 &other) const;
    mnt6_G2 mixed_add(const mnt6_G2 &other) const;
    mnt6_G2 dbl() const;
    mnt6_G2 mul_by_q() const;

    bool is_well_formed() const;

    static mnt6_G2 zero();
    static mnt6_G2 one();
    static mnt6_G2 random_element();

    static size_t size_in_bits() { return twist_field::size_in_bits() + 1; }
    static bigint<base_field::num_limbs> base_field_char() { return base_field::field_char(); }
    static bigint<scalar_field::num_limbs> order() { return scalar_field::field_char(); }

    friend std::ostream& operator<<(std::ostream &out, const mnt6_G2 &g);
    friend std::istream& operator>>(std::istream &in, mnt6_G2 &g);
};

template<mp_size_t m>
mnt6_G2 operator*(const bigint<m> &lhs, const mnt6_G2 &rhs)
{
    return scalar_mul<mnt6_G2, m>(rhs, lhs);
}

template<mp_size_t m, const bigint<m>& modulus_p>
mnt6_G2 operator*(const Fp_model<m,modulus_p> &lhs, const mnt6_G2 &rhs)
{
    return scalar_mul<mnt6_G2, m>(rhs, lhs.as_bigint());
}

template<typename T>
void batch_to_special_all_non_zeros(std::vector<T> &vec);
template<>
void batch_to_special_all_non_zeros<mnt6_G2>(std::vector<mnt6_G2> &vec);

} // libsnark

#endif // MNT6_G2_HPP_

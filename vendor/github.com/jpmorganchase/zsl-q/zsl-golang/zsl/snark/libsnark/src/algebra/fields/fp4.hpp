/** @file
 *****************************************************************************

 Declaration of interfaces for the (extension) field Fp4.

 The field Fp4 equals Fp2[V]/(V^2-U) where Fp2 = Fp[U]/(U^2-non_residue) and non_residue is in Fp.

 ASSUMPTION: the modulus p is 1 mod 6.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FP4_HPP_
#define FP4_HPP_

#include "algebra/fields/fp.hpp"
#include "algebra/fields/fp2.hpp"

namespace libsnark {

template<mp_size_t n, const bigint<n>& modulus>
class Fp4_model;

template<mp_size_t n, const bigint<n>& modulus>
std::ostream& operator<<(std::ostream &, const Fp4_model<n, modulus> &);

template<mp_size_t n, const bigint<n>& modulus>
std::istream& operator>>(std::istream &, Fp4_model<n, modulus> &);

template<mp_size_t n, const bigint<n>& modulus>
class Fp4_model {
public:
    typedef Fp_model<n, modulus> my_Fp;
    typedef Fp2_model<n, modulus> my_Fp2;
    typedef my_Fp2 my_Fpe;

    static my_Fp non_residue;
    static my_Fp Frobenius_coeffs_c1[4]; // non_residue^((modulus^i-1)/4) for i=0,1,2,3

    my_Fp2 c0, c1;
    Fp4_model() {};
    Fp4_model(const my_Fp2& c0, const my_Fp2& c1) : c0(c0), c1(c1) {};

    void print() const { printf("c0/c1:\n"); c0.print(); c1.print(); }
    void clear() { c0.clear(); c1.clear(); }

    static Fp4_model<n, modulus> zero();
    static Fp4_model<n, modulus> one();
    static Fp4_model<n, modulus> random_element();

    bool is_zero() const { return c0.is_zero() && c1.is_zero(); }
    bool operator==(const Fp4_model &other) const;
    bool operator!=(const Fp4_model &other) const;

    Fp4_model operator+(const Fp4_model &other) const;
    Fp4_model operator-(const Fp4_model &other) const;
    Fp4_model operator*(const Fp4_model &other) const;
    Fp4_model mul_by_023(const Fp4_model &other) const;
    Fp4_model operator-() const;
    Fp4_model squared() const;
    Fp4_model inverse() const;
    Fp4_model Frobenius_map(unsigned long power) const;
    Fp4_model unitary_inverse() const;
    Fp4_model cyclotomic_squared() const;

    static my_Fp2 mul_by_non_residue(const my_Fp2 &elt);

    template<mp_size_t m>
    Fp4_model cyclotomic_exp(const bigint<m> &exponent) const;

    static bigint<n> base_field_char() { return modulus; }
    static constexpr size_t extension_degree() { return 4; }

    friend std::ostream& operator<< <n, modulus>(std::ostream &out, const Fp4_model<n, modulus> &el);
    friend std::istream& operator>> <n, modulus>(std::istream &in, Fp4_model<n, modulus> &el);
};

template<mp_size_t n, const bigint<n>& modulus>
Fp4_model<n, modulus> operator*(const Fp_model<n, modulus> &lhs, const Fp4_model<n, modulus> &rhs);

template<mp_size_t n, const bigint<n>& modulus>
Fp4_model<n, modulus> operator*(const Fp2_model<n, modulus> &lhs, const Fp4_model<n, modulus> &rhs);

template<mp_size_t n, const bigint<n>& modulus, mp_size_t m>
Fp4_model<n, modulus> operator^(const Fp4_model<n, modulus> &self, const bigint<m> &exponent);

template<mp_size_t n, const bigint<n>& modulus, mp_size_t m, const bigint<m>& modulus_p>
Fp4_model<n, modulus> operator^(const Fp4_model<n, modulus> &self, const Fp_model<m, modulus_p> &exponent);

template<mp_size_t n, const bigint<n>& modulus>
Fp_model<n, modulus> Fp4_model<n, modulus>::non_residue;

template<mp_size_t n, const bigint<n>& modulus>
Fp_model<n, modulus> Fp4_model<n, modulus>::Frobenius_coeffs_c1[4];


} // libsnark

#include "algebra/fields/fp4.tcc"

#endif // FP4_HPP_

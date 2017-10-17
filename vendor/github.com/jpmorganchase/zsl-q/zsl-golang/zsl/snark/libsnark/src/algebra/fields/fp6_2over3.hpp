/** @file
 *****************************************************************************
 Declaration of arithmetic in the finite field F[(p^3)^2]
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FP6_2OVER3_HPP_
#define FP6_2OVER3_HPP_
#include "algebra/fields/fp.hpp"
#include "algebra/fields/fp2.hpp"
#include "algebra/fields/fp3.hpp"

namespace libsnark {

/**
 * Arithmetic in the finite field F[(p^3)^2].
 *
 * Let p := modulus. This interface provides arithmetic for the extension field
 * Fp6 = Fp3[Y]/(Y^2-X) where Fp3 = Fp[X]/(X^3-non_residue) and non_residue is in Fp.
 *
 * ASSUMPTION: p = 1 (mod 6)
 */
template<mp_size_t n, const bigint<n>& modulus>
class Fp6_2over3_model;

template<mp_size_t n, const bigint<n>& modulus>
std::ostream& operator<<(std::ostream &, const Fp6_2over3_model<n, modulus> &);

template<mp_size_t n, const bigint<n>& modulus>
std::istream& operator>>(std::istream &, Fp6_2over3_model<n, modulus> &);

template<mp_size_t n, const bigint<n>& modulus>
class Fp6_2over3_model {
public:
    typedef Fp_model<n, modulus> my_Fp;
    typedef Fp2_model<n, modulus> my_Fp2;
    typedef Fp3_model<n, modulus> my_Fp3;
    typedef my_Fp3 my_Fpe;

    static my_Fp non_residue;
    static my_Fp Frobenius_coeffs_c1[6]; // non_residue^((modulus^i-1)/6)   for i=0,1,2,3,4,5

    my_Fp3 c0, c1;
    Fp6_2over3_model() {};
    Fp6_2over3_model(const my_Fp3& c0, const my_Fp3& c1) : c0(c0), c1(c1) {};

    void print() const { printf("c0/c1:\n"); c0.print(); c1.print(); }
    void clear() { c0.clear(); c1.clear(); }

    static Fp6_2over3_model<n, modulus> zero();
    static Fp6_2over3_model<n, modulus> one();
    static Fp6_2over3_model<n, modulus> random_element();

    bool is_zero() const { return c0.is_zero() && c1.is_zero(); }
    bool operator==(const Fp6_2over3_model &other) const;
    bool operator!=(const Fp6_2over3_model &other) const;

    Fp6_2over3_model operator+(const Fp6_2over3_model &other) const;
    Fp6_2over3_model operator-(const Fp6_2over3_model &other) const;
    Fp6_2over3_model operator*(const Fp6_2over3_model &other) const;
    Fp6_2over3_model mul_by_2345(const Fp6_2over3_model &other) const;
    Fp6_2over3_model operator-() const;
    Fp6_2over3_model squared() const;
    Fp6_2over3_model inverse() const;
    Fp6_2over3_model Frobenius_map(unsigned long power) const;
    Fp6_2over3_model unitary_inverse() const;
    Fp6_2over3_model cyclotomic_squared() const;

    static my_Fp3 mul_by_non_residue(const my_Fp3 &elem);

    template<mp_size_t m>
    Fp6_2over3_model cyclotomic_exp(const bigint<m> &exponent) const;

    static bigint<n> base_field_char() { return modulus; }
    static constexpr size_t extension_degree() { return 6; }

    friend std::ostream& operator<< <n, modulus>(std::ostream &out, const Fp6_2over3_model<n, modulus> &el);
    friend std::istream& operator>> <n, modulus>(std::istream &in, Fp6_2over3_model<n, modulus> &el);
};

template<mp_size_t n, const bigint<n>& modulus>
std::ostream& operator<<(std::ostream& out, const std::vector<Fp6_2over3_model<n, modulus> > &v);

template<mp_size_t n, const bigint<n>& modulus>
std::istream& operator>>(std::istream& in, std::vector<Fp6_2over3_model<n, modulus> > &v);

template<mp_size_t n, const bigint<n>& modulus>
Fp6_2over3_model<n, modulus> operator*(const Fp_model<n, modulus> &lhs, const Fp6_2over3_model<n, modulus> &rhs);

template<mp_size_t n, const bigint<n>& modulus, mp_size_t m>
Fp6_2over3_model<n, modulus> operator^(const Fp6_2over3_model<n, modulus> &self, const bigint<m> &exponent);

template<mp_size_t n, const bigint<n>& modulus, mp_size_t m, const bigint<m>& exp_modulus>
Fp6_2over3_model<n, modulus> operator^(const Fp6_2over3_model<n, modulus> &self, const Fp_model<m, exp_modulus> &exponent);

template<mp_size_t n, const bigint<n>& modulus>
Fp_model<n, modulus> Fp6_2over3_model<n, modulus>::non_residue;

template<mp_size_t n, const bigint<n>& modulus>
Fp_model<n, modulus> Fp6_2over3_model<n, modulus>::Frobenius_coeffs_c1[6];

} // libsnark
#include "algebra/fields/fp6_2over3.tcc"

#endif // FP6_2OVER3_HPP_

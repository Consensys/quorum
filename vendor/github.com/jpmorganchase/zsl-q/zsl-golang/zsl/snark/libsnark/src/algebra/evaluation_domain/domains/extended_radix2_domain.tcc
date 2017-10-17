/** @file
 *****************************************************************************

 Implementation of interfaces for the "extended radix-2" evaluation domain.

 See extended_radix2_domain.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef EXTENDED_RADIX2_DOMAIN_TCC_

#include "algebra/evaluation_domain/domains/basic_radix2_domain_aux.hpp"

namespace libsnark {

template<typename FieldT>
extended_radix2_domain<FieldT>::extended_radix2_domain(const size_t m) : evaluation_domain<FieldT>(m)
{
    assert(m > 1);

    const size_t logm = log2(m);

    assert(logm == FieldT::s + 1);

    small_m = m/2;
    omega = get_root_of_unity<FieldT>(small_m);
    shift = coset_shift<FieldT>();
}

template<typename FieldT>
void extended_radix2_domain<FieldT>::FFT(std::vector<FieldT> &a)
{
    assert(a.size() == this->m);

    std::vector<FieldT> a0(small_m, FieldT::zero());
    std::vector<FieldT> a1(small_m, FieldT::zero());

    const FieldT shift_to_small_m = shift^bigint<1>(small_m);

    FieldT shift_i = FieldT::one();
    for (size_t i = 0; i < small_m; ++i)
    {
        a0[i] = a[i] + a[small_m + i];
        a1[i] = shift_i * (a[i] + shift_to_small_m * a[small_m + i]);

        shift_i *= shift;
    }

    _basic_radix2_FFT(a0, omega);
    _basic_radix2_FFT(a1, omega);

    for (size_t i = 0; i < small_m; ++i)
    {
        a[i] = a0[i];
        a[i+small_m] = a1[i];
    }
}

template<typename FieldT>
void extended_radix2_domain<FieldT>::iFFT(std::vector<FieldT> &a)
{
    assert(a.size() == this->m);

    // note: this is not in-place
    std::vector<FieldT> a0(a.begin(), a.begin() + small_m);
    std::vector<FieldT> a1(a.begin() + small_m, a.end());

    const FieldT omega_inverse = omega.inverse();
    _basic_radix2_FFT(a0, omega_inverse);
    _basic_radix2_FFT(a1, omega_inverse);

    const FieldT shift_to_small_m = shift^bigint<1>(small_m);
    const FieldT sconst = (FieldT(small_m) * (FieldT::one()-shift_to_small_m)).inverse();

    const FieldT shift_inverse = shift.inverse();
    FieldT shift_inverse_i = FieldT::one();

    for (size_t i = 0; i < small_m; ++i)
    {
        a[i] = sconst * (-shift_to_small_m * a0[i] + shift_inverse_i * a1[i]);
        a[i+small_m] = sconst * (a0[i] - shift_inverse_i * a1[i]);

        shift_inverse_i *= shift_inverse;
    }
}

template<typename FieldT>
void extended_radix2_domain<FieldT>::cosetFFT(std::vector<FieldT> &a, const FieldT &g)
{
    _multiply_by_coset(a, g);
    FFT(a);
}

template<typename FieldT>
void extended_radix2_domain<FieldT>::icosetFFT(std::vector<FieldT> &a, const FieldT &g)
{
    iFFT(a);
    _multiply_by_coset(a, g.inverse());
}

template<typename FieldT>
std::vector<FieldT> extended_radix2_domain<FieldT>::lagrange_coeffs(const FieldT &t)
{
    const std::vector<FieldT> T0 = _basic_radix2_lagrange_coeffs(small_m, t);
    const std::vector<FieldT> T1 = _basic_radix2_lagrange_coeffs(small_m, t * shift.inverse());

    std::vector<FieldT> result(this->m, FieldT::zero());

    const FieldT t_to_small_m = t ^ bigint<1>(small_m);
    const FieldT shift_to_small_m = shift ^ bigint<1>(small_m);
    const FieldT one_over_denom = (shift_to_small_m - FieldT::one()).inverse();
    const FieldT T0_coeff = (t_to_small_m - shift_to_small_m) * (-one_over_denom);
    const FieldT T1_coeff = (t_to_small_m - FieldT::one()) * one_over_denom;
    for (size_t i = 0; i < small_m; ++i)
    {
        result[i] = T0[i] * T0_coeff;
        result[i+small_m] = T1[i] * T1_coeff;
    }

    return result;
}

template<typename FieldT>
FieldT extended_radix2_domain<FieldT>::get_element(const size_t idx)
{
    if (idx < small_m)
    {
        return omega^idx;
    }
    else
    {
        return shift*(omega^(idx-small_m));
    }
}

template<typename FieldT>
FieldT extended_radix2_domain<FieldT>::compute_Z(const FieldT &t)
{
    return ((t^small_m) - FieldT::one()) * ((t^small_m) - (shift^small_m));
}

template<typename FieldT>
void extended_radix2_domain<FieldT>::add_poly_Z(const FieldT &coeff, std::vector<FieldT> &H)
{
    assert(H.size() == this->m+1);
    const FieldT shift_to_small_m = shift^small_m;

    H[this->m] += coeff;
    H[small_m] -= coeff * (shift_to_small_m + FieldT::one());
    H[0] += coeff * shift_to_small_m;
}

template<typename FieldT>
void extended_radix2_domain<FieldT>::divide_by_Z_on_coset(std::vector<FieldT> &P)
{
    const FieldT coset = FieldT::multiplicative_generator;

    const FieldT coset_to_small_m = coset^small_m;
    const FieldT shift_to_small_m = shift^small_m;

    const FieldT Z0 = (coset_to_small_m - FieldT::one()) * (coset_to_small_m - shift_to_small_m);
    const FieldT Z1 = (coset_to_small_m*shift_to_small_m - FieldT::one()) * (coset_to_small_m * shift_to_small_m - shift_to_small_m);

    const FieldT Z0_inverse = Z0.inverse();
    const FieldT Z1_inverse = Z1.inverse();

    for (size_t i = 0; i < small_m; ++i)
    {
        P[i] *= Z0_inverse;
        P[i+small_m] *= Z1_inverse;
    }
}

} // libsnark

#endif // EXTENDED_RADIX2_DOMAIN_TCC_

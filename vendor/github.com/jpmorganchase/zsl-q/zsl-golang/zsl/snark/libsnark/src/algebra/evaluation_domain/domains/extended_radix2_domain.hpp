/** @file
 *****************************************************************************

 Declaration of interfaces for the "extended radix-2" evaluation domain.

 Roughly, the domain has size m = 2^{k+1} and consists of
 "the m-th roots of unity" union "a coset of these roots".

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef EXTENDED_RADIX2_DOMAIN_HPP_
#define EXTENDED_RADIX2_DOMAIN_HPP_

#include "algebra/evaluation_domain/evaluation_domain.hpp"

namespace libsnark {

template<typename FieldT>
class extended_radix2_domain : public evaluation_domain<FieldT> {
public:

    size_t small_m;
    FieldT omega;
    FieldT shift;

    extended_radix2_domain(const size_t m);

    void FFT(std::vector<FieldT> &a);
    void iFFT(std::vector<FieldT> &a);
    void cosetFFT(std::vector<FieldT> &a, const FieldT &g);
    void icosetFFT(std::vector<FieldT> &a, const FieldT &g);
    std::vector<FieldT> lagrange_coeffs(const FieldT &t);
    FieldT get_element(const size_t idx);
    FieldT compute_Z(const FieldT &t);
    void add_poly_Z(const FieldT &coeff, std::vector<FieldT> &H);
    void divide_by_Z_on_coset(std::vector<FieldT> &P);

};

} // libsnark

#include "algebra/evaluation_domain/domains/extended_radix2_domain.tcc"

#endif // EXTENDED_RADIX2_DOMAIN_HPP_

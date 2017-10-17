/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BN128_GT_HPP_
#define BN128_GT_HPP_
#include "algebra/fields/fp.hpp"
#include "algebra/fields/field_utils.hpp"
#include <iostream>
#include "bn.h"

namespace libsnark {

class bn128_GT;
std::ostream& operator<<(std::ostream &, const bn128_GT&);
std::istream& operator>>(std::istream &, bn128_GT&);

class bn128_GT {
public:
    static bn128_GT GT_one;
    bn::Fp12 elem;

    bn128_GT();
    bool operator==(const bn128_GT &other) const;
    bool operator!=(const bn128_GT &other) const;

    bn128_GT operator*(const bn128_GT &other) const;
    bn128_GT unitary_inverse() const;

    static bn128_GT one();

    void print() { std::cout << this->elem << "\n"; };

    friend std::ostream& operator<<(std::ostream &out, const bn128_GT &g);
    friend std::istream& operator>>(std::istream &in, bn128_GT &g);
};

template<mp_size_t m>
bn128_GT operator^(const bn128_GT &rhs, const bigint<m> &lhs)
{
    return power<bn128_GT, m>(rhs, lhs);
}


template<mp_size_t m, const bigint<m>& modulus_p>
bn128_GT operator^(const bn128_GT &rhs, const Fp_model<m,modulus_p> &lhs)
{
    return power<bn128_GT, m>(rhs, lhs.as_bigint());
}

} // libsnark
#endif // BN128_GT_HPP_

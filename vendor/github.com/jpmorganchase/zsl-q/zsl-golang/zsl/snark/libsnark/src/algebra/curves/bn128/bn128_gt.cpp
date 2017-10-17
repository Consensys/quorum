/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/bn128/bn128_gt.hpp"

namespace libsnark {

bn128_GT bn128_GT::GT_one;
bn128_GT::bn128_GT()
{
    this->elem.clear();
}

bool bn128_GT::operator==(const bn128_GT &other) const
{
    return (this->elem == other.elem);
}

bool bn128_GT::operator!=(const bn128_GT& other) const
{
    return !(operator==(other));
}

bn128_GT bn128_GT::operator*(const bn128_GT &other) const
{
    bn128_GT result;
    bn::Fp12::mul(result.elem, this->elem, other.elem);
    return result;
}

bn128_GT bn128_GT::unitary_inverse() const
{
    bn128_GT result(*this);
    bn::Fp6::neg(result.elem.b_, result.elem.b_);
    return result;
}

bn128_GT bn128_GT::one()
{
    return GT_one;
}

std::ostream& operator<<(std::ostream &out, const bn128_GT &g)
{
#ifndef BINARY_OUTPUT
    out << g.elem.a_ << OUTPUT_SEPARATOR << g.elem.b_;
#else
    out.write((char*) &g.elem.a_, sizeof(g.elem.a_));
    out.write((char*) &g.elem.b_, sizeof(g.elem.b_));
#endif
    return out;
}

std::istream& operator>>(std::istream &in, bn128_GT &g)
{
#ifndef BINARY_OUTPUT
    in >> g.elem.a_;
    consume_OUTPUT_SEPARATOR(in);
    in >> g.elem.b_;
#else
    in.read((char*) &g.elem.a_, sizeof(g.elem.a_));
    in.read((char*) &g.elem.b_, sizeof(g.elem.b_));
#endif
    return in;
}
} // libsnark

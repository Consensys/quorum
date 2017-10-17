/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/bn128/bn128_g2.hpp"
#include "algebra/curves/bn128/bn_utils.hpp"

namespace libsnark {

#ifdef PROFILE_OP_COUNTS
long long bn128_G2::add_cnt = 0;
long long bn128_G2::dbl_cnt = 0;
#endif

std::vector<size_t> bn128_G2::wnaf_window_table;
std::vector<size_t> bn128_G2::fixed_base_exp_window_table;
bn128_G2 bn128_G2::G2_zero;
bn128_G2 bn128_G2::G2_one;

bn::Fp2 bn128_G2::sqrt(const bn::Fp2 &el)
{
    size_t v = bn128_Fq2_s;
    bn::Fp2 z = bn128_Fq2_nqr_to_t;
    bn::Fp2 w = mie::power(el, bn128_Fq2_t_minus_1_over_2);
    bn::Fp2 x = el * w;
    bn::Fp2 b = x * w;

#if DEBUG
    // check if square with Euler's criterion
    bn::Fp2 check = b;
    for (size_t i = 0; i < v-1; ++i)
    {
        bn::Fp2::square(check, check);
    }

    assert(check == bn::Fp2(bn::Fp(1), bn::Fp(0)));
#endif

    // compute square root with Tonelli--Shanks
    // (does not terminate if not a square!)

    while (b != bn::Fp2(1))
    {
        size_t m = 0;
        bn::Fp2 b2m = b;
        while (b2m != bn::Fp2(bn::Fp(1), bn::Fp(0)))
        {
            // invariant: b2m = b^(2^m) after entering this loop
            bn::Fp2::square(b2m, b2m);
            m += 1;
        }

        int j = v-m-1;
        w = z;
        while (j > 0)
        {
            bn::Fp2::square(w, w);
            --j;
        } // w = z^2^(v-m-1)

        z = w * w;
        b = b * z;
        x = x * w;
        v = m;
    }

    return x;
}

bn128_G2::bn128_G2()
{
    this->coord[0] = G2_zero.coord[0];
    this->coord[1] = G2_zero.coord[1];
    this->coord[2] = G2_zero.coord[2];
}

void bn128_G2::print() const
{
    if (this->is_zero())
    {
        printf("O\n");
    }
    else
    {
        bn128_G2 copy(*this);
        copy.to_affine_coordinates();
        std::cout << "(" << copy.coord[0].toString(10) << " : " << copy.coord[1].toString(10) << " : " << copy.coord[2].toString(10) << ")\n";
    }
}

void bn128_G2::print_coordinates() const
{
    if (this->is_zero())
    {
        printf("O\n");
    }
    else
    {
        std::cout << "(" << coord[0].toString(10) << " : " << coord[1].toString(10) << " : " << coord[2].toString(10) << ")\n";
    }
}

void bn128_G2::to_affine_coordinates()
{
    if (this->is_zero())
    {
        coord[0] = 0;
        coord[1] = 1;
        coord[2] = 0;
    }
    else
    {
        bn::Fp2 r;
        r = coord[2];
        r.inverse();
        bn::Fp2::square(coord[2], r);
        coord[0] *= coord[2];
        r *= coord[2];
        coord[1] *= r;
        coord[2] = 1;
    }
}

void bn128_G2::to_special()
{
    this->to_affine_coordinates();
}

bool bn128_G2::is_special() const
{
    return (this->is_zero() || this->coord[2] == 1);
}

bool bn128_G2::is_zero() const
{
    return coord[2].isZero();
}

bool bn128_G2::operator==(const bn128_G2 &other) const
{
    if (this->is_zero())
    {
        return other.is_zero();
    }

    if (other.is_zero())
    {
        return false;
    }

    /* now neither is O */

    bn::Fp2 Z1sq, Z2sq, lhs, rhs;
    bn::Fp2::square(Z1sq, this->coord[2]);
    bn::Fp2::square(Z2sq, other.coord[2]);
    bn::Fp2::mul(lhs, Z2sq, this->coord[0]);
    bn::Fp2::mul(rhs, Z1sq, other.coord[0]);

    if (lhs != rhs)
    {
        return false;
    }

    bn::Fp2 Z1cubed, Z2cubed;
    bn::Fp2::mul(Z1cubed, Z1sq, this->coord[2]);
    bn::Fp2::mul(Z2cubed, Z2sq, other.coord[2]);
    bn::Fp2::mul(lhs, Z2cubed, this->coord[1]);
    bn::Fp2::mul(rhs, Z1cubed, other.coord[1]);

    return (lhs == rhs);
}

bool bn128_G2::operator!=(const bn128_G2& other) const
{
    return !(operator==(other));
}

bn128_G2 bn128_G2::operator+(const bn128_G2 &other) const
{
    // handle special cases having to do with O
    if (this->is_zero())
    {
        return other;
    }

    if (other.is_zero())
    {
        return *this;
    }

    // no need to handle points of order 2,4
    // (they cannot exist in a prime-order subgroup)

    // handle double case, and then all the rest
    if (this->operator==(other))
    {
        return this->dbl();
    }
    else
    {
        return this->add(other);
    }
}

bn128_G2 bn128_G2::operator-() const
{
    bn128_G2 result(*this);
    bn::Fp2::neg(result.coord[1], result.coord[1]);
    return result;
}

bn128_G2 bn128_G2::operator-(const bn128_G2 &other) const
{
    return (*this) + (-other);
}

bn128_G2 bn128_G2::add(const bn128_G2 &other) const
{
#ifdef PROFILE_OP_COUNTS
    this->add_cnt++;
#endif

    bn128_G2 result;
    bn::ecop::ECAdd(result.coord, this->coord, other.coord);
    return result;
}

bn128_G2 bn128_G2::mixed_add(const bn128_G2 &other) const
{
    if (this->is_zero())
    {
        return other;
    }

    if (other.is_zero())
    {
        return *this;
    }

    // no need to handle points of order 2,4
    // (they cannot exist in a prime-order subgroup)

#ifdef DEBUG
    assert(other.is_special());
#endif

    // check for doubling case

    // using Jacobian coordinates so:
    // (X1:Y1:Z1) = (X2:Y2:Z2)
    // iff
    // X1/Z1^2 == X2/Z2^2 and Y1/Z1^3 == Y2/Z2^3
    // iff
    // X1 * Z2^2 == X2 * Z1^2 and Y1 * Z2^3 == Y2 * Z1^3

    // we know that Z2 = 1

    bn::Fp2 Z1Z1;
    bn::Fp2::square(Z1Z1, this->coord[2]);
    const bn::Fp2 &U1 = this->coord[0];
    bn::Fp2 U2;
    bn::Fp2::mul(U2, other.coord[0], Z1Z1);
    bn::Fp2 Z1_cubed;
    bn::Fp2::mul(Z1_cubed, this->coord[2], Z1Z1);

    const bn::Fp2 &S1 = this->coord[1];
    bn::Fp2 S2;
    bn::Fp2::mul(S2, other.coord[1], Z1_cubed); // S2 = Y2*Z1*Z1Z1

    if (U1 == U2 && S1 == S2)
    {
        // dbl case; nothing of above can be reused
        return this->dbl();
    }

#ifdef PROFILE_OP_COUNTS
    this->add_cnt++;
#endif

    bn128_G2 result;
    bn::Fp2 H, HH, I, J, r, V, tmp;
    // H = U2-X1
    bn::Fp2::sub(H, U2, this->coord[0]);
    // HH = H^2
    bn::Fp2::square(HH, H);
    // I = 4*HH
    bn::Fp2::add(tmp, HH, HH);
    bn::Fp2::add(I, tmp, tmp);
    // J = H*I
    bn::Fp2::mul(J, H, I);
    // r = 2*(S2-Y1)
    bn::Fp2::sub(tmp, S2, this->coord[1]);
    bn::Fp2::add(r, tmp, tmp);
    // V = X1*I
    bn::Fp2::mul(V, this->coord[0], I);
    // X3 = r^2-J-2*V
    bn::Fp2::square(result.coord[0], r);
    bn::Fp2::sub(result.coord[0], result.coord[0], J);
    bn::Fp2::sub(result.coord[0], result.coord[0], V);
    bn::Fp2::sub(result.coord[0], result.coord[0], V);
    // Y3 = r*(V-X3)-2*Y1*J
    bn::Fp2::sub(tmp, V, result.coord[0]);
    bn::Fp2::mul(result.coord[1], r, tmp);
    bn::Fp2::mul(tmp, this->coord[1], J);
    bn::Fp2::sub(result.coord[1], result.coord[1], tmp);
    bn::Fp2::sub(result.coord[1], result.coord[1], tmp);
    // Z3 = (Z1+H)^2-Z1Z1-HH
    bn::Fp2::add(tmp, this->coord[2], H);
    bn::Fp2::square(result.coord[2], tmp);
    bn::Fp2::sub(result.coord[2], result.coord[2], Z1Z1);
    bn::Fp2::sub(result.coord[2], result.coord[2], HH);
    return result;
}

bn128_G2 bn128_G2::dbl() const
{
#ifdef PROFILE_OP_COUNTS
    this->dbl_cnt++;
#endif
    bn128_G2 result;
    bn::ecop::ECDouble(result.coord, this->coord);
    return result;
}

bool bn128_G2::is_well_formed() const
{
    if (this->is_zero())
    {
        return true;
    }
    else
    {
        /*
          y^2 = x^3 + b

          We are using Jacobian coordinates, so equation we need to check is actually

          (y/z^3)^2 = (x/z^2)^3 + b
          y^2 / z^6 = x^3 / z^6 + b
          y^2 = x^3 + b z^6
        */
        bn::Fp2 X2, Y2, Z2;
        bn::Fp2::square(X2, this->coord[0]);
        bn::Fp2::square(Y2, this->coord[1]);
        bn::Fp2::square(Z2, this->coord[2]);

        bn::Fp2 X3, Z3, Z6;
        bn::Fp2::mul(X3, X2, this->coord[0]);
        bn::Fp2::mul(Z3, Z2, this->coord[2]);
        bn::Fp2::square(Z6, Z3);

        return (Y2 == X3 + bn128_twist_coeff_b * Z6);
    }
}

bn128_G2 bn128_G2::zero()
{
    return G2_zero;
}

bn128_G2 bn128_G2::one()
{
    return G2_one;
}

bn128_G2 bn128_G2::random_element()
{
    return bn128_Fr::random_element().as_bigint() * G2_one;
}

std::ostream& operator<<(std::ostream &out, const bn128_G2 &g)
{
    bn128_G2 gcopy(g);
    gcopy.to_affine_coordinates();

    out << (gcopy.is_zero() ? '1' : '0') << OUTPUT_SEPARATOR;

#ifdef NO_PT_COMPRESSION
    /* no point compression case */
#ifndef BINARY_OUTPUT
    out << gcopy.coord[0].a_ << OUTPUT_SEPARATOR << gcopy.coord[0].b_ << OUTPUT_SEPARATOR;
    out << gcopy.coord[1].a_ << OUTPUT_SEPARATOR << gcopy.coord[1].b_;
#else
    out.write((char*) &gcopy.coord[0].a_, sizeof(gcopy.coord[0].a_));
    out.write((char*) &gcopy.coord[0].b_, sizeof(gcopy.coord[0].b_));
    out.write((char*) &gcopy.coord[1].a_, sizeof(gcopy.coord[1].a_));
    out.write((char*) &gcopy.coord[1].b_, sizeof(gcopy.coord[1].b_));
#endif

#else
    /* point compression case */
#ifndef BINARY_OUTPUT
    out << gcopy.coord[0].a_ << OUTPUT_SEPARATOR << gcopy.coord[0].b_;
#else
    out.write((char*) &gcopy.coord[0].a_, sizeof(gcopy.coord[0].a_));
    out.write((char*) &gcopy.coord[0].b_, sizeof(gcopy.coord[0].b_));
#endif
    out << OUTPUT_SEPARATOR << (((unsigned char*)&gcopy.coord[1].a_)[0] & 1 ? '1' : '0');
#endif

    return out;
}

std::istream& operator>>(std::istream &in, bn128_G2 &g)
{
    char is_zero;
    in.read((char*)&is_zero, 1); // this reads is_zero;
    is_zero -= '0';
    consume_OUTPUT_SEPARATOR(in);

#ifdef NO_PT_COMPRESSION
    /* no point compression case */
#ifndef BINARY_OUTPUT
    in >> g.coord[0].a_;
    consume_OUTPUT_SEPARATOR(in);
    in >> g.coord[0].b_;
    consume_OUTPUT_SEPARATOR(in);
    in >> g.coord[1].a_;
    consume_OUTPUT_SEPARATOR(in);
    in >> g.coord[1].b_;
#else
    in.read((char*) &g.coord[0].a_, sizeof(g.coord[0].a_));
    in.read((char*) &g.coord[0].b_, sizeof(g.coord[0].b_));
    in.read((char*) &g.coord[1].a_, sizeof(g.coord[1].a_));
    in.read((char*) &g.coord[1].b_, sizeof(g.coord[1].b_));
#endif

#else
    /* point compression case */
    bn::Fp2 tX;
#ifndef BINARY_OUTPUT
    in >> tX.a_;
    consume_OUTPUT_SEPARATOR(in);
    in >> tX.b_;
#else
    in.read((char*)&tX.a_, sizeof(tX.a_));
    in.read((char*)&tX.b_, sizeof(tX.b_));
#endif
    consume_OUTPUT_SEPARATOR(in);
    unsigned char Y_lsb;
    in.read((char*)&Y_lsb, 1);
    Y_lsb -= '0';

    // y = +/- sqrt(x^3 + b)
    if (!is_zero)
    {
        g.coord[0] = tX;
        bn::Fp2 tX2, tY2;
        bn::Fp2::square(tX2, tX);
        bn::Fp2::mul(tY2, tX2, tX);
        bn::Fp2::add(tY2, tY2, bn128_twist_coeff_b);

        g.coord[1] = bn128_G2::sqrt(tY2);
        if ((((unsigned char*)&g.coord[1].a_)[0] & 1) != Y_lsb)
        {
            bn::Fp2::neg(g.coord[1], g.coord[1]);
        }
    }
#endif

    /* finalize */
    if (!is_zero)
    {
        g.coord[2] = bn::Fp2(bn::Fp(1), bn::Fp(0));
    }
    else
    {
        g = bn128_G2::zero();
    }

    return in;
}

template<>
void batch_to_special_all_non_zeros<bn128_G2>(std::vector<bn128_G2> &vec)
{
    std::vector<bn::Fp2> Z_vec;
    Z_vec.reserve(vec.size());

    for (auto &el: vec)
    {
        Z_vec.emplace_back(el.coord[2]);
    }
    bn_batch_invert<bn::Fp2>(Z_vec);

    const bn::Fp2 one = 1;

    for (size_t i = 0; i < vec.size(); ++i)
    {
        bn::Fp2 Z2, Z3;
        bn::Fp2::square(Z2, Z_vec[i]);
        bn::Fp2::mul(Z3, Z2, Z_vec[i]);

        bn::Fp2::mul(vec[i].coord[0], vec[i].coord[0], Z2);
        bn::Fp2::mul(vec[i].coord[1], vec[i].coord[1], Z3);
        vec[i].coord[2] = one;
    }
}

} // libsnark

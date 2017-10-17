/** @file
 *****************************************************************************

 Implementation of interfaces for the MNT4 G2 group.

 See mnt4_g2.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/mnt/mnt4/mnt4_g2.hpp"

namespace libsnark {

#ifdef PROFILE_OP_COUNTS
long long mnt4_G2::add_cnt = 0;
long long mnt4_G2::dbl_cnt = 0;
#endif

std::vector<size_t> mnt4_G2::wnaf_window_table;
std::vector<size_t> mnt4_G2::fixed_base_exp_window_table;
mnt4_Fq2 mnt4_G2::twist;
mnt4_Fq2 mnt4_G2::coeff_a;
mnt4_Fq2 mnt4_G2::coeff_b;
mnt4_G2 mnt4_G2::G2_zero;
mnt4_G2 mnt4_G2::G2_one;

mnt4_Fq2 mnt4_G2::mul_by_a(const mnt4_Fq2 &elt)
{
    return mnt4_Fq2(mnt4_twist_mul_by_a_c0 * elt.c0, mnt4_twist_mul_by_a_c1 * elt.c1);
}

mnt4_Fq2 mnt4_G2::mul_by_b(const mnt4_Fq2 &elt)
{
    return mnt4_Fq2(mnt4_twist_mul_by_b_c0 * elt.c1, mnt4_twist_mul_by_b_c1 * elt.c0);
}

mnt4_G2::mnt4_G2()
{
    this->X_ = G2_zero.X_;
    this->Y_ = G2_zero.Y_;
    this->Z_ = G2_zero.Z_;
}

void mnt4_G2::print() const
{
    if (this->is_zero())
    {
        printf("O\n");
    }
    else
    {
        mnt4_G2 copy(*this);
        copy.to_affine_coordinates();
        gmp_printf("(%Nd*z + %Nd , %Nd*z + %Nd)\n",
                   copy.X_.c1.as_bigint().data, mnt4_Fq::num_limbs,
                   copy.X_.c0.as_bigint().data, mnt4_Fq::num_limbs,
                   copy.Y_.c1.as_bigint().data, mnt4_Fq::num_limbs,
                   copy.Y_.c0.as_bigint().data, mnt4_Fq::num_limbs);
    }
}

void mnt4_G2::print_coordinates() const
{
    if (this->is_zero())
    {
        printf("O\n");
    }
    else
    {
        gmp_printf("(%Nd*z + %Nd : %Nd*z + %Nd : %Nd*z + %Nd)\n",
                   this->X_.c1.as_bigint().data, mnt4_Fq::num_limbs,
                   this->X_.c0.as_bigint().data, mnt4_Fq::num_limbs,
                   this->Y_.c1.as_bigint().data, mnt4_Fq::num_limbs,
                   this->Y_.c0.as_bigint().data, mnt4_Fq::num_limbs,
                   this->Z_.c1.as_bigint().data, mnt4_Fq::num_limbs,
                   this->Z_.c0.as_bigint().data, mnt4_Fq::num_limbs);
    }
}

void mnt4_G2::to_affine_coordinates()
{
    if (this->is_zero())
    {
        this->X_ = mnt4_Fq2::zero();
        this->Y_ = mnt4_Fq2::one();
        this->Z_ = mnt4_Fq2::zero();
    }
    else
    {
        const mnt4_Fq2 Z_inv = Z_.inverse();
        X_ = X_ * Z_inv;
        Y_ = Y_ * Z_inv;
        Z_ = mnt4_Fq2::one();
    }
}

void mnt4_G2::to_special()
{
    this->to_affine_coordinates();
}

bool mnt4_G2::is_special() const
{
    return (this->is_zero() || this->Z_ == mnt4_Fq2::one());
}

bool mnt4_G2::is_zero() const
{
    return (this->X_.is_zero() && this->Z_.is_zero());
}

bool mnt4_G2::operator==(const mnt4_G2 &other) const
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

    // X1/Z1 = X2/Z2 <=> X1*Z2 = X2*Z1
    if ((this->X_ * other.Z_) != (other.X_ * this->Z_))
    {
        return false;
    }

    // Y1/Z1 = Y2/Z2 <=> Y1*Z2 = Y2*Z1
    if ((this->Y_ * other.Z_) != (other.Y_ * this->Z_))
    {
        return false;
    }

    return true;
}

bool mnt4_G2::operator!=(const mnt4_G2& other) const
{
    return !(operator==(other));
}

mnt4_G2 mnt4_G2::operator+(const mnt4_G2 &other) const
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
    /*
      The code below is equivalent to (but faster than) the snippet below:

      if (this->operator==(other))
      {
      return this->dbl();
      }
      else
      {
      return this->add(other);
      }
    */

    const mnt4_Fq2 X1Z2 = (this->X_) * (other.Z_);        // X1Z2 = X1*Z2
    const mnt4_Fq2 X2Z1 = (this->Z_) * (other.X_);        // X2Z1 = X2*Z1

    // (used both in add and double checks)

    const mnt4_Fq2 Y1Z2 = (this->Y_) * (other.Z_);        // Y1Z2 = Y1*Z2
    const mnt4_Fq2 Y2Z1 = (this->Z_) * (other.Y_);        // Y2Z1 = Y2*Z1

    if (X1Z2 == X2Z1 && Y1Z2 == Y2Z1)
    {
        // perform dbl case
        const mnt4_Fq2 XX   = (this->X_).squared();                   // XX  = X1^2
        const mnt4_Fq2 ZZ   = (this->Z_).squared();                   // ZZ  = Z1^2
        const mnt4_Fq2 w    = mnt4_G2::mul_by_a(ZZ) + (XX + XX + XX); // w   = a*ZZ + 3*XX
        const mnt4_Fq2 Y1Z1 = (this->Y_) * (this->Z_);
        const mnt4_Fq2 s    = Y1Z1 + Y1Z1;                            // s   = 2*Y1*Z1
        const mnt4_Fq2 ss   = s.squared();                            // ss  = s^2
        const mnt4_Fq2 sss  = s * ss;                                 // sss = s*ss
        const mnt4_Fq2 R    = (this->Y_) * s;                         // R   = Y1*s
        const mnt4_Fq2 RR   = R.squared();                            // RR  = R^2
        const mnt4_Fq2 B    = ((this->X_)+R).squared()-XX-RR;         // B   = (X1+R)^2 - XX - RR
        const mnt4_Fq2 h    = w.squared() - (B+B);                    // h   = w^2 - 2*B
        const mnt4_Fq2 X3   = h * s;                                  // X3  = h*s
        const mnt4_Fq2 Y3   = w * (B-h)-(RR+RR);                      // Y3  = w*(B-h) - 2*RR
        const mnt4_Fq2 Z3   = sss;                                    // Z3  = sss

        return mnt4_G2(X3, Y3, Z3);
    }

    // if we have arrived here we are in the add case
    const mnt4_Fq2 Z1Z2 = (this->Z_) * (other.Z_);      // Z1Z2 = Z1*Z2
    const mnt4_Fq2 u    = Y2Z1 - Y1Z2;                  // u    = Y2*Z1-Y1Z2
    const mnt4_Fq2 uu   = u.squared();                  // uu   = u^2
    const mnt4_Fq2 v    = X2Z1 - X1Z2;                  // v    = X2*Z1-X1Z2
    const mnt4_Fq2 vv   = v.squared();                  // vv   = v^2
    const mnt4_Fq2 vvv  = v * vv;                       // vvv  = v*vv
    const mnt4_Fq2 R    = vv * X1Z2;                    // R    = vv*X1Z2
    const mnt4_Fq2 A    = uu * Z1Z2 - (vvv + R + R);    // A    = uu*Z1Z2 - vvv - 2*R
    const mnt4_Fq2 X3   = v * A;                        // X3   = v*A
    const mnt4_Fq2 Y3   = u * (R-A) - vvv * Y1Z2;       // Y3   = u*(R-A) - vvv*Y1Z2
    const mnt4_Fq2 Z3   = vvv * Z1Z2;                   // Z3   = vvv*Z1Z2

    return mnt4_G2(X3, Y3, Z3);
}

mnt4_G2 mnt4_G2::operator-() const
{
    return mnt4_G2(this->X_, -(this->Y_), this->Z_);
}


mnt4_G2 mnt4_G2::operator-(const mnt4_G2 &other) const
{
    return (*this) + (-other);
}

mnt4_G2 mnt4_G2::add(const mnt4_G2 &other) const
{
    // handle special cases having to do with O
    if (this->is_zero())
    {
        return other;
    }

    if (other.is_zero())
    {
        return (*this);
    }

    // no need to handle points of order 2,4
    // (they cannot exist in a prime-order subgroup)

    // handle double case
    if (this->operator==(other))
    {
        return this->dbl();
    }

#ifdef PROFILE_OP_COUNTS
    this->add_cnt++;
#endif
    // NOTE: does not handle O and pts of order 2,4
    // http://www.hyperelliptic.org/EFD/g1p/auto-shortw-projective.html#addition-add-1998-cmo-2

    const mnt4_Fq2 Y1Z2 = (this->Y_) * (other.Z_);        // Y1Z2 = Y1*Z2
    const mnt4_Fq2 X1Z2 = (this->X_) * (other.Z_);        // X1Z2 = X1*Z2
    const mnt4_Fq2 Z1Z2 = (this->Z_) * (other.Z_);        // Z1Z2 = Z1*Z2
    const mnt4_Fq2 u    = (other.Y_) * (this->Z_) - Y1Z2; // u    = Y2*Z1-Y1Z2
    const mnt4_Fq2 uu   = u.squared();                    // uu   = u^2
    const mnt4_Fq2 v    = (other.X_) * (this->Z_) - X1Z2; // v    = X2*Z1-X1Z2
    const mnt4_Fq2 vv   = v.squared();                    // vv   = v^2
    const mnt4_Fq2 vvv  = v * vv;                         // vvv  = v*vv
    const mnt4_Fq2 R    = vv * X1Z2;                      // R    = vv*X1Z2
    const mnt4_Fq2 A    = uu * Z1Z2 - (vvv + R + R);      // A    = uu*Z1Z2 - vvv - 2*R
    const mnt4_Fq2 X3   = v * A;                          // X3   = v*A
    const mnt4_Fq2 Y3   = u * (R-A) - vvv * Y1Z2;         // Y3   = u*(R-A) - vvv*Y1Z2
    const mnt4_Fq2 Z3   = vvv * Z1Z2;                     // Z3   = vvv*Z1Z2

    return mnt4_G2(X3, Y3, Z3);
}

mnt4_G2 mnt4_G2::mixed_add(const mnt4_G2 &other) const
{
#ifdef PROFILE_OP_COUNTS
    this->add_cnt++;
#endif
    // NOTE: does not handle O and pts of order 2,4
    // http://www.hyperelliptic.org/EFD/g1p/auto-shortw-projective.html#addition-add-1998-cmo-2
    //assert(other.Z == mnt4_Fq2::one());

    if (this->is_zero())
    {
        return other;
    }

    if (other.is_zero())
    {
        return (*this);
    }

#ifdef DEBUG
    assert(other.is_special());
#endif

    const mnt4_Fq2 &X1Z2 = (this->X_);                   // X1Z2 = X1*Z2 (but other is special and not zero)
    const mnt4_Fq2 X2Z1 = (this->Z_) * (other.X_);       // X2Z1 = X2*Z1

    // (used both in add and double checks)

    const mnt4_Fq2 &Y1Z2 = (this->Y_);                   // Y1Z2 = Y1*Z2 (but other is special and not zero)
    const mnt4_Fq2 Y2Z1 = (this->Z_) * (other.Y_);       // Y2Z1 = Y2*Z1

    if (X1Z2 == X2Z1 && Y1Z2 == Y2Z1)
    {
        return this->dbl();
    }

    const mnt4_Fq2 u = Y2Z1 - this->Y_;              // u = Y2*Z1-Y1
    const mnt4_Fq2 uu = u.squared();                 // uu = u2
    const mnt4_Fq2 v = X2Z1 - this->X_;              // v = X2*Z1-X1
    const mnt4_Fq2 vv = v.squared();                 // vv = v2
    const mnt4_Fq2 vvv = v*vv;                       // vvv = v*vv
    const mnt4_Fq2 R = vv * this->X_;                // R = vv*X1
    const mnt4_Fq2 A = uu * this->Z_ - vvv - R - R;  // A = uu*Z1-vvv-2*R
    const mnt4_Fq2 X3 = v * A;                       // X3 = v*A
    const mnt4_Fq2 Y3 = u*(R-A) - vvv * this->Y_;    // Y3 = u*(R-A)-vvv*Y1
    const mnt4_Fq2 Z3 = vvv * this->Z_;              // Z3 = vvv*Z1

    return mnt4_G2(X3, Y3, Z3);
}

mnt4_G2 mnt4_G2::dbl() const
{
#ifdef PROFILE_OP_COUNTS
    this->dbl_cnt++;
#endif
    if (this->is_zero())
    {
        return (*this);
    }
    else
    {
        // NOTE: does not handle O and pts of order 2,4
        // http://www.hyperelliptic.org/EFD/g1p/auto-shortw-projective.html#doubling-dbl-2007-bl

        const mnt4_Fq2 XX   = (this->X_).squared();                     // XX  = X1^2
        const mnt4_Fq2 ZZ   = (this->Z_).squared();                     // ZZ  = Z1^2
        const mnt4_Fq2 w    = mnt4_G2::mul_by_a(ZZ) + (XX + XX + XX);   // w   = a*ZZ + 3*XX
        const mnt4_Fq2 Y1Z1 = (this->Y_) * (this->Z_);
        const mnt4_Fq2 s    = Y1Z1 + Y1Z1;                              // s   = 2*Y1*Z1
        const mnt4_Fq2 ss   = s.squared();                              // ss  = s^2
        const mnt4_Fq2 sss  = s * ss;                                   // sss = s*ss
        const mnt4_Fq2 R    = (this->Y_) * s;                           // R   = Y1*s
        const mnt4_Fq2 RR   = R.squared();                              // RR  = R^2
        const mnt4_Fq2 B    = ((this->X_)+R).squared()-XX-RR;           // B   = (X1+R)^2 - XX - RR
        const mnt4_Fq2 h    = w.squared() - (B+B);                      // h   = w^2-2*B
        const mnt4_Fq2 X3   = h * s;                                    // X3  = h*s
        const mnt4_Fq2 Y3   = w * (B-h)-(RR+RR);                        // Y3  = w*(B-h) - 2*RR
        const mnt4_Fq2 Z3   = sss;                                      // Z3  = sss

        return mnt4_G2(X3, Y3, Z3);
    }
}

mnt4_G2 mnt4_G2::mul_by_q() const
{
    return mnt4_G2(mnt4_twist_mul_by_q_X * (this->X_).Frobenius_map(1),
                   mnt4_twist_mul_by_q_Y * (this->Y_).Frobenius_map(1),
                   (this->Z_).Frobenius_map(1));
}

bool mnt4_G2::is_well_formed() const
{
    if (this->is_zero())
    {
        return true;
    }
    else
    {
        /*
          y^2 = x^3 + ax + b

          We are using projective, so equation we need to check is actually

          (y/z)^2 = (x/z)^3 + a (x/z) + b
          z y^2 = x^3  + a z^2 x + b z^3

          z (y^2 - b z^2) = x ( x^2 + a z^2)
        */
        const mnt4_Fq2 X2 = this->X_.squared();
        const mnt4_Fq2 Y2 = this->Y_.squared();
        const mnt4_Fq2 Z2 = this->Z_.squared();
        const mnt4_Fq2 aZ2 =  mnt4_twist_coeff_a * Z2;

        return (this->Z_ * (Y2 - mnt4_twist_coeff_b * Z2) == this->X_ * (X2 + aZ2));
    }
}

mnt4_G2 mnt4_G2::zero()
{
    return G2_zero;
}

mnt4_G2 mnt4_G2::one()
{
    return G2_one;
}

mnt4_G2 mnt4_G2::random_element()
{
    return (mnt4_Fr::random_element().as_bigint()) * G2_one;
}

std::ostream& operator<<(std::ostream &out, const mnt4_G2 &g)
{
    mnt4_G2 copy(g);
    copy.to_affine_coordinates();

    out << (copy.is_zero() ? 1 : 0) << OUTPUT_SEPARATOR;
#ifdef NO_PT_COMPRESSION
    out << copy.X_ << OUTPUT_SEPARATOR << copy.Y_;
#else
    /* storing LSB of Y */
    out << copy.X_ << OUTPUT_SEPARATOR << (copy.Y_.c0.as_bigint().data[0] & 1);
#endif

    return out;
}

std::istream& operator>>(std::istream &in, mnt4_G2 &g)
{
    char is_zero;
    mnt4_Fq2 tX, tY;

#ifdef NO_PT_COMPRESSION
    in >> is_zero >> tX >> tY;
    is_zero -= '0';
#else
    in.read((char*)&is_zero, 1); // this reads is_zero;
    is_zero -= '0';
    consume_OUTPUT_SEPARATOR(in);

    unsigned char Y_lsb;
    in >> tX;
    consume_OUTPUT_SEPARATOR(in);
    in.read((char*)&Y_lsb, 1);
    Y_lsb -= '0';

    // y = +/- sqrt(x^3 + a*x + b)
    if (!is_zero)
    {
        mnt4_Fq2 tX2 = tX.squared();
        mnt4_Fq2 tY2 = (tX2 + mnt4_twist_coeff_a ) * tX + mnt4_twist_coeff_b;
        tY = tY2.sqrt();

        if ((tY.c0.as_bigint().data[0] & 1) != Y_lsb)
        {
            tY = -tY;
        }
    }
#endif
    // using projective coordinates
    if (!is_zero)
    {
        g.X_ = tX;
        g.Y_ = tY;
        g.Z_ = mnt4_Fq2::one();
    }
    else
    {
        g = mnt4_G2::zero();
    }

    return in;
}

template<>
void batch_to_special_all_non_zeros<mnt4_G2>(std::vector<mnt4_G2> &vec)
{
    std::vector<mnt4_Fq2> Z_vec;
    Z_vec.reserve(vec.size());

    for (auto &el: vec)
    {
        Z_vec.emplace_back(el.Z());
    }
    batch_invert<mnt4_Fq2>(Z_vec);

    const mnt4_Fq2 one = mnt4_Fq2::one();

    for (size_t i = 0; i < vec.size(); ++i)
    {
        vec[i] = mnt4_G2(vec[i].X() * Z_vec[i], vec[i].Y() * Z_vec[i], one);
    }
}

} // libsnark

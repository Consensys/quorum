/** @file
 *****************************************************************************

 Implementation of functions for generating randomness.

 See rng.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RNG_TCC_
#define RNG_TCC_

#include <openssl/sha.h>

namespace libsnark {

template<typename FieldT>
FieldT SHA512_rng(const uint64_t idx)
{
    assert(GMP_NUMB_BITS == 64); // current Python code cannot handle larger values, so testing here for some assumptions.
    assert(is_little_endian());

    assert(FieldT::size_in_bits() <= SHA512_DIGEST_LENGTH * 8);

    bigint<FieldT::num_limbs> rval;
    uint64_t iter = 0;
    do
    {
        mp_limb_t hash[((SHA512_DIGEST_LENGTH*8) + GMP_NUMB_BITS - 1)/GMP_NUMB_BITS];

        SHA512_CTX sha512;
        SHA512_Init(&sha512);
        SHA512_Update(&sha512, &idx, sizeof(idx));
        SHA512_Update(&sha512, &iter, sizeof(iter));
        SHA512_Final((unsigned char*)hash, &sha512);

        for (mp_size_t i = 0; i < FieldT::num_limbs; ++i)
        {
            rval.data[i] = hash[i];
        }

        /* clear all bits higher than MSB of modulus */
        size_t bitno = GMP_NUMB_BITS * FieldT::num_limbs;
        while (FieldT::mod.test_bit(bitno) == false)
        {
            const std::size_t part = bitno/GMP_NUMB_BITS;
            const std::size_t bit = bitno - (GMP_NUMB_BITS*part);

            rval.data[part] &= ~(1ul<<bit);

            bitno--;
        }

        ++iter;
    }

    /* if r.data is still >= modulus -- repeat (rejection sampling) */
    while (mpn_cmp(rval.data, FieldT::mod.data, FieldT::num_limbs) >= 0);

    return FieldT(rval);
}

} // libsnark

#endif // RNG_TCC_

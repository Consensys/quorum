/** @file
 *****************************************************************************

 AES-Based PRF for ADSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "common/default_types/r1cs_ppzkadsnark_pp.hpp"
#include "supercop/crypto_core_aes128encrypt.h"
#include "supercop/randombytes.h"
#include "gmp.h"

namespace libsnark {

template <>
aesPrfKeyT prfGen<default_r1cs_ppzkadsnark_pp>() {
    aesPrfKeyT key;
    randombytes(key.key_bytes,32);
    return key;
}

template<>
Fr<snark_pp<default_r1cs_ppzkadsnark_pp>> prfCompute<default_r1cs_ppzkadsnark_pp>(
    const aesPrfKeyT &key,  const labelT &label) {
    unsigned char seed_bytes[16];
    mpz_t aux,Fr_mod;
    unsigned char random_bytes[16*3];
    size_t exp_len;

    mpz_init (aux);
    mpz_init (Fr_mod);

    // compute random seed using AES as PRF
    crypto_core_aes128encrypt_openssl(seed_bytes,label.label_bytes,key.key_bytes,NULL);

    // use first 128 bits of output to seed AES-CTR
    // PRG to expand to 3*128 bits
    crypto_core_aes128encrypt_openssl(random_bytes,seed_bytes,key.key_bytes+16,NULL);

    mpz_import(aux, 16, 0, 1, 0, 0, seed_bytes);
    mpz_add_ui(aux,aux,1);
    mpz_export(seed_bytes, &exp_len, 0, 1, 0, 0, aux);
    while (exp_len < 16)
        seed_bytes[exp_len++] = 0;

    crypto_core_aes128encrypt_openssl(random_bytes+16,seed_bytes,key.key_bytes+16,NULL);

    mpz_add_ui(aux,aux,1);
    mpz_export(seed_bytes, &exp_len, 0, 1, 0, 0, aux);
    while (exp_len < 16)
        seed_bytes[exp_len++] = 0;

    crypto_core_aes128encrypt_openssl(random_bytes+32,seed_bytes,key.key_bytes+16,NULL);

    // see output as integer and reduce modulo r
    mpz_import(aux, 16*3, 0, 1, 0, 0, random_bytes);
    Fr<snark_pp<default_r1cs_ppzkadsnark_pp>>::mod.to_mpz(Fr_mod);
    mpz_mod(aux,aux,Fr_mod);

    return Fr<snark_pp<default_r1cs_ppzkadsnark_pp>>(
        bigint<Fr<snark_pp<default_r1cs_ppzkadsnark_pp>>::num_limbs>(aux));
}

} // libsnark

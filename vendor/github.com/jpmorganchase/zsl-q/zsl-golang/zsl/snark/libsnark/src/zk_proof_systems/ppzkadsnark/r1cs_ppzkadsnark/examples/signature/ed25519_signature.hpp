/** @file
 *****************************************************************************

 Fast batch verification signature for ADSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

/** @file
 *****************************************************************************
 * @author     This file was deed to libsnark by Manuel Barbosa.
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef ED25519SIG_HPP_
#define ED25519SIG_HPP_

#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark_signature.hpp"

namespace libsnark {

class ed25519_sigT {
public:
    unsigned char sig_bytes[64];
};

class ed25519_vkT {
public:
    unsigned char vk_bytes[32];
};

class ed25519_skT {
public:
    unsigned char sk_bytes[64];
};

} // libsnark

#endif // ED25519SIG_HPP_

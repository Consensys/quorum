/** @file
 *****************************************************************************

 Generic signature interface for ADSNARK.

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

#ifndef SIGNATURE_HPP_
#define SIGNATURE_HPP_

#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark_params.hpp"

namespace libsnark {

template<typename ppT>
class kpT {
public:
    r1cs_ppzkadsnark_skT<ppT> sk;
    r1cs_ppzkadsnark_vkT<ppT> vk;
};

template<typename ppT>
kpT<ppT> sigGen(void);

template<typename ppT>
r1cs_ppzkadsnark_sigT<ppT> sigSign(const r1cs_ppzkadsnark_skT<ppT> &sk, const labelT &label, const G2<snark_pp<ppT>> &Lambda);

template<typename ppT>
bool sigVerif(const r1cs_ppzkadsnark_vkT<ppT> &vk, const labelT &label, const G2<snark_pp<ppT>> &Lambda, const r1cs_ppzkadsnark_sigT<ppT> &sig);

template<typename ppT>
bool sigBatchVerif(const r1cs_ppzkadsnark_vkT<ppT> &vk, const std::vector<labelT> &labels, const std::vector<G2<snark_pp<ppT>>> &Lambdas, const std::vector<r1cs_ppzkadsnark_sigT<ppT>> &sigs);

} // libsnark

#endif // SIGNATURE_HPP_

/** @file
 *****************************************************************************

 Generic PRF interface for ADSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef PRF_HPP_
#define PRF_HPP_

#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/r1cs_ppzkadsnark_params.hpp"

namespace libsnark {

template <typename ppT>
r1cs_ppzkadsnark_prfKeyT<ppT> prfGen();

template<typename ppT>
Fr<snark_pp<ppT>> prfCompute(const r1cs_ppzkadsnark_prfKeyT<ppT> &key, const labelT &label);

} // libsnark

#endif // PRF_HPP_

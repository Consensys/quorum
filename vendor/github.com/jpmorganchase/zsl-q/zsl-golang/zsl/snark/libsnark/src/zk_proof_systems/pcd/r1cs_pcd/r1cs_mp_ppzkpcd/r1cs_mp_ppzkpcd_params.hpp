/** @file
 *****************************************************************************

 Parameters for *multi-predicate* ppzkPCD for R1CS.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_MP_PPZKPCD_PARAMS_HPP_
#define R1CS_MP_PPZKPCD_PARAMS_HPP_

#include "algebra/curves/public_params.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/compliance_predicate.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_pcd_params.hpp"

namespace libsnark {

template<typename PCD_ppT>
using r1cs_mp_ppzkpcd_compliance_predicate = r1cs_pcd_compliance_predicate<Fr<typename PCD_ppT::curve_A_pp> >;

template<typename PCD_ppT>
using r1cs_mp_ppzkpcd_message = r1cs_pcd_message<Fr<typename PCD_ppT::curve_A_pp> >;

template<typename PCD_ppT>
using r1cs_mp_ppzkpcd_local_data = r1cs_pcd_local_data<Fr<typename PCD_ppT::curve_A_pp> >;

template<typename PCD_ppT>
using r1cs_mp_ppzkpcd_primary_input = r1cs_pcd_compliance_predicate_primary_input<Fr<typename PCD_ppT::curve_A_pp> >;

template<typename PCD_ppT>
using r1cs_mp_ppzkpcd_auxiliary_input = r1cs_pcd_compliance_predicate_auxiliary_input<Fr<typename PCD_ppT::curve_A_pp> >;

} // libsnark

#endif // R1CS_MP_PPZKPCD_PARAMS_HPP_

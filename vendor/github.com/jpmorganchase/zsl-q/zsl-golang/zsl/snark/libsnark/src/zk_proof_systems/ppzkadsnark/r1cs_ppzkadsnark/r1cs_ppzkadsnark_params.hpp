/** @file
 *****************************************************************************

 Declaration of public-parameter selector for the R1CS ppzkADSNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef R1CS_PPZKADSNARK_PARAMS_HPP_
#define R1CS_PPZKADSNARK_PARAMS_HPP_

#include "relations/constraint_satisfaction_problems/r1cs/r1cs.hpp"

namespace libsnark {

class labelT {
public:
    unsigned char label_bytes[16];
    labelT() {};
};

/**
 * Below are various template aliases (used for convenience).
 */

template<typename r1cs_ppzkadsnark_ppT>
using snark_pp = typename r1cs_ppzkadsnark_ppT::snark_pp;

template<typename r1cs_ppzkadsnark_ppT>
using r1cs_ppzkadsnark_constraint_system = r1cs_constraint_system<Fr<snark_pp<r1cs_ppzkadsnark_ppT>>>;

template<typename r1cs_ppzkadsnark_ppT>
using r1cs_ppzkadsnark_primary_input = r1cs_primary_input<Fr<snark_pp<r1cs_ppzkadsnark_ppT>> >;

template<typename r1cs_ppzkadsnark_ppT>
using r1cs_ppzkadsnark_auxiliary_input = r1cs_auxiliary_input<Fr<snark_pp<r1cs_ppzkadsnark_ppT>> >;

template<typename r1cs_ppzkadsnark_ppT>
using r1cs_ppzkadsnark_skT = typename r1cs_ppzkadsnark_ppT::skT;

template<typename r1cs_ppzkadsnark_ppT>
using r1cs_ppzkadsnark_vkT = typename r1cs_ppzkadsnark_ppT::vkT;

template<typename r1cs_ppzkadsnark_ppT>
using r1cs_ppzkadsnark_sigT = typename r1cs_ppzkadsnark_ppT::sigT;

template<typename r1cs_ppzkadsnark_ppT>
using r1cs_ppzkadsnark_prfKeyT = typename r1cs_ppzkadsnark_ppT::prfKeyT;


} // libsnark

#endif // R1CS_PPZKADSNARK_PARAMS_HPP_


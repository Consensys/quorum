/** @file
 *****************************************************************************

 This file defines the default choices of TinyRAM zk-SNARK.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TINYRAM_PPZKSNARK_PP_HPP_
#define TINYRAM_PPZKSNARK_PP_HPP_

#include "common/default_types/r1cs_ppzkpcd_pp.hpp"
#include "relations/ram_computations/rams/tinyram/tinyram_params.hpp"

namespace libsnark {

class default_tinyram_zksnark_pp {
public:
    typedef default_r1cs_ppzkpcd_pp PCD_pp;
    typedef typename PCD_pp::scalar_field_A FieldT;
    typedef ram_tinyram<FieldT> machine_pp;

    static void init_public_params();
};

} // libsnark

#endif // TINYRAM_PPZKSNARK_PP_HPP_

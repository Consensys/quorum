/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include "common/default_types/r1cs_ppzkpcd_pp.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_mp_ppzkpcd/examples/run_r1cs_mp_ppzkpcd.hpp"

using namespace libsnark;

template<typename PCD_ppT>
void test_tally(const size_t arity, const size_t max_layer, const bool test_multi_type, const bool test_same_type_optimization)

{
    const size_t wordsize = 32;
    const bool test_serialization = true;
    const bool bit = run_r1cs_mp_ppzkpcd_tally_example<PCD_ppT>(wordsize, arity, max_layer, test_serialization, test_multi_type, test_same_type_optimization);
    assert(bit);
}

int main(void)
{
    start_profiling();
    default_r1cs_ppzkpcd_pp::init_public_params();

    const size_t max_arity = 2;
    const size_t max_layer = 2;

    test_tally<default_r1cs_ppzkpcd_pp>(max_arity, max_layer, false, false);
    test_tally<default_r1cs_ppzkpcd_pp>(max_arity, max_layer, false, true);
    test_tally<default_r1cs_ppzkpcd_pp>(max_arity, max_layer, true, false);
    test_tally<default_r1cs_ppzkpcd_pp>(max_arity, max_layer, true, true);
}

/** @file
 *****************************************************************************
 Test program that exercises the ppzkSNARK (first generator, then
 prover, then verifier) on a synthetic USCS instance.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <cassert>
#include <cstdio>

#include "common/default_types/uscs_ppzksnark_pp.hpp"
#include "common/profiling.hpp"
#include "common/utils.hpp"
#include "relations/constraint_satisfaction_problems/uscs/examples/uscs_examples.hpp"
#include "zk_proof_systems/ppzksnark/uscs_ppzksnark/examples/run_uscs_ppzksnark.hpp"

using namespace libsnark;

template<typename ppT>
void test_uscs_ppzksnark(size_t num_constraints,
                         size_t input_size)
{
    print_header("(enter) Test USCS ppzkSNARK");

    const bool test_serialization = true;
    uscs_example<Fr<ppT> > example = generate_uscs_example_with_binary_input<Fr<ppT> >(num_constraints, input_size);
    const bool bit = run_uscs_ppzksnark<ppT>(example, test_serialization);
    assert(bit);

    print_header("(leave) Test USCS ppzkSNARK");
}

int main()
{
    default_uscs_ppzksnark_pp::init_public_params();
    start_profiling();

    test_uscs_ppzksnark<default_uscs_ppzksnark_pp>(1000, 100);
}

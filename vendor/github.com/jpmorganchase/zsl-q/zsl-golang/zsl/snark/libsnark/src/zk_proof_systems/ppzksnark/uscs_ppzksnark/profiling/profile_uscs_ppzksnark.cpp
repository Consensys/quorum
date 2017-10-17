/** @file
 *****************************************************************************
 Profiling program that exercises the ppzkSNARK (first generator, then prover,
 then verifier) on a synthetic USCS instance.

 The command

     $ src/zk_proof_systems/ppzksnark/uscs_ppzksnark/profiling/profile_uscs_ppzksnark 1000 10 Fr

 exercises the ppzkSNARK (first generator, then prover, then verifier) on an USCS instance with 1000 equations and an input consisting of 10 field elements.

 (If you get the error `zmInit ERR:can't protect`, see the discussion [above](#elliptic-curve-choices).)

 The command

     $ src/zk_proof_systems/ppzksnark/uscs_ppzksnark/profiling/profile_uscs_ppzksnark 1000 10 bytes

 does the same but now the input consists of 10 bytes.

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

int main(int argc, const char * argv[])
{
    default_uscs_ppzksnark_pp::init_public_params();
    start_profiling();

    if (argc == 2 && strcmp(argv[1], "-v") == 0)
    {
        print_compilation_info();
        return 0;
    }

    if (argc != 3)
    {
        printf("usage: %s num_constraints input_size\n", argv[0]);
        return 1;
    }

    const int num_constraints = atoi(argv[1]);
    const int input_size = atoi(argv[2]);

    enter_block("Generate USCS example");
    uscs_example<Fr<default_uscs_ppzksnark_pp> > example = generate_uscs_example_with_field_input<Fr<default_uscs_ppzksnark_pp> >(num_constraints, input_size);
    leave_block("Generate USCS example");

    print_header("(enter) Profile USCS ppzkSNARK");
    const bool test_serialization = true;
    run_uscs_ppzksnark<default_uscs_ppzksnark_pp>(example, test_serialization);
    print_header("(leave) Profile USCS ppzkSNARK");
}

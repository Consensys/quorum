/** @file
 *****************************************************************************
 Profiling program that exercises the ppzkSNARK (first generator, then prover,
 then verifier) on a synthetic BACS instance.

 The command

     $ src/zk_proof_systems/bacs_ppzksnark/profiling/profile_bacs_ppzksnark 1000 10

 exercises the ppzkSNARK (first generator, then prover, then verifier) on an BACS instance with 1000 gates and an input consisting of 10 Field elements

 (If you get the error `zmInit ERR:can't protect`, see the discussion [above](#elliptic-curve-choices).)

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <cstdio>

#include "common/default_types/bacs_ppzksnark_pp.hpp"
#include "common/profiling.hpp"
#include "relations/circuit_satisfaction_problems/bacs/examples/bacs_examples.hpp"
#include "zk_proof_systems/ppzksnark/bacs_ppzksnark/examples/run_bacs_ppzksnark.hpp"

using namespace libsnark;

int main(int argc, const char * argv[])
{
    default_bacs_ppzksnark_pp::init_public_params();
    start_profiling();

    if (argc == 2 && strcmp(argv[1], "-v") == 0)
    {
        print_compilation_info();
        return 0;
    }

    if (argc != 3)
    {
        printf("usage: %s num_gates primary_input_size\n", argv[0]);
        return 1;
    }
    const int num_gates = atoi(argv[1]);
    int primary_input_size = atoi(argv[2]);

    const size_t auxiliary_input_size = 0;
    const size_t num_outputs = num_gates / 2;

    enter_block("Generate BACS example");
    bacs_example<Fr<default_bacs_ppzksnark_pp> > example = generate_bacs_example<Fr<default_bacs_ppzksnark_pp> >(primary_input_size, auxiliary_input_size, num_gates, num_outputs);
    leave_block("Generate BACS example");

    print_header("(enter) Profile BACS ppzkSNARK");
    const bool test_serialization = true;
    run_bacs_ppzksnark<default_bacs_ppzksnark_pp>(example, test_serialization);
    print_header("(leave) Profile BACS ppzkSNARK");
}

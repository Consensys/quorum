/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <algorithm>
#include <cstdio>
#include <cassert>
#include <cstring>
#include <vector>

#include "common/default_types/r1cs_ppzkadsnark_pp.hpp"
#include "common/profiling.hpp"
#include "zk_proof_systems/ppzkadsnark/r1cs_ppzkadsnark/examples/run_r1cs_ppzkadsnark.hpp"

using namespace libsnark;

int main(int argc, const char * argv[])
{
    default_r1cs_ppzkadsnark_pp::init_public_params();
    start_profiling();

    if (argc == 2 && strcmp(argv[1], "-v") == 0)
    {
        print_compilation_info();
        return 0;
    }

    if (argc != 3 && argc != 4)
    {
        printf("usage: %s num_constraints input_size [Fr|bytes]\n", argv[0]);
        return 1;
    }
    const int num_constraints = atoi(argv[1]);
    int input_size = atoi(argv[2]);
    if (argc == 4)
    {
        assert(strcmp(argv[3], "Fr") == 0 || strcmp(argv[3], "bytes") == 0);
        if (strcmp(argv[3], "bytes") == 0)
        {
            input_size = div_ceil(8 * input_size, Fr<snark_pp<default_r1cs_ppzkadsnark_pp>>::num_bits - 1);
        }
    }

    enter_block("Generate R1CS example");
    r1cs_example<Fr<snark_pp<default_r1cs_ppzkadsnark_pp>>> example =
        generate_r1cs_example_with_field_input<Fr<snark_pp<default_r1cs_ppzkadsnark_pp>>>
        (num_constraints, input_size);
    leave_block("Generate R1CS example");

    print_header("(enter) Profile R1CS ppzkADSNARK");
    const bool test_serialization = true;
    run_r1cs_ppzkadsnark<default_r1cs_ppzkadsnark_pp>(example, test_serialization);
    print_header("(leave) Profile R1CS ppzkADSNARK");
}

/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <algorithm>
#include <cstring>
#include <iostream>
#include <fstream>
#include <sstream>
#include <string>

#include "common/default_types/ram_ppzksnark_pp.hpp"
#include "common/profiling.hpp"
#include "relations/ram_computations/rams/examples/ram_examples.hpp"
#include "zk_proof_systems/ppzksnark/ram_ppzksnark/examples/run_ram_ppzksnark.hpp"
#include "relations/ram_computations/rams/tinyram/tinyram_params.hpp"

using namespace libsnark;

int main(int argc, const char * argv[])
{
    ram_ppzksnark_snark_pp<default_ram_ppzksnark_pp>::init_public_params();
    start_profiling();

    if (argc == 2 && strcmp(argv[1], "-v") == 0)
    {
        print_compilation_info();
        return 0;
    }

    if (argc != 6)
    {
        printf("usage: %s word_size reg_count program_size input_size time_bound\n", argv[0]);
        return 1;
    }

    const size_t w = atoi(argv[1]),
                 k = atoi(argv[2]),
                 program_size = atoi(argv[3]),
                 input_size = atoi(argv[4]),
                 time_bound = atoi(argv[5]);

    typedef ram_ppzksnark_machine_pp<default_ram_ppzksnark_pp> machine_ppT;

    const ram_ppzksnark_architecture_params<default_ram_ppzksnark_pp> ap(w, k);

    enter_block("Generate RAM example");
    const size_t boot_trace_size_bound = program_size + input_size;
    const bool satisfiable = true;
    ram_example<machine_ppT> example = gen_ram_example_complex<machine_ppT>(ap, boot_trace_size_bound, time_bound, satisfiable);
    enter_block("Generate RAM example");

    print_header("(enter) Profile RAM ppzkSNARK");
    const bool test_serialization = true;
    run_ram_ppzksnark<default_ram_ppzksnark_pp>(example, test_serialization);
    print_header("(leave) Profile RAM ppzkSNARK");
}

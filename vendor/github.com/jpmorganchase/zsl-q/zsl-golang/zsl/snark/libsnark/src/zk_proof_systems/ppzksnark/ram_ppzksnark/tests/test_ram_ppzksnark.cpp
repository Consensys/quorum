/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <algorithm>
#include <iostream>
#include <fstream>
#include <sstream>
#include <string>

#include "common/default_types/ram_ppzksnark_pp.hpp"
#include "common/profiling.hpp"
#include "relations/ram_computations/rams/examples/ram_examples.hpp"
#include "zk_proof_systems/ppzksnark/ram_ppzksnark/examples/run_ram_ppzksnark.hpp"

using namespace libsnark;

template<typename ppT>
void test_ram_ppzksnark(const size_t w,
                        const size_t k,
                        const size_t program_size,
                        const size_t input_size,
                        const size_t time_bound)
{
    print_header("(enter) Test RAM ppzkSNARK");

    typedef ram_ppzksnark_machine_pp<ppT> machine_ppT;
    const size_t boot_trace_size_bound = program_size + input_size;
    const bool satisfiable = true;

    const ram_ppzksnark_architecture_params<ppT> ap(w, k);
    const ram_example<machine_ppT> example = gen_ram_example_complex<machine_ppT>(ap, boot_trace_size_bound, time_bound, satisfiable);

    const bool test_serialization = true;
    const bool bit = run_ram_ppzksnark<ppT>(example, test_serialization);
    assert(bit);

    print_header("(leave) Test RAM ppzkSNARK");
}

int main()
{
    ram_ppzksnark_snark_pp<default_ram_ppzksnark_pp>::init_public_params();
    start_profiling();

    const size_t program_size = 100;
    const size_t input_size = 2;
    const size_t time_bound = 20;

    // 16-bit TinyRAM with 16 registers
    test_ram_ppzksnark<default_ram_ppzksnark_pp>(16, 16, program_size, input_size, time_bound);

    // 32-bit TinyRAM with 16 registers
    test_ram_ppzksnark<default_ram_ppzksnark_pp>(32, 16, program_size, input_size, time_bound);
}

/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <sstream>
#include "common/default_types/ram_zksnark_pp.hpp"
#include "relations/ram_computations/rams/tinyram/tinyram_params.hpp"
#include "zk_proof_systems/zksnark/ram_zksnark/examples/run_ram_zksnark.hpp"
#include "relations/ram_computations/rams/examples/ram_examples.hpp"

using namespace libsnark;

template<typename ppT>
void test_ram_zksnark(const size_t w,
                      const size_t k,
                      const size_t boot_trace_size_bound,
                      const size_t time_bound)
{
    typedef ram_zksnark_machine_pp<ppT> ramT;
    const ram_architecture_params<ramT> ap(w, k);
    const ram_example<ramT> example = gen_ram_example_complex<ramT>(ap, boot_trace_size_bound, time_bound, true);
    const bool test_serialization = true;
    const bool ans = run_ram_zksnark<ppT>(example, test_serialization);
    assert(ans);
}

int main(void)
{
    start_profiling();
    ram_zksnark_PCD_pp<default_ram_zksnark_pp>::init_public_params();

    const size_t w = 32;
    const size_t k = 16;

    const size_t boot_trace_size_bound = 20;
    const size_t time_bound = 10;

    test_ram_zksnark<default_ram_zksnark_pp>(w, k, boot_trace_size_bound, time_bound);
}

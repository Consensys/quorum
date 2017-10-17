/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include "common/default_types/r1cs_ppzksnark_pp.hpp"
#include "common/default_types/r1cs_ppzkpcd_pp.hpp"
#include "common/utils.hpp"
#include "zk_proof_systems/zksnark/ram_zksnark/examples/run_ram_zksnark.hpp"
#include "zk_proof_systems/ppzksnark/ram_ppzksnark/examples/run_ram_ppzksnark.hpp"

#include "relations/ram_computations/rams/fooram/fooram_params.hpp"

namespace libsnark {

class default_fooram_zksnark_pp {
public:
    typedef default_r1cs_ppzkpcd_pp PCD_pp;
    typedef typename PCD_pp::scalar_field_A FieldT;
    typedef ram_fooram<FieldT> machine_pp;

    static void init_public_params() { PCD_pp::init_public_params(); }
};

class default_fooram_ppzksnark_pp {
public:
    typedef default_r1cs_ppzksnark_pp snark_pp;
    typedef Fr<default_r1cs_ppzksnark_pp> FieldT;
    typedef ram_fooram<FieldT> machine_pp;

    static void init_public_params() { snark_pp::init_public_params(); }
};

} // libsnark

using namespace libsnark;

template<typename ppT>
void profile_ram_zksnark(const size_t w)
{
    typedef ram_zksnark_machine_pp<ppT> ramT;

    ram_example<ramT> example;
    example.ap = ram_architecture_params<ramT>(w);
    example.boot_trace_size_bound = 0;
    example.time_bound = 10;
    const bool test_serialization = true;
    const bool bit = run_ram_zksnark<ppT>(example, test_serialization);
    assert(bit);
}

template<typename ppT>
void profile_ram_ppzksnark(const size_t w)
{
    typedef ram_ppzksnark_machine_pp<ppT> ramT;

    ram_example<ramT> example;
    example.ap = ram_architecture_params<ramT>(w);
    example.boot_trace_size_bound = 0;
    example.time_bound = 100;
    const bool test_serialization = true;
    const bool bit = run_ram_ppzksnark<ppT>(example, test_serialization);
    assert(bit);
}

int main(int argc, const char* argv[])
{
    UNUSED(argv);
    start_profiling();
    default_fooram_ppzksnark_pp::init_public_params();
    default_fooram_zksnark_pp::init_public_params();

    if (argc == 1)
    {
        profile_ram_zksnark<default_fooram_zksnark_pp>(32);
    }
    else
    {
        profile_ram_ppzksnark<default_fooram_ppzksnark_pp>(8);
    }
}

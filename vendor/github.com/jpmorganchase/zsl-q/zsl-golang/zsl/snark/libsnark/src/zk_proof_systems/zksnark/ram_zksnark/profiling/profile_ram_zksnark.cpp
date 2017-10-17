/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include "common/default_types/ram_zksnark_pp.hpp"
#include "relations/ram_computations/memory/examples/memory_contents_examples.hpp"
#include "relations/ram_computations/rams/examples/ram_examples.hpp"
#include "zk_proof_systems/zksnark/ram_zksnark/ram_zksnark.hpp"
#include "zk_proof_systems/zksnark/ram_zksnark/examples/run_ram_zksnark.hpp"
#include "relations/ram_computations/rams/tinyram/tinyram_params.hpp"

#include <boost/program_options.hpp>

using namespace libsnark;

template<typename FieldT>
void simulate_random_memory_contents(const tinyram_architecture_params &ap, const size_t input_size, const size_t program_size)
{
    const size_t num_addresses = 1ul<<ap.dwaddr_len();
    const size_t value_size = 2 * ap.w;
    memory_contents init_random = random_memory_contents(num_addresses, value_size, program_size + (input_size + 1)/2);

    enter_block("Initialize random delegated memory");
    delegated_ra_memory<FieldT> dm_random(num_addresses, value_size, init_random);
    leave_block("Initialize random delegated memory");
}

template<typename ppT>
void profile_ram_zksnark_verifier(const tinyram_architecture_params &ap, const size_t input_size, const size_t program_size)
{
    typedef ram_zksnark_machine_pp<ppT> ramT;
    const size_t time_bound  = 10;

    const size_t boot_trace_size_bound = program_size + input_size;
    const ram_example<ramT> example = gen_ram_example_complex<ramT>(ap, boot_trace_size_bound, time_bound, true);

    ram_zksnark_proof<ppT> pi;
    ram_zksnark_verification_key<ppT> vk = ram_zksnark_verification_key<ppT>::dummy_verification_key(ap);

    enter_block("Verify fake proof");
    ram_zksnark_verifier<ppT>(vk, example.boot_trace, time_bound, pi);
    leave_block("Verify fake proof");
}

template<typename ppT>
void print_ram_zksnark_verifier_profiling()
{
    inhibit_profiling_info = true;
    for (size_t w : { 16, 32 })
    {
        const size_t k = 16;

        for (size_t input_size : { 0, 10, 100 })
        {
            for (size_t program_size = 10; program_size <= 10000; program_size *= 10)
            {
                const tinyram_architecture_params ap(w, k);

                profile_ram_zksnark_verifier<ppT>(ap, input_size, program_size);

                const double input_map = last_times["Call to ram_zksnark_verifier_input_map"];
                const double preprocessing = last_times["Call to r1cs_ppzksnark_verifier_process_vk"];
                const double accumulate = last_times["Call to r1cs_ppzksnark_IC_query::accumulate"];
                const double pairings = last_times["Online pairing computations"];
                const double total = last_times["Call to ram_zksnark_verifier"];
                const double rest = total - (input_map + preprocessing + accumulate + pairings);

                const double delegated_ra_memory_init = last_times["Construct delegated_ra_memory from memory map"];
                simulate_random_memory_contents<Fr<typename ppT::curve_A_pp> >(ap, input_size, program_size);
                const double delegated_ra_memory_init_random = last_times["Initialize random delegated memory"];
                const double input_map_random = input_map - delegated_ra_memory_init + delegated_ra_memory_init_random;
                const double total_random = total - delegated_ra_memory_init + delegated_ra_memory_init_random;

                printf("w = %zu, k = %zu, program_size = %zu, input_size = %zu, input_map = %0.2fms, preprocessing = %0.2fms, accumulate = %0.2fms, pairings = %0.2fms, rest = %0.2fms, total = %0.2fms (input_map_random = %0.2fms, total_random = %0.2fms)\n",
                       w, k, program_size, input_size, input_map * 1e-6, preprocessing * 1e-6, accumulate * 1e-6, pairings * 1e-6, rest * 1e-6, total * 1e-6, input_map_random * 1e-6, total_random * 1e-6);
            }
        }
    }
}

template<typename ppT>
void profile_ram_zksnark(const tinyram_architecture_params &ap, const size_t program_size, const size_t input_size, const size_t time_bound)
{
    typedef ram_zksnark_machine_pp<ppT> ramT;

    const size_t boot_trace_size_bound = program_size + input_size;
    const ram_example<ramT> example = gen_ram_example_complex<ramT>(ap, boot_trace_size_bound, time_bound, true);
    const bool test_serialization = true;
    const bool bit = run_ram_zksnark<ppT>(example, test_serialization);
    assert(bit);
}

namespace po = boost::program_options;

bool process_command_line(const int argc, const char** argv,
                          bool &profile_gp,
                          size_t &w,
                          size_t &k,
                          bool &profile_v,
                          size_t &l)
{
    try
    {
        po::options_description desc("Usage");
        desc.add_options()
            ("help", "print this help message")
            ("profile_gp", "profile generator and prover")
            ("w", po::value<size_t>(&w)->default_value(16), "word size")
            ("k", po::value<size_t>(&k)->default_value(16), "register count")
            ("profile_v", "profile verifier")
            ("v", "print version info")
            ("l", po::value<size_t>(&l)->default_value(10), "program length");

        po::variables_map vm;
        po::store(po::parse_command_line(argc, argv, desc), vm);

        if (vm.count("v"))
        {
            print_compilation_info();
            exit(0);
        }

        if (vm.count("help"))
        {
            std::cout << desc << "\n";
            return false;
        }

        profile_gp = vm.count("profile_gp");
        profile_v = vm.count("profile_v");

        if (!(vm.count("profile_gp") ^ vm.count("profile_v")))
        {
            std::cout << "Must choose between profiling generator/prover and profiling verifier (see --help)\n";
            return false;
        }

        po::notify(vm);
    }
    catch(std::exception& e)
    {
        std::cerr << "Error: " << e.what() << "\n";
        return false;
    }

    return true;
}

int main(int argc, const char* argv[])
{
    start_profiling();
    ram_zksnark_PCD_pp<default_ram_zksnark_pp>::init_public_params();

    bool profile_gp;
    size_t w;
    size_t k;
    bool profile_v;
    size_t l;

    if (!process_command_line(argc, argv, profile_gp, w, k, profile_v, l))
    {
        return 1;
    }

    tinyram_architecture_params ap(w, k);

    if (profile_gp)
    {
        profile_ram_zksnark<default_ram_zksnark_pp>(ap, 100, 100, 10); // w, k, l, n, T
    }

    if (profile_v)
    {
        profile_ram_zksnark_verifier<default_ram_zksnark_pp>(ap, l/2, l/2);
    }
}

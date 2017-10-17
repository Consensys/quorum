/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <fstream>
#include <iostream>
#ifndef MINDEPS
#include <boost/program_options.hpp>
#endif

#include "common/default_types/tinyram_ppzksnark_pp.hpp"
#include "zk_proof_systems/ppzksnark/ram_ppzksnark/ram_ppzksnark.hpp"
#include "relations/ram_computations/rams/tinyram/tinyram_params.hpp"

#ifndef MINDEPS
namespace po = boost::program_options;

bool process_verifier_command_line(const int argc, const char** argv,
                                   std::string &processed_assembly_fn,
                                   std::string &verification_key_fn,
                                   std::string &primary_input_fn,
                                   std::string &proof_fn,
                                   std::string &verification_result_fn)
{
    try
    {
        po::options_description desc("Usage");
        desc.add_options()
            ("help", "print this help message")
            ("processed_assembly", po::value<std::string>(&processed_assembly_fn)->required())
            ("verification_key", po::value<std::string>(&verification_key_fn)->required())
            ("primary_input", po::value<std::string>(&primary_input_fn)->required())
            ("proof", po::value<std::string>(&proof_fn)->required())
            ("verification_result", po::value<std::string>(&verification_result_fn)->required());

        po::variables_map vm;
        po::store(po::parse_command_line(argc, argv, desc), vm);

        if (vm.count("help"))
        {
            std::cout << desc << "\n";
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
#endif

using namespace libsnark;

int main(int argc, const char * argv[])
{
    default_tinyram_ppzksnark_pp::init_public_params();

#ifdef MINDEPS
    std::string processed_assembly_fn = "processed.txt";
    std::string verification_key_fn = "verification_key.txt";
    std::string proof_fn = "proof.txt";
    std::string primary_input_fn = "primary_input.txt";
    std::string verification_result_fn = "verification_result.txt";
#else
    std::string processed_assembly_fn;
    std::string verification_key_fn;
    std::string proof_fn;
    std::string primary_input_fn;
    std::string verification_result_fn;

    if (!process_verifier_command_line(argc, argv, processed_assembly_fn, verification_key_fn, primary_input_fn, proof_fn, verification_result_fn))
    {
        return 1;
    }
#endif
    start_profiling();

    ram_ppzksnark_verification_key<default_tinyram_ppzksnark_pp> vk;
    std::ifstream vk_file(verification_key_fn);
    vk_file >> vk;
    vk_file.close();

    std::ifstream processed(processed_assembly_fn);
    tinyram_program program = load_preprocessed_program(vk.ap, processed);

    std::ifstream f_primary_input(primary_input_fn);
    tinyram_input_tape primary_input = load_tape(f_primary_input);

    std::ifstream proof_file(proof_fn);
    ram_ppzksnark_proof<default_tinyram_ppzksnark_pp> pi;
    proof_file >> pi;
    proof_file.close();

    const ram_boot_trace<default_tinyram_ppzksnark_pp> boot_trace = tinyram_boot_trace_from_program_and_input(vk.ap, vk.primary_input_size_bound, program, primary_input);
    const bool bit = ram_ppzksnark_verifier<default_tinyram_ppzksnark_pp>(vk, boot_trace, pi);

    printf("================================================================================\n");
    printf("The verification result is: %s\n", (bit ? "PASS" : "FAIL"));
    printf("================================================================================\n");
    std::ofstream vr_file(verification_result_fn);
    vr_file << bit << "\n";
    vr_file.close();
}

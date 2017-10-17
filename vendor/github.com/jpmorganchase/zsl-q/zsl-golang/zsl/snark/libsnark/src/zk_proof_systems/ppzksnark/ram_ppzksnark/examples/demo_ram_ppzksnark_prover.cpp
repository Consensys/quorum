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

bool process_prover_command_line(const int argc, const char** argv,
                                 std::string &processed_assembly_fn,
                                 std::string &proving_key_fn,
                                 std::string &primary_input_fn,
                                 std::string &auxiliary_input_fn,
                                 std::string &proof_fn)
{
    try
    {
        po::options_description desc("Usage");
        desc.add_options()
            ("help", "print this help message")
            ("processed_assembly", po::value<std::string>(&processed_assembly_fn)->required())
            ("proving_key", po::value<std::string>(&proving_key_fn)->required())
            ("primary_input", po::value<std::string>(&primary_input_fn)->required())
            ("auxiliary_input", po::value<std::string>(&auxiliary_input_fn)->required())
            ("proof", po::value<std::string>(&proof_fn)->required());

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
    std::string proving_key_fn = "proving_key.txt";
    std::string primary_input_fn = "primary_input.txt";
    std::string auxiliary_input_fn = "auxiliary_input.txt";
    std::string proof_fn = "proof.txt";
#else
    std::string processed_assembly_fn;
    std::string proving_key_fn;
    std::string primary_input_fn;
    std::string auxiliary_input_fn;
    std::string proof_fn;

    if (!process_prover_command_line(argc, argv, processed_assembly_fn,
                                     proving_key_fn, primary_input_fn, auxiliary_input_fn, proof_fn))
    {
        return 1;
    }
#endif
    start_profiling();

    /* load everything */
    ram_ppzksnark_proving_key<default_tinyram_ppzksnark_pp> pk;
    std::ifstream pk_file(proving_key_fn);
    pk_file >> pk;
    pk_file.close();

    std::ifstream processed(processed_assembly_fn);
    tinyram_program program = load_preprocessed_program(pk.ap, processed);

    std::ifstream f_primary_input(primary_input_fn);
    std::ifstream f_auxiliary_input(auxiliary_input_fn);
    tinyram_input_tape primary_input = load_tape(f_primary_input);
    tinyram_input_tape auxiliary_input = load_tape(f_auxiliary_input);

    const ram_boot_trace<default_tinyram_ppzksnark_pp> boot_trace = tinyram_boot_trace_from_program_and_input(pk.ap, pk.primary_input_size_bound, program, primary_input);
    const ram_ppzksnark_proof<default_tinyram_ppzksnark_pp> proof = ram_ppzksnark_prover<default_tinyram_ppzksnark_pp>(pk, boot_trace,  auxiliary_input);

    std::ofstream proof_file(proof_fn);
    proof_file << proof;
    proof_file.close();
}

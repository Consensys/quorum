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
#ifndef MINDEPS
#include <boost/program_options.hpp>
#endif

#include "common/default_types/tinyram_ppzksnark_pp.hpp"
#include "common/profiling.hpp"
#include "reductions/ram_to_r1cs/ram_to_r1cs.hpp"
#include "relations/ram_computations/rams/tinyram/tinyram_params.hpp"
#include "zk_proof_systems/ppzksnark/ram_ppzksnark/ram_ppzksnark.hpp"

#ifndef MINDEPS
namespace po = boost::program_options;

bool process_arithm_command_line(const int argc, const char** argv,
                                 std::string &assembly_fn,
                                 std::string &processed_assembly_fn,
                                 std::string &architecture_params_fn,
                                 std::string &computation_bounds_fn,
                                 std::string &primary_input_fn,
                                 std::string &auxiliary_input_fn)
{
    try
    {
        po::options_description desc("Usage");
        desc.add_options()
            ("help", "print this help message")
            ("assembly", po::value<std::string>(&assembly_fn)->required())
            ("processed_assembly", po::value<std::string>(&processed_assembly_fn)->required())
            ("architecture_params", po::value<std::string>(&architecture_params_fn)->required())
            ("computation_bounds", po::value<std::string>(&computation_bounds_fn)->required())
            ("primary_input", po::value<std::string>(&primary_input_fn)->required())
            ("auxiliary_input", po::value<std::string>(&auxiliary_input_fn)->required());

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
    typedef Fr<default_ec_pp> FieldT;
    typedef ram_tinyram<FieldT> default_ram;

    default_ec_pp::init_public_params();

#ifdef MINDEPS
    std::string assembly_fn = "assembly.s";
    std::string processed_assembly_fn = "processed.txt";
    std::string architecture_params_fn = "architecture_params.txt";
    std::string computation_bounds_fn = "computation_bounds.txt";
    std::string primary_input_fn = "primary_input.txt";
    std::string auxiliary_input_fn = "auxiliary_input.txt";
#else
    std::string assembly_fn;
    std::string processed_assembly_fn;
    std::string architecture_params_fn;
    std::string computation_bounds_fn;
    std::string primary_input_fn;
    std::string auxiliary_input_fn;

    if (!process_arithm_command_line(argc, argv, assembly_fn, processed_assembly_fn, architecture_params_fn,
                                     computation_bounds_fn, primary_input_fn, auxiliary_input_fn))
    {
        return 1;
    }
#endif
    start_profiling();

    printf("================================================================================\n");
    printf("TinyRAM example loader\n");
    printf("================================================================================\n\n");

    /* load everything */
    ram_architecture_params<default_ram> ap;
    std::ifstream f_ap(architecture_params_fn);
    f_ap >> ap;

    printf("Will run on %zu register machine (word size = %zu)\n", ap.k, ap.w);

    std::ifstream f_rp(computation_bounds_fn);
    size_t tinyram_input_size_bound, tinyram_program_size_bound, time_bound;
    f_rp >> tinyram_input_size_bound >> tinyram_program_size_bound >> time_bound;

    std::ifstream processed(processed_assembly_fn);
    std::ifstream raw(assembly_fn);
    tinyram_program program = load_preprocessed_program(ap, processed);
    printf("Program:\n%s\n", std::string((std::istreambuf_iterator<char>(raw)),
                                         std::istreambuf_iterator<char>()).c_str());

    std::ifstream f_primary_input(primary_input_fn);
    std::ifstream f_auxiliary_input(auxiliary_input_fn);

    enter_block("Loading primary input");
    tinyram_input_tape primary_input(load_tape(f_primary_input));
    leave_block("Loading primary input");

    enter_block("Loading auxiliary input");
    tinyram_input_tape auxiliary_input = load_tape(f_auxiliary_input);
    leave_block("Loading auxiliary input");

    const size_t boot_trace_size_bound = tinyram_input_size_bound + tinyram_program_size_bound;
    const ram_boot_trace<default_ram> boot_trace = tinyram_boot_trace_from_program_and_input(ap, boot_trace_size_bound, program, primary_input);

    typedef ram_ppzksnark_machine_pp<default_tinyram_ppzksnark_pp> default_ram;

    ram_to_r1cs<default_ram> r(ap, boot_trace_size_bound, time_bound);
    r.instance_map();

    const r1cs_primary_input<FieldT> r1cs_primary_input = ram_to_r1cs<default_ram>::primary_input_map(ap, boot_trace_size_bound, boot_trace);
    const r1cs_auxiliary_input<FieldT> r1cs_auxiliary_input = r.auxiliary_input_map(boot_trace, auxiliary_input);
    const r1cs_constraint_system<FieldT> constraint_system = r.get_constraint_system();

    r.print_execution_trace();
    assert(constraint_system.is_satisfied(r1cs_primary_input, r1cs_auxiliary_input));
}

/** @file
 *****************************************************************************
 Test program that exercises the ppzkSNARK (first generator, then
 prover, then verifier) on a synthetic TBCS instance.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <cassert>
#include <cstdio>

#include "common/default_types/tbcs_ppzksnark_pp.hpp"
#include "common/profiling.hpp"
#include "relations/circuit_satisfaction_problems/tbcs/examples/tbcs_examples.hpp"
#include "zk_proof_systems/ppzksnark/tbcs_ppzksnark/examples/run_tbcs_ppzksnark.hpp"

using namespace libsnark;

template<typename ppT>
void test_tbcs_ppzksnark(const size_t primary_input_size,
                         const size_t auxiliary_input_size,
                         const size_t num_gates,
                         const size_t num_outputs)
{
    print_header("(enter) Test TBCS ppzkSNARK");

    const bool test_serialization = true;
    const tbcs_example example = generate_tbcs_example(primary_input_size, auxiliary_input_size, num_gates, num_outputs);
#ifdef DEBUG
    example.circuit.print();
#endif
    const bool bit = run_tbcs_ppzksnark<ppT>(example, test_serialization);
    assert(bit);

    print_header("(leave) Test TBCS ppzkSNARK");
}

int main()
{
    default_tbcs_ppzksnark_pp::init_public_params();
    start_profiling();

    test_tbcs_ppzksnark<default_tbcs_ppzksnark_pp>(10, 10, 20, 5);
}

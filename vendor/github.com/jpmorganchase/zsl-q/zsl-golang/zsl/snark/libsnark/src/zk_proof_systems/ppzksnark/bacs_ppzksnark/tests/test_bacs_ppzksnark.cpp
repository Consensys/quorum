/** @file
 *****************************************************************************
 Test program that exercises the ppzkSNARK (first generator, then
 prover, then verifier) on a synthetic BACS instance.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <cassert>
#include <cstdio>

#include "common/default_types/bacs_ppzksnark_pp.hpp"
#include "common/profiling.hpp"
#include "relations/circuit_satisfaction_problems/bacs/examples/bacs_examples.hpp"
#include "zk_proof_systems/ppzksnark/bacs_ppzksnark/examples/run_bacs_ppzksnark.hpp"

using namespace libsnark;

template<typename ppT>
void test_bacs_ppzksnark(const size_t primary_input_size,
                         const size_t auxiliary_input_size,
                         const size_t num_gates,
                         const size_t num_outputs)
{
    print_header("(enter) Test BACS ppzkSNARK");

    const bool test_serialization = true;
    const bacs_example<Fr<ppT> > example = generate_bacs_example<Fr<ppT> >(primary_input_size, auxiliary_input_size, num_gates, num_outputs);
#ifdef DEBUG
    example.circuit.print();
#endif
    const bool bit = run_bacs_ppzksnark<ppT>(example, test_serialization);
    assert(bit);

    print_header("(leave) Test BACS ppzkSNARK");
}

int main()
{
    default_bacs_ppzksnark_pp::init_public_params();
    start_profiling();

    test_bacs_ppzksnark<default_bacs_ppzksnark_pp>(10, 10, 20, 5);
}

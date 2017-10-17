/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include <algorithm>
#include <cassert>
#include <cstdio>
#include <cstring>
#include <vector>

#include "algebra/curves/mnt/mnt6/mnt6_pp.hpp"
#include "algebra/fields/field_utils.hpp"
#include "common/profiling.hpp"
#include "common/utils.hpp"
#include "reductions/uscs_to_ssp/uscs_to_ssp.hpp"
#include "relations/constraint_satisfaction_problems/uscs/examples/uscs_examples.hpp"

using namespace libsnark;

template<typename FieldT>
void test_ssp(const size_t num_constraints, const size_t num_inputs, const bool binary_input)
{
    enter_block("Call to test_ssp");

    print_indent(); printf("* Number of constraints: %zu\n", num_constraints);
    print_indent(); printf("* Number of inputs: %zu\n", num_inputs);
    print_indent(); printf("* Input type: %s\n", binary_input ? "binary" : "field");

    enter_block("Generate constraint system and assignment");
    uscs_example<FieldT> example;
    if (binary_input)
    {
        example = generate_uscs_example_with_binary_input<FieldT>(num_constraints, num_inputs);
    }
    else
    {
        example = generate_uscs_example_with_field_input<FieldT>(num_constraints, num_inputs);
    }
    leave_block("Generate constraint system and assignment");

    enter_block("Check satisfiability of constraint system");
    assert(example.constraint_system.is_satisfied(example.primary_input, example.auxiliary_input));
    leave_block("Check satisfiability of constraint system");

    const FieldT t = FieldT::random_element(),
                 d = FieldT::random_element();

    enter_block("Compute SSP instance 1");
    ssp_instance<FieldT> ssp_inst_1 = uscs_to_ssp_instance_map(example.constraint_system);
    leave_block("Compute SSP instance 1");

    enter_block("Compute SSP instance 2");
    ssp_instance_evaluation<FieldT> ssp_inst_2 = uscs_to_ssp_instance_map_with_evaluation(example.constraint_system, t);
    leave_block("Compute SSP instance 2");

    enter_block("Compute SSP witness");
    ssp_witness<FieldT> ssp_wit = uscs_to_ssp_witness_map(example.constraint_system, example.primary_input, example.auxiliary_input, d);
    leave_block("Compute SSP witness");

    enter_block("Check satisfiability of SSP instance 1");
    assert(ssp_inst_1.is_satisfied(ssp_wit));
    leave_block("Check satisfiability of SSP instance 1");

    enter_block("Check satisfiability of SSP instance 2");
    assert(ssp_inst_2.is_satisfied(ssp_wit));
    leave_block("Check satisfiability of SSP instance 2");

    leave_block("Call to test_ssp");
}

int main()
{
    start_profiling();

    mnt6_pp::init_public_params();

    const size_t num_inputs = 10;

    const size_t basic_domain_size = 1ul<<mnt6_Fr::s;
    const size_t step_domain_size = (1ul<<10) + (1ul<<8);
    const size_t extended_domain_size = 1ul<<(mnt6_Fr::s+1);
    const size_t extended_domain_size_special = extended_domain_size-1;

    enter_block("Test SSP for binary inputs");

    test_ssp<Fr<mnt6_pp> >(basic_domain_size, num_inputs, true);
    test_ssp<Fr<mnt6_pp> >(step_domain_size, num_inputs, true);
    test_ssp<Fr<mnt6_pp> >(extended_domain_size, num_inputs, true);
    test_ssp<Fr<mnt6_pp> >(extended_domain_size_special, num_inputs, true);

    leave_block("Test SSP for binary inputs");

    enter_block("Test SSP for field inputs");

    test_ssp<Fr<mnt6_pp> >(basic_domain_size, num_inputs, false);
    test_ssp<Fr<mnt6_pp> >(step_domain_size, num_inputs, false);
    test_ssp<Fr<mnt6_pp> >(extended_domain_size, num_inputs, false);
    test_ssp<Fr<mnt6_pp> >(extended_domain_size_special, num_inputs, false);

    leave_block("Test SSP for field inputs");
}

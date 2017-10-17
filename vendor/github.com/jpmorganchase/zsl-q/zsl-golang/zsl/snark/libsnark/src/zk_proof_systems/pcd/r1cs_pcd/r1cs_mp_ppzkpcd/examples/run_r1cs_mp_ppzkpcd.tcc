/** @file
 *****************************************************************************

 Implementation of functionality that runs the R1CS multi-predicate ppzkPCD
 for a compliance predicate example.

 See run_r1cs_mp_ppzkpcd.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_R1CS_MP_PPZKPCD_TCC_
#define RUN_R1CS_MP_PPZKPCD_TCC_

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_mp_ppzkpcd/r1cs_mp_ppzkpcd.hpp"

#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/examples/tally_cp.hpp"

namespace libsnark {

template<typename PCD_ppT>
bool run_r1cs_mp_ppzkpcd_tally_example(const size_t wordsize,
                                       const size_t max_arity,
                                       const size_t depth,
                                       const bool test_serialization,
                                       const bool test_multi_type,
                                       const bool test_same_type_optimization)
{
    enter_block("Call to run_r1cs_mp_ppzkpcd_tally_example");

    typedef Fr<typename PCD_ppT::curve_A_pp> FieldT;

    bool all_accept = true;

    enter_block("Generate all messages");
    size_t tree_size = 0;
    size_t nodes_in_layer = 1;
    for (size_t layer = 0; layer <= depth; ++layer)
    {
        tree_size += nodes_in_layer;
        nodes_in_layer *= max_arity;
    }

    std::vector<size_t> tree_types(tree_size);
    std::vector<size_t> tree_elems(tree_size);
    std::vector<size_t> tree_arity(tree_size);

    nodes_in_layer = 1;
    size_t node_idx = 0;
    for (size_t layer = 0; layer <= depth; ++layer)
    {
        for (size_t id_in_layer = 0; id_in_layer < nodes_in_layer; ++id_in_layer, ++node_idx)
        {
            if (!test_multi_type)
            {
                tree_types[node_idx] = 1;
            }
            else
            {
                if (test_same_type_optimization)
                {
                    tree_types[node_idx] = 1 + ((depth-layer) & 1);
                }
                else
                {
                    tree_types[node_idx] = 1 + (std::rand() % 2);
                }
            }

            tree_elems[node_idx] = std::rand() % 100;
            tree_arity[node_idx] = 1 + (std::rand() % max_arity); /* we will just skip below this threshold */
            printf("tree_types[%zu] = %zu\n", node_idx, tree_types[node_idx]);
            printf("tree_elems[%zu] = %zu\n", node_idx, tree_elems[node_idx]);
            printf("tree_arity[%zu] = %zu\n", node_idx, tree_arity[node_idx]);

        }
        nodes_in_layer *= max_arity;
    }

    leave_block("Generate all messages");

    std::vector<r1cs_mp_ppzkpcd_proof<PCD_ppT> > tree_proofs(tree_size);
    std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > tree_messages(tree_size);

    enter_block("Generate compliance predicates");
    std::set<size_t> tally_1_accepted_types, tally_2_accepted_types;
    if (test_same_type_optimization)
    {
        if (!test_multi_type)
        {
            /* only tally 1 is going to be used */
            tally_1_accepted_types.insert(1);
        }
        else
        {
            tally_1_accepted_types.insert(2);
            tally_2_accepted_types.insert(1);
        }
    }

    tally_cp_handler<FieldT> tally_1(1, max_arity, wordsize, test_same_type_optimization, tally_1_accepted_types);
    tally_cp_handler<FieldT> tally_2(2, max_arity, wordsize, test_same_type_optimization, tally_2_accepted_types);
    tally_1.generate_r1cs_constraints();
    tally_2.generate_r1cs_constraints();
    r1cs_pcd_compliance_predicate<FieldT> cp_1 = tally_1.get_compliance_predicate();
    r1cs_pcd_compliance_predicate<FieldT> cp_2 = tally_2.get_compliance_predicate();
    leave_block("Generate compliance predicates");

    print_header("R1CS ppzkPCD Generator");
    r1cs_mp_ppzkpcd_keypair<PCD_ppT> keypair = r1cs_mp_ppzkpcd_generator<PCD_ppT>({ cp_1, cp_2 });

    print_header("Process verification key");
    r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> pvk = r1cs_mp_ppzkpcd_process_vk<PCD_ppT>(keypair.vk);

    if (test_serialization)
    {
        enter_block("Test serialization of keys");
        keypair.pk = reserialize<r1cs_mp_ppzkpcd_proving_key<PCD_ppT> >(keypair.pk);
        keypair.vk = reserialize<r1cs_mp_ppzkpcd_verification_key<PCD_ppT> >(keypair.vk);
        pvk = reserialize<r1cs_mp_ppzkpcd_processed_verification_key<PCD_ppT> >(pvk);
        leave_block("Test serialization of keys");
    }

    std::shared_ptr<r1cs_pcd_message<FieldT> > base_msg = tally_1.get_base_case_message(); /* we choose the base to always be tally_1 */
    nodes_in_layer /= max_arity;
    for (long layer = depth; layer >= 0; --layer, nodes_in_layer /= max_arity)
    {
        for (size_t i = 0; i < nodes_in_layer; ++i)
        {
            const size_t cur_idx = (nodes_in_layer - 1) / (max_arity - 1) + i;

            tally_cp_handler<FieldT> &cur_tally = (tree_types[cur_idx] == 1 ? tally_1 : tally_2);
            r1cs_pcd_compliance_predicate<FieldT> &cur_cp = (tree_types[cur_idx] == 1 ? cp_1 : cp_2);

            const bool base_case = (max_arity * cur_idx + max_arity >= tree_size);

            std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > msgs(max_arity, base_msg);
            std::vector<r1cs_mp_ppzkpcd_proof<PCD_ppT> > proofs(max_arity);

            if (!base_case)
            {
                for (size_t i = 0; i < max_arity; ++i)
                {
                    msgs[i] = tree_messages[max_arity*cur_idx + i + 1];
                    proofs[i] = tree_proofs[max_arity*cur_idx + i + 1];
                }
            }
            msgs.resize(tree_arity[i]);
            proofs.resize(tree_arity[i]);

            std::shared_ptr<r1cs_pcd_local_data<FieldT> > ld;
            ld.reset(new tally_pcd_local_data<FieldT>(tree_elems[cur_idx]));
            cur_tally.generate_r1cs_witness(msgs, ld);

            const r1cs_pcd_compliance_predicate_primary_input<FieldT> tally_primary_input(cur_tally.get_outgoing_message());
            const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> tally_auxiliary_input(msgs, ld, cur_tally.get_witness());

            print_header("R1CS ppzkPCD Prover");
            r1cs_mp_ppzkpcd_proof<PCD_ppT> proof = r1cs_mp_ppzkpcd_prover<PCD_ppT>(keypair.pk, cur_cp.name, tally_primary_input, tally_auxiliary_input, proofs);

            if (test_serialization)
            {
                enter_block("Test serialization of proof");
                proof = reserialize<r1cs_mp_ppzkpcd_proof<PCD_ppT> >(proof);
                leave_block("Test serialization of proof");
            }

            tree_proofs[cur_idx] = proof;
            tree_messages[cur_idx] = cur_tally.get_outgoing_message();

            print_header("R1CS ppzkPCD Verifier");
            const r1cs_mp_ppzkpcd_primary_input<PCD_ppT> pcd_verifier_input(tree_messages[cur_idx]);
            const bool ans = r1cs_mp_ppzkpcd_verifier<PCD_ppT>(keypair.vk, pcd_verifier_input, tree_proofs[cur_idx]);

            print_header("R1CS ppzkPCD Online Verifier");
            const bool ans2 = r1cs_mp_ppzkpcd_online_verifier<PCD_ppT>(pvk, pcd_verifier_input, tree_proofs[cur_idx]);
            assert(ans == ans2);

            all_accept = all_accept && ans;

            printf("\n");
            for (size_t i = 0; i < msgs.size(); ++i)
            {
                printf("Message %zu was:\n", i);
                msgs[i]->print();
            }
            printf("Summand at this node:\n%zu\n", tree_elems[cur_idx]);
            printf("Outgoing message is:\n");
            tree_messages[cur_idx]->print();
            printf("\n");
            printf("Current node = %zu. Current proof verifies = %s\n", cur_idx, ans ? "YES" : "NO");
            printf("\n\n\n ================================================================================\n\n\n");
        }
    }

    leave_block("Call to run_r1cs_mp_ppzkpcd_tally_example");

    return all_accept;
}

} // libsnark

#endif // RUN_R1CS_MP_PPZKPCD_TCC_

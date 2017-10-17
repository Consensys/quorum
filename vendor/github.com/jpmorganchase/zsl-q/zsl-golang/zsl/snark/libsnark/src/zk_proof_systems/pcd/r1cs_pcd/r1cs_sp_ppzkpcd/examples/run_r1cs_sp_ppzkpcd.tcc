/** @file
 *****************************************************************************

 Implementation of functionality that runs the R1CS single-predicate ppzkPCD
 for a compliance predicate example.

 See run_r1cs_sp_ppzkpcd.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RUN_R1CS_SP_PPZKPCD_TCC_
#define RUN_R1CS_SP_PPZKPCD_TCC_

#include "zk_proof_systems/pcd/r1cs_pcd/r1cs_sp_ppzkpcd/r1cs_sp_ppzkpcd.hpp"

#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/examples/tally_cp.hpp"

namespace libsnark {

template<typename PCD_ppT>
bool run_r1cs_sp_ppzkpcd_tally_example(const size_t wordsize,
                                       const size_t arity,
                                       const size_t depth,
                                       const bool test_serialization)
{
    enter_block("Call to run_r1cs_sp_ppzkpcd_tally_example");

    typedef Fr<typename PCD_ppT::curve_A_pp> FieldT;

    bool all_accept = true;

    enter_block("Generate all messages");
    size_t tree_size = 0;
    size_t nodes_in_layer = 1;
    for (size_t layer = 0; layer <= depth; ++layer)
    {
        tree_size += nodes_in_layer;
        nodes_in_layer *= arity;
    }
    std::vector<size_t> tree_elems(tree_size);
    for (size_t i = 0; i < tree_size; ++i)
    {
        tree_elems[i] = std::rand() % 10;
        printf("tree_elems[%zu] = %zu\n", i, tree_elems[i]);
    }
    leave_block("Generate all messages");

    std::vector<r1cs_sp_ppzkpcd_proof<PCD_ppT> > tree_proofs(tree_size);
    std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > tree_messages(tree_size);

    enter_block("Generate compliance predicate");
    const size_t type = 1;
    tally_cp_handler<FieldT> tally(type, arity, wordsize);
    tally.generate_r1cs_constraints();
    r1cs_pcd_compliance_predicate<FieldT> tally_cp = tally.get_compliance_predicate();
    leave_block("Generate compliance predicate");

    print_header("R1CS ppzkPCD Generator");
    r1cs_sp_ppzkpcd_keypair<PCD_ppT> keypair = r1cs_sp_ppzkpcd_generator<PCD_ppT>(tally_cp);

    print_header("Process verification key");
    r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> pvk = r1cs_sp_ppzkpcd_process_vk<PCD_ppT>(keypair.vk);

    if (test_serialization)
    {
        enter_block("Test serialization of keys");
        keypair.pk = reserialize<r1cs_sp_ppzkpcd_proving_key<PCD_ppT> >(keypair.pk);
        keypair.vk = reserialize<r1cs_sp_ppzkpcd_verification_key<PCD_ppT> >(keypair.vk);
        pvk = reserialize<r1cs_sp_ppzkpcd_processed_verification_key<PCD_ppT> >(pvk);
        leave_block("Test serialization of keys");
    }

    std::shared_ptr<r1cs_pcd_message<FieldT> > base_msg = tally.get_base_case_message();
    nodes_in_layer /= arity;
    for (long layer = depth; layer >= 0; --layer, nodes_in_layer /= arity)
    {
        for (size_t i = 0; i < nodes_in_layer; ++i)
        {
            const size_t cur_idx = (nodes_in_layer - 1) / (arity - 1) + i;

            std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > msgs(arity, base_msg);
            std::vector<r1cs_sp_ppzkpcd_proof<PCD_ppT> > proofs(arity);

            const bool base_case = (arity * cur_idx + arity >= tree_size);

            if (!base_case)
            {
                for (size_t i = 0; i < arity; ++i)
                {
                    msgs[i] = tree_messages[arity*cur_idx + i + 1];
                    proofs[i] = tree_proofs[arity*cur_idx + i + 1];
                }
            }

            std::shared_ptr<r1cs_pcd_local_data<FieldT> > ld;
            ld.reset(new tally_pcd_local_data<FieldT>(tree_elems[cur_idx]));
            tally.generate_r1cs_witness(msgs, ld);

            const r1cs_pcd_compliance_predicate_primary_input<FieldT> tally_primary_input(tally.get_outgoing_message());
            const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> tally_auxiliary_input(msgs, ld, tally.get_witness());

            print_header("R1CS ppzkPCD Prover");
            r1cs_sp_ppzkpcd_proof<PCD_ppT> proof = r1cs_sp_ppzkpcd_prover<PCD_ppT>(keypair.pk, tally_primary_input, tally_auxiliary_input, proofs);

            if (test_serialization)
            {
                enter_block("Test serialization of proof");
                proof = reserialize<r1cs_sp_ppzkpcd_proof<PCD_ppT> >(proof);
                leave_block("Test serialization of proof");
            }

            tree_proofs[cur_idx] = proof;
            tree_messages[cur_idx] = tally.get_outgoing_message();

            print_header("R1CS ppzkPCD Verifier");
            const r1cs_sp_ppzkpcd_primary_input<PCD_ppT> pcd_verifier_input(tree_messages[cur_idx]);
            const bool ans = r1cs_sp_ppzkpcd_verifier<PCD_ppT>(keypair.vk, pcd_verifier_input, tree_proofs[cur_idx]);

            print_header("R1CS ppzkPCD Online Verifier");
            const bool ans2 = r1cs_sp_ppzkpcd_online_verifier<PCD_ppT>(pvk, pcd_verifier_input, tree_proofs[cur_idx]);
            assert(ans == ans2);

            all_accept = all_accept && ans;

            printf("\n");
            for (size_t i = 0; i < arity; ++i)
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

    leave_block("Call to run_r1cs_sp_ppzkpcd_tally_example");

    return all_accept;
}

} // libsnark

#endif // RUN_R1CS_SP_PPZKPCD_TCC_

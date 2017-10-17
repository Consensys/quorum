/** @file
 *****************************************************************************

 Implementation of interfaces for ram_universal_gadget.

 See ram_universal_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_UNIVERSAL_GADGET_TCC_
#define RAM_UNIVERSAL_GADGET_TCC_

#include "common/data_structures/integer_permutation.hpp"
#include "common/profiling.hpp"
#include "common/utils.hpp"
#include "algebra/fields/field_utils.hpp"
#include "relations/ram_computations/memory/ra_memory.hpp"

namespace libsnark {

template<typename ramT>
ram_universal_gadget<ramT>::ram_universal_gadget(ram_protoboard<ramT> &pb,
                                                 const size_t boot_trace_size_bound,
                                                 const size_t time_bound,
                                                 const pb_variable_array<FieldT> &packed_input,
                                                 const std::string &annotation_prefix) :
    ram_gadget_base<ramT>(pb, annotation_prefix),
    boot_trace_size_bound(boot_trace_size_bound),
    time_bound(time_bound),
    packed_input(packed_input)
{
    num_memory_lines = boot_trace_size_bound + (time_bound + 1) + time_bound; /* boot lines, (time_bound + 1) execution lines (including initial) and time_bound load instruction lines */
    const size_t timestamp_size = log2(num_memory_lines);

    /* allocate all lines on the execution side of the routing network */
    enter_block("Allocate initial state line");
    execution_lines.reserve(1 + time_bound);
    execution_lines.emplace_back(execution_line_variable_gadget<ramT>(pb, timestamp_size, pb.ap, FMT(annotation_prefix, " execution_lines_%zu", 0)));
    unrouted_memory_lines.emplace_back(&execution_lines[0]);
    leave_block("Allocate initial state line");

    enter_block("Allocate boot lines");
    boot_lines.reserve(boot_trace_size_bound);
    for (size_t i = 0; i < boot_trace_size_bound; ++i)
    {
        boot_lines.emplace_back(memory_line_variable_gadget<ramT>(pb, timestamp_size, pb.ap, FMT(annotation_prefix, " boot_lines_%zu", i)));
        unrouted_memory_lines.emplace_back(&boot_lines[i]);
    }
    leave_block("Allocate boot lines");

    enter_block("Allocate instruction fetch and execution lines");
    load_instruction_lines.reserve(time_bound+1); /* the last line is NOT a memory line, but here just for uniform coding (i.e. the (unusued) result of next PC) */
    for (size_t i = 0; i < time_bound; ++i)
    {
        load_instruction_lines.emplace_back(memory_line_variable_gadget<ramT>(pb, timestamp_size, pb.ap, FMT(annotation_prefix, " load_instruction_lines_%zu", i)));
        unrouted_memory_lines.emplace_back(&load_instruction_lines[i]);

        execution_lines.emplace_back(execution_line_variable_gadget<ramT>(pb, timestamp_size, pb.ap, FMT(annotation_prefix, " execution_lines_%zu", i+1)));
        unrouted_memory_lines.emplace_back(&execution_lines[i+1]);
    }
    load_instruction_lines.emplace_back(memory_line_variable_gadget<ramT>(pb, timestamp_size, pb.ap, FMT(annotation_prefix, " load_instruction_lines_%zu", time_bound)));
    leave_block("Allocate instruction fetch and execution lines");

    /* deal with packing of the input */
    enter_block("Pack input");
    const size_t line_size_bits = pb.ap.address_size() + pb.ap.value_size();
    const size_t max_chunk_size = FieldT::capacity();
    const size_t packed_line_size = div_ceil(line_size_bits, max_chunk_size);
    assert(packed_input.size() == packed_line_size * boot_trace_size_bound);

    auto input_it = packed_input.begin();
    for (size_t i = 0; i < boot_trace_size_bound; ++i)
    {
        /* note the reversed order */
        pb_variable_array<FieldT> boot_line_bits;
        boot_line_bits.insert(boot_line_bits.end(), boot_lines[boot_trace_size_bound-1-i].address->bits.begin(), boot_lines[boot_trace_size_bound-1-i].address->bits.end());
        boot_line_bits.insert(boot_line_bits.end(), boot_lines[boot_trace_size_bound-1-i].contents_after->bits.begin(), boot_lines[boot_trace_size_bound-1-i].contents_after->bits.end());

        pb_variable_array<FieldT> packed_boot_line = pb_variable_array<FieldT>(input_it, input_it + packed_line_size);
        std::advance(input_it, packed_line_size);

        unpack_boot_lines.emplace_back(multipacking_gadget<FieldT>(pb, boot_line_bits, packed_boot_line, max_chunk_size, FMT(annotation_prefix, " unpack_boot_lines_%zu", i)));
    }
    leave_block("Pack input");

    /* deal with routing */
    enter_block("Allocate routed memory lines");
    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        routed_memory_lines.emplace_back(memory_line_variable_gadget<ramT>(pb, timestamp_size, pb.ap, FMT(annotation_prefix, " routed_memory_lines_%zu", i)));
    }
    leave_block("Allocate routed memory lines");

    enter_block("Collect inputs/outputs for the routing network");
    routing_inputs.reserve(num_memory_lines);
    routing_outputs.reserve(num_memory_lines);

    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        routing_inputs.emplace_back(unrouted_memory_lines[i]->all_vars());
        routing_outputs.emplace_back(routed_memory_lines[i].all_vars());
    }
    leave_block("Collect inputs/outputs for the routing network");

    enter_block("Allocate routing network");
    routing_network.reset(new as_waksman_routing_gadget<FieldT>(pb, num_memory_lines, routing_inputs, routing_outputs, FMT(this->annotation_prefix, " routing_network")));
    leave_block("Allocate routing network");

    /* deal with all checkers */
    enter_block("Allocate execution checkers");
    execution_checkers.reserve(time_bound);
    for (size_t i = 0; i < time_bound; ++i)
    {
        execution_checkers.emplace_back(ram_cpu_checker<ramT>(pb,
                                                              load_instruction_lines[i].address->bits, // prev_pc_addr
                                                              load_instruction_lines[i].contents_after->bits, // prev_pc_val
                                                              execution_lines[i].cpu_state, // prev_state
                                                              execution_lines[i+1].address->bits, // ls_addr,
                                                              execution_lines[i+1].contents_before->bits, // ls_prev_val
                                                              execution_lines[i+1].contents_after->bits, // ls_next_val
                                                              execution_lines[i+1].cpu_state, // next_state
                                                              load_instruction_lines[i+1].address->bits, // next_pc_addr
                                                              execution_lines[i+1].has_accepted, // next_has_accepted
                                                              FMT(annotation_prefix, " execution_checkers_%zu", i)));
    }
    leave_block("Allocate execution checkers");

    enter_block("Allocate all memory checkers");
    memory_checkers.reserve(num_memory_lines);
    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        memory_checkers.emplace_back(memory_checker_gadget<ramT>(pb,
                                                                 timestamp_size,
                                                                 *unrouted_memory_lines[i],
                                                                 routed_memory_lines[i],
                                                                 FMT(this->annotation_prefix, " memory_checkers_%zu", i)));
    }
    leave_block("Allocate all memory checkers");

    /* done */
}

template<typename ramT>
void ram_universal_gadget<ramT>::generate_r1cs_constraints()
{
    enter_block("Call to generate_r1cs_constraints of ram_universal_gadget");
    for (size_t i = 0; i < boot_trace_size_bound; ++i)
    {
        unpack_boot_lines[i].generate_r1cs_constraints(false);
    }

    /* ensure that we start with all zeros state */
    for (size_t i = 0; i < this->pb.ap.cpu_state_size(); ++i)
    {
        generate_r1cs_equals_const_constraint<FieldT>(this->pb, execution_lines[0].cpu_state[i], FieldT::zero());
    }

    /* ensure increasing timestamps */
    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        generate_r1cs_equals_const_constraint<FieldT>(this->pb, unrouted_memory_lines[i]->timestamp->packed, FieldT(i));
    }

    /* ensure bitness of trace lines on the time side */
    for (size_t i = 0; i < boot_trace_size_bound; ++i)
    {
        boot_lines[i].generate_r1cs_constraints(true);
    }

    execution_lines[0].generate_r1cs_constraints(true);
    for (size_t i = 0; i < time_bound; ++i)
    {
        load_instruction_lines[i].generate_r1cs_constraints(true);
        execution_lines[i+1].generate_r1cs_constraints(true);
    }

    /* ensure bitness of trace lines on the memory side */
    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        routed_memory_lines[i].generate_r1cs_constraints();
    }

    /* ensure that load instruction lines really do loads */
    for (size_t i = 0; i < time_bound; ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, load_instruction_lines[i].contents_before->packed,
                                                             load_instruction_lines[i].contents_after->packed),
                                     FMT(this->annotation_prefix, " load_instruction_%zu_is_a_load", i));
    }

    /* ensure correct execution */
    for (size_t i = 0; i < time_bound; ++i)
    {
        execution_checkers[i].generate_r1cs_constraints();
    }

    /* check memory */
    routing_network->generate_r1cs_constraints();

    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        memory_checkers[i].generate_r1cs_constraints();
    }

    /* ensure that PC started at the prescribed value */
    generate_r1cs_equals_const_constraint<FieldT>(this->pb, load_instruction_lines[0].address->packed, FieldT(this->pb.ap.initial_pc_addr()));

    /* ensure that the last state was an accepting one */
    generate_r1cs_equals_const_constraint<FieldT>(this->pb, execution_lines[time_bound].has_accepted, FieldT::one(), "last_state_must_be_accepting");

    /* print constraint profiling */
    const size_t num_constraints = this->pb.num_constraints();
    const size_t num_variables = this->pb.num_variables();

    if (!inhibit_profiling_info)
    {
        print_indent(); printf("* Number of constraints: %zu\n", num_constraints);
        print_indent(); printf("* Number of constraints / cycle: %0.1f\n", 1.*num_constraints/this->time_bound);

        print_indent(); printf("* Number of variables: %zu\n", num_variables);
        print_indent(); printf("* Number of variables / cycle: %0.1f\n", 1.*num_variables/this->time_bound);
    }
    leave_block("Call to generate_r1cs_constraints of ram_universal_gadget");
}

template<typename ramT>
void ram_universal_gadget<ramT>::generate_r1cs_witness(const ram_boot_trace<ramT> &boot_trace,
                                                       const ram_input_tape<ramT> &auxiliary_input)
{
    /* assign correct timestamps to all lines */
    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        this->pb.val(unrouted_memory_lines[i]->timestamp->packed) = FieldT(i);
        unrouted_memory_lines[i]->timestamp->generate_r1cs_witness_from_packed();
    }

    /* fill in the initial state */
    const ram_cpu_state<ramT> initial_state = this->pb.ap.initial_cpu_state();
    execution_lines[0].cpu_state.fill_with_bits(this->pb, initial_state);

    /* fill in the boot section */
    memory_contents memory_after_boot;

    for (auto it : boot_trace.get_all_trace_entries())
    {
        const size_t boot_pos = it.first;
        assert(boot_pos < boot_trace_size_bound);
        const size_t address = it.second.first;
        const size_t contents = it.second.second;

        this->pb.val(boot_lines[boot_pos].address->packed) = FieldT(address, true);
        this->pb.val(boot_lines[boot_pos].contents_after->packed) = FieldT(contents, true);
        boot_lines[boot_pos].generate_r1cs_witness_from_packed();

        memory_after_boot[address] = contents;
    }

    /* do the actual execution */
    ra_memory mem_backend(1ul<<(this->pb.ap.address_size()), this->pb.ap.value_size(), memory_after_boot);
    typename ram_input_tape<ramT>::const_iterator auxiliary_input_it = auxiliary_input.begin();

    this->pb.val(load_instruction_lines[0].address->packed) = FieldT(this->pb.ap.initial_pc_addr(), true);
    load_instruction_lines[0].address->generate_r1cs_witness_from_packed();

    for (size_t i = 0; i < time_bound; ++i)
    {
        /* load instruction */
        const size_t pc_addr = this->pb.val(load_instruction_lines[i].address->packed).as_ulong();
        const size_t pc_val = mem_backend.get_value(pc_addr);

        this->pb.val(load_instruction_lines[i].contents_before->packed) = FieldT(pc_val, true);
        this->pb.val(load_instruction_lines[i].contents_after->packed) = FieldT(pc_val, true);
        load_instruction_lines[i].generate_r1cs_witness_from_packed();

        /* first fetch the address part of the memory */
        execution_checkers[i].generate_r1cs_witness_address();
        execution_lines[i+1].address->generate_r1cs_witness_from_bits();

        /* fill it in */
        const size_t load_store_addr = this->pb.val(execution_lines[i+1].address->packed).as_ulong();
        const size_t load_store_prev_val = mem_backend.get_value(load_store_addr);

        this->pb.val(execution_lines[i+1].contents_before->packed) = FieldT(load_store_prev_val, true);
        execution_lines[i+1].contents_before->generate_r1cs_witness_from_packed();

        /* then execute the rest of the instruction */
        execution_checkers[i].generate_r1cs_witness_other(auxiliary_input_it, auxiliary_input.end());

        /* update the memory possibly changed by the CPU checker */
        execution_lines[i+1].contents_after->generate_r1cs_witness_from_bits();
        const size_t load_store_next_val = this->pb.val(execution_lines[i+1].contents_after->packed).as_ulong();
        mem_backend.set_value(load_store_addr, load_store_next_val);

        /* the next PC address was passed in a bit form, so maintain packed form as well */
        load_instruction_lines[i+1].address->generate_r1cs_witness_from_bits();
    }

    /*
      Get the correct memory permutation.

      We sort all memory accesses by address breaking ties by
      timestamp. In our routing configuration we pair each memory
      access with subsequent access in this ordering.

      That way num_memory_pairs of memory checkers will do a full
      cycle over all memory accesses, enforced by the proper ordering
      property.
    */

    typedef std::pair<size_t, size_t> mem_pair; /* a pair of address, timestamp */
    std::vector<mem_pair> mem_pairs;

    for (size_t i = 0; i < this->num_memory_lines; ++i)
    {
        mem_pairs.emplace_back(std::make_pair(this->pb.val(unrouted_memory_lines[i]->address->packed).as_ulong(),
                                              this->pb.val(unrouted_memory_lines[i]->timestamp->packed).as_ulong()));
    }

    std::sort(mem_pairs.begin(), mem_pairs.end());

    integer_permutation pi(this->num_memory_lines);
    for (size_t i = 0; i < this->num_memory_lines; ++i)
    {
        const size_t timestamp = this->pb.val(unrouted_memory_lines[i]->timestamp->packed).as_ulong();
        const size_t address = this->pb.val(unrouted_memory_lines[i]->address->packed).as_ulong();

        const auto it = std::upper_bound(mem_pairs.begin(), mem_pairs.end(), std::make_pair(address, timestamp));
        const size_t prev = (it == mem_pairs.end() ? 0 : it->second);
        pi.set(prev, i);
    }

    /* route according to the memory permutation */
    routing_network->generate_r1cs_witness(pi);

    for (size_t i = 0; i < this->num_memory_lines; ++i)
    {
        routed_memory_lines[i].generate_r1cs_witness_from_bits();
    }

    /* generate witness for memory checkers */
    for (size_t i = 0; i < this->num_memory_lines; ++i)
    {
        memory_checkers[i].generate_r1cs_witness();
    }

    /* repack back the input */
    for (size_t i = 0; i < boot_trace_size_bound; ++i)
    {
        unpack_boot_lines[i].generate_r1cs_witness_from_bits();
    }

    /* print debugging information */
    if (!inhibit_profiling_info)
    {
        print_indent();
        printf("* Protoboard satisfied: %s\n", (this->pb.is_satisfied() ? "YES" : "no"));
    }
}

template<typename ramT>
void ram_universal_gadget<ramT>::print_execution_trace() const
{
    for (size_t i = 0; i < boot_trace_size_bound; ++i)
    {
        printf("Boot process at t=#%zu: store %zu at %zu\n",
               i,
               this->pb.val(boot_lines[i].contents_after->packed).as_ulong(),
               this->pb.val(boot_lines[i].address->packed).as_ulong());
    }

    for (size_t i = 0; i < time_bound; ++i)
    {
        printf("Execution step %zu:\n", i);
        printf("  Loaded instruction %zu from address %zu (ts = %zu)\n",
               this->pb.val(load_instruction_lines[i].contents_after->packed).as_ulong(),
               this->pb.val(load_instruction_lines[i].address->packed).as_ulong(),
               this->pb.val(load_instruction_lines[i].timestamp->packed).as_ulong());

        printf("  Debugging information from the transition function:\n");
        execution_checkers[i].dump();

        printf("  Memory operation executed: addr = %zu, contents_before = %zu, contents_after = %zu (ts_before = %zu, ts_after = %zu)\n",
               this->pb.val(execution_lines[i+1].address->packed).as_ulong(),
               this->pb.val(execution_lines[i+1].contents_before->packed).as_ulong(),
               this->pb.val(execution_lines[i+1].contents_after->packed).as_ulong(),
               this->pb.val(execution_lines[i].timestamp->packed).as_ulong(),
               this->pb.val(execution_lines[i+1].timestamp->packed).as_ulong());
    }
}

template<typename ramT>
void ram_universal_gadget<ramT>::print_memory_trace() const
{
    for (size_t i = 0; i < num_memory_lines; ++i)
    {
        printf("Memory access #%zu:\n", i);
        printf("  Time side  : ts = %zu, address = %zu, contents_before = %zu, contents_after = %zu\n",
               this->pb.val(unrouted_memory_lines[i]->timestamp->packed).as_ulong(),
               this->pb.val(unrouted_memory_lines[i]->address->packed).as_ulong(),
               this->pb.val(unrouted_memory_lines[i]->contents_before->packed).as_ulong(),
               this->pb.val(unrouted_memory_lines[i]->contents_after->packed).as_ulong());
        printf("  Memory side: ts = %zu, address = %zu, contents_before = %zu, contents_after = %zu\n",
               this->pb.val(routed_memory_lines[i].timestamp->packed).as_ulong(),
               this->pb.val(routed_memory_lines[i].address->packed).as_ulong(),
               this->pb.val(routed_memory_lines[i].contents_before->packed).as_ulong(),
               this->pb.val(routed_memory_lines[i].contents_after->packed).as_ulong());
    }
}

template<typename ramT>
size_t ram_universal_gadget<ramT>::packed_input_element_size(const ram_architecture_params<ramT> &ap)
{
    const size_t line_size_bits = ap.address_size() + ap.value_size();
    const size_t max_chunk_size = FieldT::capacity();
    const size_t packed_line_size = div_ceil(line_size_bits, max_chunk_size);

    return packed_line_size;
}

template<typename ramT>
size_t ram_universal_gadget<ramT>::packed_input_size(const ram_architecture_params<ramT> &ap,
                                                     const size_t boot_trace_size_bound)
{
    return packed_input_element_size(ap) * boot_trace_size_bound;
}

} // libsnark

#endif // RAM_UNIVERSAL_GADGET_TCC_

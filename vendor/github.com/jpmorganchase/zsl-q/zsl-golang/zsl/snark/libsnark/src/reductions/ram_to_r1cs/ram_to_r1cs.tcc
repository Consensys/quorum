/** @file
 *****************************************************************************

 Implementation of interfaces for a RAM-to-R1CS reduction, that is, constructing
 a R1CS ("Rank-1 Constraint System") from a RAM ("Random-Access Machine").

 See ram_to_r1cs.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_TO_R1CS_TCC_
#define RAM_TO_R1CS_TCC_

#include <set>

namespace libsnark {

template<typename ramT>
ram_to_r1cs<ramT>::ram_to_r1cs(const ram_architecture_params<ramT> &ap,
                               const size_t boot_trace_size_bound,
                               const size_t time_bound) :
    boot_trace_size_bound(boot_trace_size_bound),
    main_protoboard(ap)
{
    const size_t r1cs_input_size = ram_universal_gadget<ramT>::packed_input_size(ap, boot_trace_size_bound);
    r1cs_input.allocate(main_protoboard, r1cs_input_size, "r1cs_input");
    universal_gadget.reset(new ram_universal_gadget<ramT>(main_protoboard,
                                                          boot_trace_size_bound,
                                                          time_bound,
                                                          r1cs_input,
                                                          "universal_gadget"));
    main_protoboard.set_input_sizes(r1cs_input_size);
}

template<typename ramT>
void ram_to_r1cs<ramT>::instance_map()
{
    enter_block("Call to instance_map of ram_to_r1cs");
    universal_gadget->generate_r1cs_constraints();
    leave_block("Call to instance_map of ram_to_r1cs");
}

template<typename ramT>
r1cs_constraint_system<ram_base_field<ramT> > ram_to_r1cs<ramT>::get_constraint_system() const
{
    return main_protoboard.get_constraint_system();
}

template<typename ramT>
r1cs_primary_input<ram_base_field<ramT> > ram_to_r1cs<ramT>::auxiliary_input_map(const ram_boot_trace<ramT> &boot_trace,
                                                                                 const ram_input_tape<ramT> &auxiliary_input)
{
    enter_block("Call to witness_map of ram_to_r1cs");
    universal_gadget->generate_r1cs_witness(boot_trace, auxiliary_input);
#ifdef DEBUG
    const r1cs_primary_input<FieldT> primary_input_from_input_map = ram_to_r1cs<ramT>::primary_input_map(main_protoboard.ap, boot_trace_size_bound, boot_trace);
    const r1cs_primary_input<FieldT> primary_input_from_witness_map = main_protoboard.primary_input();
    assert(primary_input_from_input_map == primary_input_from_witness_map);
#endif
    leave_block("Call to witness_map of ram_to_r1cs");
    return main_protoboard.auxiliary_input();
}

template<typename ramT>
void ram_to_r1cs<ramT>::print_execution_trace() const
{
    universal_gadget->print_execution_trace();
}

template<typename ramT>
void ram_to_r1cs<ramT>::print_memory_trace() const
{
    universal_gadget->print_memory_trace();
}

template<typename ramT>
std::vector<ram_base_field<ramT> > ram_to_r1cs<ramT>::pack_primary_input_address_and_value(const ram_architecture_params<ramT> &ap,
                                                                                           const address_and_value &av)
{
    typedef ram_base_field<ramT> FieldT;

    const size_t address = av.first;
    const size_t contents = av.second;

    const bit_vector address_bits = convert_field_element_to_bit_vector<FieldT>(FieldT(address, true), ap.address_size());
    const bit_vector contents_bits = convert_field_element_to_bit_vector<FieldT>(FieldT(contents, true), ap.value_size());

    bit_vector trace_element_bits;
    trace_element_bits.insert(trace_element_bits.end(), address_bits.begin(), address_bits.end());
    trace_element_bits.insert(trace_element_bits.end(), contents_bits.begin(), contents_bits.end());

    const std::vector<FieldT> trace_element = pack_bit_vector_into_field_element_vector<FieldT>(trace_element_bits);

    return trace_element;
}


template<typename ramT>
r1cs_primary_input<ram_base_field<ramT> > ram_to_r1cs<ramT>::primary_input_map(const ram_architecture_params<ramT> &ap,
                                                                               const size_t boot_trace_size_bound,
                                                                               const ram_boot_trace<ramT>& boot_trace)
{
    typedef ram_base_field<ramT> FieldT;

    const size_t packed_input_element_size = ram_universal_gadget<ramT>::packed_input_element_size(ap);
    r1cs_primary_input<FieldT > result(ram_universal_gadget<ramT>::packed_input_size(ap, boot_trace_size_bound));

    std::set<size_t> bound_input_locations;

    for (auto it : boot_trace.get_all_trace_entries())
    {
        const size_t input_pos = it.first;
        const address_and_value av = it.second;

        assert(input_pos < boot_trace_size_bound);
        assert(bound_input_locations.find(input_pos) == bound_input_locations.end());

        const std::vector<FieldT> packed_input_element = ram_to_r1cs<ramT>::pack_primary_input_address_and_value(ap, av);
        std::copy(packed_input_element.begin(), packed_input_element.end(), result.begin() + packed_input_element_size * (boot_trace_size_bound - 1 - input_pos));

        bound_input_locations.insert(input_pos);
    }

    return result;
}

} // libsnark

#endif // RAM_TO_R1CS_TCC_

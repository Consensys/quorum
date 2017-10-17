/** @file
 *****************************************************************************

 Implementation of interfaces for trace-line variables.

 See trace_lines.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TRACE_LINES_TCC_
#define TRACE_LINES_TCC_

namespace libsnark {

template<typename ramT>
memory_line_variable_gadget<ramT>::memory_line_variable_gadget(ram_protoboard<ramT> &pb,
                                                               const size_t timestamp_size,
                                                               const ram_architecture_params<ramT> &ap,
                                                               const std::string &annotation_prefix) :
    ram_gadget_base<ramT>(pb, annotation_prefix)
{
    const size_t address_size = ap.address_size();
    const size_t value_size = ap.value_size();

    timestamp.reset(new dual_variable_gadget<FieldT>(pb, timestamp_size, FMT(this->annotation_prefix, " timestamp")));
    address.reset(new dual_variable_gadget<FieldT>(pb, address_size, FMT(this->annotation_prefix, " address")));
    contents_before.reset(new dual_variable_gadget<FieldT>(pb, value_size, FMT(this->annotation_prefix, " contents_before")));
    contents_after.reset(new dual_variable_gadget<FieldT>(pb, value_size, FMT(this->annotation_prefix, " contents_after")));
}

template<typename ramT>
void memory_line_variable_gadget<ramT>::generate_r1cs_constraints(const bool enforce_bitness)
{
    timestamp->generate_r1cs_constraints(enforce_bitness);
    address->generate_r1cs_constraints(enforce_bitness);
    contents_before->generate_r1cs_constraints(enforce_bitness);
    contents_after->generate_r1cs_constraints(enforce_bitness);
}

template<typename ramT>
void memory_line_variable_gadget<ramT>::generate_r1cs_witness_from_bits()
{
    timestamp->generate_r1cs_witness_from_bits();
    address->generate_r1cs_witness_from_bits();
    contents_before->generate_r1cs_witness_from_bits();
    contents_after->generate_r1cs_witness_from_bits();
}

template<typename ramT>
void memory_line_variable_gadget<ramT>::generate_r1cs_witness_from_packed()
{
    timestamp->generate_r1cs_witness_from_packed();
    address->generate_r1cs_witness_from_packed();
    contents_before->generate_r1cs_witness_from_packed();
    contents_after->generate_r1cs_witness_from_packed();
}

template<typename ramT>
pb_variable_array<ram_base_field<ramT> > memory_line_variable_gadget<ramT>::all_vars() const
{
    pb_variable_array<FieldT> r;
    r.insert(r.end(), timestamp->bits.begin(), timestamp->bits.end());
    r.insert(r.end(), address->bits.begin(), address->bits.end());
    r.insert(r.end(), contents_before->bits.begin(), contents_before->bits.end());
    r.insert(r.end(), contents_after->bits.begin(), contents_after->bits.end());

    return r;
}

template<typename ramT>
execution_line_variable_gadget<ramT>::execution_line_variable_gadget(ram_protoboard<ramT> &pb,
                                                                     const size_t timestamp_size,
                                                                     const ram_architecture_params<ramT> &ap,
                                                                     const std::string &annotation_prefix) :
    memory_line_variable_gadget<ramT>(pb, timestamp_size, ap, annotation_prefix)
{
    const size_t cpu_state_size = ap.cpu_state_size();

    cpu_state.allocate(pb, cpu_state_size, FMT(annotation_prefix, " cpu_state"));
    has_accepted.allocate(pb, FMT(annotation_prefix, " has_accepted"));
}


} // libsnark

#endif // TRACE_LINES_TCC_

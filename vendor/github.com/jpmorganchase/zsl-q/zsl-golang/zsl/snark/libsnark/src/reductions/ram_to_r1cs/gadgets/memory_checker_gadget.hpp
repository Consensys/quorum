/** @file
 *****************************************************************************

 Declaration of interfaces for memory_checker_gadget, a gadget that verifies the
 consistency of two accesses to memory that are adjacent in a "memory sort".

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MEMORY_CHECKER_GADGET_HPP_
#define MEMORY_CHECKER_GADGET_HPP_

#include "reductions/ram_to_r1cs/gadgets/trace_lines.hpp"

namespace libsnark {

template<typename ramT>
class memory_checker_gadget : public ram_gadget_base<ramT> {
private:

    typedef ram_base_field<ramT> FieldT;

    pb_variable<FieldT> timestamps_leq;
    pb_variable<FieldT> timestamps_less;
    std::shared_ptr<comparison_gadget<FieldT> > compare_timestamps;

    pb_variable<FieldT> addresses_eq;
    pb_variable<FieldT> addresses_leq;
    pb_variable<FieldT> addresses_less;
    std::shared_ptr<comparison_gadget<FieldT> > compare_addresses;

    pb_variable<FieldT> loose_contents_after1_equals_contents_before2;
    pb_variable<FieldT> loose_contents_before2_equals_zero;
    pb_variable<FieldT> loose_timestamp2_is_zero;

public:

    memory_line_variable_gadget<ramT> line1;
    memory_line_variable_gadget<ramT> line2;

    memory_checker_gadget(ram_protoboard<ramT> &pb,
                          const size_t timestamp_size,
                          const memory_line_variable_gadget<ramT> &line1,
                          const memory_line_variable_gadget<ramT> &line2,
                          const std::string& annotation_prefix="");

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "reductions/ram_to_r1cs/gadgets/memory_checker_gadget.tcc"

#endif // MEMORY_CHECKER_GADGET_HPP_

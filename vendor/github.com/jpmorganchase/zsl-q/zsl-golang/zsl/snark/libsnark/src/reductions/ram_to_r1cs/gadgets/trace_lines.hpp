/** @file
 *****************************************************************************

 Declaration of interfaces for trace-line variables.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TRACE_LINES_HPP_
#define TRACE_LINES_HPP_

#include <memory>
#include "relations/ram_computations/rams/ram_params.hpp"

namespace libsnark {

/**
 * A memory line contains variables for the following:
 * - timestamp
 * - address
 * - contents_before
 * - contents_after
 *
 * Memory lines are used by memory_checker_gadget.
 */
template<typename ramT>
class memory_line_variable_gadget : public ram_gadget_base<ramT> {
public:

    typedef ram_base_field<ramT> FieldT;

    std::shared_ptr<dual_variable_gadget<FieldT> > timestamp;
    std::shared_ptr<dual_variable_gadget<FieldT> > address;
    std::shared_ptr<dual_variable_gadget<FieldT> > contents_before;
    std::shared_ptr<dual_variable_gadget<FieldT> > contents_after;

public:

    memory_line_variable_gadget(ram_protoboard<ramT> &pb,
                                const size_t timestamp_size,
                                const ram_architecture_params<ramT> &ap,
                                const std::string &annotation_prefix="");

    void generate_r1cs_constraints(const bool enforce_bitness=false);
    void generate_r1cs_witness_from_bits();
    void generate_r1cs_witness_from_packed();

    pb_variable_array<FieldT> all_vars() const;
};

/**
 * An execution line inherits from a memory line and, in addition, contains
 * variables for a CPU state and for a flag denoting if the machine has accepted.
 *
 * Execution lines are used by execution_checker_gadget.
 */
template<typename ramT>
class execution_line_variable_gadget : public memory_line_variable_gadget<ramT> {
public:

    typedef ram_base_field<ramT> FieldT;

    pb_variable_array<FieldT> cpu_state;
    pb_variable<FieldT> has_accepted;

    execution_line_variable_gadget(ram_protoboard<ramT> &pb,
                                   const size_t timestamp_size,
                                   const ram_architecture_params<ramT> &ap,
                                   const std::string &annotation_prefix="");
};

} // libsnark

#include "reductions/ram_to_r1cs/gadgets/trace_lines.tcc"

#endif // TRACE_LINES_HPP_

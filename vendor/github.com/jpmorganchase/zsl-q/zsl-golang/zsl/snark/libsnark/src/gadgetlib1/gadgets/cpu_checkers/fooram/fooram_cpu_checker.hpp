/** @file
 *****************************************************************************

 Declaration of interfaces for the FOORAM CPU checker gadget.

 The gadget checks the correct operation for the CPU of the FOORAM architecture.

 In FOORAM, the only instruction is FOO(x) and its encoding is x.
 The instruction FOO(x) has the following semantics:
 - if x is odd: reg <- [2*x+(pc+1)]
 - if x is even: [pc+x] <- reg+pc
 - increment pc by 1

 Starting from empty memory, FOORAM performs non-trivial pseudo-random computation
 that exercises both loads, stores, and instruction fetches.

 E.g. for the first 200 steps on 16 cell machine we get 93 different memory configurations.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FOORAM_CPU_CHECKER_HPP_
#define FOORAM_CPU_CHECKER_HPP_

#include <cstddef>
#include <memory>

#include "common/serialization.hpp"
#include "gadgetlib1/gadget.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/fooram/components/fooram_protoboard.hpp"
#include "gadgetlib1/gadgets/cpu_checkers/fooram/components/bar_gadget.hpp"
#include "relations/ram_computations/memory/memory_interface.hpp"

namespace libsnark {

template<typename FieldT>
class fooram_cpu_checker : public fooram_gadget<FieldT> {
public:
    pb_variable_array<FieldT> prev_pc_addr;
    pb_variable_array<FieldT> prev_pc_val;
    pb_variable_array<FieldT> prev_state;
    pb_variable_array<FieldT> guess;
    pb_variable_array<FieldT> ls_addr;
    pb_variable_array<FieldT> ls_prev_val;
    pb_variable_array<FieldT> ls_next_val;
    pb_variable_array<FieldT> next_state;
    pb_variable_array<FieldT> next_pc_addr;
    pb_variable<FieldT> next_has_accepted;

    pb_variable<FieldT> zero;
    pb_variable<FieldT> packed_next_pc_addr;
    pb_linear_combination_array<FieldT> one_as_addr;
    std::shared_ptr<packing_gadget<FieldT> > pack_next_pc_addr;

    pb_variable<FieldT> packed_load_addr;
    pb_variable<FieldT> packed_store_addr;
    pb_variable<FieldT> packed_store_val;

    std::shared_ptr<bar_gadget<FieldT> > increment_pc;
    std::shared_ptr<bar_gadget<FieldT> > compute_packed_load_addr;
    std::shared_ptr<bar_gadget<FieldT> > compute_packed_store_addr;
    std::shared_ptr<bar_gadget<FieldT> > compute_packed_store_val;

    pb_variable<FieldT> packed_ls_addr;
    pb_variable<FieldT> packed_ls_prev_val;
    pb_variable<FieldT> packed_ls_next_val;
    pb_variable<FieldT> packed_prev_state;
    pb_variable<FieldT> packed_next_state;
    std::shared_ptr<packing_gadget<FieldT> > pack_ls_addr;
    std::shared_ptr<packing_gadget<FieldT> > pack_ls_prev_val;
    std::shared_ptr<packing_gadget<FieldT> > pack_ls_next_val;
    std::shared_ptr<packing_gadget<FieldT> > pack_prev_state;
    std::shared_ptr<packing_gadget<FieldT> > pack_next_state;

    fooram_cpu_checker(fooram_protoboard<FieldT> &pb,
                       pb_variable_array<FieldT> &prev_pc_addr,
                       pb_variable_array<FieldT> &prev_pc_val,
                       pb_variable_array<FieldT> &prev_state,
                       pb_variable_array<FieldT> &ls_addr,
                       pb_variable_array<FieldT> &ls_prev_val,
                       pb_variable_array<FieldT> &ls_next_val,
                       pb_variable_array<FieldT> &next_state,
                       pb_variable_array<FieldT> &next_pc_addr,
                       pb_variable<FieldT> &next_has_accepted,
                       const std::string &annotation_prefix);

    void generate_r1cs_constraints();

    void generate_r1cs_witness() { assert(0); }

    void generate_r1cs_witness_address();

    void generate_r1cs_witness_other(fooram_input_tape_iterator &aux_it,
                                     const fooram_input_tape_iterator &aux_end);

    void dump() const;
};

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/fooram/fooram_cpu_checker.tcc"

#endif // FORAM_CPU_CHECKER_HPP_

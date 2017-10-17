/** @file
 *****************************************************************************

 Implementation of interfaces for the FOORAM CPU checker gadget.

 See fooram_cpu_checker.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FOORAM_CPU_CHECKER_TCC_
#define FOORAM_CPU_CHECKER_TCC_

namespace libsnark {

template<typename FieldT>
fooram_cpu_checker<FieldT>::fooram_cpu_checker(fooram_protoboard<FieldT> &pb,
                                               pb_variable_array<FieldT> &prev_pc_addr,
                                               pb_variable_array<FieldT> &prev_pc_val,
                                               pb_variable_array<FieldT> &prev_state,
                                               pb_variable_array<FieldT> &ls_addr,
                                               pb_variable_array<FieldT> &ls_prev_val,
                                               pb_variable_array<FieldT> &ls_next_val,
                                               pb_variable_array<FieldT> &next_state,
                                               pb_variable_array<FieldT> &next_pc_addr,
                                               pb_variable<FieldT> &next_has_accepted,
                                               const std::string &annotation_prefix) :
    fooram_gadget<FieldT>(pb, annotation_prefix),
    prev_pc_addr(prev_pc_addr),
    prev_pc_val(prev_pc_val),
    prev_state(prev_state),
    ls_addr(ls_addr),
    ls_prev_val(ls_prev_val),
    ls_next_val(ls_next_val),
    next_state(next_state),
    next_pc_addr(next_pc_addr),
    next_has_accepted(next_has_accepted)
{
    /* increment PC */
    packed_next_pc_addr.allocate(pb, FMT(annotation_prefix, " packed_next_pc_addr"));
    pack_next_pc_addr.reset(new packing_gadget<FieldT>(pb, next_pc_addr, packed_next_pc_addr, FMT(annotation_prefix, " pack_next_pc_addr")));

    one_as_addr.resize(next_pc_addr.size());
    one_as_addr[0].assign(this->pb, 1);
    for (size_t i = 1; i < next_pc_addr.size(); ++i)
    {
        one_as_addr[i].assign(this->pb, 0);
    }

    /* packed_next_pc_addr = prev_pc_addr + one_as_addr */
    increment_pc.reset(new bar_gadget<FieldT>(pb, prev_pc_addr, FieldT::one(), one_as_addr, FieldT::one(), packed_next_pc_addr, FMT(annotation_prefix, " increment_pc")));

    /* packed_store_addr = prev_pc_addr + prev_pc_val */
    packed_store_addr.allocate(pb, FMT(annotation_prefix, " packed_store_addr"));
    compute_packed_store_addr.reset(new bar_gadget<FieldT>(pb, prev_pc_addr, FieldT::one(), prev_pc_val, FieldT::one(), packed_store_addr, FMT(annotation_prefix, " compute_packed_store_addr")));

    /* packed_load_addr = 2 * x + next_pc_addr */
    packed_load_addr.allocate(pb, FMT(annotation_prefix, " packed_load_addr"));
    compute_packed_load_addr.reset(new bar_gadget<FieldT>(pb, prev_pc_val, FieldT(2), next_pc_addr, FieldT::one(), packed_load_addr, FMT(annotation_prefix, " compute_packed_load_addr")));

    /*
      packed_ls_addr = x0 * packed_load_addr + (1-x0) * packed_store_addr
      packed_ls_addr ~ ls_addr
    */
    packed_ls_addr.allocate(pb, FMT(annotation_prefix, " packed_ls_addr"));
    pack_ls_addr.reset(new packing_gadget<FieldT>(pb, ls_addr, packed_ls_addr, " pack_ls_addr"));

    /* packed_store_val = prev_state_bits + prev_pc_addr */
    packed_store_val.allocate(pb, FMT(annotation_prefix, " packed_store_val"));
    compute_packed_store_val.reset(new bar_gadget<FieldT>(pb, prev_state, FieldT::one(), prev_pc_addr, FieldT::one(), packed_store_val, FMT(annotation_prefix, " compute_packed_store_val")));

    /*
      packed_ls_next_val = x0 * packed_ls_prev_val + (1-x0) * packed_store_val
      packed_ls_next_val ~ ls_next_val
    */
    packed_ls_prev_val.allocate(pb, FMT(annotation_prefix, " packed_ls_prev_val"));
    pack_ls_prev_val.reset(new packing_gadget<FieldT>(this->pb, ls_prev_val, packed_ls_prev_val, FMT(annotation_prefix, " pack_ls_prev_val")));
    packed_ls_next_val.allocate(pb, FMT(annotation_prefix, " packed_ls_next_val"));
    pack_ls_next_val.reset(new packing_gadget<FieldT>(this->pb, ls_next_val, packed_ls_next_val, FMT(annotation_prefix, " pack_ls_next_val")));

    /*
      packed_next_state = x0 * packed_ls_prev_val + (1-x0) * packed_prev_state
      packed_next_state ~ next_state
      packed_prev_state ~ prev_state
    */
    packed_prev_state.allocate(pb, FMT(annotation_prefix, " packed_prev_state"));
    pack_prev_state.reset(new packing_gadget<FieldT>(pb, prev_state, packed_prev_state, " pack_prev_state"));

    packed_next_state.allocate(pb, FMT(annotation_prefix, " packed_next_state"));
    pack_next_state.reset(new packing_gadget<FieldT>(pb, next_state, packed_next_state, " pack_next_state"));

    /* next_has_accepted = 1 */
}

template<typename FieldT>
void fooram_cpu_checker<FieldT>::generate_r1cs_constraints()
{
    /* packed_next_pc_addr = prev_pc_addr + one_as_addr */
    pack_next_pc_addr->generate_r1cs_constraints(false);
    increment_pc->generate_r1cs_constraints();

    /* packed_store_addr = prev_pc_addr + prev_pc_val */
    compute_packed_store_addr->generate_r1cs_constraints();

    /* packed_load_addr = 2 * x + next_pc_addr */
    compute_packed_load_addr->generate_r1cs_constraints();

    /*
      packed_ls_addr = x0 * packed_load_addr + (1-x0) * packed_store_addr
      packed_ls_addr - packed_store_addr = x0 * (packed_load_addr - packed_store_addr)
      packed_ls_addr ~ ls_addr
    */
    pack_ls_addr->generate_r1cs_constraints(false);
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(prev_pc_val[0],
                                                         packed_load_addr - packed_store_addr,
                                                         packed_ls_addr - packed_store_addr),
                                 FMT(this->annotation_prefix, " compute_ls_addr_packed"));

    /* packed_store_val = prev_state_bits + prev_pc_addr */
    compute_packed_store_val->generate_r1cs_constraints();

    /*
      packed_ls_next_val = x0 * packed_ls_prev_val + (1-x0) * packed_store_val
      packed_ls_next_val - packed_store_val = x0 * (packed_ls_prev_val - packed_store_val)
      packed_ls_next_val ~ ls_next_val
    */
    pack_ls_prev_val->generate_r1cs_constraints(false);
    pack_ls_next_val->generate_r1cs_constraints(false);
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(prev_pc_val[0],
                                                         packed_ls_prev_val - packed_store_val,
                                                         packed_ls_next_val - packed_store_val),
                                 FMT(this->annotation_prefix, " compute_packed_ls_next_val"));

    /*
      packed_next_state = x0 * packed_ls_prev_val + (1-x0) * packed_prev_state
      packed_next_state - packed_prev_state = x0 * (packed_ls_prev_val - packed_prev_state)
      packed_next_state ~ next_state
      packed_prev_state ~ prev_state
    */
    pack_prev_state->generate_r1cs_constraints(false);
    pack_next_state->generate_r1cs_constraints(false);
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(prev_pc_val[0],
                                                         packed_ls_prev_val - packed_prev_state,
                                                         packed_next_state - packed_prev_state),
                                 FMT(this->annotation_prefix, " compute_packed_next_state"));

    /* next_has_accepted = 1 */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, next_has_accepted, 1), FMT(this->annotation_prefix, " always_accepted"));
}

template<typename FieldT>
void fooram_cpu_checker<FieldT>::generate_r1cs_witness_address()
{
    one_as_addr.evaluate(this->pb);

    /* packed_next_pc_addr = prev_pc_addr + one_as_addr */
    increment_pc->generate_r1cs_witness();
    pack_next_pc_addr->generate_r1cs_witness_from_packed();

    /* packed_store_addr = prev_pc_addr + prev_pc_val */
    compute_packed_store_addr->generate_r1cs_witness();

    /* packed_load_addr = 2 * x + next_pc_addr */
    compute_packed_load_addr->generate_r1cs_witness();

    /*
      packed_ls_addr = x0 * packed_load_addr + (1-x0) * packed_store_addr
      packed_ls_addr - packed_store_addr = x0 * (packed_load_addr - packed_store_addr)
      packed_ls_addr ~ ls_addr
    */
    this->pb.val(packed_ls_addr) = (this->pb.val(prev_pc_val[0]) * this->pb.val(packed_load_addr) +
                                    (FieldT::one()-this->pb.val(prev_pc_val[0])) * this->pb.val(packed_store_addr));
    pack_ls_addr->generate_r1cs_witness_from_packed();
}

template<typename FieldT>
void fooram_cpu_checker<FieldT>::generate_r1cs_witness_other(fooram_input_tape_iterator &aux_it,
                                                             const fooram_input_tape_iterator &aux_end)
{
    /* fooram memory contents do not depend on program/input. */
    UNUSED(aux_it, aux_end);
    /* packed_store_val = prev_state_bits + prev_pc_addr */
    compute_packed_store_val->generate_r1cs_witness();

    /*
      packed_ls_next_val = x0 * packed_ls_prev_val + (1-x0) * packed_store_val
      packed_ls_next_val - packed_store_val = x0 * (packed_ls_prev_val - packed_store_val)
      packed_ls_next_val ~ ls_next_val
    */
    pack_ls_prev_val->generate_r1cs_witness_from_bits();
    this->pb.val(packed_ls_next_val) = (this->pb.val(prev_pc_val[0]) * this->pb.val(packed_ls_prev_val) +
                                        (FieldT::one() - this->pb.val(prev_pc_val[0])) * this->pb.val(packed_store_val));
    pack_ls_next_val->generate_r1cs_witness_from_packed();

    /*
      packed_next_state = x0 * packed_ls_prev_val + (1-x0) * packed_prev_state
      packed_next_state - packed_prev_state = x0 * (packed_ls_prev_val - packed_prev_state)
      packed_next_state ~ next_state
      packed_prev_state ~ prev_state
    */
    pack_prev_state->generate_r1cs_witness_from_bits();
    this->pb.val(packed_next_state) = (this->pb.val(prev_pc_val[0]) * this->pb.val(packed_ls_prev_val) +
                                       (FieldT::one() - this->pb.val(prev_pc_val[0])) * this->pb.val(packed_prev_state));
    pack_next_state->generate_r1cs_witness_from_packed();

    /* next_has_accepted = 1 */
    this->pb.val(next_has_accepted) = FieldT::one();
}

template<typename FieldT>
void fooram_cpu_checker<FieldT>::dump() const
{
    printf("packed_store_addr: ");
    this->pb.val(packed_store_addr).print();
    printf("packed_load_addr: ");
    this->pb.val(packed_load_addr).print();
    printf("packed_ls_addr: ");
    this->pb.val(packed_ls_addr).print();
    printf("packed_store_val: ");
    this->pb.val(packed_store_val).print();
    printf("packed_ls_next_val: ");
    this->pb.val(packed_ls_next_val).print();
    printf("packed_next_state: ");
    this->pb.val(packed_next_state).print();
}

} // libsnark

#endif // FOORAM_CPU_CHECKER_TCC

/** @file
 *****************************************************************************

 Implementation of interfaces for memory_checker_gadget.

 See memory_checker_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef MEMORY_CHECKER_GADGET_TCC_
#define MEMORY_CHECKER_GADGET_TCC_

namespace libsnark {

template<typename ramT>
memory_checker_gadget<ramT>::memory_checker_gadget(ram_protoboard<ramT> &pb,
                                                   const size_t timestamp_size,
                                                   const memory_line_variable_gadget<ramT> &line1,
                                                   const memory_line_variable_gadget<ramT> &line2,
                                                   const std::string& annotation_prefix) :
    ram_gadget_base<ramT>(pb, annotation_prefix), line1(line1), line2(line2)
{
    /* compare the two timestamps */
    timestamps_leq.allocate(pb, FMT(this->annotation_prefix, " timestamps_leq"));
    timestamps_less.allocate(pb, FMT(this->annotation_prefix, " timestamps_less"));
    compare_timestamps.reset(new comparison_gadget<FieldT>(pb, timestamp_size, line1.timestamp->packed, line2.timestamp->packed, timestamps_less, timestamps_leq,
                                                           FMT(this->annotation_prefix, " compare_ts")));


    /* compare the two addresses */
    const size_t address_size = pb.ap.address_size();
    addresses_eq.allocate(pb, FMT(this->annotation_prefix, " addresses_eq"));
    addresses_leq.allocate(pb, FMT(this->annotation_prefix, " addresses_leq"));
    addresses_less.allocate(pb, FMT(this->annotation_prefix, " addresses_less"));
    compare_addresses.reset(new comparison_gadget<FieldT>(pb, address_size, line1.address->packed, line2.address->packed, addresses_less, addresses_leq,
                                                          FMT(this->annotation_prefix, " compare_addresses")));

    /*
      Add variables that will contain flags representing the following relations:
      - "line1.contents_after = line2.contents_before" (to check that contents do not change between instructions);
      - "line2.contents_before = 0" (for the first access at an address); and
      - "line2.timestamp = 0" (for wrap-around checks to ensure only one 'cycle' in the memory sort).

      More precisely, each of the above flags is "loose" (i.e., it equals 0 if
      the relation holds, but can be either 0 or 1 if the relation does not hold).
     */
    loose_contents_after1_equals_contents_before2.allocate(pb, FMT(this->annotation_prefix, " loose_contents_after1_equals_contents_before2"));
    loose_contents_before2_equals_zero.allocate(pb, FMT(this->annotation_prefix, " loose_contents_before2_equals_zero"));
    loose_timestamp2_is_zero.allocate(pb, FMT(this->annotation_prefix, " loose_timestamp2_is_zero"));
}

template<typename ramT>
void memory_checker_gadget<ramT>::generate_r1cs_constraints()
{
    /* compare the two timestamps */
    compare_timestamps->generate_r1cs_constraints();

    /* compare the two addresses */
    compare_addresses->generate_r1cs_constraints();
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(addresses_leq, 1 - addresses_less, addresses_eq), FMT(this->annotation_prefix, " addresses_eq"));

    /*
      Add constraints for the following three flags:
       - loose_contents_after1_equals_contents_before2;
       - loose_contents_before2_equals_zero;
       - loose_timestamp2_is_zero.
     */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(loose_contents_after1_equals_contents_before2,
                                                         line1.contents_after->packed - line2.contents_before->packed, 0),
                                 FMT(this->annotation_prefix, " loose_contents_after1_equals_contents_before2"));
    generate_boolean_r1cs_constraint<FieldT>(this->pb, loose_contents_after1_equals_contents_before2, FMT(this->annotation_prefix, " loose_contents_after1_equals_contents_before2"));

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(loose_contents_before2_equals_zero,
                                                         line2.contents_before->packed, 0),
                                 FMT(this->annotation_prefix, " loose_contents_before2_equals_zero"));
    generate_boolean_r1cs_constraint<FieldT>(this->pb, loose_contents_before2_equals_zero, FMT(this->annotation_prefix, " loose_contents_before2_equals_zero"));

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(loose_timestamp2_is_zero,
                                                         line2.timestamp->packed, 0),
                                 FMT(this->annotation_prefix, " loose_timestamp2_is_zero"));
    generate_boolean_r1cs_constraint<FieldT>(this->pb, loose_timestamp2_is_zero, FMT(this->annotation_prefix, " loose_timestamp2_is_zero"));

    /*
      The three cases that need to be checked are:

      line1.address = line2.address => line1.contents_after = line2.contents_before
      (i.e. contents do not change between accesses to the same address)

      line1.address < line2.address => line2.contents_before = 0
      (i.e. access to new address has the "before" value set to 0)

      line1.address > line2.address => line2.timestamp = 0
      (i.e. there is only one cycle with non-decreasing addresses, except
      for the case where we go back to a unique pre-set timestamp; we choose
      timestamp 0 to be the one that touches address 0)

      As usual, we implement "A => B" as "NOT (A AND (NOT B))".
    */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(addresses_eq, 1 - loose_contents_after1_equals_contents_before2, 0),
                                 FMT(this->annotation_prefix, " memory_retains_contents_between_accesses"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(addresses_less, 1 - loose_contents_before2_equals_zero, 0),
                                 FMT(this->annotation_prefix, " new_address_starts_at_zero"));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1 - addresses_leq, 1 - loose_timestamp2_is_zero, 0),
                                 FMT(this->annotation_prefix, " only_one_cycle"));
}

template<typename ramT>
void memory_checker_gadget<ramT>::generate_r1cs_witness()
{
    /* compare the two addresses */
    compare_addresses->generate_r1cs_witness();
    this->pb.val(addresses_eq) = this->pb.val(addresses_leq) * (FieldT::one() - this->pb.val(addresses_less));

    /* compare the two timestamps */
    compare_timestamps->generate_r1cs_witness();

    /*
      compare the values of:
      - loose_contents_after1_equals_contents_before2;
      - loose_contents_before2_equals_zero;
      - loose_timestamp2_is_zero.
     */
    this->pb.val(loose_contents_after1_equals_contents_before2) = (this->pb.val(line1.contents_after->packed) == this->pb.val(line2.contents_before->packed)) ? FieldT::one() : FieldT::zero();
    this->pb.val(loose_contents_before2_equals_zero) = this->pb.val(line2.contents_before->packed).is_zero() ? FieldT::one() : FieldT::zero();
    this->pb.val(loose_timestamp2_is_zero) = (this->pb.val(line2.timestamp->packed) == FieldT::zero() ? FieldT::one() : FieldT::zero());
}

} // libsnark

#endif // MEMORY_CHECKER_GADGET_TCC_

/** @file
 *****************************************************************************

 Implementation of interfaces for the TinyRAM ALU control-flow gadgets.

 See alu_control_flow.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef ALU_CONTROL_FLOW_TCC_
#define ALU_CONTROL_FLOW_TCC_

#include "common/profiling.hpp"

namespace libsnark {

/* jmp */
template<typename FieldT>
void ALU_jmp_gadget<FieldT>::generate_r1cs_constraints()
{
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->argval2.packed },
            { this->result }),
        FMT(this->annotation_prefix, " jmp_result"));
}

template<typename FieldT>
void ALU_jmp_gadget<FieldT>::generate_r1cs_witness()
{
    this->pb.val(this->result) = this->pb.val(this->argval2.packed);
}

template<typename FieldT>
void test_ALU_jmp_gadget()
{
    print_time("starting jmp test");

    tinyram_architecture_params ap(16, 16);
    tinyram_program P; P.instructions = generate_tinyram_prelude(ap);
    tinyram_protoboard<FieldT> pb(ap, P.size(), 0, 10);

    word_variable_gadget<FieldT> pc(pb, "pc"), argval2(pb, "argval2");
    pb_variable<FieldT> flag, result;

    pc.generate_r1cs_constraints(true);
    argval2.generate_r1cs_constraints(true);
    flag.allocate(pb, "flag");
    result.allocate(pb, "result");

    ALU_jmp_gadget<FieldT> jmp(pb, pc, argval2, flag, result, "jmp");
    jmp.generate_r1cs_constraints();

    pb.val(argval2.packed) = FieldT(123);
    argval2.generate_r1cs_witness_from_packed();

    jmp.generate_r1cs_witness();

    assert(pb.val(result) == FieldT(123));
    assert(pb.is_satisfied());
    print_time("positive jmp test successful");

    pb.val(result) = FieldT(1);
    assert(!pb.is_satisfied());
    print_time("negative jmp test successful");
}

/* cjmp */
template<typename FieldT>
void ALU_cjmp_gadget<FieldT>::generate_r1cs_constraints()
{
    /*
      flag1 * argval2 + (1-flag1) * (pc1 + 1) = cjmp_result
      flag1 * (argval2 - pc1 - 1) = cjmp_result - pc1 - 1

      Note that instruction fetch semantics require program counter to
      be aligned to the double word by rounding down, and pc_addr in
      the outer reduction is expressed as a double word address. To
      achieve this we just discard the first ap.subaddr_len() bits of
      the byte address of the PC.
    */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            this->flag,
            pb_packing_sum<FieldT>(pb_variable_array<FieldT>(this->argval2.bits.begin() + this->pb.ap.subaddr_len(), this->argval2.bits.end())) - this->pc.packed - 1,
            this->result - this->pc.packed - 1),
        FMT(this->annotation_prefix, " cjmp_result"));
}

template<typename FieldT>
void ALU_cjmp_gadget<FieldT>::generate_r1cs_witness()
{
    this->pb.val(this->result) = ((this->pb.val(this->flag) == FieldT::one()) ?
                                  FieldT(this->pb.val(this->argval2.packed).as_ulong() >> this->pb.ap.subaddr_len()) :
                                  this->pb.val(this->pc.packed) + FieldT::one());
}

template<typename FieldT>
void test_ALU_cjmp_gadget()
{
    // TODO: update
    print_time("starting cjmp test");

    tinyram_architecture_params ap(16, 16);
    tinyram_program P; P.instructions = generate_tinyram_prelude(ap);
    tinyram_protoboard<FieldT> pb(ap, P.size(), 0, 10);

    word_variable_gadget<FieldT> pc(pb, "pc"), argval2(pb, "argval2");
    pb_variable<FieldT> flag, result;

    pc.generate_r1cs_constraints(true);
    argval2.generate_r1cs_constraints(true);
    flag.allocate(pb, "flag");
    result.allocate(pb, "result");

    ALU_cjmp_gadget<FieldT> cjmp(pb, pc, argval2, flag, result, "cjmp");
    cjmp.generate_r1cs_constraints();

    pb.val(argval2.packed) = FieldT(123);
    argval2.generate_r1cs_witness_from_packed();
    pb.val(pc.packed) = FieldT(456);
    pc.generate_r1cs_witness_from_packed();

    pb.val(flag) = FieldT(1);
    cjmp.generate_r1cs_witness();

    assert(pb.val(result) == FieldT(123));
    assert(pb.is_satisfied());
    print_time("positive cjmp test successful");

    pb.val(flag) = FieldT(0);
    assert(!pb.is_satisfied());
    print_time("negative cjmp test successful");

    pb.val(flag) = FieldT(0);
    cjmp.generate_r1cs_witness();

    assert(pb.val(result) == FieldT(456+2*ap.w/8));
    assert(pb.is_satisfied());
    print_time("positive cjmp test successful");

    pb.val(flag) = FieldT(1);
    assert(!pb.is_satisfied());
    print_time("negative cjmp test successful");
}

/* cnjmp */
template<typename FieldT>
void ALU_cnjmp_gadget<FieldT>::generate_r1cs_constraints()
{
    /*
      flag1 * (pc1 + inc) + (1-flag1) * argval2 = cnjmp_result
      flag1 * (pc1 + inc - argval2) = cnjmp_result - argval2

      Note that instruction fetch semantics require program counter to
      be aligned to the double word by rounding down, and pc_addr in
      the outer reduction is expressed as a double word address. To
      achieve this we just discard the first ap.subaddr_len() bits of
      the byte address of the PC.
    */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            this->flag,
            this->pc.packed + 1 - pb_packing_sum<FieldT>(pb_variable_array<FieldT>(this->argval2.bits.begin() + this->pb.ap.subaddr_len(), this->argval2.bits.end())),
            this->result - pb_packing_sum<FieldT>(pb_variable_array<FieldT>(this->argval2.bits.begin() + this->pb.ap.subaddr_len(), this->argval2.bits.end()))),
        FMT(this->annotation_prefix, " cnjmp_result"));
}

template<typename FieldT>
void ALU_cnjmp_gadget<FieldT>::generate_r1cs_witness()
{
    this->pb.val(this->result) = ((this->pb.val(this->flag) == FieldT::one()) ?
                                  this->pb.val(this->pc.packed) + FieldT::one() :
                                  FieldT(this->pb.val(this->argval2.packed).as_ulong() >> this->pb.ap.subaddr_len()));
}

template<typename FieldT>
void test_ALU_cnjmp_gadget()
{
    // TODO: update
    print_time("starting cnjmp test");

    tinyram_architecture_params ap(16, 16);
    tinyram_program P; P.instructions = generate_tinyram_prelude(ap);
    tinyram_protoboard<FieldT> pb(ap, P.size(), 0, 10);

    word_variable_gadget<FieldT> pc(pb, "pc"), argval2(pb, "argval2");
    pb_variable<FieldT> flag, result;

    pc.generate_r1cs_constraints(true);
    argval2.generate_r1cs_constraints(true);
    flag.allocate(pb, "flag");
    result.allocate(pb, "result");

    ALU_cnjmp_gadget<FieldT> cnjmp(pb, pc, argval2, flag, result, "cjmp");
    cnjmp.generate_r1cs_constraints();

    pb.val(argval2.packed) = FieldT(123);
    argval2.generate_r1cs_witness_from_packed();
    pb.val(pc.packed) = FieldT(456);
    pc.generate_r1cs_witness_from_packed();

    pb.val(flag) = FieldT(0);
    cnjmp.generate_r1cs_witness();

    assert(pb.val(result) == FieldT(123));
    assert(pb.is_satisfied());
    print_time("positive cnjmp test successful");

    pb.val(flag) = FieldT(1);
    assert(!pb.is_satisfied());
    print_time("negative cnjmp test successful");

    pb.val(flag) = FieldT(1);
    cnjmp.generate_r1cs_witness();

    assert(pb.val(result) == FieldT(456 + (2*pb.ap.w/8)));
    assert(pb.is_satisfied());
    print_time("positive cnjmp test successful");

    pb.val(flag) = FieldT(0);
    assert(!pb.is_satisfied());
    print_time("negative cnjmp test successful");
}

} // libsnark

#endif // ALU_CONTROL_FLOW_TCC_

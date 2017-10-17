/** @file
 *****************************************************************************

 Implementation of interfaces for the TinyRAM ALU arithmetic gadgets.

 See alu_arithmetic.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef ALU_ARITHMETIC_TCC_
#define ALU_ARITHMETIC_TCC_

#include "common/profiling.hpp"
#include "common/utils.hpp"

namespace libsnark {

/* the code here is full of template lambda magic, but it is better to
   have limited presence of such code than to have code duplication in
   testing functions, which basically do the same thing: brute force
   the range of inputs which different success predicates */

template<class T, typename FieldT>
using initializer_fn =
    std::function<T*
                  (tinyram_protoboard<FieldT>&,    // pb
                   pb_variable_array<FieldT>&,       // opcode_indicators
                   word_variable_gadget<FieldT>&, // desval
                   word_variable_gadget<FieldT>&, // arg1val
                   word_variable_gadget<FieldT>&, // arg2val
                   pb_variable<FieldT>&,             // flag
                   pb_variable<FieldT>&,             // result
                   pb_variable<FieldT>&              // result_flag
                  )>;

template<class T, typename FieldT>
void brute_force_arithmetic_gadget(const size_t w,
                                   const size_t opcode,
                                   initializer_fn<T, FieldT> initializer,
                                   std::function<size_t(size_t,bool,size_t,size_t)> res_function,
                                   std::function<bool(size_t,bool,size_t,size_t)> flag_function)
/* parameters for res_function and flag_function are both desval, flag, arg1val, arg2val */
{
    printf("testing on all %zu bit inputs\n", w);

    tinyram_architecture_params ap(w, 16);
    tinyram_program P; P.instructions = generate_tinyram_prelude(ap);
    tinyram_protoboard<FieldT> pb(ap, P.size(), 0, 10);

    pb_variable_array<FieldT> opcode_indicators;
    opcode_indicators.allocate(pb, 1ul<<ap.opcode_width(), "opcode_indicators");
    for (size_t i = 0; i < 1ul<<ap.opcode_width(); ++i)
    {
        pb.val(opcode_indicators[i]) = (i == opcode ? FieldT::one() : FieldT::zero());
    }

    word_variable_gadget<FieldT> desval(pb, "desval");
    desval.generate_r1cs_constraints(true);
    word_variable_gadget<FieldT> arg1val(pb, "arg1val");
    arg1val.generate_r1cs_constraints(true);
    word_variable_gadget<FieldT> arg2val(pb, "arg2val");
    arg2val.generate_r1cs_constraints(true);
    pb_variable<FieldT> flag; flag.allocate(pb, "flag");
    pb_variable<FieldT> result; result.allocate(pb, "result");
    pb_variable<FieldT> result_flag; result_flag.allocate(pb, "result_flag");

    std::unique_ptr<T> g;
    g.reset(initializer(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag));
    g->generate_r1cs_constraints();

    for (size_t des = 0; des < (1u << w); ++des)
    {
        pb.val(desval.packed) = FieldT(des);
        desval.generate_r1cs_witness_from_packed();

        for (char f = 0; f <= 1; ++f)
        {
            pb.val(flag) = (f ? FieldT::one() : FieldT::zero());

            for (size_t arg1 = 0; arg1 < (1u << w); ++arg1)
            {
                pb.val(arg1val.packed) = FieldT(arg1);
                arg1val.generate_r1cs_witness_from_packed();

                for (size_t arg2 = 0; arg2 < (1u << w); ++arg2)
                {
                    pb.val(arg2val.packed) = FieldT(arg2);
                    arg2val.generate_r1cs_witness_from_packed();

                    size_t res = res_function(des, f, arg1, arg2);
                    bool res_f = flag_function(des, f, arg1, arg2);
#ifdef DEBUG
                    printf("with the following parameters: flag = %d"
                           ", desval = %zu (%d)"
                           ", arg1val = %zu (%d)"
                           ", arg2val = %zu (%d)"
                           ". expected result: %zu (%d), expected flag: %d\n",
                           f,
                           des, from_twos_complement(des, w),
                           arg1, from_twos_complement(arg1, w),
                           arg2, from_twos_complement(arg2, w),
                           res, from_twos_complement(res, w), res_f);
#endif
                    g->generate_r1cs_witness();
#ifdef DEBUG
                    printf("result: ");
                    pb.val(result).print();
                    printf("flag: ");
                    pb.val(result_flag).print();
#endif
                    assert(pb.is_satisfied());
                    assert(pb.val(result) == FieldT(res));
                    assert(pb.val(result_flag) == (res_f ? FieldT::one() : FieldT::zero()));
                }
            }
        }
    }
}

/* and */
template<typename FieldT>
void ALU_and_gadget<FieldT>::generate_r1cs_constraints()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        this->pb.add_r1cs_constraint(
            r1cs_constraint<FieldT>(
                { this->arg1val.bits[i] },
                { this->arg2val.bits[i] },
                { this->res_word[i] }),
            FMT(this->annotation_prefix, " res_word_%zu", i));
    }

    /* generate result */
    pack_result->generate_r1cs_constraints(false);
    not_all_zeros->generate_r1cs_constraints();

    /* result_flag = 1 - not_all_zeros = result is 0^w */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, this->not_all_zeros_result * (-1) },
            { this->result_flag }),
        FMT(this->annotation_prefix, " result_flag"));
}

template<typename FieldT>
void ALU_and_gadget<FieldT>::generate_r1cs_witness()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        bool b1 = this->pb.val(this->arg1val.bits[i]) == FieldT::one();
        bool b2 = this->pb.val(this->arg2val.bits[i]) == FieldT::one();

        this->pb.val(this->res_word[i]) = (b1 && b2 ? FieldT::one() : FieldT::zero());
    }

    pack_result->generate_r1cs_witness_from_bits();
    not_all_zeros->generate_r1cs_witness();
    this->pb.val(this->result_flag) = FieldT::one() - this->pb.val(not_all_zeros_result);
}

template<typename FieldT>
void test_ALU_and_gadget(const size_t w)
{
    print_time("starting and test");
    brute_force_arithmetic_gadget<ALU_and_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_AND,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_and_gadget<FieldT>* {
                                                                      return new ALU_and_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, "ALU_and_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return x & y; },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return (x & y) == 0; });
    print_time("and tests successful");
}

/* or */
template<typename FieldT>
void ALU_or_gadget<FieldT>::generate_r1cs_constraints()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        this->pb.add_r1cs_constraint(
            r1cs_constraint<FieldT>(
                { ONE, this->arg1val.bits[i] * (-1) },
                { ONE, this->arg2val.bits[i] * (-1) },
                { ONE, this->res_word[i] * (-1) }),
            FMT(this->annotation_prefix, " res_word_%zu", i));
    }

    /* generate result */
    pack_result->generate_r1cs_constraints(false);
    not_all_zeros->generate_r1cs_constraints();

    /* result_flag = 1 - not_all_zeros = result is 0^w */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, this->not_all_zeros_result * (-1) },
            { this->result_flag }),
        FMT(this->annotation_prefix, " result_flag"));
}

template<typename FieldT>
void ALU_or_gadget<FieldT>::generate_r1cs_witness()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        bool b1 = this->pb.val(this->arg1val.bits[i]) == FieldT::one();
        bool b2 = this->pb.val(this->arg2val.bits[i]) == FieldT::one();

        this->pb.val(this->res_word[i]) = (b1 || b2 ? FieldT::one() : FieldT::zero());
    }

    pack_result->generate_r1cs_witness_from_bits();
    not_all_zeros->generate_r1cs_witness();
    this->pb.val(this->result_flag) = FieldT::one() - this->pb.val(this->not_all_zeros_result);
}

template<typename FieldT>
void test_ALU_or_gadget(const size_t w)
{
    print_time("starting or test");
    brute_force_arithmetic_gadget<ALU_or_gadget<FieldT>, FieldT>(w,
                                                                 tinyram_opcode_OR,
                                                                 [] (tinyram_protoboard<FieldT> &pb,
                                                                     pb_variable_array<FieldT> &opcode_indicators,
                                                                     word_variable_gadget<FieldT> &desval,
                                                                     word_variable_gadget<FieldT> &arg1val,
                                                                     word_variable_gadget<FieldT> &arg2val,
                                                                     pb_variable<FieldT> &flag,
                                                                     pb_variable<FieldT> &result,
                                                                     pb_variable<FieldT> &result_flag) ->
                                                                 ALU_or_gadget<FieldT>* {
                                                                     return new ALU_or_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, "ALU_or_gadget");
                                                                 },
                                                                 [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return x | y; },
                                                                 [w] (size_t des, bool f, size_t x, size_t y) -> bool { return (x | y) == 0; });
    print_time("or tests successful");
}

/* xor */
template<typename FieldT>
void ALU_xor_gadget<FieldT>::generate_r1cs_constraints()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        /* a = b ^ c <=> a = b + c - 2*b*c, (2*b)*c = b+c - a */
        this->pb.add_r1cs_constraint(
            r1cs_constraint<FieldT>(
                { this->arg1val.bits[i] * 2},
                { this->arg2val.bits[i] },
                { this->arg1val.bits[i], this->arg2val.bits[i], this->res_word[i] * (-1) }),
            FMT(this->annotation_prefix, " res_word_%zu", i));
    }

    /* generate result */
    pack_result->generate_r1cs_constraints(false);
    not_all_zeros->generate_r1cs_constraints();

    /* result_flag = 1 - not_all_zeros = result is 0^w */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, this->not_all_zeros_result * (-1) },
            { this->result_flag }),
        FMT(this->annotation_prefix, " result_flag"));
}

template<typename FieldT>
void ALU_xor_gadget<FieldT>::generate_r1cs_witness()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        bool b1 = this->pb.val(this->arg1val.bits[i]) == FieldT::one();
        bool b2 = this->pb.val(this->arg2val.bits[i]) == FieldT::one();

        this->pb.val(this->res_word[i]) = (b1 ^ b2 ? FieldT::one() : FieldT::zero());
    }

    pack_result->generate_r1cs_witness_from_bits();
    not_all_zeros->generate_r1cs_witness();
    this->pb.val(this->result_flag) = FieldT::one() - this->pb.val(this->not_all_zeros_result);
}

template<typename FieldT>
void test_ALU_xor_gadget(const size_t w)
{
    print_time("starting xor test");
    brute_force_arithmetic_gadget<ALU_xor_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_XOR,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_xor_gadget<FieldT>* {
                                                                      return new ALU_xor_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, "ALU_xor_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return x ^ y; },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return (x ^ y) == 0; });
    print_time("xor tests successful");
}

/* not */
template<typename FieldT>
void ALU_not_gadget<FieldT>::generate_r1cs_constraints()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        this->pb.add_r1cs_constraint(
            r1cs_constraint<FieldT>(
                { ONE },
                { ONE, this->arg2val.bits[i] * (-1) },
                { this->res_word[i] }),
            FMT(this->annotation_prefix, " res_word_%zu", i));
    }

    /* generate result */
    pack_result->generate_r1cs_constraints(false);
    not_all_zeros->generate_r1cs_constraints();

    /* result_flag = 1 - not_all_zeros = result is 0^w */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, this->not_all_zeros_result * (-1) },
            { this->result_flag }),
        FMT(this->annotation_prefix, " result_flag"));
}

template<typename FieldT>
void ALU_not_gadget<FieldT>::generate_r1cs_witness()
{
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        bool b2 = this->pb.val(this->arg2val.bits[i]) == FieldT::one();

        this->pb.val(this->res_word[i]) = (!b2 ? FieldT::one() : FieldT::zero());
    }

    pack_result->generate_r1cs_witness_from_bits();
    not_all_zeros->generate_r1cs_witness();
    this->pb.val(this->result_flag) = FieldT::one() - this->pb.val(this->not_all_zeros_result);
}

template<typename FieldT>
void test_ALU_not_gadget(const size_t w)
{
    print_time("starting not test");
    brute_force_arithmetic_gadget<ALU_not_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_NOT,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_not_gadget<FieldT>* {
                                                                      return new ALU_not_gadget<FieldT>(pb, opcode_indicators,desval, arg1val, arg2val, flag, result, result_flag, "ALU_not_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return (1ul<<w)-1-y; },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return ((1ul<<w)-1-y) == 0; });
    print_time("not tests successful");
}

/* add */
template<typename FieldT>
void ALU_add_gadget<FieldT>::generate_r1cs_constraints()
{
    /* addition_result = 1 * (arg1val + arg2val) */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->arg1val.packed, this->arg2val.packed },
            { this->addition_result }),
        FMT(this->annotation_prefix, " addition_result"));

    /* unpack into bits */
    unpack_addition->generate_r1cs_constraints(true);

    /* generate result */
    pack_result->generate_r1cs_constraints(false);
}

template<typename FieldT>
void ALU_add_gadget<FieldT>::generate_r1cs_witness()
{
    this->pb.val(addition_result) = this->pb.val(this->arg1val.packed) + this->pb.val(this->arg2val.packed);
    unpack_addition->generate_r1cs_witness_from_packed();
    pack_result->generate_r1cs_witness_from_bits();
}

template<typename FieldT>
void test_ALU_add_gadget(const size_t w)
{
    print_time("starting add test");
    brute_force_arithmetic_gadget<ALU_add_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_ADD,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_add_gadget<FieldT>* {
                                                                      return new ALU_add_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, "ALU_add_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return (x+y) % (1ul<<w); },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return (x+y) >= (1ul<<w); });
    print_time("add tests successful");
}

/* sub */
template<typename FieldT>
void ALU_sub_gadget<FieldT>::generate_r1cs_constraints()
{
    /* intermediate_result = 2^w + (arg1val - arg2val) */
    FieldT twoi = FieldT::one();

    linear_combination<FieldT> a, b, c;

    a.add_term(0, 1);
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        twoi = twoi + twoi;
    }
    b.add_term(0, twoi);
    b.add_term(this->arg1val.packed, 1);
    b.add_term(this->arg2val.packed, -1);
    c.add_term(intermediate_result, 1);

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a, b, c), FMT(this->annotation_prefix, " main_constraint"));

    /* unpack into bits */
    unpack_intermediate->generate_r1cs_constraints(true);

    /* generate result */
    pack_result->generate_r1cs_constraints(false);
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, this->negated_flag * (-1) },
            { this->result_flag }),
        FMT(this->annotation_prefix, " result_flag"));
}

template<typename FieldT>
void ALU_sub_gadget<FieldT>::generate_r1cs_witness()
{
    FieldT twoi = FieldT::one();
    for (size_t i = 0; i < this->pb.ap.w; ++i)
    {
        twoi = twoi + twoi;
    }

    this->pb.val(intermediate_result) = twoi + this->pb.val(this->arg1val.packed) - this->pb.val(this->arg2val.packed);
    unpack_intermediate->generate_r1cs_witness_from_packed();
    pack_result->generate_r1cs_witness_from_bits();
    this->pb.val(this->result_flag) = FieldT::one() - this->pb.val(this->negated_flag);
}

template<typename FieldT>
void test_ALU_sub_gadget(const size_t w)
{
    print_time("starting sub test");
    brute_force_arithmetic_gadget<ALU_sub_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_SUB,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_sub_gadget<FieldT>* {
                                                                      return new ALU_sub_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, "ALU_sub_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t {
                                                                      const size_t unsigned_result = ((1ul<<w) + x - y) % (1ul<<w);
                                                                      return unsigned_result;
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                      const size_t msb = ((1ul<<w) + x - y) >> w;
                                                                      return (msb == 0);
                                                                  });
    print_time("sub tests successful");
}

/* mov */
template<typename FieldT>
void ALU_mov_gadget<FieldT>::generate_r1cs_constraints()
{
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->arg2val.packed },
            { this->result }),
        FMT(this->annotation_prefix, " mov_result"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->flag },
            { this->result_flag }),
        FMT(this->annotation_prefix, " mov_result_flag"));
}

template<typename FieldT>
void ALU_mov_gadget<FieldT>::generate_r1cs_witness()
{
    this->pb.val(this->result) = this->pb.val(this->arg2val.packed);
    this->pb.val(this->result_flag) = this->pb.val(this->flag);
}

template<typename FieldT>
void test_ALU_mov_gadget(const size_t w)
{
    print_time("starting mov test");
    brute_force_arithmetic_gadget<ALU_mov_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_MOV,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_mov_gadget<FieldT>* {
                                                                      return new ALU_mov_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, "ALU_mov_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return y; },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return f; });
    print_time("mov tests successful");
}

/* cmov */
template<typename FieldT>
void ALU_cmov_gadget<FieldT>::generate_r1cs_constraints()
{
    /*
      flag1 * arg2val + (1-flag1) * desval = result
      flag1 * (arg2val - desval) = result - desval
    */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { this->flag },
            { this->arg2val.packed, this->desval.packed * (-1) },
            { this->result, this->desval.packed * (-1) }),
        FMT(this->annotation_prefix, " cmov_result"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->flag },
            { this->result_flag }),
        FMT(this->annotation_prefix, " cmov_result_flag"));
}

template<typename FieldT>
void ALU_cmov_gadget<FieldT>::generate_r1cs_witness()
{
    this->pb.val(this->result) = ((this->pb.val(this->flag) == FieldT::one()) ?
                                  this->pb.val(this->arg2val.packed) :
                                  this->pb.val(this->desval.packed));
    this->pb.val(this->result_flag) = this->pb.val(this->flag);
}

template<typename FieldT>
void test_ALU_cmov_gadget(const size_t w)
{
    print_time("starting cmov test");
    brute_force_arithmetic_gadget<ALU_cmov_gadget<FieldT>, FieldT>(w,
                                                                   tinyram_opcode_CMOV,
                                                                   [] (tinyram_protoboard<FieldT> &pb,
                                                                       pb_variable_array<FieldT> &opcode_indicators,
                                                                       word_variable_gadget<FieldT> &desval,
                                                                       word_variable_gadget<FieldT> &arg1val,
                                                                       word_variable_gadget<FieldT> &arg2val,
                                                                       pb_variable<FieldT> &flag,
                                                                       pb_variable<FieldT> &result,
                                                                       pb_variable<FieldT> &result_flag) ->
                                                                   ALU_cmov_gadget<FieldT>* {
                                                                       return new ALU_cmov_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, "ALU_cmov_gadget");
                                                                   },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return f ? y : des; },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> bool { return f; });
    print_time("cmov tests successful");
}

/* unsigned comparison */
template<typename FieldT>
void ALU_cmp_gadget<FieldT>::generate_r1cs_constraints()
{
    comparator.generate_r1cs_constraints();
    /*
      cmpe = cmpae * (1-cmpa)
    */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { cmpae_result_flag },
            { ONE, cmpa_result_flag * (-1) },
            { cmpe_result_flag }),
        FMT(this->annotation_prefix, " cmpa_result_flag"));

    /* copy over results */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->desval.packed },
            { cmpe_result }),
        FMT(this->annotation_prefix, " cmpe_result"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->desval.packed },
            { cmpa_result }),
        FMT(this->annotation_prefix, " cmpa_result"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->desval.packed },
            { cmpae_result }),
        FMT(this->annotation_prefix, " cmpae_result"));
}

template<typename FieldT>
void ALU_cmp_gadget<FieldT>::generate_r1cs_witness()
{
    comparator.generate_r1cs_witness();

    this->pb.val(cmpe_result) = this->pb.val(this->desval.packed);
    this->pb.val(cmpa_result) = this->pb.val(this->desval.packed);
    this->pb.val(cmpae_result) = this->pb.val(this->desval.packed);

    this->pb.val(cmpe_result_flag) = ((this->pb.val(cmpae_result_flag) == FieldT::one()) &&
                                      (this->pb.val(cmpa_result_flag) == FieldT::zero()) ?
                                      FieldT::one() :
                                      FieldT::zero());
}

template<typename FieldT>
void test_ALU_cmpe_gadget(const size_t w)
{
    print_time("starting cmpe test");
    brute_force_arithmetic_gadget<ALU_cmp_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_CMPE,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_cmp_gadget<FieldT>* {
                                                                      pb_variable<FieldT> cmpa_result; cmpa_result.allocate(pb, "cmpa_result");
                                                                      pb_variable<FieldT> cmpa_result_flag; cmpa_result_flag.allocate(pb, "cmpa_result_flag");
                                                                      pb_variable<FieldT> cmpae_result; cmpae_result.allocate(pb, "cmpae_result");
                                                                      pb_variable<FieldT> cmpae_result_flag; cmpae_result_flag.allocate(pb, "cmpae_result_flag");
                                                                      return new ALU_cmp_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                        result, result_flag,
                                                                                                        cmpa_result, cmpa_result_flag,
                                                                                                        cmpae_result, cmpae_result_flag, "ALU_cmp_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return des; },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return x == y; });
    print_time("cmpe tests successful");
}

template<typename FieldT>
void test_ALU_cmpa_gadget(const size_t w)
{
    print_time("starting cmpa test");
    brute_force_arithmetic_gadget<ALU_cmp_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_CMPA,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_cmp_gadget<FieldT>* {
                                                                      pb_variable<FieldT> cmpe_result; cmpe_result.allocate(pb, "cmpe_result");
                                                                      pb_variable<FieldT> cmpe_result_flag; cmpe_result_flag.allocate(pb, "cmpe_result_flag");
                                                                      pb_variable<FieldT> cmpae_result; cmpae_result.allocate(pb, "cmpae_result");
                                                                      pb_variable<FieldT> cmpae_result_flag; cmpae_result_flag.allocate(pb, "cmpae_result_flag");
                                                                      return new ALU_cmp_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                        cmpe_result, cmpe_result_flag,
                                                                                                        result, result_flag,
                                                                                                        cmpae_result, cmpae_result_flag, "ALU_cmp_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return des; },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return x > y; });
    print_time("cmpa tests successful");
}

template<typename FieldT>
void test_ALU_cmpae_gadget(const size_t w)
{
    print_time("starting cmpae test");
    brute_force_arithmetic_gadget<ALU_cmp_gadget<FieldT>, FieldT>(w,
                                                                  tinyram_opcode_CMPAE,
                                                                  [] (tinyram_protoboard<FieldT> &pb,
                                                                      pb_variable_array<FieldT> &opcode_indicators,
                                                                      word_variable_gadget<FieldT> &desval,
                                                                      word_variable_gadget<FieldT> &arg1val,
                                                                      word_variable_gadget<FieldT> &arg2val,
                                                                      pb_variable<FieldT> &flag,
                                                                      pb_variable<FieldT> &result,
                                                                      pb_variable<FieldT> &result_flag) ->
                                                                  ALU_cmp_gadget<FieldT>* {
                                                                      pb_variable<FieldT> cmpe_result; cmpe_result.allocate(pb, "cmpe_result");
                                                                      pb_variable<FieldT> cmpe_result_flag; cmpe_result_flag.allocate(pb, "cmpe_result_flag");
                                                                      pb_variable<FieldT> cmpa_result; cmpa_result.allocate(pb, "cmpa_result");
                                                                      pb_variable<FieldT> cmpa_result_flag; cmpa_result_flag.allocate(pb, "cmpa_result_flag");
                                                                      return new ALU_cmp_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                        cmpe_result, cmpe_result_flag,
                                                                                                        cmpa_result, cmpa_result_flag,
                                                                                                        result, result_flag, "ALU_cmp_gadget");
                                                                  },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return des; },
                                                                  [w] (size_t des, bool f, size_t x, size_t y) -> bool { return x >= y; });
    print_time("cmpae tests successful");
}

/* signed comparison */
template<typename FieldT>
void ALU_cmps_gadget<FieldT>::generate_r1cs_constraints()
{
    /* negate sign bits */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, this->arg1val.bits[this->pb.ap.w-1] * (-1) },
            { negated_arg1val_sign }),
        FMT(this->annotation_prefix, " negated_arg1val_sign"));
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, this->arg2val.bits[this->pb.ap.w-1] * (-1) },
            { negated_arg2val_sign }),
        FMT(this->annotation_prefix, " negated_arg2val_sign"));

    /* pack */
    pack_modified_arg1->generate_r1cs_constraints(false);
    pack_modified_arg2->generate_r1cs_constraints(false);

    /* compare */
    comparator->generate_r1cs_constraints();

    /* copy over results */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->desval.packed },
            { cmpg_result }),
        FMT(this->annotation_prefix, " cmpg_result"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->desval.packed },
            { cmpge_result }),
        FMT(this->annotation_prefix, " cmpge_result"));
}

template<typename FieldT>
void ALU_cmps_gadget<FieldT>::generate_r1cs_witness()
{
    /* negate sign bits */
    this->pb.val(negated_arg1val_sign) = FieldT::one() - this->pb.val(this->arg1val.bits[this->pb.ap.w-1]);
    this->pb.val(negated_arg2val_sign) = FieldT::one() - this->pb.val(this->arg2val.bits[this->pb.ap.w-1]);

    /* pack */
    pack_modified_arg1->generate_r1cs_witness_from_bits();
    pack_modified_arg2->generate_r1cs_witness_from_bits();

    /* produce result */
    comparator->generate_r1cs_witness();

    this->pb.val(cmpg_result) = this->pb.val(this->desval.packed);
    this->pb.val(cmpge_result) = this->pb.val(this->desval.packed);
}

template<typename FieldT>
void test_ALU_cmpg_gadget(const size_t w)
{
    print_time("starting cmpg test");
    brute_force_arithmetic_gadget<ALU_cmps_gadget<FieldT>, FieldT>(w,
                                                                   tinyram_opcode_CMPG,
                                                                   [] (tinyram_protoboard<FieldT> &pb,
                                                                       pb_variable_array<FieldT> &opcode_indicators,
                                                                       word_variable_gadget<FieldT> &desval,
                                                                       word_variable_gadget<FieldT> &arg1val,
                                                                       word_variable_gadget<FieldT> &arg2val,
                                                                       pb_variable<FieldT> &flag,
                                                                       pb_variable<FieldT> &result,
                                                                       pb_variable<FieldT> &result_flag) ->
                                                                   ALU_cmps_gadget<FieldT>* {
                                                                       pb_variable<FieldT> cmpge_result; cmpge_result.allocate(pb, "cmpge_result");
                                                                       pb_variable<FieldT> cmpge_result_flag; cmpge_result_flag.allocate(pb, "cmpge_result_flag");
                                                                       return new ALU_cmps_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                          result, result_flag,
                                                                                                          cmpge_result, cmpge_result_flag, "ALU_cmps_gadget");
                                                                   },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return des; },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                       return (from_twos_complement(x, w) >
                                                                               from_twos_complement(y, w));
                                                                   });
    print_time("cmpg tests successful");
}

template<typename FieldT>
void test_ALU_cmpge_gadget(const size_t w)
{
    print_time("starting cmpge test");
    brute_force_arithmetic_gadget<ALU_cmps_gadget<FieldT>, FieldT>(w,
                                                                   tinyram_opcode_CMPGE,
                                                                   [] (tinyram_protoboard<FieldT> &pb,
                                                                       pb_variable_array<FieldT> &opcode_indicators,
                                                                       word_variable_gadget<FieldT> &desval,
                                                                       word_variable_gadget<FieldT> &arg1val,
                                                                       word_variable_gadget<FieldT> &arg2val,
                                                                       pb_variable<FieldT> &flag,
                                                                       pb_variable<FieldT> &result,
                                                                       pb_variable<FieldT> &result_flag) ->
                                                                   ALU_cmps_gadget<FieldT>* {
                                                                       pb_variable<FieldT> cmpg_result; cmpg_result.allocate(pb, "cmpg_result");
                                                                       pb_variable<FieldT> cmpg_result_flag; cmpg_result_flag.allocate(pb, "cmpg_result_flag");
                                                                       return new ALU_cmps_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                          cmpg_result, cmpg_result_flag,
                                                                                                          result, result_flag, "ALU_cmps_gadget");
                                                                   },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return des; },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                       return (from_twos_complement(x, w) >=
                                                                               from_twos_complement(y, w));
                                                                   });
    print_time("cmpge tests successful");
}

template<typename FieldT>
void ALU_umul_gadget<FieldT>::generate_r1cs_constraints()
{
    /* do multiplication */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { this->arg1val.packed },
            { this->arg2val.packed },
            { mul_result.packed }),
        FMT(this->annotation_prefix, " main_constraint"));
    mul_result.generate_r1cs_constraints(true);

    /* pack result */
    pack_mull_result->generate_r1cs_constraints(false);
    pack_umulh_result->generate_r1cs_constraints(false);

    /* compute flag */
    compute_flag->generate_r1cs_constraints();

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->result_flag },
            { mull_flag }),
        FMT(this->annotation_prefix, " mull_flag"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->result_flag },
            { umulh_flag }),
        FMT(this->annotation_prefix, " umulh_flag"));
}

template<typename FieldT>
void ALU_umul_gadget<FieldT>::generate_r1cs_witness()
{
    /* do multiplication */
    this->pb.val(mul_result.packed) = this->pb.val(this->arg1val.packed) * this->pb.val(this->arg2val.packed);
    mul_result.generate_r1cs_witness_from_packed();

    /* pack result */
    pack_mull_result->generate_r1cs_witness_from_bits();
    pack_umulh_result->generate_r1cs_witness_from_bits();

    /* compute flag */
    compute_flag->generate_r1cs_witness();

    this->pb.val(mull_flag) = this->pb.val(this->result_flag);
    this->pb.val(umulh_flag) = this->pb.val(this->result_flag);
}

template<typename FieldT>
void test_ALU_mull_gadget(const size_t w)
{
    print_time("starting mull test");
    brute_force_arithmetic_gadget<ALU_umul_gadget<FieldT>, FieldT>(w,
                                                                   tinyram_opcode_MULL,
                                                                   [] (tinyram_protoboard<FieldT> &pb,
                                                                       pb_variable_array<FieldT> &opcode_indicators,
                                                                       word_variable_gadget<FieldT> &desval,
                                                                       word_variable_gadget<FieldT> &arg1val,
                                                                       word_variable_gadget<FieldT> &arg2val,
                                                                       pb_variable<FieldT> &flag,
                                                                       pb_variable<FieldT> &result,
                                                                       pb_variable<FieldT> &result_flag) ->
                                                                   ALU_umul_gadget<FieldT>* {
                                                                       pb_variable<FieldT> umulh_result; umulh_result.allocate(pb, "umulh_result");
                                                                       pb_variable<FieldT> umulh_flag; umulh_flag.allocate(pb, "umulh_flag");
                                                                       return new ALU_umul_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                          result, result_flag,
                                                                                                          umulh_result, umulh_flag,
                                                                                                          "ALU_umul_gadget");
                                                                   },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return (x*y) % (1ul<<w); },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                       return ((x*y) >> w) != 0;
                                                                   });
    print_time("mull tests successful");
}

template<typename FieldT>
void test_ALU_umulh_gadget(const size_t w)
{
    print_time("starting umulh test");
    brute_force_arithmetic_gadget<ALU_umul_gadget<FieldT>, FieldT>(w,
                                                                   tinyram_opcode_UMULH,
                                                                   [] (tinyram_protoboard<FieldT> &pb,
                                                                       pb_variable_array<FieldT> &opcode_indicators,
                                                                       word_variable_gadget<FieldT> &desval,
                                                                       word_variable_gadget<FieldT> &arg1val,
                                                                       word_variable_gadget<FieldT> &arg2val,
                                                                       pb_variable<FieldT> &flag,
                                                                       pb_variable<FieldT> &result,
                                                                       pb_variable<FieldT> &result_flag) ->
                                                                   ALU_umul_gadget<FieldT>* {
                                                                       pb_variable<FieldT> mull_result; mull_result.allocate(pb, "mull_result");
                                                                       pb_variable<FieldT> mull_flag; mull_flag.allocate(pb, "mull_flag");
                                                                       return new ALU_umul_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                          mull_result, mull_flag,
                                                                                                          result, result_flag,
                                                                                                          "ALU_umul_gadget");
                                                                   },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> size_t { return (x*y) >> w; },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                       return ((x*y) >> w) != 0;
                                                                   });
    print_time("umulh tests successful");
}

template<typename FieldT>
void ALU_smul_gadget<FieldT>::generate_r1cs_constraints()
{
    /* do multiplication */
    /*
      from two's complement: (packed - 2^w * bits[w-1])
      to two's complement: lower order bits of 2^{2w} + result_of_*
    */

    linear_combination<FieldT> a, b, c;
    a.add_term(this->arg1val.packed, 1);
    a.add_term(this->arg1val.bits[this->pb.ap.w-1], -(FieldT(2)^this->pb.ap.w));
    b.add_term(this->arg2val.packed, 1);
    b.add_term(this->arg2val.bits[this->pb.ap.w-1], -(FieldT(2)^this->pb.ap.w));
    c.add_term(mul_result.packed, 1);
    c.add_term(ONE, -(FieldT(2)^(2*this->pb.ap.w)));
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a, b, c), FMT(this->annotation_prefix, " main_constraint"));

    mul_result.generate_r1cs_constraints(true);

    /* pack result */
    pack_smulh_result->generate_r1cs_constraints(false);

    /* compute flag */
    pack_top->generate_r1cs_constraints(false);

    /*
      the gadgets below are FieldT specific:
      I * X = (1-R)
      R * X = 0

      if X = 0 then R = 1
      if X != 0 then R = 0 and I = X^{-1}
    */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { is_top_empty_aux },
            { top },
            { ONE, is_top_empty * (-1) }),
        FMT(this->annotation_prefix, " I*X=1-R (is_top_empty)"));
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { is_top_empty },
            { top },
            { ONE * 0 }),
        FMT(this->annotation_prefix, " R*X=0 (is_top_full)"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { is_top_full_aux },
            { top, ONE * (1l-(1ul<<(this->pb.ap.w+1))) },
            { ONE, is_top_full * (-1) }),
        FMT(this->annotation_prefix, " I*X=1-R (is_top_full)"));
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { is_top_full },
            { top, ONE * (1l-(1ul<<(this->pb.ap.w+1))) },
            { ONE * 0 }),
        FMT(this->annotation_prefix, " R*X=0 (is_top_full)"));

    /* smulh_flag = 1 - (is_top_full + is_top_empty) */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { ONE, is_top_full * (-1), is_top_empty * (-1) },
            { smulh_flag }),
        FMT(this->annotation_prefix, " smulh_flag"));
}

template<typename FieldT>
void ALU_smul_gadget<FieldT>::generate_r1cs_witness()
{
    /* do multiplication */
    /*
      from two's complement: (packed - 2^w * bits[w-1])
      to two's complement: lower order bits of (2^{2w} + result_of_mul)
    */
    this->pb.val(mul_result.packed) =
        (this->pb.val(this->arg1val.packed) - (this->pb.val(this->arg1val.bits[this->pb.ap.w-1])*(FieldT(2)^this->pb.ap.w))) *
        (this->pb.val(this->arg2val.packed) - (this->pb.val(this->arg2val.bits[this->pb.ap.w-1])*(FieldT(2)^this->pb.ap.w))) +
        (FieldT(2)^(2*this->pb.ap.w));

    mul_result.generate_r1cs_witness_from_packed();

    /* pack result */
    pack_smulh_result->generate_r1cs_witness_from_bits();

    /* compute flag */
    pack_top->generate_r1cs_witness_from_bits();
    size_t topval = this->pb.val(top).as_ulong();

    if (topval == 0)
    {
        this->pb.val(is_top_empty) = FieldT::one();
        this->pb.val(is_top_empty_aux) = FieldT::zero();
    }
    else
    {
        this->pb.val(is_top_empty) = FieldT::zero();
        this->pb.val(is_top_empty_aux) = this->pb.val(top).inverse();
    }

    if (topval == ((1ul<<(this->pb.ap.w+1))-1))
    {
        this->pb.val(is_top_full) = FieldT::one();
        this->pb.val(is_top_full_aux) = FieldT::zero();
    }
    else
    {
        this->pb.val(is_top_full) = FieldT::zero();
        this->pb.val(is_top_full_aux) = (this->pb.val(top)-FieldT((1ul<<(this->pb.ap.w+1))-1)).inverse();
    }

    /* smulh_flag = 1 - (is_top_full + is_top_empty) */
    this->pb.val(smulh_flag) = FieldT::one() - (this->pb.val(is_top_full) + this->pb.val(is_top_empty));
}

template<typename FieldT>
void test_ALU_smulh_gadget(const size_t w)
{
    print_time("starting smulh test");
    brute_force_arithmetic_gadget<ALU_smul_gadget<FieldT>, FieldT>(w,
                                                                   tinyram_opcode_SMULH,
                                                                   [] (tinyram_protoboard<FieldT> &pb,
                                                                       pb_variable_array<FieldT> &opcode_indicators,
                                                                       word_variable_gadget<FieldT> &desval,
                                                                       word_variable_gadget<FieldT> &arg1val,
                                                                       word_variable_gadget<FieldT> &arg2val,
                                                                       pb_variable<FieldT> &flag,
                                                                       pb_variable<FieldT> &result,
                                                                       pb_variable<FieldT> &result_flag) ->
                                                                   ALU_smul_gadget<FieldT>* {
                                                                       return new ALU_smul_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                          result, result_flag,
                                                                                                          "ALU_smul_gadget");
                                                                   },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> size_t {
                                                                       const size_t res = to_twos_complement((from_twos_complement(x, w) * from_twos_complement(y, w)), 2*w);
                                                                       return res >> w;
                                                                   },
                                                                   [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                       const int res = from_twos_complement(x, w) * from_twos_complement(y, w);
                                                                       const int truncated_res = from_twos_complement(to_twos_complement(res, 2*w) & ((1ul<<w)-1), w);
                                                                       return (res != truncated_res);
                                                                   });
    print_time("smulh tests successful");
}

template<typename FieldT>
void ALU_divmod_gadget<FieldT>::generate_r1cs_constraints()
{
    /* B_inv * B = B_nonzero */
    linear_combination<FieldT> a1, b1, c1;
    a1.add_term(B_inv, 1);
    b1.add_term(this->arg2val.packed, 1);
    c1.add_term(B_nonzero, 1);

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a1, b1, c1), FMT(this->annotation_prefix, " B_inv*B=B_nonzero"));

    /* (1-B_nonzero) * B = 0 */
    linear_combination<FieldT> a2, b2, c2;
    a2.add_term(ONE, 1);
    a2.add_term(B_nonzero, -1);
    b2.add_term(this->arg2val.packed, 1);
    c2.add_term(ONE, 0);

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a2, b2, c2), FMT(this->annotation_prefix, " (1-B_nonzero)*B=0"));

    /* B * q + r = A_aux = A * B_nonzero */
    linear_combination<FieldT> a3, b3, c3;
    a3.add_term(this->arg2val.packed, 1);
    b3.add_term(udiv_result, 1);
    c3.add_term(A_aux, 1);
    c3.add_term(umod_result, -1);

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a3, b3, c3), FMT(this->annotation_prefix, " B*q+r=A_aux"));

    linear_combination<FieldT> a4, b4, c4;
    a4.add_term(this->arg1val.packed, 1);
    b4.add_term(B_nonzero, 1);
    c4.add_term(A_aux, 1);

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a4, b4, c4), FMT(this->annotation_prefix, " A_aux=A*B_nonzero"));

    /* q * (1-B_nonzero) = 0 */
    linear_combination<FieldT> a5, b5, c5;
    a5.add_term(udiv_result, 1);
    b5.add_term(ONE, 1);
    b5.add_term(B_nonzero, -1);
    c5.add_term(ONE, 0);

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a5, b5, c5), FMT(this->annotation_prefix, " q*B_nonzero=0"));

    /* A<B_gadget<FieldT>(B, r, less=B_nonzero, leq=ONE) */
    r_less_B->generate_r1cs_constraints();
}

template<typename FieldT>
void ALU_divmod_gadget<FieldT>::generate_r1cs_witness()
{
    if (this->pb.val(this->arg2val.packed) == FieldT::zero())
    {
        this->pb.val(B_inv) = FieldT::zero();
        this->pb.val(B_nonzero) = FieldT::zero();

        this->pb.val(A_aux) = FieldT::zero();

        this->pb.val(udiv_result) = FieldT::zero();
        this->pb.val(umod_result) = FieldT::zero();

        this->pb.val(udiv_flag) = FieldT::one();
        this->pb.val(umod_flag) = FieldT::one();
    }
    else
    {
        this->pb.val(B_inv) = this->pb.val(this->arg2val.packed).inverse();
        this->pb.val(B_nonzero) = FieldT::one();

        const size_t A = this->pb.val(this->arg1val.packed).as_ulong();
        const size_t B = this->pb.val(this->arg2val.packed).as_ulong();

        this->pb.val(A_aux) = this->pb.val(this->arg1val.packed);

        this->pb.val(udiv_result) = FieldT(A / B);
        this->pb.val(umod_result) = FieldT(A % B);

        this->pb.val(udiv_flag) = FieldT::zero();
        this->pb.val(umod_flag) = FieldT::zero();
    }

    r_less_B->generate_r1cs_witness();
}

template<typename FieldT>
void test_ALU_udiv_gadget(const size_t w)
{
    print_time("starting udiv test");
    brute_force_arithmetic_gadget<ALU_divmod_gadget<FieldT>, FieldT>(w,
                                                                     tinyram_opcode_UDIV,
                                                                     [] (tinyram_protoboard<FieldT> &pb,
                                                                         pb_variable_array<FieldT> &opcode_indicators,
                                                                         word_variable_gadget<FieldT> &desval,
                                                                         word_variable_gadget<FieldT> &arg1val,
                                                                         word_variable_gadget<FieldT> &arg2val,
                                                                         pb_variable<FieldT> &flag,
                                                                         pb_variable<FieldT> &result,
                                                                         pb_variable<FieldT> &result_flag) ->
                                                                     ALU_divmod_gadget<FieldT>* {
                                                                         pb_variable<FieldT> umod_result; umod_result.allocate(pb, "umod_result");
                                                                         pb_variable<FieldT> umod_flag; umod_flag.allocate(pb, "umod_flag");
                                                                         return new ALU_divmod_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                              result, result_flag,
                                                                                                              umod_result, umod_flag,
                                                                                                              "ALU_divmod_gadget");
                                                                     },
                                                                     [w] (size_t des, bool f, size_t x, size_t y) -> size_t {
                                                                         return (y == 0 ? 0 : x / y);
                                                                     },
                                                                     [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                         return (y == 0);
                                                                     });
    print_time("udiv tests successful");
}

template<typename FieldT>
void test_ALU_umod_gadget(const size_t w)
{
    print_time("starting umod test");
    brute_force_arithmetic_gadget<ALU_divmod_gadget<FieldT>, FieldT>(w,
                                                                     tinyram_opcode_UMOD,
                                                                     [] (tinyram_protoboard<FieldT> &pb,
                                                                         pb_variable_array<FieldT> &opcode_indicators,
                                                                         word_variable_gadget<FieldT> &desval,
                                                                         word_variable_gadget<FieldT> &arg1val,
                                                                         word_variable_gadget<FieldT> &arg2val,
                                                                         pb_variable<FieldT> &flag,
                                                                         pb_variable<FieldT> &result,
                                                                         pb_variable<FieldT> &result_flag) ->
                                                                     ALU_divmod_gadget<FieldT>* {
                                                                         pb_variable<FieldT> udiv_result; udiv_result.allocate(pb, "udiv_result");
                                                                         pb_variable<FieldT> udiv_flag; udiv_flag.allocate(pb, "udiv_flag");
                                                                         return new ALU_divmod_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                              udiv_result, udiv_flag,
                                                                                                              result, result_flag,
                                                                                                              "ALU_divmod_gadget");
                                                                     },
                                                                     [w] (size_t des, bool f, size_t x, size_t y) -> size_t {
                                                                         return (y == 0 ? 0 : x % y);
                                                                     },
                                                                     [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                         return (y == 0);
                                                                     });
    print_time("umod tests successful");
}

template<typename FieldT>
void ALU_shr_shl_gadget<FieldT>::generate_r1cs_constraints()
{
    /*
      select the input for barrel shifter:

      r = arg1val * opcode_indicators[SHR] + reverse(arg1val) * (1-opcode_indicators[SHR])
      r - reverse(arg1val) = (arg1val - reverse(arg1val)) * opcode_indicators[SHR]
    */
    pack_reversed_input->generate_r1cs_constraints(false);

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { this->arg1val.packed, reversed_input * (-1) },
            { this->opcode_indicators[tinyram_opcode_SHR] },
            { barrel_right_internal[0], reversed_input * (-1) }),
        FMT(this->annotation_prefix, " select_arg1val_or_reversed"));

    /*
      do logw iterations of barrel shifts
    */
    for (size_t i = 0; i < logw; ++i)
    {
        /* assert that shifted out part is bits */
        for (size_t j = 0; j < 1ul<<i; ++j)
        {
            generate_boolean_r1cs_constraint<FieldT>(this->pb, shifted_out_bits[i][j], FMT(this->annotation_prefix, " shifted_out_bits_%zu_%zu", i, j));
        }

        /*
          add main shifting constraint


          old_result =
          (shifted_result * 2^(i+1) + shifted_out_part) * need_to_shift +
          (shfited_result) * (1-need_to_shift)

          old_result - shifted_result = (shifted_result * (2^(i+1) - 1) + shifted_out_part) * need_to_shift
        */
        linear_combination<FieldT> a, b, c;

        a.add_term(barrel_right_internal[i+1], (FieldT(2)^(i+1)) - FieldT::one());
        for (size_t j = 0; j < 1ul<<i; ++j)
        {
            a.add_term(shifted_out_bits[i][j], (FieldT(2)^j));
        }

        b.add_term(this->arg2val.bits[i], 1);

        c.add_term(barrel_right_internal[i], 1);
        c.add_term(barrel_right_internal[i+1], -1);

        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(a, b, c), FMT(this->annotation_prefix, " barrel_shift_%zu", i));
    }

    /*
      get result as the logw iterations or zero if shift was oversized

      result = (1-is_oversize_shift) * barrel_right_internal[logw]
    */
    check_oversize_shift->generate_r1cs_constraints();
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE, is_oversize_shift * (-1) },
            { barrel_right_internal[logw] },
            { this->result }),
        FMT(this->annotation_prefix, " result"));

    /*
      get reversed result for SHL
    */
    unpack_result->generate_r1cs_constraints(true);
    pack_reversed_result->generate_r1cs_constraints(false);

    /*
      select the correct output:
      r = result * opcode_indicators[SHR] + reverse(result) * (1-opcode_indicators[SHR])
      r - reverse(result) = (result - reverse(result)) * opcode_indicators[SHR]
    */
    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { this->result, reversed_result * (-1) },
            { this->opcode_indicators[tinyram_opcode_SHR] },
            { shr_result, reversed_result * (-1) }),
        FMT(this->annotation_prefix, " shr_result"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { this->result, reversed_result * (-1) },
            { this->opcode_indicators[tinyram_opcode_SHR] },
            { shr_result, reversed_result * (-1) }),
        FMT(this->annotation_prefix, " shl_result"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->arg1val.bits[0] },
            { shr_flag }),
        FMT(this->annotation_prefix, " shr_flag"));

    this->pb.add_r1cs_constraint(
        r1cs_constraint<FieldT>(
            { ONE },
            { this->arg1val.bits[this->pb.ap.w-1] },
            { shl_flag }),
        FMT(this->annotation_prefix, " shl_flag"));
}

template<typename FieldT>
void ALU_shr_shl_gadget<FieldT>::generate_r1cs_witness()
{
    /* select the input for barrel shifter */
    pack_reversed_input->generate_r1cs_witness_from_bits();

    this->pb.val(barrel_right_internal[0]) =
        (this->pb.val(this->opcode_indicators[tinyram_opcode_SHR]) == FieldT::one() ?
         this->pb.val(this->arg1val.packed) : this->pb.val(reversed_input));

    /*
      do logw iterations of barrel shifts.

      old_result =
      (shifted_result * 2^i + shifted_out_part) * need_to_shift +
      (shfited_result) * (1-need_to_shift)
    */

    for (size_t i = 0; i < logw; ++i)
    {
        this->pb.val(barrel_right_internal[i+1]) =
            (this->pb.val(this->arg2val.bits[i]) == FieldT::zero()) ? this->pb.val(barrel_right_internal[i]) :
            FieldT(this->pb.val(barrel_right_internal[i]).as_ulong() >> (i+1));

        shifted_out_bits[i].fill_with_bits_of_ulong(this->pb, this->pb.val(barrel_right_internal[i]).as_ulong() % (2u<<i));
    }

    /*
      get result as the logw iterations or zero if shift was oversized

      result = (1-is_oversize_shift) * barrel_right_internal[logw]
    */
    check_oversize_shift->generate_r1cs_witness();
    this->pb.val(this->result) = (FieldT::one() - this->pb.val(is_oversize_shift)) * this->pb.val(barrel_right_internal[logw]);

    /*
      get reversed result for SHL
    */
    unpack_result->generate_r1cs_witness_from_packed();
    pack_reversed_result->generate_r1cs_witness_from_bits();

    /*
      select the correct output:
      r = result * opcode_indicators[SHR] + reverse(result) * (1-opcode_indicators[SHR])
      r - reverse(result) = (result - reverse(result)) * opcode_indicators[SHR]
    */
    this->pb.val(shr_result) = (this->pb.val(this->opcode_indicators[tinyram_opcode_SHR]) == FieldT::one()) ?
        this->pb.val(this->result) : this->pb.val(reversed_result);

    this->pb.val(shl_result) = this->pb.val(shr_result);
    this->pb.val(shr_flag) = this->pb.val(this->arg1val.bits[0]);
    this->pb.val(shl_flag) = this->pb.val(this->arg1val.bits[this->pb.ap.w-1]);
}

template<typename FieldT>
void test_ALU_shr_gadget(const size_t w)
{
    print_time("starting shr test");
    brute_force_arithmetic_gadget<ALU_shr_shl_gadget<FieldT>, FieldT>(w,
                                                                      tinyram_opcode_SHR,
                                                                      [] (tinyram_protoboard<FieldT> &pb,
                                                                          pb_variable_array<FieldT> &opcode_indicators,
                                                                          word_variable_gadget<FieldT> &desval,
                                                                          word_variable_gadget<FieldT> &arg1val,
                                                                          word_variable_gadget<FieldT> &arg2val,
                                                                          pb_variable<FieldT> &flag,
                                                                          pb_variable<FieldT> &result,
                                                                          pb_variable<FieldT> &result_flag) ->
                                                                      ALU_shr_shl_gadget<FieldT>* {
                                                                          pb_variable<FieldT> shl_result; shl_result.allocate(pb, "shl_result");
                                                                          pb_variable<FieldT> shl_flag; shl_flag.allocate(pb, "shl_flag");
                                                                          return new ALU_shr_shl_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                                result, result_flag,
                                                                                                                shl_result, shl_flag,
                                                                                                                "ALU_shr_shl_gadget");
                                                                      },
                                                                      [w] (size_t des, bool f, size_t x, size_t y) -> size_t {
                                                                          return (x >> y);
                                                                      },
                                                                      [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                          return (x & 1);
                                                                      });
    print_time("shr tests successful");
}

template<typename FieldT>
void test_ALU_shl_gadget(const size_t w)
{
    print_time("starting shl test");
    brute_force_arithmetic_gadget<ALU_shr_shl_gadget<FieldT>, FieldT>(w,
                                                                      tinyram_opcode_SHL,
                                                                      [] (tinyram_protoboard<FieldT> &pb,
                                                                          pb_variable_array<FieldT> &opcode_indicators,
                                                                          word_variable_gadget<FieldT> &desval,
                                                                          word_variable_gadget<FieldT> &arg1val,
                                                                          word_variable_gadget<FieldT> &arg2val,
                                                                          pb_variable<FieldT> &flag,
                                                                          pb_variable<FieldT> &result,
                                                                          pb_variable<FieldT> &result_flag) ->
                                                                      ALU_shr_shl_gadget<FieldT>* {
                                                                          pb_variable<FieldT> shr_result; shr_result.allocate(pb, "shr_result");
                                                                          pb_variable<FieldT> shr_flag; shr_flag.allocate(pb, "shr_flag");
                                                                          return new ALU_shr_shl_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag,
                                                                                                                shr_result, shr_flag,
                                                                                                                result, result_flag,
                                                                                                                "ALU_shr_shl_gadget");
                                                                      },
                                                                      [w] (size_t des, bool f, size_t x, size_t y) -> size_t {
                                                                          return (x << y) & ((1ul<<w)-1);
                                                                      },
                                                                      [w] (size_t des, bool f, size_t x, size_t y) -> bool {
                                                                          return (x >> (w-1));
                                                                      });
    print_time("shl tests successful");
}

} // libsnark

#endif

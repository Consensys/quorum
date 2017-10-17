/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM ALU arithmetic gadgets.

 These gadget check the correct execution of arithmetic TinyRAM instructions.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef ALU_ARITHMETIC_HPP_
#define ALU_ARITHMETIC_HPP_
#include <memory>

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/tinyram_protoboard.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"

namespace libsnark {

/* arithmetic gadgets */
template<typename FieldT>
class ALU_arithmetic_gadget : public tinyram_standard_gadget<FieldT> {
public:
    const pb_variable_array<FieldT> opcode_indicators;
    const word_variable_gadget<FieldT> desval;
    const word_variable_gadget<FieldT> arg1val;
    const word_variable_gadget<FieldT> arg2val;
    const pb_variable<FieldT> flag;
    const pb_variable<FieldT> result;
    const pb_variable<FieldT> result_flag;

    ALU_arithmetic_gadget(tinyram_protoboard<FieldT> &pb,
                          const pb_variable_array<FieldT> &opcode_indicators,
                          const word_variable_gadget<FieldT> &desval,
                          const word_variable_gadget<FieldT> &arg1val,
                          const word_variable_gadget<FieldT> &arg2val,
                          const pb_variable<FieldT> &flag,
                          const pb_variable<FieldT> &result,
                          const pb_variable<FieldT> &result_flag,
                          const std::string &annotation_prefix="") :
        tinyram_standard_gadget<FieldT>(pb, annotation_prefix),
        opcode_indicators(opcode_indicators),
        desval(desval),
        arg1val(arg1val),
        arg2val(arg2val),
        flag(flag),
        result(result),
        result_flag(result_flag) {}
};

template<typename FieldT>
class ALU_and_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    pb_variable_array<FieldT> res_word;
    std::shared_ptr<packing_gadget<FieldT> > pack_result;
    std::shared_ptr<disjunction_gadget<FieldT> > not_all_zeros;
    pb_variable<FieldT> not_all_zeros_result;
public:
    ALU_and_gadget(tinyram_protoboard<FieldT> &pb,
                   const pb_variable_array<FieldT> &opcode_indicators,
                   const word_variable_gadget<FieldT> &desval,
                   const word_variable_gadget<FieldT> &arg1val,
                   const word_variable_gadget<FieldT> &arg2val,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &result,
                   const pb_variable<FieldT> &result_flag,
                   const std::string &annotation_prefix="") :
        ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix)
    {
        res_word.allocate(pb, pb.ap.w, FMT(this->annotation_prefix, " res_bit"));
        not_all_zeros_result.allocate(pb, FMT(this->annotation_prefix, " not_all_zeros_result"));

        pack_result.reset(
            new packing_gadget<FieldT>(pb, res_word, result,
                                       FMT(this->annotation_prefix, " pack_result")));
        not_all_zeros.reset(
            new disjunction_gadget<FieldT>(pb, res_word, not_all_zeros_result,
                                           FMT(this->annotation_prefix, "not_all_zeros")));
    }

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_and_gadget(const size_t w);

template<typename FieldT>
class ALU_or_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    pb_variable_array<FieldT> res_word;
    std::shared_ptr<packing_gadget<FieldT> > pack_result;
    std::shared_ptr<disjunction_gadget<FieldT> > not_all_zeros;
    pb_variable<FieldT> not_all_zeros_result;
public:
    ALU_or_gadget(tinyram_protoboard<FieldT> &pb,
                  const pb_variable_array<FieldT> &opcode_indicators,
                  const word_variable_gadget<FieldT> &desval,
                  const word_variable_gadget<FieldT> &arg1val,
                  const word_variable_gadget<FieldT> &arg2val,
                  const pb_variable<FieldT> &flag,
                  const pb_variable<FieldT> &result,
                  const pb_variable<FieldT> &result_flag,
                  const std::string &annotation_prefix="") :
        ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix)
    {
        res_word.allocate(pb, pb.ap.w, FMT(this->annotation_prefix, " res_bit"));
        not_all_zeros_result.allocate(pb, FMT(this->annotation_prefix, " not_all_zeros_result"));

        pack_result.reset(
            new packing_gadget<FieldT>(pb, res_word, result,
                                       FMT(this->annotation_prefix, " pack_result")));
        not_all_zeros.reset(
            new disjunction_gadget<FieldT>(pb, res_word, not_all_zeros_result,
                                           FMT(this->annotation_prefix, "not_all_zeros")));
    }

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_or_gadget(const size_t w);

template<typename FieldT>
class ALU_xor_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    pb_variable_array<FieldT> res_word;
    std::shared_ptr<packing_gadget<FieldT> > pack_result;
    std::shared_ptr<disjunction_gadget<FieldT> > not_all_zeros;
    pb_variable<FieldT> not_all_zeros_result;
public:
    ALU_xor_gadget(tinyram_protoboard<FieldT> &pb,
                   const pb_variable_array<FieldT> &opcode_indicators,
                   const word_variable_gadget<FieldT> &desval,
                   const word_variable_gadget<FieldT> &arg1val,
                   const word_variable_gadget<FieldT> &arg2val,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &result,
                   const pb_variable<FieldT> &result_flag,
                   const std::string &annotation_prefix="") :
        ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix)
    {
        res_word.allocate(pb, pb.ap.w, FMT(this->annotation_prefix, " res_bit"));
        not_all_zeros_result.allocate(pb, FMT(this->annotation_prefix, " not_all_zeros_result"));

        pack_result.reset(
            new packing_gadget<FieldT>(pb, res_word, result,
                                       FMT(this->annotation_prefix, " pack_result")));
        not_all_zeros.reset(
            new disjunction_gadget<FieldT>(pb, res_word, not_all_zeros_result,
                                           FMT(this->annotation_prefix, "not_all_zeros")));
    }

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_xor_gadget(const size_t w);

template<typename FieldT>
class ALU_not_gadget : public ALU_arithmetic_gadget<FieldT> {
/* we do bitwise not, because we need to compute flag */
private:
    pb_variable_array<FieldT> res_word;
    std::shared_ptr<packing_gadget<FieldT> > pack_result;
    std::shared_ptr<disjunction_gadget<FieldT> > not_all_zeros;
    pb_variable<FieldT> not_all_zeros_result;
public:
    ALU_not_gadget(tinyram_protoboard<FieldT> &pb,
                   const pb_variable_array<FieldT> &opcode_indicators,
                   const word_variable_gadget<FieldT> &desval,
                   const word_variable_gadget<FieldT> &arg1val,
                   const word_variable_gadget<FieldT> &arg2val,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &result,
                   const pb_variable<FieldT> &result_flag,
                   const std::string &annotation_prefix="") :
        ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix)
    {
        res_word.allocate(pb, pb.ap.w, FMT(this->annotation_prefix, " res_bit"));
        not_all_zeros_result.allocate(pb, FMT(this->annotation_prefix, " not_all_zeros_result"));

        pack_result.reset(
            new packing_gadget<FieldT>(pb, res_word, result,
                                       FMT(this->annotation_prefix, " pack_result")));
        not_all_zeros.reset(
            new disjunction_gadget<FieldT>(pb, res_word, not_all_zeros_result,
                                           FMT(this->annotation_prefix, "not_all_zeros")));
    }

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_not_gadget(const size_t w);

template<typename FieldT>
class ALU_add_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    pb_variable<FieldT> addition_result;
    pb_variable_array<FieldT> res_word;
    pb_variable_array<FieldT> res_word_and_flag;
    std::shared_ptr<packing_gadget<FieldT> > unpack_addition, pack_result;
public:
    ALU_add_gadget(tinyram_protoboard<FieldT> &pb,
                   const pb_variable_array<FieldT> &opcode_indicators,
                   const word_variable_gadget<FieldT> &desval,
                   const word_variable_gadget<FieldT> &arg1val,
                   const word_variable_gadget<FieldT> &arg2val,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &result,
                   const pb_variable<FieldT> &result_flag,
                   const std::string &annotation_prefix="") :
        ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix)
    {
        addition_result.allocate(pb, FMT(this->annotation_prefix, " addition_result"));
        res_word.allocate(pb, pb.ap.w, FMT(this->annotation_prefix, " res_word"));

        res_word_and_flag = res_word;
        res_word_and_flag.emplace_back(result_flag);

        unpack_addition.reset(
            new packing_gadget<FieldT>(pb, res_word_and_flag, addition_result,
                                       FMT(this->annotation_prefix, " unpack_addition")));
        pack_result.reset(
            new packing_gadget<FieldT>(pb, res_word, result,
                                       FMT(this->annotation_prefix, " pack_result")));
    }

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

void test_ALU_add_gadget(const size_t w);

template<typename FieldT>
class ALU_sub_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    pb_variable<FieldT> intermediate_result;
    pb_variable<FieldT> negated_flag;
    pb_variable_array<FieldT> res_word;
    pb_variable_array<FieldT> res_word_and_negated_flag;

    std::shared_ptr<packing_gadget<FieldT> > unpack_intermediate, pack_result;
public:
    ALU_sub_gadget(tinyram_protoboard<FieldT> &pb,
                   const pb_variable_array<FieldT> &opcode_indicators,
                   const word_variable_gadget<FieldT> &desval,
                   const word_variable_gadget<FieldT> &arg1val,
                   const word_variable_gadget<FieldT> &arg2val,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &result,
                   const pb_variable<FieldT> &result_flag,
                   const std::string &annotation_prefix="") :
        ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix)
    {
        intermediate_result.allocate(pb, FMT(this->annotation_prefix, " intermediate_result"));
        negated_flag.allocate(pb, FMT(this->annotation_prefix, " negated_flag"));
        res_word.allocate(pb, pb.ap.w, FMT(this->annotation_prefix, " res_word"));

        res_word_and_negated_flag = res_word;
        res_word_and_negated_flag.emplace_back(negated_flag);

        unpack_intermediate.reset(
            new packing_gadget<FieldT>(pb, res_word_and_negated_flag, intermediate_result,
                                       FMT(this->annotation_prefix, " unpack_intermediate")));
        pack_result.reset(
            new packing_gadget<FieldT>(pb, res_word, result,
                                       FMT(this->annotation_prefix, " pack_result")));
    }

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

void test_ALU_sub_gadget(const size_t w);

template<typename FieldT>
class ALU_mov_gadget : public ALU_arithmetic_gadget<FieldT> {
public:
    ALU_mov_gadget(tinyram_protoboard<FieldT> &pb,
                   const pb_variable_array<FieldT> &opcode_indicators,
                   const word_variable_gadget<FieldT> &desval,
                   const word_variable_gadget<FieldT> &arg1val,
                   const word_variable_gadget<FieldT> &arg2val,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &result,
                   const pb_variable<FieldT> &result_flag,
                   const std::string &annotation_prefix="") :
        ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix) {}

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_mov_gadget(const size_t w);

template<typename FieldT>
class ALU_cmov_gadget : public ALU_arithmetic_gadget<FieldT> {
public:
    ALU_cmov_gadget(tinyram_protoboard<FieldT> &pb,
                    const pb_variable_array<FieldT> &opcode_indicators,
                    const word_variable_gadget<FieldT> &desval,
                    const word_variable_gadget<FieldT> &arg1val,
                    const word_variable_gadget<FieldT> &arg2val,
                    const pb_variable<FieldT> &flag,
                    const pb_variable<FieldT> &result,
                    const pb_variable<FieldT> &result_flag,
                    const std::string &annotation_prefix="") :
    ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, result, result_flag, annotation_prefix)
    {
    }

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_cmov_gadget(const size_t w);

template<typename FieldT>
class ALU_cmp_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    comparison_gadget<FieldT> comparator;
public:
    const pb_variable<FieldT> cmpe_result;
    const pb_variable<FieldT> cmpe_result_flag;
    const pb_variable<FieldT> cmpa_result;
    const pb_variable<FieldT> cmpa_result_flag;
    const pb_variable<FieldT> cmpae_result;
    const pb_variable<FieldT> cmpae_result_flag;

    ALU_cmp_gadget(tinyram_protoboard<FieldT> &pb,
                   const pb_variable_array<FieldT> &opcode_indicators,
                   const word_variable_gadget<FieldT> &desval,
                   const word_variable_gadget<FieldT> &arg1val,
                   const word_variable_gadget<FieldT> &arg2val,
                   const pb_variable<FieldT> &flag,
                   const pb_variable<FieldT> &cmpe_result,
                   const pb_variable<FieldT> &cmpe_result_flag,
                   const pb_variable<FieldT> &cmpa_result,
                   const pb_variable<FieldT> &cmpa_result_flag,
                   const pb_variable<FieldT> &cmpae_result,
                   const pb_variable<FieldT> &cmpae_result_flag,
                   const std::string &annotation_prefix="") :
    ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, cmpa_result, cmpa_result_flag, annotation_prefix),
        comparator(pb, pb.ap.w, arg2val.packed, arg1val.packed, cmpa_result_flag, cmpae_result_flag,
                   FMT(this->annotation_prefix, " comparator")),
        cmpe_result(cmpe_result), cmpe_result_flag(cmpe_result_flag),
        cmpa_result(cmpa_result), cmpa_result_flag(cmpa_result_flag),
        cmpae_result(cmpae_result), cmpae_result_flag(cmpae_result_flag) {}

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_cmpe_gadget(const size_t w);

template<typename FieldT>
void test_ALU_cmpa_gadget(const size_t w);

template<typename FieldT>
void test_ALU_cmpae_gadget(const size_t w);

template<typename FieldT>
class ALU_cmps_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    pb_variable<FieldT> negated_arg1val_sign;
    pb_variable<FieldT> negated_arg2val_sign;
    pb_variable_array<FieldT> modified_arg1;
    pb_variable_array<FieldT> modified_arg2;
    pb_variable<FieldT> packed_modified_arg1;
    pb_variable<FieldT> packed_modified_arg2;
    std::shared_ptr<packing_gadget<FieldT> > pack_modified_arg1;
    std::shared_ptr<packing_gadget<FieldT> > pack_modified_arg2;
    std::shared_ptr<comparison_gadget<FieldT> > comparator;
public:
    const pb_variable<FieldT> cmpg_result;
    const pb_variable<FieldT> cmpg_result_flag;
    const pb_variable<FieldT> cmpge_result;
    const pb_variable<FieldT> cmpge_result_flag;

    ALU_cmps_gadget(tinyram_protoboard<FieldT> &pb,
                    const pb_variable_array<FieldT> &opcode_indicators,
                    const word_variable_gadget<FieldT> &desval,
                    const word_variable_gadget<FieldT> &arg1val,
                    const word_variable_gadget<FieldT> &arg2val,
                    const pb_variable<FieldT> &flag,
                    const pb_variable<FieldT> &cmpg_result,
                    const pb_variable<FieldT> &cmpg_result_flag,
                    const pb_variable<FieldT> &cmpge_result,
                    const pb_variable<FieldT> &cmpge_result_flag,
                    const std::string &annotation_prefix="") :
    ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, cmpg_result, cmpg_result_flag, annotation_prefix),
        cmpg_result(cmpg_result), cmpg_result_flag(cmpg_result_flag),
        cmpge_result(cmpge_result), cmpge_result_flag(cmpge_result_flag)
    {
        negated_arg1val_sign.allocate(pb, FMT(this->annotation_prefix, " negated_arg1val_sign"));
        negated_arg2val_sign.allocate(pb, FMT(this->annotation_prefix, " negated_arg2val_sign"));

        modified_arg1 = pb_variable_array<FieldT>(arg1val.bits.begin(), --arg1val.bits.end());
        modified_arg1.emplace_back(negated_arg1val_sign);

        modified_arg2 = pb_variable_array<FieldT>(arg2val.bits.begin(), --arg2val.bits.end());
        modified_arg2.emplace_back(negated_arg2val_sign);

        packed_modified_arg1.allocate(pb, FMT(this->annotation_prefix, " packed_modified_arg1"));
        packed_modified_arg2.allocate(pb, FMT(this->annotation_prefix, " packed_modified_arg2"));

        pack_modified_arg1.reset(new packing_gadget<FieldT>(pb, modified_arg1, packed_modified_arg1,
                                                            FMT(this->annotation_prefix, " pack_modified_arg1")));
        pack_modified_arg2.reset(new packing_gadget<FieldT>(pb, modified_arg2, packed_modified_arg2,
                                                            FMT(this->annotation_prefix, " pack_modified_arg2")));

        comparator.reset(new comparison_gadget<FieldT>(pb, pb.ap.w,
                                                       packed_modified_arg2, packed_modified_arg1,
                                                       cmpg_result_flag, cmpge_result_flag,
                                                       FMT(this->annotation_prefix, " comparator")));
    }
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_cmpg_gadget(const size_t w);

template<typename FieldT>
void test_ALU_cmpge_gadget(const size_t w);

template<typename FieldT>
class ALU_umul_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    dual_variable_gadget<FieldT> mul_result;
    pb_variable_array<FieldT> mull_bits;
    pb_variable_array<FieldT> umulh_bits;
    pb_variable<FieldT> result_flag;
    std::shared_ptr<packing_gadget<FieldT> > pack_mull_result;
    std::shared_ptr<packing_gadget<FieldT> > pack_umulh_result;
    std::shared_ptr<disjunction_gadget<FieldT> > compute_flag;
public:
    const pb_variable<FieldT> mull_result;
    const pb_variable<FieldT> mull_flag;
    const pb_variable<FieldT> umulh_result;
    const pb_variable<FieldT> umulh_flag;

    ALU_umul_gadget(tinyram_protoboard<FieldT> &pb,
                    const pb_variable_array<FieldT> &opcode_indicators,
                    const word_variable_gadget<FieldT> &desval,
                    const word_variable_gadget<FieldT> &arg1val,
                    const word_variable_gadget<FieldT> &arg2val,
                    const pb_variable<FieldT> &flag,
                    const pb_variable<FieldT> &mull_result,
                    const pb_variable<FieldT> &mull_flag,
                    const pb_variable<FieldT> &umulh_result,
                    const pb_variable<FieldT> &umulh_flag,
                    const std::string &annotation_prefix="") :
    ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, mull_result, mull_flag, annotation_prefix),
        mul_result(pb, 2*pb.ap.w, FMT(this->annotation_prefix, " mul_result")),
        mull_result(mull_result), mull_flag(mull_flag), umulh_result(umulh_result), umulh_flag(umulh_flag)
    {
        mull_bits.insert(mull_bits.end(), mul_result.bits.begin(), mul_result.bits.begin()+pb.ap.w);
        umulh_bits.insert(umulh_bits.end(), mul_result.bits.begin()+pb.ap.w, mul_result.bits.begin()+2*pb.ap.w);

        pack_mull_result.reset(new packing_gadget<FieldT>(pb, mull_bits, mull_result, FMT(this->annotation_prefix, " pack_mull_result")));
        pack_umulh_result.reset(new packing_gadget<FieldT>(pb, umulh_bits, umulh_result, FMT(this->annotation_prefix, " pack_umulh_result")));

        result_flag.allocate(pb, FMT(this->annotation_prefix, " result_flag"));
        compute_flag.reset(new disjunction_gadget<FieldT>(pb, umulh_bits, result_flag, FMT(this->annotation_prefix, " compute_flag")));
    }
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_mull_gadget(const size_t w);

template<typename FieldT>
void test_ALU_umulh_gadget(const size_t w);

template<typename FieldT>
class ALU_smul_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    dual_variable_gadget<FieldT> mul_result;
    pb_variable_array<FieldT> smulh_bits;

    pb_variable<FieldT> top;
    std::shared_ptr<packing_gadget<FieldT> > pack_top;

    pb_variable<FieldT> is_top_empty, is_top_empty_aux;
    pb_variable<FieldT> is_top_full, is_top_full_aux;

    pb_variable<FieldT> result_flag;
    std::shared_ptr<packing_gadget<FieldT> > pack_smulh_result;
public:
    const pb_variable<FieldT> smulh_result;
    const pb_variable<FieldT> smulh_flag;

    ALU_smul_gadget(tinyram_protoboard<FieldT> &pb,
                    const pb_variable_array<FieldT> &opcode_indicators,
                    const word_variable_gadget<FieldT> &desval,
                    const word_variable_gadget<FieldT> &arg1val,
                    const word_variable_gadget<FieldT> &arg2val,
                    const pb_variable<FieldT> &flag,
                    const pb_variable<FieldT> &smulh_result,
                    const pb_variable<FieldT> &smulh_flag,
                    const std::string &annotation_prefix="") :
    ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, smulh_result, smulh_flag, annotation_prefix),
        mul_result(pb, 2*pb.ap.w+1, FMT(this->annotation_prefix, " mul_result")), /* see witness map for explanation for 2w+1 */
        smulh_result(smulh_result), smulh_flag(smulh_flag)
    {
        smulh_bits.insert(smulh_bits.end(), mul_result.bits.begin()+pb.ap.w, mul_result.bits.begin()+2*pb.ap.w);

        pack_smulh_result.reset(new packing_gadget<FieldT>(pb, smulh_bits, smulh_result, FMT(this->annotation_prefix, " pack_smulh_result")));

        top.allocate(pb, FMT(this->annotation_prefix, " top"));
        pack_top.reset(new packing_gadget<FieldT>(pb, pb_variable_array<FieldT>(mul_result.bits.begin() + pb.ap.w-1, mul_result.bits.begin() + 2*pb.ap.w), top,
                                                  FMT(this->annotation_prefix, " pack_top")));

        is_top_empty.allocate(pb, FMT(this->annotation_prefix, " is_top_empty"));
        is_top_empty_aux.allocate(pb, FMT(this->annotation_prefix, " is_top_empty_aux"));

        is_top_full.allocate(pb, FMT(this->annotation_prefix, " is_top_full"));
        is_top_full_aux.allocate(pb, FMT(this->annotation_prefix, " is_top_full_aux"));

        result_flag.allocate(pb, FMT(this->annotation_prefix, " result_flag"));
    }
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_smulh_gadget(const size_t w);

template<typename FieldT>
class ALU_divmod_gadget : public ALU_arithmetic_gadget<FieldT> {
/*
  <<<<<<< Updated upstream
  B * q + r = A_aux = A * B_nonzero
  q * (1-B_nonzero) = 0
  A<B_gadget<FieldT>(r < B, less=B_nonzero, leq=ONE)
  =======
  B * q + r = A

  r <= B
  >>>>>>> Stashed changes
*/
private:
    pb_variable<FieldT> B_inv;
    pb_variable<FieldT> B_nonzero;
    pb_variable<FieldT> A_aux;
    std::shared_ptr<comparison_gadget<FieldT> > r_less_B;
public:
    const pb_variable<FieldT> udiv_result;
    const pb_variable<FieldT> udiv_flag;
    const pb_variable<FieldT> umod_result;
    const pb_variable<FieldT> umod_flag;

    ALU_divmod_gadget(tinyram_protoboard<FieldT> &pb,
                      const pb_variable_array<FieldT> &opcode_indicators,
                      const word_variable_gadget<FieldT> &desval,
                      const word_variable_gadget<FieldT> &arg1val,
                      const word_variable_gadget<FieldT> &arg2val,
                      const pb_variable<FieldT> &flag,
                      const pb_variable<FieldT> &udiv_result,
                      const pb_variable<FieldT> &udiv_flag,
                      const pb_variable<FieldT> &umod_result,
                      const pb_variable<FieldT> &umod_flag,
                      const std::string &annotation_prefix="") :
    ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, udiv_result, udiv_flag, annotation_prefix),
        udiv_result(udiv_result), udiv_flag(udiv_flag), umod_result(umod_result), umod_flag(umod_flag)
    {
        B_inv.allocate(pb, FMT(this->annotation_prefix, " B_inv"));
        B_nonzero.allocate(pb, FMT(this->annotation_prefix, " B_nonzer"));
        A_aux.allocate(pb, FMT(this->annotation_prefix, " A_aux"));
        r_less_B.reset(new comparison_gadget<FieldT>(pb, pb.ap.w, umod_result, arg2val.packed,
                                                     B_nonzero, ONE, FMT(this->annotation_prefix, " r_less_B")));
    }
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_udiv_gadget(const size_t w);

template<typename FieldT>
void test_ALU_umod_gadget(const size_t w);

template<typename FieldT>
class ALU_shr_shl_gadget : public ALU_arithmetic_gadget<FieldT> {
private:
    pb_variable<FieldT> reversed_input;
    std::shared_ptr<packing_gadget<FieldT> > pack_reversed_input;

    pb_variable_array<FieldT> barrel_right_internal;
    std::vector<pb_variable_array<FieldT> > shifted_out_bits;

    pb_variable<FieldT> is_oversize_shift;
    std::shared_ptr<disjunction_gadget<FieldT> > check_oversize_shift;
    pb_variable<FieldT> result;

    pb_variable_array<FieldT> result_bits;
    std::shared_ptr<packing_gadget<FieldT> > unpack_result;
    pb_variable<FieldT> reversed_result;
    std::shared_ptr<packing_gadget<FieldT> > pack_reversed_result;
public:
    pb_variable<FieldT> shr_result;
    pb_variable<FieldT> shr_flag;
    pb_variable<FieldT> shl_result;
    pb_variable<FieldT> shl_flag;

    size_t logw;

    ALU_shr_shl_gadget(tinyram_protoboard<FieldT> &pb,
                       const pb_variable_array<FieldT> &opcode_indicators,
                       const word_variable_gadget<FieldT> &desval,
                       const word_variable_gadget<FieldT> &arg1val,
                       const word_variable_gadget<FieldT> &arg2val,
                       const pb_variable<FieldT> &flag,
                       const pb_variable<FieldT> &shr_result,
                       const pb_variable<FieldT> &shr_flag,
                       const pb_variable<FieldT> &shl_result,
                       const pb_variable<FieldT> &shl_flag,
                       const std::string &annotation_prefix="") :
    ALU_arithmetic_gadget<FieldT>(pb, opcode_indicators, desval, arg1val, arg2val, flag, shr_result, shr_flag, annotation_prefix),
        shr_result(shr_result), shr_flag(shr_flag), shl_result(shl_result), shl_flag(shl_flag)
    {
        logw = log2(pb.ap.w);

        reversed_input.allocate(pb, FMT(this->annotation_prefix, " reversed_input"));
        pack_reversed_input.reset(
            new packing_gadget<FieldT>(pb, pb_variable_array<FieldT>(arg1val.bits.rbegin(), arg1val.bits.rend()),
                                       reversed_input,
                                       FMT(this->annotation_prefix, " pack_reversed_input")));

        barrel_right_internal.allocate(pb, logw+1, FMT(this->annotation_prefix, " barrel_right_internal"));

        shifted_out_bits.resize(logw);
        for (size_t i = 0; i < logw; ++i)
        {
            shifted_out_bits[i].allocate(pb, 1ul<<i, FMT(this->annotation_prefix, " shifted_out_bits_%zu", i));
        }

        is_oversize_shift.allocate(pb, FMT(this->annotation_prefix, " is_oversize_shift"));
        check_oversize_shift.reset(
            new disjunction_gadget<FieldT>(pb,
                                           pb_variable_array<FieldT>(arg2val.bits.begin()+logw, arg2val.bits.end()),
                                           is_oversize_shift,
                                           FMT(this->annotation_prefix, " check_oversize_shift")));
        result.allocate(pb, FMT(this->annotation_prefix, " result"));

        result_bits.allocate(pb, pb.ap.w, FMT(this->annotation_prefix, " result_bits"));
        unpack_result.reset(
            new packing_gadget<FieldT>(pb, result_bits, result, //barrel_right_internal[logw],
                                       FMT(this->annotation_prefix, " unpack_result")));

        reversed_result.allocate(pb, FMT(this->annotation_prefix, " reversed_result"));
        pack_reversed_result.reset(
            new packing_gadget<FieldT>(pb, pb_variable_array<FieldT>(result_bits.rbegin(), result_bits.rend()),
                                       reversed_result,
                                       FMT(this->annotation_prefix, " pack_reversed_result")));
    }
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_ALU_shr_gadget(const size_t w);

template<typename FieldT>
void test_ALU_shl_gadget(const size_t w);

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/alu_arithmetic.tcc"

#endif // ALU_ARITHMETIC_HPP_

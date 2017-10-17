/** @file
 *****************************************************************************

 Implementation of interfaces for the TinyRAM instruction packing gadget.

 See instruction_packing_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef INSTRUCTION_PACKING_GADGET_TCC_
#define INSTRUCTION_PACKING_GADGET_TCC_

namespace libsnark {

template<typename FieldT>
tinyram_instruction_packing_gadget<FieldT>::tinyram_instruction_packing_gadget(tinyram_protoboard<FieldT> &pb,
                                                                               const pb_variable_array<FieldT> &opcode,
                                                                               const pb_variable<FieldT> &arg2_is_imm,
                                                                               const pb_variable_array<FieldT> &desidx,
                                                                               const pb_variable_array<FieldT> &arg1idx,
                                                                               const pb_variable_array<FieldT> &arg2idx,
                                                                               const pb_variable<FieldT> &packed_instruction,
                                                                               const std::string &annotation_prefix) :
    tinyram_gadget<FieldT>(pb, annotation_prefix),
    opcode(opcode),
    arg2_is_imm(arg2_is_imm),
    desidx(desidx),
    arg1idx(arg1idx),
    arg2idx(arg2idx),
    packed_instruction(packed_instruction)
{
    all_bits.reserve(2*pb.ap.w);

    all_bits.insert(all_bits.begin(), opcode.begin(), opcode.end());
    all_bits.insert(all_bits.begin(), arg2_is_imm);
    all_bits.insert(all_bits.begin(), desidx.begin(), desidx.end());
    all_bits.insert(all_bits.begin(), arg1idx.begin(), arg1idx.end());
    dummy.allocate(pb, pb.ap.w-all_bits.size(), FMT(this->annotation_prefix, " dummy"));
    all_bits.insert(all_bits.begin(), dummy.begin(), dummy.end());
    all_bits.insert(all_bits.begin(), arg2idx.begin(), arg2idx.end());

    assert(all_bits.size() == 2*pb.ap.w);

    pack_instruction.reset(
        new packing_gadget<FieldT>(pb, all_bits, packed_instruction, FMT(this->annotation_prefix, " pack_instruction")));
}


template<typename FieldT>
void tinyram_instruction_packing_gadget<FieldT>::generate_r1cs_constraints(const bool enforce_bitness)
{
    pack_instruction->generate_r1cs_constraints(enforce_bitness);
}

template<typename FieldT>
void tinyram_instruction_packing_gadget<FieldT>::generate_r1cs_witness_from_packed()
{
    pack_instruction->generate_r1cs_witness_from_packed();
}

template<typename FieldT>
void tinyram_instruction_packing_gadget<FieldT>::generate_r1cs_witness_from_bits()
{
    pack_instruction->generate_r1cs_witness_from_bits();
}

template<typename FieldT>
FieldT pack_instruction(const tinyram_architecture_params &ap,
                        const tinyram_instruction &instr)
{
    tinyram_program P; P.instructions = generate_tinyram_prelude(ap);
    tinyram_protoboard<FieldT> pb(ap, P.size(), 0, 10);

    pb_variable_array<FieldT> v_opcode;
    pb_variable<FieldT> v_arg2_is_imm;
    pb_variable_array<FieldT> v_desidx;
    pb_variable_array<FieldT> v_arg1idx;
    pb_variable_array<FieldT> v_arg2idx;

    v_opcode.allocate(pb, ap.s, "opcode");
    v_arg2_is_imm.allocate(pb, "arg2_is_imm");
    v_desidx.allocate(pb, ap.reg_arg_width(), "desidx");
    v_arg1idx.allocate(pb, ap.reg_arg_width(), "arg1idx");
    v_arg2idx.allocate(pb, ap.reg_arg_or_imm_width(), "arg2idx");

    v_opcode.fill_with_bits_of_ulong(pb, instr.opcode);
    pb.val(v_arg2_is_imm) = instr.arg2_is_imm ? FieldT::one() : FieldT::zero();
    v_desidx.fill_with_bits_of_ulong(pb, instr.desidx);
    v_arg1idx.fill_with_bits_of_ulong(pb, instr.arg1idx);
    v_arg2idx.fill_with_bits_of_ulong(pb, instr.arg2idx_or_imm);

    pb_variable<FieldT> packed;
    packed.allocate(pb, "packed");

    tinyram_instruction_packing_gadget<FieldT> g(pb, v_opcode, v_arg2_is_imm, v_desidx, v_arg1idx, v_arg2idx, packed, "g");
    g.generate_r1cs_constraints(true);
    g.generate_r1cs_witness_from_bits();

    return pb.val(packed);
}

template<typename FieldT>
void test_instruction_packing()
{
    print_time("starting instruction packing test");

    tinyram_architecture_params ap(16, 16);
    tinyram_program P; P.instructions = generate_tinyram_prelude(ap);
    tinyram_protoboard<FieldT> pb(ap, P.size(), 0, 10);

    pb_variable_array<FieldT> opcode[2];
    pb_variable<FieldT> arg2_is_imm[2];
    pb_variable_array<FieldT> desidx[2];
    pb_variable_array<FieldT> arg1idx[2];
    pb_variable_array<FieldT> arg2idx[2];

    for (size_t i = 0; i < 2; ++i)
    {
        opcode[i].allocate(pb, ap.s, FMT("", "opcode_%zu", i));
        arg2_is_imm[i].allocate(pb, FMT("", "arg2_is_imm_%zu", i));
        desidx[i].allocate(pb, ap.reg_arg_width(), FMT("", "desidx_%zu", i));
        arg1idx[i].allocate(pb, ap.reg_arg_width(), FMT("", "arg1idx_%zu", i));
        arg2idx[i].allocate(pb, ap.reg_arg_or_imm_width(), FMT("", "arg2idx_%zu", i));
    }

    pb_variable<FieldT> packed_instr;
    packed_instr.allocate(pb, "packed_instr");

    instruction_packing_gadget<FieldT> pack(pb, opcode[0], arg2_is_imm[0], desidx[0], arg1idx[0], arg2idx[0], packed_instr, "pack");
    instruction_packing_gadget<FieldT> unpack(pb, opcode[1], arg2_is_imm[1], desidx[1], arg1idx[1], arg2idx[1], packed_instr, "unpack");

    pack.generate_r1cs_constraints(true);
    unpack.generate_r1cs_constraints(true);

    for (size_t k = 0; k < 100; ++k)
    {
        tinyram_opcode oc = static_cast<tinyram_opcode>(std::rand() % (1u << ap.s));
        bool imm = std::rand() % 2;
        size_t des = rand() % (1u << ap.reg_arg_width());
        size_t arg1 = rand() % (1u << ap.reg_arg_width());
        size_t arg2 = rand() % (1u << ap.reg_arg_or_imm_width());

        for (size_t i = 0; i < ap.s; ++i)
        {
            pb.val(opcode[0][i]) = (oc & (1ul<<i) ? FieldT::one() : FieldT::zero());
        }

        pb.val(arg2_is_imm[0]) = (imm ? FieldT::one() : FieldT::zero());

        for (size_t i = 0; i < ap.reg_arg_width(); ++i)
        {
            pb.val(desidx[0][i]) = (des & (1ul<<i) ? FieldT::one() : FieldT::zero());
            pb.val(arg1idx[0][i]) = (arg1 & (1ul<<i) ? FieldT::one() : FieldT::zero());
        }

        for (size_t i = 0; i < ap.reg_arg_or_imm_width(); ++i)
        {
            pb.val(arg2idx[0][i]) = (arg2 & (1ul<<i) ? FieldT::one() : FieldT::zero());
        }

        pack.generate_r1cs_witness_from_bits();
        unpack.generate_r1cs_witness_from_packed();

        for (size_t i = 0; i < ap.s; ++i)
        {
            assert(pb.val(opcode[0][i]) == pb.val(opcode[1][i]));
        }

        assert(pb.val(arg2_is_imm[0]) == pb.val(arg2_is_imm[1]));

        for (size_t i = 0; i < ap.reg_arg_width(); ++i)
        {
            assert(pb.val(desidx[0][i]) == pb.val(desidx[1][i]));
            assert(pb.val(arg1idx[1][i]) == pb.val(arg1idx[1][i]));
        }

        for (size_t i = 0; i < ap.reg_arg_or_imm_width(); ++i)
        {
            assert(pb.val(arg2idx[0][i]) == pb.val(arg2idx[1][i]));
        }

        assert(pb.val(packed_instr) == pack_instruction<FieldT>(ap, tinyram_instruction(oc, imm, des, arg1, arg2)));
        assert(pb.is_satisfied());
    }

    print_time("instruction packing tests successful");
}

} // libsnark

#endif // INSTRUCTION_PACKING_GADGET_TCC_

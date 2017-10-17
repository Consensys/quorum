/** @file
 *****************************************************************************

 Declaration of interfaces for the TinyRAM argument decoder gadget.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef ARGUMENT_DECODER_GADGET_HPP_
#define ARGUMENT_DECODER_GADGET_HPP_

namespace libsnark {

template<typename FieldT>
class argument_decoder_gadget : public tinyram_standard_gadget<FieldT> {
private:
    pb_variable<FieldT> packed_desidx;
    pb_variable<FieldT> packed_arg1idx;
    pb_variable<FieldT> packed_arg2idx;

    std::shared_ptr<packing_gadget<FieldT> > pack_desidx;
    std::shared_ptr<packing_gadget<FieldT> > pack_arg1idx;
    std::shared_ptr<packing_gadget<FieldT> > pack_arg2idx;

    pb_variable<FieldT> arg2_demux_result;
    pb_variable<FieldT> arg2_demux_success;

    std::shared_ptr<loose_multiplexing_gadget<FieldT> > demux_des;
    std::shared_ptr<loose_multiplexing_gadget<FieldT> > demux_arg1;
    std::shared_ptr<loose_multiplexing_gadget<FieldT> > demux_arg2;
public:
    pb_variable<FieldT> arg2_is_imm;
    pb_variable_array<FieldT> desidx;
    pb_variable_array<FieldT> arg1idx;
    pb_variable_array<FieldT> arg2idx;
    pb_variable_array<FieldT> packed_registers;
    pb_variable<FieldT> packed_desval;
    pb_variable<FieldT> packed_arg1val;
    pb_variable<FieldT> packed_arg2val;

    argument_decoder_gadget(tinyram_protoboard<FieldT> &pb,
                            const pb_variable<FieldT> &arg2_is_imm,
                            const pb_variable_array<FieldT> &desidx,
                            const pb_variable_array<FieldT> &arg1idx,
                            const pb_variable_array<FieldT> &arg2idx,
                            const pb_variable_array<FieldT> &packed_registers,
                            const pb_variable<FieldT> &packed_desval,
                            const pb_variable<FieldT> &packed_arg1val,
                            const pb_variable<FieldT> &packed_arg2val,
                            const std::string &annotation_prefix="");

    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FieldT>
void test_argument_decoder_gadget();

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/argument_decoder_gadget.tcc"

#endif // ARGUMENT_DECODER_GADGET_HPP_

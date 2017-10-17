/** @file
 *****************************************************************************

 Implementation of interfaces for the tally compliance predicate.

 See tally_cp.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TALLY_CP_TCC_
#define TALLY_CP_TCC_

#include <algorithm>
#include <functional>

namespace libsnark {

template<typename FieldT>
tally_pcd_message<FieldT>::tally_pcd_message(const size_t type,
                                             const size_t wordsize,
                                             const size_t sum,
                                             const size_t count) :
    r1cs_pcd_message<FieldT>(type), wordsize(wordsize), sum(sum), count(count)
{
}

template<typename FieldT>
r1cs_variable_assignment<FieldT> tally_pcd_message<FieldT>::payload_as_r1cs_variable_assignment() const
{
    std::function<FieldT(bool)> bit_to_FieldT = [] (const bool bit) { return bit ? FieldT::one() : FieldT::zero(); };

    const bit_vector sum_bits = convert_field_element_to_bit_vector<FieldT>(sum, wordsize);
    const bit_vector count_bits = convert_field_element_to_bit_vector<FieldT>(count, wordsize);

    r1cs_variable_assignment<FieldT> result(2 * wordsize);
    std::transform(sum_bits.begin(), sum_bits.end(), result.begin() , bit_to_FieldT);
    std::transform(count_bits.begin(), count_bits.end(), result.begin() + wordsize, bit_to_FieldT);

    return result;
}

template<typename FieldT>
void tally_pcd_message<FieldT>::print() const
{
    printf("Tally message of type %zu:\n", this->type);
    printf("  wordsize: %zu\n", wordsize);
    printf("  sum: %zu\n", sum);
    printf("  count: %zu\n", count);
}

template<typename FieldT>
tally_pcd_local_data<FieldT>::tally_pcd_local_data(const size_t summand) :
    summand(summand)
{
}

template<typename FieldT>
r1cs_variable_assignment<FieldT> tally_pcd_local_data<FieldT>::as_r1cs_variable_assignment() const
{
    const r1cs_variable_assignment<FieldT> result = { FieldT(summand) };
    return result;
}

template<typename FieldT>
void tally_pcd_local_data<FieldT>::print() const
{
    printf("Tally PCD local data:\n");
    printf("  summand: %zu\n", summand);
}

template<typename FieldT>
class tally_pcd_message_variable: public r1cs_pcd_message_variable<FieldT> {
public:
    pb_variable_array<FieldT> sum_bits;
    pb_variable_array<FieldT> count_bits;
    size_t wordsize;

    tally_pcd_message_variable(protoboard<FieldT> &pb,
                               const size_t wordsize,
                               const std::string &annotation_prefix) :
        r1cs_pcd_message_variable<FieldT>(pb, annotation_prefix), wordsize(wordsize)
    {
        sum_bits.allocate(pb, wordsize, FMT(annotation_prefix, " sum_bits"));
        count_bits.allocate(pb, wordsize, FMT(annotation_prefix, " count_bits"));

        this->update_all_vars();
    }

    std::shared_ptr<r1cs_pcd_message<FieldT> > get_message() const
    {
        const size_t type_val = this->pb.val(this->type).as_ulong();
        const size_t sum_val = sum_bits.get_field_element_from_bits(this->pb).as_ulong();
        const size_t count_val = count_bits.get_field_element_from_bits(this->pb).as_ulong();

        std::shared_ptr<r1cs_pcd_message<FieldT> > result;
        result.reset(new tally_pcd_message<FieldT>(type_val, wordsize, sum_val, count_val));
        return result;
    }

    ~tally_pcd_message_variable() = default;
};

template<typename FieldT>
class tally_pcd_local_data_variable : public r1cs_pcd_local_data_variable<FieldT> {
public:

    pb_variable<FieldT> summand;

    tally_pcd_local_data_variable(protoboard<FieldT> &pb,
                                  const std::string &annotation_prefix) :
        r1cs_pcd_local_data_variable<FieldT>(pb, annotation_prefix)
    {
        summand.allocate(pb, FMT(annotation_prefix, " summand"));

        this->update_all_vars();
    }

    std::shared_ptr<r1cs_pcd_local_data<FieldT> > get_local_data() const
    {
        const size_t summand_val = this->pb.val(summand).as_ulong();

        std::shared_ptr<r1cs_pcd_local_data<FieldT> > result;
        result.reset(new tally_pcd_local_data<FieldT>(summand_val));
        return result;
    }

    ~tally_pcd_local_data_variable() = default;
};

template<typename FieldT>
tally_cp_handler<FieldT>::tally_cp_handler(const size_t type, const size_t max_arity, const size_t wordsize,
                                           const bool relies_on_same_type_inputs,
                                           const std::set<size_t> accepted_input_types) :
    compliance_predicate_handler<FieldT, protoboard<FieldT> >(protoboard<FieldT>(),
                                                              type*100,
                                                              type,
                                                              max_arity,
                                                              relies_on_same_type_inputs,
                                                              accepted_input_types),
    wordsize(wordsize)
{
    this->outgoing_message.reset(new tally_pcd_message_variable<FieldT>(this->pb, wordsize, "outgoing_message"));
    this->arity.allocate(this->pb, "arity");

    for (size_t i = 0; i < max_arity; ++i)
    {
        this->incoming_messages[i].reset(new tally_pcd_message_variable<FieldT>(this->pb, wordsize, FMT("", "incoming_messages_%zu", i)));
    }

    this->local_data.reset(new tally_pcd_local_data_variable<FieldT>(this->pb, "local_data"));

    sum_out_packed.allocate(this->pb, "sum_out_packed");
    count_out_packed.allocate(this->pb, "count_out_packed");

    sum_in_packed.allocate(this->pb, max_arity, "sum_in_packed");
    count_in_packed.allocate(this->pb, max_arity, "count_in_packed");

    sum_in_packed_aux.allocate(this->pb, max_arity, "sum_in_packed_aux");
    count_in_packed_aux.allocate(this->pb, max_arity, "count_in_packed_aux");

    type_val_inner_product.allocate(this->pb, "type_val_inner_product");
    for (auto &msg : this->incoming_messages)
    {
        incoming_types.emplace_back(msg->type);
    }

    compute_type_val_inner_product.reset(new inner_product_gadget<FieldT>(this->pb, incoming_types, sum_in_packed, type_val_inner_product, "compute_type_val_inner_product"));

    unpack_sum_out.reset(new packing_gadget<FieldT>(this->pb, std::dynamic_pointer_cast<tally_pcd_message_variable<FieldT> >(this->outgoing_message)->sum_bits, sum_out_packed, "pack_sum_out"));
    unpack_count_out.reset(new packing_gadget<FieldT>(this->pb, std::dynamic_pointer_cast<tally_pcd_message_variable<FieldT> >(this->outgoing_message)->count_bits, count_out_packed, "pack_count_out"));

    for (size_t i = 0; i < max_arity; ++i)
    {
        pack_sum_in.emplace_back(packing_gadget<FieldT>(this->pb, std::dynamic_pointer_cast<tally_pcd_message_variable<FieldT> >(this->incoming_messages[i])->sum_bits, sum_in_packed[i], FMT("", "pack_sum_in_%zu", i)));
        pack_count_in.emplace_back(packing_gadget<FieldT>(this->pb, std::dynamic_pointer_cast<tally_pcd_message_variable<FieldT> >(this->incoming_messages[i])->sum_bits, count_in_packed[i], FMT("", "pack_count_in_%zu", i)));
    }

    arity_indicators.allocate(this->pb, max_arity+1, "arity_indicators");
}

template<typename FieldT>
void tally_cp_handler<FieldT>::generate_r1cs_constraints()
{
    unpack_sum_out->generate_r1cs_constraints(true);
    unpack_count_out->generate_r1cs_constraints(true);

    for (size_t i = 0; i < this->max_arity; ++i)
    {
        pack_sum_in[i].generate_r1cs_constraints(true);
        pack_count_in[i].generate_r1cs_constraints(true);
    }

    for (size_t i = 0; i < this->max_arity; ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(incoming_types[i], sum_in_packed_aux[i], sum_in_packed[i]), FMT("", "initial_sum_%zu_is_zero", i));
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(incoming_types[i], count_in_packed_aux[i], count_in_packed[i]), FMT("", "initial_sum_%zu_is_zero", i));
    }

    /* constrain arity indicator variables so that arity_indicators[arity] = 1 and arity_indicators[i] = 0 for any other i */
    for (size_t i = 0; i < this->max_arity; ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(this->arity - FieldT(i), arity_indicators[i], 0), FMT("", "arity_indicators_%zu", i));
    }

    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, pb_sum<FieldT>(arity_indicators), 1), "arity_indicators");

    /* require that types of messages that are past arity (i.e. unbound wires) carry 0 */
    for (size_t i = 0; i < this->max_arity; ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(0 + pb_sum<FieldT>(pb_variable_array<FieldT>(arity_indicators.begin(), arity_indicators.begin() + i)), incoming_types[i], 0), FMT("", "unbound_types_%zu", i));
    }

    /* sum_out = local_data + \sum_i type[i] * sum_in[i] */
    compute_type_val_inner_product->generate_r1cs_constraints();
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, type_val_inner_product + std::dynamic_pointer_cast<tally_pcd_local_data_variable<FieldT> >(this->local_data)->summand, sum_out_packed), "update_sum");

    /* count_out = 1 + \sum_i count_in[i] */
    this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1, 1 + pb_sum<FieldT>(count_in_packed), count_out_packed), "update_count");
}

template<typename FieldT>
void tally_cp_handler<FieldT>::generate_r1cs_witness(const std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > &incoming_messages,
                                                     const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data)
{
    base_handler::generate_r1cs_witness(incoming_messages, local_data);

    for (size_t i = 0; i < this->max_arity; ++i)
    {
        pack_sum_in[i].generate_r1cs_witness_from_bits();
        pack_count_in[i].generate_r1cs_witness_from_bits();

        if (!this->pb.val(incoming_types[i]).is_zero())
        {
            this->pb.val(sum_in_packed_aux[i]) = this->pb.val(sum_in_packed[i]) * this->pb.val(incoming_types[i]).inverse();
            this->pb.val(count_in_packed_aux[i]) = this->pb.val(count_in_packed[i]) * this->pb.val(incoming_types[i]).inverse();
        }
    }

    for (size_t i = 0; i < this->max_arity + 1; ++i)
    {
        this->pb.val(arity_indicators[i]) = (incoming_messages.size() == i ? FieldT::one() : FieldT::zero());
    }

    compute_type_val_inner_product->generate_r1cs_witness();
    this->pb.val(sum_out_packed) = this->pb.val(std::dynamic_pointer_cast<tally_pcd_local_data_variable<FieldT> >(this->local_data)->summand) + this->pb.val(type_val_inner_product);

    this->pb.val(count_out_packed) = FieldT::one();
    for (size_t i = 0; i < this->max_arity; ++i)
    {
        this->pb.val(count_out_packed) += this->pb.val(count_in_packed[i]);
    }

    unpack_sum_out->generate_r1cs_witness_from_packed();
    unpack_count_out->generate_r1cs_witness_from_packed();
}

template<typename FieldT>
std::shared_ptr<r1cs_pcd_message<FieldT> > tally_cp_handler<FieldT>::get_base_case_message() const
{
    const size_t type = 0;
    const size_t sum = 0;
    const size_t count = 0;

    std::shared_ptr<r1cs_pcd_message<FieldT> > result;
    result.reset(new tally_pcd_message<FieldT>(type, wordsize, sum, count));

    return result;
}

} // libsnark

#endif // TALLY_CP_TCC_

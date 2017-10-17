/** @file
 *****************************************************************************

 Implementation of interfaces for the knapsack gadget.

 See knapsack_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef KNAPSACK_GADGET_TCC_
#define KNAPSACK_GADGET_TCC_

#include "algebra/fields/field_utils.hpp"
#include "common/rng.hpp"

namespace libsnark {

template<typename FieldT>
std::vector<FieldT> knapsack_CRH_with_field_out_gadget<FieldT>::knapsack_coefficients;
template<typename FieldT>
size_t knapsack_CRH_with_field_out_gadget<FieldT>::num_cached_coefficients;

template<typename FieldT>
knapsack_CRH_with_field_out_gadget<FieldT>::knapsack_CRH_with_field_out_gadget(protoboard<FieldT> &pb,
                                                                               const size_t input_len,
                                                                               const block_variable<FieldT> &input_block,
                                                                               const pb_linear_combination_array<FieldT> &output,
                                                                               const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    input_len(input_len),
    dimension(knapsack_dimension<FieldT>::dimension),
    input_block(input_block),
    output(output)
{
    assert(input_block.bits.size() == input_len);
    if (num_cached_coefficients < dimension * input_len)
    {
        sample_randomness(input_len);
    }
    assert(output.size() == this->get_digest_len());
}

template<typename FieldT>
void knapsack_CRH_with_field_out_gadget<FieldT>::generate_r1cs_constraints()
{
    for (size_t i = 0; i < dimension; ++i)
    {
        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(1,
                                                             pb_coeff_sum<FieldT>(input_block.bits,
                                                                                  std::vector<FieldT>(knapsack_coefficients.begin() + input_len * i,
                                                                                                      knapsack_coefficients.begin() + input_len * (i+1))),
                                                             output[i]), FMT(this->annotation_prefix, " knapsack_%zu", i));
    }
}

template<typename FieldT>
void knapsack_CRH_with_field_out_gadget<FieldT>::generate_r1cs_witness()
{
    const bit_vector input = input_block.get_block();

    for (size_t i = 0; i < dimension; ++i)
    {
        FieldT sum = FieldT::zero();
        for (size_t k = 0; k < input_len; ++k)
        {
            if (input[k])
            {
                sum += knapsack_coefficients[input_len*i + k];
            }
        }

        this->pb.lc_val(output[i]) = sum;
    }
}

template<typename FieldT>
size_t knapsack_CRH_with_field_out_gadget<FieldT>::get_digest_len()
{
    return knapsack_dimension<FieldT>::dimension;
}

template<typename FieldT>
size_t knapsack_CRH_with_field_out_gadget<FieldT>::get_block_len()
{
    return 0;
}

template<typename FieldT>
std::vector<FieldT> knapsack_CRH_with_field_out_gadget<FieldT>::get_hash(const bit_vector &input)
{
    const size_t dimension = knapsack_dimension<FieldT>::dimension;
    if (num_cached_coefficients < dimension * input.size())
    {
        sample_randomness(input.size());
    }

    std::vector<FieldT> result(dimension, FieldT::zero());

    for (size_t i = 0; i < dimension; ++i)
    {
        for (size_t k = 0; k < input.size(); ++k)
        {
            if (input[k])
            {
                result[i] += knapsack_coefficients[input.size()*i + k];
            }
        }
    }

    return result;
}

template<typename FieldT>
size_t knapsack_CRH_with_field_out_gadget<FieldT>::expected_constraints()
{
    return knapsack_dimension<FieldT>::dimension;
}

template<typename FieldT>
void knapsack_CRH_with_field_out_gadget<FieldT>::sample_randomness(const size_t input_len)
{
    const size_t num_coefficients = knapsack_dimension<FieldT>::dimension * input_len;
    if (num_coefficients > num_cached_coefficients)
    {
        knapsack_coefficients.resize(num_coefficients);
        for (size_t i = num_cached_coefficients; i < num_coefficients; ++i)
        {
            knapsack_coefficients[i] = SHA512_rng<FieldT>(i);
        }
        num_cached_coefficients = num_coefficients;
    }
}

template<typename FieldT>
knapsack_CRH_with_bit_out_gadget<FieldT>::knapsack_CRH_with_bit_out_gadget(protoboard<FieldT> &pb,
                                                                           const size_t input_len,
                                                                           const block_variable<FieldT> &input_block,
                                                                           const digest_variable<FieldT> &output_digest,
                                                                           const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix),
    input_len(input_len),
    dimension(knapsack_dimension<FieldT>::dimension),
    input_block(input_block),
    output_digest(output_digest)
{
    assert(output_digest.bits.size() == this->get_digest_len());

    output.resize(dimension);

    for (size_t i = 0; i < dimension; ++i)
    {
        output[i].assign(pb, pb_packing_sum<FieldT>(pb_variable_array<FieldT>(output_digest.bits.begin() + i * FieldT::size_in_bits(),
                                                                              output_digest.bits.begin() + (i + 1) * FieldT::size_in_bits())));
    }

    hasher.reset(new knapsack_CRH_with_field_out_gadget<FieldT>(pb, input_len, input_block, output, FMT(annotation_prefix, " hasher")));
}


template<typename FieldT>
void knapsack_CRH_with_bit_out_gadget<FieldT>::generate_r1cs_constraints(const bool enforce_bitness)
{
    hasher->generate_r1cs_constraints();

    if (enforce_bitness)
    {
        for (size_t k = 0; k < output_digest.bits.size(); ++k)
        {
            generate_boolean_r1cs_constraint<FieldT>(this->pb, output_digest.bits[k], FMT(this->annotation_prefix, " output_digest_%zu", k));
        }
    }
}

template<typename FieldT>
void knapsack_CRH_with_bit_out_gadget<FieldT>::generate_r1cs_witness()
{
    hasher->generate_r1cs_witness();

    /* do unpacking in place */
    const bit_vector input = input_block.bits.get_bits(this->pb);
    for (size_t i = 0; i < dimension; ++i)
    {
        pb_variable_array<FieldT> va(output_digest.bits.begin() + i * FieldT::size_in_bits(),
                                     output_digest.bits.begin() + (i + 1) * FieldT::size_in_bits());
        va.fill_with_bits_of_field_element(this->pb, this->pb.lc_val(output[i]));
    }
}

template<typename FieldT>
size_t knapsack_CRH_with_bit_out_gadget<FieldT>::get_digest_len()
{
    return knapsack_dimension<FieldT>::dimension * FieldT::size_in_bits();
}

template<typename FieldT>
size_t knapsack_CRH_with_bit_out_gadget<FieldT>::get_block_len()
{
     return 0;
}

template<typename FieldT>
bit_vector knapsack_CRH_with_bit_out_gadget<FieldT>::get_hash(const bit_vector &input)
{
    const std::vector<FieldT> hash_elems = knapsack_CRH_with_field_out_gadget<FieldT>::get_hash(input);
    hash_value_type result;

    for (const FieldT &elt : hash_elems)
    {
        bit_vector elt_bits = convert_field_element_to_bit_vector<FieldT>(elt);
        result.insert(result.end(), elt_bits.begin(), elt_bits.end());
    }

    return result;
}

template<typename FieldT>
size_t knapsack_CRH_with_bit_out_gadget<FieldT>::expected_constraints(const bool enforce_bitness)
{
    const size_t hasher_constraints = knapsack_CRH_with_field_out_gadget<FieldT>::expected_constraints();
    const size_t bitness_constraints = (enforce_bitness ? get_digest_len() : 0);
    return hasher_constraints + bitness_constraints;
}

template<typename FieldT>
void knapsack_CRH_with_bit_out_gadget<FieldT>::sample_randomness(const size_t input_len)
{
    knapsack_CRH_with_field_out_gadget<FieldT>::sample_randomness(input_len);
}

template<typename FieldT>
void test_knapsack_CRH_with_bit_out_gadget_internal(const size_t dimension, const bit_vector &input_bits, const bit_vector &digest_bits)
{
    assert(knapsack_dimension<FieldT>::dimension == dimension);
    knapsack_CRH_with_bit_out_gadget<FieldT>::sample_randomness(input_bits.size());
    protoboard<FieldT> pb;

    block_variable<FieldT> input_block(pb, input_bits.size(), "input_block");
    digest_variable<FieldT> output_digest(pb, knapsack_CRH_with_bit_out_gadget<FieldT>::get_digest_len(), "output_digest");
    knapsack_CRH_with_bit_out_gadget<FieldT> H(pb, input_bits.size(), input_block, output_digest, "H");

    input_block.generate_r1cs_witness(input_bits);
    H.generate_r1cs_constraints();
    H.generate_r1cs_witness();

    assert(output_digest.get_digest().size() == digest_bits.size());
    assert(pb.is_satisfied());

    const size_t num_constraints = pb.num_constraints();
    const size_t expected_constraints = knapsack_CRH_with_bit_out_gadget<FieldT>::expected_constraints();
    assert(num_constraints == expected_constraints);
}

} // libsnark

#endif // KNAPSACK_GADGET_TCC_

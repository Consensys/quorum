/** @file
 *****************************************************************************

 Declaration of interfaces for the knapsack gadget.

 The gadget checks the correct execution of a knapsack (modular subset-sum) over
 the field specified in the template parameter. With suitable choices of parameters
 such knapsacks are collision-resistant hashes (CRHs). See \[Ajt96] and \[GGH96].

 Given two positive integers m (the input length) and d (the dimension),
 and a matrix M over the field F and of dimension dxm, the hash H_M maps {0,1}^m
 to F^d by sending x to M*x. Security of the function (very roughly) depends on
 d*log(|F|).

 Below, we give two different gadgets:
 - knapsack_CRH_with_field_out_gadget, which verifies H_M
 - knapsack_CRH_with_bit_out_gadget, which verifies H_M when its output is "expanded" to bits.
 In both cases, a method ("sample_randomness") allows to sample M.

 The parameter d (the dimension) is fixed at compile time in the struct
 knapsack_dimension below. The parameter m (the input lenght) can be chosen
 at run time (in either gadget).


 References:

 \[Ajt96]:
 "Generating hard instances of lattice problems",
 Miklos Ajtai,
 STOC 1996

 \[GGH96]:
 "Collision-free hashing from lattice problems",
 Oded Goldreich, Shafi Goldwasser, Shai Halevi,
 ECCC TR95-042

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef KNAPSACK_GADGET_HPP_
#define KNAPSACK_GADGET_HPP_

#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "gadgetlib1/gadgets/hashes/hash_io.hpp"
#include "common/data_structures/merkle_tree.hpp"

namespace libsnark {

/************************** Choice of dimension ******************************/

template<typename FieldT>
struct knapsack_dimension {
    // the size of FieldT should be (approximately) at least 200 bits
    static const size_t dimension = 1;
};

/*********************** Knapsack with field output **************************/

template<typename FieldT>
class knapsack_CRH_with_field_out_gadget : public gadget<FieldT> {
private:
    static std::vector<FieldT> knapsack_coefficients;
    static size_t num_cached_coefficients;

public:
    size_t input_len;
    size_t dimension;

    block_variable<FieldT> input_block;
    pb_linear_combination_array<FieldT> output;

    knapsack_CRH_with_field_out_gadget(protoboard<FieldT> &pb,
                                       const size_t input_len,
                                       const block_variable<FieldT> &input_block,
                                       const pb_linear_combination_array<FieldT> &output,
                                       const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();

    static size_t get_digest_len();
    static size_t get_block_len(); /* return 0 as block length, as the hash function is variable-input */
    static std::vector<FieldT> get_hash(const bit_vector &input);
    static void sample_randomness(const size_t input_len);

    /* for debugging */
    static size_t expected_constraints();
};

/********************** Knapsack with binary output **************************/

template<typename FieldT>
class knapsack_CRH_with_bit_out_gadget : public gadget<FieldT> {
public:
    typedef bit_vector hash_value_type;
    typedef merkle_authentication_path merkle_authentication_path_type;

    size_t input_len;
    size_t dimension;

    pb_linear_combination_array<FieldT> output;

    std::shared_ptr<knapsack_CRH_with_field_out_gadget<FieldT> > hasher;

    block_variable<FieldT> input_block;
    digest_variable<FieldT> output_digest;

    knapsack_CRH_with_bit_out_gadget(protoboard<FieldT> &pb,
                                     const size_t input_len,
                                     const block_variable<FieldT> &input_block,
                                     const digest_variable<FieldT> &output_digest,
                                     const std::string &annotation_prefix);
    void generate_r1cs_constraints(const bool enforce_bitness=true);
    void generate_r1cs_witness();

    static size_t get_digest_len();
    static size_t get_block_len(); /* return 0 as block length, as the hash function is variable-input */
    static hash_value_type get_hash(const bit_vector &input);
    static void sample_randomness(const size_t input_len);

    /* for debugging */
    static size_t expected_constraints(const bool enforce_bitness=true);
};

template<typename FieldT>
void test_knapsack_CRH_with_bit_out_gadget();

} // libsnark

#include "gadgetlib1/gadgets/hashes/knapsack/knapsack_gadget.tcc"

#endif // KNAPSACK_GADGET_HPP_

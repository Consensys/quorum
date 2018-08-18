// Copyright 2017 Zerocoin Electric Coin Company LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "common/default_types/r1cs_ppzksnark_pp.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp"
#include "gadgetlib1/gadgets/hashes/sha256/sha256_gadget.hpp"
#include "gadgetlib1/gadgets/merkle_tree/merkle_tree_check_read_gadget.hpp"

#include <boost/foreach.hpp>

template<typename FieldT>
pb_variable_array<FieldT> from_bits(std::vector<bool> bits, pb_variable<FieldT>& ZERO) {
    pb_variable_array<FieldT> acc;

    BOOST_FOREACH(bool bit, bits) {
        acc.emplace_back(bit ? ONE : ZERO);
    }

    return acc;
}

std::vector<unsigned char> convertIntToVectorLE(const uint64_t val_int) {
    std::vector<unsigned char> bytes;

    for(size_t i = 0; i < 8; i++) {
        bytes.push_back(val_int >> (i * 8));
    }

    return bytes;
}

// Convert bytes into boolean vector. (MSB to LSB)
std::vector<bool> convertBytesVectorToVector(const std::vector<unsigned char>& bytes) {
    std::vector<bool> ret;
    ret.resize(bytes.size() * 8);

    unsigned char c;
    for (size_t i = 0; i < bytes.size(); i++) {
        c = bytes.at(i);
        for (size_t j = 0; j < 8; j++) {
            ret.at((i*8)+j) = (c >> (7-j)) & 1;
        }
    }

    return ret;
}

// Convert boolean vector (big endian) to integer
uint64_t convertVectorToInt(const std::vector<bool>& v) {
    if (v.size() > 64) {
        throw std::length_error ("boolean vector can't be larger than 64 bits");
    }

    uint64_t result = 0;
    for (size_t i=0; i<v.size();i++) {
        if (v.at(i)) {
            result |= (uint64_t)1 << ((v.size() - 1) - i);
        }
    }

    return result;
}

std::vector<bool> uint64_to_bool_vector(uint64_t input) {
    auto num_bv = convertIntToVectorLE(input);

    return convertBytesVectorToVector(num_bv);
}

template<typename FieldT>
class KeyHasher : gadget<FieldT> {
private:
    std::shared_ptr<block_variable<FieldT>> block;
    std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher;

public:
    KeyHasher(
        protoboard<FieldT> &pb,
        pb_variable<FieldT>& ZERO,
        pb_variable_array<FieldT> sk,
        std::shared_ptr<digest_variable<FieldT>> pk
    ) : gadget<FieldT>(pb) {
        pb_linear_combination_array<FieldT> IV = SHA256_default_IV(pb);

        pb_variable_array<FieldT> length_padding =
            from_bits({
                // padding
                1,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,1,
                0,0,0,0,0,0,0,0
        }, ZERO);

        block.reset(new block_variable<FieldT>(pb, {
            sk,
            length_padding
        }, ""));

        hasher.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV,
            block->bits,
            *pk,
        ""));
    }

    void generate_r1cs_constraints() {
        hasher->generate_r1cs_constraints();
    }

    void generate_r1cs_witness() {
        hasher->generate_r1cs_witness();
    }
};

template<typename FieldT>
class SendNullifier : gadget<FieldT> {
private:
	std::shared_ptr<block_variable<FieldT>> block;
	std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher;

public:
	SendNullifier(
		protoboard<FieldT> &pb,
		pb_variable<FieldT>& ZERO,
		pb_variable_array<FieldT> rho,
		std::shared_ptr<digest_variable<FieldT>> result
	) : gadget<FieldT>(pb) {
		pb_linear_combination_array<FieldT> IV = SHA256_default_IV(pb);

		pb_variable_array<FieldT> discriminants;
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);

        pb_variable_array<FieldT> length_padding =
            from_bits({
                // padding
                1,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,1,
                0,0,0,0,1,0,0,0
		}, ZERO);

		block.reset(new block_variable<FieldT>(pb, {
            discriminants,
            rho,
            length_padding
        }, ""));

        hasher.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV,
            block->bits,
            *result,
        ""));
	}

	void generate_r1cs_constraints() {
        hasher->generate_r1cs_constraints();
    }

    void generate_r1cs_witness() {
        hasher->generate_r1cs_witness();
    }
};

template<typename FieldT>
class SpendNullifier : gadget<FieldT> {
private:
    std::shared_ptr<block_variable<FieldT>> block1;
    std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher1;
    std::shared_ptr<digest_variable<FieldT>> intermediate;
    std::shared_ptr<block_variable<FieldT>> block2;
    std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher2;

public:
    SpendNullifier(
        protoboard<FieldT> &pb,
        pb_variable<FieldT>& ZERO,
        pb_variable_array<FieldT> rho,
        pb_variable_array<FieldT> sk,
        std::shared_ptr<digest_variable<FieldT>> result
    ) : gadget<FieldT>(pb) {
        pb_linear_combination_array<FieldT> IV = SHA256_default_IV(pb);

        pb_variable_array<FieldT> discriminants;
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ONE);

        block1.reset(new block_variable<FieldT>(pb, {
            discriminants,
            rho,
            pb_variable_array<FieldT>(sk.begin(), sk.begin() + 248)
        }, ""));

        intermediate.reset(new digest_variable<FieldT>(pb, 256, ""));

        hasher1.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV,
            block1->bits,
            *intermediate,
        ""));

        pb_variable_array<FieldT> length_padding =
            from_bits({
                // padding
                1,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,1,0,
                0,0,0,0,1,0,0,0
        }, ZERO);

        block2.reset(new block_variable<FieldT>(pb, {
            pb_variable_array<FieldT>(sk.begin() + 248, sk.end()),
            length_padding
        }, ""));

        pb_linear_combination_array<FieldT> IV2(intermediate->bits);

        hasher2.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV2,
            block2->bits,
            *result,
        ""));
    }

    void generate_r1cs_constraints() {
        hasher1->generate_r1cs_constraints();
        hasher2->generate_r1cs_constraints();
    }

    void generate_r1cs_witness() {
        hasher1->generate_r1cs_witness();
        hasher2->generate_r1cs_witness();
    }
};

template<typename FieldT>
class SpendNullifierAuthenticated : gadget<FieldT> {
private:
    std::shared_ptr<block_variable<FieldT>> block1;
    std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher1;
    std::shared_ptr<digest_variable<FieldT>> intermediate;
    std::shared_ptr<block_variable<FieldT>> block2;
    std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher2;

public:
    SpendNullifierAuthenticated(
        protoboard<FieldT> &pb,
        pb_variable<FieldT>& ZERO,
        pb_variable_array<FieldT> rho,
        pb_variable_array<FieldT> sk,
        pb_variable_array<FieldT> addr,
        std::shared_ptr<digest_variable<FieldT>> result
    ) : gadget<FieldT>(pb) {
        pb_linear_combination_array<FieldT> IV = SHA256_default_IV(pb);

        pb_variable_array<FieldT> discriminants;
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ZERO);
        discriminants.emplace_back(ONE);

        block1.reset(new block_variable<FieldT>(pb, {
            discriminants,
            rho,
            pb_variable_array<FieldT>(sk.begin(), sk.begin() + 248)
        }, ""));

        intermediate.reset(new digest_variable<FieldT>(pb, 256, ""));

        hasher1.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV,
            block1->bits,
            *intermediate,
        ""));

        pb_variable_array<FieldT> length_padding =
            from_bits({
                // padding
                1,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,1,0,
                1,0,1,0,1,0,0,0
        }, ZERO);

        block2.reset(new block_variable<FieldT>(pb, {
            pb_variable_array<FieldT>(sk.begin() + 248, sk.end()),
            addr,
            length_padding
        }, ""));

        pb_linear_combination_array<FieldT> IV2(intermediate->bits);

        hasher2.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV2,
            block2->bits,
            *result,
        ""));
    }

    void generate_r1cs_constraints() {
        hasher1->generate_r1cs_constraints();
        hasher2->generate_r1cs_constraints();
    }

    void generate_r1cs_witness() {
        hasher1->generate_r1cs_witness();
        hasher2->generate_r1cs_witness();
    }
};

template<typename FieldT>
class NoteCommitment : gadget<FieldT> {
private:
	std::shared_ptr<block_variable<FieldT>> block1;
	std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher1;
	std::shared_ptr<digest_variable<FieldT>> intermediate;
	std::shared_ptr<block_variable<FieldT>> block2;
	std::shared_ptr<sha256_compression_function_gadget<FieldT>> hasher2;

public:
	NoteCommitment(
		protoboard<FieldT> &pb,
		pb_variable<FieldT>& ZERO,
		pb_variable_array<FieldT> rho,
		pb_variable_array<FieldT> pk,
		pb_variable_array<FieldT> value,
		std::shared_ptr<digest_variable<FieldT>> result
	) : gadget<FieldT>(pb) {
		pb_linear_combination_array<FieldT> IV = SHA256_default_IV(pb);

		block1.reset(new block_variable<FieldT>(pb, {
            rho,
            pk
        }, ""));

        intermediate.reset(new digest_variable<FieldT>(pb, 256, ""));

        hasher1.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV,
            block1->bits,
            *intermediate,
        ""));

        pb_variable_array<FieldT> length_padding =
            from_bits({
                // padding
                1,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,0,0,
                0,0,0,0,0,0,1,0,
                0,1,0,0,0,0,0,0
		}, ZERO);

        block2.reset(new block_variable<FieldT>(pb, {
            value,
            length_padding
        }, ""));

        pb_linear_combination_array<FieldT> IV2(intermediate->bits);

        hasher2.reset(new sha256_compression_function_gadget<FieldT>(
            pb,
            IV2,
            block2->bits,
            *result,
        ""));
	}

	void generate_r1cs_constraints() {
        hasher1->generate_r1cs_constraints();
        hasher2->generate_r1cs_constraints();
    }

    void generate_r1cs_witness() {
        hasher1->generate_r1cs_witness();
        hasher2->generate_r1cs_witness();
    }
};

template<typename FieldT>
class ShieldingCircuit : gadget<FieldT> {
private:
    // Verifier inputs
    pb_variable_array<FieldT> zk_packed_inputs;
    pb_variable_array<FieldT> zk_unpacked_inputs;
    std::shared_ptr<multipacking_gadget<FieldT>> unpacker;

    // SHA256(0x00 | rho)
    std::shared_ptr<digest_variable<FieldT>> send_nullifier;
    // The note commitment: SHA256(rho | pk | value)
    std::shared_ptr<digest_variable<FieldT>> cm;
    // 64-bit value
    pb_variable_array<FieldT> value;

    // Aux inputs
    pb_variable<FieldT> ZERO;
    std::shared_ptr<digest_variable<FieldT>> rho;
    std::shared_ptr<digest_variable<FieldT>> pk;

    // Note commitment hasher
    std::shared_ptr<NoteCommitment<FieldT>> cm_hasher;

    // Send nullifier hasher
    std::shared_ptr<SendNullifier<FieldT>> nf_hasher;

public:
    ShieldingCircuit(protoboard<FieldT> &pb) : gadget<FieldT>(pb) {
    	// Inputs
    	{
	    	zk_packed_inputs.allocate(pb, verifying_field_element_size());
	    	pb.set_input_sizes(verifying_field_element_size());

	    	send_nullifier.reset(new digest_variable<FieldT>(pb, 256, ""));
	    	zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), send_nullifier->bits.begin(), send_nullifier->bits.end());

	    	cm.reset(new digest_variable<FieldT>(pb, 256, ""));
	    	zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), cm->bits.begin(), cm->bits.end());

	    	value.allocate(pb, 64, "");
	    	zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), value.begin(), value.end());

	    	assert(zk_unpacked_inputs.size() == verifying_input_bit_size());

	    	unpacker.reset(new multipacking_gadget<FieldT>(
	            pb,
	            zk_unpacked_inputs,
	            zk_packed_inputs,
	            FieldT::capacity(),
	            "unpacker"
	        ));
	    }

	    // Aux
	    ZERO.allocate(pb);
	    rho.reset(new digest_variable<FieldT>(pb, 256, ""));
	    pk.reset(new digest_variable<FieldT>(pb, 256, ""));

	    cm_hasher.reset(new NoteCommitment<FieldT>(pb, ZERO, rho->bits, pk->bits, value, cm));
	    nf_hasher.reset(new SendNullifier<FieldT>(pb, ZERO, rho->bits, send_nullifier));
    }

    void generate_r1cs_constraints() {
        unpacker->generate_r1cs_constraints(true);
        generate_r1cs_equals_const_constraint<FieldT>(this->pb, ZERO, FieldT::zero(), "ZERO");

        rho->generate_r1cs_constraints();
        pk->generate_r1cs_constraints();

        cm_hasher->generate_r1cs_constraints();
        nf_hasher->generate_r1cs_constraints();
    }

    void generate_r1cs_witness(
        const std::vector<unsigned char>& witness_rho,
        const std::vector<unsigned char>& witness_pk,
        uint64_t witness_value
    ) {
        this->pb.val(ZERO) = FieldT::zero();

        rho->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_rho)
        );

        pk->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_pk)
        );

        value.fill_with_bits(
            this->pb,
            uint64_to_bool_vector(witness_value)
        );

        cm_hasher->generate_r1cs_witness();
        nf_hasher->generate_r1cs_witness();

        unpacker->generate_r1cs_witness_from_bits();
    }

    static r1cs_primary_input<FieldT> witness_map(
        const std::vector<unsigned char> &witness_nf,
        const std::vector<unsigned char> &witness_cm,
        uint64_t witness_value
    ) {
        std::vector<bool> verify_inputs;

        std::vector<bool> nf_bits = convertBytesVectorToVector(witness_nf);
        std::vector<bool> cm_bits = convertBytesVectorToVector(witness_cm);
        std::vector<bool> value_bits = uint64_to_bool_vector(witness_value);

        verify_inputs.insert(verify_inputs.end(), nf_bits.begin(), nf_bits.end());
        verify_inputs.insert(verify_inputs.end(), cm_bits.begin(), cm_bits.end());
        verify_inputs.insert(verify_inputs.end(), value_bits.begin(), value_bits.end());

        assert(verify_inputs.size() == verifying_input_bit_size());
        auto verify_field_elements = pack_bit_vector_into_field_element_vector<FieldT>(verify_inputs);
        assert(verify_field_elements.size() == verifying_field_element_size());
        return verify_field_elements;
    }

    static size_t verifying_field_element_size() {
        return div_ceil(verifying_input_bit_size(), FieldT::capacity());
    }

    static size_t verifying_input_bit_size() {
        size_t acc = 0;

        acc += 256; // the nullifier
        acc += 256; // the note commitment
        acc += 64; // the value of the note

        return acc;
    }
};

#define INCREMENTAL_MERKLE_TREE_DEPTH 29

template<typename FieldT>
class merkle_tree_gadget : gadget<FieldT> {
private:
    typedef sha256_two_to_one_hash_gadget<FieldT> sha256_gadget;

    pb_variable_array<FieldT> positions;
    std::shared_ptr<merkle_authentication_path_variable<FieldT, sha256_gadget>> authvars;
    std::shared_ptr<merkle_tree_check_read_gadget<FieldT, sha256_gadget>> auth;

public:
    merkle_tree_gadget(
        protoboard<FieldT>& pb,
        digest_variable<FieldT> leaf,
        digest_variable<FieldT> root,
        pb_variable<FieldT>& enforce
    ) : gadget<FieldT>(pb) {
        positions.allocate(pb, INCREMENTAL_MERKLE_TREE_DEPTH);
        authvars.reset(new merkle_authentication_path_variable<FieldT, sha256_gadget>(
            pb, INCREMENTAL_MERKLE_TREE_DEPTH, "auth"
        ));
        auth.reset(new merkle_tree_check_read_gadget<FieldT, sha256_gadget>(
            pb,
            INCREMENTAL_MERKLE_TREE_DEPTH,
            positions,
            leaf,
            root,
            *authvars,
            enforce,
            ""
        ));
    }

    void generate_r1cs_constraints() {
        for (size_t i = 0; i < INCREMENTAL_MERKLE_TREE_DEPTH; i++) {
            generate_boolean_r1cs_constraint<FieldT>(
                this->pb,
                positions[i],
                "boolean_positions"
            );
        }

        authvars->generate_r1cs_constraints();
        auth->generate_r1cs_constraints();
    }

    void generate_r1cs_witness(
        size_t path_index,
        const std::vector<std::vector<bool>>& authentication_path
    ) {
        positions.fill_with_bits_of_ulong(this->pb, path_index);

        authvars->generate_r1cs_witness(path_index, authentication_path);
        auth->generate_r1cs_witness();
    }
};

template<typename FieldT>
class UnshieldingCircuit : gadget<FieldT> {
private:
    // Verifier inputs
    pb_variable_array<FieldT> zk_packed_inputs;
    pb_variable_array<FieldT> zk_unpacked_inputs;
    std::shared_ptr<multipacking_gadget<FieldT>> unpacker;

    // 64-bit value
    pb_variable_array<FieldT> value;

    // Aux inputs
    pb_variable<FieldT> ZERO;
    std::shared_ptr<digest_variable<FieldT>> cm; // Note commitment
    std::shared_ptr<digest_variable<FieldT>> rho;
    std::shared_ptr<digest_variable<FieldT>> pk;
    std::shared_ptr<digest_variable<FieldT>> sk;

    // Key hasher
    std::shared_ptr<KeyHasher<FieldT>> key_hasher;

    // Note commitment hasher
    std::shared_ptr<NoteCommitment<FieldT>> cm_hasher;

    // Spend nullifier hasher
    std::shared_ptr<SpendNullifierAuthenticated<FieldT>> nf_hasher;

    // Merkle tree lookup
    std::shared_ptr<merkle_tree_gadget<FieldT>> merkle_lookup;

public:
    // The anchor of the tree
    std::shared_ptr<digest_variable<FieldT>> anchor;

    // SHA256(0x01 | rho | sk | addr)
    std::shared_ptr<digest_variable<FieldT>> spend_nullifier;

    std::shared_ptr<digest_variable<FieldT>> addr;

    UnshieldingCircuit(protoboard<FieldT> &pb) : gadget<FieldT>(pb) {
        // Inputs
        {
            zk_packed_inputs.allocate(pb, verifying_field_element_size());
            pb.set_input_sizes(verifying_field_element_size());

            spend_nullifier.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), spend_nullifier->bits.begin(), spend_nullifier->bits.end());

            anchor.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), anchor->bits.begin(), anchor->bits.end());

            addr.reset(new digest_variable<FieldT>(pb, 160, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), addr->bits.begin(), addr->bits.end());

            value.allocate(pb, 64, "");
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), value.begin(), value.end());

            assert(zk_unpacked_inputs.size() == verifying_input_bit_size());

            unpacker.reset(new multipacking_gadget<FieldT>(
                pb,
                zk_unpacked_inputs,
                zk_packed_inputs,
                FieldT::capacity(),
                "unpacker"
            ));
        }

        // Aux
        ZERO.allocate(pb);
        rho.reset(new digest_variable<FieldT>(pb, 256, ""));
        sk.reset(new digest_variable<FieldT>(pb, 256, ""));
        pk.reset(new digest_variable<FieldT>(pb, 256, ""));
        cm.reset(new digest_variable<FieldT>(pb, 256, ""));

        key_hasher.reset(new KeyHasher<FieldT>(pb, ZERO, sk->bits, pk));
        cm_hasher.reset(new NoteCommitment<FieldT>(pb, ZERO, rho->bits, pk->bits, value, cm));
        nf_hasher.reset(new SpendNullifierAuthenticated<FieldT>(pb, ZERO, rho->bits, sk->bits, addr->bits, spend_nullifier));
        auto test = ONE;
        merkle_lookup.reset(new merkle_tree_gadget<FieldT>(pb, *cm, *anchor, test));
    }

    void generate_r1cs_constraints() {
        unpacker->generate_r1cs_constraints(true);
        generate_r1cs_equals_const_constraint<FieldT>(this->pb, ZERO, FieldT::zero(), "ZERO");

        rho->generate_r1cs_constraints();
        sk->generate_r1cs_constraints();
        addr->generate_r1cs_constraints();

        key_hasher->generate_r1cs_constraints();
        cm_hasher->generate_r1cs_constraints();
        nf_hasher->generate_r1cs_constraints();
        merkle_lookup->generate_r1cs_constraints();
    }

    void generate_r1cs_witness(
        const std::vector<unsigned char>& witness_rho,
        const std::vector<unsigned char>& witness_sk,
        const std::vector<unsigned char>& witness_addr,
        uint64_t witness_value,
        size_t path_index,
        const std::vector<std::vector<bool>>& authentication_path
    ) {
        this->pb.val(ZERO) = FieldT::zero();

        rho->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_rho)
        );

        sk->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_sk)
        );

        addr->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_addr)
        );

        value.fill_with_bits(
            this->pb,
            uint64_to_bool_vector(witness_value)
        );

        key_hasher->generate_r1cs_witness();
        cm_hasher->generate_r1cs_witness();
        nf_hasher->generate_r1cs_witness();
        merkle_lookup->generate_r1cs_witness(path_index, authentication_path);

        unpacker->generate_r1cs_witness_from_bits();
    }

    static r1cs_primary_input<FieldT> witness_map(
        const std::vector<unsigned char> &witness_nf,
        const std::vector<unsigned char> &witness_anchor,
        const std::vector<unsigned char> &witness_addr,
        uint64_t witness_value
    ) {
        std::vector<bool> verify_inputs;

        std::vector<bool> nf_bits = convertBytesVectorToVector(witness_nf);
        std::vector<bool> anchor_bits = convertBytesVectorToVector(witness_anchor);
        std::vector<bool> addr_bits = convertBytesVectorToVector(witness_addr);
        std::vector<bool> value_bits = uint64_to_bool_vector(witness_value);

        verify_inputs.insert(verify_inputs.end(), nf_bits.begin(), nf_bits.end());
        verify_inputs.insert(verify_inputs.end(), anchor_bits.begin(), anchor_bits.end());
        verify_inputs.insert(verify_inputs.end(), addr_bits.begin(), addr_bits.end());
        verify_inputs.insert(verify_inputs.end(), value_bits.begin(), value_bits.end());

        assert(verify_inputs.size() == verifying_input_bit_size());
        auto verify_field_elements = pack_bit_vector_into_field_element_vector<FieldT>(verify_inputs);
        assert(verify_field_elements.size() == verifying_field_element_size());
        return verify_field_elements;
    }

    static size_t verifying_field_element_size() {
        return div_ceil(verifying_input_bit_size(), FieldT::capacity());
    }

    static size_t verifying_input_bit_size() {
        size_t acc = 0;

        acc += 256; // the nullifier
        acc += 256; // the anchor
        acc += 160; // the address
        acc += 64; // the value of the note

        return acc;
    }
};

template<typename T>
T swap_endianness_u64(T v) {
    if (v.size() != 64) {
        throw std::length_error("invalid bit length for 64-bit unsigned integer");
    }

    for (size_t i = 0; i < 4; i++) {
        for (size_t j = 0; j < 8; j++) {
            std::swap(v[i*8 + j], v[((7-i)*8)+j]);
        }
    }

    return v;
}

template<typename FieldT>
linear_combination<FieldT> packed_addition(pb_variable_array<FieldT> input) {
    auto input_swapped = swap_endianness_u64(input);

    return pb_packing_sum<FieldT>(pb_variable_array<FieldT>(
        input_swapped.rbegin(), input_swapped.rend()
    ));
}

template<typename FieldT>
class TransferCircuit : gadget<FieldT> {
private:
    // Verifier inputs
    pb_variable_array<FieldT> zk_packed_inputs;
    pb_variable_array<FieldT> zk_unpacked_inputs;
    std::shared_ptr<multipacking_gadget<FieldT>> unpacker;

    // Aux inputs
    pb_variable<FieldT> ZERO;

    // Verifier inputs
    std::shared_ptr<digest_variable<FieldT>> anchor;
    std::shared_ptr<digest_variable<FieldT>> spend_nullifier_input_1;
    std::shared_ptr<digest_variable<FieldT>> spend_nullifier_input_2;
    std::shared_ptr<digest_variable<FieldT>> send_nullifier_output_1;
    std::shared_ptr<digest_variable<FieldT>> send_nullifier_output_2;

    // Input stuff.
    std::shared_ptr<digest_variable<FieldT>> input_sk_1;
    std::shared_ptr<digest_variable<FieldT>> input_sk_2;
    std::shared_ptr<digest_variable<FieldT>> input_pk_1;
    std::shared_ptr<digest_variable<FieldT>> input_pk_2;
    std::shared_ptr<KeyHasher<FieldT>> key_hasher_1;
    std::shared_ptr<KeyHasher<FieldT>> key_hasher_2;
    std::shared_ptr<digest_variable<FieldT>> input_rho_1;
    std::shared_ptr<digest_variable<FieldT>> input_rho_2;
    std::shared_ptr<SpendNullifier<FieldT>> input_nf_hasher_1;
    std::shared_ptr<SpendNullifier<FieldT>> input_nf_hasher_2;
    pb_variable_array<FieldT> input_value_1;
    pb_variable_array<FieldT> input_value_2;
    std::shared_ptr<digest_variable<FieldT>> input_cm_1;
    std::shared_ptr<digest_variable<FieldT>> input_cm_2;
    std::shared_ptr<NoteCommitment<FieldT>> input_cm_hasher_1;
    std::shared_ptr<NoteCommitment<FieldT>> input_cm_hasher_2;
    pb_variable<FieldT> enforce_input_1;
    pb_variable<FieldT> enforce_input_2;
    std::shared_ptr<merkle_tree_gadget<FieldT>> merkle_lookup_1;
    std::shared_ptr<merkle_tree_gadget<FieldT>> merkle_lookup_2;

    // Output stuff.
    std::shared_ptr<digest_variable<FieldT>> output_cm_1;
    pb_variable_array<FieldT> output_value_1;
    std::shared_ptr<digest_variable<FieldT>> output_rho_1;
    std::shared_ptr<digest_variable<FieldT>> output_pk_1;
    std::shared_ptr<NoteCommitment<FieldT>> output_cm_hasher_1;
    std::shared_ptr<SendNullifier<FieldT>> output_nf_hasher_1;

    std::shared_ptr<digest_variable<FieldT>> output_cm_2;
    pb_variable_array<FieldT> output_value_2;
    std::shared_ptr<digest_variable<FieldT>> output_rho_2;
    std::shared_ptr<digest_variable<FieldT>> output_pk_2;
    std::shared_ptr<NoteCommitment<FieldT>> output_cm_hasher_2;
    std::shared_ptr<SendNullifier<FieldT>> output_nf_hasher_2;

public:

    TransferCircuit(protoboard<FieldT> &pb) : gadget<FieldT>(pb) {
        // Inputs
        {
            zk_packed_inputs.allocate(pb, verifying_field_element_size());
            pb.set_input_sizes(verifying_field_element_size());

            anchor.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), anchor->bits.begin(), anchor->bits.end());

            spend_nullifier_input_1.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), spend_nullifier_input_1->bits.begin(), spend_nullifier_input_1->bits.end());

            spend_nullifier_input_2.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), spend_nullifier_input_2->bits.begin(), spend_nullifier_input_2->bits.end());

            send_nullifier_output_1.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), send_nullifier_output_1->bits.begin(), send_nullifier_output_1->bits.end());

            send_nullifier_output_2.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), send_nullifier_output_2->bits.begin(), send_nullifier_output_2->bits.end());

            output_cm_1.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), output_cm_1->bits.begin(), output_cm_1->bits.end());

            output_cm_2.reset(new digest_variable<FieldT>(pb, 256, ""));
            zk_unpacked_inputs.insert(zk_unpacked_inputs.end(), output_cm_2->bits.begin(), output_cm_2->bits.end());

            assert(zk_unpacked_inputs.size() == verifying_input_bit_size());

            unpacker.reset(new multipacking_gadget<FieldT>(
                pb,
                zk_unpacked_inputs,
                zk_packed_inputs,
                FieldT::capacity(),
                "unpacker"
            ));
        }

        // Aux
        ZERO.allocate(pb);
        input_cm_1.reset(new digest_variable<FieldT>(pb, 256, ""));
        input_cm_2.reset(new digest_variable<FieldT>(pb, 256, ""));
        input_sk_1.reset(new digest_variable<FieldT>(pb, 256, ""));
        input_sk_2.reset(new digest_variable<FieldT>(pb, 256, ""));
        input_pk_1.reset(new digest_variable<FieldT>(pb, 256, ""));
        input_pk_2.reset(new digest_variable<FieldT>(pb, 256, ""));
        input_rho_1.reset(new digest_variable<FieldT>(pb, 256, ""));
        input_rho_2.reset(new digest_variable<FieldT>(pb, 256, ""));
        key_hasher_1.reset(new KeyHasher<FieldT>(pb, ZERO, input_sk_1->bits, input_pk_1));
        key_hasher_2.reset(new KeyHasher<FieldT>(pb, ZERO, input_sk_2->bits, input_pk_2));
        input_nf_hasher_1.reset(new SpendNullifier<FieldT>(pb, ZERO, input_rho_1->bits, input_sk_1->bits, spend_nullifier_input_1));
        input_nf_hasher_2.reset(new SpendNullifier<FieldT>(pb, ZERO, input_rho_2->bits, input_sk_2->bits, spend_nullifier_input_2));

        input_value_1.allocate(pb, 64, "");
        input_value_2.allocate(pb, 64, "");

        input_cm_hasher_1.reset(new NoteCommitment<FieldT>(pb, ZERO, input_rho_1->bits, input_pk_1->bits, input_value_1, input_cm_1));
        input_cm_hasher_2.reset(new NoteCommitment<FieldT>(pb, ZERO, input_rho_2->bits, input_pk_2->bits, input_value_2, input_cm_2));

        enforce_input_1.allocate(pb);
        enforce_input_2.allocate(pb);

        merkle_lookup_1.reset(new merkle_tree_gadget<FieldT>(pb, *input_cm_1, *anchor, enforce_input_1));
        merkle_lookup_2.reset(new merkle_tree_gadget<FieldT>(pb, *input_cm_2, *anchor, enforce_input_2));

        output_value_1.allocate(pb, 64, "");
        output_value_2.allocate(pb, 64, "");

        output_rho_1.reset(new digest_variable<FieldT>(pb, 256, ""));
        output_rho_2.reset(new digest_variable<FieldT>(pb, 256, ""));
        output_pk_1.reset(new digest_variable<FieldT>(pb, 256, ""));
        output_pk_2.reset(new digest_variable<FieldT>(pb, 256, ""));

        output_cm_hasher_1.reset(new NoteCommitment<FieldT>(pb, ZERO, output_rho_1->bits, output_pk_1->bits, output_value_1, output_cm_1));
        output_cm_hasher_2.reset(new NoteCommitment<FieldT>(pb, ZERO, output_rho_2->bits, output_pk_2->bits, output_value_2, output_cm_2));

        output_nf_hasher_1.reset(new SendNullifier<FieldT>(pb, ZERO, output_rho_1->bits, send_nullifier_output_1));
        output_nf_hasher_2.reset(new SendNullifier<FieldT>(pb, ZERO, output_rho_2->bits, send_nullifier_output_2));
    }

    void generate_r1cs_constraints() {
        unpacker->generate_r1cs_constraints(true);
        generate_r1cs_equals_const_constraint<FieldT>(this->pb, ZERO, FieldT::zero(), "ZERO");

        input_sk_1->generate_r1cs_constraints();
        input_sk_2->generate_r1cs_constraints();
        input_rho_1->generate_r1cs_constraints();
        input_rho_2->generate_r1cs_constraints();
        key_hasher_1->generate_r1cs_constraints();
        key_hasher_2->generate_r1cs_constraints();
        input_nf_hasher_1->generate_r1cs_constraints();
        input_nf_hasher_2->generate_r1cs_constraints();

        for (size_t i = 0; i < 64; i++) {
            generate_boolean_r1cs_constraint<FieldT>(
                this->pb,
                input_value_1[i],
                ""
            );
            generate_boolean_r1cs_constraint<FieldT>(
                this->pb,
                input_value_2[i],
                ""
            );
        }

        input_cm_hasher_1->generate_r1cs_constraints();
        input_cm_hasher_2->generate_r1cs_constraints();

        generate_boolean_r1cs_constraint<FieldT>(this->pb, enforce_input_1, "");

        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(
                    packed_addition(input_value_1),
                    (1 - enforce_input_1),
                    0
        ), "");

        generate_boolean_r1cs_constraint<FieldT>(this->pb, enforce_input_2, "");

        this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(
                    packed_addition(input_value_2),
                    (1 - enforce_input_2),
                    0
        ), "");

        merkle_lookup_1->generate_r1cs_constraints();
        merkle_lookup_2->generate_r1cs_constraints();

        for (size_t i = 0; i < 64; i++) {
            generate_boolean_r1cs_constraint<FieldT>(
                this->pb,
                output_value_1[i],
                ""
            );
            generate_boolean_r1cs_constraint<FieldT>(
                this->pb,
                output_value_2[i],
                ""
            );
        }

        output_rho_1->generate_r1cs_constraints();
        output_rho_2->generate_r1cs_constraints();
        output_pk_1->generate_r1cs_constraints();
        output_pk_2->generate_r1cs_constraints();

        output_cm_hasher_1->generate_r1cs_constraints();
        output_cm_hasher_2->generate_r1cs_constraints();

        output_nf_hasher_1->generate_r1cs_constraints();
        output_nf_hasher_2->generate_r1cs_constraints();

        {
            linear_combination<FieldT> left_side =
                packed_addition(input_value_1) + packed_addition(input_value_2);

            linear_combination<FieldT> right_side =
                packed_addition(output_value_1) + packed_addition(output_value_2);

            // Ensure that both sides are equal
            this->pb.add_r1cs_constraint(r1cs_constraint<FieldT>(
                left_side,
                1,
                right_side
            ));
        }
    }

    void generate_r1cs_witness(
        const std::vector<unsigned char>& witness_rho_1,
        const std::vector<unsigned char>& witness_sk_1,
        uint64_t witness_value_1,
        size_t path_index_1,
        const std::vector<std::vector<bool>>& authentication_path_1,
        const std::vector<unsigned char>& witness_rho_2,
        const std::vector<unsigned char>& witness_sk_2,
        uint64_t witness_value_2,
        size_t path_index_2,
        const std::vector<std::vector<bool>>& authentication_path_2,
        const std::vector<unsigned char>& output_witness_rho_1,
        const std::vector<unsigned char>& output_witness_pk_1,
        uint64_t output_witness_value_1,
        const std::vector<unsigned char>& output_witness_rho_2,
        const std::vector<unsigned char>& output_witness_pk_2,
        uint64_t output_witness_value_2
    ) {
        this->pb.val(ZERO) = FieldT::zero();

        this->pb.val(enforce_input_1) = (witness_value_1 != 0) ? FieldT::one() : FieldT::zero();
        this->pb.val(enforce_input_2) = (witness_value_2 != 0) ? FieldT::one() : FieldT::zero();

        input_rho_1->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_rho_1)
        );

        input_sk_1->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_sk_1)
        );

        input_value_1.fill_with_bits(
            this->pb,
            uint64_to_bool_vector(witness_value_1)
        );

        key_hasher_1->generate_r1cs_witness();
        input_cm_hasher_1->generate_r1cs_witness();
        input_nf_hasher_1->generate_r1cs_witness();
        merkle_lookup_1->generate_r1cs_witness(path_index_1, authentication_path_1);

        input_rho_2->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_rho_2)
        );

        input_sk_2->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(witness_sk_2)
        );

        input_value_2.fill_with_bits(
            this->pb,
            uint64_to_bool_vector(witness_value_2)
        );

        key_hasher_2->generate_r1cs_witness();
        input_cm_hasher_2->generate_r1cs_witness();
        input_nf_hasher_2->generate_r1cs_witness();
        merkle_lookup_2->generate_r1cs_witness(path_index_2, authentication_path_2);

        output_rho_1->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(output_witness_rho_1)
        );

        output_pk_1->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(output_witness_pk_1)
        );

        output_value_1.fill_with_bits(
            this->pb,
            uint64_to_bool_vector(output_witness_value_1)
        );

        output_cm_hasher_1->generate_r1cs_witness();
        output_nf_hasher_1->generate_r1cs_witness();

        output_rho_2->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(output_witness_rho_2)
        );

        output_pk_2->bits.fill_with_bits(
            this->pb,
            convertBytesVectorToVector(output_witness_pk_2)
        );

        output_value_2.fill_with_bits(
            this->pb,
            uint64_to_bool_vector(output_witness_value_2)
        );

        output_cm_hasher_2->generate_r1cs_witness();
        output_nf_hasher_2->generate_r1cs_witness();

        unpacker->generate_r1cs_witness_from_bits();
    }

    static r1cs_primary_input<FieldT> witness_map(
        const std::vector<unsigned char> &witness_anchor,
        const std::vector<unsigned char> &input_nf_1,
        const std::vector<unsigned char> &input_nf_2,
        const std::vector<unsigned char> &output_nf_1,
        const std::vector<unsigned char> &output_nf_2,
        const std::vector<unsigned char> &output_cm_1,
        const std::vector<unsigned char> &output_cm_2
    )
    {
        std::vector<bool> verify_inputs;

        std::vector<bool> anchor_bits = convertBytesVectorToVector(witness_anchor);
        std::vector<bool> input_nf1_bits = convertBytesVectorToVector(input_nf_1);
        std::vector<bool> input_nf2_bits = convertBytesVectorToVector(input_nf_2);
        std::vector<bool> output_nf1_bits = convertBytesVectorToVector(output_nf_1);
        std::vector<bool> output_nf2_bits = convertBytesVectorToVector(output_nf_2);
        std::vector<bool> output_cm1_bits = convertBytesVectorToVector(output_cm_1);
        std::vector<bool> output_cm2_bits = convertBytesVectorToVector(output_cm_2);

        verify_inputs.insert(verify_inputs.end(), anchor_bits.begin(), anchor_bits.end());
        verify_inputs.insert(verify_inputs.end(), input_nf1_bits.begin(), input_nf1_bits.end());
        verify_inputs.insert(verify_inputs.end(), input_nf2_bits.begin(), input_nf2_bits.end());
        verify_inputs.insert(verify_inputs.end(), output_nf1_bits.begin(), output_nf1_bits.end());
        verify_inputs.insert(verify_inputs.end(), output_nf2_bits.begin(), output_nf2_bits.end());
        verify_inputs.insert(verify_inputs.end(), output_cm1_bits.begin(), output_cm1_bits.end());
        verify_inputs.insert(verify_inputs.end(), output_cm2_bits.begin(), output_cm2_bits.end());

        assert(verify_inputs.size() == verifying_input_bit_size());
        auto verify_field_elements = pack_bit_vector_into_field_element_vector<FieldT>(verify_inputs);
        assert(verify_field_elements.size() == verifying_field_element_size());
        return verify_field_elements;
    }

    static size_t verifying_field_element_size() {
        return div_ceil(verifying_input_bit_size(), FieldT::capacity());
    }

    static size_t verifying_input_bit_size() {
        size_t acc = 0;

        acc += 256; // the anchor
        acc += 256; // input 1 nullifier
        acc += 256; // input 2 nullifier
        acc += 256; // output 1 nullifier
        acc += 256; // output 2 nullifier
        acc += 256; // output 1 commitment
        acc += 256; // output 2 commitment

        return acc;
    }
};

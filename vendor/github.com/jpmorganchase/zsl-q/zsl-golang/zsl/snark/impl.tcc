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

#include "zsl.h"

#include <iostream>
#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp"
#include "common/default_types/r1cs_ppzksnark_pp.hpp"
#include "common/utils.hpp"
#include "common/profiling.hpp"

using namespace libsnark;
using namespace std;

#include "gadgets.tcc"

typedef Fr<default_r1cs_ppzksnark_pp> FieldT;

#include <fstream>

template<typename T>
void saveToFile(std::string path, T& obj) {
    std::stringstream ss;
    ss << obj;
    std::ofstream fh;
    fh.open(path, std::ios::binary);
    ss.rdbuf()->pubseekpos(0, std::ios_base::out);
    fh << ss.rdbuf();
    fh.flush();
    fh.close();
}

template<typename T>
void loadFromFile(std::string path, T& objIn) {
    std::stringstream ss;
    std::ifstream fh(path, std::ios::binary);

    ss << fh.rdbuf();
    fh.close();

    ss.rdbuf()->pubseekpos(0, std::ios_base::in);

    ss >> objIn;
}

void zsl_initialize()
{
    default_r1cs_ppzksnark_pp::init_public_params();
    inhibit_profiling_info = true;
    inhibit_profiling_counters = true;
}

bool zsl_verify_shielding(
    void *proof_ptr,
    void *send_nf_ptr,
    void *cm_ptr,
    uint64_t value
)
{
    unsigned char *send_nf = reinterpret_cast<unsigned char *>(send_nf_ptr);
    unsigned char *cm = reinterpret_cast<unsigned char *>(cm_ptr);
    unsigned char *proof = reinterpret_cast<unsigned char *>(proof_ptr);

    std::vector<unsigned char> proof_v(proof, proof+584);

    std::stringstream proof_data;
    for (int i = 0; i < 584; i++) {
        proof_data << proof_v[i];
    }

    assert(proof_data.str().size() == 584);

    proof_data.rdbuf()->pubseekpos(0, std::ios_base::in);

    r1cs_ppzksnark_proof<default_r1cs_ppzksnark_pp> proof_obj;
    proof_data >> proof_obj;

    auto witness_map = ShieldingCircuit<FieldT>::witness_map(
        std::vector<unsigned char>(send_nf, send_nf+32),
        std::vector<unsigned char>(cm, cm+32),
        value
    );

    r1cs_ppzksnark_verification_key<default_r1cs_ppzksnark_pp> verification_key;
    loadFromFile("shielding.vk", verification_key);

    if (!r1cs_ppzksnark_verifier_strong_IC<default_r1cs_ppzksnark_pp>(verification_key, witness_map, proof_obj)) {
        return false;
    } else {
        return true;
    }
}

bool zsl_verify_unshielding(
    void *proof_ptr,
    void *spend_nf_ptr,
    void *rt_ptr,
    uint64_t value
)
{
    unsigned char *spend_nf = reinterpret_cast<unsigned char *>(spend_nf_ptr);
    unsigned char *rt = reinterpret_cast<unsigned char *>(rt_ptr);
    unsigned char *proof = reinterpret_cast<unsigned char *>(proof_ptr);

    std::vector<unsigned char> proof_v(proof, proof+584);

    std::stringstream proof_data;
    for (int i = 0; i < 584; i++) {
        proof_data << proof_v[i];
    }

    assert(proof_data.str().size() == 584);

    proof_data.rdbuf()->pubseekpos(0, std::ios_base::in);

    r1cs_ppzksnark_proof<default_r1cs_ppzksnark_pp> proof_obj;
    proof_data >> proof_obj;

    auto witness_map = UnshieldingCircuit<FieldT>::witness_map(
        std::vector<unsigned char>(spend_nf, spend_nf+32),
        std::vector<unsigned char>(rt, rt+32),
        value
    );

    r1cs_ppzksnark_verification_key<default_r1cs_ppzksnark_pp> verification_key;
    loadFromFile("unshielding.vk", verification_key);

    if (!r1cs_ppzksnark_verifier_strong_IC<default_r1cs_ppzksnark_pp>(verification_key, witness_map, proof_obj)) {
        return false;
    } else {
        return true;
    }
}

void zsl_prove_unshielding(
    void *rho_ptr,
    void *pk_ptr,
    uint64_t value,
    uint64_t tree_position,
    void *authentication_path_ptr,
    void *output_proof_ptr
)
{
    unsigned char *rho = reinterpret_cast<unsigned char *>(rho_ptr);
    unsigned char *pk = reinterpret_cast<unsigned char *>(pk_ptr);
    unsigned char *output_proof = reinterpret_cast<unsigned char *>(output_proof_ptr);
    unsigned char *authentication_path = reinterpret_cast<unsigned char *>(authentication_path_ptr);

    protoboard<FieldT> pb;
    UnshieldingCircuit<FieldT> g(pb);
    g.generate_r1cs_constraints();

    std::vector<std::vector<bool>> auth_path;
    for (int i = 0; i < 29; i++) {
        auth_path.push_back(convertBytesVectorToVector(std::vector<unsigned char>(authentication_path + i*32, authentication_path + i*32 + 32)));
    }

    std::reverse(std::begin(auth_path), std::end(auth_path));

    g.generate_r1cs_witness(
        std::vector<unsigned char>(rho, rho + 32),
        std::vector<unsigned char>(pk, pk + 32),
        value,
        tree_position,
        auth_path
    );
    pb.constraint_system.swap_AB_if_beneficial();
    assert(pb.is_satisfied());

    r1cs_ppzksnark_proving_key<default_r1cs_ppzksnark_pp> proving_key;
    loadFromFile("unshielding.pk", proving_key);

    auto proof = r1cs_ppzksnark_prover<default_r1cs_ppzksnark_pp>(proving_key, pb.primary_input(), pb.auxiliary_input(), pb.constraint_system);

    std::stringstream proof_data;
    proof_data << proof;
    auto proof_str = proof_data.str();
    assert(proof_str.size() == 584);

    for (int i = 0; i < 584; i++) {
        output_proof[i] = proof_str[i];
    }
}

void zsl_prove_shielding(
    void *rho_ptr,
    void *pk_ptr,
    uint64_t value,
    void *output_proof_ptr
)
{
    unsigned char *rho = reinterpret_cast<unsigned char *>(rho_ptr);
    unsigned char *pk = reinterpret_cast<unsigned char *>(pk_ptr);
    unsigned char *output_proof = reinterpret_cast<unsigned char *>(output_proof_ptr);

    protoboard<FieldT> pb;
    ShieldingCircuit<FieldT> g(pb);
    g.generate_r1cs_constraints();
    g.generate_r1cs_witness(
        // rho
        std::vector<unsigned char>(rho, rho + 32),
        // pk
        std::vector<unsigned char>(pk, pk + 32),
        // value
        value
    );
    pb.constraint_system.swap_AB_if_beneficial();
    assert(pb.is_satisfied());

    r1cs_ppzksnark_proving_key<default_r1cs_ppzksnark_pp> proving_key;
    loadFromFile("shielding.pk", proving_key);

    auto proof = r1cs_ppzksnark_prover<default_r1cs_ppzksnark_pp>(proving_key, pb.primary_input(), pb.auxiliary_input(), pb.constraint_system);

    std::stringstream proof_data;
    proof_data << proof;
    auto proof_str = proof_data.str();
    assert(proof_str.size() == 584);

    for (int i = 0; i < 584; i++) {
        output_proof[i] = proof_str[i];
    }
}

bool zsl_verify_transfer(
    void *proof_ptr,
    void *anchor_ptr,
    void *spend_nf_ptr_1,
    void *spend_nf_ptr_2,
    void *send_nf_ptr_1,
    void *send_nf_ptr_2,
    void *cm_ptr_1,
    void *cm_ptr_2
)
{
    unsigned char *anchor = reinterpret_cast<unsigned char *>(anchor_ptr);
    unsigned char *spend_nf_1 = reinterpret_cast<unsigned char *>(spend_nf_ptr_1);
    unsigned char *spend_nf_2 = reinterpret_cast<unsigned char *>(spend_nf_ptr_2);
    unsigned char *send_nf_1 = reinterpret_cast<unsigned char *>(send_nf_ptr_1);
    unsigned char *send_nf_2 = reinterpret_cast<unsigned char *>(send_nf_ptr_2);
    unsigned char *cm_1 = reinterpret_cast<unsigned char *>(cm_ptr_1);
    unsigned char *cm_2 = reinterpret_cast<unsigned char *>(cm_ptr_2);
    unsigned char *proof = reinterpret_cast<unsigned char *>(proof_ptr);

    std::vector<unsigned char> proof_v(proof, proof+584);

    std::stringstream proof_data;
    for (int i = 0; i < 584; i++) {
        proof_data << proof_v[i];
    }

    assert(proof_data.str().size() == 584);

    proof_data.rdbuf()->pubseekpos(0, std::ios_base::in);

    r1cs_ppzksnark_proof<default_r1cs_ppzksnark_pp> proof_obj;
    proof_data >> proof_obj;

    auto witness_map = TransferCircuit<FieldT>::witness_map(
        std::vector<unsigned char>(anchor, anchor+32),
        std::vector<unsigned char>(spend_nf_1, spend_nf_1+32),
        std::vector<unsigned char>(spend_nf_2, spend_nf_2+32),
        std::vector<unsigned char>(send_nf_1, send_nf_1+32),
        std::vector<unsigned char>(send_nf_2, send_nf_2+32),
        std::vector<unsigned char>(cm_1, cm_1+32),
        std::vector<unsigned char>(cm_2, cm_2+32)
    );

    r1cs_ppzksnark_verification_key<default_r1cs_ppzksnark_pp> verification_key;
    loadFromFile("transfer.vk", verification_key);

    if (!r1cs_ppzksnark_verifier_strong_IC<default_r1cs_ppzksnark_pp>(verification_key, witness_map, proof_obj)) {
        return false;
    } else {
        return true;
    }
}

void zsl_prove_transfer(
    void *output_proof_ptr,
    void *input_rho_ptr_1,
    void *input_pk_ptr_1,
    uint64_t input_value_1,
    uint64_t input_tree_position_1,
    void *input_authentication_path_ptr_1,
    void *input_rho_ptr_2,
    void *input_pk_ptr_2,
    uint64_t input_value_2,
    uint64_t input_tree_position_2,
    void *input_authentication_path_ptr_2,
    void *output_rho_ptr_1,
    void *output_pk_ptr_1,
    uint64_t output_value_1,
    void *output_rho_ptr_2,
    void *output_pk_ptr_2,
    uint64_t output_value_2
)
{
    unsigned char *output_proof = reinterpret_cast<unsigned char *>(output_proof_ptr);

    unsigned char *input_rho_1 = reinterpret_cast<unsigned char *>(input_rho_ptr_1);
    unsigned char *input_pk_1 = reinterpret_cast<unsigned char *>(input_pk_ptr_1);
    unsigned char *authentication_path_1 = reinterpret_cast<unsigned char *>(input_authentication_path_ptr_1);

    unsigned char *input_rho_2 = reinterpret_cast<unsigned char *>(input_rho_ptr_2);
    unsigned char *input_pk_2 = reinterpret_cast<unsigned char *>(input_pk_ptr_2);
    unsigned char *authentication_path_2 = reinterpret_cast<unsigned char *>(input_authentication_path_ptr_2);

    unsigned char *output_rho_1 = reinterpret_cast<unsigned char *>(output_rho_ptr_1);
    unsigned char *output_pk_1 = reinterpret_cast<unsigned char *>(output_pk_ptr_1);
    unsigned char *output_rho_2 = reinterpret_cast<unsigned char *>(output_rho_ptr_2);
    unsigned char *output_pk_2 = reinterpret_cast<unsigned char *>(output_pk_ptr_2);

    std::vector<std::vector<bool>> auth_path_1;
    for (int i = 0; i < 29; i++) {
        auth_path_1.push_back(convertBytesVectorToVector(std::vector<unsigned char>(authentication_path_1 + i*32, authentication_path_1 + i*32 + 32)));
    }

    std::reverse(std::begin(auth_path_1), std::end(auth_path_1));

    std::vector<std::vector<bool>> auth_path_2;
    for (int i = 0; i < 29; i++) {
        auth_path_2.push_back(convertBytesVectorToVector(std::vector<unsigned char>(authentication_path_2 + i*32, authentication_path_2 + i*32 + 32)));
    }

    std::reverse(std::begin(auth_path_2), std::end(auth_path_2));

    protoboard<FieldT> pb;
    TransferCircuit<FieldT> g(pb);
    g.generate_r1cs_constraints();
    g.generate_r1cs_witness(
        std::vector<unsigned char>(input_rho_1, input_rho_1 + 32),
        std::vector<unsigned char>(input_pk_1, input_pk_1 + 32),
        input_value_1,
        input_tree_position_1,
        auth_path_1,
        std::vector<unsigned char>(input_rho_2, input_rho_2 + 32),
        std::vector<unsigned char>(input_pk_2, input_pk_2 + 32),
        input_value_2,
        input_tree_position_2,
        auth_path_2,
        std::vector<unsigned char>(output_rho_1, output_rho_1 + 32),
        std::vector<unsigned char>(output_pk_1, output_pk_1 + 32),
        output_value_1,
        std::vector<unsigned char>(output_rho_2, output_rho_2 + 32),
        std::vector<unsigned char>(output_pk_2, output_pk_2 + 32),
        output_value_2
    );
    pb.constraint_system.swap_AB_if_beneficial();
    assert(pb.is_satisfied());

    r1cs_ppzksnark_proving_key<default_r1cs_ppzksnark_pp> proving_key;
    loadFromFile("transfer.pk", proving_key);

    auto proof = r1cs_ppzksnark_prover<default_r1cs_ppzksnark_pp>(proving_key, pb.primary_input(), pb.auxiliary_input(), pb.constraint_system);

    std::stringstream proof_data;
    proof_data << proof;
    auto proof_str = proof_data.str();
    assert(proof_str.size() == 584);

    for (int i = 0; i < 584; i++) {
        output_proof[i] = proof_str[i];
    }
}

void zsl_paramgen_transfer()
{
    protoboard<FieldT> pb;
    TransferCircuit<FieldT> g(pb);
    g.generate_r1cs_constraints();

    const r1cs_constraint_system<FieldT> constraint_system = pb.get_constraint_system();
    cout << "Number of R1CS constraints: " << constraint_system.num_constraints() << endl;
    auto crs = r1cs_ppzksnark_generator<default_r1cs_ppzksnark_pp>(constraint_system);

    saveToFile("transfer.pk", crs.pk);
    saveToFile("transfer.vk", crs.vk);
}

void zsl_paramgen_shielding()
{
    protoboard<FieldT> pb;
    ShieldingCircuit<FieldT> g(pb);
    g.generate_r1cs_constraints();

    const r1cs_constraint_system<FieldT> constraint_system = pb.get_constraint_system();
    cout << "Number of R1CS constraints: " << constraint_system.num_constraints() << endl;
    auto crs = r1cs_ppzksnark_generator<default_r1cs_ppzksnark_pp>(constraint_system);

    saveToFile("shielding.pk", crs.pk);
    saveToFile("shielding.vk", crs.vk);
}

void zsl_paramgen_unshielding()
{
    protoboard<FieldT> pb;
    UnshieldingCircuit<FieldT> g(pb);
    g.generate_r1cs_constraints();

    const r1cs_constraint_system<FieldT> constraint_system = pb.get_constraint_system();
    cout << "Number of R1CS constraints: " << constraint_system.num_constraints() << endl;
    auto crs = r1cs_ppzksnark_generator<default_r1cs_ppzksnark_pp>(constraint_system);

    saveToFile("unshielding.pk", crs.pk);
    saveToFile("unshielding.vk", crs.vk);
}

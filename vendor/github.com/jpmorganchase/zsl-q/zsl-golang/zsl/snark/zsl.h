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

#ifndef _ZSL_H_
#define _ZSL_H_

#include <stdlib.h>
#include <inttypes.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

    void zsl_initialize();
    bool zsl_verify_shielding(
        void *proof,
        void *send_nf,
        void *cm,
        uint64_t value
    );
    void zsl_prove_shielding(
        void *rho,
        void *pk,
        uint64_t value,
        void *output_proof
    );
    void zsl_paramgen_shielding();

    void zsl_prove_unshielding(
        void *rho,
        void *sk,
        void *addr,
        uint64_t value,
        uint64_t tree_position,
        void *authentication_path,
        void *output_proof
    );
    bool zsl_verify_unshielding(
        void *proof_ptr,
        void *spend_nf_ptr,
        void *addr_ptr,
        void *rt_ptr,
        uint64_t value
    );

    void zsl_paramgen_unshielding();

    void zsl_paramgen_transfer();

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
    );

    bool zsl_verify_transfer(
        void *proof_ptr,
        void *anchor_ptr,
        void *spend_nf_ptr_1,
        void *spend_nf_ptr_2,
        void *send_nf_ptr_1,
        void *send_nf_ptr_2,
        void *cm_ptr_1,
        void *cm_ptr_2
    );


#ifdef __cplusplus
}
#endif

#endif

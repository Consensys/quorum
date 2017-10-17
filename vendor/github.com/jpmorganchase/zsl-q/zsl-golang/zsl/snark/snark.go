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

package snark

// #cgo LDFLAGS: -L${SRCDIR} -lzsl -lstdc++ -lgmp -lgomp
// #include <zsl.h>
import "C"
import (
	"sync"
	"unsafe"
)

// Init() is only ever called once
var onceInit sync.Once

func Init() {
	onceInit.Do(func() {
		C.zsl_initialize()
	})
}

func ProveTransfer(input_rho_1 [32]byte,
	input_sk_1 [32]byte,
	input_value_1 uint64,
	input_tree_position_1 uint64,
	input_authentication_path_1 [29][32]byte,
	input_rho_2 [32]byte,
	input_sk_2 [32]byte,
	input_value_2 uint64,
	input_tree_position_2 uint64,
	input_authentication_path_2 [29][32]byte,
	output_rho_1 [32]byte,
	output_pk_1 [32]byte,
	output_value_1 uint64,
	output_rho_2 [32]byte,
	output_pk_2 [32]byte,
	output_value_2 uint64) [584]byte {
	var proof_buf [584]byte

	C.zsl_prove_transfer(unsafe.Pointer(&proof_buf[0]),
		unsafe.Pointer(&input_rho_1[0]),
		unsafe.Pointer(&input_sk_1[0]),
		C.uint64_t(input_value_1),
		C.uint64_t(input_tree_position_1),
		unsafe.Pointer(&input_authentication_path_1[0][0]),
		unsafe.Pointer(&input_rho_2[0]),
		unsafe.Pointer(&input_sk_2[0]),
		C.uint64_t(input_value_2),
		C.uint64_t(input_tree_position_2),
		unsafe.Pointer(&input_authentication_path_2[0][0]),
		unsafe.Pointer(&output_rho_1[0]),
		unsafe.Pointer(&output_pk_1[0]),
		C.uint64_t(output_value_1),
		unsafe.Pointer(&output_rho_2[0]),
		unsafe.Pointer(&output_pk_2[0]),
		C.uint64_t(output_value_2))

	return proof_buf
}

func VerifyTransfer(proof [584]byte,
	anchor [32]byte,
	spend_nf_1 [32]byte,
	spend_nf_2 [32]byte,
	send_nf_1 [32]byte,
	send_nf_2 [32]byte,
	cm_1 [32]byte,
	cm_2 [32]byte) bool {
	ret := C.zsl_verify_transfer(unsafe.Pointer(&proof[0]),
		unsafe.Pointer(&anchor[0]),
		unsafe.Pointer(&spend_nf_1[0]),
		unsafe.Pointer(&spend_nf_2[0]),
		unsafe.Pointer(&send_nf_1[0]),
		unsafe.Pointer(&send_nf_2[0]),
		unsafe.Pointer(&cm_1[0]),
		unsafe.Pointer(&cm_2[0]))

	if ret {
		return true
	} else {
		return false
	}
}

func ProveShielding(rho [32]byte, pk [32]byte, value uint64) [584]byte {
	var proof_buf [584]byte

	rho_ptr := C.CBytes(rho[:])
	pk_ptr := C.CBytes(pk[:])

	C.zsl_prove_shielding(rho_ptr, pk_ptr, C.uint64_t(value), unsafe.Pointer(&proof_buf[0]))

	C.free(rho_ptr)
	C.free(pk_ptr)

	return proof_buf
}

func VerifyShielding(proof [584]byte, send_nf [32]byte, cm [32]byte, value uint64) bool {
	send_nf_ptr := C.CBytes(send_nf[:])
	cm_ptr := C.CBytes(cm[:])
	ret := C.zsl_verify_shielding(unsafe.Pointer(&proof[0]), send_nf_ptr, cm_ptr, C.uint64_t(value))

	C.free(send_nf_ptr)
	C.free(cm_ptr)

	if ret {
		return true
	} else {
		return false
	}
}

func ProveUnshielding(rho [32]byte,
	sk [32]byte,
	value uint64,
	tree_position uint64,
	authentication_path [29][32]byte) [584]byte {
	var proof_buf [584]byte

	C.zsl_prove_unshielding(unsafe.Pointer(&rho[0]),
		unsafe.Pointer(&sk[0]),
		C.uint64_t(value),
		C.uint64_t(tree_position),
		unsafe.Pointer(&authentication_path[0][0]),
		unsafe.Pointer(&proof_buf[0]))

	return proof_buf
}

func VerifyUnshielding(proof [584]byte, spend_nf [32]byte, rt [32]byte, value uint64) bool {
	ret := C.zsl_verify_unshielding(unsafe.Pointer(&proof[0]),
		unsafe.Pointer(&spend_nf[0]),
		unsafe.Pointer(&rt[0]),
		C.uint64_t(value))

	if ret {
		return true
	} else {
		return false
	}
}

func CreateParamsShielding() {
	C.zsl_initialize()
	C.zsl_paramgen_shielding()
}

func CreateParamsUnshielding() {
	C.zsl_initialize()
	C.zsl_paramgen_unshielding()
}

func CreateParamsTransfer() {
	C.zsl_initialize()
	C.zsl_paramgen_transfer()
}

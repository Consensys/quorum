// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bn256"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/ripemd160"

	// ZSL START
	sha256compress "github.com/jpmorganchase/zsl-q/zsl-golang/zsl/sha256"
	"github.com/jpmorganchase/zsl-q/zsl-golang/zsl/snark"
	// ZSL END
)

// ZSL START
const ZSL_PROOF_SIZE uint64 = 584

// ZSL END

var errBadPrecompileInput = errors.New("bad pre compile input")

// Precompiled contract is the basic interface for native Go contracts. The implementation
// requires a deterministic gas count based on the input size of the Run method of the
// contract.
type PrecompiledContract interface {
	RequiredGas(input []byte) uint64  // RequiredPrice calculates the contract gas use
	Run(input []byte) ([]byte, error) // Run runs the precompiled contract
}

// PrecompiledContractsHomestead contains the default set of pre-compiled Ethereum
// contracts used in the Frontier and Homestead releases.
var PrecompiledContractsHomestead = map[common.Address]PrecompiledContract{
	common.BytesToAddress([]byte{1}): &ecrecover{},
	common.BytesToAddress([]byte{2}): &sha256hash{},
	common.BytesToAddress([]byte{3}): &ripemd160hash{},
	common.BytesToAddress([]byte{4}): &dataCopy{},
	// ZSL START
	common.BytesToAddress([]byte{0x88, 0x01}): &sha256Compress{},
	common.BytesToAddress([]byte{0x88, 0x02}): &verifyShieldedTransfer{},
	common.BytesToAddress([]byte{0x88, 0x03}): &verifyShielding{},
	common.BytesToAddress([]byte{0x88, 0x04}): &verifyUnshielding{},
	// ZSL END
}

// PrecompiledContractsMetropolis contains the default set of pre-compiled Ethereum
// contracts used in the Metropolis release.
var PrecompiledContractsMetropolis = map[common.Address]PrecompiledContract{
	common.BytesToAddress([]byte{1}): &ecrecover{},
	common.BytesToAddress([]byte{2}): &sha256hash{},
	common.BytesToAddress([]byte{3}): &ripemd160hash{},
	common.BytesToAddress([]byte{4}): &dataCopy{},
	common.BytesToAddress([]byte{5}): &bigModExp{},
	common.BytesToAddress([]byte{6}): &bn256Add{},
	common.BytesToAddress([]byte{7}): &bn256ScalarMul{},
	common.BytesToAddress([]byte{8}): &bn256Pairing{},
	// ZSL START
	common.BytesToAddress([]byte{0x88, 0x01}): &sha256Compress{},
	common.BytesToAddress([]byte{0x88, 0x02}): &verifyShieldedTransfer{},
	common.BytesToAddress([]byte{0x88, 0x03}): &verifyShielding{},
	common.BytesToAddress([]byte{0x88, 0x04}): &verifyUnshielding{},
	// ZSL END
}

// RunPrecompiledContract runs and evaluates the output of a precompiled contract.
func RunPrecompiledContract(p PrecompiledContract, input []byte, contract *Contract) (ret []byte, err error) {
	gas := p.RequiredGas(input)
	if contract.UseGas(gas) {
		return p.Run(input)
	}
	return nil, ErrOutOfGas
}

// ECRECOVER implemented as a native contract.
type ecrecover struct{}

func (c *ecrecover) RequiredGas(input []byte) uint64 {
	return params.EcrecoverGas
}

func (c *ecrecover) Run(input []byte) ([]byte, error) {
	const ecRecoverInputLength = 128

	input = common.RightPadBytes(input, ecRecoverInputLength)
	// "input" is (hash, v, r, s), each 32 bytes
	// but for ecrecover we want (r, s, v)

	r := new(big.Int).SetBytes(input[64:96])
	s := new(big.Int).SetBytes(input[96:128])
	v := input[63] - 27

	// tighter sig s values input homestead only apply to tx sigs
	if !allZero(input[32:63]) || !crypto.ValidateSignatureValues(v, r, s, false) {
		return nil, nil
	}
	// v needs to be at the end for libsecp256k1
	pubKey, err := crypto.Ecrecover(input[:32], append(input[64:128], v))
	// make sure the public key is a valid one
	if err != nil {
		return nil, nil
	}

	// the first byte of pubkey is bitcoin heritage
	return common.LeftPadBytes(crypto.Keccak256(pubKey[1:])[12:], 32), nil
}

// SHA256 implemented as a native contract.
type sha256hash struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *sha256hash) RequiredGas(input []byte) uint64 {
	return uint64(len(input)+31)/32*params.Sha256PerWordGas + params.Sha256BaseGas
}
func (c *sha256hash) Run(input []byte) ([]byte, error) {
	h := sha256.Sum256(input)
	return h[:], nil
}

// RIPMED160 implemented as a native contract.
type ripemd160hash struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *ripemd160hash) RequiredGas(input []byte) uint64 {
	return uint64(len(input)+31)/32*params.Ripemd160PerWordGas + params.Ripemd160BaseGas
}
func (c *ripemd160hash) Run(input []byte) ([]byte, error) {
	ripemd := ripemd160.New()
	ripemd.Write(input)
	return common.LeftPadBytes(ripemd.Sum(nil), 32), nil
}

// data copy implemented as a native contract.
type dataCopy struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *dataCopy) RequiredGas(input []byte) uint64 {
	return uint64(len(input)+31)/32*params.IdentityPerWordGas + params.IdentityBaseGas
}
func (c *dataCopy) Run(in []byte) ([]byte, error) {
	return in, nil
}

// bigModExp implements a native big integer exponential modular operation.
type bigModExp struct{}

var (
	big1      = big.NewInt(1)
	big4      = big.NewInt(4)
	big8      = big.NewInt(8)
	big16     = big.NewInt(16)
	big32     = big.NewInt(32)
	big64     = big.NewInt(64)
	big96     = big.NewInt(96)
	big480    = big.NewInt(480)
	big1024   = big.NewInt(1024)
	big3072   = big.NewInt(3072)
	big199680 = big.NewInt(199680)
)

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bigModExp) RequiredGas(input []byte) uint64 {
	var (
		baseLen = new(big.Int).SetBytes(getData(input, 0, 32))
		expLen  = new(big.Int).SetBytes(getData(input, 32, 32))
		modLen  = new(big.Int).SetBytes(getData(input, 64, 32))
	)
	if len(input) > 96 {
		input = input[96:]
	} else {
		input = input[:0]
	}
	// Retrieve the head 32 bytes of exp for the adjusted exponent length
	var expHead *big.Int
	if big.NewInt(int64(len(input))).Cmp(baseLen) <= 0 {
		expHead = new(big.Int)
	} else {
		if expLen.Cmp(big32) > 0 {
			expHead = new(big.Int).SetBytes(getData(input, baseLen.Uint64(), 32))
		} else {
			expHead = new(big.Int).SetBytes(getData(input, baseLen.Uint64(), expLen.Uint64()))
		}
	}
	// Calculate the adjusted exponent length
	var msb int
	if bitlen := expHead.BitLen(); bitlen > 0 {
		msb = bitlen - 1
	}
	adjExpLen := new(big.Int)
	if expLen.Cmp(big32) > 0 {
		adjExpLen.Sub(expLen, big32)
		adjExpLen.Mul(big8, adjExpLen)
	}
	adjExpLen.Add(adjExpLen, big.NewInt(int64(msb)))

	// Calculate the gas cost of the operation
	gas := new(big.Int).Set(math.BigMax(modLen, baseLen))
	switch {
	case gas.Cmp(big64) <= 0:
		gas.Mul(gas, gas)
	case gas.Cmp(big1024) <= 0:
		gas = new(big.Int).Add(
			new(big.Int).Div(new(big.Int).Mul(gas, gas), big4),
			new(big.Int).Sub(new(big.Int).Mul(big96, gas), big3072),
		)
	default:
		gas = new(big.Int).Add(
			new(big.Int).Div(new(big.Int).Mul(gas, gas), big16),
			new(big.Int).Sub(new(big.Int).Mul(big480, gas), big199680),
		)
	}
	gas.Mul(gas, math.BigMax(adjExpLen, big1))
	gas.Div(gas, new(big.Int).SetUint64(params.ModExpQuadCoeffDiv))

	if gas.BitLen() > 64 {
		return math.MaxUint64
	}
	return gas.Uint64()
}

func (c *bigModExp) Run(input []byte) ([]byte, error) {
	var (
		baseLen = new(big.Int).SetBytes(getData(input, 0, 32)).Uint64()
		expLen  = new(big.Int).SetBytes(getData(input, 32, 32)).Uint64()
		modLen  = new(big.Int).SetBytes(getData(input, 64, 32)).Uint64()
	)
	if len(input) > 96 {
		input = input[96:]
	} else {
		input = input[:0]
	}
	// Handle a special case when both the base and mod length is zero
	if baseLen == 0 && modLen == 0 {
		return []byte{}, nil
	}
	// Retrieve the operands and execute the exponentiation
	var (
		base = new(big.Int).SetBytes(getData(input, 0, baseLen))
		exp  = new(big.Int).SetBytes(getData(input, baseLen, expLen))
		mod  = new(big.Int).SetBytes(getData(input, baseLen+expLen, modLen))
	)
	if mod.BitLen() == 0 {
		// Modulo 0 is undefined, return zero
		return common.LeftPadBytes([]byte{}, int(modLen)), nil
	}
	return common.LeftPadBytes(base.Exp(base, exp, mod).Bytes(), int(modLen)), nil
}

var (
	// errNotOnCurve is returned if a point being unmarshalled as a bn256 elliptic
	// curve point is not on the curve.
	errNotOnCurve = errors.New("point not on elliptic curve")

	// errInvalidCurvePoint is returned if a point being unmarshalled as a bn256
	// elliptic curve point is invalid.
	errInvalidCurvePoint = errors.New("invalid elliptic curve point")
)

// newCurvePoint unmarshals a binary blob into a bn256 elliptic curve point,
// returning it, or an error if the point is invalid.
func newCurvePoint(blob []byte) (*bn256.G1, error) {
	p, onCurve := new(bn256.G1).Unmarshal(blob)
	if !onCurve {
		return nil, errNotOnCurve
	}
	gx, gy, _, _ := p.CurvePoints()
	if gx.Cmp(bn256.P) >= 0 || gy.Cmp(bn256.P) >= 0 {
		return nil, errInvalidCurvePoint
	}
	return p, nil
}

// newTwistPoint unmarshals a binary blob into a bn256 elliptic curve point,
// returning it, or an error if the point is invalid.
func newTwistPoint(blob []byte) (*bn256.G2, error) {
	p, onCurve := new(bn256.G2).Unmarshal(blob)
	if !onCurve {
		return nil, errNotOnCurve
	}
	x2, y2, _, _ := p.CurvePoints()
	if x2.Real().Cmp(bn256.P) >= 0 || x2.Imag().Cmp(bn256.P) >= 0 ||
		y2.Real().Cmp(bn256.P) >= 0 || y2.Imag().Cmp(bn256.P) >= 0 {
		return nil, errInvalidCurvePoint
	}
	return p, nil
}

// bn256Add implements a native elliptic curve point addition.
type bn256Add struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256Add) RequiredGas(input []byte) uint64 {
	return params.Bn256AddGas
}

func (c *bn256Add) Run(input []byte) ([]byte, error) {
	x, err := newCurvePoint(getData(input, 0, 64))
	if err != nil {
		return nil, err
	}
	y, err := newCurvePoint(getData(input, 64, 64))
	if err != nil {
		return nil, err
	}
	res := new(bn256.G1)
	res.Add(x, y)
	return res.Marshal(), nil
}

// bn256ScalarMul implements a native elliptic curve scalar multiplication.
type bn256ScalarMul struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256ScalarMul) RequiredGas(input []byte) uint64 {
	return params.Bn256ScalarMulGas
}

func (c *bn256ScalarMul) Run(input []byte) ([]byte, error) {
	p, err := newCurvePoint(getData(input, 0, 64))
	if err != nil {
		return nil, err
	}
	res := new(bn256.G1)
	res.ScalarMult(p, new(big.Int).SetBytes(getData(input, 64, 32)))
	return res.Marshal(), nil
}

var (
	// true32Byte is returned if the bn256 pairing check succeeds.
	true32Byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	// false32Byte is returned if the bn256 pairing check fails.
	false32Byte = make([]byte, 32)

	// errBadPairingInput is returned if the bn256 pairing input is invalid.
	errBadPairingInput = errors.New("bad elliptic curve pairing size")
)

// bn256Pairing implements a pairing pre-compile for the bn256 curve
type bn256Pairing struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256Pairing) RequiredGas(input []byte) uint64 {
	return params.Bn256PairingBaseGas + uint64(len(input)/192)*params.Bn256PairingPerPointGas
}

func (c *bn256Pairing) Run(input []byte) ([]byte, error) {
	// Handle some corner cases cheaply
	if len(input)%192 > 0 {
		return nil, errBadPairingInput
	}
	// Convert the input into a set of coordinates
	var (
		cs []*bn256.G1
		ts []*bn256.G2
	)
	for i := 0; i < len(input); i += 192 {
		c, err := newCurvePoint(input[i : i+64])
		if err != nil {
			return nil, err
		}
		t, err := newTwistPoint(input[i+64 : i+192])
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
		ts = append(ts, t)
	}
	// Execute the pairing checks and return the results
	if bn256.PairingCheck(cs, ts) {
		return true32Byte, nil
	}
	return false32Byte, nil
}

// ZSL START

// SHA256Compress implemented as a native contract
type sha256Compress struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *sha256Compress) RequiredGas(input []byte) uint64 {
	return params.ZSLGas
}

/*
	Input bytes when the precompile is called with string "hello":
	0000000000000000000000000000000000000000000000000000000000000020
	0000000000000000000000000000000000000000000000000000000000000005
	68656c6c6f000000000000000000000000000000000000000000000000000000
*/
func (c *sha256Compress) Run(in []byte) ([]byte, error) {
	// ignore keccac
	in = in[4:]

	// ignore next 32 bytes
	in = in[32:]

	// check payload size
	n := binary.BigEndian.Uint64(in[24:32])
	if n != 64 {
		msg := "ZSL input must have size of 64 bytes (512 bits)"
		log.Error(msg)
		return []byte{}, errors.New(msg)
	}

	// skip payload size
	in = in[32:]

	h := sha256compress.NewCompress()
	h.Write(in[0:n])
	return h.Compress(), nil
}

// verifyShielding implemented as a native contract
type verifyShielding struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *verifyShielding) RequiredGas(input []byte) uint64 {
	return params.ZSLGas
}

/**
In geth:
zslprecompile.VerifyShielding("0x001122", "0x08dbb5c1357d05e5178c9f8b88b590e0728d36f1a2e04ae93e963d5174fc4d35", "0xff2c9bdc59089c8d3aa313e9394a19ea17dbfa6f8b2520c7165734b6da615dc4", 12345)

Data passed into function:
4e320263000000000000000000000000000000000000000000000000000000000000000808dbb5c1357d05e5178c9f8b88b590e0728d36f1a2e04ae93e963d5174fc4d35ff2c9bdc59089c8d3aa313e9394a19ea17dbfa6f8b2520c7165734b6da615dc400000000000000000000000000000000000000000000000000000000000030390000000000000000000000000000000000000000000000000000000000000003001122
*/
func (c *verifyShielding) Run(in []byte) ([]byte, error) {
	snark.Init()

	// ignore keccac
	in = in[4:]

	// ignore next 32 bytes
	in = in[32:]

	var send_nf [32]byte
	var cm [32]byte
	copy(send_nf[:], in[:32])
	copy(cm[:], in[32:64])
	noteValue := binary.BigEndian.Uint64(in[88:96])
	proofSize := binary.BigEndian.Uint64(in[120:128]) // should be 584

	if proofSize != ZSL_PROOF_SIZE {
		msg := fmt.Sprintf("ZSL error, proof must have size of %d bytes, not %d.\n", ZSL_PROOF_SIZE, proofSize)
		log.Error(msg)
		return []byte{}, errors.New(msg)
	}

	var proof [ZSL_PROOF_SIZE]byte
	copy(proof[:], in[128:])

	result := snark.VerifyShielding(proof, send_nf, cm, noteValue)
	var b byte
	if result {
		b = 1
	}

	log.Info("verifyShieldingFunc: ", hex.EncodeToString(in))
	log.Info("send_nf: ", hex.EncodeToString(send_nf[:]))
	log.Info("     cm: ", hex.EncodeToString(cm[:]))
	log.Info("  value: ", noteValue)
	log.Info("   size: ", proofSize)
	log.Info("  proof: ", hex.EncodeToString(in[128:]))
	log.Info(" result: ", result)

	return []byte{b}, nil
}

// bn256Add implements a native elliptic curve point addition.
type verifyUnshielding struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *verifyUnshielding) RequiredGas(input []byte) uint64 {
	return params.ZSLGas
}

func (c *verifyUnshielding) Run(in []byte) ([]byte, error) {
	snark.Init()

	// ignore keccac
	in = in[4:]

	// ignore next 32 bytes
	in = in[32:]

	var spend_nf [32]byte
	var rt [32]byte
	var addr [20]byte
	copy(spend_nf[:], in[:32])
	copy(rt[:], in[32:64])
	copy(addr[:], in[76:96])  // type address === uint160
	noteValue := binary.BigEndian.Uint64(in[120:128])
	proofSize := binary.BigEndian.Uint64(in[152:160]) // should be 584

	if proofSize != ZSL_PROOF_SIZE {
		msg := fmt.Sprintf("ZSL error, proof must have size of %d bytes, not %d.\n", ZSL_PROOF_SIZE, proofSize)
		log.Error(msg)
		return []byte{}, errors.New(msg)
	}

	var proof [ZSL_PROOF_SIZE]byte
	copy(proof[:], in[160:])

	result := snark.VerifyUnshielding(proof, spend_nf, rt, addr, noteValue)
	var b byte
	if result {
		b = 1
	}

	log.Info("verifyUnshieldingFunc: ", hex.EncodeToString(in))
	log.Info("spend_nf: ", hex.EncodeToString(spend_nf[:]))
	log.Info("      rt: ", hex.EncodeToString(rt[:]))
	log.Info("   value: ", noteValue)
	log.Info("    size: ", proofSize)
	log.Info("   proof: ", hex.EncodeToString(in[128:]))
	log.Info("  result: ", result)

	return []byte{b}, nil
}

// verifyShieldedTransfer implements a native elliptic curve point addition.
type verifyShieldedTransfer struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *verifyShieldedTransfer) RequiredGas(input []byte) uint64 {
	return params.ZSLGas
}

func (c *verifyShieldedTransfer) Run(in []byte) ([]byte, error) {
	// ignore keccac
	in = in[4:]

	// ignore next 32 bytes
	in = in[32:]

	var anchor [32]byte
	var spend_nf_1 [32]byte
	var spend_nf_2 [32]byte
	var send_nf_1 [32]byte
	var send_nf_2 [32]byte
	var cm_1 [32]byte
	var cm_2 [32]byte
	copy(anchor[:], in[:32])
	copy(spend_nf_1[:], in[32:64])
	copy(spend_nf_2[:], in[64:96])
	copy(send_nf_1[:], in[96:128])
	copy(send_nf_2[:], in[128:160])
	copy(cm_1[:], in[160:192])
	copy(cm_2[:], in[192:224])
	proofSize := binary.BigEndian.Uint64(in[248:256]) // should be 584

	if proofSize != ZSL_PROOF_SIZE {
		msg := fmt.Sprintf("ZSL error, proof must have size of %d bytes, not %d.\n", ZSL_PROOF_SIZE, proofSize)
		log.Error(msg)
		return []byte{}, errors.New(msg)
	}

	var proof [ZSL_PROOF_SIZE]byte
	copy(proof[:], in[256:])

	snark.Init()
	result := snark.VerifyTransfer(proof, anchor, spend_nf_1, spend_nf_2, send_nf_1, send_nf_2, cm_1, cm_2)
	var b byte
	if result {
		b = 1
	}

	log.Info("verifyShieldedTransferFunc: ", hex.EncodeToString(in))
	log.Info("spend_nf_1: ", hex.EncodeToString(spend_nf_1[:]))
	log.Info("spend_nf_2: ", hex.EncodeToString(spend_nf_2[:]))
	log.Info(" send_nf_1: ", hex.EncodeToString(send_nf_1[:]))
	log.Info(" send_nf_2: ", hex.EncodeToString(send_nf_2[:]))
	log.Info("      cm_1: ", hex.EncodeToString(cm_1[:]))
	log.Info("      cm_2: ", hex.EncodeToString(cm_2[:]))
	log.Info("    anchor: ", hex.EncodeToString(anchor[:]))
	log.Info("      size: ", proofSize)
	log.Info("     proof: ", hex.EncodeToString(proof[:]))
	log.Info("    result: ", result)

	return []byte{b}, nil
}

// ZSL END

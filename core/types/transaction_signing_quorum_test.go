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

package types

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	testifyassert "github.com/stretchr/testify/assert"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// run all the tests in this file
// $> go test $(go list ./...) -run TestSignQuorum

// private key material to test both 0 and 1 bit for the recoveryId (v).
// key with v sign == 28 (Homestead)
var k0v, _ = new(big.Int).SetString("25807260602402504536675820444142779248993100028628438487502323668296269534891", 10)

// key with v sign == 27 (Homestead)
var k1v, _ = new(big.Int).SetString("10148397294747000913768625849546502595195728826990639993137198410557736548965", 10)

// helper to deterministically create an ECDSA key from an int.
func createKey(c elliptic.Curve, k *big.Int) (*ecdsa.PrivateKey, error) {
	sk := new(ecdsa.PrivateKey)
	sk.PublicKey.Curve = c
	sk.D = k
	sk.PublicKey.X, sk.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return sk, nil
}

func signTx(key *ecdsa.PrivateKey, signer Signer) (*Transaction, common.Address, error) {
	addr := crypto.PubkeyToAddress(key.PublicKey)
	tx := NewTransaction(0, addr, new(big.Int), 0, new(big.Int), nil)
	signedTx, err := SignTx(tx, signer, key)
	//fmt.Printf("\ntx.data.V signTx after sign [%v] \n", signedTx.data.V)
	return signedTx, addr, err
}

/**
 * As of quorum v2.2.3 commit be7cc31ce208525ea1822e7d0fee88bf7f14500b 30 April 2019 behavior
 *
 * Test public transactions signed by homestead Signer. Homestead sets the v param on a signed transaction to
 * either 27 or 28. The v parameter is used for recovering the sender of the signed transation.
 *
 *  1. Homestead: should be 27, 28
 * $> go test -run TestSignQuorumHomesteadPublic
 */
func TestSignQuorumHomesteadPublic(t *testing.T) {

	assert := testifyassert.New(t)

	k0, _ := createKey(crypto.S256(), k0v)
	k1, _ := createKey(crypto.S256(), k1v)

	homeSinger := HomesteadSigner{}

	// odd parity should be 27 for Homestead
	signedTx, addr, _ := signTx(k1, homeSinger)

	assert.True(signedTx.data.V.Cmp(big.NewInt(27)) == 0, fmt.Sprintf("v wasn't 27 it was [%v]", signedTx.data.V))

	// recover address from signed TX
	from, _ := Sender(homeSinger, signedTx)
	//fmt.Printf("from [%v] == addr [%v]\n\n", from, from == addr)
	assert.True(from == addr, fmt.Sprintf("Expected from and address to be equal. Got %x want %x", from, addr))

	// even parity should be 28 for Homestead
	signedTx, addr, _ = signTx(k0, homeSinger)
	assert.True(signedTx.data.V.Cmp(big.NewInt(28)) == 0, fmt.Sprintf("v wasn't 28 it was [%v]\n", signedTx.data.V))

	// recover address from signed TX
	from, _ = Sender(homeSinger, signedTx)
	//fmt.Printf("from [%v] == addr [%v]\n", from, from == addr)
	assert.True(from == addr, fmt.Sprintf("Expected from and address to be equal. Got %x want %x", from, addr))

}

/**
 * As of quorum v2.2.3 commit be7cc31ce208525ea1822e7d0fee88bf7f14500b 30 April 2019 behavior
 *
 * Test the public transactions signed by the EIP155Signer.
 * The EIP155Signer was introduced to protect against replay
 * attacks https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md and stores
 * the CHAINID in the signed transaction's `v` parameter as `v = chainId * 2 + 35`.
 *
 * The EthEIP155Signer change breaks private quorum transactions when the chainId == 1 (mainnet chainId),
 * as the v parameter on a public transaction and on a private transaction will both be 37, 38.
 *
 *  $> go test -run TestSignQuorumEIP155Public
 */
func TestSignQuorumEIP155Public(t *testing.T) {

	assert := testifyassert.New(t)

	k0, _ := createKey(crypto.S256(), k0v)
	k1, _ := createKey(crypto.S256(), k1v)

	// chainId 1 even EIP155Signer should be 37 conflicts with private transaction
	var chainId int64
	chainId = 2 // 7 2 10

	v0 := chainId*2 + 35 // sig[64] + 35 .. where sig[64] == 0
	v1 := chainId*2 + 36 // sig[64] + 35 .. where sig[64] == 1

	// Will calculate v to be `v = CHAINID * 2 + 35`
	// To compute V:
	// 	2 * 2 + 35 == 39
	// 	2 * 2 + 36 == 40
	// To retrieve Sender, pull out 27, 28 Eth Frontier / Homestead values.
	// 	39 - (2 * 2) - 8 == 27
	// 	40 - (2 * 2) - 8 == 28
	EIPsigner := NewEIP155Signer(big.NewInt(chainId))

	signedTx, addr, _ := signTx(k0, EIPsigner)

	//fmt.Printf("After signing V is [%v] \n", signedTx.data.V)
	assert.True(signedTx.data.V.Cmp(big.NewInt(v0)) == 0, fmt.Sprintf("v wasn't [%v] it was [%v]\n", v0, signedTx.data.V))
	from, _ := Sender(EIPsigner, signedTx)

	assert.True(from == addr, fmt.Sprintf("Expected from and address to be equal. Got %x want %x", from, addr))

	// chainId 1 even  EIP155Signer should be 38 conflicts with private transaction
	assert.False(signedTx.IsPrivate(), fmt.Sprintf("Public transaction is set to a private transation v == [%v]", signedTx.data.V))

	signedTx, addr, _ = signTx(k1, EIPsigner)

	assert.True(signedTx.data.V.Cmp(big.NewInt(v1)) == 0, fmt.Sprintf("v wasn't [%v], it was [%v]\n", v1, signedTx.data.V))
	from, _ = Sender(EIPsigner, signedTx)

	assert.True(from == addr, fmt.Sprintf("Expected from and address to be equal. Got %x want %x", from, addr))

}

/**
 *  As of quorum v2.2.3 commit be7cc31ce208525ea1822e7d0fee88bf7f14500b 30 April 2019 behavior
 *
 * When the signer is EIP155Signer, chainId == 1 cannot be used because the EIP155 computed `v` value conflicts
 * with the private `v` value that quorum uses to indicate a private transaction: v == 37 and v == 38.
 *
 *  $> go test -run TestSignQuorumEIP155FailPublicChain1
 */
func TestSignQuorumEIP155FailPublicChain1(t *testing.T) {

	assert := testifyassert.New(t)

	k0, _ := createKey(crypto.S256(), k0v)
	k1, _ := createKey(crypto.S256(), k1v)

	// chainId 1 even  EIP155Signer should be 37.38 which conflicts with private transaction
	var chainId int64
	chainId = 1

	v0 := chainId*2 + 35 // sig[64] + 35 .. where sig[64] == 0
	v1 := chainId*2 + 36 // sig[64] + 35 .. where sig[64] == 1

	// Will calculate v to be `v = CHAINID * 2 + 35`
	// To compute V:
	// 	2 * 1 + 35 == 37
	// 	2 * 1 + 36 == 38
	// To retrieve Sender, pull out 27, 28 Eth Frontier / Homestead values.
	// 	37 - (1 * 2) - 8 == 27
	// 	38 - (1 * 2) - 8 == 28
	EIPsigner := NewEIP155Signer(big.NewInt(chainId))

	signedTx, addr, _ := signTx(k0, EIPsigner)

	// the calculated v value should equal `chainId * 2 + 35 `
	assert.True(signedTx.data.V.Cmp(big.NewInt(v0)) == 0, fmt.Sprintf("v wasn't [%v] it was "+
		"[%v]\n", v0, signedTx.data.V))
	// the sender will not be equal as HomesteadSigner{}.Sender(tx) is used because IsPrivate() will be true
	// although it is a public tx.
	// This is test to catch when / if this behavior changes.
	assert.True(signedTx.IsPrivate(), "A public transaction with EIP155 and chainID 1 is expected to be "+
		"considered private, as its v param conflict with a private transaction. signedTx.IsPrivate() == [%v]", signedTx.IsPrivate())
	from, _ := Sender(EIPsigner, signedTx)

	assert.False(from == addr, fmt.Sprintf("Expected the sender of a public TX from chainId 1, \n "+
		"should not be recoverable from [%x] addr [%v] ", from, addr))

	signedTx, addr, _ = signTx(k1, EIPsigner)

	// the calculated v value should equal `chainId * 2 + 35`
	assert.True(signedTx.data.V.Cmp(big.NewInt(v1)) == 0,
		fmt.Sprintf("v wasn't [%v] it was [%v]", v1, signedTx.data.V))

	// the sender will not be equal as HomesteadSigner{}.Sender(tx) is used because IsPrivate() will be true
	// although it is a public tx.
	// This is test to catch when / if this behavior changes.
	// we are signing the data with EIPsigner and chainID 1, so this would be considered a private tx.
	assert.True(signedTx.IsPrivate(), "A public transaction with EIP155 and chainID 1 is expected to "+
		"to be considered private, as its v param conflict with a private transaction. "+
		"signedTx.IsPrivate() == [%v]", signedTx.IsPrivate())
	from, _ = Sender(EIPsigner, signedTx)

	assert.False(from == addr, fmt.Sprintf("Expected the sender of a public TX from chainId 1, "+
		"should not be recoverable from [%x] addr [%v] ", from, addr))

}

/**
*  As of quorum v2.2.3 commit be7cc31ce208525ea1822e7d0fee88bf7f14500b 30 April 2019 behavior
*
*  Use Homestead to sign and EIPSigner to recover.
*
*  SendTransaction creates a transaction for the given argument, signs it and submit it to the transaction pool.
*  func (s *PublicTransactionPoolAPI) SendTransaction(ctx context.Context, args SendTxArgs) (common.Hash, error) {
*  Current implementation in `internal/ethapi/api.go`
*
*  accounts/keystore/keystore.SignTx(): would hash and sign with homestead
*
*  When a private tx (obtained from json params PrivateFor) is submitted `internal/ethapi/api.go`:
*
*  1. sign with HomesteadSigner, this will set the v parameter to
*     27 or 28. // there is no indication that this is a private tx yet.
*
*  2. when submitting a transaction `submitTransaction(ctx context.Context, b Backend, tx *types.Transaction, isPrivate bool)`
      check isPrivate param, and call `tx.SetPrivate()`, this will update the `v` signature param (recoveryID)
*     from 27 -> 37, 28 -> 38. // this is now considered a private tx.
*
*  $> go test -run TestSignQuorumHomesteadEIP155SigningPrivateQuorum
*/
func TestSignQuorumHomesteadEIP155SigningPrivateQuorum(t *testing.T) {

	assert := testifyassert.New(t)

	keys := []*big.Int{k0v, k1v}

	homeSinger := HomesteadSigner{}
	recoverySigner := NewEIP155Signer(big.NewInt(18))

	// check for both sig[64] == 0, and sig[64] == 1
	for i := 0; i < len(keys); i++ {
		key, _ := createKey(crypto.S256(), keys[i])
		signedTx, addr, err := signTx(key, homeSinger)

		assert.Nil(err, err)
		// set to privateTX after the intial signing, this explicitly sets the v param.
		// Note: only works when the tx was signed with the homesteadSinger (v==27 | 28).
		signedTx.SetPrivate()

		assert.True(signedTx.IsPrivate(), fmt.Sprintf("Expected the transaction to be private [%v]", signedTx.IsPrivate()))
		// Try to recover Sender
		from, err := Sender(recoverySigner, signedTx)

		assert.Nil(err, err)
		assert.True(from == addr, fmt.Sprintf("Expected from and address to be equal. Got %x want %x", from, addr))
	}

}

/*
 * As of quorum v2.2.3 commit be7cc31ce208525ea1822e7d0fee88bf7f14500b 30 April 2019 behavior
 * Use Homestead to sign and Homestead to recover.
 *
 * Signing private transactions with HomesteadSigner, and recovering a private transaction with
 * HomesteadSigner works, but the transaction has to be set to private `signedTx.SetPrivate()` after
 * the signature and before recovering the address.
 *
 *  $> go test -run TestSignQuorumHomesteadOnlyPrivateQuorum
 */
func TestSignQuorumHomesteadOnlyPrivateQuorum(t *testing.T) {

	assert := testifyassert.New(t)

	// check even and odd parity
	keys := []*big.Int{k0v, k1v}

	homeSinger := HomesteadSigner{}
	recoverySigner := HomesteadSigner{}

	for i := 0; i < len(keys); i++ {
		key, _ := createKey(crypto.S256(), keys[i])
		signedTx, addr, err := signTx(key, homeSinger)

		assert.Nil(err, err)

		//fmt.Printf("Private tx.data.V Home [%v] \n", signedTx.data.V)
		// set to privateTX after the initial signing.
		signedTx.SetPrivate()
		assert.True(signedTx.IsPrivate(), fmt.Sprintf("Expected the transaction to be "+
			"private [%v]", signedTx.IsPrivate()))
		//fmt.Printf("Private tx.data.V Home [%v] \n", signedTx.data.V)

		// Try to recover Sender
		from, err := Sender(recoverySigner, signedTx)

		assert.Nil(err, err)
		assert.True(from == addr, fmt.Sprintf("Expected from and address to be equal. "+
			" Got %x want %x", from, addr))
	}

}

/*
 * As of quorum v2.2.3 commit be7cc31ce208525ea1822e7d0fee88bf7f14500b 30 April 2019 behavior
 *
 * Use EIP155 to sign and EIP155 to recover (This is not a valid combination and does **not** work).
 *
 * Signing private transactions with EIP155Signer, and recovering a private transaction with
 * EIP155Signer does **not** work.
 * note: deriveChainId only checks for 27, 28 when using EIP155
 * note: In the case where the v param is not 27 or 28 when setting private it will always be set to 37
 *
 *  $> go test -run TestSignQuorumEIP155OnlyPrivateQuorum
 */
func TestSignQuorumEIP155OnlyPrivateQuorum(t *testing.T) {

	assert := testifyassert.New(t)

	// check even and odd parity
	keys := []*big.Int{k0v, k1v}

	EIP155Signer := NewEIP155Signer(big.NewInt(0))

	for i := 0; i < len(keys); i++ {
		key, _ := createKey(crypto.S256(), keys[i])
		signedTx, addr, err := signTx(key, EIP155Signer)

		assert.Nil(err, err)
		//fmt.Printf("Private tx.data.V Home [%v] \n", signedTx.data.V)

		// set to privateTX after the initial signing.
		signedTx.SetPrivate()

		assert.True(signedTx.IsPrivate(), fmt.Sprintf("Expected the transaction to be private [%v]", signedTx.IsPrivate()))
		//fmt.Printf("Private tx.data.V Home [%v] \n", signedTx.data.V)

		// Try to recover Sender
		from, err := Sender(EIP155Signer, signedTx)

		assert.Nil(err, err)
		assert.False(from == addr, fmt.Sprintf("Expected recovery to fail. from [%x] should not equal "+
			"addr [%x]", from, addr))

	}

}

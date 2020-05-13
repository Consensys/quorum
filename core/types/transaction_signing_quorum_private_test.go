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
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	testifyassert "github.com/stretchr/testify/assert"
)

func signTxWithSigner(signer Signer, key *ecdsa.PrivateKey) (*Transaction, common.Address, error) {
	addr := crypto.PubkeyToAddress(key.PublicKey)
	tx := NewTransaction(0, addr, new(big.Int), 0, new(big.Int), nil)
	signedTx, err := SignTx(tx, signer, key)
	return signedTx, addr, err
}

// run all the tests in this file
// $> go test $(go list ./...) -run QuorumSignPrivate

// test with QuorumPrivateSigner
/*
*  $> go test -run TestQuorumSignPrivateQuorum
 */
func TestQuorumSignPrivateQuorum(t *testing.T) {

	assert := testifyassert.New(t)
	keys := []*big.Int{k0v, k1v}

	for i := 0; i < len(keys); i++ {
		key, _ := createKey(crypto.S256(), keys[i])
		qpPrivateSigner := QuorumPrivateTxSigner{HomesteadSigner{}}

		signedTx, addr, err := signTxWithSigner(qpPrivateSigner, key)
		assert.Nil(err, err)
		assert.True(signedTx.IsPrivate(),
			fmt.Sprintf("The signed transaction is not private, signedTx.data.V is [%v]", signedTx.data.V))
		from, err := Sender(qpPrivateSigner, signedTx)
		assert.Nil(err, err)
		assert.True(from == addr, fmt.Sprintf("Expected from == address, [%x] == [%x]", from, addr))
	}

}

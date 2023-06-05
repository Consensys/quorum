// Copyright 2016 The go-ethereum Authors
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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Signs with Homestead
// obtains sender from EIP55Signer
type QuorumPrivateTxSigner struct{ HomesteadSigner }

func (s QuorumPrivateTxSigner) Sender(tx *Transaction) (common.Address, error) {
	return HomesteadSigner{}.Sender(tx)
}

// SignatureValues returns signature values. This signature
// needs to be in the [R || S || V] format where V is 0 or 1.
func (qs QuorumPrivateTxSigner) SignatureValues(tx *Transaction, sig []byte) (R, S, V *big.Int, err error) {
	r, s, _, _ := HomesteadSigner{}.SignatureValues(tx, sig)
	// update v for private transaction marker: needs to be 37 (0+37) or 38 (1+37) for a private transaction.
	v := new(big.Int).SetBytes([]byte{sig[64] + 37})
	return r, s, v, nil
}

// Hash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (s QuorumPrivateTxSigner) Hash(tx *Transaction) common.Hash {
	return s.HomesteadSigner.Hash(tx)
}

func (s QuorumPrivateTxSigner) Equal(s2 Signer) bool {
	_, ok := s2.(QuorumPrivateTxSigner)
	return ok
}

/*
 * If v is `37` or `38` that marks the transaction as private in Quorum.
 * Note: this means quorum chains cannot have a public ethereum chainId == 1, as the EIP155 v
 * param is `37` and `38` for the public Ethereum chain. Having a private chain with a chainId ==1
 * is discouraged in the general Ethereum ecosystem.
 */
func isPrivate(v *big.Int) bool {
	return v.Cmp(big.NewInt(37)) == 0 || v.Cmp(big.NewInt(38)) == 0
}

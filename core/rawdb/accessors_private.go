// Copyright 2018 The go-ethereum Authors
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

package rawdb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// The fields below define the low level database schema prefixing.
var (
	privateRootPrefix          = []byte("P")
	privateblockReceiptsPrefix = []byte("Pr") // blockReceiptsPrefix + num (uint64 big endian) + hash -> block receipts
	privateReceiptPrefix       = []byte("Prs")
	privateBloomPrefix         = []byte("Pb")
)

//GetPrivateStateRoot utility to get private hate root hash
func GetPrivateStateRoot(db DatabaseReader, blockRoot common.Hash) common.Hash {
	root, _ := db.Get(append(privateRootPrefix, blockRoot[:]...))
	return common.BytesToHash(root)
}

//WritePrivateStateRoot utility to write private root hash
func WritePrivateStateRoot(db DatabaseWriter, blockRoot, root common.Hash) error {
	return db.Put(append(privateRootPrefix, blockRoot[:]...), root[:])
}

// WritePrivateBlockBloom creates a bloom filter for the given receipts and saves it to the database
// with the number given as identifier (i.e. block number).
func WritePrivateBlockBloom(db DatabaseWriter, number uint64, receipts types.Receipts) error {
	rbloom := types.CreateBloom(receipts)
	return db.Put(append(privateBloomPrefix, encodeBlockNumber(number)...), rbloom[:])
}

// GetPrivateBlockBloom retrieves the private bloom associated with the given number.
func GetPrivateBlockBloom(db DatabaseReader, number uint64) (bloom types.Bloom) {
	data, _ := db.Get(append(privateBloomPrefix, encodeBlockNumber(number)...))
	if len(data) > 0 {
		bloom = types.BytesToBloom(data)
	}
	return bloom
}

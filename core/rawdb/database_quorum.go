/*


// Copyright 2015 The go-ethereum Authors

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
*/
package rawdb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	privateRootPrefix           = []byte("P")
	privateBloomPrefix          = []byte("Pb")
	quorumEIP155ActivatedPrefix = []byte("quorum155active")
)

//returns whether we have a chain configuration that can't be updated
//after the EIP155 HF has happened
func GetIsQuorumEIP155Activated(db ethdb.KeyValueReader) bool {
	data, _ := db.Get(quorumEIP155ActivatedPrefix)
	return len(data) == 1
}

// WriteQuorumEIP155Activation writes a flag to the database saying EIP155 HF is enforced
func WriteQuorumEIP155Activation(db ethdb.KeyValueWriter) error {
	return db.Put(quorumEIP155ActivatedPrefix, []byte{1})
}

func GetPrivateStateRoot(db ethdb.Database, blockRoot common.Hash) common.Hash {
	root, _ := db.Get(append(privateRootPrefix, blockRoot[:]...))
	return common.BytesToHash(root)
}

func WritePrivateStateRoot(db ethdb.Database, blockRoot, root common.Hash) error {
	return db.Put(append(privateRootPrefix, blockRoot[:]...), root[:])
}

// WritePrivateBlockBloom creates a bloom filter for the given receipts and saves it to the database
// with the number given as identifier (i.e. block number).
func WritePrivateBlockBloom(db ethdb.Database, number uint64, receipts types.Receipts) error {
	rbloom := types.CreateBloom(receipts)
	return db.Put(append(privateBloomPrefix, encodeBlockNumber(number)...), rbloom[:])
}

// GetPrivateBlockBloom retrieves the private bloom associated with the given number.
func GetPrivateBlockBloom(db ethdb.Database, number uint64) (bloom types.Bloom) {
	data, _ := db.Get(append(privateBloomPrefix, encodeBlockNumber(number)...))
	if len(data) > 0 {
		bloom = types.BytesToBloom(data)
	}
	return bloom
}

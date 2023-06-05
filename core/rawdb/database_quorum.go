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
	privateStatesTrieRootPrefix = []byte("PSTP")
	privateBloomPrefix          = []byte("Pb")
	quorumEIP155ActivatedPrefix = []byte("quorum155active")
	// Quorum
	// we introduce a generic approach to store extra data for an account. PrivacyMetadata is wrapped.
	// However, this value is kept as-is to support backward compatibility
	stateRootToExtraDataRootPrefix = []byte("PSR2PMDR")
	// emptyRoot is the known root hash of an empty trie. Duplicate from `trie/trie.go#emptyRoot`
	emptyRoot = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
)

// returns whether we have a chain configuration that can't be updated
// after the EIP155 HF has happened
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

func GetPrivateStatesTrieRoot(db ethdb.Database, blockRoot common.Hash) common.Hash {
	root, _ := db.Get(append(privateStatesTrieRootPrefix, blockRoot[:]...))
	return common.BytesToHash(root)
}

func GetAccountExtraDataRoot(db ethdb.KeyValueReader, stateRoot common.Hash) common.Hash {
	root, _ := db.Get(append(stateRootToExtraDataRootPrefix, stateRoot[:]...))
	return common.BytesToHash(root)
}

func WritePrivateStateRoot(db ethdb.Database, blockRoot, root common.Hash) error {
	return db.Put(append(privateRootPrefix, blockRoot[:]...), root[:])
}

func WritePrivateStatesTrieRoot(db ethdb.Database, blockRoot, root common.Hash) error {
	return db.Put(append(privateStatesTrieRootPrefix, blockRoot[:]...), root[:])
}

// WriteRootHashMapping stores the mapping between root hash of state trie and
// root hash of state.AccountExtraData trie to persistent storage
func WriteRootHashMapping(db ethdb.KeyValueWriter, stateRoot, extraDataRoot common.Hash) error {
	return db.Put(append(stateRootToExtraDataRootPrefix, stateRoot[:]...), extraDataRoot[:])
}

// WritePrivateBlockBloom creates a bloom filter for the given receipts and saves it to the database
// with the number given as identifier (i.e. block number).
func WritePrivateBlockBloom(db ethdb.Database, number uint64, receipts types.Receipts) error {
	rbloom := types.CreateBloom(receipts.Flatten())
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

// AccountExtraDataLinker maintains mapping between root hash of the state trie
// and root hash of state.AccountExtraData trie
type AccountExtraDataLinker interface {
	// GetAccountExtraDataRoot returns the root hash of the state.AccountExtraData trie from
	// the given root hash of the state trie.
	//
	// It returns an empty hash if not found.
	GetAccountExtraDataRoot(stateRoot common.Hash) common.Hash
	// Link saves the mapping between root hash of the state trie and
	// root hash of state.AccountExtraData trie to the persistent storage.
	// Don't write the mapping if extraDataRoot is an emptyRoot
	Link(stateRoot, extraDataRoot common.Hash) error
}

// ethdbAccountExtraDataLinker implements AccountExtraDataLinker using ethdb.Database
// as the persistence storage
type ethdbAccountExtraDataLinker struct {
	db ethdb.Database
}

func NewAccountExtraDataLinker(db ethdb.Database) AccountExtraDataLinker {
	return &ethdbAccountExtraDataLinker{
		db: db,
	}
}

func (pml *ethdbAccountExtraDataLinker) GetAccountExtraDataRoot(stateRoot common.Hash) common.Hash {
	return GetAccountExtraDataRoot(pml.db, stateRoot)
}

func (pml *ethdbAccountExtraDataLinker) Link(stateRoot, extraDataRoot common.Hash) error {
	if extraDataRoot != emptyRoot {
		return WriteRootHashMapping(pml.db, stateRoot, extraDataRoot)
	}
	return nil
}

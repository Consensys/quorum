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

package rawdb

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/stretchr/testify/assert"
)

// Tests that setting the flag for Quorum EIP155 activation read values correctly
func TestIsQuorumEIP155Active(t *testing.T) {
	db := NewMemoryDatabase()

	isQuorumEIP155Active := GetIsQuorumEIP155Activated(db)
	if isQuorumEIP155Active {
		t.Fatal("Quorum EIP155 active read to be set, but wasn't set beforehand")
	}

	dbSet := NewMemoryDatabase()
	err := WriteQuorumEIP155Activation(dbSet)

	if err != nil {
		t.Fatal("unable to write quorum EIP155 activation")
	}

	isQuorumEIP155ActiveAfterSetting := GetIsQuorumEIP155Activated(dbSet)
	if !isQuorumEIP155ActiveAfterSetting {
		t.Fatal("Quorum EIP155 active read to be unset, but was set beforehand")
	}
}

func TestAccountExtraDataLinker_whenLinkingEmptyRoot(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}

	linker := NewAccountExtraDataLinker(db)

	err := linker.Link(psr, emptyRoot)

	if err != nil {
		t.Fatal("unable to store the link")
	}

	value, _ := db.Get(append(stateRootToExtraDataRootPrefix, psr[:]...))

	if value != nil {
		t.Fatal("the mapping should not have been stored")
	}
}

func TestAccountExtraDataLinker_whenLinkingRoots(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}
	pmr := common.Hash{2}

	linker := NewAccountExtraDataLinker(db)

	err := linker.Link(psr, pmr)

	if err != nil {
		t.Fatal("unable to store the link")
	}

	value, _ := db.Get(append(stateRootToExtraDataRootPrefix, psr[:]...))

	if value == nil {
		t.Fatal("the mapping should have been stored")
	}

	valueHash := common.BytesToHash(value)

	if pmr != valueHash {
		t.Fatal("the privacy metadata root does not have the expected value")
	}
}

var errReadOnly = errors.New("unable to write")

type ReadOnlyDB struct {
	memorydb.Database
}

func (t *ReadOnlyDB) Put(key []byte, value []byte) error {
	return errReadOnly
}

func TestAccountExtraDataLinker_whenError(t *testing.T) {
	db := NewDatabase(&ReadOnlyDB{})
	psr := common.Hash{1}
	pmr := common.Hash{2}

	linker := NewAccountExtraDataLinker(db)

	err := linker.Link(psr, pmr)

	if err == nil {
		t.Fatal("expecting a read only error to be returned")
	}

	if err != errReadOnly {
		t.Fatal("expecting the read only error to be returned")
	}
}

func TestAccountExtraDataLinker_whenFinding(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}
	pmr := common.Hash{2}

	err := db.Put(append(stateRootToExtraDataRootPrefix, psr[:]...), pmr[:])

	if err != nil {
		t.Fatal("unable to write to db")
	}

	pml := NewAccountExtraDataLinker(db)

	pmrRetrieved := pml.GetAccountExtraDataRoot(psr)

	if pmrRetrieved != pmr {
		t.Fatal("the mapping should have been retrieved")
	}
}

func TestAccountExtraDataLinker_whenNotFound(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}

	pml := NewAccountExtraDataLinker(db)

	pmrRetrieved := pml.GetAccountExtraDataRoot(psr)

	if !common.EmptyHash(pmrRetrieved) {
		t.Fatal("the retrieved privacy metadata root should be the empty hash")
	}
}

func TestPrivateStatesTrieRoot(t *testing.T) {
	db := NewMemoryDatabase()
	blockRoot := common.HexToHash("0x4c50c7d11e58e5c6f40fa1a630ffcb3a017453e7f9d0ec8ccb01033fcf9f2210")
	mtRoot := common.HexToHash("0x5c46375b6b333983077e152d1b6ca101d0586a6565fa75750deb1b07154bbdca")

	err := WritePrivateStatesTrieRoot(db, blockRoot, mtRoot)
	assert.Nil(t, err)

	retrievedRoot := GetPrivateStatesTrieRoot(db, blockRoot)
	assert.Equal(t, mtRoot, retrievedRoot)

	retrievedEmptyRoot := GetPrivateStatesTrieRoot(db, common.Hash{})
	assert.Equal(t, common.Hash{}, retrievedEmptyRoot)
}

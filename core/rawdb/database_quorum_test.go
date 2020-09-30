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

func TestPrivacyMedatadaLinkEmptyRoot(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}

	pml := NewPrivacyMetadataLinker(db)

	err := pml.LinkPrivacyMetadataRootToPrivateStateRoot(psr, emptyRoot)

	if err != nil {
		t.Fatal("unable to store the link")
	}

	value, _ := db.Get(append(privateRootToPrivacyMetadataRootPrefix, psr[:]...))

	if value != nil {
		t.Fatal("the mapping should not have been stored")
	}
}

func TestPrivacyMedatadaLinkRoot(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}
	pmr := common.Hash{2}

	pml := NewPrivacyMetadataLinker(db)

	err := pml.LinkPrivacyMetadataRootToPrivateStateRoot(psr, pmr)

	if err != nil {
		t.Fatal("unable to store the link")
	}

	value, _ := db.Get(append(privateRootToPrivacyMetadataRootPrefix, psr[:]...))

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

func TestPrivacyMedatadaLinkRootErrorWrapping(t *testing.T) {
	db := NewDatabase(&ReadOnlyDB{})
	psr := common.Hash{1}
	pmr := common.Hash{2}

	pml := NewPrivacyMetadataLinker(db)

	err := pml.LinkPrivacyMetadataRootToPrivateStateRoot(psr, pmr)

	if err == nil {
		t.Fatal("expecting a read only error to be returned")
	}

	if err != errReadOnly {
		t.Fatal("expecting the read only error to be returned")
	}
}

func TestPrivacyMedatadaRetrievePrivacyMetadataRoot(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}
	pmr := common.Hash{2}

	err := db.Put(append(privateRootToPrivacyMetadataRootPrefix, psr[:]...), pmr[:])

	if err != nil {
		t.Fatal("unable to write to db")
	}

	pml := NewPrivacyMetadataLinker(db)

	pmrRetrieved := pml.PrivacyMetadataRootForPrivateStateRoot(psr)

	if pmrRetrieved != pmr {
		t.Fatal("the mapping should have been retrieved")
	}
}

func TestPrivacyMedatadaRetrieveEmptyPrivacyMetadataRoot(t *testing.T) {
	db := NewMemoryDatabase()
	psr := common.Hash{1}

	pml := NewPrivacyMetadataLinker(db)

	pmrRetrieved := pml.PrivacyMetadataRootForPrivateStateRoot(psr)

	if !common.EmptyHash(pmrRetrieved) {
		t.Fatal("the retrieved privacy metadata root should be thg empty hash")
	}
}

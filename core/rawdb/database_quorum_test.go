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
	"testing"
)

// Tests that setting the flag for Quorum EIP155 activation read values correctly
func TestIsQuorumEIP155Active(t *testing.T) {
	db := NewMemoryDatabase()

	isQuorumEIP155Active := GetIsQuorumEIP155Activated(db)
	if isQuorumEIP155Active {
		t.Fatal("Quorum EIP155 active read to be set, but wasn't set beforehand")
	}

	dbSet := NewMemoryDatabase()
	WriteQuorumEIP155Activation(dbSet)

	isQuorumEIP155ActiveAfterSetting := GetIsQuorumEIP155Activated(dbSet)
	if !isQuorumEIP155ActiveAfterSetting {
		t.Fatal("Quorum EIP155 active read to be unset, but was set beforehand")
	}
}

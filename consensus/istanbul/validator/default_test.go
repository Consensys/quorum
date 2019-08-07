// Copyright 2017 The go-ethereum Authors
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

package validator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	testAddress  = "70524d664ffe731100208a0154e556f9bb679ae6"
	testAddress2 = "b37866a925bccd69cfa98d43b510f1d23d78a851"
	testAddress3 = "b37866a925bccd69cfa98d43b510f1d23d78a852"
	testAddress4 = "70524d664ffe731100208a0154e556f9bb679ae7"
)

func TestValidatorSet(t *testing.T) {
	testNewValidatorSet(t)
	testNormalValSet(t)
	testEmptyValSet(t)
	testStickyProposer(t)
	testAddAndRemoveValidator(t)
	testQuorumSize(t)
}

func testNewValidatorSet(t *testing.T) {
	var validators []istanbul.Validator
	const ValCnt = 100

	// Create 100 validators with random addresses
	b := []byte{}
	for i := 0; i < ValCnt; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		val := New(addr)
		validators = append(validators, val)
		b = append(b, val.Address().Bytes()...)
	}

	// Create ValidatorSet
	valSet := NewSet(ExtractValidators(b), istanbul.RoundRobin)
	if valSet == nil {
		t.Errorf("the validator byte array cannot be parsed")
		t.FailNow()
	}

	// Check validators sorting: should be in ascending order
	for i := 0; i < ValCnt-1; i++ {
		val := valSet.GetByIndex(uint64(i))
		nextVal := valSet.GetByIndex(uint64(i + 1))
		if strings.Compare(val.String(), nextVal.String()) >= 0 {
			t.Errorf("validator set is not sorted in ascending order")
		}
	}
}

func testNormalValSet(t *testing.T) {
	b1 := common.Hex2Bytes(testAddress)
	b2 := common.Hex2Bytes(testAddress2)
	addr1 := common.BytesToAddress(b1)
	addr2 := common.BytesToAddress(b2)
	val1 := New(addr1)
	val2 := New(addr2)

	valSet := newDefaultSet([]common.Address{addr1, addr2}, istanbul.RoundRobin)
	if valSet == nil {
		t.Errorf("the format of validator set is invalid")
		t.FailNow()
	}

	// check size
	if size := valSet.Size(); size != 2 {
		t.Errorf("the size of validator set is wrong: have %v, want 2", size)
	}
	// test get by index
	if val := valSet.GetByIndex(uint64(0)); !reflect.DeepEqual(val, val1) {
		t.Errorf("validator mismatch: have %v, want %v", val, val1)
	}
	// test get by invalid index
	if val := valSet.GetByIndex(uint64(2)); val != nil {
		t.Errorf("validator mismatch: have %v, want nil", val)
	}
	// test get by address
	if _, val := valSet.GetByAddress(addr2); !reflect.DeepEqual(val, val2) {
		t.Errorf("validator mismatch: have %v, want %v", val, val2)
	}
	// test get by invalid address
	invalidAddr := common.HexToAddress("0x9535b2e7faaba5288511d89341d94a38063a349b")
	if _, val := valSet.GetByAddress(invalidAddr); val != nil {
		t.Errorf("validator mismatch: have %v, want nil", val)
	}
	// test get proposer
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val1) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val1)
	}
	// test calculate proposer
	lastProposer := addr1
	valSet.CalcProposer(lastProposer, uint64(0))
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val2) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val2)
	}
	valSet.CalcProposer(lastProposer, uint64(3))
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val1) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val1)
	}
	// test empty last proposer
	lastProposer = common.Address{}
	valSet.CalcProposer(lastProposer, uint64(3))
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val2) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val2)
	}
}

func testEmptyValSet(t *testing.T) {
	valSet := NewSet(ExtractValidators([]byte{}), istanbul.RoundRobin)
	if valSet == nil {
		t.Errorf("validator set should not be nil")
	}
}

func testAddAndRemoveValidator(t *testing.T) {
	valSet := NewSet(ExtractValidators([]byte{}), istanbul.RoundRobin)
	if !valSet.AddValidator(common.StringToAddress(string(2))) {
		t.Error("the validator should be added")
	}
	if valSet.AddValidator(common.StringToAddress(string(2))) {
		t.Error("the existing validator should not be added")
	}
	valSet.AddValidator(common.StringToAddress(string(1)))
	valSet.AddValidator(common.StringToAddress(string(0)))
	if len(valSet.List()) != 3 {
		t.Error("the size of validator set should be 3")
	}

	for i, v := range valSet.List() {
		expected := common.StringToAddress(string(i))
		if v.Address() != expected {
			t.Errorf("the order of validators is wrong: have %v, want %v", v.Address().Hex(), expected.Hex())
		}
	}

	if !valSet.RemoveValidator(common.StringToAddress(string(2))) {
		t.Error("the validator should be removed")
	}
	if valSet.RemoveValidator(common.StringToAddress(string(2))) {
		t.Error("the non-existing validator should not be removed")
	}
	if len(valSet.List()) != 2 {
		t.Error("the size of validator set should be 2")
	}
	valSet.RemoveValidator(common.StringToAddress(string(1)))
	if len(valSet.List()) != 1 {
		t.Error("the size of validator set should be 1")
	}
	valSet.RemoveValidator(common.StringToAddress(string(0)))
	if len(valSet.List()) != 0 {
		t.Error("the size of validator set should be 0")
	}
}

func testStickyProposer(t *testing.T) {
	b1 := common.Hex2Bytes(testAddress)
	b2 := common.Hex2Bytes(testAddress2)
	addr1 := common.BytesToAddress(b1)
	addr2 := common.BytesToAddress(b2)
	val1 := New(addr1)
	val2 := New(addr2)

	valSet := newDefaultSet([]common.Address{addr1, addr2}, istanbul.Sticky)

	// test get proposer
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val1) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val1)
	}
	// test calculate proposer
	lastProposer := addr1
	valSet.CalcProposer(lastProposer, uint64(0))
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val1) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val1)
	}

	valSet.CalcProposer(lastProposer, uint64(1))
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val2) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val2)
	}
	// test empty last proposer
	lastProposer = common.Address{}
	valSet.CalcProposer(lastProposer, uint64(3))
	if val := valSet.GetProposer(); !reflect.DeepEqual(val, val2) {
		t.Errorf("proposer mismatch: have %v, want %v", val, val2)
	}
}

func testQuorumSize(t *testing.T) {
	b1 := common.Hex2Bytes(testAddress)
	b2 := common.Hex2Bytes(testAddress2)
	addr1 := common.BytesToAddress(b1)
	addr2 := common.BytesToAddress(b2)
	
	valSet := newDefaultSet([]common.Address{addr1, addr2}, istanbul.RoundRobin)
	// N==2
	// formulaType = 0, default 2f()+1 
	if valSet.Size() != 2 {
		t.Errorf("valSet.Size() expected: %v, got: %v", 2, valSet.Size())
	}
	if valSet.F() != 0 {
		t.Errorf("valSet.F() expected: %v, got: %v", 0, valSet.F())
	}
	if valSet.QuorumSize(0) != 1 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 0, 2, 1, valSet.QuorumSize(0))
	}
	// formulaType = 1, proposed update Ceil(2N/3)
	if valSet.QuorumSize(1) != 2 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 1, 2, 2, valSet.QuorumSize(1))
	}
	// formulaType = 2, proposed update N-f() 
	if valSet.QuorumSize(2) != 2 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 2, 2, 2, valSet.QuorumSize(2))
	}
	// N==3
	b3 := common.Hex2Bytes(testAddress3)
	addr3 := common.BytesToAddress(b3)
	valSet.AddValidator(addr3)
	if valSet.Size() != 3 {
		t.Errorf("valSet.Size() expected: %v, got: %v", 3, valSet.Size())
	}
	if valSet.F() != 0 {
		t.Errorf("valSet.F() expected: %v, got: %v", 0, valSet.F())
	}
	// formulaType = 0, default 2f()+1 
	if valSet.QuorumSize(0) != 1 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 0, 3, 1, valSet.QuorumSize(0))
	}
	// formulaType = 1, proposed update Ceil(2N/3)
	if valSet.QuorumSize(1) != 2 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 1, 3, 2, valSet.QuorumSize(1))
	}
	// formulaType = 2, proposed update N-f() 
	if valSet.QuorumSize(2) != 3 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 2, 3, 3, valSet.QuorumSize(2))
	}

	// N==4
	b4 := common.Hex2Bytes(testAddress4)
	addr4 := common.BytesToAddress(b4)
	valSet.AddValidator(addr4)
	if valSet.Size() != 4 {
		t.Errorf("valSet.Size() expected: %v, got: %v", 4, valSet.Size())
	}
	if valSet.F() != 1 {
		t.Errorf("valSet.F() expected: %v, got: %v", 1, valSet.F())
	}
	// formulaType = 0, default 2f()+1 
	if valSet.QuorumSize(0) != 3 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 0, 4, 3, valSet.QuorumSize(0))
	}
	// formulaType = 1, proposed update Ceil(2N/3)
	if valSet.QuorumSize(1) != 3 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 1, 4, 3, valSet.QuorumSize(1))
	}
	// formulaType = 2, proposed update N-f() 
	if valSet.QuorumSize(2) != 3 {
		t.Errorf("QuorumSize wrong for formulaType: %v,  N: %v, expected: %v, got: %v", 2, 4, 3, valSet.QuorumSize(2))
	}
}

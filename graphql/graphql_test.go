// Copyright 2019 The go-ethereum Authors
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

package graphql

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/private/engine"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/private"
)

func TestBuildSchema(t *testing.T) {
	// Make sure the schema can be parsed and matched up to the object model.
	if _, err := newHandler(nil); err != nil {
		t.Errorf("Could not construct GraphQL handler: %v", err)
	}
}

// Quorum
// Test Quorum specific GraphQL schema for private transaction
func TestQuorumSchema(t *testing.T) {
	saved := private.P
	defer func() {
		private.P = saved
	}()
	key := common.BytesToEncryptedPayloadHash([]byte("key"))
	private.P = &StubPrivateTransactionManager{
		responses: map[string][]interface{}{
			key.String(): {
				[]byte("private payload"), // equals to 0x70726976617465207061796c6f6164 after converting to bytes
				nil,
			},
		},
	}
	// Test private transaction
	privateTx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), key.Bytes())
	privateTx.SetPrivate()
	privateTxQuery := &Transaction{tx: privateTx}
	isPrivate, err := privateTxQuery.IsPrivate(context.Background())
	if err != nil {
		t.Fatalf("Expect no error: %v", err)
	}
	if !*isPrivate {
		t.Fatalf("Expect isPrivate to be true for private TX")
	}
	privateInputData, err := privateTxQuery.PrivateInputData(context.Background())
	if err != nil {
		t.Fatalf("Expect no error: %v", err)
	}
	if privateInputData.String() != "0x70726976617465207061796c6f6164" {
		t.Fatalf("Expect privateInputData to be: \"0x70726976617465207061796c6f6164\" for private TX, actual: %v", privateInputData.String())
	}
	// Test public transaction
	publicTx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), []byte("key"))
	publicTxQuery := &Transaction{tx: publicTx}
	isPrivate, err = publicTxQuery.IsPrivate(context.Background())
	if err != nil {
		t.Fatalf("Expect no error: %v", err)
	}
	if *isPrivate {
		t.Fatalf("Expect isPrivate to be false for public TX")
	}
	privateInputData, err = publicTxQuery.PrivateInputData(context.Background())
	if err != nil {
		t.Fatalf("Expect no error: %v", err)
	}
	if privateInputData.String() != "0x" {
		t.Fatalf("Expect privateInputData to be: \"0x\" for public TX, actual: %v", privateInputData.String())
	}
}

type StubPrivateTransactionManager struct {
	responses map[string][]interface{}
}

func (spm *StubPrivateTransactionManager) Name() string {
	return "to be implemented"
}

func (spm *StubPrivateTransactionManager) Features() engine.PTMFeatures {
	return engine.NewPTMFeatures(engine.PrivacyEnhancements)
}

func (spm *StubPrivateTransactionManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, fmt.Errorf("to be implemented")
}

func (spm *StubPrivateTransactionManager) StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, fmt.Errorf("to be implemented")
}

func (spm *StubPrivateTransactionManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) ([]byte, error) {
	return nil, fmt.Errorf("to be implemented")
}

func (spm *StubPrivateTransactionManager) Receive(data common.EncryptedPayloadHash) ([]byte, *engine.ExtraMetadata, error) {
	res := spm.responses[data.String()]
	if err, ok := res[1].(error); ok {
		return nil, nil, err
	}
	if ret, ok := res[0].([]byte); ok {
		return ret, &engine.ExtraMetadata{
			PrivacyFlag: engine.PrivacyFlagStandardPrivate,
		}, nil
	}
	return nil, nil, nil
}

func (spm *StubPrivateTransactionManager) ReceiveRaw(data common.EncryptedPayloadHash) ([]byte, *engine.ExtraMetadata, error) {
	return spm.Receive(data)
}

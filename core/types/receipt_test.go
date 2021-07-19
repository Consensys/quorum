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

package types

import (
	"bytes"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

func TestLegacyReceiptDecoding(t *testing.T) {
	tests := []struct {
		name   string
		encode func(*Receipt) ([]byte, error)
	}{
		{
			"StoredReceiptRLP",
			encodeAsStoredReceiptRLP,
		},
		{
			"V4StoredReceiptRLP",
			encodeAsV4StoredReceiptRLP,
		},
		{
			"V3StoredReceiptRLP",
			encodeAsV3StoredReceiptRLP,
		},
	}

	tx := NewTransaction(1, common.HexToAddress("0x1"), big.NewInt(1), 1, big.NewInt(1), nil)
	receipt := &Receipt{
		Status:            ReceiptStatusFailed,
		CumulativeGasUsed: 1,
		Logs: []*Log{
			{
				Address: common.BytesToAddress([]byte{0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
			},
			{
				Address: common.BytesToAddress([]byte{0x01, 0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
			},
		},
		TxHash:          tx.Hash(),
		ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
		GasUsed:         111111,
	}
	receipt.Bloom = CreateBloom(Receipts{receipt})

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			enc, err := tc.encode(receipt)
			if err != nil {
				t.Fatalf("Error encoding receipt: %v", err)
			}
			var dec ReceiptForStorage
			if err := rlp.DecodeBytes(enc, &dec); err != nil {
				t.Fatalf("Error decoding RLP receipt: %v", err)
			}
			// Check whether all consensus fields are correct.
			testConsensusFields(t, dec, receipt)
		})
	}
}

func testConsensusFields(t *testing.T, dec ReceiptForStorage, receipt *Receipt) {
	if dec.Status != receipt.Status {
		t.Fatalf("Receipt status mismatch, want %v, have %v", receipt.Status, dec.Status)
	}
	if dec.CumulativeGasUsed != receipt.CumulativeGasUsed {
		t.Fatalf("Receipt CumulativeGasUsed mismatch, want %v, have %v", receipt.CumulativeGasUsed, dec.CumulativeGasUsed)
	}
	if dec.Bloom != receipt.Bloom {
		t.Fatalf("Bloom data mismatch, want %v, have %v", receipt.Bloom, dec.Bloom)
	}
	if len(dec.Logs) != len(receipt.Logs) {
		t.Fatalf("Receipt log number mismatch, want %v, have %v", len(receipt.Logs), len(dec.Logs))
	}
	for i := 0; i < len(dec.Logs); i++ {
		if dec.Logs[i].Address != receipt.Logs[i].Address {
			t.Fatalf("Receipt log %d address mismatch, want %v, have %v", i, receipt.Logs[i].Address, dec.Logs[i].Address)
		}
		if !reflect.DeepEqual(dec.Logs[i].Topics, receipt.Logs[i].Topics) {
			t.Fatalf("Receipt log %d topics mismatch, want %v, have %v", i, receipt.Logs[i].Topics, dec.Logs[i].Topics)
		}
		if !bytes.Equal(dec.Logs[i].Data, receipt.Logs[i].Data) {
			t.Fatalf("Receipt log %d data mismatch, want %v, have %v", i, receipt.Logs[i].Data, dec.Logs[i].Data)
		}
	}

	if !bytes.Equal(dec.RevertReason, receipt.RevertReason) {
		t.Fatalf("RevertReason data mismatch, want %v, have %v", receipt.RevertReason, dec.RevertReason)
	}
}

func encodeAsStoredReceiptRLP(want *Receipt) ([]byte, error) {
	receiptForStorage := (*ReceiptForStorage)(want)
	return rlp.EncodeToBytes(receiptForStorage)
}

func encodeAsV4StoredReceiptRLP(want *Receipt) ([]byte, error) {
	stored := &v4StoredReceiptRLP{
		PostStateOrStatus: want.statusEncoding(),
		CumulativeGasUsed: want.CumulativeGasUsed,
		TxHash:            want.TxHash,
		ContractAddress:   want.ContractAddress,
		Logs:              make([]*LogForStorage, len(want.Logs)),
		GasUsed:           want.GasUsed,
	}
	for i, log := range want.Logs {
		stored.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.EncodeToBytes(stored)
}

func encodeAsV3StoredReceiptRLP(want *Receipt) ([]byte, error) {
	stored := &v3StoredReceiptRLP{
		PostStateOrStatus: want.statusEncoding(),
		CumulativeGasUsed: want.CumulativeGasUsed,
		Bloom:             want.Bloom,
		TxHash:            want.TxHash,
		ContractAddress:   want.ContractAddress,
		Logs:              make([]*LogForStorage, len(want.Logs)),
		GasUsed:           want.GasUsed,
	}
	for i, log := range want.Logs {
		stored.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.EncodeToBytes(stored)
}

// Tests that receipt data can be correctly derived from the contextual infos
func TestDeriveFields(t *testing.T) {
	// Create a few transactions to have receipts for
	txs := Transactions{
		NewContractCreation(1, big.NewInt(1), 1, big.NewInt(1), nil),
		NewTransaction(2, common.HexToAddress("0x2"), big.NewInt(2), 2, big.NewInt(2), nil),
	}
	// Create the corresponding receipts
	receipts := Receipts{
		&Receipt{
			Status:            ReceiptStatusFailed,
			CumulativeGasUsed: 1,
			Logs: []*Log{
				{Address: common.BytesToAddress([]byte{0x11})},
				{Address: common.BytesToAddress([]byte{0x01, 0x11})},
			},
			TxHash:          txs[0].Hash(),
			ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
			GasUsed:         1,
		},
		&Receipt{
			PostState:         common.Hash{2}.Bytes(),
			CumulativeGasUsed: 3,
			Logs: []*Log{
				{Address: common.BytesToAddress([]byte{0x22})},
				{Address: common.BytesToAddress([]byte{0x02, 0x22})},
			},
			TxHash:          txs[1].Hash(),
			ContractAddress: common.BytesToAddress([]byte{0x02, 0x22, 0x22}),
			GasUsed:         2,
		},
	}
	// Clear all the computed fields and re-derive them
	number := big.NewInt(1)
	hash := common.BytesToHash([]byte{0x03, 0x14})

	clearComputedFieldsOnReceipts(t, receipts)
	if err := receipts.DeriveFields(params.TestChainConfig, hash, number.Uint64(), txs); err != nil {
		t.Fatalf("DeriveFields(...) = %v, want <nil>", err)
	}
	// Iterate over all the computed fields and check that they're correct
	signer := MakeSigner(params.TestChainConfig, number)

	for i := range receipts {
		testReceiptFields(t, receipts[i], txs, i, "receipt"+strconv.Itoa(i), hash, number, signer)
	}
}

// Tests that receipt data can be correctly derived from the contextual infos
// Tests public, private, and private mps txs/receipts
func TestDeriveFieldsMPS(t *testing.T) {
	// Create a public tx, private tx, psi tx
	pubT := NewContractCreation(1, big.NewInt(1), 1, big.NewInt(1), nil)
	privT := NewContractCreation(2, big.NewInt(2), 2, big.NewInt(2), nil)
	privT.SetPrivate()
	psiT := NewTransaction(3, common.HexToAddress("0x3"), big.NewInt(3), 3, big.NewInt(3), nil)
	psiT.SetPrivate()
	//3 transactions: public, private, and private with mps
	txs := Transactions{
		pubT,
		privT,
		psiT,
	}
	publicReceipt := &Receipt{
		Status:            ReceiptStatusFailed,
		CumulativeGasUsed: 1,
		Logs: []*Log{
			{Address: common.BytesToAddress([]byte{0x11})},
			{Address: common.BytesToAddress([]byte{0x01, 0x11})},
		},
		TxHash:          txs[0].Hash(),
		ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
		GasUsed:         1,
	}
	innerPrivateReceipt := &Receipt{
		PostState:         common.Hash{2}.Bytes(),
		CumulativeGasUsed: 3,
		Logs: []*Log{
			{Address: common.BytesToAddress([]byte{0x22})},
			{Address: common.BytesToAddress([]byte{0x02, 0x22})},
		},
		TxHash:          txs[1].Hash(),
		ContractAddress: common.BytesToAddress([]byte{0x02, 0x22, 0x22}),
		GasUsed:         2,
	}
	innerPSIReceipt := &Receipt{
		PostState:         common.Hash{3}.Bytes(),
		CumulativeGasUsed: 6,
		Logs: []*Log{
			{Address: common.BytesToAddress([]byte{0x33})},
			{Address: common.BytesToAddress([]byte{0x03, 0x33})},
		},
		TxHash:          txs[2].Hash(),
		ContractAddress: common.BytesToAddress([]byte{0x03, 0x33, 0x33}),
		GasUsed:         1,
	}
	psiReceipt := innerPSIReceipt
	psiReceipt.PSReceipts = make(map[PrivateStateIdentifier]*Receipt)
	psiReceipt.PSReceipts[PrivateStateIdentifier("psi1")] = innerPSIReceipt
	psiReceipt.PSReceipts[EmptyPrivateStateIdentifier] = innerPSIReceipt

	privateReceipt := innerPrivateReceipt
	privateReceipt.PSReceipts = make(map[PrivateStateIdentifier]*Receipt)
	privateReceipt.PSReceipts[EmptyPrivateStateIdentifier] = innerPrivateReceipt
	// Create the corresponding receipts: public, private, psi
	receipts := Receipts{
		publicReceipt,
		privateReceipt,
		psiReceipt,
	}
	// Clear all the computed fields and re-derive them
	number := big.NewInt(1)
	hash := common.BytesToHash([]byte{0x03, 0x14})

	clearComputedFieldsOnReceipts(t, receipts)
	if err := receipts.DeriveFields(params.QuorumMPSTestChainConfig, hash, number.Uint64(), txs); err != nil {
		t.Fatalf("DeriveFields(...) = %v, want <nil>", err)
	}
	// Iterate over all the computed fields and check that they're correct
	signer := MakeSigner(params.QuorumMPSTestChainConfig, number)

	for i := range receipts {
		testReceiptFields(t, receipts[i], txs, i, "receipt"+strconv.Itoa(i), hash, number, signer)
	}
	//check psi info on public and private receipt
	assert.Empty(t, receipts[0].PSReceipts)
	privRec := receipts[1]
	assert.Equal(t, 1, len(privRec.PSReceipts))
	assert.Contains(t, privRec.PSReceipts, EmptyPrivateStateIdentifier)
	for _, pR := range privRec.PSReceipts {
		testReceiptFields(t, pR, txs, 1, "privateReceipt", hash, number, signer)
	}

	//check psi info on private mps receipt
	psiRec := receipts[2]
	assert.Equal(t, 2, len(psiRec.PSReceipts))
	assert.Contains(t, psiRec.PSReceipts, EmptyPrivateStateIdentifier)
	assert.Contains(t, psiRec.PSReceipts, PrivateStateIdentifier("psi1"))
	for _, pR := range psiRec.PSReceipts {
		testReceiptFields(t, pR, txs, 2, "psiReceipt", hash, number, signer)
	}
}

func testReceiptFields(t *testing.T, receipt *Receipt, txs Transactions, txIndex int, receiptName string, hash common.Hash, number *big.Int, signer Signer) {
	if receipt.TxHash != txs[txIndex].Hash() {
		t.Errorf("%s.TxHash = %s, want %s", receiptName, receipt.TxHash.String(), txs[1].Hash().String())
	}
	if receipt.BlockHash != hash {
		t.Errorf("%s.BlockHash = %s, want %s", receiptName, receipt.BlockHash.String(), hash.String())
	}
	if receipt.BlockNumber.Cmp(number) != 0 {
		t.Errorf("%s.BlockNumber = %s, want %s", receiptName, receipt.BlockNumber.String(), number.String())
	}
	if receipt.TransactionIndex != uint(txIndex) {
		t.Errorf("%s.TransactionIndex = %d, want %d", receiptName, receipt.TransactionIndex, txIndex)
	}
	if receipt.GasUsed != txs[txIndex].Gas() {
		t.Errorf("%s.GasUsed = %d, want %d", receiptName, receipt.GasUsed, txs[txIndex].Gas())
	}
	if txs[txIndex].To() != nil && receipt.ContractAddress != (common.Address{}) {
		t.Errorf("%s.ContractAddress = %s, want %s", receiptName, receipt.ContractAddress.String(), (common.Address{}).String())
	}
	from, _ := Sender(signer, txs[txIndex])
	contractAddress := crypto.CreateAddress(from, txs[txIndex].Nonce())
	if txs[txIndex].To() == nil && receipt.ContractAddress != contractAddress {
		t.Errorf("%s.ContractAddress = %s, want %s", receiptName, receipt.ContractAddress.String(), contractAddress.String())
	}
	for j := range receipt.Logs {
		if receipt.Logs[j].BlockNumber != number.Uint64() {
			t.Errorf("%s.Logs[%d].BlockNumber = %d, want %d", receiptName, j, receipt.Logs[j].BlockNumber, number.Uint64())
		}
		if receipt.Logs[j].BlockHash != hash {
			t.Errorf("%s.Logs[%d].BlockHash = %s, want %s", receiptName, j, receipt.Logs[j].BlockHash.String(), hash.String())
		}
		if receipt.Logs[j].TxHash != txs[txIndex].Hash() {
			t.Errorf("%s.Logs[%d].TxHash = %s, want %s", receiptName, j, receipt.Logs[j].TxHash.String(), txs[txIndex].Hash().String())
		}
		if receipt.Logs[j].TxHash != txs[txIndex].Hash() {
			t.Errorf("%s.Logs[%d].TxHash = %s, want %s", receiptName, j, receipt.Logs[j].TxHash.String(), txs[txIndex].Hash().String())
		}
		if receipt.Logs[j].TxIndex != uint(txIndex) {
			t.Errorf("%s.Logs[%d].TransactionIndex = %d, want %d", receiptName, j, receipt.Logs[j].TxIndex, txIndex)
		}
	}
}

func clearComputedFieldsOnReceipts(t *testing.T, receipts Receipts) {
	t.Helper()

	for _, receipt := range receipts {
		clearComputedFieldsOnReceipt(t, receipt)
	}
}

func clearComputedFieldsOnReceipt(t *testing.T, receipt *Receipt) {
	t.Helper()

	receipt.TxHash = common.Hash{}
	receipt.BlockHash = common.Hash{}
	receipt.BlockNumber = big.NewInt(math.MaxUint32)
	receipt.TransactionIndex = math.MaxUint32
	receipt.ContractAddress = common.Address{}
	receipt.GasUsed = 0

	clearComputedFieldsOnLogs(t, receipt.Logs)
}

func clearComputedFieldsOnLogs(t *testing.T, logs []*Log) {
	t.Helper()

	for _, log := range logs {
		clearComputedFieldsOnLog(t, log)
	}
}

func clearComputedFieldsOnLog(t *testing.T, log *Log) {
	t.Helper()

	log.BlockNumber = math.MaxUint32
	log.BlockHash = common.Hash{}
	log.TxHash = common.Hash{}
	log.TxIndex = math.MaxUint32
	log.Index = math.MaxUint32
}

func TestQuorumReceiptExtraDataDecodingSuccess(t *testing.T) {
	tx := NewTransaction(1, common.HexToAddress("0x1"), big.NewInt(1), 1, big.NewInt(1), nil)
	receipt := &Receipt{
		Status:            ReceiptStatusFailed,
		CumulativeGasUsed: 1,
		Logs: []*Log{
			{
				Address: common.BytesToAddress([]byte{0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
			},
			{
				Address: common.BytesToAddress([]byte{0x01, 0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
			},
		},
		TxHash:          tx.Hash(),
		ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
		GasUsed:         111111,
	}
	receipt.Bloom = CreateBloom(Receipts{receipt})

	extraData := &QuorumReceiptExtraData{
		RevertReason: []byte("arbitrary reason"),
		PSReceipts:   map[PrivateStateIdentifier]*Receipt{PrivateStateIdentifier("psi1"): receipt},
	}
	rlpData, err := rlp.EncodeToBytes(extraData)
	assert.Nil(t, err)
	var decodedExtraData QuorumReceiptExtraData
	err = rlp.DecodeBytes(rlpData, &decodedExtraData)
	assert.Nil(t, err)
	assert.Equal(t, decodedExtraData.RevertReason, []byte("arbitrary reason"))
	assert.Contains(t, decodedExtraData.PSReceipts, PrivateStateIdentifier("psi1"))
	decodedReceipt := decodedExtraData.PSReceipts[PrivateStateIdentifier("psi1")]
	assert.NotNil(t, decodedReceipt)
	testConsensusFields(t, ReceiptForStorage(*decodedReceipt), receipt)
}

func TestQuorumReceiptExtraDataDecodingFailDueToUnknownVersion(t *testing.T) {
	rlpData, err := rlp.EncodeToBytes(&storedQuorumReceiptExtraDataV1RLP{
		Version:      2,
		RevertReason: []byte("arbitrary reason"),
	})
	assert.Nil(t, err)
	var decodedExtraData QuorumReceiptExtraData
	err = rlp.DecodeBytes(rlpData, &decodedExtraData)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "unknown version 2")
}

func TestQuorumReceiptExtraDataDecodingFailDueToGarbageData(t *testing.T) {
	var decodedExtraData QuorumReceiptExtraData
	err := rlp.DecodeBytes([]byte("arbitrary data"), &decodedExtraData)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "unexpected content type (expecting list) 0")
}

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

func TestDecodeEmptyTypedReceipt(t *testing.T) {
	input := []byte{0x80}
	var r Receipt
	err := rlp.DecodeBytes(input, &r)
	if err != errEmptyTypedReceipt {
		t.Fatal("wrong error:", err)
	}
}

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
			testConsensusFields(t, dec, receipt, false, false, false)
		})
	}
}

func testConsensusFields(t *testing.T, decReceipt ReceiptForStorage, encReceipt *Receipt, expectRevertReason bool, expectPSReceipts bool, supportPTM bool) {
	if encReceipt.Status != decReceipt.Status {
		t.Errorf("Status mismatch, want %v, have %v", encReceipt.Status, decReceipt.Status)
	}
	if encReceipt.CumulativeGasUsed != decReceipt.CumulativeGasUsed {
		t.Errorf("CumulativeGasUsed mismatch, want %v, have %v", encReceipt.CumulativeGasUsed, decReceipt.CumulativeGasUsed)
	}
	if encReceipt.Bloom != decReceipt.Bloom {
		t.Errorf("Bloom mismatch, want %v, have %v", encReceipt.Bloom, decReceipt.Bloom)
	}

	if !bytes.Equal(encReceipt.PostState, decReceipt.PostState) {
		t.Errorf("PostState mismatch, want %v, have %v", encReceipt.PostState, decReceipt.PostState)
	}
	compareLogsConsensusFields(t, encReceipt.Logs, decReceipt.Logs)

	if expectRevertReason {
		if !bytes.Equal(encReceipt.RevertReason, decReceipt.RevertReason) {
			t.Errorf("RevertReason mismatch, want %v, have %v", encReceipt.RevertReason, decReceipt.RevertReason)
		}
	} else if decReceipt.RevertReason != nil {
		t.Errorf("RevertReason mismatch, expecting nil, have %v", decReceipt.RevertReason)
	}

	if expectPSReceipts {
		comparePSReceipts(t, encReceipt.PSReceipts, decReceipt.PSReceipts, supportPTM)
	} else if decReceipt.PSReceipts != nil {
		t.Errorf("PSReceipts mismatch, expecting nil, have %v", decReceipt.PSReceipts)
	}
}

func compareLogsConsensusFields(t *testing.T, encLogs []*Log, decLogs []*Log) {
	if len(encLogs) != len(encLogs) {
		t.Fatalf("Logs[] length mismatch, want %v, have %v", len(encLogs), len(encLogs))
	}

	for i := 0; i < len(encLogs); i++ {
		if encLogs[i].Address != decLogs[i].Address {
			t.Errorf("Logs[%d].Address mismatch, want %v, have %v", i, encLogs[i].Address, decLogs[i].Address)
		}
		if !reflect.DeepEqual(encLogs[i].Topics, decLogs[i].Topics) {
			t.Errorf("Logs[%d].Topics mismatch, want %v, have %v", i, encLogs[i].Topics, decLogs[i].Topics)
		}
		if !bytes.Equal(encLogs[i].Data, decLogs[i].Data) {
			t.Errorf("Logs[%d].Data mismatch, want %v, have %v", i, encLogs[i].Data, decLogs[i].Data)
		}
	}
}

func comparePSReceipts(t *testing.T, encPSReceipts map[PrivateStateIdentifier]*Receipt, decPSReceipts map[PrivateStateIdentifier]*Receipt, supportPTM bool) {
	if encPSReceipts == nil {
		t.Fatalf("Receipt is missing the expected psReceipts[]")
	}
	if len(encPSReceipts) != len(decPSReceipts) {
		t.Fatalf("Receipt psi number mismatch, want %v, have %v", len(encPSReceipts), len(decPSReceipts))
	}

	for psi, decPsiReceipt := range decPSReceipts {
		wantedPsiReceipt := encPSReceipts[psi]
		if decPsiReceipt.Status != wantedPsiReceipt.Status {
			t.Errorf("PSReceipts[%s].Status mismatch, want %v, have %v", psi.String(), wantedPsiReceipt.Status, decPsiReceipt.Status)
		}
		if decPsiReceipt.CumulativeGasUsed != wantedPsiReceipt.CumulativeGasUsed {
			t.Errorf("PSReceipts[%s].CumulativeGasUsed mismatch, want %v, have %v", psi.String(), wantedPsiReceipt.CumulativeGasUsed, decPsiReceipt.CumulativeGasUsed)
		}
		if decPsiReceipt.Bloom != wantedPsiReceipt.Bloom {
			t.Errorf("PSReceipts[%s].Bloom mismatch, want %v, have %v", psi.String(), wantedPsiReceipt.Bloom, decPsiReceipt.Bloom)
		}
		if len(decPsiReceipt.Logs) != len(wantedPsiReceipt.Logs) {
			t.Errorf("PSReceipts[%s].Logs mismatch, want %v, have %v", psi.String(), wantedPsiReceipt.Logs, decPsiReceipt.Logs)
		}
		if supportPTM {
			// TxHash & ContractAddress are only encoded/decoded if PTM support is enabled
			if decPsiReceipt.TxHash != wantedPsiReceipt.TxHash {
				t.Errorf("PSReceipts[%s].TxHash mismatch, want %v, have %v", psi.String(), wantedPsiReceipt.TxHash, decPsiReceipt.TxHash)
			}
			if decPsiReceipt.ContractAddress != wantedPsiReceipt.ContractAddress {
				t.Errorf("PSReceipts[%s].ContractAddress mismatch, want %v, have %v", psi.String(), wantedPsiReceipt.ContractAddress, decPsiReceipt.ContractAddress)
			}
		}

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

func TestReceiptForStorage_OnlySubsetOfFieldsPreservedDuringSerialisation(t *testing.T) {

	fullReceipt := newFullReceipt(false, false, false)

	buf := new(bytes.Buffer)
	if err := rlp.Encode(buf, fullReceipt); err != nil {
		t.Fatalf("Error RLP encoding receipt: %v", err)
	}

	got := new(ReceiptForStorage)
	if err := rlp.Decode(buf, got); err != nil {
		t.Fatalf("Error RLP encoding receipt: %v", err)
	}

	// only a subset of fields are to be encoded, the rest are derived after decoding
	want := &ReceiptForStorage{
		Status:            ReceiptStatusSuccessful,
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
		Bloom: fullReceipt.Bloom,
	}

	assert.Equal(t, want, got)
}

var (
	stubHash = common.HexToHash("0xabcdef")
)

// newFullReceipt returns a new receipt with non-zero values in all fields to test field preservation during encode/decode
func newFullReceipt(withPSReceipts, withTopLevelReceiptRevertReason, withPSReceiptRevertReason bool) *ReceiptForStorage {
	fullReceipt := &ReceiptForStorage{
		Status:            ReceiptStatusSuccessful,
		CumulativeGasUsed: 1,
		Logs: []*Log{
			{
				Address: common.BytesToAddress([]byte{0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
				// log fields that shouldn't be encoded/decoded
				BlockNumber: uint64(5),
				TxHash:      stubHash,
				TxIndex:     uint(3),
				BlockHash:   stubHash,
				Index:       uint(54),
				Removed:     true,
			},
			{
				Address: common.BytesToAddress([]byte{0x01, 0x11}),
				Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
				Data:    []byte{0x01, 0x00, 0xff},
				// log fields that shouldn't be encoded/decoded
				BlockNumber: uint64(5),
				TxHash:      stubHash,
				TxIndex:     uint(3),
				BlockHash:   stubHash,
				Index:       uint(54),
				Removed:     true,
			},
		},
		// receipt fields that shouldn't be encoded/decoded
		TxHash:           stubHash,
		ContractAddress:  common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
		GasUsed:          111111,
		BlockHash:        stubHash,
		BlockNumber:      big.NewInt(14),
		TransactionIndex: uint(4),
	}
	topLevelReceiptBloom := CreateBloom(Receipts{(*Receipt)(fullReceipt)})
	fullReceipt.Bloom = topLevelReceiptBloom

	if withTopLevelReceiptRevertReason {
		fullReceipt.RevertReason = []byte{0x01, 0x00, 0xff}
	}

	if withPSReceipts {
		fullReceipt.PSReceipts = map[PrivateStateIdentifier]*Receipt{
			"myPSI": {
				Status:            ReceiptStatusSuccessful,
				CumulativeGasUsed: 1,
				Logs: []*Log{
					{
						Address: common.BytesToAddress([]byte{0x11}),
						Topics:  []common.Hash{common.HexToHash("dead"), common.HexToHash("beef")},
						Data:    []byte{0x01, 0x00, 0xff},
						// log fields that shouldn't be encoded/decoded
						BlockNumber: uint64(5),
						TxHash:      stubHash,
						TxIndex:     uint(3),
						BlockHash:   stubHash,
						Index:       uint(54),
						Removed:     true,
					},
				},
				// txhash and contractAddress should be encoded/decoded for PSReceipts only
				TxHash:          stubHash,
				ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
				// psReceipt fields that shouldn't be encoded/decoded
				GasUsed:          111111,
				BlockHash:        stubHash,
				BlockNumber:      big.NewInt(14),
				TransactionIndex: uint(4),
			},
		}
		psReceiptBloom := CreateBloom(Receipts{(fullReceipt.PSReceipts["myPSI"])})
		fullReceipt.PSReceipts["myPSI"].Bloom = psReceiptBloom

		if withPSReceiptRevertReason {
			fullReceipt.PSReceipts["myPSI"].RevertReason = []byte{0x01, 0x00, 0xff}
		}
	}
	return fullReceipt
}

// Tests that receipt data can be correctly derived from the contextual infos
func TestDeriveFields(t *testing.T) {
	// Create a few transactions to have receipts for
	to2 := common.HexToAddress("0x2")
	to3 := common.HexToAddress("0x3")
	txs := Transactions{
		NewTx(&LegacyTx{
			Nonce:    1,
			Value:    big.NewInt(1),
			Gas:      1,
			GasPrice: big.NewInt(1),
		}),
		NewTx(&LegacyTx{
			To:       &to2,
			Nonce:    2,
			Value:    big.NewInt(2),
			Gas:      2,
			GasPrice: big.NewInt(2),
		}),
		NewTx(&AccessListTx{
			To:       &to3,
			Nonce:    3,
			Value:    big.NewInt(3),
			Gas:      3,
			GasPrice: big.NewInt(3),
		}),
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
		&Receipt{
			Type:              AccessListTxType,
			PostState:         common.Hash{3}.Bytes(),
			CumulativeGasUsed: 6,
			Logs: []*Log{
				{Address: common.BytesToAddress([]byte{0x33})},
				{Address: common.BytesToAddress([]byte{0x03, 0x33})},
			},
			TxHash:          txs[2].Hash(),
			ContractAddress: common.BytesToAddress([]byte{0x03, 0x33, 0x33}),
			GasUsed:         3,
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

func testReceiptFields(t *testing.T, receipt *Receipt, txs Transactions, txIndex int, receiptName string, blockHash common.Hash, blockNumber *big.Int, signer Signer) {
	if receipt.BlockNumber.Cmp(blockNumber) != 0 {
		t.Errorf("%s.BlockNumber = %s, want %s", receiptName, receipt.BlockNumber.String(), blockNumber.String())
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
		if receipt.Logs[j].BlockNumber != blockNumber.Uint64() {
			t.Errorf("%s.Logs[%d].BlockNumber = %d, want %d", receiptName, j, receipt.Logs[j].BlockNumber, blockNumber.Uint64())
		}
		if receipt.Logs[j].BlockHash != blockHash {
			t.Errorf("%s.Logs[%d].BlockHash = %s, want %s", receiptName, j, receipt.Logs[j].BlockHash.String(), blockHash.String())
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
		// TODO: @achraf have a look how to include this
		//if receipts[i].Logs[j].Index != logIndex {
		//	t.Errorf("receipts[%d].Logs[%d].Index = %d, want %d", i, j, receipts[i].Logs[j].Index, logIndex)
		//}
		//logIndex++
	}
}

// TestTypedReceiptEncodingDecoding reproduces a flaw that existed in the receipt
// rlp decoder, which failed due to a shadowing error.
func TestTypedReceiptEncodingDecoding(t *testing.T) {
	var payload = common.FromHex("f9043eb9010c01f90108018262d4b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0b9010c01f901080182cd14b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0b9010d01f901090183013754b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0b9010d01f90109018301a194b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0")
	check := func(bundle []*Receipt) {
		t.Helper()
		for i, receipt := range bundle {
			if got, want := receipt.Type, uint8(1); got != want {
				t.Fatalf("bundle %d: got %x, want %x", i, got, want)
			}
		}
	}
	{
		var bundle []*Receipt
		rlp.DecodeBytes(payload, &bundle)
		check(bundle)
	}
	{
		var bundle []*Receipt
		r := bytes.NewReader(payload)
		s := rlp.NewStream(r, uint64(len(payload)))
		if err := s.Decode(&bundle); err != nil {
			t.Fatal(err)
		}
		check(bundle)
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
	ps1Receipt := &Receipt{
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
	ps1Receipt.Bloom = CreateBloom(Receipts{ps1Receipt})

	extraData := &QuorumReceiptExtraData{
		RevertReason: []byte("arbitrary reason"),
		PSReceipts:   map[PrivateStateIdentifier]*Receipt{PrivateStateIdentifier("psi1"): ps1Receipt},
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
	comparePSReceipts(t, decodedExtraData.PSReceipts, extraData.PSReceipts, true)
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

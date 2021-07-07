// Copyright 2014 The go-ethereum Authors
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
	"errors"
	"fmt"
	"io"
	"math/big"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
)

//go:generate gencodec -type Receipt -field-override receiptMarshaling -out gen_receipt_json.go

var (
	receiptStatusFailedRLP     = []byte{}
	receiptStatusSuccessfulRLP = []byte{0x01}
)

const (
	// ReceiptStatusFailed is the status code of a transaction if execution failed.
	ReceiptStatusFailed = uint64(0)

	// ReceiptStatusSuccessful is the status code of a transaction if execution succeeded.
	ReceiptStatusSuccessful = uint64(1)
)

// Receipt represents the results of a transaction.
type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	PostState         []byte `json:"root"`
	Status            uint64 `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" gencodec:"required"`
	Bloom             Bloom  `json:"logsBloom"         gencodec:"required"`
	Logs              []*Log `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash          common.Hash    `json:"transactionHash" gencodec:"required"`
	ContractAddress common.Address `json:"contractAddress"`
	GasUsed         uint64         `json:"gasUsed" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex uint        `json:"transactionIndex"`

	// Quorum
	//
	// This is to support execution of a private transaction on multiple private states,
	// in which receipts are produced per PSI. It is also used by privacy marker transactions.
	// PSReceipts will hold a receipt for each PSI that is managed by this node.
	// Note that for MPS, the parent receipt will be an auxiliary receipt, whereas for PMT the parent
	// will be the privacy marker transaction receipt.
	//
	// This nested structure would not have more than 2 levels.
	PSReceipts map[PrivateStateIdentifier]*Receipt `json:"-"`
	// support saving the revert reason into the receipt itself for later consultation.
	RevertReason []byte `json:"revertReason,omitempty"`
	// End Quorum
}

type receiptMarshaling struct {
	PostState         hexutil.Bytes
	Status            hexutil.Uint64
	CumulativeGasUsed hexutil.Uint64
	GasUsed           hexutil.Uint64
	BlockNumber       *hexutil.Big
	TransactionIndex  hexutil.Uint
}

// receiptRLP is the consensus encoding of a receipt.
type receiptRLP struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Bloom             Bloom
	Logs              []*Log
}

// storedReceiptRLP is the storage encoding of a receipt.
type storedReceiptRLP struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Logs              []*LogForStorage
}

// v4StoredReceiptRLP is the storage encoding of a receipt used in database version 4.
type v4StoredReceiptRLP struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	TxHash            common.Hash
	ContractAddress   common.Address
	Logs              []*LogForStorage
	GasUsed           uint64
}

// v3StoredReceiptRLP is the original storage encoding of a receipt including some unnecessary fields.
type v3StoredReceiptRLP struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Bloom             Bloom
	TxHash            common.Hash
	ContractAddress   common.Address
	Logs              []*LogForStorage
	GasUsed           uint64
}

// NewReceipt creates a barebone transaction receipt, copying the init fields.
func NewReceipt(root []byte, failed bool, cumulativeGasUsed uint64) *Receipt {
	r := &Receipt{PostState: common.CopyBytes(root), CumulativeGasUsed: cumulativeGasUsed}
	if failed {
		r.Status = ReceiptStatusFailed
	} else {
		r.Status = ReceiptStatusSuccessful
	}
	return r
}

// EncodeRLP implements rlp.Encoder, and flattens the consensus fields of a receipt
// into an RLP stream. If no post state is present, byzantium fork is assumed.
func (r *Receipt) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &receiptRLP{r.statusEncoding(), r.CumulativeGasUsed, r.Bloom, r.Logs})
}

// DecodeRLP implements rlp.Decoder, and loads the consensus fields of a receipt
// from an RLP stream.
func (r *Receipt) DecodeRLP(s *rlp.Stream) error {
	var dec receiptRLP
	if err := s.Decode(&dec); err != nil {
		return err
	}
	if err := r.setStatus(dec.PostStateOrStatus); err != nil {
		return err
	}
	r.CumulativeGasUsed, r.Bloom, r.Logs = dec.CumulativeGasUsed, dec.Bloom, dec.Logs
	return nil
}

func (r *Receipt) setStatus(postStateOrStatus []byte) error {
	switch {
	case bytes.Equal(postStateOrStatus, receiptStatusSuccessfulRLP):
		r.Status = ReceiptStatusSuccessful
	case bytes.Equal(postStateOrStatus, receiptStatusFailedRLP):
		r.Status = ReceiptStatusFailed
	case len(postStateOrStatus) == len(common.Hash{}):
		r.PostState = postStateOrStatus
	default:
		return fmt.Errorf("invalid receipt status %x", postStateOrStatus)
	}
	return nil
}

func (r *Receipt) statusEncoding() []byte {
	if len(r.PostState) == 0 {
		if r.Status == ReceiptStatusFailed {
			return receiptStatusFailedRLP
		}
		return receiptStatusSuccessfulRLP
	}
	return r.PostState
}

// Size returns the approximate memory used by all internal contents. It is used
// to approximate and limit the memory consumption of various caches.
func (r *Receipt) Size() common.StorageSize {
	size := common.StorageSize(unsafe.Sizeof(*r)) + common.StorageSize(len(r.PostState))

	size += common.StorageSize(len(r.Logs)) * common.StorageSize(unsafe.Sizeof(Log{}))
	for _, log := range r.Logs {
		size += common.StorageSize(len(log.Topics)*common.HashLength + len(log.Data))
	}
	return size
}

// ReceiptForStorage is a wrapper around a Receipt that flattens and parses the
// entire content of a receipt, as opposed to only the consensus fields originally.
type ReceiptForStorage Receipt

// EncodeRLP implements rlp.Encoder, and flattens all content fields of a receipt
// into an RLP stream.
// Quorum:
// - added logic to support multiple private state and revert reason
// - original EncodeRLP is now encodeRLPOriginal
// Note that PMTReceipts also have TxHash & ContractAddress encoded, as needed for privacy marker transactions
func (r *ReceiptForStorage) EncodeRLP(w io.Writer) error {
	if r.PSReceipts == nil {
		if hasRevertReason((*Receipt)(r)) {
			return r.encodeRLPOriginalWithRevertReason(w)
		} else {
			return r.encodeRLPOriginal(w)
		}
	}
	return r.encodeRLPWithPSReceipts(w)
}

func (r *ReceiptForStorage) encodeRLPWithPSReceipts(w io.Writer) error {
	// if any PSReceipts have a RevertReason then encode all with RevertReason
	if hasRevertReason((*Receipt)(r)) {
		return r.encodeRLPForMPSWithRevertReason(w)
	}
	for _, psr := range r.PSReceipts {
		if hasRevertReason(psr) {
			return r.encodeRLPForMPSWithRevertReason(w)
		}
	}

	return r.encodeRLPForMPS(w)
}

func hasRevertReason(r *Receipt) bool {
	return r.RevertReason != nil && len(r.RevertReason) > 0
}

// Quorum - MPS
// encodeRLPForMPS includes Multiple Private State support
func (r *ReceiptForStorage) encodeRLPForMPS(w io.Writer) error {
	enc := &storedMPSReceiptRLP{
		PostStateOrStatus: (*Receipt)(r).statusEncoding(),
		CumulativeGasUsed: r.CumulativeGasUsed,
		Logs:              convertLogsForEncoding(r.Logs),
		PSReceipts:        convertPrivateReceiptsForEncoding(r.PSReceipts),
	}
	return rlp.Encode(w, enc)
}

// Quorum
// encodeRLPForMPSWithRevertReason includes Multiple Private State support & Revert Reason
func (r *ReceiptForStorage) encodeRLPForMPSWithRevertReason(w io.Writer) error {
	enc := &storedMPSReceiptRLPWithRevertReason{
		PostStateOrStatus: (*Receipt)(r).statusEncoding(),
		CumulativeGasUsed: r.CumulativeGasUsed,
		Logs:              convertLogsForEncoding(r.Logs),
		PSReceipts:        convertPrivateReceiptsWithRevertReasonForEncoding(r.PSReceipts),
		RevertReason:      r.RevertReason,
	}
	return rlp.Encode(w, enc)
}

// Quorum
// encodeRLPOriginalWithRevertReason includes Revert Reason
func (r *ReceiptForStorage) encodeRLPOriginalWithRevertReason(w io.Writer) error {
	enc := &storedReceiptRLPWithRevertReason{
		PostStateOrStatus: (*Receipt)(r).statusEncoding(),
		CumulativeGasUsed: r.CumulativeGasUsed,
		Logs:              convertLogsForEncoding(r.Logs),
		RevertReason:      r.RevertReason,
	}
	return rlp.Encode(w, enc)
}

// encodeRLPOriginal is the original from upstream
func (r *ReceiptForStorage) encodeRLPOriginal(w io.Writer) error {
	enc := &storedReceiptRLP{
		PostStateOrStatus: (*Receipt)(r).statusEncoding(),
		CumulativeGasUsed: r.CumulativeGasUsed,
		Logs:              make([]*LogForStorage, len(r.Logs)),
	}
	for i, log := range r.Logs {
		enc.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.Encode(w, enc)
}

// DecodeRLP implements rlp.Decoder, and loads both consensus and implementation
// fields of a receipt from an RLP stream.
func (r *ReceiptForStorage) DecodeRLP(s *rlp.Stream) error {
	// Retrieve the entire receipt blob as we need to try multiple decoders
	blob, err := s.Raw()
	if err != nil {
		return err
	}
	// Try decoding from the newest format for future proofness, then the older one
	// for old nodes that just upgraded. V4 was an intermediate unreleased format so
	// we do need to decode it, but it's not common (try last).

	if err := decodeStoredMPSReceiptRLPWithRevertReason(r, blob); err == nil {
		return nil
	}
	if err := decodeStoredMPSReceiptRLP(r, blob); err == nil {
		return nil
	}
	if err := decodeStoredReceiptRLPWithRevertReason(r, blob); err == nil {
		return nil
	}
	if err := decodeStoredReceiptRLP(r, blob); err == nil {
		return nil
	}
	if err := decodeV3StoredReceiptRLP(r, blob); err == nil {
		return nil
	}
	return decodeV4StoredReceiptRLP(r, blob)
}

func decodeStoredReceiptRLP(r *ReceiptForStorage, blob []byte) error {
	var stored storedReceiptRLP
	if err := rlp.DecodeBytes(blob, &stored); err != nil {
		return err
	}
	if err := (*Receipt)(r).setStatus(stored.PostStateOrStatus); err != nil {
		return err
	}
	r.CumulativeGasUsed = stored.CumulativeGasUsed
	r.Logs = make([]*Log, len(stored.Logs))
	for i, log := range stored.Logs {
		r.Logs[i] = (*Log)(log)
	}
	r.Bloom = CreateBloom(Receipts{(*Receipt)(r)})

	return nil
}

func decodeV4StoredReceiptRLP(r *ReceiptForStorage, blob []byte) error {
	var stored v4StoredReceiptRLP
	if err := rlp.DecodeBytes(blob, &stored); err != nil {
		return err
	}
	if err := (*Receipt)(r).setStatus(stored.PostStateOrStatus); err != nil {
		return err
	}
	r.CumulativeGasUsed = stored.CumulativeGasUsed
	r.TxHash = stored.TxHash
	r.ContractAddress = stored.ContractAddress
	r.GasUsed = stored.GasUsed
	r.Logs = make([]*Log, len(stored.Logs))
	for i, log := range stored.Logs {
		r.Logs[i] = (*Log)(log)
	}
	r.Bloom = CreateBloom(Receipts{(*Receipt)(r)})

	return nil
}

func decodeV3StoredReceiptRLP(r *ReceiptForStorage, blob []byte) error {
	var stored v3StoredReceiptRLP
	if err := rlp.DecodeBytes(blob, &stored); err != nil {
		return err
	}
	if err := (*Receipt)(r).setStatus(stored.PostStateOrStatus); err != nil {
		return err
	}
	r.CumulativeGasUsed = stored.CumulativeGasUsed
	r.Bloom = stored.Bloom
	r.TxHash = stored.TxHash
	r.ContractAddress = stored.ContractAddress
	r.GasUsed = stored.GasUsed
	r.Logs = make([]*Log, len(stored.Logs))
	for i, log := range stored.Logs {
		r.Logs[i] = (*Log)(log)
	}
	return nil
}

// Receipts is a wrapper around a Receipt array to implement DerivableList.
type Receipts []*Receipt

// Len returns the number of receipts in this list.
func (r Receipts) Len() int { return len(r) }

// GetRlp returns the RLP encoding of one receipt from the list.
func (r Receipts) GetRlp(i int) []byte {
	bytes, err := rlp.EncodeToBytes(r[i])
	if err != nil {
		panic(err)
	}
	return bytes
}

// Quorum
// CopyReceipts makes a deep copy of the given receipts.
func CopyReceipts(receipts []*Receipt) []*Receipt {
	result := make([]*Receipt, len(receipts))
	for i, receiptOrig := range receipts {
		receiptCopy := *receiptOrig
		result[i] = &receiptCopy

		if receiptOrig.PSReceipts != nil {
			receiptCopy.PSReceipts = make(map[PrivateStateIdentifier]*Receipt)
			for psi, psReceiptOrig := range receiptOrig.PSReceipts {
				psReceiptCpy := *psReceiptOrig
				result[i].PSReceipts[psi] = &psReceiptCpy
			}
		}
	}

	return result
}

// DeriveFields fills the receipts with their computed fields based on consensus
// data and contextual infos like containing block and transactions.
// Quorum:
// - Provide additional support for Multiple Private State and Privacy Marker Transactions,
//   where the private receipts are held under the relevant receipt.PSReceipts
// - Original DeriveFields func is now deriveFieldsOrig
func (r Receipts) DeriveFields(config *params.ChainConfig, hash common.Hash, number uint64, txs Transactions) error {
	// Will work on a copy of Receipts so we don't modify the original receipts until the end
	receiptsCopy := CopyReceipts(r)

	// flatten all the MPS receipts, so that we have a flat array for deriveFieldsOrig()
	allReceipts := make(map[PrivateStateIdentifier][]*Receipt) // Holds all public and private receipts, for each PSI
	allPublic := make([]*Receipt, 0)                           // All public receipts, & private receipts if MPS disabled
	for i := 0; i < len(receiptsCopy); i++ {
		receipt := receiptsCopy[i]
		tx := txs[i]

		// if receipt is public, append to all known PSIs
		// if private, append to all attached PSIs
		//    if new PSI, attach public version of all previous receipts
		// append public to all other PSIs

		if !tx.IsPrivate() {
			for psi := range allReceipts {
				allReceipts[psi] = append(allReceipts[psi], receipt)
			}
		}

		// if this is a private tx or a privacy marker tx then receipt.PSReceipts must be processed to
		// add the PSI version of the receipt to all the relevant PSI arrays
		for psi, privateReceipt := range receipt.PSReceipts {
			// if this PSI doesn't yet exist in allReceipts then add it
			if _, ok := allReceipts[psi]; !ok {
				allReceipts[psi] = append(make([]*Receipt, 0), allPublic...)
			}
			allReceipts[psi] = append(allReceipts[psi], privateReceipt)
		}
		// add the empty PSI receipt to all the currently tracked PSIs
		emptyReceipt := receipt.PSReceipts[EmptyPrivateStateIdentifier]
		for psi := range allReceipts {
			if len(allReceipts[psi]) < i+1 {
				allReceipts[psi] = append(allReceipts[psi], emptyReceipt)
			}
		}

		allPublic = append(allPublic, receipt)
	}

	// now we have all the receipts, so derive all their fields
	for _, receipts := range allReceipts {
		casted := Receipts(receipts)
		if err := casted.deriveFieldsOrig(config, hash, number, txs); err != nil {
			return err
		}
	}
	if err := Receipts(allPublic).deriveFieldsOrig(config, hash, number, txs); err != nil {
		return err
	}

	// fields now derived, put back into correct order
	tmp := make([]*Receipt, len(receiptsCopy))
	for i := 0; i < len(receiptsCopy); i++ {
		tmp[i] = allPublic[i]
		oldPsis := tmp[i].PSReceipts
		tmp[i].PSReceipts = nil
		if oldPsis != nil {
			tmp[i].PSReceipts = make(map[PrivateStateIdentifier]*Receipt)
		}
		for psi := range oldPsis {
			psiReceipt := allReceipts[psi][i]
			// check original receipt, so see if TxnHash was populated, if so then it was a PMT receipt
			if r[i].PSReceipts != nil && r[i].PSReceipts[psi] != nil && r[i].PSReceipts[psi].TxHash != (common.Hash{}) {
				// PMT private receipts - TxnHash & ContractAddress were decoded from store, so preserve those
				psiReceipt.TxHash = r[i].PSReceipts[psi].TxHash
				psiReceipt.ContractAddress = r[i].PSReceipts[psi].ContractAddress
			}
			tmp[i].PSReceipts[psi] = psiReceipt
		}
	}

	for i := 0; i < len(r); i++ {
		r[i] = tmp[i]
	}
	return nil
}

// deriveFieldsOrig is the original DeriveFields from upstream
func (r Receipts) deriveFieldsOrig(config *params.ChainConfig, hash common.Hash, number uint64, txs Transactions) error {
	signer := MakeSigner(config, new(big.Int).SetUint64(number))

	logIndex := uint(0)
	if len(txs) != len(r) {
		return errors.New("transaction and receipt count mismatch")
	}
	for i := 0; i < len(r); i++ {
		// The transaction hash can be retrieved from the transaction itself
		r[i].TxHash = txs[i].Hash()

		// block location fields
		r[i].BlockHash = hash
		r[i].BlockNumber = new(big.Int).SetUint64(number)
		r[i].TransactionIndex = uint(i)

		// The contract address can be derived from the transaction itself
		if txs[i].To() == nil {
			// Deriving the signer is expensive, only do if it's actually needed
			from, _ := Sender(signer, txs[i])
			r[i].ContractAddress = crypto.CreateAddress(from, txs[i].Nonce())
		}
		// The used gas can be calculated based on previous r
		if i == 0 {
			r[i].GasUsed = r[i].CumulativeGasUsed
		} else {
			r[i].GasUsed = r[i].CumulativeGasUsed - r[i-1].CumulativeGasUsed
		}
		// The derived log fields can simply be set from the block and transaction
		for j := 0; j < len(r[i].Logs); j++ {
			r[i].Logs[j].BlockNumber = number
			r[i].Logs[j].BlockHash = hash
			r[i].Logs[j].TxHash = r[i].TxHash
			r[i].Logs[j].TxIndex = uint(i)
			r[i].Logs[j].Index = logIndex
			logIndex++
		}
	}
	return nil
}

// Quorum

// storedMPSReceiptRLPWithRevertReason is the storage encoding of a receipt which contains
// receipts per PSI, with added revert reason,
// plus TxHash & ContractAddress (needed for privacy marker transactions)
type storedMPSReceiptRLPWithRevertReason struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Logs              []*LogForStorage
	RevertReason      []byte
	TxHash            common.Hash
	ContractAddress   common.Address
	PSReceipts        []storedPSIToReceiptMapEntryWithRevertReason
}

type storedPSIToReceiptMapEntryWithRevertReason struct {
	Key   PrivateStateIdentifier
	Value storedMPSReceiptRLPWithRevertReason
}

// storedMPSReceiptRLP is the storage encoding of a receipt which contains
// receipts per PSI
// plus TxHash & ContractAddress (needed for privacy marker transactions)
type storedMPSReceiptRLP struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Logs              []*LogForStorage
	TxHash            common.Hash
	ContractAddress   common.Address
	PSReceipts        []storedPSIToReceiptMapEntry
}

type storedPSIToReceiptMapEntry struct {
	Key   PrivateStateIdentifier
	Value storedMPSReceiptRLP
}

// storedReceiptRLPWithRevertReason is the storage encoding of a receipt from geth upstream, with added revert reason
type storedReceiptRLPWithRevertReason struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Logs              []*LogForStorage
	RevertReason      []byte
}

// Flatten takes a list of private receipts, which will be the "private" PSI receipt,
// and flatten all the MPS receipts into a single list, which the bloom can work with
func (r Receipts) Flatten() []*Receipt {
	var flattenedReceipts []*Receipt
	for _, privReceipt := range r {
		flattenedReceipts = append(flattenedReceipts, privReceipt)
		for _, psReceipt := range privReceipt.PSReceipts {
			flattenedReceipts = append(flattenedReceipts, psReceipt)
		}
	}
	return flattenedReceipts
}

func convertPrivateReceiptsWithRevertReasonForEncoding(psReceipts map[PrivateStateIdentifier]*Receipt) []storedPSIToReceiptMapEntryWithRevertReason {
	result := make([]storedPSIToReceiptMapEntryWithRevertReason, len(psReceipts))
	idx := 0
	for key, val := range psReceipts {
		rec := storedMPSReceiptRLPWithRevertReason{
			PostStateOrStatus: val.statusEncoding(),
			CumulativeGasUsed: val.CumulativeGasUsed,
			Logs:              make([]*LogForStorage, len(val.Logs)),
			RevertReason:      val.RevertReason,
			TxHash:            val.TxHash,
			ContractAddress:   val.ContractAddress,
		}
		for i, log := range val.Logs {
			rec.Logs[i] = (*LogForStorage)(log)
		}
		result[idx] = storedPSIToReceiptMapEntryWithRevertReason{Key: key, Value: rec}
		idx++
	}
	return result
}

func convertPrivateReceiptsForEncoding(psReceipts map[PrivateStateIdentifier]*Receipt) []storedPSIToReceiptMapEntry {
	result := make([]storedPSIToReceiptMapEntry, len(psReceipts))
	idx := 0
	for key, val := range psReceipts {
		rec := storedMPSReceiptRLP{
			PostStateOrStatus: val.statusEncoding(),
			CumulativeGasUsed: val.CumulativeGasUsed,
			Logs:              make([]*LogForStorage, len(val.Logs)),
			TxHash:            val.TxHash,
			ContractAddress:   val.ContractAddress,
		}
		for i, log := range val.Logs {
			rec.Logs[i] = (*LogForStorage)(log)
		}
		result[idx] = storedPSIToReceiptMapEntry{Key: key, Value: rec}
		idx++
	}
	return result
}

func convertPrivateReceiptsForDecoding(storedPSReceipts []storedPSIToReceiptMapEntry) (map[PrivateStateIdentifier]*Receipt, error) {
	if len(storedPSReceipts) <= 0 {
		return nil, nil
	}

	result := make(map[PrivateStateIdentifier]*Receipt)
	for _, entry := range storedPSReceipts {
		rec := &Receipt{}
		if err := rec.setStatus(entry.Value.PostStateOrStatus); err != nil {
			return nil, err
		}
		rec.CumulativeGasUsed = entry.Value.CumulativeGasUsed
		rec.Logs = make([]*Log, len(entry.Value.Logs))
		for i, log := range entry.Value.Logs {
			rec.Logs[i] = (*Log)(log)
			rec.Logs[i].PSI = entry.Key
		}
		rec.Bloom = CreateBloom(Receipts{rec})
		rec.TxHash = entry.Value.TxHash
		rec.ContractAddress = entry.Value.ContractAddress
		result[entry.Key] = rec
	}
	return result, nil
}

func convertPrivateReceiptsWithRevertReasonForDecoding(storedPSReceipts []storedPSIToReceiptMapEntryWithRevertReason) (map[PrivateStateIdentifier]*Receipt, error) {
	if len(storedPSReceipts) <= 0 {
		return nil, nil
	}

	result := make(map[PrivateStateIdentifier]*Receipt)
	for _, entry := range storedPSReceipts {
		rec := &Receipt{}
		if err := rec.setStatus(entry.Value.PostStateOrStatus); err != nil {
			return nil, err
		}
		rec.CumulativeGasUsed = entry.Value.CumulativeGasUsed
		rec.Logs = make([]*Log, len(entry.Value.Logs))
		for i, log := range entry.Value.Logs {
			rec.Logs[i] = (*Log)(log)
			rec.Logs[i].PSI = entry.Key
		}
		rec.Bloom = CreateBloom(Receipts{rec})
		rec.RevertReason = entry.Value.RevertReason
		rec.TxHash = entry.Value.TxHash
		rec.ContractAddress = entry.Value.ContractAddress
		result[entry.Key] = rec
	}
	return result, nil
}

func convertLogsForEncoding(logs []*Log) []*LogForStorage {
	result := make([]*LogForStorage, len(logs))
	for i, log := range logs {
		result[i] = (*LogForStorage)(log)
	}
	return result
}

func convertLogsForDecoding(storedLogs []*LogForStorage) []*Log {
	result := make([]*Log, len(storedLogs))
	for i, log := range storedLogs {
		result[i] = (*Log)(log)
	}
	return result
}

// Includes logic to support multiple private state.
// Note that PMTReceipts entries TxHash & ContractAddress are also decoded, as needed for privacy marker transactions
func decodeStoredMPSReceiptRLP(r *ReceiptForStorage, blob []byte) error {
	var stored storedMPSReceiptRLP
	if err := rlp.DecodeBytes(blob, &stored); err != nil {
		return err
	}
	if err := (*Receipt)(r).setStatus(stored.PostStateOrStatus); err != nil {
		return err
	}
	r.CumulativeGasUsed = stored.CumulativeGasUsed
	r.Logs = convertLogsForDecoding(stored.Logs)
	r.Bloom = CreateBloom(Receipts{(*Receipt)(r)})
	psReceipts, err := convertPrivateReceiptsForDecoding(stored.PSReceipts)
	if err != nil {
		return err
	}
	r.PSReceipts = psReceipts
	return nil
}

func decodeStoredMPSReceiptRLPWithRevertReason(r *ReceiptForStorage, blob []byte) error {
	var stored storedMPSReceiptRLPWithRevertReason
	if err := rlp.DecodeBytes(blob, &stored); err != nil {
		return err
	}
	if err := (*Receipt)(r).setStatus(stored.PostStateOrStatus); err != nil {
		return err
	}
	r.CumulativeGasUsed = stored.CumulativeGasUsed
	r.Logs = convertLogsForDecoding(stored.Logs)
	r.Bloom = CreateBloom(Receipts{(*Receipt)(r)})
	psReceipts, err := convertPrivateReceiptsWithRevertReasonForDecoding(stored.PSReceipts)
	if err != nil {
		return err
	}
	r.PSReceipts = psReceipts
	r.RevertReason = stored.RevertReason
	return nil
}

func decodeStoredReceiptRLPWithRevertReason(r *ReceiptForStorage, blob []byte) error {
	var stored storedReceiptRLPWithRevertReason
	if err := rlp.DecodeBytes(blob, &stored); err != nil {
		return err
	}
	if err := (*Receipt)(r).setStatus(stored.PostStateOrStatus); err != nil {
		return err
	}
	r.CumulativeGasUsed = stored.CumulativeGasUsed
	r.Logs = convertLogsForDecoding(stored.Logs)
	r.Bloom = CreateBloom(Receipts{(*Receipt)(r)})
	r.RevertReason = stored.RevertReason
	return nil
}

// End Quorum

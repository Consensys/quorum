package lc

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	bindings "github.com/ethereum/go-ethereum/lc/bind"
)

type LcServiceApi struct {
	routerServiceSession bindings.RouterServiceSession
	stdLcFacSession      bindings.StandardLCFactorySession
	upasLcFacSession     bindings.UPASLCFactorySession
	routerServiceAddress common.Address
	stdLcFacAddress      common.Address
	upasLcFacAddres      common.Address
}

type TransactionInput struct {
	Data string `json:"data"` // Optional resource locator within a backend
}

func toTransactionInput(tx *types.Transaction) TransactionInput {
	return TransactionInput{
		Data: fmt.Sprintf("0x%s", common.Bytes2Hex(tx.Data())),
	}
}

type SliceOfHex []string
type SliceOfHash []common.Hash

func (s SliceOfHex) toSliceOfHashes() SliceOfHash {
	hashes := make(SliceOfHash, len(s))
	for i, hex := range s {
		hashes[i] = common.HexToHash(hex)
	}
	return hashes
}

func (s SliceOfHash) toSliceByte32() [][32]byte {
	byte32s := make([][32]byte, len(s))
	for i, hash := range s {
		byte32s[i] = hash
	}
	return byte32s
}

type IStageContractContentParams struct {
	RootHash       string
	SignedTime     *big.Int
	PrevHash       string
	NumOfDocuments *big.Int
	ContentHash    SliceOfHex
	Url            string
	Acknowledge    []byte
	Signature      []byte
}

type IAmendRequestAmendStageParams struct {
	Stage    *big.Int
	SubStage *big.Int
	Content  IStageContractContentParams
}

func (i IStageContractContentParams) toBindingStageContractContent() bindings.IStageContractContent {
	return bindings.IStageContractContent{
		RootHash:       common.HexToHash(i.RootHash),
		SignedTime:     i.SignedTime,
		PrevHash:       common.HexToHash(i.PrevHash),
		NumOfDocuments: i.NumOfDocuments,
		ContentHash:    i.ContentHash.toSliceOfHashes().toSliceByte32(),
		Url:            i.Url,
		Acknowledge:    i.Acknowledge,
		Signature:      i.Signature,
	}
}

func (i IAmendRequestAmendStageParams) toAmendRequestAmendStage() bindings.IAmendRequestAmendStage {
	return bindings.IAmendRequestAmendStage{
		Stage:    i.Stage,
		SubStage: i.SubStage,
		Content:  i.Content.toBindingStageContractContent(),
	}
}

func (s *LcServiceApi) RouterService() common.Address {
	return s.routerServiceAddress
}

func (s *LcServiceApi) StandardLcFactory() common.Address {
	return s.stdLcFacAddress
}

func (s *LcServiceApi) UpasLcFactory() common.Address {
	return s.upasLcFacAddres
}

func (s *LcServiceApi) Amc() (common.Address, error) {
	return s.routerServiceSession.Amc()
}

func (s *LcServiceApi) GetLcAddress(_documentId big.Int) (common.Address, error) {
	result, err := s.routerServiceSession.GetAddress(&_documentId)
	return result.Contract, err
}

func (s *LcServiceApi) GetRootHash(_documentId big.Int) ([32]byte, error) {
	return s.routerServiceSession.GetRootHash(&_documentId)
}

func (s *LcServiceApi) GetStageContent(_documentId big.Int, _stage big.Int, _subStage big.Int) (bindings.IStageContractContent, error) {
	return s.routerServiceSession.GetStageContent(&_documentId, &_stage, &_subStage)
}

func (s *LcServiceApi) GetAmendmentRequest(_documentId big.Int, _requestId big.Int) (bindings.IAmendRequestRequest, error) {
	return s.routerServiceSession.GetAmendmentRequest(&_documentId, &_requestId)
}

func (s *LcServiceApi) IsAmendApproved(_documentId big.Int, _requestId big.Int) (bool, error) {
	return s.routerServiceSession.IsAmendApproved(&_documentId, &_requestId)
}

func (s *LcServiceApi) Approve(_documentId big.Int, _stage big.Int, _subStage big.Int, _content bindings.IStageContractContent) (TransactionInput, error) {
	tx, err := s.routerServiceSession.Approve(&_documentId, &_stage, &_subStage, _content)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) ApproveAmendment(_documentId big.Int, _requestId big.Int, _signature []byte) (TransactionInput, error) {
	tx, err := s.routerServiceSession.ApproveAmendment(&_documentId, &_requestId, _signature)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CloseLC(_documentId big.Int) (TransactionInput, error) {
	tx, err := s.routerServiceSession.CloseLC(&_documentId)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) FulfillAmendment(_documentId big.Int, _requestId big.Int) (TransactionInput, error) {
	tx, err := s.routerServiceSession.FulfillAmendment(&_documentId, &_requestId)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) SetAMC(_amc common.Address) (TransactionInput, error) {
	tx, err := s.routerServiceSession.SetAMC(_amc)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) SubmitAmendment(_documentId big.Int, _migratingStages SliceOfHex, _amendStage IAmendRequestAmendStageParams, _signature []byte) (TransactionInput, error) {
	tx, err := s.routerServiceSession.SubmitAmendment(&_documentId, _migratingStages.toSliceOfHashes().toSliceByte32(), _amendStage.toAmendRequestAmendStage(), _signature)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CreateLc(_parties SliceOfHex, _content IStageContractContentParams) (TransactionInput, error) {
	tx, err := s.stdLcFacSession.Create(_parties.toSliceOfHashes().toSliceByte32(), _content.toBindingStageContractContent())
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CreateUpasLc(_parties SliceOfHex, _content IStageContractContentParams) (TransactionInput, error) {
	tx, err := s.upasLcFacSession.Create(_parties.toSliceOfHashes().toSliceByte32(), _content.toBindingStageContractContent())
	return toTransactionInput(tx), err
}

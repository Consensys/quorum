package lc

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	bindings "github.com/ethereum/go-ethereum/lc/bind"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
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

func toTransactionInput(tx *types.Transaction) *TransactionInput {
	return &TransactionInput{
		Data: fmt.Sprintf("0x%s", common.Bytes2Hex(tx.Data())),
	}
}

type IStageContractContentParams struct {
	RootHash       common.Hash   `validate:"required"`
	SignedTime     *big.Int      `validate:"required"`
	PrevHash       common.Hash   `validate:"required"`
	NumOfDocuments *big.Int      `validate:"required"`
	ContentHash    []common.Hash `validate:"required,gt=1,dive,required"`
	Url            string        `validate:"url"`
	Acknowledge    hexutil.Bytes `validate:"required"`
	Signature      hexutil.Bytes `validate:"required"`
}

type IAmendRequestAmendStageParams struct {
	Stage    *big.Int
	SubStage *big.Int
	Content  IStageContractContentParams
}

func (i IStageContractContentParams) toBindingStageContractContent() bindings.IStageContractContent {
	return bindings.IStageContractContent{
		RootHash:       i.RootHash,
		SignedTime:     i.SignedTime,
		PrevHash:       i.PrevHash,
		NumOfDocuments: i.NumOfDocuments,
		ContentHash:    commonHashToSliceByte32(i.ContentHash),
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

func sliceByte32ToCommonHash(input [][common.HashLength]byte) []common.Hash {
	hashes := make([]common.Hash, len(input))
	for i, hex := range input {
		hashes[i] = hex
	}
	return hashes
}

func commonHashToSliceByte32(input []common.Hash) [][common.HashLength]byte {
	byte32s := make([][common.HashLength]byte, len(input))
	for i, hash := range input {
		byte32s[i] = hash
	}
	return byte32s
}

func bindingStageContractContent2IStageContractContentParams(i bindings.IStageContractContent) IStageContractContentParams {
	return IStageContractContentParams{
		RootHash:       i.RootHash,
		SignedTime:     i.SignedTime,
		PrevHash:       i.PrevHash,
		NumOfDocuments: i.NumOfDocuments,
		ContentHash: 		sliceByte32ToCommonHash(i.ContentHash),
		Url:            i.Url,
		Acknowledge:    i.Acknowledge,
		Signature:      i.Signature,
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

func (s *LcServiceApi) GetRootHash(_documentId big.Int) (string, error) {
	result, err := s.routerServiceSession.GetRootHash(&_documentId)
	return hexutil.Encode(result[:]), err
}

func (s *LcServiceApi) GetStageContent(_documentId big.Int, _stage big.Int, _subStage big.Int) (IStageContractContentParams, error) {
	result, err := s.routerServiceSession.GetStageContent(&_documentId, &_stage, &_subStage)
	return bindingStageContractContent2IStageContractContentParams(result), err
}

func (s *LcServiceApi) GetAmendmentRequest(_documentId big.Int, _requestId big.Int) (bindings.IAmendRequestRequest, error) {
	return s.routerServiceSession.GetAmendmentRequest(&_documentId, &_requestId)
}

func (s *LcServiceApi) IsAmendApproved(_documentId big.Int, _requestId big.Int) (bool, error) {
	return s.routerServiceSession.IsAmendApproved(&_documentId, &_requestId)
}

func (s *LcServiceApi) Approve(_documentId big.Int, _stage big.Int, _subStage big.Int, _content bindings.IStageContractContent) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.Approve(&_documentId, &_stage, &_subStage, _content)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) ApproveAmendment(_documentId big.Int, _requestId big.Int, _signature []byte) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.ApproveAmendment(&_documentId, &_requestId, _signature)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CloseLC(_documentId big.Int) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.CloseLC(&_documentId)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) FulfillAmendment(_documentId big.Int, _requestId big.Int) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.FulfillAmendment(&_documentId, &_requestId)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) SetAMC(_amc common.Address) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.SetAMC(_amc)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) SubmitAmendment(_documentId big.Int, _migratingStages []common.Hash, _amendStage IAmendRequestAmendStageParams, _signature []byte) (*TransactionInput, error) {
	err := validate.Struct(_amendStage)
	if err != nil {
		return nil, err
	}

	tx, err := s.routerServiceSession.SubmitAmendment(&_documentId, commonHashToSliceByte32(_migratingStages), _amendStage.toAmendRequestAmendStage(), _signature)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CreateLc(_parties []common.Hash, _content IStageContractContentParams) (*TransactionInput, error) {
	err := validate.Struct(_content)
	if err != nil {
		return nil, err
	}
	tx, err := s.stdLcFacSession.Create(commonHashToSliceByte32(_parties), _content.toBindingStageContractContent())
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CreateUpasLc(_parties []common.Hash, _content IStageContractContentParams) (*TransactionInput, error) {
	err := validate.Struct(_content)
	if err != nil {
		return nil, err
	}
	tx, err := s.upasLcFacSession.Create(commonHashToSliceByte32(_parties), _content.toBindingStageContractContent())
	return toTransactionInput(tx), err
}

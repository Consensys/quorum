package lc

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	bindings "github.com/ethereum/go-ethereum/lc/bind"
	pbindings "github.com/ethereum/go-ethereum/permission/v2/bind"
	"github.com/go-playground/validator/v10"
)

var (
	validate                           = validator.New()
	ErrRootHashNotEqual                = errors.New("root hash not equal to 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470")
	ErrRequiredAcknowledgeMessage      = errors.New("require acknowledge message signature at stage 1 4 5")
	ErrDocumentIdAreUsed               = errors.New("documentId has been used by one LC contract")
	ErrNewDocumentIdUnMatch            = errors.New("new documentId must match with 1st element in content hash")
	ErrNoDocsShouldEqOrLessContentHash = errors.New("number of documents should equal or smaller than length of content hash")
	ErrOutOfBound                      = errors.New("out of bound stage, stage must > 0 and <= 7")
)

type LcServiceApi struct {
	lcManagementSession  bindings.LCManagementSession
	routerServiceSession bindings.RouterServiceSession
	stdLcFacSession      bindings.StandardLCFactorySession
	upasLcFacSession     bindings.UPASLCFactorySession
	modeSession          bindings.ModeSession
	amendSession         bindings.AmendRequestSession
	lcSession            func(lcAddr common.Address) (bindings.LCSession, error)
	permInfSession       pbindings.PermInterfaceSession
	addressConfig        Config
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
	Acknowledge    hexutil.Bytes
	Signature      hexutil.Bytes `validate:"required"`
}

type IStageContractContentParamsCreateLC struct {
	RootHash       *common.Hash
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

type IConfirmedAmendment struct {
	IssuingBank        string        `validate:"required"`
	AdvisingBank       string        `validate:"required"`
	ReimbursingBank    string        `validate:"required"`
	IssuingBankSig     hexutil.Bytes `validate:"required"`
	AdvisingBankSig    hexutil.Bytes `validate:"required"`
	ReimbursingBankSig hexutil.Bytes `validate:"required"`
}

type IAmendmentRequest struct {
	TypeOf          *big.Int                      `validate:"required"`
	Proposer        common.Address                `validate:"required"`
	MigratingStages []common.Hash                 `validate:"required"`
	AmendStage      IAmendRequestAmendStageParams `validate:"required,gt=1,dive,required"`
	Confirmed       IConfirmedAmendment           `validate:"required"`
	IsFulfilled     bool                          `validate:"required"`
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

func (i IStageContractContentParamsCreateLC) toBindingStageContractContent() bindings.IStageContractContent {
	return bindings.IStageContractContent{
		RootHash:       *i.RootHash,
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
		ContentHash:    sliceByte32ToCommonHash(i.ContentHash),
		Url:            i.Url,
		Acknowledge:    i.Acknowledge,
		Signature:      i.Signature,
	}
}

func bindingAmendRequestRequest2IAmendRequestAmendStageParams(i bindings.IAmendRequestRequest) IAmendmentRequest {
	return IAmendmentRequest{
		TypeOf:          i.TypeOf,
		Proposer:        i.Proposer,
		MigratingStages: sliceByte32ToCommonHash(i.MigratingStages),
		AmendStage: IAmendRequestAmendStageParams{
			Stage:    i.AmendStage.Stage,
			SubStage: i.AmendStage.SubStage,
			Content: IStageContractContentParams{
				RootHash:       common.BytesToHash(i.AmendStage.Content.RootHash[:]),
				SignedTime:     i.AmendStage.Content.SignedTime,
				PrevHash:       common.BytesToHash(i.AmendStage.Content.PrevHash[:]),
				NumOfDocuments: i.AmendStage.Content.NumOfDocuments,
				ContentHash:    sliceByte32ToCommonHash(i.AmendStage.Content.ContentHash),
				Url:            i.AmendStage.Content.Url,
				Acknowledge:    i.AmendStage.Content.Acknowledge,
				Signature:      i.AmendStage.Content.Signature,
			},
		},
		Confirmed: IConfirmedAmendment{
			IssuingBank:        i.Confirmed.IssuingBank,
			AdvisingBank:       i.Confirmed.AdvisingBank,
			ReimbursingBank:    i.Confirmed.ReimbursingBank,
			IssuingBankSig:     i.Confirmed.IssuingBankSig,
			AdvisingBankSig:    i.Confirmed.AdvisingBankSig,
			ReimbursingBankSig: i.Confirmed.ReimbursingBankSig,
		},
		IsFulfilled: i.IsFulfilled,
	}
}

func (s *LcServiceApi) Addresses() Config {
	return s.addressConfig
}

func (s *LcServiceApi) GetLcAddress(_documentId common.Hash) (common.Address, error) {
	result, err := s.routerServiceSession.GetAddress(_documentId.Big())
	return result.Contract, err
}

func (s *LcServiceApi) GetRootHash(_documentId common.Hash) (string, error) {
	result, err := s.routerServiceSession.GetRootHash(_documentId.Big())
	return hexutil.Encode(result[:]), err
}

func (s *LcServiceApi) GetStageContent(_documentId common.Hash, _stage big.Int, _subStage big.Int) (IStageContractContentParams, error) {
	result, err := s.routerServiceSession.GetStageContent(_documentId.Big(), &_stage, &_subStage)
	return bindingStageContractContent2IStageContractContentParams(result), err
}

func (s *LcServiceApi) GetAmendmentRequest(_documentId common.Hash, _requestId common.Hash) (IAmendmentRequest, error) {
	result, err := s.routerServiceSession.GetAmendmentRequest(_documentId.Big(), _requestId.Big())
	return bindingAmendRequestRequest2IAmendRequestAmendStageParams(result), err
}

func (s *LcServiceApi) IsAmendApproved(_documentId common.Hash, _requestId common.Hash) (bool, error) {
	return s.routerServiceSession.IsAmendApproved(_documentId.Big(), _requestId.Big())
}

func (s *LcServiceApi) Mode() (string, error) {
	bool, err := s.modeSession.SwitchedToDAO()
	if bool {
		return "dao", err
	}
	return "admin", err
}

func (s *LcServiceApi) GetNonce(_proposer common.Address) (*big.Int, error) {
	return s.amendSession.Nonces(_proposer)
}

func (s *LcServiceApi) GetCounter(_documentId common.Hash) (*big.Int, error) {
	lcAddr, err := s.routerServiceSession.GetAddress(_documentId.Big())
	if err != nil {
		return nil, fmt.Errorf("documentId not found %v", err)
	}

	lcSession, err := s.lcSession(lcAddr.Contract)
	if err != nil {
		return nil, fmt.Errorf("unable to load lc contract %v", err)
	}

	return lcSession.GetCounter()
}

// Transactions
func (s *LcServiceApi) Whitelist(orgs []string) (*TransactionInput, error) {
	tx, err := s.lcManagementSession.Whitelist(orgs)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) UnWhitelist(orgs []string) (*TransactionInput, error) {
	tx, err := s.lcManagementSession.Unwhitelist(orgs)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) Approve(_documentId common.Hash, _stage big.Int, _subStage big.Int, _content IStageContractContentParams) (*TransactionInput, error) {
	// Stage = 4 or stage = 5
	if _stage.Cmp(big.NewInt(4)) == 0 || _stage.Cmp(big.NewInt(5)) == 0 {
		if len(_content.Acknowledge) == 0 {
			return nil, ErrRequiredAcknowledgeMessage
		}
	}

	// Stage > 7 or Stage == 0
	if _stage.Cmp(big.NewInt(7)) == 1 || _stage.Cmp(big.NewInt(0)) == 0 {
		return nil, ErrOutOfBound
	}

	// TODO add more check

	tx, err := s.routerServiceSession.Approve(_documentId.Big(), &_stage, &_subStage, _content.toBindingStageContractContent())
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) ApproveAmendment(_documentId common.Hash, _requestId common.Hash, _signature hexutil.Bytes) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.ApproveAmendment(_documentId.Big(), _requestId.Big(), _signature)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CloseLC(_documentId common.Hash) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.CloseLC(_documentId.Big())
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) FulfillAmendment(_documentId common.Hash, _requestId common.Hash) (*TransactionInput, error) {
	tx, err := s.routerServiceSession.FulfillAmendment(_documentId.Big(), _requestId.Big())
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) SubmitAmendment(_documentId common.Hash, _migratingStages []common.Hash, _amendStage IAmendRequestAmendStageParams, _signature hexutil.Bytes) (*TransactionInput, error) {
	err := validate.Struct(_amendStage)
	if err != nil {
		return nil, err
	}

	tx, err := s.routerServiceSession.SubmitAmendment(_documentId.Big(), commonHashToSliceByte32(_migratingStages), _amendStage.toAmendRequestAmendStage(), _signature)
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CreateLc(_parties []string, _content IStageContractContentParamsCreateLC) (*TransactionInput, error) {
	if _content.PrevHash.Hex() != _content.ContentHash[0].Hex() {
		return nil, ErrNewDocumentIdUnMatch
	}

	if int(_content.NumOfDocuments.Int64()) > len(_content.ContentHash) {
		return nil, ErrNoDocsShouldEqOrLessContentHash
	}

	err := validate.Struct(_content)
	if err != nil {
		return nil, err
	}

	addresses, err := s.stdLcFacSession.GetLCAddress(big.NewInt(0).SetBytes(_content.ContentHash[0][:]))
	if err != nil {
		return nil, err
	}

	if len(addresses) > 0 {
		return nil, ErrDocumentIdAreUsed
	}

	tx, err := s.stdLcFacSession.Create(_parties, _content.toBindingStageContractContent())
	return toTransactionInput(tx), err
}

func (s *LcServiceApi) CreateUpasLc(_parties []string, _content IStageContractContentParamsCreateLC) (*TransactionInput, error) {
	if _content.PrevHash.Hex() != _content.ContentHash[0].Hex() {
		return nil, ErrNewDocumentIdUnMatch
	}

	if int(_content.NumOfDocuments.Int64()) > len(_content.ContentHash) {
		return nil, ErrNoDocsShouldEqOrLessContentHash
	}

	err := validate.Struct(_content)
	if err != nil {
		return nil, err
	}

	addresses, err := s.upasLcFacSession.GetLCAddress(big.NewInt(0).SetBytes(_content.ContentHash[0][:]))
	if err != nil {
		return nil, err
	}

	if len(addresses) > 0 {
		return nil, ErrDocumentIdAreUsed
	}

	if _content.RootHash != nil {
		if _content.RootHash.Hex() != "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470" {
			return nil, ErrRootHashNotEqual
		}
	} else {
		hash := common.HexToHash("0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470")
		_content.RootHash = &hash
	}

	tx, err := s.upasLcFacSession.Create(_parties, _content.toBindingStageContractContent())
	return toTransactionInput(tx), err
}

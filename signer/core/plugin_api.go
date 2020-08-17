package core

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/plugin/account"
)

// <Quorum>

type approvalCreatorService struct {
	creator account.CreatorService
	ui      UIClientAPI
}

// NewApprovalCreatorService adds a wrapper to the provided creator service which requires UI approval before executing the service's methods
func NewApprovalCreatorService(creator account.CreatorService, ui UIClientAPI) account.CreatorService {
	return &approvalCreatorService{
		creator: creator,
		ui:      ui,
	}
}

func (s *approvalCreatorService) NewAccount(ctx context.Context, newAccountConfig interface{}) (accounts.Account, error) {
	if resp, err := s.ui.ApproveNewAccount(&NewAccountRequest{MetadataFromContext(ctx)}); err != nil {
		return accounts.Account{}, err
	} else if !resp.Approved {
		return accounts.Account{}, ErrRequestDenied
	}

	return s.creator.NewAccount(ctx, newAccountConfig)
}

// ImportRawKey is unsupported in the clef external API for parity with the available keystore account functionality
func (s *approvalCreatorService) ImportRawKey(_ context.Context, _ string, _ interface{}) (accounts.Account, error) {
	return accounts.Account{}, errors.New("not supported")
}

package qlight

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/multitenancy"
	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/rpc"
)

type privateBlockDataResolverImpl struct {
	privateStateManager mps.PrivateStateManager
	ptm                 private.PrivateTransactionManager
}

func NewPrivateBlockDataResolver(privateStateManager mps.PrivateStateManager, ptm private.PrivateTransactionManager) PrivateBlockDataResolver {
	return &privateBlockDataResolverImpl{privateStateManager: privateStateManager, ptm: ptm}
}

func (p *privateBlockDataResolverImpl) PrepareBlockPrivateData(block *types.Block, psi string) (*BlockPrivateData, error) {
	PSI := types.PrivateStateIdentifier(psi)
	var pvtTxs []PrivateTransactionData
	psm, err := p.privateStateManager.ResolveForUserContext(rpc.WithPrivateStateIdentifier(context.Background(), PSI))
	if err != nil {
		return nil, err
	}
	for _, tx := range block.Transactions() {
		if tx.IsPrivacyMarker() {
			ptd, err := p.fetchPrivateData(tx.Data(), psm)
			if err != nil {
				return nil, err
			}
			if ptd != nil {
				pvtTxs = append(pvtTxs, *ptd)
			}

			innerTx, _, _, _ := private.FetchPrivateTransactionWithPTM(tx.Data(), p.ptm)
			if innerTx != nil {
				tx = innerTx
			}
		}

		if tx.IsPrivate() {
			ptd, err := p.fetchPrivateData(tx.Data(), psm)
			if err != nil {
				return nil, err
			}
			if ptd != nil {
				pvtTxs = append(pvtTxs, *ptd)
			}
		}
	}
	if len(pvtTxs) == 0 {
		return nil, nil
	}

	var privateStateRoot = common.Hash{}

	stateRepo, err := p.privateStateManager.StateRepository(block.Root())
	if err != nil {
		log.Debug("Unable to retrieve private state repo while preparing the private block data", "block.No", block.Number(), "psi", psi, "err", err)
	} else {
		privateStateRoot, err = stateRepo.PrivateStateRoot(PSI)
		if err != nil {
			log.Debug("Unable to retrieve private state root while preparing the private block data", "block.No", block.Number(), "psi", psi, "err", err)
		}
	}

	return &BlockPrivateData{
		BlockHash:           block.Hash(),
		PSI:                 PSI,
		PrivateStateRoot:    privateStateRoot,
		PrivateTransactions: pvtTxs,
	}, nil
}

func (p *privateBlockDataResolverImpl) fetchPrivateData(encryptedPayloadHash []byte, psm *mps.PrivateStateMetadata) (*PrivateTransactionData, error) {
	txHash := common.BytesToEncryptedPayloadHash(encryptedPayloadHash)
	_, _, privateTx, extra, err := p.ptm.Receive(txHash)
	if err != nil {
		return nil, err
	}
	// we're not party to this transaction
	if privateTx == nil {
		return nil, nil
	}
	if p.privateStateManager.NotIncludeAny(psm, extra.ManagedParties...) {
		return nil, nil
	}

	extra.ManagedParties = psm.FilterAddresses(extra.ManagedParties...)

	ptd := PrivateTransactionData{
		Hash:     &txHash,
		Payload:  privateTx,
		Extra:    extra,
		IsSender: false,
	}
	if len(psm.Addresses) == 0 {
		// this is not an MPS node so we have to ask tessera
		ptd.IsSender, err = p.ptm.IsSender(txHash)
		if err != nil {
			return nil, err
		}
	} else {
		// this is an MPS node so we can speed up the IsSender logic by checking the addresses in the private state metadata
		ptd.IsSender = !psm.NotIncludeAny(extra.Sender)
	}

	return &ptd, nil
}

type authProviderImpl struct {
	privateStateManager mps.PrivateStateManager
	authManagerProvider AuthManagerProvider
	authManager         security.AuthenticationManager
	enabled             bool
}

func NewAuthProvider(privateStateManager mps.PrivateStateManager, authManagerProvider AuthManagerProvider) AuthProvider {
	return &authProviderImpl{
		privateStateManager: privateStateManager,
		authManagerProvider: authManagerProvider,
		enabled:             false,
	}
}

func (a *authProviderImpl) Initialize() error {
	if a.authManagerProvider != nil {
		a.authManager = a.authManagerProvider()
		if a.authManager == nil {
			return nil
		}
		authEnabled, err := a.authManager.IsEnabled(context.Background())
		if err != nil {
			return err
		}
		a.enabled = authEnabled
	}
	return nil
}

func (a *authProviderImpl) Authorize(token string, psi string) error {
	if !a.enabled {
		return nil
	}

	authToken, err := a.authManager.Authenticate(context.Background(), token)
	if err != nil {
		return err
	}
	PSI := types.PrivateStateIdentifier(psi)
	// check that we have access to the relevant PSI
	psiAuth, err := multitenancy.IsPSIAuthorized(authToken, PSI)
	if err != nil {
		return err
	}
	if !psiAuth {
		return fmt.Errorf("PSI not authorized")
	}
	// check that we have access to  qlight://p2p , rpc://eth_*
	qlightP2P := false
	rpcETH := false
	for _, ga := range authToken.GetAuthorities() {
		if ga.GetRaw() == "p2p://qlight" {
			qlightP2P = true
		}
		if ga.GetRaw() == "rpc://eth_*" {
			rpcETH = true
		}
	}
	if !qlightP2P || !rpcETH {
		return fmt.Errorf("The P2P token does not have the necessary authorization p2p=%v rpcETH=%v", qlightP2P, rpcETH)
	}
	// try to resolve the PSI
	_, err = a.privateStateManager.ResolveForUserContext(rpc.WithPrivateStateIdentifier(context.Background(), PSI))
	if err != nil {
		return fmt.Errorf("QLight auth error: %w", err)
	}
	return nil
}

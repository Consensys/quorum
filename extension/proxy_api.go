package extension

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

type PrivateExtensionProxyAPI struct {
	PrivateExtensionAPI
	proxyClient *rpc.Client
}

func NewPrivateExtensionProxyAPI(privacyService *PrivacyService) interface{} {
	apiSupport, ok := privacyService.apiBackendHelper.(ethapi.ProxyAPISupport)
	if ok {
		if apiSupport.ProxyEnabled() {
			return &PrivateExtensionProxyAPI{
				PrivateExtensionAPI{privacyService},
				apiSupport.ProxyClient(),
			}
		}
	}
	return NewPrivateExtensionAPI(privacyService)
}

// ActiveExtensionContracts returns the list of all currently outstanding extension contracts
func (api *PrivateExtensionProxyAPI) ActiveExtensionContracts(ctx context.Context) []ExtensionContract {
	api.privacyService.mu.Lock()
	defer api.privacyService.mu.Unlock()

	psi, err := api.privacyService.apiBackendHelper.PSMR().ResolveForUserContext(ctx)
	if err != nil {
		return nil
	}

	extracted := make([]ExtensionContract, 0)
	for _, contract := range api.privacyService.psiContracts[psi.ID] {
		extracted = append(extracted, *contract)
	}

	return extracted
}

// ApproveContractExtension submits the vote to the specified extension management contract. The vote indicates whether to extend
// a given contract to a new participant or not
func (api *PrivateExtensionProxyAPI) ApproveExtension(ctx context.Context, addressToVoteOn common.Address, vote bool, txa ethapi.SendTxArgs) (string, error) {
	log.Info("QLight - proxy enabled")
	var result string
	err := api.proxyClient.CallContext(ctx, &result, "quorumExtension_approveExtension", addressToVoteOn, vote, txa)
	return result, err
}

func (api *PrivateExtensionProxyAPI) ExtendContract(ctx context.Context, toExtend common.Address, newRecipientPtmPublicKey string, recipientAddr common.Address, txa ethapi.SendTxArgs) (string, error) {
	log.Info("QLight - proxy enabled")
	var result string
	err := api.proxyClient.CallContext(ctx, &result, "quorumExtension_extendContract", toExtend, newRecipientPtmPublicKey, recipientAddr, txa)
	return result, err
}

func (api *PrivateExtensionProxyAPI) CancelExtension(ctx context.Context, extensionContract common.Address, txa ethapi.SendTxArgs) (string, error) {
	log.Info("QLight - proxy enabled")
	var result string
	err := api.proxyClient.CallContext(ctx, &result, "quorumExtension_cancelExtension", extensionContract, txa)
	return result, err
}

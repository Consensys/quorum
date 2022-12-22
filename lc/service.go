package lc

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	bindings "github.com/ethereum/go-ethereum/lc/bind"
	"github.com/ethereum/go-ethereum/node"
	ptypes "github.com/ethereum/go-ethereum/permission/core/types"
	pbindings "github.com/ethereum/go-ethereum/permission/v2/bind"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	//default gas limit to use if not passed in sendTxArgs
	defaultGasLimit = uint64(4712384)
	//default gas price to use if not passed in sendTxArgs
	defaultGasPrice = big.NewInt(0)
)

var nullSigner = func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return transaction, nil
}

type LcService struct {
	Cfg     *Config
	PerCfg  ptypes.PermissionConfig
	Node    *node.Node
	api     *LcServiceApi
	ethClnt bind.ContractBackend
}

func NewLcService(cfg Config, perCfg ptypes.PermissionConfig, stack *node.Node) (*LcService, error) {
	s := &LcService{Cfg: &cfg, PerCfg: perCfg, Node: stack}
	client, err := s.Node.Attach()
	if err != nil {
		return nil, fmt.Errorf("unable to create rpc client: %v", err)
	}
	s.ethClnt = ethclient.NewClient(client)

	routerSvc, err := bindings.NewRouterService(s.Cfg.RouterAddress, s.ethClnt)
	if err != nil {
		return nil, fmt.Errorf("unable to create RouterService contract: %v", err)
	}
	s.api = &LcServiceApi{}
	s.api.routerServiceSession = bindings.RouterServiceSession{
		Contract:     routerSvc,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
	}

	stdLcSvc, err := bindings.NewStandardLCFactory(s.Cfg.StandardFactoryAddress, s.ethClnt)
	if err != nil {
		return nil, fmt.Errorf("unable to create NewStandardLCFactory contract: %v", err)
	}
	s.api.stdLcFacSession = bindings.StandardLCFactorySession{
		Contract:     stdLcSvc,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
	}

	upasLcSvc, err := bindings.NewUPASLCFactory(s.Cfg.UPASFactoryAddress, s.ethClnt)
	if err != nil {
		return nil, fmt.Errorf("unable to create UPASLCFactorySession contract: %v", err)
	}
	s.api.upasLcFacSession = bindings.UPASLCFactorySession{
		Contract:     upasLcSvc,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
	}

	lcManagement, err := bindings.NewLCManagement(s.Cfg.LCManagementAddress, s.ethClnt)
	if err != nil {
		return nil, fmt.Errorf("unable to create NewLCManagement contract: %v", err)
	}
	s.api.lcManagementSession = bindings.LCManagementSession{
		Contract:     lcManagement,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
	}

	mode, err := bindings.NewMode(s.Cfg.ModeAddress, s.ethClnt)
	if err != nil {
		return nil, fmt.Errorf("unable to create NewMode contract: %v", err)
	}
	s.api.modeSession = bindings.ModeSession{
		Contract:     mode,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
	}

	amend, err := bindings.NewAmendRequest(s.Cfg.AmendRequestAddress, s.ethClnt)
	if err != nil {
		return nil, fmt.Errorf("unable to create NewAmendRequest contract: %v", err)
	}
	s.api.amendSession = bindings.AmendRequestSession{
		Contract:     amend,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
	}

	permInterf, err := pbindings.NewPermInterface(s.PerCfg.InterfAddress, s.ethClnt)
	if err != nil {
		return nil, fmt.Errorf("unable to create NewAmendRequest contract: %v", err)
	}

	s.api.permInfSession = pbindings.PermInterfaceSession{
		Contract:     permInterf,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
	}

	s.api.lcSession = func(lcAddr common.Address) (bindings.LCSession, error){
		lc, err := bindings.NewLC(lcAddr, s.ethClnt)

		lcSession := bindings.LCSession{
			Contract:     lc,
			CallOpts:     bind.CallOpts{Pending: false},
			TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
		}

		return lcSession, err
	}

	s.api.addressConfig = cfg
	return s, nil
}

func (s *LcService) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "lc",
			Service:   s.api,
			Version:   "0.3",
			Public:    false,
		},
	}
}

func (s LcService) Start() error {
	return nil
}

func (s LcService) Stop() error {
	return nil
}

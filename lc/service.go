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
	Node    *node.Node
	api     *LcServiceApi
	ethClnt bind.ContractBackend
}

func NewLcService(cfg Config, stack *node.Node) (*LcService, error) {
	s := &LcService{Cfg: &cfg, Node: stack}
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
		return nil, fmt.Errorf("unable to create NewStandardLCFactory contract: %v", err)
	}
	s.api.upasLcFacSession = bindings.UPASLCFactorySession{
		Contract:     upasLcSvc,
		CallOpts:     bind.CallOpts{Pending: false},
		TransactOpts: bind.TransactOpts{NoSend: true, GasPrice: defaultGasPrice, GasLimit: defaultGasLimit, Signer: nullSigner},
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

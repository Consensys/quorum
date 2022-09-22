package lc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	bindings "github.com/ethereum/go-ethereum/lc/bind"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
)

type LcService struct {
	Cfg     *Config
	Node    *node.Node
	api     *LcServiceApi
	ethClnt bind.ContractBackend
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
	client, err := s.Node.Attach()
	if err != nil {
		return fmt.Errorf("unable to create rpc client: %v", err)
	}
	s.ethClnt = ethclient.NewClient(client)

	routerSvc, err := bindings.NewRouterService(s.Cfg.RouterAddress, s.ethClnt)
	if err != nil {
		return fmt.Errorf("unable to create RouterService contract: %v", err)
	}
	s.api = &LcServiceApi{}
	s.api.routerServiceSession = bindings.RouterServiceSession{
		Contract: routerSvc,
	}

	stdLcSvc, err := bindings.NewStandardLCFactory(s.Cfg.StandardFactoryAddress, s.ethClnt)
	if err != nil {
		return fmt.Errorf("unable to create NewStandardLCFactory contract: %v", err)
	}
	s.api.stdLcFacSession = bindings.StandardLCFactorySession{
		Contract: stdLcSvc,
	}

	upasLcSvc, err := bindings.NewUPASLCFactory(s.Cfg.UPASFactoryAddress, s.ethClnt)
	if err != nil {
		return fmt.Errorf("unable to create NewStandardLCFactory contract: %v", err)
	}
	s.api.upasLcFacSession = bindings.UPASLCFactorySession{
		Contract: upasLcSvc,
	}
	return nil
}

func (s LcService) Stop() error {
	return nil
}

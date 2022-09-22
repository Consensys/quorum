package lc

import (
	"github.com/ethereum/go-ethereum/common"
	bindings "github.com/ethereum/go-ethereum/lc/bind"
)

type LcServiceApi struct {
	routerServiceSession bindings.RouterServiceSession
	stdLcFacSession      bindings.StandardLCFactorySession
	upasLcFacSession     bindings.UPASLCFactorySession
}

func (s *LcServiceApi) Amc() (common.Address, error) {
	return s.routerServiceSession.Amc()
}

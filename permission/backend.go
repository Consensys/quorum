package permission

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/permission/basic"
	bind2 "github.com/ethereum/go-ethereum/permission/basic/bind"
	"github.com/ethereum/go-ethereum/permission/eea"
	"github.com/ethereum/go-ethereum/permission/eea/bind"
	ptype "github.com/ethereum/go-ethereum/permission/types"
)

func NewPermissionContractService(ethClnt bind.ContractBackend, eeaFlag bool, key *ecdsa.PrivateKey,
	permConfig *types.PermissionConfig) ptype.ContractService {
	if eeaFlag {
		return &eea.Contract{EthClnt: ethClnt, Key: key, PermConfig: permConfig}
	}
	return &basic.Contract{EthClnt: ethClnt, Key: key, PermConfig: permConfig}
}

func NewPermissionContractServiceForApi(p *PermissionCtrl, transactOpts *bind.TransactOpts) ptype.ContractService {
	switch p.eeaFlag {
	case true:
		pc := p.contract.(*eea.Contract)
		ps := &permission.PermInterfaceSession{
			Contract: pc.PermInterf,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: *transactOpts,
		}
		return &eea.Contract{PermInterfSession: ps, PermConfig: p.permConfig}

	default:
		pc := p.contract.(*basic.Contract)
		ps := &bind2.PermInterfaceSession{
			Contract: pc.PermInterf,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: *transactOpts,
		}
		return &basic.Contract{PermInterfSession: ps, PermConfig: p.permConfig}
	}
}

func (p *PermissionCtrl) populateBackEnd() {
	isRaft := false
	if p.eth != nil {
		isRaft = p.eth.BlockChain().Config().Istanbul == nil && p.eth.BlockChain().Config().Clique == nil
	}
	switch p.eeaFlag {
	case true:
		p.backend = &eea.Backend{
			Node:    p.node,
			IsRaft:  isRaft,
			DataDir: p.dataDir}

	default:
		p.backend = &basic.Backend{
			Node:    p.node,
			IsRaft:  isRaft,
			DataDir: p.dataDir}
	}
}

func (p *PermissionCtrl) populateContractInterface() {
	p.contract = NewPermissionContractService(p.ethClnt, p.eeaFlag, p.key, p.permConfig)
	switch p.eeaFlag {
	case true:
		p.backend.(*eea.Backend).Contr = p.contract.(*eea.Contract)
	default:
		p.backend.(*basic.Backend).Contr = p.contract.(*basic.Contract)
	}
}

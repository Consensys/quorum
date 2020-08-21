package eea

import (
	"crypto/ecdsa"
	"math/big"

	ptype "github.com/ethereum/go-ethereum/permission/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	binding "github.com/ethereum/go-ethereum/permission/eea/bind"
)

type Contract struct {
	EthClnt    bind.ContractBackend
	Key        *ecdsa.PrivateKey
	PermConfig *types.PermissionConfig

	//binding contracts
	PermUpgr   *binding.EeaPermUpgr
	PermInterf *binding.EeaPermInterface
	PermNode   *binding.EeaNodeManager
	PermAcct   *binding.EeaAcctManager
	PermRole   *binding.EeaRoleManager
	PermOrg    *binding.EeaOrgManager
	//sessions
	PermInterfSession *binding.EeaPermInterfaceSession
	permOrgSession    *binding.EeaOrgManagerSession
	permNodeSession   *binding.EeaNodeManagerSession
	permRoleSession   *binding.EeaRoleManagerSession
	permAcctSession   *binding.EeaAcctManagerSession
}

func (p *Contract) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {

	return p.PermInterfSession.RemoveRole(_roleId, _orgId)
}

func (p *Contract) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {

	return p.PermInterfSession.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}

func (p *Contract) ConnectionAllowedImpl(_enodeId string, _ip string, _port uint16, _raftport uint16) (bool, error) {

	return p.PermInterfSession.ConnectionAllowedImpl(_enodeId, _ip, _port, _raftport)
}

func (p *Contract) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {

	return p.PermInterfSession.TransactionAllowed(_srcaccount, _tgtaccount)
}

func (p *Contract) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {

	return p.PermInterfSession.AssignAccountRole(_account, _orgId, _roleId)
}

func (p *Contract) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {

	return p.PermInterfSession.UpdateAccountStatus(_orgId, _account, _action)
}

func (p *Contract) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.PermInterfSession.ApproveBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *Contract) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.PermInterfSession.StartBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *Contract) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.PermInterfSession.StartBlacklistedAccountRecovery(_orgId, _account)
}

func (p *Contract) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.PermInterfSession.ApproveBlacklistedAccountRecovery(_orgId, _account)
}

func (p *Contract) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return p.PermInterfSession.GetPendingOp(_orgId)
}

func (p *Contract) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.PermInterfSession.ApproveAdminRole(_orgId, _account)
}

func (p *Contract) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {

	return p.PermInterfSession.AssignAdminRole(_orgId, _account, _roleId)
}

func (p *Contract) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.PermInterfSession.AddNode(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *Contract) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string, _action *big.Int) (*types.Transaction, error) {

	return p.PermInterfSession.UpdateNodeStatus(_orgId, _enodeId, _ip, _port, _raftport, _action)

}

func (p *Contract) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.PermInterfSession.ApproveOrgStatus(_orgId, _action)
}

func (p *Contract) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.PermInterfSession.UpdateOrgStatus(_orgId, _action)
}

func (p *Contract) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	return p.PermInterfSession.ApproveOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}

func (p *Contract) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.PermInterfSession.AddSubOrg(_pOrgId, _orgId, _enodeId, _ip, _port, _raftport)
}

func (p *Contract) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {

	return p.PermInterfSession.AddOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}

func (p *Contract) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {

	return p.permAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (p *Contract) GetNumberOfAccounts() (*big.Int, error) {

	return p.permAcctSession.GetNumberOfAccounts()
}

func (p *Contract) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (p *Contract) GetNumberOfRoles() (*big.Int, error) {
	return p.permRoleSession.GetNumberOfRoles()
}

func (p *Contract) GetNumberOfOrgs() (*big.Int, error) {
	return p.permOrgSession.GetNumberOfOrgs()
}

func (p *Contract) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return p.PermInterfSession.UpdateNetworkBootStatus()
}

func (p *Contract) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return p.PermInterfSession.AddAdminAccount(_acct)
}

func (p *Contract) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return p.PermInterfSession.AddAdminNode(_enodeId, _ip, _port, _raftport)
}

func (p *Contract) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return p.PermInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (p *Contract) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return p.PermInterfSession.Init(_breadth, _depth)
}

func (p *Contract) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return p.permAcctSession.GetAccountDetails(_account)
}

func (p *Contract) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, types.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport), r.NodeStatus, err
}

func (p *Contract) GetNumberOfNodes() (*big.Int, error) {
	return p.permNodeSession.GetNumberOfNodes()
}

func (p *Contract) GetNodeDetails(enodeId string) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetails(enodeId)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, types.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport), r.NodeStatus, err
}

func (p *Contract) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (p *Contract) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return p.permOrgSession.GetSubOrgIndexes(_orgId)
}

func (p *Contract) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgInfo(_orgIndex)
}

func (p *Contract) GetNetworkBootStatus() (bool, error) {
	return p.PermInterfSession.GetNetworkBootStatus()
}

func (p *Contract) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *Contract) AfterStart() error {
	log.Debug("permission service: binding contracts")

	err := p.eeaBindContract()
	if err != nil {
		return err
	}

	p.initSession()
	return nil
}

func (p *Contract) eeaBindContract() error {
	if err := ptype.BindContract(&p.PermUpgr, func() (interface{}, error) { return binding.NewEeaPermUpgr(p.PermConfig.UpgrdAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermInterf, func() (interface{}, error) { return binding.NewEeaPermInterface(p.PermConfig.InterfAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermAcct, func() (interface{}, error) { return binding.NewEeaAcctManager(p.PermConfig.AccountAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermNode, func() (interface{}, error) { return binding.NewEeaNodeManager(p.PermConfig.NodeAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermRole, func() (interface{}, error) { return binding.NewEeaRoleManager(p.PermConfig.RoleAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermOrg, func() (interface{}, error) { return binding.NewEeaOrgManager(p.PermConfig.OrgAddress, p.EthClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *Contract) initSession() {
	auth := bind.NewKeyedTransactor(p.Key)
	p.PermInterfSession = &binding.EeaPermInterfaceSession{
		Contract: p.PermInterf,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 47000000,
			GasPrice: big.NewInt(0),
		},
	}

	p.permOrgSession = &binding.EeaOrgManagerSession{
		Contract: p.PermOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSession = &binding.EeaNodeManagerSession{
		Contract: p.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSession = &binding.EeaRoleManagerSession{
		Contract: p.PermRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSession = &binding.EeaAcctManagerSession{
		Contract: p.PermAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

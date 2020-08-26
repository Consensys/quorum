package eea

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	binding "github.com/ethereum/go-ethereum/permission/eea/bind"
	ptype "github.com/ethereum/go-ethereum/permission/types"
)

type Contract struct {
	EthClnt    bind.ContractBackend
	Key        *ecdsa.PrivateKey
	PermConfig *types.PermissionConfig
	IsRaft     bool
	UseDns     bool

	//binding contracts
	PermUpgr   *binding.PermUpgr
	PermInterf *binding.PermInterface
	PermNode   *binding.NodeManager
	PermAcct   *binding.AcctManager
	PermRole   *binding.RoleManager
	PermOrg    *binding.OrgManager
	//sessions
	PermInterfSession *binding.PermInterfaceSession
	permOrgSession    *binding.OrgManagerSession
	permNodeSession   *binding.NodeManagerSession
	permRoleSession   *binding.RoleManagerSession
	permAcctSession   *binding.AcctManagerSession
}

func (p *Contract) RemoveRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.RemoveRole(_args.RoleId, _args.OrgId)
}

func (p *Contract) AddNewRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.AddNewRole(_args.RoleId, _args.OrgId, big.NewInt(int64(_args.AccessType)), _args.IsVoter, _args.IsAdmin)
}

func (p *Contract) ConnectionAllowedImpl(url string) (bool, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(url)
	if err != nil {
		return false, err
	}

	return p.PermInterfSession.ConnectionAllowedImpl(enodeId, ip, port, raftPort)
}

func (p *Contract) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {
	return p.PermInterfSession.TransactionAllowed(_srcaccount, _tgtaccount)
}

func (p *Contract) AssignAccountRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.AssignAccountRole(_args.AcctId, _args.OrgId, _args.RoleId)
}

func (p *Contract) UpdateAccountStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.UpdateAccountStatus(_args.OrgId, _args.AcctId, big.NewInt(int64(_args.Action)))
}

func (p *Contract) ApproveBlacklistedNodeRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(_args.Url)
	if err != nil {
		return nil, err
	}
	return p.PermInterfSession.ApproveBlacklistedNodeRecovery(_args.OrgId, enodeId, ip, port, raftPort)
}

func (p *Contract) StartBlacklistedNodeRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(_args.Url)
	if err != nil {
		return nil, err
	}
	return p.PermInterfSession.StartBlacklistedNodeRecovery(_args.OrgId, enodeId, ip, port, raftPort)
}

func (p *Contract) StartBlacklistedAccountRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.StartBlacklistedAccountRecovery(_args.OrgId, _args.AcctId)
}

func (p *Contract) ApproveBlacklistedAccountRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.ApproveBlacklistedAccountRecovery(_args.OrgId, _args.AcctId)
}

func (p *Contract) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return p.PermInterfSession.GetPendingOp(_orgId)
}

func (p *Contract) ApproveAdminRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.ApproveAdminRole(_args.OrgId, _args.AcctId)
}

func (p *Contract) AssignAdminRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.AssignAdminRole(_args.OrgId, _args.AcctId, _args.RoleId)
}

func (p *Contract) AddNode(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(_args.Url)
	if err != nil {
		return nil, err
	}

	return p.PermInterfSession.AddNode(_args.OrgId, enodeId, ip, port, raftPort)
}

func (p *Contract) UpdateNodeStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(_args.Url)
	if err != nil {
		return nil, err
	}
	return p.PermInterfSession.UpdateNodeStatus(_args.OrgId, enodeId, ip, port, raftPort, big.NewInt(int64(_args.Action)))
}

func (p *Contract) ApproveOrgStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.ApproveOrgStatus(_args.OrgId, big.NewInt(int64(_args.Action)))
}

func (p *Contract) UpdateOrgStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	return p.PermInterfSession.UpdateOrgStatus(_args.OrgId, big.NewInt(int64(_args.Action)))
}

func (p *Contract) ApproveOrg(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(_args.Url)
	if err != nil {
		return nil, err
	}
	return p.PermInterfSession.ApproveOrg(_args.OrgId, enodeId, ip, port, raftPort, _args.AcctId)
}

func (p *Contract) AddSubOrg(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(_args.Url)
	if err != nil {
		return nil, err
	}
	return p.PermInterfSession.AddSubOrg(_args.POrgId, _args.OrgId, enodeId, ip, port, raftPort)
}

func (p *Contract) AddOrg(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := p.getNodeDetails(_args.Url)
	if err != nil {
		return nil, err
	}
	return p.PermInterfSession.AddOrg(_args.OrgId, enodeId, ip, port, raftPort, _args.AcctId)
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
	if err := ptype.BindContract(&p.PermUpgr, func() (interface{}, error) { return binding.NewPermUpgr(p.PermConfig.UpgrdAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermInterf, func() (interface{}, error) { return binding.NewPermInterface(p.PermConfig.InterfAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermAcct, func() (interface{}, error) { return binding.NewAcctManager(p.PermConfig.AccountAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermNode, func() (interface{}, error) { return binding.NewNodeManager(p.PermConfig.NodeAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermRole, func() (interface{}, error) { return binding.NewRoleManager(p.PermConfig.RoleAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := ptype.BindContract(&p.PermOrg, func() (interface{}, error) { return binding.NewOrgManager(p.PermConfig.OrgAddress, p.EthClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *Contract) initSession() {
	auth := bind.NewKeyedTransactor(p.Key)
	p.PermInterfSession = &binding.PermInterfaceSession{
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

	p.permOrgSession = &binding.OrgManagerSession{
		Contract: p.PermOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSession = &binding.NodeManagerSession{
		Contract: p.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSession = &binding.RoleManagerSession{
		Contract: p.PermRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSession = &binding.AcctManagerSession{
		Contract: p.PermAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

// checks if the passed URL is no nil and then calls GetNodeDetails
func (p *Contract) getNodeDetails(url string) (string, string, uint16, uint16, error) {
	if len(url) > 0 {
		return ptype.GetNodeDetails(url, p.IsRaft, p.UseDns)
	}

	return "", "", uint16(0), uint16(0), nil
}

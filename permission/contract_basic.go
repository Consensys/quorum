package permission

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	binding "github.com/ethereum/go-ethereum/permission/basic/bind"
)

type BasicContract struct {
	ethClnt    bind.ContractBackend
	key        *ecdsa.PrivateKey
	permConfig *types.PermissionConfig

	//binding contracts
	permUpgr   *binding.PermUpgr
	permInterf *binding.PermInterface
	permNode   *binding.NodeManager
	permAcct   *binding.AcctManager
	permRole   *binding.RoleManager
	permOrg    *binding.OrgManager
	//sessions
	permInterfSession *binding.PermInterfaceSession
	permOrgSession    *binding.OrgManagerSession
	permNodeSession   *binding.NodeManagerSession
	permRoleSession   *binding.RoleManagerSession
	permAcctSession   *binding.AcctManagerSession
}

func (p *BasicContract) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return p.permInterfSession.RemoveRole(_roleId, _orgId)
}
func (p *BasicContract) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	return p.permInterfSession.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}
func (p *BasicContract) ConnectionAllowedImpl(_enodeId string, _ip string, _port uint16, _raftport uint16) (bool, error) {
	return false, fmt.Errorf("not implemented for binding contract")
}
func (p *BasicContract) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {
	return false, fmt.Errorf("not implemented for binding contract")
}
func (p *BasicContract) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return p.permInterfSession.AssignAccountRole(_account, _orgId, _roleId)
}
func (p *BasicContract) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateAccountStatus(_orgId, _account, _action)
}
func (p *BasicContract) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.ApproveBlacklistedNodeRecovery(_orgId, _url)
}
func (p *BasicContract) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.StartBlacklistedNodeRecovery(_orgId, _url)
}
func (p *BasicContract) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.StartBlacklistedAccountRecovery(_orgId, _account)
}
func (p *BasicContract) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveBlacklistedAccountRecovery(_orgId, _account)
}
func (p *BasicContract) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return p.permInterfSession.GetPendingOp(_orgId)
}

func (p *BasicContract) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveAdminRole(_orgId, _account)
}

func (p *BasicContract) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	return p.permInterfSession.AssignAdminRole(_orgId, _account, _roleId)
}

func (p *BasicContract) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddNode(_orgId, _url)
}

func (p *BasicContract) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateNodeStatus(_orgId, _url, _action)

}

func (p *BasicContract) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrgStatus(_orgId, _action)
}

func (p *BasicContract) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateOrgStatus(_orgId, _action)
}

func (p *BasicContract) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrg(_orgId, _url, _account)
}

func (p *BasicContract) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddSubOrg(_pOrgId, _orgId, _url)
}

func (p *BasicContract) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddOrg(_orgId, _url, _account)
}

func (p *BasicContract) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	return p.permAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (p *BasicContract) GetNumberOfAccounts() (*big.Int, error) {
	return p.permAcctSession.GetNumberOfAccounts()
}

func (p *BasicContract) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (p *BasicContract) GetNumberOfRoles() (*big.Int, error) {
	return p.permRoleSession.GetNumberOfRoles()
}

func (p *BasicContract) GetNumberOfOrgs() (*big.Int, error) {
	return p.permOrgSession.GetNumberOfOrgs()
}

func (p *BasicContract) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return p.permInterfSession.UpdateNetworkBootStatus()
}

func (p *BasicContract) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminAccount(_acct)
}

func (p *BasicContract) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminNode(types.GetNodeUrl(_enodeId, _ip[:], _port, _raftport))
}

func (p *BasicContract) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return p.permInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (p *BasicContract) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.Init(_breadth, _depth)
}

func (p *BasicContract) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return p.permAcctSession.GetAccountDetails(_account)
}

func (p *BasicContract) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	return r.OrgId, r.EnodeId, r.NodeStatus, err
}

func (p *BasicContract) GetNumberOfNodes() (*big.Int, error) {
	return p.permNodeSession.GetNumberOfNodes()
}

func (p *BasicContract) GetNodeDetails(enodeId string) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetails(enodeId)
	return r.OrgId, r.EnodeId, r.NodeStatus, err
}

func (p *BasicContract) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (p *BasicContract) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return p.permOrgSession.GetSubOrgIndexes(_orgId)
}

func (p *BasicContract) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgInfo(_orgIndex)
}

func (p *BasicContract) GetNetworkBootStatus() (bool, error) {
	return p.permInterfSession.GetNetworkBootStatus()
}

func (p *BasicContract) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *BasicContract) AfterStart() error {
	log.Debug("permission service: binding contracts")
	err := p.basicBindContract()
	if err != nil {
		return err
	}
	p.initSession()
	return nil
}

func (p *BasicContract) basicBindContract() error {
	if err := bindContract(&p.permUpgr, func() (interface{}, error) { return binding.NewPermUpgr(p.permConfig.UpgrdAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permInterf, func() (interface{}, error) { return binding.NewPermInterface(p.permConfig.InterfAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permAcct, func() (interface{}, error) { return binding.NewAcctManager(p.permConfig.AccountAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permNode, func() (interface{}, error) { return binding.NewNodeManager(p.permConfig.NodeAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permRole, func() (interface{}, error) { return binding.NewRoleManager(p.permConfig.RoleAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permOrg, func() (interface{}, error) { return binding.NewOrgManager(p.permConfig.OrgAddress, p.ethClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *BasicContract) initSession() {
	auth := bind.NewKeyedTransactor(p.key)
	p.permInterfSession = &binding.PermInterfaceSession{
		Contract: p.permInterf,
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
		Contract: p.permOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSession = &binding.NodeManagerSession{
		Contract: p.permNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSession = &binding.RoleManagerSession{
		Contract: p.permRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSession = &binding.AcctManagerSession{
		Contract: p.permAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

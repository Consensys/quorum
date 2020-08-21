package permission

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	basic "github.com/ethereum/go-ethereum/permission/basic/bind"
)

type PermissionContractBasic struct {
	ethClnt    bind.ContractBackend
	key        *ecdsa.PrivateKey
	permConfig *types.PermissionConfig

	//basic contracts
	permUpgr   *basic.PermUpgr
	permInterf *basic.PermInterface
	permNode   *basic.NodeManager
	permAcct   *basic.AcctManager
	permRole   *basic.RoleManager
	permOrg    *basic.OrgManager
	//sessions
	permInterfSession *basic.PermInterfaceSession
	permOrgSession    *basic.OrgManagerSession
	permNodeSession   *basic.NodeManagerSession
	permRoleSession   *basic.RoleManagerSession
	permAcctSession   *basic.AcctManagerSession
}

func (p *PermissionContractBasic) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return p.permInterfSession.RemoveRole(_roleId, _orgId)
}
func (p *PermissionContractBasic) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	return p.permInterfSession.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}
func (p *PermissionContractBasic) ConnectionAllowedImpl(_enodeId string, _ip string, _port uint16, _raftport uint16) (bool, error) {
	return false, fmt.Errorf("not implemented for basic contract")
}
func (p *PermissionContractBasic) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {
	return false, fmt.Errorf("not implemented for basic contract")
}
func (p *PermissionContractBasic) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return p.permInterfSession.AssignAccountRole(_account, _orgId, _roleId)
}
func (p *PermissionContractBasic) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateAccountStatus(_orgId, _account, _action)
}
func (p *PermissionContractBasic) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.ApproveBlacklistedNodeRecovery(_orgId, _url)
}
func (p *PermissionContractBasic) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.StartBlacklistedNodeRecovery(_orgId, _url)
}
func (p *PermissionContractBasic) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.StartBlacklistedAccountRecovery(_orgId, _account)
}
func (p *PermissionContractBasic) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveBlacklistedAccountRecovery(_orgId, _account)
}
func (p *PermissionContractBasic) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return p.permInterfSession.GetPendingOp(_orgId)
}

func (p *PermissionContractBasic) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveAdminRole(_orgId, _account)
}

func (p *PermissionContractBasic) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	return p.permInterfSession.AssignAdminRole(_orgId, _account, _roleId)
}

func (p *PermissionContractBasic) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddNode(_orgId, _url)
}

func (p *PermissionContractBasic) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateNodeStatus(_orgId, _url, _action)

}

func (p *PermissionContractBasic) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrgStatus(_orgId, _action)
}

func (p *PermissionContractBasic) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateOrgStatus(_orgId, _action)
}

func (p *PermissionContractBasic) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrg(_orgId, _url, _account)
}

func (p *PermissionContractBasic) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddSubOrg(_pOrgId, _orgId, _url)
}

func (p *PermissionContractBasic) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddOrg(_orgId, _url, _account)
}

func (p *PermissionContractBasic) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	return p.permAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (p *PermissionContractBasic) GetNumberOfAccounts() (*big.Int, error) {
	return p.permAcctSession.GetNumberOfAccounts()
}

func (p *PermissionContractBasic) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (p *PermissionContractBasic) GetNumberOfRoles() (*big.Int, error) {
	return p.permRoleSession.GetNumberOfRoles()
}

func (p *PermissionContractBasic) GetNumberOfOrgs() (*big.Int, error) {
	return p.permOrgSession.GetNumberOfOrgs()
}

func (p *PermissionContractBasic) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return p.permInterfSession.UpdateNetworkBootStatus()
}

func (p *PermissionContractBasic) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminAccount(_acct)
}

func (p *PermissionContractBasic) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminNode(types.GetNodeUrl(_enodeId, _ip[:], _port, _raftport))
}

func (p *PermissionContractBasic) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return p.permInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (p *PermissionContractBasic) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.Init(_breadth, _depth)
}

func (p *PermissionContractBasic) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return p.permAcctSession.GetAccountDetails(_account)
}

func (p *PermissionContractBasic) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	return r.OrgId, r.EnodeId, r.NodeStatus, err
}

func (p *PermissionContractBasic) GetNumberOfNodes() (*big.Int, error) {
	return p.permNodeSession.GetNumberOfNodes()
}

func (p *PermissionContractBasic) GetNodeDetails(enodeId string) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetails(enodeId)
	return r.OrgId, r.EnodeId, r.NodeStatus, err
}

func (p *PermissionContractBasic) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (p *PermissionContractBasic) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return p.permOrgSession.GetSubOrgIndexes(_orgId)
}

func (p *PermissionContractBasic) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgInfo(_orgIndex)
}

func (p *PermissionContractBasic) GetNetworkBootStatus() (bool, error) {
	return p.permInterfSession.GetNetworkBootStatus()
}

func (p *PermissionContractBasic) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *PermissionContractBasic) AfterStart() error {
	log.Debug("permission service: binding contracts")
	err := p.basicBindContract()
	if err != nil {
		return err
	}
	p.initSession()
	return nil
}

func (p *PermissionContractBasic) basicBindContract() error {
	if err := bindContract(&p.permUpgr, func() (interface{}, error) { return basic.NewPermUpgr(p.permConfig.UpgrdAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permInterf, func() (interface{}, error) { return basic.NewPermInterface(p.permConfig.InterfAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permAcct, func() (interface{}, error) { return basic.NewAcctManager(p.permConfig.AccountAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permNode, func() (interface{}, error) { return basic.NewNodeManager(p.permConfig.NodeAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permRole, func() (interface{}, error) { return basic.NewRoleManager(p.permConfig.RoleAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permOrg, func() (interface{}, error) { return basic.NewOrgManager(p.permConfig.OrgAddress, p.ethClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *PermissionContractBasic) initSession() {
	auth := bind.NewKeyedTransactor(p.key)
	p.permInterfSession = &basic.PermInterfaceSession{
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

	p.permOrgSession = &basic.OrgManagerSession{
		Contract: p.permOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSession = &basic.NodeManagerSession{
		Contract: p.permNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSession = &basic.RoleManagerSession{
		Contract: p.permRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSession = &basic.AcctManagerSession{
		Contract: p.permAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

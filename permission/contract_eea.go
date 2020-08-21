package permission

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	binding "github.com/ethereum/go-ethereum/permission/eea/bind"
)

type EeaContract struct {
	ethClnt    bind.ContractBackend
	key        *ecdsa.PrivateKey
	permConfig *types.PermissionConfig

	//binding contracts
	permUpgr   *binding.EeaPermUpgr
	permInterf *binding.EeaPermInterface
	permNode   *binding.EeaNodeManager
	permAcct   *binding.EeaAcctManager
	permRole   *binding.EeaRoleManager
	permOrg    *binding.EeaOrgManager
	//sessions
	permInterfSession *binding.EeaPermInterfaceSession
	permOrgSession    *binding.EeaOrgManagerSession
	permNodeSession   *binding.EeaNodeManagerSession
	permRoleSession   *binding.EeaRoleManagerSession
	permAcctSession   *binding.EeaAcctManagerSession
}

func (p *EeaContract) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {

	return p.permInterfSession.RemoveRole(_roleId, _orgId)
}

func (p *EeaContract) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {

	return p.permInterfSession.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}

func (p *EeaContract) ConnectionAllowedImpl(_enodeId string, _ip string, _port uint16, _raftport uint16) (bool, error) {

	return p.permInterfSession.ConnectionAllowedImpl(_enodeId, _ip, _port, _raftport)
}

func (p *EeaContract) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {

	return p.permInterfSession.TransactionAllowed(_srcaccount, _tgtaccount)
}

func (p *EeaContract) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {

	return p.permInterfSession.AssignAccountRole(_account, _orgId, _roleId)
}

func (p *EeaContract) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {

	return p.permInterfSession.UpdateAccountStatus(_orgId, _account, _action)
}

func (p *EeaContract) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.permInterfSession.ApproveBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *EeaContract) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.permInterfSession.StartBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *EeaContract) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.StartBlacklistedAccountRecovery(_orgId, _account)
}

func (p *EeaContract) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveBlacklistedAccountRecovery(_orgId, _account)
}

func (p *EeaContract) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return p.permInterfSession.GetPendingOp(_orgId)
}

func (p *EeaContract) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveAdminRole(_orgId, _account)
}

func (p *EeaContract) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {

	return p.permInterfSession.AssignAdminRole(_orgId, _account, _roleId)
}

func (p *EeaContract) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.permInterfSession.AddNode(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *EeaContract) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string, _action *big.Int) (*types.Transaction, error) {

	return p.permInterfSession.UpdateNodeStatus(_orgId, _enodeId, _ip, _port, _raftport, _action)

}

func (p *EeaContract) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrgStatus(_orgId, _action)
}

func (p *EeaContract) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateOrgStatus(_orgId, _action)
}

func (p *EeaContract) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}

func (p *EeaContract) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddSubOrg(_pOrgId, _orgId, _enodeId, _ip, _port, _raftport)
}

func (p *EeaContract) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {

	return p.permInterfSession.AddOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}

func (p *EeaContract) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {

	return p.permAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (p *EeaContract) GetNumberOfAccounts() (*big.Int, error) {

	return p.permAcctSession.GetNumberOfAccounts()
}

func (p *EeaContract) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (p *EeaContract) GetNumberOfRoles() (*big.Int, error) {
	return p.permRoleSession.GetNumberOfRoles()
}

func (p *EeaContract) GetNumberOfOrgs() (*big.Int, error) {
	return p.permOrgSession.GetNumberOfOrgs()
}

func (p *EeaContract) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return p.permInterfSession.UpdateNetworkBootStatus()
}

func (p *EeaContract) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminAccount(_acct)
}

func (p *EeaContract) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminNode(_enodeId, _ip, _port, _raftport)
}

func (p *EeaContract) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return p.permInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (p *EeaContract) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.Init(_breadth, _depth)
}

func (p *EeaContract) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return p.permAcctSession.GetAccountDetails(_account)
}

func (p *EeaContract) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, types.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport), r.NodeStatus, err
}

func (p *EeaContract) GetNumberOfNodes() (*big.Int, error) {
	return p.permNodeSession.GetNumberOfNodes()
}

func (p *EeaContract) GetNodeDetails(enodeId string) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetails(enodeId)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, types.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport), r.NodeStatus, err
}

func (p *EeaContract) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (p *EeaContract) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return p.permOrgSession.GetSubOrgIndexes(_orgId)
}

func (p *EeaContract) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgInfo(_orgIndex)
}

func (p *EeaContract) GetNetworkBootStatus() (bool, error) {
	return p.permInterfSession.GetNetworkBootStatus()
}

func (p *EeaContract) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *EeaContract) AfterStart() error {
	log.Debug("permission service: binding contracts")

	err := p.eeaBindContract()
	if err != nil {
		return err
	}

	p.initSession()
	return nil
}

func (p *EeaContract) eeaBindContract() error {
	if err := bindContract(&p.permUpgr, func() (interface{}, error) { return binding.NewEeaPermUpgr(p.permConfig.UpgrdAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permInterf, func() (interface{}, error) { return binding.NewEeaPermInterface(p.permConfig.InterfAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permAcct, func() (interface{}, error) { return binding.NewEeaAcctManager(p.permConfig.AccountAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permNode, func() (interface{}, error) { return binding.NewEeaNodeManager(p.permConfig.NodeAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permRole, func() (interface{}, error) { return binding.NewEeaRoleManager(p.permConfig.RoleAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permOrg, func() (interface{}, error) { return binding.NewEeaOrgManager(p.permConfig.OrgAddress, p.ethClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *EeaContract) initSession() {
	auth := bind.NewKeyedTransactor(p.key)
	p.permInterfSession = &binding.EeaPermInterfaceSession{
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

	p.permOrgSession = &binding.EeaOrgManagerSession{
		Contract: p.permOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSession = &binding.EeaNodeManagerSession{
		Contract: p.permNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSession = &binding.EeaRoleManagerSession{
		Contract: p.permRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSession = &binding.EeaAcctManagerSession{
		Contract: p.permAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

package permission

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/permission/bind/eea"
)

type PermissionContractEea struct {
	ethClnt    bind.ContractBackend
	key        *ecdsa.PrivateKey
	permConfig *types.PermissionConfig

	//eea contracts
	permUpgr   *eea.EeaPermUpgr
	permInterf *eea.EeaPermInterface
	permNode   *eea.EeaNodeManager
	permAcct   *eea.EeaAcctManager
	permRole   *eea.EeaRoleManager
	permOrg    *eea.EeaOrgManager
	//sessions
	permInterfSession *eea.EeaPermInterfaceSession
	permOrgSession    *eea.EeaOrgManagerSession
	permNodeSession   *eea.EeaNodeManagerSession
	permRoleSession   *eea.EeaRoleManagerSession
	permAcctSession   *eea.EeaAcctManagerSession
}

func (p *PermissionContractEea) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {

	return p.permInterfSession.RemoveRole(_roleId, _orgId)
}

func (p *PermissionContractEea) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {

	return p.permInterfSession.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}

func (p *PermissionContractEea) ConnectionAllowedImpl(_enodeId string, _ip string, _port uint16, _raftport uint16) (bool, error) {

	return p.permInterfSession.ConnectionAllowedImpl(_enodeId, _ip, _port, _raftport)
}

func (p *PermissionContractEea) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {

	return p.permInterfSession.TransactionAllowed(_srcaccount, _tgtaccount)
}

func (p *PermissionContractEea) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {

	return p.permInterfSession.AssignAccountRole(_account, _orgId, _roleId)
}

func (p *PermissionContractEea) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {

	return p.permInterfSession.UpdateAccountStatus(_orgId, _account, _action)
}

func (p *PermissionContractEea) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.permInterfSession.ApproveBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *PermissionContractEea) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.permInterfSession.StartBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *PermissionContractEea) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.StartBlacklistedAccountRecovery(_orgId, _account)
}

func (p *PermissionContractEea) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveBlacklistedAccountRecovery(_orgId, _account)
}

func (p *PermissionContractEea) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return p.permInterfSession.GetPendingOp(_orgId)
}

func (p *PermissionContractEea) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return p.permInterfSession.ApproveAdminRole(_orgId, _account)
}

func (p *PermissionContractEea) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {

	return p.permInterfSession.AssignAdminRole(_orgId, _account, _roleId)
}

func (p *PermissionContractEea) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {

	return p.permInterfSession.AddNode(_orgId, _enodeId, _ip, _port, _raftport)
}

func (p *PermissionContractEea) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string, _action *big.Int) (*types.Transaction, error) {

	return p.permInterfSession.UpdateNodeStatus(_orgId, _enodeId, _ip, _port, _raftport, _action)

}

func (p *PermissionContractEea) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrgStatus(_orgId, _action)
}

func (p *PermissionContractEea) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.UpdateOrgStatus(_orgId, _action)
}

func (p *PermissionContractEea) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	return p.permInterfSession.ApproveOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}

func (p *PermissionContractEea) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	return p.permInterfSession.AddSubOrg(_pOrgId, _orgId, _enodeId, _ip, _port, _raftport)
}

func (p *PermissionContractEea) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {

	return p.permInterfSession.AddOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}

func (p *PermissionContractEea) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {

	return p.permAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (p *PermissionContractEea) GetNumberOfAccounts() (*big.Int, error) {

	return p.permAcctSession.GetNumberOfAccounts()
}

func (p *PermissionContractEea) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (p *PermissionContractEea) GetNumberOfRoles() (*big.Int, error) {
	return p.permRoleSession.GetNumberOfRoles()
}

func (p *PermissionContractEea) GetNumberOfOrgs() (*big.Int, error) {
	return p.permOrgSession.GetNumberOfOrgs()
}

func (p *PermissionContractEea) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return p.permInterfSession.UpdateNetworkBootStatus()
}

func (p *PermissionContractEea) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminAccount(_acct)
}

func (p *PermissionContractEea) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return p.permInterfSession.AddAdminNode(_enodeId, _ip, _port, _raftport)
}

func (p *PermissionContractEea) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return p.permInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (p *PermissionContractEea) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return p.permInterfSession.Init(_breadth, _depth)
}

func (p *PermissionContractEea) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return p.permAcctSession.GetAccountDetails(_account)
}

func (p *PermissionContractEea) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, types.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport), r.NodeStatus, err
}

func (p *PermissionContractEea) GetNumberOfNodes() (*big.Int, error) {
	return p.permNodeSession.GetNumberOfNodes()
}

func (p *PermissionContractEea) GetNodeDetails(enodeId string) (string, string, *big.Int, error) {
	r, err := p.permNodeSession.GetNodeDetails(enodeId)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, types.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport), r.NodeStatus, err
}

func (p *PermissionContractEea) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return p.permRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (p *PermissionContractEea) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return p.permOrgSession.GetSubOrgIndexes(_orgId)
}

func (p *PermissionContractEea) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgInfo(_orgIndex)
}

func (p *PermissionContractEea) GetNetworkBootStatus() (bool, error) {
	return p.permInterfSession.GetNetworkBootStatus()
}

func (p *PermissionContractEea) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return p.permOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *PermissionContractEea) AfterStart() error {
	log.Debug("permission service: binding contracts")

	err := p.eeaBindContract()
	if err != nil {
		return err
	}

	p.initSession()
	return nil
}

func (p *PermissionContractEea) eeaBindContract() error {
	if err := bindContract(&p.permUpgr, func() (interface{}, error) { return eea.NewEeaPermUpgr(p.permConfig.UpgrdAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permInterf, func() (interface{}, error) { return eea.NewEeaPermInterface(p.permConfig.InterfAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permAcct, func() (interface{}, error) { return eea.NewEeaAcctManager(p.permConfig.AccountAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permNode, func() (interface{}, error) { return eea.NewEeaNodeManager(p.permConfig.NodeAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permRole, func() (interface{}, error) { return eea.NewEeaRoleManager(p.permConfig.RoleAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permOrg, func() (interface{}, error) { return eea.NewEeaOrgManager(p.permConfig.OrgAddress, p.ethClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *PermissionContractEea) initSession() {
	auth := bind.NewKeyedTransactor(p.key)
	p.permInterfSession = &eea.EeaPermInterfaceSession{
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

	p.permOrgSession = &eea.EeaOrgManagerSession{
		Contract: p.permOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSession = &eea.EeaNodeManagerSession{
		Contract: p.permNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSession = &eea.EeaRoleManagerSession{
		Contract: p.permRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSession = &eea.EeaAcctManagerSession{
		Contract: p.permAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

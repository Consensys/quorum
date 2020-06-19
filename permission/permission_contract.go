package permission

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	pbind "github.com/ethereum/go-ethereum/permission/bind"
	"math/big"
)

type PermissionContractService struct {
	ethClnt    bind.ContractBackend
	key        *ecdsa.PrivateKey
	permConfig *types.PermissionConfig
	eeaFlag    bool

	//basic contracts
	permUpgr   *pbind.PermUpgr
	permInterf *pbind.PermInterface
	permNode   *pbind.NodeManager
	permAcct   *pbind.AcctManager
	permRole   *pbind.RoleManager
	permOrg    *pbind.OrgManager
	//sessions
	permInterfSession *pbind.PermInterfaceSession
	permOrgSession    *pbind.OrgManagerSession
	permNodeSession   *pbind.NodeManagerSession
	permRoleSession   *pbind.RoleManagerSession
	permAcctSession   *pbind.AcctManagerSession

	//eea contracts
	permUpgrE   *pbind.EeaPermUpgr
	permInterfE *pbind.EeaPermInterface
	permNodeE   *pbind.EeaNodeManager
	permAcctE   *pbind.EeaAcctManager
	permRoleE   *pbind.EeaRoleManager
	permOrgE    *pbind.EeaOrgManager
	//sessions
	permInterfSessionE *pbind.EeaPermInterfaceSession
	permOrgSessionE    *pbind.EeaOrgManagerSession
	permNodeSessionE   *pbind.EeaNodeManagerSession
	permRoleSessionE   *pbind.EeaRoleManagerSession
	permAcctSessionE   *pbind.EeaAcctManagerSession
}

func (p *PermissionContractService) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.RemoveRole(_roleId, _orgId)
	}
	return p.permInterfSession.RemoveRole(_roleId, _orgId)
}
func (p *PermissionContractService) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
	}
	return p.permInterfSession.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}
func (p *PermissionContractService) ConnectionAllowedImpl(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (bool, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.ConnectionAllowedImpl(_enodeId, _ip, _port, _raftport)
	}
	return false, fmt.Errorf("not implemented for basic contract")
}
func (p *PermissionContractService) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.TransactionAllowed(_srcaccount, _tgtaccount)
	}
	return false, fmt.Errorf("not implemented for basic contract")
}
func (p *PermissionContractService) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AssignAccountRole(_account, _orgId, _roleId)
	}
	return p.permInterfSession.AssignAccountRole(_account, _orgId, _roleId)
}
func (p *PermissionContractService) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.UpdateAccountStatus(_orgId, _account, _action)
	}
	return p.permInterfSession.UpdateAccountStatus(_orgId, _account, _action)
}
func (p *PermissionContractService) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.ApproveBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.permInterfSession.ApproveBlacklistedNodeRecovery(_orgId, _url)
}
func (p *PermissionContractService) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.StartBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.permInterfSession.StartBlacklistedNodeRecovery(_orgId, _url)
}
func (p *PermissionContractService) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.StartBlacklistedAccountRecovery(_orgId, _account)
	}
	return p.permInterfSession.StartBlacklistedAccountRecovery(_orgId, _account)
}
func (p *PermissionContractService) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.ApproveBlacklistedAccountRecovery(_orgId, _account)
	}
	return p.permInterfSession.ApproveBlacklistedAccountRecovery(_orgId, _account)
}
func (p *PermissionContractService) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.GetPendingOp(_orgId)
	}
	return p.permInterfSession.GetPendingOp(_orgId)
}

func (p *PermissionContractService) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.ApproveAdminRole(_orgId, _account)
	}
	return p.permInterfSession.ApproveAdminRole(_orgId, _account)
}

func (p *PermissionContractService) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AssignAdminRole(_orgId, _account, _roleId)
	}
	return p.permInterfSession.AssignAdminRole(_orgId, _account, _roleId)
}

func (p *PermissionContractService) AddNode(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AddNode(_orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.permInterfSession.AddNode(_orgId, _url)
}

func (p *PermissionContractService) UpdateNodeStatus(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string, _action *big.Int) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.UpdateNodeStatus(_orgId, _enodeId, _ip, _port, _raftport, _action)
	}
	return p.permInterfSession.UpdateNodeStatus(_orgId, _url, _action)

}

func (p *PermissionContractService) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.ApproveOrgStatus(_orgId, _action)
	}
	return p.permInterfSession.ApproveOrgStatus(_orgId, _action)
}

func (p *PermissionContractService) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.UpdateOrgStatus(_orgId, _action)
	}
	return p.permInterfSession.UpdateOrgStatus(_orgId, _action)
}

func (p *PermissionContractService) ApproveOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.ApproveOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
	}
	return p.permInterfSession.ApproveOrg(_orgId, _url, _account)
}

func (p *PermissionContractService) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AddSubOrg(_pOrgId, _orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.permInterfSession.AddSubOrg(_pOrgId, _orgId, _url)
}

func (p *PermissionContractService) AddOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AddOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
	}
	return p.permInterfSession.AddOrg(_orgId, _url, _account)
}

func (p *PermissionContractService) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	if p.eeaFlag {
		return p.permAcctSessionE.GetAccountDetailsFromIndex(_aIndex)
	}
	return p.permAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (p *PermissionContractService) GetNumberOfAccounts() (*big.Int, error) {
	if p.eeaFlag {
		return p.permAcctSessionE.GetNumberOfAccounts()
	}
	return p.permAcctSession.GetNumberOfAccounts()
}

func (p *PermissionContractService) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	if p.eeaFlag {
		return p.permRoleSessionE.GetRoleDetailsFromIndex(_rIndex)
	}
	return p.permRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (p *PermissionContractService) GetNumberOfRoles() (*big.Int, error) {
	if p.eeaFlag {
		return p.permRoleSessionE.GetNumberOfRoles()
	}
	return p.permRoleSession.GetNumberOfRoles()
}

func (p *PermissionContractService) GetNumberOfOrgs() (*big.Int, error) {
	if p.eeaFlag {
		return p.permOrgSessionE.GetNumberOfOrgs()
	}
	return p.permOrgSession.GetNumberOfOrgs()
}

func (p *PermissionContractService) UpdateNetworkBootStatus() (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.UpdateNetworkBootStatus()
	}
	return p.permInterfSession.UpdateNetworkBootStatus()
}

func (p *PermissionContractService) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AddAdminAccount(_acct)
	}
	return p.permInterfSession.AddAdminAccount(_acct)
}

func (p *PermissionContractService) AddAdminNode(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.AddAdminNode(_enodeId, _ip, _port, _raftport)
	}
	return p.permInterfSession.AddAdminNode(types.GetNodeUrl(_enodeId, string(_ip[:]), _port, _raftport))
}

func (p *PermissionContractService) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
	}
	return p.permInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (p *PermissionContractService) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.Init(_breadth, _depth)
	}
	return p.permInterfSession.Init(_breadth, _depth)
}

func (p *PermissionContractService) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	if p.eeaFlag {
		return p.permAcctSessionE.GetAccountDetails(_account)
	}
	return p.permAcctSession.GetAccountDetails(_account)
}

func (p *PermissionContractService) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error) {

	if p.eeaFlag {
		r, err := p.permNodeSessionE.GetNodeDetailsFromIndex(_nodeIndex)
		if err != nil {
			return "", "", big.NewInt(0), err
		}
		return r.OrgId, types.GetNodeUrl(r.EnodeId, string(r.Ip[:]), r.Port, r.Raftport), r.NodeStatus, err
	}
	r, err := p.permNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	return r.OrgId, r.EnodeId, r.NodeStatus, err
}

func (p *PermissionContractService) GetNumberOfNodes() (*big.Int, error) {
	if p.eeaFlag {
		return p.permNodeSessionE.GetNumberOfNodes()
	}
	return p.permNodeSession.GetNumberOfNodes()
}

func (p *PermissionContractService) GetNodeDetails(enodeId string) (string, string, *big.Int, error) {
	if p.eeaFlag {
		r, err := p.permNodeSessionE.GetNodeDetails(enodeId)
		if err != nil {
			return "", "", big.NewInt(0), err
		}
		return r.OrgId, types.GetNodeUrl(r.EnodeId, string(r.Ip[:]), r.Port, r.Raftport), r.NodeStatus, err
	}
	r, err := p.permNodeSession.GetNodeDetails(enodeId)
	return r.OrgId, r.EnodeId, r.NodeStatus, err
}

func (p *PermissionContractService) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	if p.eeaFlag {
		return p.permRoleSessionE.GetRoleDetails(_roleId, _orgId)
	}
	return p.permRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (p *PermissionContractService) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	if p.eeaFlag {
		return p.permOrgSessionE.GetSubOrgIndexes(_orgId)
	}
	return p.permOrgSession.GetSubOrgIndexes(_orgId)
}

func (p *PermissionContractService) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	if p.eeaFlag {
		return p.permOrgSessionE.GetOrgInfo(_orgIndex)
	}
	return p.permOrgSession.GetOrgInfo(_orgIndex)
}

func (p *PermissionContractService) GetNetworkBootStatus() (bool, error) {
	if p.eeaFlag {
		return p.permInterfSessionE.GetNetworkBootStatus()
	}
	return p.permInterfSession.GetNetworkBootStatus()
}

func (p *PermissionContractService) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	if p.eeaFlag {
		return p.permOrgSessionE.GetOrgDetails(_orgId)
	}
	return p.permOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *PermissionContractService) AfterStart() error {
	log.Debug("permission service: binding contracts")
	if p.eeaFlag {
		err := p.eeaBindContract()
		if err != nil {
			return err
		}
	} else {
		err := p.basicBindContract()
		if err != nil {
			return err
		}
	}
	p.InitSession()
	return nil
}

func (p *PermissionContractService) basicBindContract() error {
	if err := bindContract(&p.permUpgr, func() (interface{}, error) { return pbind.NewPermUpgr(p.permConfig.UpgrdAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permInterf, func() (interface{}, error) { return pbind.NewPermInterface(p.permConfig.InterfAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permAcct, func() (interface{}, error) { return pbind.NewAcctManager(p.permConfig.AccountAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permNode, func() (interface{}, error) { return pbind.NewNodeManager(p.permConfig.NodeAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permRole, func() (interface{}, error) { return pbind.NewRoleManager(p.permConfig.RoleAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permOrg, func() (interface{}, error) { return pbind.NewOrgManager(p.permConfig.OrgAddress, p.ethClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *PermissionContractService) eeaBindContract() error {
	if err := bindContract(&p.permUpgrE, func() (interface{}, error) { return pbind.NewEeaPermUpgr(p.permConfig.UpgrdAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permInterfE, func() (interface{}, error) { return pbind.NewEeaPermInterface(p.permConfig.InterfAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permAcctE, func() (interface{}, error) { return pbind.NewEeaAcctManager(p.permConfig.AccountAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permNodeE, func() (interface{}, error) { return pbind.NewEeaNodeManager(p.permConfig.NodeAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permRoleE, func() (interface{}, error) { return pbind.NewEeaRoleManager(p.permConfig.RoleAddress, p.ethClnt) }); err != nil {
		return err
	}
	if err := bindContract(&p.permOrgE, func() (interface{}, error) { return pbind.NewEeaOrgManager(p.permConfig.OrgAddress, p.ethClnt) }); err != nil {
		return err
	}
	return nil
}

///   ------ old code //// -------

func (p *PermissionContractService) InitSession() {
	if p.eeaFlag {
		p.eeaSession()
	} else {
		p.basicSession()
	}
}

func (p *PermissionContractService) eeaSession() {
	auth := bind.NewKeyedTransactor(p.key)
	p.permInterfSessionE = &pbind.EeaPermInterfaceSession{
		Contract: p.permInterfE,
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

	p.permOrgSessionE = &pbind.EeaOrgManagerSession{
		Contract: p.permOrgE,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSessionE = &pbind.EeaNodeManagerSession{
		Contract: p.permNodeE,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSessionE = &pbind.EeaRoleManagerSession{
		Contract: p.permRoleE,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSessionE = &pbind.EeaAcctManagerSession{
		Contract: p.permAcctE,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}
func (p *PermissionContractService) basicSession() {
	auth := bind.NewKeyedTransactor(p.key)
	p.permInterfSession = &pbind.PermInterfaceSession{
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

	p.permOrgSession = &pbind.OrgManagerSession{
		Contract: p.permOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.permNodeSession = &pbind.NodeManagerSession{
		Contract: p.permNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.permRoleSession = &pbind.RoleManagerSession{
		Contract: p.permRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.permAcctSession = &pbind.AcctManagerSession{
		Contract: p.permAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

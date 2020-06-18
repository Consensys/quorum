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
	EthClnt    bind.ContractBackend
	Key        *ecdsa.PrivateKey
	PermConfig *types.PermissionConfig
	EeaFlag    bool

	//basic contracts
	PermUpgr   *pbind.PermUpgr
	PermInterf *pbind.PermInterface
	PermNode   *pbind.NodeManager
	PermAcct   *pbind.AcctManager
	PermRole   *pbind.RoleManager
	PermOrg    *pbind.OrgManager
	//sessions
	PermInterfSession *pbind.PermInterfaceSession
	PermOrgSession    *pbind.OrgManagerSession
	PermNodeSession   *pbind.NodeManagerSession
	PermRoleSession   *pbind.RoleManagerSession
	PermAcctSession   *pbind.AcctManagerSession

	//eea contracts
	PermUpgrE   *pbind.EeaPermUpgr
	PermInterfE *pbind.EeaPermInterface
	PermNodeE   *pbind.EeaNodeManager
	PermAcctE   *pbind.EeaAcctManager
	PermRoleE   *pbind.EeaRoleManager
	PermOrgE    *pbind.EeaOrgManager
	//sessions
	PermInterfSessionE *pbind.EeaPermInterfaceSession
	PermOrgSessionE    *pbind.EeaOrgManagerSession
	PermNodeSessionE   *pbind.EeaNodeManagerSession
	PermRoleSessionE   *pbind.EeaRoleManagerSession
	PermAcctSessionE   *pbind.EeaAcctManagerSession
}

func (p *PermissionContractService) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.RemoveRole(_roleId, _orgId)
	}
	return p.PermInterfSession.RemoveRole(_roleId, _orgId)
}
func (p *PermissionContractService) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
	}
	return p.PermInterfSession.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}
func (p *PermissionContractService) ConnectionAllowedImpl(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (bool, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.ConnectionAllowedImpl(_enodeId, _ip, _port, _raftport)
	}
	return false, fmt.Errorf("not implemented for basic contract")
}
func (p *PermissionContractService) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.TransactionAllowed(_srcaccount, _tgtaccount)
	}
	return false, fmt.Errorf("not implemented for basic contract")
}
func (p *PermissionContractService) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.AssignAccountRole(_account, _orgId, _roleId)
	}
	return p.PermInterfSession.AssignAccountRole(_account, _orgId, _roleId)
}
func (p *PermissionContractService) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.UpdateAccountStatus(_orgId, _account, _action)
	}
	return p.PermInterfSession.UpdateAccountStatus(_orgId, _account, _action)
}
func (p *PermissionContractService) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.ApproveBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.PermInterfSession.ApproveBlacklistedNodeRecovery(_orgId, _url)
}
func (p *PermissionContractService) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.EeaFlag{
		return p.PermInterfSessionE.StartBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.PermInterfSession.StartBlacklistedNodeRecovery(_orgId, _url)
}
func (p *PermissionContractService) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.StartBlacklistedAccountRecovery(_orgId, _account)
	}
	return p.PermInterfSession.StartBlacklistedAccountRecovery(_orgId, _account)
}
func (p *PermissionContractService) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.ApproveBlacklistedAccountRecovery(_orgId, _account)
	}
	return p.PermInterfSession.ApproveBlacklistedAccountRecovery(_orgId, _account)
}
func (p *PermissionContractService) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.GetPendingOp(_orgId)
	}
	return p.PermInterfSession.GetPendingOp(_orgId)
}


func (p *PermissionContractService) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	if p.EeaFlag{
		return p.PermInterfSessionE.ApproveAdminRole(_orgId, _account)
	}
	return p.PermInterfSession.ApproveAdminRole(_orgId, _account)
}

func (p *PermissionContractService) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	if p.EeaFlag{
		return p.PermInterfSessionE.AssignAdminRole(_orgId, _account, _roleId)
	}
	return p.PermInterfSession.AssignAdminRole(_orgId, _account, _roleId)
}

func (p *PermissionContractService) AddNode(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.EeaFlag{
		return p.PermInterfSessionE.AddNode(_orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.PermInterfSession.AddNode(_orgId, _url)
}

func (p *PermissionContractService) UpdateNodeStatus(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string, _action *big.Int) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.UpdateNodeStatus(_orgId, _enodeId, _ip, _port, _raftport, _action)
	}
	return p.PermInterfSession.UpdateNodeStatus(_orgId, _url, _action)

}

func (p *PermissionContractService) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	if p.EeaFlag{
		return p.PermInterfSessionE.ApproveOrgStatus(_orgId, _action)
	}
	return p.PermInterfSession.ApproveOrgStatus(_orgId, _action)
}

func (p *PermissionContractService) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	if p.EeaFlag{
		return p.PermInterfSessionE.UpdateOrgStatus(_orgId, _action)
	}
	return p.PermInterfSession.UpdateOrgStatus(_orgId, _action)
}

func (p *PermissionContractService) ApproveOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address, _url string) (*types.Transaction, error) {
	if p.EeaFlag{
		return p.PermInterfSessionE.ApproveOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
	}
	return p.PermInterfSession.ApproveOrg(_orgId, _url, _account)
}

func (p *PermissionContractService) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _url string) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.AddSubOrg(_pOrgId, _orgId, _enodeId, _ip, _port, _raftport)
	}
	return p.PermInterfSession.AddSubOrg(_pOrgId, _orgId, _url)
}

func (p *PermissionContractService) AddOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.AddOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
	}
	return p.PermInterfSession.AddOrg(_orgId, _enodeId, _account)
}

func (p *PermissionContractService) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	if p.EeaFlag {
		return p.PermAcctSessionE.GetAccountDetailsFromIndex(_aIndex)
	}
	return p.PermAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (p *PermissionContractService) GetNumberOfAccounts() (*big.Int, error) {
	if p.EeaFlag {
		return p.PermAcctSessionE.GetNumberOfAccounts()
	}
	return p.PermAcctSession.GetNumberOfAccounts()
}

func (p *PermissionContractService) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	if p.EeaFlag {
		return p.PermRoleSessionE.GetRoleDetailsFromIndex(_rIndex)
	}
	return p.PermRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (p *PermissionContractService) GetNumberOfRoles() (*big.Int, error) {
	if p.EeaFlag {
		return p.PermRoleSessionE.GetNumberOfRoles()
	}
	return p.PermRoleSession.GetNumberOfRoles()
}

func (p *PermissionContractService) GetNumberOfOrgs() (*big.Int, error) {
	if p.EeaFlag {
		return p.PermOrgSessionE.GetNumberOfOrgs()
	}
	return p.PermOrgSession.GetNumberOfOrgs()
}

func (p *PermissionContractService) UpdateNetworkBootStatus() (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.UpdateNetworkBootStatus()
	}
	return p.PermInterfSession.UpdateNetworkBootStatus()
}

func (p *PermissionContractService) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.AddAdminAccount(_acct)
	}
	return p.PermInterfSession.AddAdminAccount(_acct)
}

func (p *PermissionContractService) AddAdminNode(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.AddAdminNode(_enodeId, _ip, _port, _raftport)
	}
	// TODO Amal
	return p.PermInterfSession.AddAdminNode(_enodeId)
}

func (p *PermissionContractService) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
	}
	return p.PermInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (p *PermissionContractService) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.Init(_breadth, _depth)
	}
	return p.PermInterfSession.Init(_breadth, _depth)
}

func (p *PermissionContractService) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	if p.EeaFlag {
		return p.PermAcctSessionE.GetAccountDetails(_account)
	}
	return p.PermAcctSession.GetAccountDetails(_account)
}

// TODO : Amal
func (p *PermissionContractService) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	Ip         [32]byte
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}, error) {
	if p.EeaFlag {
		return p.PermNodeSessionE.GetNodeDetailsFromIndex(_nodeIndex)
	}
	ret := new(struct {
		OrgId      string
		EnodeId    string
		Ip         [32]byte
		Port       uint16
		Raftport   uint16
		NodeStatus *big.Int
	})
	r, err := p.PermNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	if err == nil {
		ret.EnodeId = r.EnodeId
		ret.OrgId = r.OrgId
		ret.NodeStatus = r.NodeStatus
	}
	return *ret, err
}

func (p *PermissionContractService) GetNumberOfNodes() (*big.Int, error) {
	if p.EeaFlag {
		return p.PermNodeSessionE.GetNumberOfNodes()
	}
	return p.PermNodeSession.GetNumberOfNodes()
}

// TODO Amal
func (p *PermissionContractService) GetNodeDetails(enodeId string) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	if p.EeaFlag {
		return p.PermNodeSessionE.GetNodeDetails(enodeId)
	} else {
		return p.PermNodeSession.GetNodeDetails(enodeId)
	}
}

func (p *PermissionContractService) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	if p.EeaFlag {
		return p.PermRoleSessionE.GetRoleDetails(_roleId, _orgId)
	}
	return p.PermRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (p *PermissionContractService) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	if p.EeaFlag {
		return p.PermOrgSessionE.GetSubOrgIndexes(_orgId)
	}
	return p.PermOrgSession.GetSubOrgIndexes(_orgId)
}

func (p *PermissionContractService) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	if p.EeaFlag {
		return p.PermOrgSessionE.GetOrgInfo(_orgIndex)
	}
	return p.PermOrgSession.GetOrgInfo(_orgIndex)
}

func (p *PermissionContractService) GetNetworkBootStatus() (bool, error) {
	if p.EeaFlag {
		return p.PermInterfSessionE.GetNetworkBootStatus()
	}
	return p.PermInterfSession.GetNetworkBootStatus()
}

func (p *PermissionContractService) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	if p.EeaFlag {
		return p.PermOrgSessionE.GetOrgDetails(_orgId)
	}
	return p.PermOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *PermissionContractService) AfterStart() error {
	log.Debug("permission service: binding contracts")
	if p.EeaFlag {
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
	if err := types.BindContract(&p.PermUpgr, func() (interface{}, error) { return pbind.NewPermUpgr(p.PermConfig.UpgrdAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermInterf, func() (interface{}, error) { return pbind.NewPermInterface(p.PermConfig.InterfAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermAcct, func() (interface{}, error) { return pbind.NewAcctManager(p.PermConfig.AccountAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermNode, func() (interface{}, error) { return pbind.NewNodeManager(p.PermConfig.NodeAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermRole, func() (interface{}, error) { return pbind.NewRoleManager(p.PermConfig.RoleAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermOrg, func() (interface{}, error) { return pbind.NewOrgManager(p.PermConfig.OrgAddress, p.EthClnt) }); err != nil {
		return err
	}
	return nil
}

func (p *PermissionContractService) eeaBindContract() error {
	if err := types.BindContract(&p.PermUpgrE, func() (interface{}, error) { return pbind.NewEeaPermUpgr(p.PermConfig.UpgrdAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermInterfE, func() (interface{}, error) { return pbind.NewEeaPermInterface(p.PermConfig.InterfAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermAcctE, func() (interface{}, error) { return pbind.NewEeaAcctManager(p.PermConfig.AccountAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermNodeE, func() (interface{}, error) { return pbind.NewEeaNodeManager(p.PermConfig.NodeAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermRoleE, func() (interface{}, error) { return pbind.NewEeaRoleManager(p.PermConfig.RoleAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermOrgE, func() (interface{}, error) { return pbind.NewEeaOrgManager(p.PermConfig.OrgAddress, p.EthClnt) }); err != nil {
		return err
	}
	return nil
}

///   ------ old code //// -------

func (p *PermissionContractService) InitSession() {
	if p.EeaFlag {
		p.eeaSession()
	} else {
		p.basicSession()
	}
}

func (p *PermissionContractService) eeaSession() {
	auth := bind.NewKeyedTransactor(p.Key)
	p.PermInterfSession = &pbind.PermInterfaceSession{
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

	p.PermOrgSession = &pbind.OrgManagerSession{
		Contract: p.PermOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.PermNodeSession = &pbind.NodeManagerSession{
		Contract: p.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.PermRoleSession = &pbind.RoleManagerSession{
		Contract: p.PermRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.PermAcctSession = &pbind.AcctManagerSession{
		Contract: p.PermAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}
func (p *PermissionContractService) basicSession() {
	auth := bind.NewKeyedTransactor(p.Key)
	p.PermInterfSession = &pbind.PermInterfaceSession{
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

	p.PermOrgSession = &pbind.OrgManagerSession{
		Contract: p.PermOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	p.PermNodeSession = &pbind.NodeManagerSession{
		Contract: p.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	p.PermRoleSession = &pbind.RoleManagerSession{
		Contract: p.PermRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	p.PermAcctSession = &pbind.AcctManagerSession{
		Contract: p.PermAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}



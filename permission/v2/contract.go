package v2

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/permission/core"
	ptype "github.com/ethereum/go-ethereum/permission/core/types"
	binding "github.com/ethereum/go-ethereum/permission/v2/bind"
)

// definitions for v2 permissions model which is aligned with eea specs

type PermissionModelV2 struct {
	ContractBackend   ptype.ContractBackend
	PermInterf        *binding.PermInterface
	PermInterfSession *binding.PermInterfaceSession
}

type Audit struct {
	Backend *PermissionModelV2
}

type Role struct {
	Backend *PermissionModelV2
}

type Account struct {
	Backend *PermissionModelV2
}

type Control struct {
	Backend *PermissionModelV2
}

type Org struct {
	Backend *PermissionModelV2
}

type Node struct {
	Backend *PermissionModelV2
}

type Init struct {
	Backend ptype.ContractBackend
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

func (a *Account) AssignAccountRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return a.Backend.PermInterfSession.AssignAccountRole(_args.AcctId, _args.OrgId, _args.RoleId)
}

func (a *Account) UpdateAccountStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	return a.Backend.PermInterfSession.UpdateAccountStatus(_args.OrgId, _args.AcctId, big.NewInt(int64(_args.Action)))
}

func (a *Account) StartBlacklistedAccountRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	return a.Backend.PermInterfSession.StartBlacklistedAccountRecovery(_args.OrgId, _args.AcctId)
}

func (a *Account) ApproveBlacklistedAccountRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	return a.Backend.PermInterfSession.ApproveBlacklistedAccountRecovery(_args.OrgId, _args.AcctId)
}

func (a *Account) ApproveAdminRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return a.Backend.PermInterfSession.ApproveAdminRole(_args.OrgId, _args.AcctId)
}

func (a *Account) AssignAdminRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return a.Backend.PermInterfSession.AssignAdminRole(_args.OrgId, _args.AcctId, _args.RoleId)
}

func (i *Init) GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	return i.permAcctSession.GetAccountDetailsFromIndex(_aIndex)
}

func (i *Init) GetNumberOfAccounts() (*big.Int, error) {
	return i.permAcctSession.GetNumberOfAccounts()
}

func (i *Init) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return i.permRoleSession.GetRoleDetailsFromIndex(_rIndex)
}

func (i *Init) GetNumberOfRoles() (*big.Int, error) {
	return i.permRoleSession.GetNumberOfRoles()
}

func (i *Init) GetNumberOfOrgs() (*big.Int, error) {
	return i.permOrgSession.GetNumberOfOrgs()
}

func (i *Init) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return i.PermInterfSession.UpdateNetworkBootStatus()
}

func (i *Init) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return i.PermInterfSession.AddAdminAccount(_acct)
}

func (i *Init) AddAdminNode(url string) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(url, i.Backend.IsRaft, i.Backend.UseDns)
	if err != nil {
		return nil, err
	}
	return i.PermInterfSession.AddAdminNode(enodeId, ip, port, raftPort)
}

func (i *Init) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return i.PermInterfSession.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}

func (i *Init) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return i.PermInterfSession.Init(_breadth, _depth)
}

func (i *Init) GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return i.permAcctSession.GetAccountDetails(_account)
}

func (i *Init) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error) {
	r, err := i.permNodeSession.GetNodeDetailsFromIndex(_nodeIndex)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, core.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport, i.Backend.IsRaft), r.NodeStatus, err
}

func (i *Init) GetNumberOfNodes() (*big.Int, error) {
	return i.permNodeSession.GetNumberOfNodes()
}

func (i *Init) GetNodeDetails(enodeId string) (string, string, *big.Int, error) {
	r, err := i.permNodeSession.GetNodeDetails(enodeId)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	return r.OrgId, core.GetNodeUrl(r.EnodeId, r.Ip[:], r.Port, r.Raftport, i.Backend.IsRaft), r.NodeStatus, err
}

func (i *Init) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return i.permRoleSession.GetRoleDetails(_roleId, _orgId)
}

func (i *Init) GetSubOrgIndexes(_orgId string) ([]*big.Int, error) {
	return i.permOrgSession.GetSubOrgIndexes(_orgId)
}

func (i *Init) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return i.permOrgSession.GetOrgInfo(_orgIndex)
}

func (i *Init) GetNetworkBootStatus() (bool, error) {
	return i.PermInterfSession.GetNetworkBootStatus()
}

func (i *Init) GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error) {
	return i.permOrgSession.GetOrgDetails(_orgId)
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (i *Init) BindContracts() error {
	log.Debug("permission service: binding contracts")

	err := i.bindContract()
	if err != nil {
		return err
	}

	i.initSession()
	return nil
}

func (a *Audit) ValidatePendingOp(_authOrg, _orgId, _url string, _account common.Address, _pendingOp int64) bool {
	var enodeId string
	var err error
	if _url != "" {
		enodeId, _, _, _, err = getNodeDetails(_url, a.Backend.ContractBackend.IsRaft, a.Backend.ContractBackend.UseDns)
		if err != nil {
			log.Error("permission: encountered error while checking for pending operations", "err", err)
			return false
		}
	}
	pOrg, pUrl, pAcct, op, err := a.Backend.PermInterfSession.GetPendingOp(_authOrg)
	return err == nil && (op.Int64() == _pendingOp && pOrg == _orgId && pUrl == enodeId && pAcct == _account)
}

func (a *Audit) CheckPendingOp(_orgId string) bool {
	_, _, _, op, err := a.Backend.PermInterfSession.GetPendingOp(_orgId)
	return err == nil && op.Int64() != 0
}

func (c *Control) ConnectionAllowed(_enodeId, _ip string, _port, _raftPort uint16) (bool, error) {
	url := core.GetNodeUrl(_enodeId, _ip, _port, _raftPort, c.Backend.ContractBackend.IsRaft)
	enodeId, ip, port, _, err := getNodeDetails(url, c.Backend.ContractBackend.IsRaft, c.Backend.ContractBackend.UseDns)
	if err != nil {
		return false, err
	}

	return c.Backend.PermInterfSession.ConnectionAllowed(enodeId, ip, port)
}

func (c *Control) TransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte, _transactionType core.TransactionType) error {
	if allowed, err := c.Backend.PermInterfSession.TransactionAllowed(_sender, _target, _value, _gasPrice, _gasLimit, _payload); err != nil {
		return err
	} else if !allowed {
		return ptype.ErrNoPermissionForTxn
	}
	return nil
}

func (r *Role) RemoveRole(_args ptype.TxArgs) (*types.Transaction, error) {
	return r.Backend.PermInterfSession.RemoveRole(_args.RoleId, _args.OrgId)
}

func (r *Role) AddNewRole(_args ptype.TxArgs) (*types.Transaction, error) {
	if _args.AccessType > 7 {
		return nil, fmt.Errorf("invalid access type given")
	}
	return r.Backend.PermInterfSession.AddNewRole(_args.RoleId, _args.OrgId, big.NewInt(int64(_args.AccessType)), _args.IsVoter, _args.IsAdmin)
}

func (o *Org) ApproveOrgStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	return o.Backend.PermInterfSession.ApproveOrgStatus(_args.OrgId, big.NewInt(int64(_args.Action)))
}

func (o *Org) UpdateOrgStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	return o.Backend.PermInterfSession.UpdateOrgStatus(_args.OrgId, big.NewInt(int64(_args.Action)))
}

func (o *Org) ApproveOrg(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(_args.Url, o.Backend.ContractBackend.IsRaft, o.Backend.ContractBackend.UseDns)
	if err != nil {
		return nil, err
	}
	return o.Backend.PermInterfSession.ApproveOrg(_args.OrgId, enodeId, ip, port, raftPort, _args.AcctId)
}

func (o *Org) AddSubOrg(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(_args.Url, o.Backend.ContractBackend.IsRaft, o.Backend.ContractBackend.UseDns)
	if err != nil {
		return nil, err
	}
	return o.Backend.PermInterfSession.AddSubOrg(_args.POrgId, _args.OrgId, enodeId, ip, port, raftPort)
}

func (o *Org) AddOrg(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(_args.Url, o.Backend.ContractBackend.IsRaft, o.Backend.ContractBackend.UseDns)
	if err != nil {
		return nil, err
	}
	return o.Backend.PermInterfSession.AddOrg(_args.OrgId, enodeId, ip, port, raftPort, _args.AcctId)
}

func (n *Node) ApproveBlacklistedNodeRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(_args.Url, n.Backend.ContractBackend.IsRaft, n.Backend.ContractBackend.UseDns)
	if err != nil {
		return nil, err
	}
	return n.Backend.PermInterfSession.ApproveBlacklistedNodeRecovery(_args.OrgId, enodeId, ip, port, raftPort)
}

func (n *Node) StartBlacklistedNodeRecovery(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(_args.Url, n.Backend.ContractBackend.IsRaft, n.Backend.ContractBackend.UseDns)
	if err != nil {
		return nil, err
	}
	return n.Backend.PermInterfSession.StartBlacklistedNodeRecovery(_args.OrgId, enodeId, ip, port, raftPort)
}

func (n *Node) AddNode(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(_args.Url, n.Backend.ContractBackend.IsRaft, n.Backend.ContractBackend.UseDns)
	if err != nil {
		return nil, err
	}

	return n.Backend.PermInterfSession.AddNode(_args.OrgId, enodeId, ip, port, raftPort)
}

func (n *Node) UpdateNodeStatus(_args ptype.TxArgs) (*types.Transaction, error) {
	enodeId, ip, port, raftPort, err := getNodeDetails(_args.Url, n.Backend.ContractBackend.IsRaft, n.Backend.ContractBackend.UseDns)
	if err != nil {
		return nil, err
	}
	return n.Backend.PermInterfSession.UpdateNodeStatus(_args.OrgId, enodeId, ip, port, raftPort, big.NewInt(int64(_args.Action)))
}

func (i *Init) bindContract() error {
	if err := ptype.BindContract(&i.PermUpgr, func() (interface{}, error) {
		return binding.NewPermUpgr(i.Backend.PermConfig.UpgrdAddress, i.Backend.EthClnt)
	}); err != nil {
		return err
	}
	if err := ptype.BindContract(&i.PermInterf, func() (interface{}, error) {
		return binding.NewPermInterface(i.Backend.PermConfig.InterfAddress, i.Backend.EthClnt)
	}); err != nil {
		return err
	}
	if err := ptype.BindContract(&i.PermAcct, func() (interface{}, error) {
		return binding.NewAcctManager(i.Backend.PermConfig.AccountAddress, i.Backend.EthClnt)
	}); err != nil {
		return err
	}
	if err := ptype.BindContract(&i.PermNode, func() (interface{}, error) {
		return binding.NewNodeManager(i.Backend.PermConfig.NodeAddress, i.Backend.EthClnt)
	}); err != nil {
		return err
	}
	if err := ptype.BindContract(&i.PermRole, func() (interface{}, error) {
		return binding.NewRoleManager(i.Backend.PermConfig.RoleAddress, i.Backend.EthClnt)
	}); err != nil {
		return err
	}
	if err := ptype.BindContract(&i.PermOrg, func() (interface{}, error) {
		return binding.NewOrgManager(i.Backend.PermConfig.OrgAddress, i.Backend.EthClnt)
	}); err != nil {
		return err
	}
	return nil
}

func (i *Init) initSession() {
	auth := bind.NewKeyedTransactor(i.Backend.Key)
	log.Debug("NodeAccount V2", "nodeAcc", auth.From)
	i.PermInterfSession = &binding.PermInterfaceSession{
		Contract: i.PermInterf,
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

	i.permOrgSession = &binding.OrgManagerSession{
		Contract: i.PermOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	i.permNodeSession = &binding.NodeManagerSession{
		Contract: i.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate roles
	i.permRoleSession = &binding.RoleManagerSession{
		Contract: i.PermRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	//populate accounts
	i.permAcctSession = &binding.AcctManagerSession{
		Contract: i.PermAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
}

// checks if the passed URL is no nil and then calls GetNodeDetails
func getNodeDetails(url string, isRaft, useDns bool) (string, string, uint16, uint16, error) {
	if len(url) > 0 {
		return ptype.GetNodeDetails(url, isRaft, useDns)
	}

	return "", "", uint16(0), uint16(0), nil
}

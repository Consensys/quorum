package permission

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/permission/bind"
	"math/big"
)

type EeaPermInterfaceService struct {
	ps *permission.EeaPermInterfaceSession
}

func (pi *EeaPermInterfaceService) ConnectionAllowed(_enodeId string, _ip [32]byte, _port uint16) (bool, error) {
	return pi.ps.ConnectionAllowed(_enodeId, _ip, _port)
}
func (pi *EeaPermInterfaceService) ConnectionAllowedImpl(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (bool, error) {
	return pi.ps.ConnectionAllowedImpl(_enodeId, _ip, _port, _raftport)
}
func (pi *EeaPermInterfaceService) GetNetworkBootStatus() (bool, error) {
	return pi.ps.GetNetworkBootStatus()
}
func (pi *EeaPermInterfaceService) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return pi.ps.GetPendingOp(_orgId)
}
func (pi *EeaPermInterfaceService) GetPermissionsImpl() (common.Address, error) {
	return pi.ps.GetPermissionsImpl()
}
func (pi *EeaPermInterfaceService) IsNetworkAdmin(_account common.Address) (bool, error) {
	return pi.ps.IsNetworkAdmin(_account)
}
func (pi *EeaPermInterfaceService) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return pi.ps.IsOrgAdmin(_account, _orgId)
}
func (pi *EeaPermInterfaceService) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {
	return pi.ps.TransactionAllowed(_srcaccount, _tgtaccount)
}
func (pi *EeaPermInterfaceService) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return pi.ps.ValidateAccount(_account, _orgId)
}
func (pi *EeaPermInterfaceService) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return pi.ps.AddAdminAccount(_acct)
}
func (pi *EeaPermInterfaceService) AddAdminNode(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.AddAdminNode(_enodeId, _ip, _port, _raftport)
}
func (pi *EeaPermInterfaceService) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	return pi.ps.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}
func (pi *EeaPermInterfaceService) AddNode(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.AddNode(_orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *EeaPermInterfaceService) AddOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address) (*types.Transaction, error) {
	return pi.ps.AddOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}
func (pi *EeaPermInterfaceService) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.AddSubOrg(_pOrgId, _orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *EeaPermInterfaceService) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return pi.ps.ApproveAdminRole(_orgId, _account)
}
func (pi *EeaPermInterfaceService) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return pi.ps.ApproveBlacklistedAccountRecovery(_orgId, _account)
}
func (pi *EeaPermInterfaceService) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.ApproveBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *EeaPermInterfaceService) ApproveOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address) (*types.Transaction, error) {
	return pi.ps.ApproveOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}
func (pi *EeaPermInterfaceService) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.ApproveOrgStatus(_orgId, _action)
}
func (pi *EeaPermInterfaceService) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return pi.ps.AssignAccountRole(_account, _orgId, _roleId)
}
func (pi *EeaPermInterfaceService) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	return pi.ps.AssignAdminRole(_orgId, _account, _roleId)
}
func (pi *EeaPermInterfaceService) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return pi.ps.Init(_breadth, _depth)
}
func (pi *EeaPermInterfaceService) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return pi.ps.RemoveRole(_roleId, _orgId)
}
func (pi *EeaPermInterfaceService) SetPermImplementation(_permImplementation common.Address) (*types.Transaction, error) {
	return pi.ps.SetPermImplementation(_permImplementation)
}
func (pi *EeaPermInterfaceService) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return pi.ps.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}
func (pi *EeaPermInterfaceService) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return pi.ps.StartBlacklistedAccountRecovery(_orgId, _account)
}
func (pi *EeaPermInterfaceService) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.StartBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *EeaPermInterfaceService) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.UpdateAccountStatus(_orgId, _account, _action)
}
func (pi *EeaPermInterfaceService) UpdateNetworkBootStatus() (*types.Transaction, error) {

	return pi.ps.UpdateNetworkBootStatus()
}
func (pi *EeaPermInterfaceService) UpdateNodeStatus(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.UpdateNodeStatus(_orgId, _enodeId, _ip, _port, _raftport, _action)
}
func (pi *EeaPermInterfaceService) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.UpdateOrgStatus(_orgId, _action)
}

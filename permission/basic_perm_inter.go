package permission

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/permission/bind"
	"math/big"
)

type BasicPermInterfaceService struct {
	ps *permission.PermInterfaceSession
}

func (pi *BasicPermInterfaceService) ConnectionAllowed(_enodeId string, _ip [32]byte, _port uint16) (bool, error) {
	panic("not implemented")
}
func (pi *BasicPermInterfaceService) ConnectionAllowedImpl(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (bool, error) {
	panic("not implemented")
}
func (pi *BasicPermInterfaceService) GetNetworkBootStatus() (bool, error) {
	return pi.ps.GetNetworkBootStatus()
}
func (pi *BasicPermInterfaceService) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return pi.ps.GetPendingOp(_orgId)
}
func (pi *BasicPermInterfaceService) GetPermissionsImpl() (common.Address, error) {
	return pi.ps.GetPermissionsImpl()
}
func (pi *BasicPermInterfaceService) IsNetworkAdmin(_account common.Address) (bool, error) {
	return pi.ps.IsNetworkAdmin(_account)
}
func (pi *BasicPermInterfaceService) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return pi.ps.IsOrgAdmin(_account, _orgId)
}
func (pi *BasicPermInterfaceService) TransactionAllowed(_srcaccount common.Address, _tgtaccount common.Address) (bool, error) {
	panic("not implemented")
}
func (pi *BasicPermInterfaceService) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return pi.ps.ValidateAccount(_account, _orgId)
}
func (pi *BasicPermInterfaceService) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return pi.ps.AddAdminAccount(_acct)
}
func (pi *BasicPermInterfaceService) AddAdminNode(_enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.AddAdminNode(_enodeId, _ip, _port, _raftport)
}
func (pi *BasicPermInterfaceService) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	return pi.ps.AddNewRole(_roleId, _orgId, _access, _voter, _admin)
}
func (pi *BasicPermInterfaceService) AddNode(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.AddNode(_orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *BasicPermInterfaceService) AddOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address) (*types.Transaction, error) {
	return pi.ps.AddOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}
func (pi *BasicPermInterfaceService) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.AddSubOrg(_pOrgId, _orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *BasicPermInterfaceService) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return pi.ps.ApproveAdminRole(_orgId, _account)
}
func (pi *BasicPermInterfaceService) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return pi.ps.ApproveBlacklistedAccountRecovery(_orgId, _account)
}
func (pi *BasicPermInterfaceService) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.ApproveBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *BasicPermInterfaceService) ApproveOrg(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _account common.Address) (*types.Transaction, error) {
	return pi.ps.ApproveOrg(_orgId, _enodeId, _ip, _port, _raftport, _account)
}
func (pi *BasicPermInterfaceService) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.ApproveOrgStatus(_orgId, _action)
}
func (pi *BasicPermInterfaceService) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return pi.ps.AssignAccountRole(_account, _orgId, _roleId)
}
func (pi *BasicPermInterfaceService) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	return pi.ps.AssignAdminRole(_orgId, _account, _roleId)
}
func (pi *BasicPermInterfaceService) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return pi.ps.Init(_breadth, _depth)
}
func (pi *BasicPermInterfaceService) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return pi.ps.RemoveRole(_roleId, _orgId)
}
func (pi *BasicPermInterfaceService) SetPermImplementation(_permImplementation common.Address) (*types.Transaction, error) {
	return pi.ps.SetPermImplementation(_permImplementation)
}
func (pi *BasicPermInterfaceService) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return pi.ps.SetPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole)
}
func (pi *BasicPermInterfaceService) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return pi.ps.StartBlacklistedAccountRecovery(_orgId, _account)
}
func (pi *BasicPermInterfaceService) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return pi.ps.StartBlacklistedNodeRecovery(_orgId, _enodeId, _ip, _port, _raftport)
}
func (pi *BasicPermInterfaceService) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.UpdateAccountStatus(_orgId, _account, _action)
}
func (pi *BasicPermInterfaceService) UpdateNetworkBootStatus() (*types.Transaction, error) {

	return pi.ps.UpdateNetworkBootStatus()
}
func (pi *BasicPermInterfaceService) UpdateNodeStatus(_orgId string, _enodeId string, _ip [32]byte, _port uint16, _raftport uint16, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.UpdateNodeStatus(_orgId, _enodeId, _ip, _port, _raftport, _action)
}
func (pi *BasicPermInterfaceService) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return pi.ps.UpdateOrgStatus(_orgId, _action)
}

package types

// to signal all watches when service is stopped
type StopEvent struct {
}

type NodeOperation uint8

const (
	NodeAdd NodeOperation = iota
	NodeDelete
)

type Backend interface {
	// Monitors account access related events and updates the cache accordingly
	ManageAccountPermissions() error
	// Monitors Node management events and updates cache accordingly
	ManageNodePermissions() error
	// monitors org management related events happening via smart contracts
	// and updates cache accordingly
	ManageOrgPermissions() error
	// monitors role management related events and updated cache
	ManageRolePermissions() error
}

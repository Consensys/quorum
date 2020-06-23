package permission

type backend interface {
	// Monitors account access related events and updates the cache accordingly
	manageAccountPermissions() error
	// Monitors Node management events and updates cache accordingly
	manageNodePermissions() error
	// monitors org management related events happening via smart contracts
	// and updates cache accordingly
	manageOrgPermissions() error
	// monitors role management related events and updated cache
	manageRolePermissions() error
}

func newBackend(p *PermissionCtrl) backend {
	if p.eeaFlag {
		return &backendEea{p}
	}
	return &backendBasic{p}
}

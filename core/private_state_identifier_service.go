package core

import (
	"context"
)

type PrivateStateIdentifierService interface {
	ResolveForManagedParty(managedParty string) (string, error)
	ResolveForUserContext(ctx context.Context) (string, error)
}

type TesseraPKPSISImpl struct {
}

func (t *TesseraPKPSISImpl) ResolveForManagedParty(managedParty string) (string, error) {
	return managedParty, nil
}

func (t *TesseraPKPSISImpl) ResolveForUserContext(ctx context.Context) (string, error) {
	psi, ok := ctx.Value("PSI").(string)
	if !ok {
		psi = "private"
	}
	return psi, nil
}

var PSIS PrivateStateIdentifierService = &TesseraPKPSISImpl{}

package qlight

import "context"

type PluginTokenManager interface {
	TokenRefresh(ctx context.Context, currentToken, psi string) (string, error)
	PluginTokenManager(ctx context.Context) (int32, error)
}

//go:generate mockgen -source=service.go -destination service_mockery.go -package qlight
var _ PluginTokenManager = &MockPluginTokenManager{}

type PluginTokenManagerDeferFunc func() (PluginTokenManager, error)

type ReloadablePluginTokenManager struct {
	DeferFunc PluginTokenManagerDeferFunc
}

func (d *ReloadablePluginTokenManager) TokenRefresh(ctx context.Context, currentToken, psi string) (string, error) {
	p, err := d.DeferFunc()
	if err != nil {
		return "", err
	}
	return p.TokenRefresh(ctx, currentToken, psi)
}

func (d *ReloadablePluginTokenManager) PluginTokenManager(ctx context.Context) (int32, error) {
	p, err := d.DeferFunc()
	if err != nil {
		return 0, err
	}
	return p.PluginTokenManager(ctx)
}

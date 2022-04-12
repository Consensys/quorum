package qlight

import "context"

type PluginTokenManager interface {
	TokenRefresh(ctx context.Context, currentToken string) (string, error)
}

type PluginTokenManagerDeferFunc func() (PluginTokenManager, error)

type ReloadablePluginTokenManager struct {
	DeferFunc PluginTokenManagerDeferFunc
}

func (d *ReloadablePluginTokenManager) TokenRefresh(ctx context.Context, currentToken string) (string, error) {
	p, err := d.DeferFunc()
	if err != nil {
		return "", err
	}
	return p.TokenRefresh(ctx, currentToken)
}

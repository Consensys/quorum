package qlight

import (
	"context"

	"github.com/baptiste-b-pegasys/quorum-plugin-qlight-token-manager/proto"
)

type PluginGateway struct {
	client proto.PluginQLightTokenRefresherClient
}

var _ PluginTokenManager = &PluginGateway{}

func (p *PluginGateway) TokenRefresh(ctx context.Context, currentToken, psi string) (string, error) {
	resp, err := p.client.TokenRefresh(ctx, &proto.PluginQLightTokenManager_Request{
		CurrentToken: currentToken,
		Psi:          psi,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

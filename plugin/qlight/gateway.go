package qlight

import (
	"context"
	"fmt"

	"github.com/ConsenSys/quorum-qlight-token-manager-plugin-sdk-go/proto"
)

type PluginGateway struct {
	client proto.PluginQLightTokenRefresherClient
}

var _ PluginTokenManager = &PluginGateway{}

func (p *PluginGateway) TokenRefresh(ctx context.Context, currentToken, psi string) (string, error) {
	resp, err := p.client.TokenRefresh(ctx, &proto.TokenRefresh_Request{
		CurrentToken: currentToken,
		Psi:          psi,
	})
	if err != nil {
		return "", fmt.Errorf("refresh token: %w", err)
	}
	return resp.Token, nil
}

func (p *PluginGateway) PluginTokenManager(ctx context.Context) (int32, error) {
	resp, err := p.client.PluginQLightTokenManager(ctx, &proto.PluginQLightTokenManager_Request{})
	if err != nil {
		return 0, fmt.Errorf("refresh token: %w", err)
	}
	return resp.RefreshAnticipationInMillisecond, nil
}

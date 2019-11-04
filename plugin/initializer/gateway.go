package initializer

import (
	"context"

	"github.com/ethereum/go-ethereum/plugin/proto"
)

type PluginGateway struct {
	client proto.PluginInitializerClient
}

func (g *PluginGateway) Init(ctx context.Context, nodeIdentity string, rawConfiguration []byte) error {
	_, err := g.client.Init(ctx, &proto.PluginInitialization_Request{
		HostIdentity:     nodeIdentity,
		RawConfiguration: rawConfiguration,
	})
	if err != nil {
		return err
	}
	return nil
}

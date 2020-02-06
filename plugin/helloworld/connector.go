package helloworld

import (
	"context"

	iplugin "github.com/ethereum/go-ethereum/internal/plugin"
	"github.com/hashicorp/go-plugin"
	"github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go/proto"
	"google.golang.org/grpc"
)

const ConnectorName = "ping"

type PluginConnector struct {
	plugin.Plugin
}

func (p *PluginConnector) GRPCServer(b *plugin.GRPCBroker, s *grpc.Server) error {
	return iplugin.ErrNotSupported
}

func (p *PluginConnector) GRPCClient(ctx context.Context, b *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &PluginGateway{
		client: proto.NewPluginGreetingClient(cc),
	}, nil
}

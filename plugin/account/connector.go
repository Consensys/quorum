package account

import (
	"context"

	iplugin "github.com/ethereum/go-ethereum/internal/plugin"
	"github.com/hashicorp/go-plugin"
	"github.com/jpmorganchase/quorum-account-plugin-sdk-go/proto"
	"google.golang.org/grpc"
)

const ConnectorName = "account"

type PluginConnector struct {
	plugin.Plugin
}

func (*PluginConnector) GRPCServer(_ *plugin.GRPCBroker, _ *grpc.Server) error {
	return iplugin.ErrNotSupported
}

func (*PluginConnector) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &service{
		client: proto.NewAccountServiceClient(cc),
	}, nil
}

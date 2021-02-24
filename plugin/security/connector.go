package security

import (
	"context"

	iplugin "github.com/ethereum/go-ethereum/internal/plugin"
	"github.com/hashicorp/go-plugin"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	"google.golang.org/grpc"
)

const (
	TLSConfigurationConnectorName = "tls"
	AuthenticationConnectorName   = "auth"
)

type TLSConfigurationSourcePluginConnector struct {
	plugin.Plugin
}

func (*TLSConfigurationSourcePluginConnector) GRPCServer(b *plugin.GRPCBroker, s *grpc.Server) error {
	return iplugin.ErrNotSupported
}

func (*TLSConfigurationSourcePluginConnector) GRPCClient(ctx context.Context, b *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &TLSConfigurationSourcePluginGateway{
		client: proto.NewTLSConfigurationSourceClient(cc),
	}, nil
}

type AuthenticationManagerPluginConnector struct {
	plugin.Plugin
}

func (*AuthenticationManagerPluginConnector) GRPCServer(b *plugin.GRPCBroker, s *grpc.Server) error {
	return iplugin.ErrNotSupported
}

func (*AuthenticationManagerPluginConnector) GRPCClient(ctx context.Context, b *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &AuthenticationManagerPluginGateway{
		client: proto.NewAuthenticationManagerClient(cc),
	}, nil
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	iplugin "github.com/ethereum/go-ethereum/internal/plugin"
	"github.com/ethereum/go-ethereum/plugin/proto"
	"github.com/hashicorp/go-plugin"
)

// this is to demonstrate how to write a plugin that implements HelloWorld plugin interface
func main() {
	log.SetFlags(0) // don't display time
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: iplugin.DefaultHandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"impl": &HelloWorldPluginImpl{},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

// implements 2 interfaces:
// 1. Initializer plugin interface - mandatory
// 2. HelloWorld plugin interface
// 3. GRPC Plugin from go-plugin
type HelloWorldPluginImpl struct {
	plugin.Plugin
	cfg *config
}

func (h *HelloWorldPluginImpl) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginInitializerServer(s, h)
	proto.RegisterPluginGreetingServer(s, h)
	return nil
}

func (h *HelloWorldPluginImpl) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (interface{}, error) {
	return nil, iplugin.ErrNotSupported
}

type config struct {
	Language string
}

func (c *config) validate() error {
	switch l := c.Language; l {
	case "en", "es":
		return nil
	default:
		return fmt.Errorf("unsupported language: [%s]", l)
	}
}

func (h *HelloWorldPluginImpl) Init(_ context.Context, req *proto.PluginInitialization_Request) (*proto.PluginInitialization_Response, error) {
	var cfg config
	if err := json.Unmarshal(req.RawConfiguration, &cfg); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid config: %s, err: %s", string(req.RawConfiguration), err.Error()))
	}
	if err := cfg.validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	h.cfg = &cfg
	return &proto.PluginInitialization_Response{}, nil
}

func (h *HelloWorldPluginImpl) Greeting(_ context.Context, req *proto.PluginHelloWorld_Request) (*proto.PluginHelloWorld_Response, error) {
	switch l := h.cfg.Language; l {
	case "en":
		return &proto.PluginHelloWorld_Response{Msg: fmt.Sprintf("Hello %s!", req.Msg)}, nil
	case "es":
		return &proto.PluginHelloWorld_Response{Msg: fmt.Sprintf("Hola %s!", req.Msg)}, nil
	default:
		return nil, status.Error(codes.Internal, fmt.Sprintf("language [%s] not supported", l))
	}
}

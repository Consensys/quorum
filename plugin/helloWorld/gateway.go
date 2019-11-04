package helloWorld

import (
	"context"

	"github.com/ethereum/go-ethereum/plugin/proto"
)

type PluginGateway struct {
	client proto.PluginGreetingClient
}

func (p *PluginGateway) Greeting(ctx context.Context, msg string) (string, error) {
	resp, err := p.client.Greeting(ctx, &proto.PluginHelloWorld_Request{
		Msg: msg,
	})
	if err != nil {
		return "", err
	}
	return resp.Msg, nil
}

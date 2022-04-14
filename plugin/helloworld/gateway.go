package helloworld

import (
	"context"

	"github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go/proto"
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

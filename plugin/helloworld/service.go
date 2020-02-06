package helloworld

import "context"

type PluginHelloWorld interface {
	Greeting(ctx context.Context, msg string) (string, error)
}

type PluginHelloWorldDeferFunc func() (PluginHelloWorld, error)

type ReloadablePluginHelloWorld struct {
	DeferFunc PluginHelloWorldDeferFunc
}

func (d *ReloadablePluginHelloWorld) Greeting(ctx context.Context, msg string) (string, error) {
	p, err := d.DeferFunc()
	if err != nil {
		return "", err
	}
	return p.Greeting(ctx, msg)
}

package helloWorld

import "context"

type PluginHelloWorld interface {
	Greeting(ctx context.Context, msg string) (string, error)
}

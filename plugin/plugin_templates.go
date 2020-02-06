package plugin

import "github.com/ethereum/go-ethereum/plugin/helloworld"

// a template that returns the hello world plugin instance
type HelloWorldPluginTemplate struct {
	*basePlugin
}

func (p *HelloWorldPluginTemplate) Get() (helloworld.PluginHelloWorld, error) {
	return &helloworld.ReloadablePluginHelloWorld{
		DeferFunc: func() (helloworld.PluginHelloWorld, error) {
			raw, err := p.dispense(helloworld.ConnectorName)
			if err != nil {
				return nil, err
			}
			return raw.(helloworld.PluginHelloWorld), nil
		},
	}, nil
}

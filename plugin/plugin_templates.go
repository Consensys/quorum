package plugin

import "github.com/ethereum/go-ethereum/plugin/helloWorld"

// a template that returns the hello world plugin instance
type HellowWorldPluginTemplate struct {
	*basePlugin
}

func (p *HellowWorldPluginTemplate) Get() (helloWorld.PluginHelloWorld, error) {
	return &helloWorld.ReloadablePluginHelloWorld{
		DeferFunc: func() (helloWorld.PluginHelloWorld, error) {
			raw, err := p.dispense(helloWorld.ConnectorName)
			if err != nil {
				return nil, err
			}
			return raw.(helloWorld.PluginHelloWorld), nil
		},
	}, nil
}

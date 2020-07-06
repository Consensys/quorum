package plugin

import (
	"github.com/ethereum/go-ethereum/plugin/helloworld"
	"github.com/ethereum/go-ethereum/plugin/security"
)

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

type SecurityPluginTemplate struct {
	*basePlugin
}

func (sp *SecurityPluginTemplate) TLSConfigurationSource() (security.TLSConfigurationSource, error) {
	raw, err := sp.dispense(security.TLSConfigurationConnectorName)
	if err != nil {
		return nil, err
	}
	return raw.(security.TLSConfigurationSource), nil
}

func (sp *SecurityPluginTemplate) AuthenticationManager() (security.AuthenticationManager, error) {
	return security.NewDeferredAuthenticationManager(func() (security.AuthenticationManager, error) {
		raw, err := sp.dispense(security.AuthenticationConnectorName)
		if err != nil {
			return nil, err
		}
		return raw.(security.AuthenticationManager), nil
	}), nil
}

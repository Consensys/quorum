package plugin

import "fmt"

type PluginManagerAPI struct {
	pm *PluginManager
}

func NewPluginManagerAPI(pm *PluginManager) *PluginManagerAPI {
	return &PluginManagerAPI{
		pm: pm,
	}
}

func (pmapi *PluginManagerAPI) ReloadPlugin(name PluginInterfaceName) (bool, error) {
	p, ok := pmapi.pm.getPlugin(name)
	if !ok {
		return false, fmt.Errorf("no such plugin provider: %s", name)
	}
	_ = p.Stop()
	if err := p.Start(); err != nil {
		return false, err
	}
	return true, nil
}

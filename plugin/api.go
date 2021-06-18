package plugin

type PluginManagerAPI struct {
	pm *PluginManager
}

func NewPluginManagerAPI(pm *PluginManager) *PluginManagerAPI {
	return &PluginManagerAPI{
		pm: pm,
	}
}

func (pmapi *PluginManagerAPI) ReloadPlugin(name PluginInterfaceName) (bool, error) {
	return pmapi.pm.Reload(name)
}

package qlight

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plugin"
	qlightplugin "github.com/ethereum/go-ethereum/plugin/qlight"
)

type TokenHolder struct {
	token  string
	psi    string
	plugin qlightplugin.PluginTokenManager
}

func NewTokenHolder(psi string, pluginManager *plugin.PluginManager) (*TokenHolder, error) {
	plugin, err := getPlugin(pluginManager)
	if err != nil {
		return nil, err
	}
	return NewTokenHolderWithPlugin(psi, plugin), nil
}

func getPlugin(pluginManager *plugin.PluginManager) (plugin qlightplugin.PluginTokenManager, err error) {
	pluginTemplate, err := tokenManager(pluginManager)
	if err != nil {
		return
	}
	if pluginTemplate != nil {
		err = pluginTemplate.Start()
		if err != nil {
			return
		}
		plugin, err = pluginTemplate.Get()
	}
	return
}

func NewTokenHolderWithPlugin(psi string, plugin qlightplugin.PluginTokenManager) *TokenHolder {
	return &TokenHolder{
		psi:    psi,
		plugin: plugin,
	}
}

func (h *TokenHolder) RefreshPlugin(pluginManager *plugin.PluginManager) (err error) {
	h.plugin, err = getPlugin(pluginManager)
	return
}

func (h *TokenHolder) HttpCredentialsProvider(ctx context.Context) (string, error) {
	if h.plugin != nil {
		log.Debug("HttpCredentialsProvider using", "plugin", h.plugin)
		return h.plugin.TokenRefresh(ctx, h.token, h.psi)
	}
	log.Debug("HttpCredentialsProvider using", "token", h.token)
	return h.token, nil
}

func (h *TokenHolder) GetCurrentToken() string {
	if h == nil {
		log.Debug("token holder nil")
		return ""
	}
	if h.plugin != nil {
		returnedToken, err := h.plugin.TokenRefresh(context.Background(), h.token, h.psi)
		if err != nil {
			log.Error("get token from plugin", "err", err)
		} else {
			log.Debug("new token from plugin", "old", h.token, "new", returnedToken)
			h.token = returnedToken
		}
	} else {
		log.Debug("token plugin is missing", "token", h.token)
	}
	return h.token
}

func (h *TokenHolder) SetCurrentToken(newToken string) {
	h.token = newToken
	log.Debug("token set", "token", newToken)
}

func tokenManager(pluginManager *plugin.PluginManager) (tmp *plugin.QLightTokenManagerPluginTemplate, err error) {
	name := plugin.QLightTokenManagerPluginInterfaceName
	if pluginManager.IsEnabled(name) {
		tmp = new(plugin.QLightTokenManagerPluginTemplate)
		err = pluginManager.GetPluginTemplate(name, tmp)
		return
	}
	log.Info("Token Manager Plugin is not enabled")
	return
}

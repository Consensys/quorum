package qlight

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plugin"
	qlightplugin "github.com/ethereum/go-ethereum/plugin/qlight"
)

type TokenHolder struct {
	token               string
	psi                 string
	refreshAnticipation int32
	plugin              qlightplugin.PluginTokenManager
	pluginManager       *plugin.PluginManager
	peerUpdater         RunningPeerAuthUpdater
}

func NewTokenHolder(psi string, peerUpdater RunningPeerAuthUpdater, pluginManager *plugin.PluginManager) (*TokenHolder, error) {
	plugin, err := getPlugin(pluginManager)
	if err != nil {
		return nil, fmt.Errorf("get plugin: %w", err)
	}
	var refreshAnticipation int32
	if plugin != nil {
		refreshAnticipation, err = plugin.PluginTokenManager(context.Background())
		if err != nil {
			return nil, fmt.Errorf("fetch refresh anticipation value: %w", err)
		}
	}
	return NewTokenHolderWithPlugin(psi, refreshAnticipation, peerUpdater, plugin, pluginManager), nil
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

func NewTokenHolderWithPlugin(psi string, refreshAnticipation int32, peerUpdater RunningPeerAuthUpdater, plugin qlightplugin.PluginTokenManager, pluginManager *plugin.PluginManager) *TokenHolder {
	return &TokenHolder{
		psi:                 psi,
		plugin:              plugin,
		pluginManager:       pluginManager,
		peerUpdater:         peerUpdater,
		refreshAnticipation: refreshAnticipation,
	}
}

func (h *TokenHolder) RefreshPlugin(pluginManager *plugin.PluginManager) (err error) {
	h.plugin, err = getPlugin(pluginManager)
	if err != nil {
		return
	}
	var refreshAnticipation int32
	if h.plugin != nil {
		refreshAnticipation, err = h.plugin.PluginTokenManager(context.Background())
		if err != nil {
			return
		}
	}
	h.refreshAnticipation = refreshAnticipation
	return
}

func (h *TokenHolder) HttpCredentialsProvider(ctx context.Context) (string, error) {
	if h.plugin != nil {
		log.Debug("HttpCredentialsProvider using plugin")
		return h.plugin.TokenRefresh(ctx, h.token, h.psi)
	}
	log.Debug("HttpCredentialsProvider using token")
	return h.token, nil
}

func (h *TokenHolder) ReloadPlugin() error {
	plugin, err := getPlugin(h.pluginManager)
	if err != nil {
		return err
	}
	refreshAnticipation, err := plugin.PluginTokenManager(context.Background())
	if err != nil {
		return fmt.Errorf("fetch refresh anticipation value: %w", err)
	}
	h.plugin = plugin
	h.refreshAnticipation = refreshAnticipation
	return nil
}

func (h *TokenHolder) CurrentToken() string {
	if h == nil {
		log.Warn("token holder nil, returning empty token")
		return ""
	}
	if h.plugin == nil {
		log.Warn("token plugin is missing, no update possible")
		return h.token
	}
	expired, err := h.tokenExpired()
	if err != nil {
		log.Warn("error while checking if token is expired", "err", err)
	}
	if !expired {
		return h.token
	}
	returnedToken, err := h.plugin.TokenRefresh(context.Background(), h.token, h.psi)
	if err != nil {
		log.Error("get token from plugin", "err", err)
	} else {
		if h.token != returnedToken {
			log.Debug("new token from plugin")
			err = h.peerUpdater.UpdateTokenForRunningQPeers(returnedToken)
			if err != nil {
				log.Warn("update token to QPeers", "err", err)
			}
		}
		h.token = returnedToken
	}
	return h.token
}

type JWT struct {
	ExpireAt int64 `json:"exp"`
}

func (h *TokenHolder) tokenExpired() (bool, error) {
	token := h.token
	idx := strings.Index(token, " ")
	if idx >= 0 {
		token = token[idx+1:]
	}
	split := strings.Split(token, ".")
	data, err := base64.RawStdEncoding.DecodeString(split[1])
	if err != nil {
		return false, fmt.Errorf("decode Base64: %w", err)
	}
	jwt := &JWT{}
	err = json.Unmarshal(data, jwt)
	if err != nil {
		return false, fmt.Errorf("unmarshal JSON: %w", err)
	}
	expireAt := time.Unix(jwt.ExpireAt, 0)
	return time.Since(expireAt) >= -time.Duration(h.refreshAnticipation)*time.Millisecond, nil
}

func (h *TokenHolder) SetCurrentToken(v string) {
	h.token = v
}

func (h *TokenHolder) SetExpirationAnticipation(v int32) {
	h.refreshAnticipation = v
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

package qlight

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
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
	timer               *time.Timer
	eta                 time.Time
	lock                sync.Mutex
}

func NewTokenHolder(psi string, pluginManager *plugin.PluginManager) (*TokenHolder, error) {
	plugin, err := getPlugin(pluginManager, new(plugin.QLightTokenManagerPluginTemplate))
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
	return NewTokenHolderWithPlugin(psi, refreshAnticipation, plugin, pluginManager), nil
}

func getPlugin(pluginManager plugin.PluginManagerInterface, pluginTemplate plugin.QLightTokenManagerPluginTemplateInterface) (plugin qlightplugin.PluginTokenManager, err error) {
	pluginTemplate, err = tokenManager(pluginManager, pluginTemplate)
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

func NewTokenHolderWithPlugin(psi string, refreshAnticipation int32, plugin qlightplugin.PluginTokenManager, pluginManager *plugin.PluginManager) *TokenHolder {
	return &TokenHolder{
		psi:                 psi,
		plugin:              plugin,
		pluginManager:       pluginManager,
		refreshAnticipation: refreshAnticipation,
	}
}

func (h *TokenHolder) SetPeerUpdater(peerUpdater RunningPeerAuthUpdater) {
	if h == nil || peerUpdater == nil {
		return
	}
	h.peerUpdater = peerUpdater
}

func (h *TokenHolder) RefreshPlugin(pluginManager plugin.PluginManagerInterface) error {
	return h.refreshPlugin(pluginManager, new(plugin.QLightTokenManagerPluginTemplate))
}

func (h *TokenHolder) refreshPlugin(pluginManager plugin.PluginManagerInterface, template plugin.QLightTokenManagerPluginTemplateInterface) (err error) {
	h.plugin, err = getPlugin(pluginManager, template)
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
	err = h.updateTimer()
	return
}

func (h *TokenHolder) HttpCredentialsProvider(ctx context.Context) (string, error) {
	return h.CurrentToken(), nil
}

func (h *TokenHolder) ReloadPlugin() error {
	plugin, err := getPlugin(h.pluginManager, new(plugin.QLightTokenManagerPluginTemplate))
	if err != nil {
		return err
	}
	refreshAnticipation, err := plugin.PluginTokenManager(context.Background())
	if err != nil {
		return fmt.Errorf("fetch refresh anticipation value: %w", err)
	}
	h.plugin = plugin
	h.refreshAnticipation = refreshAnticipation
	return h.updateTimer()
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
	h.lock.Lock()
	defer h.lock.Unlock()
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
		err = h.updateTimer()
		if err != nil {
			log.Warn("update token timer", "err", err)
		}
	}
	return h.token
}

// updateTimer updates the expiration timer that will trigger automatically a token refreshment
func (h *TokenHolder) updateTimer() error {
	if h == nil && h.plugin == nil {
		return nil
	}
	expireIn, err := h.tokenExpirationDelay()
	if err != nil {
		return err
	}
	if h.timer != nil && time.Now().Add(expireIn).After(h.eta) {
		if !h.timer.Stop() {
			log.Debug("token updateTimer read timer.C", "expire in", expireIn)
			<-h.timer.C
		}
		h.timer = nil
	}
	if h.timer == nil {
		if expireIn <= 0 { // automatic refresh one second after if already expired
			expireIn = time.Second
		} else {
			expireIn = expireIn - time.Duration(h.refreshAnticipation)*time.Millisecond
		}
		h.eta = time.Now().Add(expireIn)
		log.Debug("token updateTimer new", "expire in", expireIn, "eta", h.eta)
		h.timer = time.AfterFunc(expireIn, func() {
			log.Debug("token updateTimer triggered", "expire in", expireIn, "eta", h.eta)
			h.timer = nil
			h.eta = time.Now()
			h.CurrentToken()
		})
	}
	return nil
}

type JWT struct {
	ExpireAt int64 `json:"exp"`
}

func (h *TokenHolder) tokenExpirationDelay() (time.Duration, error) {
	if len(h.token) == 0 {
		return 0, nil
	}
	token := h.token
	idx := strings.Index(token, " ")
	if idx >= 0 {
		token = token[idx+1:]
	}
	split := strings.Split(token, ".")
	if len(split) <= 1 {
		return 0, nil
	}
	data, err := base64.RawStdEncoding.DecodeString(split[1])
	if err != nil {
		return 0, fmt.Errorf("decode Base64: %w", err)
	}
	jwt := &JWT{}
	err = json.Unmarshal(data, jwt)
	if err != nil {
		return 0, fmt.Errorf("unmarshal JSON: %w", err)
	}
	expireAt := time.Unix(jwt.ExpireAt, 0)
	return -time.Since(expireAt), nil // transform negative value to positive as expiration date is in future and time.Since measure in the past
}

func (h *TokenHolder) tokenExpired() (bool, error) {
	expireIn, err := h.tokenExpirationDelay()
	if err != nil {
		return true, err
	}
	return expireIn < time.Duration(h.refreshAnticipation)*time.Millisecond, nil
}

func (h *TokenHolder) SetCurrentToken(v string) {
	h.token = v
}

func (h *TokenHolder) SetExpirationAnticipation(v int32) {
	h.refreshAnticipation = v
}

func tokenManager(pluginManager plugin.PluginManagerInterface, template plugin.QLightTokenManagerPluginTemplateInterface) (plugin.QLightTokenManagerPluginTemplateInterface, error) {
	name := plugin.QLightTokenManagerPluginInterfaceName
	if pluginManager.IsEnabled(name) {
		managedPlugin := template.ManagedPlugin()
		log.Warn("template plugin", "tmp", template, "managed", managedPlugin)
		err := pluginManager.GetPluginTemplate(name, managedPlugin)
		return template, err
	}
	log.Info("Token Manager Plugin is not enabled")
	return nil, nil
}

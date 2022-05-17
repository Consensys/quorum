package qlight

import (
	"testing"

	"github.com/ethereum/go-ethereum/plugin"
	"github.com/ethereum/go-ethereum/plugin/qlight"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenHolder(t *testing.T) {
	th := NewTokenHolderWithPlugin("test", 0, nil, nil)
	require.NotNil(t, th)

	// API mode
	expected := "token"
	th.SetCurrentToken(expected)
	value := th.CurrentToken()
	assert.Equal(t, expected, value)

	expected = "token2"
	th.SetCurrentToken(expected)
	value = th.CurrentToken()
	assert.Equal(t, expected, value)

	// Plugin Token Refresher Mode
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	pluginManager := plugin.NewMockPluginManagerInterface(ctrl)
	pluginManager.
		EXPECT().
		IsEnabled(gomock.Eq(plugin.QLightTokenManagerPluginInterfaceName)).
		Return(true)
	var err error
	pluginManager.
		EXPECT().
		GetPluginTemplate(gomock.Eq(plugin.QLightTokenManagerPluginInterfaceName), gomock.Any()).
		Return(err)

	mockPlugin := qlight.NewMockPluginTokenManager(ctrl)
	mockPlugin.EXPECT().PluginTokenManager(gomock.Any()).Return(int32(1), nil)

	template := plugin.NewMockQLightTokenManagerPluginTemplateInterface(ctrl)
	template.EXPECT().Start().Return(err)
	template.EXPECT().Get().Return(mockPlugin, nil)
	template.EXPECT().ManagedPlugin().Return(nil)

	err = th.refreshPlugin(pluginManager, template)
	require.NoError(t, err)
}

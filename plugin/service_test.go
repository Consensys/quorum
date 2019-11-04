package plugin

import (
	"testing"

	testifyassert "github.com/stretchr/testify/assert"
)

func TestPluginManager_GetPluginTemplate_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	testObject, _ := NewPluginManager("arbitraryName", &Settings{
		Providers: map[PluginInterfaceName]PluginDefinition{
			HelloWorldPluginInterfaceName: {
				Name:    "arbitrary-helloWorld",
				Version: "1.0.0",
				Config:  "arbitrary config",
			},
		},
	}, false, false, "")

	p := new(HellowWorldPluginTemplate)
	err := testObject.GetPluginTemplate(HelloWorldPluginInterfaceName, p)

	assert.NoError(err)
	assert.NotNil(p)
}

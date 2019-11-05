package plugin

import (
	"testing"

	testifyassert "github.com/stretchr/testify/assert"
)

func typicalPluginManager(t *testing.T) *PluginManager {
	testObject, err := NewPluginManager("arbitraryName", &Settings{
		Providers: map[PluginInterfaceName]PluginDefinition{
			HelloWorldPluginInterfaceName: {
				Name:    "arbitrary-helloWorld",
				Version: "1.0.0",
				Config:  "arbitrary config",
			},
		},
	}, false, false, "")

	testifyassert.NoError(t, err)
	return testObject
}

func TestPluginManager_GetPluginTemplate_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)
	testObject := typicalPluginManager(t)

	p := new(HellowWorldPluginTemplate)
	err := testObject.GetPluginTemplate(HelloWorldPluginInterfaceName, p)

	assert.NoError(err)
	assert.NotNil(p)
}

func TestPluginManager_GetPlugin_whenReadFromCache(t *testing.T) {
	assert := testifyassert.New(t)
	testObject := typicalPluginManager(t)
	p := new(HellowWorldPluginTemplate)
	err := testObject.GetPluginTemplate(HelloWorldPluginInterfaceName, p)
	assert.NoError(err)
	assert.NotNil(p)

	actual, ok := testObject.getPlugin(HelloWorldPluginInterfaceName)

	assert.True(ok)
	assert.Equal(p, actual)
}

func TestPluginManager_GetPluginTemplate_whenReadFromCache(t *testing.T) {
	assert := testifyassert.New(t)
	testObject := typicalPluginManager(t)
	p := new(HellowWorldPluginTemplate)
	err := testObject.GetPluginTemplate(HelloWorldPluginInterfaceName, p)
	assert.NoError(err)
	assert.NotNil(p)

	actual := new(HellowWorldPluginTemplate)
	err = testObject.GetPluginTemplate(HelloWorldPluginInterfaceName, actual)

	assert.NoError(err)
	assert.Equal(p, actual)
}

func TestPluginManager_GetPluginTemplate_whenPluginTemplateNotExtendBasePlugin(t *testing.T) {
	assert := testifyassert.New(t)
	testObject := typicalPluginManager(t)

	invalid := new(invalidPluginTemplate)
	err := testObject.GetPluginTemplate(HelloWorldPluginInterfaceName, invalid)

	t.Log(err)
	assert.Error(err)
}

func TestPluginManager_GetPluginTemplate_whenPluginTemplateNotExtendPointerBasePlugin(t *testing.T) {
	assert := testifyassert.New(t)
	testObject := typicalPluginManager(t)

	invalid := new(invalidPluginTemplateNoPointer)
	err := testObject.GetPluginTemplate(HelloWorldPluginInterfaceName, invalid)

	t.Log(err)
	assert.Error(err)
}

type invalidPluginTemplateNoPointer struct {
	basePlugin
}

type invalidPluginTemplate struct {
	someField int
}

func (i invalidPluginTemplate) Start() error {
	panic("implement me")
}

func (i invalidPluginTemplate) Stop() error {
	panic("implement me")
}

func (i invalidPluginTemplate) Info() (PluginInterfaceName, interface{}) {
	panic("implement me")
}

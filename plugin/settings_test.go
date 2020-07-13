package plugin

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/ethereum/go-ethereum/plugin/account"
	"github.com/naoina/toml"
	testifyassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadMultiFormatConfig_whenConfigEmbeddedAsArray(t *testing.T) {
	assert := testifyassert.New(t)

	av1 := "arbitrary value1"
	av2 := "arbitrary value2"

	cfg, err := ReadMultiFormatConfig([]string{av1, av2})

	assert.NoError(err)
	assert.Contains(string(cfg), av1)
	assert.Contains(string(cfg), av2)
}

func TestReadMultiFormatConfig_whenConfigEmbeddedAsFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Remove(tmpFile.Name())
	}()
	av1 := "arbitrary value1"
	_, err = tmpFile.WriteString(av1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("wrote tmp file: " + tmpFile.Name())
	assert := testifyassert.New(t)

	cfg, err := ReadMultiFormatConfig("file://" + tmpFile.Name())

	assert.NoError(err)
	assert.Equal(av1, string(cfg))
}

func TestReadMultiFormatConfig_whenConfigEmbeddedAsString(t *testing.T) {
	av1 := "arbitrary value1"
	assert := testifyassert.New(t)

	cfg, err := ReadMultiFormatConfig(av1)

	assert.NoError(err)
	assert.Equal(av1, string(cfg))
}

func TestReadMultiFormatConfig_whenFromEnvVariable(t *testing.T) {
	assert := testifyassert.New(t)

	arbitraryString := "arbitrary config string"
	if err := os.Setenv("KEY1", arbitraryString); err != nil {
		t.Fatal(err)
	}
	cfg, err := ReadMultiFormatConfig("env://KEY1")

	assert.NoError(err)
	assert.Equal(arbitraryString, string(cfg))
}

func TestReadMultiFormatConfig_whenFromEnvFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Remove(tmpFile.Name())
	}()
	av1 := "arbitrary value1"
	_, err = tmpFile.WriteString(av1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("wrote tmp file: " + tmpFile.Name())
	if err := os.Setenv("KEY1", tmpFile.Name()); err != nil {
		t.Fatal(err)
	}

	assert := testifyassert.New(t)
	cfg, err := ReadMultiFormatConfig("env://KEY1?type=file")

	assert.NoError(err)
	assert.Equal(av1, string(cfg))
}

func TestEnvironmentAwaredValue_UnmarshalJSON_whenValueFromEnvVariable(t *testing.T) {
	assert := testifyassert.New(t)

	if err := os.Setenv("KEY1", "foo"); err != nil {
		t.Fatal(err)
	}

	var value struct {
		Vinstance EnvironmentAwaredValue
		Vpointer  *EnvironmentAwaredValue
	}
	assert.NoError(json.Unmarshal([]byte(`{"Vinstance": "env://KEY1", "Vpointer": "env://KEY1"}`), &value))
	assert.Equal("foo", value.Vinstance.String())
	assert.Equal("foo", value.Vpointer.String())
}

func TestEnvironmentAwaredValue_UnmarshalJSON_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	var value struct {
		Vinstance EnvironmentAwaredValue
		Vpointer  *EnvironmentAwaredValue
	}
	assert.NoError(json.Unmarshal([]byte(`{"Vinstance": "foo", "Vpointer": "bar"}`), &value))
	assert.Equal("foo", value.Vinstance.String())
	assert.Equal("bar", value.Vpointer.String())
}

func TestEnvironmentAwaredValue_UnmarshalTOML_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	var value struct {
		Vinstance EnvironmentAwaredValue
		Vpointer  *EnvironmentAwaredValue
	}
	assert.NoError(toml.Unmarshal([]byte(`
Vinstance = "foo"
Vpointer = "bar"`), &value))
	assert.Equal("foo", value.Vinstance.String())
	assert.Equal("bar", value.Vpointer.String())
}

func TestEnvironmentAwaredValue_UnmarshalTOML_whenValueFromEnvVariable(t *testing.T) {
	assert := testifyassert.New(t)

	if err := os.Setenv("KEY1", "foo"); err != nil {
		t.Fatal(err)
	}

	var value struct {
		Vinstance EnvironmentAwaredValue
		Vpointer  *EnvironmentAwaredValue
	}
	assert.NoError(toml.Unmarshal([]byte(`
Vinstance = "env://KEY1"
Vpointer = "env://KEY1"`), &value))
	assert.Equal("foo", value.Vinstance.String())
	assert.Equal("foo", value.Vpointer.String())
}

func TestPluginInterfaceName_UnmarshalTOML_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	var value struct {
		MyMap map[PluginInterfaceName]string
	}
	assert.NoError(toml.Unmarshal([]byte(`
[MyMap]
Foo = "a1"
BAR = "a2"
`), &value))
	assert.Contains(value.MyMap, PluginInterfaceName("foo"))
	assert.Contains(value.MyMap, PluginInterfaceName("bar"))
}

// For JSON, keys are not being changed. Might be a bug in the decoder
func TestPluginInterfaceName_UnmarshalJSON_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	var value struct {
		MyMap map[PluginInterfaceName]string
	}
	assert.NoError(json.Unmarshal([]byte(`
{
	"MyMap": {
		"Foo" : "a1",
		"BAR" : "a2"
	}
}
`), &value))
	assert.Contains(value.MyMap, PluginInterfaceName("Foo"))
	assert.Contains(value.MyMap, PluginInterfaceName("BAR"))
}

func TestAccountAPIProviderFunc_OnlyExposeAccountCreationAPI(t *testing.T) {
	pm, err := NewPluginManager(
		"arbitraryName",
		&Settings{
			Providers: map[PluginInterfaceName]PluginDefinition{
				AccountPluginInterfaceName: {
					Name:    "arbitrary-account",
					Version: "1.0.0",
					Config:  "arbitrary config",
				},
			},
		},
		false,
		false,
		"",
	)
	require.NoError(t, err)

	provider, ok := pluginProviders[AccountPluginInterfaceName]
	require.True(t, ok)

	api, err := provider.apiProviderFunc("namespace", pm)
	require.NoError(t, err)
	require.Len(t, api, 1)
	require.Equal(t, "namespace", api[0].Namespace)
	require.Implements(t, (*account.CreatorService)(nil), api[0].Service)

	_, ok = api[0].Service.(account.Service)
	require.False(t, ok)
}

func TestSettings_CheckSettingsAreSupported_AllSupported(t *testing.T) {
	s := Settings{
		Providers: map[PluginInterfaceName]PluginDefinition{
			AccountPluginInterfaceName:    {},
			HelloWorldPluginInterfaceName: {},
		},
	}
	supported := []PluginInterfaceName{AccountPluginInterfaceName, HelloWorldPluginInterfaceName}

	err := s.CheckSettingsAreSupported(supported)

	require.NoError(t, err)
}

func TestSettings_CheckSettingsAreSupported_NoneSupported(t *testing.T) {
	s := Settings{
		Providers: map[PluginInterfaceName]PluginDefinition{
			AccountPluginInterfaceName:    {},
			HelloWorldPluginInterfaceName: {},
		},
	}
	supported := []PluginInterfaceName{}

	err := s.CheckSettingsAreSupported(supported)

	require.Error(t, err)

	wantMsgPattern := regexp.MustCompile(`^unsupported plugins configured: \[(account|helloworld) (account|helloworld)\]$`)
	matches := wantMsgPattern.FindStringSubmatch(err.Error())

	// make sure the msg matches the pattern and the same plugin is not listed twice
	require.Regexp(t, wantMsgPattern, err.Error())

	require.NotNil(t, matches, "error message did not match wanted pattern")
	require.Len(t, matches, 3)
	require.NotEmpty(t, matches[1])
	require.NotEmpty(t, matches[2])
	require.NotEqualf(t, matches[1], matches[2], "\"%v\" listed twice", matches[1])
}

func TestSettings_CheckSettingsAreSupported_SomeSupported(t *testing.T) {
	s := Settings{
		Providers: map[PluginInterfaceName]PluginDefinition{
			AccountPluginInterfaceName:    {},
			HelloWorldPluginInterfaceName: {},
		},
	}
	supported := []PluginInterfaceName{AccountPluginInterfaceName}

	err := s.CheckSettingsAreSupported(supported)

	require.EqualError(t, err, "unsupported plugins configured: [helloworld]")
}

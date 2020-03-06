package plugin

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/naoina/toml"

	testifyassert "github.com/stretchr/testify/assert"
)

func TestPluginDefinition_ReadConfig_whenConfigEmbeddedAsArray(t *testing.T) {
	assert := testifyassert.New(t)

	av1 := "arbitrary value1"
	av2 := "arbitrary value2"

	testObject := &PluginDefinition{
		Config: []string{av1, av2},
	}

	cfg, err := testObject.ReadConfig()

	assert.NoError(err)
	assert.Contains(string(cfg), av1)
	assert.Contains(string(cfg), av2)
}

func TestPluginDefinition_ReadConfig_whenConfigEmbeddedAsFile(t *testing.T) {
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

	testObject := &PluginDefinition{
		Config: "file://" + tmpFile.Name(),
	}

	cfg, err := testObject.ReadConfig()

	assert.NoError(err)
	assert.Equal(av1, string(cfg))
}

func TestPluginDefinition_ReadConfig_whenConfigEmbeddedAsString(t *testing.T) {
	av1 := "arbitrary value1"
	assert := testifyassert.New(t)

	testObject := &PluginDefinition{
		Config: av1,
	}

	cfg, err := testObject.ReadConfig()

	assert.NoError(err)
	assert.Equal(av1, string(cfg))
}

func TestPluginDefinition_ReadConfig_whenFromEnvVariable(t *testing.T) {
	assert := testifyassert.New(t)

	arbitraryString := "arbitrary config string"
	if err := os.Setenv("KEY1", arbitraryString); err != nil {
		t.Fatal(err)
	}
	testObject := &PluginDefinition{
		Config: "env://KEY1",
	}

	cfg, err := testObject.ReadConfig()

	assert.NoError(err)
	assert.Equal(arbitraryString, string(cfg))
}

func TestPluginDefinition_ReadConfig_whenFromEnvFile(t *testing.T) {
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

	testObject := &PluginDefinition{
		Config: "env://KEY1?type=file",
	}

	cfg, err := testObject.ReadConfig()

	assert.NoError(err)
	assert.Equal(av1, string(cfg))
	assert.Equal(tmpFile.Name(), testObject.Config)
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

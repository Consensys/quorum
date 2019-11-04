package utils

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/node"
	"github.com/stretchr/testify/assert"
	"gopkg.in/urfave/cli.v1"
)

func TestSetPlugins_whenPluginsNotEnabled(t *testing.T) {
	arbitraryNodeConfig := &node.Config{}
	arbitraryCLIContext := cli.NewContext(nil, &flag.FlagSet{}, nil)

	setPlugins(arbitraryCLIContext, arbitraryNodeConfig)

	assert.Nil(t, arbitraryNodeConfig.Plugins)
}

func TestSetPlugins_whenInvalidFlagsCombination(t *testing.T) {
	arbitraryNodeConfig := &node.Config{}
	fs := &flag.FlagSet{}
	fs.String(PluginSettingsFlag.Name, "", "")
	fs.Bool(PluginSkipVerifyFlag.Name, true, "")
	fs.Bool(PluginLocalVerifyFlag.Name, true, "")
	fs.String(PluginPublicKeyFlag.Name, "", "")
	arbitraryCLIContext := cli.NewContext(nil, fs, nil)
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginSettingsFlag.Name, "arbitrary value"))

	verifyFatalMessage(t, arbitraryCLIContext, arbitraryNodeConfig, "Only --plugins.skipverify or --plugins.localverify must be set")

	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginSkipVerifyFlag.Name, "false"))
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginLocalVerifyFlag.Name, "false"))
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginPublicKeyFlag.Name, "arbitry value"))

	verifyFatalMessage(t, arbitraryCLIContext, arbitraryNodeConfig, "--plugins.localverify is required for setting --plugins.publickey")
}

func TestSetPlugins_whenInvalidPluginSettingsURL(t *testing.T) {
	arbitraryNodeConfig := &node.Config{}
	fs := &flag.FlagSet{}
	fs.String(PluginSettingsFlag.Name, "", "")
	arbitraryCLIContext := cli.NewContext(nil, fs, nil)
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginSettingsFlag.Name, "arbitrary value"))

	verifyFatalMessage(t, arbitraryCLIContext, arbitraryNodeConfig, "plugins: unable to create reader due to unsupported scheme ")
}

func TestSetPlugins_whenTypical(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()
	arbitraryJSONFile := path.Join(tmpDir, "arbitary.json")
	if err := ioutil.WriteFile(arbitraryJSONFile, []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}
	arbitraryNodeConfig := &node.Config{}
	fs := &flag.FlagSet{}
	fs.String(PluginSettingsFlag.Name, "", "")
	arbitraryCLIContext := cli.NewContext(nil, fs, nil)
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginSettingsFlag.Name, "file://"+arbitraryJSONFile))

	setPlugins(arbitraryCLIContext, arbitraryNodeConfig)

	assert.NotNil(t, arbitraryNodeConfig.Plugins)
}

func verifyFatalMessage(t *testing.T, ctx *cli.Context, cfg *node.Config, expectedMsg string) {
	msgCaptor := newFatalMessageCaptor()
	saved := fatalfFunc
	defer func() {
		fatalfFunc = saved
		recover() // as tests would result fatal, we need to assert here
		assert.Equal(t, expectedMsg, msgCaptor.capturedValue)
	}()
	fatalfFunc = msgCaptor.fatalfFunc

	setPlugins(ctx, cfg)
}

type fatalMessageCaptor struct {
	fatalfFunc    func(format string, args ...interface{})
	capturedValue string
}

func newFatalMessageCaptor() *fatalMessageCaptor {
	captor := &fatalMessageCaptor{}
	captor.fatalfFunc = func(format string, args ...interface{}) {
		captor.capturedValue = fmt.Sprintf(format, args...)
		panic("suspending")
	}
	return captor
}

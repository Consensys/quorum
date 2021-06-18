package main

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/plugin"
	"github.com/stretchr/testify/require"
	"gopkg.in/urfave/cli.v1"
)

// newAccountPluginCLIContext creates a cli.Context setup with the core account plugin CLI flags.
// args sets the values of the flags.
func newAccountPluginCLIContext(args []string) *cli.Context {
	fs := &flag.FlagSet{}
	fs.String(utils.PluginSettingsFlag.Name, "", "")
	fs.String(utils.AccountPluginNewAccountConfigFlag.Name, "", "")
	_ = fs.Parse(args)

	return cli.NewContext(nil, fs, nil)
}

type mockConfigNodeMaker struct {
	do func(ctx *cli.Context) (*node.Node, gethConfig)
}

func (m *mockConfigNodeMaker) makeConfigNode(ctx *cli.Context) (*node.Node, gethConfig) {
	return m.do(ctx)
}

func TestListPluginAccounts_ErrIfCLIFlagNotSet(t *testing.T) {
	var args []string
	ctx := newAccountPluginCLIContext(args)

	_, err := listPluginAccounts(ctx)
	require.EqualError(t, err, "--plugins required")
}

func TestListPluginAccounts_ErrIfUnsupportedPluginInConfig(t *testing.T) {
	var unsupportedPlugin plugin.PluginInterfaceName = "somename"
	pluginSettings := plugin.Settings{
		Providers: map[plugin.PluginInterfaceName]plugin.PluginDefinition{
			unsupportedPlugin: {},
		},
	}

	args := []string{
		"--plugins", "/path/to/config.json",
	}
	ctx := newAccountPluginCLIContext(args)

	makeConfigNodeDelegate = &mockConfigNodeMaker{
		do: func(ctx *cli.Context) (*node.Node, gethConfig) {
			return nil, gethConfig{
				Node: node.Config{
					Plugins: &pluginSettings,
				},
			}
		},
	}

	_, err := listPluginAccounts(ctx)
	require.EqualError(t, err, "unsupported plugins configured: [somename]")
}

func TestCreatePluginAccount_ErrIfCLIFlagsNotSet(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "no plugin flags",
			args: []string{},
		},
		{
			name: "only plugin settings flag",
			args: []string{"--plugins", "/path/to/config.json"},
		},
		{
			name: "only new plugin account config settings flag",
			args: []string{"--plugins.account.config", "/path/to/new-acct-config.json"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := newAccountPluginCLIContext(tt.args)

			_, err := createPluginAccount(ctx)
			require.EqualError(t, err, "--plugins and --plugins.account.config flags must be set")
		})
	}
}

func TestCreatePluginAccount_ErrIfInvalidNewAccountConfig(t *testing.T) {
	tests := []struct {
		name       string
		flagValue  string
		wantErrMsg string
	}{
		{
			name:       "json: invalid json",
			flagValue:  "{invalidjson: abc}",
			wantErrMsg: "invalid account creation config provided: invalid character 'i' looking for beginning of object key string",
		},
		{
			name:       "file: does not exist",
			flagValue:  "file://doesnotexist",
			wantErrMsg: "invalid account creation config provided: open doesnotexist: no such file or directory",
		},
		{
			name:       "env: not set",
			flagValue:  "env://notset",
			wantErrMsg: "invalid account creation config provided: env variable notset not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []string{
				"--plugins", "/path/to/config.json",
				"--plugins.account.config", tt.flagValue,
			}
			ctx := newAccountPluginCLIContext(args)
			_, err := createPluginAccount(ctx)
			require.EqualError(t, err, tt.wantErrMsg)
		})
	}
}

func TestCreatePluginAccount_ErrIfUnsupportedPluginInConfig(t *testing.T) {
	var unsupportedPlugin plugin.PluginInterfaceName = "somename"
	pluginSettings := plugin.Settings{
		Providers: map[plugin.PluginInterfaceName]plugin.PluginDefinition{
			unsupportedPlugin: {},
		},
	}

	args := []string{
		"--plugins", "/path/to/config.json",
		"--plugins.account.config", "{}",
	}
	ctx := newAccountPluginCLIContext(args)

	makeConfigNodeDelegate = &mockConfigNodeMaker{
		do: func(ctx *cli.Context) (*node.Node, gethConfig) {
			return nil, gethConfig{
				Node: node.Config{
					Plugins: &pluginSettings,
				},
			}
		},
	}

	_, err := createPluginAccount(ctx)
	require.EqualError(t, err, "unsupported plugins configured: [somename]")
}

func TestImportPluginAccount_ErrIfNoArg(t *testing.T) {
	var args []string
	ctx := newAccountPluginCLIContext(args)

	_, err := importPluginAccount(ctx)
	require.EqualError(t, err, "keyfile must be given as argument")
}

func TestImportPluginAccount_ErrIfInvalidRawkey(t *testing.T) {
	args := []string{"/incorrect/path/to/file.key"}
	ctx := newAccountPluginCLIContext(args)

	_, err := importPluginAccount(ctx)
	require.EqualError(t, err, "Failed to load the private key: open /incorrect/path/to/file.key: no such file or directory")
}

func TestImportPluginAccount_ErrIfCLIFlagsNotSet(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "rawkey")
	require.NoError(t, err)
	t.Log("creating tmp file", "path", tmpfile.Name())
	defer os.Remove(tmpfile.Name())
	_, err = tmpfile.Write([]byte("1fe8f1ad4053326db20529257ac9401f2e6c769ef1d736b8c2f5aba5f787c72b"))
	require.NoError(t, err)
	err = tmpfile.Close()
	require.NoError(t, err)

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "no plugin flags",
			args: []string{tmpfile.Name()},
		},
		{
			name: "only plugin settings flag",
			args: []string{"--plugins", "/path/to/config.json", tmpfile.Name()},
		},
		{
			name: "only new plugin account config settings flag",
			args: []string{"--plugins.account.config", "/path/to/new-acct-config.json", tmpfile.Name()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := newAccountPluginCLIContext(tt.args)

			_, err := importPluginAccount(ctx)
			require.EqualError(t, err, "--plugins and --plugins.account.config flags must be set")
		})
	}
}

func TestImportPluginAccount_ErrIfInvalidNewAccountConfig(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "rawkey")
	require.NoError(t, err)
	t.Log("creating tmp file", "path", tmpfile.Name())
	defer os.Remove(tmpfile.Name())
	_, err = tmpfile.Write([]byte("1fe8f1ad4053326db20529257ac9401f2e6c769ef1d736b8c2f5aba5f787c72b"))
	require.NoError(t, err)
	err = tmpfile.Close()
	require.NoError(t, err)

	tests := []struct {
		name       string
		flagValue  string
		wantErrMsg string
	}{
		{
			name:       "json: invalid json",
			flagValue:  "{invalidjson: abc}",
			wantErrMsg: "invalid account creation config provided: invalid character 'i' looking for beginning of object key string",
		},
		{
			name:       "file: does not exist",
			flagValue:  "file://doesnotexist",
			wantErrMsg: "invalid account creation config provided: open doesnotexist: no such file or directory",
		},
		{
			name:       "env: not set",
			flagValue:  "env://notset",
			wantErrMsg: "invalid account creation config provided: env variable notset not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []string{
				"--plugins", "/path/to/config.json",
				"--plugins.account.config", tt.flagValue,
				tmpfile.Name(),
			}
			ctx := newAccountPluginCLIContext(args)
			_, err := importPluginAccount(ctx)
			require.EqualError(t, err, tt.wantErrMsg)
		})
	}
}

func TestImportPluginAccount_ErrIfUnsupportedPluginInConfig(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "rawkey")
	require.NoError(t, err)
	t.Log("creating tmp file", "path", tmpfile.Name())
	defer os.Remove(tmpfile.Name())
	_, err = tmpfile.Write([]byte("1fe8f1ad4053326db20529257ac9401f2e6c769ef1d736b8c2f5aba5f787c72b"))
	require.NoError(t, err)
	err = tmpfile.Close()
	require.NoError(t, err)

	var unsupportedPlugin plugin.PluginInterfaceName = "somename"
	pluginSettings := plugin.Settings{
		Providers: map[plugin.PluginInterfaceName]plugin.PluginDefinition{
			unsupportedPlugin: {},
		},
	}

	args := []string{
		"--plugins", "/path/to/config.json",
		"--plugins.account.config", "{}",
		tmpfile.Name(),
	}
	ctx := newAccountPluginCLIContext(args)

	makeConfigNodeDelegate = &mockConfigNodeMaker{
		do: func(ctx *cli.Context) (*node.Node, gethConfig) {
			return nil, gethConfig{
				Node: node.Config{
					Plugins: &pluginSettings,
				},
			}
		},
	}

	_, err = importPluginAccount(ctx)
	require.EqualError(t, err, "unsupported plugins configured: [somename]")
}

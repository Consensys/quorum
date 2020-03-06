// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package node

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/plugin"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
)

// Tests that datadirs can be successfully created, be them manually configured
// ones or automatically generated temporary ones.
func TestDatadirCreation(t *testing.T) {
	// Create a temporary data dir and check that it can be used by a node
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("failed to create manual data dir: %v", err)
	}
	defer os.RemoveAll(dir)

	if _, err := New(&Config{DataDir: dir}); err != nil {
		t.Fatalf("failed to create stack with existing datadir: %v", err)
	}
	// Generate a long non-existing datadir path and check that it gets created by a node
	dir = filepath.Join(dir, "a", "b", "c", "d", "e", "f")
	if _, err := New(&Config{DataDir: dir}); err != nil {
		t.Fatalf("failed to create stack with creatable datadir: %v", err)
	}
	if _, err := os.Stat(dir); err != nil {
		t.Fatalf("freshly created datadir not accessible: %v", err)
	}
	// Verify that an impossible datadir fails creation
	file, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	dir = filepath.Join(file.Name(), "invalid/path")
	if _, err := New(&Config{DataDir: dir}); err == nil {
		t.Fatalf("protocol stack created with an invalid datadir")
	}
}

// Tests that IPC paths are correctly resolved to valid endpoints of different
// platforms.
func TestIPCPathResolution(t *testing.T) {
	var tests = []struct {
		DataDir  string
		IPCPath  string
		Windows  bool
		Endpoint string
	}{
		{"", "", false, ""},
		{"data", "", false, ""},
		{"", "geth.ipc", false, filepath.Join(os.TempDir(), "geth.ipc")},
		{"data", "geth.ipc", false, "data/geth.ipc"},
		{"data", "./geth.ipc", false, "./geth.ipc"},
		{"data", "/geth.ipc", false, "/geth.ipc"},
		{"", "", true, ``},
		{"data", "", true, ``},
		{"", "geth.ipc", true, `\\.\pipe\geth.ipc`},
		{"data", "geth.ipc", true, `\\.\pipe\geth.ipc`},
		{"data", `\\.\pipe\geth.ipc`, true, `\\.\pipe\geth.ipc`},
	}
	for i, test := range tests {
		// Only run when platform/test match
		if (runtime.GOOS == "windows") == test.Windows {
			if endpoint := (&Config{DataDir: test.DataDir, IPCPath: test.IPCPath}).IPCEndpoint(); endpoint != test.Endpoint {
				t.Errorf("test %d: IPC endpoint mismatch: have %s, want %s", i, endpoint, test.Endpoint)
			}
		}
	}
}

// Tests that node keys can be correctly created, persisted, loaded and/or made
// ephemeral.
func TestNodeKeyPersistency(t *testing.T) {
	// Create a temporary folder and make sure no key is present
	dir, err := ioutil.TempDir("", "node-test")
	if err != nil {
		t.Fatalf("failed to create temporary data directory: %v", err)
	}
	defer os.RemoveAll(dir)

	keyfile := filepath.Join(dir, "unit-test", datadirPrivateKey)

	// Configure a node with a preset key and ensure it's not persisted
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate one-shot node key: %v", err)
	}
	config := &Config{Name: "unit-test", DataDir: dir, P2P: p2p.Config{PrivateKey: key}}
	config.NodeKey()
	if _, err := os.Stat(filepath.Join(keyfile)); err == nil {
		t.Fatalf("one-shot node key persisted to data directory")
	}

	// Configure a node with no preset key and ensure it is persisted this time
	config = &Config{Name: "unit-test", DataDir: dir}
	config.NodeKey()
	if _, err := os.Stat(keyfile); err != nil {
		t.Fatalf("node key not persisted to data directory: %v", err)
	}
	if _, err = crypto.LoadECDSA(keyfile); err != nil {
		t.Fatalf("failed to load freshly persisted node key: %v", err)
	}
	blob1, err := ioutil.ReadFile(keyfile)
	if err != nil {
		t.Fatalf("failed to read freshly persisted node key: %v", err)
	}

	// Configure a new node and ensure the previously persisted key is loaded
	config = &Config{Name: "unit-test", DataDir: dir}
	config.NodeKey()
	blob2, err := ioutil.ReadFile(filepath.Join(keyfile))
	if err != nil {
		t.Fatalf("failed to read previously persisted node key: %v", err)
	}
	if !bytes.Equal(blob1, blob2) {
		t.Fatalf("persisted node key mismatch: have %x, want %x", blob2, blob1)
	}

	// Configure ephemeral node and ensure no key is dumped locally
	config = &Config{Name: "unit-test", DataDir: ""}
	config.NodeKey()
	if _, err := os.Stat(filepath.Join(".", "unit-test", datadirPrivateKey)); err == nil {
		t.Fatalf("ephemeral node key persisted to disk")
	}
}

func TestConfig_ResolvePluginBaseDir_whenPluginFeatureIsDisabled(t *testing.T) {
	testObject := &Config{}

	assert.NoError(t, testObject.ResolvePluginBaseDir())
}

func TestConfig_ResolvePluginBaseDir_whenBaseDirDoesNotExist(t *testing.T) {
	arbitraryBaseDir := path.Join(os.TempDir(), fmt.Sprintf("foo-%d", time.Now().Unix()))
	defer func() {
		_ = os.RemoveAll(arbitraryBaseDir)
	}()
	testObject := &Config{
		Plugins: &plugin.Settings{
			BaseDir: plugin.EnvironmentAwaredValue(arbitraryBaseDir),
		},
	}

	assert.NoError(t, testObject.ResolvePluginBaseDir())
	assert.True(t, common.FileExist(arbitraryBaseDir))
	assert.True(t, path.IsAbs(testObject.Plugins.BaseDir.String()))
}

func TestConfig_ResolvePluginBaseDir_whenBaseDirExists(t *testing.T) {
	arbitraryBaseDir, err := ioutil.TempDir("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(arbitraryBaseDir)
	}()
	testObject := &Config{
		Plugins: &plugin.Settings{
			BaseDir: plugin.EnvironmentAwaredValue(arbitraryBaseDir),
		},
	}

	assert.NoError(t, testObject.ResolvePluginBaseDir())
	assert.True(t, path.IsAbs(testObject.Plugins.BaseDir.String()))
}

// Quorum
//
func TestConfig_IsPermissionEnabled_whenTypical(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpdir)
	}()
	if err := ioutil.WriteFile(path.Join(tmpdir, params.PERMISSION_MODEL_CONFIG), []byte("foo"), 0644); err != nil {
		t.Fatal(err)
	}
	testObject := &Config{
		EnableNodePermission: true,
		DataDir:              tmpdir,
	}

	assert.True(t, testObject.IsPermissionEnabled())
}

// Quorum
//
func TestConfig_IsPermissionEnabled_whenPermissionedFlagIsFalse(t *testing.T) {
	testObject := &Config{
		EnableNodePermission: false,
	}

	assert.False(t, testObject.IsPermissionEnabled())
}

// Quorum
//
func TestConfig_IsPermissionEnabled_whenPermissionConfigIsNotAvailable(t *testing.T) {
	testObject := &Config{
		EnableNodePermission: true,
		DataDir:              os.TempDir(),
	}

	assert.False(t, testObject.IsPermissionEnabled())
}

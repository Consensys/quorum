// Copyright 2019 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// Package utils contains internal helper functions for go-ethereum commands.
package utils

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/stretchr/testify/assert"
	"gopkg.in/urfave/cli.v1"
)

func TestSetPlugins_whenPluginsNotEnabled(t *testing.T) {
	arbitraryNodeConfig := &node.Config{}
	arbitraryCLIContext := cli.NewContext(nil, &flag.FlagSet{}, nil)

	assert.NoError(t, SetPlugins(arbitraryCLIContext, arbitraryNodeConfig))

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

	verifyErrorMessage(t, arbitraryCLIContext, arbitraryNodeConfig, "only --plugins.skipverify or --plugins.localverify must be set")

	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginSkipVerifyFlag.Name, "false"))
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginLocalVerifyFlag.Name, "false"))
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginPublicKeyFlag.Name, "arbitrary value"))

	verifyErrorMessage(t, arbitraryCLIContext, arbitraryNodeConfig, "--plugins.localverify is required for setting --plugins.publickey")
}

func TestSetPlugins_whenInvalidPluginSettingsURL(t *testing.T) {
	arbitraryNodeConfig := &node.Config{}
	fs := &flag.FlagSet{}
	fs.String(PluginSettingsFlag.Name, "", "")
	arbitraryCLIContext := cli.NewContext(nil, fs, nil)
	assert.NoError(t, arbitraryCLIContext.GlobalSet(PluginSettingsFlag.Name, "arbitrary value"))

	verifyErrorMessage(t, arbitraryCLIContext, arbitraryNodeConfig, "plugins: unable to create reader due to unsupported scheme ")
}

func TestSetImmutabilityThreshold(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Int(QuorumImmutabilityThreshold.Name, 0, "")
	arbitraryCLIContext := cli.NewContext(nil, fs, nil)
	assert.NoError(t, arbitraryCLIContext.GlobalSet(QuorumImmutabilityThreshold.Name, strconv.Itoa(100000)))
	assert.True(t, arbitraryCLIContext.GlobalIsSet(QuorumImmutabilityThreshold.Name), "immutability threshold flag not set")
	assert.Equal(t, 100000, arbitraryCLIContext.GlobalInt(QuorumImmutabilityThreshold.Name), "immutability threshold value not set")
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

	assert.NoError(t, SetPlugins(arbitraryCLIContext, arbitraryNodeConfig))

	assert.NotNil(t, arbitraryNodeConfig.Plugins)
}

func verifyErrorMessage(t *testing.T, ctx *cli.Context, cfg *node.Config, expectedMsg string) {
	err := SetPlugins(ctx, cfg)
	assert.EqualError(t, err, expectedMsg)
}

func Test_SplitTagsFlag(t *testing.T) {
	tests := []struct {
		name string
		args string
		want map[string]string
	}{
		{
			"2 tags case",
			"host=localhost,bzzkey=123",
			map[string]string{
				"host":   "localhost",
				"bzzkey": "123",
			},
		},
		{
			"1 tag case",
			"host=localhost123",
			map[string]string{
				"host": "localhost123",
			},
		},
		{
			"empty case",
			"",
			map[string]string{},
		},
		{
			"garbage",
			"smth=smthelse=123",
			map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitTagsFlag(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitTagsFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuorumConfigFlags(t *testing.T) {
	fs := &flag.FlagSet{}
	arbitraryCLIContext := cli.NewContext(nil, fs, nil)
	arbitraryEthConfig := &eth.Config{}

	fs.Int(EVMCallTimeOutFlag.Name, 0, "")
	assert.NoError(t, arbitraryCLIContext.GlobalSet(EVMCallTimeOutFlag.Name, strconv.Itoa(12)))
	fs.Bool(MultitenancyFlag.Name, false, "")
	assert.NoError(t, arbitraryCLIContext.GlobalSet(MultitenancyFlag.Name, "true"))
	fs.Bool(QuorumEnablePrivacyMarker.Name, true, "")
	assert.NoError(t, arbitraryCLIContext.GlobalSet(QuorumEnablePrivacyMarker.Name, "true"))
	fs.Uint64(IstanbulRequestTimeoutFlag.Name, 0, "")
	assert.NoError(t, arbitraryCLIContext.GlobalSet(IstanbulRequestTimeoutFlag.Name, "23"))
	fs.Uint64(IstanbulBlockPeriodFlag.Name, 0, "")
	assert.NoError(t, arbitraryCLIContext.GlobalSet(IstanbulBlockPeriodFlag.Name, "34"))
	fs.Bool(RaftModeFlag.Name, false, "")
	assert.NoError(t, arbitraryCLIContext.GlobalSet(RaftModeFlag.Name, "true"))

	if !assert.NoError(t, setQuorumConfig(arbitraryCLIContext, arbitraryEthConfig)) {
		t.FailNow()
	}

	assert.True(t, arbitraryCLIContext.GlobalIsSet(EVMCallTimeOutFlag.Name), "EVMCallTimeOutFlag not set")
	assert.True(t, arbitraryCLIContext.GlobalIsSet(MultitenancyFlag.Name), "MultitenancyFlag not set")
	assert.True(t, arbitraryCLIContext.GlobalIsSet(RaftModeFlag.Name), "RaftModeFlag not set")

	assert.Equal(t, 12*time.Second, arbitraryEthConfig.EVMCallTimeOut, "EVMCallTimeOut value is incorrect")
	assert.Equal(t, true, arbitraryEthConfig.EnableMultitenancy, "MultitenancyFlag value is incorrect")
	assert.Equal(t, true, arbitraryEthConfig.QuorumPrivacyMarkerTransactionsEnabled, "QuorumEnablePrivacyMarker value is incorrect")
	assert.Equal(t, uint64(23), arbitraryEthConfig.Istanbul.RequestTimeout, "IstanbulRequestTimeoutFlag value is incorrect")
	assert.Equal(t, uint64(34), arbitraryEthConfig.Istanbul.BlockPeriod, "IstanbulBlockPeriodFlag value is incorrect")
	assert.Equal(t, true, arbitraryEthConfig.RaftMode, "RaftModeFlag value is incorrect")
}

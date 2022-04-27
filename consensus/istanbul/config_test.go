// Copyright 2017 The go-ethereum Authors
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

package istanbul

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/naoina/toml"
	"github.com/stretchr/testify/assert"
)

func TestProposerPolicy_UnmarshalTOML(t *testing.T) {
	input := `id = 2
`
	expectedId := ProposerPolicyId(2)
	var p proposerPolicyToml
	assert.NoError(t, toml.Unmarshal([]byte(input), &p))
	assert.Equal(t, expectedId, p.Id, "ProposerPolicyId mismatch")
}

func TestProposerPolicy_MarshalTOML(t *testing.T) {
	output := `id = 1
`
	p := &ProposerPolicy{Id: 1}
	b, err := p.MarshalTOML()
	if err != nil {
		t.Errorf("error marshalling ProposerPolicy: %v", err)
	}
	assert.Equal(t, output, b, "ProposerPolicy MarshalTOML mismatch")
}

func TestGetConfig(t *testing.T) {
	if !reflect.DeepEqual(DefaultConfig.GetConfig(nil), *DefaultConfig) {
		t.Errorf("error default config:\nexpected: %v\n", DefaultConfig)
	}

	config := DefaultConfig
	config.Transitions = []params.Transition{{
		Block:       big.NewInt(1),
		EpochLength: 40000,
	}, {
		Block:              big.NewInt(3),
		BlockPeriodSeconds: 5,
	}, {
		Block:                 big.NewInt(5),
		RequestTimeoutSeconds: 15000,
	}}
	config1 := *DefaultConfig
	config1.Epoch = 40000
	config3 := config1
	config3.BlockPeriod = 5
	config5 := config3
	config5.RequestTimeout = 15000

	type test struct {
		blockNumber    int64
		expectedConfig Config
	}
	tests := []test{
		{1, config1},
		{2, config1},
		{3, config3},
		{4, config3},
		{5, config5},
		{10, config5},
		{100, config5},
	}

	for _, test := range tests {
		c := config.GetConfig(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(c, test.expectedConfig) {
			t.Errorf("error mismatch:\nexpected: %v\ngot: %v\n", test.expectedConfig, c)
		}
	}
}

func TestIsQBFTConsensusAt(t *testing.T) {
	config1 := *DefaultConfig
	config1.TestQBFTBlock = nil
	config2 := *DefaultConfig
	config2.TestQBFTBlock = big.NewInt(5)
	config3 := *DefaultConfig
	config3.TestQBFTBlock = nil
	config3.Transitions = []params.Transition{
		{Block: big.NewInt(10), Algorithm: params.QBFT},
	}
	type test struct {
		config      Config
		blockNumber int64
		isQBFT      bool
	}
	tests := []test{
		{*DefaultConfig, 0, true},
		{*DefaultConfig, 10, true},
		{config1, 0, false},
		{config1, 10, false},
		{config2, 4, false},
		{config2, 5, true},
		{config2, 7, true},
		{config3, 0, false},
		{config3, 7, false},
		{config3, 10, true},
		{config3, 11, true},
	}
	for _, test := range tests {
		isQbft := test.config.IsQBFTConsensusAt(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(isQbft, test.isQBFT) {
			t.Errorf("error mismatch:\nexpected: %v\ngot: %v\n", test.isQBFT, isQbft)
		}
	}
}

func TestEmptyBlock(t *testing.T) {
	vFalse := false
	vTrue := true
	config1 := *DefaultConfig
	config1.EmptyBlock = nil
	config2 := *DefaultConfig
	config2.EmptyBlock = &vFalse
	config3 := *DefaultConfig
	config3.EmptyBlock = &vTrue
	config4 := *DefaultConfig
	config4.EmptyBlock = &vTrue
	config4.Transitions = []params.Transition{
		{Block: big.NewInt(5), EmptyBlock: nil},
		{Block: big.NewInt(10), EmptyBlock: &vFalse},
	}
	type test struct {
		config      Config
		blockNumber int64
		emptyBlock  bool
	}
	tests := []test{
		{*DefaultConfig, 0, true},
		{*DefaultConfig, 10, true},
		{config1, 0, true},
		{config1, 10, true},
		{config2, 0, false},
		{config2, 10, false},
		{config3, 0, true},
		{config3, 10, true},
		{config4, 0, true},
		{config4, 9, true},
		{config4, 10, false},
		{config4, 11, false},
	}
	for idx, test := range tests {
		emptyBlock := test.config.EmptyBlockGeneration(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(emptyBlock, test.emptyBlock) {
			t.Errorf("error mismatch: test=%d\nexpected: %v\ngot: %v\n", idx, test.emptyBlock, emptyBlock)
		}
	}
}

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
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/naoina/toml"
)

type ProposerPolicyId uint64

const (
	RoundRobin ProposerPolicyId = iota
	Sticky
)

// ProposerPolicy represents the Validator Proposer Policy
type ProposerPolicy struct {
	Id         ProposerPolicyId    // Could be RoundRobin or Sticky
	By         ValidatorSortByFunc // func that defines how the ValidatorSet should be sorted
	registry   []ValidatorSet      // Holds the ValidatorSet for a given block height
	registryMU *sync.Mutex         // Mutex to lock access to changes to Registry
}

// NewRoundRobinProposerPolicy returns a RoundRobin ProposerPolicy with ValidatorSortByString as default sort function
func NewRoundRobinProposerPolicy() *ProposerPolicy {
	return NewProposerPolicy(RoundRobin)
}

// NewStickyProposerPolicy return a Sticky ProposerPolicy with ValidatorSortByString as default sort function
func NewStickyProposerPolicy() *ProposerPolicy {
	return NewProposerPolicy(Sticky)
}

func NewProposerPolicy(id ProposerPolicyId) *ProposerPolicy {
	return NewProposerPolicyByIdAndSortFunc(id, ValidatorSortByString())
}

func NewProposerPolicyByIdAndSortFunc(id ProposerPolicyId, by ValidatorSortByFunc) *ProposerPolicy {
	return &ProposerPolicy{Id: id, By: by, registryMU: new(sync.Mutex)}
}

type proposerPolicyToml struct {
	Id ProposerPolicyId
}

func (p *ProposerPolicy) MarshalTOML() (interface{}, error) {
	if p == nil {
		return nil, nil
	}
	pp := &proposerPolicyToml{Id: p.Id}
	data, err := toml.Marshal(pp)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (p *ProposerPolicy) UnmarshalTOML(decode func(interface{}) error) error {
	var innerToml string
	err := decode(&innerToml)
	if err != nil {
		return err
	}
	var pp proposerPolicyToml
	err = toml.Unmarshal([]byte(innerToml), &pp)
	if err != nil {
		return err
	}
	p.Id = pp.Id
	p.By = ValidatorSortByString()
	return nil
}

// Use sets the ValidatorSortByFunc for the given ProposerPolicy and sorts the validatorSets according to it
func (p *ProposerPolicy) Use(v ValidatorSortByFunc) {
	p.By = v

	for _, validatorSet := range p.registry {
		validatorSet.SortValidators()
	}
}

// RegisterValidatorSet stores the given ValidatorSet in the policy registry
func (p *ProposerPolicy) RegisterValidatorSet(valSet ValidatorSet) {
	p.registryMU.Lock()
	defer p.registryMU.Unlock()

	if len(p.registry) == 0 {
		p.registry = []ValidatorSet{valSet}
	} else {
		p.registry = append(p.registry, valSet)
	}
}

// ClearRegistry removes any ValidatorSet from the ProposerPolicy registry
func (p *ProposerPolicy) ClearRegistry() {
	p.registryMU.Lock()
	defer p.registryMU.Unlock()

	p.registry = nil
}

type Config struct {
	RequestTimeout         uint64          `toml:",omitempty"` // The timeout for each Istanbul round in milliseconds.
	BlockPeriod            uint64          `toml:",omitempty"` // Default minimum difference between two consecutive block's timestamps in second
	ProposerPolicy         *ProposerPolicy `toml:",omitempty"` // The policy for proposer selection
	Epoch                  uint64          `toml:",omitempty"` // The number of blocks after which to checkpoint and reset the pending votes
	Ceil2Nby3Block         *big.Int        `toml:",omitempty"` // Number of confirmations required to move from one state to next [2F + 1 to Ceil(2N/3)]
	AllowedFutureBlockTime uint64          `toml:",omitempty"` // Max time (in seconds) from current time allowed for blocks, before they're considered future blocks
	TestQBFTBlock          *big.Int        `toml:",omitempty"` // Fork block at which block confirmations are done using qbft consensus instead of ibft
	Transitions            []params.Transition
	ValidatorContract      common.Address
	Client                 bind.ContractCaller `toml:",omitempty"`
}

var DefaultConfig = &Config{
	RequestTimeout:         10000,
	BlockPeriod:            1,
	ProposerPolicy:         NewRoundRobinProposerPolicy(),
	Epoch:                  30000,
	Ceil2Nby3Block:         big.NewInt(0),
	AllowedFutureBlockTime: 0,
	TestQBFTBlock:          big.NewInt(0),
}

// QBFTBlockNumber returns the qbftBlock fork block number, returns -1 if qbftBlock is not defined
func (c Config) QBFTBlockNumber() int64 {
	if c.TestQBFTBlock == nil {
		return -1
	}
	return c.TestQBFTBlock.Int64()
}

// IsQBFTConsensusAt checks if qbft consensus is enabled for the block height identified by the given header
func (c *Config) IsQBFTConsensusAt(blockNumber *big.Int) bool {
	if c.TestQBFTBlock != nil {
		if c.TestQBFTBlock.Uint64() == 0 {
			return true
		}

		if blockNumber.Cmp(c.TestQBFTBlock) >= 0 {
			return true
		}
	}
	for i := 0; c.Transitions != nil && i < len(c.Transitions) && c.Transitions[i].Block.Cmp(blockNumber) <= 0; i++ {
		if c.Transitions[i].Algorithm == params.QBFT {
			return true
		}
	}
	return false
}

func (c Config) GetConfig(blockNumber *big.Int) Config {
	newConfig := c
	for i := 0; c.Transitions != nil && i < len(c.Transitions) && c.Transitions[i].Block.Cmp(blockNumber) <= 0; i++ {
		if c.Transitions[i].RequestTimeoutSeconds != 0 {
			newConfig.RequestTimeout = c.Transitions[i].RequestTimeoutSeconds
		}
		if c.Transitions[i].EpochLength != 0 {
			newConfig.Epoch = c.Transitions[i].EpochLength
		}
		if c.Transitions[i].BlockPeriodSeconds != 0 {
			newConfig.BlockPeriod = c.Transitions[i].BlockPeriodSeconds
		}
	}
	return newConfig
}

func (c Config) GetValidatorContractAddress(blockNumber *big.Int) common.Address {
	validatorContractAddress := c.ValidatorContract
	for i := 0; c.Transitions != nil && i < len(c.Transitions) && c.Transitions[i].Block.Cmp(blockNumber) <= 0; i++ {
		if c.Transitions[i].ValidatorContractAddress != (common.Address{}) {
			validatorContractAddress = c.Transitions[i].ValidatorContractAddress
		}
	}
	return validatorContractAddress
}

func (c Config) GetValidatorSelectionMode(blockNumber *big.Int) string {
	mode := params.BlockHeaderMode
	for i := 0; c.Transitions != nil && i < len(c.Transitions) && c.Transitions[i].Block.Cmp(blockNumber) <= 0; i++ {
		if c.Transitions[i].ValidatorSelectionMode != "" {
			mode = c.Transitions[i].ValidatorSelectionMode
		}
	}
	return mode
}

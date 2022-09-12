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

package params

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// Quorum - test code size and transaction size limit in chain config
func TestMaxCodeSizeAndTransactionSizeLimit(t *testing.T) {
	type testData struct {
		size  uint64
		valid bool
		err   string
	}
	type testDataType struct {
		isCodeSize bool
		data       []testData
	}

	const codeSizeErr = "Genesis max code size must be between 24 and 128"
	const txSizeErr = "Genesis transaction size limit must be between 32 and 128"
	var codeSizeData = []testData{
		{23, false, codeSizeErr},
		{24, true, ""},
		{50, true, ""},
		{128, true, ""},
		{129, false, codeSizeErr},
	}

	var txSizeData = []testData{
		{31, false, txSizeErr},
		{32, true, ""},
		{50, true, ""},
		{128, true, ""},
		{129, false, txSizeErr},
	}

	var testDataArr = []testDataType{
		{true, codeSizeData},
		{false, txSizeData},
	}

	for _, td := range testDataArr {
		var ccfg *ChainConfig
		for _, d := range td.data {
			var msgPrefix string
			if td.isCodeSize {
				ccfg = &ChainConfig{MaxCodeSize: d.size, TransactionSizeLimit: 50}
				msgPrefix = "max code size"
			} else {
				ccfg = &ChainConfig{MaxCodeSize: 50, TransactionSizeLimit: d.size}
				msgPrefix = "transaction size limit"
			}
			err := ccfg.IsValid()
			if d.valid {
				if err != nil {
					t.Errorf(msgPrefix+" %d, expected no error but got %v", d.size, err)
				}
			} else {
				if err == nil {
					t.Errorf(msgPrefix+" %d, expected error but got none", d.size)
				} else {
					if err.Error() != d.err {
						t.Errorf(msgPrefix+" %d, expected error but got %v", d.size, err.Error())
					}
				}
			}
		}
	}
}

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	var storedMaxCodeConfig0, storedMaxCodeConfig1, storedMaxCodeConfig2 []MaxCodeConfigStruct
	defaultRec := MaxCodeConfigStruct{big.NewInt(0), 24}
	rec1 := MaxCodeConfigStruct{big.NewInt(5), 32}
	rec2 := MaxCodeConfigStruct{big.NewInt(10), 40}
	rec3 := MaxCodeConfigStruct{big.NewInt(8), 40}

	storedMaxCodeConfig0 = append(storedMaxCodeConfig0, defaultRec)

	storedMaxCodeConfig1 = append(storedMaxCodeConfig1, defaultRec)
	storedMaxCodeConfig1 = append(storedMaxCodeConfig1, rec1)
	storedMaxCodeConfig1 = append(storedMaxCodeConfig1, rec2)

	storedMaxCodeConfig2 = append(storedMaxCodeConfig2, rec1)
	storedMaxCodeConfig2 = append(storedMaxCodeConfig2, rec2)

	var passedValidMaxConfig0 []MaxCodeConfigStruct
	passedValidMaxConfig0 = append(passedValidMaxConfig0, defaultRec)
	passedValidMaxConfig0 = append(passedValidMaxConfig0, rec1)

	var passedValidMaxConfig1 []MaxCodeConfigStruct
	passedValidMaxConfig1 = append(passedValidMaxConfig1, defaultRec)
	passedValidMaxConfig1 = append(passedValidMaxConfig1, rec1)
	passedValidMaxConfig1 = append(passedValidMaxConfig1, rec3)

	tests := []test{
		{stored: AllEthashProtocolChanges, new: AllEthashProtocolChanges, head: 0, wantErr: nil},
		{stored: AllEthashProtocolChanges, new: AllEthashProtocolChanges, head: 100, wantErr: nil},
		{
			stored:  &ChainConfig{EIP150Block: big.NewInt(10)},
			new:     &ChainConfig{EIP150Block: big.NewInt(20)},
			head:    9,
			wantErr: nil,
		},
		{
			stored: AllEthashProtocolChanges,
			new:    &ChainConfig{HomesteadBlock: nil},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Homestead fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    nil,
				RewindTo:     0,
			},
		},
		{
			stored: AllEthashProtocolChanges,
			new:    &ChainConfig{HomesteadBlock: big.NewInt(1)},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Homestead fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    big.NewInt(1),
				RewindTo:     0,
			},
		},
		{
			stored: &ChainConfig{HomesteadBlock: big.NewInt(30), EIP150Block: big.NewInt(10)},
			new:    &ChainConfig{HomesteadBlock: big.NewInt(25), EIP150Block: big.NewInt(20)},
			head:   25,
			wantErr: &ConfigCompatError{
				What:         "EIP150 fork block",
				StoredConfig: big.NewInt(10),
				NewConfig:    big.NewInt(20),
				RewindTo:     9,
			},
		},
		{
			stored:  &ChainConfig{ConstantinopleBlock: big.NewInt(30)},
			new:     &ChainConfig{ConstantinopleBlock: big.NewInt(30), PetersburgBlock: big.NewInt(30)},
			head:    40,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{ConstantinopleBlock: big.NewInt(30)},
			new:    &ChainConfig{ConstantinopleBlock: big.NewInt(30), PetersburgBlock: big.NewInt(31)},
			head:   40,
			wantErr: &ConfigCompatError{
				What:         "Petersburg fork block",
				StoredConfig: nil,
				NewConfig:    big.NewInt(31),
				RewindTo:     30,
			},
		},
		{
			stored:  &ChainConfig{Istanbul: &IstanbulConfig{Ceil2Nby3Block: big.NewInt(10)}},
			new:     &ChainConfig{Istanbul: &IstanbulConfig{Ceil2Nby3Block: big.NewInt(20)}},
			head:    4,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{Istanbul: &IstanbulConfig{Ceil2Nby3Block: big.NewInt(10)}},
			new:    &ChainConfig{Istanbul: &IstanbulConfig{Ceil2Nby3Block: big.NewInt(20)}},
			head:   30,
			wantErr: &ConfigCompatError{
				What:         "Ceil 2N/3 fork block",
				StoredConfig: big.NewInt(10),
				NewConfig:    big.NewInt(20),
				RewindTo:     9,
			},
		},
		{
			stored:  &ChainConfig{Istanbul: &IstanbulConfig{TestQBFTBlock: big.NewInt(50)}},
			new:     &ChainConfig{Istanbul: &IstanbulConfig{TestQBFTBlock: big.NewInt(60)}},
			head:    40,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{Istanbul: &IstanbulConfig{TestQBFTBlock: big.NewInt(20)}},
			new:    &ChainConfig{Istanbul: &IstanbulConfig{TestQBFTBlock: big.NewInt(30)}},
			head:   20,
			wantErr: &ConfigCompatError{
				What:         "Test QBFT fork block",
				StoredConfig: big.NewInt(20),
				NewConfig:    big.NewInt(30),
				RewindTo:     19,
			},
		},
		{
			stored: &ChainConfig{MaxCodeSizeChangeBlock: big.NewInt(10)},
			new:    &ChainConfig{MaxCodeSizeChangeBlock: big.NewInt(20)},
			head:   30,
			wantErr: &ConfigCompatError{
				What:         "max code size change fork block",
				StoredConfig: big.NewInt(10),
				NewConfig:    big.NewInt(20),
				RewindTo:     9,
			},
		},
		{
			stored:  &ChainConfig{MaxCodeSizeChangeBlock: big.NewInt(10)},
			new:     &ChainConfig{MaxCodeSizeChangeBlock: big.NewInt(20)},
			head:    4,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{QIP714Block: big.NewInt(10)},
			new:    &ChainConfig{QIP714Block: big.NewInt(20)},
			head:   30,
			wantErr: &ConfigCompatError{
				What:         "permissions fork block",
				StoredConfig: big.NewInt(10),
				NewConfig:    big.NewInt(20),
				RewindTo:     9,
			},
		},
		{
			stored:  &ChainConfig{QIP714Block: big.NewInt(10)},
			new:     &ChainConfig{QIP714Block: big.NewInt(20)},
			head:    4,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig0},
			new:    &ChainConfig{MaxCodeSizeConfig: nil},
			head:   4,
			wantErr: &ConfigCompatError{
				What:         "genesis file missing max code size information",
				StoredConfig: big.NewInt(4),
				NewConfig:    big.NewInt(4),
				RewindTo:     3,
			},
		},
		{
			stored:  &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig0},
			new:     &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig0},
			head:    4,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig0},
			new:    &ChainConfig{MaxCodeSizeConfig: passedValidMaxConfig0},
			head:   10,
			wantErr: &ConfigCompatError{
				What:         "maxCodeSizeConfig data incompatible. updating maxCodeSize for past",
				StoredConfig: big.NewInt(10),
				NewConfig:    big.NewInt(10),
				RewindTo:     9,
			},
		},
		{
			stored:  &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig0},
			new:     &ChainConfig{MaxCodeSizeConfig: passedValidMaxConfig0},
			head:    4,
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig1},
			new:     &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig1},
			head:    12,
			wantErr: nil,
		},
		{
			stored: &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig1},
			new:    &ChainConfig{MaxCodeSizeConfig: passedValidMaxConfig1},
			head:   12,
			wantErr: &ConfigCompatError{
				What:         "maxCodeSizeConfig data incompatible. maxCodeSize historical data does not match",
				StoredConfig: big.NewInt(12),
				NewConfig:    big.NewInt(12),
				RewindTo:     11,
			},
		},
		{
			stored:  &ChainConfig{MaxCodeSize: 32},
			new:     &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig2},
			head:    8,
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{MaxCodeSize: 32},
			new:     &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig2},
			head:    15,
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{MaxCodeSize: 32, MaxCodeSizeChangeBlock: big.NewInt(10)},
			new:     &ChainConfig{MaxCodeSizeConfig: storedMaxCodeConfig1},
			head:    15,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head, false)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}

func TestCheckTransitionsData(t *testing.T) {
	type test struct {
		stored  *ChainConfig
		wantErr error
	}
	var ibftTransitionsConfig, qbftTransitionsConfig, invalidTransition, invalidBlockOrder []Transition

	tranI0 := Transition{big.NewInt(0), IBFT, 30000, 5, 5, 10, 50, common.Address{}, "", nil, nil, nil, nil, 0, nil, 0, nil, nil}
	tranQ5 := Transition{big.NewInt(5), QBFT, 30000, 5, 10, 10, 50, common.Address{}, "", nil, nil, nil, nil, 0, nil, 0, nil, nil}
	tranI10 := Transition{big.NewInt(10), IBFT, 30000, 5, 5, 10, 50, common.Address{}, "", nil, nil, nil, nil, 0, nil, 0, nil, nil}
	tranQ8 := Transition{big.NewInt(8), QBFT, 30000, 5, 10, 10, 50, common.Address{}, "", nil, nil, nil, nil, 0, nil, 0, nil, nil}

	ibftTransitionsConfig = append(ibftTransitionsConfig, tranI0, tranI10)
	qbftTransitionsConfig = append(qbftTransitionsConfig, tranQ5, tranQ8)

	invalidTransition = append(invalidTransition, tranI0, tranQ5, tranI10)
	invalidBlockOrder = append(invalidBlockOrder, tranQ8, tranQ5)

	tests := []test{
		{stored: MainnetChainConfig, wantErr: nil},
		{stored: RopstenChainConfig, wantErr: nil},
		{stored: RinkebyChainConfig, wantErr: nil},
		{stored: GoerliChainConfig, wantErr: nil},
		{stored: YoloV3ChainConfig, wantErr: nil},
		{stored: AllEthashProtocolChanges, wantErr: nil},
		{stored: AllCliqueProtocolChanges, wantErr: nil},
		{stored: TestChainConfig, wantErr: nil},
		{stored: QuorumTestChainConfig, wantErr: nil},
		{stored: QuorumMPSTestChainConfig, wantErr: nil},
		{
			stored:  &ChainConfig{IBFT: &IBFTConfig{}},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{IBFT: &IBFTConfig{}, Transitions: ibftTransitionsConfig},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{QBFT: &QBFTConfig{}},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{QBFT: &QBFTConfig{}, Transitions: qbftTransitionsConfig},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{IBFT: &IBFTConfig{}, Transitions: qbftTransitionsConfig},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{Transitions: ibftTransitionsConfig},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{Transitions: qbftTransitionsConfig},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{IBFT: &IBFTConfig{}, Transitions: invalidTransition},
			wantErr: ErrTransition,
		},
		{
			stored:  &ChainConfig{QBFT: &QBFTConfig{}, Transitions: ibftTransitionsConfig},
			wantErr: ErrTransition,
		},
		{
			stored:  &ChainConfig{Transitions: invalidBlockOrder},
			wantErr: ErrBlockOrder,
		},
		{
			stored:  &ChainConfig{Transitions: []Transition{{nil, IBFT, 30000, 5, 10, 10, 50, common.Address{}, "", nil, nil, nil, nil, 0, nil, 0, nil, nil}}},
			wantErr: ErrBlockNumberMissing,
		},
		{
			stored:  &ChainConfig{Transitions: []Transition{{Block: big.NewInt(0), Algorithm: "AA"}}},
			wantErr: ErrTransitionAlgorithm,
		},
		{
			stored:  &ChainConfig{Transitions: []Transition{{Block: big.NewInt(0), Algorithm: ""}}},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{MaxCodeSizeConfig: []MaxCodeConfigStruct{{big.NewInt(10), 24}}, Transitions: []Transition{{Block: big.NewInt(0), ContractSizeLimit: 50}}},
			wantErr: ErrMaxCodeSizeConfigAndTransitions,
		},
		{
			stored:  &ChainConfig{Transitions: []Transition{{Block: big.NewInt(0), ContractSizeLimit: 23}}},
			wantErr: ErrContractSizeLimit,
		},
		{
			stored:  &ChainConfig{Transitions: []Transition{{Block: big.NewInt(0), ContractSizeLimit: 129}}},
			wantErr: ErrContractSizeLimit,
		},
		{
			stored:  &ChainConfig{Transitions: []Transition{{Block: big.NewInt(0), ContractSizeLimit: 50}}},
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{Transitions: []Transition{{Block: big.NewInt(0)}}},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		err := test.stored.CheckTransitionsData()
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nerr: %v\nwant: %v", test.stored, err, test.wantErr)
		}
	}
}

func TestGetMaxCodeSize(t *testing.T) {
	type test struct {
		config      *ChainConfig
		blockNumber int64
		maxCode     int
	}
	config1, config2, config3 := *TestChainConfig, *TestChainConfig, *TestChainConfig
	config1.MaxCodeSizeConfig = []MaxCodeConfigStruct{
		{big.NewInt(2), 28},
		{big.NewInt(4), 32},
	}
	config1.MaxCodeSize = 34
	config2.MaxCodeSize = 36
	config2.MaxCodeSizeChangeBlock = big.NewInt(2)
	config3.MaxCodeSize = 0
	config3.Transitions = []Transition{
		{Block: big.NewInt(2), ContractSizeLimit: 50},
		{Block: big.NewInt(4), ContractSizeLimit: 54},
	}
	maxCodeDefault := 32 * 1024
	tests := []test{
		{MainnetChainConfig, 0, MaxCodeSize},
		{RopstenChainConfig, 0, MaxCodeSize},
		{RinkebyChainConfig, 0, MaxCodeSize},
		{GoerliChainConfig, 0, MaxCodeSize},
		{YoloV3ChainConfig, 0, MaxCodeSize},
		{AllEthashProtocolChanges, 0, 35 * 1024},
		{AllCliqueProtocolChanges, 0, maxCodeDefault},
		{TestChainConfig, 0, maxCodeDefault},
		{QuorumTestChainConfig, 0, maxCodeDefault},
		{QuorumMPSTestChainConfig, 0, maxCodeDefault},
		{&config1, 0, MaxCodeSize},
		{&config1, 1, MaxCodeSize},
		{&config1, 2, 28 * 1024},
		{&config1, 3, 28 * 1024},
		{&config1, 4, 32 * 1024},
		{&config2, 0, MaxCodeSize},
		{&config2, 1, MaxCodeSize},
		{&config2, 2, 36 * 1024},
		{&config2, 3, 36 * 1024},
		{&config3, 0, MaxCodeSize},
		{&config3, 1, MaxCodeSize},
		{&config3, 2, 50 * 1024},
		{&config3, 3, 50 * 1024},
		{&config3, 4, 54 * 1024},
		{&config3, 8, 54 * 1024},
	}
	for _, test := range tests {
		maxCodeSize := test.config.GetMaxCodeSize(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(maxCodeSize, test.maxCode) {
			t.Errorf("error mismatch:\nexpected: %v\nreceived: %v\n", test.maxCode, maxCodeSize)
		}
	}
}

func TestIsQIP714(t *testing.T) {
	type test struct {
		config      *ChainConfig
		blockNumber int64
		IsQIP714    bool
	}

	config1, config2 := *TestChainConfig, *TestChainConfig
	config1.QIP714Block = big.NewInt(11)

	config2.QIP714Block = nil
	config2.Transitions = []Transition{
		{Block: big.NewInt(21), EnhancedPermissioningEnabled: newPBool(true)},
	}

	tests := []test{
		{MainnetChainConfig, 0, false},
		{&config1, 10, false},
		{&config1, 11, true},
		{&config2, 20, false},
		{&config2, 21, true},
		{&config2, 22, true},
	}

	for _, test := range tests {
		isQIP714 := test.config.IsQIP714(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(isQIP714, test.IsQIP714) {
			t.Errorf("error mismatch on %v:\nexpected: %v\nreceived: %v\n", test.blockNumber, test.IsQIP714, isQIP714)
		}
	}
}

func TestIsPrivacyEnhancementsEnabled(t *testing.T) {
	type test struct {
		config                     *ChainConfig
		blockNumber                int64
		PrivacyEnhancementsEnabled bool
	}

	config1, config2 := *TestChainConfig, *TestChainConfig
	config1.PrivacyEnhancementsBlock = big.NewInt(11)

	config2.PrivacyEnhancementsBlock = nil
	config2.Transitions = []Transition{
		{Block: big.NewInt(21), PrivacyEnhancementsEnabled: newPBool(true)},
	}

	tests := []test{
		{MainnetChainConfig, 0, false},
		{&config1, 10, false},
		{&config1, 11, true},
		{&config2, 20, false},
		{&config2, 21, true},
		{&config2, 22, true},
	}

	for _, test := range tests {
		isPrivacyEnhancementsEnabled := test.config.IsPrivacyEnhancementsEnabled(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(isPrivacyEnhancementsEnabled, test.PrivacyEnhancementsEnabled) {
			t.Errorf("error mismatch on %v:\nexpected: %v\nreceived: %v\n", test.blockNumber, test.PrivacyEnhancementsEnabled, isPrivacyEnhancementsEnabled)
		}
	}
}

func TestIsPrivacyPrecompileEnabled(t *testing.T) {
	type test struct {
		config                   *ChainConfig
		blockNumber              int64
		PrivacyPrecompileEnabled bool
	}

	config1, config2 := *TestChainConfig, *TestChainConfig
	config1.PrivacyPrecompileBlock = big.NewInt(11)

	config2.PrivacyPrecompileBlock = nil
	config2.Transitions = []Transition{
		{Block: big.NewInt(21), PrivacyPrecompileEnabled: newPBool(true)},
	}

	tests := []test{
		{MainnetChainConfig, 0, false},
		{&config1, 10, false},
		{&config1, 11, true},
		{&config2, 20, false},
		{&config2, 21, true},
		{&config2, 22, true},
	}

	for _, test := range tests {
		isPrivacyPrecompileEnabled := test.config.IsPrivacyPrecompileEnabled(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(isPrivacyPrecompileEnabled, test.PrivacyPrecompileEnabled) {
			t.Errorf("error mismatch on %v:\nexpected: %v\nreceived: %v\n", test.blockNumber, test.PrivacyPrecompileEnabled, isPrivacyPrecompileEnabled)
		}
	}
}

func TestIsGasPriceEnabled(t *testing.T) {
	type test struct {
		config          *ChainConfig
		blockNumber     int64
		GasPriceEnabled bool
	}

	config1, config2 := *TestChainConfig, *TestChainConfig
	config1.EnableGasPriceBlock = big.NewInt(11)

	config2.EnableGasPriceBlock = nil
	config2.Transitions = []Transition{
		{Block: big.NewInt(21), GasPriceEnabled: newPBool(true)},
	}

	tests := []test{
		{MainnetChainConfig, 0, false},
		{&config1, 10, false},
		{&config1, 11, true},
		{&config2, 20, false},
		{&config2, 21, true},
		{&config2, 22, true},
	}

	for _, test := range tests {
		isGasPriceEnabled := test.config.IsGasPriceEnabled(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(isGasPriceEnabled, test.GasPriceEnabled) {
			t.Errorf("error mismatch on %v:\nexpected: %v\nreceived: %v\n", test.blockNumber, test.GasPriceEnabled, isGasPriceEnabled)
		}
	}
}

func newPBool(b bool) *bool {
	return &b
}

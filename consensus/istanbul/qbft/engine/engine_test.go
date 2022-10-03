package qbftengine

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	istanbulcommon "github.com/ethereum/go-ethereum/consensus/istanbul/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrepareExtra(t *testing.T) {
	validators := make([]common.Address, 4)
	validators[0] = common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f067778eaf8a"))
	validators[1] = common.BytesToAddress(hexutil.MustDecode("0x294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212"))
	validators[2] = common.BytesToAddress(hexutil.MustDecode("0x6beaaed781d2d2ab6350f5c4566a2c6eaac407a6"))
	validators[3] = common.BytesToAddress(hexutil.MustDecode("0x8be76812f765c24641ec63dc2852b378aba2b440"))

	expectedResult := "0xf87aa00000000000000000000000000000000000000000000000000000000000000000f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b4408080c0"

	h := &types.Header{}
	err := ApplyHeaderQBFTExtra(
		h,
		WriteValidators(validators),
	)
	if err != nil {
		t.Errorf("error mismatch: have %v, want: nil", err)
	}
	result := hexutil.Encode(h.Extra)
	assert.Equal(t, expectedResult, result)
}

func TestWriteCommittedSeals(t *testing.T) {
	istRawData := hexutil.MustDecode("0xf85a80f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b4408080c0")
	expectedCommittedSeal := append([]byte{1, 2, 3}, bytes.Repeat([]byte{0x00}, types.IstanbulExtraSeal-3)...)
	expectedIstExtra := &types.QBFTExtra{
		VanityData: []byte{},
		Validators: []common.Address{
			common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f067778eaf8a")),
			common.BytesToAddress(hexutil.MustDecode("0x294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212")),
			common.BytesToAddress(hexutil.MustDecode("0x6beaaed781d2d2ab6350f5c4566a2c6eaac407a6")),
			common.BytesToAddress(hexutil.MustDecode("0x8be76812f765c24641ec63dc2852b378aba2b440")),
		},
		CommittedSeal: [][]byte{expectedCommittedSeal},
		Round:         make([]byte, 0),
		Vote:          nil,
	}

	h := &types.Header{
		Extra: istRawData,
	}

	// normal case
	err := ApplyHeaderQBFTExtra(
		h,
		writeCommittedSeals([][]byte{expectedCommittedSeal}),
	)
	if err != nil {
		t.Errorf("error mismatch: have %v, want: nil", err)
	}

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}
	if !reflect.DeepEqual(istExtra, expectedIstExtra) {
		t.Errorf("extra data mismatch: have %v, want %v", istExtra, expectedIstExtra)
	}

	// invalid seal
	unexpectedCommittedSeal := append(expectedCommittedSeal, make([]byte, 1)...)
	err = ApplyHeaderQBFTExtra(
		h,
		writeCommittedSeals([][]byte{unexpectedCommittedSeal}),
	)
	if err != istanbulcommon.ErrInvalidCommittedSeals {
		t.Errorf("error mismatch: have %v, want %v", err, istanbulcommon.ErrInvalidCommittedSeals)
	}
}

func TestWriteRoundNumber(t *testing.T) {
	istRawData := hexutil.MustDecode("0xf85a80f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b4408005c0")
	round := make([]byte, 4)
	binary.BigEndian.PutUint32(round, 5)
	expectedIstExtra := &types.QBFTExtra{
		VanityData: []byte{},
		Validators: []common.Address{
			common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f067778eaf8a")),
			common.BytesToAddress(hexutil.MustDecode("0x294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212")),
			common.BytesToAddress(hexutil.MustDecode("0x6beaaed781d2d2ab6350f5c4566a2c6eaac407a6")),
			common.BytesToAddress(hexutil.MustDecode("0x8be76812f765c24641ec63dc2852b378aba2b440")),
		},
		CommittedSeal: [][]byte{},
		Round:         round,
		Vote:          nil,
	}

	var expectedErr error

	h := &types.Header{
		Extra: istRawData,
	}

	// normal case
	err := ApplyHeaderQBFTExtra(
		h,
		writeRoundNumber(big.NewInt(5)),
	)
	if err != expectedErr {
		t.Errorf("error mismatch: have %v, want %v", err, expectedErr)
	}

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}
	if !reflect.DeepEqual(istExtra, expectedIstExtra) {
		t.Errorf("extra data mismatch: have %v, want %v", istExtra.VanityData, expectedIstExtra.VanityData)
	}
}

func TestWriteValidatorVote(t *testing.T) {
	vanity := bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity)
	istRawData := hexutil.MustDecode("0xf85a80f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b4408005c0")
	vote := &types.ValidatorVote{RecipientAddress: common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f06777123456")), VoteType: types.QBFTAuthVote}
	expectedIstExtra := &types.QBFTExtra{
		VanityData:    vanity,
		Validators:    []common.Address{},
		CommittedSeal: [][]byte{},
		Round:         make([]byte, 0),
		Vote:          vote,
	}

	var expectedErr error

	h := &types.Header{
		Extra: istRawData,
	}

	// normal case
	err := ApplyHeaderQBFTExtra(
		h,
		WriteVote(common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f06777123456")), true),
	)
	if err != expectedErr {
		t.Errorf("error mismatch: have %v, want %v", err, expectedErr)
	}

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}
	if !reflect.DeepEqual(istExtra.Vote, expectedIstExtra.Vote) {
		t.Errorf("extra data mismatch: have %v, want %v", istExtra, expectedIstExtra)
	}
}

func TestAccumulateRewards(t *testing.T) {
	addr := common.HexToAddress("0xed9d02e382b34818e88b88a309c7fe71e65f419d")
	fixedMode := "fixed"
	listMode := "list"
	validatorsMode := "validators"
	emptyMode := ""
	dummyMode := "dummy"
	m := []struct {
		addr              common.Address
		miningBeneficiary *common.Address
		balance           *big.Int
		blockReward       *math.HexOrDecimal256
		mode              *string
		list              []common.Address
		expectedBalance   *big.Int
	}{
		{ // off
			addr:              addr,
			miningBeneficiary: nil,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              nil,
			list:              nil,
			expectedBalance:   big.NewInt(1),
		},
		{ // auto/default
			addr:              addr,
			miningBeneficiary: &addr,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              nil,
			list:              nil,
			expectedBalance:   big.NewInt(2),
		},
		{ // failing
			addr:              addr,
			miningBeneficiary: nil,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              &fixedMode,
			list:              nil,
			expectedBalance:   big.NewInt(1),
		},
		{
			addr:              addr,
			miningBeneficiary: &addr,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              &fixedMode,
			list:              nil,
			expectedBalance:   big.NewInt(2),
		},
		{ // failing
			addr:              addr,
			miningBeneficiary: nil,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              &listMode,
			list:              nil,
			expectedBalance:   big.NewInt(1),
		},
		{
			addr:              addr,
			miningBeneficiary: nil,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              &listMode,
			list:              []common.Address{addr},
			expectedBalance:   big.NewInt(2),
		},
		{
			addr:              addr,
			miningBeneficiary: nil,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              &validatorsMode,
			expectedBalance:   big.NewInt(1),
		},
		{
			addr:              addr,
			miningBeneficiary: nil,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              &emptyMode,
			expectedBalance:   big.NewInt(1),
		},
		{
			addr:              addr,
			miningBeneficiary: nil,
			balance:           big.NewInt(1),
			blockReward:       math.NewHexOrDecimal256(1),
			mode:              &dummyMode,
			expectedBalance:   big.NewInt(1),
		},
	}
	var e *Engine
	chain := &core.BlockChain{}
	db := state.NewDatabaseWithConfig(rawdb.NewMemoryDatabase(), nil)
	state, err := state.New(common.Hash{}, db, nil)
	require.NoError(t, err)

	header := &types.Header{
		Number: big.NewInt(1),
	}
	for idx, te := range m {
		if te.mode == &validatorsMode {
			continue // skip, it's not testable yet
		}
		state.SetBalance(te.addr, te.balance)
		cfg := istanbul.Config{
			BlockReward:       te.blockReward,
			BeneficiaryMode:   te.mode,
			BeneficiaryList:   te.list,
			MiningBeneficiary: te.miningBeneficiary,
		}
		e.accumulateRewards(chain, state, header, nil, cfg)
		balance := state.GetBalance(te.addr)
		assert.Equal(t, te.expectedBalance, balance, fmt.Sprintf("index: %d", idx), te)
	}
}

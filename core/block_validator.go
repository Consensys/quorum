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

package core

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/params"
	"gopkg.in/fatih/set.v0"
)

func forceParseRfc3339(str string) time.Time {
	time, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic("unexpected failure to parse rfc3339 timestamp: " + str)
	}
	return time
}

var (
	ExpDiffPeriod = big.NewInt(100000)
	big10         = big.NewInt(10)
	bigMinus99    = big.NewInt(-99)
	nanosecond2017Timestamp = forceParseRfc3339("2017-01-01T00:00:00+00:00").UnixNano()
)

// BlockValidator is responsible for validating block headers, uncles and
// processed state.
//
// BlockValidator implements Validator.
type BlockValidator struct {
	chaindb            ethdb.Database
	config             *ChainConfig // Chain configuration options
	bc                 *BlockChain  // Canonical block chain
	enableQuorumChecks bool         // indication if the signature and vote count is checked (disabled for testing puposes)
}

// NewBlockValidator returns a new block validator which is safe for re-use
func NewBlockValidator(chaindb ethdb.Database, config *ChainConfig, blockchain *BlockChain, enableQuorumChecks bool) *BlockValidator {
	validator := &BlockValidator{
		chaindb:            chaindb,
		config:             config,
		bc:                 blockchain,
		enableQuorumChecks: enableQuorumChecks,
	}
	return validator
}

// ValidateBlock validates the given block's header and uncles and verifies the
// the block header's transaction and uncle roots.
//
// ValidateBlock does not validate the header's finaliser. The finaliser work validated
// separately so we can process them in parallel.
//
// ValidateBlock also validates and makes sure that any previous state (or present)
// state that might or might not be present is checked to make sure that fast
// sync has done it's job proper. This prevents the block validator form accepting
// false positives where a header is present but the state is not.
func (v *BlockValidator) ValidateBlock(block *types.Block) error {
	if v.bc.HasBlock(block.Hash()) {
		if _, err := state.New(block.Root(), v.bc.chainDb); err == nil {
			return &KnownBlockError{block.Number(), block.Hash()}
		}
	}
	parent := v.bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		return ParentError(block.ParentHash())
	}
	if _, err := state.New(parent.Root(), v.bc.chainDb); err != nil {
		return ParentError(block.ParentHash())
	}

	header := block.Header()
	// validate the block header
	if err := ValidateHeader(v.chaindb, v.bc, v.config, header, parent.Header(), false, v.enableQuorumChecks); err != nil {
		return err
	}
	// verify the uncles are correctly rewarded
	if err := v.VerifyUncles(block, parent); err != nil {
		return err
	}

	// Verify UncleHash before running other uncle validations
	unclesSha := types.CalcUncleHash(block.Uncles())
	if unclesSha != header.UncleHash {
		return fmt.Errorf("invalid uncles root hash. received=%x calculated=%x", header.UncleHash, unclesSha)
	}

	// The transactions Trie's root (R = (Tr [[i, RLP(T1)], [i, RLP(T2)], ... [n, RLP(Tn)]]))
	// can be used by light clients to make sure they've received the correct Txs
	txSha := types.DeriveSha(block.Transactions())
	if txSha != header.TxHash {
		return fmt.Errorf("invalid transaction root hash. received=%x calculated=%x", header.TxHash, txSha)
	}

	return nil
}

// callmsg is the message type used for call transactions.
type callmsg struct {
	from          *state.StateObject
	to            *common.Address
	gas, gasPrice *big.Int
	value         *big.Int
	data          []byte
}

// accessor boilerplate to implement core.Message
func (m callmsg) From() (common.Address, error)         { return m.from.Address(), nil }
func (m callmsg) FromFrontier() (common.Address, error) { return m.from.Address(), nil }
func (m callmsg) Nonce() uint64                         { return m.from.Nonce() }
func (m callmsg) To() *common.Address                   { return m.to }
func (m callmsg) GasPrice() *big.Int                    { return m.gasPrice }
func (m callmsg) Gas() *big.Int                         { return m.gas }
func (m callmsg) Value() *big.Int                       { return m.value }
func (m callmsg) Data() []byte                          { return m.data }
func (m callmsg) CheckNonce() bool                      { return true }

// ValidateState validates the various changes that happen after a state
// transition, such as amount of used gas, the receipt roots and the state root
// itself. For quorum it also verifies if the canonical hash in the blocks state
// points to a valid parent hash.
//
// ValidateState returns a database batch if the validation was a success
// otherwise nil and an error is returned.
func (v *BlockValidator) ValidateState(block, parent *types.Block, statedb *state.StateDB, receipts types.Receipts, usedGas *big.Int) (err error) {
	header := block.Header()
	if block.GasUsed().Cmp(usedGas) != 0 {
		return ValidationError(fmt.Sprintf("gas used error (%v / %v)", block.GasUsed(), usedGas))
	}
	// Validate the received block's bloom with the one derived from the generated receipts.
	// For valid blocks this should always validate to true.
	rbloom := types.CreateBloom(receipts)
	if rbloom != header.Bloom {
		return fmt.Errorf("unable to replicate block's bloom=%x vs calculated bloom=%x", header.Bloom, rbloom)
	}
	// Tre receipt Trie's root (R = (Tr [[H1, R1], ... [Hn, R1]]))
	receiptSha := types.DeriveSha(receipts)
	if receiptSha != header.ReceiptHash {
		return fmt.Errorf("invalid receipt root hash. received=%x calculated=%x", header.ReceiptHash, receiptSha)
	}

	// Validate the state root against the received state root and throw
	// an error if they don't match.
	if root := statedb.IntermediateRoot(); header.Root != root {
		return fmt.Errorf("invalid merkle root: header=%x computed=%x", header.Root, root)
	}

	if v.enableQuorumChecks {
		// Ensure that the parent block was indeed the one that was voted for in the state of this block.
		// The contract enforces that there are enough votes and only votes from parties that are allowed to vote.
		var (
			gp        = new(GasPool).AddGas(common.MaxBig)
			to        = common.HexToAddress("0x0000000000000000000000000000000000000020")
			stateCopy = statedb.Copy()
			msg       = callmsg{
				from:     stateCopy.GetOrNewStateObject(common.HexToAddress("0x0000000000000000000000000000000000000000")),
				to:       &to,
				gas:      big.NewInt(500000),
				gasPrice: common.Big0,
				value:    common.Big0,
				data:     common.Hex2Bytes(fmt.Sprintf("559c390c%064x", block.Number())), // call getCanonHash(uint256)
			}
			vmenv = NewEnv(stateCopy, stateCopy, v.config, v.bc, msg, block.Header(), v.config.VmConfig)
		)

		result, _, _, err := NewStateTransition(vmenv, msg, gp).TransitionDb()
		if err != nil {
			return err
		}

		// result holds the hash that was the winning hash according the voting contract
		parentHash := common.BytesToHash(result)
		if parentHash == (common.Hash{}) {
			// too little votes
			return fmt.Errorf("block parent could not be verified, ignore block (%d)", block.Number())
		}
		if block.ParentHash() != parentHash {
			return fmt.Errorf("build on top of unexpected parent, expected %s, got %s", parentHash.Hex(), block.ParentHash().Hex())
		}
	}

	return nil
}

// VerifyUncles verifies the given block's uncles and applies the Ethereum
// consensus rules to the various block headers included; it will return an
// error if any of the included uncle headers were invalid. It returns an error
// if the validation failed.
func (v *BlockValidator) VerifyUncles(block, parent *types.Block) error {
	// validate that there at most 2 uncles included in this block
	if len(block.Uncles()) > 2 {
		return ValidationError("Block can only contain maximum 2 uncles (contained %v)", len(block.Uncles()))
	}

	uncles := set.New()
	ancestors := make(map[common.Hash]*types.Block)
	for _, ancestor := range v.bc.GetBlocksFromHash(block.ParentHash(), 7) {
		ancestors[ancestor.Hash()] = ancestor
		// Include ancestors uncles in the uncle set. Uncles must be unique.
		for _, uncle := range ancestor.Uncles() {
			uncles.Add(uncle.Hash())
		}
	}
	ancestors[block.Hash()] = block
	uncles.Add(block.Hash())

	for i, uncle := range block.Uncles() {
		hash := uncle.Hash()
		if uncles.Has(hash) {
			// Error not unique
			return UncleError("uncle[%d](%x) not unique", i, hash[:4])
		}
		uncles.Add(hash)

		if ancestors[hash] != nil {
			branch := fmt.Sprintf("  O - %x\n  |\n", block.Hash())
			for h := range ancestors {
				branch += fmt.Sprintf("  O - %x\n  |\n", h)
			}
			glog.Infoln(branch)
			return UncleError("uncle[%d](%x) is ancestor", i, hash[:4])
		}

		if ancestors[uncle.ParentHash] == nil || uncle.ParentHash == parent.Hash() {
			return UncleError("uncle[%d](%x)'s parent is not ancestor (%x)", i, hash[:4], uncle.ParentHash[0:4])
		}

		if err := ValidateHeader(v.chaindb, v.bc, v.config, uncle, ancestors[uncle.ParentHash].Header(), true, v.enableQuorumChecks); err != nil {
			return ValidationError(fmt.Sprintf("uncle[%d](%x) header invalid: %v", i, hash[:4], err))
		}
	}

	return nil
}

// ValidateHeader validates the given header and, depending on the finaliser arg,
// checks the proof of work of the given header. Returns an error if the
// validation failed.
func (v *BlockValidator) ValidateHeader(chaindb ethdb.Database, header, parent *types.Header) error {
	// Short circuit if the parent is missing.
	if parent == nil {
		return ParentError(header.ParentHash)
	}
	// Short circuit if the header's already known or its parent missing
	if v.bc.HasHeader(header.Hash()) {
		return nil
	}
	return ValidateHeader(chaindb, v.bc, v.config, header, parent, false, v.enableQuorumChecks)
}

// Validates a header. Returns an error if the header is invalid and verify if the
// block signature is from an allowed block creator.
//
// See YP section 4.3.4. "Block Header Validity"
func ValidateHeader(chaindb ethdb.Database, bc *BlockChain, config *ChainConfig, header *types.Header, parent *types.Header, uncle, validateSignature bool) error {
	if big.NewInt(int64(len(header.Extra))).Cmp(params.MaximumExtraDataSize) == 1 {
		return fmt.Errorf("Header extra data too long (%d)", len(header.Extra))
	}

	if uncle {
		if header.Time.Cmp(common.MaxBig) == 1 {
			return BlockTSTooBigErr
		}
	} else {
		// We disable future checking if we're in --raft mode. This is crucial
		// because block validation in the raft setting needs to be deterministic.
		// There is no forking of the chain, and we need each node to only perform
		// validation as a pure function of block contents with respect to the
		// previous database state.
		//
		// NOTE: whereas we are currently checking whether the timestamp field has
		// nanosecond semantics to detect --raft mode, we could also use a special
		// "raft" sentinel in the Extra field, or pass a boolean for raftMode from
		// all call sites of this function.
		if raftMode := time.Now().UnixNano() > nanosecond2017Timestamp; !raftMode {
			if header.Time.Cmp(big.NewInt(time.Now().Unix())) == 1 {
				return BlockFutureErr
			}
		}
	}
	if header.Time.Cmp(parent.Time) != 1 {
		return BlockEqualTSErr
	}

	expd := CalcDifficulty(config, header.Time.Uint64(), parent.Time.Uint64(), parent.Number, parent.Difficulty)
	if expd.Cmp(header.Difficulty) != 0 {
		return fmt.Errorf("Difficulty check failed for header %v, %v", header.Difficulty, expd)
	}

	a := new(big.Int).Set(parent.GasLimit)
	a = a.Sub(a, header.GasLimit)
	a.Abs(a)
	b := new(big.Int).Set(parent.GasLimit)
	b = b.Div(b, params.GasLimitBoundDivisor)
	if !(a.Cmp(b) < 0) || (header.GasLimit.Cmp(params.MinGasLimit) == -1) {
		return fmt.Errorf("GasLimit check failed for header %v (%v > %v)", header.GasLimit, a, b)
	}

	num := new(big.Int).Set(parent.Number)
	num.Sub(header.Number, num)
	if num.Cmp(big.NewInt(1)) != 0 {
		return BlockNumberErr
	}

	if validateSignature {
		return ValidateExtraData(chaindb, bc, config, parent, header)
	}
	return nil
}

// ValidateExtraData verifies the signature in the extra data field and ensures the signer is allowed to create blocks.
//
// In Quorum blocks the Extra data field contains a signature created by the block creator.
// This signature is used to verify that the block is created by a party that is allowed to create blocks.
func ValidateExtraData(chaindb ethdb.Database, bc *BlockChain, config *ChainConfig, parent, header *types.Header) error {
	var (
		hash      = header.QuorumHash()
		signature = header.Extra
		addr      = header.Coinbase
	)

	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return err
	}

	signerAddr := crypto.PubkeyToAddress(*pubKey)
	if signerAddr != addr {
		return fmt.Errorf("invalid header signature %s != %s", signerAddr.Hex(), addr.Hex())
	}

	// Ensure that the recovered address belongs to an account this is allowed to create blocks.
	var (
		state, _ = state.New(parent.Root, chaindb)
		gp       = new(GasPool).AddGas(common.MaxBig)
		to       = common.HexToAddress("0x0000000000000000000000000000000000000020")
		msg      = callmsg{
			from:     state.GetOrNewStateObject(common.HexToAddress("0x0000000000000000000000000000000000000000")),
			to:       &to,
			gas:      big.NewInt(500000),
			gasPrice: common.Big0,
			value:    common.Big0,
			data:     common.Hex2Bytes(fmt.Sprintf("e814d1c7%064x", signerAddr.Bytes())), // call isBlockMaker(address)
		}
		vmenv = NewEnv(state, state, config, bc, msg, header, config.VmConfig)
	)

	result, _, _, err := NewStateTransition(vmenv, msg, gp).TransitionDb()
	if err != nil {
		return err
	}

	if config.HomesteadGasRepriceBlock != nil && config.HomesteadGasRepriceBlock.Cmp(header.Number) == 0 {
		if config.HomesteadGasRepriceHash != (common.Hash{}) && config.HomesteadGasRepriceHash != header.Hash() {
			return ValidationError("Homestead gas reprice fork hash mismatch: have 0x%x, want 0x%x", header.Hash(), config.HomesteadGasRepriceHash)
		}
	}

	if new(big.Int).SetBytes(result).Cmp(common.Big1) == 0 {
		return nil
	}

	return fmt.Errorf("Invalid header: %s isn't allowed to create blocks", signerAddr.Hex())
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time
// given the parent block's time and difficulty.
func CalcDifficulty(config *ChainConfig, time, parentTime uint64, parentNumber, parentDiff *big.Int) *big.Int {
	if config.IsHomestead(new(big.Int).Add(parentNumber, common.Big1)) {
		return calcDifficultyHomestead(time, parentTime, parentNumber, parentDiff)
	} else {
		return calcDifficultyFrontier(time, parentTime, parentNumber, parentDiff)
	}
}

func calcDifficultyHomestead(time, parentTime uint64, parentNumber, parentDiff *big.Int) *big.Int {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.mediawiki
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	//        ) + 2^(periodCount - 2)

	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).SetUint64(parentTime)

	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)

	// 1 - (block_timestamp -parent_timestamp) // 10
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big10)
	x.Sub(common.Big1, x)

	// max(1 - (block_timestamp - parent_timestamp) // 10, -99)))
	if x.Cmp(bigMinus99) < 0 {
		x.Set(bigMinus99)
	}

	// (parent_diff + parent_diff // 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	y.Div(parentDiff, params.DifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parentDiff, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}

	// for the exponential factor
	periodCount := new(big.Int).Add(parentNumber, common.Big1)
	periodCount.Div(periodCount, ExpDiffPeriod)

	// the exponential factor, commonly referred to as "the bomb"
	// diff = diff + 2^(periodCount - 2)
	if periodCount.Cmp(common.Big1) > 0 {
		y.Sub(periodCount, common.Big2)
		y.Exp(common.Big2, y, nil)
		x.Add(x, y)
	}

	return x
}

func calcDifficultyFrontier(time, parentTime uint64, parentNumber, parentDiff *big.Int) *big.Int {
	diff := new(big.Int)
	adjust := new(big.Int).Div(parentDiff, params.DifficultyBoundDivisor)
	bigTime := new(big.Int)
	bigParentTime := new(big.Int)

	bigTime.SetUint64(time)
	bigParentTime.SetUint64(parentTime)

	if bigTime.Sub(bigTime, bigParentTime).Cmp(params.DurationLimit) < 0 {
		diff.Add(parentDiff, adjust)
	} else {
		diff.Sub(parentDiff, adjust)
	}
	if diff.Cmp(params.MinimumDifficulty) < 0 {
		diff.Set(params.MinimumDifficulty)
	}

	periodCount := new(big.Int).Add(parentNumber, common.Big1)
	periodCount.Div(periodCount, ExpDiffPeriod)
	if periodCount.Cmp(common.Big1) > 0 {
		// diff = diff + 2^(periodCount - 2)
		expDiff := periodCount.Sub(periodCount, common.Big2)
		expDiff.Exp(common.Big2, expDiff, nil)
		diff.Add(diff, expDiff)
		diff = common.BigMax(diff, params.MinimumDifficulty)
	}

	return diff
}

// CalcGasLimit computes the gas limit of the next block after parent.
// The result may be modified by the caller.
// This is miner strategy, not consensus protocol.
func CalcGasLimit(parent *types.Block) *big.Int {
	// contrib = (parentGasUsed * 3 / 2) / 4096
	contrib := new(big.Int).Mul(parent.GasUsed(), big.NewInt(3))
	contrib = contrib.Div(contrib, big.NewInt(2))
	contrib = contrib.Div(contrib, params.GasLimitBoundDivisor)

	// decay = parentGasLimit / 1024 -1
	decay := new(big.Int).Div(parent.GasLimit(), params.GasLimitBoundDivisor)
	decay.Sub(decay, big.NewInt(1))

	/*
		strategy: gasLimit of block-to-mine is set based on parent's
		gasUsed value.  if parentGasUsed > parentGasLimit * (2/3) then we
		increase it, otherwise lower it (or leave it unchanged if it's right
		at that usage) the amount increased/decreased depends on how far away
		from parentGasLimit * (2/3) parentGasUsed is.
	*/
	gl := new(big.Int).Sub(parent.GasLimit(), decay)
	gl = gl.Add(gl, contrib)
	gl.Set(common.BigMax(gl, params.MinGasLimit))

	// however, if we're now below the target (TargetGasLimit) we increase the
	// limit as much as we can (parentGasLimit / 1024 -1)
	if gl.Cmp(params.TargetGasLimit) < 0 {
		gl.Add(parent.GasLimit(), decay)
		gl.Set(common.BigMin(gl, params.TargetGasLimit))
	}
	return gl
}

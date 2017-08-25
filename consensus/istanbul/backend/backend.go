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

package backend

import (
	"crypto/ecdsa"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	istanbulCore "github.com/ethereum/go-ethereum/consensus/istanbul/core"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
)

// New creates an Ethereum backend for Istanbul core engine.
func New(config *istanbul.Config, privateKey *ecdsa.PrivateKey, db ethdb.Database) consensus.Istanbul {
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	backend := &backend{
		config:           config,
		istanbulEventMux: new(event.TypeMux),
		privateKey:       privateKey,
		address:          crypto.PubkeyToAddress(privateKey.PublicKey),
		logger:           log.New(),
		db:               db,
		commitCh:         make(chan *types.Block, 1),
		recents:          recents,
		candidates:       make(map[common.Address]bool),
		coreStarted:      false,
	}
	backend.core = istanbulCore.New(backend, backend.config)
	return backend
}

// ----------------------------------------------------------------------------

type backend struct {
	config           *istanbul.Config
	istanbulEventMux *event.TypeMux
	privateKey       *ecdsa.PrivateKey
	address          common.Address
	core             istanbulCore.Engine
	logger           log.Logger
	db               ethdb.Database
	chain            consensus.ChainReader
	inserter         func(types.Blocks) (int, error)

	// the channels for istanbul engine notifications
	commitCh          chan *types.Block
	proposedBlockHash common.Hash
	sealMu            sync.Mutex
	coreStarted       bool
	coreMu            sync.Mutex

	// Current list of candidates we are pushing
	candidates map[common.Address]bool
	// Protects the signer fields
	candidatesLock sync.RWMutex
	// Snapshots for recent block to speed up reorgs
	recents *lru.ARCCache

	// event subscription for ChainHeadEvent event
	broadcaster consensus.Broadcaster
}

// Address implements istanbul.Backend.Address
func (sb *backend) Address() common.Address {
	return sb.address
}

// Validators implements istanbul.Backend.Validators
func (sb *backend) Validators(proposal istanbul.Proposal) istanbul.ValidatorSet {
	return sb.getValidators(proposal.Number().Uint64(), proposal.Hash())
}

// Broadcast implements istanbul.Backend.Broadcast
func (sb *backend) Broadcast(valSet istanbul.ValidatorSet, payload []byte) error {
	// send to others
	sb.Gossip(valSet, payload)
	// send to self
	msg := istanbul.MessageEvent{
		Payload: payload,
	}
	go sb.istanbulEventMux.Post(msg)
	return nil
}

func (sb *backend) Gossip(valSet istanbul.ValidatorSet, payload []byte) error {
	targets := make(map[common.Address]bool)
	for _, val := range valSet.List() {
		if val.Address() != sb.Address() {
			targets[val.Address()] = true
		}
	}

	if sb.broadcaster != nil && len(targets) > 0 {
		ps := sb.broadcaster.FindPeers(targets)
		for _, p := range ps {
			go p.Send(istanbulMsg, payload)
		}
	}
	return nil
}

// Commit implements istanbul.Backend.Commit
func (sb *backend) Commit(proposal istanbul.Proposal, seals [][]byte) error {
	// Check if the proposal is a valid block
	block := &types.Block{}
	block, ok := proposal.(*types.Block)
	if !ok {
		sb.logger.Error("Invalid proposal, %v", proposal)
		return errInvalidProposal
	}

	h := block.Header()
	// Append seals into extra-data
	err := writeCommittedSeals(h, seals)
	if err != nil {
		return err
	}
	// update block's header
	block = block.WithSeal(h)

	sb.logger.Info("Committed", "address", sb.Address(), "hash", proposal.Hash(), "number", proposal.Number().Uint64())
	// - if the proposed and committed blocks are the same, send the proposed hash
	//   to commit channel, which is being watched inside the engine.Seal() function.
	// - otherwise, we try to insert the block.
	// -- if success, the ChainHeadEvent event will be broadcasted, try to build
	//    the next block and the previous Seal() will be stopped.
	// -- otherwise, a error will be returned and a round change event will be fired.
	if sb.proposedBlockHash == block.Hash() {
		// feed block hash to Seal() and wait the Seal() result
		sb.commitCh <- block
		// TODO: how do we check the block is inserted correctly?
		return nil
	}
	// if I'm not a proposer, insert the block directly and broadcast NewCommittedEvent
	if _, err := sb.inserter(types.Blocks{block}); err != nil && err != core.ErrKnownBlock {
		return err
	}

	if sb.broadcaster != nil {
		go sb.broadcaster.BroadcastBlock(block, false)
	}
	return nil
}

// EventMux implements istanbul.Backend.EventMux
func (sb *backend) EventMux() *event.TypeMux {
	return sb.istanbulEventMux
}

// Verify implements istanbul.Backend.Verify
func (sb *backend) Verify(proposal istanbul.Proposal) (time.Duration, error) {
	// Check if the proposal is a valid block
	block := &types.Block{}
	block, ok := proposal.(*types.Block)
	if !ok {
		sb.logger.Error("Invalid proposal, %v", proposal)
		return 0, errInvalidProposal
	}
	// verify the header of proposed block
	err := sb.VerifyHeader(sb.chain, block.Header(), false)
	// ignore errEmptyCommittedSeals error because we don't have the committed seals yet
	if err == nil || err == errEmptyCommittedSeals {
		return 0, nil
	} else if err == consensus.ErrFutureBlock {
		return time.Unix(block.Header().Time.Int64(), 0).Sub(now()), consensus.ErrFutureBlock
	}
	return 0, err
}

// Sign implements istanbul.Backend.Sign
func (sb *backend) Sign(data []byte) ([]byte, error) {
	hashData := crypto.Keccak256([]byte(data))
	return crypto.Sign(hashData, sb.privateKey)
}

// CheckSignature implements istanbul.Backend.CheckSignature
func (sb *backend) CheckSignature(data []byte, address common.Address, sig []byte) error {
	signer, err := istanbul.GetSignatureAddress(data, sig)
	if err != nil {
		log.Error("Failed to get signer address", "err", err)
		return err
	}
	// Compare derived addresses
	if signer != address {
		return errInvalidSignature
	}
	return nil
}

// GetProposer implements istanbul.Backend.GetProposer
func (sb *backend) GetProposer(number uint64) common.Address {
	if h := sb.chain.GetHeaderByNumber(number); h != nil {
		a, _ := sb.Author(h)
		return a
	}
	return common.Address{}
}

func (sb *backend) getValidators(number uint64, hash common.Hash) istanbul.ValidatorSet {
	snap, err := sb.snapshot(sb.chain, number, hash, nil)
	if err != nil {
		return validator.NewSet(nil, sb.config.ProposerPolicy)
	}
	return snap.ValSet
}

func (sb *backend) LastProposal() (istanbul.Proposal, common.Address) {
	if sb.chain == nil {
		sb.logger.Error("Failed to access blockchain")
		return nil, common.Address{}
	}

	h := sb.chain.CurrentHeader()

	var proposer common.Address
	if h.Number.Cmp(common.Big0) > 0 {
		var err error
		proposer, err = sb.Author(h)
		if err != nil {
			sb.logger.Error("Failed to get block proposer", "err", err)
			return nil, common.Address{}
		}
	}

	// Return header only block here since we don't need block body
	return types.NewBlockWithHeader(h), proposer
}

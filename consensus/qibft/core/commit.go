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

package core

import (
	"reflect"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

func (c *core) sendCommit() {
	sub := c.current.Subject()
	c.broadcastCommitMsg(sub)
}

func (c *core) sendCommitForOldBlock(view *View, digest common.Hash) {
	sub := &Subject{
		View:   view,
		Digest: digest,
	}
	c.broadcastCommitMsg(sub)
}

func (c *core) broadcastCommit(sub *Subject) {
	logger := c.logger.New("state", c.state)

	encodedSubject, err := Encode(sub)
	if err != nil {
		logger.Error("Failed to encode", "subject", sub)
		return
	}

	logger.Info("QBFT: sendCommit", "m", sub)
	c.broadcast(&message{
		Code: msgCommit,
		Msg:  encodedSubject,
	})
}

func (c *core) broadcastCommitMsg(sub *Subject) {
	var err error

	logger := c.logger.New("state", c.state)

	commitMsg := &CommitMsg{
		CommonMsg:  CommonMsg{
			code:           commitMsgCode,
			source: c.address,
			Sequence:       sub.View.Sequence,
			Round:          sub.View.Round,
		},
		Digest:     sub.Digest,
	}

	// Add Commit Seal
	seal := PrepareCommittedSeal(sub.Digest)
	commitMsg.CommitSeal, err = c.backend.Sign(seal)
	if err != nil {
		logger.Error("QBFT: Failed to create commit seal", "msg", commitMsg, "err", err)
		return
	}

	// Sign Message
	encodedPayload, err := commitMsg.EncodedPayload()
	if err != nil {
		logger.Error("QBFT: Failed to encode payload of commit message", "msg", commitMsg, "err", err)
		return
	}
	commitMsg.Signature, err = c.backend.Sign(encodedPayload)
	if err != nil {
		logger.Error("QBFT: Failed to sign commit message", "msg", commitMsg, "err", err)
		return
	}

	// RLP-encode message
	payload, err := rlp.EncodeToBytes(&commitMsg)
	if err != nil {
		logger.Error("QBFT: Failed to encode commit message", "msg", commitMsg, "err", err)
		return
	}

	logger.Info("QBFT: broadcastCommitMsg", "m", sub, "payload", payload)
	// Broadcast RLP-encoded message
	if err = c.backend.Broadcast(c.valSet, commitMsgCode, payload); err != nil {
		logger.Error("QBFT: Failed to broadcast message", "msg", commitMsg, "err", err)
		return
	}
}

func (c *core) handleCommitMsg(commit *CommitMsg) error {
	logger := c.logger.New("state", c.state)

	logger.Info("QBFT: handleCommitMsg", "msg", &commit)

	// Verify signature
	var err error
	var payload []byte
	if payload, err = commit.EncodedPayload(); err != nil {
		logger.Error("QBFT: Error encoding payload", "err", err)
		return errInvalidMessage
	}

	var source common.Address
	if source, err = c.validateFn(payload, commit.Signature); err != nil {
		logger.Error("QBFT: Error checking signature", "err", err)
		return errInvalidSigner
	}
	commit.source = source

	// Check view and state
	view := commit.View()
	if err := c.checkMessage(msgCommit, &view); err != nil {
		logger.Error("QBFT: Failed to check message", "err", err)
		return err
	}

	// Check digest
	if commit.Digest != c.current.Proposal().Hash() {
		logger.Error("QBFT: Failed to check digest")
		return errInvalidMessage
	}

	// Add to received msgs
	if err := c.current.QBFTCommits.Add(commit); err != nil {
		c.logger.Error("QBFT: Failed to save commit message", "msg", commit, "err", err)
		return err
	}

	logger.Info("QBFT: commit threshold", "commits", c.current.QBFTCommits.Size(), "quorum", c.QuorumSize())
	// Check threshold and decide
	if c.current.QBFTCommits.Size() >= c.QuorumSize() {
		logger.Info("QBFT: Reached commit threshold")
		c.commitQBFT()
	}

	return nil
}

func (c *core) handleCommit(msg *message, src istanbul.Validator) error {
	// Decode COMMIT message
	var commit *Subject
	err := msg.Decode(&commit)
	if err != nil {
		return errFailedDecodeCommit
	}

	if err := c.checkMessage(msgCommit, commit.View); err != nil {
		return err
	}

	// Check if a quorum of prepare message have been received corresponding to the commit digest. If not, then return a errInvalidMessage error
	if commit.Digest != c.current.Proposal().Hash() {
		c.logger.Error("commit digest does not match proposal hash", "proposal", c.current.Proposal().Hash(), "commit", commit.Digest, "state", c.state)
		return errInvalidMessage
	}

	if err := c.verifyCommit(commit, src); err != nil {
		return err
	}

	c.acceptCommit(msg, src)

	// Commit the proposal once we have enough COMMIT messages and we are not in the Committed state.
	//
	// If we already have a proposal, we may have chance to speed up the consensus process
	// by committing the proposal without PREPARE messages.
	if c.current.Commits.Size() >= c.QuorumSize() && c.state.Cmp(StateCommitted) < 0 {
		c.commit()
	}

	return nil
}

// verifyCommit verifies if the received COMMIT message is equivalent to our subject
func (c *core) verifyCommit(commit *Subject, src istanbul.Validator) error {
	logger := c.logger.New("from", src, "state", c.state)

	sub := c.current.Subject()
	if !reflect.DeepEqual(commit.View, sub.View) || commit.Digest.Hex() != sub.Digest.Hex() {
		logger.Warn("Inconsistent subjects between commit and proposal", "expected", sub, "got", commit)
		return errInconsistentSubject
	}

	return nil
}

func (c *core) acceptCommit(msg *message, src istanbul.Validator) error {
	logger := c.logger.New("from", src, "state", c.state)

	// Add the COMMIT message to current round state
	if err := c.current.Commits.Add(msg); err != nil {
		logger.Error("Failed to record commit message", "msg", msg, "err", err)
		return err
	}

	return nil
}

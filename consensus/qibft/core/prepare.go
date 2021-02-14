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
	"github.com/ethereum/go-ethereum/rlp"
)

func (c *core) broadcastPrepare() {
	var err error

	logger := c.logger.New("state", c.state)

	sub := c.current.Subject()
	prepareMsg := &PrepareMsg{
		CommonMsg:  CommonMsg{
			code:           prepareMsgCode,
			source: c.address,
			Sequence:       sub.View.Sequence,
			Round:          sub.View.Round,
		},
		Digest:     sub.Digest,
	}

	// Sign Message
	encodedPayload, err := prepareMsg.EncodePayload()
	if err != nil {
		logger.Error("QBFT: Failed to encode payload of prepare message", "msg", prepareMsg, "err", err)
		return
	}
	prepareMsg.signature, err = c.backend.Sign(encodedPayload)
	if err != nil {
		logger.Error("QBFT: Failed to sign commit message", "msg", prepareMsg, "err", err)
		return
	}

	// RLP-encode message
	payload, err := rlp.EncodeToBytes(&prepareMsg)
	if err != nil {
		logger.Error("QBFT: Failed to encode commit message", "msg", prepareMsg, "err", err)
		return
	}

	logger.Info("QBFT: broadcastPrepare", "m", sub, "payload", payload)
	// Broadcast RLP-encoded message
	if err = c.backend.Broadcast(c.valSet, prepareMsgCode, payload); err != nil {
		logger.Error("QBFT: Failed to broadcast message", "msg", prepareMsg, "err", err)
		return
	}
}

func (c *core) handlePrepare(prepare *PrepareMsg) error {
	logger := c.logger.New("state", c.state)

	logger.Info("QBFT: handlePrepare", "msg", &prepare)

	// Check digest
	if prepare.Digest != c.current.Proposal().Hash() {
		logger.Error("QBFT: Failed to check digest")
		return errInvalidMessage
	}

	// Add to received msgs
	if err := c.current.QBFTPrepares.Add(prepare); err != nil {
		c.logger.Error("QBFT: Failed to save prepare message", "msg", prepare, "err", err)
		return err
	}

	// Change to Prepared state if we've received enough PREPARE messages
	// and we are in earlier state before Prepared state.
	if (c.current.QBFTPrepares.Size() >= c.QuorumSize()) && c.state.Cmp(StatePrepared) < 0 {

		// IBFT REDUX
		c.current.preparedRound = c.currentView().Round
		c.QBFTPreparedPrepares = c.current.QBFTPrepares
		if c.current.Proposal() != nil && c.current.Proposal().Hash() == prepare.Digest {
			c.current.preparedBlock = c.current.Proposal()
		}

		c.setState(StatePrepared)
		c.broadcastCommit()
	}

	return nil
}

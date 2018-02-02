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
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

func (c *core) handleFinalCommitted(proposal istanbul.Proposal, proposer common.Address) error {
	logger := c.logger.New("state", c.state, "number", proposal.Number(), "hash", proposal.Hash())
	logger.Trace("Received a final committed proposal")

	// Catch up the sequence number
	if proposal.Number().Cmp(c.current.Sequence()) >= 0 {
		// Remember to store the proposer since we've accpetted the proposal
		diff := new(big.Int).Sub(proposal.Number(), c.current.Sequence())
		c.sequenceMeter.Mark(new(big.Int).Add(diff, common.Big1).Int64())

		if !c.consensusTimestamp.IsZero() {
			c.consensusTimer.UpdateSince(c.consensusTimestamp)
			c.consensusTimestamp = time.Time{}
		}

		c.lastProposer = proposer
		c.lastProposal = proposal
		c.startNewRound(&istanbul.View{
			Sequence: new(big.Int).Add(proposal.Number(), common.Big1),
			Round:    new(big.Int).Set(common.Big0),
		}, proposal, proposer, false)
	}

	return nil
}

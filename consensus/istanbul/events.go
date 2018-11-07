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
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

// RequestEvent is posted to propose a proposal
type RequestEvent struct {
	Proposal Proposal
}

// MessageEvent is posted for Istanbul engine communication
type MessageEvent struct {
	Payload []byte
}

// FinalCommittedEvent is posted when a proposal is committed
type FinalCommittedEvent struct {
}

// cachUp is the information to report of a time stamp of proposer time for a block.
type CatchUpEvent struct {
	Action string      `json: "action"`
	Data   DataCatchUp `json: "data"`
}
type DataCatchUp struct {
	Address       common.Address  `json: "address"`
	Block         *big.Int        `json: "block"`
	OldProposer   common.Address  `json: "old_proposer"`
	NewProposer   *common.Address `json: "new_proposer"`
	Validators    []string        `json: "validators"`
	ValidatorSize int             `json: "validator_size"`
}

func (cu *CatchUpEvent) Str() (str string) {
	str = "{action: '" + cu.Action + "', data: {address: '" + cu.Data.Address.Hex()
	str += "', block: " + fmt.Sprint(cu.Data.Block) + ", old_proposer: '"
	str += cu.Data.OldProposer.Hex() + "', "
	if cu.Data.NewProposer != nil {
		str += "new_proposer: '" + cu.Data.NewProposer.Hex() + "', "
	}
	str += "validators: ["
	for cont := 0; cont < cu.Data.ValidatorSize; cont++ {
		str += "'" + cu.Data.Validators[cont] + "'"
		if cont < (cu.Data.ValidatorSize - 1) {
			str += ", "
		}
	}
	str += "], validator_size: " + fmt.Sprint(cu.Data.ValidatorSize) + "}"
	return str
}

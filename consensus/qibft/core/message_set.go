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
	"fmt"
	"io"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/rlp"
)

// Construct a new message_deprecated set to accumulate messages for given sequence/view number.
func newMessageSet(valSet istanbul.ValidatorSet) *messageSet {
	return &messageSet{
		view: &View{
			Round:    new(big.Int),
			Sequence: new(big.Int),
		},
		messagesMu: new(sync.Mutex),
		messages:   make(map[common.Address]*message_deprecated),
		valSet:     valSet,
	}
}

// ----------------------------------------------------------------------------

type messageSet struct {
	view       *View
	valSet     istanbul.ValidatorSet
	messagesMu *sync.Mutex
	messages   map[common.Address]*message_deprecated
}

// messageMapAsStruct is a temporary holder struct to convert messages map to a slice when Encoding and Decoding messageSet
type messageMapAsStruct struct {
	Address common.Address
	Msg     *message_deprecated
}

func (ms *messageSet) View() *View {
	return ms.view
}

func (ms *messageSet) Add(msg *message_deprecated) error {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()

	if err := ms.verify(msg); err != nil {
		return err
	}

	return ms.addVerifiedMessage(msg)
}

func (ms *messageSet) Values() (result []*message_deprecated) {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()

	for _, v := range ms.messages {
		result = append(result, v)
	}

	return result
}

func (ms *messageSet) Size() int {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	return len(ms.messages)
}

func (ms *messageSet) Get(addr common.Address) *message_deprecated {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	return ms.messages[addr]
}

// ----------------------------------------------------------------------------

func (ms *messageSet) verify(msg *message_deprecated) error {
	// verify if the message_deprecated comes from one of the validators
	if _, v := ms.valSet.GetByAddress(msg.Address); v == nil {
		return istanbul.ErrUnauthorizedAddress
	}

	// TODO: check view number and sequence number

	return nil
}

func (ms *messageSet) addVerifiedMessage(msg *message_deprecated) error {
	ms.messages[msg.Address] = msg
	return nil
}

func (ms *messageSet) String() string {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	addresses := make([]string, 0, len(ms.messages))
	for _, v := range ms.messages {
		addresses = append(addresses, v.Address.String())
	}
	return fmt.Sprintf("[%v]", strings.Join(addresses, ", "))
}

// EncodeRLP serializes messageSet into Ethereum RLP format
// valSet is currently not being encoded.
func (ms *messageSet) EncodeRLP(w io.Writer) error {
	if ms == nil {
		return nil
	}
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()

	// maps cannot be RLP encoded, convert the map into a slice of struct and then encode it
	var messagesAsSlice []messageMapAsStruct
	for k, v := range ms.messages {
		msgMapAsStruct := messageMapAsStruct{
			Address: k,
			Msg:     v,
		}
		messagesAsSlice = append(messagesAsSlice, msgMapAsStruct)
	}

	return rlp.Encode(w, []interface{}{
		ms.view,
		//		ms.valSet,
		messagesAsSlice,
	})
}

// DecodeRLP deserializes rlp stream into messageSet
// valSet is currently not being decoded
func (ms *messageSet) DecodeRLP(stream *rlp.Stream) error {
	// Don't decode messageSet if the size of the stream is 0
	_, size, _ := stream.Kind()
	if size == 0 {
		return nil
	}
	var msgSet struct {
		MsgView *View
		//		valSet        istanbul.ValidatorSet
		MessagesSlice []messageMapAsStruct
	}
	if err := stream.Decode(&msgSet); err != nil {
		return err
	}

	// convert the messages struct slice back to map
	messages := make(map[common.Address]*message_deprecated)
	for _, msgStruct := range msgSet.MessagesSlice {
		messages[msgStruct.Address] = msgStruct.Msg
	}

	ms.view = msgSet.MsgView
	//	ms.valSet = msgSet.valSet
	ms.messages = messages
	ms.messagesMu = new(sync.Mutex)

	return nil
}

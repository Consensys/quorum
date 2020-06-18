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
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

func TestCheckMessage(t *testing.T) {
	c := &core{
		state: StateAcceptRequest,
		current: newRoundState(&View{
			Sequence: big.NewInt(1),
			Round:    big.NewInt(0),
		}, newTestValidatorSet(4), nil, nil, nil),
	}

	// invalid view format
	err := c.checkMessage(msgPreprepare, nil)
	if err != errInvalidMessage {
		t.Errorf("error mismatch: have %v, want %v", err, errInvalidMessage)
	}

	testStates := []State{StateAcceptRequest, StatePreprepared, StatePrepared, StateCommitted}
	testCode := []uint64{msgPreprepare, msgPrepare, msgCommit, msgRoundChange}

	// future sequence
	v := &View{
		Sequence: big.NewInt(2),
		Round:    big.NewInt(0),
	}
	for i := 0; i < len(testStates); i++ {
		c.state = testStates[i]
		for j := 0; j < len(testCode); j++ {
			err := c.checkMessage(testCode[j], v)
			if err != errFutureMessage {
				t.Errorf("error mismatch: have %v, want %v", err, errFutureMessage)
			}
		}
	}

	// future round
	v = &View{
		Sequence: big.NewInt(1),
		Round:    big.NewInt(1),
	}
	for i := 0; i < len(testStates); i++ {
		c.state = testStates[i]
		for j := 0; j < len(testCode); j++ {
			err := c.checkMessage(testCode[j], v)
			if testCode[j] == msgRoundChange {
				if err != nil {
					t.Errorf("error mismatch: have %v, want nil", err)
				}
			} else if err != errFutureMessage {
				t.Errorf("error mismatch: have %v, want %v", err, errFutureMessage)
			}
		}
	}

	// current view but waiting for round change
	v = &View{
		Sequence: big.NewInt(1),
		Round:    big.NewInt(0),
	}
	for i := 0; i < len(testStates); i++ {
		c.state = testStates[i]
		for j := 0; j < len(testCode); j++ {
			err := c.checkMessage(testCode[j], v)
			if testCode[j] == msgRoundChange {
				if err != nil {
					t.Errorf("error mismatch: have %v, want nil", err)
				}
			} else if testStates[i] == StateAcceptRequest && testCode[j] > msgPreprepare {
				if err != errFutureMessage {
					t.Errorf("error mismatch: have %v, want %v", err, errFutureMessage)
				}
			} else if err != nil {
				t.Errorf("error mismatch: have %v, want %v", err, nil)
			}
		}
	}

	v = c.currentView()
	// current view, state = StateAcceptRequest
	c.state = StateAcceptRequest
	for i := 0; i < len(testCode); i++ {
		err = c.checkMessage(testCode[i], v)
		if testCode[i] == msgRoundChange {
			if err != nil {
				t.Errorf("error mismatch: have %v, want nil", err)
			}
		} else if testCode[i] == msgPreprepare {
			if err != nil {
				t.Errorf("error mismatch: have %v, want nil", err)
			}
		} else {
			if err != errFutureMessage {
				t.Errorf("error mismatch: have %v, want %v", err, errFutureMessage)
			}
		}
	}

	// current view, state = StatePreprepared
	c.state = StatePreprepared
	for i := 0; i < len(testCode); i++ {
		err = c.checkMessage(testCode[i], v)
		if testCode[i] == msgRoundChange {
			if err != nil {
				t.Errorf("error mismatch: have %v, want nil", err)
			}
		} else if err != nil {
			t.Errorf("error mismatch: have %v, want nil", err)
		}
	}

	// current view, state = StatePrepared
	c.state = StatePrepared
	for i := 0; i < len(testCode); i++ {
		err = c.checkMessage(testCode[i], v)
		if testCode[i] == msgRoundChange {
			if err != nil {
				t.Errorf("error mismatch: have %v, want nil", err)
			}
		} else if err != nil {
			t.Errorf("error mismatch: have %v, want nil", err)
		}
	}

	// current view, state = StateCommitted
	c.state = StateCommitted
	for i := 0; i < len(testCode); i++ {
		err = c.checkMessage(testCode[i], v)
		if testCode[i] == msgRoundChange {
			if err != nil {
				t.Errorf("error mismatch: have %v, want nil", err)
			}
		} else if err != nil {
			t.Errorf("error mismatch: have %v, want nil", err)
		}
	}

}

func TestStoreBacklog(t *testing.T) {
	c := &core{
		logger:     log.New("backend", "test", "id", 0),
		valSet:     newTestValidatorSet(1),
		backlogs:   make(map[common.Address]*prque.Prque),
		backlogsMu: new(sync.Mutex),
	}
	v := &View{
		Round:    big.NewInt(10),
		Sequence: big.NewInt(10),
	}
	p := c.valSet.GetByIndex(0)
	// push preprepare msg
	proposal := makeBlock(1)
	preprepareWithPB := &PreprepareWithPiggybackMsgs{
		Preprepare: &Preprepare{
			View:     v,
			Proposal: proposal,
		},
		PiggybackMessages: &PiggybackMessages{
			PreparedMessages: newMessageSet(c.valSet),
			RCMessages:       newMessageSet(c.valSet),
		},
	}
	prepreparePayload, _ := Encode(preprepareWithPB)
	m := &message{
		Code: msgPreprepare,
		Msg:  prepreparePayload,
	}
	c.storeBacklog(m, p)
	msg := c.backlogs[p.Address()].PopItem()
	if !reflect.DeepEqual(msg, m) {
		t.Errorf("message mismatch: have %v, want %v", msg, m)
	}

	preparedBlock := newTestProposal()
	// push prepare msg
	subject := &Subject{
		View:   v,
		Digest: preparedBlock,
	}
	subjectPayload, _ := Encode(subject)

	// round change message
	rcMessage := &RoundChangePiggybackMsgs{
		RoundChangeMessage: &RoundChangeMessage{
			View:          v,
			PreparedRound: v.Round,
			PreparedBlock: preparedBlock,
		},
	}

	rcPayload, _ := Encode(rcMessage)

	m = &message{
		Code: msgPrepare,
		Msg:  subjectPayload,
	}
	c.storeBacklog(m, p)
	msg = c.backlogs[p.Address()].PopItem()
	if !reflect.DeepEqual(msg, m) {
		t.Errorf("message mismatch: have %v, want %v", msg, m)
	}

	// push commit msg
	m = &message{
		Code: msgCommit,
		Msg:  subjectPayload,
	}
	c.storeBacklog(m, p)
	msg = c.backlogs[p.Address()].PopItem()
	if !reflect.DeepEqual(msg, m) {
		t.Errorf("message mismatch: have %v, want %v", msg, m)
	}

	// push roundChange msg
	m = &message{
		Code: msgRoundChange,
		Msg:  rcPayload,
	}
	c.storeBacklog(m, p)
	msg = c.backlogs[p.Address()].PopItem()
	if !reflect.DeepEqual(msg, m) {
		t.Errorf("message mismatch: have %v, want %v", msg, m)
	}
}

func TestProcessFutureBacklog(t *testing.T) {
	backend := &testSystemBackend{
		events: new(event.TypeMux),
	}
	c := &core{
		logger:     log.New("backend", "test", "id", 0),
		valSet:     newTestValidatorSet(1),
		backlogs:   make(map[common.Address]*prque.Prque),
		backlogsMu: new(sync.Mutex),
		backend:    backend,
		current: newRoundState(&View{
			Sequence: big.NewInt(1),
			Round:    big.NewInt(0),
		}, newTestValidatorSet(4), nil, nil, nil),
		state: StateAcceptRequest,
	}
	c.subscribeEvents()
	defer c.unsubscribeEvents()

	v := &View{
		Round:    big.NewInt(10),
		Sequence: big.NewInt(10),
	}
	p := c.valSet.GetByIndex(0)
	// push a future msg
	subject := &Subject{
		View:   v,
		Digest: makeBlock(5),
	}
	subjectPayload, _ := Encode(subject)
	m := &message{
		Code: msgCommit,
		Msg:  subjectPayload,
	}
	c.storeBacklog(m, p)
	c.processBacklog()

	const timeoutDura = 2 * time.Second
	timeout := time.NewTimer(timeoutDura)
	select {
	case e, ok := <-c.events.Chan():
		if !ok {
			return
		}
		t.Errorf("unexpected events comes: %v", e)
	case <-timeout.C:
		// success
	}
}

func TestProcessBacklog(t *testing.T) {
	vset := newTestValidatorSet(1)
	v := &View{
		Round:    big.NewInt(0),
		Sequence: big.NewInt(1),
	}
	proposal := makeBlock(1)
	preprepareWithPB := &PreprepareWithPiggybackMsgs{
		Preprepare: &Preprepare{
			View:     v,
			Proposal: proposal,
		},
		PiggybackMessages: &PiggybackMessages{
			PreparedMessages: newMessageSet(vset),
			RCMessages:       newMessageSet(vset),
		},
	}
	prepreparePayload, _ := Encode(preprepareWithPB)

	subject := &Subject{
		View:   v,
		Digest: makeBlock(5),
	}
	subjectPayload, _ := Encode(subject)

	roundChangeWithPB := &RoundChangePiggybackMsgs{
		RoundChangeMessage: &RoundChangeMessage{
			View:          v,
			PreparedRound: v.Round,
			PreparedBlock: makeBlock(5),
		},
	}

	roundChangePayload, _ := Encode(roundChangeWithPB)

	msgs := []*message{
		{
			Code: msgPreprepare,
			Msg:  prepreparePayload,
		},
		{
			Code: msgPrepare,
			Msg:  subjectPayload,
		},
		{
			Code: msgCommit,
			Msg:  subjectPayload,
		},
		{
			Code: msgRoundChange,
			Msg:  roundChangePayload,
		},
	}
	for i := 0; i < len(msgs); i++ {
		testProcessBacklog(t, msgs[i], vset)
	}
}

func testProcessBacklog(t *testing.T, msg *message, vset istanbul.ValidatorSet) {
	backend := &testSystemBackend{
		events: new(event.TypeMux),
		peers:  vset,
	}
	c := &core{
		logger:     log.New("backend", "test", "id", 0),
		backlogs:   make(map[common.Address]*prque.Prque),
		backlogsMu: new(sync.Mutex),
		valSet:     vset,
		backend:    backend,
		state:      State(msg.Code),
		current: newRoundState(&View{
			Sequence: big.NewInt(1),
			Round:    big.NewInt(0),
		}, newTestValidatorSet(4), nil, nil, nil),
	}
	c.subscribeEvents()
	defer c.unsubscribeEvents()

	c.storeBacklog(msg, vset.GetByIndex(0))
	c.processBacklog()

	const timeoutDura = 2 * time.Second
	timeout := time.NewTimer(timeoutDura)
	select {
	case ev := <-c.events.Chan():
		e, ok := ev.Data.(backlogEvent)
		if !ok {
			t.Errorf("unexpected event comes: %v", reflect.TypeOf(ev.Data))
		}
		if e.msg.Code != msg.Code {
			t.Errorf("message code mismatch: have %v, want %v", e.msg.Code, msg.Code)
		}
		// success
	case <-timeout.C:
		t.Errorf("unexpected timeout occurs for msg: %v", msg.Code)
	}
}

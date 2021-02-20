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
	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

var (
	// msgPriority is defined for calculating processing priority to speedup consensus
	// msgPreprepare > msgCommit > msgPrepare
	msgPriority = map[uint64]int{
		preprepareMsgCode: 1,
		commitMsgCode:     2,
		prepareMsgCode:    3,
	}
)

// checkMessage checks the message_deprecated state
// return errInvalidMessage if the message_deprecated is invalid
// return errFutureMessage if the message_deprecated view is larger than current view
// return errOldMessage if the message_deprecated view is smaller than current view
func (c *core) checkMessage(msgCode uint64, view *View) error {
	if view == nil || view.Sequence == nil || view.Round == nil {
		return errInvalidMessage
	}

	if msgCode == roundChangeMsgCode {
		if view.Sequence.Cmp(c.currentView().Sequence) > 0 {
			return errFutureMessage
		} else if view.Cmp(c.currentView()) < 0 {
			return errOldMessage
		}
		return nil
	}

	if view.Cmp(c.currentView()) > 0 {
		return errFutureMessage
	}

	if view.Cmp(c.currentView()) < 0 {
		return errOldMessage
	}

	switch c.state {
	case StateAcceptRequest:
		// StateAcceptRequest only accepts msgPreprepare and msgRoundChange
		// other messages are future messages
		if msgCode > preprepareMsgCode {
			return errFutureMessage
		}
		return nil
	case StatePreprepared:
		// StatePreprepared only accepts msgPrepare and msgRoundChange
		// message_deprecated less than msgPrepare are invalid and greater are future messages
		if msgCode < prepareMsgCode {
			return errInvalidMessage
		} else if msgCode > prepareMsgCode {
			return errFutureMessage
		}
		return nil
	case StatePrepared:
		// StatePrepared only accepts msgCommit and msgRoundChange
		// other messages are invalid messages
		if msgCode < commitMsgCode {
			return errInvalidMessage
		}
		return nil
	case StateCommitted:
		// StateCommit rejects all messages other than msgRoundChange
		return errInvalidMessage
	}
	return nil
}

func (c *core) storeQBFTBacklog(msg QBFTMessage) {
	src := msg.Source()
	logger := c.logger.New("from", src, "state", c.state)

	if src == c.Address() {
		logger.Warn("Backlog from self")
		return
	}

	logger.Trace("Store future message_deprecated")

	c.backlogsMu.Lock()
	defer c.backlogsMu.Unlock()

	logger.Debug("Retrieving backlog queue", "for", src, "backlogs_size", len(c.backlogs))
	backlog := c.backlogs[src]
	if backlog == nil {
		backlog = prque.New()
	}
	view := msg.View()
	backlog.Push(msg, toPriority(msg.Code(), &view))
	c.backlogs[src] = backlog
}

func (c *core) processBacklog() {
	c.backlogsMu.Lock()
	defer c.backlogsMu.Unlock()

	for srcAddress, backlog := range c.backlogs {
		if backlog == nil {
			continue
		}
		_, src := c.valSet.GetByAddress(srcAddress)
		if src == nil {
			// validator is not available
			delete(c.backlogs, srcAddress)
			continue
		}
		logger := c.logger.New("from", src, "state", c.state)
		isFuture := false

		// We stop processing if
		//   1. backlog is empty
		//   2. The first message_deprecated in queue is a future message_deprecated
		for !(backlog.Empty() || isFuture) {
			m, prio := backlog.Pop()

			var code uint64
			var view View
			var event backlogEvent


			msg := m.(QBFTMessage)
			code = msg.Code()
			view = msg.View()
			event.msg = msg

			// Push back if it's a future message_deprecated
			err := c.checkMessage(code, &view)
			if err != nil {
				if err == errFutureMessage {
					logger.Trace("Stop processing backlog", "msg", m)
					backlog.Push(m, prio)
					isFuture = true
					break
				}
				logger.Trace("Skip the backlog event", "msg", m, "err", err)
				continue
			}
			logger.Trace("Post backlog event", "msg", m)

			event.src = src
			go c.sendEvent(event)
		}
	}
}

func toPriority(msgCode uint64, view *View) float32 {
	if msgCode == roundChangeMsgCode {
		// For msgRoundChange, set the message_deprecated priority based on its sequence
		return -float32(view.Sequence.Uint64() * 1000)
	}
	// FIXME: round will be reset as 0 while new sequence
	// 10 * Round limits the range of message_deprecated code is from 0 to 9
	// 1000 * Sequence limits the range of round is from 0 to 99
	return -float32(view.Sequence.Uint64()*1000 + view.Round.Uint64()*10 + uint64(msgPriority[msgCode]))
}

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

package eth

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
)

// ethPeerInfo represents a short summary of the `eth` sub-protocol metadata known
// about a connected peer.
type ethPeerInfo struct {
	Version    uint     `json:"version"`    // Ethereum protocol version negotiated
	Difficulty *big.Int `json:"difficulty"` // Total difficulty of the peer's blockchain
	Head       string   `json:"head"`       // Hex hash of the peer's best owned block
}

// ethPeer is a wrapper around eth.Peer to maintain a few extra metadata.
type ethPeer struct {
	*eth.Peer
	snapExt *snapPeer // Satellite `snap` connection

	syncDrop *time.Timer   // Connection dropper if `eth` sync progress isn't validated in time
	snapWait chan struct{} // Notification channel for snap connections
	lock     sync.RWMutex  // Mutex protecting the internal fields
	// TODO qlight - consider whether it is worth duplicating the peer structure and the surrounding zoo
	qlightServer bool
	qlightPSI    string
	qlightToken  string
}

// info gathers and returns some `eth` protocol metadata known about a peer.
func (p *ethPeer) info() *ethPeerInfo {
	hash, td := p.Head()

	return &ethPeerInfo{
		Version:    p.Version(),
		Difficulty: td,
		Head:       hash.Hex(),
	}
}

// snapPeerInfo represents a short summary of the `snap` sub-protocol metadata known
// about a connected peer.
type snapPeerInfo struct {
	Version uint `json:"version"` // Snapshot protocol version negotiated
}

// snapPeer is a wrapper around snap.Peer to maintain a few extra metadata.
type snapPeer struct {
	*snap.Peer
}

// info gathers and returns some `snap` protocol metadata known about a peer.
func (p *snapPeer) info() *snapPeerInfo {
	return &snapPeerInfo{
		Version: p.Version(),
	}
}


// TODO qlight rebase
//case prop := <-p.queuedBlocks:
//if len(p.qlightPSI) > 0 && prop.privateTransactionsData != nil {
//p.Log().Info("Sending new block private data msg")
//err := p2p.Send(p.rw, QLightNewBlockPrivateDataMsg, prop.privateTransactionsData)
//if err != nil {
//p.Log().Error("Error occurred while sending private data msg", "err", err)
//removePeer(p.id)
//return
//}
//}

//case prop := <-p.queuedBlocks:
//if pbd := prop.privateBlockData; len(p.qlightPSI) > 0 && pbd != nil {
//var toBroadcast []engine.BlockPrivatePayloads
//
//if p.qlightPSI != pbd.PSI.String() {
//// TODO(cjh) review log levels for all new logs - if the decision is to use the qlight prefix make sure this is consistently present
//p.Log().Info("qlight: PSI not used by the qlight client, skipping")
//continue
//}
//
//result := &engine.BlockPrivatePayloads{
//BlockHash:        pbd.BlockHash.ToBase64(),
//PrivateStateRoot: pbd.PrivateStateRoot.ToBase64(),
//Payloads:         make([]engine.RLPPrivateTx, len(pbd.PrivateTransactions)),
//}
//for i, privTxData := range pbd.PrivateTransactions {
//result.Payloads[i].EncryptedPayloadHashB64 = privTxData.Hash.ToBase64()
//result.Payloads[i].QuorumPrivateTxData = engine.QuorumPayloadExtra{
//Payload:       fmt.Sprintf("0x%x", privTxData.Payload),
//ExtraMetaData: privTxData.Extra,
//IsSender:      privTxData.IsSender,
//}
//}
//toBroadcast = append(toBroadcast, *result)
//
//p.Log().Info("Sending new block private data msg", "len(toBroadcast)", len(toBroadcast))
//err := p2p.Send(p.rw, QLightNewBlockPrivateDataMsg, toBroadcast)
//if err != nil {
//p.Log().Error("Error occurred while sending private data msg", "err", err)
//removePeer(p.id)
//return
//}
//}



type PrivateTransactionsData []PrivateTransactionData

type PrivateTransactionData struct {
	Hash     *common.EncryptedPayloadHash
	Payload  []byte
	Extra    *engine.ExtraMetadata
	IsSender bool
}


func (p *peer) QLightHandshake(server bool, psi string, token string) error {
	// Send out own handshake in a new thread
	errc := make(chan error, 2)

	var (
		status qLightStatusData // safe to read after two values have been received from errc
	)
	go func() {
		errc <- p2p.Send(p.rw, QLightStatusMsg, &qLightStatusData{
			ProtocolVersion: uint32(p.version),
			Server:          server,
			PSI:             psi,
			Token:           token,
		})
	}()
	go func() {
		errc <- p.readQLightStatus(&status)
	}()
	timeout := time.NewTimer(handshakeTimeout)
	defer timeout.Stop()
	for i := 0; i < 2; i++ {
		select {
		case err := <-errc:
			if err != nil {
				return err
			}
		case <-timeout.C:
			return p2p.DiscReadTimeout
		}
	}
	p.qlightServer, p.qlightPSI, p.qlightToken = status.Server, status.PSI, status.Token
	return nil
}


func (p *peer) readQLightStatus(qligtStatus *qLightStatusData) error {
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Code != QLightStatusMsg {
		return errResp(ErrNoStatusMsg, "second msg has code %x (!= %x)", msg.Code, QLightStatusMsg)
	}
	if msg.Size > protocolMaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocolMaxMsgSize)
	}
	// Decode the handshake and make sure everything matches
	if err := msg.Decode(&qligtStatus); err != nil {
		return errResp(ErrDecode, "msg %v: %v", msg, err)
	}
	if !qligtStatus.Server && len(qligtStatus.PSI) == 0 {
		return errResp(ErrDecode, "client connected without specifying PSI")
	}
	// TODO qlight - check that the PSI exists
	// TODO qlight - check that if multi tenancy is enabled the token matches the PSI
	return nil
}

func (ps *peerSet) RegisterIdlePeer(p *peer, removePeer func(string), protoName string) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if ps.closed {
		return errClosed
	}
	if _, ok := ps.peers[p.id]; ok {
		return errAlreadyRegistered
	}
	ps.peers[p.id] = p
	return nil
}
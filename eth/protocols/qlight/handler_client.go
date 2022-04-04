package qlight

import (
	"fmt"

	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

// MakeProtocols constructs the P2P protocol definitions for `eth`.
func MakeProtocolsClient(backend Backend, network uint64, dnsdisc enode.Iterator) []p2p.Protocol {
	protocols := make([]p2p.Protocol, 1)
	version := uint(QLIGHT65)
	protocols[0] = p2p.Protocol{
		Name:    ProtocolName,
		Version: QLIGHT65,
		Length:  QLightProtocolLength,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			ethPeer := eth.NewPeerNoBroadcast(version, p, rw, backend.TxPool())
			peer := NewPeer(version, p, rw, ethPeer)
			defer ethPeer.Close()
			defer peer.Close()

			return backend.RunQPeer(peer, func(peer *Peer) error {
				return HandleClient(backend, peer)
			})
		},
		NodeInfo: func() interface{} {
			return eth.NodeInfoFunc(backend.Chain(), network)
		},
		PeerInfo: func(id enode.ID) interface{} {
			return backend.PeerInfo(id)
		},
		Attributes:     []enr.Entry{eth.CurrentENREntry(backend.Chain())},
		DialCandidates: dnsdisc,
	}
	return protocols
}

// Handle is invoked whenever an `eth` connection is made that successfully passes
// the protocol handshake. This method will keep processing messages until the
// connection is torn down.
func HandleClient(backend Backend, peer *Peer) error {
	for {
		if err := handleMessageClient(backend, peer); err != nil {
			peer.Log().Debug("Message handling failed in `eth`", "err", err)
			return err
		}
	}
}

// handleMessage is invoked whenever an inbound message is received from a remote
// peer. The remote connection is torn down upon returning any error.
func handleMessageClient(backend Backend, peer *Peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := peer.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > maxMessageSize {
		return fmt.Errorf("%w: %v > %v", errMsgTooLarge, msg.Size, maxMessageSize)
	}
	defer msg.Discard()

	peer.Log().Info("QLight client message received", "msg", msg.Code)

	var handlers = eth.ETH_65_FULL_SYNC

	switch msg.Code {
	case eth.BlockHeadersMsg:
		if handler := handlers[msg.Code]; handler != nil {
			return handler(backend, msg, peer.EthPeer)
		}
	case eth.TransactionsMsg:
		return qlightClientHandleTransactions(backend, msg, peer)
	case eth.BlockBodiesMsg:
		res := new(eth.BlockBodiesPacket)
		if err := msg.Decode(res); err != nil {
			return fmt.Errorf("%w: message %v: %v", errDecode, msg, err)
		}
		return backend.QHandle(peer, res)
	case eth.NewBlockHashesMsg:
		if handler := handlers[msg.Code]; handler != nil {
			return handler(backend, msg, peer.EthPeer)
		}
	case QLightNewBlockPrivateDataMsg:
		peer.Log().Info("QLight Received block private data message", "msg", msg.Code)
		return qlightClientHandleNewBlockPrivateData(backend, msg, peer)
	case eth.NewBlockMsg:
		return qlightClientHandleNewBlock(backend, msg, peer)
	case eth.NewPooledTransactionHashesMsg:
		if handler := handlers[msg.Code]; handler != nil {
			return handler(backend, msg, peer.EthPeer)
		}
	}
	peer.Log().Info("QLight Unable to find handler for received message", "msg", msg.Code)
	return fmt.Errorf("%w: %v", errInvalidMsgCode, msg.Code)
}

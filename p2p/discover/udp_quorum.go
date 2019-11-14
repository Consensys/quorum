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

package discover

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/ethereum/go-ethereum/rlp"
)

// RPC request structures
type (
	pingQuorum struct {
		Version    uint
		From, To   rpcEndpoint
		Expiration uint64
		Hostname   string
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// pong is the reply to ping.
	pongQuorum struct {
		// This field should mirror the UDP envelope address
		// of the ping packet, which provides a way to discover the
		// the external address (after NAT).
		To rpcEndpoint

		ReplyTok   []byte // This contains the hash of the ping packet.
		Expiration uint64 // Absolute timestamp at which the packet becomes invalid.
		// Ignore additional fields (for forward compatibility).

		Rest []rlp.RawValue `rlp:"tail"`
	}

	// findnode is a query for nodes close to the given target.
	findnodeQuorum struct {
		Target     encPubkey
		Expiration uint64
		// Ignore additional fields (for forward compatibility).

		IsQuorum   bool

		Rest []rlp.RawValue `rlp:"tail"`
	}

	// reply to findnode
	neighborsQuorum struct {
		Nodes      []rpcNodeQuorum
		Expiration uint64

		IsQuorum   bool

		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// this message is the same as the vanilla discv4 neighbors
	// to allow older nodes to send replies to our findnode messages
	// reply to findnode
	neighborsVanilla struct {
		Nodes      []rpcNode
		Expiration uint64

		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	rpcNodeQuorum struct {
		IP  net.IP // len 4 for IPv4 or 16 for IPv6
		UDP uint16 // for discovery protocol
		TCP uint16 // for RLPx protocol
		ID  encPubkey
		Hostname string
	}

)

type packetQuorum interface {
	handle(t *udpQuorum, from *net.UDPAddr, fromKey encPubkey, mac []byte) error
	name() string
}

func (t *udpQuorum) nodeFromRPC(sender *net.UDPAddr, rn rpcNodeQuorum) (*node, error) {
	if rn.UDP <= 1024 {
		return nil, errors.New("low port")
	}
	if err := netutil.CheckRelayIP(sender.IP, rn.IP); err != nil {
		return nil, err
	}
	if t.netrestrict != nil && !t.netrestrict.Contains(rn.IP) {
		return nil, errors.New("not contained in netrestrict whitelist")
	}
	key, err := decodePubkey(rn.ID)
	if err != nil {
		return nil, err
	}

	var n *node
	if rn.Hostname == "" {
		n = wrapNode(enode.NewV4(key, rn.IP, int(rn.TCP), int(rn.UDP), 0))
	} else {
		n = wrapNode(enode.NewV4Hostname(key, rn.Hostname, int(rn.TCP), int(rn.UDP), 0))
	}

	err = n.ValidateComplete()
	return n, err
}

func nodeToRPCQuorum(n *node) rpcNodeQuorum {
	var key ecdsa.PublicKey
	var ekey encPubkey
	if err := n.Load((*enode.Secp256k1)(&key)); err == nil {
		ekey = encodePubkey(&key)
	}
	return rpcNodeQuorum{ID: ekey, IP: n.IP(), UDP: uint16(n.UDP()), TCP: uint16(n.TCP()), Hostname: n.Host()}
}

// udp implements the discovery v4 UDP wire protocol.
type udpQuorum struct {
	udp
}

// ListenUDPQuorum returns a new table that listens for UDP packets on laddr.
func ListenUDPQuorum(c conn, ln *enode.LocalNode, cfg Config) (*Table, error) {
	tab, _, err := newUDPQuorum(c, ln, cfg)
	if err != nil {
		return nil, err
	}
	return tab, nil
}

func newUDPQuorum(c conn, ln *enode.LocalNode, cfg Config) (*Table, *udpQuorum, error) {
	udp := &udp{
		conn:        c,
		priv:        cfg.PrivateKey,
		netrestrict: cfg.NetRestrict,
		localNode:   ln,
		db:          ln.Database(),
		closing:     make(chan struct{}),
		gotreply:    make(chan reply),
		addpending:  make(chan *pending),
	}
	udpQ := &udpQuorum{*udp}
	tab, err := newTable(udpQ, ln.Database(), cfg.Bootnodes)
	if err != nil {
		return nil, nil, err
	}
	udpQ.tab = tab

	udpQ.wg.Add(2)
	go udpQ.loop()
	go udpQ.readLoop(cfg.Unhandled)
	return udpQ.tab, udpQ, nil
}

// ping sends a ping message to the given node and waits for a reply.
func (t *udpQuorum) ping(toid enode.ID, toaddr *net.UDPAddr) error {
	return <-t.sendPing(toid, toaddr, nil)
}

// sendPing sends a ping message to the given node and invokes the callback
// when the reply arrives.
func (t *udpQuorum) sendPing(toid enode.ID, toaddr *net.UDPAddr, callback func()) <-chan error {
	req := &pingQuorum{
		Version:    4,
		From:       t.ourEndpoint(),
		To:         makeEndpoint(toaddr, 0), // TODO: maybe use known TCP port from DB
		Expiration: uint64(time.Now().Add(expiration).Unix()),
		Hostname:   t.localNode.Node().Host(),
	}
	packet, hash, err := encodePacket(t.priv, pingPacket, req)
	if err != nil {
		errc := make(chan error, 1)
		errc <- err
		return errc
	}
	errc := t.pending(toid, pongPacket, func(p interface{}) bool {
		ok := bytes.Equal(p.(*pongQuorum).ReplyTok, hash)
		if ok && callback != nil {
			callback()
		}
		return ok
	})
	t.localNode.UDPContact(toaddr)
	t.write(toaddr, req.name(), packet)
	return errc
}

// findnode sends a findnode request to the given node and waits until
// the node has sent up to k neighbors.
func (t *udpQuorum) findnode(toid enode.ID, toaddr *net.UDPAddr, target encPubkey) ([]*node, error) {
	// If we haven't seen a ping from the destination node for a while, it won't remember
	// our endpoint proof and reject findnode. Solicit a ping first.
	if time.Since(t.db.LastPingReceived(toid)) > bondExpiration {
		t.ping(toid, toaddr)
		t.waitping(toid)
	}

	nodes := make([]*node, 0, bucketSize)
	errc := t.pending(toid, neighborsPacket, func(r interface{}) bool {
		reply := r.(*neighborsQuorum)
		for _, rn := range reply.Nodes {
			n, err := t.nodeFromRPC(toaddr, rn)
			if err != nil {
				log.Trace("Invalid neighbor node received", "ip", rn.IP, "addr", toaddr, "err", err)
				continue
			}
			nodes = append(nodes, n)
		}
		return len(reply.Nodes) >= bucketSize
	})
	t.send(toaddr, findnodePacket, &findnodeQuorum{
		Target:     target,
		Expiration: uint64(time.Now().Add(expiration).Unix()),
		IsQuorum:   true,
	})
	return nodes, <-errc
}

func (t *udpQuorum) handleReply(from enode.ID, ptype byte, req packetQuorum) bool {
	matched := make(chan bool, 1)
	select {
	case t.gotreply <- reply{from, ptype, req, matched}:
		// loop will handle it
		return <-matched
	case <-t.closing:
		return false
	}
}

var (
	// Neighbors replies are sent across multiple packets to
	// stay below the 1280 byte limit. We compute the maximum number
	// of entries by stuffing a packet until it grows too large.
	maxNeighborsQuorum = 3
)

func (t *udpQuorum) send(toaddr *net.UDPAddr, ptype byte, req packetQuorum) ([]byte, error) {
	packet, hash, err := encodePacket(t.priv, ptype, req)
	if err != nil {
		return hash, err
	}
	return hash, t.write(toaddr, req.name(), packet)
}

// readLoop runs in its own goroutine. it handles incoming UDP packets.
func (t *udpQuorum) readLoop(unhandled chan<- ReadPacket) {
	defer t.wg.Done()
	if unhandled != nil {
		defer close(unhandled)
	}

	// Discovery packets are defined to be no larger than 1280 bytes.
	// Packets larger than this size will be cut at the end and treated
	// as invalid because their hash won't match.
	buf := make([]byte, 1280)
	for {
		nbytes, from, err := t.conn.ReadFromUDP(buf)
		if netutil.IsTemporaryError(err) {
			// Ignore temporary read errors.
			log.Debug("Temporary UDP read error", "err", err)
			continue
		} else if err != nil {
			// Shut down the loop for permament errors.
			log.Debug("UDP read error", "err", err)
			return
		}

		errHandle := t.handlePacket(from, buf[:nbytes])
		if errHandle != nil && unhandled != nil {
			select {
			case unhandled <- ReadPacket{buf[:nbytes], from}:
			default:
			}
		}
	}
}

func (t *udpQuorum) handlePacket(from *net.UDPAddr, buf []byte) error {
	packet, fromID, hash, err := decodePacketQuorum(buf)
	if err != nil {
		log.Debug("Bad discv4 packet", "addr", from, "err", err)
		return err
	}
	err = packet.handle(t, from, fromID, hash)
	log.Trace("<< "+packet.name(), "addr", from, "err", err)
	return err
}

func decodePacketQuorum(buf []byte) (packetQuorum, encPubkey, []byte, error) {
	if len(buf) < headSize+1 {
		return nil, encPubkey{}, nil, errPacketTooSmall
	}
	hash, sig, sigdata := buf[:macSize], buf[macSize:headSize], buf[headSize:]
	shouldhash := crypto.Keccak256(buf[macSize:])
	if !bytes.Equal(hash, shouldhash) {
		return nil, encPubkey{}, nil, errBadHash
	}
	fromKey, err := recoverNodeKey(crypto.Keccak256(buf[headSize:]), sig)
	if err != nil {
		return nil, fromKey, hash, err
	}

	var req packetQuorum
	switch ptype := sigdata[0]; ptype {
	case pingPacket:
		req = new(pingQuorum)
	case pongPacket:
		req = new(pongQuorum)
	case findnodePacket:
		req = new(findnodeQuorum)
	case neighborsPacket:
		req = new(neighborsQuorum)
	default:
		return nil, fromKey, hash, fmt.Errorf("unknown type: %d", ptype)
	}
	s := rlp.NewStream(bytes.NewReader(sigdata[1:]), 0)
	err = s.Decode(req)

	if (err != nil) && (sigdata[0] == neighborsPacket) {
		req = new(neighborsVanilla)
		s := rlp.NewStream(bytes.NewReader(sigdata[1:]), 0)
		err = s.Decode(req)
	}

	return req, fromKey, hash, err
}

func (req *pingQuorum) handle(t *udpQuorum, from *net.UDPAddr, fromKey encPubkey, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}
	key, err := decodePubkey(fromKey)
	if err != nil {
		return fmt.Errorf("invalid public key: %v", err)
	}
	t.send(from, pongPacket, &pongQuorum{
		To:         makeEndpoint(from, req.From.TCP),
		ReplyTok:   mac,
		Expiration: uint64(time.Now().Add(expiration).Unix()),
	})

	var n *node
	if req.Hostname == "" {
		n = wrapNode(enode.NewV4(key, from.IP, int(req.From.TCP), from.Port, 0))
	} else {
		n = wrapNode(enode.NewV4Hostname(key, req.Hostname, int(req.From.TCP), from.Port, 0))
	}
	t.handleReply(n.ID(), pingPacket, req)
	if time.Since(t.db.LastPongReceived(n.ID())) > bondExpiration {
		t.sendPing(n.ID(), from, func() { t.tab.addThroughPing(n) })
	} else {
		t.tab.addThroughPing(n)
	}
	t.localNode.UDPEndpointStatement(from, &net.UDPAddr{IP: req.To.IP, Port: int(req.To.UDP)})
	t.db.UpdateLastPingReceived(n.ID(), time.Now())
	return nil
}

func (req *pingQuorum) name() string { return "PING/v4Quorum" }

func (req *pongQuorum) handle(t *udpQuorum, from *net.UDPAddr, fromKey encPubkey, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}
	fromID := fromKey.id()
	if !t.handleReply(fromID, pongPacket, req) {
		return errUnsolicitedReply
	}
	t.localNode.UDPEndpointStatement(from, &net.UDPAddr{IP: req.To.IP, Port: int(req.To.UDP)})
	t.db.UpdateLastPongReceived(fromID, time.Now())
	return nil
}

func (req *pongQuorum) name() string { return "PONG/v4Quorum" }

func (req *findnodeQuorum) handle(t *udpQuorum, from *net.UDPAddr, fromKey encPubkey, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}
	fromID := fromKey.id()
	if time.Since(t.db.LastPongReceived(fromID)) > bondExpiration {
		// No endpoint proof pong exists, we don't process the packet. This prevents an
		// attack vector where the discovery protocol could be used to amplify traffic in a
		// DDOS attack. A malicious actor would send a findnode request with the IP address
		// and UDP port of the target as the source address. The recipient of the findnode
		// packet would then send a neighbors packet (which is a much bigger packet than
		// findnode) to the victim.
		return errUnknownNode
	}
	target := enode.ID(crypto.Keccak256Hash(req.Target[:]))
	t.tab.mutex.Lock()
	closest := t.tab.closest(target, bucketSize).entries
	t.tab.mutex.Unlock()

	p := neighborsQuorum{Expiration: uint64(time.Now().Add(expiration).Unix()), IsQuorum: true}
	var sent bool
	// Send neighbors in chunks with at most maxNeighbors per packet
	// to stay below the 1280 byte limit.
	for _, n := range closest {
		if netutil.CheckRelayIP(from.IP, n.IP()) == nil {
			p.Nodes = append(p.Nodes, nodeToRPCQuorum(n))
		}
		if len(p.Nodes) == maxNeighborsQuorum {
			t.send(from, neighborsPacket, &p)
			p.Nodes = p.Nodes[:0]
			sent = true
		}
	}
	if len(p.Nodes) > 0 || !sent {
		t.send(from, neighborsPacket, &p)
	}
	return nil
}

func (req *findnodeQuorum) name() string { return "FINDNODE/v4Quorum" }

func (req *neighborsQuorum) handle(t *udpQuorum, from *net.UDPAddr, fromKey encPubkey, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}
	if !t.handleReply(fromKey.id(), neighborsPacket, req) {
		return errUnsolicitedReply
	}
	return nil
}

func (req *neighborsQuorum) name() string { return "NEIGHBORS/v4Quorum" }

func (req *neighborsVanilla) handle(t *udpQuorum, from *net.UDPAddr, fromKey encPubkey, mac []byte) error {
	convertedReq := new(neighborsQuorum)
	convertedReq.IsQuorum = true
	convertedReq.Expiration = req.Expiration
	convertedReq.Rest = req.Rest
	convertedReq.Nodes = make([]rpcNodeQuorum, 0)
	for _, node := range req.Nodes {
		convertedNode := new(rpcNodeQuorum)
		convertedNode.IP = node.IP
		convertedNode.Hostname = node.IP.String()
		convertedNode.TCP = node.TCP
		convertedNode.ID = node.ID
		convertedNode.UDP = node.UDP
		convertedReq.Nodes = append(convertedReq.Nodes, *convertedNode)
	}

	if expired(req.Expiration) {
		return errExpired
	}
	if !t.handleReply(fromKey.id(), neighborsPacket, convertedReq) {
		return errUnsolicitedReply
	}
	return nil
}

func (req *neighborsVanilla) name() string { return "NEIGHBORSVANILLA/v4Quorum" }

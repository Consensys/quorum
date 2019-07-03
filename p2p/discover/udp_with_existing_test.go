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
	"encoding/binary"
	"math/rand"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

func init() {
	spew.Config.DisableMethods = true
}

func newUDPTestUsingExistingTable(t *testing.T) *udpTest {
	test := &udpTest{
		t:          t,
		pipe:       newpipe(),
		localkey:   newkey(),
		remotekey:  newkey(),
		remoteaddr: &net.UDPAddr{IP: net.IP{10, 0, 1, 99}, Port: 30303},
	}
	db, _ := enode.OpenDB("")
	ln := enode.NewLocalNode(db, test.localkey)
	table, _, _ := newUDP(test.pipe, ln, Config{PrivateKey: test.localkey})
	test.table, test.udp, _ = newUDPExisting(test.pipe, ln, Config{PrivateKey: test.localkey}, table)

	// Wait for initial refresh so the table doesn't send unexpected findnode.
	<-test.table.initDone
	return test
}

func TestUDPExisting_packetErrors(t *testing.T) {
	test := newUDPTestUsingExistingTable(t)
	defer test.table.Close()

	test.packetIn(errExpired, pingPacket, &ping{From: testRemote, To: testLocalAnnounced, Version: 4})
	test.packetIn(errUnsolicitedReply, pongPacket, &pong{ReplyTok: []byte{}, Expiration: futureExp})
	test.packetIn(errUnknownNode, findnodePacket, &findnode{Expiration: futureExp})
	test.packetIn(errUnsolicitedReply, neighborsPacket, &neighbors{Expiration: futureExp})
}

func TestUDPExisting_pingTimeout(t *testing.T) {
	t.Parallel()
	test := newUDPTestUsingExistingTable(t)
	defer test.table.Close()

	toaddr := &net.UDPAddr{IP: net.ParseIP("1.2.3.4"), Port: 2222}
	toid := enode.ID{1, 2, 3, 4}
	if err := test.udp.ping(toid, toaddr); err != errTimeout {
		t.Error("expected timeout error, got", err)
	}
}

func TestUDPExisting_responseTimeouts(t *testing.T) {
	t.Parallel()
	test := newUDPTestUsingExistingTable(t)
	defer test.table.Close()

	rand.Seed(time.Now().UnixNano())
	randomDuration := func(max time.Duration) time.Duration {
		return time.Duration(rand.Int63n(int64(max)))
	}

	var (
		nReqs      = 200
		nTimeouts  = 0                       // number of requests with ptype > 128
		nilErr     = make(chan error, nReqs) // for requests that get a reply
		timeoutErr = make(chan error, nReqs) // for requests that time out
	)
	for i := 0; i < nReqs; i++ {
		// Create a matcher for a random request in udp.loop. Requests
		// with ptype <= 128 will not get a reply and should time out.
		// For all other requests, a reply is scheduled to arrive
		// within the timeout window.
		p := &pending{
			ptype:    byte(rand.Intn(255)),
			callback: func(interface{}) bool { return true },
		}
		binary.BigEndian.PutUint64(p.from[:], uint64(i))
		if p.ptype <= 128 {
			p.errc = timeoutErr
			test.udp.addpending <- p
			nTimeouts++
		} else {
			p.errc = nilErr
			test.udp.addpending <- p
			time.AfterFunc(randomDuration(60*time.Millisecond), func() {
				if !test.udp.handleReply(p.from, p.ptype, nil) {
					t.Logf("not matched: %v", p)
				}
			})
		}
		time.Sleep(randomDuration(30 * time.Millisecond))
	}

	// Check that all timeouts were delivered and that the rest got nil errors.
	// The replies must be delivered.
	var (
		recvDeadline        = time.After(20 * time.Second)
		nTimeoutsRecv, nNil = 0, 0
	)
	for i := 0; i < nReqs; i++ {
		select {
		case err := <-timeoutErr:
			if err != errTimeout {
				t.Fatalf("got non-timeout error on timeoutErr %d: %v", i, err)
			}
			nTimeoutsRecv++
		case err := <-nilErr:
			if err != nil {
				t.Fatalf("got non-nil error on nilErr %d: %v", i, err)
			}
			nNil++
		case <-recvDeadline:
			t.Fatalf("exceeded recv deadline")
		}
	}
	if nTimeoutsRecv != nTimeouts {
		t.Errorf("wrong number of timeout errors received: got %d, want %d", nTimeoutsRecv, nTimeouts)
	}
	if nNil != nReqs-nTimeouts {
		t.Errorf("wrong number of successful replies: got %d, want %d", nNil, nReqs-nTimeouts)
	}
}

func TestUDPExisting_findnodeTimeout(t *testing.T) {
	t.Parallel()
	test := newUDPTestUsingExistingTable(t)
	defer test.table.Close()

	toaddr := &net.UDPAddr{IP: net.ParseIP("1.2.3.4"), Port: 2222}
	toid := enode.ID{1, 2, 3, 4}
	target := encPubkey{4, 5, 6, 7}
	result, err := test.udp.findnode(toid, toaddr, target)
	if err != errTimeout {
		t.Error("expected timeout error, got", err)
	}
	if len(result) > 0 {
		t.Error("expected empty result, got", result)
	}
}

func TestUDPExisting_findnode(t *testing.T) {
	test := newUDPTestUsingExistingTable(t)
	defer test.table.Close()

	// put a few nodes into the table. their exact
	// distribution shouldn't matter much, although we need to
	// take care not to overflow any bucket.
	nodes := &nodesByDistance{target: testTarget.id()}
	for i := 0; i < bucketSize; i++ {
		key := newkey()
		n := wrapNode(enode.NewV4(&key.PublicKey, net.IP{10, 13, 0, 1}, 0, i, 0))
		nodes.push(n, bucketSize)
	}
	test.table.stuff(nodes.entries)

	// ensure there's a bond with the test node,
	// findnode won't be accepted otherwise.
	remoteID := encodePubkey(&test.remotekey.PublicKey).id()
	test.table.db.UpdateLastPongReceived(remoteID, time.Now())

	// check that closest neighbors are returned.
	test.packetIn(nil, findnodePacket, &findnode{Target: testTarget, Expiration: futureExp})
	expected := test.table.closest(testTarget.id(), bucketSize)

	waitNeighbors := func(want []*node) {
		test.waitPacketOut(func(p *neighbors) {
			if len(p.Nodes) != len(want) {
				t.Errorf("wrong number of results: got %d, want %d", len(p.Nodes), bucketSize)
			}
			for i := range p.Nodes {
				if p.Nodes[i].ID.id() != want[i].ID() {
					t.Errorf("result mismatch at %d:\n  got:  %v\n  want: %v", i, p.Nodes[i], expected.entries[i])
				}
			}
		})
	}
	foundNodes := 0
	for foundNodes < bucketSize-maxNeighbors {
		waitNeighbors(expected.entries[foundNodes:foundNodes+maxNeighbors])
		foundNodes += maxNeighbors
	}
	waitNeighbors(expected.entries[foundNodes:])
}

func TestUDPExisting_findnodeMultiReply(t *testing.T) {
	test := newUDPTestUsingExistingTable(t)
	defer test.table.Close()

	rid := enode.PubkeyToIDV4(&test.remotekey.PublicKey)
	test.table.db.UpdateLastPingReceived(rid, time.Now())

	// queue a pending findnode request
	resultc, errc := make(chan []*node), make(chan error)
	go func() {
		rid := encodePubkey(&test.remotekey.PublicKey).id()
		ns, err := test.udp.findnode(rid, test.remoteaddr, testTarget)
		if err != nil && len(ns) == 0 {
			errc <- err
		} else {
			resultc <- ns
		}
	}()

	// wait for the findnode to be sent.
	// after it is sent, the transport is waiting for a reply
	test.waitPacketOut(func(p *findnode) {
		if p.Target != testTarget {
			t.Errorf("wrong target: got %v, want %v", p.Target, testTarget)
		}
	})

	// send the reply as two packets.
	list := []*node{
		wrapNode(enode.NewV4(parsePubkey("ba85011c70bcc5c04d8607d3a0ed29aa6179c092cbdda10d5d32684fb33ed01bd94f588ca8f91ac48318087dcb02eaf36773a7a453f0eedd6742af668097b29c"), net.ParseIP("10.0.1.16"), 30303, 30304, 0)),
		wrapNode(enode.NewV4(parsePubkey("81fa361d25f157cd421c60dcc28d8dac5ef6a89476633339c5df30287474520caca09627da18543d9079b5b288698b542d56167aa5c09111e55acdbbdf2ef799"), net.ParseIP("10.0.1.16"), 30303, 30303, 0)),
		wrapNode(enode.NewV4(parsePubkey("9bffefd833d53fac8e652415f4973bee289e8b1a5c6c4cbe70abf817ce8a64cee11b823b66a987f51aaa9fba0d6a91b3e6bf0d5a5d1042de8e9eeea057b217f8"), net.ParseIP("10.0.1.36"), 30301, 17, 0)),
		wrapNode(enode.NewV4(parsePubkey("1b5b4aa662d7cb44a7221bfba67302590b643028197a7d5214790f3bac7aaa4a3241be9e83c09cf1f6c69d007c634faae3dc1b1221793e8446c0b3a09de65960"), net.ParseIP("10.0.1.16"), 30303, 30303, 0)),
	}
	rpclist := make([]rpcNode, len(list))
	for i := range list {
		rpclist[i] = nodeToRPC(list[i])
	}
	test.packetIn(nil, neighborsPacket, &neighbors{Expiration: futureExp, Nodes: rpclist[:2]})
	test.packetIn(nil, neighborsPacket, &neighbors{Expiration: futureExp, Nodes: rpclist[2:]})

	// check that the sent neighbors are all returned by findnode
	select {
	case result := <-resultc:
		want := append(list[:2], list[3:]...)
		if !reflect.DeepEqual(result, want) {
			t.Errorf("neighbors mismatch:\n  got:  %v\n  want: %v", result, want)
		}
	case err := <-errc:
		t.Errorf("findnode error: %v", err)
	case <-time.After(5 * time.Second):
		t.Error("findnode did not return within 5 seconds")
	}
}



func TestUDPExisting_successfulPing(t *testing.T) {
	test := newUDPTestUsingExistingTable(t)
	added := make(chan *node, 1)
	test.table.nodeAddedHook = func(n *node) { added <- n }
	defer test.table.Close()

	// The remote side sends a ping packet to initiate the exchange.
	go test.packetIn(nil, pingPacket, &ping{From: testRemote, To: testLocalAnnounced, Version: 4, Expiration: futureExp})

	// the ping is replied to.
	test.waitPacketOut(func(p *pong) {
		pinghash := test.sent[0][:macSize]
		if !bytes.Equal(p.ReplyTok, pinghash) {
			t.Errorf("got pong.ReplyTok %x, want %x", p.ReplyTok, pinghash)
		}
		wantTo := rpcEndpoint{
			// The mirrored UDP address is the UDP packet sender
			IP: test.remoteaddr.IP, UDP: uint16(test.remoteaddr.Port),
			// The mirrored TCP port is the one from the ping packet
			TCP: testRemote.TCP,
		}
		if !reflect.DeepEqual(p.To, wantTo) {
			t.Errorf("got pong.To %v, want %v", p.To, wantTo)
		}
	})

	// remote is unknown, the table pings back.
	hash, _ := test.waitPacketOut(func(p *ping) error {
		if !reflect.DeepEqual(p.From, test.udp.ourEndpoint()) {
			t.Errorf("got ping.From %#v, want %#v", p.From, test.udp.ourEndpoint())
		}
		wantTo := rpcEndpoint{
			// The mirrored UDP address is the UDP packet sender.
			IP:  test.remoteaddr.IP,
			UDP: uint16(test.remoteaddr.Port),
			TCP: 0,
		}
		if !reflect.DeepEqual(p.To, wantTo) {
			t.Errorf("got ping.To %v, want %v", p.To, wantTo)
		}
		return nil
	})
	test.packetIn(nil, pongPacket, &pong{ReplyTok: hash, Expiration: futureExp})

	// the node should be added to the table shortly after getting the
	// pong packet.
	select {
	case n := <-added:
		rid := encodePubkey(&test.remotekey.PublicKey).id()
		if n.ID() != rid {
			t.Errorf("node has wrong ID: got %v, want %v", n.ID(), rid)
		}
		if !n.IP().Equal(test.remoteaddr.IP) {
			t.Errorf("node has wrong IP: got %v, want: %v", n.IP(), test.remoteaddr.IP)
		}
		if int(n.UDP()) != test.remoteaddr.Port {
			t.Errorf("node has wrong UDP port: got %v, want: %v", n.UDP(), test.remoteaddr.Port)
		}
		if n.TCP() != int(testRemote.TCP) {
			t.Errorf("node has wrong TCP port: got %v, want: %v", n.TCP(), testRemote.TCP)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("node was not added within 2 seconds")
	}
}

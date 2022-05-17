// Copyright 2018 The go-ethereum Authors
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

package enode

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
	"strings"
	"testing"
	"testing/quick"

	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

var pyRecord, _ = hex.DecodeString("f884b8407098ad865b00a582051940cb9cf36836572411a47278783077011599ed5cd16b76f2635f4e234738f30813a89eb9137e3e3df5266e3a1f11df72ecf1145ccb9c01826964827634826970847f00000189736563703235366b31a103ca634cae0d49acb401d8a4c6b6fe8c55b70d115bf400769cc1400f3258cd31388375647082765f")

// TestPythonInterop checks that we can decode and verify a record produced by the Python
// implementation.
func TestPythonInterop(t *testing.T) {
	var r enr.Record
	if err := rlp.DecodeBytes(pyRecord, &r); err != nil {
		t.Fatalf("can't decode: %v", err)
	}
	n, err := New(ValidSchemes, &r)
	if err != nil {
		t.Fatalf("can't verify record: %v", err)
	}

	var (
		wantID  = HexID("a448f24c6d18e575453db13171562b71999873db5b286df957af199ec94617f7")
		wantSeq = uint64(1)
		wantIP  = enr.IPv4{127, 0, 0, 1}
		wantUDP = enr.UDP(30303)
	)
	if n.Seq() != wantSeq {
		t.Errorf("wrong seq: got %d, want %d", n.Seq(), wantSeq)
	}
	if n.ID() != wantID {
		t.Errorf("wrong id: got %x, want %x", n.ID(), wantID)
	}
	want := map[enr.Entry]interface{}{new(enr.IPv4): &wantIP, new(enr.UDP): &wantUDP}
	for k, v := range want {
		desc := fmt.Sprintf("loading key %q", k.ENRKey())
		if assert.NoError(t, n.Load(k), desc) {
			assert.Equal(t, k, v, desc)
		}
	}
}

func TestHexID(t *testing.T) {
	ref := ID{0, 0, 0, 0, 0, 0, 0, 128, 106, 217, 182, 31, 165, 174, 1, 67, 7, 235, 220, 150, 66, 83, 173, 205, 159, 44, 10, 57, 42, 161, 26, 188}
	id1 := HexID("0x00000000000000806ad9b61fa5ae014307ebdc964253adcd9f2c0a392aa11abc")
	id2 := HexID("00000000000000806ad9b61fa5ae014307ebdc964253adcd9f2c0a392aa11abc")

	if id1 != ref {
		t.Errorf("wrong id1\ngot  %v\nwant %v", id1[:], ref[:])
	}
	if id2 != ref {
		t.Errorf("wrong id2\ngot  %v\nwant %v", id2[:], ref[:])
	}
}

func TestID_textEncoding(t *testing.T) {
	ref := ID{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30,
		0x31, 0x32,
	}
	hex := "0102030405060708091011121314151617181920212223242526272829303132"

	text, err := ref.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(text, []byte(hex)) {
		t.Fatalf("text encoding did not match\nexpected: %s\ngot:      %s", hex, text)
	}

	id := new(ID)
	if err := id.UnmarshalText(text); err != nil {
		t.Fatal(err)
	}
	if *id != ref {
		t.Fatalf("text decoding did not match\nexpected: %s\ngot:      %s", ref, id)
	}
}

func TestID_distcmp(t *testing.T) {
	distcmpBig := func(target, a, b ID) int {
		tbig := new(big.Int).SetBytes(target[:])
		abig := new(big.Int).SetBytes(a[:])
		bbig := new(big.Int).SetBytes(b[:])
		return new(big.Int).Xor(tbig, abig).Cmp(new(big.Int).Xor(tbig, bbig))
	}
	if err := quick.CheckEqual(DistCmp, distcmpBig, nil); err != nil {
		t.Error(err)
	}
}

// The random tests is likely to miss the case where a and b are equal,
// this test checks it explicitly.
func TestID_distcmpEqual(t *testing.T) {
	base := ID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	x := ID{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if DistCmp(base, x, x) != 0 {
		t.Errorf("DistCmp(base, x, x) != 0")
	}
}

func TestID_logdist(t *testing.T) {
	logdistBig := func(a, b ID) int {
		abig, bbig := new(big.Int).SetBytes(a[:]), new(big.Int).SetBytes(b[:])
		return new(big.Int).Xor(abig, bbig).BitLen()
	}
	if err := quick.CheckEqual(LogDist, logdistBig, nil); err != nil {
		t.Error(err)
	}
}

// The random tests is likely to miss the case where a and b are equal,
// this test checks it explicitly.
func TestID_logdistEqual(t *testing.T) {
	x := ID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	if LogDist(x, x) != 0 {
		t.Errorf("LogDist(x, x) != 0")
	}
}

// Quorum
//
// test raft port in node detail
func TestNodeInfoForRaftPort(t *testing.T) {
	node := NewV4Hostname(
		hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
		"192.168.0.1",
		30302,
		30303,
		2021,
	)
	wantIP := enr.IPv4{192, 168, 0, 1}
	wantUdp := 30303
	wantTcp := 30302
	wantRaftPort := 2021
	assert.Equal(t, wantRaftPort, node.RaftPort(), "node raft port mismatch")
	assert.Equal(t, net.IP(wantIP), node.IP(), "node ip mismatch")
	assert.Equal(t, wantUdp, node.UDP(), "node UDP port mismatch")
	assert.Equal(t, wantTcp, node.TCP(), "node TCP port mismatch")

}

// Quorum - test parsing url with hostname (if host is FQDN)
func TestNodeParseUrlWithHostnameForQuorum(t *testing.T) {
	var url = "enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@localhost:21000?discport=0&raftport=50401"
	n, err := ParseV4(url)
	if err != nil {
		t.Errorf("parsing host failed %v", err)
	}
	assert.Equal(t, 50401, n.RaftPort())

	url = "enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@localhost1:21000?discport=0&raftport=50401"
	_, err = ParseV4(url)
	if err != nil {
		errMsg := err.Error()
		hasError := strings.Contains(errMsg, "no such host")

		assert.Equal(t, true, hasError, err.Error())
	}
}

// Quorum
// test Incomplete
func TestIncomplete(t *testing.T) {
	var testCases = []struct {
		n            *Node
		isIncomplete bool
	}{
		{
			n: NewV4(
				hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
				net.IP{127, 0, 0, 1},
				52150,
				52150,
			),
			isIncomplete: false,
		},
		{
			n: NewV4(hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
				net.ParseIP("::"),
				52150,
				52150,
			),
			isIncomplete: false,
		},
		{
			n: NewV4(
				hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
				net.ParseIP("2001:db8:3c4d:15::abcd:ef12"),
				52150,
				52150,
			),
			isIncomplete: false,
		},
		{
			n: NewV4(
				hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
				nil,
				52150,
				52150,
			),
			isIncomplete: true,
		},
		{
			n: NewV4Hostname(
				hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
				"hostname",
				52150,
				52150,
				50400,
			),
			isIncomplete: false,
		},
		{
			n: NewV4Hostname(
				hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
				"hostname",
				52150,
				52150,
				0,
			),
			isIncomplete: true,
		},
		{
			n: NewV4Hostname(
				hexPubkey("1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439"),
				"",
				52150,
				52150,
				50400,
			),
			isIncomplete: true,
		},
	}

	for i, test := range testCases {
		if test.n.Incomplete() != test.isIncomplete {
			t.Errorf("test %d: Node.Incomplete() mismatch:\ngot:  %v\nwant: %v", i, test.n.Incomplete(), test.isIncomplete)
		}
	}
}

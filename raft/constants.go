package gethRaft

import (
	"github.com/coreos/etcd/raft"
)

const (
	protocolName           = "raft"
	protocolVersion uint64 = 0x01

	raftMsg = 0x00

	minterRole   = raft.LEADER
	verifierRole = raft.NOT_LEADER

	// Raft's ticker interval
	tickerMS = 100

	// We use a bounded channel of constant size buffering incoming messages
	msgChanSize = 1000

	// Snapshot after this many messages
	// Our snapshots are *super* cheap -- much cheaper than raft naively assumes
	// -- because we only store the latest block hash.
	// We might be able to get away with snapshotting *every* entry.
	defaultSnapCount = 100

	// checkpoints
	PEER_CONNECTED = "PEER-CONNECTED"
	PEER_DISCONNECTED = "PEER-DISCONNECTED"
	TX_CREATED = "TX-CREATED"
	TX_ACCEPTED = "TX-ACCEPTED"
	BECAME_MINTER = "BECAME-MINTER"
	BECAME_VERIFIER = "BECAME-VERIFIER"
)

package raft

import (
	etcdRaft "github.com/coreos/etcd/raft"
)

const (
	protocolName           = "raft"
	protocolVersion uint64 = 0x01

	raftMsg = 0x00

	minterRole   = etcdRaft.LEADER
	verifierRole = etcdRaft.NOT_LEADER

	// Raft's ticker interval
	tickerMS = 100

	// We use a bounded channel of constant size buffering incoming messages
	msgChanSize = 1000

	// Snapshot after this many raft messages
	//
	// TODO: measure and get this as low as possible without affecting performance
	//
	snapshotPeriod = 250

	peerUrlKeyPrefix = "peerUrl-"

	// checkpoints
	peerConnected    = "PEER-CONNECTED"
	peerDisconnected = "PEER-DISCONNECTED"
	txCreated        = "TX-CREATED"
	txAccepted       = "TX-ACCEPTED"
	becameMinter     = "BECAME-MINTER"
	becameVerifier   = "BECAME-VERIFIER"
)

var (
	appliedDbKey = []byte("applied")
)

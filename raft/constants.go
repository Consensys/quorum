package raft

import etcd "github.com/ethereum/go-ethereum/raft/etcd/raft"

const (
	//protocolName           = "raft"
	//protocolVersion uint64 = 0x01

	//raftMsg = 0x00

	minterRole = etcd.LEADER
	//verifierRole = etcdRaft.NOT_LEADER

	// Raft's ticker interval
	tickerMS = 100

	// We use a bounded channel of constant size buffering incoming messages
	//msgChanSize = 1000

	// Snapshot after this many raft messages
	//
	// TODO: measure and get this as low as possible without affecting performance
	//
	snapshotPeriod = 250

	//peerUrlKeyPrefix = "peerUrl-"
)

var (
	appliedDbKey = []byte("applied")
)

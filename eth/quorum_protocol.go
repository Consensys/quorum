package eth

import (
	"errors"

	"github.com/ethereum/go-ethereum/p2p"
)

// Quorum: quorum_protocol enables the eth service to return two different protocols, one for the eth mainnet "eth" service,
//         and one for the quorum specific consensus algo, obtained from engine.consensus
//         2021 Jan in the future consensus (istanbul) may run from its own service and use a single subprotocol there,
//         instead of overloading the eth service.

var (
	// errEthPeerNil is returned when no eth peer is found to be associated with a p2p peer.
	errEthPeerNil           = errors.New("eth peer was nil")
	errEthPeerNotRegistered = errors.New("eth peer was not registered")
)

// quorum consensus Protocol variables are optionally set in addition to the "eth" protocol variables (eth/protocol.go).
var quorumConsensusProtocolName = ""

// ProtocolVersions are the supported versions of the quorum consensus protocol (first is primary), e.g. []uint{Istanbul64, Istanbul99, Istanbul100}.
var quorumConsensusProtocolVersions []uint

// protocol Length describe the number of messages support by the protocol/version map[uint]uint64{Istanbul64: 18, Istanbul99: 18, Istanbul100: 18}
var quorumConsensusProtocolLengths map[uint]uint64

func (s *Ethereum) quorumConsensusProtocols() []p2p.Protocol {
	protos := make([]p2p.Protocol, len(quorumConsensusProtocolVersions))
	for i, vsn := range quorumConsensusProtocolVersions {
		// if we have a legacy protocol, e.g. istanbul/99, istanbul/64 then the protocol handler is will be the "eth"
		// protocol handler, and the subprotocol "eth" will not be used, but rather the legacy subprotocol will handle
		// both eth messages and consensus messages.
		if isLegacyProtocol(quorumConsensusProtocolName, vsn) {
			length, ok := quorumConsensusProtocolLengths[vsn]
			if !ok {
				panic("makeProtocol for unknown version")
			}
			lp := s.handler.makeLegacyProtocol(quorumConsensusProtocolName, vsn, length)
			protos[i] = lp
		} else {
			length, ok := quorumConsensusProtocolLengths[vsn]
			if !ok {
				panic("makeQuorumConsensusProtocol for unknown version")
			}
			protos[i] = s.handler.makeQuorumConsensusProtocol(quorumConsensusProtocolName, vsn, length)
		}
	}
	return protos
}

// istanbul/64, istanbul/99, clique/63, clique/64 all override the "eth" subprotocol.
func isLegacyProtocol(name string, version uint) bool {
	// protocols that override "eth" subprotocol and run only the quorum subprotocol.
	quorumLegacyProtocols := map[string][]uint{"istanbul": {64, 99}, "clique": {63, 64}}
	for lpName, lpVersions := range quorumLegacyProtocols {
		if lpName == name {
			for _, v := range lpVersions {
				if v == version {
					return true
				}
			}
		}
	}
	return false
}

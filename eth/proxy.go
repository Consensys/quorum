package eth

// We need access to a few private ethereum backen members from raft. These
// getters allow raft to access them.

import (
	"math/big"

	"github.com/ethereum/go-ethereum/eth/downloader"
)

type RaftEthProxy struct {
	Downloader *downloader.Downloader
	// Find the best peer (furthest ahead), but filtered down to only the fields
	// that raft needs to synchronize with it
	GetBestRaftPeer func() (string, *big.Int)
}

func (s *Ethereum) GetProxy() RaftEthProxy {
	pm := s.protocolManager

	return RaftEthProxy{
		Downloader: pm.downloader,
		GetBestRaftPeer: func() (string, *big.Int) {
			peer := pm.peers.BestPeer()
			return peer.id, peer.td
		},
	}
}

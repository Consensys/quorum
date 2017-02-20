package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/core/quorum"
	"github.com/ethereum/go-ethereum/rpc"
)

func (s *Ethereum) StartBlockVoting(client *rpc.Client, voteKey, blockMakerKey *ecdsa.PrivateKey) error {
	s.blockMakerStrat = quorum.NewRandomDeadelineStrategy(s.eventMux, s.minBlockTime, s.maxBlockTime, s.minVoteTime, s.maxVoteTime)
	quorum.Strategy = s.blockMakerStrat
	return s.blockVoting.Start(client, s.blockMakerStrat, voteKey, blockMakerKey)
}

package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/core/quorum"
	"github.com/ethereum/go-ethereum/rpc"
)

func (s *Ethereum) StartBlockVoting(client *rpc.Client, voteKey, blockMakerKey *ecdsa.PrivateKey) error {
	activateVoting, activateBlockCreation := voteKey != nil, blockMakerKey != nil
	strat := quorum.NewRandomDeadelineStrategy(s.eventMux, s.minBlockTime, s.maxBlockTime, s.minVoteTime, s.maxVoteTime, activateVoting, activateBlockCreation)

	s.blockMakerStrat = strat
	quorum.Strategy = strat

	return s.blockVoting.Start(client, s.blockMakerStrat, voteKey, blockMakerKey)
}

package quorum

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	Strategy BlockMakerStrategy
)

type PublicQuorumAPI struct {
	bv *BlockVoting
}

func NewPublicQuorumAPI(bv *BlockVoting) *PublicQuorumAPI {
	return &PublicQuorumAPI{bv}
}

func (api *PublicQuorumAPI) CanonicalHash(height rpc.HexNumber) (common.Hash, error) {
	return api.bv.canonHash(height.Uint64())
}

func (api *PublicQuorumAPI) Vote(blockHash common.Hash) (common.Hash, error) {
	pBlock, _, _ := api.bv.Pending()

	req := Vote{
		Hash:   blockHash,
		Number: pBlock.Number(),
		TxHash: make(chan common.Hash),
		Err:    make(chan error),
	}

	if err := api.bv.mux.Post(req); err != nil {
		return common.Hash{}, err
	}

	select {
	case h := <-req.TxHash:
		return h, nil
	case err := <-req.Err:
		return common.Hash{}, err
	case <-time.NewTimer(30 * time.Second).C:
		return common.Hash{}, fmt.Errorf("timeout vote request")
	}
}

func (api *PublicQuorumAPI) NodeInfo() map[string]interface{} {
	result := make(map[string]interface{})

	if api.bv.bmk != nil {
		addr := crypto.PubkeyToAddress(api.bv.bmk.PublicKey)
		allowed, _ := api.bv.callContract.IsBlockMaker(nil, addr)
		result["blockMakerAccount"] = addr
		result["canCreateBlocks"] = allowed
		if Strategy != nil {
			result["blockmakestrategy"] = Strategy
		}
	}

	if api.bv.vk != nil {
		addr := crypto.PubkeyToAddress(api.bv.vk.PublicKey)
		allowed, _ := api.bv.callContract.IsVoter(nil, addr)
		result["voteAccount"] = addr
		result["canVote"] = allowed
	}

	return result
}

func (api *PublicQuorumAPI) MakeBlock() (common.Hash, error) {
	req := CreateBlock{
		Hash: make(chan common.Hash),
		Err:  make(chan error),
	}

	if err := api.bv.mux.Post(req); err != nil {
		return common.Hash{}, err
	}

	select {
	case h := <-req.Hash:
		return h, nil
	case err := <-req.Err:
		return common.Hash{}, err
	case <-time.NewTimer(30 * time.Second).C:
		return common.Hash{}, fmt.Errorf("timeout block make request")
	}
}

func (api *PublicQuorumAPI) IsVoter(addr common.Address) (bool, error) {
	return api.bv.isVoter(addr)
}

func (api *PublicQuorumAPI) IsBlockMaker(addr common.Address) (bool, error) {
	return api.bv.isBlockMaker(addr)
}

func (api *PublicQuorumAPI) PauseBlockMaker() error {
	if Strategy != nil {
		return Strategy.Pause()
	}
	return nil
}

func (api PublicQuorumAPI) ResumeBlockMaker() error {
	if Strategy != nil {
		return Strategy.Resume()
	}
	return nil
}

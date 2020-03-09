package extension

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/extension/extensionContracts"
)

var (
	//Log queries
	newExtensionQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.NewContractExtensionContractCreatedTopicHash)}},
		Addresses: []common.Address{},
	}

	finishedExtensionQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.ExtensionFinishedTopicHash)}},
		Addresses: []common.Address{},
	}

	voteCompletedQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.AllNodesHaveVotedTopicHash)}},
		Addresses: []common.Address{},
	}

	newVoteQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.NewVoteTopicHash)}},
		Addresses: []common.Address{},
	}

	canPerformStateShareQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.CanPerformStateShareTopicHash)}},
		Addresses: []common.Address{},
	}
)

type ExtensionContract struct {
	Address                   common.Address `json:"address"`
	Initiator                 common.Address `json:"initiator"`
	ManagementContractAddress common.Address `json:"managementcontractaddress"`
	CreationData              []byte         `json:"creationData"`
}

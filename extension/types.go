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

	canPerformStateShareQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.CanPerformStateShareTopicHash)}},
		Addresses: []common.Address{},
	}
)

type ExtensionContract struct {
	ContractExtended          common.Address `json:"contractExtended"`
	Initiator                 common.Address `json:"initiator"`
	Recipient                 common.Address `json:"recipient"`
	ManagementContractAddress common.Address `json:"managementContractAddress"`
	RecipientPtmKey           string         `json:"recipientPtmKey"`
	CreationData              []byte         `json:"creationData"`
}

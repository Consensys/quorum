package extension

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/private"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/extension/extensionContracts"
)

// generateUuid sends some data to the linked Private Transaction Manager which
// uses a randomly generated key to encrypt the data and then hash it this
// means we get a effectively random hash, whilst also having a reference
// transaction inside the PTM
func generateUuid(contractAddress common.Address, privateFrom string, ptm private.PrivateTransactionManager) (string, error) {
	hash, err := ptm.Send(contractAddress.Bytes(), privateFrom, []string{})
	if err != nil {
		return "", err
	}
	return common.BytesToEncryptedPayloadHash(hash).String(), nil
}

func getAllVoters(addressToVoteOn common.Address, client *ethclient.Client) ([]common.Address, error){
	caller, err := extensionContracts.NewContractExtenderCaller(addressToVoteOn, client)
	if err != nil {
		return nil, err
	}
	numberOfVoters, err := caller.TotalNumberOfVoters(nil)
	if err != nil {
		return nil, err
	}
	var i int64
	var voters []common.Address
	for i = 0; i < numberOfVoters.Int64(); i++ {
		voter, err := caller.WalletAddressesToVote(nil, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		voters = append(voters, voter)
	}
	return voters, nil
}

func checkAddressInList(addressToFind common.Address, addressList []common.Address) bool {
	for _, addr := range addressList {
		if addressToFind == addr {
			return true
		}
	}
	return false
}

func unpackNewExtension(data []byte) (*extensionContracts.ContractExtenderNewContractExtensionContractCreated, error){
	newExtensionEvent := new(extensionContracts.ContractExtenderNewContractExtensionContractCreated)
	err := extensionContracts.ContractExtenderParsedABI.Unpack(newExtensionEvent, "NewContractExtensionContractCreated", data)

	return newExtensionEvent, err
}
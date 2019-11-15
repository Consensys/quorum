package extension

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/private"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
)

func getAddressState(privateState *state.StateDB, addressToShare common.Address) []byte {
	keepAddresses := make(map[string]extensionContracts.AccountWithMetadata)

	if account, found := privateState.DumpAddress(addressToShare); found {
		keepAddresses[addressToShare.Hex()] = extensionContracts.AccountWithMetadata{
			State: account,
		}
	}

	//types can be marshalled, so errors can't occur
	out, _ := json.Marshal(&keepAddresses)
	return out
}

func generateTransactOpts(accountManager *accounts.Manager, txa ethapi.SendTxArgs) (*bind.TransactOpts, error) {
	//Find the account we plan to send the transaction from
	frmAcct := accounts.Account{Address: txa.From}
	wallet, err := accountManager.Find(frmAcct)
	if err != nil {
		return nil, fmt.Errorf("no wallet found for account %s", txa.From.String())
	}

	txArgs := bind.NewWalletTransactor(wallet, frmAcct)
	txArgs.PrivateFrom = txa.PrivateFrom
	txArgs.PrivateFor = txa.PrivateFor
	txArgs.GasLimit = defaultGasLimit
	txArgs.GasPrice = defaultGasPrice

	if txa.GasPrice != nil {
		txArgs.GasPrice = txa.GasPrice.ToInt()
	}
	if txa.Gas != nil {
		txArgs.GasLimit = uint64(*txa.Gas)
	}
	return txArgs, nil
}

func writeContentsToFile(extensionContracts map[common.Address]*ExtensionContract, datadir string) error {
	//no unmarshallable types, so can't error
	output, _ := json.Marshal(&extensionContracts)

	path := filepath.Join(datadir, extensionContractData)
	if errSaving := ioutil.WriteFile(path, output, 0644); errSaving != nil {
		log.Error("Couldn't save outstanding extension contract details")
		return errSaving
	}
	return nil
}

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
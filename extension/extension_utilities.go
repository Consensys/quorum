package extension

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
)

// generateUuid sends some data to the linked Private Transaction Manager which
// uses a randomly generated key to encrypt the data and then hash it this
// means we get a effectively random hash, whilst also having a reference
// transaction inside the PTM
func generateUuid(contractAddress common.Address, privateFrom string, privateFor []string, ptm private.PrivateTransactionManager) (string, error) {

	// to ensure recoverability , the UUID generation logic is as below:
	// 1. Call Tessera to encrypt the management contract address
	// 2. Send the encrypted payload to all participants on the contract extension
	// 3. Use the received hash as the UUID
	payloadHash, err := ptm.EncryptPayload(contractAddress.Bytes(), privateFrom, []string{}, &engine.ExtraMetadata{})
	if err != nil {
		return "", err
	}

	hash, err := ptm.Send(payloadHash, privateFrom, privateFor, &engine.ExtraMetadata{})
	if err != nil {
		return "", err
	}
	return hash.String(), nil
}

func checkAddressInList(addressToFind common.Address, addressList []common.Address) bool {
	for _, addr := range addressList {
		if addressToFind == addr {
			return true
		}
	}
	return false
}

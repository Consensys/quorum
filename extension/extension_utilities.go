package extension

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private"
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

func checkAddressInList(addressToFind common.Address, addressList []common.Address) bool {
	for _, addr := range addressList {
		if addressToFind == addr {
			return true
		}
	}
	return false
}

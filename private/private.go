package private

import (
	"os"

	"github.com/ethereum/go-ethereum/private/constellation"
)

// PrivateTransactionManager is the interface to the private transaction data
// stored in Quorum. It exposes a simple API to send and receive data. This type
// specifies the main way quorum interacts with constellation.
type PrivateTransactionManager interface {
	// Send is used to send data to store in the PrivateTransactionManager.
	// from is the sending party's base64-encoded public key to use.
	// to is an array of the recipients' base64-encoded public keys.
	// the private transaction manager will be responsible for storing and
	// distributing the data to all the relevant recipient nodes, and should only return after having done so.
	// Will return the hash key of
	// the data to be used to look it up.
	Send(data []byte, from string, to []string) ([]byte, error)
	// Used to get data from the private transaction manager.  The data []bytes
	// key specifies the digest of the data to extract. Will only successfully
	// extract the data if it is stored in the current nodes transaction manager.
	Receive(data []byte) ([]byte, error)
}

// FromEnvironmentOrNil performs a lookup using the config in the specified
// environment variable from the provided name string parameter.
// Internally this is done using the PRIVATE_CONFIG environment variable.
func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	return constellation.MustNew(cfgPath)
}

// Load the constant constellation P from the environment variable PRIVATE_CONFIG.
// This specifies the transaction manager used throughout Quorum.
var P = FromEnvironmentOrNil("PRIVATE_CONFIG")

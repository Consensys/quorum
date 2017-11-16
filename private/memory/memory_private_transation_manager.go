package memory

import (
	"crypto/sha512"
	"encoding/base64"
)

// MemoryPrivateTransactionManger implements an in memory representation of the Private Transaction Manager Interface
type MemoryPrivateTransactionManger struct {
	privDB map[string][]byte
}

// Send Sends a payload to list of "to" addresses
// Implements the PrivateTransactionManager.Send() interface
// Store payload into local repository, returning key (hash of payload)
func (g *MemoryPrivateTransactionManger) Send(data []byte, from string, to []string) (out []byte, err error) {
	h := sha512.New512_256()

	h.Write(data)

	out = h.Sum(nil)

	b64Key := base64.StdEncoding.EncodeToString(out)
	g.privDB[b64Key] = data

	//Out is hash key for retrieval of the payload
	return out, nil
}

// Receive Retrieve Payload for the key (data).
// Implements the PrivateTransactionManager.Receive() interface
func (g *MemoryPrivateTransactionManger) Receive(data []byte) ([]byte, error) {

	b64Key := base64.StdEncoding.EncodeToString(data)
	pl := g.privDB[b64Key]

	return pl, nil
}

// MustNew Instantiates the in memory database
func MustNew() *MemoryPrivateTransactionManger {
	return &MemoryPrivateTransactionManger{
		privDB: make(map[string][]byte),
	}
}

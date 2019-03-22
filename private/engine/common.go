package engine

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrPrivateTxManagerNotinUse     = errors.New("private transaction manager is not in use")
	ErrPrivateTxManagerNotReady     = errors.New("private transaction manager is not ready")
	ErrPrivateTxManagerNotSupported = errors.New("private transaction manager does not suppor this operation")
)

type NotInUsePrivateTxManager struct{}

func (dn *NotInUsePrivateTxManager) Send(data []byte, from string, to []string, extra *ExtraMetadata) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, ErrPrivateTxManagerNotinUse
}

func (dn *NotInUsePrivateTxManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *ExtraMetadata) ([]byte, error) {
	return nil, ErrPrivateTxManagerNotinUse
}

func (dn *NotInUsePrivateTxManager) Receive(data common.EncryptedPayloadHash) ([]byte, *ExtraMetadata, error) {
	return nil, nil, ErrPrivateTxManagerNotinUse
}

func (dn *NotInUsePrivateTxManager) Name() string {
	return "NotInUse"
}

// Additional information for the private transaction that Private Transaction Manager carries
type ExtraMetadata struct {
	// Hashes of affected Contracts
	ACHashes common.EncryptedPayloadHashes
	// Root Hash of a Merkle Trie containing all affected contract account in state objects
	ACMerkleRoot           common.Hash
	PrivateStateValidation bool
}

type Client struct {
	HttpClient *http.Client
	BaseURL    string
}

func (c *Client) FullPath(path string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, path)
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.HttpClient.Get(c.FullPath(path))
}

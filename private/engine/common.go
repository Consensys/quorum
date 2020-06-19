package engine

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrPrivateTxManagerNotinUse                          = errors.New("private transaction manager is not in use")
	ErrPrivateTxManagerNotReady                          = errors.New("private transaction manager is not ready")
	ErrPrivateTxManagerNotSupported                      = errors.New("private transaction manager does not support this operation")
	ErrPrivateTxManagerDoesNotSupportPrivacyEnhancements = errors.New("private transaction manager does not support privacy enhancements")
)

// Additional information for the private transaction that Private Transaction Manager carries
type ExtraMetadata struct {
	// Hashes of affected Contracts
	ACHashes common.EncryptedPayloadHashes
	// Root Hash of a Merkle Trie containing all affected contract account in state objects
	ACMerkleRoot common.Hash
	//Privacy flag for contract: standardPrivate, partyProtection, psv
	PrivacyFlag PrivacyFlagType
}

type Client struct {
	HttpClient *http.Client
	BaseURL    string
}

func (c *Client) FullPath(path string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, path)
}

func (c *Client) Get(path string) (*http.Response, error) {
	response, err := c.HttpClient.Get(c.FullPath(path))
	if err != nil {
		return response, fmt.Errorf("unable to submit request (method:%s,path:%s). Cause: %v", "GET", path, err)
	}
	return response, err
}

type PrivacyFlagType uint64

const (
	PrivacyFlagStandardPrivate PrivacyFlagType = iota                              // 0
	PrivacyFlagPartyProtection PrivacyFlagType = 1 << PrivacyFlagType(iota-1)      // 1
	PrivacyFlagStateValidation                 = iota | PrivacyFlagPartyProtection // 3 which includes PrivacyFlagPartyProtection
)

func (f PrivacyFlagType) IsNotStandardPrivate() bool {
	return !f.IsStandardPrivate()
}

func (f PrivacyFlagType) IsStandardPrivate() bool {
	return f == PrivacyFlagStandardPrivate
}

func (f PrivacyFlagType) Has(other PrivacyFlagType) bool {
	return other&f == other
}

func (f PrivacyFlagType) HasAll(others ...PrivacyFlagType) bool {
	var all PrivacyFlagType
	for _, flg := range others {
		all = all | flg
	}
	return f.Has(all)
}

func (f PrivacyFlagType) Validate() error {
	if f == PrivacyFlagStandardPrivate || f == PrivacyFlagPartyProtection || f == PrivacyFlagStateValidation {
		return nil
	}
	return fmt.Errorf("invalid privacy flag")
}

type PrivateTransactionManagerFeature uint64

const (
	None                PrivateTransactionManagerFeature = iota                                          // 0
	PrivacyEnhancements PrivateTransactionManagerFeature = 1 << PrivateTransactionManagerFeature(iota-1) // 1
)

type FeatureSet struct {
	features uint64
}

func NewFeatureSet(features ...PrivateTransactionManagerFeature) *FeatureSet {
	var all uint64 = 0
	for _, feature := range features {
		all = all | uint64(feature)
	}
	return &FeatureSet{features: all}
}

func (p *FeatureSet) HasFeature(feature PrivateTransactionManagerFeature) bool {
	return uint64(feature)&p.features != 0
}

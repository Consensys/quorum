package engine

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	PrivacyGroupResident = "RESIDENT"
	PrivacyGroupLegacy   = "LEGACY"
	PrivacyGroupPantheon = "PANTHEON"
)

var (
	ErrPrivateTxManagerNotinUse                          = errors.New("private transaction manager is not in use")
	ErrPrivateTxManagerNotReady                          = errors.New("private transaction manager is not ready")
	ErrPrivateTxManagerNotSupported                      = errors.New("private transaction manager does not support this operation")
	ErrPrivateTxManagerDoesNotSupportPrivacyEnhancements = errors.New("private transaction manager does not support privacy enhancements")
	ErrPrivateTxManagerDoesNotSupportMandatoryRecipients = errors.New("private transaction manager does not support mandatory recipients")
)

type PrivacyGroup struct {
	Type           string   `json:"type"`
	Name           string   `json:"name"`
	PrivacyGroupId string   `json:"privacyGroupId"`
	Description    string   `json:"description"`
	From           string   `json:"from"`
	Members        []string `json:"members"`
}

// Additional information for the private transaction that Private Transaction Manager carries
type ExtraMetadata struct {
	// Hashes of affected Contracts
	ACHashes common.EncryptedPayloadHashes
	// Root Hash of a Merkle Trie containing all affected contract account in state objects
	ACMerkleRoot common.Hash
	// Privacy flag for contract: standardPrivate, partyProtection, psv
	PrivacyFlag PrivacyFlagType
	// Contract participants that are managed by the corresponding Tessera.
	// Being used in Multi Tenancy
	ManagedParties []string
	// The sender of the transaction
	Sender string
	// Recipients that are mandated to be included
	MandatoryRecipients []string
}

type QuorumPayloadExtra struct {
	Payload       string
	ExtraMetaData *ExtraMetadata
	IsSender      bool
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
	PrivacyFlagStandardPrivate PrivacyFlagType = iota
	PrivacyFlagPartyProtection
	PrivacyFlagMandatoryRecipients
	PrivacyFlagStateValidation
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
	if f == PrivacyFlagStandardPrivate || f == PrivacyFlagPartyProtection || f == PrivacyFlagMandatoryRecipients || f == PrivacyFlagStateValidation {
		return nil
	}
	return fmt.Errorf("invalid privacy flag")
}

type PrivateTransactionManagerFeature uint64

const (
	None                  PrivateTransactionManagerFeature = iota                                          // 0
	PrivacyEnhancements   PrivateTransactionManagerFeature = 1 << PrivateTransactionManagerFeature(iota-1) // 1
	MultiTenancy          PrivateTransactionManagerFeature = 1 << PrivateTransactionManagerFeature(iota-1) // 2
	MultiplePrivateStates PrivateTransactionManagerFeature = 1 << PrivateTransactionManagerFeature(iota-1) // 4
	MandatoryRecipients   PrivateTransactionManagerFeature = 1 << PrivateTransactionManagerFeature(iota-1) // 8
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

type ExtraMetaDataRLP ExtraMetadata

func (emd *ExtraMetadata) DecodeRLP(stream *rlp.Stream) error {
	// initialize the ACHashes map before passing it into the rlp decoder
	emd.ACHashes = make(map[common.EncryptedPayloadHash]struct{})
	emdRLP := (*ExtraMetaDataRLP)(emd)
	if err := stream.Decode(emdRLP); err != nil {
		return err
	}
	return nil
}

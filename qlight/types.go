package qlight

import (
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/private/engine/qlightptm"
	"github.com/ethereum/go-ethereum/rlp"
)

type PrivateStateRootHashValidator interface {
	ValidatePrivateStateRoot(blockHash common.Hash, blockPublicStateRoot common.Hash) error
}

type PrivateClientCache interface {
	PrivateStateRootHashValidator
	AddPrivateBlock(blockPrivateData BlockPrivateData) error
	CheckAndAddEmptyEntry(hash common.EncryptedPayloadHash)
}

type PrivateBlockDataResolver interface {
	PrepareBlockPrivateData(block *types.Block, psi string) (*BlockPrivateData, error)
}

type AuthManagerProvider func() security.AuthenticationManager

type AuthProvider interface {
	Initialize() error
	Authorize(token string, psi string) error
}

type CacheWithEmpty interface {
	Cache(privateTxData *qlightptm.CachablePrivateTransactionData) error
	CheckAndAddEmptyToCache(hash common.EncryptedPayloadHash)
}

type BlockPrivateData struct {
	BlockHash           common.Hash
	PSI                 types.PrivateStateIdentifier
	PrivateStateRoot    common.Hash
	PrivateTransactions []PrivateTransactionData
}

type QLightCacheKey struct {
	BlockHash common.Hash
	PSI       types.PrivateStateIdentifier
}

func (k *QLightCacheKey) String() string {
	bytes, err := rlp.EncodeToBytes(k)
	if err != nil {
		return err.Error()
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

type PrivateTransactionData struct {
	Hash     *common.EncryptedPayloadHash
	Payload  []byte
	Extra    *engine.ExtraMetadata
	IsSender bool
}

func (d *PrivateTransactionData) ToCachable() *qlightptm.CachablePrivateTransactionData {
	return &qlightptm.CachablePrivateTransactionData{
		Hash: *d.Hash,
		QuorumPrivateTxData: engine.QuorumPayloadExtra{
			Payload:       fmt.Sprintf("0x%x", d.Payload),
			ExtraMetaData: d.Extra,
			IsSender:      d.IsSender,
		},
	}
}

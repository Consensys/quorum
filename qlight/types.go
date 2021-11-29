package qlight

import (
	"encoding/base64"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
)

type PrivateStateRootHashValidator interface {
	ValidatePrivateStateRoot(blockHash common.Hash, blockPrivateStateRoot common.Hash) error
}

type PrivateClientCache interface {
	PrivateStateRootHashValidator
	AddPrivateBlock(blockPrivateData engine.BlockPrivatePayloads) error
	CheckAndAddEmptyEntry(hash common.EncryptedPayloadHash)
}

type ServerPrivateBlockData struct {
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

type PrivateBlockData struct {
	PrivateStateRoot    common.Hash
	PrivateTransactions []PrivateTransactionData
}

type PrivateTransactionData struct {
	Hash     *common.EncryptedPayloadHash
	Payload  []byte
	Extra    *engine.ExtraMetadata
	IsSender bool
}

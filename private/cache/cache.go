package cache

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	DefaultExpiration = 5 * time.Minute
	CleanupInterval   = 5 * time.Minute
)

type PrivateCacheItem struct {
	Payload      []byte
	ACHashes     common.EncryptedPayloadHashes // hashes of affected contracts
	ACMerkleRoot common.Hash                   // merkle root of all affected contracts
}

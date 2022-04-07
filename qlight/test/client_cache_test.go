package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/private/cache"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/private/engine/qlightptm"
	"github.com/ethereum/go-ethereum/qlight"
	"github.com/golang/mock/gomock"
	gocache "github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestClientCache_AddPrivateBlock(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	memDB := rawdb.NewMemoryDatabase()
	cacheWithEmpty := NewMockCacheWithEmpty(ctrl)
	gocache := gocache.New(cache.DefaultExpiration, cache.CleanupInterval)

	clientCache, _ := qlight.NewClientCacheWithEmpty(memDB, cacheWithEmpty, gocache)

	txHash1 := common.BytesToEncryptedPayloadHash([]byte("TXHash1"))
	ptd1 := qlight.PrivateTransactionData{
		Hash:    &txHash1,
		Payload: []byte("payload"),
		Extra: &engine.ExtraMetadata{
			ACHashes:            nil,
			ACMerkleRoot:        common.Hash{},
			PrivacyFlag:         0,
			ManagedParties:      nil,
			Sender:              "",
			MandatoryRecipients: nil,
		},
		IsSender: false,
	}
	blockPrivateData := qlight.BlockPrivateData{
		BlockHash:           common.StringToHash("BlockHash"),
		PSI:                 "",
		PrivateStateRoot:    common.StringToHash("PrivateStateRoot"),
		PrivateTransactions: []qlight.PrivateTransactionData{ptd1},
	}

	var capturedCacheItem *qlightptm.CachablePrivateTransactionData
	cacheWithEmpty.EXPECT().Cache(gomock.Any()).DoAndReturn(func(privateTxData *qlightptm.CachablePrivateTransactionData) error {
		capturedCacheItem = privateTxData
		return nil
	})

	clientCache.AddPrivateBlock(blockPrivateData)

	assert.Equal(fmt.Sprintf("0x%x", ptd1.Payload), capturedCacheItem.QuorumPrivateTxData.Payload)
	assert.Equal(ptd1.Hash, &capturedCacheItem.Hash)

	psr, _ := gocache.Get(blockPrivateData.BlockHash.ToBase64())
	assert.Equal(blockPrivateData.PrivateStateRoot.ToBase64(), psr)
}

func TestClientCache_ValidatePrivateStateRootSuccess(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	memDB := rawdb.NewMemoryDatabase()
	cacheWithEmpty := NewMockCacheWithEmpty(ctrl)
	gocache := gocache.New(cache.DefaultExpiration, cache.CleanupInterval)

	clientCache, _ := qlight.NewClientCacheWithEmpty(memDB, cacheWithEmpty, gocache)

	publicStateRoot := common.StringToHash("PublicStateRoot")
	blockPrivateData := qlight.BlockPrivateData{
		BlockHash:           common.StringToHash("BlockHash"),
		PSI:                 "",
		PrivateStateRoot:    common.StringToHash("PrivateStateRoot"),
		PrivateTransactions: []qlight.PrivateTransactionData{},
	}

	clientCache.AddPrivateBlock(blockPrivateData)
	rawdb.WritePrivateStateRoot(memDB, publicStateRoot, blockPrivateData.PrivateStateRoot)

	err := clientCache.ValidatePrivateStateRoot(blockPrivateData.BlockHash, publicStateRoot)

	assert.Nil(err)
}

func TestClientCache_ValidatePrivateStateRootMismatch(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	memDB := rawdb.NewMemoryDatabase()
	cacheWithEmpty := NewMockCacheWithEmpty(ctrl)
	gocache := gocache.New(cache.DefaultExpiration, cache.CleanupInterval)

	clientCache, _ := qlight.NewClientCacheWithEmpty(memDB, cacheWithEmpty, gocache)

	publicStateRoot := common.StringToHash("PublicStateRoot")
	blockPrivateData := qlight.BlockPrivateData{
		BlockHash:           common.StringToHash("BlockHash"),
		PSI:                 "",
		PrivateStateRoot:    common.StringToHash("PrivateStateRoot"),
		PrivateTransactions: []qlight.PrivateTransactionData{},
	}

	clientCache.AddPrivateBlock(blockPrivateData)
	rawdb.WritePrivateStateRoot(memDB, publicStateRoot, common.StringToHash("Mismatch"))

	err := clientCache.ValidatePrivateStateRoot(blockPrivateData.BlockHash, publicStateRoot)

	assert.Error(err)
}

func TestClientCache_ValidatePrivateStateRootNoDataInClientCache(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	memDB := rawdb.NewMemoryDatabase()
	cacheWithEmpty := NewMockCacheWithEmpty(ctrl)
	gocache := gocache.New(cache.DefaultExpiration, cache.CleanupInterval)

	clientCache, _ := qlight.NewClientCacheWithEmpty(memDB, cacheWithEmpty, gocache)

	publicStateRoot := common.StringToHash("PublicStateRoot")
	blockPrivateData := qlight.BlockPrivateData{
		BlockHash:           common.StringToHash("BlockHash"),
		PSI:                 "",
		PrivateStateRoot:    common.StringToHash("PrivateStateRoot"),
		PrivateTransactions: []qlight.PrivateTransactionData{},
	}

	rawdb.WritePrivateStateRoot(memDB, publicStateRoot, blockPrivateData.PrivateStateRoot)

	err := clientCache.ValidatePrivateStateRoot(blockPrivateData.BlockHash, publicStateRoot)

	assert.Nil(err)
}

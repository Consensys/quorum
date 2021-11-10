package qlight

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine/qlightptm"
)

var cachingTXManager *qlightptm.CachingProxyTxManager

func InitializeClientCache() (err error) {
	var ok bool
	cachingTXManager, ok = private.P.(*qlightptm.CachingProxyTxManager)
	if !ok {
		return fmt.Errorf("unable to initialize cachingTXManager")
	}
	return nil
}

func AddPrivateBlockToClientCache(key QLightCacheKey) error {
	return cachingTXManager.AddPrivateBlockToCache(key.String())
}

func CheckAndAddEmptyToClientCache(hash common.EncryptedPayloadHash) {
	cachingTXManager.CheckAndAddEmptyToCache(hash)
}

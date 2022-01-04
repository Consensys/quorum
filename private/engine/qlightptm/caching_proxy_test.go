package qlightptm

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCachingProxy_ReceiveWithDataAvailableInCache(t *testing.T) {
	assert := assert.New(t)

	cpTM := New()
	cpTM.Cache(&CachablePrivateTransactionData{
		Hash: common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")),
		QuorumPrivateTxData: engine.QuorumPayloadExtra{
			Payload: fmt.Sprintf("0x%x", []byte("payload")),
			ExtraMetaData: &engine.ExtraMetadata{
				ACHashes:            nil,
				ACMerkleRoot:        common.Hash{},
				PrivacyFlag:         0,
				ManagedParties:      nil,
				Sender:              "sender1",
				MandatoryRecipients: nil,
			},
			IsSender: false,
		},
	})

	cpTM.Cache(&CachablePrivateTransactionData{
		Hash: common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash2")),
		QuorumPrivateTxData: engine.QuorumPayloadExtra{
			Payload: fmt.Sprintf("0x%x", []byte("payload")),
			ExtraMetaData: &engine.ExtraMetadata{
				ACHashes:            nil,
				ACMerkleRoot:        common.Hash{},
				PrivacyFlag:         0,
				ManagedParties:      nil,
				Sender:              "sender2",
				MandatoryRecipients: nil,
			},
			IsSender: true,
		},
	})

	sender, _, payload, extraMetaData, err := cpTM.Receive(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	assert.Nil(err)
	assert.Equal("sender1", sender)
	assert.Equal([]byte("payload"), payload)
	assert.NotNil(extraMetaData)

	isSender, err := cpTM.IsSender(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	assert.Nil(err)
	assert.False(isSender)

	isSender, err = cpTM.IsSender(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash2")))

	assert.Nil(err)
	assert.True(isSender)
}

func TestCachingProxy_ReceiveWithDataNotAvailableInCache(t *testing.T) {
	assert := assert.New(t)

	cpTM := New()
	cpTM.CheckAndAddEmptyToCache(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	_, _, payload, extraMetaData, err := cpTM.Receive(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	assert.Nil(err)
	assert.Nil(payload)
	assert.Nil(extraMetaData)

	isSender, err := cpTM.IsSender(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	assert.Nil(err)
	assert.False(isSender)
}

func TestCachingProxy_ReceiveWithDataMissingFromCacheAvailableRemotely(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cpTM := New()
	mockRPCClient := NewMockRPCClientCaller(ctrl)
	mockRPCClient.EXPECT().Call(gomock.Any(), gomock.Eq("eth_getQuorumPayloadExtra"),
		gomock.Eq(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")).Hex())).DoAndReturn(
		func(result interface{}, method string, args ...interface{}) error {
			res, _ := result.(*engine.QuorumPayloadExtra)
			res.IsSender = false
			res.ExtraMetaData = &engine.ExtraMetadata{
				ACHashes:            nil,
				ACMerkleRoot:        common.Hash{},
				PrivacyFlag:         0,
				ManagedParties:      nil,
				Sender:              "sender1",
				MandatoryRecipients: nil,
			}
			res.Payload = fmt.Sprintf("0x%x", []byte("payload"))
			return nil
		})

	cpTM.SetRPCClientCaller(mockRPCClient)

	sender, _, payload, extraMetaData, err := cpTM.Receive(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	assert.Nil(err)
	assert.Equal("sender1", sender)
	assert.Equal([]byte("payload"), payload)
	assert.NotNil(extraMetaData)
}

func TestCachingProxy_ReceiveWithDataMissingFromCacheUnavailableRemotely(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cpTM := New()
	mockRPCClient := NewMockRPCClientCaller(ctrl)
	mockRPCClient.EXPECT().Call(gomock.Any(), gomock.Eq("eth_getQuorumPayloadExtra"),
		gomock.Eq(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")).Hex())).Return(nil)

	cpTM.SetRPCClientCaller(mockRPCClient)

	sender, _, payload, extraMetaData, err := cpTM.Receive(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	assert.Nil(err)
	assert.Equal("", sender)
	assert.Nil(payload)
	assert.Nil(extraMetaData)
}

func TestCachingProxy_ReceiveWithDataMissingFromCacheAndRPCError(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cpTM := New()
	mockRPCClient := NewMockRPCClientCaller(ctrl)
	mockRPCClient.EXPECT().Call(gomock.Any(), gomock.Eq("eth_getQuorumPayloadExtra"),
		gomock.Eq(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")).Hex())).Return(fmt.Errorf("RPC Error"))

	cpTM.SetRPCClientCaller(mockRPCClient)

	_, _, _, _, err := cpTM.Receive(common.BytesToEncryptedPayloadHash([]byte("encryptedpayloadhash1")))

	assert.EqualError(err, "RPC Error")
}

func TestCachingProxy_DecryptPayloadSuccess(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cpTM := New()
	mockRPCClient := NewMockRPCClientCaller(ctrl)
	mockRPCClient.EXPECT().Call(gomock.Any(), gomock.Eq("eth_decryptQuorumPayload"),
		gomock.Any()).DoAndReturn(
		func(result interface{}, method string, args ...interface{}) error {
			res, _ := result.(*engine.QuorumPayloadExtra)
			res.IsSender = false
			res.ExtraMetaData = &engine.ExtraMetadata{
				ACHashes:            nil,
				ACMerkleRoot:        common.Hash{},
				PrivacyFlag:         0,
				ManagedParties:      nil,
				Sender:              "sender1",
				MandatoryRecipients: nil,
			}
			res.Payload = fmt.Sprintf("0x%x", []byte("payload"))
			return nil
		})

	cpTM.SetRPCClientCaller(mockRPCClient)

	payload, extraMetaData, err := cpTM.DecryptPayload(common.DecryptRequest{
		SenderKey:       []byte("sender1"),
		CipherText:      []byte("ciphertext"),
		CipherTextNonce: []byte("nonce"),
		RecipientBoxes:  nil,
		RecipientNonce:  []byte("nonce"),
		RecipientKeys:   nil,
	})

	assert.Nil(err)
	assert.Equal("sender1", extraMetaData.Sender)
	assert.Equal([]byte("payload"), payload)
	assert.NotNil(extraMetaData)
}

func TestCachingProxy_DecryptPayloadErrorInCall(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cpTM := New()
	mockRPCClient := NewMockRPCClientCaller(ctrl)
	mockRPCClient.EXPECT().Call(gomock.Any(), gomock.Eq("eth_decryptQuorumPayload"),
		gomock.Any()).Return(fmt.Errorf("RPC Error"))

	cpTM.SetRPCClientCaller(mockRPCClient)

	_, _, err := cpTM.DecryptPayload(common.DecryptRequest{
		SenderKey:       []byte("sender1"),
		CipherText:      []byte("ciphertext"),
		CipherTextNonce: []byte("nonce"),
		RecipientBoxes:  nil,
		RecipientNonce:  []byte("nonce"),
		RecipientKeys:   nil,
	})

	assert.EqualError(err, "RPC Error")
}

type HasRPCClient interface {
	SetRPCClient(client *rpc.Client)
}

func TestCachingProxy_HasRPCClient(t *testing.T) {
	assert := assert.New(t)
	var cpTM interface{}
	cpTM = New()

	_, ok := cpTM.(HasRPCClient)
	assert.True(ok)
}

package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/qlight"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

func TestBlockPrivateData_RLPEncodeDecode(t *testing.T) {
	txHash := common.BytesToEncryptedPayloadHash([]byte("EPH1"))
	data := qlight.BlockPrivateData{
		BlockHash:        common.BytesToHash([]byte("BlockHash")),
		PSI:              "PS1",
		PrivateStateRoot: common.BytesToHash([]byte("PSR")),
		PrivateTransactions: []qlight.PrivateTransactionData{
			{
				Hash:    &txHash,
				Payload: []byte("data"),
				Extra: &engine.ExtraMetadata{
					ACHashes:            common.EncryptedPayloadHashes{common.BytesToEncryptedPayloadHash([]byte("ACEPH1")): struct{}{}},
					ACMerkleRoot:        common.BytesToHash([]byte("root")),
					PrivacyFlag:         engine.PrivacyFlagPartyProtection,
					ManagedParties:      []string{"party1", "party2"},
					Sender:              "party3",
					MandatoryRecipients: []string{"party1"},
				},
				IsSender: false,
			}},
	}

	bytes, err := rlp.EncodeToBytes(data)
	assert.NoError(t, err)
	var decodedData qlight.BlockPrivateData
	err = rlp.DecodeBytes(bytes, &decodedData)
	assert.NoError(t, err)
	assert.Equal(t, data.PSI, decodedData.PSI)
	assert.Equal(t, data.BlockHash, decodedData.BlockHash)
	assert.Equal(t, data.PrivateStateRoot, decodedData.PrivateStateRoot)
	assert.Len(t, decodedData.PrivateTransactions, 1)
	privateTx := decodedData.PrivateTransactions[0]
	assert.Equal(t, &txHash, privateTx.Hash)
	assert.Equal(t, data.PrivateTransactions[0].Payload, privateTx.Payload)
	assert.Equal(t, data.PrivateTransactions[0].IsSender, privateTx.IsSender)
	assert.Equal(t, data.PrivateTransactions[0].Hash, privateTx.Hash)
	assert.Equal(t, data.PrivateTransactions[0].Extra.Sender, privateTx.Extra.Sender)
	assert.Equal(t, data.PrivateTransactions[0].Extra.ACMerkleRoot, privateTx.Extra.ACMerkleRoot)
	assert.Equal(t, data.PrivateTransactions[0].Extra.PrivacyFlag, privateTx.Extra.PrivacyFlag)
	assert.Equal(t, data.PrivateTransactions[0].Extra.MandatoryRecipients, privateTx.Extra.MandatoryRecipients)
	assert.Len(t, decodedData.PrivateTransactions[0].Extra.ACHashes, 1)
	_, found := decodedData.PrivateTransactions[0].Extra.ACHashes[common.BytesToEncryptedPayloadHash([]byte("ACEPH1"))]
	assert.True(t, found)
}

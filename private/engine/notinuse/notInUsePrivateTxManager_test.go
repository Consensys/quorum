package notinuse

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	ptm := &PrivateTransactionManager{}
	name := ptm.Name()

	assert.Equal(t, name, "NotInUse", "got wrong name for NotInUsePrivateTxManager")
}

func TestSendReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.Send([]byte{}, "", []string{}, nil)

	assert.Equal(t, err, engine.ErrPrivateTxManagerNotinUse, "got wrong error in 'send'")
}

func TestStoreRawReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.StoreRaw([]byte{}, "")

	assert.Equal(t, err, engine.ErrPrivateTxManagerNotinUse, "got wrong error in 'storeraw'")
}

func TestReceiveReturnsEmpty(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	data, metadata, err := ptm.Receive(common.EncryptedPayloadHash{})

	assert.Nil(t, err, "got unexpected error in 'receive'")
	assert.Nil(t, data, "got unexpected data in 'receive'")
	assert.Nil(t, metadata, "got unexpected metadata in 'receive'")
}

func TestReceiveRawReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, _, err := ptm.ReceiveRaw(common.EncryptedPayloadHash{})

	assert.Equal(t, err, engine.ErrPrivateTxManagerNotinUse, "got wrong error in 'send'")
}

func TestSendSignedTxReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.SendSignedTx(common.EncryptedPayloadHash{}, []string{}, nil)

	assert.Equal(t, err, engine.ErrPrivateTxManagerNotinUse, "got wrong error in 'SendSignedTx'")
}

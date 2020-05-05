package notinuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	ptm := &PrivateTransactionManager{}
	name := ptm.Name()

	assert.Equal(t, name, "NotInUse", "got wrong name for NotInUsePrivateTxManager")
}

func TestSendReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.Send([]byte{}, "", []string{})

	assert.Equal(t, err, ErrPrivateTxManagerNotInUse, "got wrong error in 'send'")
}

func TestStoreRawReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.StoreRaw([]byte{}, "")

	assert.Equal(t, err, ErrPrivateTxManagerNotInUse, "got wrong error in 'storeraw'")
}

func TestReceiveReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.Receive([]byte{})

	assert.Nil(t, err, "got unexpected error in 'receive'")
}

func TestSendSignedTxReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.SendSignedTx([]byte{}, []string{})

	assert.Equal(t, err, ErrPrivateTxManagerNotInUse, "got wrong error in 'SendSignedTx'")
}

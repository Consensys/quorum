package notinuse

import "testing"

func TestName(t *testing.T) {
	ptm := &PrivateTransactionManager{}
	name := ptm.Name()

	if name != "NotInUse" {
		t.Errorf("got wrong name for NotInUsePrivateTxManager. Expected '%s' but got '%s'", "NotInUse", name)
	}
}

func TestSendReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.Send([]byte{}, "", []string{})

	if err != ErrPrivateTxManagerNotInUse {
		t.Errorf("got wrong error in send. Expected '%s' but got '%s'", ErrPrivateTxManagerNotInUse, err)
	}
}

func TestReceiveReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.Receive([]byte{})

	if err != nil {
		t.Errorf("got wrong error in Receive. Expected nil but got '%s'", err)
	}
}

func TestSendSignedTxReturnsError(t *testing.T) {
	ptm := &PrivateTransactionManager{}

	_, err := ptm.SendSignedTx([]byte{}, []string{})

	if err != ErrPrivateTxManagerNotInUse {
		t.Errorf("got wrong error in SendSignedTx. Expected '%s' but got '%s'", ErrPrivateTxManagerNotInUse, err)
	}
}

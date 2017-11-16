package memory

import (
	"bytes"
	"testing"
)

func TestSendAndReceive(t *testing.T) {
	g := MustNew()

	data := []byte{1, 2, 3, 4, 5}
	to := []string{"A", "B", "C"}
	from := "sendNode"

	key, _ := g.Send(data, from, to)

	result, _ := g.Receive(key)

	if !bytes.Equal(data, result) {
		t.Errorf("Expected % x got % x", data, result)
	}
}

func TestReceiveWithoutSend(t *testing.T) {
	g := MustNew()

	key := []byte{1, 2}

	result, _ := g.Receive(key)

	if !bytes.Equal(nil, result) {
		t.Errorf("Expected nil got % x", result)
	}
}

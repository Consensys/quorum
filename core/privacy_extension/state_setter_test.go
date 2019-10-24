package privacy_extension

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

func TestLogContainsExtensionTopicWithWrongLengthReturnsFalse(t *testing.T) {
	testLog := &types.Log{
		Topics:      []common.Hash{{}, {}},
	}

	contained := LogContainsExtensionTopic(testLog)

	if contained {
		t.Errorf("expected value '%t', but got '%t'", false, contained)
	}
}

func TestLogContainsExtensionTopicWithWrongHashReturnsFalse(t *testing.T) {
	testLog := &types.Log{
		Topics:      []common.Hash{common.HexToHash("0xc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a2904")},
	}

	contained := LogContainsExtensionTopic(testLog)

	if contained {
		t.Errorf("expected value '%t', but got '%t'", false, contained)
	}
}

func TestLogContainsExtensionTopicWithCorrectHashReturnsTrue(t *testing.T) {
	testLog := &types.Log{
		Topics:      []common.Hash{common.HexToHash("0x40b79448ff8678eac1487385427aa682ee6ee831ce0702c09f95255645428531")},
	}

	contained := LogContainsExtensionTopic(testLog)

	if !contained {
		t.Errorf("expected value '%t', but got '%t'", true, contained)
	}
}
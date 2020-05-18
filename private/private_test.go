package private_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/private/engine/notinuse"

	"github.com/ethereum/go-ethereum/private"
)

func TestDummyPrivateTxManagerUsedWhenIgnoreSpecified(t *testing.T) {
	os.Setenv("PRIVATE_CONFIG", "ignore")

	ptm := private.FromEnvironmentOrNil("PRIVATE_CONFIG")

	if _, ok := ptm.(*notinuse.PrivateTransactionManager); !ok {
		expectedType := reflect.TypeOf(&notinuse.PrivateTransactionManager{}).String()
		actualType := reflect.TypeOf(ptm).String()

		t.Errorf("got wrong tx manager type. Expected '%s' but got '%s'", expectedType, actualType)
	}
}

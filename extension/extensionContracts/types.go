package extensionContracts

import (
	"github.com/ethereum/go-ethereum/core/state"
)

type AccountWithMetadata struct {
	State state.DumpAccount `json:"state"`
}

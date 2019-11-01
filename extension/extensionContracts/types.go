package extensionContracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/state"
)

var (
	//error is ignored here since it cannot happen (the ABI is generated and thus correct)
	ContractExtensionABI, _ = abi.JSON(strings.NewReader(ContractExtenderABI))
)

type AccountWithMetadata struct {
	State state.DumpAccount `json:"state"`
}

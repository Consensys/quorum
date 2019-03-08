package params

import "github.com/ethereum/go-ethereum/common"

var (
	QuorumPermissionsContract          = common.Address{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32}
	QuorumPrivateKeyManagementContract = common.Address{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 34}
)

const (
	PERMISSIONED_CONFIG = "permissioned-nodes.json"
	BLACKLIST_CONFIG    = "disallowed-nodes.json"
)

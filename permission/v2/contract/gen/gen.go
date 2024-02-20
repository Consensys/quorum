// Quorum
//
// this is to generate go binding for smart contracts used in permissioning
//
// Require:
// 1. solc 0.8.1
// 2. abigen (make all from root)

//go:generate solc --abi --bin -o . --overwrite ../AccountManager.sol
//go:generate solc --abi --bin -o . --overwrite ../NodeManager.sol
//go:generate solc --abi --bin -o . --overwrite ../OrgManager.sol
//go:generate solc --abi --bin -o . --overwrite ../PermissionsImplementation.sol
//go:generate solc --abi --bin -o . --overwrite ../PermissionsInterface.sol
//go:generate solc --abi --bin -o . --overwrite ../PermissionsUpgradable.sol
//go:generate solc --abi --bin -o . --overwrite ../RoleManager.sol
//go:generate solc --abi --bin -o . --overwrite ../VoterManager.sol
//go:generate solc --abi --bin -o . --overwrite ../ContractWhitelistManager.sol
//go:generate solc --abi --bin -o . --overwrite ../openzeppelin-v5/Initializable.sol

//go:generate abigen -pkg bind -abi  ./AccountManager.abi            -bin  ./AccountManager.bin            -type AcctManager               -out ../../bind/accounts.go 					--alias Initialized=InitializedAccs
//go:generate abigen -pkg bind -abi  ./NodeManager.abi               -bin  ./NodeManager.bin               -type NodeManager               -out ../../bind/nodes.go						--alias Initialized=InitializedNodes
//go:generate abigen -pkg bind -abi  ./OrgManager.abi                -bin  ./OrgManager.bin                -type OrgManager                -out ../../bind/org.go						--alias Initialized=InitializedOrg
//go:generate abigen -pkg bind -abi  ./PermissionsImplementation.abi -bin  ./PermissionsImplementation.bin -type PermImpl                  -out ../../bind/permission_impl.go			--alias Initialized=InitializedPermImpl
//go:generate abigen -pkg bind -abi  ./PermissionsInterface.abi      -bin  ./PermissionsInterface.bin      -type PermInterface             -out ../../bind/permission_interface.go		--alias Initialized=InitializedPermIface
//go:generate abigen -pkg bind -abi  ./PermissionsUpgradable.abi     -bin  ./PermissionsUpgradable.bin     -type permUpgr                  -out ../../bind/permission_upgr.go			--alias Initialized=Initialized
//go:generate abigen -pkg bind -abi  ./RoleManager.abi               -bin  ./RoleManager.bin               -type RoleManager               -out ../../bind/roles.go						--alias Initialized=InitializedRoles
//go:generate abigen -pkg bind -abi  ./VoterManager.abi              -bin  ./VoterManager.bin              -type VoterManager              -out ../../bind/voter.go						--alias Initialized=InitializedVoter
//go:generate abigen -pkg bind -abi  ./ContractWhitelistManager.abi  -bin  ./ContractWhitelistManager.bin  -type ContractWhitelistManager  -out ../../bind/whitelist.go					--alias Initialized=InitializedWhitelist

package gen

// Quorum
//
// this is to generate go binding for smart contracts used in permissioning
//
// Require:
// 1. solc 0.5.4
// 2. abigen (make all from root)

//go:generate solc --abi --bin -o . --overwrite ../AccountManager.sol
//go:generate solc --abi --bin -o . --overwrite ../NodeManager.sol
//go:generate solc --abi --bin -o . --overwrite ../OrgManager.sol
//go:generate solc --abi --bin -o . --overwrite ../PermissionsImplementation.sol
//go:generate solc --abi --bin -o . --overwrite ../PermissionsInterface.sol
//go:generate solc --abi --bin -o . --overwrite ../PermissionsUpgradable.sol
//go:generate solc --abi --bin -o . --overwrite ../RoleManager.sol
//go:generate solc --abi --bin -o . --overwrite ../VoterManager.sol

//go:generate abigen -pkg bind -abi  ./AccountManager.abi            -bin  ./AccountManager.bin            -type AcctManager   -out ../../bind/accounts.go
//go:generate abigen -pkg bind -abi  ./NodeManager.abi               -bin  ./NodeManager.bin               -type NodeManager   -out ../../bind/nodes.go
//go:generate abigen -pkg bind -abi  ./OrgManager.abi                -bin  ./OrgManager.bin                -type OrgManager    -out ../../bind/org.go
//go:generate abigen -pkg bind -abi  ./PermissionsImplementation.abi -bin  ./PermissionsImplementation.bin -type PermImpl      -out ../../bind/permission_impl.go
//go:generate abigen -pkg bind -abi  ./PermissionsInterface.abi      -bin  ./PermissionsInterface.bin      -type PermInterface -out ../../bind/permission_interface.go
//go:generate abigen -pkg bind -abi  ./PermissionsUpgradable.abi     -bin  ./PermissionsUpgradable.bin     -type permUpgr      -out ../../bind/permission_upgr.go
//go:generate abigen -pkg bind -abi  ./RoleManager.abi               -bin  ./RoleManager.bin               -type RoleManager   -out ../../bind/roles.go
//go:generate abigen -pkg bind -abi  ./VoterManager.abi              -bin  ./VoterManager.bin              -type VoterManager  -out ../../bind/voter.go

package gen

// Quorum
//
// this is to generate go binding for the validators smart contract
//
// Require:
// 1. solc 0.5.4
// 2. abigen (make all from root)
//go:generate solc --abi --bin -o . --overwrite ./ValidatorSmartContractInterface.sol
//go:generate abigen -pkg contract -abi  ./ValidatorSmartContractInterface.abi            -bin  ./ValidatorSmartContractInterface.bin            -type  ValidatorContractInterface  -out ./validator_contract_interface.go
//go:generate rm ValidatorSmartContractInterface.abi ValidatorSmartContractInterface.bin

package contract

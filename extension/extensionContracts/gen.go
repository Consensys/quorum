package extensionContracts

//go:generate solc --abi --bin -o . contract_extender.sol
//go:generate abigen -pkg extensionContract -abi ./ContractExtender.abi -bin ./ContractExtender.bin -type ContractExtender -out ./contract_extender.go
//go:generate rm ContractExtender.abi ContractExtender.bin

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractWhitelistManagerMetaData contains all meta data concerning the ContractWhitelistManager contract.
var ContractWhitelistManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_contractKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"ContractWhitelistModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_contractKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"ContractWhitelistRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cIndex\",\"type\":\"uint256\"}],\"name\":\"getContractWhitelistDetailsFromIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfWhitelistedContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"revokeWhitelistByAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"revokeWhitelistByKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50611a5f8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630199abce1461006457806307020b6814610080578063b20f4fa51461009c578063b5a93e26146100ce578063c4d66de8146100ec578063cdee5b5e14610108575b5f80fd5b61007e60048036038101906100799190610f50565b610124565b005b61009a60048036038101906100959190610fad565b610532565b005b6100b660048036038101906100b1919061100b565b610779565b6040516100c5939291906110de565b60405180910390f35b6100d66108d6565b6040516100e3919061111a565b60405180910390f35b61010660048036038101906101019190610fad565b6108eb565b005b610122600480360381019061011d9190611133565b610b19565b005b61012c610d66565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa158015610196573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101ba9190611192565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610227576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021e90611207565b60405180910390fd5b5f610230610d66565b90505f816003018585604051610247929190611261565b90815260200160405180910390205414610343575f6102668585610dcc565b90508282600101828154811061027f5761027e611279565b5b905f5260205f2090600302015f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084848360010183815481106102e2576102e1611279565b5b905f5260205f20906003020160010191826102fe9291906114d7565b50600182600101828154811061031757610316611279565b5b905f5260205f2090600302016002015f6101000a81548160ff021916908360ff160217905550506104ee565b806004015f815480929190610357906115d1565b91905055508060040154816003018585604051610375929190611261565b9081526020016040518091039020819055508060040154816002015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060010160405180606001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001600160ff16815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816104ca9190611618565b506040820151816002015f6101000a81548160ff021916908360ff16021790555050505b7f2a0a8be827bc10fe2f9cb3bc17d27e87a5d8b735015ac992aedb6155a0c72d0382858560016040516105249493929190611758565b60405180910390a150505050565b61053a610d66565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105a4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105c89190611192565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062c90611207565b60405180910390fd5b5f61063e610d66565b6002015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036106bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b4906117e0565b60405180910390fd5b5f6106c782610e0b565b905060026106d3610d66565b60010182815481106106e8576106e7611279565b5b905f5260205f2090600302016002015f6101000a81548160ff021916908360ff1602179055507fdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f482610738610d66565b600101838154811061074d5761074c611279565b5b905f5260205f209060030201600101600260405161076d939291906118b8565b60405180910390a15050565b5f60605f80610786610d66565b600101858154811061079b5761079a611279565b5b905f5260205f2090600302016040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546108169061130a565b80601f01602080910402602001604051908101604052809291908181526020018280546108429061130a565b801561088d5780601f106108645761010080835404028352916020019161088d565b820191905f5260205f20905b81548152906001019060200180831161087057829003601f168201915b50505050508152602001600282015f9054906101000a900460ff1660ff1660ff16815250509050805f0151816020015182604001518060ff169050935093509350509193909250565b5f6108df610d66565b60010180549050905090565b5f6108f4610e66565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff1614801561093c5750825b90505f60018367ffffffffffffffff1614801561096f57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b90508115801561097d575080155b156109b4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610a01576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1603610a6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a669061193e565b60405180910390fd5b85610a78610d66565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508315610b11575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610b08919061199f565b60405180910390a15b505050505050565b610b21610d66565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b8b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610baf9190611192565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1390611207565b60405180910390fd5b5f610c25610d66565b6003018383604051610c38929190611261565b90815260200160405180910390205403610c87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7e906117e0565b60405180910390fd5b5f610c928383610dcc565b90506002610c9e610d66565b6001018281548110610cb357610cb2611279565b5b905f5260205f2090600302016002015f6101000a81548160ff021916908360ff1602179055507fdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f4610d02610d66565b6001018281548110610d1757610d16611279565b5b905f5260205f2090600302015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684846002604051610d5994939291906119b8565b60405180910390a1505050565b5f8060ff5f1b1960017fdc0a0fb9b8c3742858130ce0eafb2fa7793d4ff4fec8654c10918f0e0dfd8c765f1c610d9c91906119f6565b604051602001610dac919061111a565b604051602081830303815290604052805190602001201690508091505090565b5f6001610dd7610d66565b6003018484604051610dea929190611261565b908152602001604051809103902054610e0391906119f6565b905092915050565b5f6001610e16610d66565b6002015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610e5f91906119f6565b9050919050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610eb657610eb5610e95565b5b8235905067ffffffffffffffff811115610ed357610ed2610e99565b5b602083019150836001820283011115610eef57610eee610e9d565b5b9250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f1f82610ef6565b9050919050565b610f2f81610f15565b8114610f39575f80fd5b50565b5f81359050610f4a81610f26565b92915050565b5f805f60408486031215610f6757610f66610e8d565b5b5f84013567ffffffffffffffff811115610f8457610f83610e91565b5b610f9086828701610ea1565b93509350506020610fa386828701610f3c565b9150509250925092565b5f60208284031215610fc257610fc1610e8d565b5b5f610fcf84828501610f3c565b91505092915050565b5f819050919050565b610fea81610fd8565b8114610ff4575f80fd5b50565b5f8135905061100581610fe1565b92915050565b5f602082840312156110205761101f610e8d565b5b5f61102d84828501610ff7565b91505092915050565b61103f81610f15565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561107c578082015181840152602081019050611061565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6110a182611045565b6110ab818561104f565b93506110bb81856020860161105f565b6110c481611087565b840191505092915050565b6110d881610fd8565b82525050565b5f6060820190506110f15f830186611036565b81810360208301526111038185611097565b905061111260408301846110cf565b949350505050565b5f60208201905061112d5f8301846110cf565b92915050565b5f806020838503121561114957611148610e8d565b5b5f83013567ffffffffffffffff81111561116657611165610e91565b5b61117285828601610ea1565b92509250509250929050565b5f8151905061118c81610f26565b92915050565b5f602082840312156111a7576111a6610e8d565b5b5f6111b48482850161117e565b91505092915050565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f6111f1600e8361104f565b91506111fc826111bd565b602082019050919050565b5f6020820190508181035f83015261121e816111e5565b9050919050565b5f81905092915050565b828183375f83830152505050565b5f6112488385611225565b935061125583858461122f565b82840190509392505050565b5f61126d82848661123d565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061132157607f821691505b602082108103611334576113336112dd565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261135b565b6113a0868361135b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6113db6113d66113d184610fd8565b6113b8565b610fd8565b9050919050565b5f819050919050565b6113f4836113c1565b611408611400826113e2565b848454611367565b825550505050565b5f90565b61141c611410565b6114278184846113eb565b505050565b5b8181101561144a5761143f5f82611414565b60018101905061142d565b5050565b601f82111561148f576114608161133a565b6114698461134c565b81016020851015611478578190505b61148c6114848561134c565b83018261142c565b50505b505050565b5f82821c905092915050565b5f6114af5f1984600802611494565b1980831691505092915050565b5f6114c783836114a0565b9150826002028217905092915050565b6114e183836112a6565b67ffffffffffffffff8111156114fa576114f96112b0565b5b611504825461130a565b61150f82828561144e565b5f601f83116001811461153c575f841561152a578287013590505b61153485826114bc565b86555061159b565b601f19841661154a8661133a565b5f5b828110156115715784890135825560018201915060208501945060208101905061154c565b8683101561158e578489013561158a601f8916826114a0565b8355505b6001600288020188555050505b50505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6115db82610fd8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361160d5761160c6115a4565b5b600182019050919050565b61162182611045565b67ffffffffffffffff81111561163a576116396112b0565b5b611644825461130a565b61164f82828561144e565b5f60209050601f831160018114611680575f841561166e578287015190505b61167885826114bc565b8655506116df565b601f19841661168e8661133a565b5f5b828110156116b557848901518255600182019150602085019450602081019050611690565b868310156116d257848901516116ce601f8916826114a0565b8355505b6001600288020188555050505b505050505050565b5f6116f2838561104f565b93506116ff83858461122f565b61170883611087565b840190509392505050565b5f819050919050565b5f60ff82169050919050565b5f61174261173d61173884611713565b6113b8565b61171c565b9050919050565b61175281611728565b82525050565b5f60608201905061176b5f830187611036565b818103602083015261177e8185876116e7565b905061178d6040830184611749565b95945050505050565b7f77686974656c69737420646f6573206e6f7420657869737473000000000000005f82015250565b5f6117ca60198361104f565b91506117d582611796565b602082019050919050565b5f6020820190508181035f8301526117f7816117be565b9050919050565b5f815461180a8161130a565b611814818661104f565b9450600182165f811461182e576001811461184457611876565b60ff198316865281151560200286019350611876565b61184d8561133a565b5f5b8381101561186e5781548189015260018201915060208101905061184f565b808801955050505b50505092915050565b5f819050919050565b5f6118a261189d6118988461187f565b6113b8565b61171c565b9050919050565b6118b281611888565b82525050565b5f6060820190506118cb5f830186611036565b81810360208301526118dd81856117fe565b90506118ec60408301846118a9565b949350505050565b7f43616e6e6f742073657420746f20656d707479206164647265737300000000005f82015250565b5f611928601b8361104f565b9150611933826118f4565b602082019050919050565b5f6020820190508181035f8301526119558161191c565b9050919050565b5f67ffffffffffffffff82169050919050565b5f61198961198461197f84611713565b6113b8565b61195c565b9050919050565b6119998161196f565b82525050565b5f6020820190506119b25f830184611990565b92915050565b5f6060820190506119cb5f830187611036565b81810360208301526119de8185876116e7565b90506119ed60408301846118a9565b95945050505050565b5f611a0082610fd8565b9150611a0b83610fd8565b9250828203905081811115611a2357611a226115a4565b5b9291505056fea2646970667358221220b2f27cea3f0b9dbf8de2b964a4713ea7b369efc5b9113ede70a880fe72fe296564736f6c63430008180033",
}

// ContractWhitelistManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractWhitelistManagerMetaData.ABI instead.
var ContractWhitelistManagerABI = ContractWhitelistManagerMetaData.ABI

var ContractWhitelistManagerParsedABI, _ = abi.JSON(strings.NewReader(ContractWhitelistManagerABI))

// ContractWhitelistManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractWhitelistManagerMetaData.Bin instead.
var ContractWhitelistManagerBin = ContractWhitelistManagerMetaData.Bin

// DeployContractWhitelistManager deploys a new Ethereum contract, binding an instance of ContractWhitelistManager to it.
func DeployContractWhitelistManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractWhitelistManager, error) {
	parsed, err := ContractWhitelistManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractWhitelistManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractWhitelistManager{ContractWhitelistManagerCaller: ContractWhitelistManagerCaller{contract: contract}, ContractWhitelistManagerTransactor: ContractWhitelistManagerTransactor{contract: contract}, ContractWhitelistManagerFilterer: ContractWhitelistManagerFilterer{contract: contract}}, nil
}

// ContractWhitelistManager is an auto generated Go binding around an Ethereum contract.
type ContractWhitelistManager struct {
	ContractWhitelistManagerCaller     // Read-only binding to the contract
	ContractWhitelistManagerTransactor // Write-only binding to the contract
	ContractWhitelistManagerFilterer   // Log filterer for contract events
}

// ContractWhitelistManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractWhitelistManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWhitelistManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractWhitelistManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWhitelistManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractWhitelistManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWhitelistManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractWhitelistManagerSession struct {
	Contract     *ContractWhitelistManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractWhitelistManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractWhitelistManagerCallerSession struct {
	Contract *ContractWhitelistManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractWhitelistManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractWhitelistManagerTransactorSession struct {
	Contract     *ContractWhitelistManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractWhitelistManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractWhitelistManagerRaw struct {
	Contract *ContractWhitelistManager // Generic contract binding to access the raw methods on
}

// ContractWhitelistManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractWhitelistManagerCallerRaw struct {
	Contract *ContractWhitelistManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractWhitelistManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractWhitelistManagerTransactorRaw struct {
	Contract *ContractWhitelistManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractWhitelistManager creates a new instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManager(address common.Address, backend bind.ContractBackend) (*ContractWhitelistManager, error) {
	contract, err := bindContractWhitelistManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManager{ContractWhitelistManagerCaller: ContractWhitelistManagerCaller{contract: contract}, ContractWhitelistManagerTransactor: ContractWhitelistManagerTransactor{contract: contract}, ContractWhitelistManagerFilterer: ContractWhitelistManagerFilterer{contract: contract}}, nil
}

// NewContractWhitelistManagerCaller creates a new read-only instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractWhitelistManagerCaller, error) {
	contract, err := bindContractWhitelistManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerCaller{contract: contract}, nil
}

// NewContractWhitelistManagerTransactor creates a new write-only instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractWhitelistManagerTransactor, error) {
	contract, err := bindContractWhitelistManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerTransactor{contract: contract}, nil
}

// NewContractWhitelistManagerFilterer creates a new log filterer instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractWhitelistManagerFilterer, error) {
	contract, err := bindContractWhitelistManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerFilterer{contract: contract}, nil
}

// bindContractWhitelistManager binds a generic wrapper to an already deployed contract.
func bindContractWhitelistManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractWhitelistManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWhitelistManager *ContractWhitelistManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractWhitelistManager.Contract.ContractWhitelistManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWhitelistManager *ContractWhitelistManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.ContractWhitelistManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWhitelistManager *ContractWhitelistManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.ContractWhitelistManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWhitelistManager *ContractWhitelistManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractWhitelistManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.contract.Transact(opts, method, params...)
}

// GetContractWhitelistDetailsFromIndex is a free data retrieval call binding the contract method 0xb20f4fa5.
//
// Solidity: function getContractWhitelistDetailsFromIndex(uint256 _cIndex) view returns(address, string, uint256 status)
func (_ContractWhitelistManager *ContractWhitelistManagerCaller) GetContractWhitelistDetailsFromIndex(opts *bind.CallOpts, _cIndex *big.Int) (common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _ContractWhitelistManager.contract.Call(opts, &out, "getContractWhitelistDetailsFromIndex", _cIndex)

	if err != nil {
		return *new(common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetContractWhitelistDetailsFromIndex is a free data retrieval call binding the contract method 0xb20f4fa5.
//
// Solidity: function getContractWhitelistDetailsFromIndex(uint256 _cIndex) view returns(address, string, uint256 status)
func (_ContractWhitelistManager *ContractWhitelistManagerSession) GetContractWhitelistDetailsFromIndex(_cIndex *big.Int) (common.Address, string, *big.Int, error) {
	return _ContractWhitelistManager.Contract.GetContractWhitelistDetailsFromIndex(&_ContractWhitelistManager.CallOpts, _cIndex)
}

// GetContractWhitelistDetailsFromIndex is a free data retrieval call binding the contract method 0xb20f4fa5.
//
// Solidity: function getContractWhitelistDetailsFromIndex(uint256 _cIndex) view returns(address, string, uint256 status)
func (_ContractWhitelistManager *ContractWhitelistManagerCallerSession) GetContractWhitelistDetailsFromIndex(_cIndex *big.Int) (common.Address, string, *big.Int, error) {
	return _ContractWhitelistManager.Contract.GetContractWhitelistDetailsFromIndex(&_ContractWhitelistManager.CallOpts, _cIndex)
}

// GetNumberOfWhitelistedContracts is a free data retrieval call binding the contract method 0xb5a93e26.
//
// Solidity: function getNumberOfWhitelistedContracts() view returns(uint256)
func (_ContractWhitelistManager *ContractWhitelistManagerCaller) GetNumberOfWhitelistedContracts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractWhitelistManager.contract.Call(opts, &out, "getNumberOfWhitelistedContracts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfWhitelistedContracts is a free data retrieval call binding the contract method 0xb5a93e26.
//
// Solidity: function getNumberOfWhitelistedContracts() view returns(uint256)
func (_ContractWhitelistManager *ContractWhitelistManagerSession) GetNumberOfWhitelistedContracts() (*big.Int, error) {
	return _ContractWhitelistManager.Contract.GetNumberOfWhitelistedContracts(&_ContractWhitelistManager.CallOpts)
}

// GetNumberOfWhitelistedContracts is a free data retrieval call binding the contract method 0xb5a93e26.
//
// Solidity: function getNumberOfWhitelistedContracts() view returns(uint256)
func (_ContractWhitelistManager *ContractWhitelistManagerCallerSession) GetNumberOfWhitelistedContracts() (*big.Int, error) {
	return _ContractWhitelistManager.Contract.GetNumberOfWhitelistedContracts(&_ContractWhitelistManager.CallOpts)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0x0199abce.
//
// Solidity: function addWhitelist(string _key, address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) AddWhitelist(opts *bind.TransactOpts, _key string, _contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "addWhitelist", _key, _contract)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0x0199abce.
//
// Solidity: function addWhitelist(string _key, address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) AddWhitelist(_key string, _contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.AddWhitelist(&_ContractWhitelistManager.TransactOpts, _key, _contract)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0x0199abce.
//
// Solidity: function addWhitelist(string _key, address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) AddWhitelist(_key string, _contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.AddWhitelist(&_ContractWhitelistManager.TransactOpts, _key, _contract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permUpgradable) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) Initialize(opts *bind.TransactOpts, _permUpgradable common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "initialize", _permUpgradable)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permUpgradable) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) Initialize(_permUpgradable common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.Initialize(&_ContractWhitelistManager.TransactOpts, _permUpgradable)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permUpgradable) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) Initialize(_permUpgradable common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.Initialize(&_ContractWhitelistManager.TransactOpts, _permUpgradable)
}

// RevokeWhitelistByAddress is a paid mutator transaction binding the contract method 0x07020b68.
//
// Solidity: function revokeWhitelistByAddress(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) RevokeWhitelistByAddress(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "revokeWhitelistByAddress", _contract)
}

// RevokeWhitelistByAddress is a paid mutator transaction binding the contract method 0x07020b68.
//
// Solidity: function revokeWhitelistByAddress(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) RevokeWhitelistByAddress(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.RevokeWhitelistByAddress(&_ContractWhitelistManager.TransactOpts, _contract)
}

// RevokeWhitelistByAddress is a paid mutator transaction binding the contract method 0x07020b68.
//
// Solidity: function revokeWhitelistByAddress(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) RevokeWhitelistByAddress(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.RevokeWhitelistByAddress(&_ContractWhitelistManager.TransactOpts, _contract)
}

// RevokeWhitelistByKey is a paid mutator transaction binding the contract method 0xcdee5b5e.
//
// Solidity: function revokeWhitelistByKey(string _key) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) RevokeWhitelistByKey(opts *bind.TransactOpts, _key string) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "revokeWhitelistByKey", _key)
}

// RevokeWhitelistByKey is a paid mutator transaction binding the contract method 0xcdee5b5e.
//
// Solidity: function revokeWhitelistByKey(string _key) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) RevokeWhitelistByKey(_key string) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.RevokeWhitelistByKey(&_ContractWhitelistManager.TransactOpts, _key)
}

// RevokeWhitelistByKey is a paid mutator transaction binding the contract method 0xcdee5b5e.
//
// Solidity: function revokeWhitelistByKey(string _key) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) RevokeWhitelistByKey(_key string) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.RevokeWhitelistByKey(&_ContractWhitelistManager.TransactOpts, _key)
}

// ContractWhitelistManagerContractWhitelistModifiedIterator is returned from FilterContractWhitelistModified and is used to iterate over the raw logs and unpacked data for ContractWhitelistModified events raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistModifiedIterator struct {
	Event *ContractWhitelistManagerContractWhitelistModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWhitelistManagerContractWhitelistModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWhitelistManagerContractWhitelistModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWhitelistManagerContractWhitelistModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWhitelistManagerContractWhitelistModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWhitelistManagerContractWhitelistModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWhitelistManagerContractWhitelistModified represents a ContractWhitelistModified event raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistModified struct {
	ContractAddr common.Address
	ContractKey  string
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractWhitelistModified is a free log retrieval operation binding the contract event 0x2a0a8be827bc10fe2f9cb3bc17d27e87a5d8b735015ac992aedb6155a0c72d03.
//
// Solidity: event ContractWhitelistModified(address _contractAddr, string _contractKey, uint8 _status)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) FilterContractWhitelistModified(opts *bind.FilterOpts) (*ContractWhitelistManagerContractWhitelistModifiedIterator, error) {

	logs, sub, err := _ContractWhitelistManager.contract.FilterLogs(opts, "ContractWhitelistModified")
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerContractWhitelistModifiedIterator{contract: _ContractWhitelistManager.contract, event: "ContractWhitelistModified", logs: logs, sub: sub}, nil
}

var ContractWhitelistModifiedTopicHash = "0x2a0a8be827bc10fe2f9cb3bc17d27e87a5d8b735015ac992aedb6155a0c72d03"

// WatchContractWhitelistModified is a free log subscription operation binding the contract event 0x2a0a8be827bc10fe2f9cb3bc17d27e87a5d8b735015ac992aedb6155a0c72d03.
//
// Solidity: event ContractWhitelistModified(address _contractAddr, string _contractKey, uint8 _status)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) WatchContractWhitelistModified(opts *bind.WatchOpts, sink chan<- *ContractWhitelistManagerContractWhitelistModified) (event.Subscription, error) {

	logs, sub, err := _ContractWhitelistManager.contract.WatchLogs(opts, "ContractWhitelistModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWhitelistManagerContractWhitelistModified)
				if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractWhitelistModified is a log parse operation binding the contract event 0x2a0a8be827bc10fe2f9cb3bc17d27e87a5d8b735015ac992aedb6155a0c72d03.
//
// Solidity: event ContractWhitelistModified(address _contractAddr, string _contractKey, uint8 _status)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) ParseContractWhitelistModified(log types.Log) (*ContractWhitelistManagerContractWhitelistModified, error) {
	event := new(ContractWhitelistManagerContractWhitelistModified)
	if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWhitelistManagerContractWhitelistRevokedIterator is returned from FilterContractWhitelistRevoked and is used to iterate over the raw logs and unpacked data for ContractWhitelistRevoked events raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistRevokedIterator struct {
	Event *ContractWhitelistManagerContractWhitelistRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWhitelistManagerContractWhitelistRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWhitelistManagerContractWhitelistRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWhitelistManagerContractWhitelistRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWhitelistManagerContractWhitelistRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWhitelistManagerContractWhitelistRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWhitelistManagerContractWhitelistRevoked represents a ContractWhitelistRevoked event raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistRevoked struct {
	ContractAddr common.Address
	ContractKey  string
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractWhitelistRevoked is a free log retrieval operation binding the contract event 0xdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f4.
//
// Solidity: event ContractWhitelistRevoked(address _contractAddr, string _contractKey, uint8 _status)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) FilterContractWhitelistRevoked(opts *bind.FilterOpts) (*ContractWhitelistManagerContractWhitelistRevokedIterator, error) {

	logs, sub, err := _ContractWhitelistManager.contract.FilterLogs(opts, "ContractWhitelistRevoked")
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerContractWhitelistRevokedIterator{contract: _ContractWhitelistManager.contract, event: "ContractWhitelistRevoked", logs: logs, sub: sub}, nil
}

var ContractWhitelistRevokedTopicHash = "0xdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f4"

// WatchContractWhitelistRevoked is a free log subscription operation binding the contract event 0xdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f4.
//
// Solidity: event ContractWhitelistRevoked(address _contractAddr, string _contractKey, uint8 _status)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) WatchContractWhitelistRevoked(opts *bind.WatchOpts, sink chan<- *ContractWhitelistManagerContractWhitelistRevoked) (event.Subscription, error) {

	logs, sub, err := _ContractWhitelistManager.contract.WatchLogs(opts, "ContractWhitelistRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWhitelistManagerContractWhitelistRevoked)
				if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractWhitelistRevoked is a log parse operation binding the contract event 0xdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f4.
//
// Solidity: event ContractWhitelistRevoked(address _contractAddr, string _contractKey, uint8 _status)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) ParseContractWhitelistRevoked(log types.Log) (*ContractWhitelistManagerContractWhitelistRevoked, error) {
	event := new(ContractWhitelistManagerContractWhitelistRevoked)
	if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWhitelistManagerInitializedWhitelistIterator is returned from FilterInitializedWhitelist and is used to iterate over the raw logs and unpacked data for InitializedWhitelist events raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerInitializedWhitelistIterator struct {
	Event *ContractWhitelistManagerInitializedWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWhitelistManagerInitializedWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWhitelistManagerInitializedWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWhitelistManagerInitializedWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWhitelistManagerInitializedWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWhitelistManagerInitializedWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWhitelistManagerInitializedWhitelist represents a InitializedWhitelist event raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerInitializedWhitelist struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitializedWhitelist is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) FilterInitializedWhitelist(opts *bind.FilterOpts) (*ContractWhitelistManagerInitializedWhitelistIterator, error) {

	logs, sub, err := _ContractWhitelistManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerInitializedWhitelistIterator{contract: _ContractWhitelistManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

var InitializedWhitelistTopicHash = "0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2"

// WatchInitializedWhitelist is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) WatchInitializedWhitelist(opts *bind.WatchOpts, sink chan<- *ContractWhitelistManagerInitializedWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractWhitelistManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWhitelistManagerInitializedWhitelist)
				if err := _ContractWhitelistManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitializedWhitelist is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) ParseInitializedWhitelist(log types.Log) (*ContractWhitelistManagerInitializedWhitelist, error) {
	event := new(ContractWhitelistManagerInitializedWhitelist)
	if err := _ContractWhitelistManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

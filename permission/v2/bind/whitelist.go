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
	Bin: "0x608060405234801561000f575f80fd5b506119398061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630199abce1461006457806307020b6814610080578063b20f4fa51461009c578063b5a93e26146100ce578063c4d66de8146100ec578063cdee5b5e14610108575b5f80fd5b61007e60048036038101906100799190610e2a565b610124565b005b61009a60048036038101906100959190610e87565b610507565b005b6100b660048036038101906100b19190610ee5565b610729565b6040516100c593929190610fb8565b60405180910390f35b6100d661085f565b6040516100e39190610ff4565b60405180910390f35b61010660048036038101906101019190610e87565b61086b565b005b610122600480360381019061011d919061100d565b610a90565b005b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa15801561018c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101b0919061106c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461021d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610214906110e1565b60405180910390fd5b5f6003848460405161023092919061113b565b90815260200160405180910390205414610325575f61024f8484610cb8565b9050816001828154811061026657610265611153565b5b905f5260205f2090600302015f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508383600183815481106102c7576102c6611153565b5b905f5260205f20906003020160010191826102e39291906113b1565b5060018082815481106102f9576102f8611153565b5b905f5260205f2090600302016002015f6101000a81548160ff021916908360ff160217905550506104c4565b60045f815480929190610337906114ab565b91905055506004546003848460405161035192919061113b565b90815260200160405180910390208190555060045460025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600160405180606001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001600160ff16815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816104a091906114f2565b506040820151816002015f6101000a81548160ff021916908360ff16021790555050505b7f2a0a8be827bc10fe2f9cb3bc17d27e87a5d8b735015ac992aedb6155a0c72d0381848460016040516104fa9493929190611632565b60405180910390a1505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa15801561056f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610593919061106c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610600576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f7906110e1565b60405180910390fd5b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20540361067f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610676906116ba565b60405180910390fd5b5f61068982610cee565b90506002600182815481106106a1576106a0611153565b5b905f5260205f2090600302016002015f6101000a81548160ff021916908360ff1602179055507fdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f482600183815481106106fd576106fc611153565b5b905f5260205f209060030201600101600260405161071d93929190611792565b60405180910390a15050565b5f60605f6001848154811061074157610740611153565b5b905f5260205f2090600302015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001858154811061078357610782611153565b5b905f5260205f209060030201600101600186815481106107a6576107a5611153565b5b905f5260205f2090600302016002015f9054906101000a900460ff168180546107ce906111e4565b80601f01602080910402602001604051908101604052809291908181526020018280546107fa906111e4565b80156108455780601f1061081c57610100808354040283529160200191610845565b820191905f5260205f20905b81548152906001019060200180831161082857829003601f168201915b505050505091508060ff1690509250925092509193909250565b5f600180549050905090565b5f610874610d40565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156108bc5750825b90505f60018367ffffffffffffffff161480156108ef57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156108fd575080155b15610934576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610981576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16036109ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e690611818565b60405180910390fd5b855f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508315610a88575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610a7f9190611879565b60405180910390a15b505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa158015610af8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b1c919061106c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b80906110e1565b60405180910390fd5b5f60038383604051610b9c92919061113b565b90815260200160405180910390205403610beb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be2906116ba565b60405180910390fd5b5f610bf68383610cb8565b9050600260018281548110610c0e57610c0d611153565b5b905f5260205f2090600302016002015f6101000a81548160ff021916908360ff1602179055507fdf1fbdc3e0d6d8354b33b1edf0e98496e372aebab012e761210a67b323b9f4f460018281548110610c6957610c68611153565b5b905f5260205f2090600302015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684846002604051610cab9493929190611892565b60405180910390a1505050565b5f600160038484604051610ccd92919061113b565b908152602001604051809103902054610ce691906118d0565b905092915050565b5f600160025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610d3991906118d0565b9050919050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610d9057610d8f610d6f565b5b8235905067ffffffffffffffff811115610dad57610dac610d73565b5b602083019150836001820283011115610dc957610dc8610d77565b5b9250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610df982610dd0565b9050919050565b610e0981610def565b8114610e13575f80fd5b50565b5f81359050610e2481610e00565b92915050565b5f805f60408486031215610e4157610e40610d67565b5b5f84013567ffffffffffffffff811115610e5e57610e5d610d6b565b5b610e6a86828701610d7b565b93509350506020610e7d86828701610e16565b9150509250925092565b5f60208284031215610e9c57610e9b610d67565b5b5f610ea984828501610e16565b91505092915050565b5f819050919050565b610ec481610eb2565b8114610ece575f80fd5b50565b5f81359050610edf81610ebb565b92915050565b5f60208284031215610efa57610ef9610d67565b5b5f610f0784828501610ed1565b91505092915050565b610f1981610def565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610f56578082015181840152602081019050610f3b565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610f7b82610f1f565b610f858185610f29565b9350610f95818560208601610f39565b610f9e81610f61565b840191505092915050565b610fb281610eb2565b82525050565b5f606082019050610fcb5f830186610f10565b8181036020830152610fdd8185610f71565b9050610fec6040830184610fa9565b949350505050565b5f6020820190506110075f830184610fa9565b92915050565b5f806020838503121561102357611022610d67565b5b5f83013567ffffffffffffffff8111156110405761103f610d6b565b5b61104c85828601610d7b565b92509250509250929050565b5f8151905061106681610e00565b92915050565b5f6020828403121561108157611080610d67565b5b5f61108e84828501611058565b91505092915050565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f6110cb600e83610f29565b91506110d682611097565b602082019050919050565b5f6020820190508181035f8301526110f8816110bf565b9050919050565b5f81905092915050565b828183375f83830152505050565b5f61112283856110ff565b935061112f838584611109565b82840190509392505050565b5f611147828486611117565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111fb57607f821691505b60208210810361120e5761120d6111b7565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026112707fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611235565b61127a8683611235565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6112b56112b06112ab84610eb2565b611292565b610eb2565b9050919050565b5f819050919050565b6112ce8361129b565b6112e26112da826112bc565b848454611241565b825550505050565b5f90565b6112f66112ea565b6113018184846112c5565b505050565b5b81811015611324576113195f826112ee565b600181019050611307565b5050565b601f8211156113695761133a81611214565b61134384611226565b81016020851015611352578190505b61136661135e85611226565b830182611306565b50505b505050565b5f82821c905092915050565b5f6113895f198460080261136e565b1980831691505092915050565b5f6113a1838361137a565b9150826002028217905092915050565b6113bb8383611180565b67ffffffffffffffff8111156113d4576113d361118a565b5b6113de82546111e4565b6113e9828285611328565b5f601f831160018114611416575f8415611404578287013590505b61140e8582611396565b865550611475565b601f19841661142486611214565b5f5b8281101561144b57848901358255600182019150602085019450602081019050611426565b868310156114685784890135611464601f89168261137a565b8355505b6001600288020188555050505b50505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6114b582610eb2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114e7576114e661147e565b5b600182019050919050565b6114fb82610f1f565b67ffffffffffffffff8111156115145761151361118a565b5b61151e82546111e4565b611529828285611328565b5f60209050601f83116001811461155a575f8415611548578287015190505b6115528582611396565b8655506115b9565b601f19841661156886611214565b5f5b8281101561158f5784890151825560018201915060208501945060208101905061156a565b868310156115ac57848901516115a8601f89168261137a565b8355505b6001600288020188555050505b505050505050565b5f6115cc8385610f29565b93506115d9838584611109565b6115e283610f61565b840190509392505050565b5f819050919050565b5f60ff82169050919050565b5f61161c611617611612846115ed565b611292565b6115f6565b9050919050565b61162c81611602565b82525050565b5f6060820190506116455f830187610f10565b81810360208301526116588185876115c1565b90506116676040830184611623565b95945050505050565b7f77686974656c69737420646f6573206e6f7420657869737473000000000000005f82015250565b5f6116a4601983610f29565b91506116af82611670565b602082019050919050565b5f6020820190508181035f8301526116d181611698565b9050919050565b5f81546116e4816111e4565b6116ee8186610f29565b9450600182165f8114611708576001811461171e57611750565b60ff198316865281151560200286019350611750565b61172785611214565b5f5b8381101561174857815481890152600182019150602081019050611729565b808801955050505b50505092915050565b5f819050919050565b5f61177c61177761177284611759565b611292565b6115f6565b9050919050565b61178c81611762565b82525050565b5f6060820190506117a55f830186610f10565b81810360208301526117b781856116d8565b90506117c66040830184611783565b949350505050565b7f43616e6e6f742073657420746f20656d707479206164647265737300000000005f82015250565b5f611802601b83610f29565b915061180d826117ce565b602082019050919050565b5f6020820190508181035f83015261182f816117f6565b9050919050565b5f67ffffffffffffffff82169050919050565b5f61186361185e611859846115ed565b611292565b611836565b9050919050565b61187381611849565b82525050565b5f60208201905061188c5f83018461186a565b92915050565b5f6060820190506118a55f830187610f10565b81810360208301526118b88185876115c1565b90506118c76040830184611783565b95945050505050565b5f6118da82610eb2565b91506118e583610eb2565b92508282039050818111156118fd576118fc61147e565b5b9291505056fea26469706673582212202cd9c2ccf0918c37b4cc6dccdf5b477f1b83a0485460bb501391128789ca71a364736f6c63430008180033",
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

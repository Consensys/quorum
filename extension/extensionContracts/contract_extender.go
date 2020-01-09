// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extensionContracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractExtenderABI is the input ABI used to generate the binding from.
const ContractExtenderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractToExtend\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkIfExtensionFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNumberOfVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetRecipientPublicKeyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"walletAddressesToVote\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"setUuid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sharedDataHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setSharedStateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updatePartyMembers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteOutcome\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkIfVoted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vote\",\"type\":\"bool\"},{\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"doVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"haveAllNodesVoted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"walletAddresses\",\"type\":\"address[]\"},{\"name\":\"recipientHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toExtend\",\"type\":\"address\"}],\"name\":\"NewContractExtensionContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"outcome\",\"type\":\"bool\"}],\"name\":\"AllNodesHaveVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CanPerformStateShare\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ExtensionFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vote\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"NewVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tesserahash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"StateShared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UpdateMembers\",\"type\":\"event\"}]"

var ContractExtenderParsedABI, _ = abi.JSON(strings.NewReader(ContractExtenderABI))

// ContractExtenderBin is the compiled bytecode used for deploying new contracts.
const ContractExtenderBin = `60806040523480156200001157600080fd5b506040516200141c3803806200141c833981018060405260608110156200003757600080fd5b8151602083018051919392830192916401000000008111156200005957600080fd5b820160208101848111156200006d57600080fd5b81518560208202830111640100000000821117156200008b57600080fd5b50509291906020018051640100000000811115620000a857600080fd5b82016020810184811115620000bc57600080fd5b8151640100000000811182820187101715620000d757600080fd5b505060008054600160a060020a031916331790558051909350620001059250600191506020840190620002a8565b5060028054600160a060020a031916600160a060020a0385161790558151620001369060039060208501906200032d565b506040805160208101918290526000908190526200015791600a91620002a8565b506009805460ff1916600117905560006006819055805b8351811015620001f85760016005600086848151811015156200018d57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff191691151591909117905583513390859083908110620001d057fe5b90602001906020020151600160a060020a03161415620001ef57600191505b6001016200016e565b508015156200025c576003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018054600160a060020a031916339081179091556000908152600560205260409020805460ff191690911790555b60035460045560408051600160a060020a038616815290517f1bb7909ad96bc757f60de4d9ce11daf7b006e8f398ce028dceb10ce7fdca0f689181900360200190a150505050620003da565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002eb57805160ff19168380011785556200031b565b828001600101855582156200031b579182015b828111156200031b578251825591602001919060010190620002fe565b506200032992915062000393565b5090565b82805482825590600052602060002090810192821562000385579160200282015b82811115620003855782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906200034e565b5062000329929150620003b3565b620003b091905b808211156200032957600081556001016200039a565b90565b620003b091905b8082111562000329578054600160a060020a0319168155600101620003ba565b61103280620003ea6000396000f3fe608060405234801561001057600080fd5b5060043610610128576000357c01000000000000000000000000000000000000000000000000000000009004806388f520a0116100bf578063cb2805ec1161008e578063cb2805ec14610397578063d56b28891461039f578063d8bff5a5146103a7578063de5828cb146103cd578063f57077d81461047c57610128565b806388f520a0146102d9578063893971ba146102e1578063ac8b920514610387578063b5da45bb1461038f57610128565b80634688a5e3116100fb5780634688a5e31461018f57806379d41b8f1461020c5780637b35296214610229578063821e93da1461023157610128565b806302d05d3f1461012d57806315e56a6a146101515780631962cb9b146101595780633852772714610175575b600080fd5b610135610484565b60408051600160a060020a039092168252519081900360200190f35b610135610493565b6101616104a2565b604080519115158252519081900360200190f35b61017d6104ac565b60408051918252519081900360200190f35b6101976104b2565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101d15781810151838201526020016101b9565b50505050905090810190601f1680156101fe5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101356004803603602081101561022257600080fd5b503561053f565b610161610567565b6102d76004803603602081101561024757600080fd5b81019060208101813564010000000081111561026257600080fd5b82018360208201111561027457600080fd5b8035906020019184600183028401116401000000008311171561029657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610570945050505050565b005b6101976105fe565b6102d7600480360360208110156102f757600080fd5b81019060208101813564010000000081111561031257600080fd5b82018360208201111561032457600080fd5b8035906020019184600183028401116401000000008311171561034657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610659945050505050565b6102d76109c3565b610161610abe565b610161610ac7565b6102d7610add565b610161600480360360208110156103bd57600080fd5b5035600160a060020a0316610b78565b6102d7600480360360408110156103e357600080fd5b81351515919081019060408101602082013564010000000081111561040757600080fd5b82018360208201111561041957600080fd5b8035906020019184600183028401116401000000008311171561043b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b8d945050505050565b610161610c32565b600054600160a060020a031681565b600254600160a060020a031681565b600c5460ff165b90565b60045481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105375780601f1061050c57610100808354040283529160200191610537565b820191906000526020600020905b81548152906001019060200180831161051a57829003601f168201915b505050505081565b600380548290811061054d57fe5b600091825260209091200154600160a060020a0316905081565b600c5460ff1681565b600c5460ff16156105b55760405160e560020a62461bcd028152600401808060200182810382526025815260200180610fe26025913960400191505060405180910390fd5b600b80546001810180835560009290925282516105f9917f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db901906020850190610efe565b505050565b600a805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105375780601f1061050c57610100808354040283529160200191610537565b600054600160a060020a031633146106a55760405160e560020a62461bcd028152600401808060200182810382526023815260200180610fbf6023913960400191505060405180910390fd5b600c5460ff16156106ea5760405160e560020a62461bcd028152600401808060200182810382526025815260200180610fe26025913960400191505060405180910390fd5b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156107765780601f1061074b57610100808354040283529160200191610776565b820191906000526020600020905b81548152906001019060200180831161075957829003601f168201915b505085519394508593151592506107da915050576040805160e560020a62461bcd02815260206004820152601860248201527f6e657720686173682063616e6e6f7420626520656d7074790000000000000000604482015290519081900360640190fd5b815115610831576040805160e560020a62461bcd02815260206004820152601660248201527f7374617465206861736820616c72656164792073657400000000000000000000604482015290519081900360640190fd5b825161084490600a906020860190610efe565b5060005b600b548110156109ba57600254600b80547f67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e92600160a060020a031691600a918590811061089257fe5b600091825260209182902060408051600160a060020a0387168152606094810185815286546002600019610100600184161502019091160495820186905292909301939083019060808401908690801561092d5780601f106109025761010080835404028352916020019161092d565b820191906000526020600020905b81548152906001019060200180831161091057829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156109a15780601f10610976576101008083540402835291602001916109a1565b820191906000526020600020905b81548152906001019060200180831161098457829003601f168201915b50509550505050505060405180910390a1600101610848565b506105f9610add565b60005b600b54811015610abb57600254600b80547f8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b4492600160a060020a0316919084908110610a0e57fe5b600091825260209182902060408051600160a060020a0386168152938401818152919092018054600260001961010060018416150201909116049284018390529291606083019084908015610aa45780601f10610a7957610100808354040283529160200191610aa4565b820191906000526020600020905b815481529060010190602001808311610a8757829003601f168201915b5050935050505060405180910390a16001016109c6565b50565b60095460ff1681565b3360009081526007602052604090205460ff1690565b600c5460ff1615610b225760405160e560020a62461bcd028152600401808060200182810382526025815260200180610fe26025913960400191505060405180910390fd5b600054600160a060020a03163314610b6e5760405160e560020a62461bcd028152600401808060200182810382526023815260200180610fbf6023913960400191505060405180910390fd5b610b76610c3c565b565b60086020526000908152604090205460ff1681565b600c5460ff1615610bd25760405160e560020a62461bcd028152600401808060200182810382526025815260200180610fe26025913960400191505060405180910390fd5b610bdb82610c74565b8115610bea57610bea81610570565b610bf2610e45565b60408051831515815233602082015281517f225708d30006b0cc86d855ab91047edb5fe9c2e416412f36c18c6e90fe4e461f929181900390910190a15050565b6006546003541490565b600c805460ff191660011790556040517f79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede990600090a1565b600c5460ff1615610cb95760405160e560020a62461bcd028152600401808060200182810382526028815260200180610f976028913960400191505060405180910390fd5b3360009081526005602052604090205460ff161515610d22576040805160e560020a62461bcd02815260206004820152601360248201527f6e6f7420616c6c6f77656420746f20766f746500000000000000000000000000604482015290519081900360640190fd5b3360009081526007602052604090205460ff1615610d8a576040805160e560020a62461bcd02815260206004820152600d60248201527f616c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b60095460ff161515610de6576040805160e560020a62461bcd02815260206004820152601760248201527f766f74696e6720616c7265616479206465636c696e6564000000000000000000604482015290519081900360640190fd5b3360009081526007602090815260408083208054600160ff19918216811790925560089093529220805490911683151517905560068054909101905560095460ff168015610e315750805b6009805460ff191691151591909117905550565b60095460ff161515610e9257604080516000815290517fc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a29049181900360200190a1610e8d610c3c565b610b76565b610e9a610c32565b15610b7657604080516001815290517fc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a29049181900360200190a16040517ffd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d2890600090a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f3f57805160ff1916838001178555610f6c565b82800160010185558215610f6c579182015b82811115610f6c578251825591602001919060010190610f51565b50610f78929150610f7c565b5090565b6104a991905b80821115610f785760008155600101610f8256fe657874656e73696f6e2070726f6365737320636f6d706c657465642e2063616e6e6f7420766f74656f6e6c79206c6561646572206d617920706572666f726d207468697320616374696f6e657874656e73696f6e20686173206265656e206d61726b65642061732066696e6973686564a165627a7a7230582023105643f95ec239fb859a1df870740273a1086185e144297de775eab4b70e7a0029`

// DeployContractExtender deploys a new Ethereum contract, binding an instance of ContractExtender to it.
func DeployContractExtender(auth *bind.TransactOpts, backend bind.ContractBackend, contractAddress common.Address, walletAddresses []common.Address, recipientHash string) (common.Address, *types.Transaction, *ContractExtender, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractExtenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractExtenderBin), backend, contractAddress, walletAddresses, recipientHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractExtender{ContractExtenderCaller: ContractExtenderCaller{contract: contract}, ContractExtenderTransactor: ContractExtenderTransactor{contract: contract}, ContractExtenderFilterer: ContractExtenderFilterer{contract: contract}}, nil
}

// ContractExtender is an auto generated Go binding around an Ethereum contract.
type ContractExtender struct {
	ContractExtenderCaller     // Read-only binding to the contract
	ContractExtenderTransactor // Write-only binding to the contract
	ContractExtenderFilterer   // Log filterer for contract events
}

// ContractExtenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractExtenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractExtenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractExtenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractExtenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractExtenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractExtenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractExtenderSession struct {
	Contract     *ContractExtender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractExtenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractExtenderCallerSession struct {
	Contract *ContractExtenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractExtenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractExtenderTransactorSession struct {
	Contract     *ContractExtenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractExtenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractExtenderRaw struct {
	Contract *ContractExtender // Generic contract binding to access the raw methods on
}

// ContractExtenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractExtenderCallerRaw struct {
	Contract *ContractExtenderCaller // Generic read-only contract binding to access the raw methods on
}

// ContractExtenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractExtenderTransactorRaw struct {
	Contract *ContractExtenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractExtender creates a new instance of ContractExtender, bound to a specific deployed contract.
func NewContractExtender(address common.Address, backend bind.ContractBackend) (*ContractExtender, error) {
	contract, err := bindContractExtender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractExtender{ContractExtenderCaller: ContractExtenderCaller{contract: contract}, ContractExtenderTransactor: ContractExtenderTransactor{contract: contract}, ContractExtenderFilterer: ContractExtenderFilterer{contract: contract}}, nil
}

// NewContractExtenderCaller creates a new read-only instance of ContractExtender, bound to a specific deployed contract.
func NewContractExtenderCaller(address common.Address, caller bind.ContractCaller) (*ContractExtenderCaller, error) {
	contract, err := bindContractExtender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractExtenderCaller{contract: contract}, nil
}

// NewContractExtenderTransactor creates a new write-only instance of ContractExtender, bound to a specific deployed contract.
func NewContractExtenderTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractExtenderTransactor, error) {
	contract, err := bindContractExtender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractExtenderTransactor{contract: contract}, nil
}

// NewContractExtenderFilterer creates a new log filterer instance of ContractExtender, bound to a specific deployed contract.
func NewContractExtenderFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractExtenderFilterer, error) {
	contract, err := bindContractExtender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractExtenderFilterer{contract: contract}, nil
}

// bindContractExtender binds a generic wrapper to an already deployed contract.
func bindContractExtender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractExtenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractExtender *ContractExtenderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractExtender.Contract.ContractExtenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractExtender *ContractExtenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractExtender.Contract.ContractExtenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractExtender *ContractExtenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractExtender.Contract.ContractExtenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractExtender *ContractExtenderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractExtender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractExtender *ContractExtenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractExtender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractExtender *ContractExtenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractExtender.Contract.contract.Transact(opts, method, params...)
}

// CheckIfExtensionFinished is a free data retrieval call binding the contract method 0x1962cb9b.
//
// Solidity: function checkIfExtensionFinished() constant returns(bool)
func (_ContractExtender *ContractExtenderCaller) CheckIfExtensionFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "checkIfExtensionFinished")
	return *ret0, err
}

// CheckIfExtensionFinished is a free data retrieval call binding the contract method 0x1962cb9b.
//
// Solidity: function checkIfExtensionFinished() constant returns(bool)
func (_ContractExtender *ContractExtenderSession) CheckIfExtensionFinished() (bool, error) {
	return _ContractExtender.Contract.CheckIfExtensionFinished(&_ContractExtender.CallOpts)
}

// CheckIfExtensionFinished is a free data retrieval call binding the contract method 0x1962cb9b.
//
// Solidity: function checkIfExtensionFinished() constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) CheckIfExtensionFinished() (bool, error) {
	return _ContractExtender.Contract.CheckIfExtensionFinished(&_ContractExtender.CallOpts)
}

// CheckIfVoted is a free data retrieval call binding the contract method 0xcb2805ec.
//
// Solidity: function checkIfVoted() constant returns(bool)
func (_ContractExtender *ContractExtenderCaller) CheckIfVoted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "checkIfVoted")
	return *ret0, err
}

// CheckIfVoted is a free data retrieval call binding the contract method 0xcb2805ec.
//
// Solidity: function checkIfVoted() constant returns(bool)
func (_ContractExtender *ContractExtenderSession) CheckIfVoted() (bool, error) {
	return _ContractExtender.Contract.CheckIfVoted(&_ContractExtender.CallOpts)
}

// CheckIfVoted is a free data retrieval call binding the contract method 0xcb2805ec.
//
// Solidity: function checkIfVoted() constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) CheckIfVoted() (bool, error) {
	return _ContractExtender.Contract.CheckIfVoted(&_ContractExtender.CallOpts)
}

// ContractToExtend is a free data retrieval call binding the contract method 0x15e56a6a.
//
// Solidity: function contractToExtend() constant returns(address)
func (_ContractExtender *ContractExtenderCaller) ContractToExtend(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "contractToExtend")
	return *ret0, err
}

// ContractToExtend is a free data retrieval call binding the contract method 0x15e56a6a.
//
// Solidity: function contractToExtend() constant returns(address)
func (_ContractExtender *ContractExtenderSession) ContractToExtend() (common.Address, error) {
	return _ContractExtender.Contract.ContractToExtend(&_ContractExtender.CallOpts)
}

// ContractToExtend is a free data retrieval call binding the contract method 0x15e56a6a.
//
// Solidity: function contractToExtend() constant returns(address)
func (_ContractExtender *ContractExtenderCallerSession) ContractToExtend() (common.Address, error) {
	return _ContractExtender.Contract.ContractToExtend(&_ContractExtender.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_ContractExtender *ContractExtenderCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_ContractExtender *ContractExtenderSession) Creator() (common.Address, error) {
	return _ContractExtender.Contract.Creator(&_ContractExtender.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_ContractExtender *ContractExtenderCallerSession) Creator() (common.Address, error) {
	return _ContractExtender.Contract.Creator(&_ContractExtender.CallOpts)
}

// HaveAllNodesVoted is a free data retrieval call binding the contract method 0xf57077d8.
//
// Solidity: function haveAllNodesVoted() constant returns(bool)
func (_ContractExtender *ContractExtenderCaller) HaveAllNodesVoted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "haveAllNodesVoted")
	return *ret0, err
}

// HaveAllNodesVoted is a free data retrieval call binding the contract method 0xf57077d8.
//
// Solidity: function haveAllNodesVoted() constant returns(bool)
func (_ContractExtender *ContractExtenderSession) HaveAllNodesVoted() (bool, error) {
	return _ContractExtender.Contract.HaveAllNodesVoted(&_ContractExtender.CallOpts)
}

// HaveAllNodesVoted is a free data retrieval call binding the contract method 0xf57077d8.
//
// Solidity: function haveAllNodesVoted() constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) HaveAllNodesVoted() (bool, error) {
	return _ContractExtender.Contract.HaveAllNodesVoted(&_ContractExtender.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ContractExtender *ContractExtenderCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ContractExtender *ContractExtenderSession) IsFinished() (bool, error) {
	return _ContractExtender.Contract.IsFinished(&_ContractExtender.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) IsFinished() (bool, error) {
	return _ContractExtender.Contract.IsFinished(&_ContractExtender.CallOpts)
}

// SharedDataHash is a free data retrieval call binding the contract method 0x88f520a0.
//
// Solidity: function sharedDataHash() constant returns(string)
func (_ContractExtender *ContractExtenderCaller) SharedDataHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "sharedDataHash")
	return *ret0, err
}

// SharedDataHash is a free data retrieval call binding the contract method 0x88f520a0.
//
// Solidity: function sharedDataHash() constant returns(string)
func (_ContractExtender *ContractExtenderSession) SharedDataHash() (string, error) {
	return _ContractExtender.Contract.SharedDataHash(&_ContractExtender.CallOpts)
}

// SharedDataHash is a free data retrieval call binding the contract method 0x88f520a0.
//
// Solidity: function sharedDataHash() constant returns(string)
func (_ContractExtender *ContractExtenderCallerSession) SharedDataHash() (string, error) {
	return _ContractExtender.Contract.SharedDataHash(&_ContractExtender.CallOpts)
}

// TargetRecipientPublicKeyHash is a free data retrieval call binding the contract method 0x4688a5e3.
//
// Solidity: function targetRecipientPublicKeyHash() constant returns(string)
func (_ContractExtender *ContractExtenderCaller) TargetRecipientPublicKeyHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "targetRecipientPublicKeyHash")
	return *ret0, err
}

// TargetRecipientPublicKeyHash is a free data retrieval call binding the contract method 0x4688a5e3.
//
// Solidity: function targetRecipientPublicKeyHash() constant returns(string)
func (_ContractExtender *ContractExtenderSession) TargetRecipientPublicKeyHash() (string, error) {
	return _ContractExtender.Contract.TargetRecipientPublicKeyHash(&_ContractExtender.CallOpts)
}

// TargetRecipientPublicKeyHash is a free data retrieval call binding the contract method 0x4688a5e3.
//
// Solidity: function targetRecipientPublicKeyHash() constant returns(string)
func (_ContractExtender *ContractExtenderCallerSession) TargetRecipientPublicKeyHash() (string, error) {
	return _ContractExtender.Contract.TargetRecipientPublicKeyHash(&_ContractExtender.CallOpts)
}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x38527727.
//
// Solidity: function totalNumberOfVoters() constant returns(uint256)
func (_ContractExtender *ContractExtenderCaller) TotalNumberOfVoters(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "totalNumberOfVoters")
	return *ret0, err
}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x38527727.
//
// Solidity: function totalNumberOfVoters() constant returns(uint256)
func (_ContractExtender *ContractExtenderSession) TotalNumberOfVoters() (*big.Int, error) {
	return _ContractExtender.Contract.TotalNumberOfVoters(&_ContractExtender.CallOpts)
}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x38527727.
//
// Solidity: function totalNumberOfVoters() constant returns(uint256)
func (_ContractExtender *ContractExtenderCallerSession) TotalNumberOfVoters() (*big.Int, error) {
	return _ContractExtender.Contract.TotalNumberOfVoters(&_ContractExtender.CallOpts)
}

// VoteOutcome is a free data retrieval call binding the contract method 0xb5da45bb.
//
// Solidity: function voteOutcome() constant returns(bool)
func (_ContractExtender *ContractExtenderCaller) VoteOutcome(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "voteOutcome")
	return *ret0, err
}

// VoteOutcome is a free data retrieval call binding the contract method 0xb5da45bb.
//
// Solidity: function voteOutcome() constant returns(bool)
func (_ContractExtender *ContractExtenderSession) VoteOutcome() (bool, error) {
	return _ContractExtender.Contract.VoteOutcome(&_ContractExtender.CallOpts)
}

// VoteOutcome is a free data retrieval call binding the contract method 0xb5da45bb.
//
// Solidity: function voteOutcome() constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) VoteOutcome() (bool, error) {
	return _ContractExtender.Contract.VoteOutcome(&_ContractExtender.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes( address) constant returns(bool)
func (_ContractExtender *ContractExtenderCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "votes", arg0)
	return *ret0, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes( address) constant returns(bool)
func (_ContractExtender *ContractExtenderSession) Votes(arg0 common.Address) (bool, error) {
	return _ContractExtender.Contract.Votes(&_ContractExtender.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes( address) constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) Votes(arg0 common.Address) (bool, error) {
	return _ContractExtender.Contract.Votes(&_ContractExtender.CallOpts, arg0)
}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote( uint256) constant returns(address)
func (_ContractExtender *ContractExtenderCaller) WalletAddressesToVote(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "walletAddressesToVote", arg0)
	return *ret0, err
}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote( uint256) constant returns(address)
func (_ContractExtender *ContractExtenderSession) WalletAddressesToVote(arg0 *big.Int) (common.Address, error) {
	return _ContractExtender.Contract.WalletAddressesToVote(&_ContractExtender.CallOpts, arg0)
}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote( uint256) constant returns(address)
func (_ContractExtender *ContractExtenderCallerSession) WalletAddressesToVote(arg0 *big.Int) (common.Address, error) {
	return _ContractExtender.Contract.WalletAddressesToVote(&_ContractExtender.CallOpts, arg0)
}

// DoVote is a paid mutator transaction binding the contract method 0xde5828cb.
//
// Solidity: function doVote(vote bool, nextuuid string) returns()
func (_ContractExtender *ContractExtenderTransactor) DoVote(opts *bind.TransactOpts, vote bool, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "doVote", vote, nextuuid)
}

// DoVote is a paid mutator transaction binding the contract method 0xde5828cb.
//
// Solidity: function doVote(vote bool, nextuuid string) returns()
func (_ContractExtender *ContractExtenderSession) DoVote(vote bool, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.DoVote(&_ContractExtender.TransactOpts, vote, nextuuid)
}

// DoVote is a paid mutator transaction binding the contract method 0xde5828cb.
//
// Solidity: function doVote(vote bool, nextuuid string) returns()
func (_ContractExtender *ContractExtenderTransactorSession) DoVote(vote bool, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.DoVote(&_ContractExtender.TransactOpts, vote, nextuuid)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ContractExtender *ContractExtenderTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ContractExtender *ContractExtenderSession) Finish() (*types.Transaction, error) {
	return _ContractExtender.Contract.Finish(&_ContractExtender.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ContractExtender *ContractExtenderTransactorSession) Finish() (*types.Transaction, error) {
	return _ContractExtender.Contract.Finish(&_ContractExtender.TransactOpts)
}

// SetSharedStateHash is a paid mutator transaction binding the contract method 0x893971ba.
//
// Solidity: function setSharedStateHash(hash string) returns()
func (_ContractExtender *ContractExtenderTransactor) SetSharedStateHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "setSharedStateHash", hash)
}

// SetSharedStateHash is a paid mutator transaction binding the contract method 0x893971ba.
//
// Solidity: function setSharedStateHash(hash string) returns()
func (_ContractExtender *ContractExtenderSession) SetSharedStateHash(hash string) (*types.Transaction, error) {
	return _ContractExtender.Contract.SetSharedStateHash(&_ContractExtender.TransactOpts, hash)
}

// SetSharedStateHash is a paid mutator transaction binding the contract method 0x893971ba.
//
// Solidity: function setSharedStateHash(hash string) returns()
func (_ContractExtender *ContractExtenderTransactorSession) SetSharedStateHash(hash string) (*types.Transaction, error) {
	return _ContractExtender.Contract.SetSharedStateHash(&_ContractExtender.TransactOpts, hash)
}

// SetUuid is a paid mutator transaction binding the contract method 0x821e93da.
//
// Solidity: function setUuid(nextuuid string) returns()
func (_ContractExtender *ContractExtenderTransactor) SetUuid(opts *bind.TransactOpts, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "setUuid", nextuuid)
}

// SetUuid is a paid mutator transaction binding the contract method 0x821e93da.
//
// Solidity: function setUuid(nextuuid string) returns()
func (_ContractExtender *ContractExtenderSession) SetUuid(nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.SetUuid(&_ContractExtender.TransactOpts, nextuuid)
}

// SetUuid is a paid mutator transaction binding the contract method 0x821e93da.
//
// Solidity: function setUuid(nextuuid string) returns()
func (_ContractExtender *ContractExtenderTransactorSession) SetUuid(nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.SetUuid(&_ContractExtender.TransactOpts, nextuuid)
}

// UpdatePartyMembers is a paid mutator transaction binding the contract method 0xac8b9205.
//
// Solidity: function updatePartyMembers() returns()
func (_ContractExtender *ContractExtenderTransactor) UpdatePartyMembers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "updatePartyMembers")
}

// UpdatePartyMembers is a paid mutator transaction binding the contract method 0xac8b9205.
//
// Solidity: function updatePartyMembers() returns()
func (_ContractExtender *ContractExtenderSession) UpdatePartyMembers() (*types.Transaction, error) {
	return _ContractExtender.Contract.UpdatePartyMembers(&_ContractExtender.TransactOpts)
}

// UpdatePartyMembers is a paid mutator transaction binding the contract method 0xac8b9205.
//
// Solidity: function updatePartyMembers() returns()
func (_ContractExtender *ContractExtenderTransactorSession) UpdatePartyMembers() (*types.Transaction, error) {
	return _ContractExtender.Contract.UpdatePartyMembers(&_ContractExtender.TransactOpts)
}

// ContractExtenderAllNodesHaveVotedIterator is returned from FilterAllNodesHaveVoted and is used to iterate over the raw logs and unpacked data for AllNodesHaveVoted events raised by the ContractExtender contract.
type ContractExtenderAllNodesHaveVotedIterator struct {
	Event *ContractExtenderAllNodesHaveVoted // Event containing the contract specifics and raw log

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
func (it *ContractExtenderAllNodesHaveVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderAllNodesHaveVoted)
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
		it.Event = new(ContractExtenderAllNodesHaveVoted)
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
func (it *ContractExtenderAllNodesHaveVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderAllNodesHaveVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderAllNodesHaveVoted represents a AllNodesHaveVoted event raised by the ContractExtender contract.
type ContractExtenderAllNodesHaveVoted struct {
	Outcome bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAllNodesHaveVoted is a free log retrieval operation binding the contract event 0xc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a2904.
//
// Solidity: e AllNodesHaveVoted(outcome bool)
func (_ContractExtender *ContractExtenderFilterer) FilterAllNodesHaveVoted(opts *bind.FilterOpts) (*ContractExtenderAllNodesHaveVotedIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "AllNodesHaveVoted")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderAllNodesHaveVotedIterator{contract: _ContractExtender.contract, event: "AllNodesHaveVoted", logs: logs, sub: sub}, nil
}

var AllNodesHaveVotedTopicHash = "0xc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a2904"

// WatchAllNodesHaveVoted is a free log subscription operation binding the contract event 0xc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a2904.
//
// Solidity: e AllNodesHaveVoted(outcome bool)
func (_ContractExtender *ContractExtenderFilterer) WatchAllNodesHaveVoted(opts *bind.WatchOpts, sink chan<- *ContractExtenderAllNodesHaveVoted) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "AllNodesHaveVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderAllNodesHaveVoted)
				if err := _ContractExtender.contract.UnpackLog(event, "AllNodesHaveVoted", log); err != nil {
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

// ContractExtenderCanPerformStateShareIterator is returned from FilterCanPerformStateShare and is used to iterate over the raw logs and unpacked data for CanPerformStateShare events raised by the ContractExtender contract.
type ContractExtenderCanPerformStateShareIterator struct {
	Event *ContractExtenderCanPerformStateShare // Event containing the contract specifics and raw log

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
func (it *ContractExtenderCanPerformStateShareIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderCanPerformStateShare)
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
		it.Event = new(ContractExtenderCanPerformStateShare)
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
func (it *ContractExtenderCanPerformStateShareIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderCanPerformStateShareIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderCanPerformStateShare represents a CanPerformStateShare event raised by the ContractExtender contract.
type ContractExtenderCanPerformStateShare struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCanPerformStateShare is a free log retrieval operation binding the contract event 0xfd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d28.
//
// Solidity: e CanPerformStateShare()
func (_ContractExtender *ContractExtenderFilterer) FilterCanPerformStateShare(opts *bind.FilterOpts) (*ContractExtenderCanPerformStateShareIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "CanPerformStateShare")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderCanPerformStateShareIterator{contract: _ContractExtender.contract, event: "CanPerformStateShare", logs: logs, sub: sub}, nil
}

var CanPerformStateShareTopicHash = "0xfd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d28"

// WatchCanPerformStateShare is a free log subscription operation binding the contract event 0xfd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d28.
//
// Solidity: e CanPerformStateShare()
func (_ContractExtender *ContractExtenderFilterer) WatchCanPerformStateShare(opts *bind.WatchOpts, sink chan<- *ContractExtenderCanPerformStateShare) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "CanPerformStateShare")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderCanPerformStateShare)
				if err := _ContractExtender.contract.UnpackLog(event, "CanPerformStateShare", log); err != nil {
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

// ContractExtenderExtensionFinishedIterator is returned from FilterExtensionFinished and is used to iterate over the raw logs and unpacked data for ExtensionFinished events raised by the ContractExtender contract.
type ContractExtenderExtensionFinishedIterator struct {
	Event *ContractExtenderExtensionFinished // Event containing the contract specifics and raw log

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
func (it *ContractExtenderExtensionFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderExtensionFinished)
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
		it.Event = new(ContractExtenderExtensionFinished)
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
func (it *ContractExtenderExtensionFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderExtensionFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderExtensionFinished represents a ExtensionFinished event raised by the ContractExtender contract.
type ContractExtenderExtensionFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExtensionFinished is a free log retrieval operation binding the contract event 0x79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede9.
//
// Solidity: e ExtensionFinished()
func (_ContractExtender *ContractExtenderFilterer) FilterExtensionFinished(opts *bind.FilterOpts) (*ContractExtenderExtensionFinishedIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "ExtensionFinished")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderExtensionFinishedIterator{contract: _ContractExtender.contract, event: "ExtensionFinished", logs: logs, sub: sub}, nil
}

var ExtensionFinishedTopicHash = "0x79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede9"

// WatchExtensionFinished is a free log subscription operation binding the contract event 0x79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede9.
//
// Solidity: e ExtensionFinished()
func (_ContractExtender *ContractExtenderFilterer) WatchExtensionFinished(opts *bind.WatchOpts, sink chan<- *ContractExtenderExtensionFinished) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "ExtensionFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderExtensionFinished)
				if err := _ContractExtender.contract.UnpackLog(event, "ExtensionFinished", log); err != nil {
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

// ContractExtenderNewContractExtensionContractCreatedIterator is returned from FilterNewContractExtensionContractCreated and is used to iterate over the raw logs and unpacked data for NewContractExtensionContractCreated events raised by the ContractExtender contract.
type ContractExtenderNewContractExtensionContractCreatedIterator struct {
	Event *ContractExtenderNewContractExtensionContractCreated // Event containing the contract specifics and raw log

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
func (it *ContractExtenderNewContractExtensionContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderNewContractExtensionContractCreated)
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
		it.Event = new(ContractExtenderNewContractExtensionContractCreated)
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
func (it *ContractExtenderNewContractExtensionContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderNewContractExtensionContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderNewContractExtensionContractCreated represents a NewContractExtensionContractCreated event raised by the ContractExtender contract.
type ContractExtenderNewContractExtensionContractCreated struct {
	ToExtend common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewContractExtensionContractCreated is a free log retrieval operation binding the contract event 0x1bb7909ad96bc757f60de4d9ce11daf7b006e8f398ce028dceb10ce7fdca0f68.
//
// Solidity: e NewContractExtensionContractCreated(toExtend address)
func (_ContractExtender *ContractExtenderFilterer) FilterNewContractExtensionContractCreated(opts *bind.FilterOpts) (*ContractExtenderNewContractExtensionContractCreatedIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "NewContractExtensionContractCreated")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderNewContractExtensionContractCreatedIterator{contract: _ContractExtender.contract, event: "NewContractExtensionContractCreated", logs: logs, sub: sub}, nil
}

var NewContractExtensionContractCreatedTopicHash = "0x1bb7909ad96bc757f60de4d9ce11daf7b006e8f398ce028dceb10ce7fdca0f68"

// WatchNewContractExtensionContractCreated is a free log subscription operation binding the contract event 0x1bb7909ad96bc757f60de4d9ce11daf7b006e8f398ce028dceb10ce7fdca0f68.
//
// Solidity: e NewContractExtensionContractCreated(toExtend address)
func (_ContractExtender *ContractExtenderFilterer) WatchNewContractExtensionContractCreated(opts *bind.WatchOpts, sink chan<- *ContractExtenderNewContractExtensionContractCreated) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "NewContractExtensionContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderNewContractExtensionContractCreated)
				if err := _ContractExtender.contract.UnpackLog(event, "NewContractExtensionContractCreated", log); err != nil {
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

// ContractExtenderNewVoteIterator is returned from FilterNewVote and is used to iterate over the raw logs and unpacked data for NewVote events raised by the ContractExtender contract.
type ContractExtenderNewVoteIterator struct {
	Event *ContractExtenderNewVote // Event containing the contract specifics and raw log

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
func (it *ContractExtenderNewVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderNewVote)
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
		it.Event = new(ContractExtenderNewVote)
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
func (it *ContractExtenderNewVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderNewVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderNewVote represents a NewVote event raised by the ContractExtender contract.
type ContractExtenderNewVote struct {
	Vote  bool
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewVote is a free log retrieval operation binding the contract event 0x225708d30006b0cc86d855ab91047edb5fe9c2e416412f36c18c6e90fe4e461f.
//
// Solidity: e NewVote(vote bool, voter address)
func (_ContractExtender *ContractExtenderFilterer) FilterNewVote(opts *bind.FilterOpts) (*ContractExtenderNewVoteIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "NewVote")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderNewVoteIterator{contract: _ContractExtender.contract, event: "NewVote", logs: logs, sub: sub}, nil
}

var NewVoteTopicHash = "0x225708d30006b0cc86d855ab91047edb5fe9c2e416412f36c18c6e90fe4e461f"

// WatchNewVote is a free log subscription operation binding the contract event 0x225708d30006b0cc86d855ab91047edb5fe9c2e416412f36c18c6e90fe4e461f.
//
// Solidity: e NewVote(vote bool, voter address)
func (_ContractExtender *ContractExtenderFilterer) WatchNewVote(opts *bind.WatchOpts, sink chan<- *ContractExtenderNewVote) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "NewVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderNewVote)
				if err := _ContractExtender.contract.UnpackLog(event, "NewVote", log); err != nil {
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

// ContractExtenderStateSharedIterator is returned from FilterStateShared and is used to iterate over the raw logs and unpacked data for StateShared events raised by the ContractExtender contract.
type ContractExtenderStateSharedIterator struct {
	Event *ContractExtenderStateShared // Event containing the contract specifics and raw log

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
func (it *ContractExtenderStateSharedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderStateShared)
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
		it.Event = new(ContractExtenderStateShared)
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
func (it *ContractExtenderStateSharedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderStateSharedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderStateShared represents a StateShared event raised by the ContractExtender contract.
type ContractExtenderStateShared struct {
	ToExtend    common.Address
	Tesserahash string
	Uuid        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStateShared is a free log retrieval operation binding the contract event 0x67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e.
//
// Solidity: e StateShared(toExtend address, tesserahash string, uuid string)
func (_ContractExtender *ContractExtenderFilterer) FilterStateShared(opts *bind.FilterOpts) (*ContractExtenderStateSharedIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "StateShared")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderStateSharedIterator{contract: _ContractExtender.contract, event: "StateShared", logs: logs, sub: sub}, nil
}

var StateSharedTopicHash = "0x67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e"

// WatchStateShared is a free log subscription operation binding the contract event 0x67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e.
//
// Solidity: e StateShared(toExtend address, tesserahash string, uuid string)
func (_ContractExtender *ContractExtenderFilterer) WatchStateShared(opts *bind.WatchOpts, sink chan<- *ContractExtenderStateShared) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "StateShared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderStateShared)
				if err := _ContractExtender.contract.UnpackLog(event, "StateShared", log); err != nil {
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

// ContractExtenderUpdateMembersIterator is returned from FilterUpdateMembers and is used to iterate over the raw logs and unpacked data for UpdateMembers events raised by the ContractExtender contract.
type ContractExtenderUpdateMembersIterator struct {
	Event *ContractExtenderUpdateMembers // Event containing the contract specifics and raw log

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
func (it *ContractExtenderUpdateMembersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderUpdateMembers)
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
		it.Event = new(ContractExtenderUpdateMembers)
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
func (it *ContractExtenderUpdateMembersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderUpdateMembersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderUpdateMembers represents a UpdateMembers event raised by the ContractExtender contract.
type ContractExtenderUpdateMembers struct {
	ToExtend common.Address
	Uuid     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateMembers is a free log retrieval operation binding the contract event 0x8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b44.
//
// Solidity: e UpdateMembers(toExtend address, uuid string)
func (_ContractExtender *ContractExtenderFilterer) FilterUpdateMembers(opts *bind.FilterOpts) (*ContractExtenderUpdateMembersIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "UpdateMembers")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderUpdateMembersIterator{contract: _ContractExtender.contract, event: "UpdateMembers", logs: logs, sub: sub}, nil
}

var UpdateMembersTopicHash = "0x8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b44"

// WatchUpdateMembers is a free log subscription operation binding the contract event 0x8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b44.
//
// Solidity: e UpdateMembers(toExtend address, uuid string)
func (_ContractExtender *ContractExtenderFilterer) WatchUpdateMembers(opts *bind.WatchOpts, sink chan<- *ContractExtenderUpdateMembers) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "UpdateMembers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderUpdateMembers)
				if err := _ContractExtender.contract.UnpackLog(event, "UpdateMembers", log); err != nil {
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

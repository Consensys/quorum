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
const ContractExtenderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractToExtend\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkIfExtensionFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNumberOfVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"walletAddressesToVote\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"setUuid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sharedDataHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setSharedStateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updatePartyMembers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteOutcome\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkIfVoted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vote\",\"type\":\"bool\"},{\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"doVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetRecipientPTMKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creationTxHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"haveAllNodesVoted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"recipientAddress\",\"type\":\"address\"},{\"name\":\"recipientPTMKey\",\"type\":\"string\"},{\"name\":\"creTxHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipientPTMKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"NewContractExtensionContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"outcome\",\"type\":\"bool\"}],\"name\":\"AllNodesHaveVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CanPerformStateShare\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ExtensionFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vote\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"NewVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tesserahash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"StateShared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UpdateMembers\",\"type\":\"event\"}]"

var ContractExtenderParsedABI, _ = abi.JSON(strings.NewReader(ContractExtenderABI))

// ContractExtenderBin is the compiled bytecode used for deploying new contracts.
var ContractExtenderBin = "0x60806040523480156200001157600080fd5b506040516200142538038062001425833981018060405260808110156200003757600080fd5b81516020830151604084018051929491938201926401000000008111156200005e57600080fd5b820160208101848111156200007257600080fd5b81516401000000008111828201871017156200008d57600080fd5b50509291906020018051640100000000811115620000aa57600080fd5b82016020810184811115620000be57600080fd5b8151640100000000811182820187101715620000d957600080fd5b5050600080546001600160a01b031916331790558051909350620001079250600191506020840190620002f0565b5081516200011d906002906020850190620002f0565b50600380546001600160a01b038087166001600160a01b031992831617909255600480546001818101835560008381527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9283018054861633179055835491820190935501805493871693909216929092179055604080516020810191829052829052620001af91600b9190620002f0565b50600a805460ff19166001179055600060078190555b6004548110156200022357600160066000600484815481101515620001e657fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff1916911515919091179055600101620001c5565b50600454600555604080516001600160a01b0380871682528516918101919091526060602080830182815285519284019290925284517f04576ede6057794ada68966eebc285c98a2726cbc4929ffd1ad9900336728d9393889387938993608084019186019080838360005b83811015620002a95781810151838201526020016200028f565b50505050905090810190601f168015620002d75780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15050505062000395565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200033357805160ff191683800117855562000363565b8280016001018555821562000363579182015b828111156200036357825182559160200191906001019062000346565b506200037192915062000375565b5090565b6200039291905b808211156200037157600081556001016200037c565b90565b61108080620003a56000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063ac8b9205116100a2578063d8bff5a511610071578063d8bff5a51461038d578063de5828cb146103b3578063e5af0f3014610462578063ec2fa2c11461046a578063f57077d81461047257610116565b8063ac8b92051461036d578063b5da45bb14610375578063cb2805ec1461037d578063d56b28891461038557610116565b806379d41b8f116100e957806379d41b8f1461017d5780637b3529621461019a578063821e93da146101a257806388f520a01461024a578063893971ba146102c757610116565b806302d05d3f1461011b57806315e56a6a1461013f5780631962cb9b146101475780633852772714610163575b600080fd5b61012361047a565b604080516001600160a01b039092168252519081900360200190f35b610123610489565b61014f610498565b604080519115158252519081900360200190f35b61016b6104a2565b60408051918252519081900360200190f35b6101236004803603602081101561019357600080fd5b50356104a8565b61014f6104d0565b610248600480360360208110156101b857600080fd5b8101906020810181356401000000008111156101d357600080fd5b8201836020820111156101e557600080fd5b8035906020019184600183028401116401000000008311171561020757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104d9945050505050565b005b610252610567565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561028c578181015183820152602001610274565b50505050905090810190601f1680156102b95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610248600480360360208110156102dd57600080fd5b8101906020810181356401000000008111156102f857600080fd5b82018360208201111561030a57600080fd5b8035906020019184600183028401116401000000008311171561032c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105f5945050505050565b61024861095f565b61014f610a5a565b61014f610a63565b610248610a79565b61014f600480360360208110156103a357600080fd5b50356001600160a01b0316610b14565b610248600480360360408110156103c957600080fd5b8135151591908101906040810160208201356401000000008111156103ed57600080fd5b8201836020820111156103ff57600080fd5b8035906020019184600183028401116401000000008311171561042157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b29945050505050565b610252610bce565b610252610c26565b61014f610c80565b6000546001600160a01b031681565b6003546001600160a01b031681565b600d5460ff165b90565b60055481565b60048054829081106104b657fe5b6000918252602090912001546001600160a01b0316905081565b600d5460ff1681565b600d5460ff161561051e57604051600160e51b62461bcd0281526004018080602001828103825260258152602001806110306025913960400191505060405180910390fd5b600c8054600181018083556000929092528251610562917fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c701906020850190610f4c565b505050565b600b805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105ed5780601f106105c2576101008083540402835291602001916105ed565b820191906000526020600020905b8154815290600101906020018083116105d057829003601f168201915b505050505081565b6000546001600160a01b0316331461064157604051600160e51b62461bcd02815260040180806020018281038252602381526020018061100d6023913960400191505060405180910390fd5b600d5460ff161561068657604051600160e51b62461bcd0281526004018080602001828103825260258152602001806110306025913960400191505060405180910390fd5b600b8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156107125780601f106106e757610100808354040283529160200191610712565b820191906000526020600020905b8154815290600101906020018083116106f557829003601f168201915b505085519394508593151592506107769150505760408051600160e51b62461bcd02815260206004820152601860248201527f6e657720686173682063616e6e6f7420626520656d7074790000000000000000604482015290519081900360640190fd5b8151156107cd5760408051600160e51b62461bcd02815260206004820152601660248201527f7374617465206861736820616c72656164792073657400000000000000000000604482015290519081900360640190fd5b82516107e090600b906020860190610f4c565b5060005b600c5481101561095657600354600c80547f67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e926001600160a01b031691600b918590811061082e57fe5b6000918252602091829020604080516001600160a01b038716815260609481018581528654600260001961010060018416150201909116049582018690529290930193908301906080840190869080156108c95780601f1061089e576101008083540402835291602001916108c9565b820191906000526020600020905b8154815290600101906020018083116108ac57829003601f168201915b505083810382528454600260001961010060018416150201909116048082526020909101908590801561093d5780601f106109125761010080835404028352916020019161093d565b820191906000526020600020905b81548152906001019060200180831161092057829003601f168201915b50509550505050505060405180910390a16001016107e4565b50610562610a79565b60005b600c54811015610a5757600354600c80547f8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b44926001600160a01b03169190849081106109aa57fe5b6000918252602091829020604080516001600160a01b0386168152938401818152919092018054600260001961010060018416150201909116049284018390529291606083019084908015610a405780601f10610a1557610100808354040283529160200191610a40565b820191906000526020600020905b815481529060010190602001808311610a2357829003601f168201915b5050935050505060405180910390a1600101610962565b50565b600a5460ff1681565b3360009081526008602052604090205460ff1690565b600d5460ff1615610abe57604051600160e51b62461bcd0281526004018080602001828103825260258152602001806110306025913960400191505060405180910390fd5b6000546001600160a01b03163314610b0a57604051600160e51b62461bcd02815260040180806020018281038252602381526020018061100d6023913960400191505060405180910390fd5b610b12610c8a565b565b60096020526000908152604090205460ff1681565b600d5460ff1615610b6e57604051600160e51b62461bcd0281526004018080602001828103825260258152602001806110306025913960400191505060405180910390fd5b610b7782610cc2565b8115610b8657610b86816104d9565b610b8e610e93565b60408051831515815233602082015281517f225708d30006b0cc86d855ab91047edb5fe9c2e416412f36c18c6e90fe4e461f929181900390910190a15050565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156105ed5780601f106105c2576101008083540402835291602001916105ed565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105ed5780601f106105c2576101008083540402835291602001916105ed565b6007546004541490565b600d805460ff191660011790556040517f79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede990600090a1565b600d5460ff1615610d0757604051600160e51b62461bcd028152600401808060200182810382526028815260200180610fe56028913960400191505060405180910390fd5b3360009081526006602052604090205460ff161515610d705760408051600160e51b62461bcd02815260206004820152601360248201527f6e6f7420616c6c6f77656420746f20766f746500000000000000000000000000604482015290519081900360640190fd5b3360009081526008602052604090205460ff1615610dd85760408051600160e51b62461bcd02815260206004820152600d60248201527f616c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b600a5460ff161515610e345760408051600160e51b62461bcd02815260206004820152601760248201527f766f74696e6720616c7265616479206465636c696e6564000000000000000000604482015290519081900360640190fd5b3360009081526008602090815260408083208054600160ff199182168117909255600990935292208054909116831515179055600780549091019055600a5460ff168015610e7f5750805b600a805460ff191691151591909117905550565b600a5460ff161515610ee057604080516000815290517fc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a29049181900360200190a1610edb610c8a565b610b12565b610ee8610c80565b15610b1257604080516001815290517fc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a29049181900360200190a16040517ffd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d2890600090a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f8d57805160ff1916838001178555610fba565b82800160010185558215610fba579182015b82811115610fba578251825591602001919060010190610f9f565b50610fc6929150610fca565b5090565b61049f91905b80821115610fc65760008155600101610fd056fe657874656e73696f6e2070726f6365737320636f6d706c657465642e2063616e6e6f7420766f74656f6e6c79206c6561646572206d617920706572666f726d207468697320616374696f6e657874656e73696f6e20686173206265656e206d61726b65642061732066696e6973686564a165627a7a72305820b12876502879651ce75b1d6723d37d71bd531f4d87db300e0348c68c6c1bbfed0029"

// DeployContractExtender deploys a new Ethereum contract, binding an instance of ContractExtender to it.
func DeployContractExtender(auth *bind.TransactOpts, backend bind.ContractBackend, contractAddress common.Address, recipientAddress common.Address, recipientPTMKey string, creTxHash string) (common.Address, *types.Transaction, *ContractExtender, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractExtenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractExtenderBin), backend, contractAddress, recipientAddress, recipientPTMKey, creTxHash)
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

// CreationTxHash is a free data retrieval call binding the contract method 0xec2fa2c1.
//
// Solidity: function creationTxHash() constant returns(string)
func (_ContractExtender *ContractExtenderCaller) CreationTxHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "creationTxHash")
	return *ret0, err
}

// CreationTxHash is a free data retrieval call binding the contract method 0xec2fa2c1.
//
// Solidity: function creationTxHash() constant returns(string)
func (_ContractExtender *ContractExtenderSession) CreationTxHash() (string, error) {
	return _ContractExtender.Contract.CreationTxHash(&_ContractExtender.CallOpts)
}

// CreationTxHash is a free data retrieval call binding the contract method 0xec2fa2c1.
//
// Solidity: function creationTxHash() constant returns(string)
func (_ContractExtender *ContractExtenderCallerSession) CreationTxHash() (string, error) {
	return _ContractExtender.Contract.CreationTxHash(&_ContractExtender.CallOpts)
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

// TargetRecipientPTMKey is a free data retrieval call binding the contract method 0xe5af0f30.
//
// Solidity: function targetRecipientPTMKey() constant returns(string)
func (_ContractExtender *ContractExtenderCaller) TargetRecipientPTMKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "targetRecipientPTMKey")
	return *ret0, err
}

// TargetRecipientPTMKey is a free data retrieval call binding the contract method 0xe5af0f30.
//
// Solidity: function targetRecipientPTMKey() constant returns(string)
func (_ContractExtender *ContractExtenderSession) TargetRecipientPTMKey() (string, error) {
	return _ContractExtender.Contract.TargetRecipientPTMKey(&_ContractExtender.CallOpts)
}

// TargetRecipientPTMKey is a free data retrieval call binding the contract method 0xe5af0f30.
//
// Solidity: function targetRecipientPTMKey() constant returns(string)
func (_ContractExtender *ContractExtenderCallerSession) TargetRecipientPTMKey() (string, error) {
	return _ContractExtender.Contract.TargetRecipientPTMKey(&_ContractExtender.CallOpts)
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
// Solidity: function votes(address ) constant returns(bool)
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
// Solidity: function votes(address ) constant returns(bool)
func (_ContractExtender *ContractExtenderSession) Votes(arg0 common.Address) (bool, error) {
	return _ContractExtender.Contract.Votes(&_ContractExtender.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) Votes(arg0 common.Address) (bool, error) {
	return _ContractExtender.Contract.Votes(&_ContractExtender.CallOpts, arg0)
}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote(uint256 ) constant returns(address)
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
// Solidity: function walletAddressesToVote(uint256 ) constant returns(address)
func (_ContractExtender *ContractExtenderSession) WalletAddressesToVote(arg0 *big.Int) (common.Address, error) {
	return _ContractExtender.Contract.WalletAddressesToVote(&_ContractExtender.CallOpts, arg0)
}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote(uint256 ) constant returns(address)
func (_ContractExtender *ContractExtenderCallerSession) WalletAddressesToVote(arg0 *big.Int) (common.Address, error) {
	return _ContractExtender.Contract.WalletAddressesToVote(&_ContractExtender.CallOpts, arg0)
}

// DoVote is a paid mutator transaction binding the contract method 0xde5828cb.
//
// Solidity: function doVote(bool vote, string nextuuid) returns()
func (_ContractExtender *ContractExtenderTransactor) DoVote(opts *bind.TransactOpts, vote bool, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "doVote", vote, nextuuid)
}

// DoVote is a paid mutator transaction binding the contract method 0xde5828cb.
//
// Solidity: function doVote(bool vote, string nextuuid) returns()
func (_ContractExtender *ContractExtenderSession) DoVote(vote bool, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.DoVote(&_ContractExtender.TransactOpts, vote, nextuuid)
}

// DoVote is a paid mutator transaction binding the contract method 0xde5828cb.
//
// Solidity: function doVote(bool vote, string nextuuid) returns()
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
// Solidity: function setSharedStateHash(string hash) returns()
func (_ContractExtender *ContractExtenderTransactor) SetSharedStateHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "setSharedStateHash", hash)
}

// SetSharedStateHash is a paid mutator transaction binding the contract method 0x893971ba.
//
// Solidity: function setSharedStateHash(string hash) returns()
func (_ContractExtender *ContractExtenderSession) SetSharedStateHash(hash string) (*types.Transaction, error) {
	return _ContractExtender.Contract.SetSharedStateHash(&_ContractExtender.TransactOpts, hash)
}

// SetSharedStateHash is a paid mutator transaction binding the contract method 0x893971ba.
//
// Solidity: function setSharedStateHash(string hash) returns()
func (_ContractExtender *ContractExtenderTransactorSession) SetSharedStateHash(hash string) (*types.Transaction, error) {
	return _ContractExtender.Contract.SetSharedStateHash(&_ContractExtender.TransactOpts, hash)
}

// SetUuid is a paid mutator transaction binding the contract method 0x821e93da.
//
// Solidity: function setUuid(string nextuuid) returns()
func (_ContractExtender *ContractExtenderTransactor) SetUuid(opts *bind.TransactOpts, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "setUuid", nextuuid)
}

// SetUuid is a paid mutator transaction binding the contract method 0x821e93da.
//
// Solidity: function setUuid(string nextuuid) returns()
func (_ContractExtender *ContractExtenderSession) SetUuid(nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.SetUuid(&_ContractExtender.TransactOpts, nextuuid)
}

// SetUuid is a paid mutator transaction binding the contract method 0x821e93da.
//
// Solidity: function setUuid(string nextuuid) returns()
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
// Solidity: event AllNodesHaveVoted(bool outcome)
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
// Solidity: event AllNodesHaveVoted(bool outcome)
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

// ParseAllNodesHaveVoted is a log parse operation binding the contract event 0xc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a2904.
//
// Solidity: event AllNodesHaveVoted(bool outcome)
func (_ContractExtender *ContractExtenderFilterer) ParseAllNodesHaveVoted(log types.Log) (*ContractExtenderAllNodesHaveVoted, error) {
	event := new(ContractExtenderAllNodesHaveVoted)
	if err := _ContractExtender.contract.UnpackLog(event, "AllNodesHaveVoted", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event CanPerformStateShare()
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
// Solidity: event CanPerformStateShare()
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

// ParseCanPerformStateShare is a log parse operation binding the contract event 0xfd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d28.
//
// Solidity: event CanPerformStateShare()
func (_ContractExtender *ContractExtenderFilterer) ParseCanPerformStateShare(log types.Log) (*ContractExtenderCanPerformStateShare, error) {
	event := new(ContractExtenderCanPerformStateShare)
	if err := _ContractExtender.contract.UnpackLog(event, "CanPerformStateShare", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event ExtensionFinished()
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
// Solidity: event ExtensionFinished()
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

// ParseExtensionFinished is a log parse operation binding the contract event 0x79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede9.
//
// Solidity: event ExtensionFinished()
func (_ContractExtender *ContractExtenderFilterer) ParseExtensionFinished(log types.Log) (*ContractExtenderExtensionFinished, error) {
	event := new(ContractExtenderExtensionFinished)
	if err := _ContractExtender.contract.UnpackLog(event, "ExtensionFinished", log); err != nil {
		return nil, err
	}
	return event, nil
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
	ToExtend         common.Address
	RecipientPTMKey  string
	RecipientAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewContractExtensionContractCreated is a free log retrieval operation binding the contract event 0x04576ede6057794ada68966eebc285c98a2726cbc4929ffd1ad9900336728d93.
//
// Solidity: event NewContractExtensionContractCreated(address toExtend, string recipientPTMKey, address recipientAddress)
func (_ContractExtender *ContractExtenderFilterer) FilterNewContractExtensionContractCreated(opts *bind.FilterOpts) (*ContractExtenderNewContractExtensionContractCreatedIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "NewContractExtensionContractCreated")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderNewContractExtensionContractCreatedIterator{contract: _ContractExtender.contract, event: "NewContractExtensionContractCreated", logs: logs, sub: sub}, nil
}

var NewContractExtensionContractCreatedTopicHash = "0x04576ede6057794ada68966eebc285c98a2726cbc4929ffd1ad9900336728d93"

// WatchNewContractExtensionContractCreated is a free log subscription operation binding the contract event 0x04576ede6057794ada68966eebc285c98a2726cbc4929ffd1ad9900336728d93.
//
// Solidity: event NewContractExtensionContractCreated(address toExtend, string recipientPTMKey, address recipientAddress)
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

// ParseNewContractExtensionContractCreated is a log parse operation binding the contract event 0x04576ede6057794ada68966eebc285c98a2726cbc4929ffd1ad9900336728d93.
//
// Solidity: event NewContractExtensionContractCreated(address toExtend, string recipientPTMKey, address recipientAddress)
func (_ContractExtender *ContractExtenderFilterer) ParseNewContractExtensionContractCreated(log types.Log) (*ContractExtenderNewContractExtensionContractCreated, error) {
	event := new(ContractExtenderNewContractExtensionContractCreated)
	if err := _ContractExtender.contract.UnpackLog(event, "NewContractExtensionContractCreated", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event NewVote(bool vote, address voter)
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
// Solidity: event NewVote(bool vote, address voter)
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

// ParseNewVote is a log parse operation binding the contract event 0x225708d30006b0cc86d855ab91047edb5fe9c2e416412f36c18c6e90fe4e461f.
//
// Solidity: event NewVote(bool vote, address voter)
func (_ContractExtender *ContractExtenderFilterer) ParseNewVote(log types.Log) (*ContractExtenderNewVote, error) {
	event := new(ContractExtenderNewVote)
	if err := _ContractExtender.contract.UnpackLog(event, "NewVote", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event StateShared(address toExtend, string tesserahash, string uuid)
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
// Solidity: event StateShared(address toExtend, string tesserahash, string uuid)
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

// ParseStateShared is a log parse operation binding the contract event 0x67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e.
//
// Solidity: event StateShared(address toExtend, string tesserahash, string uuid)
func (_ContractExtender *ContractExtenderFilterer) ParseStateShared(log types.Log) (*ContractExtenderStateShared, error) {
	event := new(ContractExtenderStateShared)
	if err := _ContractExtender.contract.UnpackLog(event, "StateShared", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event UpdateMembers(address toExtend, string uuid)
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
// Solidity: event UpdateMembers(address toExtend, string uuid)
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

// ParseUpdateMembers is a log parse operation binding the contract event 0x8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b44.
//
// Solidity: event UpdateMembers(address toExtend, string uuid)
func (_ContractExtender *ContractExtenderFilterer) ParseUpdateMembers(log types.Log) (*ContractExtenderUpdateMembers, error) {
	event := new(ContractExtenderUpdateMembers)
	if err := _ContractExtender.contract.UnpackLog(event, "UpdateMembers", log); err != nil {
		return nil, err
	}
	return event, nil
}

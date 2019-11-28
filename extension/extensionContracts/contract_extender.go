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
const ContractExtenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"walletAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"recipientHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"outcome\",\"type\":\"bool\"}],\"name\":\"AllNodesHaveVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CanPerformStateShare\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ExtensionFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toExtend\",\"type\":\"address\"}],\"name\":\"NewContractExtensionContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tesserahash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"StateShared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UpdateMembers\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractToExtend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"vote\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"doVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"haveAllNodesVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setSharedStateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"setUuid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"shareAcceptStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sharedDataHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetHasAccepted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetRecipientPublicKeyHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNumberOfVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updatePartyMembers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteOutcome\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"walletAddressesToVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var ContractExtenderParsedABI, _ = abi.JSON(strings.NewReader(ContractExtenderABI))

// ContractExtenderBin is the compiled bytecode used for deploying new contracts.
const ContractExtenderBin = `0x60806040523480156200001157600080fd5b506040516200187c3803806200187c833981810160405260808110156200003757600080fd5b8151602083018051604051929492938301929190846401000000008211156200005f57600080fd5b9083019060208201858111156200007557600080fd5b82518660208202830111640100000000821117156200009357600080fd5b82525081516020918201928201910280838360005b83811015620000c2578181015183820152602001620000a8565b5050505090500160405260200180516040519392919084640100000000821115620000ec57600080fd5b9083019060208201858111156200010257600080fd5b82516401000000008111828201881017156200011d57600080fd5b82525081516020918201929091019080838360005b838110156200014c57818101518382015260200162000132565b50505050905090810190601f1680156200017a5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200019e57600080fd5b908301906020820185811115620001b457600080fd5b8251640100000000811182820188101715620001cf57600080fd5b82525081516020918201929091019080838360005b83811015620001fe578181015183820152602001620001e4565b50505050905090810190601f1680156200022c5780820380516001836020036101000a031916815260200191505b506040525050600080546001600160a01b031916331790555081516200025a90600190602085019062000724565b5060028054610100600160a81b0319166101006001600160a01b03871602179055825162000290906003906020860190620007a9565b50604080516020810191829052600090819052620002b191600a9162000724565b506009805460ff1916600117905560006006819055805b84518110156200036157600160056000878481518110620002e557fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550336001600160a01b03168582815181106200033b57fe5b60200260200101516001600160a01b031614156200035857600191505b600101620002c8565b5080620003c3576003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b031916339081179091556000908152600560205260409020805460ff191690911790555b600354600455604080516001600160a01b038716815290517f1bb7909ad96bc757f60de4d9ce11daf7b006e8f398ce028dceb10ce7fdca0f689181900360200190a16200041a60016001600160e01b036200043816565b6200042d6001600160e01b036200060916565b505050505062000853565b3360009081526005602052604090205460ff16620004b757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f6e6f7420616c6c6f77656420746f20766f746500000000000000000000000000604482015290519081900360640190fd5b3360009081526007602052604090205460ff16156200053757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f616c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b60095460ff16620005a957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f766f74696e6720616c7265616479206465636c696e6564000000000000000000604482015290519081900360640190fd5b3360009081526007602090815260408083208054600160ff19918216811790925560089093529220805490911683151517905560068054909101905560095460ff168015620005f55750805b6009805460ff191691151591909117905550565b60095460ff166200065057604080516000815290516000805160206200185c8339815191529181900360200190a16200064a6001600160e01b03620006e116565b620006df565b620006636001600160e01b036200071916565b156200068d57604080516001815290516000805160206200185c8339815191529181900360200190a15b620006a06001600160e01b036200071916565b8015620006af575060025460ff165b15620006df576040517ffd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d2890600090a15b565b600c805460ff191660011790556040517f79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede990600090a1565b600654600354145b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200076757805160ff191683800117855562000797565b8280016001018555821562000797579182015b82811115620007975782518255916020019190600101906200077a565b50620007a59291506200080f565b5090565b82805482825590600052602060002090810192821562000801579160200282015b828111156200080157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620007ca565b50620007a59291506200082c565b6200072191905b80821115620007a5576000815560010162000816565b6200072191905b80821115620007a55780546001600160a01b031916815560010162000833565b610ff980620008636000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063821e93da116100a2578063b5da45bb11610071578063b5da45bb14610412578063d56b28891461041a578063d8bff5a514610422578063de5828cb14610448578063f57077d8146104f55761010b565b8063821e93da146102ba57806388f520a01461035e578063893971ba14610366578063ac8b92051461040a5761010b565b80634688a5e3116100de5780634688a5e314610172578063666fd2a4146101ef57806379d41b8f146102955780637b352962146102b25761010b565b806302d05d3f1461011057806315e56a6a14610134578063385277271461013c5780634242b7a614610156575b600080fd5b6101186104fd565b604080516001600160a01b039092168252519081900360200190f35b61011861050c565b610144610520565b60408051918252519081900360200190f35b61015e610526565b604080519115158252519081900360200190f35b61017a61052f565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101b457818101518382015260200161019c565b50505050905090810190601f1680156101e15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102936004803603602081101561020557600080fd5b810190602081018135600160201b81111561021f57600080fd5b82018360208201111561023157600080fd5b803590602001918460018302840111600160201b8311171561025257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105bc945050505050565b005b610118600480360360208110156102ab57600080fd5b50356105dd565b61015e610604565b610293600480360360208110156102d057600080fd5b810190602081018135600160201b8111156102ea57600080fd5b8201836020820111156102fc57600080fd5b803590602001918460018302840111600160201b8311171561031d57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061060d945050505050565b61017a610698565b6102936004803603602081101561037c57600080fd5b810190602081018135600160201b81111561039657600080fd5b8201836020820111156103a857600080fd5b803590602001918460018302840111600160201b831117156103c957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106f3945050505050565b610293610a52565b61015e610b50565b610293610b59565b61015e6004803603602081101561043857600080fd5b50356001600160a01b0316610bee565b6102936004803603604081101561045e57600080fd5b813515159190810190604081016020820135600160201b81111561048157600080fd5b82018360208201111561049357600080fd5b803590602001918460018302840111600160201b831117156104b457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c03945050505050565b61015e610c69565b6000546001600160a01b031681565b60025461010090046001600160a01b031681565b60045481565b60025460ff1681565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105b45780601f10610589576101008083540402835291602001916105b4565b820191906000526020600020905b81548152906001019060200180831161059757829003601f168201915b505050505081565b6105c58161060d565b6002805460ff191660011790556105da610c74565b50565b600381815481106105ea57fe5b6000918252602090912001546001600160a01b0316905081565b600c5460ff1681565b600c5460ff161561064f5760405162461bcd60e51b8152600401808060200182810382526025815260200180610fa06025913960400191505060405180910390fd5b600b8054600181018083556000929092528251610693917f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db901906020850190610ee4565b505050565b600a805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105b45780601f10610589576101008083540402835291602001916105b4565b6000546001600160a01b0316331461073c5760405162461bcd60e51b8152600401808060200182810382526023815260200180610f7d6023913960400191505060405180910390fd5b600c5460ff161561077e5760405162461bcd60e51b8152600401808060200182810382526025815260200180610fa06025913960400191505060405180910390fd5b600a8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561080a5780601f106107df5761010080835404028352916020019161080a565b820191906000526020600020905b8154815290600101906020018083116107ed57829003601f168201915b50505050509050606082905080516000141561086d576040805162461bcd60e51b815260206004820152601860248201527f6e657720686173682063616e6e6f7420626520656d7074790000000000000000604482015290519081900360640190fd5b8151156108ba576040805162461bcd60e51b81526020600482015260166024820152751cdd185d19481a185cda08185b1c9958591e481cd95d60521b604482015290519081900360640190fd5b82516108cd90600a906020860190610ee4565b5060005b600b54811015610a49577f67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e600260019054906101000a90046001600160a01b0316600a600b848154811061092157fe5b6000918252602091829020604080516001600160a01b038716815260609481018581528654600260001961010060018416150201909116049582018690529290930193908301906080840190869080156109bc5780601f10610991576101008083540402835291602001916109bc565b820191906000526020600020905b81548152906001019060200180831161099f57829003601f168201915b5050838103825284546002600019610100600184161502019091160480825260209091019085908015610a305780601f10610a0557610100808354040283529160200191610a30565b820191906000526020600020905b815481529060010190602001808311610a1357829003601f168201915b50509550505050505060405180910390a16001016108d1565b50610693610b59565b60005b600b548110156105da577f8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b44600260019054906101000a90046001600160a01b0316600b8381548110610aa357fe5b6000918252602091829020604080516001600160a01b0386168152938401818152919092018054600260001961010060018416150201909116049284018390529291606083019084908015610b395780601f10610b0e57610100808354040283529160200191610b39565b820191906000526020600020905b815481529060010190602001808311610b1c57829003601f168201915b5050935050505060405180910390a1600101610a55565b60095460ff1681565b600c5460ff1615610b9b5760405162461bcd60e51b8152600401808060200182810382526025815260200180610fa06025913960400191505060405180910390fd5b6000546001600160a01b03163314610be45760405162461bcd60e51b8152600401808060200182810382526023815260200180610f7d6023913960400191505060405180910390fd5b610bec610d47565b565b60086020526000908152604090205460ff1681565b600c5460ff1615610c455760405162461bcd60e51b8152600401808060200182810382526025815260200180610fa06025913960400191505060405180910390fd5b610c4e82610d7f565b8115610c5d57610c5d8161060d565b610c65610c74565b5050565b600654600354145b90565b60095460ff16610cbf57604080516000815290517fc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a29049181900360200190a1610cba610d47565b610bec565b610cc7610c69565b15610d0157604080516001815290517fc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a29049181900360200190a15b610d09610c69565b8015610d17575060025460ff165b15610bec576040517ffd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d2890600090a1565b600c805460ff191660011790556040517f79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede990600090a1565b3360009081526005602052604090205460ff16610dd9576040805162461bcd60e51b81526020600482015260136024820152726e6f7420616c6c6f77656420746f20766f746560681b604482015290519081900360640190fd5b3360009081526007602052604090205460ff1615610e2e576040805162461bcd60e51b815260206004820152600d60248201526c185b1c9958591e481d9bdd1959609a1b604482015290519081900360640190fd5b60095460ff16610e85576040805162461bcd60e51b815260206004820152601760248201527f766f74696e6720616c7265616479206465636c696e6564000000000000000000604482015290519081900360640190fd5b3360009081526007602090815260408083208054600160ff19918216811790925560089093529220805490911683151517905560068054909101905560095460ff168015610ed05750805b6009805460ff191691151591909117905550565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f2557805160ff1916838001178555610f52565b82800160010185558215610f52579182015b82811115610f52578251825591602001919060010190610f37565b50610f5e929150610f62565b5090565b610c7191905b80821115610f5e5760008155600101610f6856fe6f6e6c79206c6561646572206d617920706572666f726d207468697320616374696f6e657874656e73696f6e20686173206265656e206d61726b65642061732066696e6973686564a265627a7a72315820063c59db268c7d8d27c406fde9408580ebad979655ac2a6075e4ad01cfa5057b64736f6c634300050d0032c05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a2904`

// DeployContractExtender deploys a new Ethereum contract, binding an instance of ContractExtender to it.
func DeployContractExtender(auth *bind.TransactOpts, backend bind.ContractBackend, contractAddress common.Address, walletAddresses []common.Address, recipientHash string, uuid string) (common.Address, *types.Transaction, *ContractExtender, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractExtenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractExtenderBin), backend, contractAddress, walletAddresses, recipientHash, uuid)
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

// TargetHasAccepted is a free data retrieval call binding the contract method 0x4242b7a6.
//
// Solidity: function targetHasAccepted() constant returns(bool)
func (_ContractExtender *ContractExtenderCaller) TargetHasAccepted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractExtender.contract.Call(opts, out, "targetHasAccepted")
	return *ret0, err
}

// TargetHasAccepted is a free data retrieval call binding the contract method 0x4242b7a6.
//
// Solidity: function targetHasAccepted() constant returns(bool)
func (_ContractExtender *ContractExtenderSession) TargetHasAccepted() (bool, error) {
	return _ContractExtender.Contract.TargetHasAccepted(&_ContractExtender.CallOpts)
}

// TargetHasAccepted is a free data retrieval call binding the contract method 0x4242b7a6.
//
// Solidity: function targetHasAccepted() constant returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) TargetHasAccepted() (bool, error) {
	return _ContractExtender.Contract.TargetHasAccepted(&_ContractExtender.CallOpts)
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

// ShareAcceptStatus is a paid mutator transaction binding the contract method 0x666fd2a4.
//
// Solidity: function shareAcceptStatus(nextuuid string) returns()
func (_ContractExtender *ContractExtenderTransactor) ShareAcceptStatus(opts *bind.TransactOpts, nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.contract.Transact(opts, "shareAcceptStatus", nextuuid)
}

// ShareAcceptStatus is a paid mutator transaction binding the contract method 0x666fd2a4.
//
// Solidity: function shareAcceptStatus(nextuuid string) returns()
func (_ContractExtender *ContractExtenderSession) ShareAcceptStatus(nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.ShareAcceptStatus(&_ContractExtender.TransactOpts, nextuuid)
}

// ShareAcceptStatus is a paid mutator transaction binding the contract method 0x666fd2a4.
//
// Solidity: function shareAcceptStatus(nextuuid string) returns()
func (_ContractExtender *ContractExtenderTransactorSession) ShareAcceptStatus(nextuuid string) (*types.Transaction, error) {
	return _ContractExtender.Contract.ShareAcceptStatus(&_ContractExtender.TransactOpts, nextuuid)
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
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewVote is a free log retrieval operation binding the contract event 0x7dd2f5e995795c1d5e48803234b25a9a1dab98dd14e968ebe9bf3ae069ac9e7e.
//
// Solidity: e NewVote()
func (_ContractExtender *ContractExtenderFilterer) FilterNewVote(opts *bind.FilterOpts) (*ContractExtenderNewVoteIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "NewVote")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderNewVoteIterator{contract: _ContractExtender.contract, event: "NewVote", logs: logs, sub: sub}, nil
}

var NewVoteTopicHash = "0x7dd2f5e995795c1d5e48803234b25a9a1dab98dd14e968ebe9bf3ae069ac9e7e"

// WatchNewVote is a free log subscription operation binding the contract event 0x7dd2f5e995795c1d5e48803234b25a9a1dab98dd14e968ebe9bf3ae069ac9e7e.
//
// Solidity: e NewVote()
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

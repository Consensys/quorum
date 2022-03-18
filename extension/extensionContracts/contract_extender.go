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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractExtenderABI is the input ABI used to generate the binding from.
const ContractExtenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"recipientPTMKey\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"outcome\",\"type\":\"bool\"}],\"name\":\"AllNodesHaveAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CanPerformStateShare\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ExtensionFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipientPTMKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"NewContractExtensionContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"vote\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"NewVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tesserahash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"StateShared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toExtend\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UpdateMembers\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkIfExtensionFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkIfVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractToExtend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"vote\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"doVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"haveAllNodesVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setSharedStateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"nextuuid\",\"type\":\"string\"}],\"name\":\"setUuid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sharedDataHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetRecipientPTMKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNumberOfVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updatePartyMembers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteOutcome\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"walletAddressesToVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var ContractExtenderParsedABI, _ = abi.JSON(strings.NewReader(ContractExtenderABI))

// ContractExtenderBin is the compiled bytecode used for deploying new contracts.
var ContractExtenderBin = "0x60806040523480156200001157600080fd5b5060405162001cbb38038062001cbb833981810160405260608110156200003757600080fd5b810190808051906020019092919080519060200190929190805160405193929190846401000000008211156200006c57600080fd5b838201915060208201858111156200008357600080fd5b8251866001820283011164010000000082111715620000a157600080fd5b8083526020830192505050908051906020019080838360005b83811015620000d7578082015181840152602081019050620000ba565b50505050905090810190601f168015620001055780820380516001836020036101000a031916815260200191505b50604052505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060019080519060200190620001649291906200048c565b5082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060033390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060038290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060405180602001604052806000815250600a9080519060200190620002999291906200048c565b506001600960006101000a81548160ff021916908315150217905550600060068190555060008090505b6003805490508110156200036f5760016005600060038481548110620002e557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050620002c3565b506003805490506004819055507f04576ede6057794ada68966eebc285c98a2726cbc4929ffd1ad9900336728d93838284604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156200044657808201518184015260208101905062000429565b50505050905090810190601f168015620004745780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15050506200053b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004cf57805160ff191683800117855562000500565b8280016001018555821562000500579182015b82811115620004ff578251825591602001919060010190620004e2565b5b5090506200050f919062000513565b5090565b6200053891905b80821115620005345760008160009055506001016200051a565b5090565b90565b611770806200054b6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063893971ba116100a2578063d56b288911610071578063d56b2889146104bb578063d8bff5a5146104c5578063de5828cb14610521578063e5af0f30146105e8578063f57077d81461066b5761010b565b8063893971ba146103b2578063ac8b92051461046d578063b5da45bb14610477578063cb2805ec146104995761010b565b806379d41b8f116100de57806379d41b8f146101e45780637b35296214610252578063821e93da1461027457806388f520a01461032f5761010b565b806302d05d3f1461011057806315e56a6a1461015a5780631962cb9b146101a457806338527727146101c6575b600080fd5b61011861068d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101626106b2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101ac6106d8565b604051808215151515815260200191505060405180910390f35b6101ce6106ef565b6040518082815260200191505060405180910390f35b610210600480360360208110156101fa57600080fd5b81019080803590602001909291905050506106f5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61025a610731565b604051808215151515815260200191505060405180910390f35b61032d6004803603602081101561028a57600080fd5b81019080803590602001906401000000008111156102a757600080fd5b8201836020820111156102b957600080fd5b803590602001918460018302840111640100000000831117156102db57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610744565b005b6103376107ec565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561037757808201518184015260208101905061035c565b50505050905090810190601f1680156103a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61046b600480360360208110156103c857600080fd5b81019080803590602001906401000000008111156103e557600080fd5b8201836020820111156103f757600080fd5b8035906020019184600183028401116401000000008311171561041957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061088a565b005b610475610d1d565b005b61047f610e65565b604051808215151515815260200191505060405180910390f35b6104a1610e78565b604051808215151515815260200191505060405180910390f35b6104c3610ecc565b005b610507600480360360208110156104db57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fe1565b604051808215151515815260200191505060405180910390f35b6105e66004803603604081101561053757600080fd5b810190808035151590602001909291908035906020019064010000000081111561056057600080fd5b82018360208201111561057257600080fd5b8035906020019184600183028401116401000000008311171561059457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611001565b005b6105f06110fb565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610630578082015181840152602081019050610615565b50505050905090810190601f16801561065d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610673611199565b604051808215151515815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600c60009054906101000a900460ff16905090565b60045481565b6003818154811061070257fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c60009054906101000a900460ff1681565b600c60009054906101000a900460ff16156107aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806117176025913960400191505060405180910390fd5b600b8190806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906107e7929190611626565b505050565b600a8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108825780601f1061085757610100808354040283529160200191610882565b820191906000526020600020905b81548152906001019060200180831161086557829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461092f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806116f46023913960400191505060405180910390fd5b600c60009054906101000a900460ff1615610995576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806117176025913960400191505060405180910390fd5b6060600a8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a2d5780601f10610a0257610100808354040283529160200191610a2d565b820191906000526020600020905b815481529060010190602001808311610a1057829003601f168201915b505050505090506060829050600081511415610ab1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6e657720686173682063616e6e6f7420626520656d707479000000000000000081525060200191505060405180910390fd5b6000825114610b28576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f7374617465206861736820616c7265616479207365740000000000000000000081525060200191505060405180910390fd5b82600a9080519060200190610b3e929190611626565b5060008090505b600b80549050811015610d0f577f67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600a600b8481548110610ba557fe5b90600052602060002001604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001838103835285818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610c6e5780601f10610c4357610100808354040283529160200191610c6e565b820191906000526020600020905b815481529060010190602001808311610c5157829003601f168201915b5050838103825284818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610cf15780601f10610cc657610100808354040283529160200191610cf1565b820191906000526020600020905b815481529060010190602001808311610cd457829003601f168201915b50509550505050505060405180910390a18080600101915050610b45565b50610d18610ecc565b505050565b60008090505b600b80549050811015610e62577f8adc4573f947f9930560525736f61b116be55049125cb63a36887a40f92f3b44600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600b8381548110610d8157fe5b90600052602060002001604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610e465780601f10610e1b57610100808354040283529160200191610e46565b820191906000526020600020905b815481529060010190602001808311610e2957829003601f168201915b5050935050505060405180910390a18080600101915050610d23565b50565b600960009054906101000a900460ff1681565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905090565b600c60009054906101000a900460ff1615610f32576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806117176025913960400191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fd7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806116f46023913960400191505060405180910390fd5b610fdf6111aa565b565b60086020528060005260406000206000915054906101000a900460ff1681565b600c60009054906101000a900460ff1615611067576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806117176025913960400191505060405180910390fd5b611070826111f3565b81156110805761107f81610744565b5b611088611550565b7f225708d30006b0cc86d855ab91047edb5fe9c2e416412f36c18c6e90fe4e461f823360405180831515151581526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a15050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111915780601f1061116657610100808354040283529160200191611191565b820191906000526020600020905b81548152906001019060200180831161117457829003601f168201915b505050505081565b600060065460038054905014905090565b6001600c60006101000a81548160ff0219169083151502179055507f79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede960405160405180910390a1565b600c60009054906101000a900460ff1615611259576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806116cc6028913960400191505060405180910390fd5b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611318576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6e6f7420616c6c6f77656420746f20766f74650000000000000000000000000081525060200191505060405180910390fd5b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156113d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f616c726561647920766f7465640000000000000000000000000000000000000081525060200191505060405180910390fd5b600960009054906101000a900460ff1661145a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f766f74696e6720616c7265616479206465636c696e656400000000000000000081525060200191505060405180910390fd5b6001600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600660008154809291906001019190505550600960009054906101000a900460ff1680156115345750805b600960006101000a81548160ff02191690831515021790555050565b600960009054906101000a900460ff166115ad577ff20540914db019dd7c8d05ed165316a58d1583642772ac46f3d0c29b8644bd366000604051808215151515815260200191505060405180910390a16115a86111aa565b611624565b6115b5611199565b15611623577ff20540914db019dd7c8d05ed165316a58d1583642772ac46f3d0c29b8644bd366001604051808215151515815260200191505060405180910390a17ffd46cafaa71d87561071b8095703a7f081265fad232945049f5cf2d2c39b3d2860405160405180910390a15b5b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061166757805160ff1916838001178555611695565b82800160010185558215611695579182015b82811115611694578251825591602001919060010190611679565b5b5090506116a291906116a6565b5090565b6116c891905b808211156116c45760008160009055506001016116ac565b5090565b9056fe657874656e73696f6e2070726f6365737320636f6d706c657465642e2063616e6e6f7420766f74656f6e6c79206c6561646572206d617920706572666f726d207468697320616374696f6e657874656e73696f6e20686173206265656e206d61726b65642061732066696e6973686564a265627a7a72315820625108b92f7ff30d44757ae1bb19335828b2892b67a277794ea401fa969f7bdf64736f6c63430005110032"

// DeployContractExtender deploys a new Ethereum contract, binding an instance of ContractExtender to it.
func DeployContractExtender(auth *bind.TransactOpts, backend bind.ContractBackend, contractAddress common.Address, recipientAddress common.Address, recipientPTMKey string) (common.Address, *types.Transaction, *ContractExtender, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractExtenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractExtenderBin), backend, contractAddress, recipientAddress, recipientPTMKey)
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
func (_ContractExtender *ContractExtenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_ContractExtender *ContractExtenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function checkIfExtensionFinished() view returns(bool)
func (_ContractExtender *ContractExtenderCaller) CheckIfExtensionFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "checkIfExtensionFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfExtensionFinished is a free data retrieval call binding the contract method 0x1962cb9b.
//
// Solidity: function checkIfExtensionFinished() view returns(bool)
func (_ContractExtender *ContractExtenderSession) CheckIfExtensionFinished() (bool, error) {
	return _ContractExtender.Contract.CheckIfExtensionFinished(&_ContractExtender.CallOpts)
}

// CheckIfExtensionFinished is a free data retrieval call binding the contract method 0x1962cb9b.
//
// Solidity: function checkIfExtensionFinished() view returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) CheckIfExtensionFinished() (bool, error) {
	return _ContractExtender.Contract.CheckIfExtensionFinished(&_ContractExtender.CallOpts)
}

// CheckIfVoted is a free data retrieval call binding the contract method 0xcb2805ec.
//
// Solidity: function checkIfVoted() view returns(bool)
func (_ContractExtender *ContractExtenderCaller) CheckIfVoted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "checkIfVoted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfVoted is a free data retrieval call binding the contract method 0xcb2805ec.
//
// Solidity: function checkIfVoted() view returns(bool)
func (_ContractExtender *ContractExtenderSession) CheckIfVoted() (bool, error) {
	return _ContractExtender.Contract.CheckIfVoted(&_ContractExtender.CallOpts)
}

// CheckIfVoted is a free data retrieval call binding the contract method 0xcb2805ec.
//
// Solidity: function checkIfVoted() view returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) CheckIfVoted() (bool, error) {
	return _ContractExtender.Contract.CheckIfVoted(&_ContractExtender.CallOpts)
}

// ContractToExtend is a free data retrieval call binding the contract method 0x15e56a6a.
//
// Solidity: function contractToExtend() view returns(address)
func (_ContractExtender *ContractExtenderCaller) ContractToExtend(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "contractToExtend")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractToExtend is a free data retrieval call binding the contract method 0x15e56a6a.
//
// Solidity: function contractToExtend() view returns(address)
func (_ContractExtender *ContractExtenderSession) ContractToExtend() (common.Address, error) {
	return _ContractExtender.Contract.ContractToExtend(&_ContractExtender.CallOpts)
}

// ContractToExtend is a free data retrieval call binding the contract method 0x15e56a6a.
//
// Solidity: function contractToExtend() view returns(address)
func (_ContractExtender *ContractExtenderCallerSession) ContractToExtend() (common.Address, error) {
	return _ContractExtender.Contract.ContractToExtend(&_ContractExtender.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ContractExtender *ContractExtenderCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ContractExtender *ContractExtenderSession) Creator() (common.Address, error) {
	return _ContractExtender.Contract.Creator(&_ContractExtender.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_ContractExtender *ContractExtenderCallerSession) Creator() (common.Address, error) {
	return _ContractExtender.Contract.Creator(&_ContractExtender.CallOpts)
}

// HaveAllNodesVoted is a free data retrieval call binding the contract method 0xf57077d8.
//
// Solidity: function haveAllNodesVoted() view returns(bool)
func (_ContractExtender *ContractExtenderCaller) HaveAllNodesVoted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "haveAllNodesVoted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HaveAllNodesVoted is a free data retrieval call binding the contract method 0xf57077d8.
//
// Solidity: function haveAllNodesVoted() view returns(bool)
func (_ContractExtender *ContractExtenderSession) HaveAllNodesVoted() (bool, error) {
	return _ContractExtender.Contract.HaveAllNodesVoted(&_ContractExtender.CallOpts)
}

// HaveAllNodesVoted is a free data retrieval call binding the contract method 0xf57077d8.
//
// Solidity: function haveAllNodesVoted() view returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) HaveAllNodesVoted() (bool, error) {
	return _ContractExtender.Contract.HaveAllNodesVoted(&_ContractExtender.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_ContractExtender *ContractExtenderCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "isFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_ContractExtender *ContractExtenderSession) IsFinished() (bool, error) {
	return _ContractExtender.Contract.IsFinished(&_ContractExtender.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) IsFinished() (bool, error) {
	return _ContractExtender.Contract.IsFinished(&_ContractExtender.CallOpts)
}

// SharedDataHash is a free data retrieval call binding the contract method 0x88f520a0.
//
// Solidity: function sharedDataHash() view returns(string)
func (_ContractExtender *ContractExtenderCaller) SharedDataHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "sharedDataHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SharedDataHash is a free data retrieval call binding the contract method 0x88f520a0.
//
// Solidity: function sharedDataHash() view returns(string)
func (_ContractExtender *ContractExtenderSession) SharedDataHash() (string, error) {
	return _ContractExtender.Contract.SharedDataHash(&_ContractExtender.CallOpts)
}

// SharedDataHash is a free data retrieval call binding the contract method 0x88f520a0.
//
// Solidity: function sharedDataHash() view returns(string)
func (_ContractExtender *ContractExtenderCallerSession) SharedDataHash() (string, error) {
	return _ContractExtender.Contract.SharedDataHash(&_ContractExtender.CallOpts)
}

// TargetRecipientPTMKey is a free data retrieval call binding the contract method 0xe5af0f30.
//
// Solidity: function targetRecipientPTMKey() view returns(string)
func (_ContractExtender *ContractExtenderCaller) TargetRecipientPTMKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "targetRecipientPTMKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TargetRecipientPTMKey is a free data retrieval call binding the contract method 0xe5af0f30.
//
// Solidity: function targetRecipientPTMKey() view returns(string)
func (_ContractExtender *ContractExtenderSession) TargetRecipientPTMKey() (string, error) {
	return _ContractExtender.Contract.TargetRecipientPTMKey(&_ContractExtender.CallOpts)
}

// TargetRecipientPTMKey is a free data retrieval call binding the contract method 0xe5af0f30.
//
// Solidity: function targetRecipientPTMKey() view returns(string)
func (_ContractExtender *ContractExtenderCallerSession) TargetRecipientPTMKey() (string, error) {
	return _ContractExtender.Contract.TargetRecipientPTMKey(&_ContractExtender.CallOpts)
}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x38527727.
//
// Solidity: function totalNumberOfVoters() view returns(uint256)
func (_ContractExtender *ContractExtenderCaller) TotalNumberOfVoters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "totalNumberOfVoters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x38527727.
//
// Solidity: function totalNumberOfVoters() view returns(uint256)
func (_ContractExtender *ContractExtenderSession) TotalNumberOfVoters() (*big.Int, error) {
	return _ContractExtender.Contract.TotalNumberOfVoters(&_ContractExtender.CallOpts)
}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x38527727.
//
// Solidity: function totalNumberOfVoters() view returns(uint256)
func (_ContractExtender *ContractExtenderCallerSession) TotalNumberOfVoters() (*big.Int, error) {
	return _ContractExtender.Contract.TotalNumberOfVoters(&_ContractExtender.CallOpts)
}

// VoteOutcome is a free data retrieval call binding the contract method 0xb5da45bb.
//
// Solidity: function voteOutcome() view returns(bool)
func (_ContractExtender *ContractExtenderCaller) VoteOutcome(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "voteOutcome")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VoteOutcome is a free data retrieval call binding the contract method 0xb5da45bb.
//
// Solidity: function voteOutcome() view returns(bool)
func (_ContractExtender *ContractExtenderSession) VoteOutcome() (bool, error) {
	return _ContractExtender.Contract.VoteOutcome(&_ContractExtender.CallOpts)
}

// VoteOutcome is a free data retrieval call binding the contract method 0xb5da45bb.
//
// Solidity: function voteOutcome() view returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) VoteOutcome() (bool, error) {
	return _ContractExtender.Contract.VoteOutcome(&_ContractExtender.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(bool)
func (_ContractExtender *ContractExtenderCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(bool)
func (_ContractExtender *ContractExtenderSession) Votes(arg0 common.Address) (bool, error) {
	return _ContractExtender.Contract.Votes(&_ContractExtender.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(bool)
func (_ContractExtender *ContractExtenderCallerSession) Votes(arg0 common.Address) (bool, error) {
	return _ContractExtender.Contract.Votes(&_ContractExtender.CallOpts, arg0)
}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote(uint256 ) view returns(address)
func (_ContractExtender *ContractExtenderCaller) WalletAddressesToVote(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractExtender.contract.Call(opts, &out, "walletAddressesToVote", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote(uint256 ) view returns(address)
func (_ContractExtender *ContractExtenderSession) WalletAddressesToVote(arg0 *big.Int) (common.Address, error) {
	return _ContractExtender.Contract.WalletAddressesToVote(&_ContractExtender.CallOpts, arg0)
}

// WalletAddressesToVote is a free data retrieval call binding the contract method 0x79d41b8f.
//
// Solidity: function walletAddressesToVote(uint256 ) view returns(address)
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

// ContractExtenderAllNodesHaveAcceptedIterator is returned from FilterAllNodesHaveAccepted and is used to iterate over the raw logs and unpacked data for AllNodesHaveAccepted events raised by the ContractExtender contract.
type ContractExtenderAllNodesHaveAcceptedIterator struct {
	Event *ContractExtenderAllNodesHaveAccepted // Event containing the contract specifics and raw log

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
func (it *ContractExtenderAllNodesHaveAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtenderAllNodesHaveAccepted)
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
		it.Event = new(ContractExtenderAllNodesHaveAccepted)
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
func (it *ContractExtenderAllNodesHaveAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtenderAllNodesHaveAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtenderAllNodesHaveAccepted represents a AllNodesHaveAccepted event raised by the ContractExtender contract.
type ContractExtenderAllNodesHaveAccepted struct {
	Outcome bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAllNodesHaveAccepted is a free log retrieval operation binding the contract event 0xf20540914db019dd7c8d05ed165316a58d1583642772ac46f3d0c29b8644bd36.
//
// Solidity: event AllNodesHaveAccepted(bool outcome)
func (_ContractExtender *ContractExtenderFilterer) FilterAllNodesHaveAccepted(opts *bind.FilterOpts) (*ContractExtenderAllNodesHaveAcceptedIterator, error) {

	logs, sub, err := _ContractExtender.contract.FilterLogs(opts, "AllNodesHaveAccepted")
	if err != nil {
		return nil, err
	}
	return &ContractExtenderAllNodesHaveAcceptedIterator{contract: _ContractExtender.contract, event: "AllNodesHaveAccepted", logs: logs, sub: sub}, nil
}

var AllNodesHaveAcceptedTopicHash = "0xf20540914db019dd7c8d05ed165316a58d1583642772ac46f3d0c29b8644bd36"

// WatchAllNodesHaveAccepted is a free log subscription operation binding the contract event 0xf20540914db019dd7c8d05ed165316a58d1583642772ac46f3d0c29b8644bd36.
//
// Solidity: event AllNodesHaveAccepted(bool outcome)
func (_ContractExtender *ContractExtenderFilterer) WatchAllNodesHaveAccepted(opts *bind.WatchOpts, sink chan<- *ContractExtenderAllNodesHaveAccepted) (event.Subscription, error) {

	logs, sub, err := _ContractExtender.contract.WatchLogs(opts, "AllNodesHaveAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtenderAllNodesHaveAccepted)
				if err := _ContractExtender.contract.UnpackLog(event, "AllNodesHaveAccepted", log); err != nil {
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

// ParseAllNodesHaveAccepted is a log parse operation binding the contract event 0xf20540914db019dd7c8d05ed165316a58d1583642772ac46f3d0c29b8644bd36.
//
// Solidity: event AllNodesHaveAccepted(bool outcome)
func (_ContractExtender *ContractExtenderFilterer) ParseAllNodesHaveAccepted(log types.Log) (*ContractExtenderAllNodesHaveAccepted, error) {
	event := new(ContractExtenderAllNodesHaveAccepted)
	if err := _ContractExtender.contract.UnpackLog(event, "AllNodesHaveAccepted", log); err != nil {
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

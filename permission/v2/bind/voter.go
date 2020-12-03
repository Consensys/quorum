// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permission

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

// VoterManagerABI is the input ABI used to generate the binding from.
const VoterManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOpDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"deleteVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authOrg\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"},{\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"processVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authOrg\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"addVotingItem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VotingItemAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VoteProcessed\",\"type\":\"event\"}]"

var VoterManagerParsedABI, _ = abi.JSON(strings.NewReader(VoterManagerABI))

// VoterManagerBin is the compiled bytecode used for deploying new contracts.
var VoterManagerBin = "0x6080604052600060035534801561001557600080fd5b50604051602080611fe48339810180604052602081101561003557600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055611f7d806100676000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063014e6acc1461005c5780635607395b146101c857806359cbd6fe14610241578063b0213864146102b8578063e98ac22d14610349575b600080fd5b6100ca6004803603602081101561007257600080fd5b810190602081018135600160201b81111561008c57600080fd5b82018360208201111561009e57600080fd5b803590602001918460018302840111600160201b831117156100bf57600080fd5b509092509050610466565b604051808060200180602001856001600160a01b03166001600160a01b03168152602001848152602001838103835287818151815260200191508051906020019080838360005b83811015610129578181015183820152602001610111565b50505050905090810190601f1680156101565780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b83811015610189578181015183820152602001610171565b50505050905090810190601f1680156101b65780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b61023f600480360360408110156101de57600080fd5b810190602081018135600160201b8111156101f857600080fd5b82018360208201111561020a57600080fd5b803590602001918460018302840111600160201b8311171561022b57600080fd5b9193509150356001600160a01b0316610740565b005b61023f6004803603604081101561025757600080fd5b810190602081018135600160201b81111561027157600080fd5b82018360208201111561028357600080fd5b803590602001918460018302840111600160201b831117156102a457600080fd5b9193509150356001600160a01b0316610f06565b610335600480360360608110156102ce57600080fd5b810190602081018135600160201b8111156102e857600080fd5b8201836020820111156102fa57600080fd5b803590602001918460018302840111600160201b8311171561031b57600080fd5b91935091506001600160a01b0381351690602001356111e8565b604080519115158252519081900360200190f35b61023f600480360360a081101561035f57600080fd5b810190602081018135600160201b81111561037957600080fd5b82018360208201111561038b57600080fd5b803590602001918460018302840111600160201b831117156103ac57600080fd5b919390929091602081019035600160201b8111156103c957600080fd5b8201836020820111156103db57600080fd5b803590602001918460018302840111600160201b831117156103fc57600080fd5b919390929091602081019035600160201b81111561041957600080fd5b82018360208201111561042b57600080fd5b803590602001918460018302840111600160201b8311171561044c57600080fd5b91935091506001600160a01b0381351690602001356116f8565b6060806000806000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b1580156104b957600080fd5b505afa1580156104cd573d6000803e3d6000fd5b505050506040513d60208110156104e357600080fd5b50516001600160a01b031633146105385760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b600061057987878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611afd92505050565b905060018181548110151561058a57fe5b90600052602060002090600b02016004016000016001828154811015156105ad57fe5b90600052602060002090600b02016004016001016001838154811015156105d057fe5b600091825260209091206006600b909202010154600180546001600160a01b0390921691859081106105fe57fe5b60009182526020918290206007600b909202010154845460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815291928691908301828280156106995780601f1061066e57610100808354040283529160200191610699565b820191906000526020600020905b81548152906001019060200180831161067c57829003601f168201915b5050865460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959950889450925084019050828280156107275780601f106106fc57610100808354040283529160200191610727565b820191906000526020600020905b81548152906001019060200180831161070a57829003601f168201915b5050505050925094509450945094505092959194509250565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561078d57600080fd5b505afa1580156107a1573d6000803e3d6000fd5b505050506040513d60208110156107b757600080fd5b50516001600160a01b0316331461080c5760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b60026000848460405160200180806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505093505050506040516020818303038152906040528051906020012081526020019081526020016000205460001415610b78576003805460010190819055604080516020808201908152918101859052600291600091879187918190606001848480828437600081840152601f19601f8201169050808301925050509350505050604051602081830303815290604052805190602001208152602001908152602001600020819055506000600180548091906001016109069190611cdd565b9050838360018381548110151561091957fe5b6000918252602090912061093393600b9092020191611d0e565b506001808281548110151561094457fe5b90600052602060002090600b0201600101819055506001808281548110151561096957fe5b90600052602060002090600b020160020181905550600060018281548110151561098f57fe5b90600052602060002090600b020160030181905550604051806020016040528060008152506001828154811015156109c357fe5b90600052602060002090600b020160040160000190805190602001906109ea929190611d8c565b506040805160208101909152600081526001805483908110610a0857fe5b90600052602060002090600b02016004016001019080519060200190610a2f929190611d8c565b506000600182815481101515610a4157fe5b600091825260208220600b919091020160060180546001600160a01b0319166001600160a01b0393909316929092179091556001805483908110610a8157fe5b600091825260209091206007600b9092020101556001805482908110610aa357fe5b90600052602060002090600b020160010154600182815481101515610ac457fe5b600091825260208083206001600160a01b03871684526009600b9093020191909101905260409020556001805482908110610afb57fe5b60009182526020808320604080518082019091526001600160a01b0387811682526001828501818152600b969096029093016008018054938401815586529290942093519301805492516001600160a01b03199093169390911692909217600160a01b60ff021916600160a01b9115159190910217905550610e90565b6000610bb984848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611afd92505050565b9050600181815481101515610bca57fe5b600091825260208083206001600160a01b03861684526009600b9093020191909101905260409020541515610d2c576001805482908110610c0757fe5b600091825260209091206001600b909202018101805482019055805482908110610c2d57fe5b90600052602060002090600b020160010154600182815481101515610c4e57fe5b600091825260208083206001600160a01b03871684526009600b9093020191909101905260409020556001805482908110610c8557fe5b60009182526020808320604080518082019091526001600160a01b0387811682526001828501818152600b96909602909301600801805480850182559087529390952090519201805493516001600160a01b03199094169290941691909117600160a01b60ff021916600160a01b9215159290920291909117909155805482908110610d0d57fe5b600091825260209091206002600b909202010180546001019055610e8e565b6000610d6f85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250611ba5915050565b9050600182815481101515610d8057fe5b90600052602060002090600b020160080181815481101515610d9e57fe5b600091825260209091200154600160a01b900460ff16151560011415610e0e5760408051600160e51b62461bcd02815260206004820152600f60248201527f616c7265616479206120766f7465720000000000000000000000000000000000604482015290519081900360640190fd5b60018083815481101515610e1e57fe5b90600052602060002090600b020160080182815481101515610e3c57fe5b60009182526020909120018054911515600160a01b02600160a01b60ff02199092169190911790556001805483908110610e7257fe5b600091825260209091206002600b909202010180546001019055505b505b604080516001600160a01b03831660208201528181529081018390527f424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574908490849084908060608101858580828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b158015610f5357600080fd5b505afa158015610f67573d6000803e3d6000fd5b505050506040513d6020811015610f7d57600080fd5b50516001600160a01b03163314610fd25760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b82828080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250849250611016915083905082611bf7565b15156001146110645760408051600160e51b62461bcd02815260206004820152600f6024820152600160891b6e36bab9ba1031329030903b37ba32b902604482015290519081900360640190fd5b60006110a586868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611afd92505050565b905060006110ea87878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250899250611ba5915050565b90506001828154811015156110fb57fe5b6000918252602082206002600b90920201018054600019019055600180548490811061112357fe5b90600052602060002090600b02016008018281548110151561114157fe5b9060005260206000200160000160146101000a81548160ff0219169083151502179055507f654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b68787876040518080602001836001600160a01b03166001600160a01b031681526020018281038252858582818152602001925080828437600083820152604051601f909101601f1916909201829003965090945050505050a150505050505050565b60008060009054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561123757600080fd5b505afa15801561124b573d6000803e3d6000fd5b505050506040513d602081101561126157600080fd5b50516001600160a01b031633146112b65760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506112fa915083905082611bf7565b15156001146113485760408051600160e51b62461bcd02815260206004820152600f6024820152600160891b6e36bab9ba1031329030903b37ba32b902604482015290519081900360640190fd5b61138987878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250611ca7915050565b15156001146113e25760408051600160e51b62461bcd02815260206004820152601260248201527f6e6f7468696e6720746f20617070726f76650000000000000000000000000000604482015290519081900360640190fd5b600061142388888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611afd92505050565b905060018181548110151561143457fe5b60009182526020808320848452600a600b9093020191909101815260408083206001600160a01b038a16845290915290205460ff161515600114156114c35760408051600160e51b62461bcd02815260206004820152601260248201527f63616e6e6f7420646f75626c6520766f74650000000000000000000000000000604482015290519081900360640190fd5b60018054829081106114d157fe5b600091825260209091206003600b90920201018054600190810190915580548190839081106114fc57fe5b60009182526020808320858452600b92909202909101600a01815260408083206001600160a01b038b168452825291829020805460ff19169315159390931790925580518281529182018990527f87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508918a918a919081908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a160026001828154811015156115ac57fe5b90600052602060002090600b0201600201548115156115c757fe5b046001828154811015156115d757fe5b90600052602060002090600b02016003015411156116e857604080516020810190915260008152600180548390811061160c57fe5b90600052602060002090600b02016004016000019080519060200190611633929190611d8c565b50604080516020810190915260008152600180548390811061165157fe5b90600052602060002090600b02016004016001019080519060200190611678929190611d8c565b50600060018281548110151561168a57fe5b600091825260208220600b919091020160060180546001600160a01b0319166001600160a01b03939093169290921790915560018054839081106116ca57fe5b600091825260209091206007600b90920201015550600192506116ee565b60009350505b5050949350505050565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561174557600080fd5b505afa158015611759573d6000803e3d6000fd5b505050506040513d602081101561176f57600080fd5b50516001600160a01b031633146117c45760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b61180388888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250611ca7915050565b151561184357604051600160e51b62461bcd028152600401808060200182810382526034815260200180611f1e6034913960400191505060405180910390fd5b600061188489898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611afd92505050565b9050868660018381548110151561189757fe5b90600052602060002090600b020160040160000191906118b8929190611d0e565b5084846001838154811015156118ca57fe5b90600052602060002090600b020160040160010191906118eb929190611d0e565b50826001828154811015156118fc57fe5b90600052602060002090600b020160040160020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160018281548110151561194457fe5b6000918252602082206007600b9092020101919091555b600180548390811061196957fe5b90600052602060002090600b020160080180549050811015611a6b57600180548390811061199357fe5b90600052602060002090600b0201600801818154811015156119b157fe5b600091825260209091200154600160a01b900460ff1615611a635760006001838154811015156119dd57fe5b90600052602060002090600b0201600a0160008481526020019081526020016000206000600185815481101515611a1057fe5b90600052602060002090600b020160080184815481101515611a2e57fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790555b60010161195b565b506000600182815481101515611a7d57fe5b90600052602060002090600b0201600301819055507f5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3898960405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a1505050505050505050565b6000600160026000846040516020018080602001828103825283818151815260200191508051906020019080838360005b83811015611b46578181015183820152602001611b2e565b50505050905090810190601f168015611b735780820380516001836020036101000a031916815260200191505b509250505060405160208183030381529060405280519060200120815260200190815260200160002054039050919050565b600080611bb184611afd565b905060018082815481101515611bc357fe5b600091825260208083206001600160a01b03881684526009600b909302019190910190526040902054039150505b92915050565b600080611c0384611afd565b9050600181815481101515611c1457fe5b600091825260208083206001600160a01b03871684526009600b9093020191909101905260409020541515611c4d576000915050611bf1565b6000611c598585611ba5565b9050600182815481101515611c6a57fe5b90600052602060002090600b020160080181815481101515611c8857fe5b600091825260209091200154600160a01b900460ff1695945050505050565b6000816001611cb585611afd565b81548110611cbf57fe5b90600052602060002090600b02016004016003015414905092915050565b815481835581811115611d0957600b0281600b028360005260206000209182019101611d099190611dfa565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611d4f5782800160ff19823516178555611d7c565b82800160010185558215611d7c579182015b82811115611d7c578235825591602001919060010190611d61565b50611d88929150611e7f565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611dcd57805160ff1916838001178555611d7c565b82800160010185558215611d7c579182015b82811115611d7c578251825591602001919060010190611ddf565b611e7c91905b80821115611d88576000611e148282611e99565b60006001830181905560028301819055600383018190556004830190611e3a8282611e99565b611e48600183016000611e99565b506002810180546001600160a01b031916905560006003909101819055611e73906008840190611ee0565b50600b01611e00565b90565b611e7c91905b80821115611d885760008155600101611e85565b50805460018160011615610100020316600290046000825580601f10611ebf5750611edd565b601f016020900490600052602060002090810190611edd9190611e7f565b50565b5080546000825590600052602060002090810190611edd9190611e7c91905b80821115611d885780546001600160a81b0319168155600101611eff56fe6974656d732070656e64696e6720666f7220617070726f76616c2e206e6577206974656d2063616e6e6f74206265206164646564a165627a7a723058207dc9a68c6931494f043f7420dfa8288733d9a3a676ed30b4ac8e9cc704bd928b0029"

// DeployVoterManager deploys a new Ethereum contract, binding an instance of VoterManager to it.
func DeployVoterManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permUpgradable common.Address) (common.Address, *types.Transaction, *VoterManager, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VoterManagerBin), backend, _permUpgradable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VoterManager{VoterManagerCaller: VoterManagerCaller{contract: contract}, VoterManagerTransactor: VoterManagerTransactor{contract: contract}, VoterManagerFilterer: VoterManagerFilterer{contract: contract}}, nil
}

// VoterManager is an auto generated Go binding around an Ethereum contract.
type VoterManager struct {
	VoterManagerCaller     // Read-only binding to the contract
	VoterManagerTransactor // Write-only binding to the contract
	VoterManagerFilterer   // Log filterer for contract events
}

// VoterManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterManagerSession struct {
	Contract     *VoterManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterManagerCallerSession struct {
	Contract *VoterManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VoterManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterManagerTransactorSession struct {
	Contract     *VoterManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VoterManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterManagerRaw struct {
	Contract *VoterManager // Generic contract binding to access the raw methods on
}

// VoterManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterManagerCallerRaw struct {
	Contract *VoterManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VoterManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterManagerTransactorRaw struct {
	Contract *VoterManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterManager creates a new instance of VoterManager, bound to a specific deployed contract.
func NewVoterManager(address common.Address, backend bind.ContractBackend) (*VoterManager, error) {
	contract, err := bindVoterManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterManager{VoterManagerCaller: VoterManagerCaller{contract: contract}, VoterManagerTransactor: VoterManagerTransactor{contract: contract}, VoterManagerFilterer: VoterManagerFilterer{contract: contract}}, nil
}

// NewVoterManagerCaller creates a new read-only instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerCaller(address common.Address, caller bind.ContractCaller) (*VoterManagerCaller, error) {
	contract, err := bindVoterManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterManagerCaller{contract: contract}, nil
}

// NewVoterManagerTransactor creates a new write-only instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterManagerTransactor, error) {
	contract, err := bindVoterManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterManagerTransactor{contract: contract}, nil
}

// NewVoterManagerFilterer creates a new log filterer instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterManagerFilterer, error) {
	contract, err := bindVoterManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterManagerFilterer{contract: contract}, nil
}

// bindVoterManager binds a generic wrapper to an already deployed contract.
func bindVoterManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterManager *VoterManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterManager.Contract.VoterManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterManager *VoterManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterManager.Contract.VoterManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterManager *VoterManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterManager.Contract.VoterManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterManager *VoterManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterManager *VoterManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterManager *VoterManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterManager.Contract.contract.Transact(opts, method, params...)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerCaller) GetPendingOpDetails(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _VoterManager.contract.Call(opts, out, "getPendingOpDetails", _orgId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _VoterManager.Contract.GetPendingOpDetails(&_VoterManager.CallOpts, _orgId)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerCallerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _VoterManager.Contract.GetPendingOpDetails(&_VoterManager.CallOpts, _orgId)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactor) AddVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "addVoter", _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactorSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_VoterManager *VoterManagerTransactor) AddVotingItem(opts *bind.TransactOpts, _authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "addVotingItem", _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_VoterManager *VoterManagerSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVotingItem(&_VoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_VoterManager *VoterManagerTransactorSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVotingItem(&_VoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactor) DeleteVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "deleteVoter", _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.DeleteVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_VoterManager *VoterManagerTransactorSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.DeleteVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_VoterManager *VoterManagerTransactor) ProcessVote(opts *bind.TransactOpts, _authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "processVote", _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_VoterManager *VoterManagerSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.ProcessVote(&_VoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_VoterManager *VoterManagerTransactorSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.ProcessVote(&_VoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// VoterManagerVoteProcessedIterator is returned from FilterVoteProcessed and is used to iterate over the raw logs and unpacked data for VoteProcessed events raised by the VoterManager contract.
type VoterManagerVoteProcessedIterator struct {
	Event *VoterManagerVoteProcessed // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoteProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoteProcessed)
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
		it.Event = new(VoterManagerVoteProcessed)
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
func (it *VoterManagerVoteProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoteProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoteProcessed represents a VoteProcessed event raised by the VoterManager contract.
type VoterManagerVoteProcessed struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteProcessed is a free log retrieval operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_VoterManager *VoterManagerFilterer) FilterVoteProcessed(opts *bind.FilterOpts) (*VoterManagerVoteProcessedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoteProcessedIterator{contract: _VoterManager.contract, event: "VoteProcessed", logs: logs, sub: sub}, nil
}

var VoteProcessedTopicHash = "0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508"

// WatchVoteProcessed is a free log subscription operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_VoterManager *VoterManagerFilterer) WatchVoteProcessed(opts *bind.WatchOpts, sink chan<- *VoterManagerVoteProcessed) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoteProcessed)
				if err := _VoterManager.contract.UnpackLog(event, "VoteProcessed", log); err != nil {
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

// ParseVoteProcessed is a log parse operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_VoterManager *VoterManagerFilterer) ParseVoteProcessed(log types.Log) (*VoterManagerVoteProcessed, error) {
	event := new(VoterManagerVoteProcessed)
	if err := _VoterManager.contract.UnpackLog(event, "VoteProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VoterManagerVoterAddedIterator is returned from FilterVoterAdded and is used to iterate over the raw logs and unpacked data for VoterAdded events raised by the VoterManager contract.
type VoterManagerVoterAddedIterator struct {
	Event *VoterManagerVoterAdded // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoterAdded)
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
		it.Event = new(VoterManagerVoterAdded)
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
func (it *VoterManagerVoterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoterAdded represents a VoterAdded event raised by the VoterManager contract.
type VoterManagerVoterAdded struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterAdded is a free log retrieval operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) FilterVoterAdded(opts *bind.FilterOpts) (*VoterManagerVoterAddedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoterAddedIterator{contract: _VoterManager.contract, event: "VoterAdded", logs: logs, sub: sub}, nil
}

var VoterAddedTopicHash = "0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574"

// WatchVoterAdded is a free log subscription operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) WatchVoterAdded(opts *bind.WatchOpts, sink chan<- *VoterManagerVoterAdded) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoterAdded)
				if err := _VoterManager.contract.UnpackLog(event, "VoterAdded", log); err != nil {
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

// ParseVoterAdded is a log parse operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) ParseVoterAdded(log types.Log) (*VoterManagerVoterAdded, error) {
	event := new(VoterManagerVoterAdded)
	if err := _VoterManager.contract.UnpackLog(event, "VoterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VoterManagerVoterDeletedIterator is returned from FilterVoterDeleted and is used to iterate over the raw logs and unpacked data for VoterDeleted events raised by the VoterManager contract.
type VoterManagerVoterDeletedIterator struct {
	Event *VoterManagerVoterDeleted // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoterDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoterDeleted)
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
		it.Event = new(VoterManagerVoterDeleted)
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
func (it *VoterManagerVoterDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoterDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoterDeleted represents a VoterDeleted event raised by the VoterManager contract.
type VoterManagerVoterDeleted struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterDeleted is a free log retrieval operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) FilterVoterDeleted(opts *bind.FilterOpts) (*VoterManagerVoterDeletedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoterDeletedIterator{contract: _VoterManager.contract, event: "VoterDeleted", logs: logs, sub: sub}, nil
}

var VoterDeletedTopicHash = "0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6"

// WatchVoterDeleted is a free log subscription operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) WatchVoterDeleted(opts *bind.WatchOpts, sink chan<- *VoterManagerVoterDeleted) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoterDeleted)
				if err := _VoterManager.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
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

// ParseVoterDeleted is a log parse operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_VoterManager *VoterManagerFilterer) ParseVoterDeleted(log types.Log) (*VoterManagerVoterDeleted, error) {
	event := new(VoterManagerVoterDeleted)
	if err := _VoterManager.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VoterManagerVotingItemAddedIterator is returned from FilterVotingItemAdded and is used to iterate over the raw logs and unpacked data for VotingItemAdded events raised by the VoterManager contract.
type VoterManagerVotingItemAddedIterator struct {
	Event *VoterManagerVotingItemAdded // Event containing the contract specifics and raw log

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
func (it *VoterManagerVotingItemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVotingItemAdded)
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
		it.Event = new(VoterManagerVotingItemAdded)
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
func (it *VoterManagerVotingItemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVotingItemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVotingItemAdded represents a VotingItemAdded event raised by the VoterManager contract.
type VoterManagerVotingItemAdded struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVotingItemAdded is a free log retrieval operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_VoterManager *VoterManagerFilterer) FilterVotingItemAdded(opts *bind.FilterOpts) (*VoterManagerVotingItemAddedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVotingItemAddedIterator{contract: _VoterManager.contract, event: "VotingItemAdded", logs: logs, sub: sub}, nil
}

var VotingItemAddedTopicHash = "0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3"

// WatchVotingItemAdded is a free log subscription operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_VoterManager *VoterManagerFilterer) WatchVotingItemAdded(opts *bind.WatchOpts, sink chan<- *VoterManagerVotingItemAdded) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVotingItemAdded)
				if err := _VoterManager.contract.UnpackLog(event, "VotingItemAdded", log); err != nil {
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

// ParseVotingItemAdded is a log parse operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_VoterManager *VoterManagerFilterer) ParseVotingItemAdded(log types.Log) (*VoterManagerVotingItemAdded, error) {
	event := new(VoterManagerVotingItemAdded)
	if err := _VoterManager.contract.UnpackLog(event, "VotingItemAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

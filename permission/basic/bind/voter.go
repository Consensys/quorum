// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

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

// VoterManagerBin is the compiled bytecode used for deploying new contracts.
var VoterManagerBin = "0x6080604052600060035534801561001557600080fd5b506040516020806120968339810180604052602081101561003557600080fd5b505160008054600160a060020a03909216600160a060020a031990921691909117905561202f806100676000396000f3fe608060405234801561001057600080fd5b506004361061005a5760e060020a6000350463014e6acc811461005f5780635607395b146101cd57806359cbd6fe14610248578063b0213864146102c1578063e98ac22d14610354575b600080fd5b6100cf6004803603602081101561007557600080fd5b81019060208101813564010000000081111561009057600080fd5b8201836020820111156100a257600080fd5b803590602001918460018302840111640100000000831117156100c457600080fd5b509092509050610477565b60405180806020018060200185600160a060020a0316600160a060020a03168152602001848152602001838103835287818151815260200191508051906020019080838360005b8381101561012e578181015183820152602001610116565b50505050905090810190601f16801561015b5780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b8381101561018e578181015183820152602001610176565b50505050905090810190601f1680156101bb5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b610246600480360360408110156101e357600080fd5b8101906020810181356401000000008111156101fe57600080fd5b82018360208201111561021057600080fd5b8035906020019184600183028401116401000000008311171561023257600080fd5b919350915035600160a060020a031661074e565b005b6102466004803603604081101561025e57600080fd5b81019060208101813564010000000081111561027957600080fd5b82018360208201111561028b57600080fd5b803590602001918460018302840111640100000000831117156102ad57600080fd5b919350915035600160a060020a0316610f63565b610340600480360360608110156102d757600080fd5b8101906020810181356401000000008111156102f257600080fd5b82018360208201111561030457600080fd5b8035906020019184600183028401116401000000008311171561032657600080fd5b9193509150600160a060020a03813516906020013561124d565b604080519115158252519081900360200190f35b610246600480360360a081101561036a57600080fd5b81019060208101813564010000000081111561038557600080fd5b82018360208201111561039757600080fd5b803590602001918460018302840111640100000000831117156103b957600080fd5b9193909290916020810190356401000000008111156103d757600080fd5b8201836020820111156103e957600080fd5b8035906020019184600183028401116401000000008311171561040b57600080fd5b91939092909160208101903564010000000081111561042957600080fd5b82018360208201111561043b57600080fd5b8035906020019184600183028401116401000000008311171561045d57600080fd5b9193509150600160a060020a038135169060200135611772565b6060806000806000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156104cd57600080fd5b505afa1580156104e1573d6000803e3d6000fd5b505050506040513d60208110156104f757600080fd5b5051600160a060020a03163314610546576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020611fb0833981519152604482015290519081900360640190fd5b600061058787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b7492505050565b905060018181548110151561059857fe5b90600052602060002090600b02016004016000016001828154811015156105bb57fe5b90600052602060002090600b02016004016001016001838154811015156105de57fe5b600091825260209091206006600b90920201015460018054600160a060020a03909216918590811061060c57fe5b60009182526020918290206007600b909202010154845460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815291928691908301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b5050865460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959950889450925084019050828280156107355780601f1061070a57610100808354040283529160200191610735565b820191906000526020600020905b81548152906001019060200180831161071857829003601f168201915b5050505050925094509450945094505092959194509250565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561079e57600080fd5b505afa1580156107b2573d6000803e3d6000fd5b505050506040513d60208110156107c857600080fd5b5051600160a060020a03163314610817576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020611fb0833981519152604482015290519081900360640190fd5b60026000848460405160200180806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505093505050506040516020818303038152906040528051906020012081526020019081526020016000205460001415610bac576003805460010190819055604080516020808201908152918101859052600291600091879187918190606001848480828437600081840152601f19601f8201169050808301925050509350505050604051602081830303815290604052805190602001208152602001908152602001600020819055506000600180548091906001016109119190611d54565b9050838360018381548110151561092457fe5b6000918252602090912061093e93600b9092020191611d85565b506001808281548110151561094f57fe5b90600052602060002090600b0201600101819055506001808281548110151561097457fe5b90600052602060002090600b020160020181905550600060018281548110151561099a57fe5b90600052602060002090600b02016003018190555060206040519081016040528060008152506001828154811015156109cf57fe5b90600052602060002090600b020160040160000190805190602001906109f6929190611e03565b506040805160208101909152600081526001805483908110610a1457fe5b90600052602060002090600b02016004016001019080519060200190610a3b929190611e03565b506000600182815481101515610a4d57fe5b600091825260208220600b9190910201600601805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0393909316929092179091556001805483908110610a9a57fe5b600091825260209091206007600b9092020101556001805482908110610abc57fe5b90600052602060002090600b020160010154600182815481101515610add57fe5b60009182526020808320600160a060020a03871684526009600b9093020191909101905260409020556001805482908110610b1457fe5b6000918252602080832060408051808201909152600160a060020a0387811682526001828501818152600b9690960290930160080180549384018155865292909420935193018054925173ffffffffffffffffffffffffffffffffffffffff19909316939091169290921774ff0000000000000000000000000000000000000000191660a060020a9115159190910217905550610eed565b6000610bed84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b7492505050565b9050600181815481101515610bfe57fe5b60009182526020808320600160a060020a03861684526009600b9093020191909101905260409020541515610d7b576001805482908110610c3b57fe5b600091825260209091206001600b909202018101805482019055805482908110610c6157fe5b90600052602060002090600b020160010154600182815481101515610c8257fe5b60009182526020808320600160a060020a03871684526009600b9093020191909101905260409020556001805482908110610cb957fe5b6000918252602080832060408051808201909152600160a060020a0387811682526001828501818152600b969096029093016008018054808501825590875293909520905192018054935173ffffffffffffffffffffffffffffffffffffffff19909416929094169190911774ff0000000000000000000000000000000000000000191660a060020a9215159290920291909117909155805482908110610d5c57fe5b600091825260209091206002600b909202010180546001019055610eeb565b6000610dbe85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250611c1c915050565b9050600182815481101515610dcf57fe5b90600052602060002090600b020160080181815481101515610ded57fe5b60009182526020909120015460a060020a900460ff16151560011415610e5d576040805160e560020a62461bcd02815260206004820152600f60248201527f616c7265616479206120766f7465720000000000000000000000000000000000604482015290519081900360640190fd5b60018083815481101515610e6d57fe5b90600052602060002090600b020160080182815481101515610e8b57fe5b6000918252602090912001805491151560a060020a0274ff0000000000000000000000000000000000000000199092169190911790556001805483908110610ecf57fe5b600091825260209091206002600b909202010180546001019055505b505b60408051600160a060020a03831660208201528181529081018390527f424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574908490849084908060608101858580828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610fb357600080fd5b505afa158015610fc7573d6000803e3d6000fd5b505050506040513d6020811015610fdd57600080fd5b5051600160a060020a0316331461102c576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020611fb0833981519152604482015290519081900360640190fd5b82828080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250849250611070915083905082611c6e565b15156001146110c9576040805160e560020a62461bcd02815260206004820152600f60248201527f6d757374206265206120766f7465720000000000000000000000000000000000604482015290519081900360640190fd5b600061110a86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b7492505050565b9050600061114f87878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250899250611c1c915050565b905060018281548110151561116057fe5b6000918252602082206002600b90920201018054600019019055600180548490811061118857fe5b90600052602060002090600b0201600801828154811015156111a657fe5b9060005260206000200160000160146101000a81548160ff0219169083151502179055507f654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6878787604051808060200183600160a060020a0316600160a060020a031681526020018281038252858582818152602001925080828437600083820152604051601f909101601f1916909201829003965090945050505050a150505050505050565b60008060009054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561129f57600080fd5b505afa1580156112b3573d6000803e3d6000fd5b505050506040513d60208110156112c957600080fd5b5051600160a060020a03163314611318576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020611fb0833981519152604482015290519081900360640190fd5b84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525086925061135c915083905082611c6e565b15156001146113b5576040805160e560020a62461bcd02815260206004820152600f60248201527f6d757374206265206120766f7465720000000000000000000000000000000000604482015290519081900360640190fd5b6113f687878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250611d1e915050565b151560011461144f576040805160e560020a62461bcd02815260206004820152601260248201527f6e6f7468696e6720746f20617070726f76650000000000000000000000000000604482015290519081900360640190fd5b600061149088888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b7492505050565b90506001818154811015156114a157fe5b60009182526020808320848452600a600b909302019190910181526040808320600160a060020a038a16845290915290205460ff16151560011415611530576040805160e560020a62461bcd02815260206004820152601260248201527f63616e6e6f7420646f75626c6520766f74650000000000000000000000000000604482015290519081900360640190fd5b600180548290811061153e57fe5b600091825260209091206003600b909202010180546001908101909155805481908390811061156957fe5b60009182526020808320858452600b92909202909101600a0181526040808320600160a060020a038b168452825291829020805460ff19169315159390931790925580518281529182018990527f87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508918a918a919081908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a1600260018281548110151561161957fe5b90600052602060002090600b02016002015481151561163457fe5b0460018281548110151561164457fe5b90600052602060002090600b020160030154111561176257604080516020810190915260008152600180548390811061167957fe5b90600052602060002090600b020160040160000190805190602001906116a0929190611e03565b5060408051602081019091526000815260018054839081106116be57fe5b90600052602060002090600b020160040160010190805190602001906116e5929190611e03565b5060006001828154811015156116f757fe5b600091825260208220600b9190910201600601805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039390931692909217909155600180548390811061174457fe5b600091825260209091206007600b9092020101555060019250611768565b60009350505b5050949350505050565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156117c257600080fd5b505afa1580156117d6573d6000803e3d6000fd5b505050506040513d60208110156117ec57600080fd5b5051600160a060020a0316331461183b576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020611fb0833981519152604482015290519081900360640190fd5b61187a88888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250611d1e915050565b15156118ba5760405160e560020a62461bcd028152600401808060200182810382526034815260200180611fd06034913960400191505060405180910390fd5b60006118fb89898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b7492505050565b9050868660018381548110151561190e57fe5b90600052602060002090600b0201600401600001919061192f929190611d85565b50848460018381548110151561194157fe5b90600052602060002090600b02016004016001019190611962929190611d85565b508260018281548110151561197357fe5b90600052602060002090600b020160040160020160006101000a815481600160a060020a030219169083600160a060020a03160217905550816001828154811015156119bb57fe5b6000918252602082206007600b9092020101919091555b60018054839081106119e057fe5b90600052602060002090600b020160080180549050811015611ae2576001805483908110611a0a57fe5b90600052602060002090600b020160080181815481101515611a2857fe5b60009182526020909120015460a060020a900460ff1615611ada576000600183815481101515611a5457fe5b90600052602060002090600b0201600a0160008481526020019081526020016000206000600185815481101515611a8757fe5b90600052602060002090600b020160080184815481101515611aa557fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff19169115159190911790555b6001016119d2565b506000600182815481101515611af457fe5b90600052602060002090600b0201600301819055507f5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3898960405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a1505050505050505050565b6000600160026000846040516020018080602001828103825283818151815260200191508051906020019080838360005b83811015611bbd578181015183820152602001611ba5565b50505050905090810190601f168015611bea5780820380516001836020036101000a031916815260200191505b509250505060405160208183030381529060405280519060200120815260200190815260200160002054039050919050565b600080611c2884611b74565b905060018082815481101515611c3a57fe5b60009182526020808320600160a060020a03881684526009600b909302019190910190526040902054039150505b92915050565b600080611c7a84611b74565b9050600181815481101515611c8b57fe5b60009182526020808320600160a060020a03871684526009600b9093020191909101905260409020541515611cc4576000915050611c68565b6000611cd08585611c1c565b9050600182815481101515611ce157fe5b90600052602060002090600b020160080181815481101515611cff57fe5b60009182526020909120015460a060020a900460ff1695945050505050565b6000816001611d2c85611b74565b81548110611d3657fe5b90600052602060002090600b02016004016003015414905092915050565b815481835581811115611d8057600b0281600b028360005260206000209182019101611d809190611e71565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611dc65782800160ff19823516178555611df3565b82800160010185558215611df3579182015b82811115611df3578235825591602001919060010190611dd8565b50611dff929150611f03565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611e4457805160ff1916838001178555611df3565b82800160010185558215611df3579182015b82811115611df3578251825591602001919060010190611e56565b611f0091905b80821115611dff576000611e8b8282611f1d565b60006001830181905560028301819055600383018190556004830190611eb18282611f1d565b611ebf600183016000611f1d565b5060028101805473ffffffffffffffffffffffffffffffffffffffff1916905560006003909101819055611ef7906008840190611f64565b50600b01611e77565b90565b611f0091905b80821115611dff5760008155600101611f09565b50805460018160011615610100020316600290046000825580601f10611f435750611f61565b601f016020900490600052602060002090810190611f619190611f03565b50565b5080546000825590600052602060002090810190611f619190611f0091905b80821115611dff57805474ffffffffffffffffffffffffffffffffffffffffff19168155600101611f8356fe696e76616c69642063616c6c65720000000000000000000000000000000000006974656d732070656e64696e6720666f7220617070726f76616c2e206e6577206974656d2063616e6e6f74206265206164646564a165627a7a72305820853b884fe72210949b9c5307ef60e47eb2f01504e26db430be567e00ee7527a70029"

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

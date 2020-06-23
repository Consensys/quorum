// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eea

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

// EeaVoterManagerABI is the input ABI used to generate the binding from.
const EeaVoterManagerABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOpDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"deleteVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_authOrg\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"processVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_authOrg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"addVotingItem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VotingItemAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VoteProcessed\",\"type\":\"event\"}]"

// EeaVoterManagerBin is the compiled bytecode used for deploying new contracts.
var EeaVoterManagerBin = "0x6080604052600060035534801561001557600080fd5b50604051611f3b380380611f3b8339818101604052602081101561003857600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055611ed18061006a6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063014e6acc1461005c5780635607395b146101c857806359cbd6fe14610241578063b0213864146102b8578063e98ac22d14610349575b600080fd5b6100ca6004803603602081101561007257600080fd5b810190602081018135600160201b81111561008c57600080fd5b82018360208201111561009e57600080fd5b803590602001918460018302840111600160201b831117156100bf57600080fd5b509092509050610466565b604051808060200180602001856001600160a01b03166001600160a01b03168152602001848152602001838103835287818151815260200191508051906020019080838360005b83811015610129578181015183820152602001610111565b50505050905090810190601f1680156101565780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b83811015610189578181015183820152602001610171565b50505050905090810190601f1680156101b65780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b61023f600480360360408110156101de57600080fd5b810190602081018135600160201b8111156101f857600080fd5b82018360208201111561020a57600080fd5b803590602001918460018302840111600160201b8311171561022b57600080fd5b9193509150356001600160a01b0316610734565b005b61023f6004803603604081101561025757600080fd5b810190602081018135600160201b81111561027157600080fd5b82018360208201111561028357600080fd5b803590602001918460018302840111600160201b831117156102a457600080fd5b9193509150356001600160a01b0316610ebd565b610335600480360360608110156102ce57600080fd5b810190602081018135600160201b8111156102e857600080fd5b8201836020820111156102fa57600080fd5b803590602001918460018302840111600160201b8311171561031b57600080fd5b91935091506001600160a01b03813516906020013561118f565b604080519115158252519081900360200190f35b61023f600480360360a081101561035f57600080fd5b810190602081018135600160201b81111561037957600080fd5b82018360208201111561038b57600080fd5b803590602001918460018302840111600160201b831117156103ac57600080fd5b919390929091602081019035600160201b8111156103c957600080fd5b8201836020820111156103db57600080fd5b803590602001918460018302840111600160201b831117156103fc57600080fd5b919390929091602081019035600160201b81111561041957600080fd5b82018360208201111561042b57600080fd5b803590602001918460018302840111600160201b8311171561044c57600080fd5b91935091506001600160a01b03813516906020013561166c565b6060806000806000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b1580156104b957600080fd5b505afa1580156104cd573d6000803e3d6000fd5b505050506040513d60208110156104e357600080fd5b50516001600160a01b03163314610532576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b604482015290519081900360640190fd5b600061057387878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a5292505050565b90506001818154811061058257fe5b90600052602060002090600b0201600401600001600182815481106105a357fe5b90600052602060002090600b0201600401600101600183815481106105c457fe5b600091825260209091206006600b909202010154600180546001600160a01b0390921691859081106105f257fe5b60009182526020918290206007600b909202010154845460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152919286919083018282801561068d5780601f106106625761010080835404028352916020019161068d565b820191906000526020600020905b81548152906001019060200180831161067057829003601f168201915b5050865460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529599508894509250840190508282801561071b5780601f106106f05761010080835404028352916020019161071b565b820191906000526020600020905b8154815290600101906020018083116106fe57829003601f168201915b5050505050925094509450945094505092959194509250565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561078157600080fd5b505afa158015610795573d6000803e3d6000fd5b505050506040513d60208110156107ab57600080fd5b50516001600160a01b031633146107fa576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b604482015290519081900360640190fd5b60026000848460405160200180806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505093505050506040516020818303038152906040528051906020012081526020019081526020016000205460001415610b55576003805460010190819055604080516020808201908152918101859052600291600091879187918190606001848480828437600081840152601f19601f8201169050808301925050509350505050604051602081830303815290604052805190602001208152602001908152602001600020819055506000600180548091906001016108f49190611c28565b905083836001838154811061090557fe5b6000918252602090912061091f93600b9092020191611c59565b50600180828154811061092e57fe5b90600052602060002090600b020160010181905550600180828154811061095157fe5b90600052602060002090600b02016002018190555060006001828154811061097557fe5b90600052602060002090600b02016003018190555060405180602001604052806000815250600182815481106109a757fe5b90600052602060002090600b020160040160000190805190602001906109ce929190611cd7565b5060405180602001604052806000815250600182815481106109ec57fe5b90600052602060002090600b02016004016001019080519060200190610a13929190611cd7565b50600060018281548110610a2357fe5b600091825260208220600b919091020160060180546001600160a01b0319166001600160a01b0393909316929092179091556001805483908110610a6357fe5b600091825260209091206007600b9092020101556001805482908110610a8557fe5b90600052602060002090600b02016001015460018281548110610aa457fe5b600091825260208083206001600160a01b03871684526009600b9093020191909101905260409020556001805482908110610adb57fe5b60009182526020808320604080518082019091526001600160a01b0387811682526001828501818152600b969096029093016008018054938401815586529290942093519301805492516001600160a01b0319909316939091169290921760ff60a01b1916600160a01b9115159190910217905550610e47565b6000610b9684848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a5292505050565b905060018181548110610ba557fe5b600091825260208083206001600160a01b03861684526009600b909302019190910190526040902054610cff5760018181548110610bdf57fe5b600091825260209091206001600b909202018101805482019055805482908110610c0557fe5b90600052602060002090600b02016001015460018281548110610c2457fe5b600091825260208083206001600160a01b03871684526009600b9093020191909101905260409020556001805482908110610c5b57fe5b60009182526020808320604080518082019091526001600160a01b0387811682526001828501818152600b96909602909301600801805480850182559087529390952090519201805493516001600160a01b0319909416929094169190911760ff60a01b1916600160a01b9215159290920291909117909155805482908110610ce057fe5b600091825260209091206002600b909202010180546001019055610e45565b6000610d4285858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250611afa915050565b905060018281548110610d5157fe5b90600052602060002090600b02016008018181548110610d6d57fe5b600091825260209091200154600160a01b900460ff16151560011415610dcc576040805162461bcd60e51b815260206004820152600f60248201526e30b63932b0b23c9030903b37ba32b960891b604482015290519081900360640190fd5b6001808381548110610dda57fe5b90600052602060002090600b02016008018281548110610df657fe5b60009182526020909120018054911515600160a01b0260ff60a01b199092169190911790556001805483908110610e2957fe5b600091825260209091206002600b909202010180546001019055505b505b604080516001600160a01b03831660208201528181529081018390527f424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574908490849084908060608101858580828437600083820152604051601f909101601f1916909201829003965090945050505050a1505050565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b158015610f0a57600080fd5b505afa158015610f1e573d6000803e3d6000fd5b505050506040513d6020811015610f3457600080fd5b50516001600160a01b03163314610f83576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b604482015290519081900360640190fd5b82828080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250849250610fc7915083905082611b4a565b151560011461100f576040805162461bcd60e51b815260206004820152600f60248201526e36bab9ba1031329030903b37ba32b960891b604482015290519081900360640190fd5b600061105086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a5292505050565b9050600061109587878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250899250611afa915050565b9050600182815481106110a457fe5b6000918252602082206002600b9092020101805460001901905560018054849081106110cc57fe5b90600052602060002090600b020160080182815481106110e857fe5b9060005260206000200160000160146101000a81548160ff0219169083151502179055507f654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b68787876040518080602001836001600160a01b03166001600160a01b031681526020018281038252858582818152602001925080828437600083820152604051601f909101601f1916909201829003965090945050505050a150505050505050565b60008060009054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b1580156111de57600080fd5b505afa1580156111f2573d6000803e3d6000fd5b505050506040513d602081101561120857600080fd5b50516001600160a01b03163314611257576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b604482015290519081900360640190fd5b84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525086925061129b915083905082611b4a565b15156001146112e3576040805162461bcd60e51b815260206004820152600f60248201526e36bab9ba1031329030903b37ba32b960891b604482015290519081900360640190fd5b61132487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250611bf2915050565b151560011461136f576040805162461bcd60e51b81526020600482015260126024820152716e6f7468696e6720746f20617070726f766560701b604482015290519081900360640190fd5b60006113b088888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a5292505050565b9050600181815481106113bf57fe5b60009182526020808320848452600a600b9093020191909101815260408083206001600160a01b038a16845290915290205460ff16151560011415611440576040805162461bcd60e51b815260206004820152601260248201527163616e6e6f7420646f75626c6520766f746560701b604482015290519081900360640190fd5b6001818154811061144d57fe5b600091825260209091206003600b909202010180546001908101909155805481908390811061147857fe5b60009182526020808320858452600b92909202909101600a01815260408083206001600160a01b038b168452825291829020805460ff19169315159390931790925580518281529182018990527f87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508918a918a919081908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a160026001828154811061152657fe5b90600052602060002090600b0201600201548161153f57fe5b046001828154811061154d57fe5b90600052602060002090600b020160030154111561165c57604051806020016040528060008152506001828154811061158257fe5b90600052602060002090600b020160040160000190805190602001906115a9929190611cd7565b5060405180602001604052806000815250600182815481106115c757fe5b90600052602060002090600b020160040160010190805190602001906115ee929190611cd7565b506000600182815481106115fe57fe5b600091825260208220600b919091020160060180546001600160a01b0319166001600160a01b039390931692909217909155600180548390811061163e57fe5b600091825260209091206007600b9092020101555060019250611662565b60009350505b5050949350505050565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b1580156116b957600080fd5b505afa1580156116cd573d6000803e3d6000fd5b505050506040513d60208110156116e357600080fd5b50516001600160a01b03163314611732576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b604482015290519081900360640190fd5b61177188888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509250611bf2915050565b6117ac5760405162461bcd60e51b8152600401808060200182810382526034815260200180611e696034913960400191505060405180910390fd5b60006117ed89898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a5292505050565b90508686600183815481106117fe57fe5b90600052602060002090600b0201600401600001919061181f929190611c59565b5084846001838154811061182f57fe5b90600052602060002090600b02016004016001019190611850929190611c59565b50826001828154811061185f57fe5b90600052602060002090600b020160040160020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600182815481106118a557fe5b6000918252602082206007600b9092020101919091555b600182815481106118c957fe5b90600052602060002090600b0201600801805490508110156119c257600182815481106118f257fe5b90600052602060002090600b0201600801818154811061190e57fe5b600091825260209091200154600160a01b900460ff16156119ba5760006001838154811061193857fe5b90600052602060002090600b0201600a01600084815260200190815260200160002060006001858154811061196957fe5b90600052602060002090600b0201600801848154811061198557fe5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff19169115159190911790555b6001016118bc565b506000600182815481106119d257fe5b90600052602060002090600b0201600301819055507f5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3898960405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a1505050505050505050565b6000600160026000846040516020018080602001828103825283818151815260200191508051906020019080838360005b83811015611a9b578181015183820152602001611a83565b50505050905090810190601f168015611ac85780820380516001836020036101000a031916815260200191505b509250505060405160208183030381529060405280519060200120815260200190815260200160002054039050919050565b600080611b0684611a52565b90506001808281548110611b1657fe5b600091825260208083206001600160a01b03881684526009600b909302019190910190526040902054039150505b92915050565b600080611b5684611a52565b905060018181548110611b6557fe5b600091825260208083206001600160a01b03871684526009600b909302019190910190526040902054611b9c576000915050611b44565b6000611ba88585611afa565b905060018281548110611bb757fe5b90600052602060002090600b02016008018181548110611bd357fe5b600091825260209091200154600160a01b900460ff1695945050505050565b6000816001611c0085611a52565b81548110611c0a57fe5b90600052602060002090600b02016004016003015414905092915050565b815481835581811115611c5457600b0281600b028360005260206000209182019101611c549190611d45565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611c9a5782800160ff19823516178555611cc7565b82800160010185558215611cc7579182015b82811115611cc7578235825591602001919060010190611cac565b50611cd3929150611dca565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611d1857805160ff1916838001178555611cc7565b82800160010185558215611cc7579182015b82811115611cc7578251825591602001919060010190611d2a565b611dc791905b80821115611cd3576000611d5f8282611de4565b60006001830181905560028301819055600383018190556004830190611d858282611de4565b611d93600183016000611de4565b506002810180546001600160a01b031916905560006003909101819055611dbe906008840190611e2b565b50600b01611d4b565b90565b611dc791905b80821115611cd35760008155600101611dd0565b50805460018160011615610100020316600290046000825580601f10611e0a5750611e28565b601f016020900490600052602060002090810190611e289190611dca565b50565b5080546000825590600052602060002090810190611e289190611dc791905b80821115611cd35780546001600160a81b0319168155600101611e4a56fe6974656d732070656e64696e6720666f7220617070726f76616c2e206e6577206974656d2063616e6e6f74206265206164646564a265627a7a72315820888112f67efda532c62d71744f9367722c2e10a6fce5ef503e7ba6b180855bbf64736f6c634300050b0032"

// DeployEeaVoterManager deploys a new Ethereum contract, binding an instance of EeaVoterManager to it.
func DeployEeaVoterManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permUpgradable common.Address) (common.Address, *types.Transaction, *EeaVoterManager, error) {
	parsed, err := abi.JSON(strings.NewReader(EeaVoterManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EeaVoterManagerBin), backend, _permUpgradable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EeaVoterManager{EeaVoterManagerCaller: EeaVoterManagerCaller{contract: contract}, EeaVoterManagerTransactor: EeaVoterManagerTransactor{contract: contract}, EeaVoterManagerFilterer: EeaVoterManagerFilterer{contract: contract}}, nil
}

// EeaVoterManager is an auto generated Go binding around an Ethereum contract.
type EeaVoterManager struct {
	EeaVoterManagerCaller     // Read-only binding to the contract
	EeaVoterManagerTransactor // Write-only binding to the contract
	EeaVoterManagerFilterer   // Log filterer for contract events
}

// EeaVoterManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EeaVoterManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EeaVoterManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EeaVoterManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EeaVoterManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EeaVoterManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EeaVoterManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EeaVoterManagerSession struct {
	Contract     *EeaVoterManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EeaVoterManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EeaVoterManagerCallerSession struct {
	Contract *EeaVoterManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EeaVoterManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EeaVoterManagerTransactorSession struct {
	Contract     *EeaVoterManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EeaVoterManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EeaVoterManagerRaw struct {
	Contract *EeaVoterManager // Generic contract binding to access the raw methods on
}

// EeaVoterManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EeaVoterManagerCallerRaw struct {
	Contract *EeaVoterManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EeaVoterManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EeaVoterManagerTransactorRaw struct {
	Contract *EeaVoterManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEeaVoterManager creates a new instance of EeaVoterManager, bound to a specific deployed contract.
func NewEeaVoterManager(address common.Address, backend bind.ContractBackend) (*EeaVoterManager, error) {
	contract, err := bindEeaVoterManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EeaVoterManager{EeaVoterManagerCaller: EeaVoterManagerCaller{contract: contract}, EeaVoterManagerTransactor: EeaVoterManagerTransactor{contract: contract}, EeaVoterManagerFilterer: EeaVoterManagerFilterer{contract: contract}}, nil
}

// NewEeaVoterManagerCaller creates a new read-only instance of EeaVoterManager, bound to a specific deployed contract.
func NewEeaVoterManagerCaller(address common.Address, caller bind.ContractCaller) (*EeaVoterManagerCaller, error) {
	contract, err := bindEeaVoterManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EeaVoterManagerCaller{contract: contract}, nil
}

// NewEeaVoterManagerTransactor creates a new write-only instance of EeaVoterManager, bound to a specific deployed contract.
func NewEeaVoterManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EeaVoterManagerTransactor, error) {
	contract, err := bindEeaVoterManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EeaVoterManagerTransactor{contract: contract}, nil
}

// NewEeaVoterManagerFilterer creates a new log filterer instance of EeaVoterManager, bound to a specific deployed contract.
func NewEeaVoterManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EeaVoterManagerFilterer, error) {
	contract, err := bindEeaVoterManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EeaVoterManagerFilterer{contract: contract}, nil
}

// bindEeaVoterManager binds a generic wrapper to an already deployed contract.
func bindEeaVoterManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EeaVoterManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EeaVoterManager *EeaVoterManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EeaVoterManager.Contract.EeaVoterManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EeaVoterManager *EeaVoterManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.EeaVoterManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EeaVoterManager *EeaVoterManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.EeaVoterManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EeaVoterManager *EeaVoterManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EeaVoterManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EeaVoterManager *EeaVoterManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EeaVoterManager *EeaVoterManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.contract.Transact(opts, method, params...)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_EeaVoterManager *EeaVoterManagerCaller) GetPendingOpDetails(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
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
	err := _EeaVoterManager.contract.Call(opts, out, "getPendingOpDetails", _orgId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_EeaVoterManager *EeaVoterManagerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _EeaVoterManager.Contract.GetPendingOpDetails(&_EeaVoterManager.CallOpts, _orgId)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(string _orgId) constant returns(string, string, address, uint256)
func (_EeaVoterManager *EeaVoterManagerCallerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _EeaVoterManager.Contract.GetPendingOpDetails(&_EeaVoterManager.CallOpts, _orgId)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_EeaVoterManager *EeaVoterManagerTransactor) AddVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _EeaVoterManager.contract.Transact(opts, "addVoter", _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_EeaVoterManager *EeaVoterManagerSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.AddVoter(&_EeaVoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(string _orgId, address _vAccount) returns()
func (_EeaVoterManager *EeaVoterManagerTransactorSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.AddVoter(&_EeaVoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_EeaVoterManager *EeaVoterManagerTransactor) AddVotingItem(opts *bind.TransactOpts, _authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _EeaVoterManager.contract.Transact(opts, "addVotingItem", _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_EeaVoterManager *EeaVoterManagerSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.AddVotingItem(&_EeaVoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(string _authOrg, string _orgId, string _enodeId, address _account, uint256 _pendingOp) returns()
func (_EeaVoterManager *EeaVoterManagerTransactorSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.AddVotingItem(&_EeaVoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_EeaVoterManager *EeaVoterManagerTransactor) DeleteVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _EeaVoterManager.contract.Transact(opts, "deleteVoter", _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_EeaVoterManager *EeaVoterManagerSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.DeleteVoter(&_EeaVoterManager.TransactOpts, _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(string _orgId, address _vAccount) returns()
func (_EeaVoterManager *EeaVoterManagerTransactorSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.DeleteVoter(&_EeaVoterManager.TransactOpts, _orgId, _vAccount)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_EeaVoterManager *EeaVoterManagerTransactor) ProcessVote(opts *bind.TransactOpts, _authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _EeaVoterManager.contract.Transact(opts, "processVote", _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_EeaVoterManager *EeaVoterManagerSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.ProcessVote(&_EeaVoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(string _authOrg, address _vAccount, uint256 _pendingOp) returns(bool)
func (_EeaVoterManager *EeaVoterManagerTransactorSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _EeaVoterManager.Contract.ProcessVote(&_EeaVoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// EeaVoterManagerVoteProcessedIterator is returned from FilterVoteProcessed and is used to iterate over the raw logs and unpacked data for VoteProcessed events raised by the EeaVoterManager contract.
type EeaVoterManagerVoteProcessedIterator struct {
	Event *EeaVoterManagerVoteProcessed // Event containing the contract specifics and raw log

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
func (it *EeaVoterManagerVoteProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EeaVoterManagerVoteProcessed)
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
		it.Event = new(EeaVoterManagerVoteProcessed)
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
func (it *EeaVoterManagerVoteProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EeaVoterManagerVoteProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EeaVoterManagerVoteProcessed represents a VoteProcessed event raised by the EeaVoterManager contract.
type EeaVoterManagerVoteProcessed struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteProcessed is a free log retrieval operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_EeaVoterManager *EeaVoterManagerFilterer) FilterVoteProcessed(opts *bind.FilterOpts) (*EeaVoterManagerVoteProcessedIterator, error) {

	logs, sub, err := _EeaVoterManager.contract.FilterLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return &EeaVoterManagerVoteProcessedIterator{contract: _EeaVoterManager.contract, event: "VoteProcessed", logs: logs, sub: sub}, nil
}

// WatchVoteProcessed is a free log subscription operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: event VoteProcessed(string _orgId)
func (_EeaVoterManager *EeaVoterManagerFilterer) WatchVoteProcessed(opts *bind.WatchOpts, sink chan<- *EeaVoterManagerVoteProcessed) (event.Subscription, error) {

	logs, sub, err := _EeaVoterManager.contract.WatchLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EeaVoterManagerVoteProcessed)
				if err := _EeaVoterManager.contract.UnpackLog(event, "VoteProcessed", log); err != nil {
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
func (_EeaVoterManager *EeaVoterManagerFilterer) ParseVoteProcessed(log types.Log) (*EeaVoterManagerVoteProcessed, error) {
	event := new(EeaVoterManagerVoteProcessed)
	if err := _EeaVoterManager.contract.UnpackLog(event, "VoteProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EeaVoterManagerVoterAddedIterator is returned from FilterVoterAdded and is used to iterate over the raw logs and unpacked data for VoterAdded events raised by the EeaVoterManager contract.
type EeaVoterManagerVoterAddedIterator struct {
	Event *EeaVoterManagerVoterAdded // Event containing the contract specifics and raw log

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
func (it *EeaVoterManagerVoterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EeaVoterManagerVoterAdded)
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
		it.Event = new(EeaVoterManagerVoterAdded)
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
func (it *EeaVoterManagerVoterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EeaVoterManagerVoterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EeaVoterManagerVoterAdded represents a VoterAdded event raised by the EeaVoterManager contract.
type EeaVoterManagerVoterAdded struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterAdded is a free log retrieval operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_EeaVoterManager *EeaVoterManagerFilterer) FilterVoterAdded(opts *bind.FilterOpts) (*EeaVoterManagerVoterAddedIterator, error) {

	logs, sub, err := _EeaVoterManager.contract.FilterLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return &EeaVoterManagerVoterAddedIterator{contract: _EeaVoterManager.contract, event: "VoterAdded", logs: logs, sub: sub}, nil
}

// WatchVoterAdded is a free log subscription operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: event VoterAdded(string _orgId, address _vAccount)
func (_EeaVoterManager *EeaVoterManagerFilterer) WatchVoterAdded(opts *bind.WatchOpts, sink chan<- *EeaVoterManagerVoterAdded) (event.Subscription, error) {

	logs, sub, err := _EeaVoterManager.contract.WatchLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EeaVoterManagerVoterAdded)
				if err := _EeaVoterManager.contract.UnpackLog(event, "VoterAdded", log); err != nil {
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
func (_EeaVoterManager *EeaVoterManagerFilterer) ParseVoterAdded(log types.Log) (*EeaVoterManagerVoterAdded, error) {
	event := new(EeaVoterManagerVoterAdded)
	if err := _EeaVoterManager.contract.UnpackLog(event, "VoterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EeaVoterManagerVoterDeletedIterator is returned from FilterVoterDeleted and is used to iterate over the raw logs and unpacked data for VoterDeleted events raised by the EeaVoterManager contract.
type EeaVoterManagerVoterDeletedIterator struct {
	Event *EeaVoterManagerVoterDeleted // Event containing the contract specifics and raw log

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
func (it *EeaVoterManagerVoterDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EeaVoterManagerVoterDeleted)
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
		it.Event = new(EeaVoterManagerVoterDeleted)
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
func (it *EeaVoterManagerVoterDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EeaVoterManagerVoterDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EeaVoterManagerVoterDeleted represents a VoterDeleted event raised by the EeaVoterManager contract.
type EeaVoterManagerVoterDeleted struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterDeleted is a free log retrieval operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_EeaVoterManager *EeaVoterManagerFilterer) FilterVoterDeleted(opts *bind.FilterOpts) (*EeaVoterManagerVoterDeletedIterator, error) {

	logs, sub, err := _EeaVoterManager.contract.FilterLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return &EeaVoterManagerVoterDeletedIterator{contract: _EeaVoterManager.contract, event: "VoterDeleted", logs: logs, sub: sub}, nil
}

// WatchVoterDeleted is a free log subscription operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: event VoterDeleted(string _orgId, address _vAccount)
func (_EeaVoterManager *EeaVoterManagerFilterer) WatchVoterDeleted(opts *bind.WatchOpts, sink chan<- *EeaVoterManagerVoterDeleted) (event.Subscription, error) {

	logs, sub, err := _EeaVoterManager.contract.WatchLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EeaVoterManagerVoterDeleted)
				if err := _EeaVoterManager.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
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
func (_EeaVoterManager *EeaVoterManagerFilterer) ParseVoterDeleted(log types.Log) (*EeaVoterManagerVoterDeleted, error) {
	event := new(EeaVoterManagerVoterDeleted)
	if err := _EeaVoterManager.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EeaVoterManagerVotingItemAddedIterator is returned from FilterVotingItemAdded and is used to iterate over the raw logs and unpacked data for VotingItemAdded events raised by the EeaVoterManager contract.
type EeaVoterManagerVotingItemAddedIterator struct {
	Event *EeaVoterManagerVotingItemAdded // Event containing the contract specifics and raw log

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
func (it *EeaVoterManagerVotingItemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EeaVoterManagerVotingItemAdded)
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
		it.Event = new(EeaVoterManagerVotingItemAdded)
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
func (it *EeaVoterManagerVotingItemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EeaVoterManagerVotingItemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EeaVoterManagerVotingItemAdded represents a VotingItemAdded event raised by the EeaVoterManager contract.
type EeaVoterManagerVotingItemAdded struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVotingItemAdded is a free log retrieval operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_EeaVoterManager *EeaVoterManagerFilterer) FilterVotingItemAdded(opts *bind.FilterOpts) (*EeaVoterManagerVotingItemAddedIterator, error) {

	logs, sub, err := _EeaVoterManager.contract.FilterLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return &EeaVoterManagerVotingItemAddedIterator{contract: _EeaVoterManager.contract, event: "VotingItemAdded", logs: logs, sub: sub}, nil
}

// WatchVotingItemAdded is a free log subscription operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: event VotingItemAdded(string _orgId)
func (_EeaVoterManager *EeaVoterManagerFilterer) WatchVotingItemAdded(opts *bind.WatchOpts, sink chan<- *EeaVoterManagerVotingItemAdded) (event.Subscription, error) {

	logs, sub, err := _EeaVoterManager.contract.WatchLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EeaVoterManagerVotingItemAdded)
				if err := _EeaVoterManager.contract.UnpackLog(event, "VotingItemAdded", log); err != nil {
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
func (_EeaVoterManager *EeaVoterManagerFilterer) ParseVotingItemAdded(log types.Log) (*EeaVoterManagerVotingItemAdded, error) {
	event := new(EeaVoterManagerVotingItemAdded)
	if err := _EeaVoterManager.contract.UnpackLog(event, "VotingItemAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

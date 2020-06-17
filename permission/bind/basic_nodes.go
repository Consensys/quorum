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

// NodeManagerABI is the input ABI used to generate the binding from.
const NodeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enodeId\",\"type\":\"string\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addOrgNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeIndex\",\"type\":\"uint256\"}],\"name\":\"getNodeDetailsFromIndex\",\"outputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addAdminNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeRecoveryInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeRecoveryCompleted\",\"type\":\"event\"}]"

// NodeManagerBin is the compiled bytecode used for deploying new contracts.
var NodeManagerBin = "0x608060405234801561001057600080fd5b506040516020806125388339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a03199092169190911790556124d6806100626000396000f3fe608060405234801561001057600080fd5b506004361061008c5760003560e060020a9004806397c07a9b1161005f57806397c07a9b1461042e578063a97a44061461044b578063b81c806a1461050d578063e3b09d84146102aa5761008c565b80630cc50146146100915780633f0e0e47146101555780633f5e1a45146102aa57806386bc36521461036c575b600080fd5b610153600480360360608110156100a757600080fd5b8101906020810181356401000000008111156100c257600080fd5b8201836020820111156100d457600080fd5b803590602001918460018302840111640100000000831117156100f657600080fd5b91939092909160208101903564010000000081111561011457600080fd5b82018360208201111561012657600080fd5b8035906020019184600183028401116401000000008311171561014857600080fd5b919350915035610527565b005b6101c56004803603602081101561016b57600080fd5b81019060208101813564010000000081111561018657600080fd5b82018360208201111561019857600080fd5b803590602001918460018302840111640100000000831117156101ba57600080fd5b509092509050610f34565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561020c5781810151838201526020016101f4565b50505050905090810190601f1680156102395780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561026c578181015183820152602001610254565b50505050905090810190601f1680156102995780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b610153600480360360408110156102c057600080fd5b8101906020810181356401000000008111156102db57600080fd5b8201836020820111156102ed57600080fd5b8035906020019184600183028401116401000000008311171561030f57600080fd5b91939092909160208101903564010000000081111561032d57600080fd5b82018360208201111561033f57600080fd5b8035906020019184600183028401116401000000008311171561036157600080fd5b50909250905061120b565b6101536004803603604081101561038257600080fd5b81019060208101813564010000000081111561039d57600080fd5b8201836020820111156103af57600080fd5b803590602001918460018302840111640100000000831117156103d157600080fd5b9193909290916020810190356401000000008111156103ef57600080fd5b82018360208201111561040157600080fd5b8035906020019184600183028401116401000000008311171561042357600080fd5b5090925090506115e2565b6101c56004803603602081101561044457600080fd5b5035611b01565b6101536004803603604081101561046157600080fd5b81019060208101813564010000000081111561047c57600080fd5b82018360208201111561048e57600080fd5b803590602001918460018302840111640100000000831117156104b057600080fd5b9193909290916020810190356401000000008111156104ce57600080fd5b8201836020820111156104e057600080fd5b8035906020019184600183028401116401000000008311171561050257600080fd5b509092509050611c90565b610515612067565b60408051918252519081900360200190f35b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561057757600080fd5b505afa15801561058b573d6000803e3d6000fd5b505050506040513d60208110156105a157600080fd5b5051600160a060020a031633146105f0576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612438833981519152604482015290519081900360640190fd5b84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b8381101561066d578181015183820152602001610655565b50505050905090810190601f16801561069a5780820380516001836020036101000a031916815260200191505b5060408051601f1981840301815291815281516020928301208652908501959095525050500160002054151561071a576040805160e560020a62461bcd02815260206004820152601e60248201527f70617373656420656e6f646520696420646f6573206e6f742065786973740000604482015290519081900360640190fd5b61078d86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a01819004810282018101909252888152925088915087908190840183828082843760009201919091525061206e92505050565b15156107cd5760405160e560020a62461bcd02815260040180806020018281038252602a8152602001806123ee602a913960400191505060405180910390fd5b81600114806107dc5750816002145b806107e75750816003145b806107f25750816004145b806107fd5750816005145b151561083d5760405160e560020a62461bcd0281526004018080602001828103825260268152602001806124586026913960400191505060405180910390fd5b81600114156109bd5761088586868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121d092505050565b6002146108ca576040805160e560020a62461bcd02815260206004820152601d6024820152600080516020612418833981519152604482015290519081900360640190fd5b6003600161090d88888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122ad92505050565b8154811061091757fe5b9060005260206000209060030201600201819055507fc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f2c565b8160021415610b3d57610a0586868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121d092505050565b600314610a4a576040805160e560020a62461bcd02815260206004820152601d6024820152600080516020612418833981519152604482015290519081900360640190fd5b60026001610a8d88888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122ad92505050565b81548110610a9757fe5b9060005260206000209060030201600201819055507f49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f2c565b8160031415610c395760046001610b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122ad92505050565b81548110610b9357fe5b9060005260206000209060030201600201819055507f4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f2c565b8160041415610db957610c8186868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121d092505050565b600414610cc6576040805160e560020a62461bcd02815260206004820152601d6024820152600080516020612418833981519152604482015290519081900360640190fd5b60056001610d0988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122ad92505050565b81548110610d1357fe5b9060005260206000209060030201600201819055507ffd385c618a1e89d01fb9a21780846793e282e8bc0b60caf6ccb3e422d543fbfb868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f2c565b610df886868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121d092505050565b600514610e3d576040805160e560020a62461bcd02815260206004820152601d6024820152600080516020612418833981519152604482015290519081900360640190fd5b60026001610e8088888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122ad92505050565b81548110610e8a57fe5b9060005260206000209060030201600201819055507f787d7bc525e7c4658c64e3e456d974a1be21cc196e8162a4bf1337a12cb38dac868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15b505050505050565b606080600060026000836040516020018080602001828103825283818151815260200191508051906020019080838360005b83811015610f7e578181015183820152602001610f66565b50505050905090810190601f168015610fab5780820380516001836020036101000a031916815260200191505b5092505050604051602081830303815290604052805190602001208152602001908152602001600020546000141561103c5784846000602060405190810160405280600081525092919082828080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250969950919750919550611204945050505050565b600061107d86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122ad92505050565b905060018181548110151561108e57fe5b90600052602060002090600302016001016001828154811015156110ae57fe5b90600052602060002090600302016000016001838154811015156110ce57fe5b60009182526020918290206002600390920201810154845460408051601f6000196101006001861615020190931694909404918201859004850284018501905280835290928591908301828280156111675780601f1061113c57610100808354040283529160200191611167565b820191906000526020600020905b81548152906001019060200180831161114a57829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959850879450925084019050828280156111f55780601f106111ca576101008083540402835291602001916111f5565b820191906000526020600020905b8154815290600101906020018083116111d857829003601f168201915b50505050509150935093509350505b9250925092565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561125b57600080fd5b505afa15801561126f573d6000803e3d6000fd5b505050506040513d602081101561128557600080fd5b5051600160a060020a031633146112d4576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612438833981519152604482015290519081900360640190fd5b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b83811015611351578181015183820152602001611339565b50505050905090810190601f16801561137e5780820380516001836020036101000a031916815260200191505b5060408051601f1981840301815291815281516020928301208652908501959095525050500160002054156113fd576040805160e560020a62461bcd02815260206004820152601660248201527f70617373656420656e6f64652069642065786973747300000000000000000000604482015290519081900360640190fd5b6003805460010190819055604080516020808201908152918101879052600291600091899189918190606001848480828437600081840152601f19601f820116905080830192505050935050505060405160208183030381529060405280519060200120815260200190815260200160002081905550600160606040519081016040528087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f88018190048102820181019092528681529181019190879087908190840183828082843760009201829052509385525050600260209384015250835460018101808655948252908290208351805160039093029091019261152892849290910190612355565b5060208281015180516115419260018501920190612355565b50604082015181600201555050507f0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d858585856040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561163257600080fd5b505afa158015611646573d6000803e3d6000fd5b505050506040513d602081101561165c57600080fd5b5051600160a060020a031633146116ab576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612438833981519152604482015290519081900360640190fd5b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b83811015611728578181015183820152602001611710565b50505050905090810190601f1680156117555780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181528151602092830120865290850195909552505050016000205415156117d5576040805160e560020a62461bcd02815260206004820152601e60248201527f70617373656420656e6f646520696420646f6573206e6f742065786973740000604482015290519081900360640190fd5b61184885858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8901819004810282018101909252878152925087915086908190840183828082843760009201919091525061206e92505050565b15156118885760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061247e602d913960400191505060405180910390fd5b6118c785858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121d092505050565b60011461191e576040805160e560020a62461bcd02815260206004820152601c60248201527f6e6f7468696e672070656e64696e6720666f7220617070726f76616c00000000604482015290519081900360640190fd5b600061195f86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122ad92505050565b9050600260018281548110151561197257fe5b9060005260206000209060030201600201819055507f0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d6001828154811015156119b757fe5b90600052602060002090600302016000016001838154811015156119d757fe5b9060005260206000209060030201600101604051808060200180602001838103835285818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611a755780601f10611a4a57610100808354040283529160200191611a75565b820191906000526020600020905b815481529060010190602001808311611a5857829003601f168201915b5050838103825284546002600019610100600184161502019091160480825260209091019085908015611ae95780601f10611abe57610100808354040283529160200191611ae9565b820191906000526020600020905b815481529060010190602001808311611acc57829003601f168201915b505094505050505060405180910390a1505050505050565b6060806000600184815481101515611b1557fe5b9060005260206000209060030201600101600185815481101515611b3557fe5b9060005260206000209060030201600001600186815481101515611b5557fe5b60009182526020918290206002600390920201810154845460408051601f600019610100600186161502019093169490940491820185900485028401850190528083529092859190830182828015611bee5780601f10611bc357610100808354040283529160200191611bee565b820191906000526020600020905b815481529060010190602001808311611bd157829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815295985087945092508401905082828015611c7c5780601f10611c5157610100808354040283529160200191611c7c565b820191906000526020600020905b815481529060010190602001808311611c5f57829003601f168201915b505050505091509250925092509193909250565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b158015611ce057600080fd5b505afa158015611cf4573d6000803e3d6000fd5b505050506040513d6020811015611d0a57600080fd5b5051600160a060020a03163314611d59576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612438833981519152604482015290519081900360640190fd5b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b83811015611dd6578181015183820152602001611dbe565b50505050905090810190601f168015611e035780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181528151602092830120865290850195909552505050016000205415611e82576040805160e560020a62461bcd02815260206004820152601660248201527f70617373656420656e6f64652069642065786973747300000000000000000000604482015290519081900360640190fd5b6003805460010190819055604080516020808201908152918101879052600291600091899189918190606001848480828437600081840152601f19601f820116905080830192505050935050505060405160208183030381529060405280519060200120815260200190815260200160002081905550600160606040519081016040528087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f880181900481028201810190925286815291810191908790879081908401838280828437600092018290525093855250506001602093840181905285549081018087559583529183902084518051600390940290910193611fad93859350910190612355565b506020828101518051611fc69260018501920190612355565b50604082015181600201555050507fb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f858585856040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050565b6003545b90565b6000816040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156120b1578181015183820152602001612099565b50505050905090810190601f1680156120de5780820380516001836020036101000a031916815260200191505b5092505050604051602081830303815290604052805190602001206001612104856122ad565b8154811061210e57fe5b906000526020600020906003020160010160405160200180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156121ab5780601f10612180576101008083540402835291602001916121ab565b820191906000526020600020905b81548152906001019060200180831161218e57829003601f168201915b5050925050506040516020818303038152906040528051906020012014905092915050565b600060026000836040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156122175781810151838201526020016121ff565b50505050905090810190601f1680156122445780820380516001836020036101000a031916815260200191505b5092505050604051602081830303815290604052805190602001208152602001908152602001600020546000141561227e575060006122a8565b6001612289836122ad565b8154811061229357fe5b90600052602060002090600302016002015490505b919050565b6000600160026000846040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156122f65781810151838201526020016122de565b50505050905090810190601f1680156123235780820380516001836020036101000a031916815260200191505b509250505060405160208183030381529060405280519060200120815260200190815260200160002054039050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061239657805160ff19168380011785556123c3565b828001600101855582156123c3579182015b828111156123c35782518255916020019190600101906123a8565b506123cf9291506123d3565b5090565b61206b91905b808211156123cf57600081556001016123d956fe656e6f646520696420646f6573206e6f742062656c6f6e6720746f2074686520706173736564206f72676f7065726174696f6e2063616e6e6f7420626520706572666f726d6564000000696e76616c69642063616c6c6572000000000000000000000000000000000000696e76616c6964206f7065726174696f6e2e2077726f6e6720616374696f6e20706173736564656e6f646520696420646f6573206e6f742062656c6f6e6720746f2074686520706173736564206f7267206964a165627a7a7230582056e58e3bf58b0c0c90d997e73401ac42acb39bacd2aafecf52dca6e05fb2f3250029"

// DeployNodeManager deploys a new Ethereum contract, binding an instance of NodeManager to it.
func DeployNodeManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permUpgradable common.Address) (common.Address, *types.Transaction, *NodeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeManagerBin), backend, _permUpgradable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeManager{NodeManagerCaller: NodeManagerCaller{contract: contract}, NodeManagerTransactor: NodeManagerTransactor{contract: contract}, NodeManagerFilterer: NodeManagerFilterer{contract: contract}}, nil
}

// NodeManager is an auto generated Go binding around an Ethereum contract.
type NodeManager struct {
	NodeManagerCaller     // Read-only binding to the contract
	NodeManagerTransactor // Write-only binding to the contract
	NodeManagerFilterer   // Log filterer for contract events
}

// NodeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeManagerSession struct {
	Contract     *NodeManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeManagerCallerSession struct {
	Contract *NodeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeManagerTransactorSession struct {
	Contract     *NodeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeManagerRaw struct {
	Contract *NodeManager // Generic contract binding to access the raw methods on
}

// NodeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeManagerCallerRaw struct {
	Contract *NodeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NodeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeManagerTransactorRaw struct {
	Contract *NodeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeManager creates a new instance of NodeManager, bound to a specific deployed contract.
func NewNodeManager(address common.Address, backend bind.ContractBackend) (*NodeManager, error) {
	contract, err := bindNodeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeManager{NodeManagerCaller: NodeManagerCaller{contract: contract}, NodeManagerTransactor: NodeManagerTransactor{contract: contract}, NodeManagerFilterer: NodeManagerFilterer{contract: contract}}, nil
}

// NewNodeManagerCaller creates a new read-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerCaller(address common.Address, caller bind.ContractCaller) (*NodeManagerCaller, error) {
	contract, err := bindNodeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerCaller{contract: contract}, nil
}

// NewNodeManagerTransactor creates a new write-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeManagerTransactor, error) {
	contract, err := bindNodeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerTransactor{contract: contract}, nil
}

// NewNodeManagerFilterer creates a new log filterer instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeManagerFilterer, error) {
	contract, err := bindNodeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeManagerFilterer{contract: contract}, nil
}

// bindNodeManager binds a generic wrapper to an already deployed contract.
func bindNodeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.NodeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transact(opts, method, params...)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(string enodeId) constant returns(string _orgId, string _enodeId, uint256 _nodeStatus)
func (_NodeManager *NodeManagerCaller) GetNodeDetails(opts *bind.CallOpts, enodeId string) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	ret := new(struct {
		OrgId      string
		EnodeId    string
		NodeStatus *big.Int
	})
	out := ret
	err := _NodeManager.contract.Call(opts, out, "getNodeDetails", enodeId)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(string enodeId) constant returns(string _orgId, string _enodeId, uint256 _nodeStatus)
func (_NodeManager *NodeManagerSession) GetNodeDetails(enodeId string) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetails(&_NodeManager.CallOpts, enodeId)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(string enodeId) constant returns(string _orgId, string _enodeId, uint256 _nodeStatus)
func (_NodeManager *NodeManagerCallerSession) GetNodeDetails(enodeId string) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetails(&_NodeManager.CallOpts, enodeId)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(uint256 _nodeIndex) constant returns(string _orgId, string _enodeId, uint256 _nodeStatus)
func (_NodeManager *NodeManagerCaller) GetNodeDetailsFromIndex(opts *bind.CallOpts, _nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	ret := new(struct {
		OrgId      string
		EnodeId    string
		NodeStatus *big.Int
	})
	out := ret
	err := _NodeManager.contract.Call(opts, out, "getNodeDetailsFromIndex", _nodeIndex)
	return *ret, err
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(uint256 _nodeIndex) constant returns(string _orgId, string _enodeId, uint256 _nodeStatus)
func (_NodeManager *NodeManagerSession) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetailsFromIndex(&_NodeManager.CallOpts, _nodeIndex)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(uint256 _nodeIndex) constant returns(string _orgId, string _enodeId, uint256 _nodeStatus)
func (_NodeManager *NodeManagerCallerSession) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetailsFromIndex(&_NodeManager.CallOpts, _nodeIndex)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerCaller) GetNumberOfNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeManager.contract.Call(opts, out, "getNumberOfNodes")
	return *ret0, err
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManager.Contract.GetNumberOfNodes(&_NodeManager.CallOpts)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerCallerSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManager.Contract.GetNumberOfNodes(&_NodeManager.CallOpts)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0xe3b09d84.
//
// Solidity: function addAdminNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactor) AddAdminNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addAdminNode", _enodeId, _orgId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0xe3b09d84.
//
// Solidity: function addAdminNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerSession) AddAdminNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddAdminNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0xe3b09d84.
//
// Solidity: function addAdminNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactorSession) AddAdminNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddAdminNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactor) AddNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addNode", _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerSession) AddNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactorSession) AddNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactor) AddOrgNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addOrgNode", _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerSession) AddOrgNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddOrgNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactorSession) AddOrgNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddOrgNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x86bc3652.
//
// Solidity: function approveNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactor) ApproveNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "approveNode", _enodeId, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x86bc3652.
//
// Solidity: function approveNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerSession) ApproveNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.ApproveNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x86bc3652.
//
// Solidity: function approveNode(string _enodeId, string _orgId) returns()
func (_NodeManager *NodeManagerTransactorSession) ApproveNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.ApproveNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(string _enodeId, string _orgId, uint256 _action) returns()
func (_NodeManager *NodeManagerTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _enodeId string, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "updateNodeStatus", _enodeId, _orgId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(string _enodeId, string _orgId, uint256 _action) returns()
func (_NodeManager *NodeManagerSession) UpdateNodeStatus(_enodeId string, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.UpdateNodeStatus(&_NodeManager.TransactOpts, _enodeId, _orgId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(string _enodeId, string _orgId, uint256 _action) returns()
func (_NodeManager *NodeManagerTransactorSession) UpdateNodeStatus(_enodeId string, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.UpdateNodeStatus(&_NodeManager.TransactOpts, _enodeId, _orgId, _action)
}

// NodeManagerNodeActivatedIterator is returned from FilterNodeActivated and is used to iterate over the raw logs and unpacked data for NodeActivated events raised by the NodeManager contract.
type NodeManagerNodeActivatedIterator struct {
	Event *NodeManagerNodeActivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeActivated)
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
		it.Event = new(NodeManagerNodeActivated)
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
func (it *NodeManagerNodeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeActivated represents a NodeActivated event raised by the NodeManager contract.
type NodeManagerNodeActivated struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeActivated is a free log retrieval operation binding the contract event 0x49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f.
//
// Solidity: event NodeActivated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) FilterNodeActivated(opts *bind.FilterOpts) (*NodeManagerNodeActivatedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeActivatedIterator{contract: _NodeManager.contract, event: "NodeActivated", logs: logs, sub: sub}, nil
}

// WatchNodeActivated is a free log subscription operation binding the contract event 0x49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f.
//
// Solidity: event NodeActivated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) WatchNodeActivated(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeActivated) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeActivated)
				if err := _NodeManager.contract.UnpackLog(event, "NodeActivated", log); err != nil {
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

// ParseNodeActivated is a log parse operation binding the contract event 0x49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f.
//
// Solidity: event NodeActivated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) ParseNodeActivated(log types.Log) (*NodeManagerNodeActivated, error) {
	event := new(NodeManagerNodeActivated)
	if err := _NodeManager.contract.UnpackLog(event, "NodeActivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeManagerNodeApprovedIterator is returned from FilterNodeApproved and is used to iterate over the raw logs and unpacked data for NodeApproved events raised by the NodeManager contract.
type NodeManagerNodeApprovedIterator struct {
	Event *NodeManagerNodeApproved // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeApproved)
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
		it.Event = new(NodeManagerNodeApproved)
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
func (it *NodeManagerNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeApproved represents a NodeApproved event raised by the NodeManager contract.
type NodeManagerNodeApproved struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeApproved is a free log retrieval operation binding the contract event 0x0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d.
//
// Solidity: event NodeApproved(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) FilterNodeApproved(opts *bind.FilterOpts) (*NodeManagerNodeApprovedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeApprovedIterator{contract: _NodeManager.contract, event: "NodeApproved", logs: logs, sub: sub}, nil
}

// WatchNodeApproved is a free log subscription operation binding the contract event 0x0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d.
//
// Solidity: event NodeApproved(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) WatchNodeApproved(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeApproved) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeApproved)
				if err := _NodeManager.contract.UnpackLog(event, "NodeApproved", log); err != nil {
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

// ParseNodeApproved is a log parse operation binding the contract event 0x0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d.
//
// Solidity: event NodeApproved(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) ParseNodeApproved(log types.Log) (*NodeManagerNodeApproved, error) {
	event := new(NodeManagerNodeApproved)
	if err := _NodeManager.contract.UnpackLog(event, "NodeApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeManagerNodeBlacklistedIterator is returned from FilterNodeBlacklisted and is used to iterate over the raw logs and unpacked data for NodeBlacklisted events raised by the NodeManager contract.
type NodeManagerNodeBlacklistedIterator struct {
	Event *NodeManagerNodeBlacklisted // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeBlacklisted)
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
		it.Event = new(NodeManagerNodeBlacklisted)
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
func (it *NodeManagerNodeBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeBlacklisted represents a NodeBlacklisted event raised by the NodeManager contract.
type NodeManagerNodeBlacklisted struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeBlacklisted is a free log retrieval operation binding the contract event 0x4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a.
//
// Solidity: event NodeBlacklisted(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) FilterNodeBlacklisted(opts *bind.FilterOpts) (*NodeManagerNodeBlacklistedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeBlacklistedIterator{contract: _NodeManager.contract, event: "NodeBlacklisted", logs: logs, sub: sub}, nil
}

// WatchNodeBlacklisted is a free log subscription operation binding the contract event 0x4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a.
//
// Solidity: event NodeBlacklisted(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) WatchNodeBlacklisted(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeBlacklisted) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeBlacklisted)
				if err := _NodeManager.contract.UnpackLog(event, "NodeBlacklisted", log); err != nil {
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

// ParseNodeBlacklisted is a log parse operation binding the contract event 0x4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a.
//
// Solidity: event NodeBlacklisted(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) ParseNodeBlacklisted(log types.Log) (*NodeManagerNodeBlacklisted, error) {
	event := new(NodeManagerNodeBlacklisted)
	if err := _NodeManager.contract.UnpackLog(event, "NodeBlacklisted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeManagerNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the NodeManager contract.
type NodeManagerNodeDeactivatedIterator struct {
	Event *NodeManagerNodeDeactivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeDeactivated)
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
		it.Event = new(NodeManagerNodeDeactivated)
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
func (it *NodeManagerNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeDeactivated represents a NodeDeactivated event raised by the NodeManager contract.
type NodeManagerNodeDeactivated struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a.
//
// Solidity: event NodeDeactivated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) FilterNodeDeactivated(opts *bind.FilterOpts) (*NodeManagerNodeDeactivatedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeDeactivatedIterator{contract: _NodeManager.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a.
//
// Solidity: event NodeDeactivated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeDeactivated) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeDeactivated)
				if err := _NodeManager.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
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

// ParseNodeDeactivated is a log parse operation binding the contract event 0xc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a.
//
// Solidity: event NodeDeactivated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) ParseNodeDeactivated(log types.Log) (*NodeManagerNodeDeactivated, error) {
	event := new(NodeManagerNodeDeactivated)
	if err := _NodeManager.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeManagerNodeProposedIterator is returned from FilterNodeProposed and is used to iterate over the raw logs and unpacked data for NodeProposed events raised by the NodeManager contract.
type NodeManagerNodeProposedIterator struct {
	Event *NodeManagerNodeProposed // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeProposed)
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
		it.Event = new(NodeManagerNodeProposed)
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
func (it *NodeManagerNodeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeProposed represents a NodeProposed event raised by the NodeManager contract.
type NodeManagerNodeProposed struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeProposed is a free log retrieval operation binding the contract event 0xb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f.
//
// Solidity: event NodeProposed(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) FilterNodeProposed(opts *bind.FilterOpts) (*NodeManagerNodeProposedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeProposedIterator{contract: _NodeManager.contract, event: "NodeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeProposed is a free log subscription operation binding the contract event 0xb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f.
//
// Solidity: event NodeProposed(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) WatchNodeProposed(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeProposed) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeProposed)
				if err := _NodeManager.contract.UnpackLog(event, "NodeProposed", log); err != nil {
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

// ParseNodeProposed is a log parse operation binding the contract event 0xb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f.
//
// Solidity: event NodeProposed(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) ParseNodeProposed(log types.Log) (*NodeManagerNodeProposed, error) {
	event := new(NodeManagerNodeProposed)
	if err := _NodeManager.contract.UnpackLog(event, "NodeProposed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeManagerNodeRecoveryCompletedIterator is returned from FilterNodeRecoveryCompleted and is used to iterate over the raw logs and unpacked data for NodeRecoveryCompleted events raised by the NodeManager contract.
type NodeManagerNodeRecoveryCompletedIterator struct {
	Event *NodeManagerNodeRecoveryCompleted // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeRecoveryCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeRecoveryCompleted)
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
		it.Event = new(NodeManagerNodeRecoveryCompleted)
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
func (it *NodeManagerNodeRecoveryCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeRecoveryCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeRecoveryCompleted represents a NodeRecoveryCompleted event raised by the NodeManager contract.
type NodeManagerNodeRecoveryCompleted struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRecoveryCompleted is a free log retrieval operation binding the contract event 0x787d7bc525e7c4658c64e3e456d974a1be21cc196e8162a4bf1337a12cb38dac.
//
// Solidity: event NodeRecoveryCompleted(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) FilterNodeRecoveryCompleted(opts *bind.FilterOpts) (*NodeManagerNodeRecoveryCompletedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeRecoveryCompleted")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeRecoveryCompletedIterator{contract: _NodeManager.contract, event: "NodeRecoveryCompleted", logs: logs, sub: sub}, nil
}

// WatchNodeRecoveryCompleted is a free log subscription operation binding the contract event 0x787d7bc525e7c4658c64e3e456d974a1be21cc196e8162a4bf1337a12cb38dac.
//
// Solidity: event NodeRecoveryCompleted(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) WatchNodeRecoveryCompleted(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeRecoveryCompleted) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeRecoveryCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeRecoveryCompleted)
				if err := _NodeManager.contract.UnpackLog(event, "NodeRecoveryCompleted", log); err != nil {
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

// ParseNodeRecoveryCompleted is a log parse operation binding the contract event 0x787d7bc525e7c4658c64e3e456d974a1be21cc196e8162a4bf1337a12cb38dac.
//
// Solidity: event NodeRecoveryCompleted(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) ParseNodeRecoveryCompleted(log types.Log) (*NodeManagerNodeRecoveryCompleted, error) {
	event := new(NodeManagerNodeRecoveryCompleted)
	if err := _NodeManager.contract.UnpackLog(event, "NodeRecoveryCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeManagerNodeRecoveryInitiatedIterator is returned from FilterNodeRecoveryInitiated and is used to iterate over the raw logs and unpacked data for NodeRecoveryInitiated events raised by the NodeManager contract.
type NodeManagerNodeRecoveryInitiatedIterator struct {
	Event *NodeManagerNodeRecoveryInitiated // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeRecoveryInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeRecoveryInitiated)
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
		it.Event = new(NodeManagerNodeRecoveryInitiated)
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
func (it *NodeManagerNodeRecoveryInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeRecoveryInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeRecoveryInitiated represents a NodeRecoveryInitiated event raised by the NodeManager contract.
type NodeManagerNodeRecoveryInitiated struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRecoveryInitiated is a free log retrieval operation binding the contract event 0xfd385c618a1e89d01fb9a21780846793e282e8bc0b60caf6ccb3e422d543fbfb.
//
// Solidity: event NodeRecoveryInitiated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) FilterNodeRecoveryInitiated(opts *bind.FilterOpts) (*NodeManagerNodeRecoveryInitiatedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeRecoveryInitiated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeRecoveryInitiatedIterator{contract: _NodeManager.contract, event: "NodeRecoveryInitiated", logs: logs, sub: sub}, nil
}

// WatchNodeRecoveryInitiated is a free log subscription operation binding the contract event 0xfd385c618a1e89d01fb9a21780846793e282e8bc0b60caf6ccb3e422d543fbfb.
//
// Solidity: event NodeRecoveryInitiated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) WatchNodeRecoveryInitiated(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeRecoveryInitiated) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeRecoveryInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeRecoveryInitiated)
				if err := _NodeManager.contract.UnpackLog(event, "NodeRecoveryInitiated", log); err != nil {
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

// ParseNodeRecoveryInitiated is a log parse operation binding the contract event 0xfd385c618a1e89d01fb9a21780846793e282e8bc0b60caf6ccb3e422d543fbfb.
//
// Solidity: event NodeRecoveryInitiated(string _enodeId, string _orgId)
func (_NodeManager *NodeManagerFilterer) ParseNodeRecoveryInitiated(log types.Log) (*NodeManagerNodeRecoveryInitiated, error) {
	event := new(NodeManagerNodeRecoveryInitiated)
	if err := _NodeManager.contract.UnpackLog(event, "NodeRecoveryInitiated", log); err != nil {
		return nil, err
	}
	return event, nil
}

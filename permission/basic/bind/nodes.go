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

var NodeManagerParsedABI, _ = abi.JSON(strings.NewReader(NodeManagerABI))

// NodeManagerBin is the compiled bytecode used for deploying new contracts.
var NodeManagerBin = "0x608060405234801561001057600080fd5b5060405160208061250b8339810180604052602081101561003057600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556124a9806100626000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806397c07a9b1161005b57806397c07a9b1461041c578063a97a440614610439578063b81c806a146104f7578063e3b09d84146102a057610088565b80630cc501461461008d5780633f0e0e471461014d5780633f5e1a45146102a057806386bc36521461035e575b600080fd5b61014b600480360360608110156100a357600080fd5b810190602081018135600160201b8111156100bd57600080fd5b8201836020820111156100cf57600080fd5b803590602001918460018302840111600160201b831117156100f057600080fd5b919390929091602081019035600160201b81111561010d57600080fd5b82018360208201111561011f57600080fd5b803590602001918460018302840111600160201b8311171561014057600080fd5b919350915035610511565b005b6101bb6004803603602081101561016357600080fd5b810190602081018135600160201b81111561017d57600080fd5b82018360208201111561018f57600080fd5b803590602001918460018302840111600160201b831117156101b057600080fd5b509092509050610f21565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156102025781810151838201526020016101ea565b50505050905090810190601f16801561022f5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561026257818101518382015260200161024a565b50505050905090810190601f16801561028f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b61014b600480360360408110156102b657600080fd5b810190602081018135600160201b8111156102d057600080fd5b8201836020820111156102e257600080fd5b803590602001918460018302840111600160201b8311171561030357600080fd5b919390929091602081019035600160201b81111561032057600080fd5b82018360208201111561033257600080fd5b803590602001918460018302840111600160201b8311171561035357600080fd5b5090925090506111f7565b61014b6004803603604081101561037457600080fd5b810190602081018135600160201b81111561038e57600080fd5b8201836020820111156103a057600080fd5b803590602001918460018302840111600160201b831117156103c157600080fd5b919390929091602081019035600160201b8111156103de57600080fd5b8201836020820111156103f057600080fd5b803590602001918460018302840111600160201b8311171561041157600080fd5b5090925090506115d0565b6101bb6004803603602081101561043257600080fd5b5035611af2565b61014b6004803603604081101561044f57600080fd5b810190602081018135600160201b81111561046957600080fd5b82018360208201111561047b57600080fd5b803590602001918460018302840111600160201b8311171561049c57600080fd5b919390929091602081019035600160201b8111156104b957600080fd5b8201836020820111156104cb57600080fd5b803590602001918460018302840111600160201b831117156104ec57600080fd5b509092509050611c81565b6104ff61205a565b60408051918252519081900360200190f35b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561055e57600080fd5b505afa158015610572573d6000803e3d6000fd5b505050506040513d602081101561058857600080fd5b50516001600160a01b031633146105dd5760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b8381101561065a578181015183820152602001610642565b50505050905090810190601f1680156106875780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181528151602092830120865290850195909552505050016000205415156107075760408051600160e51b62461bcd02815260206004820152601e60248201527f70617373656420656e6f646520696420646f6573206e6f742065786973740000604482015290519081900360640190fd5b61077a86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a01819004810282018101909252888152925088915087908190840183828082843760009201919091525061206192505050565b15156107ba57604051600160e51b62461bcd02815260040180806020018281038252602a8152602001806123e1602a913960400191505060405180910390fd5b81600114806107c95750816002145b806107d45750816003145b806107df5750816004145b806107ea5750816005145b151561082a57604051600160e51b62461bcd02815260040180806020018281038252602681526020018061242b6026913960400191505060405180910390fd5b81600114156109aa5761087286868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121c392505050565b6002146108b75760408051600160e51b62461bcd02815260206004820152601d602482015260008051602061240b833981519152604482015290519081900360640190fd5b600360016108fa88888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122a092505050565b8154811061090457fe5b9060005260206000209060030201600201819055507fc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f19565b8160021415610b2a576109f286868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121c392505050565b600314610a375760408051600160e51b62461bcd02815260206004820152601d602482015260008051602061240b833981519152604482015290519081900360640190fd5b60026001610a7a88888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122a092505050565b81548110610a8457fe5b9060005260206000209060030201600201819055507f49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f19565b8160031415610c265760046001610b7688888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122a092505050565b81548110610b8057fe5b9060005260206000209060030201600201819055507f4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f19565b8160041415610da657610c6e86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121c392505050565b600414610cb35760408051600160e51b62461bcd02815260206004820152601d602482015260008051602061240b833981519152604482015290519081900360640190fd5b60056001610cf688888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122a092505050565b81548110610d0057fe5b9060005260206000209060030201600201819055507ffd385c618a1e89d01fb9a21780846793e282e8bc0b60caf6ccb3e422d543fbfb868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a1610f19565b610de586868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121c392505050565b600514610e2a5760408051600160e51b62461bcd02815260206004820152601d602482015260008051602061240b833981519152604482015290519081900360640190fd5b60026001610e6d88888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122a092505050565b81548110610e7757fe5b9060005260206000209060030201600201819055507f787d7bc525e7c4658c64e3e456d974a1be21cc196e8162a4bf1337a12cb38dac868686866040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15b505050505050565b606080600060026000836040516020018080602001828103825283818151815260200191508051906020019080838360005b83811015610f6b578181015183820152602001610f53565b50505050905090810190601f168015610f985780820380516001836020036101000a031916815260200191505b5092505050604051602081830303815290604052805190602001208152602001908152602001600020546000141561102857848460006040518060200160405280600081525092919082828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509699509197509195506111f0945050505050565b600061106986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122a092505050565b905060018181548110151561107a57fe5b906000526020600020906003020160010160018281548110151561109a57fe5b90600052602060002090600302016000016001838154811015156110ba57fe5b60009182526020918290206002600390920201810154845460408051601f6000196101006001861615020190931694909404918201859004850284018501905280835290928591908301828280156111535780601f1061112857610100808354040283529160200191611153565b820191906000526020600020905b81548152906001019060200180831161113657829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959850879450925084019050828280156111e15780601f106111b6576101008083540402835291602001916111e1565b820191906000526020600020905b8154815290600101906020018083116111c457829003601f168201915b50505050509150935093509350505b9250925092565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561124457600080fd5b505afa158015611258573d6000803e3d6000fd5b505050506040513d602081101561126e57600080fd5b50516001600160a01b031633146112c35760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b83811015611340578181015183820152602001611328565b50505050905090810190601f16801561136d5780820380516001836020036101000a031916815260200191505b5060408051601f1981840301815291815281516020928301208652908501959095525050500160002054156113ec5760408051600160e51b62461bcd02815260206004820152601660248201527f70617373656420656e6f64652069642065786973747300000000000000000000604482015290519081900360640190fd5b6003805460010190819055604080516020808201908152918101879052600291600091899189918190606001848480828437600081840152601f19601f8201169050808301925050509350505050604051602081830303815290604052805190602001208152602001908152602001600020819055506001604051806060016040528087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f88018190048102820181019092528681529181019190879087908190840183828082843760009201829052509385525050600260209384015250835460018101808655948252908290208351805160039093029091019261151692849290910190612348565b50602082810151805161152f9260018501920190612348565b50604082015181600201555050507f0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d858585856040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561161d57600080fd5b505afa158015611631573d6000803e3d6000fd5b505050506040513d602081101561164757600080fd5b50516001600160a01b0316331461169c5760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b83811015611719578181015183820152602001611701565b50505050905090810190601f1680156117465780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181528151602092830120865290850195909552505050016000205415156117c65760408051600160e51b62461bcd02815260206004820152601e60248201527f70617373656420656e6f646520696420646f6573206e6f742065786973740000604482015290519081900360640190fd5b61183985858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8901819004810282018101909252878152925087915086908190840183828082843760009201919091525061206192505050565b151561187957604051600160e51b62461bcd02815260040180806020018281038252602d815260200180612451602d913960400191505060405180910390fd5b6118b885858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121c392505050565b60011461190f5760408051600160e51b62461bcd02815260206004820152601c60248201527f6e6f7468696e672070656e64696e6720666f7220617070726f76616c00000000604482015290519081900360640190fd5b600061195086868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506122a092505050565b9050600260018281548110151561196357fe5b9060005260206000209060030201600201819055507f0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d6001828154811015156119a857fe5b90600052602060002090600302016000016001838154811015156119c857fe5b9060005260206000209060030201600101604051808060200180602001838103835285818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611a665780601f10611a3b57610100808354040283529160200191611a66565b820191906000526020600020905b815481529060010190602001808311611a4957829003601f168201915b5050838103825284546002600019610100600184161502019091160480825260209091019085908015611ada5780601f10611aaf57610100808354040283529160200191611ada565b820191906000526020600020905b815481529060010190602001808311611abd57829003601f168201915b505094505050505060405180910390a1505050505050565b6060806000600184815481101515611b0657fe5b9060005260206000209060030201600101600185815481101515611b2657fe5b9060005260206000209060030201600001600186815481101515611b4657fe5b60009182526020918290206002600390920201810154845460408051601f600019610100600186161502019093169490940491820185900485028401850190528083529092859190830182828015611bdf5780601f10611bb457610100808354040283529160200191611bdf565b820191906000526020600020905b815481529060010190602001808311611bc257829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815295985087945092508401905082828015611c6d5780601f10611c4257610100808354040283529160200191611c6d565b820191906000526020600020905b815481529060010190602001808311611c5057829003601f168201915b505050505091509250925092509193909250565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b158015611cce57600080fd5b505afa158015611ce2573d6000803e3d6000fd5b505050506040513d6020811015611cf857600080fd5b50516001600160a01b03163314611d4d5760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060408051602080820181815288519383019390935287516002975093955087945091928392606090920191850190808383895b83811015611dca578181015183820152602001611db2565b50505050905090810190601f168015611df75780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181528151602092830120865290850195909552505050016000205415611e765760408051600160e51b62461bcd02815260206004820152601660248201527f70617373656420656e6f64652069642065786973747300000000000000000000604482015290519081900360640190fd5b6003805460010190819055604080516020808201908152918101879052600291600091899189918190606001848480828437600081840152601f19601f8201169050808301925050509350505050604051602081830303815290604052805190602001208152602001908152602001600020819055506001604051806060016040528087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f880181900481028201810190925286815291810191908790879081908401838280828437600092018290525093855250506001602093840181905285549081018087559583529183902084518051600390940290910193611fa093859350910190612348565b506020828101518051611fb99260018501920190612348565b50604082015181600201555050507fb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f858585856040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050565b6003545b90565b6000816040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156120a457818101518382015260200161208c565b50505050905090810190601f1680156120d15780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040528051906020012060016120f7856122a0565b8154811061210157fe5b9060005260206000209060030201600101604051602001808060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561219e5780601f106121735761010080835404028352916020019161219e565b820191906000526020600020905b81548152906001019060200180831161218157829003601f168201915b5050925050506040516020818303038152906040528051906020012014905092915050565b600060026000836040516020018080602001828103825283818151815260200191508051906020019080838360005b8381101561220a5781810151838201526020016121f2565b50505050905090810190601f1680156122375780820380516001836020036101000a031916815260200191505b509250505060405160208183030381529060405280519060200120815260200190815260200160002054600014156122715750600061229b565b600161227c836122a0565b8154811061228657fe5b90600052602060002090600302016002015490505b919050565b6000600160026000846040516020018080602001828103825283818151815260200191508051906020019080838360005b838110156122e95781810151838201526020016122d1565b50505050905090810190601f1680156123165780820380516001836020036101000a031916815260200191505b509250505060405160208183030381529060405280519060200120815260200190815260200160002054039050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061238957805160ff19168380011785556123b6565b828001600101855582156123b6579182015b828111156123b657825182559160200191906001019061239b565b506123c29291506123c6565b5090565b61205e91905b808211156123c257600081556001016123cc56fe656e6f646520696420646f6573206e6f742062656c6f6e6720746f2074686520706173736564206f72676f7065726174696f6e2063616e6e6f7420626520706572666f726d6564000000696e76616c6964206f7065726174696f6e2e2077726f6e6720616374696f6e20706173736564656e6f646520696420646f6573206e6f742062656c6f6e6720746f2074686520706173736564206f7267206964a165627a7a723058207ca0dd787547cf61d1f16df314986310b2a2c8f853fdca9e4a4c784046b0864c0029"

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

var NodeActivatedTopicHash = "0x49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f"

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

var NodeApprovedTopicHash = "0x0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d"

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

var NodeBlacklistedTopicHash = "0x4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a"

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

var NodeDeactivatedTopicHash = "0xc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a"

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

var NodeProposedTopicHash = "0xb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f"

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

var NodeRecoveryCompletedTopicHash = "0x787d7bc525e7c4658c64e3e456d974a1be21cc196e8162a4bf1337a12cb38dac"

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

var NodeRecoveryInitiatedTopicHash = "0xfd385c618a1e89d01fb9a21780846793e282e8bc0b60caf6ccb3e422d543fbfb"

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

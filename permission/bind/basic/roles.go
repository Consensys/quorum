// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic

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

// RoleManagerABI is the input ABI used to generate the binding from.
const RoleManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getRoleDetails\",\"outputs\":[{\"name\":\"roleId\",\"type\":\"string\"},{\"name\":\"orgId\",\"type\":\"string\"},{\"name\":\"accessType\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"admin\",\"type\":\"bool\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"name\":\"_isVoter\",\"type\":\"bool\"},{\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"addRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfRoles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_rIndex\",\"type\":\"uint256\"}],\"name\":\"getRoleDetailsFromIndex\",\"outputs\":[{\"name\":\"roleId\",\"type\":\"string\"},{\"name\":\"orgId\",\"type\":\"string\"},{\"name\":\"accessType\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"admin\",\"type\":\"bool\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"roleExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"isAdminRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"isVoterRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_isVoter\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"RoleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"}]"

// RoleManagerBin is the compiled bytecode used for deploying new contracts.
var RoleManagerBin = "0x608060405234801561001057600080fd5b5060405160208061221f8339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a03199092169190911790556121bd806100626000396000f3fe608060405234801561001057600080fd5b506004361061008c5760003560e060020a90048063a63430121161005f578063a6343012146103c6578063abf5739f14610488578063be322e5414610650578063deb16ba7146107645761008c565b80631870aba3146100915780637b7135791461025157806387f55d311461038f578063a451d4a8146103a9575b600080fd5b610153600480360360408110156100a757600080fd5b8101906020810181356401000000008111156100c257600080fd5b8201836020820111156100d457600080fd5b803590602001918460018302840111640100000000831117156100f657600080fd5b91939092909160208101903564010000000081111561011457600080fd5b82018360208201111561012657600080fd5b8035906020019184600183028401116401000000008311171561014857600080fd5b509092509050610878565b604080519081018590528315156060820152821515608082015281151560a082015260c08082528751908201528651819060208083019160e08401918b019080838360005b838110156101b0578181015183820152602001610198565b50505050905090810190601f1680156101dd5780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b838110156102105781810151838201526020016101f8565b50505050905090810190601f16801561023d5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b61038d600480360360a081101561026757600080fd5b81019060208101813564010000000081111561028257600080fd5b82018360208201111561029457600080fd5b803590602001918460018302840111640100000000831117156102b657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561030957600080fd5b82018360208201111561031b57600080fd5b8035906020019184600183028401116401000000008311171561033d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135151590604001351515610bfe565b005b610397611149565b60408051918252519081900360200190f35b610153600480360360208110156103bf57600080fd5b5035611150565b61038d600480360360408110156103dc57600080fd5b8101906020810181356401000000008111156103f757600080fd5b82018360208201111561040957600080fd5b8035906020019184600183028401116401000000008311171561042b57600080fd5b91939092909160208101903564010000000081111561044957600080fd5b82018360208201111561045b57600080fd5b8035906020019184600183028401116401000000008311171561047d57600080fd5b50909250905061136e565b61063c6004803603606081101561049e57600080fd5b8101906020810181356401000000008111156104b957600080fd5b8201836020820111156104cb57600080fd5b803590602001918460018302840111640100000000831117156104ed57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561054057600080fd5b82018360208201111561055257600080fd5b8035906020019184600183028401116401000000008311171561057457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156105c757600080fd5b8201836020820111156105d957600080fd5b803590602001918460018302840111640100000000831117156105fb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611666945050505050565b604080519115158252519081900360200190f35b61063c6004803603606081101561066657600080fd5b81019060208101813564010000000081111561068157600080fd5b82018360208201111561069357600080fd5b803590602001918460018302840111640100000000831117156106b557600080fd5b9193909290916020810190356401000000008111156106d357600080fd5b8201836020820111156106e557600080fd5b8035906020019184600183028401116401000000008311171561070757600080fd5b91939092909160208101903564010000000081111561072557600080fd5b82018360208201111561073757600080fd5b8035906020019184600183028401116401000000008311171561075957600080fd5b5090925090506118da565b61063c6004803603606081101561077a57600080fd5b81019060208101813564010000000081111561079557600080fd5b8201836020820111156107a757600080fd5b803590602001918460018302840111640100000000831117156107c957600080fd5b9193909290916020810190356401000000008111156107e757600080fd5b8201836020820111156107f957600080fd5b8035906020019184600183028401116401000000008311171561081b57600080fd5b91939092909160208101903564010000000081111561083957600080fd5b82018360208201111561084b57600080fd5b8035906020019184600183028401116401000000008311171561086d57600080fd5b509092509050611c57565b6060806000806000806109028a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b9081908401838280828437600092018290525060408051602081019091529081529250611666915050565b151561096c57898960008060008085858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052506040805160208101909152908152939f50929d50959b509399509197509550610bf1945050505050565b60006109e18b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8f018190048102820181019092528d815292508d91508c9081908401838280828437600092019190915250611fc992505050565b90506001818154811015156109f257fe5b9060005260206000209060040201600001600182815481101515610a1257fe5b9060005260206000209060040201600101600183815481101515610a3257fe5b906000526020600020906004020160020154600184815481101515610a5357fe5b60009182526020909120600360049092020101546001805460ff9092169186908110610a7b57fe5b906000526020600020906004020160030160019054906101000a900460ff16600186815481101515610aa957fe5b6000918252602091829020600491909102016003015486546040805160026101006001851615026000190190931692909204601f81018590048502830185019091528082526201000090920460ff169290918891830182828015610b4e5780601f10610b2357610100808354040283529160200191610b4e565b820191906000526020600020905b815481529060010190602001808311610b3157829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a945092508401905082828015610bdc5780601f10610bb157610100808354040283529160200191610bdc565b820191906000526020600020905b815481529060010190602001808311610bbf57829003601f168201915b50505050509450965096509650965096509650505b9499939850945094509450565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610c4e57600080fd5b505afa158015610c62573d6000803e3d6000fd5b505050506040513d6020811015610c7857600080fd5b5051600160a060020a03163314610cc7576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612172833981519152604482015290519081900360640190fd5b600260008686604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610d11578181015183820152602001610cf9565b50505050905090810190601f168015610d3e5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610d71578181015183820152602001610d59565b50505050905090810190601f168015610d9e5780820380516001836020036101000a031916815260200191505b50945050505050604051602081830303815290604052805190602001208152602001908152602001600020546000141515610e23576040805160e560020a62461bcd02815260206004820152601760248201527f726f6c652065786973747320666f7220746865206f7267000000000000000000604482015290519081900360640190fd5b60038054600101908190556040805160208082018381528951606084015289516002946000948c948c94938493830192608001918701908083838b5b83811015610e77578181015183820152602001610e5f565b50505050905090810190601f168015610ea45780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610ed7578181015183820152602001610ebf565b50505050905090810190601f168015610f045780820380516001836020036101000a031916815260200191505b5060408051601f1981840301815291815281516020928301208852878201989098529587016000908120989098555050845160c0810186528b81528085018b905294850189905250505084151560608301528315156080830152600160a083018190528054808201808355919094528251805191946004027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60192610fae928492909101906120d9565b506020828101518051610fc792600185019201906120d9565b5060408281015160028301556060808401516003909301805460808087015160a09788015160ff199093169615159690961761ff001916610100961515969096029590951762ff0000191662010000911515919091021790558151918201889052861515908201528415159181019190915281815287519181019190915286517fefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c92508791879187918791879190819060208083019160c08401918a019080838360005b838110156110a357818101518382015260200161108b565b50505050905090810190601f1680156110d05780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b838110156111035781810151838201526020016110eb565b50505050905090810190601f1680156111305780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15050505050565b6001545b90565b60608060008060008060018781548110151561116857fe5b906000526020600020906004020160000160018881548110151561118857fe5b90600052602060002090600402016001016001898154811015156111a857fe5b90600052602060002090600402016002015460018a8154811015156111c957fe5b60009182526020909120600360049092020101546001805460ff909216918c9081106111f157fe5b906000526020600020906004020160030160019054906101000a900460ff1660018c81548110151561121f57fe5b6000918252602091829020600491909102016003015486546040805160026101006001851615026000190190931692909204601f81018590048502830185019091528082526201000090920460ff1692909188918301828280156112c45780601f10611299576101008083540402835291602001916112c4565b820191906000526020600020905b8154815290600101906020018083116112a757829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a9450925084019050828280156113525780601f1061132757610100808354040283529160200191611352565b820191906000526020600020905b81548152906001019060200180831161133557829003601f168201915b5050505050945095509550955095509550955091939550919395565b6000809054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156113be57600080fd5b505afa1580156113d2573d6000803e3d6000fd5b505050506040513d60208110156113e857600080fd5b5051600160a060020a03163314611437576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612172833981519152604482015290519081900360640190fd5b60026000858585856040516020018080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f820116905080830192505050965050505050505060405160208183030381529060405280519060200120815260200190815260200160002054600014151515611523576040805160e560020a62461bcd02815260206004820152601360248201527f726f6c6520646f6573206e6f7420657869737400000000000000000000000000604482015290519081900360640190fd5b600061159885858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f89018190048102820181019092528781529250879150869081908401838280828437600092019190915250611fc992505050565b905060006001828154811015156115ab57fe5b906000526020600020906004020160030160026101000a81548160ff0219169083151502179055507f1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6858585856040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050565b600080600260008686604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b838110156116b357818101518382015260200161169b565b50505050905090810190601f1680156116e05780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156117135781810151838201526020016116fb565b50505050905090810190601f1680156117405780820380516001836020036101000a031916815260200191505b509450505050506040516020818303038152906040528051906020012081526020019081526020016000205460001415156117b75761177f8585611fc9565b905060018181548110151561179057fe5b906000526020600020906004020160030160029054906101000a900460ff169150506118d3565b600260008685604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b838110156118015781810151838201526020016117e9565b50505050905090810190601f16801561182e5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015611861578181015183820152602001611849565b50505050905090810190601f16801561188e5780820380516001836020036101000a031916815260200191505b509450505050506040516020818303038152906040528051906020012081526020019081526020016000205460001415156118cd5761177f8584611fc9565b60009150505b9392505050565b60008060009054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561192c57600080fd5b505afa158015611940573d6000803e3d6000fd5b505050506040513d602081101561195657600080fd5b5051600160a060020a031633146119a5576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612172833981519152604482015290519081900360640190fd5b611a4c87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b01819004810282018101909252898152925089915088908190840183828082843760009201919091525050604080516020601f8a01819004810282018101909252888152925088915087908190840183828082843760009201919091525061166692505050565b1515611a5a57506000611c4d565b600060026000898989896040516020018080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f8201169050808301925050509650505050505050604051602081830303815290604052805190602001208152602001908152602001600020546000141515611b7157611b6a88888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a9150899081908401838280828437600092019190915250611fc992505050565b9050611be7565b611be488888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a018190048102820181019092528881529250889150879081908401838280828437600092019190915250611fc992505050565b90505b6001805482908110611bf557fe5b906000526020600020906004020160030160029054906101000a900460ff168015611c4957506001805482908110611c2957fe5b906000526020600020906004020160030160019054906101000a900460ff165b9150505b9695505050505050565b60008060009054906101000a9004600160a060020a0316600160a060020a0316630e32cf906040518163ffffffff1660e060020a02815260040160206040518083038186803b158015611ca957600080fd5b505afa158015611cbd573d6000803e3d6000fd5b505050506040513d6020811015611cd357600080fd5b5051600160a060020a03163314611d22576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020612172833981519152604482015290519081900360640190fd5b611dc987878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b01819004810282018101909252898152925089915088908190840183828082843760009201919091525050604080516020601f8a01819004810282018101909252888152925088915087908190840183828082843760009201919091525061166692505050565b1515611dd757506000611c4d565b600060026000898989896040516020018080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f8201169050808301925050509650505050505050604051602081830303815290604052805190602001208152602001908152602001600020546000141515611eee57611ee788888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a9150899081908401838280828437600092019190915250611fc992505050565b9050611f64565b611f6188888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a018190048102820181019092528881529250889150879081908401838280828437600092019190915250611fc992505050565b90505b6001805482908110611f7257fe5b906000526020600020906004020160030160029054906101000a900460ff168015611c4957506001805482908110611fa657fe5b600091825260209091206004909102016003015460ff1698975050505050505050565b60006001600260008585604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b83811015612017578181015183820152602001611fff565b50505050905090810190601f1680156120445780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561207757818101518382015260200161205f565b50505050905090810190601f1680156120a45780820380516001836020036101000a031916815260200191505b509450505050506040516020818303038152906040528051906020012081526020019081526020016000205403905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061211a57805160ff1916838001178555612147565b82800160010185558215612147579182015b8281111561214757825182559160200191906001019061212c565b50612153929150612157565b5090565b61114d91905b80821115612153576000815560010161215d56fe696e76616c69642063616c6c6572000000000000000000000000000000000000a165627a7a72305820f8334f4c837fc860658ba7f6cbdcd114ed4d2ad6e79527389616d6999c588de60029"

// DeployRoleManager deploys a new Ethereum contract, binding an instance of RoleManager to it.
func DeployRoleManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permUpgradable common.Address) (common.Address, *types.Transaction, *RoleManager, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RoleManagerBin), backend, _permUpgradable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoleManager{RoleManagerCaller: RoleManagerCaller{contract: contract}, RoleManagerTransactor: RoleManagerTransactor{contract: contract}, RoleManagerFilterer: RoleManagerFilterer{contract: contract}}, nil
}

// RoleManager is an auto generated Go binding around an Ethereum contract.
type RoleManager struct {
	RoleManagerCaller     // Read-only binding to the contract
	RoleManagerTransactor // Write-only binding to the contract
	RoleManagerFilterer   // Log filterer for contract events
}

// RoleManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleManagerSession struct {
	Contract     *RoleManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleManagerCallerSession struct {
	Contract *RoleManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RoleManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleManagerTransactorSession struct {
	Contract     *RoleManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RoleManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleManagerRaw struct {
	Contract *RoleManager // Generic contract binding to access the raw methods on
}

// RoleManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleManagerCallerRaw struct {
	Contract *RoleManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RoleManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleManagerTransactorRaw struct {
	Contract *RoleManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleManager creates a new instance of RoleManager, bound to a specific deployed contract.
func NewRoleManager(address common.Address, backend bind.ContractBackend) (*RoleManager, error) {
	contract, err := bindRoleManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleManager{RoleManagerCaller: RoleManagerCaller{contract: contract}, RoleManagerTransactor: RoleManagerTransactor{contract: contract}, RoleManagerFilterer: RoleManagerFilterer{contract: contract}}, nil
}

// NewRoleManagerCaller creates a new read-only instance of RoleManager, bound to a specific deployed contract.
func NewRoleManagerCaller(address common.Address, caller bind.ContractCaller) (*RoleManagerCaller, error) {
	contract, err := bindRoleManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagerCaller{contract: contract}, nil
}

// NewRoleManagerTransactor creates a new write-only instance of RoleManager, bound to a specific deployed contract.
func NewRoleManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleManagerTransactor, error) {
	contract, err := bindRoleManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagerTransactor{contract: contract}, nil
}

// NewRoleManagerFilterer creates a new log filterer instance of RoleManager, bound to a specific deployed contract.
func NewRoleManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleManagerFilterer, error) {
	contract, err := bindRoleManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleManagerFilterer{contract: contract}, nil
}

// bindRoleManager binds a generic wrapper to an already deployed contract.
func bindRoleManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManager *RoleManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RoleManager.Contract.RoleManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManager *RoleManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManager.Contract.RoleManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManager *RoleManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManager.Contract.RoleManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManager *RoleManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RoleManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManager *RoleManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManager *RoleManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManager.Contract.contract.Transact(opts, method, params...)
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() constant returns(uint256)
func (_RoleManager *RoleManagerCaller) GetNumberOfRoles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "getNumberOfRoles")
	return *ret0, err
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() constant returns(uint256)
func (_RoleManager *RoleManagerSession) GetNumberOfRoles() (*big.Int, error) {
	return _RoleManager.Contract.GetNumberOfRoles(&_RoleManager.CallOpts)
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() constant returns(uint256)
func (_RoleManager *RoleManagerCallerSession) GetNumberOfRoles() (*big.Int, error) {
	return _RoleManager.Contract.GetNumberOfRoles(&_RoleManager.CallOpts)
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(string _roleId, string _orgId) constant returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManager *RoleManagerCaller) GetRoleDetails(opts *bind.CallOpts, _roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	ret := new(struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Admin      bool
		Active     bool
	})
	out := ret
	err := _RoleManager.contract.Call(opts, out, "getRoleDetails", _roleId, _orgId)
	return *ret, err
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(string _roleId, string _orgId) constant returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManager *RoleManagerSession) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetails(&_RoleManager.CallOpts, _roleId, _orgId)
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(string _roleId, string _orgId) constant returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManager *RoleManagerCallerSession) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetails(&_RoleManager.CallOpts, _roleId, _orgId)
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(uint256 _rIndex) constant returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManager *RoleManagerCaller) GetRoleDetailsFromIndex(opts *bind.CallOpts, _rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	ret := new(struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Admin      bool
		Active     bool
	})
	out := ret
	err := _RoleManager.contract.Call(opts, out, "getRoleDetailsFromIndex", _rIndex)
	return *ret, err
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(uint256 _rIndex) constant returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManager *RoleManagerSession) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetailsFromIndex(&_RoleManager.CallOpts, _rIndex)
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(uint256 _rIndex) constant returns(string roleId, string orgId, uint256 accessType, bool voter, bool admin, bool active)
func (_RoleManager *RoleManagerCallerSession) GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetailsFromIndex(&_RoleManager.CallOpts, _rIndex)
}

// IsAdminRole is a free data retrieval call binding the contract method 0xbe322e54.
//
// Solidity: function isAdminRole(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerCaller) IsAdminRole(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "isAdminRole", _roleId, _orgId, _ultParent)
	return *ret0, err
}

// IsAdminRole is a free data retrieval call binding the contract method 0xbe322e54.
//
// Solidity: function isAdminRole(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerSession) IsAdminRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManager.Contract.IsAdminRole(&_RoleManager.CallOpts, _roleId, _orgId, _ultParent)
}

// IsAdminRole is a free data retrieval call binding the contract method 0xbe322e54.
//
// Solidity: function isAdminRole(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerCallerSession) IsAdminRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManager.Contract.IsAdminRole(&_RoleManager.CallOpts, _roleId, _orgId, _ultParent)
}

// IsVoterRole is a free data retrieval call binding the contract method 0xdeb16ba7.
//
// Solidity: function isVoterRole(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerCaller) IsVoterRole(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "isVoterRole", _roleId, _orgId, _ultParent)
	return *ret0, err
}

// IsVoterRole is a free data retrieval call binding the contract method 0xdeb16ba7.
//
// Solidity: function isVoterRole(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerSession) IsVoterRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManager.Contract.IsVoterRole(&_RoleManager.CallOpts, _roleId, _orgId, _ultParent)
}

// IsVoterRole is a free data retrieval call binding the contract method 0xdeb16ba7.
//
// Solidity: function isVoterRole(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerCallerSession) IsVoterRole(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManager.Contract.IsVoterRole(&_RoleManager.CallOpts, _roleId, _orgId, _ultParent)
}

// RoleExists is a free data retrieval call binding the contract method 0xabf5739f.
//
// Solidity: function roleExists(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerCaller) RoleExists(opts *bind.CallOpts, _roleId string, _orgId string, _ultParent string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "roleExists", _roleId, _orgId, _ultParent)
	return *ret0, err
}

// RoleExists is a free data retrieval call binding the contract method 0xabf5739f.
//
// Solidity: function roleExists(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerSession) RoleExists(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManager.Contract.RoleExists(&_RoleManager.CallOpts, _roleId, _orgId, _ultParent)
}

// RoleExists is a free data retrieval call binding the contract method 0xabf5739f.
//
// Solidity: function roleExists(string _roleId, string _orgId, string _ultParent) constant returns(bool)
func (_RoleManager *RoleManagerCallerSession) RoleExists(_roleId string, _orgId string, _ultParent string) (bool, error) {
	return _RoleManager.Contract.RoleExists(&_RoleManager.CallOpts, _roleId, _orgId, _ultParent)
}

// AddRole is a paid mutator transaction binding the contract method 0x7b713579.
//
// Solidity: function addRole(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin) returns()
func (_RoleManager *RoleManagerTransactor) AddRole(opts *bind.TransactOpts, _roleId string, _orgId string, _baseAccess *big.Int, _isVoter bool, _isAdmin bool) (*types.Transaction, error) {
	return _RoleManager.contract.Transact(opts, "addRole", _roleId, _orgId, _baseAccess, _isVoter, _isAdmin)
}

// AddRole is a paid mutator transaction binding the contract method 0x7b713579.
//
// Solidity: function addRole(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin) returns()
func (_RoleManager *RoleManagerSession) AddRole(_roleId string, _orgId string, _baseAccess *big.Int, _isVoter bool, _isAdmin bool) (*types.Transaction, error) {
	return _RoleManager.Contract.AddRole(&_RoleManager.TransactOpts, _roleId, _orgId, _baseAccess, _isVoter, _isAdmin)
}

// AddRole is a paid mutator transaction binding the contract method 0x7b713579.
//
// Solidity: function addRole(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin) returns()
func (_RoleManager *RoleManagerTransactorSession) AddRole(_roleId string, _orgId string, _baseAccess *big.Int, _isVoter bool, _isAdmin bool) (*types.Transaction, error) {
	return _RoleManager.Contract.AddRole(&_RoleManager.TransactOpts, _roleId, _orgId, _baseAccess, _isVoter, _isAdmin)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_RoleManager *RoleManagerTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManager.contract.Transact(opts, "removeRole", _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_RoleManager *RoleManagerSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManager.Contract.RemoveRole(&_RoleManager.TransactOpts, _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_RoleManager *RoleManagerTransactorSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManager.Contract.RemoveRole(&_RoleManager.TransactOpts, _roleId, _orgId)
}

// RoleManagerRoleCreatedIterator is returned from FilterRoleCreated and is used to iterate over the raw logs and unpacked data for RoleCreated events raised by the RoleManager contract.
type RoleManagerRoleCreatedIterator struct {
	Event *RoleManagerRoleCreated // Event containing the contract specifics and raw log

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
func (it *RoleManagerRoleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagerRoleCreated)
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
		it.Event = new(RoleManagerRoleCreated)
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
func (it *RoleManagerRoleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagerRoleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagerRoleCreated represents a RoleCreated event raised by the RoleManager contract.
type RoleManagerRoleCreated struct {
	RoleId     string
	OrgId      string
	BaseAccess *big.Int
	IsVoter    bool
	IsAdmin    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoleCreated is a free log retrieval operation binding the contract event 0xefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c.
//
// Solidity: event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin)
func (_RoleManager *RoleManagerFilterer) FilterRoleCreated(opts *bind.FilterOpts) (*RoleManagerRoleCreatedIterator, error) {

	logs, sub, err := _RoleManager.contract.FilterLogs(opts, "RoleCreated")
	if err != nil {
		return nil, err
	}
	return &RoleManagerRoleCreatedIterator{contract: _RoleManager.contract, event: "RoleCreated", logs: logs, sub: sub}, nil
}

// WatchRoleCreated is a free log subscription operation binding the contract event 0xefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c.
//
// Solidity: event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin)
func (_RoleManager *RoleManagerFilterer) WatchRoleCreated(opts *bind.WatchOpts, sink chan<- *RoleManagerRoleCreated) (event.Subscription, error) {

	logs, sub, err := _RoleManager.contract.WatchLogs(opts, "RoleCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagerRoleCreated)
				if err := _RoleManager.contract.UnpackLog(event, "RoleCreated", log); err != nil {
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

// ParseRoleCreated is a log parse operation binding the contract event 0xefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c.
//
// Solidity: event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess, bool _isVoter, bool _isAdmin)
func (_RoleManager *RoleManagerFilterer) ParseRoleCreated(log types.Log) (*RoleManagerRoleCreated, error) {
	event := new(RoleManagerRoleCreated)
	if err := _RoleManager.contract.UnpackLog(event, "RoleCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RoleManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RoleManager contract.
type RoleManagerRoleRevokedIterator struct {
	Event *RoleManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoleManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagerRoleRevoked)
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
		it.Event = new(RoleManagerRoleRevoked)
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
func (it *RoleManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagerRoleRevoked represents a RoleRevoked event raised by the RoleManager contract.
type RoleManagerRoleRevoked struct {
	RoleId string
	OrgId  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: event RoleRevoked(string _roleId, string _orgId)
func (_RoleManager *RoleManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts) (*RoleManagerRoleRevokedIterator, error) {

	logs, sub, err := _RoleManager.contract.FilterLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return &RoleManagerRoleRevokedIterator{contract: _RoleManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: event RoleRevoked(string _roleId, string _orgId)
func (_RoleManager *RoleManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoleManagerRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _RoleManager.contract.WatchLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagerRoleRevoked)
				if err := _RoleManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: event RoleRevoked(string _roleId, string _orgId)
func (_RoleManager *RoleManagerFilterer) ParseRoleRevoked(log types.Log) (*RoleManagerRoleRevoked, error) {
	event := new(RoleManagerRoleRevoked)
	if err := _RoleManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

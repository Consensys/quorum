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

// RoleManagerABI is the input ABI used to generate the binding from.
const RoleManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getRoleDetails\",\"outputs\":[{\"name\":\"roleId\",\"type\":\"string\"},{\"name\":\"orgId\",\"type\":\"string\"},{\"name\":\"accessType\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"admin\",\"type\":\"bool\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"name\":\"_isVoter\",\"type\":\"bool\"},{\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"addRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfRoles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_rIndex\",\"type\":\"uint256\"}],\"name\":\"getRoleDetailsFromIndex\",\"outputs\":[{\"name\":\"roleId\",\"type\":\"string\"},{\"name\":\"orgId\",\"type\":\"string\"},{\"name\":\"accessType\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"admin\",\"type\":\"bool\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"roleExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"isAdminRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"isVoterRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_isVoter\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"RoleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"}]"

var RoleManagerParsedABI, _ = abi.JSON(strings.NewReader(RoleManagerABI))

// RoleManagerBin is the compiled bytecode used for deploying new contracts.
var RoleManagerBin = "0x608060405234801561001057600080fd5b506040516020806122418339810180604052602081101561003057600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556121df806100626000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a63430121161005b578063a6343012146103ba578063abf5739f14610478578063be322e541461063a578063deb16ba71461074857610088565b80631870aba31461008d5780637b7135791461024957806387f55d3114610383578063a451d4a81461039d575b600080fd5b61014b600480360360408110156100a357600080fd5b810190602081018135600160201b8111156100bd57600080fd5b8201836020820111156100cf57600080fd5b803590602001918460018302840111600160201b831117156100f057600080fd5b919390929091602081019035600160201b81111561010d57600080fd5b82018360208201111561011f57600080fd5b803590602001918460018302840111600160201b8311171561014057600080fd5b509092509050610856565b604080519081018590528315156060820152821515608082015281151560a082015260c08082528751908201528651819060208083019160e08401918b019080838360005b838110156101a8578181015183820152602001610190565b50505050905090810190601f1680156101d55780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b838110156102085781810151838201526020016101f0565b50505050905090810190601f1680156102355780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b610381600480360360a081101561025f57600080fd5b810190602081018135600160201b81111561027957600080fd5b82018360208201111561028b57600080fd5b803590602001918460018302840111600160201b831117156102ac57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102fe57600080fd5b82018360208201111561031057600080fd5b803590602001918460018302840111600160201b8311171561033157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135151590604001351515610bdc565b005b61038b611182565b60408051918252519081900360200190f35b61014b600480360360208110156103b357600080fd5b5035611189565b610381600480360360408110156103d057600080fd5b810190602081018135600160201b8111156103ea57600080fd5b8201836020820111156103fc57600080fd5b803590602001918460018302840111600160201b8311171561041d57600080fd5b919390929091602081019035600160201b81111561043a57600080fd5b82018360208201111561044c57600080fd5b803590602001918460018302840111600160201b8311171561046d57600080fd5b5090925090506113a7565b6106266004803603606081101561048e57600080fd5b810190602081018135600160201b8111156104a857600080fd5b8201836020820111156104ba57600080fd5b803590602001918460018302840111600160201b831117156104db57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561052d57600080fd5b82018360208201111561053f57600080fd5b803590602001918460018302840111600160201b8311171561056057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156105b257600080fd5b8201836020820111156105c457600080fd5b803590602001918460018302840111600160201b831117156105e557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506116a2945050505050565b604080519115158252519081900360200190f35b6106266004803603606081101561065057600080fd5b810190602081018135600160201b81111561066a57600080fd5b82018360208201111561067c57600080fd5b803590602001918460018302840111600160201b8311171561069d57600080fd5b919390929091602081019035600160201b8111156106ba57600080fd5b8201836020820111156106cc57600080fd5b803590602001918460018302840111600160201b831117156106ed57600080fd5b919390929091602081019035600160201b81111561070a57600080fd5b82018360208201111561071c57600080fd5b803590602001918460018302840111600160201b8311171561073d57600080fd5b509092509050611916565b6106266004803603606081101561075e57600080fd5b810190602081018135600160201b81111561077857600080fd5b82018360208201111561078a57600080fd5b803590602001918460018302840111600160201b831117156107ab57600080fd5b919390929091602081019035600160201b8111156107c857600080fd5b8201836020820111156107da57600080fd5b803590602001918460018302840111600160201b831117156107fb57600080fd5b919390929091602081019035600160201b81111561081857600080fd5b82018360208201111561082a57600080fd5b803590602001918460018302840111600160201b8311171561084b57600080fd5b509092509050611c96565b6060806000806000806108e08a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b90819084018382808284376000920182905250604080516020810190915290815292506116a2915050565b151561094a57898960008060008085858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052506040805160208101909152908152939f50929d50959b509399509197509550610bcf945050505050565b60006109bf8b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8f018190048102820181019092528d815292508d91508c908190840183828082843760009201919091525061200b92505050565b90506001818154811015156109d057fe5b90600052602060002090600402016000016001828154811015156109f057fe5b9060005260206000209060040201600101600183815481101515610a1057fe5b906000526020600020906004020160020154600184815481101515610a3157fe5b60009182526020909120600360049092020101546001805460ff9092169186908110610a5957fe5b906000526020600020906004020160030160019054906101000a900460ff16600186815481101515610a8757fe5b6000918252602091829020600491909102016003015486546040805160026101006001851615026000190190931692909204601f81018590048502830185019091528082526201000090920460ff169290918891830182828015610b2c5780601f10610b0157610100808354040283529160200191610b2c565b820191906000526020600020905b815481529060010190602001808311610b0f57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a945092508401905082828015610bba5780601f10610b8f57610100808354040283529160200191610bba565b820191906000526020600020905b815481529060010190602001808311610b9d57829003601f168201915b50505050509450965096509650965096509650505b9499939850945094509450565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b158015610c2957600080fd5b505afa158015610c3d573d6000803e3d6000fd5b505050506040513d6020811015610c5357600080fd5b50516001600160a01b03163314610ca85760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b60048310610d005760408051600160e51b62461bcd02815260206004820152601460248201527f696e76616c6964206163636573732076616c7565000000000000000000000000604482015290519081900360640190fd5b600260008686604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610d4a578181015183820152602001610d32565b50505050905090810190601f168015610d775780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610daa578181015183820152602001610d92565b50505050905090810190601f168015610dd75780820380516001836020036101000a031916815260200191505b50945050505050604051602081830303815290604052805190602001208152602001908152602001600020546000141515610e5c5760408051600160e51b62461bcd02815260206004820152601760248201527f726f6c652065786973747320666f7220746865206f7267000000000000000000604482015290519081900360640190fd5b60038054600101908190556040805160208082018381528951606084015289516002946000948c948c94938493830192608001918701908083838b5b83811015610eb0578181015183820152602001610e98565b50505050905090810190601f168015610edd5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610f10578181015183820152602001610ef8565b50505050905090810190601f168015610f3d5780820380516001836020036101000a031916815260200191505b5060408051601f1981840301815291815281516020928301208852878201989098529587016000908120989098555050845160c0810186528b81528085018b905294850189905250505084151560608301528315156080830152600160a083018190528054808201808355919094528251805191946004027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60192610fe79284929091019061211b565b506020828101518051611000926001850192019061211b565b5060408281015160028301556060808401516003909301805460808087015160a09788015160ff199093169615159690961761ff001916610100961515969096029590951762ff0000191662010000911515919091021790558151918201889052861515908201528415159181019190915281815287519181019190915286517fefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c92508791879187918791879190819060208083019160c08401918a019080838360005b838110156110dc5781810151838201526020016110c4565b50505050905090810190601f1680156111095780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b8381101561113c578181015183820152602001611124565b50505050905090810190601f1680156111695780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15050505050565b6001545b90565b6060806000806000806001878154811015156111a157fe5b90600052602060002090600402016000016001888154811015156111c157fe5b90600052602060002090600402016001016001898154811015156111e157fe5b90600052602060002090600402016002015460018a81548110151561120257fe5b60009182526020909120600360049092020101546001805460ff909216918c90811061122a57fe5b906000526020600020906004020160030160019054906101000a900460ff1660018c81548110151561125857fe5b6000918252602091829020600491909102016003015486546040805160026101006001851615026000190190931692909204601f81018590048502830185019091528082526201000090920460ff1692909188918301828280156112fd5780601f106112d2576101008083540402835291602001916112fd565b820191906000526020600020905b8154815290600101906020018083116112e057829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a94509250840190508282801561138b5780601f106113605761010080835404028352916020019161138b565b820191906000526020600020905b81548152906001019060200180831161136e57829003601f168201915b5050505050945095509550955095509550955091939550919395565b6000809054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b1580156113f457600080fd5b505afa158015611408573d6000803e3d6000fd5b505050506040513d602081101561141e57600080fd5b50516001600160a01b031633146114735760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b60026000858585856040516020018080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f82011690508083019250505096505050505050506040516020818303038152906040528051906020012081526020019081526020016000205460001415151561155f5760408051600160e51b62461bcd02815260206004820152601360248201527f726f6c6520646f6573206e6f7420657869737400000000000000000000000000604482015290519081900360640190fd5b60006115d485858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8901819004810282018101909252878152925087915086908190840183828082843760009201919091525061200b92505050565b905060006001828154811015156115e757fe5b906000526020600020906004020160030160026101000a81548160ff0219169083151502179055507f1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6858585856040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050565b600080600260008686604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b838110156116ef5781810151838201526020016116d7565b50505050905090810190601f16801561171c5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561174f578181015183820152602001611737565b50505050905090810190601f16801561177c5780820380516001836020036101000a031916815260200191505b509450505050506040516020818303038152906040528051906020012081526020019081526020016000205460001415156117f3576117bb858561200b565b90506001818154811015156117cc57fe5b906000526020600020906004020160030160029054906101000a900460ff1691505061190f565b600260008685604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561183d578181015183820152602001611825565b50505050905090810190601f16801561186a5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561189d578181015183820152602001611885565b50505050905090810190601f1680156118ca5780820380516001836020036101000a031916815260200191505b50945050505050604051602081830303815290604052805190602001208152602001908152602001600020546000141515611909576117bb858461200b565b60009150505b9392505050565b60008060009054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b15801561196557600080fd5b505afa158015611979573d6000803e3d6000fd5b505050506040513d602081101561198f57600080fd5b50516001600160a01b031633146119e45760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b611a8b87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b01819004810282018101909252898152925089915088908190840183828082843760009201919091525050604080516020601f8a0181900481028201810190925288815292508891508790819084018382808284376000920191909152506116a292505050565b1515611a9957506000611c8c565b600060026000898989896040516020018080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f8201169050808301925050509650505050505050604051602081830303815290604052805190602001208152602001908152602001600020546000141515611bb057611ba988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a915089908190840183828082843760009201919091525061200b92505050565b9050611c26565b611c2388888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a01819004810282018101909252888152925088915087908190840183828082843760009201919091525061200b92505050565b90505b6001805482908110611c3457fe5b906000526020600020906004020160030160029054906101000a900460ff168015611c8857506001805482908110611c6857fe5b906000526020600020906004020160030160019054906101000a900460ff165b9150505b9695505050505050565b60008060009054906101000a90046001600160a01b03166001600160a01b0316630e32cf906040518163ffffffff1660e01b815260040160206040518083038186803b158015611ce557600080fd5b505afa158015611cf9573d6000803e3d6000fd5b505050506040513d6020811015611d0f57600080fd5b50516001600160a01b03163314611d645760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b611e0b87878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b01819004810282018101909252898152925089915088908190840183828082843760009201919091525050604080516020601f8a0181900481028201810190925288815292508891508790819084018382808284376000920191909152506116a292505050565b1515611e1957506000611c8c565b600060026000898989896040516020018080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f8201169050808301925050509650505050505050604051602081830303815290604052805190602001208152602001908152602001600020546000141515611f3057611f2988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a915089908190840183828082843760009201919091525061200b92505050565b9050611fa6565b611fa388888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a01819004810282018101909252888152925088915087908190840183828082843760009201919091525061200b92505050565b90505b6001805482908110611fb457fe5b906000526020600020906004020160030160029054906101000a900460ff168015611c8857506001805482908110611fe857fe5b600091825260209091206004909102016003015460ff1698975050505050505050565b60006001600260008585604051602001808060200180602001838103835285818151815260200191508051906020019080838360005b83811015612059578181015183820152602001612041565b50505050905090810190601f1680156120865780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156120b95781810151838201526020016120a1565b50505050905090810190601f1680156120e65780820380516001836020036101000a031916815260200191505b509450505050506040516020818303038152906040528051906020012081526020019081526020016000205403905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061215c57805160ff1916838001178555612189565b82800160010185558215612189579182015b8281111561218957825182559160200191906001019061216e565b50612195929150612199565b5090565b61118691905b80821115612195576000815560010161219f56fea165627a7a723058209059a9af47943da0750b529cb5cf17b9f0745cfb3bea00dad68345c815bbec800029"

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

var RoleCreatedTopicHash = "0xefa5bc1bedbee25b04b00855c15a0c180ecb4a2440d4d08296e49561655e2b1c"

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

var RoleRevokedTopicHash = "0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6"

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

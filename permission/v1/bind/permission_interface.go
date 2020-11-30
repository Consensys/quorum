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

// PermInterfaceABI is the input ABI used to generate the binding from.
const PermInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getPermissionsImpl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"approveAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setPolicy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_roleId\",\"type\":\"string\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"approveBlacklistedAccountRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"addAdminNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_roleId\",\"type\":\"string\"}],\"name\":\"assignAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pOrgId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"addAdminAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_permImplementation\",\"type\":\"address\"}],\"name\":\"setPermImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_access\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"bool\"},{\"name\":\"_admin\",\"type\":\"bool\"}],\"name\":\"addNewRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"approveBlacklistedNodeRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"approveOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"approveOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateAccountStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"startBlacklistedNodeRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isOrgAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_breadth\",\"type\":\"uint256\"},{\"name\":\"_depth\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"startBlacklistedAccountRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permImplUpgradeable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

var PermInterfaceParsedABI, _ = abi.JSON(strings.NewReader(PermInterfaceABI))

// PermInterfaceBin is the compiled bytecode used for deploying new contracts.
var PermInterfaceBin = "0x608060405234801561001057600080fd5b506040516020806125aa8339810180604052602081101561003057600080fd5b5051600280546001600160a01b0319166001600160a01b0390921691909117905561254a806100606000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80635adbfa7a116101045780639bd38101116100a2578063a97a440611610071578063a97a440614610f50578063bb3b6e801461100e578063d1aa0c201461107c578063f346a3a7146110a2576101cf565b80639bd3810114610d7a578063a5843f0814610df8578063a634301214610e1b578063a97914bf14610ed9576101cf565b80637e461258116100de5780637e46125814610ab157806384b7a84a14610b785780638cb58ef314610bf55780638f362a3e14610cb3576101cf565b80635adbfa7a146109075780635be9672c146109c55780636b568d7614610a33576101cf565b806343de646c116101715780634cff819e1161014b5780634cff819e146106df5780634fe57e7a146107ed578063511bbd9f1461081357806351f604c314610839576101cf565b806343de646c146105f057806344478e79146106bb5780634cbfa82e146106d7576101cf565b80631b610220116101ad5780631b6102201461032f5780632f7f0a121461043d5780633e239b231461050b5780633f25c28814610582576101cf565b806303ed6933146101d45780630cc50146146101f857806316724c44146102b8575b600080fd5b6101dc61120e565b604080516001600160a01b039092168252519081900360200190f35b6102b66004803603606081101561020e57600080fd5b810190602081018135600160201b81111561022857600080fd5b82018360208201111561023a57600080fd5b803590602001918460018302840111600160201b8311171561025b57600080fd5b919390929091602081019035600160201b81111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111600160201b831117156102ab57600080fd5b91935091503561121d565b005b6102b6600480360360408110156102ce57600080fd5b810190602081018135600160201b8111156102e857600080fd5b8201836020820111156102fa57600080fd5b803590602001918460018302840111600160201b8311171561031b57600080fd5b9193509150356001600160a01b03166112fa565b6102b66004803603606081101561034557600080fd5b810190602081018135600160201b81111561035f57600080fd5b82018360208201111561037157600080fd5b803590602001918460018302840111600160201b8311171561039257600080fd5b919390929091602081019035600160201b8111156103af57600080fd5b8201836020820111156103c157600080fd5b803590602001918460018302840111600160201b831117156103e257600080fd5b919390929091602081019035600160201b8111156103ff57600080fd5b82018360208201111561041157600080fd5b803590602001918460018302840111600160201b8311171561043257600080fd5b5090925090506113a9565b6102b66004803603606081101561045357600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561047d57600080fd5b82018360208201111561048f57600080fd5b803590602001918460018302840111600160201b831117156104b057600080fd5b919390929091602081019035600160201b8111156104cd57600080fd5b8201836020820111156104df57600080fd5b803590602001918460018302840111600160201b8311171561050057600080fd5b5090925090506114a0565b6102b66004803603604081101561052157600080fd5b810190602081018135600160201b81111561053b57600080fd5b82018360208201111561054d57600080fd5b803590602001918460018302840111600160201b8311171561056e57600080fd5b9193509150356001600160a01b0316611562565b6102b66004803603602081101561059857600080fd5b810190602081018135600160201b8111156105b257600080fd5b8201836020820111156105c457600080fd5b803590602001918460018302840111600160201b831117156105e557600080fd5b5090925090506115f4565b6102b66004803603606081101561060657600080fd5b810190602081018135600160201b81111561062057600080fd5b82018360208201111561063257600080fd5b803590602001918460018302840111600160201b8311171561065357600080fd5b919390926001600160a01b0383351692604081019060200135600160201b81111561067d57600080fd5b82018360208201111561068f57600080fd5b803590602001918460018302840111600160201b831117156106b057600080fd5b50909250905061168a565b6106c3611749565b604080519115158252519081900360200190f35b6106c36117cb565b6102b6600480360360608110156106f557600080fd5b810190602081018135600160201b81111561070f57600080fd5b82018360208201111561072157600080fd5b803590602001918460018302840111600160201b8311171561074257600080fd5b919390929091602081019035600160201b81111561075f57600080fd5b82018360208201111561077157600080fd5b803590602001918460018302840111600160201b8311171561079257600080fd5b919390929091602081019035600160201b8111156107af57600080fd5b8201836020820111156107c157600080fd5b803590602001918460018302840111600160201b831117156107e257600080fd5b50909250905061182e565b6102b66004803603602081101561080357600080fd5b50356001600160a01b031661190f565b6102b66004803603602081101561082957600080fd5b50356001600160a01b0316611978565b6102b6600480360360a081101561084f57600080fd5b810190602081018135600160201b81111561086957600080fd5b82018360208201111561087b57600080fd5b803590602001918460018302840111600160201b8311171561089c57600080fd5b919390929091602081019035600160201b8111156108b957600080fd5b8201836020820111156108cb57600080fd5b803590602001918460018302840111600160201b831117156108ec57600080fd5b919350915080359060208101351515906040013515156119fc565b6102b66004803603604081101561091d57600080fd5b810190602081018135600160201b81111561093757600080fd5b82018360208201111561094957600080fd5b803590602001918460018302840111600160201b8311171561096a57600080fd5b919390929091602081019035600160201b81111561098757600080fd5b82018360208201111561099957600080fd5b803590602001918460018302840111600160201b831117156109ba57600080fd5b509092509050611af1565b6102b6600480360360408110156109db57600080fd5b810190602081018135600160201b8111156109f557600080fd5b820183602082011115610a0757600080fd5b803590602001918460018302840111600160201b83111715610a2857600080fd5b919350915035611bc3565b6106c360048036036040811015610a4957600080fd5b6001600160a01b038235169190810190604081016020820135600160201b811115610a7357600080fd5b820183602082011115610a8557600080fd5b803590602001918460018302840111600160201b83111715610aa657600080fd5b509092509050611c50565b6102b660048036036060811015610ac757600080fd5b810190602081018135600160201b811115610ae157600080fd5b820183602082011115610af357600080fd5b803590602001918460018302840111600160201b83111715610b1457600080fd5b919390929091602081019035600160201b811115610b3157600080fd5b820183602082011115610b4357600080fd5b803590602001918460018302840111600160201b83111715610b6457600080fd5b9193509150356001600160a01b0316611d08565b6102b660048036036060811015610b8e57600080fd5b810190602081018135600160201b811115610ba857600080fd5b820183602082011115610bba57600080fd5b803590602001918460018302840111600160201b83111715610bdb57600080fd5b91935091506001600160a01b038135169060200135611dc7565b6102b660048036036040811015610c0b57600080fd5b810190602081018135600160201b811115610c2557600080fd5b820183602082011115610c3757600080fd5b803590602001918460018302840111600160201b83111715610c5857600080fd5b919390929091602081019035600160201b811115610c7557600080fd5b820183602082011115610c8757600080fd5b803590602001918460018302840111600160201b83111715610ca857600080fd5b509092509050611e61565b6102b660048036036060811015610cc957600080fd5b810190602081018135600160201b811115610ce357600080fd5b820183602082011115610cf557600080fd5b803590602001918460018302840111600160201b83111715610d1657600080fd5b919390929091602081019035600160201b811115610d3357600080fd5b820183602082011115610d4557600080fd5b803590602001918460018302840111600160201b83111715610d6657600080fd5b9193509150356001600160a01b0316611f15565b6106c360048036036040811015610d9057600080fd5b6001600160a01b038235169190810190604081016020820135600160201b811115610dba57600080fd5b820183602082011115610dcc57600080fd5b803590602001918460018302840111600160201b83111715610ded57600080fd5b509092509050611fd4565b6102b660048036036040811015610e0e57600080fd5b5080359060200135612058565b6102b660048036036040811015610e3157600080fd5b810190602081018135600160201b811115610e4b57600080fd5b820183602082011115610e5d57600080fd5b803590602001918460018302840111600160201b83111715610e7e57600080fd5b919390929091602081019035600160201b811115610e9b57600080fd5b820183602082011115610ead57600080fd5b803590602001918460018302840111600160201b83111715610ece57600080fd5b5090925090506120aa565b6102b660048036036040811015610eef57600080fd5b810190602081018135600160201b811115610f0957600080fd5b820183602082011115610f1b57600080fd5b803590602001918460018302840111600160201b83111715610f3c57600080fd5b9193509150356001600160a01b031661215e565b6102b660048036036040811015610f6657600080fd5b810190602081018135600160201b811115610f8057600080fd5b820183602082011115610f9257600080fd5b803590602001918460018302840111600160201b83111715610fb357600080fd5b919390929091602081019035600160201b811115610fd057600080fd5b820183602082011115610fe257600080fd5b803590602001918460018302840111600160201b8311171561100357600080fd5b5090925090506121f0565b6102b66004803603604081101561102457600080fd5b810190602081018135600160201b81111561103e57600080fd5b82018360208201111561105057600080fd5b803590602001918460018302840111600160201b8311171561107157600080fd5b9193509150356122a4565b6106c36004803603602081101561109257600080fd5b50356001600160a01b0316612331565b611110600480360360208110156110b857600080fd5b810190602081018135600160201b8111156110d257600080fd5b8201836020820111156110e457600080fd5b803590602001918460018302840111600160201b8311171561110557600080fd5b5090925090506123b4565b604051808060200180602001856001600160a01b03166001600160a01b03168152602001848152602001838103835287818151815260200191508051906020019080838360005b8381101561116f578181015183820152602001611157565b50505050905090810190601f16801561119c5780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b838110156111cf5781810151838201526020016111b7565b50505050905090810190601f1680156111fc5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6000546001600160a01b031690565b600054604051600160e01b63dbfad711028152604481018390523360648201819052608060048301908152608483018890526001600160a01b039093169263dbfad711928992899289928992899290918190602481019060a401898980828437600083820152601f01601f191690910184810383528781526020019050878780828437600081840152601f19601f82011690508083019250505098505050505050505050600060405180830381600087803b1580156112db57600080fd5b505af11580156112ef573d6000803e3d6000fd5b505050505050505050565b600054604051600160e01b63888430410281526001600160a01b03838116602483015233604483018190526060600484019081526064840187905291909316926388843041928792879287929091908190608401868680828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561138c57600080fd5b505af11580156113a0573d6000803e3d6000fd5b50505050505050565b600054604051600160e51b62db0811028152606060048201908152606482018890526001600160a01b0390921691631b610220918991899189918991899189918190602481019060448101906084018a8a80828437600083820152601f01601f191690910185810384528881526020019050888880828437600083820152601f01601f191690910185810383528681526020019050868680828437600081840152601f19601f8201169050808301925050509950505050505050505050600060405180830381600087803b15801561148057600080fd5b505af1158015611494573d6000803e3d6000fd5b50505050505050505050565b600054604051600160e01b638baa81910281526001600160a01b03878116600483019081523360648401819052608060248501908152608485018990529290941693638baa8191938a938a938a938a938a9391929190604481019060a401888880828437600083820152601f01601f191690910184810383528681526020019050868680828437600081840152601f19601f82011690508083019250505098505050505050505050600060405180830381600087803b1580156112db57600080fd5b600054604051600160e01b634b20f45f0281526001600160a01b0383811660248301523360448301819052606060048401908152606484018790529190931692634b20f45f928792879287929091908190608401868680828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561138c57600080fd5b600054604051600160e31b6307e4b851028152602060048201908152602482018490526001600160a01b0390921691633f25c28891859185918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561166e57600080fd5b505af1158015611682573d6000803e3d6000fd5b505050505050565b600054604051600160e01b63404bf3eb0281526001600160a01b038581166024830152336064830181905260806004840190815260848401899052919093169263404bf3eb9289928992899289928992918190604481019060a401898980828437600083820152601f01601f191690910184810383528681526020019050868680828437600081840152601f19601f82011690508083019250505098505050505050505050600060405180830381600087803b1580156112db57600080fd5b60008060009054906101000a90046001600160a01b03166001600160a01b03166344478e796040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561179a57600080fd5b505af11580156117ae573d6000803e3d6000fd5b505050506040513d60208110156117c457600080fd5b5051905090565b60008060009054906101000a90046001600160a01b03166001600160a01b0316634cbfa82e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561181a57600080fd5b505afa1580156117ae573d6000803e3d6000fd5b600054604051600160e51b63053269430281523360648201819052608060048301908152608483018990526001600160a01b039093169263a64d2860928a928a928a928a928a928a9281906024810190604481019060a4018b8b80828437600083820152601f01601f191690910185810384528981526020019050898980828437600083820152601f01601f191690910185810383528781526020019050878780828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b15801561148057600080fd5b6000805460408051600160e11b6327f2bf3d0281526001600160a01b03858116600483015291519190921692634fe57e7a926024808201939182900301818387803b15801561195d57600080fd5b505af1158015611971573d6000803e3d6000fd5b5050505050565b6002546001600160a01b031633146119da5760408051600160e51b62461bcd02815260206004820152600e60248201527f696e76616c69642063616c6c6572000000000000000000000000000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600054604051600160e11b630d82613b02815260448101859052831515606482015282151560848201523360a4820181905260c06004830190815260c483018a90526001600160a01b0390931692631b04c276928b928b928b928b928b928b928b9291908190602481019060e4018b8b80828437600083820152601f01601f191690910184810383528981526020019050898980828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b158015611ad057600080fd5b505af1158015611ae4573d6000803e3d6000fd5b5050505050505050505050565b600054604051600160e01b63655a8ef50281523360448201819052606060048301908152606483018790526001600160a01b039093169263655a8ef5928892889288928892919081906024810190608401888880828437600083820152601f01601f191690910184810383528681526020019050868680828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b158015611ba557600080fd5b505af1158015611bb9573d6000803e3d6000fd5b5050505050505050565b600054604051600160e21b632d551959028152602481018390523360448201819052606060048301908152606483018690526001600160a01b039093169263b5546564928792879287928190608401868680828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561138c57600080fd5b6000805460408051600160e11b6335ab46bb0281526001600160a01b03878116600483019081526024830193845260448301879052931692636b568d76928892889288929091606401848480828437600083820152604051601f909101601f1916909201965060209550909350505081840390508186803b158015611cd457600080fd5b505afa158015611ce8573d6000803e3d6000fd5b505050506040513d6020811015611cfe57600080fd5b5051949350505050565b600054604051600160e11b631de03ef50281526001600160a01b0383811660448301523360648301819052608060048401908152608484018990529190931692633bc07dea9289928992899289928992918190602481019060a401898980828437600083820152601f01601f191690910184810383528781526020019050878780828437600081840152601f19601f82011690508083019250505098505050505050505050600060405180830381600087803b1580156112db57600080fd5b600054604051600160e11b6302740f8f0281526001600160a01b0384811660248301526044820184905233606483018190526080600484019081526084840188905291909316926304e81f1e92889288928892889290819060a401878780828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b158015611ba557600080fd5b600054604051600160e01b63c3dc8e090281523360448201819052606060048301908152606483018790526001600160a01b039093169263c3dc8e09928892889288928892919081906024810190608401888880828437600083820152601f01601f191690910184810383528681526020019050868680828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b158015611ba557600080fd5b600054604051600160e11b637c917c010281526001600160a01b038381166044830152336064830181905260806004840190815260848401899052919093169263f922f8029289928992899289928992918190602481019060a401898980828437600083820152601f01601f191690910184810383528781526020019050878780828437600081840152601f19601f82011690508083019250505098505050505050505050600060405180830381600087803b1580156112db57600080fd5b6000805460408051600160e01b639bd381010281526001600160a01b03878116600483019081526024830193845260448301879052931692639bd38101928892889288929091606401848480828437600083820152604051601f909101601f1916909201965060209550909350505081840390508186803b158015611cd457600080fd5b6000805460408051600160e31b6314b087e1028152600481018690526024810185905290516001600160a01b039092169263a5843f089260448084019382900301818387803b15801561166e57600080fd5b600054604051600160e11b632e52d6df0281523360448201819052606060048301908152606483018790526001600160a01b0390931692635ca5adbe928892889288928892919081906024810190608401888880828437600083820152601f01601f191690910184810383528681526020019050868680828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b158015611ba557600080fd5b600054604051600160e11b630e124c890281526001600160a01b0383811660248301523360448301819052606060048401908152606484018790529190931692631c249912928792879287929091908190608401868680828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561138c57600080fd5b600054604051600160e01b6359a260a30281523360448201819052606060048301908152606483018790526001600160a01b03909316926359a260a3928892889288928892919081906024810190608401888880828437600083820152601f01601f191690910184810383528681526020019050868680828437600081840152601f19601f820116905080830192505050975050505050505050600060405180830381600087803b158015611ba557600080fd5b600054604051600160e01b633cf5f33b028152602481018390523360448201819052606060048301908152606483018690526001600160a01b0390931692633cf5f33b928792879287928190608401868680828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561138c57600080fd5b6000805460408051600160e51b63068d50610281526001600160a01b0385811660048301529151919092169163d1aa0c20916024808301926020929190829003018186803b15801561238257600080fd5b505afa158015612396573d6000803e3d6000fd5b505050506040513d60208110156123ac57600080fd5b505192915050565b60008054604051600160e01b63f346a3a7028152602060048201908152602482018590526060938493909283926001600160a01b039092169163f346a3a791899189918190604401848480828437600081840152601f19601f820116905080830192505050935050505060006040518083038186803b15801561243657600080fd5b505afa15801561244a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052608081101561247357600080fd5b810190808051600160201b81111561248a57600080fd5b8201602081018481111561249d57600080fd5b8151600160201b8111828201871017156124b657600080fd5b50509291906020018051600160201b8111156124d157600080fd5b820160208101848111156124e457600080fd5b8151600160201b8111828201871017156124fd57600080fd5b50506020820151604090920151949b909a509098509296509194505050505056fea165627a7a72305820c59bf7b1eb3a15d1406b140bc566b70353e3ef021637abb4ecb03c63261f92b10029"

// DeployPermInterface deploys a new Ethereum contract, binding an instance of PermInterface to it.
func DeployPermInterface(auth *bind.TransactOpts, backend bind.ContractBackend, _permImplUpgradeable common.Address) (common.Address, *types.Transaction, *PermInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(PermInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PermInterfaceBin), backend, _permImplUpgradeable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermInterface{PermInterfaceCaller: PermInterfaceCaller{contract: contract}, PermInterfaceTransactor: PermInterfaceTransactor{contract: contract}, PermInterfaceFilterer: PermInterfaceFilterer{contract: contract}}, nil
}

// PermInterface is an auto generated Go binding around an Ethereum contract.
type PermInterface struct {
	PermInterfaceCaller     // Read-only binding to the contract
	PermInterfaceTransactor // Write-only binding to the contract
	PermInterfaceFilterer   // Log filterer for contract events
}

// PermInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermInterfaceSession struct {
	Contract     *PermInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermInterfaceCallerSession struct {
	Contract *PermInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PermInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermInterfaceTransactorSession struct {
	Contract     *PermInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PermInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermInterfaceRaw struct {
	Contract *PermInterface // Generic contract binding to access the raw methods on
}

// PermInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermInterfaceCallerRaw struct {
	Contract *PermInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PermInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermInterfaceTransactorRaw struct {
	Contract *PermInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermInterface creates a new instance of PermInterface, bound to a specific deployed contract.
func NewPermInterface(address common.Address, backend bind.ContractBackend) (*PermInterface, error) {
	contract, err := bindPermInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermInterface{PermInterfaceCaller: PermInterfaceCaller{contract: contract}, PermInterfaceTransactor: PermInterfaceTransactor{contract: contract}, PermInterfaceFilterer: PermInterfaceFilterer{contract: contract}}, nil
}

// NewPermInterfaceCaller creates a new read-only instance of PermInterface, bound to a specific deployed contract.
func NewPermInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PermInterfaceCaller, error) {
	contract, err := bindPermInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermInterfaceCaller{contract: contract}, nil
}

// NewPermInterfaceTransactor creates a new write-only instance of PermInterface, bound to a specific deployed contract.
func NewPermInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PermInterfaceTransactor, error) {
	contract, err := bindPermInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermInterfaceTransactor{contract: contract}, nil
}

// NewPermInterfaceFilterer creates a new log filterer instance of PermInterface, bound to a specific deployed contract.
func NewPermInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PermInterfaceFilterer, error) {
	contract, err := bindPermInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermInterfaceFilterer{contract: contract}, nil
}

// bindPermInterface binds a generic wrapper to an already deployed contract.
func bindPermInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermInterface *PermInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermInterface.Contract.PermInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermInterface *PermInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermInterface.Contract.PermInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermInterface *PermInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermInterface.Contract.PermInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermInterface *PermInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermInterface *PermInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermInterface *PermInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermInterface.Contract.contract.Transact(opts, method, params...)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermInterface *PermInterfaceCaller) GetNetworkBootStatus(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "getNetworkBootStatus")
	return *ret0, err
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermInterface *PermInterfaceSession) GetNetworkBootStatus() (bool, error) {
	return _PermInterface.Contract.GetNetworkBootStatus(&_PermInterface.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) GetNetworkBootStatus() (bool, error) {
	return _PermInterface.Contract.GetNetworkBootStatus(&_PermInterface.CallOpts)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) constant returns(string, string, address, uint256)
func (_PermInterface *PermInterfaceCaller) GetPendingOp(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
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
	err := _PermInterface.contract.Call(opts, out, "getPendingOp", _orgId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) constant returns(string, string, address, uint256)
func (_PermInterface *PermInterfaceSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermInterface.Contract.GetPendingOp(&_PermInterface.CallOpts, _orgId)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) constant returns(string, string, address, uint256)
func (_PermInterface *PermInterfaceCallerSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermInterface.Contract.GetPendingOp(&_PermInterface.CallOpts, _orgId)
}

// GetPermissionsImpl is a free data retrieval call binding the contract method 0x03ed6933.
//
// Solidity: function getPermissionsImpl() constant returns(address)
func (_PermInterface *PermInterfaceCaller) GetPermissionsImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "getPermissionsImpl")
	return *ret0, err
}

// GetPermissionsImpl is a free data retrieval call binding the contract method 0x03ed6933.
//
// Solidity: function getPermissionsImpl() constant returns(address)
func (_PermInterface *PermInterfaceSession) GetPermissionsImpl() (common.Address, error) {
	return _PermInterface.Contract.GetPermissionsImpl(&_PermInterface.CallOpts)
}

// GetPermissionsImpl is a free data retrieval call binding the contract method 0x03ed6933.
//
// Solidity: function getPermissionsImpl() constant returns(address)
func (_PermInterface *PermInterfaceCallerSession) GetPermissionsImpl() (common.Address, error) {
	return _PermInterface.Contract.GetPermissionsImpl(&_PermInterface.CallOpts)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) constant returns(bool)
func (_PermInterface *PermInterfaceCaller) IsNetworkAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "isNetworkAdmin", _account)
	return *ret0, err
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) constant returns(bool)
func (_PermInterface *PermInterfaceSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermInterface.Contract.IsNetworkAdmin(&_PermInterface.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermInterface.Contract.IsNetworkAdmin(&_PermInterface.CallOpts, _account)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) constant returns(bool)
func (_PermInterface *PermInterfaceCaller) IsOrgAdmin(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "isOrgAdmin", _account, _orgId)
	return *ret0, err
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) constant returns(bool)
func (_PermInterface *PermInterfaceSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.IsOrgAdmin(&_PermInterface.CallOpts, _account, _orgId)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.IsOrgAdmin(&_PermInterface.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) constant returns(bool)
func (_PermInterface *PermInterfaceCaller) ValidateAccount(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "validateAccount", _account, _orgId)
	return *ret0, err
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) constant returns(bool)
func (_PermInterface *PermInterfaceSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.ValidateAccount(&_PermInterface.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.ValidateAccount(&_PermInterface.CallOpts, _account, _orgId)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _acct) returns()
func (_PermInterface *PermInterfaceTransactor) AddAdminAccount(opts *bind.TransactOpts, _acct common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addAdminAccount", _acct)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _acct) returns()
func (_PermInterface *PermInterfaceSession) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminAccount(&_PermInterface.TransactOpts, _acct)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _acct) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddAdminAccount(_acct common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminAccount(&_PermInterface.TransactOpts, _acct)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x3f25c288.
//
// Solidity: function addAdminNode(string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactor) AddAdminNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addAdminNode", _enodeId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x3f25c288.
//
// Solidity: function addAdminNode(string _enodeId) returns()
func (_PermInterface *PermInterfaceSession) AddAdminNode(_enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminNode(&_PermInterface.TransactOpts, _enodeId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x3f25c288.
//
// Solidity: function addAdminNode(string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddAdminNode(_enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminNode(&_PermInterface.TransactOpts, _enodeId)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x51f604c3.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin) returns()
func (_PermInterface *PermInterfaceTransactor) AddNewRole(opts *bind.TransactOpts, _roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addNewRole", _roleId, _orgId, _access, _voter, _admin)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x51f604c3.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin) returns()
func (_PermInterface *PermInterfaceSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNewRole(&_PermInterface.TransactOpts, _roleId, _orgId, _access, _voter, _admin)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x51f604c3.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNewRole(&_PermInterface.TransactOpts, _roleId, _orgId, _access, _voter, _admin)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactor) AddNode(opts *bind.TransactOpts, _orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addNode", _orgId, _enodeId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceSession) AddNode(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNode(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddNode(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNode(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(string _orgId, string _enodeId, address _account) returns()
func (_PermInterface *PermInterfaceTransactor) AddOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addOrg", _orgId, _enodeId, _account)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(string _orgId, string _enodeId, address _account) returns()
func (_PermInterface *PermInterfaceSession) AddOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(string _orgId, string _enodeId, address _account) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x4cff819e.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactor) AddSubOrg(opts *bind.TransactOpts, _pOrgId string, _orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addSubOrg", _pOrgId, _orgId, _enodeId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x4cff819e.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceSession) AddSubOrg(_pOrgId string, _orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddSubOrg(&_PermInterface.TransactOpts, _pOrgId, _orgId, _enodeId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x4cff819e.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddSubOrg(_pOrgId string, _orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddSubOrg(&_PermInterface.TransactOpts, _pOrgId, _orgId, _enodeId)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x16724c44.
//
// Solidity: function approveAdminRole(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveAdminRole", _orgId, _account)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x16724c44.
//
// Solidity: function approveAdminRole(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceSession) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveAdminRole(&_PermInterface.TransactOpts, _orgId, _account)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x16724c44.
//
// Solidity: function approveAdminRole(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveAdminRole(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveAdminRole(&_PermInterface.TransactOpts, _orgId, _account)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x3e239b23.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveBlacklistedAccountRecovery(opts *bind.TransactOpts, _orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveBlacklistedAccountRecovery", _orgId, _account)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x3e239b23.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceSession) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveBlacklistedAccountRecovery(&_PermInterface.TransactOpts, _orgId, _account)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x3e239b23.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveBlacklistedAccountRecovery(&_PermInterface.TransactOpts, _orgId, _account)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0x5adbfa7a.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveBlacklistedNodeRecovery(opts *bind.TransactOpts, _orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveBlacklistedNodeRecovery", _orgId, _enodeId)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0x5adbfa7a.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceSession) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveBlacklistedNodeRecovery(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0x5adbfa7a.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveBlacklistedNodeRecovery(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, address _account) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveOrg", _orgId, _enodeId, _account)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, address _account) returns()
func (_PermInterface *PermInterfaceSession) ApproveOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, address _account) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0x5be9672c.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveOrgStatus(opts *bind.TransactOpts, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveOrgStatus", _orgId, _action)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0x5be9672c.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action) returns()
func (_PermInterface *PermInterfaceSession) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrgStatus(&_PermInterface.TransactOpts, _orgId, _action)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0x5be9672c.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrgStatus(&_PermInterface.TransactOpts, _orgId, _action)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId) returns()
func (_PermInterface *PermInterfaceTransactor) AssignAccountRole(opts *bind.TransactOpts, _account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "assignAccountRole", _account, _orgId, _roleId)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId) returns()
func (_PermInterface *PermInterfaceSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignAccountRole(&_PermInterface.TransactOpts, _account, _orgId, _roleId)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId) returns()
func (_PermInterface *PermInterfaceTransactorSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignAccountRole(&_PermInterface.TransactOpts, _account, _orgId, _roleId)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x43de646c.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId) returns()
func (_PermInterface *PermInterfaceTransactor) AssignAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "assignAdminRole", _orgId, _account, _roleId)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x43de646c.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId) returns()
func (_PermInterface *PermInterfaceSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignAdminRole(&_PermInterface.TransactOpts, _orgId, _account, _roleId)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x43de646c.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId) returns()
func (_PermInterface *PermInterfaceTransactorSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignAdminRole(&_PermInterface.TransactOpts, _orgId, _account, _roleId)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermInterface *PermInterfaceTransactor) Init(opts *bind.TransactOpts, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "init", _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermInterface *PermInterfaceSession) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.Init(&_PermInterface.TransactOpts, _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermInterface *PermInterfaceTransactorSession) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.Init(&_PermInterface.TransactOpts, _breadth, _depth)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_PermInterface *PermInterfaceTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "removeRole", _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_PermInterface *PermInterfaceSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _PermInterface.Contract.RemoveRole(&_PermInterface.TransactOpts, _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(string _roleId, string _orgId) returns()
func (_PermInterface *PermInterfaceTransactorSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _PermInterface.Contract.RemoveRole(&_PermInterface.TransactOpts, _roleId, _orgId)
}

// SetPermImplementation is a paid mutator transaction binding the contract method 0x511bbd9f.
//
// Solidity: function setPermImplementation(address _permImplementation) returns()
func (_PermInterface *PermInterfaceTransactor) SetPermImplementation(opts *bind.TransactOpts, _permImplementation common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "setPermImplementation", _permImplementation)
}

// SetPermImplementation is a paid mutator transaction binding the contract method 0x511bbd9f.
//
// Solidity: function setPermImplementation(address _permImplementation) returns()
func (_PermInterface *PermInterfaceSession) SetPermImplementation(_permImplementation common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPermImplementation(&_PermInterface.TransactOpts, _permImplementation)
}

// SetPermImplementation is a paid mutator transaction binding the contract method 0x511bbd9f.
//
// Solidity: function setPermImplementation(address _permImplementation) returns()
func (_PermInterface *PermInterfaceTransactorSession) SetPermImplementation(_permImplementation common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPermImplementation(&_PermInterface.TransactOpts, _permImplementation)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermInterface *PermInterfaceTransactor) SetPolicy(opts *bind.TransactOpts, _nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "setPolicy", _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermInterface *PermInterfaceSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPolicy(&_PermInterface.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermInterface *PermInterfaceTransactorSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPolicy(&_PermInterface.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0xa97914bf.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceTransactor) StartBlacklistedAccountRecovery(opts *bind.TransactOpts, _orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "startBlacklistedAccountRecovery", _orgId, _account)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0xa97914bf.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceSession) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.StartBlacklistedAccountRecovery(&_PermInterface.TransactOpts, _orgId, _account)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0xa97914bf.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account) returns()
func (_PermInterface *PermInterfaceTransactorSession) StartBlacklistedAccountRecovery(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.StartBlacklistedAccountRecovery(&_PermInterface.TransactOpts, _orgId, _account)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0x8cb58ef3.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactor) StartBlacklistedNodeRecovery(opts *bind.TransactOpts, _orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "startBlacklistedNodeRecovery", _orgId, _enodeId)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0x8cb58ef3.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceSession) StartBlacklistedNodeRecovery(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.StartBlacklistedNodeRecovery(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0x8cb58ef3.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId) returns()
func (_PermInterface *PermInterfaceTransactorSession) StartBlacklistedNodeRecovery(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.StartBlacklistedNodeRecovery(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactor) UpdateAccountStatus(opts *bind.TransactOpts, _orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "updateAccountStatus", _orgId, _account, _action)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action) returns()
func (_PermInterface *PermInterfaceSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateAccountStatus(&_PermInterface.TransactOpts, _orgId, _account, _action)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactorSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateAccountStatus(&_PermInterface.TransactOpts, _orgId, _account, _action)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermInterface *PermInterfaceTransactor) UpdateNetworkBootStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "updateNetworkBootStatus")
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermInterface *PermInterfaceSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNetworkBootStatus(&_PermInterface.TransactOpts)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermInterface *PermInterfaceTransactorSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNetworkBootStatus(&_PermInterface.TransactOpts)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _orgId string, _enodeId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "updateNodeStatus", _orgId, _enodeId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, uint256 _action) returns()
func (_PermInterface *PermInterfaceSession) UpdateNodeStatus(_orgId string, _enodeId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNodeStatus(&_PermInterface.TransactOpts, _orgId, _enodeId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactorSession) UpdateNodeStatus(_orgId string, _enodeId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNodeStatus(&_PermInterface.TransactOpts, _orgId, _enodeId, _action)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0xbb3b6e80.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactor) UpdateOrgStatus(opts *bind.TransactOpts, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "updateOrgStatus", _orgId, _action)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0xbb3b6e80.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action) returns()
func (_PermInterface *PermInterfaceSession) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateOrgStatus(&_PermInterface.TransactOpts, _orgId, _action)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0xbb3b6e80.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action) returns()
func (_PermInterface *PermInterfaceTransactorSession) UpdateOrgStatus(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateOrgStatus(&_PermInterface.TransactOpts, _orgId, _action)
}

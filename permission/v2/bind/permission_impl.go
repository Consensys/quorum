// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PermImplMetaData contains all meta data concerning the PermImpl contract.
var PermImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_networkBootStatus\",\"type\":\"bool\"}],\"name\":\"PermissionsInitialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAdminAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"}],\"name\":\"addAdminNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contractKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addContractWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_access\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_admin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNewRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pOrgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveBlacklistedAccountRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveBlacklistedNodeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrgStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"}],\"name\":\"connectionAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccessLevelForUnconfiguredAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetworkBootStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOp\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolicyDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_breadth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depth\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_orgManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rolesManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractWhitelistManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isOrgAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roleId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"removeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"revokeContractWhitelistByAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contractKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"revokeContractWhitelistByKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accessLevel\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"setAccessLevelForUnconfiguredAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isIpValidationEnabled\",\"type\":\"bool\"}],\"name\":\"setIpValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_oAdminRole\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_networkBootStatus\",\"type\":\"bool\"}],\"name\":\"setMigrationPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"startBlacklistedAccountRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"startBlacklistedNodeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"transactionAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateAccountStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateNetworkBootStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_enodeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_port\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_raftport\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_action\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateOrgStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526003600a555f600b5f6101000a81548160ff0219169083151502179055503480156200002e575f80fd5b506198ef806200003d5f395ff3fe608060405234801561000f575f80fd5b5060043610610230575f3560e01c80638683c7fe1161012e578063b9b7fe6c116100b6578063e91b0e191161007a578063e91b0e191461063f578063ecad01d51461065b578063f346a3a714610677578063f5ad584a146106aa578063f75f0a06146106c657610230565b8063b9b7fe6c1461059a578063cc9ba6fa146105b6578063d1aa0c20146105d7578063d43815f814610607578063d621d9571461062357610230565b80639bd38101116100fd5780639bd38101146104fa5780639fc666b21461052a578063a042bf4014610546578063a5843f0814610562578063b55465641461057e57610230565b80638683c7fe1461047657806388843041146104925780638baa8191146104ae578063936421d5146104ca57610230565b8063404bf3eb116101bc5780634fe57e7a116101805780634fe57e7a146103d65780635ca5adbe146103f257806368a612731461040e57806368f808e51461042a5780636b568d761461044657610230565b8063404bf3eb1461033257806344478e791461034e57806345a59e5b1461036c5780634b20f45f1461039c5780634cbfa82e146103b857610230565b80631c249912116102035780631c249912146102a657806327bb2cad146102c25780632a768da2146102de57806335876476146102fa5780633cf5f33b1461031657610230565b806304e81f1e146102345780630dca12f6146102505780631b04c2761461026e5780631b6102201461028a575b5f80fd5b61024e600480360381019061024991906168b7565b6106e2565b005b61025861095f565b604051610265919061694a565b60405180910390f35b61028860048036038101906102839190616998565b6109f3565b005b6102a4600480360381019061029f9190616a62565b610caf565b005b6102c060048036038101906102bb9190616b12565b610e3e565b005b6102dc60048036038101906102d79190616b12565b6110ad565b005b6102f860048036038101906102f39190616b83565b61128a565b005b610314600480360381019061030f9190616be0565b611464565b005b610330600480360381019061032b9190616c7d565b611aa3565b005b61034c60048036038101906103479190616cee565b611d27565b005b610356612029565b6040516103639190616da0565b60405180910390f35b61038660048036038101906103819190616df0565b6121ee565b6040516103939190616da0565b60405180910390f35b6103b660048036038101906103b19190616b12565b6122b8565b005b6103c0612530565b6040516103cd9190616da0565b60405180910390f35b6103f060048036038101906103eb9190616e81565b612545565b005b61040c60048036038101906104079190616eac565b6127bb565b005b61042860048036038101906104239190616f3d565b612b5b565b005b610444600480360381019061043f919061705a565b612edd565b005b610460600480360381019061045b91906171d0565b6130b4565b60405161046d9190616da0565b60405180910390f35b610490600480360381019061048b919061722a565b613157565b005b6104ac60048036038101906104a79190616b12565b613344565b005b6104c860048036038101906104c391906172cd565b6137a7565b005b6104e460048036038101906104df91906173be565b613b1a565b6040516104f19190616da0565b60405180910390f35b610514600480360381019061050f91906171d0565b613fa6565b6040516105219190616da0565b60405180910390f35b610544600480360381019061053f9190617468565b6141a0565b005b610560600480360381019061055b9190617493565b61437b565b005b61057c60048036038101906105779190617570565b6145fb565b005b61059860048036038101906105939190616c7d565b6148f9565b005b6105b460048036038101906105af91906175ae565b614c7b565b005b6105be614eb9565b6040516105ce9493929190617719565b60405180910390f35b6105f160048036038101906105ec9190616e81565b615080565b6040516105fe9190616da0565b60405180910390f35b610621600480360381019061061c9190617771565b615171565b005b61063d60048036038101906106389190617493565b615348565b005b610659600480360381019061065491906177af565b6155bf565b005b61067560048036038101906106709190617493565b6159eb565b005b610691600480360381019061068c91906178a0565b615c17565b6040516106a194939291906178fa565b60405180910390f35b6106c460048036038101906106bf919061794b565b615ccb565b005b6106e060048036038101906106db91906177af565b615e08565b005b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561074c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107709190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d490617abe565b60405180910390fd5b8085858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050600115156108308383613fa6565b151514610872576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086990617b4c565b60405180910390fd5b60018414806108815750600284145b8061088c5750600384145b6108cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c290617bda565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166384b7a84a888888886040518563ffffffff1660e01b81526004016109299493929190617c24565b5f604051808303815f87803b158015610940575f80fd5b505af1158015610952573d5f803e3d5ffd5b5050505050505050505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630dca12f66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109ca573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109ee9190617c76565b905090565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a5d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a819190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610aee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae590617abe565b60405180910390fd5b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050610b3b8161628a565b610b7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7190617ceb565b60405180910390fd5b8187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505060011515610bcd8383613fa6565b151514610c0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0690617b4c565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637b7135798c8c8c8c8c8c8c6040518863ffffffff1660e01b8152600401610c759796959493929190617d09565b5f604051808303815f87803b158015610c8c575f80fd5b505af1158015610c9e573d5f803e3d5ffd5b505050505050505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d19573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d3d9190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610daa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da190617abe565b60405180910390fd5b5f801515600b5f9054906101000a900460ff16151514610dff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610df690617db6565b60405180910390fd5b868660079182610e10929190617fd8565b50848460089182610e22929190617fd8565b50828260099182610e34929190617fd8565b5050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ea8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ecc9190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3090617abe565b60405180910390fd5b8060011515610f4782615080565b151514610f89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8090618115565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166384b7a84a86868660046040518563ffffffff1660e01b8152600401610fe8949392919061816c565b5f604051808303815f87803b158015610fff575f80fd5b505af1158015611011573d5f803e3d5ffd5b5050505060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e98ac22d600787878760066040518663ffffffff1660e01b8152600401611079959493929190618287565b5f604051808303815f87803b158015611090575f80fd5b505af11580156110a2573d5f803e3d5ffd5b505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611117573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061113b9190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119f90617abe565b60405180910390fd5b80600115156111b682615080565b1515146111f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ef90618115565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630199abce8686866040518463ffffffff1660e01b8152600401611256939291906182ed565b5f604051808303815f87803b15801561126d575f80fd5b505af115801561127f573d5f803e3d5ffd5b505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113189190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611385576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137c90617abe565b60405180910390fd5b806001151561139382615080565b1515146113d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113cc90618115565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cdee5b5e85856040518363ffffffff1660e01b815260040161143192919061831d565b5f604051808303815f87803b158015611448575f80fd5b505af115801561145a573d5f803e3d5ffd5b5050505050505050565b5f61146d61632e565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156114b55750825b90505f60018367ffffffffffffffff161480156114e857505f3073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156114f6575080155b1561152d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561157a576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168c73ffffffffffffffffffffffffffffffffffffffff16036115e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115df906183af565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff1603611656576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164d9061843d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff16036116c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116bb906184cb565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1603611732576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161172990618559565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff16036117a0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611797906185e7565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff160361180e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161180590618675565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff160361187c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161187390618703565b60405180910390fd5b8b60065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a60045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508960015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550885f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508760025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508315611a95575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611a8c919061876d565b60405180910390a15b505050505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b0d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b319190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b9590617abe565b60405180910390fd5b8060011515611bac82615080565b151514611bee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611be590618115565b60405180910390fd5b5f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630cc274938787876040518463ffffffff1660e01b8152600401611c4d93929190618786565b6020604051808303815f875af1158015611c69573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c8d9190617c76565b905060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e98ac22d600788885f866040518663ffffffff1660e01b8152600401611cf29594939291906187b6565b5f604051808303815f87803b158015611d09575f80fd5b505af1158015611d1b573d5f803e3d5ffd5b50505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d91573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611db59190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e1990617abe565b60405180910390fd5b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050611e6f81616355565b611eae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ea590618866565b60405180910390fd5b8160011515611ebc82615080565b151514611efe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ef590618115565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3483a9d878a8a898960016040518763ffffffff1660e01b8152600401611f61969594939291906188b4565b5f604051808303815f87803b158015611f78575f80fd5b505af1158015611f8a573d5f803e3d5ffd5b5050505060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e98ac22d60078a8a8a60046040518663ffffffff1660e01b8152600401611ff2959493929190618909565b5f604051808303815f87803b158015612009575f80fd5b505af115801561201b573d5f803e3d5ffd5b505050505050505050505050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612094573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120b89190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612125576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161211c90617abe565b60405180910390fd5b5f801515600b5f9054906101000a900460ff1615151461217a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161217190617db6565b60405180910390fd5b6001600b5f6101000a81548160ff0219169083151502179055507f04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf600b5f9054906101000a900460ff166040516121d19190616da0565b60405180910390a1600b5f9054906101000a900460ff1691505090565b5f600b5f9054906101000a900460ff1661220b57600190506122af565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166345a59e5b87878787876040518663ffffffff1660e01b815260040161226d95949392919061897e565b602060405180830381865afa158015612288573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122ac91906189d9565b90505b95945050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612322573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123469190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146123b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123aa90617abe565b60405180910390fd5b80600115156123c182615080565b151514612403576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123fa90618115565b60405180910390fd5b6124976007805461241390617e0b565b80601f016020809104026020016040519081016040528092919081815260200182805461243f90617e0b565b801561248a5780601f106124615761010080835404028352916020019161248a565b820191905f5260205f20905b81548152906001019060200180831161246d57829003601f168201915b50505050508360066163f6565b15612529575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166384b7a84a86868660056040518563ffffffff1660e01b81526004016124fb9493929190618a3d565b5f604051808303815f87803b158015612512575f80fd5b505af1158015612524573d5f803e3d5ffd5b505050505b5050505050565b5f600b5f9054906101000a900460ff16905090565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125af573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125d39190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612640576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161263790617abe565b60405180910390fd5b5f801515600b5f9054906101000a900460ff16151514612695576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161268c90617db6565b60405180910390fd5b612729600780546126a590617e0b565b80601f01602080910402602001604051908101604052809291908181526020018280546126d190617e0b565b801561271c5780601f106126f35761010080835404028352916020019161271c565b820191905f5260205f20905b8154815290600101906020018083116126ff57829003601f168201915b505050505083600161649e565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3483a9d836007600860026040518563ffffffff1660e01b815260040161278a9493929190618ab4565b5f604051808303815f87803b1580156127a1575f80fd5b505af11580156127b3573d5f803e3d5ffd5b505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612825573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128499190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146128b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128ad90617abe565b60405180910390fd5b82828080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050506129038161628a565b612942576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161293990617ceb565b60405180910390fd5b8184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050600115156129958383613fa6565b1515146129d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129ce90617b4c565b60405180910390fd5b60086040516020016129e99190618b05565b604051602081830303815290604052805190602001208888604051602001612a1292919061831d565b6040516020818303038152906040528051906020012014158015612a8557506009604051602001612a439190618b05565b604051602081830303815290604052805190602001208888604051602001612a6c92919061831d565b6040516020818303038152906040528051906020012014155b612ac4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612abb90618b6f565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a6343012898989896040518563ffffffff1660e01b8152600401612b249493929190618b8d565b5f604051808303815f87803b158015612b3b575f80fd5b505af1158015612b4d573d5f803e3d5ffd5b505050505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612bc5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612be99190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612c56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c4d90617abe565b60405180910390fd5b8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050612ca381616355565b612ce2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cd990618866565b60405180910390fd5b818c8c8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505060011515612d358383613fa6565b151514612d77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d6e90617b4c565b60405180910390fd5b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631f9534808f8f8f8f6040518563ffffffff1660e01b8152600401612dd79493929190618b8d565b5f604051808303815f87803b158015612dee575f80fd5b505af1158015612e00573d5f803e3d5ffd5b505050505f8e8e8e8e604051602001612e1c9493929190618c3e565b60405160208183030381529060405290505f8b8b90501115612ecc5760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634c5733118c8c8c8c8c8c886040518863ffffffff1660e01b8152600401612e9e9796959493929190618c70565b5f604051808303815f87803b158015612eb5575f80fd5b505af1158015612ec7573d5f803e3d5ffd5b505050505b505050505050505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f47573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f6b9190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612fd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fcf90617abe565b60405180910390fd5b8060011515612fe682615080565b151514613028576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161301f90618115565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633965c04a846040518263ffffffff1660e01b8152600401613082919061694a565b5f604051808303815f87803b158015613099575f80fd5b505af11580156130ab573d5f803e3d5ffd5b50505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636b568d7684846040518363ffffffff1660e01b8152600401613110929190618cda565b602060405180830381865afa15801561312b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061314f91906189d9565b905092915050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156131c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131e59190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161324990617abe565b60405180910390fd5b5f801515600b5f9054906101000a900460ff161515146132a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161329e90617db6565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634530abe188888888888860076040518863ffffffff1660e01b815260040161330e9796959493929190618d08565b5f604051808303815f87803b158015613325575f80fd5b505af1158015613337573d5f803e3d5ffd5b5050505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133ae573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133d29190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461343f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161343690617abe565b60405180910390fd5b806001151561344d82615080565b15151461348f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161348690618115565b60405180910390fd5b6135236007805461349f90617e0b565b80601f01602080910402602001604051908101604052809291908181526020018280546134cb90617e0b565b80156135165780601f106134ed57610100808354040283529160200191613516565b820191905f5260205f20905b8154815290600101906020018083116134f957829003601f168201915b50505050508360046163f6565b156137a0575f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631d09dc9388886040518363ffffffff1660e01b815260040161358492919061831d565b60408051808303815f875af115801561359f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135c39190618d72565b91509150811561366157613660600780546135dd90617e0b565b80601f016020809104026020016040519081016040528092919081815260200182805461360990617e0b565b80156136545780601f1061362b57610100808354040283529160200191613654565b820191905f5260205f20905b81548152906001019060200180831161363757829003601f168201915b5050505050825f61649e565b5b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c214e5e58989896040518463ffffffff1660e01b81526004016136bf939291906182ed565b6020604051808303815f875af11580156136db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136ff91906189d9565b9050801561379c5761379b6007805461371790617e0b565b80601f016020809104026020016040519081016040528092919081815260200182805461374390617e0b565b801561378e5780601f106137655761010080835404028352916020019161378e565b820191905f5260205f20905b81548152906001019060200180831161377157829003601f168201915b505050505087600161649e565b5b5050505b5050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613811573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138359190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146138a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161389990617abe565b60405180910390fd5b8083600115156138b28383613fa6565b1515146138f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138eb90617b4c565b60405180910390fd5b846138fe8161628a565b61393d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161393490617ceb565b60405180910390fd5b6001151561394b88886130b4565b15151461398d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161398490618dfa565b60405180910390fd5b6001151561399b86886165c1565b1515146139dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139d490618e62565b60405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663be322e548789613a268b61666f565b6040518463ffffffff1660e01b8152600401613a4493929190618e80565b602060405180830381865afa158015613a5f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a8391906189d9565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663143a5604898989856040518563ffffffff1660e01b8152600401613ae39493929190618eca565b5f604051808303815f87803b158015613afa575f80fd5b505af1158015613b0c573d5f803e3d5ffd5b505050505050505050505050565b5f600b5f9054906101000a900460ff16613b375760019050613f9b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd4fa05a8a6040518263ffffffff1660e01b8152600401613b919190618f1b565b602060405180830381865afa158015613bac573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613bd09190617c76565b90506002811480613be057505f81145b15613f96575f600190505f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1603613c265760029050613c37565b5f858590501115613c3657600390505b5b5f8203613d6f5760015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4604b6160015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630dca12f66040518163ffffffff1660e01b8152600401602060405180830381865afa158015613ce5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613d099190617c76565b836040518363ffffffff1660e01b8152600401613d27929190618f34565b602060405180830381865afa158015613d42573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613d6691906189d9565b92505050613f9b565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636acee5fd8d6040518263ffffffff1660e01b8152600401613dc99190618f1b565b5f60405180830381865afa158015613de3573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190613e0b9190618fc9565b915091505f613e198361666f565b905060045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633fd62ae7846040518263ffffffff1660e01b8152600401613e75919061903f565b602060405180830381865afa158015613e90573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613eb491906189d9565b15613f9157613ec28d615080565b80613ed35750613ed28d84613fa6565b5b15613ee657600195505050505050613f9b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d1f77866838584886040518563ffffffff1660e01b8152600401613f46949392919061905f565b602060405180830381865afa158015613f61573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f8591906189d9565b95505050505050613f9b565b505050505b5f9150505b979650505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8b42bf48484613fee8661666f565b6040518463ffffffff1660e01b815260040161400c939291906190b7565b602060405180830381865afa158015614027573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061404b91906189d9565b15614059576001905061419a565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663be322e545f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166381d66b23866040518263ffffffff1660e01b81526004016140ee9190618f1b565b5f60405180830381865afa158015614108573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061413091906190fa565b8461413a8661666f565b6040518463ffffffff1660e01b815260040161415893929190618e80565b602060405180830381865afa158015614173573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061419791906189d9565b90505b92915050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561420a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061422e9190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461429b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161429290617abe565b60405180910390fd5b5f801515600b5f9054906101000a900460ff161515146142f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016142e790617db6565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639fc666b2836040518263ffffffff1660e01b815260040161434a9190616da0565b5f604051808303815f87803b158015614361575f80fd5b505af1158015614373573d5f803e3d5ffd5b505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156143e5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906144099190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614476576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161446d90617abe565b60405180910390fd5b806001151561448482615080565b1515146144c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016144bd90618115565b60405180910390fd5b61455a600780546144d690617e0b565b80601f016020809104026020016040519081016040528092919081815260200182805461450290617e0b565b801561454d5780601f106145245761010080835404028352916020019161454d565b820191905f5260205f20905b81548152906001019060200180831161453057829003601f168201915b50505050508360056163f6565b156145f25760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d50b27878787878c60056040518763ffffffff1660e01b81526004016145c496959493929190619141565b5f604051808303815f87803b1580156145db575f80fd5b505af11580156145ed573d5f803e3d5ffd5b505050505b50505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614665573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906146899190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146146f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016146ed90617abe565b60405180910390fd5b5f801515600b5f9054906101000a900460ff1615151461474b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161474290617db6565b60405180910390fd5b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639e58eb9f600785856040518463ffffffff1660e01b81526004016147aa939291906191b5565b5f604051808303815f87803b1580156147c1575f80fd5b505af11580156147d3573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637b71357960086007600a546001806040518663ffffffff1660e01b815260040161483e9594939291906191f1565b5f604051808303815f87803b158015614855575f80fd5b505af1158015614867573d5f803e3d5ffd5b505050505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cef7f6af600860096040518363ffffffff1660e01b81526004016148c7929190619250565b5f604051808303815f87803b1580156148de575f80fd5b505af11580156148f0573d5f803e3d5ffd5b50505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614963573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906149879190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146149f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149eb90617abe565b60405180910390fd5b8060011515614a0282615080565b151514614a44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a3b90618115565b60405180910390fd5b6001831480614a535750600283145b614a92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a89906192cf565b60405180910390fd5b5f8060018503614aa9576002915060039050614abb565b60028503614aba5760039150600590505b5b60011515614b0c88888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505083616714565b151514614b4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b4590619337565b60405180910390fd5b614be160078054614b5e90617e0b565b80601f0160208091040260200160405190810160405280929190818152602001828054614b8a90617e0b565b8015614bd55780601f10614bac57610100808354040283529160200191614bd5565b820191905f5260205f20905b815481529060010190602001808311614bb857829003601f168201915b505050505085846163f6565b15614c725760045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166314f775f98888886040518463ffffffff1660e01b8152600401614c4493929190618786565b5f604051808303815f87803b158015614c5b575f80fd5b505af1158015614c6d573d5f803e3d5ffd5b505050505b50505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614ce5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614d099190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614d76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d6d90617abe565b60405180910390fd5b60011515614d848289613fa6565b151514614dc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614dbd90617b4c565b60405180910390fd5b6001821480614dd55750600282145b80614de05750600382145b614e1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614e1690617bda565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d50b27878787878c886040518763ffffffff1660e01b8152600401614e8396959493929190619355565b5f604051808303815f87803b158015614e9a575f80fd5b505af1158015614eac573d5f803e3d5ffd5b5050505050505050505050565b60608060605f600760086009600b5f9054906101000a900460ff16838054614ee090617e0b565b80601f0160208091040260200160405190810160405280929190818152602001828054614f0c90617e0b565b8015614f575780601f10614f2e57610100808354040283529160200191614f57565b820191905f5260205f20905b815481529060010190602001808311614f3a57829003601f168201915b50505050509350828054614f6a90617e0b565b80601f0160208091040260200160405190810160405280929190818152602001828054614f9690617e0b565b8015614fe15780601f10614fb857610100808354040283529160200191614fe1565b820191905f5260205f20905b815481529060010190602001808311614fc457829003601f168201915b50505050509250818054614ff490617e0b565b80601f016020809104026020016040519081016040528092919081815260200182805461502090617e0b565b801561506b5780601f106150425761010080835404028352916020019161506b565b820191905f5260205f20905b81548152906001019060200180831161504e57829003601f168201915b50505050509150935093509350935090919293565b5f60086040516020016150939190618b05565b604051602081830303815290604052805190602001205f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166381d66b23846040518263ffffffff1660e01b81526004016151019190618f1b565b5f60405180830381865afa15801561511b573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061514391906190fa565b604051602001615153919061903f565b60405160208183030381529060405280519060200120149050919050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156151db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906151ff9190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461526c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161526390617abe565b60405180910390fd5b806001151561527a82615080565b1515146152bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016152b390618115565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166307020b68846040518263ffffffff1660e01b81526004016153169190618f1b565b5f604051808303815f87803b15801561532d575f80fd5b505af115801561533f573d5f803e3d5ffd5b50505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156153b2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906153d69190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614615443576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161543a90617abe565b60405180910390fd5b806001151561545182615080565b151514615493576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161548a90618115565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337d50b27878787878c60046040518763ffffffff1660e01b81526004016154f8969594939291906193c9565b5f604051808303815f87803b15801561550f575f80fd5b505af1158015615521573d5f803e3d5ffd5b5050505060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e98ac22d600789895f60056040518663ffffffff1660e01b815260040161558995949392919061943d565b5f604051808303815f87803b1580156155a0575f80fd5b505af11580156155b2573d5f803e3d5ffd5b5050505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615629573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061564d9190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146156ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016156b190617abe565b60405180910390fd5b60011515600b5f9054906101000a900460ff1615151461570f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161570690617db6565b60405180910390fd5b6001151561571c82615080565b15151461575e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161575590618115565b60405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e98ac22d600789898660016040518663ffffffff1660e01b81526004016157c29594939291906194a3565b5f604051808303815f87803b1580156157d9575f80fd5b505af11580156157eb573d5f803e3d5ffd5b5050505060045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f9953de5886040518263ffffffff1660e01b8152600401615849919061903f565b5f604051808303815f87803b158015615860575f80fd5b505af1158015615872573d5f803e3d5ffd5b5050505060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663549583df878787878c6040518663ffffffff1660e01b81526004016158d8959493929190619509565b5f604051808303815f87803b1580156158ef575f80fd5b505af1158015615901573d5f803e3d5ffd5b505050506001151561591383896130b4565b151514615955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161594c906195b9565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3483a9d8389600960016040518563ffffffff1660e01b81526004016159b594939291906195d7565b5f604051808303815f87803b1580156159cc575f80fd5b505af11580156159de573d5f803e3d5ffd5b5050505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615a55573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615a799190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614615ae6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615add90617abe565b60405180910390fd5b85615af08161628a565b615b2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b2690617ceb565b60405180910390fd5b60011515615b3d8389613fa6565b151514615b7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b7690617b4c565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634c573311878787878c6040518663ffffffff1660e01b8152600401615be1959493929190619509565b5f604051808303815f87803b158015615bf8575f80fd5b505af1158015615c0a573d5f803e3d5ffd5b5050505050505050505050565b6060805f8060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663014e6acc87876040518363ffffffff1660e01b8152600401615c7892919061831d565b5f60405180830381865afa158015615c92573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190615cba9190619628565b935093509350935092959194509250565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614615d5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615d519061970e565b60405180910390fd5b5f801515600b5f9054906101000a900460ff16151514615daf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615da690617db6565b60405180910390fd5b878760079182615dc0929190617fd8565b50858560089182615dd2929190617fd8565b50838360099182615de4929190617fd8565b5081600b5f6101000a81548160ff0219169083151502179055505050505050505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e572515c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615e72573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615e969190617a23565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614615f03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615efa90617abe565b60405180910390fd5b60011515615f1082615080565b151514615f52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f4990618115565b60405180910390fd5b60011515615f61886001616714565b151514615fa3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f9a90619776565b60405180910390fd5b61603760078054615fb390617e0b565b80601f0160208091040260200160405190810160405280929190818152602001828054615fdf90617e0b565b801561602a5780601f106160015761010080835404028352916020019161602a565b820191905f5260205f20905b81548152906001019060200180831161600d57829003601f168201915b50505050508260016163f6565b156162815760045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3028316886040518263ffffffff1660e01b8152600401616096919061903f565b5f604051808303815f87803b1580156160ad575f80fd5b505af11580156160bf573d5f803e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637b713579600989600a546001806040518663ffffffff1660e01b8152600401616129959493929190619794565b5f604051808303815f87803b158015616140575f80fd5b505af1158015616152573d5f803e3d5ffd5b5050505060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f82e08ac878787878c6040518663ffffffff1660e01b81526004016161b8959493929190619509565b5f604051808303815f87803b1580156161cf575f80fd5b505af11580156161e1573d5f803e3d5ffd5b505050505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c214e5e588846040518363ffffffff1660e01b815260040161623f9291906197f3565b6020604051808303815f875af115801561625b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061627f91906189d9565b505b50505050505050565b5f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638c8642df8360026040518363ffffffff1660e01b81526004016162e8929190619821565b602060405180830381865afa158015616303573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061632791906189d9565b9050919050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ffe40d1d836040518263ffffffff1660e01b81526004016163b0919061903f565b602060405180830381865afa1580156163cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906163ef91906189d9565b9050919050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b02138648585856040518463ffffffff1660e01b81526004016164559392919061984f565b6020604051808303815f875af1158015616471573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061649591906189d9565b90509392505050565b80156165325760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635607395b84846040518363ffffffff1660e01b81526004016165009291906197f3565b5f604051808303815f87803b158015616517575f80fd5b505af1158015616529573d5f803e3d5ffd5b505050506165bc565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359cbd6fe84846040518363ffffffff1660e01b815260040161658e9291906197f3565b5f604051808303815f87803b1580156165a5575f80fd5b505af11580156165b7573d5f803e3d5ffd5b505050505b505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663abf5739f848461660a8661666f565b6040518463ffffffff1660e01b815260040161662893929190618e80565b602060405180830381865afa158015616643573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061666791906189d9565b905092915050565b606060045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663177c8d8a836040518263ffffffff1660e01b81526004016166cb919061903f565b5f60405180830381865afa1580156166e5573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061670d91906190fa565b9050919050565b5f60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638c8642df84846040518363ffffffff1660e01b815260040161677192919061988b565b602060405180830381865afa15801561678c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906167b091906189d9565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126167ea576167e96167c9565b5b8235905067ffffffffffffffff811115616807576168066167cd565b5b602083019150836001820283011115616823576168226167d1565b5b9250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6168538261682a565b9050919050565b61686381616849565b811461686d575f80fd5b50565b5f8135905061687e8161685a565b92915050565b5f819050919050565b61689681616884565b81146168a0575f80fd5b50565b5f813590506168b18161688d565b92915050565b5f805f805f608086880312156168d0576168cf6167c1565b5b5f86013567ffffffffffffffff8111156168ed576168ec6167c5565b5b6168f9888289016167d5565b9550955050602061690c88828901616870565b935050604061691d888289016168a3565b925050606061692e88828901616870565b9150509295509295909350565b61694481616884565b82525050565b5f60208201905061695d5f83018461693b565b92915050565b5f8115159050919050565b61697781616963565b8114616981575f80fd5b50565b5f813590506169928161696e565b92915050565b5f805f805f805f8060c0898b0312156169b4576169b36167c1565b5b5f89013567ffffffffffffffff8111156169d1576169d06167c5565b5b6169dd8b828c016167d5565b9850985050602089013567ffffffffffffffff811115616a00576169ff6167c5565b5b616a0c8b828c016167d5565b96509650506040616a1f8b828c016168a3565b9450506060616a308b828c01616984565b9350506080616a418b828c01616984565b92505060a0616a528b828c01616870565b9150509295985092959890939650565b5f805f805f8060608789031215616a7c57616a7b6167c1565b5b5f87013567ffffffffffffffff811115616a9957616a986167c5565b5b616aa589828a016167d5565b9650965050602087013567ffffffffffffffff811115616ac857616ac76167c5565b5b616ad489828a016167d5565b9450945050604087013567ffffffffffffffff811115616af757616af66167c5565b5b616b0389828a016167d5565b92509250509295509295509295565b5f805f8060608587031215616b2a57616b296167c1565b5b5f85013567ffffffffffffffff811115616b4757616b466167c5565b5b616b53878288016167d5565b94509450506020616b6687828801616870565b9250506040616b7787828801616870565b91505092959194509250565b5f805f60408486031215616b9a57616b996167c1565b5b5f84013567ffffffffffffffff811115616bb757616bb66167c5565b5b616bc3868287016167d5565b93509350506020616bd686828701616870565b9150509250925092565b5f805f805f805f60e0888a031215616bfb57616bfa6167c1565b5b5f616c088a828b01616870565b9750506020616c198a828b01616870565b9650506040616c2a8a828b01616870565b9550506060616c3b8a828b01616870565b9450506080616c4c8a828b01616870565b93505060a0616c5d8a828b01616870565b92505060c0616c6e8a828b01616870565b91505092959891949750929550565b5f805f8060608587031215616c9557616c946167c1565b5b5f85013567ffffffffffffffff811115616cb257616cb16167c5565b5b616cbe878288016167d5565b94509450506020616cd1878288016168a3565b9250506040616ce287828801616870565b91505092959194509250565b5f805f805f8060808789031215616d0857616d076167c1565b5b5f87013567ffffffffffffffff811115616d2557616d246167c5565b5b616d3189828a016167d5565b96509650506020616d4489828a01616870565b945050604087013567ffffffffffffffff811115616d6557616d646167c5565b5b616d7189828a016167d5565b93509350506060616d8489828a01616870565b9150509295509295509295565b616d9a81616963565b82525050565b5f602082019050616db35f830184616d91565b92915050565b5f61ffff82169050919050565b616dcf81616db9565b8114616dd9575f80fd5b50565b5f81359050616dea81616dc6565b92915050565b5f805f805f60608688031215616e0957616e086167c1565b5b5f86013567ffffffffffffffff811115616e2657616e256167c5565b5b616e32888289016167d5565b9550955050602086013567ffffffffffffffff811115616e5557616e546167c5565b5b616e61888289016167d5565b93509350506040616e7488828901616ddc565b9150509295509295909350565b5f60208284031215616e9657616e956167c1565b5b5f616ea384828501616870565b91505092915050565b5f805f805f60608688031215616ec557616ec46167c1565b5b5f86013567ffffffffffffffff811115616ee257616ee16167c5565b5b616eee888289016167d5565b9550955050602086013567ffffffffffffffff811115616f1157616f106167c5565b5b616f1d888289016167d5565b93509350506040616f3088828901616870565b9150509295509295909350565b5f805f805f805f805f805f60e08c8e031215616f5c57616f5b6167c1565b5b5f8c013567ffffffffffffffff811115616f7957616f786167c5565b5b616f858e828f016167d5565b9b509b505060208c013567ffffffffffffffff811115616fa857616fa76167c5565b5b616fb48e828f016167d5565b995099505060408c013567ffffffffffffffff811115616fd757616fd66167c5565b5b616fe38e828f016167d5565b975097505060608c013567ffffffffffffffff811115617006576170056167c5565b5b6170128e828f016167d5565b955095505060806170258e828f01616ddc565b93505060a06170368e828f01616ddc565b92505060c06170478e828f01616870565b9150509295989b509295989b9093969950565b5f80604083850312156170705761706f6167c1565b5b5f61707d858286016168a3565b925050602061708e85828601616870565b9150509250929050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6170e28261709c565b810181811067ffffffffffffffff82111715617101576171006170ac565b5b80604052505050565b5f6171136167b8565b905061711f82826170d9565b919050565b5f67ffffffffffffffff82111561713e5761713d6170ac565b5b6171478261709c565b9050602081019050919050565b828183375f83830152505050565b5f61717461716f84617124565b61710a565b9050828152602081018484840111156171905761718f617098565b5b61719b848285617154565b509392505050565b5f82601f8301126171b7576171b66167c9565b5b81356171c7848260208601617162565b91505092915050565b5f80604083850312156171e6576171e56167c1565b5b5f6171f385828601616870565b925050602083013567ffffffffffffffff811115617214576172136167c5565b5b617220858286016171a3565b9150509250929050565b5f805f805f8060808789031215617244576172436167c1565b5b5f87013567ffffffffffffffff811115617261576172606167c5565b5b61726d89828a016167d5565b9650965050602087013567ffffffffffffffff8111156172905761728f6167c5565b5b61729c89828a016167d5565b945094505060406172af89828a01616ddc565b92505060606172c089828a01616ddc565b9150509295509295509295565b5f805f80608085870312156172e5576172e46167c1565b5b5f6172f287828801616870565b945050602085013567ffffffffffffffff811115617313576173126167c5565b5b61731f878288016171a3565b935050604085013567ffffffffffffffff8111156173405761733f6167c5565b5b61734c878288016171a3565b925050606061735d87828801616870565b91505092959194509250565b5f8083601f84011261737e5761737d6167c9565b5b8235905067ffffffffffffffff81111561739b5761739a6167cd565b5b6020830191508360018202830111156173b7576173b66167d1565b5b9250929050565b5f805f805f805f60c0888a0312156173d9576173d86167c1565b5b5f6173e68a828b01616870565b97505060206173f78a828b01616870565b96505060406174088a828b016168a3565b95505060606174198a828b016168a3565b945050608061742a8a828b016168a3565b93505060a088013567ffffffffffffffff81111561744b5761744a6167c5565b5b6174578a828b01617369565b925092505092959891949750929550565b5f6020828403121561747d5761747c6167c1565b5b5f61748a84828501616984565b91505092915050565b5f805f805f8060c087890312156174ad576174ac6167c1565b5b5f87013567ffffffffffffffff8111156174ca576174c96167c5565b5b6174d689828a016171a3565b965050602087013567ffffffffffffffff8111156174f7576174f66167c5565b5b61750389828a016171a3565b955050604087013567ffffffffffffffff811115617524576175236167c5565b5b61753089828a016171a3565b945050606061754189828a01616ddc565b935050608061755289828a01616ddc565b92505060a061756389828a01616870565b9150509295509295509295565b5f8060408385031215617586576175856167c1565b5b5f617593858286016168a3565b92505060206175a4858286016168a3565b9150509250929050565b5f805f805f805f60e0888a0312156175c9576175c86167c1565b5b5f88013567ffffffffffffffff8111156175e6576175e56167c5565b5b6175f28a828b016171a3565b975050602088013567ffffffffffffffff811115617613576176126167c5565b5b61761f8a828b016171a3565b965050604088013567ffffffffffffffff8111156176405761763f6167c5565b5b61764c8a828b016171a3565b955050606061765d8a828b01616ddc565b945050608061766e8a828b01616ddc565b93505060a061767f8a828b016168a3565b92505060c06176908a828b01616870565b91505092959891949750929550565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156176d65780820151818401526020810190506176bb565b5f8484015250505050565b5f6176eb8261769f565b6176f581856176a9565b93506177058185602086016176b9565b61770e8161709c565b840191505092915050565b5f6080820190508181035f83015261773181876176e1565b9050818103602083015261774581866176e1565b9050818103604083015261775981856176e1565b90506177686060830184616d91565b95945050505050565b5f8060408385031215617787576177866167c1565b5b5f61779485828601616870565b92505060206177a585828601616870565b9150509250929050565b5f805f805f805f60e0888a0312156177ca576177c96167c1565b5b5f88013567ffffffffffffffff8111156177e7576177e66167c5565b5b6177f38a828b016171a3565b975050602088013567ffffffffffffffff811115617814576178136167c5565b5b6178208a828b016171a3565b965050604088013567ffffffffffffffff811115617841576178406167c5565b5b61784d8a828b016171a3565b955050606061785e8a828b01616ddc565b945050608061786f8a828b01616ddc565b93505060a06178808a828b01616870565b92505060c06178918a828b01616870565b91505092959891949750929550565b5f80602083850312156178b6576178b56167c1565b5b5f83013567ffffffffffffffff8111156178d3576178d26167c5565b5b6178df858286016167d5565b92509250509250929050565b6178f481616849565b82525050565b5f6080820190508181035f83015261791281876176e1565b9050818103602083015261792681866176e1565b905061793560408301856178eb565b617942606083018461693b565b95945050505050565b5f805f805f805f6080888a031215617966576179656167c1565b5b5f88013567ffffffffffffffff811115617983576179826167c5565b5b61798f8a828b016167d5565b9750975050602088013567ffffffffffffffff8111156179b2576179b16167c5565b5b6179be8a828b016167d5565b9550955050604088013567ffffffffffffffff8111156179e1576179e06167c5565b5b6179ed8a828b016167d5565b93509350506060617a008a828b01616984565b91505092959891949750929550565b5f81519050617a1d8161685a565b92915050565b5f60208284031215617a3857617a376167c1565b5b5f617a4584828501617a0f565b91505092915050565b7f63616e2062652063616c6c656420627920696e7465726661636520636f6e74725f8201527f616374206f6e6c79000000000000000000000000000000000000000000000000602082015250565b5f617aa86028836176a9565b9150617ab382617a4e565b604082019050919050565b5f6020820190508181035f830152617ad581617a9c565b9050919050565b7f6163636f756e74206973206e6f742061206f72672061646d696e206163636f755f8201527f6e74000000000000000000000000000000000000000000000000000000000000602082015250565b5f617b366022836176a9565b9150617b4182617adc565b604082019050919050565b5f6020820190508181035f830152617b6381617b2a565b9050919050565b7f696e76616c696420616374696f6e2e206f7065726174696f6e206e6f7420616c5f8201527f6c6f776564000000000000000000000000000000000000000000000000000000602082015250565b5f617bc46025836176a9565b9150617bcf82617b6a565b604082019050919050565b5f6020820190508181035f830152617bf181617bb8565b9050919050565b5f617c0383856176a9565b9350617c10838584617154565b617c198361709c565b840190509392505050565b5f6060820190508181035f830152617c3d818688617bf8565b9050617c4c60208301856178eb565b617c59604083018461693b565b95945050505050565b5f81519050617c708161688d565b92915050565b5f60208284031215617c8b57617c8a6167c1565b5b5f617c9884828501617c62565b91505092915050565b7f6f7267206e6f7420696e20617070726f766564207374617475730000000000005f82015250565b5f617cd5601a836176a9565b9150617ce082617ca1565b602082019050919050565b5f6020820190508181035f830152617d0281617cc9565b9050919050565b5f60a0820190508181035f830152617d2281898b617bf8565b90508181036020830152617d37818789617bf8565b9050617d46604083018661693b565b617d536060830185616d91565b617d606080830184616d91565b98975050505050505050565b7f496e636f7272656374206e6574776f726b20626f6f74207374617475730000005f82015250565b5f617da0601d836176a9565b9150617dab82617d6c565b602082019050919050565b5f6020820190508181035f830152617dcd81617d94565b9050919050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680617e2257607f821691505b602082108103617e3557617e34617dde565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302617e977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82617e5c565b617ea18683617e5c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f617edc617ed7617ed284616884565b617eb9565b616884565b9050919050565b5f819050919050565b617ef583617ec2565b617f09617f0182617ee3565b848454617e68565b825550505050565b5f90565b617f1d617f11565b617f28818484617eec565b505050565b5b81811015617f4b57617f405f82617f15565b600181019050617f2e565b5050565b601f821115617f9057617f6181617e3b565b617f6a84617e4d565b81016020851015617f79578190505b617f8d617f8585617e4d565b830182617f2d565b50505b505050565b5f82821c905092915050565b5f617fb05f1984600802617f95565b1980831691505092915050565b5f617fc88383617fa1565b9150826002028217905092915050565b617fe28383617dd4565b67ffffffffffffffff811115617ffb57617ffa6170ac565b5b6180058254617e0b565b618010828285617f4f565b5f601f83116001811461803d575f841561802b578287013590505b6180358582617fbd565b86555061809c565b601f19841661804b86617e3b565b5f5b828110156180725784890135825560018201915060208501945060208101905061804d565b8683101561808f578489013561808b601f891682617fa1565b8355505b6001600288020188555050505b50505050505050565b7f6163636f756e74206973206e6f742061206e6574776f726b2061646d696e20615f8201527f63636f756e740000000000000000000000000000000000000000000000000000602082015250565b5f6180ff6026836176a9565b915061810a826180a5565b604082019050919050565b5f6020820190508181035f83015261812c816180f3565b9050919050565b5f819050919050565b5f61815661815161814c84618133565b617eb9565b616884565b9050919050565b6181668161813c565b82525050565b5f6060820190508181035f830152618185818688617bf8565b905061819460208301856178eb565b6181a1604083018461815d565b95945050505050565b5f81546181b681617e0b565b6181c081866176a9565b9450600182165f81146181da57600181146181f057618222565b60ff198316865281151560200286019350618222565b6181f985617e3b565b5f5b8381101561821a578154818901526001820191506020810190506181fb565b808801955050505b50505092915050565b50565b5f6182395f836176a9565b91506182448261822b565b5f82019050919050565b5f819050919050565b5f61827161826c6182678461824e565b617eb9565b616884565b9050919050565b61828181618257565b82525050565b5f60a0820190508181035f83015261829f81886181aa565b905081810360208301526182b4818688617bf8565b905081810360408301526182c78161822e565b90506182d660608301856178eb565b6182e36080830184618278565b9695505050505050565b5f6040820190508181035f830152618306818587617bf8565b905061831560208301846178eb565b949350505050565b5f6020820190508181035f830152618336818486617bf8565b90509392505050565b7f5f7065726d55706772616461626c652063616e6e6f7420626520616e20656d705f8201527f7479206164647265737300000000000000000000000000000000000000000000602082015250565b5f618399602a836176a9565b91506183a48261833f565b604082019050919050565b5f6020820190508181035f8301526183c68161838d565b9050919050565b7f5f6f72674d616e616765722063616e6e6f7420626520616e20656d70747920615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6184276026836176a9565b9150618432826183cd565b604082019050919050565b5f6020820190508181035f8301526184548161841b565b9050919050565b7f5f726f6c65734d616e616765722063616e6e6f7420626520616e20656d7074795f8201527f2061646472657373000000000000000000000000000000000000000000000000602082015250565b5f6184b56028836176a9565b91506184c08261845b565b604082019050919050565b5f6020820190508181035f8301526184e2816184a9565b9050919050565b7f5f6163636f756e744d616e616765722063616e6e6f7420626520616e20656d705f8201527f7479206164647265737300000000000000000000000000000000000000000000602082015250565b5f618543602a836176a9565b915061854e826184e9565b604082019050919050565b5f6020820190508181035f83015261857081618537565b9050919050565b7f5f766f7465724d616e616765722063616e6e6f7420626520616e20656d7074795f8201527f2061646472657373000000000000000000000000000000000000000000000000602082015250565b5f6185d16028836176a9565b91506185dc82618577565b604082019050919050565b5f6020820190508181035f8301526185fe816185c5565b9050919050565b7f5f6e6f64654d616e616765722063616e6e6f7420626520616e20656d707479205f8201527f6164647265737300000000000000000000000000000000000000000000000000602082015250565b5f61865f6027836176a9565b915061866a82618605565b604082019050919050565b5f6020820190508181035f83015261868c81618653565b9050919050565b7f5f636f6e747261637457686974656c6973744d616e616765722063616e6e6f745f8201527f20626520616e20656d7074792061646472657373000000000000000000000000602082015250565b5f6186ed6034836176a9565b91506186f882618693565b604082019050919050565b5f6020820190508181035f83015261871a816186e1565b9050919050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f61875761875261874d84618721565b617eb9565b61872a565b9050919050565b6187678161873d565b82525050565b5f6020820190506187805f83018461875e565b92915050565b5f6040820190508181035f83015261879f818587617bf8565b90506187ae602083018461693b565b949350505050565b5f60a0820190508181035f8301526187ce81886181aa565b905081810360208301526187e3818688617bf8565b905081810360408301526187f68161822e565b905061880560608301856178eb565b618812608083018461693b565b9695505050505050565b7f6f726720646f6573206e6f7420657869737400000000000000000000000000005f82015250565b5f6188506012836176a9565b915061885b8261881c565b602082019050919050565b5f6020820190508181035f83015261887d81618844565b9050919050565b5f61889e61889961889484618721565b617eb9565b616884565b9050919050565b6188ae81618884565b82525050565b5f6080820190506188c75f8301896178eb565b81810360208301526188da818789617bf8565b905081810360408301526188ef818587617bf8565b90506188fe60608301846188a5565b979650505050505050565b5f60a0820190508181035f83015261892181886181aa565b90508181036020830152618936818688617bf8565b905081810360408301526189498161822e565b905061895860608301856178eb565b618965608083018461815d565b9695505050505050565b61897881616db9565b82525050565b5f6060820190508181035f830152618997818789617bf8565b905081810360208301526189ac818587617bf8565b90506189bb604083018461896f565b9695505050505050565b5f815190506189d38161696e565b92915050565b5f602082840312156189ee576189ed6167c1565b5b5f6189fb848285016189c5565b91505092915050565b5f819050919050565b5f618a27618a22618a1d84618a04565b617eb9565b616884565b9050919050565b618a3781618a0d565b82525050565b5f6060820190508181035f830152618a56818688617bf8565b9050618a6560208301856178eb565b618a726040830184618a2e565b95945050505050565b5f819050919050565b5f618a9e618a99618a9484618a7b565b617eb9565b616884565b9050919050565b618aae81618a84565b82525050565b5f608082019050618ac75f8301876178eb565b8181036020830152618ad981866181aa565b90508181036040830152618aed81856181aa565b9050618afc6060830184618aa5565b95945050505050565b5f6020820190508181035f830152618b1d81846181aa565b905092915050565b7f61646d696e20726f6c65732063616e6e6f742062652072656d6f7665640000005f82015250565b5f618b59601d836176a9565b9150618b6482618b25565b602082019050919050565b5f6020820190508181035f830152618b8681618b4d565b9050919050565b5f6040820190508181035f830152618ba6818688617bf8565b90508181036020830152618bbb818486617bf8565b905095945050505050565b5f81905092915050565b5f618bdb8385618bc6565b9350618be8838584617154565b82840190509392505050565b7f2e000000000000000000000000000000000000000000000000000000000000005f82015250565b5f618c28600183618bc6565b9150618c3382618bf4565b600182019050919050565b5f618c4a828688618bd0565b9150618c5582618c1c565b9150618c62828486618bd0565b915081905095945050505050565b5f60a0820190508181035f830152618c8981898b617bf8565b90508181036020830152618c9e818789617bf8565b9050618cad604083018661896f565b618cba606083018561896f565b8181036080830152618ccc81846176e1565b905098975050505050505050565b5f604082019050618ced5f8301856178eb565b8181036020830152618cff81846176e1565b90509392505050565b5f60a0820190508181035f830152618d2181898b617bf8565b90508181036020830152618d36818789617bf8565b9050618d45604083018661896f565b618d52606083018561896f565b8181036080830152618d6481846181aa565b905098975050505050505050565b5f8060408385031215618d8857618d876167c1565b5b5f618d95858286016189c5565b9250506020618da685828601617a0f565b9150509250929050565b7f6f7065726174696f6e2063616e6e6f7420626520706572666f726d65640000005f82015250565b5f618de4601d836176a9565b9150618def82618db0565b602082019050919050565b5f6020820190508181035f830152618e1181618dd8565b9050919050565b7f726f6c6520646f6573206e6f74206578697374730000000000000000000000005f82015250565b5f618e4c6014836176a9565b9150618e5782618e18565b602082019050919050565b5f6020820190508181035f830152618e7981618e40565b9050919050565b5f6060820190508181035f830152618e9881866176e1565b90508181036020830152618eac81856176e1565b90508181036040830152618ec081846176e1565b9050949350505050565b5f608082019050618edd5f8301876178eb565b8181036020830152618eef81866176e1565b90508181036040830152618f0381856176e1565b9050618f126060830184616d91565b95945050505050565b5f602082019050618f2e5f8301846178eb565b92915050565b5f604082019050618f475f83018561693b565b618f54602083018461693b565b9392505050565b5f618f6d618f6884617124565b61710a565b905082815260208101848484011115618f8957618f88617098565b5b618f948482856176b9565b509392505050565b5f82601f830112618fb057618faf6167c9565b5b8151618fc0848260208601618f5b565b91505092915050565b5f8060408385031215618fdf57618fde6167c1565b5b5f83015167ffffffffffffffff811115618ffc57618ffb6167c5565b5b61900885828601618f9c565b925050602083015167ffffffffffffffff811115619029576190286167c5565b5b61903585828601618f9c565b9150509250929050565b5f6020820190508181035f83015261905781846176e1565b905092915050565b5f6080820190508181035f83015261907781876176e1565b9050818103602083015261908b81866176e1565b9050818103604083015261909f81856176e1565b90506190ae606083018461693b565b95945050505050565b5f6060820190506190ca5f8301866178eb565b81810360208301526190dc81856176e1565b905081810360408301526190f081846176e1565b9050949350505050565b5f6020828403121561910f5761910e6167c1565b5b5f82015167ffffffffffffffff81111561912c5761912b6167c5565b5b61913884828501618f9c565b91505092915050565b5f60c0820190508181035f83015261915981896176e1565b9050818103602083015261916d81886176e1565b905061917c604083018761896f565b619189606083018661896f565b818103608083015261919b81856176e1565b90506191aa60a0830184618a2e565b979650505050505050565b5f6060820190508181035f8301526191cd81866181aa565b90506191dc602083018561693b565b6191e9604083018461693b565b949350505050565b5f60a0820190508181035f83015261920981886181aa565b9050818103602083015261921d81876181aa565b905061922c604083018661693b565b6192396060830185616d91565b6192466080830184616d91565b9695505050505050565b5f6040820190508181035f83015261926881856181aa565b9050818103602083015261927c81846181aa565b90509392505050565b7f4f7065726174696f6e206e6f7420616c6c6f77656400000000000000000000005f82015250565b5f6192b96015836176a9565b91506192c482619285565b602082019050919050565b5f6020820190508181035f8301526192e6816192ad565b9050919050565b7f6f7065726174696f6e206e6f7420616c6c6f77656400000000000000000000005f82015250565b5f6193216015836176a9565b915061932c826192ed565b602082019050919050565b5f6020820190508181035f83015261934e81619315565b9050919050565b5f60c0820190508181035f83015261936d81896176e1565b9050818103602083015261938181886176e1565b9050619390604083018761896f565b61939d606083018661896f565b81810360808301526193af81856176e1565b90506193be60a083018461693b565b979650505050505050565b5f60c0820190508181035f8301526193e181896176e1565b905081810360208301526193f581886176e1565b9050619404604083018761896f565b619411606083018661896f565b818103608083015261942381856176e1565b905061943260a083018461815d565b979650505050505050565b5f60a0820190508181035f83015261945581886181aa565b9050818103602083015261946981876176e1565b9050818103604083015261947d81866176e1565b905061948c60608301856178eb565b6194996080830184618a2e565b9695505050505050565b5f60a0820190508181035f8301526194bb81886181aa565b905081810360208301526194cf81876176e1565b905081810360408301526194e381866176e1565b90506194f260608301856178eb565b6194ff60808301846188a5565b9695505050505050565b5f60a0820190508181035f83015261952181886176e1565b9050818103602083015261953581876176e1565b9050619544604083018661896f565b619551606083018561896f565b818103608083015261956381846176e1565b90509695505050505050565b7f4f7065726174696f6e2063616e6e6f7420626520706572666f726d65640000005f82015250565b5f6195a3601d836176a9565b91506195ae8261956f565b602082019050919050565b5f6020820190508181035f8301526195d081619597565b9050919050565b5f6080820190506195ea5f8301876178eb565b81810360208301526195fc81866176e1565b9050818103604083015261961081856181aa565b905061961f60608301846188a5565b95945050505050565b5f805f80608085870312156196405761963f6167c1565b5b5f85015167ffffffffffffffff81111561965d5761965c6167c5565b5b61966987828801618f9c565b945050602085015167ffffffffffffffff81111561968a576196896167c5565b5b61969687828801618f9c565b93505060406196a787828801617a0f565b92505060606196b887828801617c62565b91505092959194509250565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f6196f8600e836176a9565b9150619703826196c4565b602082019050919050565b5f6020820190508181035f830152619725816196ec565b9050919050565b7f4e6f7468696e6720746f20617070726f766500000000000000000000000000005f82015250565b5f6197606012836176a9565b915061976b8261972c565b602082019050919050565b5f6020820190508181035f83015261978d81619754565b9050919050565b5f60a0820190508181035f8301526197ac81886181aa565b905081810360208301526197c081876176e1565b90506197cf604083018661693b565b6197dc6060830185616d91565b6197e96080830184616d91565b9695505050505050565b5f6040820190508181035f83015261980b81856176e1565b905061981a60208301846178eb565b9392505050565b5f6040820190508181035f83015261983981856176e1565b90506198486020830184618aa5565b9392505050565b5f6060820190508181035f83015261986781866176e1565b905061987660208301856178eb565b619883604083018461693b565b949350505050565b5f6040820190508181035f8301526198a381856176e1565b90506198b2602083018461693b565b939250505056fea2646970667358221220d269fed2be828750f0b6524ecba11d0be3843533835c9c8b4d752b035195434f64736f6c63430008180033",
}

// PermImplABI is the input ABI used to generate the binding from.
// Deprecated: Use PermImplMetaData.ABI instead.
var PermImplABI = PermImplMetaData.ABI

var PermImplParsedABI, _ = abi.JSON(strings.NewReader(PermImplABI))

// PermImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PermImplMetaData.Bin instead.
var PermImplBin = PermImplMetaData.Bin

// DeployPermImpl deploys a new Ethereum contract, binding an instance of PermImpl to it.
func DeployPermImpl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PermImpl, error) {
	parsed, err := PermImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PermImplBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermImpl{PermImplCaller: PermImplCaller{contract: contract}, PermImplTransactor: PermImplTransactor{contract: contract}, PermImplFilterer: PermImplFilterer{contract: contract}}, nil
}

// PermImpl is an auto generated Go binding around an Ethereum contract.
type PermImpl struct {
	PermImplCaller     // Read-only binding to the contract
	PermImplTransactor // Write-only binding to the contract
	PermImplFilterer   // Log filterer for contract events
}

// PermImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermImplSession struct {
	Contract     *PermImpl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermImplCallerSession struct {
	Contract *PermImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PermImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermImplTransactorSession struct {
	Contract     *PermImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PermImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermImplRaw struct {
	Contract *PermImpl // Generic contract binding to access the raw methods on
}

// PermImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermImplCallerRaw struct {
	Contract *PermImplCaller // Generic read-only contract binding to access the raw methods on
}

// PermImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermImplTransactorRaw struct {
	Contract *PermImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermImpl creates a new instance of PermImpl, bound to a specific deployed contract.
func NewPermImpl(address common.Address, backend bind.ContractBackend) (*PermImpl, error) {
	contract, err := bindPermImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermImpl{PermImplCaller: PermImplCaller{contract: contract}, PermImplTransactor: PermImplTransactor{contract: contract}, PermImplFilterer: PermImplFilterer{contract: contract}}, nil
}

// NewPermImplCaller creates a new read-only instance of PermImpl, bound to a specific deployed contract.
func NewPermImplCaller(address common.Address, caller bind.ContractCaller) (*PermImplCaller, error) {
	contract, err := bindPermImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermImplCaller{contract: contract}, nil
}

// NewPermImplTransactor creates a new write-only instance of PermImpl, bound to a specific deployed contract.
func NewPermImplTransactor(address common.Address, transactor bind.ContractTransactor) (*PermImplTransactor, error) {
	contract, err := bindPermImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermImplTransactor{contract: contract}, nil
}

// NewPermImplFilterer creates a new log filterer instance of PermImpl, bound to a specific deployed contract.
func NewPermImplFilterer(address common.Address, filterer bind.ContractFilterer) (*PermImplFilterer, error) {
	contract, err := bindPermImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermImplFilterer{contract: contract}, nil
}

// bindPermImpl binds a generic wrapper to an already deployed contract.
func bindPermImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermImpl *PermImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermImpl.Contract.PermImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermImpl *PermImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.Contract.PermImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermImpl *PermImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermImpl.Contract.PermImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermImpl *PermImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermImpl *PermImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermImpl *PermImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermImpl.Contract.contract.Transact(opts, method, params...)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_PermImpl *PermImplCaller) ConnectionAllowed(opts *bind.CallOpts, _enodeId string, _ip string, _port uint16) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "connectionAllowed", _enodeId, _ip, _port)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_PermImpl *PermImplSession) ConnectionAllowed(_enodeId string, _ip string, _port uint16) (bool, error) {
	return _PermImpl.Contract.ConnectionAllowed(&_PermImpl.CallOpts, _enodeId, _ip, _port)
}

// ConnectionAllowed is a free data retrieval call binding the contract method 0x45a59e5b.
//
// Solidity: function connectionAllowed(string _enodeId, string _ip, uint16 _port) view returns(bool)
func (_PermImpl *PermImplCallerSession) ConnectionAllowed(_enodeId string, _ip string, _port uint16) (bool, error) {
	return _PermImpl.Contract.ConnectionAllowed(&_PermImpl.CallOpts, _enodeId, _ip, _port)
}

// GetAccessLevelForUnconfiguredAccount is a free data retrieval call binding the contract method 0x0dca12f6.
//
// Solidity: function getAccessLevelForUnconfiguredAccount() view returns(uint256)
func (_PermImpl *PermImplCaller) GetAccessLevelForUnconfiguredAccount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "getAccessLevelForUnconfiguredAccount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccessLevelForUnconfiguredAccount is a free data retrieval call binding the contract method 0x0dca12f6.
//
// Solidity: function getAccessLevelForUnconfiguredAccount() view returns(uint256)
func (_PermImpl *PermImplSession) GetAccessLevelForUnconfiguredAccount() (*big.Int, error) {
	return _PermImpl.Contract.GetAccessLevelForUnconfiguredAccount(&_PermImpl.CallOpts)
}

// GetAccessLevelForUnconfiguredAccount is a free data retrieval call binding the contract method 0x0dca12f6.
//
// Solidity: function getAccessLevelForUnconfiguredAccount() view returns(uint256)
func (_PermImpl *PermImplCallerSession) GetAccessLevelForUnconfiguredAccount() (*big.Int, error) {
	return _PermImpl.Contract.GetAccessLevelForUnconfiguredAccount(&_PermImpl.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() view returns(bool)
func (_PermImpl *PermImplCaller) GetNetworkBootStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "getNetworkBootStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() view returns(bool)
func (_PermImpl *PermImplSession) GetNetworkBootStatus() (bool, error) {
	return _PermImpl.Contract.GetNetworkBootStatus(&_PermImpl.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() view returns(bool)
func (_PermImpl *PermImplCallerSession) GetNetworkBootStatus() (bool, error) {
	return _PermImpl.Contract.GetNetworkBootStatus(&_PermImpl.CallOpts)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) view returns(string, string, address, uint256)
func (_PermImpl *PermImplCaller) GetPendingOp(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "getPendingOp", _orgId)

	if err != nil {
		return *new(string), *new(string), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) view returns(string, string, address, uint256)
func (_PermImpl *PermImplSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermImpl.Contract.GetPendingOp(&_PermImpl.CallOpts, _orgId)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(string _orgId) view returns(string, string, address, uint256)
func (_PermImpl *PermImplCallerSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermImpl.Contract.GetPendingOp(&_PermImpl.CallOpts, _orgId)
}

// GetPolicyDetails is a free data retrieval call binding the contract method 0xcc9ba6fa.
//
// Solidity: function getPolicyDetails() view returns(string, string, string, bool)
func (_PermImpl *PermImplCaller) GetPolicyDetails(opts *bind.CallOpts) (string, string, string, bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "getPolicyDetails")

	if err != nil {
		return *new(string), *new(string), *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetPolicyDetails is a free data retrieval call binding the contract method 0xcc9ba6fa.
//
// Solidity: function getPolicyDetails() view returns(string, string, string, bool)
func (_PermImpl *PermImplSession) GetPolicyDetails() (string, string, string, bool, error) {
	return _PermImpl.Contract.GetPolicyDetails(&_PermImpl.CallOpts)
}

// GetPolicyDetails is a free data retrieval call binding the contract method 0xcc9ba6fa.
//
// Solidity: function getPolicyDetails() view returns(string, string, string, bool)
func (_PermImpl *PermImplCallerSession) GetPolicyDetails() (string, string, string, bool, error) {
	return _PermImpl.Contract.GetPolicyDetails(&_PermImpl.CallOpts)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_PermImpl *PermImplCaller) IsNetworkAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "isNetworkAdmin", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_PermImpl *PermImplSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermImpl.Contract.IsNetworkAdmin(&_PermImpl.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_PermImpl *PermImplCallerSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermImpl.Contract.IsNetworkAdmin(&_PermImpl.CallOpts, _account)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) view returns(bool)
func (_PermImpl *PermImplCaller) IsOrgAdmin(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "isOrgAdmin", _account, _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) view returns(bool)
func (_PermImpl *PermImplSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.IsOrgAdmin(&_PermImpl.CallOpts, _account, _orgId)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(address _account, string _orgId) view returns(bool)
func (_PermImpl *PermImplCallerSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.IsOrgAdmin(&_PermImpl.CallOpts, _account, _orgId)
}

// TransactionAllowed is a free data retrieval call binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address _sender, address _target, uint256 _value, uint256 _gasPrice, uint256 _gasLimit, bytes _payload) view returns(bool)
func (_PermImpl *PermImplCaller) TransactionAllowed(opts *bind.CallOpts, _sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "transactionAllowed", _sender, _target, _value, _gasPrice, _gasLimit, _payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransactionAllowed is a free data retrieval call binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address _sender, address _target, uint256 _value, uint256 _gasPrice, uint256 _gasLimit, bytes _payload) view returns(bool)
func (_PermImpl *PermImplSession) TransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte) (bool, error) {
	return _PermImpl.Contract.TransactionAllowed(&_PermImpl.CallOpts, _sender, _target, _value, _gasPrice, _gasLimit, _payload)
}

// TransactionAllowed is a free data retrieval call binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address _sender, address _target, uint256 _value, uint256 _gasPrice, uint256 _gasLimit, bytes _payload) view returns(bool)
func (_PermImpl *PermImplCallerSession) TransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte) (bool, error) {
	return _PermImpl.Contract.TransactionAllowed(&_PermImpl.CallOpts, _sender, _target, _value, _gasPrice, _gasLimit, _payload)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_PermImpl *PermImplCaller) ValidateAccount(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "validateAccount", _account, _orgId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_PermImpl *PermImplSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.ValidateAccount(&_PermImpl.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(address _account, string _orgId) view returns(bool)
func (_PermImpl *PermImplCallerSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.ValidateAccount(&_PermImpl.CallOpts, _account, _orgId)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _account) returns()
func (_PermImpl *PermImplTransactor) AddAdminAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addAdminAccount", _account)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _account) returns()
func (_PermImpl *PermImplSession) AddAdminAccount(_account common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminAccount(&_PermImpl.TransactOpts, _account)
}

// AddAdminAccount is a paid mutator transaction binding the contract method 0x4fe57e7a.
//
// Solidity: function addAdminAccount(address _account) returns()
func (_PermImpl *PermImplTransactorSession) AddAdminAccount(_account common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminAccount(&_PermImpl.TransactOpts, _account)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x8683c7fe.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport) returns()
func (_PermImpl *PermImplTransactor) AddAdminNode(opts *bind.TransactOpts, _enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addAdminNode", _enodeId, _ip, _port, _raftport)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x8683c7fe.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport) returns()
func (_PermImpl *PermImplSession) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminNode(&_PermImpl.TransactOpts, _enodeId, _ip, _port, _raftport)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0x8683c7fe.
//
// Solidity: function addAdminNode(string _enodeId, string _ip, uint16 _port, uint16 _raftport) returns()
func (_PermImpl *PermImplTransactorSession) AddAdminNode(_enodeId string, _ip string, _port uint16, _raftport uint16) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminNode(&_PermImpl.TransactOpts, _enodeId, _ip, _port, _raftport)
}

// AddContractWhitelist is a paid mutator transaction binding the contract method 0x27bb2cad.
//
// Solidity: function addContractWhitelist(string _contractKey, address _contractAddress, address _caller) returns()
func (_PermImpl *PermImplTransactor) AddContractWhitelist(opts *bind.TransactOpts, _contractKey string, _contractAddress common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addContractWhitelist", _contractKey, _contractAddress, _caller)
}

// AddContractWhitelist is a paid mutator transaction binding the contract method 0x27bb2cad.
//
// Solidity: function addContractWhitelist(string _contractKey, address _contractAddress, address _caller) returns()
func (_PermImpl *PermImplSession) AddContractWhitelist(_contractKey string, _contractAddress common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddContractWhitelist(&_PermImpl.TransactOpts, _contractKey, _contractAddress, _caller)
}

// AddContractWhitelist is a paid mutator transaction binding the contract method 0x27bb2cad.
//
// Solidity: function addContractWhitelist(string _contractKey, address _contractAddress, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) AddContractWhitelist(_contractKey string, _contractAddress common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddContractWhitelist(&_PermImpl.TransactOpts, _contractKey, _contractAddress, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin, address _caller) returns()
func (_PermImpl *PermImplTransactor) AddNewRole(opts *bind.TransactOpts, _roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addNewRole", _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin, address _caller) returns()
func (_PermImpl *PermImplSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNewRole(&_PermImpl.TransactOpts, _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(string _roleId, string _orgId, uint256 _access, bool _voter, bool _admin, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNewRole(&_PermImpl.TransactOpts, _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0xecad01d5.
//
// Solidity: function addNode(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactor) AddNode(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addNode", _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0xecad01d5.
//
// Solidity: function addNode(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplSession) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNode(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0xecad01d5.
//
// Solidity: function addNode(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) AddNode(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNode(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xe91b0e19.
//
// Solidity: function addOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactor) AddOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addOrg", _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xe91b0e19.
//
// Solidity: function addOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermImpl *PermImplSession) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xe91b0e19.
//
// Solidity: function addOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) AddOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x68a61273.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactor) AddSubOrg(opts *bind.TransactOpts, _pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addSubOrg", _pOrgId, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x68a61273.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplSession) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddSubOrg(&_PermImpl.TransactOpts, _pOrgId, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x68a61273.
//
// Solidity: function addSubOrg(string _pOrgId, string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) AddSubOrg(_pOrgId string, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddSubOrg(&_PermImpl.TransactOpts, _pOrgId, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactor) ApproveAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveAdminRole", _orgId, _account, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplSession) ApproveAdminRole(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) ApproveAdminRole(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x4b20f45f.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactor) ApproveBlacklistedAccountRecovery(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveBlacklistedAccountRecovery", _orgId, _account, _caller)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x4b20f45f.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplSession) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveBlacklistedAccountRecovery(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// ApproveBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x4b20f45f.
//
// Solidity: function approveBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) ApproveBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveBlacklistedAccountRecovery(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xa042bf40.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactor) ApproveBlacklistedNodeRecovery(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveBlacklistedNodeRecovery", _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xa042bf40.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplSession) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveBlacklistedNodeRecovery(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xa042bf40.
//
// Solidity: function approveBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) ApproveBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveBlacklistedNodeRecovery(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xf75f0a06.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveOrg", _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xf75f0a06.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermImpl *PermImplSession) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xf75f0a06.
//
// Solidity: function approveOrg(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) ApproveOrg(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _account, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactor) ApproveOrgStatus(opts *bind.TransactOpts, _orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveOrgStatus", _orgId, _action, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplSession) ApproveOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrgStatus(&_PermImpl.TransactOpts, _orgId, _action, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) ApproveOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrgStatus(&_PermImpl.TransactOpts, _orgId, _action, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, address _caller) returns()
func (_PermImpl *PermImplTransactor) AssignAccountRole(opts *bind.TransactOpts, _account common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "assignAccountRole", _account, _orgId, _roleId, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, address _caller) returns()
func (_PermImpl *PermImplSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAccountRole(&_PermImpl.TransactOpts, _account, _orgId, _roleId, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(address _account, string _orgId, string _roleId, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) AssignAccountRole(_account common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAccountRole(&_PermImpl.TransactOpts, _account, _orgId, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId, address _caller) returns()
func (_PermImpl *PermImplTransactor) AssignAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "assignAdminRole", _orgId, _account, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId, address _caller) returns()
func (_PermImpl *PermImplSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(string _orgId, address _account, string _roleId, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _roleId, _caller)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermImpl *PermImplTransactor) Init(opts *bind.TransactOpts, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "init", _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermImpl *PermImplSession) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermImpl.Contract.Init(&_PermImpl.TransactOpts, _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function init(uint256 _breadth, uint256 _depth) returns()
func (_PermImpl *PermImplTransactorSession) Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermImpl.Contract.Init(&_PermImpl.TransactOpts, _breadth, _depth)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _permUpgradable, address _orgManager, address _rolesManager, address _accountManager, address _voterManager, address _nodeManager, address _contractWhitelistManager) returns()
func (_PermImpl *PermImplTransactor) Initialize(opts *bind.TransactOpts, _permUpgradable common.Address, _orgManager common.Address, _rolesManager common.Address, _accountManager common.Address, _voterManager common.Address, _nodeManager common.Address, _contractWhitelistManager common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "initialize", _permUpgradable, _orgManager, _rolesManager, _accountManager, _voterManager, _nodeManager, _contractWhitelistManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _permUpgradable, address _orgManager, address _rolesManager, address _accountManager, address _voterManager, address _nodeManager, address _contractWhitelistManager) returns()
func (_PermImpl *PermImplSession) Initialize(_permUpgradable common.Address, _orgManager common.Address, _rolesManager common.Address, _accountManager common.Address, _voterManager common.Address, _nodeManager common.Address, _contractWhitelistManager common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.Initialize(&_PermImpl.TransactOpts, _permUpgradable, _orgManager, _rolesManager, _accountManager, _voterManager, _nodeManager, _contractWhitelistManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _permUpgradable, address _orgManager, address _rolesManager, address _accountManager, address _voterManager, address _nodeManager, address _contractWhitelistManager) returns()
func (_PermImpl *PermImplTransactorSession) Initialize(_permUpgradable common.Address, _orgManager common.Address, _rolesManager common.Address, _accountManager common.Address, _voterManager common.Address, _nodeManager common.Address, _contractWhitelistManager common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.Initialize(&_PermImpl.TransactOpts, _permUpgradable, _orgManager, _rolesManager, _accountManager, _voterManager, _nodeManager, _contractWhitelistManager)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(string _roleId, string _orgId, address _caller) returns()
func (_PermImpl *PermImplTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "removeRole", _roleId, _orgId, _caller)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(string _roleId, string _orgId, address _caller) returns()
func (_PermImpl *PermImplSession) RemoveRole(_roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RemoveRole(&_PermImpl.TransactOpts, _roleId, _orgId, _caller)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(string _roleId, string _orgId, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) RemoveRole(_roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RemoveRole(&_PermImpl.TransactOpts, _roleId, _orgId, _caller)
}

// RevokeContractWhitelistByAddress is a paid mutator transaction binding the contract method 0xd43815f8.
//
// Solidity: function revokeContractWhitelistByAddress(address _contractAddress, address _caller) returns()
func (_PermImpl *PermImplTransactor) RevokeContractWhitelistByAddress(opts *bind.TransactOpts, _contractAddress common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "revokeContractWhitelistByAddress", _contractAddress, _caller)
}

// RevokeContractWhitelistByAddress is a paid mutator transaction binding the contract method 0xd43815f8.
//
// Solidity: function revokeContractWhitelistByAddress(address _contractAddress, address _caller) returns()
func (_PermImpl *PermImplSession) RevokeContractWhitelistByAddress(_contractAddress common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RevokeContractWhitelistByAddress(&_PermImpl.TransactOpts, _contractAddress, _caller)
}

// RevokeContractWhitelistByAddress is a paid mutator transaction binding the contract method 0xd43815f8.
//
// Solidity: function revokeContractWhitelistByAddress(address _contractAddress, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) RevokeContractWhitelistByAddress(_contractAddress common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RevokeContractWhitelistByAddress(&_PermImpl.TransactOpts, _contractAddress, _caller)
}

// RevokeContractWhitelistByKey is a paid mutator transaction binding the contract method 0x2a768da2.
//
// Solidity: function revokeContractWhitelistByKey(string _contractKey, address _caller) returns()
func (_PermImpl *PermImplTransactor) RevokeContractWhitelistByKey(opts *bind.TransactOpts, _contractKey string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "revokeContractWhitelistByKey", _contractKey, _caller)
}

// RevokeContractWhitelistByKey is a paid mutator transaction binding the contract method 0x2a768da2.
//
// Solidity: function revokeContractWhitelistByKey(string _contractKey, address _caller) returns()
func (_PermImpl *PermImplSession) RevokeContractWhitelistByKey(_contractKey string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RevokeContractWhitelistByKey(&_PermImpl.TransactOpts, _contractKey, _caller)
}

// RevokeContractWhitelistByKey is a paid mutator transaction binding the contract method 0x2a768da2.
//
// Solidity: function revokeContractWhitelistByKey(string _contractKey, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) RevokeContractWhitelistByKey(_contractKey string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RevokeContractWhitelistByKey(&_PermImpl.TransactOpts, _contractKey, _caller)
}

// SetAccessLevelForUnconfiguredAccount is a paid mutator transaction binding the contract method 0x68f808e5.
//
// Solidity: function setAccessLevelForUnconfiguredAccount(uint256 _accessLevel, address _caller) returns()
func (_PermImpl *PermImplTransactor) SetAccessLevelForUnconfiguredAccount(opts *bind.TransactOpts, _accessLevel *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "setAccessLevelForUnconfiguredAccount", _accessLevel, _caller)
}

// SetAccessLevelForUnconfiguredAccount is a paid mutator transaction binding the contract method 0x68f808e5.
//
// Solidity: function setAccessLevelForUnconfiguredAccount(uint256 _accessLevel, address _caller) returns()
func (_PermImpl *PermImplSession) SetAccessLevelForUnconfiguredAccount(_accessLevel *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.SetAccessLevelForUnconfiguredAccount(&_PermImpl.TransactOpts, _accessLevel, _caller)
}

// SetAccessLevelForUnconfiguredAccount is a paid mutator transaction binding the contract method 0x68f808e5.
//
// Solidity: function setAccessLevelForUnconfiguredAccount(uint256 _accessLevel, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) SetAccessLevelForUnconfiguredAccount(_accessLevel *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.SetAccessLevelForUnconfiguredAccount(&_PermImpl.TransactOpts, _accessLevel, _caller)
}

// SetIpValidation is a paid mutator transaction binding the contract method 0x9fc666b2.
//
// Solidity: function setIpValidation(bool _isIpValidationEnabled) returns()
func (_PermImpl *PermImplTransactor) SetIpValidation(opts *bind.TransactOpts, _isIpValidationEnabled bool) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "setIpValidation", _isIpValidationEnabled)
}

// SetIpValidation is a paid mutator transaction binding the contract method 0x9fc666b2.
//
// Solidity: function setIpValidation(bool _isIpValidationEnabled) returns()
func (_PermImpl *PermImplSession) SetIpValidation(_isIpValidationEnabled bool) (*types.Transaction, error) {
	return _PermImpl.Contract.SetIpValidation(&_PermImpl.TransactOpts, _isIpValidationEnabled)
}

// SetIpValidation is a paid mutator transaction binding the contract method 0x9fc666b2.
//
// Solidity: function setIpValidation(bool _isIpValidationEnabled) returns()
func (_PermImpl *PermImplTransactorSession) SetIpValidation(_isIpValidationEnabled bool) (*types.Transaction, error) {
	return _PermImpl.Contract.SetIpValidation(&_PermImpl.TransactOpts, _isIpValidationEnabled)
}

// SetMigrationPolicy is a paid mutator transaction binding the contract method 0xf5ad584a.
//
// Solidity: function setMigrationPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole, bool _networkBootStatus) returns()
func (_PermImpl *PermImplTransactor) SetMigrationPolicy(opts *bind.TransactOpts, _nwAdminOrg string, _nwAdminRole string, _oAdminRole string, _networkBootStatus bool) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "setMigrationPolicy", _nwAdminOrg, _nwAdminRole, _oAdminRole, _networkBootStatus)
}

// SetMigrationPolicy is a paid mutator transaction binding the contract method 0xf5ad584a.
//
// Solidity: function setMigrationPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole, bool _networkBootStatus) returns()
func (_PermImpl *PermImplSession) SetMigrationPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string, _networkBootStatus bool) (*types.Transaction, error) {
	return _PermImpl.Contract.SetMigrationPolicy(&_PermImpl.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole, _networkBootStatus)
}

// SetMigrationPolicy is a paid mutator transaction binding the contract method 0xf5ad584a.
//
// Solidity: function setMigrationPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole, bool _networkBootStatus) returns()
func (_PermImpl *PermImplTransactorSession) SetMigrationPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string, _networkBootStatus bool) (*types.Transaction, error) {
	return _PermImpl.Contract.SetMigrationPolicy(&_PermImpl.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole, _networkBootStatus)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermImpl *PermImplTransactor) SetPolicy(opts *bind.TransactOpts, _nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "setPolicy", _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermImpl *PermImplSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermImpl.Contract.SetPolicy(&_PermImpl.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(string _nwAdminOrg, string _nwAdminRole, string _oAdminRole) returns()
func (_PermImpl *PermImplTransactorSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermImpl.Contract.SetPolicy(&_PermImpl.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x1c249912.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactor) StartBlacklistedAccountRecovery(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "startBlacklistedAccountRecovery", _orgId, _account, _caller)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x1c249912.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplSession) StartBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.StartBlacklistedAccountRecovery(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// StartBlacklistedAccountRecovery is a paid mutator transaction binding the contract method 0x1c249912.
//
// Solidity: function startBlacklistedAccountRecovery(string _orgId, address _account, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) StartBlacklistedAccountRecovery(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.StartBlacklistedAccountRecovery(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xd621d957.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactor) StartBlacklistedNodeRecovery(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "startBlacklistedNodeRecovery", _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xd621d957.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplSession) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.StartBlacklistedNodeRecovery(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// StartBlacklistedNodeRecovery is a paid mutator transaction binding the contract method 0xd621d957.
//
// Solidity: function startBlacklistedNodeRecovery(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) StartBlacklistedNodeRecovery(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.StartBlacklistedNodeRecovery(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactor) UpdateAccountStatus(opts *bind.TransactOpts, _orgId string, _account common.Address, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateAccountStatus", _orgId, _account, _action, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateAccountStatus(&_PermImpl.TransactOpts, _orgId, _account, _action, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(string _orgId, address _account, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) UpdateAccountStatus(_orgId string, _account common.Address, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateAccountStatus(&_PermImpl.TransactOpts, _orgId, _account, _action, _caller)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermImpl *PermImplTransactor) UpdateNetworkBootStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateNetworkBootStatus")
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermImpl *PermImplSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNetworkBootStatus(&_PermImpl.TransactOpts)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermImpl *PermImplTransactorSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNetworkBootStatus(&_PermImpl.TransactOpts)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xb9b7fe6c.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateNodeStatus", _orgId, _enodeId, _ip, _port, _raftport, _action, _caller)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xb9b7fe6c.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplSession) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNodeStatus(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _action, _caller)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xb9b7fe6c.
//
// Solidity: function updateNodeStatus(string _orgId, string _enodeId, string _ip, uint16 _port, uint16 _raftport, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) UpdateNodeStatus(_orgId string, _enodeId string, _ip string, _port uint16, _raftport uint16, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNodeStatus(&_PermImpl.TransactOpts, _orgId, _enodeId, _ip, _port, _raftport, _action, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactor) UpdateOrgStatus(opts *bind.TransactOpts, _orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateOrgStatus", _orgId, _action, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplSession) UpdateOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateOrgStatus(&_PermImpl.TransactOpts, _orgId, _action, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(string _orgId, uint256 _action, address _caller) returns()
func (_PermImpl *PermImplTransactorSession) UpdateOrgStatus(_orgId string, _action *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateOrgStatus(&_PermImpl.TransactOpts, _orgId, _action, _caller)
}

// PermImplInitializedPermImplIterator is returned from FilterInitializedPermImpl and is used to iterate over the raw logs and unpacked data for InitializedPermImpl events raised by the PermImpl contract.
type PermImplInitializedPermImplIterator struct {
	Event *PermImplInitializedPermImpl // Event containing the contract specifics and raw log

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
func (it *PermImplInitializedPermImplIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermImplInitializedPermImpl)
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
		it.Event = new(PermImplInitializedPermImpl)
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
func (it *PermImplInitializedPermImplIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermImplInitializedPermImplIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermImplInitializedPermImpl represents a InitializedPermImpl event raised by the PermImpl contract.
type PermImplInitializedPermImpl struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitializedPermImpl is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PermImpl *PermImplFilterer) FilterInitializedPermImpl(opts *bind.FilterOpts) (*PermImplInitializedPermImplIterator, error) {

	logs, sub, err := _PermImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermImplInitializedPermImplIterator{contract: _PermImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

var InitializedPermImplTopicHash = "0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2"

// WatchInitializedPermImpl is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PermImpl *PermImplFilterer) WatchInitializedPermImpl(opts *bind.WatchOpts, sink chan<- *PermImplInitializedPermImpl) (event.Subscription, error) {

	logs, sub, err := _PermImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermImplInitializedPermImpl)
				if err := _PermImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitializedPermImpl is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PermImpl *PermImplFilterer) ParseInitializedPermImpl(log types.Log) (*PermImplInitializedPermImpl, error) {
	event := new(PermImplInitializedPermImpl)
	if err := _PermImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermImplPermissionsInitializedIterator is returned from FilterPermissionsInitialized and is used to iterate over the raw logs and unpacked data for PermissionsInitialized events raised by the PermImpl contract.
type PermImplPermissionsInitializedIterator struct {
	Event *PermImplPermissionsInitialized // Event containing the contract specifics and raw log

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
func (it *PermImplPermissionsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermImplPermissionsInitialized)
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
		it.Event = new(PermImplPermissionsInitialized)
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
func (it *PermImplPermissionsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermImplPermissionsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermImplPermissionsInitialized represents a PermissionsInitialized event raised by the PermImpl contract.
type PermImplPermissionsInitialized struct {
	NetworkBootStatus bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPermissionsInitialized is a free log retrieval operation binding the contract event 0x04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf.
//
// Solidity: event PermissionsInitialized(bool _networkBootStatus)
func (_PermImpl *PermImplFilterer) FilterPermissionsInitialized(opts *bind.FilterOpts) (*PermImplPermissionsInitializedIterator, error) {

	logs, sub, err := _PermImpl.contract.FilterLogs(opts, "PermissionsInitialized")
	if err != nil {
		return nil, err
	}
	return &PermImplPermissionsInitializedIterator{contract: _PermImpl.contract, event: "PermissionsInitialized", logs: logs, sub: sub}, nil
}

var PermissionsInitializedTopicHash = "0x04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf"

// WatchPermissionsInitialized is a free log subscription operation binding the contract event 0x04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf.
//
// Solidity: event PermissionsInitialized(bool _networkBootStatus)
func (_PermImpl *PermImplFilterer) WatchPermissionsInitialized(opts *bind.WatchOpts, sink chan<- *PermImplPermissionsInitialized) (event.Subscription, error) {

	logs, sub, err := _PermImpl.contract.WatchLogs(opts, "PermissionsInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermImplPermissionsInitialized)
				if err := _PermImpl.contract.UnpackLog(event, "PermissionsInitialized", log); err != nil {
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

// ParsePermissionsInitialized is a log parse operation binding the contract event 0x04f651be6fb8fc575d94591e56e9f6e66e33ef23e949765918b3bdae50c617cf.
//
// Solidity: event PermissionsInitialized(bool _networkBootStatus)
func (_PermImpl *PermImplFilterer) ParsePermissionsInitialized(log types.Log) (*PermImplPermissionsInitialized, error) {
	event := new(PermImplPermissionsInitialized)
	if err := _PermImpl.contract.UnpackLog(event, "PermissionsInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

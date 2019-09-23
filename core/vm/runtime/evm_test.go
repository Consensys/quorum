package runtime

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	testifyassert "github.com/stretchr/testify/assert"
)

/*
The following contracts are used as the samples. Bytecodes are compiled using solc 0.5.4
import "./C2.sol";
contract C1 {
    uint x;
    constructor(uint initVal) public {
        x = initVal;
    }
    function set(uint newValue) public returns (uint) {
        x = newValue;
        return x;
    }
    function get() public view returns (uint) {
        return x;
    }
    function newContractC2(uint newValue) public {
        C2 c = new C2(address(this));
        c.set(newValue);
    }
}
import "./C1.sol";
contract C2  {
   C1 c1;
   constructor(address _t) public {
       c1 = C1(_t);
   }
   function get() public view returns (uint result) {
       return c1.get();
   }
   function set(uint _val) public {
       c1.set(_val);
   }
}
------------------
pragma solidity ^0.5.4;

contract D {
  uint public n;
  address public sender;

  function callSetN(address _e, uint _n) public {
    _e.call(abi.encodeWithSignature("setN(uint256)", _n)); // E's storage is set, D is not modified
  }

  function delegatecallSetN(address _e, uint _n) public {
    _e.delegatecall(abi.encodeWithSignature("setN(uint256)", _n)); // D's storage is set, E is not modified
  }
}

pragma solidity ^0.5.4;

contract E {
  uint public n;
  address public sender;

  function setN(uint _n) public {
    n = _n;
    sender = msg.sender;
    // msg.sender is D if invoked by D's delegatecallSetN. None of E's storage is updated
    // msg.sender is C if invoked by C.foo(). None of E's storage is updated

    // the value of "this" is D, when invoked by either D's callcodeSetN or C.foo()
  }
}
*/

type contract struct {
	abi      abi.ABI
	bytecode []byte
	name     string
}

var (
	c1, c2, d, e *contract
)

func init() {
	c1 = &contract{
		name:     "c1",
		abi:      mustParse(c1AbiDefinition),
		bytecode: common.Hex2Bytes("608060405234801561001057600080fd5b506040516020806105a88339810180604052602081101561003057600080fd5b81019080805190602001909291905050508060008190555050610550806100586000396000f3fe608060405260043610610051576000357c01000000000000000000000000000000000000000000000000000000009004806360fe47b1146100565780636d4ce63c146100a5578063d7139463146100d0575b600080fd5b34801561006257600080fd5b5061008f6004803603602081101561007957600080fd5b810190808035906020019092919050505061010b565b6040518082815260200191505060405180910390f35b3480156100b157600080fd5b506100ba61011e565b6040518082815260200191505060405180910390f35b3480156100dc57600080fd5b50610109600480360360208110156100f357600080fd5b8101908080359060200190929190505050610127565b005b6000816000819055506000549050919050565b60008054905090565b600030610132610212565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610184573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff166360fe47b1836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b1580156101f657600080fd5b505af115801561020a573d6000803e3d6000fd5b505050505050565b604051610302806102238339019056fe608060405234801561001057600080fd5b506040516020806103028339810180604052602081101561003057600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610271806100916000396000f3fe608060405260043610610046576000357c01000000000000000000000000000000000000000000000000000000009004806360fe47b11461004b5780636d4ce63c14610086575b600080fd5b34801561005757600080fd5b506100846004803603602081101561006e57600080fd5b81019080803590602001909291905050506100b1565b005b34801561009257600080fd5b5061009b610180565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360fe47b1826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561014157600080fd5b505af1158015610155573d6000803e3d6000fd5b505050506040513d602081101561016b57600080fd5b81019080805190602001909291905050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d4ce63c6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561020557600080fd5b505afa158015610219573d6000803e3d6000fd5b505050506040513d602081101561022f57600080fd5b810190808051906020019092919050505090509056fea165627a7a72305820a537f4c360ce5c6f55523298e314e6456e5c3e02c170563751dfda37d3aeddb30029a165627a7a7230582060396bfff29d2dfc5a9f4216bfba5e24d031d54fd4b26ebebde1a26c59df0c1e0029"),
	}
	c2 = &contract{
		name:     "c2",
		abi:      mustParse(c2AbiDefinition),
		bytecode: common.Hex2Bytes("608060405234801561001057600080fd5b506040516020806102f58339810180604052602081101561003057600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610264806100916000396000f3fe608060405234801561001057600080fd5b5060043610610053576000357c01000000000000000000000000000000000000000000000000000000009004806360fe47b1146100585780636d4ce63c14610086575b600080fd5b6100846004803603602081101561006e57600080fd5b81019080803590602001909291905050506100a4565b005b61008e610173565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360fe47b1826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561013457600080fd5b505af1158015610148573d6000803e3d6000fd5b505050506040513d602081101561015e57600080fd5b81019080805190602001909291905050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d4ce63c6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156101f857600080fd5b505afa15801561020c573d6000803e3d6000fd5b505050506040513d602081101561022257600080fd5b810190808051906020019092919050505090509056fea165627a7a72305820dd8a5dcf693e1969289c444a282d0684a9760bac26f1e4e0139d46821ec1979b0029"),
	}
	d = &contract{
		name:     "d",
		abi:      mustParse(dAbiDefinition),
		bytecode: common.Hex2Bytes("608060405234801561001057600080fd5b50610458806100206000396000f3fe608060405234801561001057600080fd5b5060043610610069576000357c0100000000000000000000000000000000000000000000000000000000900480632e52d6061461006e57806367e404ce1461008c5780639b58bc26146100d6578063eea4c86414610124575b600080fd5b610076610172565b6040518082815260200191505060405180910390f35b610094610178565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610122600480360360408110156100ec57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061019e565b005b6101706004803603604081101561013a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506102e4565b005b60005481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8173ffffffffffffffffffffffffffffffffffffffff1681604051602401808281526020019150506040516020818303038152906040527f3f7a0270000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b6020831015156102785780518252602082019150602081019050602083039250610253565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146102d8576040519150601f19603f3d011682016040523d82523d6000602084013e6102dd565b606091505b5050505050565b8173ffffffffffffffffffffffffffffffffffffffff1681604051602401808281526020019150506040516020818303038152906040527f3f7a0270000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b6020831015156103be5780518252602082019150602081019050602083039250610399565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610420576040519150601f19603f3d011682016040523d82523d6000602084013e610425565b606091505b505050505056fea165627a7a72305820903354d4ff73c2763f7ea1812e9b45970e607706b9acdc7dd2b44a46c6726a030029"),
	}
	e = &contract{
		name:     "e",
		abi:      mustParse(cAbiDefinition),
		bytecode: common.Hex2Bytes("608060405234801561001057600080fd5b5061019c806100206000396000f3fe608060405234801561001057600080fd5b506004361061005e576000357c0100000000000000000000000000000000000000000000000000000000900480632e52d606146100635780633f7a02701461008157806367e404ce146100af575b600080fd5b61006b6100f9565b6040518082815260200191505060405180910390f35b6100ad6004803603602081101561009757600080fd5b81019080803590602001909291905050506100ff565b005b6100b761014a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60005481565b8060008190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea165627a7a72305820b38efe0156613953609c12b74f7847429cffec7bf7c4d7b4dd33be5c757680680029"),
	}
	log.PrintOrigins(true)
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
}

func TestEVM_AffectedContracts_whenCreatingC1AndCallingC1Get(t *testing.T) {
	assert := testifyassert.New(t)
	cfg := newConfig()
	initialValue := int64(42)
	var calledContracts []common.Address
	var createdContracts []common.Address
	var affectedModeFunc func(common.Address) (vm.AffectedMode, error)
	cfg.onAfterEVM = func(evm *vm.EVM) {
		calledContracts = evm.CalledContracts()
		createdContracts = evm.CreatedContracts()
		affectedModeFunc = evm.AffectedMode
	}

	c1Address := createC1(assert, cfg, initialValue)

	assert.Empty(calledContracts, "Contract C1 creation doesn't affect any other contract")
	assert.Len(createdContracts, 1, "Contract C1 creation must create C1 contract itself")
	c1Mode, err := affectedModeFunc(c1Address)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeWrite|vm.ModeRead), c1Mode)

	actualValue := callContractFunction(assert, cfg, c1, c1Address, "get")

	assert.Equal(initialValue, actualValue)
	assert.Len(calledContracts, 1, "Calling C1.get() affects 1 contract")
	assert.Equal(c1Address, calledContracts[0], "Calling C1.get() affects C1 contract itself")
	c1Mode, err = affectedModeFunc(c1Address)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead), c1Mode)
}

func TestEVM_AffectedContracts_whenCreatingC2AndCallingC2Get(t *testing.T) {
	assert := testifyassert.New(t)
	cfg := newConfig()
	initialValue := int64(30)

	c1Address := createC1(assert, cfg, initialValue)

	var calledContracts []common.Address
	var createdContracts []common.Address
	var affectedModeFunc func(common.Address) (vm.AffectedMode, error)
	cfg.onAfterEVM = func(evm *vm.EVM) {
		calledContracts = evm.CalledContracts()
		createdContracts = evm.CreatedContracts()
		affectedModeFunc = evm.AffectedMode
	}

	c2Address := createC2(assert, cfg, c1Address)

	assert.Empty(calledContracts, "Contract C2 creation doesn't affect any other contract")
	assert.Len(createdContracts, 1, "Contract C2 creation must create C2 contract itself")
	c2Mode, err := affectedModeFunc(c2Address)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeWrite|vm.ModeRead), c2Mode)

	actualValue := callContractFunction(assert, cfg, c2, c2Address, "get")

	assert.Equal(initialValue, actualValue)
	assert.Len(calledContracts, 2, "Calling C2.get() affects 2 contracts")
	assert.Contains(calledContracts, c1Address, "Calling C2.get() affects C1")
	assert.Contains(calledContracts, c2Address, "Calling C2.get() affects C2")
	c2Mode, err = affectedModeFunc(c2Address)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead), c2Mode)
	c1Mode, err := affectedModeFunc(c1Address)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead), c1Mode)
}

func TestEVM_AffectedContracts_whenCallingC2SetFunction(t *testing.T) {
	assert := testifyassert.New(t)
	cfg := newConfig()
	initialValue := int64(30)
	newValue := int64(42)

	c1Address := createC1(assert, cfg, initialValue)

	var calledContracts []common.Address
	var createdContracts []common.Address
	var affectedModeFunc func(common.Address) (vm.AffectedMode, error)
	cfg.onAfterEVM = func(evm *vm.EVM) {
		calledContracts = evm.CalledContracts()
		createdContracts = evm.CreatedContracts()
		affectedModeFunc = evm.AffectedMode
	}
	c2Address := createC2(assert, cfg, c1Address)
	assert.Empty(calledContracts, "Contract C2 creation doesn't affect any other contract")
	assert.Len(createdContracts, 1, "Contract C2 creation must create C2 contract itself")

	callContractFunction(assert, cfg, c2, c2Address, "set", big.NewInt(newValue))

	assert.Len(calledContracts, 2, "Calling C2.set() affects 2 contracts")
	assert.Contains(calledContracts, c1Address, "Calling C2.set() affects C1")
	assert.Contains(calledContracts, c2Address, "Calling C2.set() affects C2")
	c2Mode, err := affectedModeFunc(c2Address)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead), c2Mode)
	c1Mode, err := affectedModeFunc(c1Address)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead|vm.ModeWrite), c1Mode)
}

func TestEVM_AffectedContracts_whenCreatingC2ByCallingC1Function(t *testing.T) {
	assert := testifyassert.New(t)
	cfg := newConfig()
	initialValue := int64(30)
	newValue := int64(40)

	c1Address := createC1(assert, cfg, initialValue)

	var calledContracts []common.Address
	var createdContracts []common.Address
	cfg.onAfterEVM = func(evm *vm.EVM) {
		calledContracts = evm.CalledContracts()
		createdContracts = evm.CreatedContracts()
	}
	callContractFunction(assert, cfg, c1, c1Address, "newContractC2", big.NewInt(newValue))

	assert.Len(calledContracts, 1, "Calling C1.newContractC2() affects 1 contract")
	assert.Contains(calledContracts, c1Address, "Calling C1.newContractC2() affects C1")
	assert.Len(createdContracts, 1, "Calling C1.newContractC2() creates 1 contract which is C2")
}

func TestEVM_AffectedContracts_whenUsingCall(t *testing.T) {
	assert := testifyassert.New(t)
	cfg := newConfig()

	dAddress := createD(assert, cfg)
	eAddress := createE(assert, cfg)

	var calledContracts []common.Address
	var affectedModeFunc func(common.Address) (vm.AffectedMode, error)
	cfg.onAfterEVM = func(evm *vm.EVM) {
		calledContracts = evm.CalledContracts()
		affectedModeFunc = evm.AffectedMode
	}

	callContractFunction(assert, cfg, d, dAddress, "callSetN", eAddress, big.NewInt(20))

	assert.Len(calledContracts, 2, "Calling D.callSetN() affects 2 contracts")
	assert.Contains(calledContracts, dAddress, "Calling D.callSetN() affects D")
	assert.Contains(calledContracts, eAddress, "Calling D.callSetN() affects E")
	dMode, err := affectedModeFunc(dAddress)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead), dMode)
	eMode, err := affectedModeFunc(eAddress)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead|vm.ModeWrite), eMode)
}

func TestEVM_AffectedContracts_whenUsingDelegateCall(t *testing.T) {
	assert := testifyassert.New(t)
	cfg := newConfig()

	dAddress := createD(assert, cfg)
	eAddress := createE(assert, cfg)

	var calledContracts []common.Address
	var affectedModeFunc func(common.Address) (vm.AffectedMode, error)
	cfg.onAfterEVM = func(evm *vm.EVM) {
		calledContracts = evm.CalledContracts()
		affectedModeFunc = evm.AffectedMode
	}

	callContractFunction(assert, cfg, d, dAddress, "delegatecallSetN", eAddress, big.NewInt(20))

	assert.Len(calledContracts, 2, "Calling D.delegatecallSetN() affects 2 contracts")
	assert.Contains(calledContracts, dAddress, "Calling D.delegatecallSetN() affects D")
	assert.Contains(calledContracts, eAddress, "Calling D.delegatecallSetN() affects E")
	dMode, err := affectedModeFunc(dAddress)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead|vm.ModeWrite), dMode)
	eMode, err := affectedModeFunc(eAddress)
	assert.NoError(err)
	assert.Equal(vm.AffectedMode(vm.ModeRead), eMode)
}

func callContractFunction(assert *testifyassert.Assertions, cfg *extendedConfig, c *contract, address common.Address, name string, args ...interface{}) int64 {
	sig := fmt.Sprintf("%s.%s", c.name, name)
	log.Debug("running", "method", sig)
	f := mustPack(assert, c, name, args...)
	ret, _, err := call(address, f, cfg)
	assert.NoError(err, "Execute %s", sig)
	log.Debug(sig, "ret_hex", common.Bytes2Hex(ret))
	for len(ret) > 0 && ret[0] == 0 {
		ret = ret[1:]
	}
	if len(ret) == 0 {
		return -1
	}
	actualValue, err := hexutil.DecodeBig(hexutil.Encode(ret))
	assert.NoError(err)
	log.Debug(sig, "ret", actualValue)
	return actualValue.Int64()
}

func createD(assert *testifyassert.Assertions, cfg *extendedConfig) common.Address {
	constructorCode := mustPack(assert, d, "")

	_, address, _, err := create(append(d.bytecode, constructorCode...), cfg)
	assert.NoError(err, "Create contract D")

	log.Debug("Created D", "address", address.Hex())
	return address
}

func createE(assert *testifyassert.Assertions, cfg *extendedConfig) common.Address {
	constructorCode := mustPack(assert, e, "")

	_, address, _, err := create(append(e.bytecode, constructorCode...), cfg)
	assert.NoError(err, "Create contract E")

	log.Debug("Created E", "address", address.Hex())
	return address
}

func createC2(assert *testifyassert.Assertions, cfg *extendedConfig, c1Address common.Address) common.Address {
	constructorCode := mustPack(assert, c2, "", c1Address)

	_, address, _, err := create(append(c2.bytecode, constructorCode...), cfg)
	assert.NoError(err, "Create contract C2")

	log.Debug("Created C2", "address", address.Hex())
	return address
}

func createC1(assert *testifyassert.Assertions, cfg *extendedConfig, initialValue int64) common.Address {
	constructorCode := mustPack(assert, c1, "", big.NewInt(initialValue))

	_, address, _, err := create(append(c1.bytecode, constructorCode...), cfg)
	assert.NoError(err, "Create contract C1")

	log.Debug("Created C1", "address", address.Hex())
	return address
}

func mustPack(assert *testifyassert.Assertions, c *contract, name string, args ...interface{}) []byte {
	bytes, err := c.abi.Pack(name, args...)
	assert.NoError(err, "Pack method")
	return bytes
}

func newConfig() *extendedConfig {
	cfg := new(Config)
	setDefaults(cfg)
	cfg.Debug = true
	database := rawdb.NewMemoryDatabase()
	cfg.State, _ = state.New(common.Hash{}, state.NewDatabase(database))
	privateState, _ := state.New(common.Hash{}, state.NewDatabase(database))

	cfg.ChainConfig.IsQuorum = true
	cfg.ChainConfig.ByzantiumBlock = big.NewInt(0)
	return &extendedConfig{
		Config:       cfg,
		privateState: privateState,
	}
}

type extendedConfig struct {
	*Config
	privateState *state.StateDB
	onAfterEVM   func(evm *vm.EVM)
}

func newEVM(cfg *extendedConfig) *vm.EVM {
	context := vm.Context{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },

		Origin:      cfg.Origin,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		Difficulty:  cfg.Difficulty,
		GasLimit:    cfg.GasLimit,
		GasPrice:    cfg.GasPrice,
	}
	evm := vm.NewEVM(context, cfg.State, cfg.privateState, cfg.ChainConfig, cfg.EVMConfig)
	return evm
}

// Create executes the code using the EVM create method
func create(input []byte, cfg *extendedConfig) ([]byte, common.Address, uint64, error) {
	var (
		vmenv  = newEVM(cfg)
		sender = vm.AccountRef(cfg.Origin)
	)
	defer func() {
		if cfg.onAfterEVM != nil {
			cfg.onAfterEVM(vmenv)
		}
	}()

	// Call the code with the given configuration.
	code, address, leftOverGas, err := vmenv.Create(
		sender,
		input,
		cfg.GasLimit,
		cfg.Value,
	)
	return code, address, leftOverGas, err
}

// Call executes the code given by the contract's address. It will return the
// EVM's return value or an error if it failed.
//
// Call, unlike Execute, requires a config and also requires the State field to
// be set.
func call(address common.Address, input []byte, cfg *extendedConfig) ([]byte, uint64, error) {
	vmenv := newEVM(cfg)
	defer func() {
		if cfg.onAfterEVM != nil {
			cfg.onAfterEVM(vmenv)
		}
	}()

	sender := cfg.State.GetOrNewStateObject(cfg.Origin)
	// Call the code with the given configuration.
	ret, leftOverGas, err := vmenv.Call(
		sender,
		address,
		input,
		cfg.GasLimit,
		cfg.Value,
	)

	return ret, leftOverGas, err
}

func mustParse(def string) abi.ABI {
	abi, err := abi.JSON(strings.NewReader(def))
	if err != nil {
		log.Error("Can't parse ABI def", "err", err)
		os.Exit(1)
	}
	return abi
}

const (
	c1AbiDefinition = `
[
	{
		"constant": false,
		"inputs": [
			{
				"name": "newValue",
				"type": "uint256"
			}
		],
		"name": "set",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "get",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "newValue",
				"type": "uint256"
			}
		],
		"name": "newContractC2",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "initVal",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	}
]
`
	c2AbiDefinition = `
[
	{
		"constant": false,
		"inputs": [
			{
				"name": "_val",
				"type": "uint256"
			}
		],
		"name": "set",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "get",
		"outputs": [
			{
				"name": "result",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "_t",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	}
]
`
	dAbiDefinition = `
[
	{
		"constant": true,
		"inputs": [],
		"name": "n",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "sender",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_e",
				"type": "address"
			},
			{
				"name": "_n",
				"type": "uint256"
			}
		],
		"name": "delegatecallSetN",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_e",
				"type": "address"
			},
			{
				"name": "_n",
				"type": "uint256"
			}
		],
		"name": "callSetN",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`
	cAbiDefinition = `
[
	{
		"constant": true,
		"inputs": [],
		"name": "n",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_n",
				"type": "uint256"
			}
		],
		"name": "setN",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "sender",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]
`
)

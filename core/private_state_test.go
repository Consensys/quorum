package core

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// callmsg is the message type used for call transactions in the private state test
type callmsg struct {
	addr     common.Address
	to       *common.Address
	gas      uint64
	gasPrice *big.Int
	value    *big.Int
	data     []byte
}

// accessor boilerplate to implement core.Message
func (m callmsg) From() common.Address         { return m.addr }
func (m callmsg) FromFrontier() common.Address { return m.addr }
func (m callmsg) Nonce() uint64                { return 0 }
func (m callmsg) To() *common.Address          { return m.to }
func (m callmsg) GasPrice() *big.Int           { return m.gasPrice }
func (m callmsg) Gas() uint64                  { return m.gas }
func (m callmsg) Value() *big.Int              { return m.value }
func (m callmsg) Data() []byte                 { return m.data }
func (m callmsg) CheckNonce() bool             { return true }

func ExampleMakeCallHelper() {
	var (
		// setup new pair of keys for the calls
		key, _ = crypto.GenerateKey()
		// create a new helper
		helper = MakeCallHelper()
	)
	// Private contract address
	prvContractAddr := common.Address{1}
	// Initialise custom code for private contract
	helper.PrivateState.SetCode(prvContractAddr, common.Hex2Bytes("600a60005500"))
	// Public contract address
	pubContractAddr := common.Address{2}
	// Initialise custom code for public contract
	helper.PublicState.SetCode(pubContractAddr, common.Hex2Bytes("601460005500"))

	// Make a call to the private contract
	err := helper.MakeCall(true, key, prvContractAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
	// Make a call to the public contract
	err = helper.MakeCall(false, key, pubContractAddr, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Private: 10
	// Public: 20
	fmt.Println("Private:", helper.PrivateState.GetState(prvContractAddr, common.Hash{}).Big())
	fmt.Println("Public:", helper.PublicState.GetState(pubContractAddr, common.Hash{}).Big())
}

// 600a600055600060006001a1
// 60 0a, 60 00, 55,  60 00, 60 00, 60 01,  a1
// [1] (0x60) PUSH1 0x0a (store value)
// [3] (0x60) PUSH1 0x00 (store addr)
// [4] (0x55) SSTORE  (Store (k-00,v-a))

// [6] (0x60) PUSH1 0x00
// [8] (0x60) PUSH1 0x00
// [10](0x60) PUSH1 0x01
// [11](0xa1) LOG1 offset(0x01), len(0x00), topic(0x00)
//
// Store then log
func TestPrivateTransaction(t *testing.T) {
	var (
		key, _       = crypto.GenerateKey()
		helper       = MakeCallHelper()
		privateState = helper.PrivateState
		publicState  = helper.PublicState
	)

	prvContractAddr := common.Address{1}
	pubContractAddr := common.Address{2}
	// SSTORE (K,V) SSTORE(0, 10): 600a600055
	// +
	// LOG1 OFFSET LEN TOPIC,  LOG1 (a1) 01, 00, 00: 600060006001a1
	privateState.SetCode(prvContractAddr, common.Hex2Bytes("600a600055600060006001a1"))
	// SSTORE (K,V) SSTORE(0, 14): 6014600055
	publicState.SetCode(pubContractAddr, common.Hex2Bytes("6014600055"))

	if publicState.Exist(prvContractAddr) {
		t.Error("didn't expect private contract address to exist on public state")
	}

	// Private transaction 1
	err := helper.MakeCall(true, key, prvContractAddr, nil)

	if err != nil {
		t.Fatal(err)
	}
	stateEntry := privateState.GetState(prvContractAddr, common.Hash{}).Big()
	if stateEntry.Cmp(big.NewInt(10)) != 0 {
		t.Error("expected state to have 10, got", stateEntry)
	}
	if len(privateState.Logs()) != 1 {
		t.Error("expected private state to have 1 log, got", len(privateState.Logs()))
	}
	if len(publicState.Logs()) != 0 {
		t.Error("expected public state to have 0 logs, got", len(publicState.Logs()))
	}
	if publicState.Exist(prvContractAddr) {
		t.Error("didn't expect private contract address to exist on public state")
	}
	if !privateState.Exist(prvContractAddr) {
		t.Error("expected private contract address to exist on private state")
	}

	// Public transaction 1
	err = helper.MakeCall(false, key, pubContractAddr, nil)
	if err != nil {
		t.Fatal(err)
	}
	stateEntry = publicState.GetState(pubContractAddr, common.Hash{}).Big()
	if stateEntry.Cmp(big.NewInt(20)) != 0 {
		t.Error("expected state to have 20, got", stateEntry)
	}

	// Private transaction 2
	err = helper.MakeCall(true, key, prvContractAddr, nil)
	stateEntry = privateState.GetState(prvContractAddr, common.Hash{}).Big()
	if stateEntry.Cmp(big.NewInt(10)) != 0 {
		t.Error("expected state to have 10, got", stateEntry)
	}

	if publicState.Exist(prvContractAddr) {
		t.Error("didn't expect private contract address to exist on public state")
	}
	if privateState.Exist(pubContractAddr) {
		t.Error("didn't expect public contract address to exist on private state")
	}
}

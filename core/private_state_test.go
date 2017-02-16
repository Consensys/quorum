package core

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

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

func TestPrivateTransaction(t *testing.T) {
	var (
		key, _       = crypto.GenerateKey()
		helper       = MakeCallHelper()
		privateState = helper.PrivateState
		publicState  = helper.PublicState
	)

	prvContractAddr := common.Address{1}
	pubContractAddr := common.Address{2}
	privateState.SetCode(prvContractAddr, common.Hex2Bytes("600a600055600060006001a1"))
	privateState.SetState(prvContractAddr, common.Hash{}, common.Hash{9})
	publicState.SetCode(pubContractAddr, common.Hex2Bytes("6014600055"))
	publicState.SetState(pubContractAddr, common.Hash{}, common.Hash{19})

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

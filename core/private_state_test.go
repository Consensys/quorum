package core

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/big"
	"os"
	osExec "os/exec"
	"path"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/constellation"
)

// callmsg is the message type used for call transactions in the private state test
type callmsg struct {
	addr          common.Address
	to            *common.Address
	gas, gasPrice *big.Int
	value         *big.Int
	data          []byte
}

// accessor boilerplate to implement core.Message
func (m callmsg) From() common.Address         { return m.addr }
func (m callmsg) FromFrontier() common.Address { return m.addr }
func (m callmsg) Nonce() uint64                { return 0 }
func (m callmsg) To() *common.Address          { return m.to }
func (m callmsg) GasPrice() *big.Int           { return m.gasPrice }
func (m callmsg) Gas() *big.Int                { return m.gas }
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

var constellationCfgTemplate = template.Must(template.New("t").Parse(`
	url = "http://127.0.0.1:9000/"
	port = 9000
	socketPath = "{{.RootDir}}/qdata/tm1.ipc"
	otherNodeUrls = []
	publicKeyPath = "{{.RootDir}}/keys/tm1.pub"
	privateKeyPath = "{{.RootDir}}/keys/tm1.key"
	archivalPublicKeyPath = "{{.RootDir}}/keys/tm1a.pub"
	archivalPrivateKeyPath = "{{.RootDir}}/keys/tm1a.key"
	storagePath = "{{.RootDir}}/qdata/constellation1"
`))

func runConstellation() (*osExec.Cmd, error) {
	dir, err := ioutil.TempDir("", "TestPrivateTxConstellationData")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	here, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if err = os.MkdirAll(path.Join(dir, "qdata"), 0755); err != nil {
		return nil, err
	}
	if err = os.Symlink(path.Join(here, "constellation-test-keys"), path.Join(dir, "keys")); err != nil {
		return nil, err
	}
	cfgFile, err := os.Create(path.Join(dir, "constellation.cfg"))
	if err != nil {
		return nil, err
	}
	err = constellationCfgTemplate.Execute(cfgFile, map[string]string{"RootDir": dir})
	if err != nil {
		return nil, err
	}
	constellationCmd := osExec.Command("constellation-node", cfgFile.Name())
	var stdout, stderr bytes.Buffer
	constellationCmd.Stdout = &stdout
	constellationCmd.Stderr = &stderr
	var constellationErr error
	go func() {
		constellationErr = constellationCmd.Start()
	}()
	// Give the constellation subprocess some time to start.
	time.Sleep(1 * time.Second)
	if constellationErr != nil {
		fmt.Println(stdout.String() + stderr.String())
		return nil, constellationErr
	}
	private.P = constellation.MustNew(cfgFile.Name())
	return constellationCmd, nil
}

// 600a600055600060006001a1
// [1] PUSH1 0x0a (store value)
// [3] PUSH1 0x00 (store addr)
// [4] SSTORE
// [6] PUSH1 0x00
// [8] PUSH1 0x00
// [10] PUSH1 0x01
// [11] LOG1
//
// Store then log
func TestPrivateTransaction(t *testing.T) {
	var (
		key, _       = crypto.GenerateKey()
		helper       = MakeCallHelper()
		privateState = helper.PrivateState
		publicState  = helper.PublicState
	)

	constellationCmd, err := runConstellation()
	if err != nil {
		t.Fatal(err)
	}
	defer constellationCmd.Process.Kill()

	prvContractAddr := common.Address{1}
	pubContractAddr := common.Address{2}
	privateState.SetCode(prvContractAddr, common.Hex2Bytes("600a600055600060006001a1"))
	privateState.SetState(prvContractAddr, common.Hash{}, common.Hash{9})
	publicState.SetCode(pubContractAddr, common.Hex2Bytes("6014600055"))
	publicState.SetState(pubContractAddr, common.Hash{}, common.Hash{19})

	if publicState.Exist(prvContractAddr) {
		t.Error("didn't expect private contract address to exist on public state")
	}

	// Private transaction 1
	err = helper.MakeCall(true, key, prvContractAddr, nil)
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

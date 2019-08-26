package permission

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/p2p"

	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/eth"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/node"
	pbind "github.com/ethereum/go-ethereum/permission/bind"
)

func TestPermissionCtrl_AfterStart(t *testing.T) {
	guardianKey, _ := crypto.GenerateKey()
	nodeKey, _ := crypto.GenerateKey()

	guardianAddress := crypto.PubkeyToAddress(guardianKey.PublicKey)

	guardianTransactor := bind.NewKeyedTransactor(guardianKey)
	genesisAlloc := map[common.Address]core.GenesisAccount{
		guardianAddress: {
			Balance: big.NewInt(100000000000000),
		},
	}
	// Create a networkless protocol stack and start an Ethereum service within
	stack, err := node.New(&node.Config{
		DataDir:           "",
		UseLightweightKDF: true,
		P2P: p2p.Config{
			PrivateKey: nodeKey,
		},
	})
	if err != nil {
		t.Fatalf("failed to create node: %v", err)
	}
	ethConf := &eth.Config{
		Genesis:   &core.Genesis{Config: params.AllEthashProtocolChanges, GasLimit: 10000000000, Alloc: genesisAlloc},
		Etherbase: guardianAddress,
		Ethash: ethash.Config{
			PowMode: ethash.ModeTest,
		},
	}
	if err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) { return eth.New(ctx, ethConf) }); err != nil {
		t.Fatalf("failed to register Ethereum protocol: %v", err)
	}
	// Start the node and assemble the JavaScript console around it
	if err = stack.Start(); err != nil {
		t.Fatalf("failed to start test stack: %v", err)
	}
	var ethereum *eth.Ethereum
	if err := stack.Service(&ethereum); err != nil {
		t.Fatal(err)
	}
	sb := backends.NewSimulatedBackendFrom(ethereum)

	permUpgrAddress, _, permUpgrInstance, err := pbind.DeployPermUpgr(guardianTransactor, sb, guardianAddress)
	if err != nil {
		t.Fatal(err)
	}
	permInterfaceAddress, _, _, err := pbind.DeployPermInterface(guardianTransactor, sb, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	nodeManagerAddress, _, _, err := pbind.DeployNodeManager(guardianTransactor, sb, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	roleManagerAddress, _, _, err := pbind.DeployRoleManager(guardianTransactor, sb, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	accountManagerAddress, _, _, err := pbind.DeployAcctManager(guardianTransactor, sb, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	orgManagerAddress, _, _, err := pbind.DeployOrgManager(guardianTransactor, sb, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	voterManagerAddress, _, _, err := pbind.DeployVoterManager(guardianTransactor, sb, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	permImplAddress, _, _, err := pbind.DeployPermImpl(guardianTransactor, sb, permUpgrAddress, orgManagerAddress, roleManagerAddress, accountManagerAddress, voterManagerAddress, nodeManagerAddress)
	if err != nil {
		t.Fatal(err)
	}
	// call init
	if _, err := permUpgrInstance.Init(guardianTransactor, permInterfaceAddress, permImplAddress); err != nil {
		t.Fatal(err)
	}
	testObject, err := NewQuorumPermissionCtrl(stack, &types.PermissionConfig{
		UpgrdAddress:   permUpgrAddress,
		InterfAddress:  permInterfaceAddress,
		ImplAddress:    permImplAddress,
		NodeAddress:    nodeManagerAddress,
		AccountAddress: accountManagerAddress,
		RoleAddress:    roleManagerAddress,
		VoterAddress:   voterManagerAddress,
		OrgAddress:     orgManagerAddress,
		NwAdminOrg:     "NETWORK_ADMIN",
		NwAdminRole:    "NETWORK_ADMIN_ROLE",
		OrgAdminRole:   "ORG_ADMIN_ROLE",
		Accounts: []common.Address{
			guardianAddress,
		},
		SubOrgDepth:   *big.NewInt(3),
		SubOrgBreadth: *big.NewInt(3),
	})
	if err != nil {
		t.Fatal(err)
	}
	testObject.ethClnt = sb
	testObject.eth = ethereum
	go func() {
		testObject.errorChan <- nil
	}()
	fmt.Println("after start")
	err = testObject.AfterStart()

	assert.NoError(t, err)
}

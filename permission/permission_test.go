package permission

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/p2p"

	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/eth"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/node"
	pbind "github.com/ethereum/go-ethereum/permission/bind"
)

var (
	testObject  *PermissionCtrl
	guardianKey *ecdsa.PrivateKey
	backend     bind.ContractBackend
	permUpgrAddress, permInterfaceAddress, permImplAddress, voterManagerAddress,
	nodeManagerAddress, roleManagerAddress, accountManagerAddress, orgManagerAddress common.Address
	ethereum        *eth.Ethereum
	stack           *node.Node
	guardianAddress common.Address
)

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() {
	t := log.New(os.Stdout, "", 0)
	guardianKey, _ = crypto.GenerateKey()
	nodeKey, _ := crypto.GenerateKey()

	guardianAddress = crypto.PubkeyToAddress(guardianKey.PublicKey)

	guardianTransactor := bind.NewKeyedTransactor(guardianKey)
	genesisAlloc := map[common.Address]core.GenesisAccount{
		guardianAddress: {
			Balance: big.NewInt(100000000000000),
		},
	}
	var err error
	// Create a networkless protocol stack and start an Ethereum service within
	stack, err = node.New(&node.Config{
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
	if err := stack.Service(&ethereum); err != nil {
		t.Fatal(err)
	}
	backend = backends.NewSimulatedBackendFrom(ethereum)

	var permUpgrInstance *pbind.PermUpgr
	permUpgrAddress, _, permUpgrInstance, err = pbind.DeployPermUpgr(guardianTransactor, backend, guardianAddress)
	if err != nil {
		t.Fatal(err)
	}
	permInterfaceAddress, _, _, err = pbind.DeployPermInterface(guardianTransactor, backend, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	nodeManagerAddress, _, _, err = pbind.DeployNodeManager(guardianTransactor, backend, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	roleManagerAddress, _, _, err = pbind.DeployRoleManager(guardianTransactor, backend, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	accountManagerAddress, _, _, err = pbind.DeployAcctManager(guardianTransactor, backend, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	orgManagerAddress, _, _, err = pbind.DeployOrgManager(guardianTransactor, backend, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	voterManagerAddress, _, _, err = pbind.DeployVoterManager(guardianTransactor, backend, permUpgrAddress)
	if err != nil {
		t.Fatal(err)
	}
	permImplAddress, _, _, err = pbind.DeployPermImpl(guardianTransactor, backend, permUpgrAddress, orgManagerAddress, roleManagerAddress, accountManagerAddress, voterManagerAddress, nodeManagerAddress)
	if err != nil {
		t.Fatal(err)
	}
	// call init
	if _, err := permUpgrInstance.Init(guardianTransactor, permInterfaceAddress, permImplAddress); err != nil {
		t.Fatal(err)
	}
}

func teardown() {

}

func TestPermissionCtrl_AfterStart(t *testing.T) {
	testObject := typicalPermissionCtrl(t)

	err := testObject.AfterStart()

	assert.NoError(t, err)
	assert.NotNil(t, testObject.permOrg)
	assert.NotNil(t, testObject.permRole)
	assert.NotNil(t, testObject.permNode)
	assert.NotNil(t, testObject.permAcct)
	assert.NotNil(t, testObject.permInterf)
	assert.NotNil(t, testObject.permUpgr)

	networkBooted, err := testObject.permInterf.GetNetworkBootStatus(&bind.CallOpts{
		Pending: true,
	})
	assert.NoError(t, err)
	assert.True(t, networkBooted)
}

func TestPermissionCtrl_PopulateInitPermissions_whenNetworkIsInitialized(t *testing.T) {
	testObject := typicalPermissionCtrl(t)
	assert.NoError(t, testObject.AfterStart())

	err := testObject.populateInitPermissions()

	assert.NoError(t, err)

	// assert cache
	assert.Equal(t, 1, len(types.OrgInfoMap.GetOrgList()))
	cachedOrg := types.OrgInfoMap.GetOrgList()[0]
	assert.Equal(t, "NETWORK_ADMIN", cachedOrg.OrgId)
	assert.Equal(t, "NETWORK_ADMIN", cachedOrg.FullOrgId)
	assert.Equal(t, "NETWORK_ADMIN", cachedOrg.UltimateParent)
	assert.Equal(t, "", cachedOrg.ParentOrgId)
	assert.Equal(t, types.OrgApproved, cachedOrg.Status)
	assert.Equal(t, 0, len(cachedOrg.SubOrgList))
	assert.Equal(t, big.NewInt(1), cachedOrg.Level)

	assert.Equal(t, 1, len(types.RoleInfoMap.GetRoleList()))
	cachedRole := types.RoleInfoMap.GetRoleList()[0]
	assert.Equal(t, "NETWORK_ADMIN", cachedRole.OrgId)
	assert.Equal(t, "NETWORK_ADMIN_ROLE", cachedRole.RoleId)
	assert.True(t, cachedRole.Active)
	assert.True(t, cachedRole.IsAdmin)
	assert.True(t, cachedRole.IsVoter)
	assert.Equal(t, types.FullAccess, cachedRole.Access)

	assert.Equal(t, 0, len(types.NodeInfoMap.GetNodeList()))

	assert.Equal(t, 1, len(types.AcctInfoMap.GetAcctList()))
	cachedAccount := types.AcctInfoMap.GetAcctList()[0]
	assert.Equal(t, "NETWORK_ADMIN", cachedAccount.OrgId)
	assert.Equal(t, "NETWORK_ADMIN_ROLE", cachedAccount.RoleId)
	assert.Equal(t, types.AcctActive, cachedAccount.Status)
	assert.True(t, cachedAccount.IsOrgAdmin)
	assert.Equal(t, guardianAddress, cachedAccount.AcctId)
}

func typicalPermissionCtrl(t *testing.T) *PermissionCtrl {
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
	testObject.ethClnt = backend
	testObject.eth = ethereum
	go func() {
		testObject.errorChan <- nil
	}()
	return testObject
}

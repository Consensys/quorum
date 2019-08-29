package permission

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"io/ioutil"
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

const (
	arbitraryNetworkAdminOrg  = "NETWORK_ADMIN"
	arbitraryNetworkAdminRole = "NETWORK_ADMIN_ROLE"
	arbitraryOrgAdminRole     = "ORG_ADMIN_ROLE"
	arbitraryNode1 = "enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@127.0.0.1:21000?discport=0&raftport=50401"
	arbitraryNode2 = "enode://0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416@127.0.0.1:21001?discport=0&raftport=50402"
	arbitraryOrgToAdd = "ORG1"
	arbitrarySubOrg = "SUB1"
)

var ErrOperationNotAllowed = errors.New("Operation not allowed")
var ErrNodeDoesNoteExist = errors.New("Node does not exists")

var (
	testObject  *PermissionCtrl
	guardianKey *ecdsa.PrivateKey
	guardianAccount accounts.Account
	backend     bind.ContractBackend
	permUpgrAddress, permInterfaceAddress, permImplAddress, voterManagerAddress,
	nodeManagerAddress, roleManagerAddress, accountManagerAddress, orgManagerAddress common.Address
	ethereum        *eth.Ethereum
	stack           *node.Node
	guardianAddress common.Address
	ksdir string
)


func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() {
	var err error
	t := log.New(os.Stdout, "", 0)

	ksdir, _, err := tmpKeyStore (false)
	defer os.RemoveAll(ksdir)

	if err != nil {
		t.Fatal("failed to create keystore: %v\n", err)
	}
	nodeKey, _ := crypto.GenerateKey()
	guardianKey, _ = crypto.GenerateKey()

	// Create a network-less protocol stack and start an Ethereum service within
	stack, err = node.New(&node.Config{
		DataDir:           "",
		KeyStoreDir: ksdir,
		UseLightweightKDF: true,
		P2P: p2p.Config{
			PrivateKey: nodeKey,
		},
	})

	ksbackend := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	guardianAccount, err = ksbackend.ImportECDSA(guardianKey, "foo")
	guardianAddress = guardianAccount.Address
	if err != nil {
		t.Fatal(err)
	}
	err = ksbackend.TimedUnlock(guardianAccount, "foo", 0)
	if err != nil {
		t.Fatal("failed to unlock")
	}

	genesisAlloc := map[common.Address]core.GenesisAccount{
		guardianAddress: {
			Balance: big.NewInt(100000000000000),
		},
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

	guardianTransactor := bind.NewKeyedTransactor(guardianKey)

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

	fmt.Printf("current block is %v\n", ethereum.BlockChain().CurrentBlock().Number().Int64())
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

	isNetworkInitialized, err := testObject.permInterf.GetNetworkBootStatus(&bind.CallOpts{
		Pending: true,
	})
	assert.NoError(t, err)
	assert.True(t, isNetworkInitialized)
}

func TestPermissionCtrl_PopulateInitPermissions_AfterNetworkIsInitialized(t *testing.T) {
	testObject := typicalPermissionCtrl(t)
	assert.NoError(t, testObject.AfterStart())

	err := testObject.populateInitPermissions()

	assert.NoError(t, err)

	// assert cache
	assert.Equal(t, 1, len(types.OrgInfoMap.GetOrgList()))
	cachedOrg := types.OrgInfoMap.GetOrgList()[0]
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedOrg.OrgId)
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedOrg.FullOrgId)
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedOrg.UltimateParent)
	assert.Equal(t, "", cachedOrg.ParentOrgId)
	assert.Equal(t, types.OrgApproved, cachedOrg.Status)
	assert.Equal(t, 0, len(cachedOrg.SubOrgList))
	assert.Equal(t, big.NewInt(1), cachedOrg.Level)

	assert.Equal(t, 1, len(types.RoleInfoMap.GetRoleList()))
	cachedRole := types.RoleInfoMap.GetRoleList()[0]
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedRole.OrgId)
	assert.Equal(t, arbitraryNetworkAdminRole, cachedRole.RoleId)
	assert.True(t, cachedRole.Active)
	assert.True(t, cachedRole.IsAdmin)
	assert.True(t, cachedRole.IsVoter)
	assert.Equal(t, types.FullAccess, cachedRole.Access)

	assert.Equal(t, 0, len(types.NodeInfoMap.GetNodeList()))

	assert.Equal(t, 1, len(types.AcctInfoMap.GetAcctList()))
	cachedAccount := types.AcctInfoMap.GetAcctList()[0]
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedAccount.OrgId)
	assert.Equal(t, arbitraryNetworkAdminRole, cachedAccount.RoleId)
	assert.Equal(t, types.AcctActive, cachedAccount.Status)
	assert.True(t, cachedAccount.IsOrgAdmin)
	assert.Equal(t, guardianAddress, cachedAccount.AcctId)
}

func TestQuorumControlsAPI_TestAPIs(t *testing.T) {
	testObject := typicalPermissionCtrl(t)
	assert.NoError(t, testObject.AfterStart())

	err := testObject.populateInitPermissions()

	assert.NoError(t, err)
	q := NewQuorumControlsAPI(testObject)

	// test GetOrgDetails
	orgDetails, err := q.GetOrgDetails(arbitraryNetworkAdminOrg)
	assert.NoError(t, err)
	assert.Equal(t, orgDetails.AcctList[0].AcctId, guardianAddress)
	assert.Equal(t, orgDetails.RoleList[0].RoleId, arbitraryNetworkAdminRole)

	// test NodeList
	assert.Equal(t, len(q.NodeList()), 0 )
	// test AcctList
	assert.True(t, len(q.AcctList()) > 0 , fmt.Sprintf("expected non zero account list"))
	// test OrgList
	assert.True(t, len(q.OrgList()) > 0 , fmt.Sprintf("expected non zero org list"))
	// test RoleList
	assert.True(t, len(q.RoleList()) > 0 , fmt.Sprintf("expected non zero org list"))


	// test AddOrg
	orgAdminKey, _ := crypto.GenerateKey()
	orgAdminAddress := crypto.PubkeyToAddress(orgAdminKey.PublicKey)

	txa := ethapi.SendTxArgs{From:guardianAddress}
	//time.Sleep(10 * time.Second)

	_, err = q.AddOrg(arbitraryOrgToAdd, arbitraryNode1, orgAdminAddress, txa )
	assert.NoError(t, err)

	_, err = q.ApproveOrg(arbitraryOrgToAdd, arbitraryNode1, orgAdminAddress, txa )
	assert.NoError(t, err)

	types.OrgInfoMap.UpsertOrg(arbitraryOrgToAdd, "", arbitraryOrgToAdd, big.NewInt(1), types.OrgApproved )
	_, err = q.UpdateOrgStatus(arbitraryOrgToAdd, uint8(SuspendOrg), txa)
	assert.NoError(t, err)

	types.OrgInfoMap.UpsertOrg(arbitraryOrgToAdd, "", arbitraryOrgToAdd, big.NewInt(1), types.OrgSuspended )
	_, err = q.ApproveOrgStatus(arbitraryOrgToAdd, uint8(SuspendOrg), txa)
	assert.NoError(t, err)

	_, err = q.AddSubOrg(arbitraryNetworkAdminOrg, arbitrarySubOrg, "", txa )
	assert.NoError(t, err)
	types.OrgInfoMap.UpsertOrg(arbitrarySubOrg, arbitraryNetworkAdminOrg, arbitraryNetworkAdminOrg, big.NewInt(2), types.OrgApproved)


	_, err = q.AddNode(arbitraryNetworkAdminOrg, arbitraryNode2 , txa )
	assert.NoError(t, err)
	types.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, types.NodeApproved)

	_, err = q.UpdateNodeStatus(arbitraryNetworkAdminOrg, arbitraryNode2 , uint8(SuspendNode), txa )
	assert.NoError(t, err)
	types.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2,types.NodeDeactivated)

	_, err = q.UpdateNodeStatus(arbitraryNetworkAdminOrg, arbitraryNode2 , uint8(ActivateSuspendedNode), txa )
	assert.NoError(t, err)
	types.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, types.NodeApproved)

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
		NwAdminOrg:     arbitraryNetworkAdminOrg,
		NwAdminRole:    arbitraryNetworkAdminRole,
		OrgAdminRole:   arbitraryOrgAdminRole,
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

func tmpKeyStore( encrypted bool) (string, *keystore.KeyStore, error) {
	d, err := ioutil.TempDir("", "eth-keystore-test")
	if err != nil {
		return "", nil, err
	}
	new := keystore.NewPlaintextKeyStore
	if encrypted {
		new = func(kd string) *keystore.KeyStore { return keystore.NewKeyStore(kd, keystore.LightScryptN, keystore.LightScryptP) }
	}
	return d, new(d), err
}


package permission

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	pcore "github.com/ethereum/go-ethereum/permission/core"
	ptype "github.com/ethereum/go-ethereum/permission/core/types"
	"github.com/ethereum/go-ethereum/permission/v1"
	v1bind "github.com/ethereum/go-ethereum/permission/v1/bind"
	"github.com/ethereum/go-ethereum/permission/v2"
	v2bind "github.com/ethereum/go-ethereum/permission/v2/bind"
	"github.com/stretchr/testify/assert"
)

const (
	arbitraryNetworkAdminOrg   = "NETWORK_ADMIN"
	arbitraryNetworkAdminRole  = "NETWORK_ADMIN_ROLE"
	arbitraryOrgAdminRole      = "ORG_ADMIN_ROLE"
	arbitraryNode1             = "enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@127.0.0.1:21000?discport=0&raftport=50401"
	arbitraryNode2             = "enode://0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416@127.0.0.1:21001?discport=0&raftport=50402"
	arbitraryNode3             = "enode://579f786d4e2830bbcc02815a27e8a9bacccc9605df4dc6f20bcc1a6eb391e7225fff7cb83e5b4ecd1f3a94d8b733803f2f66b7e871961e7b029e22c155c3a778@127.0.0.1:21002?discport=0&raftport=50403"
	arbitraryNode4             = "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404"
	arbitraryNode4withHostName = "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@lcoalhost:21003?discport=0&raftport=50404"
	arbitraryOrgToAdd          = "ORG1"
	arbitrarySubOrg            = "SUB1"
	arbitrartNewRole1          = "NEW_ROLE_1"
	arbitrartNewRole2          = "NEW_ROLE_2"
	orgCacheSize               = 4
	roleCacheSize              = 4
	nodeCacheSize              = 2
	accountCacheSize           = 4
)

var ErrAccountsLinked = errors.New("Accounts linked to the role. Cannot be removed")
var ErrPendingApproval = errors.New("Pending approvals for the organization. Approve first")
var ErrAcctBlacklisted = errors.New("Blacklisted account. Operation not allowed")
var ErrNodeBlacklisted = errors.New("Blacklisted node. Operation not allowed")

var (
	guardianKey     *ecdsa.PrivateKey
	guardianAccount accounts.Account
	contrBackend    bind.ContractBackend
	ethereum        *eth.Ethereum
	stack           *node.Node
	guardianAddress common.Address
	v2Flag          bool

	permUpgrAddress, permInterfaceAddress, permImplAddress, voterManagerAddress,
	nodeManagerAddress, roleManagerAddress, accountManagerAddress, orgManagerAddress common.Address
)

func TestMain(m *testing.M) {
	var v2FlagVer = []bool{false}
	var ret int
	for i := range v2FlagVer {
		v2Flag = v2FlagVer[i]
		setup()
		ret = m.Run()
		teardown()
		if ret != 0 {
			os.Exit(ret)
		}
	}
	os.Exit(ret)
}

func setup() {
	var err error
	t := log.New(os.Stdout, "", 0)

	ksdir, _, err := tmpKeyStore(false)
	defer os.RemoveAll(ksdir)

	if err != nil {
		t.Fatalf("failed to create keystore: %v\n", err)
	}
	nodeKey, _ := crypto.GenerateKey()
	guardianKey, _ = crypto.GenerateKey()

	// Create a network-less protocol stack and start an Ethereum service within
	stack, err = node.New(&node.Config{
		DataDir:           "",
		KeyStoreDir:       ksdir,
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
		Genesis: &core.Genesis{Config: params.AllEthashProtocolChanges, GasLimit: 10000000000, Alloc: genesisAlloc},
		Miner:   miner.Config{Etherbase: guardianAddress},
		Ethash: ethash.Config{
			PowMode: ethash.ModeTest,
		},
	}

	if err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) { return eth.New(ctx, ethConf) }); err != nil {
		t.Fatalf("failed to register Ethereum protocol: %v", err)
	}
	// Start the Node and assemble the JavaScript console around it
	if err = stack.Start(); err != nil {
		t.Fatalf("failed to start test stack: %v", err)
	}
	if err := stack.Service(&ethereum); err != nil {
		t.Fatal(err)
	}
	contrBackend = backends.NewSimulatedBackendFrom(ethereum)

	var permUpgrInstance *v1bind.PermUpgr
	var permUpgrInstanceE *v2bind.PermUpgr

	guardianTransactor := bind.NewKeyedTransactor(guardianKey)

	if v2Flag {
		permUpgrAddress, _, permUpgrInstanceE, err = v2bind.DeployPermUpgr(guardianTransactor, contrBackend, guardianAddress)
		if err != nil {
			t.Fatal(err)
		}
		permInterfaceAddress, _, _, err = v2bind.DeployPermInterface(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		nodeManagerAddress, _, _, err = v2bind.DeployNodeManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		roleManagerAddress, _, _, err = v2bind.DeployRoleManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		accountManagerAddress, _, _, err = v2bind.DeployAcctManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		orgManagerAddress, _, _, err = v2bind.DeployOrgManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		voterManagerAddress, _, _, err = v2bind.DeployVoterManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		permImplAddress, _, _, err = v2bind.DeployPermImpl(guardianTransactor, contrBackend, permUpgrAddress, orgManagerAddress, roleManagerAddress, accountManagerAddress, voterManagerAddress, nodeManagerAddress)
		if err != nil {
			t.Fatal(err)
		}
		// call init
		if _, err := permUpgrInstanceE.Init(guardianTransactor, permInterfaceAddress, permImplAddress); err != nil {
			t.Fatal(err)
		}
	} else {
		permUpgrAddress, _, permUpgrInstance, err = v1bind.DeployPermUpgr(guardianTransactor, contrBackend, guardianAddress)
		if err != nil {
			t.Fatal(err)
		}
		permInterfaceAddress, _, _, err = v1bind.DeployPermInterface(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		nodeManagerAddress, _, _, err = v1bind.DeployNodeManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		roleManagerAddress, _, _, err = v1bind.DeployRoleManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		accountManagerAddress, _, _, err = v1bind.DeployAcctManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		orgManagerAddress, _, _, err = v1bind.DeployOrgManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		voterManagerAddress, _, _, err = v1bind.DeployVoterManager(guardianTransactor, contrBackend, permUpgrAddress)
		if err != nil {
			t.Fatal(err)
		}
		permImplAddress, _, _, err = v1bind.DeployPermImpl(guardianTransactor, contrBackend, permUpgrAddress, orgManagerAddress, roleManagerAddress, accountManagerAddress, voterManagerAddress, nodeManagerAddress)
		if err != nil {
			t.Fatal(err)
		}
		// call init
		if _, err := permUpgrInstance.Init(guardianTransactor, permInterfaceAddress, permImplAddress); err != nil {
			t.Fatal(err)
		}
	}
	fmt.Printf("current block is %v\n", ethereum.BlockChain().CurrentBlock().Number().Int64())
}

func teardown() {

}

func TestPermissionCtrl_AfterStart(t *testing.T) {
	testObject := typicalPermissionCtrl(t, v2Flag)

	err := testObject.AfterStart()

	assert.NoError(t, err)
	if testObject.IsV2Permission() {
		var contract *v2.Init
		contract, _ = testObject.contract.(*v2.Init)
		assert.NotNil(t, contract.PermOrg)
		assert.NotNil(t, contract.PermRole)
		assert.NotNil(t, contract.PermNode)
		assert.NotNil(t, contract.PermAcct)
		assert.NotNil(t, contract.PermInterf)
		assert.NotNil(t, contract.PermUpgr)
	} else {
		var contract *v1.Init
		contract, _ = testObject.contract.(*v1.Init)
		assert.NotNil(t, contract.PermOrg)
		assert.NotNil(t, contract.PermRole)
		assert.NotNil(t, contract.PermNode)
		assert.NotNil(t, contract.PermAcct)
		assert.NotNil(t, contract.PermInterf)
		assert.NotNil(t, contract.PermUpgr)
	}

	isNetworkInitialized, err := testObject.contract.GetNetworkBootStatus()
	assert.NoError(t, err)
	assert.True(t, isNetworkInitialized)
}

func TestPermissionCtrl_PopulateInitPermissions_AfterNetworkIsInitialized(t *testing.T) {
	testObject := typicalPermissionCtrl(t, v2Flag)
	assert.NoError(t, testObject.AfterStart())

	err := testObject.populateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize)

	assert.NoError(t, err)

	// assert cache
	assert.Equal(t, 1, len(pcore.OrgInfoMap.GetOrgList()))
	cachedOrg := pcore.OrgInfoMap.GetOrgList()[0]
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedOrg.OrgId)
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedOrg.FullOrgId)
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedOrg.UltimateParent)
	assert.Equal(t, "", cachedOrg.ParentOrgId)
	assert.Equal(t, pcore.OrgApproved, cachedOrg.Status)
	assert.Equal(t, 0, len(cachedOrg.SubOrgList))
	assert.Equal(t, big.NewInt(1), cachedOrg.Level)

	assert.Equal(t, 1, len(pcore.RoleInfoMap.GetRoleList()))
	cachedRole := pcore.RoleInfoMap.GetRoleList()[0]
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedRole.OrgId)
	assert.Equal(t, arbitraryNetworkAdminRole, cachedRole.RoleId)
	assert.True(t, cachedRole.Active)
	assert.True(t, cachedRole.IsAdmin)
	assert.True(t, cachedRole.IsVoter)
	assert.Equal(t, pcore.FullAccess, cachedRole.Access)

	assert.Equal(t, 0, len(pcore.NodeInfoMap.GetNodeList()))

	assert.Equal(t, 1, len(pcore.AcctInfoMap.GetAcctList()))
	cachedAccount := pcore.AcctInfoMap.GetAcctList()[0]
	assert.Equal(t, arbitraryNetworkAdminOrg, cachedAccount.OrgId)
	assert.Equal(t, arbitraryNetworkAdminRole, cachedAccount.RoleId)
	assert.Equal(t, pcore.AcctActive, cachedAccount.Status)
	assert.True(t, cachedAccount.IsOrgAdmin)
	assert.Equal(t, guardianAddress, cachedAccount.AcctId)
}

func typicalQuorumControlsAPI(t *testing.T) *QuorumControlsAPI {
	pc := typicalPermissionCtrl(t, v2Flag)
	if !assert.NoError(t, pc.AfterStart()) {
		t.Fail()
	}
	if !assert.NoError(t, pc.populateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize)) {
		t.Fail()
	}
	return NewQuorumControlsAPI(pc)
}

func TestQuorumControlsAPI_ListAPIs(t *testing.T) {
	testObject := typicalQuorumControlsAPI(t)

	orgDetails, err := testObject.GetOrgDetails(arbitraryNetworkAdminOrg)
	assert.NoError(t, err)
	assert.Equal(t, orgDetails.AcctList[0].AcctId, guardianAddress)
	assert.Equal(t, orgDetails.RoleList[0].RoleId, arbitraryNetworkAdminRole)

	orgDetails, err = testObject.GetOrgDetails("XYZ")
	assert.Equal(t, err, errors.New("Org does not exist"))

	// test NodeList
	assert.Equal(t, len(testObject.NodeList()), 0)
	// test AcctList
	assert.True(t, len(testObject.AcctList()) > 0, fmt.Sprintf("expected non zero account list"))
	// test OrgList
	assert.True(t, len(testObject.OrgList()) > 0, fmt.Sprintf("expected non zero org list"))
	// test RoleList
	assert.True(t, len(testObject.RoleList()) > 0, fmt.Sprintf("expected non zero org list"))
}

func TestQuorumControlsAPI_OrgAPIs(t *testing.T) {
	testObject := typicalQuorumControlsAPI(t)
	invalidTxa := ethapi.SendTxArgs{From: getArbitraryAccount()}

	// test AddOrg
	orgAdminKey, _ := crypto.GenerateKey()
	orgAdminAddress := crypto.PubkeyToAddress(orgAdminKey.PublicKey)

	txa := ethapi.SendTxArgs{From: guardianAddress}
	_, err := testObject.AddOrg(arbitraryOrgToAdd, arbitraryNode1, orgAdminAddress, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.AddOrg(arbitraryOrgToAdd, arbitraryNode1, orgAdminAddress, txa)
	assert.NoError(t, err)

	_, err = testObject.AddOrg(arbitraryOrgToAdd, arbitraryNode1, orgAdminAddress, txa)
	assert.Equal(t, err, ErrPendingApproval)

	_, err = testObject.ApproveOrg(arbitraryOrgToAdd, arbitraryNode1, orgAdminAddress, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.ApproveOrg("XYZ", arbitraryNode1, orgAdminAddress, txa)
	assert.Equal(t, err, errors.New("Nothing to approve"))

	_, err = testObject.ApproveOrg(arbitraryOrgToAdd, arbitraryNode1, orgAdminAddress, txa)
	assert.NoError(t, err)

	pcore.OrgInfoMap.UpsertOrg(arbitraryOrgToAdd, "", arbitraryOrgToAdd, big.NewInt(1), pcore.OrgApproved)
	_, err = testObject.UpdateOrgStatus(arbitraryOrgToAdd, uint8(SuspendOrg), invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.UpdateOrgStatus(arbitraryOrgToAdd, uint8(SuspendOrg), txa)
	assert.NoError(t, err)

	pcore.OrgInfoMap.UpsertOrg(arbitraryOrgToAdd, "", arbitraryOrgToAdd, big.NewInt(1), pcore.OrgSuspended)
	_, err = testObject.ApproveOrgStatus(arbitraryOrgToAdd, uint8(SuspendOrg), invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.ApproveOrgStatus(arbitraryOrgToAdd, uint8(SuspendOrg), txa)
	assert.NoError(t, err)

	_, err = testObject.AddSubOrg(arbitraryNetworkAdminOrg, arbitrarySubOrg, "", invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.AddSubOrg(arbitraryNetworkAdminOrg, arbitrarySubOrg, "", txa)
	assert.NoError(t, err)
	pcore.OrgInfoMap.UpsertOrg(arbitrarySubOrg, arbitraryNetworkAdminOrg, arbitraryNetworkAdminOrg, big.NewInt(2), pcore.OrgApproved)

	suborg := "ABC.12345"
	_, err = testObject.AddSubOrg(arbitraryNetworkAdminOrg, suborg, "", txa)
	assert.Equal(t, err, errors.New("Org id cannot contain special characters"))

	_, err = testObject.AddSubOrg(arbitraryNetworkAdminOrg, "", "", txa)
	assert.Equal(t, err, errors.New("Invalid input"))

	// caching tests - cache size for org is 4. add 4 sub orgs
	// this will result in cache eviction
	// get org details after this
	for i := 0; i < orgCacheSize; i++ {
		subOrgId := "TESTSUBORG" + strconv.Itoa(i)
		_, err = testObject.AddSubOrg(arbitraryNetworkAdminOrg, subOrgId, "", txa)
		assert.NoError(t, err)
		pcore.OrgInfoMap.UpsertOrg(subOrgId, arbitraryNetworkAdminOrg, arbitraryNetworkAdminOrg, big.NewInt(2), pcore.OrgApproved)
	}

	assert.Equal(t, orgCacheSize, len(pcore.OrgInfoMap.GetOrgList()))

	orgDetails, err := testObject.GetOrgDetails(arbitraryNetworkAdminOrg)
	assert.Equal(t, orgDetails.AcctList[0].AcctId, guardianAddress)
	assert.Equal(t, orgDetails.RoleList[0].RoleId, arbitraryNetworkAdminRole)
}

func testConnectionAllowed(t *testing.T, q *QuorumControlsAPI, url string, expected bool) {
	enode, ip, port, raftPort, err := ptype.GetNodeDetails(url, false, false)
	if q.permCtrl.IsV2Permission() {
		assert.NoError(t, err)
		connAllowed := q.ConnectionAllowed(enode, ip, port, raftPort)
		assert.Equal(t, expected, connAllowed)
	} else {
		assert.Equal(t, isNodePermissionedV1(url, enode, enode, "INCOMING"), expected)
	}
}

func TestQuorumControlsAPI_NodeAPIs(t *testing.T) {
	testObject := typicalQuorumControlsAPI(t)
	invalidTxa := ethapi.SendTxArgs{From: getArbitraryAccount()}
	txa := ethapi.SendTxArgs{From: guardianAddress}

	testObject.permCtrl.isRaft = true
	_, err := testObject.AddNode(arbitraryNetworkAdminOrg, arbitraryNode2, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))
	testConnectionAllowed(t, testObject, arbitraryNode2, false)

	_, err = testObject.AddNode(arbitraryNetworkAdminOrg, arbitraryNode2, txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, pcore.NodeApproved)
	testConnectionAllowed(t, testObject, arbitraryNode2, true)

	_, err = testObject.UpdateNodeStatus(arbitraryNetworkAdminOrg, arbitraryNode2, uint8(SuspendNode), invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.UpdateNodeStatus(arbitraryNetworkAdminOrg, arbitraryNode2, uint8(SuspendNode), txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, pcore.NodeDeactivated)
	testConnectionAllowed(t, testObject, arbitraryNode2, false)

	_, err = testObject.UpdateNodeStatus(arbitraryNetworkAdminOrg, arbitraryNode2, uint8(ActivateSuspendedNode), txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, pcore.NodeApproved)
	testConnectionAllowed(t, testObject, arbitraryNode2, true)

	_, err = testObject.UpdateNodeStatus(arbitraryNetworkAdminOrg, arbitraryNode2, uint8(BlacklistNode), txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, pcore.NodeBlackListed)

	_, err = testObject.UpdateNodeStatus(arbitraryNetworkAdminOrg, arbitraryNode2, uint8(ActivateSuspendedNode), txa)
	assert.Equal(t, err, ErrNodeBlacklisted)

	_, err = testObject.RecoverBlackListedNode(arbitraryNetworkAdminOrg, arbitraryNode2, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.RecoverBlackListedNode(arbitraryNetworkAdminOrg, arbitraryNode2, txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, pcore.NodeRecoveryInitiated)

	_, err = testObject.ApproveBlackListedNodeRecovery(arbitraryNetworkAdminOrg, arbitraryNode2, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.ApproveBlackListedNodeRecovery(arbitraryNetworkAdminOrg, arbitraryNode2, txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode2, pcore.NodeApproved)

	// caching tests - cache size for Node is 3. add 2 nodes which will
	// result in Node eviction from cache. get evicted Node details using api
	_, err = testObject.AddNode(arbitraryNetworkAdminOrg, arbitraryNode3, txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode3, pcore.NodeApproved)

	testObject.permCtrl.isRaft = true
	_, err = testObject.AddNode(arbitraryNetworkAdminOrg, arbitraryNode4withHostName, txa)
	assert.Equal(t, err, ptype.ErrHostNameNotSupported)

	_, err = testObject.AddNode(arbitraryNetworkAdminOrg, arbitraryNode4, txa)
	assert.NoError(t, err)
	pcore.NodeInfoMap.UpsertNode(arbitraryNetworkAdminOrg, arbitraryNode4, pcore.NodeApproved)

	assert.Equal(t, nodeCacheSize, len(pcore.NodeInfoMap.GetNodeList()))
	nodeInfo, err := pcore.NodeInfoMap.GetNodeByUrl(arbitraryNode4)
	assert.True(t, err == nil, "Node fetch returned error")
	assert.Equal(t, pcore.NodeApproved, nodeInfo.Status)
}

func testTransactionAllowed(t *testing.T, q *QuorumControlsAPI, txa ethapi.SendTxArgs, expected bool) {
	actAllowed := q.TransactionAllowed(txa)
	assert.Equal(t, expected, actAllowed)
}

func TestQuorumControlsAPI_TransactionAllowed(t *testing.T) {
	testObject := typicalQuorumControlsAPI(t)

	if testObject.permCtrl.IsV2Permission() {

		acct := getArbitraryAccount()
		txa := ethapi.SendTxArgs{From: guardianAddress}
		payload := hexutil.Bytes(([]byte("0x43d3e767000000000000000000000000000000000000000000000000000000000000000a"))[:])
		value := hexutil.Big(*(big.NewInt(10)))

		transactionTxa := ethapi.SendTxArgs{From: acct, To: &guardianAddress, Value: &value}
		contractCallTxa := ethapi.SendTxArgs{From: acct, To: &guardianAddress, Data: &payload}
		contractCreateTxa := ethapi.SendTxArgs{From: acct, To: &common.Address{}, Data: &payload}

		for i := 0; i < 8; i++ {
			roleId := arbitrartNewRole1 + strconv.Itoa(i)
			_, err := testObject.AddNewRole(arbitraryNetworkAdminOrg, roleId, uint8(i), false, false, txa)
			assert.NoError(t, err)
			pcore.RoleInfoMap.UpsertRole(arbitraryNetworkAdminOrg, roleId, false, false, pcore.AccessType(uint8(i)), true)

			if i == 0 {
				_, err = testObject.AddAccountToOrg(acct, arbitraryNetworkAdminOrg, roleId, txa)
				assert.NoError(t, err)
			} else {
				_, err = testObject.ChangeAccountRole(acct, arbitraryNetworkAdminOrg, roleId, txa)
				assert.NoError(t, err)
			}

			switch pcore.AccessType(uint8(i)) {
			case pcore.ReadOnly:
				testTransactionAllowed(t, testObject, transactionTxa, false)
				testTransactionAllowed(t, testObject, contractCallTxa, false)
				testTransactionAllowed(t, testObject, contractCreateTxa, false)

			case pcore.Transact:
				testTransactionAllowed(t, testObject, transactionTxa, true)
				testTransactionAllowed(t, testObject, contractCallTxa, false)
				testTransactionAllowed(t, testObject, contractCreateTxa, false)

			case pcore.ContractDeploy:
				testTransactionAllowed(t, testObject, transactionTxa, false)
				testTransactionAllowed(t, testObject, contractCallTxa, false)
				testTransactionAllowed(t, testObject, contractCreateTxa, true)

			case pcore.FullAccess:
				testTransactionAllowed(t, testObject, transactionTxa, true)
				testTransactionAllowed(t, testObject, contractCallTxa, true)
				testTransactionAllowed(t, testObject, contractCreateTxa, true)

			case pcore.ContractCall:
				testTransactionAllowed(t, testObject, transactionTxa, false)
				testTransactionAllowed(t, testObject, contractCallTxa, true)
				testTransactionAllowed(t, testObject, contractCreateTxa, false)

			case pcore.TransactAndContractCall:
				testTransactionAllowed(t, testObject, transactionTxa, true)
				testTransactionAllowed(t, testObject, contractCallTxa, true)
				testTransactionAllowed(t, testObject, contractCreateTxa, false)

			case pcore.TransactAndContractDeploy:
				testTransactionAllowed(t, testObject, transactionTxa, true)
				testTransactionAllowed(t, testObject, contractCallTxa, false)
				testTransactionAllowed(t, testObject, contractCreateTxa, true)
			case pcore.ContractCallAndDeploy:
				testTransactionAllowed(t, testObject, transactionTxa, false)
				testTransactionAllowed(t, testObject, contractCallTxa, true)
				testTransactionAllowed(t, testObject, contractCreateTxa, true)

			}

		}
	}

}

func TestQuorumControlsAPI_RoleAndAccountsAPIs(t *testing.T) {
	testObject := typicalQuorumControlsAPI(t)
	invalidTxa := ethapi.SendTxArgs{From: getArbitraryAccount()}
	acct := getArbitraryAccount()
	txa := ethapi.SendTxArgs{From: guardianAddress, To: &acct}

	pcore.SetNetworkBootUpCompleted()
	pcore.SetQIP714BlockReached()

	_, err := testObject.AssignAdminRole(arbitraryNetworkAdminOrg, acct, arbitraryNetworkAdminRole, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.AssignAdminRole(arbitraryNetworkAdminOrg, acct, arbitraryNetworkAdminRole, txa)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitraryNetworkAdminRole, acct, true, pcore.AcctPendingApproval)

	_, err = testObject.ApproveAdminRole(arbitraryNetworkAdminOrg, acct, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))
	testTransactionAllowed(t, testObject, ethapi.SendTxArgs{From: acct, To: &acct}, false)

	_, err = testObject.ApproveAdminRole(arbitraryNetworkAdminOrg, acct, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.ApproveAdminRole(arbitraryNetworkAdminOrg, acct, txa)
	assert.NoError(t, err)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitraryNetworkAdminRole, acct, true, pcore.AcctActive)
	testTransactionAllowed(t, testObject, ethapi.SendTxArgs{From: acct, To: &acct}, true)

	_, err = testObject.AddNewRole(arbitraryNetworkAdminOrg, arbitrartNewRole1, uint8(pcore.FullAccess), false, false, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.AddNewRole(arbitraryNetworkAdminOrg, arbitrartNewRole1, uint8(pcore.FullAccess), false, false, txa)
	assert.NoError(t, err)
	pcore.RoleInfoMap.UpsertRole(arbitraryNetworkAdminOrg, arbitrartNewRole1, false, false, pcore.FullAccess, true)

	acct = getArbitraryAccount()
	_, err = testObject.AddAccountToOrg(acct, arbitraryNetworkAdminOrg, arbitrartNewRole1, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.AddAccountToOrg(acct, arbitraryNetworkAdminOrg, arbitrartNewRole1, txa)
	assert.NoError(t, err)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitrartNewRole1, acct, true, pcore.AcctActive)

	_, err = testObject.RemoveRole(arbitraryNetworkAdminOrg, arbitrartNewRole1, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.RemoveRole(arbitraryNetworkAdminOrg, arbitrartNewRole1, txa)
	assert.Equal(t, err, ErrAccountsLinked)

	_, err = testObject.AddNewRole(arbitraryNetworkAdminOrg, arbitrartNewRole2, uint8(pcore.FullAccess), false, false, txa)
	assert.NoError(t, err)
	pcore.RoleInfoMap.UpsertRole(arbitraryNetworkAdminOrg, arbitrartNewRole2, false, false, pcore.FullAccess, true)

	_, err = testObject.ChangeAccountRole(acct, arbitraryNetworkAdminOrg, arbitrartNewRole2, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.ChangeAccountRole(acct, arbitraryNetworkAdminOrg, arbitrartNewRole2, txa)
	assert.NoError(t, err)

	_, err = testObject.RemoveRole(arbitraryNetworkAdminOrg, arbitrartNewRole1, txa)
	assert.Equal(t, err, ErrAccountsLinked)

	_, err = testObject.UpdateAccountStatus(arbitraryNetworkAdminOrg, acct, uint8(SuspendAccount), invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.UpdateAccountStatus(arbitraryNetworkAdminOrg, acct, uint8(SuspendAccount), txa)
	assert.NoError(t, err)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitrartNewRole2, acct, true, pcore.AcctSuspended)
	testTransactionAllowed(t, testObject, ethapi.SendTxArgs{From: acct, To: &acct}, false)

	_, err = testObject.UpdateAccountStatus(arbitraryNetworkAdminOrg, acct, uint8(ActivateSuspendedAccount), txa)
	assert.NoError(t, err)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitrartNewRole2, acct, true, pcore.AcctActive)

	_, err = testObject.UpdateAccountStatus(arbitraryNetworkAdminOrg, acct, uint8(BlacklistAccount), txa)
	assert.NoError(t, err)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitrartNewRole2, acct, true, pcore.AcctBlacklisted)
	testTransactionAllowed(t, testObject, ethapi.SendTxArgs{From: acct, To: &acct}, false)

	_, err = testObject.UpdateAccountStatus(arbitraryNetworkAdminOrg, acct, uint8(ActivateSuspendedAccount), txa)
	assert.Equal(t, err, ErrAcctBlacklisted)

	_, err = testObject.RecoverBlackListedAccount(arbitraryNetworkAdminOrg, acct, invalidTxa)
	assert.Equal(t, err, errors.New("Invalid account id"))

	_, err = testObject.RecoverBlackListedAccount(arbitraryNetworkAdminOrg, acct, txa)
	assert.NoError(t, err)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitrartNewRole2, acct, true, pcore.AcctRecoveryInitiated)
	_, err = testObject.ApproveBlackListedAccountRecovery(arbitraryNetworkAdminOrg, acct, txa)
	assert.NoError(t, err)
	pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitrartNewRole2, acct, true, pcore.AcctActive)

	// check role cache. the cache size is set to 4
	// insert 4 records and then retrieve the 1st role
	for i := 0; i < roleCacheSize; i++ {
		roleId := "TESTROLE" + strconv.Itoa(i)
		_, err = testObject.AddNewRole(arbitraryNetworkAdminOrg, roleId, uint8(pcore.FullAccess), false, false, txa)
		assert.NoError(t, err)
		pcore.RoleInfoMap.UpsertRole(arbitraryNetworkAdminOrg, roleId, false, false, pcore.FullAccess, true)
	}

	assert.Equal(t, roleCacheSize, len(pcore.RoleInfoMap.GetRoleList()))
	roleInfo, err := pcore.RoleInfoMap.GetRole(arbitraryNetworkAdminOrg, arbitrartNewRole1)
	assert.True(t, err == nil, "error encountered")

	assert.Equal(t, roleInfo.RoleId, arbitrartNewRole1)

	// check account cache
	var AccountArray [4]common.Address
	AccountArray[0] = common.StringToAddress("0fbdc686b912d7722dc86510934589e0aaf3b55a")
	AccountArray[1] = common.StringToAddress("9186eb3d20cbd1f5f992a950d808c4495153abd5")
	AccountArray[2] = common.StringToAddress("0638e1574728b6d862dd5d3a3e0942c3be47d996")
	AccountArray[3] = common.StringToAddress("ae9bc6cd5145e67fbd1887a5145271fd182f0ee7")

	for i := 0; i < accountCacheSize; i++ {
		_, err = testObject.AddAccountToOrg(AccountArray[i], arbitraryNetworkAdminOrg, arbitrartNewRole1, txa)
		assert.NoError(t, err)
		pcore.AcctInfoMap.UpsertAccount(arbitraryNetworkAdminOrg, arbitrartNewRole1, AccountArray[i], false, pcore.AcctActive)
	}
	assert.Equal(t, accountCacheSize, len(pcore.AcctInfoMap.GetAcctList()))

	acctInfo, err := pcore.AcctInfoMap.GetAccount(acct)
	assert.True(t, err == nil, "error encountered")
	assert.True(t, acctInfo != nil, "account details nil")
}

func getArbitraryAccount() common.Address {
	acctKey, _ := crypto.GenerateKey()
	return crypto.PubkeyToAddress(acctKey.PublicKey)
}

func typicalPermissionCtrl(t *testing.T, v2Flag bool) *PermissionCtrl {
	pconfig := &ptype.PermissionConfig{
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
		SubOrgDepth:   big.NewInt(10),
		SubOrgBreadth: big.NewInt(10),
	}
	if v2Flag {
		pconfig.PermissionsModel = ptype.PERMISSION_V2
	} else {
		pconfig.PermissionsModel = ptype.PERMISSION_V1
	}
	testObject, err := NewQuorumPermissionCtrl(stack, pconfig, false)
	if err != nil {
		t.Fatal(err)
	}

	testObject.ethClnt = contrBackend
	testObject.eth = ethereum

	// set contract and backend's contract as asyncStart won't get called
	testObject.contract = NewPermissionContractService(testObject.ethClnt, testObject.IsV2Permission(), testObject.key, testObject.permConfig, false, false)
	if v2Flag {
		b := testObject.backend.(*v2.Backend)
		b.Contr = testObject.contract.(*v2.Init)
	} else {
		b := testObject.backend.(*v1.Backend)
		b.Contr = testObject.contract.(*v1.Init)
	}

	go func() {
		testObject.errorChan <- nil
	}()
	return testObject
}

func tmpKeyStore(encrypted bool) (string, *keystore.KeyStore, error) {
	d, err := ioutil.TempDir("", "Eth-keystore-test")
	if err != nil {
		return "", nil, err
	}
	newKs := keystore.NewPlaintextKeyStore
	if encrypted {
		newKs = func(kd string) *keystore.KeyStore {
			return keystore.NewKeyStore(kd, keystore.LightScryptN, keystore.LightScryptP)
		}
	}
	return d, newKs(d), err
}

func TestPermissionCtrl_whenUpdateFile(t *testing.T) {
	testObject := typicalPermissionCtrl(t, v2Flag)
	assert.NoError(t, testObject.AfterStart())

	err := testObject.populateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize)
	assert.NoError(t, err)

	d, _ := ioutil.TempDir("", "qdata")
	defer os.RemoveAll(d)

	testObject.dataDir = d
	ptype.UpdatePermissionedNodes(testObject.node, d, arbitraryNode1, ptype.NodeAdd, true)

	permFile, _ := os.Create(d + "/" + "permissioned-nodes.json")

	ptype.UpdateFile("testFile", arbitraryNode2, ptype.NodeAdd, false)
	ptype.UpdateFile(permFile.Name(), arbitraryNode2, ptype.NodeAdd, false)
	ptype.UpdateFile(permFile.Name(), arbitraryNode2, ptype.NodeAdd, true)
	ptype.UpdateFile(permFile.Name(), arbitraryNode2, ptype.NodeAdd, true)
	ptype.UpdateFile(permFile.Name(), arbitraryNode1, ptype.NodeAdd, false)
	ptype.UpdateFile(permFile.Name(), arbitraryNode1, ptype.NodeDelete, false)
	ptype.UpdateFile(permFile.Name(), arbitraryNode1, ptype.NodeDelete, false)

	blob, err := ioutil.ReadFile(permFile.Name())
	var nodeList []string
	if err := json.Unmarshal(blob, &nodeList); err != nil {
		t.Fatal("Failed to load nodes list from file", "fileName", permFile, "err", err)
		return
	}
	assert.Equal(t, len(nodeList), 1)
	ptype.UpdatePermissionedNodes(testObject.node, d, arbitraryNode1, ptype.NodeAdd, true)
	ptype.UpdatePermissionedNodes(testObject.node, d, arbitraryNode1, ptype.NodeDelete, true)

	blob, err = ioutil.ReadFile(permFile.Name())
	if err := json.Unmarshal(blob, &nodeList); err != nil {
		t.Fatal("Failed to load nodes list from file", "fileName", permFile, "err", err)
		return
	}
	assert.Equal(t, len(nodeList), 1)

	ptype.UpdateDisallowedNodes(d, arbitraryNode2, ptype.NodeAdd)
	ptype.UpdateDisallowedNodes(d, arbitraryNode2, ptype.NodeDelete)
	blob, err = ioutil.ReadFile(d + "/" + "disallowed-nodes.json")
	if err := json.Unmarshal(blob, &nodeList); err != nil {
		t.Fatal("Failed to load nodes list from file", "fileName", permFile, "err", err)
		return
	}
	assert.Equal(t, len(nodeList), 0)

}

func TestParsePermissionConfig(t *testing.T) {
	d, _ := ioutil.TempDir("", "qdata")
	defer os.RemoveAll(d)

	_, err := ptype.ParsePermissionConfig(d)
	assert.True(t, err != nil, "expected file not there error")

	fileName := d + "/permission-config.json"
	_, err = os.Create(fileName)
	_, err = ptype.ParsePermissionConfig(d)
	assert.True(t, err != nil, "expected unmarshalling error")

	// write permission-config.json into the temp dir
	var tmpPermCofig ptype.PermissionConfig
	tmpPermCofig.NwAdminOrg = arbitraryNetworkAdminOrg
	tmpPermCofig.NwAdminRole = arbitraryNetworkAdminRole
	tmpPermCofig.OrgAdminRole = arbitraryOrgAdminRole
	tmpPermCofig.InterfAddress = common.Address{}
	tmpPermCofig.ImplAddress = common.Address{}
	tmpPermCofig.UpgrdAddress = common.Address{}
	tmpPermCofig.VoterAddress = common.Address{}
	tmpPermCofig.RoleAddress = common.Address{}
	tmpPermCofig.OrgAddress = common.Address{}
	tmpPermCofig.NodeAddress = common.Address{}
	tmpPermCofig.SubOrgBreadth = new(big.Int)
	tmpPermCofig.SubOrgDepth = new(big.Int)

	blob, err := json.Marshal(tmpPermCofig)
	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		t.Fatal("Error writing new Node info to file", "fileName", fileName, "err", err)
	}
	_, err = ptype.ParsePermissionConfig(d)
	assert.True(t, err != nil, "permission model not given error")

	_ = os.Remove(fileName)
	tmpPermCofig.PermissionsModel = "ABCD"
	blob, _ = json.Marshal(tmpPermCofig)
	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		t.Fatal("Error writing new Node info to file", "fileName", fileName, "err", err)
	}

	_, err = ptype.ParsePermissionConfig(d)
	assert.True(t, err != nil, "invalid permission model error")
	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		t.Fatal("Error writing new Node info to file", "fileName", fileName, "err", err)
	}

	_ = os.Remove(fileName)
	tmpPermCofig.PermissionsModel = "v1"
	blob, _ = json.Marshal(tmpPermCofig)
	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		t.Fatal("Error writing new Node info to file", "fileName", fileName, "err", err)
	}

	_, err = ptype.ParsePermissionConfig(d)
	assert.True(t, err != nil, "expected account not given  error")

	_ = os.Remove(fileName)
	tmpPermCofig.Accounts = append(tmpPermCofig.Accounts, common.StringToAddress("0xed9d02e382b34818e88b88a309c7fe71e65f419d"))
	blob, err = json.Marshal(tmpPermCofig)
	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		t.Fatal("Error writing new Node info to file", "fileName", fileName, "err", err)
	}

	_, err = ptype.ParsePermissionConfig(d)
	assert.True(t, err != nil, "expected sub org depth not set error")

	_ = os.Remove(fileName)
	tmpPermCofig.SubOrgBreadth.Set(big.NewInt(4))
	tmpPermCofig.SubOrgDepth.Set(big.NewInt(4))
	blob, _ = json.Marshal(tmpPermCofig)
	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		t.Fatal("Error writing new Node info to file", "fileName", fileName, "err", err)
	}

	_, err = ptype.ParsePermissionConfig(d)
	assert.True(t, err != nil, "expected contract address error")

	_ = os.Remove(fileName)
	tmpPermCofig.InterfAddress = common.StringToAddress("0xed9d02e382b34818e88b88a309c7fe71e65f419d")
	blob, err = json.Marshal(tmpPermCofig)
	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		t.Fatal("Error writing new Node info to file", "fileName", fileName, "err", err)
	}
	permConfig, err := ptype.ParsePermissionConfig(d)
	assert.False(t, permConfig.IsEmpty(), "expected non empty object")
}

func TestIsTransactionAllowed_V1(t *testing.T) {
	testObject := typicalQuorumControlsAPI(t)
	pcore.PermissionTransactionAllowedFunc = testObject.permCtrl.IsTransactionAllowed
	var Acct1 = common.BytesToAddress([]byte("permission"))
	var Acct2 = common.BytesToAddress([]byte("perm-test"))
	pcore.SetDefaults(arbitraryNetworkAdminRole, arbitraryOrgAdminRole, false)
	pcore.SetQIP714BlockReached()
	pcore.SetNetworkBootUpCompleted()
	pcore.OrgInfoMap = pcore.NewOrgCache(params.DEFAULT_ORGCACHE_SIZE)
	pcore.RoleInfoMap = pcore.NewRoleCache(params.DEFAULT_ROLECACHE_SIZE)
	pcore.AcctInfoMap = pcore.NewAcctCache(params.DEFAULT_ACCOUNTCACHE_SIZE)

	pcore.OrgInfoMap.UpsertOrg(arbitraryOrgAdminRole, "", arbitraryOrgAdminRole, big.NewInt(1), pcore.OrgApproved)
	pcore.RoleInfoMap.UpsertRole(arbitraryOrgAdminRole, "ROLE1", false, false, pcore.Transact, true)
	pcore.RoleInfoMap.UpsertRole(arbitraryOrgAdminRole, "ROLE2", false, false, pcore.ContractDeploy, true)
	pcore.RoleInfoMap.UpsertRole(arbitraryOrgAdminRole, "ROLE3", false, false, pcore.FullAccess, true)
	var Acct3 = common.BytesToAddress([]byte("permission-test1"))
	var Acct4 = common.BytesToAddress([]byte("permission-test2"))

	pcore.AcctInfoMap.UpsertAccount(arbitraryOrgAdminRole, "ROLE1", Acct1, false, pcore.AcctActive)
	pcore.AcctInfoMap.UpsertAccount(arbitraryOrgAdminRole, "ROLE2", Acct2, false, pcore.AcctActive)
	pcore.AcctInfoMap.UpsertAccount(arbitraryOrgAdminRole, "ROLE3", Acct3, false, pcore.AcctActive)

	type args struct {
		address         common.Address
		transactionType pcore.TransactionType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Account with transact permission calling value transfer",
			args:    args{address: Acct1, transactionType: pcore.ValueTransferTxn},
			wantErr: false,
		},
		{
			name:    "Account with transact permission calling value contract call transaction",
			args:    args{address: Acct1, transactionType: pcore.ContractCallTxn},
			wantErr: false,
		},
		{
			name:    "Account with transact permission calling contract deploy",
			args:    args{address: Acct1, transactionType: pcore.ContractDeployTxn},
			wantErr: true,
		},
		{
			name:    "Account with contract permission deploy calling value transfer",
			args:    args{address: Acct2, transactionType: pcore.ValueTransferTxn},
			wantErr: false,
		},
		{
			name:    "Account with contract deploy permission calling value contract call transaction",
			args:    args{address: Acct2, transactionType: pcore.ContractCallTxn},
			wantErr: false,
		},
		{
			name:    "Account with contract deploy permission calling contract deploy",
			args:    args{address: Acct2, transactionType: pcore.ContractDeployTxn},
			wantErr: false,
		},
		{
			name:    "Account with full permission calling value transfer",
			args:    args{address: Acct3, transactionType: pcore.ValueTransferTxn},
			wantErr: false,
		},
		{
			name:    "Account with full permission calling value contract call transaction",
			args:    args{address: Acct3, transactionType: pcore.ContractCallTxn},
			wantErr: false,
		},
		{
			name:    "Account with full permission calling contract deploy",
			args:    args{address: Acct3, transactionType: pcore.ContractDeployTxn},
			wantErr: false,
		},
		{
			name:    "un-permissioned account calling value transfer",
			args:    args{address: Acct4, transactionType: pcore.ValueTransferTxn},
			wantErr: true,
		},
		{
			name:    "un-permissioned account calling contract call transaction",
			args:    args{address: Acct4, transactionType: pcore.ContractCallTxn},
			wantErr: true,
		},
		{
			name:    "un-permissioned account calling contract deploy",
			args:    args{address: Acct4, transactionType: pcore.ContractDeployTxn},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pcore.IsTransactionAllowed(tt.args.address, common.Address{}, nil, nil, nil, nil, tt.args.transactionType); (err != nil) != tt.wantErr {
				t.Errorf("IsTransactionAllowed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

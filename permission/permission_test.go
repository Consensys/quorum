package permission

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	pbind "github.com/ethereum/go-ethereum/permission/bind"
)

func TestPermissionCtrl_InitializeService(t *testing.T) {
	guardianKey, _ := crypto.GenerateKey()
	guardianAddress := crypto.PubkeyToAddress(guardianKey.PublicKey)

	guardianTransactor := bind.NewKeyedTransactor(guardianKey)
	genesisAlloc := map[common.Address]core.GenesisAccount{
		guardianAddress: {
			Balance: big.NewInt(100000000000000),
		},
	}
	sb := backends.NewSimulatedBackend(genesisAlloc, 100000000000000)

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

	sNode, err := node.New(&node.Config{
		P2P: p2p.Config{
			PrivateKey: guardianKey,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	testObject, err := NewQuorumPermissionCtrl(sNode, &types.PermissionConfig{
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

	err = testObject.InitializeService()

	assert.NoError(t, err)
}

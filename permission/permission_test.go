package permission

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	permission "github.com/ethereum/go-ethereum/permission/bind"
	"github.com/stretchr/testify/assert"
)

func TestPermissionCtrl_InitializeService(t *testing.T) {
	key, _ := crypto.GenerateKey()
	senderOpts := bind.NewKeyedTransactor(key)
	genesisAlloc := map[common.Address]core.GenesisAccount{
		senderOpts.From: {
			Balance: big.NewInt(100000000000000),
		},
	}

	sb := backends.NewSimulatedBackend(genesisAlloc, 10000000)

	permUpgrAddress, _, _, err := permission.DeployPermUpgr(senderOpts, sb, senderOpts.From)
	if err != nil {
		t.Fatal(err)
	}

	p, err := permission.NewPermUpgr(permUpgrAddress, sb)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, p)
}

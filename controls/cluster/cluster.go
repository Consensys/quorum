package cluster

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/controls"
	pbind "github.com/ethereum/go-ethereum/controls/bind/cluster"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

type OrgKeyCtrl struct {
	ethClient *ethclient.Client
	key       *ecdsa.PrivateKey
	km        *pbind.Cluster
}

// Creates the controls structure for org key management
func NewOrgKeyCtrl(node *node.Node) (*OrgKeyCtrl, error) {
	stateReader, _, err := controls.CreateEthClient(node)
	if err != nil {
		log.Error("Unable to create ethereum client for cluster check", "err", err)
		return nil, err
	}
	// check if permissioning contract is there at address. If not return from here
	km, err := pbind.NewCluster(params.QuorumPrivateKeyManagementContract, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}
	return &OrgKeyCtrl{stateReader, node.GetNodeKey(), km}, nil
}

// starts the org key management services
func (k *OrgKeyCtrl) Start() error {

	_, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Cluster not enabled for the network", "err", err)
		return nil
	}

	// check if permissioning contract is there at address. If not return from here
	err = k.checkIfContractExists()
	if err != nil {
		return err
	}

	// start the service
	k.manageClusterKeys()
	return nil
}

// checks if the contract is deployed for org key management
func (k *OrgKeyCtrl) checkIfContractExists() error {
	auth := bind.NewKeyedTransactor(k.key)
	clusterSession := &pbind.ClusterSession{
		Contract: k.km,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 4700000,
			GasPrice: big.NewInt(0),
		},
	}

	// dummy call to contrat to check if the contract is deployed
	_, err := clusterSession.CheckOrgContractExists()
	if err != nil {
		return err
	}

	return nil
}

// in case of geth restart firts checks for historical key update events and
// populates the cache, then starts the key change monitoring service
func (k *OrgKeyCtrl) manageClusterKeys() error {
	//call populate nodes to populate the nodes into contract
	if err := k.populatePrivateKeys(); err != nil {
		return err
	}
	//monitor for nodes deletiin via smart contract
	k.monitorKeyChanges()
	return nil

}

// populates cache based on the historical key change events.
func (k *OrgKeyCtrl) populatePrivateKeys() error {
	cluster, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor node delete", "err", err)
		return err
	}

	opts := &bind.FilterOpts{}
	pastAddEvents, err := cluster.FilterOrgKeyAdded(opts)

	if err != nil && err.Error() == "no contract code at given address" {
		return err
	}

	recExists := true
	for recExists {
		recExists = pastAddEvents.Next()
		if recExists {
			types.AddOrgKey(pastAddEvents.Event.OrgId, pastAddEvents.Event.TmKey)
		}
	}

	opts = &bind.FilterOpts{}
	pastDeleteEvents, _ := cluster.FilterOrgKeyDeleted(opts)

	recExists = true
	for recExists {
		recExists = pastDeleteEvents.Next()
		if recExists {
			types.DeleteOrgKey(pastDeleteEvents.Event.OrgId, pastDeleteEvents.Event.TmKey)
		}
	}
	return nil
}

// service to monitor key change events
func (k *OrgKeyCtrl) monitorKeyChanges() {
	go k.monitorKeyAdd()

	go k.monitorKeyDelete()
}

// monitors for new key added event and updates caches based on the same
func (k *OrgKeyCtrl) monitorKeyAdd() {
	cluster, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor Account cluster", "err", err)
	}
	ch := make(chan *pbind.ClusterOrgKeyAdded)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *pbind.ClusterOrgKeyAdded

	_, err = cluster.WatchOrgKeyAdded(opts, ch)
	if err != nil {
		log.Info("Failed WatchOrgKeyDeleted: %v", err)
	}

	for {
		select {
		case newEvent = <-ch:
			types.AddOrgKey(newEvent.OrgId, newEvent.TmKey)
		}
	}
}

// monitors for new key delete event and updates caches based on the same
func (k *OrgKeyCtrl) monitorKeyDelete() {
	cluster, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor Account cluster", "err", err)
	}
	ch := make(chan *pbind.ClusterOrgKeyDeleted)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *pbind.ClusterOrgKeyDeleted

	_, err = cluster.WatchOrgKeyDeleted(opts, ch)
	if err != nil {
		log.Info("Failed WatchOrgKeyDeleted: %v", err)
	}

	for {
		select {
		case newEvent = <-ch:
			types.DeleteOrgKey(newEvent.OrgId, newEvent.TmKey)
		}
	}
}

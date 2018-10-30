package cluster

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/controls"
)

type OrgKeyCtrl struct {
	ethClient *ethclient.Client
}

func NewOrgKeyCtrl(node *node.Node) (*OrgKeyCtrl, error) {
	stateReader, _, err := controls.CreateEthClient(node)
	if err != nil {
		log.Error("Unable to create ethereum client for cluster check : ", "err", err)
		return nil, err
	}
	return &OrgKeyCtrl{stateReader}, nil
}

// This function first adds the node list from permissioned-nodes.json to
// the permissiones contract deployed as a precompile via genesis.json
func (k *OrgKeyCtrl) Start() error {

	_, err := NewClusterFilterer(params.PrivateKeyManagementContract, k.ethClient)
	// check if permissioning contract is there at address. If not return from here
	if err != nil {
		log.Error("Cluster not enabled for the network : ", "err", err)
		return nil
	}
	k.manageClusterKeys()
	return err
}

func (k *OrgKeyCtrl) manageClusterKeys() error {
	//call populate nodes to populate the nodes into contract
	if err := k.populatePrivateKeys(); err != nil {
		return err
	}
	//monitor for nodes deletiin via smart contract
	k.monitorKeyChanges()
	return nil

}

func (k *OrgKeyCtrl) populatePrivateKeys() error {
	cluster, err := NewClusterFilterer(params.PrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor node delete: ", "err", err)
		return err
	}

	opts := &bind.FilterOpts{}
	pastAddEvents, err := cluster.FilterOrgKeyAdded(opts)

	recExists := true
	for recExists {
		recExists = pastAddEvents.Next()
		if recExists {
			types.AddOrgKey(pastAddEvents.Event.OrgId, pastAddEvents.Event.PrivateKey)
		}
	}

	opts = &bind.FilterOpts{}
	pastDeleteEvents, err := cluster.FilterOrgKeyDeleted(opts)

	recExists = true
	for recExists {
		recExists = pastDeleteEvents.Next()
		if recExists {
			types.DeleteOrgKey(pastDeleteEvents.Event.OrgId, pastDeleteEvents.Event.PrivateKey)
		}
	}
	return nil
}

func (k *OrgKeyCtrl) monitorKeyChanges() {
	go k.monitorKeyAdd()

	go k.monitorKeyDelete()
}

func (k *OrgKeyCtrl) monitorKeyAdd() {
	cluster, err := NewClusterFilterer(params.PrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor Account cluster : ", "err", err)
	}
	ch := make(chan *ClusterOrgKeyAdded)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *ClusterOrgKeyAdded

	_, err = cluster.WatchOrgKeyAdded(opts, ch)
	if err != nil {
		log.Info("Failed WatchOrgKeyDeleted: %v", err)
	}

	for {
		select {
		case newEvent = <-ch:
			types.AddOrgKey(newEvent.OrgId, newEvent.PrivateKey)
		}
	}
}

func (k *OrgKeyCtrl) monitorKeyDelete() {
	cluster, err := NewClusterFilterer(params.PrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor Account cluster : ", "err", err)
	}
	ch := make(chan *ClusterOrgKeyDeleted)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *ClusterOrgKeyDeleted

	_, err = cluster.WatchOrgKeyDeleted(opts, ch)
	if err != nil {
		log.Info("Failed WatchOrgKeyDeleted: %v", err)
	}

	for {
		select {
		case newEvent = <-ch:
			types.DeleteOrgKey(newEvent.OrgId, newEvent.PrivateKey)
		}
	}
}

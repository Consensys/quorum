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

func NewOrgKeyCtrl(node *node.Node) (*OrgKeyCtrl, error) {
	stateReader, _, err := controls.CreateEthClient(node)
	if err != nil {
		log.Error("Unable to create ethereum client for cluster check : ", "err", err)
		return nil, err
	}
	// check if permissioning contract is there at address. If not return from here
	km, err := pbind.NewCluster(params.QuorumPrivateKeyManagementContract, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network : ", "err", err)
		return nil, err
	}
	return &OrgKeyCtrl{stateReader, node.GetNodeKey(), km}, nil
}

// This function first adds the node list from permissioned-nodes.json to
// the permissiones contract deployed as a precompile via genesis.json
func (k *OrgKeyCtrl) Start() error {

	_, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	// check if permissioning contract is there at address. If not return from here
	if err != nil {
		log.Error("Cluster not enabled for the network : ", "err", err)
		return nil
	}
	err = k.checkIfContractExists()
	if err != nil {
		return err
	}
	k.manageClusterKeys()
	return nil
}

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

	_, err := clusterSession.CheckOrgContractExists()
	if err != nil {
		return err
	}

	return nil
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
	cluster, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor node delete: ", "err", err)
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

func (k *OrgKeyCtrl) monitorKeyChanges() {
	go k.monitorKeyAdd()

	go k.monitorKeyDelete()
}

func (k *OrgKeyCtrl) monitorKeyAdd() {
	cluster, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor Account cluster : ", "err", err)
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

func (k *OrgKeyCtrl) monitorKeyDelete() {
	cluster, err := pbind.NewClusterFilterer(params.QuorumPrivateKeyManagementContract, k.ethClient)
	if err != nil {
		log.Error("Failed to monitor Account cluster : ", "err", err)
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

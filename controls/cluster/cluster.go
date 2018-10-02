package cluster

import (

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/controls"
	"gopkg.in/urfave/cli.v1"
)
// This function first adds the node list from permissioned-nodes.json to
// the permissiones contract deployed as a precompile via genesis.json
func ManageOrgKeys(ctx *cli.Context, stack *node.Node ) error {
	// Create a new ethclient to for interfacing with the contract
	stateReader, err := controls.CreateEthClient(stack)
	if err != nil {
		log.Error ("Unable to create ethereum client for cluster check : ", "err" , err)
		return err
	}

	// check if permissioning contract is there at address. If not return from here
	if _ , err = NewClusterFilterer(params.PrivateKeyManagementContract, stateReader); err != nil {
		log.Error ("Cluster not enabled for the network : ", "err" , err)
		return nil
	}
	manageClusterKeys(stack, stateReader);

	return err
}

func manageClusterKeys (stack *node.Node, stateReader *ethclient.Client ) error {
	//call populate nodes to populate the nodes into contract
	if err := populatePrivateKeys (stack, stateReader); err != nil {
		return err
	}
	//monitor for nodes deletiin via smart contract
	monitorKeyChanges(stack, stateReader)
	return nil

}

func populatePrivateKeys(stack *node.Node, stateReader *ethclient.Client) error{
	cluster, err := NewClusterFilterer(params.PrivateKeyManagementContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor node delete: ", "err" , err)
		return err
	}

	opts := &bind.FilterOpts{}
	pastAddEvents, err := cluster.FilterOrgKeyAdded(opts)

	recExists := true
	for recExists {
		recExists = pastAddEvents.Next()
		if recExists {
			types.AddOrgKey(pastAddEvents.Event.OrgId, pastAddEvents.Event.PrivateKey )
		}
	}

	opts = &bind.FilterOpts{}
	pastDeleteEvents, err := cluster.FilterOrgKeyDeleted(opts)

	recExists = true
	for recExists {
		recExists = pastDeleteEvents.Next()
		if recExists {
			types.DeleteOrgKey(pastDeleteEvents.Event.OrgId, pastDeleteEvents.Event.PrivateKey )
		}
	}
	return nil
}

func monitorKeyChanges(stack *node.Node, stateReader *ethclient.Client) {
	go monitorKeyAdd(stack, stateReader)

	go monitorKeyDelete(stack, stateReader)
}

func monitorKeyAdd(stack *node.Node, stateReader *ethclient.Client){
	cluster, err := NewClusterFilterer(params.PrivateKeyManagementContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor Account cluster : ", "err" , err)
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

func monitorKeyDelete(stack *node.Node, stateReader *ethclient.Client){
	cluster, err := NewClusterFilterer(params.PrivateKeyManagementContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor Account cluster : ", "err" , err)
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

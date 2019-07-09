package permission

import (
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
)

// Create an RPC client for the contract interface
func CreateEthClient(stack *node.Node) (*ethclient.Client, *eth.Ethereum, error) {
	var e *eth.Ethereum

	if err := stack.Service(&e); err != nil {
		return nil, nil, err
	}

	rpcClient, err := stack.Attach()
	if err != nil {
		return nil, nil, err
	}

	return ethclient.NewClient(rpcClient), e, nil
}

## Build

Just run `make all` and the binaries will be in `build\bin`

Either add that directory to your path or copy the binaries to a location that is already on the path.

## Running

Set up two nodes on a private network based on the instructions at [this link](https://docs.goquorum.consensys.net/tutorials/private-network/create-qbft-network)

Run the nodes with the following command lines:

### Node 1
`geth --datadir data --networkid 1337 --nodiscover --verbosity 5 --syncmode full --istanbul.blockperiod 5 --mine --miner.threads 1 --miner.gasprice 0 --emitcheckpoints --http --http.addr 127.0.0.1 --http.port 22000 --http.corsdomain "*" --http.vhosts "*" --ws --ws.addr 127.0.0.1 --ws.port 32000 --ws.origins "*"     --http.api admin,eth,debug,miner,net,txpool,personal,web3,istanbul --ws.api admin,eth,debug,miner,net,txpool,personal,web3,istanbul --unlock ${ADDRESS} --allow-insecure-unlock --password ./data/keystore/accountPassword --port 30300 --miner.gaslimit 100000000000000000`

### Node 2
`geth --datadir data --networkid 1337 --nodiscover --verbosity 5 --syncmode full --istanbul.blockperiod 5 --mine --miner.threads 1 --miner.gasprice 0 --emitcheckpoints --http --http.addr 127.0.0.1 --http.port 22001 --http.corsdomain "*" --http.vhosts "*" --ws --ws.addr 127.0.0.1 --ws.port 32001 --ws.origins "*" --http.api admin,eth,debug,miner,net,txpool,personal,web3,istanbul --ws.api admin,eth,debug,miner,net,txpool,personal,web3,istanbul --unlock ${ADDRESS} --allow-insecure-unlock --password ./data/keystore/accountPassword --port 30301 --miner.gaslimit 1000000000000000`

## Explanation of Behavior

When there are a quorum of nodes available that are well-behaved, QBFT appears to behave as follows:

* The validators wait for requests to arrive
* When a request is received, the block commit process starts with a pre-prepare message from the proposer
* The other validators receive the pre-prepare message and broadcast their own prepare messages
* Once a quorum of prepare messages has been received, they validators move to the commit phase
* When a quorum of commit messages has been received, the state transitions from prepared to committed
* Once in the committed state, the new block is added
* The validators go back to waiting for requests

When things are in a bad state, such as when there are no longer enough nodes for a quorum:

* The round change timer expires.
* The round increases and the timeout grows.
* While the number of nodes is under the minimum needed for consensus, no blocks can be validated/committed
* With our simple two-node network, both nodes are needed for a quorum.
* When the number of nodes reaches the threshold required for a quorum, block processing can continue.
* The round number will reset and the round timeouts will fall back to the lowest setting (10 seconds)

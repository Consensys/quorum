#!/bin/bash

#
# This is used at Container start up to run the constellation and geth nodes
#

set -e

### Configuration Options
TMCONF=/qdata/constellation/tm.conf

COMMON_ARGS="--datadir /qdata/ethereum --permissioned --rpc --rpcport 8545 --rpcaddr 0.0.0.0 --ws --wsport 8546 --wsaddr 0.0.0.0 --unlock 0 --password /qdata/ethereum/passwords.txt --verbosity 4 --bootnodes"
COMMON_APIS="admin,db,eth,debug,miner,net,shh,txpool,personal,web3"
RAFT_APIS="$COMMON_APIS,raft"
RAFT_ARGS="--raft --rpcapi $RAFT_APIS --wsapi $RAFT_APIS"
IBFT_APIS="$COMMON_APIS,istanbul"
IBFT_ARGS="--syncmode full --mine --rpcapi $IBFT_APIS --wsapi $IBFT_APIS"

wsOrigins="*"

###
### These are the arguments supported:
### bootnode=<enode> - argument is the enode URI of the bootnode
### raftInit - to indicate that this node is part of the initial raft quorum/cluster
### raftID=<number> - to indicate that this node is joining an existing quorum/cluster
### ibft - to indicate using Istanbul BFT as the consensus algorithm, instead of Raft
###
### raftInit and raftID are mutually exclusive
###
### if ibft is specified, both raftInit and raftID are ignored
###
### If the bootnode argument is omitted, the program enters a sleep loop until a file
### "bootnode.info" is found.
###
### If the raft* argument is omitted, the program assumes this is joining an existing
### cluster, and enters a sleep loop until a file "raft.id" is found.
###
while [ "$1" != "" ]; do
    PARAM=`echo $1 | awk -F= '{print $1}'`
    VALUE=`echo $1 | awk -F= '{print $2}'`
    case $PARAM in
        --ibft)
            ibft=YES
            ;;
        --raftInit)
            raftInit=YES
            ;;
        --raftID)
            raftID=$VALUE
            ;;
        --bootnode)
            bootnode=$VALUE
            ;;
        --wsorigins)
            wsOrigins=$VALUE
            ;;
        *)
            echo "ERROR: unknown parameter \"$PARAM\""
            exit 1
            ;;
    esac
    shift
done

echo "bootnode URI          = $bootnode"
echo "initial Raft cluster? = $raftInit"
echo "Raft ID               = $raftID"
echo "Istanbul BFT          = $ibft"
echo "WS Origins            = $wsOrigins"

#
# since the bootnode is required, do not proceed until
# it's found either in the argument or in the bootnode.info file
#
if [ "$bootnode" == "" ]; then
  while [ ! -f /qdata/ethereum/bootnode.info ]
  do
    sleep 2
  done

  bootnode=`cat /qdata/ethereum/bootnode.info`
fi

#
# now decide whether to use --raftjoinexisting parameter
#
if [ "$ibft" == YES ]; then
  # using IBFT consensus
  GETH_ARGS="$COMMON_ARGS $bootnode $IBFT_ARGS"
elif [ "$raftInit" == YES ]; then
  # initial Raft cluster
  GETH_ARGS="$COMMON_ARGS $bootnode $RAFT_ARGS"
else
  if [ "$raftID" == "" ]; then
    # joining and existing Raft cluster, a raft id is required
    while [ ! -f /qdata/ethereum/raft.id ]
    do
      sleep 2
    done

    raftID=`cat /qdata/ethereum/raft.id`
  fi

  GETH_ARGS="$COMMON_ARGS $bootnode $RAFT_ARGS --raftjoinexisting $raftID"
fi

GETH_ARGS="$GETH_ARGS --wsorigins=$wsOrigins"

#
# the geth node should not start until constellation started and
# generated the shared IPC
#
while [ ! -f $TMCONF ]
do
  sleep 2
done

# sleep an additional 2 seconds to give the constellation node a chance to start
sleep 2

#
# ALL SET!
#
if [ ! -d /qdata/ethereum/geth/chaindata ]; then
  echo "[*] Mining Genesis block"
  geth --datadir /qdata/ethereum init /qdata/ethereum/genesis.json
fi

echo "[*] Starting node with args $GETH_ARGS"
PRIVATE_CONFIG=$TMCONF nohup geth $GETH_ARGS 2>>/qdata/logs/geth.log

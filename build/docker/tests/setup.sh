#!/bin/bash

#
# Create all the necessary scripts, keys, configurations etc. to run
# a cluster of N Quorum nodes with Raft consensus.
#
# The nodes will be in Docker containers. List the IP addresses that
# they will run at below (arbitrary addresses are fine).
#
# Run the cluster with "docker-compose up -d"
#
# Run a console on Node N with "geth attach qdata_N/ethereum/geth.ipc"
# (assumes Geth is installed on the host.)
#
# Geth and Constellation logfiles for Node N will be in qdata_N/logs/
#

# TODO: check file access permissions, especially for keys.


#### Configuration options #############################################

# One Docker container will be configured for each IP address in $ips
subnet="172.13.0.0/16"
ips=("172.13.0.3" "172.13.0.5" "172.13.0.7" "172.13.0.9" "172.13.0.11")

cips=()
## constellation node uses IP address derived from
## the corresponding geth node IP address by subtracting 1
for ip in ${ips[*]}
do
    [[ $ip =~ (.*\.)([0-9]+)$ ]] && cip=${BASH_REMATCH[2]}
    cips+=("${BASH_REMATCH[1]}$(($cip - 1))")
done

# Docker image name
image_constellation=jpmorganchase/constellation
image_quorum=jpmorganchase/quorum

########################################################################

nnodes=${#ips[@]}

if [[ $nnodes < 2 ]]
then
    echo "ERROR: There must be more than one node IP address."
    exit 1
fi

./scripts/cleanup.sh

cd tmp
pwd=`pwd`

#### Create directories for each node's configuration ##################

echo '[1] Configuring for '$nnodes' nodes.'

n=1
for ip in ${ips[*]}
do
    qd=qdata_$n
    mkdir -p $qd/{logs,constellation}
    mkdir -p $qd/ethereum/geth

    let n++
done


#### Make static-nodes.json and store keys #############################

echo '[2] Creating Enodes and static-nodes.json.'

echo "[" > static-nodes.json
n=1
for ip in ${ips[*]}
do
    qd=qdata_$n

    # Generate the node's Enode and key
    bootnode_cmd="docker run -it -v $pwd/$qd:/qdata $image_quorum /usr/local/bin/bootnode"
    $bootnode_cmd -genkey /qdata/ethereum/nodekey
    enode=`$bootnode_cmd -nodekey /qdata/ethereum/nodekey -writeaddress | tr -d '[:space:]'`
    echo "Node $n id: $enode"

    # Add the enode to static-nodes.json
    sep=`[[ $n < $nnodes ]] && echo ","`
    echo '  "enode://'$enode'@'$ip':30303?discport=0&raftport=50400"'$sep >> static-nodes.json

    let n++
done
echo "]" >> static-nodes.json


#### Create accounts, keys and genesis.json file #######################

echo '[3] Creating Ether accounts and genesis.json.'

cat > genesis.json <<EOF
{
  "alloc": {
EOF

n=1
for ip in ${ips[*]}
do
    qd=qdata_$n

    # Generate an Ether account for the node
    touch $qd/ethereum/passwords.txt
    create_account="docker run -v $pwd/$qd:/qdata $image_quorum /usr/local/bin/geth --datadir=/qdata/ethereum --password /qdata/ethereum/passwords.txt account new"
    account1=`$create_account | cut -c 11-50`
    account2=`$create_account | cut -c 11-50`
    account3=`$create_account | cut -c 11-50`
    account4=`$create_account | cut -c 11-50`
    echo "Accounts for node $n: $account1 $account2 $account3 $account4"

    # Add the account to the genesis block so it has some Ether at start-up
    sep=`[[ $n < $nnodes ]] && echo ","`
    cat >> genesis.json <<EOF
    "${account1}": {
      "balance": "1000000000000000000000000000"
    }, "${account2}": {
      "balance": "1000000000000000000000000000"
    }, "${account3}": {
      "balance": "1000000000000000000000000000"
    }, "${account4}": {
      "balance": "1000000000000000000000000000"
    }${sep}
EOF

    let n++
done

cat >> genesis.json <<EOF
  },
  "coinbase": "0x0000000000000000000000000000000000000000",
  "config": {
    "homesteadBlock": 0
  },
  "difficulty": "0x0",
  "extraData": "0x",
  "gasLimit": "0x2FEFD800",
  "mixhash": "0x00000000000000000000000000000000000000647572616c65787365646c6578",
  "nonce": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
EOF


#### Make node list for tm.conf ########################################

nodelist=
n=1
for ip in ${cips[*]}
do
    sep=`[[ $ip != ${cips[0]} ]] && echo ","`
    nodelist=${nodelist}${sep}'"http://'${ip}':9000/"'
    let n++
done


#### Complete each node's configuration ################################

echo '[4] Creating Quorum keys and finishing configuration.'

n=1
for ip in ${cips[*]}
do
    qd=qdata_$n

    cat ../templates/tm.conf \
        | sed s/_NODEIP_/${cips[$((n-1))]}/g \
        | sed s%_NODELIST_%$nodelist%g \
              > $qd/constellation/tm.conf

    cp genesis.json $qd/ethereum/genesis.json
    cp static-nodes.json $qd/ethereum/static-nodes.json
    cp static-nodes.json $qd/ethereum/permissioned-nodes.json

    # Generate Quorum-related keys (used by Constellation)
    docker run -v $pwd/$qd:/qdata -v $pwd/../scripts:/scripts $image_constellation /scripts/generate-keys.sh
    echo 'Node '$n' public key: '`cat $qd/constellation/tm.pub`

    let n++
done
rm -rf genesis.json static-nodes.json


#### Create the docker-compose file ####################################

cat > docker-compose.yml <<EOF
version: '2'
services:
EOF

for index in ${!ips[*]}; do 
    n=$((index+1))
    qd=qdata_$n
    ip=${ips[index]}; 
    cip=${cips[index]}; 

    cat >> docker-compose.yml <<EOF
  constellation_$n:
    container_name: constellation_$n
    image: $image_constellation
    volumes:
      - './$qd:/qdata'
    networks:
      quorum_net:
        ipv4_address: '$cip'
    ipc: shareable

  node_$n:
    container_name: node_$n
    image: $image_quorum
    volumes:
      - './$qd:/qdata'
    networks:
      quorum_net:
        ipv4_address: '$ip'
    ports:
      - $((n+22000)):8545
    depends_on:
      - constellation_$n
    ipc: container:constellation_$n

EOF
done

cat >> docker-compose.yml <<EOF

networks:
  quorum_net:
    driver: bridge
    ipam:
      driver: default
      config:
      - subnet: $subnet
EOF


#### Create pre-populated contracts ####################################

# Private contract - insert Node 2 as the recipient
cat ../templates/contract_pri.js \
    | sed s:_NODEKEY_:`cat qdata_2/constellation/tm.pub`:g \
          > contract_pri.js

# Public contract - no change required
cp ../templates/contract_pub.js ./

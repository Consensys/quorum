#!/usr/bin/sh

ROOT=/tmp/gemini
GEMINI_REPO=$GOPATH/src/github.com/ethereum/go-ethereum
GEMINI_BIN=$GEMINI_REPO/build/bin/geth
NETID=87234

BOOTNODE_BIN=$GEMINI_REPO/build/bin/bootnode
BOOTNODE_KEYHEX=77bd02ffa26e3fb8f324bda24ae588066f1873d95680104de5bc2db9e7b2e510
BOOTNODE_ENODE=enode://6433e8fb82c4638a8a6d499d40eb7d8158883219600bfd49acb968e3a37ccced04c964fa87b3a78a2da1b71dc1b90275f4d055720bb67fad4a118a56925125dc@[127.0.0.1]:33445

GLOBAL_ARGS="--bootnodes $BOOTNODE_ENODE --networkid $NETID --gasprice \"0\""

mkdir -p $ROOT/logs

echo "[*] Start bootnode"
nohup $BOOTNODE_BIN --nodekeyhex "$BOOTNODE_KEYHEX" --addr="127.0.0.1:33445" 2>>$ROOT/logs/bootnode.log &
echo "wait for bootnode to start..."
sleep 5

echo "[*] Start node 1"
nohup $GEMINI_BIN --datadir $ROOT/dd1 $GLOABL_ARGS --blockmakeraccount "0xed9d02e382b34818e88b88a309c7fe71e65f419d" --blockmakerpassword "" --voteaccount "0xed9d02e382b34818e88b88a309c7fe71e65f419d" --votepassword "" --port 21000 --networkid $NETID 2>>$ROOT/logs/1.log &

echo "[*] Start node 2"
nohup $GEMINI_BIN --datadir $ROOT/dd2 $GLOBAL_ARGS --blockmakeraccount "0xca843569e3427144cead5e4d5999a3d0ccf92b8e" --blockmakerpassword "" --port 21001 2>>$ROOT/logs/2.log &

echo "[*] Start node 3"
nohup $GEMINI_BIN --datadir $ROOT/dd3 $GLOBAL_ARGS --voteaccount "0x0fbdc686b912d7722dc86510934589e0aaf3b55a" --votepassword "" --port 21002 2>>$ROOT/logs/3.log &

echo "[*] Start node 4"
nohup $GEMINI_BIN --datadir $ROOT/dd4 $GLOBAL_ARGS --voteaccount "0x9186eb3d20cbd1f5f992a950d808c4495153abd5" --votepassword "" --port 21003 2>>$ROOT/logs/4.log &

echo "[*] Start node 5"
nohup $GEMINI_BIN --datadir $ROOT/dd5 $GLOBAL_ARGS --voteaccount "0x0638e1574728b6d862dd5d3a3e0942c3be47d996" --votepassword "" --port 21004 2>>$ROOT/logs/5.log &

echo "[*] Start node 6"
nohup $GEMINI_BIN --datadir $ROOT/dd6 $GLOBAL_ARGS --port 21005 2>>$ROOT/logs/6.log &

echo "[*] Start node 7"
nohup $GEMINI_BIN --datadir $ROOT/dd7 $GLOBAL_ARGS --port 21006 2>>$ROOT/logs/7.log &

#!/usr/bin/sh

ROOT=/tmp/gemini
GEMINI_REPO=$GOPATH/src/github.com/ethereum/go-ethereum
GEMINI_BIN=$GEMINI_REPO/build/bin/geth

echo "[*] Cleanup existing data directories"
rm -rf /tmp/gemini/dd*

echo "[*] Configure node 1 as voter and block maker"
mkdir -p $ROOT/dd1/keystore
cp $GEMINI_REPO/_data/keys/UTC--2016-11-02T08-55-33.544599174Z--ed9d02e382b34818e88b88a309c7fe71e65f419d $ROOT/dd1/keystore
$GEMINI_BIN --datadir $ROOT/dd1 init $GEMINI_REPO/_data/genesis.json

echo "[*] Configure node 2 as block maker"
mkdir -p $ROOT/dd2/keystore
cp $GEMINI_REPO/_data/keys/UTC--2016-11-02T08-55-36.695601929Z--ca843569e3427144cead5e4d5999a3d0ccf92b8e $ROOT/dd2/keystore
$GEMINI_BIN --datadir $ROOT/dd2 init $GEMINI_REPO/_data/genesis.json

echo "[*] Configure node 3 as voter"
mkdir -p $ROOT/dd3/keystore
cp $GEMINI_REPO/_data/keys/UTC--2016-11-02T08-55-39.164648792Z--0fbdc686b912d7722dc86510934589e0aaf3b55a $ROOT/dd3/keystore
$GEMINI_BIN --datadir $ROOT/dd3 init $GEMINI_REPO/_data/genesis.json

echo "[*] Configure node 4 as voter"
mkdir -p $ROOT/dd4/keystore
cp $GEMINI_REPO/_data/keys/UTC--2016-11-02T08-56-07.802508523Z--9186eb3d20cbd1f5f992a950d808c4495153abd5 $ROOT/dd4/keystore
$GEMINI_BIN --datadir $ROOT/dd4 init $GEMINI_REPO/_data/genesis.json

echo "[*] Configure node 5 as voter"
mkdir -p $ROOT/dd5/keystore
cp $GEMINI_REPO/_data/keys/UTC--2016-11-02T09-05-09.535511997Z--0638e1574728b6d862dd5d3a3e0942c3be47d996 $ROOT/dd5/keystore
$GEMINI_BIN --datadir $ROOT/dd5 init $GEMINI_REPO/_data/genesis.json

echo "[*] Configure node 6"
mkdir -p $ROOT/dd6/keystore
$GEMINI_BIN --datadir $ROOT/dd6 init $GEMINI_REPO/_data/genesis.json

echo "[*] Configure node 7"
mkdir -p $ROOT/dd7/keystore
$GEMINI_BIN --datadir $ROOT/dd7 init $GEMINI_REPO/_data/genesis.json

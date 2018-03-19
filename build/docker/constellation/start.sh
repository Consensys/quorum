#!/bin/bash

#
# This is used at Container start up to run the constellation and geth nodes
#

set -u
set -e

### Configuration Options
TMCONF=/qdata/tm.conf

echo "[*] Starting Constellation node"
nohup /usr/local/bin/constellation-node $TMCONF -v3

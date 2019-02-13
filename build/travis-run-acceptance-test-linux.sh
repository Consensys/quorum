#!/bin/bash
# start network and run acceptance tests
set -e
echo "start quorum network for consensus ${TF_VAR_consensus_mechanism} ..."
cd $TRAVIS_HOME/quorum-cloud/travis/4nodes
./init.sh ${TF_VAR_consensus_mechanism}
./start.sh ${TF_VAR_consensus_mechanism} tessera
echo "network started"
set -e
cd $TRAVIS_HOME/quorum-acceptance-tests
cp config/application-local.4nodes.yml config/application-local.yml
echo "running acceptance test for consensus $consensus ..."
./src/travis/run_tests.sh
echo "acceptance test finished"
echo "stop the network..."
$TRAVIS_HOME/quorum-cloud/travis/4nodes/stop.sh
echo "network stopped"

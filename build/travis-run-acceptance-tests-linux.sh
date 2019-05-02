#!/bin/bash
set -e
# start network and run acceptance tests
echo "---> start quorum network for consensus ${TF_VAR_consensus_mechanism} ..."
export PATH=${TRAVIS_BUILD_DIR}/build/bin:$PATH
export TESSERA_JAR=${HOME}/tessera.jar
cd ${TRAVIS_HOME}/quorum-cloud/travis/4nodes
./init.sh ${TF_VAR_consensus_mechanism}
./start.sh ${TF_VAR_consensus_mechanism} tessera
echo "---> network started"
cd ${TRAVIS_HOME}/quorum-acceptance-tests
cp config/application-local.4nodes.yml config/application-local.yml
echo "---> run acceptance tests for consensus ${TF_VAR_consensus_mechanism} ..."
./src/travis/run_tests.sh
echo "---> acceptance tests finished"
echo "---> stop the network..."
${TRAVIS_HOME}/quorum-cloud/travis/4nodes/stop.sh
echo "---> network stopped"

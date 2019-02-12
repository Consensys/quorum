#!/bin/bash
set -e
cd $TRAVIS_HOME/quorum-acceptance-tests
cp config/application-local.4nodes.yml config/application-local.yml
echo "running acceptance test for consensus $consensus ..."
./src/travis/run_tests.sh
echo "stop the network..."
$TRAVIS_HOME/quorum-cloud/travis/4nodes/stop.sh
echo "acceptance test finished"
#!/bin/bash
echo "install started..."
set -e
#git clone https://github.com/jpmorganchase/quorum-acceptance-tests.git $TRAVIS_HOME/quorum-acceptance-tests
echo "cloning quorum-acceptance-test finished"
#git clone https://github.com/jpmorganchase/quorum-cloud.git $TRAVIS_HOME/quorum-cloud
git clone --branch=quorum-travis-ci-accept-test-integ https://github.com/amalrajmani/quorum-acceptance-tests.git $TRAVIS_HOME/quorum-acceptance-tests
git clone --branch=travis-ci-integ https://github.com/QuorumEngineering/quorum-cloud.git $TRAVIS_HOME/quorum-cloud
echo "cloning quorum-cloud finished"
sudo chmod 755 $TRAVIS_HOME/quorum-cloud/travis/start-network-linux.sh
sudo apt update
sudo apt -y install dpkg
echo "installing jre 8.."
sudo apt -y install openjdk-8-jre-headless
java -version
echo "installing maven.."
sudo apt -y install maven
mvn --version
sudo apt-get -y install software-properties-common
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt update
echo "installing solidity.."
sudo apt-get -y install solc
solc --version
echo "getting tessera jar..."
wget https://github.com/jpmorganchase/tessera/releases/download/tessera-0.8/tessera-app-0.8-app.jar -O tessera.jar -q
sudo cp tessera.jar $HOME
sudo chmod 755 $HOME/tessera.jar
echo "getting gauge jar..."
wget https://github.com/getgauge/gauge/releases/download/v1.0.3/gauge-1.0.3-linux.x86_64.zip -O gauge.zip -q
sudo unzip -o gauge.zip -d /usr/local/bin
echo "installing gauge..."
cd $TRAVIS_HOME/quorum-acceptance-tests
gauge telemetry off
gauge install
echo "install done"
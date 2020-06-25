#!/bin/bash
set -e
# install geth and dependencies for acceptance tests
echo "---> install started ..."
echo "---> installing tools ..."
sudo apt-get update
# upgrade dpkg to fix issue with trusty: dpkg-deb: error
sudo apt-get -y install dpkg
java -version
mvn --version

sudo wget https://github.com/ethereum/solidity/releases/download/v0.5.4/solc-static-linux -O /usr/local/bin/solc -q
sudo chmod +x /usr/local/bin/solc
solc --version
echo "---> tools installation done"

echo "---> building geth ..."
sudo modprobe fuse
sudo chmod 666 /dev/fuse
sudo chown root:${USER} /etc/fuse.conf
go run build/ci.go install
echo "---> building geth done"

echo "---> cloning quorum-cloud and quorum-acceptance-tests ..."
git clone https://github.com/jpmorganchase/quorum-acceptance-tests.git ${TRAVIS_HOME}/quorum-acceptance-tests
git clone https://github.com/jpmorganchase/quorum-cloud.git ${TRAVIS_HOME}/quorum-cloud

echo "---> cloning done"

echo "---> getting tessera jar ..."
wget https://oss.sonatype.org/service/local/repositories/releases/content/com/jpmorgan/quorum/tessera-app/0.10.4/tessera-app-0.10.4-app.jar -O $HOME/tessera.jar -q
echo "---> tessera done"

echo "---> getting gauge jar ..."
wget https://github.com/getgauge/gauge/releases/download/v1.0.8/gauge-1.0.8-linux.x86_64.zip -O gauge.zip -q
sudo unzip -o gauge.zip -d /usr/local/bin
gauge telemetry off
cd ${TRAVIS_HOME}/quorum-acceptance-tests
gauge install
echo "---> gauge installation done"

echo "---> install done"
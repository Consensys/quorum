#!/bin/bash
set -e
# install geth and dependencies for acceptance tests
echo "---> install started ..."
echo "---> installing tools ..."
sudo apt-get update
# upgrade dpkg to fix issue with trusty: dpkg-deb: error
sudo apt-get -y install dpkg
# Travis pre-installs jdk11 by default.
# However, Tessera 0.8 requires jre8 to run so we use jdk_switcher utility from Travis
if test -f ${HOME}/.jdk_switcher_rc; then
    . ${HOME}/.jdk_switcher_rc
fi
if test -f /opt/jdk_switcher/jdk_switcher.sh; then
    . /opt/jdk_switcher/jdk_switcher.sh
fi
jdk_switcher use openjdk8
java -version
mvn --version

sudo wget https://github.com/ethereum/solidity/releases/download/v0.5.4/solc-static-linux -O /usr/local/bin/solc
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
wget https://github.com/jpmorganchase/tessera/releases/download/tessera-0.8/tessera-app-0.8-app.jar -O $HOME/tessera.jar -q
echo "---> tessera done"

echo "---> getting gauge jar ..."
wget https://github.com/getgauge/gauge/releases/download/v1.0.4/gauge-1.0.4-linux.x86_64.zip -O gauge.zip -q
sudo unzip -o gauge.zip -d /usr/local/bin
gauge telemetry off
cd ${TRAVIS_HOME}/quorum-acceptance-tests
gauge install
echo "---> gauge installation done"

echo "---> install done"
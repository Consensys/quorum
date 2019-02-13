#!/bin/bash
# install geth and dependencies for acceptance test
echo "install started..."
set -e
echo "building geth..."
sudo modprobe fuse
sudo chmod 666 /dev/fuse
sudo chown root:$USER /etc/fuse.conf
go run build/ci.go install
export PATH=$TRAVIS_BUILD_DIR/build/bin:$PATH
echo "building geth finished"

#git clone https://github.com/jpmorganchase/quorum-acceptance-tests.git $TRAVIS_HOME/quorum-acceptance-tests
git clone --branch=quorum-travis-ci-accept-test-integ https://github.com/amalrajmani/quorum-acceptance-tests.git $TRAVIS_HOME/quorum-acceptance-tests
echo "cloning quorum-acceptance-test finished"

git clone https://github.com/jpmorganchase/quorum-cloud.git $TRAVIS_HOME/quorum-cloud
echo "cloning quorum-cloud finished"

sudo apt update
sudo apt -y install dpkg
echo "installing jre 8.."
sudo apt -y install openjdk-8-jre-headless
java -version
echo "jre 8 installation done"
echo "installing maven.."
sudo apt -y install maven
mvn --version
echo "maven installation done"

sudo apt-get -y install software-properties-common
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt update

echo "installing solidity.."
sudo apt-get -y install solc
solc --version
echo "solidity installation done"

echo "getting tessera jar..."
wget https://github.com/jpmorganchase/tessera/releases/download/tessera-0.8/tessera-app-0.8-app.jar -O tessera.jar -q
sudo cp tessera.jar $HOME
sudo chmod 755 $HOME/tessera.jar
echo "tessera done"

echo "getting gauge jar..."
wget https://github.com/getgauge/gauge/releases/download/v1.0.3/gauge-1.0.3-linux.x86_64.zip -O gauge.zip -q
echo "gauge done"

sudo unzip -o gauge.zip -d /usr/local/bin

echo "installing gauge..."
cd $TRAVIS_HOME/quorum-acceptance-tests
gauge telemetry off
gauge install
echo "gauge installation done"

echo "install done"
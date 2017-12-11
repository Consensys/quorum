#!/bin/bash

RELEASE=0.2.0
ZIPNAME=constellation-$RELEASE-ubuntu1604
ZIPFILE=$ZIPNAME.tar.xz

wget https://github.com/jpmorganchase/constellation/releases/download/v$RELEASE/$ZIPFILE
tar -xvf $ZIPFILE
cp $ZIPNAME/constellation-node /work/build/bin
chmod 0755 /work/build/bin/constellation-node
rm -rf $ZIPFILE $ZIPNAME
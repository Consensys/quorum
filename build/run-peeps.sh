#!/bin/sh

if ! git clone https://github.com/ConsenSys/PEEPS.git
then
  cd PEEPS
  git reset --hard HEAD
  git pull origin master
else
  cd PEEPS
fi

./gradlew --no-daemon --parallel endToEndTest


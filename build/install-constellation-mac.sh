#!/usr/bin/env bash

echo "Going into $HOME"
cd $HOME

echo "Cloning Constellation repo"
git clone https://github.com/jpmorganchase/constellation.git

cd constellation

echo "Running stack setup in $(pwd)"
stack setup

echo "Now installing Constellation from $(pwd)"
stack install
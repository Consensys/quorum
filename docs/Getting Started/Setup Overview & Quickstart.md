# Setup Overview & Quickstart

Using Quorum requires that a Quorum Node and a Constellation/Tessera Node are installed, configured and
running (see build/installation instructions for both below). An overview of the steps to follow to manually set up Quorum, including key generation, genesis block & Constellation/Tessera configuration will be available soon, but for now the best way to get started is to use the Vagrant environment that has been made available for running the [Quorum Examples](../Quorum-Examples). The Vagrant environment automatically sets up a test Quorum network that is ready for development use within minutes and is the recommended approach if you are looking to get started with Quorum.  If you don't want to use the Quorum Examples approach and instead would like to manually set up Quorum then please see below (Note: this documentation is Work In Progress)

## Building Quorum Node From Source

Clone the repository and build the source:

```
git clone https://github.com/jpmorganchase/quorum.git
cd quorum
make all
```

Binaries are placed in `$REPO_ROOT/build/bin`. Put that folder in your PATH to make `geth` and `bootnode` easily invokable, or copy those binaries to a folder already in PATH, e.g. `/usr/local/bin`.

An easy way to supplement PATH is to add `PATH=$PATH:/path/to/repository/build/bin` to your `~/.bashrc` or `~/.bash_aliases` file.

Run the tests:

```
make test
```

## Installing Constellation
Grab a package for your platform [here](https://github.com/jpmorganchase/constellation/releases), and place the extracted binaries somewhere in PATH, e.g. /usr/local/bin.

## Installing Tessera
Follow the installation instructions on the [Tessera project page](https://github.com/jpmorganchase/tessera).

## Getting Started from Scratch
Follow the instructions given [here](../Getting-Started-From-Scratch).

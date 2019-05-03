## Installation

### Prerequisites

  1. Install supporting libraries:
    - Ubuntu: `apt-get install libdb-dev libleveldb-dev libsodium-dev zlib1g-dev libtinfo-dev`
    - Red Hat: `dnf install libdb-devel leveldb-devel libsodium-devel zlib-devel ncurses-devel`
    - MacOS: `brew install berkeley-db leveldb libsodium`

### Downloading precompiled binaries

Constellation binaries for most major platforms can be downloaded [here](https://github.com/jpmorganchase/constellation/releases).

### Installation from source

  1. First time only: Install Stack:
    - Linux: `curl -sSL https://get.haskellstack.org/ | sh`
    - MacOS: `brew install haskell-stack`

  2. First time only: run `stack setup` to install GHC, the Glasgow
     Haskell Compiler

  3. Run `stack install`

## Generating keys

  1. To generate a key pair "node", run `constellation-node --generatekeys=node`

  If you choose to lock the keys with a password, they will be encrypted using
  a master key derived from the password using Argon2id. This is designed to be
  a very expensive operation to deter password cracking efforts. When
  constellation encounters a locked key, it will prompt for a password after
  which the decrypted key will live in memory until the process ends.

## Running

  1. Run `constellation-node <path to config file>` or specify configuration
     variables as command-line options (see `constellation-node --help`)

Please refer to the [Constellation client Go library](../constellation-go)
for an example of how to use Constellation. 

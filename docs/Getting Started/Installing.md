# Installing

Quorum and its Privacy Managers can be installed and used as Docker containers, by building from source, or by downloading pre-built release binaries.

## As containers
Docker containers exist for Quorum, Tessera and Constellation and can be found at the [`quorumengineering` Docker repository](https://hub.docker.com/u/quorumengineering/):

```
docker pull quorumengineering/quorum
docker pull quorumengineering/tessera
docker pull quorumengineering/constellation
```    

## From source
### Quorum

1. Clone the repository and build the source:

    ```bash
    git clone https://github.com/jpmorganchase/quorum.git
    cd quorum
    make all
    ```

    Binaries are placed in `$REPO_ROOT/build/bin`. Add that folder to `PATH` to make `geth` and `bootnode` easily invokable, or copy those binaries to a folder already in `PATH`, e.g. `/usr/local/bin`.

    An easy way to supplement `PATH` is to add `PATH=$PATH:/path/to/repository/build/bin` to your `~/.bashrc` or `~/.bash_aliases` file.

1. Run the tests:

    ```bash
    make test
    ```

### Privacy Managers

#### Tessera
Tessera is a production-ready implementation of Quorum's privacy manager.  It is undergoing active development with new features being added regularly.

Follow the installation instructions on the [Tessera project page](https://github.com/jpmorganchase/tessera).

#### Constellation
Constellation is the reference implementation of Quorum's privacy manager.  It is still supported but no longer undergoing active development of new features.  

Grab a package for your platform [here](https://github.com/jpmorganchase/constellation/releases), and place the extracted binaries somewhere in `PATH`, e.g. `/usr/local/bin`.

## As release binaries
The pre-compiled release binaries for Quorum, Tessera and Constellation can be downloaded from the following links:

* [Quorum](https://github.com/jpmorganchase/quorum/releases)
* [Tessera](https://github.com/jpmorganchase/tessera/releases)
* [Constellation](https://github.com/jpmorganchase/constellation/releases)

Once downloaded, add the binaries to `PATH` to make them easily invokable.

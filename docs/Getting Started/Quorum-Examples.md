# Quorum Examples

This section details couple of setup [examples for Quorum](https://github.com/jpmorganchase/quorum-examples.git)

Current examples include:

* [7nodes](../7Nodes): Starts up a fully-functioning Quorum environment consisting of 7 independent nodes. From this example one can test consensus, privacy, and all the expected functionality of an Ethereum platform.
* [5nodesRTGS](https://github.com/bacen/quorum-examples/tree/master/examples/5nodesRTGS): [__Note__: This links to an external repo which you will need to clone, thanks to @rsarres for this contribution!] Starts up a set of 5 nodes that simulates a Real-time Gross Setlement environment with 3 banks, one regulator (typically a central bank) and an observer that cannot access the private data.

The easiest way to get started with running the examples is to use the vagrant environment (see below).

**Important note**: Any account/encryption keys contained in this repository are for
demonstration and testing purposes only. Before running a real environment, you should
generate new ones using Geth's `account` tool and the `--generate-keys` option for Constellation (or `-keygen` option for Tessera).

## Getting Started
The 7nodes example can be run in three ways:

1. By running a preconfigured Vagrant environment which comes complete with Quorum, Constellation, Tessera and the 7nodes example (__works on any machine__).
1. By running [`docker-compose`](https://docs.docker.com/compose/) against a preconfigured `compose` file ([example](https://github.com/jpmorganchase/quorum-examples/blob/master/docker-compose.yml) from the `quorum-examples` repo) which starts 7nodes example (tested on Windows 10, macOS Mojave & Ubuntu 18.04).
1. By downloading and locally running Quorum, Tessera and the examples (__requires an Ubuntu-based/macOS machine; note that Constellation does not support running locally__)

### Setting up Vagrant
1. Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
2. Install [Vagrant](https://www.vagrantup.com/downloads.html)
3. Download and start the Vagrant instance (note: running `vagrant up` takes approx 5 mins):

    ```sh
    git clone https://github.com/jpmorganchase/quorum-examples
    cd quorum-examples
    vagrant up
    vagrant ssh
    ```

4. To shutdown the Vagrant instance, run `vagrant suspend`. To delete it, run
   `vagrant destroy`. To start from scratch, run `vagrant up` after destroying the
   instance.

#### Troubleshooting Vagrant
* If you are behind a proxy server, please see https://github.com/jpmorganchase/quorum/issues/23.
* If you are using macOS and get an error saying that the ubuntu/xenial64 image doesn't
exist, please run `sudo rm -r /opt/vagrant/embedded/bin/curl`. This is usually due to
issues with the version of curl bundled with Vagrant.
* If you receive the error `default: cp: cannot open '/path/to/geth.ipc' for reading: Operation not supported` after running `vagrant up`, run `./raft-init.sh` within the 7nodes directory on your local machine.  This will remove temporary files created after running 7nodes locally and will enable `vagrant up` to execute correctly.  

#### Troubleshooting Vagrant: Memory usage
* The Vagrant instance is allocated 6 GB of memory.  This is defined in the `Vagrantfile`, `v.memory = 6144`.  This has been deemed a suitable value to allow the VM and examples to run as expected.  The memory allocation can be changed by updating this value and running `vagrant reload` to apply the change.

* If the machine you are using has less than 8 GB memory you will likely encounter system issues such as slow down and unresponsiveness when starting the Vagrant instance as your machine will not have the capacity to run the VM.  There are several steps that can be taken to overcome this:
    1. Shutdown any running processes that are not required
    1. If running the [7nodes example](../7Nodes), reduce the number of nodes started up.  See the [7nodes: Reducing the number of nodes](../7Nodes#reducing-the-number-of-nodes) for info on how to do this.
    1. Set up and run the examples locally.  Running locally reduces the load on your memory compared to running in Vagrant.

### Setting up Docker

1. Install Docker (https://www.docker.com/get-started)
    - If your Docker distribution does not contain `docker-compose`, follow [this](https://docs.docker.com/compose/install/) to install Docker Compose
    - Make sure your Docker daemon has at least 4G memory
    - Required Docker Engine 18.02.0+ and Docker Compose 1.21+
1. Download and run `docker-compose`
   ```sh
   git clone https://github.com/jpmorganchase/quorum-examples
   cd quorum-examples
   docker-compose up -d
   ```
1. By default, Quorum Network is created using Tessera transaction manager and Istanbul BFT consensus. If you wish to change consensus configuration to Raft, set the environment variable `QUORUM_CONSENSUS=raft` before running `docker-compose`
   ```sh
   QUORUM_CONSENSUS=raft docker-compose up -d
   ```
1. Run `docker ps` to verify that all quorum-examples containers (7 nodes and 7 tx managers) are **healthy**
1. __Note__: to run the 7nodes demo, use the following snippet to open `geth` Javascript console to a desired node (using container name from `docker ps`) and send a private transaction
   ```sh
   $ docker exec -it quorum-examples_node1_1 geth attach /qdata/dd/geth.ipc
   Welcome to the Geth JavaScript console!

   instance: Geth/node1-istanbul/v1.7.2-stable/linux-amd64/go1.9.7
   coinbase: 0xd8dba507e85f116b1f7e231ca8525fc9008a6966
   at block: 70 (Thu, 18 Oct 2018 14:49:47 UTC)
    datadir: /qdata/dd
    modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

   > loadScript('/examples/private-contract.js')
   ```
1. Shutdown Quorum Network
   ```sh
   docker-compose down
   ```

#### Troubleshooting Docker

1. Docker is frozen
    - Check if your Docker daemon is allocated enough memory (minimum 4G)
1. Tessera is crashed due to missing file/directory
    - This is due to the location of `quorum-examples` folder is not shared
    - Please refer to Docker documentation for more details:
        - [Docker Desktop for Windows](https://docs.docker.com/docker-for-windows/troubleshoot/#shared-drives)
        - [Docker Desktop for Mac](https://docs.docker.com/docker-for-mac/#file-sharing)
        - [Docker Machine](https://docs.docker.com/machine/overview/): this depends on what Docker machine provider is used. Please refer to its documentation on how to configure shared folders/drives
1. If you run Docker inside Docker, make sure to run the container with `--privileged`

### Setting up locally

!!! info
    This is only possible with Tessera. Constellation is not supported when running the examples locally. To use Constellation,   the examples must be run in Vagrant.

1. Install [Golang](https://golang.org/dl/)
2. Download and build [Quorum](https://github.com/jpmorganchase/quorum/):
   
    ```sh
    git clone https://github.com/jpmorganchase/quorum
    cd quorum
    make
    GETHDIR=`pwd`; export PATH=$GETHDIR/build/bin:$PATH
    cd ..
    ```
    
3. Download and build Tessera (see [README](https://github.com/jpmorganchase/tessera) for build options)
   
    ```bash
    git clone https://github.com/jpmorganchase/tessera.git
    cd tessera
    mvn install
    ```
    
4. Download quorum-examples
    ```sh
    git clone https://github.com/jpmorganchase/quorum-examples
    ```

### Running the 7nodes example
Shell scripts are included in the examples to make it simple to configure the network and start submitting transactions.

All logs and temporary data are written to the `qdata` folder.

#### Using Raft consensus

1. Navigate to the 7nodes example, configure the Quorum nodes and initialize accounts & keystores:
    ```sh
    cd path/to/7nodes
    ./raft-init.sh
    ```
2. Start the Quorum and privacy manager nodes (Constellation or Tessera):
    - If running in Vagrant:
        ```sh
        ./raft-start.sh
        ```
        By default, Constellation will be used as the privacy manager.  To use Tessera run the following:
        ```
        ./raft-start.sh tessera
        ```
        By default, `raft-start.sh` will look in `/home/vagrant/tessera/tessera-app/target/tessera-app-{version}-app.jar` for the Tessera jar.

    - If running locally with Tessera:
        ```
        ./raft-start.sh tessera --tesseraOptions "--tesseraJar /path/to/tessera-app.jar"
        ```

        The Tessera jar location can also be specified by setting the environment variable `TESSERA_JAR`.

3. You are now ready to start sending private/public transactions between the nodes

#### Using Istanbul BFT consensus
To run the example using __Istanbul BFT__ consensus use the corresponding commands:
```sh
istanbul-init.sh
istanbul-start.sh
istanbul-start.sh tessera
stop.sh
```

#### Using Clique POA consensus
To run the example using __Clique POA__ consensus use the corresponding commands:
```sh
clique-init.sh
clique-start.sh
clique-start.sh tessera
stop.sh
```

### Next steps: Sending transactions
Some simple transaction contracts are included in quorum-examples to demonstrate the privacy features of Quorum.  To learn how to use them see the [7nodes](../7Nodes).

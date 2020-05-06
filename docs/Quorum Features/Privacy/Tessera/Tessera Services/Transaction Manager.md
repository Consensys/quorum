## Transaction Manager

### What is a transaction manager?

A transaction manager is the central piece in the lifecycle of a private transaction. It interfaces with most other parts of the network/infrastructure and manages the lifecycle of private data.

### What does a transaction manager do?

The transaction manager's duties include:

- forming a P2P network of transaction managers & broadcasting peer/key information
- interfacing with the enclave for encrypting/decrypting private payloads
- storing and retrieving saved data from the database
- providing the gateway for Quorum to distribute private information

The Transaction Manager, which handles peer management and database access, as well as Quorum communication, does not contain access to any private keys and does not perform and encryption/decryption, greatly reducing the impact an attack can have.

### Where does the transaction manager sit in the private transaction flow?

The transaction manager is the touch point for Quorum to distribute it's private payloads. It connects directly to Quorum and interfaces with the attached enclave, as well as with other transaction managers.

![Quorum Tessera Privacy Flow](https://github.com/jpmorganchase/tessera/raw/master/Tessera%20Privacy%20flow.jpeg)

## Setting up a Transaction Manager

### Running Tessera
The only mandatory parameter for running a minimal Transaction Manager is the location of the configuration file to use.
Use the `-configfile <path>` argument to specify the location of the config file.

Other CLI arguments can be passed, and details of these commands can be found in their respective pages - particularly around key vaults and key generation.

### Databases
By default, Tessera uses an H2 file-based database, but any JDBC compatible database can be used.

To do this, add the necessary drivers to the classpath, and run the `com.quorum.tessera.Launcher` class, like the following:

```
java -cp some-jdbc-driver.jar:/path/to/tessera-app.jar:. com.quorum.tessera.Launcher
```

For example, to use Oracle database: 
```
java -cp ojdbc7.jar:tessera-app.jar:. com.quorum.tessera.Launcher -configfile config.json
```

Some DDL scripts have been provided for more popular databases, but feel free to adapt these to whichever database you wish to use.

### Configuration

The configuration for the transaction manager is described in the [configuration overview](../../Configuration/Configuration Overview), as well as [sample configurations](../../Configuration/Sample Configuration).

### Flavours of transaction manager
For advanced users, you may decide on certain options for the transaction manager, or to disable other parts.

The default transaction manager comes with the standard options most setups will use, but other versions are as follows:

- GRPC communication (experimental)
- Non-remote only enclaves (named "tessera-simple")

These must be built from source and can be found inside the `tessera-dist` module.


## Data recovery

Tessera contains functionality to request transactions from other nodes in the network; this is useful if the database is lost or corrupted somehow. 
However, depending on the size of the network and the number of transactions made between peers, this can put heavy strain on the network resending all the data.

### How to enable
The data recovery mechanism is intended to be a "switch-on" feature as a startup command. The times when you will need this will be known prior to starting the application (usually after a disaster event). When starting Tessera, simply add the following property to the startup command: `-Dspring.profiles.active=enable-sync-poller`. This should go before any jar or class definitions, e.g. `java -Dspring.profiles.active=enable-sync-poller -jar tessera.jar -configfile config.json`.

### How it works
The data recovery procedure works by invoking a "resend request" to each new node it sees in the network. This request will cause the target node to resend each of its transactions to the intended recipient, meaning they will again save the transaction in their database.

The target node will not send back transactions as a response the request in order to ensure that a malicious node cannot get access to the transactions. i.e. anyone can send a request for a particular key, but it will mean that the node that holds that key will receive the transactions, not the node making the request. In normal usage, the node making the request and the node holding the public key are the same.

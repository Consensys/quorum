# Privacy Manager

A Privacy Manager is required to use private transactions in Quorum.

The Privacy Manager is a separate component that is concerned with the storing and distribution of encrypted private transaction data between recipients of a private transaction.

To enable private transactions, use the `PRIVATE_CONFIG` environment variable when starting a Quorum node to provide the node with the path to the Privacy Manager's `.ipc` socket, e.g.:

```shell
export PRIVATE_CONFIG=path/to/tm.ipc
``` 

The Privacy Manager has two components:

* **Transaction Manager**: See the [Homepage](../../#privacy-manager) and [Tessera's Transaction Manager page](../Tessera/Tessera%20Services/Transaction%20Manager) for more details on the responsibilities of the Transaction Manager
* **Enclave**: See [Homepage](../../#privacy-manager) and [Tessera's Enclave page](../Tessera/Tessera%20Services/Enclave) for more details on the responsibilities of the Enclave 

## Implementations
* [Tessera](../Tessera/Tessera) is a production-ready implementation of Quorum's privacy manager.  It is undergoing active development with new features being added regularly.

* [Constellation](../Constellation/Constellation) is the reference implementation of Quorum's privacy manager.  It is still supported but no longer undergoing active development of new features.  

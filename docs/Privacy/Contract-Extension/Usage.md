# Usage

## Setup

In order to use the contract extension, the node must have access to one of Tessera's `ThirdParty` servers, so that it
can send private transaction using the ethclient. The address of the third party server should be provided by the 
`--contractextension.server` CLI parameter, and just like the `PRIVATE_CONFIG` variable, it expects it to be a file path
to an IPC file.

!!! tips "Tessera `ThirdParty` Server"
    If there's already one `Thirdparty` server configuration using HTTP/HTTPS, please define additional `Thirdparty` server
    configuration block using IPC Unix Socket. Refer [here](/Privacy/Tessera/Configuration/Configuration Overview/#server) for more details

## Guide

1. The extender calls the API to initiate the extension process. This requires the **Ethereum addresses**
of the other nodes who the extender wants to vote, as well as the **PTM public key** of the new recipient.

    The API is invoked like so:
    ```
    quorumExtension.extendContract(
        "<address of contract to share>",
        "<new recipient public key>",
        [<ethereum addresses who can vote>],
        <normal tx args, which are used to send the transactions>
    )
    ```

1. All nodes who are required to vote should do so before the extender shares the state of the contract.
The only **required** private recipient of this vote is the creator of the contract, so they can keep track
of all the votes.

    The API is invoked like so:
    ```
    quorumExtension.approveExtension(
        "<extension management contract address>",
        <boolean, which way to vote>,
        <normal tx args, which are used to send the transactions>
    )
    ```

1. Once all votes are collected, the state of the contract is sent to the new recipient and a transaction
is sent to the contract using the same address as who is extending. At the point of this transaction, the new
recipient retrieves the state (assuming they accepted the change) and inserts the state into their private own state.


The process is now complete. The new party has the contract state and is able to interact with the
contract.

## Other APIs

### Viewing outstanding contracts

Current extension requests can be viewed by calling the `quorumExtension.activeExtensionContracts` endpoint which
will return a list of currently active contracts.

### Cancelling an extension

If the contract no longer wishes to be extended to the new participant, it can be cancelled to remove it
from the list of outstanding contracts. Be sure to include all the private recipients in this so they
all get the notification that the extension is cancelled.

Only the initiating account is allowed to cancel the extension. The voters must vote false, which will also end the 
extension proposal.

The API is invoked like so:
```
quorumExtension.cancelExtension(
    "<extension management contract address>",
    <normal tx args, which are used to send the transactions>
)
```

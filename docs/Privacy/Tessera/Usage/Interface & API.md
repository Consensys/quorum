## Interface Details

All interfaces can be set to run over HTTP, GRPC or HTTP-over-Unix-Sockets.

### gRPC for inter-node communication (Deprecated)

We currently have an implementation of gRPC for peer node communication as experiment API. This is not enabled on Quorum yet, but between Tessera nodes they can be enabled by adding in a couple of properties in the configuration file as child elements of `serverConfig`.

- `grpcPort` - when this value is specified, Tessera node will start a gRPC server listening on this port. The normal `port` value would still be used for starting REST server.

- `communicationType` - possible values are `REST`, `GRPC`. Default value is `REST`.

Please note that communication between Quorum and Tessera are still via unix socket. This communication flag provides additional options for Tessera peer-to-peer communication. If gRPC is the option specified, please ensure the peers urls are provided with the appropriate ports.

!!! info
    gRPC as a protocol for peer-to-peer communication will be removed from Tessera version 0.10.2

---

### Tessera to Tessera - Public API

Tessera nodes communicate with each other for:

- Node/network discovery
- Sending/Receiving encrypted payloads

The following endpoints are advertised on this interface:

* `/version`
* `/upcheck`
* `/push`
* `/resend`
* `/partyinfo`
* `/partyinfo/validate`

### Third Party - Public API 

Tessera nodes communicate with third parties for:

- storing encrypted payloads for external applications

The following endpoints are advertised on this interface:

* `/version`
* `/upcheck`
* `/storeraw`
* `/keys`
* `/partyinfo/keys`

### Quorum to Tessera - Private API

Quorum uses this API to:
- Check if the local Tessera node is running
- Send and receive details of private transactions

The following endpoints are advertised on this interface:

- `/version`
- `/upcheck`
- `/sendraw`
- `/send`
- `/receiveraw`
- `/receive`
- `/storeraw`
- `/sendsignedtx`
- `/transaction/{key}/isSender`
- `/transaction/{key}/participants`

### Admin API

Admins should use this API to:

- Access information about the Tessera node
- Make changes to the configuration of the Tessera node

The following endpoints are advertised on this API:

- `/peers` - Add to, and retrieve from, the Tessera node's peers list

## API Details

**`version`** - _Get Tessera version_

- Returns the version of Tessera that is running.

**`upcheck`** - _Check Tessera node is running_

- Returns the text "I'm up!"

**`push`** - _Push transactions between nodes_

- Persist encrypted payload received from another node.

**`resend`** - _Resend transaction_

- Resend all transactions for given key or given hash/recipient.

**`partyinfo`** - _Retrieve details of known nodes_

- GET: Request public keys/url of all known peer nodes.
- POST: accepts a stream that contains the caller node's network information, and returns a merged copy with the callee node's network information

**`partyinfo/validate`** - _Validates a node possesses a key_ 

- Will request a node to decrypt a transaction in order to prove that it has access to the private part of its advertised public key.

**`sendraw`** - _Send transaction bytestring_

- Send transaction payload bytestring from Quorum to Tessera node. Tessera sends the transaction hash in the response back. 

**`send`** - _Send transaction bytestring_

- Similar to sendraw however request payload is in json format. Please see our [Swagger documentation](https://jpmorganchase.github.io/tessera-swagger/index.html) for object model.

**`storeraw`** - _Store transaction bytestring_

- Store transaction bytestring from a third party to Tessera node. Tessera sends the transaction hash in the response back.

**`sendsignedtx`** - _Distribute signed transaction payload_

- Send transaction payload identified by hash (returned by storeraw) from Quorum to Tessera node. Tessera sends the transaction hash in the response back.

**`receiveraw`** - _Receive transaction bytestring_ 

- Receive decrypted bytestring of the transaction payload from Tessera to Quorum for transactions it is party to.

**`receive`** - _Receive transaction bytestring_

- Similar to receiveraw however request payload is in json format. Please see our [Swagger documentation](https://jpmorganchase.github.io/tessera-swagger/index.html) for object model.

**`delete`** - _Delete a transaction_ 

- Delete hashed encrypted payload stored in Tessera nodes.

**`/transaction/{key}/isSender`** - _Determine if a node is the sender_ 

- Tell if the local enclave is the sender of a particular transaction (by checking if the sender public key is part of the nodes enclave)

**`/transaction/{key}/participants`** - _Retrieve participants_ 

- Retrieve transaction participants directly from the database (a recipient will have no participants)

For more interactions with the API see the [Swagger documentation](https://jpmorganchase.github.io/tessera-swagger/index.html).

## How Constellation works

Each Constellation node hosts some number of key pairs, and advertises
a publicly accessible FQDN/port for other hosts to connect to.

Nodes can be started with a reference to existing nodes on the network
(with the `othernodes` configuration variable,) or without, in which
case some other node must later be pointed to this node to achieve
synchronization.

When a node starts up, it will reach out to each node in `othernodes`,
and learn about the public keys they host, as well as other nodes in
the network. In short order, the node's public key directory will be
the same as that of all other nodes, and you can start addressing
messages to any of the known public keys.

This is what happens when you use the `send` function of the Private
API to send the bytestring `foo` to the public key
`ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc=`:

  1. You send a POST API request to the Private API socket like:
     `{"payload": "foo", "from": "mypublickey", to: "ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="}`

  2. The local node generates using `/dev/urandom` (or similar):
       - A random Master Key (MK) and nonce
       - A random recipient nonce

  3. The local node encrypts the payload using NaCl `secretbox` using
     the random MK and nonce.

  4. The local node generates an MK container for each recipient
     public key; in this case, simply one container for `ROAZ...`,
     using NaCl `box` and the recipient nonce.

     NaCl `box` works by deriving a shared key based
     on your private key and the recipient's public key. This is known
     as elliptic curve key agreement.

     Note that the sender public key and recipient public key we
     specified above aren't enough to perform the
     encryption. Therefore, the node will check to see that it is
     actually hosting the private key that corresponds to the given
     public key before generating an MK container for each recipient
     based on SharedKey(yourprivatekey, recipientpublickey) and the
     recipient nonce.

     We now have:

       - An encrypted payload which is `foo` encrypted with the random
         MK and a random nonce. This is the same for all recipients.

       - A random recipient nonce that also is the same for all
         recipients.

       - For each recipient, the MK encrypted with the
         shared key of your private key and their public key. This
         MK container is unique per recipient, and is only transmitted to
         that recipient.

  5. For each recipient, the local node looks up the recipient host,
     and transmits to it:

       - The sender's (your) public key

       - The encrypted payload and nonce

       - The MK container for that recipient and the recipient nonce

  6. The recipient node returns a SHA3-512 hash digest of the
     encrypted payload, which represents its storage address.

     (Note that it is not possible for the sender to dictate the
     storage address. Every node generates it independently by hashing
     the encrypted payload.)

  7. The local node stores the payload locally, generating the same
     hash digest.

  8. The API call returns successfully once all nodes have confirmed
     receipt and storage of the payload, and returned a hash digest.

Now, through some other mechanism, you'll inform the recipient that
they have a payload waiting for them with the identifier `owqkrokwr`,
and they will make a call to the `receive` method of their Private
API:

  1. Make a call to the Private API socket `receive` method:
     `{"key": "qrqwrqwr"}`

  2. The local node will look in its storage for the key `qrqwrqwr`,
     and abort if it isn't found.

  3. When found, the node will use the information about the sender as
     well as its private key to derive SharedKey(senderpublickey,
     yourprivatekey) and decrypt the MK container using NaCl `box`
     with the recipient nonce.

  4. Using the decrypted MK, the local node will decrypt the encrypted
     payload using NaCl `secretbox` using the main nonce.

  5. The API call returns the decrypted data.


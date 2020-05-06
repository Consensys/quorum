# Constellation

Constellation is a self-managing, peer-to-peer system in which each
node:

  - Hosts a number of NaCl (Curve25519) public/private key pairs.

  - Automatically discovers other nodes on the network after
    synchronizing with as little as one other host.

  - Synchronizes a directory of public keys mapped to recipient hosts
    with other nodes on the network.

  - Exposes a public API which allows other nodes to send encrypted
    bytestrings to your node, and to synchronize, retrieving
    information about the nodes that your node knows about.

  - Exposes a private API which:

      - Allows you to send a bytestring to one or more  public keys,
        returning a content-addressable identifier. This bytestring is
        encrypted transparently and efficiently (at symmetric
        encryption speeds) before being transmitted over the wire to
        the correct recipient nodes (and only those nodes.) The
        identifier is a hash digest of the encrypted payload that
        every receipient node receives. Each recipient node also
        receives a small blob encrypted for their public key which
        contains the Master Key for the encrypted payload.

      - Allows you to receive a decrypted bytestring
        based on an identifier. Payloads which your node has sent or
        received can be decrypted and retrieved in this way.

      - Exposes methods for deletion, resynchronization, and other
        management functions.

  - Supports a number of storage backends including LevelDB,
    BerkeleyDB, SQLite, and Directory/Maildir-style file storage
    suitable for use with any FUSE adapter, e.g. for AWS S3.

  - Uses mutually-authenticated TLS with modern settings and various trust
    models including hybrid CA/tofu (default), tofu (think OpenSSH), and
    whitelist (only some set of public keys can connect.)

  - Supports access controls like an IP whitelist.

Conceptually, one can think of Constellation as an amalgamation of a
distributed key server, PGP encryption (using modern cryptography,)
and Mail Transfer Agents (MTAs.)

Constellation's current primary application is to implement the
"privacy engine" of Quorum, a fork of Ethereum with support for
private transactions that function exactly as described in this
README. Private transactions in Quorum contain only a flag indicating
that they're private and the content-addressable identifier described
here.

Constellation can be run stand-alone as a daemon via
`constellation-node`, or imported as a Haskell library, which allows
you to implement custom storage and encryption logic.

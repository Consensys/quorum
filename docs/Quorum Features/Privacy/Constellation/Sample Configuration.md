``` yaml
#####
## Constellation configuration file example
## ----------------------------------------
## Every option listed here can also be specified on the command line, e.g.
## `constellation-node --url=http://www.foo.com --port 9001 ...`
## (lists are given using comma-separated strings)
## If both command line parameters and a configuration file are given, the
## command line options will take precedence.
##
## The only strictly necessary option is `port`, however it's recommended to
## set at least the following:
##
##   --url           The URL to advertise to other nodes (reachable by them)
##   --port          The local port to listen on
##   --workdir       The folder to put stuff in (default: .)
##   --socket        IPC socket to create for access to the Private API
##   --othernodes    "Boot nodes" to connect to to discover the network
##   --publickeys    Public keys hosted by this node
##   --privatekeys   Private keys hosted by this node (in corresponding order)
##
## Example usage:
##
##   constellation-node --workdir=data --generatekeys=foo
##   (To generate a keypair foo in the data directory)
##
##   constellation-node --url=https://localhost:9000/ \
##                      --port=9000 \
##                      --workdir=data \
##                      --socket=constellation.ipc \
##                      --othernodes=https://localhost:9001/ \
##                      --publickeys=foo.pub \
##                      --privatekeys=foo.key
##
##   constellation-node sample.conf
##
##   constellation-node --port=9002 sample.conf
##   (This overrides the port value given in sample.conf)
##
## Note on defaults: "Default:" below indicates the value that will be assumed
## if the option is not present either in the configuration file or as a command
## line parameter.
##
## Note about security: In the default configuration, Constellation will
## automatically generate TLS certificates and trust other nodes' certificates
## when they're first encountered (trust-on-first-use). See the documentation
## for tlsservertrust and tlsclienttrust below. To disable TLS entirely, e.g.
## when using Constellation in conjunction with a VPN like WireGuard, set tls to
## off.
#####

## Externally accessible URL for this node's public API (this is what's
## advertised to other nodes on the network, and must be reachable by them.)
url = "http://127.0.0.1:9001/"

## Port to listen on for the public API.
port = 9001

## Directory in which to put and look for other files referenced here.
##
## Default: The current directory
workdir = "data"

## Socket file to use for the private API / IPC. If this is commented out,
## the private API will not be accessible.
##
## Default: Not set
socket = "constellation.ipc"

## Initial (not necessarily complete) list of other nodes in the network.
## Constellation will automatically connect to other nodes not in this list
## that are advertised by the nodes below, thus these can be considered the
## "boot nodes."
##
## Default: []
othernodes = ["http://127.0.0.1:9000/"]

## The set of public keys this node will host.
##
## Default: []
publickeys = ["foo.pub"]

## The corresponding set of private keys. These must correspond to the public
## keys listed above.
##
## Default: []
privatekeys = ["foo.key"]

## Optional comma-separated list of paths to public keys to add as recipients
## for every transaction sent through this node, e.g. for backup purposes.
## These keys must be advertised by some Constellation node on the network, i.e.
## be in a node's publickeys/privatekeys lists.
##
## Default: []
alwayssendto = []

## Optional file containing the passwords needed to unlock the given privatekeys
## (the file should contain one password per line -- add an empty line if any
## one key isn't locked.)
##
## Default: Not set
# passwords = "passwords"

## Storage engine used to save payloads and related information. Options:
##   - bdb:path (BerkeleyDB)
##   - dir:path (Directory/file storage - can be used with e.g. FUSE-mounted
##     file systems.)
##   - leveldb:path (LevelDB - experimental)
##   - memory (Contents are cleared when Constellation exits)
##   - sqlite:path (SQLite - experimental)
##
## Default: "dir:storage"
storage = "dir:storage"

## Verbosity level (each level includes all prior levels)
##   - 0: Only fatal errors
##   - 1: Warnings
##   - 2: Informational messages
##   - 3: Debug messages
##
## At the command line this can be specified using -v0, -v1, -v2, -v3, or
## -v (2) and -vv (3).
##
## Default: 1
verbosity = 1

## Optional IP whitelist for the public API. If unspecified/empty,
## connections from all sources will be allowed (but the private API remains
## accessible only via the IPC socket above.) To allow connections from
## localhost when a whitelist is defined, e.g. when running multiple
## Constellation nodes on the same machine, add "127.0.0.1" and "::1" to
## this list.
##
## Default: Not set
# ipwhitelist = ["10.0.0.1", "2001:0db8:85a3:0000:0000:8a2e:0370:7334"]

## TLS status. Options:
##
##   - strict: All connections to and from this node must use TLS with mutual
##     authentication. See the documentation for tlsservertrust and
##     tlsclienttrust below.
##   - off: Mutually authenticated TLS is not used for in- and outbound
##     connections, although unauthenticated connections to HTTPS hosts are
##     still possible. This should only be used if another transport security
##     mechanism like WireGuard is in place.
##
## Default: "strict"
tls = "strict"

## Path to a file containing the server's TLS certificate in Apache format.
## This is used to identify this node to other nodes in the network when they
## connect to the public API.
##
## This file will be auto-generated if it doesn't exist.
##
## Default: "tls-server-cert.pem"
tlsservercert = "tls-server-cert.pem"

## List of files that constitute the CA trust chain for the server certificate.
## This can be empty for auto-generated/non-PKI-based certificates.
##
## Default: []
tlsserverchain = []

## The private key file for the server TLS certificate.
##
## This file will be auto-generated if it doesn't exist.
##
## Default: "tls-server-key.pem"
tlsserverkey = "tls-server-key.pem"

## TLS trust mode for the server. This decides who's allowed to connect to it.
## Options:
##
##   - whitelist: Only nodes that have previously connected to this node and
##     been added to the tlsknownclients file below will be allowed to connect.
##     This mode will not add any new clients to the tlsknownclients file.
##
##   - tofu: (Trust-on-first-use) Only the first node that connects identifying
##     as a certain host will be allowed to connect as the same host in the
##     future. Note that nodes identifying as other hosts will still be able
##     to connect -- switch to whitelist after populating the tlsknownclients
##     list to restrict access.
##
##   - ca: Only nodes with a valid certificate and chain of trust to one of
##     the system root certificates will be allowed to connect. The folder
##     containing trusted root certificates can be overriden with the
##     SYSTEM_CERTIFICATE_PATH environment variable.
##
##   - ca-or-tofu: A combination of ca and tofu: If a certificate is valid,
##     it is always allowed and added to the tlsknownclients list. If it is
##     self-signed, it will be allowed only if it's the first certificate this
##     node has seen for that host.
##
##   - insecure-no-validation: Any client can connect, however they will still
##     be added to the tlsknownclients file.
##
## Default: "tofu"
tlsservertrust = "tofu"

## TLS known clients file for the server. This contains the fingerprints of
## public keys of other nodes that are allowed to connect to this one.
##
## Default: "tls-known-clients"
tlsknownclients = "tls-known-clients"

## Path to a file containing the client's TLS certificate in Apache format.
## This is used to identify this node to other nodes in the network when it is
## connecting to their public APIs.
##
## This file will be auto-generated if it doesn't exist.
##
## Default: "tls-client-cert.pem"
tlsclientcert = "tls-client-cert.pem"

## List of files that constitute the CA trust chain for the client certificate.
## This can be empty for auto-generated/non-PKI-based certificates.
##
## Default: []
tlsclientchain = []

## The private key file for the client TLS certificate.
##
## This file will be auto-generated if it doesn't exist.
##
## Default: "tls-client-key.pem"
tlsclientkey = "tls-client-key.pem"

## TLS trust mode for the client. This decides which servers it will connect to.
## Options:
##
##   - whitelist: This node will only connect to servers it has previously seen
##     and added to the tlsknownclients file below. This mode will not add
##     any new servers to the tlsknownservers file.
##
##   - tofu: (Trust-on-first-use) This node will only connect to the same
##     server for any given host. (Similar to how OpenSSH works.)
##
##   - ca: The node will only connect to servers with a valid certificate and
##     chain of trust to one of the system root certificates. The folder
##     containing trusted root certificates can be overriden with the
##     SYSTEM_CERTIFICATE_PATH environment variable.
##
##   - ca-or-tofu: A combination of ca and tofu: If a certificate is valid,
##     it is always allowed and added to the tlsknownservers list. If it is
##     self-signed, it will be allowed only if it's the first certificate this
##     node has seen for that host.
##
##   - insecure-no-validation: This node will connect to any server, regardless
##     of certificate, however it will still be added to the tlsknownservers
##     file.
##
## Default: "ca-or-tofu"
tlsclienttrust = "ca-or-tofu"

## TLS known servers file for the client. This contains the fingerprints of
## public keys of other nodes that this node has encountered.
##
## Default: "tls-known-servers"
tlsknownservers = "tls-known-servers"
```


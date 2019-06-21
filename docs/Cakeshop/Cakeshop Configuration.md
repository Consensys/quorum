## Cakeshop

The sections below describe the various configuration settings for Cakeshop itself.

Most of the configuration can be found in the file `application.properties` located in your `data` folder at `data/<env>/application.properties`. The default `env` is `local` resulting in the file `data/local/application.properties`.

Some settings (as noted below) may be overridden via Java system properties passed via `java -D<prop>=<val>`. e.g.,

```sh
$ java -Dcakeshop.database.vendor=oracle -jar cakeshop.war
```

Finally, certain properties can also be passed via environment variable, as detailed below.

Config order:

* Environment var
* System property
* Config file (application.properties)

### geth

Cakeshop can use geth in one of two modes: __managed__ or __unmanaged__.

* In __managed__ mode, Cakeshop will start a bundled version of geth which can be controlled from within the application.
* In __unmanaged__ mode, Cakeshop will connect to an externally managed version of geth (one that is started in some other way or running on another machine).

The various options are described below and can be modified via `application.properties`.

Also note that normal "public node" options are disabled by default and there is
currently no way to change this behavior:

    --nat none --nodiscover --ipcdisable

```
# Toggle managed mode: true = managed
geth.auto.start=true
geth.auto.stop=true

# RPC URL to connect to geth on [managed or unmanaged]
#
# In managed mode, the port is passed to --rpcport. --rpcaddr defaults to
# '127.0.0.1' but can be overridden via geth.params.extra, below.
#
# In unmanaged mode, simply uses this URL for issuing RPC calls.
geth.url=http://localhost:8102

# Sets the p2p node port (--port) [managed only]
geth.node.port=30303

# Network identifier (--networkid) [managed only]
geth.networkid=1006

# Data directory path (--datadir) [managed only]
#
# When set to the default value of "/.ethereum", will be rewritten to a local
# path at `data/<env>/ethereum`
geth.datadir=/.ethereum

# List of RPC APIs to enable [managed only]
#
# Only modify if you know what you are doing.
#
# NOTE: If running in `unmanaged` mode, make sure your geth is configured
#       to enable at least these APIs.
geth.rpcapi.list=admin,db,eth,debug,miner,net,shh,txpool,personal,web3

# geth log verbosity (--verbosity) [managed or unmanaged]
#
# Higher is more verbose. Can be modified at runtime without restart via API.
geth.verbosity=

# Custom node name (--identity) [managed only]
geth.identity=

# Enable Mining (--mine) [managed or unmanaged]
#
# Can be modified at runtime without restart via API. Only applies to vanilla
# geth.
#
# When using Raft-based consensus with quorum, the mining value doesn't have any affect.
geth.mining=true

# Enable CORS for geth RPC endpoint [managed only]
geth.cors.enabled=false
geth.cors.url=

# Extra geth params [managed only]
geth.params.extra=

```

### Contract Registry

On startup, Cakeshop will deploy a custom Smart Contract Registry onto the
configured chain. This registry is used to store metadata about the contract,
most importantly the ABI and the original solidity source. Once deployed, the
registry's address is stored in the config file under the
`contract.registry.addr` key. This value can be overridden by passing the
environment variable `CAKESHOP_REGISTRY_ADDR`.

#### Shared config mode

Another mode exists to facilitate locally testing a multi-node cluster by
running Cakeshop multiple times in managed mode. By setting the
`CAKESHOP_SHARED_CONFIG` environment variable, Cakeshop will read/write the
registry address from/to the shared location. This allows this address to be
easily shared among multiple Cakeshops.

Note, however, that the `CAKESHOP_REGISTRY_ADDR` env var takes highest
precedence and will override all other property file values (both local and
shared).

Example: `CAKESHOP_REGISTRY_ADDR=$HOME/data/shared_cakeshop_config`

### Cakeshop Internals

```
# Timeout between polls on contract deploy
contract.poll.delay.millis=5000

# Address of contract registry
contract.registry.addr=

# Various internal queues and thread pools
# Defaults are probably fine, modify at your own risk
cakeshop.mvc.async.pool.threads.core=250
cakeshop.mvc.async.pool.threads.max=1000
cakeshop.mvc.async.pool.queue.max=2000
```

### Database

Cakeshop stores and indexes certain information in a relational database (RDBMS) to support certain features, namely displaying transaction history for a given contract. An embedded HSQLDB is used by default for development purposes and is sufficient for most use cases.

The ability to use an external database is also available. Currently supported are Oracle, PostgreSQL, MySQL, and HSQLDB.

The following properties can be set either in `application.properties` or via Java system properties:

```
cakeshop.database.vendor              Enables the preferred db driver.
                                      Allowed options are:
                                      hsqldb|oracle|mysql|postgres
                                      Default: hsqldb

cakeshop.jndi.name                    Used for configuring an external
                                      connection pool (usually container-
                                      managed) with oracle, mysql or postgres.

cakeshop.jdbc.url                     JDBC URL
cakeshop.jdbc.user                    JDBC username
cakeshop.jdbc.pass                    JDBC password

cakeshop.hibernate.jdbc.batch_size    Hibernate tuneables
cakeshop.hibernate.hbm2ddl.auto

cakeshop.hibernate.dialect            Hibernate dialect. Will usually be set
                                      automatically by your db pref. Override
                                      it here.
```

### Spring Framework Internals

The settings below control the behavior of Spring. For detailed information, please refer to the [Spring Framework](http://docs.spring.io/spring/docs/4.2.5.RELEASE/spring-framework-reference/htmlsingle/) and [Spring Boot Actuator](http://docs.spring.io/spring-boot/docs/1.3.3.RELEASE/reference/htmlsingle/#production-ready) documentation.

```config
# spring config
spring.main.banner-mode=off

spring.mvc.view.prefix=/WEB-INF/jsp/
spring.mvc.view.suffix=.jsp

server.compression.enabled=true
server.compression.mime-types=application/json,application/xml,text/html,text/xml,text/plain

# spring boot actuator
management.context-path=/manage
endpoints.actuator.enabled=true
```

## Tomcat

Minimum requirements for tomcat (`conf/server.xml`):

```xml
  <Connector port="8080" protocol="HTTP/1.1"
             enableLookups="false"
             maxKeepAliveRequests="-1"
             maxConnections="10000"
             redirectPort="8443"
             connectionTimeout="20000"/>
```


## Building a Docker Image

```
# build custom maven image
docker build --pull -t jpmc/cakeshop-build docker/build/

# build cakeshop.war
docker run --rm -v ~/.m2:/home/cakeshop/.m2 -v $(pwd):/usr/src -w /usr/src jpmc/cakeshop-build mvn -DskipTests clean package

# build cakeshop image using war from previous step
mv cakeshop-api/target/cakeshop*.war docker/cakeshop/
docker build --pull -t cakeshop docker/cakeshop/
```
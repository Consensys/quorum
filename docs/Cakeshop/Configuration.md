# Cakeshop Configuration


Cakeshop follows standard [Spring Boot configuration patterns](https://docs.spring.io/spring-boot/docs/2.0.9.RELEASE/reference/html/boot-features-external-config.html#:~:text=Spring%20Boot%20lets%20you%20externalize,line%20arguments%20to%20externalize%20configuration.), so you may provide an `application.properties` in the location where you are running Cakeshop to customize some settings.

You may also override config options via Java system properties passed via `java -D<prop>=<val>`:

```
java -Dcakeshop.initialnodes=data/cakeshop/nodes.json -jar cakeshop.war
```

### Initial Nodes

When Cakeshop starts for the first time, it will not know which node(s) it is supposed to connect to. You will need to click on 'Manage Nodes' in the top right corner of the Cakeshop UI to add nodes by RPC url.

Alternatively, you may provide an initial set of nodes in a JSON file that Cakeshop will use to prepopulate the nodes list. This file will only be used if no nodes have previously been added to Cakeshops database.

The format of the JSON file is as follows:

```
[
  {
    "name": "node1",
    "rpcUrl": "http://localhost:22000",
    "transactionManagerUrl": "http://localhost:9081"
  },
  {
    "name": "node2",
    "rpcUrl": "http://127.0.0.1:22001",
    "transactionManagerUrl": "http://127.0.0.1:9082"
  },
  {
    "name": "node3",
    "rpcUrl": "http://localhost:22002",
    "transactionManagerUrl": "http://localhost:9083"
  }
]
```

The rpcUrl field should be the RPC endpoint on the Quorum (geth) node, and the transactionManagerUrl should be the Tessera 3rd Party API endpoint. 

Provide the location of the initial nodes file through application.properties or by using the `-D` flag mentioned above.
```sh
# inside application.properties
cakeshop.initialnodes=path/to/nodes.json
```

### Cakeshop Internals

Some other options that may be customized in application.properties:

```sh
# some systems don't call the nodejs binary 'node', in that change you can change this value
nodejs.binary=node

# Timeout between polls on contract deploy
contract.poll.delay.millis=5000

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

server.compression.enabled=true
server.compression.mime-types=application/json,application/xml,text/html,text/xml,text/plain

# spring boot actuator
management.context-path=/manage
endpoints.actuator.enabled=true
```

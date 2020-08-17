## Download

Binary packages are available on the [Github releases page](https://github.com/jpmorganchase/cakeshop/releases).

## Running via Quorum Wizard

The easiest way to use Cakeshop is to generate a Quorum network with [Quorum Wizard](../../Wizard/GettingStarted) and choose to deploy Cakeshop alongside the network.

## Running via Spring Boot

### Requirements

* Java 11+
* NodeJS (if the nodejs binary on your machine isn't called `node`, see [here](../Configuration#cakeshop-internals))

### Running

* Download WAR file
* Run `java -jar cakeshop.war`
* Navigate to [http://localhost:8080/](http://localhost:8080/)


## Running via Docker

Simple example of running via docker on port 8080:

```sh
docker run -p 8080:8080 quorumengineering/cakeshop
```

Then access the UI at [http://localhost:8080/](http://localhost:8080/)

### Docker Customizations
You can add some extra flags to the run command to further customize cakeshop.

Here is an example where you mount `./data` as a data volume for the container to use:

```sh
mkdir data
docker run -p 8080:8080 -v "$PWD/data":/opt/cakeshop/data quorumengineering/cakeshop
```

An example providing an initial nodes.json in the data directory and configuring it to be used:

```sh
# makes sure you have nodes.json at $PWD/data/nodes.json
docker run -p 8080:8080 -v "$PWD/data":/opt/cakeshop/data \
    -e JAVA_OPTS="-Dcakeshop.initialnodes=/opt/cakeshop/data/nodes.json" \
    quorumengineering/cakeshop
```

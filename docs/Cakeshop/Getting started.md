
## Download

Binary packages are available for macOS, Windows, and Linux platforms on the [releases](https://github.com/jpmorganchase/cakeshop/releases) page.


## Quickstart

### Requirements

* Java 8+
* Java app server (Tomcat, Jetty, etc) [Optional]

### Running via Spring Boot

* Download WAR file
* Run `java -jar cakeshop.war`
* Navigate to [http://localhost:8080/cakeshop/](http://localhost:8080/cakeshop/)

*Note: when running in Windows, -Dgeth.node=geth must be specified as Quorum is not yet available on Windows OS*

### Running via App Server

* Download WAR file
* Put in `/webapps` folder of your app server
* Add Java system property `-Dspring.profiles.active=local` to startup script (`setenv.sh` for tomcat)
* Start app server
* Navigate to [http://localhost:8080/cakeshop/](http://localhost:8080/cakeshop/) (default port is usually 8080)

*Note: when running in Windows, -Dgeth.node=geth must be specified as Quorum is not yet available on Windows OS*

### Running via Docker -- NEEDS UPDATE

Run via docker and access UI on [http://localhost:8080/cakeshop/](http://localhost:8080/cakeshop/)

```sh
docker run -p 8080:8080 quorumengineering/cakeshop
```

You'll probably want to mount a data volume:

```sh
mkdir data
docker run -p 8080:8080 -v "$PWD/data":/opt/cakeshop/data quorumengineering/cakeshop
```

Running under a specific environment

```sh
docker run -p 8080:8080 -v "$PWD/data":/opt/cakeshop/data \
    -e JAVA_OPTS="-Dspring.profiles.active=local" \
    quorumengineering/cakeshop
```

Note that DAG generation will take time and Cakeshop will not be available until it's complete. If you already have a DAG for epoch 0 in your `$HOME/.ethash` folder, then you can expose that to your container (or just cache it for later):

```sh
docker run -p 8080:8080 -v "$PWD/data":/opt/cakeshop/data \
    -v $HOME/.ethash:/opt/cakeshop/.ethash \
    quorumengineering/cakeshop
```


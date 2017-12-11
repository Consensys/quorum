FROM ubuntu:16.04

# Install add-apt-repository
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        libdb-dev \
        libleveldb-dev \
        libsodium-dev \
        libtinfo-dev \
        zlib1g-dev \
        libgmp10

COPY \
        ./build/bin/constellation-node \
        ./build/docker/constellation/start.sh \
    /usr/local/bin/

CMD ["/usr/local/bin/start.sh"]

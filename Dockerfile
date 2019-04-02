# Build Geth in a stock Go builder container
FROM golang:1.11-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /go-ethereum
RUN cd /go-ethereum && make geth bootnode

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates build-base git bash curl python npm make gcc 
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/
COPY --from=builder /go-ethereum/build/bin/bootnode /usr/local/bin/

RUN mkdir -p /ledgerium/governanceapp/governanceapp \
    && cd /ledgerium/governanceapp \
    && git clone -b feat/LB-101 https://github.com/ledgerium/governanceapp.git

WORKDIR /ledgerium/governanceapp/governanceapp
#RUN git checkout feat/LB-101
RUN npm install

WORKDIR /ledgerium/governanceapp/governanceapp/app

RUN npm install
RUN npm install web3@1.0.0-beta.36

EXPOSE 8545 30303 30303/udp
ENTRYPOINT ["geth"]

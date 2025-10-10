# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.22-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-ethereum/
COPY go.sum /go-ethereum/
RUN cd /go-ethereum && go mod download

ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/geth
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/bootnode

# Pull Geth into a second stage deploy alpine container
FROM alpine:3.20

# Update busybox to fix CVE (Alpine Linux Security Advisory 511246)
# Required version: busybox 1.36.1-r29 or higher
# Update OpenSSL to fix CVE (Alpine Linux Security Advisory 513818)
# Required version: openssl 3.3.1-r1, libssl3 3.3.1-r1, libcrypto3 3.3.1-r1
# Update curl to fix CVE (Alpine Linux Security Advisory 513861)
# Required version: curl 8.9.0-r0, libcurl 8.9.0-r0
RUN apk update && \
    apk upgrade --no-cache busybox busybox-binsh ssl_client libssl3 libcrypto3 openssl curl libcurl && \
    apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/
COPY --from=builder /go-ethereum/build/bin/bootnode /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["geth"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"

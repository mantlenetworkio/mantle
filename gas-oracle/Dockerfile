FROM golang:1.19-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./gas-oracle /gas-oracle
COPY ./bss-core /bss-core
COPY ./l2geth /l2geth
COPY ./metrics /metrics
COPY ./fraud-proof /fraud-proof

COPY ./gas-oracle/docker.go.work /go/go.work

RUN cd /gas-oracle && make gas-oracle

FROM alpine:3.15

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /gas-oracle/gas-oracle /usr/local/bin/
COPY --from=builder /gas-oracle/gas-oracle.sh /usr/local/bin/

WORKDIR /usr/local/bin/
ENTRYPOINT ["/bin/sh","gas-oracle.sh"]

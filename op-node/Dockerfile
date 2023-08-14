FROM --platform=$BUILDPLATFORM golang:1.19.9-alpine3.16 as builder

ARG VERSION=v0.0.0

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build op-node with the shared go.mod & go.sum files
COPY ./op-node /go/op-node
COPY ./op-chain-ops /go/op-chain-ops
COPY ./op-service /go/op-service
COPY ./op-bindings /go/op-bindings
COPY ./.git /go/.git
COPY ./op-node/docker.go.work /go/go.work

WORKDIR /go/op-node

RUN make op-node

FROM alpine:3.16

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/op-node/bin/op-node /usr/local/bin

WORKDIR /usr/local/bin
CMD ["op-node"]

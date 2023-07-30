FROM --platform=$BUILDPLATFORM golang:1.19.9-alpine3.16 as builder

ARG VERSION=v0.0.0

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build op-batcher with the shared go.mod & go.sum files
COPY ./op-batcher /go/op-batcher
COPY ./op-bindings /go/op-bindings
COPY ./op-node /go/op-node
COPY ./op-service /go/op-service
COPY ./op-signer /go/op-signer
COPY ./op-batcher/go.mod /go/op-batcher/go.mod
COPY ./op-batcher/go.sum /go/op-batcher/go.sum
COPY ./op-batcher/docker.go.work /go/go.work
COPY ./.git /go/.git

WORKDIR /go/op-batcher

RUN make op-batcher

FROM alpine:3.16

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/op-batcher/bin/op-batcher /usr/local/bin

WORKDIR /usr/local/bin
ENTRYPOINT ["op-batcher"]

FROM --platform=$BUILDPLATFORM golang:1.19.9-alpine3.16 as builder

ARG VERSION=v0.0.0

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build op-proposer with the shared go.mod & go.sum files
COPY ./op-proposer /go/op-proposer
COPY ./op-bindings /go/op-bindings
COPY ./op-node /go/op-node
COPY ./op-service /go/op-service
COPY ./op-signer /go/op-signer
COPY ./op-proposer/go.mod /go/op-proposer/go.mod
COPY ./op-proposer/go.sum /go/op-proposer/go.sum
COPY ./.git /go/.git
COPY ./op-proposer/docker.go.work /go/go.work

WORKDIR /go/op-proposer

RUN make op-proposer

FROM alpine:3.16

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/op-proposer/bin/op-proposer /usr/local/bin

WORKDIR /usr/local/bin
CMD ["op-proposer"]

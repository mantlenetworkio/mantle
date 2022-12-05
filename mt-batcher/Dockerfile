FROM golang:1.18.0-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./mt-batcher /go/mt-batcher
COPY ./bss-core /go/bss-core
COPY ./l2geth /go/l2geth
COPY ./mt-batcher/docker.go.work /go/go.work

WORKDIR /go/mt-batcher
RUN make

FROM alpine:3.15

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /go/mt-batcher/mt-batcher /usr/local/bin/

WORKDIR /usr/local/bin
COPY ./ops/scripts/mt-batcher.sh .
ENTRYPOINT ["mt-batcher"]

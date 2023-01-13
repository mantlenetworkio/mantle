FROM golang:1.18.0-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./mt-meso-service /go/mt-meso-service
COPY ./datalayr-mantle /go/datalayr-mantle
COPY ./l2geth /go/l2geth
COPY ./mt-batcher /go/mt-batcher
COPY ./bss-core /go/bss-core
COPY ./mt-meso-service/docker.go.work /go/go.work

WORKDIR /go/mt-meso-service
RUN make mt-meso-service

FROM alpine:3.15

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /go/mt-meso-service/mt-meso-service /usr/local/bin/

WORKDIR /usr/local/bin

ENTRYPOINT ["mt-meso-service"]

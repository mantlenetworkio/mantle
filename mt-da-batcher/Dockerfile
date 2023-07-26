FROM golang:1.19.0-alpine3.15 as builder


RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./mt-da-batcher /go/mt-da-batcher
COPY ./bss-core /go/bss-core
COPY ./datalayr /go/datalayr
COPY ./l2geth /go/l2geth
COPY ./fraud-proof /go/fraud-proof
COPY ./metrics  /go/metrics
COPY ./mt-da-batcher/docker.go.work /go/go.work

WORKDIR /go/mt-da-batcher
RUN make

FROM alpine:3.15

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/mt-da-batcher/mt-da-batcher /usr/local/bin/

WORKDIR /usr/local/bin
ENTRYPOINT ["mt-da-batcher"]

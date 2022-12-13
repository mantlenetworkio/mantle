FROM golang:1.19-alpine3.17 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./scheduler /go/scheduler
COPY ./bss-core /go/bss-core
COPY ./l2geth /go/l2geth
COPY ./scheduler/docker.go.work /go/go.work

WORKDIR /go/scheduler
RUN make

FROM alpine:3.17

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /go/scheduler/scheduler /usr/local/bin/

WORKDIR /usr/local/bin
COPY ./ops/scripts/scheduler.sh .
ENTRYPOINT ["scheduler"]

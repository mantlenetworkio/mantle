FROM golang:1.19.0-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./mt-da-challenger /go/mt-da-challenger
COPY ./datalayr /go/datalayr
COPY ./l2geth /go/l2geth
COPY ./mt-da-batcher /go/mt-da-batcher
COPY ./bss-core /go/bss-core
COPY ./fraud-proof /go/fraud-proof
COPY ./metrics  /go/metrics
COPY ./mt-da-challenger/docker.go.work /go/go.work

WORKDIR /go/mt-da-challenger
RUN make mt-da-challenger

FROM alpine:3.15

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/mt-da-challenger/mt-da-challenger /usr/local/bin/

WORKDIR /usr/local/bin

ENTRYPOINT ["mt-da-challenger"]

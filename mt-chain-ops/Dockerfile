FROM golang:1.18.0-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./mt-chain-ops /app/mt-chain-ops
COPY ./mt-bindings /app/mt-bindings
COPY ./mt-node /app/mt-node
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./.git /app/.git

WORKDIR /app/mt-chain-ops

RUN make mt-migrate

FROM alpine:3.15

COPY --from=builder /app/mt-chain-ops/bin/op-migrate /usr/local/bin

ENTRYPOINT ["op-migrate"]

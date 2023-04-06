FROM --platform=$BUILDPLATFORM golang:1.18.0-alpine3.15 as builder

ARG VERSION=v0.0.0

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build mt-l1batcher with the shared go.mod & go.sum files
COPY ./mt-l1batcher /app/mt-l1batcher
COPY ./mt-bindings /app/mt-bindings
COPY ./mt-node /app/mt-node
COPY ./mt-service /app/mt-service
COPY ./mt-signer /app/mt-signer
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

COPY ./.git /app/.git

WORKDIR /app/mt-l1batcher

RUN go mod download

ARG TARGETOS TARGETARCH
RUN echo $TARGETARCH

RUN make mt-l1batcher VERSION="$VERSION" GOOS=$TARGETOS GOARCH=$TARGETARCH

FROM alpine:3.15

COPY --from=builder /app/mt-l1batcher/bin/mt-l1batcher /usr/local/bin

ENTRYPOINT ["mt-l1batcher"]

FROM --platform=$BUILDPLATFORM golang:1.18.0-alpine3.15 as builder

ARG VERSION=v0.0.0

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build mt-node with the shared go.mod & go.sum files
COPY ./mt-node /app/mt-node
COPY ./mt-chain-ops /app/mt-chain-ops
COPY ./mt-service /app/mt-service
COPY ./mt-bindings /app/mt-bindings
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./.git /app/.git

WORKDIR /app/mt-node

RUN go mod download

ARG TARGETOS TARGETARCH

RUN make mt-node VERSION="$VERSION" GOOS=$TARGETOS GOARCH=$TARGETARCH

FROM alpine:3.15

COPY --from=builder /app/mt-node/bin/mt-node /usr/local/bin

CMD ["mt-node"]

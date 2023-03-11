FROM --platform=$BUILDPLATFORM golang:1.18.0-alpine3.15 as builder

ARG VERSION=v0.0.0

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build mt-proposer with the shared go.mod & go.sum files
COPY ./mt-proposer /app/mt-proposer
COPY ./mt-bindings /app/mt-bindings
COPY ./mt-node /app/mt-node
COPY ./mt-service /app/mt-service
COPY ./mt-signer /app/mt-signer
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./.git /app/.git

WORKDIR /app/mt-proposer

RUN go mod download

ARG TARGETOS TARGETARCH

RUN make mt-proposer VERSION="$VERSION" GOOS=$TARGETOS GOARCH=$TARGETARCH

FROM alpine:3.15

COPY --from=builder /app/mt-proposer/bin/mt-proposer /usr/local/bin

CMD ["mt-proposer"]

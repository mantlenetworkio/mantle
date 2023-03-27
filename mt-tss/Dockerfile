# Build tss in a stock Go builder container
FROM golang:1.19-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /app/

WORKDIR /app/tss
RUN make build

# Pull tss into a second stage deploy alpine container
FROM alpine:3.15

WORKDIR /app

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /app/tss/tss /app

# EXPOSE 8545 8546 8547
ENTRYPOINT ["./tss"]

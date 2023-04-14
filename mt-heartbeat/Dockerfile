FROM golang:1.18.0-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build mt-heartbeat with local monorepo go modules
COPY ./mt-heartbeat /app/mt-heartbeat
COPY ./.git /app/.git
COPY ./mt-heartbeat/go.mod /app/go.mod
COPY ./mt-heartbeat/go.sum /app/go.sum

WORKDIR /app/mt-heartbeat

RUN make mt-heartbeat

FROM alpine:3.15

COPY --from=builder /app/mt-heartbeat/bin/mt-heartbeat /usr/local/bin

CMD ["mt-heartbeat"]

FROM golang:1.19-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./subsidy /subsidy

RUN cd /subsidy && make subsidy

FROM alpine:3.15

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /subsidy/subsidy /usr/local/bin/
COPY --from=builder /subsidy/subsidy.sh /usr/local/bin/

WORKDIR /usr/local/bin/
ENTRYPOINT ["/bin/sh","subsidy.sh"]

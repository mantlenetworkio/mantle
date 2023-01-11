FROM golang:1.19-alpine3.17 as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash
RUN go env -w GOPROXY=https://goproxy.cn
COPY ./gas-oracle /gas-oracle

RUN cd /gas-oracle && make gas-oracle

FROM alpine:3.17

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /gas-oracle/gas-oracle /usr/local/bin/
COPY --from=builder /gas-oracle/gas-oracle.sh /usr/local/bin/

WORKDIR /usr/local/bin/
ENTRYPOINT ["/bin/sh","gas-oracle.sh"]

FROM golang:1.19 as builder
WORKDIR /base
ADD . .
RUN cd tss && CGO_ENABLED=1 make build

FROM centos:8 as tss
WORKDIR /app
COPY --from=builder /base/tss/tss /app
RUN chmod +x tss


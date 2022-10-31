from golang:1.18 as build-environment
WORKDIR /base
ADD . .
RUN cd tss && CGO_ENABLED=1 make build

from centos:8 as tss
WORKDIR /app
COPY --from=build-environment /base/tss/tss /app
RUN chmod +x tss



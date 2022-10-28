from golang:1.18 as build-environment
WORKDIR /base
ADD . .
RUN cd tss && CGO_ENABLED=1 make build

from centos:8 as tss-node
WORKDIR /app
COPY --from=build-environment /base/tss/tss /app
RUN chmod +x tss
CMD ["./tss", "node", "-c", "config/tss-node-config.toml"]


from centos:8 as tss-manager
WORKDIR /app
COPY --from=build-environment /base/tss/tss /app
RUN chmod +x tss
CMD ["./tss", "manager", "-c", "config/tss-manager-config.toml"]

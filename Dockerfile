# This Dockerfile builds all the dependencies needed by the monorepo, and should
# be used to build any of the follow-on services
#
# ### BASE: Install deps
FROM mantlenetworkio/foundry:latest as foundry
FROM node:16-alpine3.15 as base

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add curl \
    jq \
    python3 \
    ca-certificates \
    git \
    make \
    gcc \
    g++ \
    musl-dev \
    libusb-dev \
    eudev-dev \
    linux-headers \
    bash \
    build-base \
    gcompat

ENV GLIBC_KEY=https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
ENV GLIBC_KEY_FILE=/etc/apk/keys/sgerrand.rsa.pub
ENV GLIBC_RELEASE=https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.35-r0/glibc-2.35-r0.apk
ENV NODE_OPTIONS="--max-old-space-size=8192"

RUN wget -q -O ${GLIBC_KEY_FILE} ${GLIBC_KEY} \
    && wget -O glibc.apk ${GLIBC_RELEASE} \
    && apk add glibc.apk --force

COPY --from=foundry /usr/local/bin/forge /usr/local/bin/forge
COPY --from=foundry /usr/local/bin/cast /usr/local/bin/cast

# copy over the needed configs to run the dep installation
# note: this approach can be a bit unhandy to maintain, but it allows
# us to cache the installation steps
WORKDIR /opt/mantle
COPY *.json yarn.lock ./
COPY packages/sdk/package.json ./packages/sdk/package.json
COPY packages/core-utils/package.json ./packages/core-utils/package.json
COPY packages/common-ts/package.json ./packages/common-ts/package.json
COPY packages/contracts/package.json ./packages/contracts/package.json
COPY packages/data-transport-layer/package.json ./packages/data-transport-layer/package.json
COPY packages/hardhat-deploy-config/package.json ./packages/hardhat-deploy-config/package.json
COPY packages/message-relayer/package.json ./packages/message-relayer/package.json
COPY packages/replica-healthcheck/package.json ./packages/replica-healthcheck/package.json
COPY integration-tests/package.json ./integration-tests/package.json

RUN yarn install --frozen-lockfile && yarn cache clean

COPY ./packages ./packages
COPY ./integration-tests ./integration-tests

# build it!
RUN cd packages/contracts && yarn build
RUN cd ../../
RUN yarn build

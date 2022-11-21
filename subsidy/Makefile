SHELL := /bin/bash

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS :=-ldflags "$(LDFLAGSSTRING)"

CONTRACTS_PATH := "../packages/contracts/artifacts/contracts"

subsidy:
	env GO111MODULE=on go build $(LDFLAGS)
.PHONY: subsidy

clean:
	rm subsidy

test:
	go test -v ./...

lint:
	golangci-lint run ./...


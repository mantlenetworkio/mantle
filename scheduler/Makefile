GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

CTC_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/L1/rollup/CanonicalTransactionChain.sol/CanonicalTransactionChain.json
SCC_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/L1/rollup/StateCommitmentChain.sol/StateCommitmentChain.json

SEQUENCER_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/L1/sequencer/Sequencer.sol/Sequencer.json

scheduler:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/scheduler

clean:
	rm scheduler

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings: bindings-ctc bindings-scc bindings-sequencer

bindings-ctc:
	$(eval temp := $(shell mktemp))

	cat $(CTC_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(CTC_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg ctc \
		--abi - \
		--out bindings/ctc/canonical_transaction_chain.go \
		--type CanonicalTransactionChain \
		--bin $(temp)

	rm $(temp)

bindings-scc:
	$(eval temp := $(shell mktemp))

	cat $(SCC_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(SCC_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg scc \
		--abi - \
		--out bindings/scc/state_commitment_chain.go \
		--type StateCommitmentChain \
		--bin $(temp)

	rm $(temp)

bindings-sequencer:
	$(eval temp := $(shell mktemp))

	cat $(SEQUENCER_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(SEQUENCER_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg seq \
		--abi - \
		--out bindings/sequencer/sequencer.go \
		--type Sequencer \
		--bin $(temp)

	rm $(temp)

format:
	find . -name '*.go' -type f -not -path "./bindings*" | xargs gofumpt -w -l

.PHONY: \
	scheduler \
	bindings \
	bindings-ctc \
	bindings-scc \
	clean \
	test \
	lint \
  format

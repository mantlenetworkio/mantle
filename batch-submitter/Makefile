GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

CTC_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/L1/rollup/CanonicalTransactionChain.sol/CanonicalTransactionChain.json
SCC_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/L1/rollup/StateCommitmentChain.sol/StateCommitmentChain.json
DA_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/da/BVM_EigenDataLayrChain.sol/BVM_EigenDataLayrChain.json

batch-submitter:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/batch-submitter

clean:
	rm batch-submitter

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings: bindings-ctc bindings-scc bindings-da

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

bindings-da:
	$(eval temp := $(shell mktemp))

	cat $(DA_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(DA_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg da \
		--abi - \
		--out bindings/da/bvm_eigen_datalayr_chain.go \
		--type BVM_EigenDataLayrChain \
		--bin $(temp)

	rm $(temp)

.PHONY: \
	batch-submitter \
	bindings \
	bindings-ctc \
	bindings-scc \
	bindings-da \
	clean \
	test \
	lint

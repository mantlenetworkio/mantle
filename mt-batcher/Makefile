GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

EDLC_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/da/BVM_EigenDataLayrChain.sol/BVM_EigenDataLayrChain.json
EDLF_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/da/BVM_EigenDataLayrFee.sol/BVM_EigenDataLayrFee.json

mt-batcher:
	go mod tidy
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/mt-batcher

datalayr-mantle:
	cd ../datalayr-mantle/contracts && ./compile.sh compile-el && ./compile.sh compile-rollup

clean:
	rm mt-batcher

test:
	go test -v ./...

lint:
	golangci-lint run ./...

binding-edc:
	$(eval temp := $(shell mktemp))

	cat $(EDLC_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(EDLC_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/bvm_eigen_datalayr_chain.go \
		--type BVM_EigenDataLayrChain \
		--bin $(temp)

		rm $(temp)

binding-edf:
	  $(eval temp := $(shell mktemp))

	  cat $(EDLF_ABI_ARTIFACT) \
    		| jq -r .bytecode > $(temp)

		cat $(EDLF_ABI_ARTIFACT) \
    		| jq .abi \
    		| abigen --pkg bindings \
    		--abi - \
    		--out bindings/bvm_eigen_datalayr_fee.go \
    		--type BVM_EigenDataLayrFee \
    		--bin $(temp)

		rm $(temp)


.PHONY: \
	mt-batcher \
	binding \
	clean \
	test \
	lint

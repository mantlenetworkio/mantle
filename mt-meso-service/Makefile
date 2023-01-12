GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

EDLC_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/da/BVM_EigenDataLayrChain.sol/BVM_EigenDataLayrChain.json

mt-meso-service:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/mt-meso-service

datalayr-mantle:
	cd ../datalayr-mantle/contracts && ./compile.sh compile-el && ./compile.sh compile-rollup

clean:
	rm mt-meso-service

test:
	go test -v ./...

lint:
	golangci-lint run ./...

binding:
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


.PHONY: \
	mt-meso-service \
	binding \
	clean \
	test \
	lint

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

build:
	env GO111MODULE=on CGO_ENABLED=1 go build -v $(LDFLAGS) .

test:
	go test -v ./...

lint:
	golangci-lint run ./...

TSS_GROUP_MANAGER_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/L1/tss/TssGroupManager.sol/TssGroupManager.json
TSS_STAKING_SLASHING_ABI_ARTIFACT := ../packages/contracts/artifacts/contracts/L1/tss/TssStakingSlashing.sol/TssStakingSlashing.json

bindings: bindings-tss-group bindings-tss-staking-slashing

bindings-tss-group :
	$(eval temp := $(shell mktemp))

	cat $(TSS_GROUP_MANAGER_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(TSS_GROUP_MANAGER_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg tgm \
		--abi - \
		--out bindings/tgm/tss_group_manager.go \
		--type TssGroupManager \
		--bin $(temp)

	rm $(temp)

bindings-tss-staking-slashing :
	$(eval temp := $(shell mktemp))

	cat $(TSS_STAKING_SLASHING_ABI_ARTIFACT) \
		| jq -r .bytecode > $(temp)

	cat $(TSS_STAKING_SLASHING_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg tsh \
		--abi - \
		--out bindings/tsh/tss_staking-slashing.go \
		--type TssStakingSlashing \
		--bin $(temp)

	rm $(temp)

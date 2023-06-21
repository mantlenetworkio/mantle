COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545
LINT_PATH=./ci-lint
COVERAGE_PATH=./coverage

build: build-ts
.PHONY: build

build-ts: submodules
	if [ -n "$$NVM_DIR" ]; then \
		. $$NVM_DIR/nvm.sh && nvm use; \
	fi
	yarn install
	yarn build
.PHONY: build-ts

submodules:
	# CI will checkout submodules on its own (and fails on these commands)
	if [ -z "$$GITHUB_ENV" ]; then \
		git submodule init; \
		git submodule update; \
	fi
.PHONY: submodules

clean-node-modules:
	rm -rf node_modules
	rm -rf packages/**/node_modules

mod-tidy:
	cd ./batch-submitter && go mod tidy && cd .. && \
	cd ./bss-core && go mod tidy && cd ..  && \
	cd ./gas-oracle && go mod tidy && cd ..  && \
	cd ./l2geth && go mod tidy && cd ..  && \
	cd ./l2geth-exporter && go mod tidy && cd ..
	cd ./mt-batcher && go mod tidy && cd ..
.PHONY: mod-tidy

ci: ci-batch-submitter ci-fraud-proof ci-gas-oracle ci-l2geth ci-mt-batcher ci-mt-challenger ci-tss
	@echo
.PHONY: ci

ci-batch-submitter:
	mkdir -p ${LINT_PATH}
	rm -rf ${LINT_PATH}/batch-submitter.ci.out
	touch ${LINT_PATH}/batch-submitter.ci.out
	echo batch-submitter ci path: ${LINT_PATH}/batch-submitter.ci.out
	cd batch-submitter && golangci-lint run ./... > ../${LINT_PATH}/batch-submitter.ci.out || true && cd -
.PHONY: ci-batch-submitter

ci-fraud-proof:
	mkdir -p ${LINT_PATH}
	rm -rf ${LINT_PATH}/fraud-proof.ci.out
	touch ${LINT_PATH}/fraud-proof.ci.out
	echo fraud-proof ci path: ${LINT_PATH}/fraud-proof.ci.out
	cd fraud-proof && golangci-lint run ./... > ../${LINT_PATH}/fraud-proof.ci.out || true && cd -
.PHONY: ci-fraud-proof

ci-gas-oracle:
	mkdir -p ${LINT_PATH}
	rm -rf ${LINT_PATH}/gas-oracle.ci.out
	touch ${LINT_PATH}/gas-oracle.ci.out
	echo gas-oracle ci path: ${LINT_PATH}/gas-oracle.ci.out
	cd gas-oracle && golangci-lint run ./... > ../${LINT_PATH}/gas-oracle.ci.out || true && cd -
.PHONY: ci-gas-oracle

ci-l2geth:
	mkdir -p ${LINT_PATH}
	rm -rf ${LINT_PATH}/l2geth.ci.out
	touch ${LINT_PATH}/l2geth.ci.out
	echo l2geth ci path: ${LINT_PATH}/l2geth.ci.out
	cd l2geth && golangci-lint run ./... > ../${LINT_PATH}/l2geth.ci.out || true && cd -
.PHONY: ci-l2geth

ci-mt-batcher:
	mkdir -p ${LINT_PATH}
	rm -rf ${LINT_PATH}/mt-batcher.ci.out
	touch ${LINT_PATH}/mt-batcher.ci.out
	echo mt-batcher ci path: ${LINT_PATH}/mt-batcher.ci.out
	cd mt-batcher && golangci-lint run ./... > ../${LINT_PATH}/mt-batcher.ci.out || true && cd -
.PHONY: ci-mt-batcher

ci-mt-challenger:
	mkdir -p ${LINT_PATH}
	rm -rf ${LINT_PATH}/mt-challenger.ci.out
	touch ${LINT_PATH}/mt-challenger.ci.out
	echo mt-challenger ci path: ${LINT_PATH}/mt-challenger.ci.out
	cd mt-challenger && golangci-lint run ./... > ../${LINT_PATH}/mt-challenger.ci.out || true && cd -
.PHONY: ci-mt-challenger

ci-tss:
	mkdir -p ${LINT_PATH}
	rm -rf ${LINT_PATH}/tss.ci.out
	touch ${LINT_PATH}/tss.ci.out
	echo tss ci path: ${LINT_PATH}/tss.ci.out
	cd tss && golangci-lint run ./... > ../${LINT_PATH}/tss.ci.out || true && cd -
.PHONY: ci-tss

cover: cover-batch-submitter cover-fraud-proof cover-gas-oracle cover-mt-batcher cover-mt-challenger cover-tss
	@echo
.PHONY: cover

cover-batch-submitter:
	@mkdir -p ${COVERAGE_PATH}
	@rm -rf ${COVERAGE_PATH}/batch-submitter.cover.out
	@touch ${COVERAGE_PATH}/batch-submitter.cover.out
	@echo module coverage path: ${COVERAGE_PATH}/batch-submitter.cover.out
	@cd batch-submitter && go test ./... -coverprofile=cover.out && go tool cover -func=cover.out > ../${COVERAGE_PATH}/batch-submitter.cover.out || true && rm cover.out && find . -name 'logs'| xargs rm -r && cd -
.PHONY: cover-batch-submitter

cover-fraud-proof:
	@mkdir -p ${COVERAGE_PATH}
	@rm -rf ${COVERAGE_PATH}/fraud-proof.cover.out
	@touch ${COVERAGE_PATH}/fraud-proof.cover.out
	@echo module coverage path: ${COVERAGE_PATH}/fraud-proof.cover.out
	@cd fraud-proof && go test ./... -coverprofile=cover.out && go tool cover -func=cover.out > ../${COVERAGE_PATH}/fraud-proof.cover.out || true && rm cover.out && find . -name 'logs'| xargs rm -r && cd -
.PHONY: cover-fraud-proof

cover-gas-oracle:
	@mkdir ${COVERAGE_PATH}
	@rm -rf ${COVERAGE_PATH}/gas-oracle.cover.out
	@touch ${COVERAGE_PATH}/gas-oracle.cover.out
	@echo module coverage path: ${COVERAGE_PATH}/gas-oracle.cover.out
	@cd gas-oracle && go test ./... -coverprofile=cover.out && go tool cover -func=cover.out > ../${COVERAGE_PATH}/gas-oracle.cover.out || true && rm cover.out && find . -name 'logs'| xargs rm -r && cd -
.PHONY: cover-gas-oracle

cover-mt-batcher:
	@mkdir ${COVERAGE_PATH}
	@rm -rf ${COVERAGE_PATH}/mt-batcher.cover.out
	@touch ${COVERAGE_PATH}/mt-batcher.cover.out
	@echo module coverage path: ${COVERAGE_PATH}/mt-batcher.cover.out
	@cd mt-batcher && go test ./... -coverprofile=cover.out && go tool cover -func=cover.out > ../${COVERAGE_PATH}/mt-batcher.cover.out  || true && rm cover.out && find . -name 'logs'| xargs rm -r && cd -
.PHONY: cover-mt-batcher

cover-mt-challenger:
	@mkdir ${COVERAGE_PATH}
	@rm -rf ${COVERAGE_PATH}/mt-challenger.cover.out
	@touch ${COVERAGE_PATH}/mt-challenger.cover.out
	@echo module coverage path: ${COVERAGE_PATH}/mt-challenger.cover.out
	@cd mt-challenger && go test ./... -coverprofile=cover.out && go tool cover -func=cover.out > ../${COVERAGE_PATH}/mt-challenger.cover.out && rm cover.out && find . -name 'logs'| xargs rm -r && cd -
.PHONY: cover-mt-challenger

cover-tss:
	@mkdir ${COVERAGE_PATH}
	@rm -rf ${COVERAGE_PATH}/tss.cover.out
	@touch ${COVERAGE_PATH}/tss.cover.out
	@echo module coverage path: ${COVERAGE_PATH}/tss.cover.out
	@cd tss && go test ./... -coverprofile=cover.out && go tool cover -func=cover.out > ../${COVERAGE_PATH}/tss.cover.out || true && rm tss.cover.out && find . -name 'logs'| xargs rm -r && cd -
.PHONY: cover-tss

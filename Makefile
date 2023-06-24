COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545
LINT_PATH=./ci-lint
COVERAGE_PATH=./coverage

## color codes
C_RESET=\033[0m
C_RESET_UNDERLINE=\033[24m
C_RESET_REVERSE=\033[27m
C_DEFAULT=\033[39m
C_DEFAULTB=\033[49m
C_BOLD=\033[1m
C_BRIGHT=\033[2m
C_UNDERSCORE=\033[4m
C_REVERSE=\033[7m
C_BLACK=\033[30m
C_RED=\033[31m
C_GREEN=\033[32m
C_BROWN=\033[33m
C_BLUE=\033[34m
C_MAGENTA=\033[35m
C_CYAN=\033[36m
C_WHITE=\033[37m

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

getdeps:
	@printf "${C_BROWN}${C_BOLD}>> Checking golangci-lint: \n${C_RESET}"

	@printf "${C_GREEN}${C_BOLD}>> golangci-lint installed! \n${C_RESET}"
.PHONY: getdeps

before-ci:
	@which golangci-lint 1>/dev/null || (echo "Installing golangci-lint" && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3)
	@printf "${C_BROWN}${C_BOLD}>> Checking cache: \n${C_RESET}"
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint cache status
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint cache clean
	@printf "${C_GREEN}${C_BOLD}>> cache cleaned! \n${C_RESET}"
.PHONY: before-ci

ci: getdeps ci-batch-submitter ci-fraud-proof ci-gas-oracle ci-l2geth ci-mt-batcher ci-mt-challenger ci-tss
	@echo
.PHONY: ci

ci-batch-submitter: before-ci
	@printf "${C_BROWN}${C_BOLD}>> ci-batch-submitter golangci-lint... \n${C_RESET}"
	mkdir -p ${LINT_PATH}
	echo batch-submitter ci path: ${LINT_PATH}/batch-submitter.ci.out
	cd batch-submitter && ${GOPATH}/bin/golangci-lint run ./... --timeout=5m --config ../.golangci.yml > ../${LINT_PATH}/batch-submitter.ci.out || true && cd -
	@printf "${C_GREEN}${C_BOLD}>> ci-batch-submitter finished! \n${C_RESET}"
.PHONY: ci-batch-submitter

ci-fraud-proof: before-ci
	@printf "${C_BROWN}${C_BOLD}>> ci-batch-submitter... \n${C_RESET}"
	mkdir -p ${LINT_PATH}
	echo fraud-proof ci path: ${LINT_PATH}/fraud-proof.ci.out
	cd fraud-proof && ${GOPATH}/bin/golangci-lint run ./... --timeout=5m --config ../.golangci.yml > ../${LINT_PATH}/fraud-proof.ci.out || true && cd -
	@printf "${C_GREEN}${C_BOLD}>> ci-fraud-proof finished! \n${C_RESET}"
.PHONY: ci-fraud-proof

ci-gas-oracle: before-ci
	@printf "${C_BROWN}${C_BOLD}>> ci-batch-submitter... \n${C_RESET}"
	mkdir -p ${LINT_PATH}
	echo gas-oracle ci path: ${LINT_PATH}/gas-oracle.ci.out
	cd gas-oracle && ${GOPATH}/bin/golangci-lint run ./... --timeout=5m --config ../.golangci.yml > ../${LINT_PATH}/gas-oracle.ci.out || true && cd -
	@printf "${C_GREEN}${C_BOLD}>> ci-gas-oracle finished! \n${C_RESET}"
.PHONY: ci-gas-oracle

ci-l2geth: before-ci
	@printf "${C_BROWN}${C_BOLD}>> ci-batch-submitter... \n${C_RESET}"
	mkdir -p ${LINT_PATH}
	echo l2geth ci path: ${LINT_PATH}/l2geth.ci.out
	cd l2geth && ${GOPATH}/bin/golangci-lint run ./... --timeout=5m --config ../.golangci.yml > ../${LINT_PATH}/l2geth.ci.out || true && cd -
	@printf "${C_GREEN}${C_BOLD}>> ci-l2geth finished! \n${C_RESET}"
.PHONY: ci-l2geth

ci-mt-batcher: before-ci
	@printf "${C_BROWN}${C_BOLD}>> ci-batch-submitter... \n${C_RESET}"
	mkdir -p ${LINT_PATH}
	echo mt-batcher ci path: ${LINT_PATH}/mt-batcher.ci.out
	cd mt-batcher && ${GOPATH}/bin/golangci-lint run ./... --timeout=5m --config ../.golangci.yml > ../${LINT_PATH}/mt-batcher.ci.out || true && cd -
	@printf "${C_GREEN}${C_BOLD}>> ci-mt-batcher finished! \n${C_RESET}"
.PHONY: ci-mt-batcher

ci-mt-challenger: before-ci
	@printf "${C_BROWN}${C_BOLD}>> ci-batch-submitter... \n${C_RESET}"
	mkdir -p ${LINT_PATH}
	echo mt-challenger ci path: ${LINT_PATH}/mt-challenger.ci.out
	cd mt-challenger && ${GOPATH}/bin/golangci-lint run ./... --timeout=5m --config ../.golangci.yml > ../${LINT_PATH}/mt-challenger.ci.out || true && cd -
	@printf "${C_GREEN}${C_BOLD}>> ci-mt-challenger finished! \n${C_RESET}"
.PHONY: ci-mt-challenger

ci-tss: before-ci
	@printf "${C_BROWN}${C_BOLD}>> ci-batch-submitter... \n${C_RESET}"
	mkdir -p ${LINT_PATH}
	echo tss ci path: ${LINT_PATH}/tss.ci.out
	cd tss && ${GOPATH}/bin/golangci-lint run ./... --timeout=5m --config ../.golangci.yml > ../${LINT_PATH}/tss.ci.out || true && cd -
	@printf "${C_GREEN}${C_BOLD}>> ci-tss finished! \n${C_RESET}"
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
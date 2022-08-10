COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545

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

# Remove the baseline-commit to generate a base reading & show all issues
#semgrep:
#	$(eval DEV_REF := $(shell git rev-parse develop))
#	SEMGREP_REPO_NAME=bitdao-io/bitnetwork semgrep ci --baseline-commit=$(DEV_REF)
#.PHONY: semgrep


clean-node-modules:
	rm -rf node_modules
	rm -rf packages/**/node_modules

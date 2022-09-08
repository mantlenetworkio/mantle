COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545

compile:
	bash ./compile.sh
.PHONY: compile

test: compile
	go test ./
.PHONY: test

.PHONY: test

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## install the gocc binary
	go install .

test: ## run all unit tests
	go test -v ./...

gofmt: ## format all go files
	gofmt -l -s -w .

govet:
	go vet ./...

errcheck:
	go get github.com/kisielk/errcheck@v1.2.0
	errcheck -exclude .errcheck-ignore ./...

lint: govet errcheck ## run linter checks

goimports:  ## sort all imports
	go get golang.org/x/tools/cmd/goimports
	goimports -l -w .

regenerate: ## regenerate all example and test files
	make install
	make -C example regenerate
	make -C internal/test regenerate
	make gofmt
	make goimports
	make lint

check: ## regenerate, lint and run a terse version of check
	@# quietly install and regenerate
	@make --quiet install
	@make --quiet -C example regenerate
	@make --quiet -C internal/test regenerate
	@# promote formatting changes to errors
	@if [ -n "$(gofmt -l -s .)" ]; then { echo "gofmt errors:"; gofmt -d -l -s . ; exit 1; }; fi
	@if [ -n "$(gomports -l . )" ]; then { echo "goimports errors:"; goimports -l . ; exit 1; }; fi
	@make --quiet lint
	@go test ./... | grep -v "\[no test"

ci: ## run all ci checks
	make regenerate
	@git diff --exit-code . || { echo "ERROR: commit and working copy differ after 'make regenerate'"; exit 22 ; }
	make test
	golangci-lint run
	@git diff --exit-code . || { echo "ERROR: commit and working copy differ after 'make ci'" ; exit 2  ; }

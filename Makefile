.PHONY: test

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## install the gocc binary
	go install .

test: ## run all unit tests
	go test -v ./...

gofmt: ## format all go files
	gofmt -l -s -w .

govet: ## run go's code vetting on all code
	go vet ./...

ci-lint: ## see https://golangci-lint.run/, applies .golangci.yml
	golangci-lint run

lint: govet ci-lint

goclean: gofmt ## apply go style rules to source

regenerate: ## regenerate all example and test files
	make install
	make -C example regenerate
	make -C internal/test regenerate
	make goclean
	# make lint

check: ## regenerate, lint and run a terse version of check
	@# quietly install and regenerate
	@go mod tidy
	@make --quiet install
	@make --quiet -C example regenerate
	@make --quiet -C internal/test regenerate
	@# promote formatting changes to errors
	@if [ -n "$(gofmt -l -s .)" ]; then { echo "gofmt errors:"; gofmt -d -l -s . ; exit 1; }; fi
	@make --quiet lint
	@go test ./... | grep -v "\[no test"

ci: ## run all ci checks
	make regenerate
	@git diff --exit-code . || { echo "ERROR: commit and working copy differ after 'make regenerate'"; exit 22 ; }
	make test
	@git diff --exit-code . || { echo "ERROR: commit and working copy differ after 'make test'" ; exit 22 ; }

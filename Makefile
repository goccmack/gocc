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
	go vet -methods=false .

errcheck:
	go get github.com/kisielk/errcheck@v1.2.0
	errcheck -exclude .errcheck-ignore ./...

lint: govet errcheck ## run linter checks

goimports:  ## sort all imports
	go get golang.org/x/tools/cmd/goimports
	goimports -w .

regenerate: ## regenerate all example and test files
	make install
	make -C example regenerate
	make -C internal/test regenerate
	make gofmt
	make goimports

ci: ## run all ci checks
	make regenerate
	make test
	golangci-lint run
	git diff --exit-code .

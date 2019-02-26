.PHONY: test

install:
	go install .

test:
	go test -v ./...

gofmt:
	gofmt -l -s -w .

govet:
	go vet -methods=false .

goimports:
	go get golang.org/x/tools/cmd/goimports
	goimports -w .

regenerate:
	make install
	make -C example regenerate
	make -C internal/test regenerate
	make gofmt

travis:
	make install
	make regenerate
	make govet
	make goimports
	make errcheck
	make test
	git diff --exit-code .

errcheck:
	go get github.com/kisielk/errcheck
	errcheck -exclude .errcheck-ignore ./...

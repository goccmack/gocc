.PHONY: test

install:
	go install .

test:
	go test -v ./...

gofmt:
	gofmt -l -s -w .

govet:
	go tool vet -methods=false .

regenerate:
	make install
	make -C example regenerate
	make -C internal/test regenerate
	make gofmt

travis:
	make install
	make regenerate
	go get golang.org/x/tools/cmd/goimports
	goimports -w .
	make errcheck
	git diff --exit-code .

errcheck:
	go get github.com/kisielk/errcheck
	errcheck ./...

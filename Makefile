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
	(cd example && make regenerate)
	(cd test && make regenerate)
	make gofmt

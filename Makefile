.PHONY: test

build:
	go build -v ./...

test:
	go test -v ./...

gofmt:
	gofmt -l -s -w .

govet:
	go vet ./...

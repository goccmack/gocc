build:
	go build -v ./...

TESTDIRS = test example

test:
	for d in $(TESTDIRS); do $(MAKE) -B -C $$d ; done
	go test -v ./...

gofmt:
	gofmt -l -s -w .


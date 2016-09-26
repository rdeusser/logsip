.PHONY: all fmt deps test

all: fmt deps test

fmt:
	go fmt `go list ./...`

deps:
	go get -t -v ./...

test:
	go tool vet .
	go test -v -race ./...

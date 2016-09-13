.PHONY: all fmt test build

all: fmt test build

fmt:
	go fmt `go list ./...`

test:
	go get -t -v ./...
	go tool vet .
	go test -v -race ./...

TEST?=$(shell go list ./... | grep -v vendor)

all: fmt test

fmt:
	go fmt `go list ./... | grep -v vendor`

test:
	@go test
	@go vet $(TEST); if [ $$? -eq 1 ]; then \
		echo "ERROR: Vet found problems in the code."; \
		exit 1; \
	fi

.PHONY: all fmt test

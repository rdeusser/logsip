TEST?=$(shell go list ./... | grep -v vendor)

all: fmt deps test

fmt:
	go fmt `go list ./... | grep -v vendor`

deps:
	@godep restore
	@godep save

test:
	@go test
	@go vet $(TEST); if [ $$? -eq 1 ]; then \
		echo "ERROR: Vet found problems in the code."; \
		exit 1; \
	fi

.PHONY: all fmt deps test

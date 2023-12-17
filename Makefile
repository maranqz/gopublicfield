.PHONY: fmt lint clean.test test test.clean

CVPKG=go list ./... | grep -v mocks | grep -v internal/
GO_TEST=go test `$(CVPKG)` -race
COVERAGE_FILE="coverage.out"

all: fmt lint test install

fmt:
	go fmt ./...

lint:
	golangci-lint run --fix ./...

clean.test:
	go clean --testcache

test:
	go test

test.clean: clean.test test

test.coverage:
	$(GO_TEST) -covermode=atomic -coverprofile=$(COVERAGE_FILE)

install:
	go install ./cmd/gofactory

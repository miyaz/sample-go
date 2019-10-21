# メタ情報
REVISION := $(shell git rev-parse --short HEAD)
BUILD_AT := $(shell date +%Y%m%d%H%M%S)
LDFLAGS := -ldflags="-s -w -X \"main.revision=$(REVISION)\" -X \"main.buildAt=$(BUILD_AT)\" -extldflags \"-static\""

export GO111MODULE=on

# 必要なツール類をセットアップする
## Install for Development
.PHONY: devel-deps
devel-deps: deps
	GO111MODULE=off go get -u        \
	golang.org/x/lint/golint         \
	golang.org/x/tools/cmd/goimports

## Clean binaries
.PHONY: clean
clean:
	rm -rf bin/*

## Run tests
.PHONY: test
test: deps
	go test -race -cover -v ./...

## Install dependencies
.PHONY: deps
deps:
	go get ./...

## Update dependencies
.PHONY: update
update:
	go get -u -d ./...
	go mod tidy

## Run Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## Format source codes
.PHONY: fmt
fmt: devel-deps
	goimports -w ./...

## Build binaries ex. make bin/sample-go
bin/%: *.go
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o $@ $^


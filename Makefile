GO_BIN := $(GOPATH)/bin
GOIMPORTS := $(GO_BIN)/goimports
GOLINT := $(GO_BIN)/golangci-lint

all: install

install: fmt
	go install -v

test:
	go test ./... -v

lint: $(GOLINT)
	golangci-lint run

fmt: $(GOIMPORTS)
	goimports -w -l .

$(GOIMPORTS):
	go get -u golang.org/x/tools/cmd/goimports

$(GOLINT):
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: install test fmt lint

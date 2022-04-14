.PHONY: install build test

BINDIR=bin
INSTALL_PATH=github.com/ritarock/quickquery

install:
	go install $(INSTALL_PATH)

build:
	go build -o $(BINDIR)/quickquery cmd/quickquery/main.go

test:
	go test $(shell go list ./... |grep internal)

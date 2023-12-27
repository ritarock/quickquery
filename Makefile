PHONY: test build install

BINDIR=bin

test:
	go test ./...

build:
	go build -o ${BINDIR}/qq .

install:
	go install

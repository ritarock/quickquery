MOCK_GEN = go run -mod=mod go.uber.org/mock/mockgen@latest

BINDIR=bin

mock:
	$(MOCK_GEN) \
	-source=./application/port.go \
	-destination=./testing/mock/application.go \
	-package=mock

test:
	go test ./...

build:
	go build -o ${BINDIR}/quickquery .

install:
	go install .

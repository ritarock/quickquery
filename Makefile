MOCK_GEN = go run -mod=mod go.uber.org/mock/mockgen@latest

mock:
	$(MOCK_GEN) \
	-source=./application/port.go \
	-destination=./testing/mock/application.go \
	-package=mock

test:
	go test ./...

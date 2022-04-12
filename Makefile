PWD=$(shell pwd)

install:
	cd $(PWD)/cmd/quickquery;go install

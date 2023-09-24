GOCMD:=$(shell which go)

test:
	@$(GOCMD) test -v ./...

coverage:
	@$(GOCMD) test -v ./... -coverprofile=coverage.txt -covermode=atomic
	@$(GOCMD) tool cover -func=coverage.txt
	@$(GOCMD) tool cover -html=coverage.txt

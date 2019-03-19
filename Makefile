MODULES :=$(shell go list ./... | grep -v example)

default: vet fmt

ci: vet fmt

vet: 
	@go vet $(MODULES)

fmt:
	@go fmt $(MODULES)

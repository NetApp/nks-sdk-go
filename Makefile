MODULES :=$(shell go list ./... | grep -v example)

default: vet fmt

ci: vet fmt

vet: 
	@go vet $(MODULES)

fmt:
	@go fmt $(MODULES)

test:
	cd nks && NKS_TEST_ENV=mock go test ./... -v

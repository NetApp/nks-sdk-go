MODULES :=$(shell go list ./... | grep -v example)

default: vet fmt

ci: vet fmt mock_test

vet: 
	@go vet $(MODULES)

fmt:
	@go fmt $(MODULES)

mock_test:
	cd nks && NKS_TEST_ENV=mock go test ./... -v

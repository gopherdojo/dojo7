# CONST
BINARYNAME=imachan

export GO111MODULE=auto

# command
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## go build
	go build -v ./cmd/${BINARYNAME}

.PHONY: test
test: ## go test
	go test -v -cover ./pkg/${BINARYNAME}

.PHONY: clean
clean: ## go clean
	go clean -cache -testcache

.PHONY: analyze
analyze: ## do static code analysis
	goimports -l -w .
	go vet ./...
	golint ./...

.PHONY: remove
remove: ## remove binary and test output data
	rm -f ./${BINARYNAME}
	find testdata -type f -not -name *.jpg | xargs rm

.PHONY: all
all: remove clean test analyze build ## run 'build' with 'remove', 'clean', 'test' and 'analyze'

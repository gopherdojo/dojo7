.PHONY: fmt
fmt: ## Format source code
	@which goimports > /dev/null || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	@goimports -d -w $(shell find . -type f -name '*.go')
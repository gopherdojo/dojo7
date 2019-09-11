GOCMD=go
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOLINT=golint
GOTEST=$(GOCMD) test
BINARYNAME=bin/imgconv

export GO111MODULE=off

.PHONY: build
build: ## go build
	go build -o ${BINARYNAME} cmd/imgconv/main.go
fmt:
	${GOFMT} ./...
	${GOVET} ./...
	${GOLINT} ./...

check:
	${GOTEST} ./... -v
	${GOTEST} ./... -cover

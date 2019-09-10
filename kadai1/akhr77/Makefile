GOCMD=go
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOLINT=golint
GOTEST=$(GOCMD) test

export GO111MODULE=on

all:imgconv

imgconv:
	go build -o imgconv -v

fmt:
	${GOFMT} ./...
	${GOVET} ./...
	${GOLINT} ./...

check:
	${GOTEST} ./... -v
	${GOTEST} ./... -cover

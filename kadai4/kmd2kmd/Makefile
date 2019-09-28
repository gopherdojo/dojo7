.PHONY: test
test:
	@ go test -race -v ./...

.PHONY: cover
cover:
	@go test -race -v -coverprofile=cover.out ./...
	@go tool cover -html=cover.out -o cover.html

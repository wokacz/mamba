# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## format: format .go files
.PHONY: format
format:
	go fmt ./...

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## build: build the binary
.PHONY: build
build:
	go build -o ./bin/mamba ./cmd/cli/...

## sha256: generate sha256 checksum
.PHONY: sha256
sha256:
	@shasum -a 256 ./bin/mamba.tar.gz

## tar: create a tarball
.PHONY: tar
tar:
	@tar -czvf ./bin/mamba.tar.gz ./bin/mamba

## prepare: prepare the binary
.PHONY: prepare
prepare: build tar sha256

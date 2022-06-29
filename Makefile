PKG := "github.com/jettjia/go-ddd"
PKG_LIST := $(shell go list ${PKG}/...)
APP=go-ddd
VERSION=1.0.0

.PHONY: tidy
tidy:
	$(eval files=$(shell find . -name go.mod))
	@set -e; \
	for file in ${files}; do \
		goModPath=$$(dirname $$file); \
		cd $$goModPath; \
		go mod tidy; \
		cd -; \
	done

.PHONY: fmt
fmt:
	@go fmt ${PKG_LIST}

init: # install golint
	@go install golang.org/x/lint/golint@latest

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

.PHONY: vet
vet: ## Vet the files
	@go vet ${PKG_LIST}

.PHONY: test
test:
	@go test -cover ./...

race: ## Run tests with data race detector
	@go test -race ${PKG_LIST}

.PHONY: build
build:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o ./bin/linux_amd64/${APP}.${VERSION} main.go
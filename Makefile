PKG := "github.com/jett/gin-ddd"
PKG_LIST := $(shell go list ${PKG}/...)
APP=gin-ddd
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

.PHONY: vet
vet: ## Vet the files
	@go vet ${PKG_LIST}

.PHONY: test
test:
	go test -cover ./...

.PHONY: build
build:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o ./bin/linux_amd64/${APP}.${VERSION} main.go
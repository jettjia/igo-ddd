PKG := "github.com/jettjia/go-ddd-demo"
PKG_LIST := $(shell go list ${PKG}/...)
APP=GoDddDemo
DOCKER_IMG=registry.cn-hangzhou.aliyuncs.com/jett-cicd/gas
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

.PHONY: test-coverage
test-coverage:
	@go test ./... -v -coverprofile=report/cover 2>&1 | go-junit-report > report/ut_report.xml
	@gocov convert report/cover | gocov-html > report/coverage.html

.PHONY: golangci
golangci:
	@golangci-lint run --config .golangci.yml
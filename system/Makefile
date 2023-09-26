PKG := "jettjia/go-ddd-demo" # go mod 项目名称
PKG_LIST := $(shell go list ${PKG}/...)

# docker
## 镜像仓库+镜像名称,其中jidi-docker.pkg.coding.net/gas/go-ddd-demo/ 是镜像仓库; demo 是镜像名称
DOCKER_IMAGE_REPO := "jidi-docker.pkg.coding.net/gas/go-ddd-demo/demo:v1.0.0"
APP := "GoDddDemo" #go build后的可执行文件名称，和dockerfile保持一致

init: # 初始化安装 go的校验包
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/jstemmer/go-junit-report/v2@latest
	@go install github.com/matm/gocov-html/cmd/gocov-html@latest

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

.PHONY: generate
vet: ## generate the files
	@go generate ./...

.PHONY: test
test:
	@go test -cover ./...

.PHONY: race
race: ## Run tests with data race detector
	@go test -race ${PKG_LIST}

.PHONY: test-coverage
test-coverage:  ## Generate a single test report file in HTML format
	@go test ./... -v -coverprofile=report/cover 2>&1 | go-junit-report > report/ut_report.xml
	@gocov convert report/cover | gocov-html > report/coverage.html

.PHONY: golangci
golangci:
	@golangci-lint run --config .golangci.yml

.PHONY: build
build:
	@go build -o ${APP}

.PHONY: docker-build
docker-build: ## docker-build build项目镜像
	@docker build -t ${DOCKER_IMAGE_REPO}  .

.PHONY: docker-push
docker-push: ## docker-push push项目镜像到远程仓库
	@docker push ${DOCKER_IMAGE_REPO}

.PHONY: helm-build
helm-build: ## 打包helm
	@cd ${HELM_TEMP_DIR} \
 	&& tar zcvf ${HELM_IMAGE} ${HELM_DIR}

.PHONY: helm-push
helm-push: ## 推送helm到远程仓库
	@cd ${HELM_TEMP_DIR} \
	 && helm push ${HELM_IMAGE} ${HELM_IMAGE_REPO}
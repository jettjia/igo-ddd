PKG := "github.com/jettjia/ddddemo" # go mod 项目名称
PKG_LIST := $(shell go list ${PKG}/...)

# docker
DOCKER_IMAGE_REPO := "registry.cn-shanghai.aliyuncs.com/xdata-xtext/x-data:v1.0.0"
APP := "go-main" #go build后的可执行文件名称，和dockerfile保持一致

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

.PHONY: gbuild ## 构建go项目
gbuild:
	@GOWORK=off CGO_ENABLED=0 go build -mod=mod -ldflags="-s -w" -o deploy/bin/${APP} .

.PHONY: dbuild
dbuild: ## docker-build build项目镜像
	@docker build -f deploy/Dockerfile-ez -t ${DOCKER_IMAGE_REPO}  .

.PHONY: dpush
dpush: ## docker-push push项目镜像到远程仓库
	@docker push ${DOCKER_IMAGE_REPO}

.PHONY: hbuild
hbuild: ## 打包helm
	@cd ${HELM_TEMP_DIR} \
 	&& tar zcvf ${HELM_IMAGE} ${HELM_DIR}

.PHONY: hpush
hpush: ## 推送helm到远程仓库
	@cd ${HELM_TEMP_DIR} \
	 && helm push ${HELM_IMAGE} ${HELM_IMAGE_REPO}

.PHONY: buildAll
buildAll: ## 构建go包，构建镜像，推送镜像
	@make gbuild
	@make dbuild
	@make dpush
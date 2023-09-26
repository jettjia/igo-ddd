# helm
## helm地址
HELM_TEMP_DIR := "manifest/framework_tool/k8s/temp_helm" # 进入helm地址
HELM_IMAGE := "go-ddd-demo.tgz"
HELM_DIR := "go-ddd-demo"
HELM_IMAGE_REPO := "go-ddd-demo-helm" # helm远程仓库地址

.PHONY: init-helm
init-helm: ## 安装helm环境
	@cd /tmp && wget https://get.helm.sh/helm-v3.8.2-linux-amd64.tar.gz \
	&& tar -zxvf helm-v3.8.2-linux-amd64.tar.gz \
	&& cp linux-amd64/helm /usr/local/bin \
	&& helm version

.PHONY: install-rds
install-rds: ## 安装项目依赖的db
	@docker-compose -f deploy/framework_tool/mysql/docker-compose.yaml up -d

.PHONY: uninstall-rds
uninstall-rds: ## 卸载项目依赖的db
	@docker-compose -f deploy/framework_tool/mysql/docker-compose.yaml down

.PHONY: install-otel
install-otel: ## 安装项目的可观测性
	@docker-compose -f deploy/framework_tool/otel-jaeger/docker-compose.yaml up -d

.PHONY: uninstall-otel
uninstall-otel: ## 协助项目的可观测性
	@docker-compose -f deploy/framework_tool/otel-jaeger/docker-compose.yaml down
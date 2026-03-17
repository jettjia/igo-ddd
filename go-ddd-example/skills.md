# skills.md（开发速查）

本文档面向日常开发：快速理解项目结构、调用链路、扩展一个业务的落点、以及常用命令与约定。

## 1. 项目一句话

这是一个以 DDD 分层组织的 Go 服务示例，支持多入口形态（HTTP / gRPC / MCP / Event / Job），核心业务以“用户（sys_user）”为样板串起从接口层到数据落库的全链路。

入口文件：[main.go](igo-ddd/go-ddd-example/main.go)

## 2. 分层与目录（强约定）

分层说明以 [DDD.md](go-ddd-example/DDD.md) 为准，实际代码落点如下：

- api：接入层，适配协议（HTTP / gRPC / MCP / Event / Job），只做参数处理、鉴权、调用 application
- application：应用层，编排流程（组装 DTO / 调用 domain / 选择 infra）
- domain：领域层，领域模型与规则（entity / aggregate / domain service / repository interface）
- infra：基础设施层，数据库与外部依赖落地（repo 实现、PO、converter、DB/缓存/日志等封装）
- types：完全独立的通用类型与常量（错误码、常量等）
- manifest：运行时配置与国际化资源
- tools/apidoc：接口文档生成工具（可选）
- idl：proto 定义与生成代码（示例）

## 3. 启动与运行形态

### 3.1 启动链路（main）

[main.go](go-ddd-example/main.go) 依次做：

1) 读取 -env 并写入环境变量 env（默认 debug）
2) 自动建表：infra/repository/po.AutoTable
3) 初始化数据：boot.InitData
4) 启动 HTTP：api/http.InitHttp（当前仅 public HTTP）
5) 启动 gRPC：api/grpc.InitGrpc（受配置 enable_grpc 控制）
6) 启动 MCP：api/mcp.InitMCP（受配置 enable_mcp 控制）
7) 启动 Event：api/event.InitEvent（受配置 enable_event 控制）
8) 启动 Job：api/job.InitJob（受配置 enable_job 控制）

### 3.2 配置加载规则

配置入口：[config.go](go-ddd-example/config/config.go)

- 默认读取：./manifest/config/config.yaml
- env=release：读取 /sysvol/conf/public-center.yaml
- env=docker：读取 ./manifest/config/config-docker.yaml

默认配置文件：[config.yaml](go-ddd-example/manifest/config/config.yaml)

### 3.3 本地运行

```bash
go run main.go
go run main.go -env test
go run main.go -env release
```

## 4. HTTP：从路由到落库的标准链路（以 sys_user 为例）

### 4.1 路由

统一前缀在 [router.go](go-ddd-example/api/http/router/router.go)：

- /api/xtext/x-data/v1

用户路由在 [sys_router.go](go-ddd-example/api/http/router/public/sys_router.go)：

- POST   /user/user
- DELETE /user/user/:ulid
- PUT    /user/user/:ulid
- GET    /user/user/:ulid
- POST   /user/user/byQuery
- POST   /user/user/byAll
- POST   /user/userPage

### 4.2 Handler（接口层）

Handler 入口：[sys_user_handler.go](go-ddd-example/api/http/handler/public/user/sys_user_handler.go)

该层典型职责：

- BindJSON / BindUri
- validate.Validate 校验
- 将认证信息（如 user_id）写入请求 DTO（来自通用中间件上下文）
- 调用 application service
- 统一用 xresponse 返回，异常通过 c.Error(err) 进入全局 recover/错误处理链

常用中间件：

- Universal：[universal.go](go-ddd-example/api/http/middleware/universal.go)
- Token 校验（封装 hydra 方案）：[auth.go](go-ddd-example/api/http/middleware/auth.go)

HTTP Server 初始化：[http.go](go-ddd-example/api/http/http.go)

### 4.3 Application Service（流程编排）

样板：[sys_user_svc.go](go-ddd-example/application/service/user/sys_user_svc.go)

典型模式：

- assembler：DTO <-> Entity 转换
- 聚合（aggregate）：需要事务/跨 repo 的用聚合
- 领域服务（srv）：单 repo/纯查询型可走 domain service

### 4.4 Domain（领域层）

聚合样板：[sys_user_agg.go](go-ddd-example/domain/aggregate/user/sys_user_agg.go)

- 通过 idata.NewDataOptionCli 拿到 Transaction（ExecTx）
- 在事务中组合多个 repo 操作
- 领域错误用 xerror + types/apierror 统一定义

领域服务样板：[sys_user_svc.go](go-ddd-example/domain/srv/user/sys_user_svc.go)

实体样板：[sys_user_entity.go](go-ddd-example/domain/entity/user/sys_user_entity.go)

### 4.5 Infra Repo（持久化实现）

repo 实现样板：[sys_user_impl.go](go-ddd-example/infra/repository/repo/user/sys_user_impl.go)

常见结构：

- po：数据库映射对象（GORM tag / TableName / BeforeCreate）
- converter：Entity <-> PO
- repo：实现 domain/irepository 接口，内部用 data.DB(ctx) 执行

建表示例：

- 自动迁移：[auto_table.go](go-ddd-example/infra/repository/po/auto_table.go)
- sys_user PO：[sys_user_po.go](go-ddd-example/infra/repository/po/user/sys_user_po.go)

DB 连接注入：

- [data.go](go-ddd-example/infra/pkg/idata/data.go)

## 5. gRPC（可选启用）

开关：manifest/config/config.yaml 中 server.enable_grpc

Server 初始化：[grpc.go](go-ddd-example/api/grpc/grpc.go)

- 使用 go-grpc-middleware 统一接入认证与 recover
- AuthInterceptor 基于 JWT secret（见 consts）：[auth.go](go-ddd-example/api/grpc/middleware/auth.go)

proto 与生成代码示例在 idl/proto 下（不是必须依赖于 sys_user 样板）。

## 6. MCP（可选开启）

入口：[init.go](go-ddd-example/api/mcp/init.go)

主要内容：

- Tool 注册：[mcp.go](go-ddd-example/api/mcp/mcp.go)
- Tool Handler（调用 application service）：[handler_user.go](go-ddd-example/api/mcp/handler_user.go)

## 7. Event / Job（可按需扩展）

Event 开关：server.enable_event

- 初始化：[event.go](go-ddd-example/api/event/event.go)
- 订阅入口骨架：[sys_user_subscribe.go](go-ddd-example/api/event/subscribe/user/sys_user_subscribe.go)

Job 开关：server.enable_job（当前 InitJob 为空实现）

- [job.go](go-ddd-example/api/job/job.go)

## 8. 错误码与国际化

错误码定义在：[error_def.go](go-ddd-example/types/apierror/error_def.go)

建议的使用方式：

- 业务错误：在 domain（entity/aggregate/srv）构造 xerror（含 cause/solution）
- 接口层：参数错误统一映射为 BadRequestErr 等，交给全局错误处理输出结构化 JSON

国际化资源在：

- [zh-CN/error.toml](go-ddd-example/manifest/i18n/zh-CN/error.toml)
- [en/error.toml](go-ddd-example/manifest/i18n/en/error.toml)

## 9. 常用开发命令

以 [Makefile](go-ddd-example/Makefile) 为准：

```bash
make init
make tidy
make fmt
make vet
make test
make golangci
make gbuild
```

## 10. 新增一个业务模块：落点清单（按层扩展）

以“新增 foo 模块”为例，建议按以下顺序写代码（每一步都能独立编译运行）：

1) types：新增通用常量/错误码（若需要）
2) infra：
   - 新增 po（表结构）与 converter
   - 新增 repo 实现，并实现 domain/irepository 接口
   - 如需自动建表，把 po 加到 AutoMigrate
3) domain：
   - 新增 entity（领域对象）
   - 新增 irepository 接口（若还没有）
   - 新增 srv 或 aggregate（是否需要事务/跨 repo）
4) application：
   - 新增 dto（入参/出参）
   - 新增 assembler（dto <-> entity）
   - 新增 service（编排调用）
5) api：
   - HTTP：router + handler + 校验 + 返回
   - gRPC：proto + ginit 注册 + handler
   - MCP：新增 tool 定义与 handler（可选）
   - Event/Job：新增订阅/定时任务（可选）

## 生成代码的脚手架
https://github.com/jettjia/igo-ddd/tree/master/go-ddd-cli
可是生成所有的代码，包括：
- infra/repo
- domain
- application
- api

## pkg依赖
https://github.com/jettjia/igo-pkg
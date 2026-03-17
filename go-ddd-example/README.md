# 说明
此项目为DDD的演示项目，便于快速的开发

# 服务说明

公共服务的功能，主要处理登录、登出、token下发、token校验等

# 服务基础配置信息

端口：

| 外部端口 | 内部端口 | GRPC端口 | 接口前缀                         |
| -------- | -------- | -------- | -------------------------------- |
| 6666     | 6667     | 6668     | /api/v1/group-name/service-name/ |
|          |          |          |                                  |

存储配置：

| 存储类型  | 数据库                           | 表前缀/索引等举例                        |
| --------- | -------------------------------- | ---------------------------------------- |
| mysql     |                                  |                                          |
|           | databasename                     | public_cen_*                             |
|           |                                  |                                          |
| redis     | 1号库                            | public_cen_                              |
|           |                                  |                                          |
| mq        |                                  | databasename.public.center.folder.create |
|           |                                  |                                          |
| es_search | databasename_public_center_index | public_cen_                              |

# 端口和错误码

## 错误码和端口定义

| 模块大类 | 子模块        | 端口        | 中间三位错误码 | 案例      |
| -------- | ------------- | ----------- | -------------- | --------- |
| public-* | public-center | 对外：33000 | 000            | 400300100 |
|          |               | 对内：6667  | 001            | 400301100 |
|          |               | grpc：6668  | 002            | 400302100 |

## 错误码说明

错误码举例：

```json
{
    "code": 401300100,
    "cause": "token expired：xxxx",
    "message": "授权过期",
    "solution": "请刷新页面更新token或重新登录",
    "detail": {}
}
```

> - `code: 错误码（前三位：标准http错误码，中间三位为服务器特定码，后三位服务中自定义码）`
> - `cause: 错误原因，产生错误的具体原，比如错误的方法位置、行数`
> - `solution：符合国际化要求的针对当前错误的操作提示`
> - `message:  符合国际化要求的错误描述`
> - `detail:  错误码拓展信息，补充说明错误信息。`

错误码位数说明：

> 400300100
>
> 400 是 http 的状态码
>
> 300 是服务的 标记码
>
> 100 是服务定义的错误码

# 核心功能清单

登录

登出

token校验 （内部接口）

# 快速开始

## 配置文件

manifest/config/

## 国际化

manifest/i18n/

## 错误码定义

types/apierror

# 项目运行

## 配置

manifest/config/config.yaml

## 运行
- 直接运行： go run main.go
- 指定环境运行: go run main.go -env test
- 编译：go build


# 更多

## 支持多配置运行

```
go run main.go # 会使用默认的开发配置，即 debug模式
go run main.go -env test # 会使用test的配置，即测试环境
go run main.go -env release # 会使用release的配置，即正式环境
```

## 支持多协议并存

程序可以支持：http协议，又分为内部接口、外部接口。比如

外部接口: http://127.0.0.1:21800/api/pc/v1/sys/menu/1

内部接口：http://127.0.0.1:21801/private/pc/v1/sys/demo

grpc接口：127.0.0.1:21802


## 构建镜像
make gbuild (构建go二进制文件)
sudo make dbuild (构建docker镜像)
sudo make dpush (推送docker镜像)

# otel-uptrace 手册

## 安装uptrace依赖的环境

安装参考：https://uptrace.dev/get/install.html#binaries

```shell
docker-compose up -d
```


登录

http://10.4.7.71:14318/

The default login is `uptrace@localhost` with pass `uptrace`



## 官方案例运行

地址：https://github.com/uptrace/opentelemetry-go-extra

1.配置环境变量

```
OTEL_EXPORTER_JAEGER_ENDPOINT: http://10.4.7.71:14268/api/traces

UPTRACE_DSN: http://project1_secret_token@10.4.7.71:14317/1
```



```shell
export OTEL_EXPORTER_JAEGER_ENDPOINT=http://10.4.7.71:14268/api/traces
export UPTRACE_DSN=http://project1_secret_token@10.4.7.71:14317/1
```





2.运行项目

```
cd opentelemetry-go-extra/example/gin
go run main.go
```

打开：http://127.0.0.1:9999


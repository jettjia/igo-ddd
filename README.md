# go-ddd-multi

## 介绍

go-ddd-multi ，是实现ddd的多模块的案例。在业务开发中，会分common、biz1、biz2等；

```shell
├── common 公共模块，比如配置的读取，mysql的链接等
├── deploy 部署相关
│   ├── docker-compose 采用docker-compose的方式部署
│   ├── k8s 采用k8s的方式部署
├── system 系统服务，为具体的服务，可以有多个
```

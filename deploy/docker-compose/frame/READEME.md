# 安装mysql

## 目录结构和文件说明
```shell
├── conf
│   └── my.cnf  # MySQL配置文件
├── init
│   └── init.sql  # 初始化root 的密码
├── docker-compose.yml # docker-compose.yml文件
```


## 安装
```shell
docker-compose up -d
```

注意的坑！！！

第一次启动容器时候的账号密码，会记录在 ./db 目录下，所以当你想要在相同的目录下启动新容器（并设置了新密码），务必将./db 下所有文件删除掉，这样才能用docker-compose.yml中新密码连接数据库。

## mysql链接
root: root/admin123

admin: admin/admin123

## redis链接
pwd: admin123

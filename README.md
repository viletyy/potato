<!--
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:48:06
 * @FilePath: /potato/README.md
-->
# Potato

[![viletyy potato](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/viletyy/potato)

Go项目脚手架

基于gin、gorm、zap、cron的脚手架

本脚手架包含以下内容：

1. mvc结构。
2. swagger接口文档。
3. 配置、数据库、redis、日志、工具库、后台任务封装。
4. 单点登陆(jwt)。
5. 数据库版本管理

## 内容列表

- [Potato](#potato)
  - [内容列表](#内容列表)
  - [项目结构](#项目结构)
  - [安装](#安装)
  - [使用说明](#使用说明)
  - [相关仓库](#相关仓库)
  - [如何贡献](#如何贡献)
  - [使用许可](#使用许可)

## 项目结构
```
potato
├── config(配置目录)
├── docs(文档集合)
├── global(全局变量)
├── initialize(初始化函数)
├── internal(内部模块)
│   ├── controller(控制器层，用于存放控制器)
│   ├── dao(数据访问层，所有与数据相关等操作都会在dao层进行)
│   ├── job(后台任务)
│   ├── middleware(HTTP中间件)
│   ├── model(模型层，用于存放model对象)
│   ├── routers(路由相关逻辑处理)
│   └── service(项目核心业务逻辑)
├── migrations(数据库迁移文件)
├── pkg(项目相关等模块包)
├── scripts(各类构建、按照，分析等操作等脚本)
└── tmp(项目生成的临时文件)
```

## 安装

这个项目使用 [go](https://golang.org/) 、 [swag](https://github.com/swaggo/swag)、[docker](https://www.docker.com/)[下载地址](https://www.docker.com/products/docker-desktop)、[jaeger](https://www.jaegertracing.io/)。请确保你本地安装了它们。

go
```sh
$ tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin
```

swag安装
```sh
$ go get -u github.com/swaggo/swag/cmd/swag 
$ mv $GOPATH/bin/swag /usr/local/go/bin          
```

jaeger
```sh
docker run -d --name jaeger \                                                             
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:1.16
```

## 使用说明

```sh
# 启动项目
$ go build -o potato main.go
$ ./potato
# 生成api文档
$ swag init
```

## 相关仓库

- [Gin](https://github.com/gin-gonic/gin) — Web Framework
- [Gorm](https://github.com/jinzhu/gorm) — ORM
- [Swag](https://github.com/swaggo/swag) - RESTful API Doc
- [Cron](https://github.com/robfig/cron) - A cron library

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/viletyy/potato/issues/new) 或者提交一个 Pull Request。


## 使用许可

[MIT](LICENSE) © Viletyy

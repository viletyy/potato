<!--
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-06 10:36:10
 * @FilePath: /potato/README.md
-->
# Potato

[![viletyy potato](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/viletyy/potato)

Go项目脚手架

基于gin、gorm、zap的脚手架

本脚手架包含以下内容：

1. mvc结构。
2. swagger接口文档。
3. 配置、数据库、redis、日志、工具库封装。
4. 单点登陆(jwt)。

## 内容列表

- [安装](#安装)
- [使用说明](#使用说明)
- [相关仓库](#相关仓库)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可)

## 安装

这个项目使用 [go](https://golang.org/) 和 [swag](https://github.com/swaggo/swag)。请确保你本地安装了它们。

```sh
$ tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin
```

## 使用说明

```sh
$ go build -o potato main.go
$ ./potato
```

## 相关仓库

- [Gin](https://github.com/gin-gonic/gin) — Web Framework
- [Gorm](https://github.com/jinzhu/gorm) — ORM

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/viletyy/potato/issues/new) 或者提交一个 Pull Request。


## 使用许可

[MIT](LICENSE) © Viletyy

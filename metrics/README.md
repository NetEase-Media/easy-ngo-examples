# metrics demo

## 简介
此为指标观测(metrics)的demo, 用于理解和使用easy-ngo框架的metrics功能。

## 文件介绍
* app.toml 配置文件
* docker-compose.yaml docker-compose 配置文件
* Dockerfile docker镜像构建文件
* go.mod go.sum gomod相关
* main.go demo 代码
* setup.sql 建库建表语句

## 如何快速运行
1. 确保安装了 docker 以及 docker-compose
2. 启动服务 `docker-compose up`
3. 访问测试API [/hello](http://localhost:28888/hello)
4. 查看grafana视图 [Grafana UI](http://localhost:3000)



# kafka demo

## 简介
kafka 模块的使用方法用例

#### 文件描述
app.toml 程序启动的配置文件
main.go  主程序文件
docker-compose.yml 主程序需要连接kafka，因此通过docker-compse快速启动一组依赖以提供数据库服务

#### 执行过程
```
// 第一步先构建docker镜像, 并且启动以来服务kafka
// 执行命令
    docker-compose up

// 第二步编译主程序main.go
// 执行命令
    go build -o main .

// 第三步运行主程序
// 执行命令
    ./main -c app.toml
```
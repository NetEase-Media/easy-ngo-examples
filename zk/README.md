# zookeeper demo

## 简介
zookeeper 模块的使用方法用例

#### 文件描述
app.toml 程序启动的配置文件
main.go  主程序文件

#### 执行过程
```
// 第一步先构建docker镜像, 并且启动以来服务zookeeper
// 执行命令
    docker run -p 2181:2181 --restart always  zookeeper

// 第二步编译主程序main.go
// 执行命令
    go build -o main .

// 第三步运行主程序
// 执行命令
    ./main -c app.toml
```
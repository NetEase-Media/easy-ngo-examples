# redis demo

## 简介
redis 模块的使用方法用例

#### 文件描述
app.toml 程序启动的配置文件
main.go  主程序文件
Dockerfile 主程序需要连接redis，因此通过Dockerfile创建镜像，并且快速启动一组依赖以提供数据库服务
redis.conf 启动的redis容器的个性化配置，可以修改并且生效

#### 执行过程
```
// 第一步先构建docker镜像
// 执行命令
    docker build -t redis:demo .

// 第二步启动容器
// 执行命令
    docker run -p 6379:6379 redis:demo

// 第三步编译主程序main.go
// 执行命令
    go build -o main .

// 第四步运行主程序
// 执行命令
    ./main -c app.toml
```
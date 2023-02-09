# xxljob demo

## 简介
xxljob 模块的使用方法用例

#### 文件描述
app.toml 程序启动的配置文件
main.go  主程序文件
Dockerfile `main.go`的主程序也打包放到容器中执行
setup.sql 容器启动时初始化的数据库脚本

#### 执行过程
```
// 第一步先编译`docker-compose`
// 执行命令
    docker-compose build

// 第二步启动容器
// 执行命令
    docker-compose up

// 第三步进入`xxljob-admin`
// 执行命令
    在浏览器中输入：http://localhost:8080/xxl-job-admin
    输入账号/密码： admin/123456

// 第四步点击调用调用一次
// 可以在控制台观测到成功的结果
 $user 回调成功字眼
```
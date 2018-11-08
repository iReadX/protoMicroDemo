# proto 微服务架构Demo
## demo-service 服务端
```cmd
cd demo-service
go run main.go
```
## demo-cli 客户端
```cmd
cd demo-cli
go run cli.go
```

# docker 部署步骤   
此处只给出服务端的配置，客户端的改下配置即可，其他都是一样的
1. 编译生成linux运行的二进制文件
    ```gitBash
    cd demo-service
    GOOS=linux GOARCH=amd64 go build
    ```
    将会在`demo-service`文件夹中生成`demo-service`文件
2. 在`demo`目录下创建`Dockerfile`文件，代码如下：
    ```dockerfile
    # 使用最新版 debian 作为基础镜像，环境为Linux
    # alpine 为WEB基本运行环境
    FROM debian:latest
    
    # 在容器的根目录下创建 app 目录
    RUN mkdir /app
    
    # 将工作目录切换到 /app 下
    WORKDIR /app
    
    # 将微服务的服务端运行文件拷贝到 /app 下
    ADD demo-service /app/demo-service
    
    # 运行服务端
    CMD ["./demo-service/demo-service"]
    ```
3. 生成docker镜像
    ```cmd
    # 生成镜像名称：demo-service
    # 如需加上版本，直接在镜像名称后加上版本信息，如：demo-service:v1.0
    docker build -t demo-service .
    ```
4. docker运行镜像
    ```cmd
    docker run -p 50051:50051 demo-service
    ```
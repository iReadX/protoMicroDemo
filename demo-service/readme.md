# 服务端

在当前目录下运行脚本即可编译proto文件   
```cmd
protoc -I. --go_out=plugins=micro:. --micro_out=. proto/demo/demo.proto
```
## 以mdns方式启动
```cmd
MICRO_SERVER_ADDRESS=:50051 MICRO_REGISTRY=mdns go run main.go
```
## 打包docker镜像
```cmd
docker build -t demo-service .
```
## docker中mdns方式启动
```cmd
docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns demo-service
```
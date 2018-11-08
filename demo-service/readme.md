# 服务端

在当前目录下运行脚本即可编译proto文件   
```cmd
protoc -I. --go_out=plugins=grpc:. --micro_out=. proto/auditIntegral/auditIntegral.proto
```
# 服务端

在当前目录下运行脚本即可编译proto文件   
```cmd
protoc -I. --go_out=plugins=micro:. --micro_out=. proto/demo/demo.proto
```
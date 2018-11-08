build:
	cd demo-service
	# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
	GOOS=linux GOARCH=amd64 go build
	# 返回上级
	cd ..
	# 根据当前目录下的 Dockerfile 生成名为 consignment-service 的镜像
	docker build -t demo-service .
run:
	# 在 Docker alpine 容器的 50001 端口上运行 consignment-service 服务
	# 可添加 -d 参数将微服务放到后台运行
	docker run -p 50051:50051 demo-service
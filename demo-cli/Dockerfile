FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

# 将当前目录下的货物信息文件 demo.json 拷贝到 /app 目录下
ADD demo.json /app/demo.json
ADD demo-cli /app/demo-cli

CMD ["./demo-cli"]
buildGo:
	GOOS=linux GOARCH=amd64 go build
	docker build -t demo-service .
buildProto:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/auditIntegral/demo-service/proto/demo/proto/demo/demo.proto
run:
	docker run -p 50051:50051 demo-service
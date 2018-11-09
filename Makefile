build:
	cd demo-service
	GOOS=linux GOARCH=amd64 go build
	docker build -t demo-service .
run:
	docker run -p 50051:50051 demo-service
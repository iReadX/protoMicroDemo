build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t demo-cli .
run:
	docker run -e MICRO_REGISTRY=mdns demo-cli
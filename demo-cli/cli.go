package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"
	pb "protoMicroDemo/demo-service/proto/demo"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "demo.json"
)

// 读取consignment.json中记录的信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("demo.json file content error")
	}
	return consignment, nil
}

func main() {
	// 连接到gRPC服务器
	//conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("connect error: %v", err)
	//}
	//defer conn.Close()
	//conn,err := micro.new
	service := micro.NewService(micro.Name("go.micro.srv.demo"))
	service.Init()

	// 初始化gRPC客户端
	//clien := pb.NewShippingServiceClient(conn)
	// 创建微服务的客户端
	client := pb.NewShippingService("go.micro.srv.demo", service.Client())

	// 在命令行中指定新的货物信息json文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	// 调用RPC
	// 将货物存储到我们自己的仓库里
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	// 新货物是否托运成功
	log.Printf("created: %v", resp.Created)

	// 列出目前仓库中托运的货物
	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})

	if err != nil {
		log.Fatalf("failed to list consignment: %v", err)
	}

	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}
}

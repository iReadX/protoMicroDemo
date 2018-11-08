package main

import (
	pb "auditIntegral/demo-service/proto/demo"
	"context"
	"github.com/micro/go-micro"
	"log"
)

// 仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment                                   // 获取仓库中所有的货物
}

// 存放货物的仓库，实现了IRepository接口
type Repository struct {
	consignments []*pb.Consignment
}

// 创建
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	log.Printf("push consignment: %v", consignment)
	return consignment, nil
}

// 获取全部
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

/*
 * 定义微服务
 */
type service struct {
	repo Repository
}

/*
 * 实现 consignment.pb.go 中的 ShippingServiceHandler 接口
 * 使 service 作为 gRPC 的服务端
 */

// 托运新的货物
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	// 接收承运的货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil
	}
	resp = &pb.Response{Created: true, Consignment: consignment}
	return nil
}

// 获取目前所有托运的货物
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	allConsignments := s.repo.GetAll()
	resp = &pb.Response{Consignments: allConsignments}
	return nil
}

func main() {
	server := micro.NewService(
		// 必须和 demo.proto 中声明的 package 一致
		micro.Name("go.micro.srv.demo"),
		micro.Version("latest"),
	)

	// 解析命令行参数
	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

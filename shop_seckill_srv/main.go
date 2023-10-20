package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"shop_seckill_srv/handler"
	pb "shop_seckill_srv/proto"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := service.NewService(
		service.Name("shop_seckill_srv"),
		service.Address("127.0.0.1:6003"),
		service.Registry(consulReg),
	)

	// Register handler
	pb.RegisterShop_seckill_srvHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

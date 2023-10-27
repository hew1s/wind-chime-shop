package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"shop_seckill_srv/controller"
	_ "shop_seckill_srv/data_source"
	"shop_seckill_srv/proto/secpro"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := service.NewService(
		service.Name("shop_seckill_srv"),
		service.Address("127.0.0.1:6003"),
		service.Registry(consulReg),
	)
	shop_seckill_srv.RegisterSecKillHandler(srv.Server(), new(controller.SecKill))
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

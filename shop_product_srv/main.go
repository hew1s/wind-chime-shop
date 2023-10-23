package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"shop_product_srv/controller"
	_ "shop_product_srv/data_source"
	product "shop_product_srv/proto/product"
	seckill "shop_product_srv/proto/seckill"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := service.NewService(
		service.Name("shop_product_srv"),
		service.Address("127.0.0.1:6002"),
		service.Registry(consulReg),
	)

	product.RegisterProductsHandler(srv.Server(), new(controller.Products))
	seckill.RegisterSecKillsHandler(srv.Server(), new(controller.SecKills))
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

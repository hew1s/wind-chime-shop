package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"shop_product_srv/controller"
	_ "shop_product_srv/data_source"
	"shop_product_srv/proto/product"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := service.NewService(
		service.Name("shop_product_srv"),
		service.Address("127.0.0.1:6002"),
		service.Registry(consulReg),
	)

	shop_product_srv.RegisterProductsHandler(srv.Server(), new(controller.Products))
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

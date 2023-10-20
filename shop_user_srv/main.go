package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"shop_user_srv/controller"
	_ "shop_user_srv/data_source"
	"shop_user_srv/proto/admin_user"
	"shop_user_srv/proto/front_user"
)

func main() {
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	// Create service
	srv := service.NewService(
		service.Name("shop_user_srv"),
		service.Version("latest"),
		service.Address("127.0.0.1:6001"),
		service.Registry(consulReg),
	)

	// Register handler
	front_user.RegisterFrontUserHandler(srv.Server(), new(controller.FrontUser))
	admin_user.RegisterAdminUserHandler(srv.Server(), new(controller.AdminUser))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

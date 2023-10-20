package main

import (
	"github.com/gin-gonic/gin"
	"shop_web/middle_ware"
	all_router "shop_web/router"
)

func main() {
	//consulReg := consul.NewRegistry()
	r := gin.Default()
	r.Use(middle_ware.CrosMiddleWare)
	all_router.InitRouter(r)
	r.Run("127.0.0.1:6060")
}

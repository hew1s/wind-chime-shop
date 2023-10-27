package seckill

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	shop_seckill_srv "shop_seckill_srv/proto/secpro"
	"shop_web/utils"
)

var microSecProClient shop_seckill_srv.SecKillService

func registerSecProConsul() {
	consulReg := consul.NewRegistry()
	service := micro.NewService(
		micro.Address("127.0.0.1:6003"),
		micro.Registry(consulReg),
	)
	microSecProClient = shop_seckill_srv.NewSecKillService("shop_seckill_srv", service.Client())
}

// 秒杀
func SecKill(c *gin.Context) {
	id := c.PostForm("id")
	registerSecProConsul()
	rep, err := microSecProClient.FrontSecKill(context.TODO(), &shop_seckill_srv.SecKillRequest{
		Id: utils.IntStr(id),
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, rep)
}

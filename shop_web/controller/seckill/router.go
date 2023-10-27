package seckill

import (
	"github.com/gin-gonic/gin"
	"shop_web/controller/product"
	"shop_web/middle_ware"
)

func Router(r *gin.RouterGroup) {
	r.GET("/get_seckill_list", middle_ware.JwtTokenValid, GetSeckillList)
	r.GET("/get_products", middle_ware.JwtTokenValid, product.GetProductList)
	r.POST("/seckill_add", middle_ware.JwtTokenValid, SecKillAdd)
	r.POST("/seckill_del", middle_ware.JwtTokenValid, SecKillDel)
	r.GET("/seckill_to_edit", middle_ware.JwtTokenValid, GetSecKillById)
	r.POST("/seckill_do_edit", middle_ware.JwtTokenValid, UpSecKill)

	//前端列表
	r.GET("/front/get_seckill_list", GetFrontSeckillList)
	//前端详情
	r.GET("/front/seckill_detail", middle_ware.FrontJwtTokenValid, GetFrontSeckillDetail)

	//秒杀接口
	r.POST("/front/seckill", middle_ware.FrontJwtTokenValid, SecKill)
}

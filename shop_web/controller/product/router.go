package product

import (
	"github.com/gin-gonic/gin"
	"shop_web/middle_ware"
)

func Router(r *gin.RouterGroup) {
	r.GET("/get_product_list", middle_ware.JwtTokenValid, GetProductList)
	r.POST("/product_add", middle_ware.JwtTokenValid, ProductAdd)
	r.POST("/product_del", middle_ware.JwtTokenValid, ProductDel)
	r.GET("/to_product_edit", middle_ware.JwtTokenValid, GetProductById)
	r.POST("/do_product_edit", middle_ware.JwtTokenValid, EditProduct)
}

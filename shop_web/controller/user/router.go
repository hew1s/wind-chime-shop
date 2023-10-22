package user

import (
	"github.com/gin-gonic/gin"
	"shop_web/middle_ware"
)

func Router(r *gin.RouterGroup) {
	//用户端注册登录
	r.POST("/send_email", SendEmail)
	r.POST("/front_user_register", FrontUserRegister)
	r.POST("/front_user_login", FrontUserLogin)
	//管理端登录
	r.POST("/admin_login", AdminLogin)
	r.GET("/get_front_users", middle_ware.JwtTokenValid, FrontUserList)
}

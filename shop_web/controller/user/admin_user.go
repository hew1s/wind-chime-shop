package user

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_user_srv/proto/admin_user"
	"shop_web/utils"
)

var microAdminClient admin_user.AdminUserService

// 将注册服务封装成一个函数
func RegisterAdminConsul() {
	consulReg := consul.NewRegistry()
	service := micro.NewService(
		micro.Address("127.0.0.1:6001"),
		micro.Registry(consulReg),
	)
	microAdminClient = admin_user.NewAdminUserService("shop_user_srv", service.Client())
}

func AdminLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	RegisterAdminConsul()
	rep, err := microAdminClient.AdminUserLogin(context.TODO(), &admin_user.AdminUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	} else {
		admin_token, err := utils.GetToken(rep.UserName, utils.AdminUserExpireDuration, utils.AdminUserSecretKey)
		if err != nil {
			c.JSON(http.StatusOK,rep)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":        rep.Code,
			"msg":         rep.Msg,
			"admin_token": admin_token,
			"user_name":   rep.UserName,
		})
	}
}

package user

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_user_srv/proto/front_user"
	"shop_web/utils"
)

var microFrontClient front_user.FrontUserService

// 将注册服务封装成一个函数
func RegisterFrontConsul() {
	consulReg := consul.NewRegistry()
	service := micro.NewService(
		micro.Address("127.0.0.1:6001"),
		micro.Registry(consulReg),
	)
	microFrontClient = front_user.NewFrontUserService("shop_user_srv", service.Client())
}

// 发送邮箱验证码
func SendEmail(c *gin.Context) {
	email := c.PostForm("email")
	if utils.VerifyEmailFormat(email) { //校验完成进行grpc通信
		RegisterFrontConsul()
		rep, err := microFrontClient.FrontUserSendEmail(context.TODO(), &front_user.FrontUserMailRequest{Email: email})
		if err != nil {
			c.JSON(http.StatusOK, rep)
			return
		}
		c.JSON(http.StatusOK, rep)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
	}
}

// 用户注册
func FrontUserRegister(c *gin.Context) {
	email := c.PostForm("email")
	if utils.VerifyEmailFormat(email) {
		captche := c.PostForm("captche")
		password := c.PostForm("password")
		repassword := c.PostForm("repassword")
		if password != repassword {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "两次密码不一致！",
			})
		} else {
			//使用grpc通信
			RegisterFrontConsul()
			rep, err := microFrontClient.FrontUserRegister(context.TODO(), &front_user.FrontUserRequest{
				Email:    email,
				Code:     captche,
				Password: password,
			})
			if err != nil {
				c.JSON(http.StatusOK, rep)
				return
			} else {
				c.JSON(http.StatusOK, rep)
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
	}
}

// 用户登录
func FrontUserLogin(c *gin.Context) {
	mail := c.PostForm("mail")
	password := c.PostForm("password")
	is_ok := utils.VerifyEmailFormat(mail)
	if is_ok {
		RegisterFrontConsul()
		rep, err := microFrontClient.FrontUserLogin(context.TODO(), &front_user.FrontUserRequest{
			Email:    mail,
			Password: password,
		})
		if err != nil {
			c.JSON(http.StatusOK, rep)
			return
		} else {
			//生成token
			tokenStr, _ := utils.GetToken(rep.UserName, utils.FrontUserExpireDuration, utils.FrontUserSecretKey)
			c.JSON(http.StatusOK, gin.H{
				"code":     rep.Code,
				"msg":      rep.Msg,
				"token":    tokenStr,
				"username": rep.UserName,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "邮箱地址不正确",
	})
}

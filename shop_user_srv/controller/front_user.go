package controller

import (
	"context"
	"github.com/patrickmn/go-cache"
	"shop_user_srv/data_source"
	"shop_user_srv/models"
	"shop_user_srv/proto/front_user"
	"shop_user_srv/utils"
	"time"
)

type FrontUser struct {
}

var c = cache.New(60*time.Second, 10*time.Second)

// 注册服务
func (f *FrontUser) FrontUserRegister(ctx context.Context, in *front_user.FrontUserRequest, out *front_user.FrontUserResponse) error {
	//用户注册
	//获取数据
	email := in.Email
	captche := in.Code
	password := in.Password
	// 1.校验验证码是否正确
	code, is_ok := c.Get(email)
	if is_ok {
		if code != captche {
			out.Code = 500
			out.Msg = "邮箱验证码不正确"
		} else {
			//存入数据库
			front_user := models.FrontUser{
				Email:      email,
				Password:   utils.Md5pwd(password),
				Status:     1,
				CreateTime: time.Now(),
			}
			data_source.Db.Create(&front_user)
			out.Code = 200
			out.Msg = "注册成功，请登录"
		}
	} else {
		out.Code = 500
		out.Msg = "注册失败，请重新尝试"
	}
	return nil
}

// 发送验证码服务
func (f *FrontUser) FrontUserSendEmail(ctx context.Context, in *front_user.FrontUserMailRequest, out *front_user.FrontUserResponse) error {
	email := in.Email
	var front_user models.FrontUser
	var count int
	data_source.Db.Where("email = ?", email).Find(&front_user).Count(&count)
	if count < 1 {
		rand_num := utils.GenRandNum(6)
		utils.SendEmail(email, rand_num) //发送邮件
		//将数据存入缓存
		c.Set(email, rand_num, cache.DefaultExpiration)
		out.Code = 200
		out.Msg = "发送成功"
	} else {
		out.Code = 500
		out.Msg = "邮箱已存在"
	}
	return nil
}

// 登录服务
func (f *FrontUser) FrontUserLogin(ctx context.Context, in *front_user.FrontUserRequest, out *front_user.FrontUserResponse) error {
	email := in.Email
	password := in.Password
	password = utils.Md5pwd(password)
	var front_user models.FrontUser
	err := data_source.Db.Where("email =?", email).Find(&front_user).Error
	if err != nil {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return err
	}
	if password != front_user.Password {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return err
	}
	out.Code = 200
	out.Msg = "登录成功"
	out.UserName = front_user.Email
	return nil
}

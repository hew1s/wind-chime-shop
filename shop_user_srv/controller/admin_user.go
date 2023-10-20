package controller

import (
	"context"
	"shop_user_srv/data_source"
	"shop_user_srv/models"
	"shop_user_srv/proto/admin_user"
	"shop_user_srv/utils"
)

type AdminUser struct{}

func (a *AdminUser) AdminUserLogin(ctx context.Context, in *admin_user.AdminUserRequest, out *admin_user.AdminUserResponse) error {
	username := in.Username
	password := in.Password
	md5Pwd := utils.Md5pwd(password)
	var adminuser models.AdminUser
	err := data_source.Db.Where("user_name =?", username).First(&adminuser).Error
	if md5Pwd != adminuser.Password {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return err
	}
	out.Code = 200
	out.Msg = "登录成功"
	out.UserName = username
	return nil
}

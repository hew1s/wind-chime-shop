package controller

import (
	"context"
	"shop_user_srv/data_source"
	"shop_user_srv/models"
	"shop_user_srv/proto/admin_user"
	"shop_user_srv/utils"
	"strconv"
)

type AdminUser struct{}

// 管理员登录
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

// 获取用户列表
func (a *AdminUser) FrontUserList(ctx context.Context, in *admin_user.FrontUsersRequest, out *admin_user.FrontUsersResponse) error {
	currentPage := in.CurrentPage
	pageSize := in.PageSize
	var users []models.FrontUser
	var count int32
	count = int32(data_source.Db.Find(&users).RowsAffected)
	result := data_source.Db.Limit(pageSize).Offset(pageSize * (currentPage - 1)).Find(&users)
	var front_users []*admin_user.FrontUserList
	if result.Error != nil {
		out.Code = 500
		out.Msg = "服务错误，请重试"
		out.FrontUsers = front_users
		return result.Error
	}
	if result.RowsAffected < 1 {
		out.Code = 500
		out.Msg = "未查询到数据"
		out.FrontUsers = front_users
		return result.Error
	}
	for _, v := range users {
		var user admin_user.FrontUserList
		user.Email = v.Email
		user.Status = strconv.Itoa(v.Status)
		user.Desc = v.Desc
		user.CreateTime = v.CreateTime.Format("2006-01-02 15:04:05")
		front_users = append(front_users, &user)
	}
	out.Code = 200
	out.Msg = "success"
	out.Current = currentPage
	out.Pagesize = pageSize
	out.Total = count
	out.FrontUsers = front_users
	return nil
}

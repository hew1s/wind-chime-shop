package utils

import (
	"fmt"
	"github.com/astaxie/beego/utils"
	"strings"
)

// 使用beego下utils下的NewEMail
func SendEmail(to_email, msg string) {
	username := "2461098644@qq.com" // 发送者的邮箱地址
	password := "puaprdwegpnmdjdj"  // 授权密码
	host := "smtp.qq.com"           // 邮件协议
	port := "587"                   // 端口号

	emailConfig := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%s}`, username, password, host, port)
	emailConn := utils.NewEMail(emailConfig) // beego下的
	emailConn.From = strings.TrimSpace(username)
	emailConn.To = []string{strings.TrimSpace(to_email)}
	emailConn.Subject = "用户注册邮箱验证码"
	//注意这里我们发送给用户的是激活请求地址
	emailConn.Text = msg
	emailConn.Send()
}

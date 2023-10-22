package middle_ware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_web/utils"
	"strings"
)

func JwtTokenValid(c *gin.Context) {
	auth_header := c.Request.Header.Get("Authorization")
	if len(auth_header) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带Token",
		})
		c.Abort()
		return
	}
	auths := strings.Split(auth_header, " ")
	bearer := auths[0]
	token := auths[1]
	if len(token) == 0 || len(bearer) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "Token格式错误",
		})
		c.Abort()
		return
	}
	user, err := utils.AuthToken(token, utils.AdminUserSecretKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "无效的Token",
		})
		c.Abort()
		return
	}
	c.Set("userName", user.UserName)
	c.Next()
}

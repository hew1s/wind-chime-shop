package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_web/utils"
	"strings"
)

func TestToken(c *gin.Context) {
	username := "admin"
	token, _ := utils.GetToken(username, utils.FrontUserExpireDuration, utils.FrontUserSecretKey)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func TestValid(c *gin.Context) {
	header_auth := c.Request.Header.Get("Authorization")
	list_token := strings.Split(header_auth, " ")
	token, err := utils.AuthToken(list_token[1], utils.FrontUserSecretKey)
	if err != nil {
		return
	}
	fmt.Println(token.UserName)
	c.JSON(200, "success")
}

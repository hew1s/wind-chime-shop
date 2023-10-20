package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserToken struct {
	jwt.StandardClaims
	//自定义信息
	UserName string `json:"user_name"`
}

// 前端用户过期时间
var FrontUserExpireDuration = time.Hour * 1
var FrontUserSecretKey = []byte("front_user_token")

// 后端用户过期时间
var AdminUserExpireDuration = time.Hour * 2
var AdminUserSecretKey = []byte("admin_user_token")

// 生成token
func GetToken(UserName string, expireDuration time.Duration, secret_key []byte) (string, error) {
	user := UserToken{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			Issuer:    "micro_gin_vue",
		},
		UserName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(secret_key)
}

// 验证token
func AuthToken(tokenStr string, secretKey []byte) (*UserToken, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserToken{}, func(token *jwt.Token) (key interface{}, err error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, is_ok := token.Claims.(*UserToken)
	//验证token
	if is_ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token valid err")
}

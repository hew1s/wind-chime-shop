# 秒杀商城

go-micro（consul）+gin+vue

将user、product、seckill功能分别注册为微服务

后端使用gin框架来实现

## JWT方式进行身份认证

```go
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
```



## 登录、注册

### 邮箱验证码发送与缓存

将验证码缓存

```go
var c = cache.New(60*time.Second, 10*time.Second)
//将验证码插入
c.Set(email, rand_num, cache.DefaultExpiration)
//取出验证码
code, is_ok := c.Get(email)
```

### 注册

1. 发送验证码
2. 验证邮箱与验证码是否对应
3. 验证用户是否注册
4. 将用户信息插入数据库

代码

```go
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
```

### 用户登录与管理员登录

用户登录

密码经过加密后与数据库比对

返回由用户名参与的token与username

```go
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
```

管理员登录

```go
func (a *AdminUser) AdminUserLogin(ctx context.Context, in *admin_user.AdminUserRequest, out *admin_user.AdminUserResponse) error {
	username := in.Username
	password := in.Password
	md5Pwd := utils.Md5pwd(password)
	var adminuser models.AdminUser
	result := data_source.Db.Where("user_name =?", username).Where("password =?", md5Pwd).Find(&adminuser)
	if result.RowsAffected < 1 {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return result.Error
	}
	out.Code = 200
	out.Msg = "登录成功"
	out.UserName = username
	return nil
}
```


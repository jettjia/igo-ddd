package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("wUk7E@fwowuyek")

// CustomerInfo 返回给前端的数据内容
type CustomerInfo struct {
	UserId   uint64 ` json:"user_id"`  // ID
	Phone    string ` json:"phone"`    // 手机号码
	Username string ` json:"username"` // 手机号码
	RoleId   uint   `json:"role_id"`   // 角色id
}

// CustomClaims jwtClaims定义
type CustomClaims struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}

// Token 返回给客户端的token内容
type Token struct {
	AccessToken string `json:"access_token"` // token
	ExpiresIn   int64  `json:"expires_in"`   // 过期时间
	Username    string `json:"username"`     // 用户名
	Phone       string `json:"phone"`        // 手机号码
	UserId      uint64 `json:"user_id"`      // 用户id
}

// CreateToken 获取jwt token
func CreateToken(info CustomerInfo) (*Token, error) {
	expiresAt := time.Now().Add(time.Minute * 60).Unix()

	claims := &CustomClaims{
		&jwt.StandardClaims{

			ExpiresAt: expiresAt,
			Issuer:    "GasPc",
		},
		"level1",
		info,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtTokenStr, err := tokenClaims.SignedString(jwtSecret) //该方法内部生成签名字符串，再用于获取完整、已签名的token

	var token Token
	if err != nil {
		return &token, err
	}

	token.AccessToken = jwtTokenStr
	token.ExpiresIn = expiresAt
	token.Username = info.Username
	token.Phone = info.Phone
	token.UserId = info.UserId

	return &token, nil
}

// ParseToken token 校验
func ParseToken(tokenString string) (*CustomClaims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

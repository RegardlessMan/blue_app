package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const expireTime = 2 * time.Hour

var CustomSecret = []byte("我是最帅的")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	Username             string `json:"username"`
	UserId               int64  `json:"user_id"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

func GenerateToken(userId int64, username string) (string, error) {
	// 创建一个我们自己的声明
	c := CustomClaims{
		Username: username,
		UserId:   userId, // 自定义字段
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)), // 过期时间
			Issuer:    "bluebell",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, err
}

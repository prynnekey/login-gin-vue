package utils

import (
	"login-vue/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// token过期时间
var TokenExpireDuration = time.Now().Add(7 * 24 * time.Hour)

// 用于生成token的密钥
var jwtKey = []byte("a_secret_crect")

// 根据user发放token
func GetToken(user models.User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: TokenExpireDuration.Unix(), // token过期时间
			IssuedAt:  time.Now().Unix(),          // 发放时间
			Issuer:    "prynnekey",                // 发放者
			Subject:   "user token",               // 主题
		},
	}

	//使用指定的签名方法创建签名对象
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := _token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

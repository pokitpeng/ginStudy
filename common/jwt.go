package common

import (
	"ginStudy/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Cliams struct {
	UserId uint
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user model.User) (string, error) {
	// 过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	cliams := &Cliams{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pokit",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Cliams, error) {
	cliams := &Cliams{}
	token, err := jwt.ParseWithClaims(tokenString, cliams, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, cliams, err
}

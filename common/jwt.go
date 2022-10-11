package common

import (
	"github.com/dgrijalva/jwt-go"
	"mygin.com/mygin/model"
	"time"
)

var jwtkey = []byte("a_secret_crect")

// Claims 要求
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken 颁发 token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "mygin.com",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析 token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})

	return token, claims, err
}

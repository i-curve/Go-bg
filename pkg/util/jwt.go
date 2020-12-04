package util

import (
	"template/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claim := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "curve",
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaim.SignedString(jwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claims, nil
		}
	}
	return nil, err
}

/*
 * @Date: 2021-03-22 17:16:46
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-23 09:24:08
 * @FilePath: /potato/utils/jwt.go
 */
package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/viletyy/potato/global"
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func GenerateToken(userId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Minute * 30)
	loginUUID := uuid.New().String()

	claims := CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "data_govern",
			Id:        loginUUID,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenClaims.SignedString([]byte(global.GO_CONFIG.App.JwtSecret))
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("General Token Error: %v", err))
	}
	global.GO_LOG.Info("General Token:" + tokenString)
	result, err := global.GO_REDIS.Set("login:"+loginUUID, userId, 1*time.Hour).Result()
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("General Token Set To Redis Error: %v", err))
	}
	global.GO_LOG.Info("Set Token To Redis:" + result)

	return tokenString, err
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.GO_CONFIG.App.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

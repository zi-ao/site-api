package users

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zi-ao/site-api/pkg/config"
	"github.com/zi-ao/site-api/pkg/logger"
	"strconv"
	"time"
)

type standardClaims struct {
	jwt.StandardClaims
	*User
}

// GenerateToken 生成 Token
func GenerateToken(user *User, duration time.Duration) (string, error) {
	expireTime := time.Now().Add(duration)
	stdClaims := jwt.StandardClaims{
		Audience:  "WEB", // 使用端 WEB，WAP，APP，MINI_PROGRAM
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        strconv.Itoa(int(user.ID)),
		Issuer:    config.Global.Name,
		Subject:   "API",
	}

	//user.Password = ""
	uClaims := standardClaims{
		StandardClaims: stdClaims,
		User:           user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Global.Key))
	if err != nil {
		logger.Fatal("config is wrong, can not generate jwt")
	}
	return tokenString, err
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := standardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Global.Key), nil
	})
	if err != nil {
		return nil, err
	}
	return claims.User, err
}

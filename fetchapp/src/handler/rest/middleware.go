package restserver

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Auth struct {
	opt AuthOptions
}

type AuthOptions struct {
	Secret string
}

type AuthInfo struct {
	Phone string
	Role  string
}

func InitAuth(opt AuthOptions) *Auth {
	return &Auth{
		opt: opt,
	}
}

func (a *Auth) GetAuthInfo(tokenString string) (AuthInfo, error) {
	var result AuthInfo
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.opt.Secret), nil
	})

	if err != nil {
		return result, err
	}
	if !token.Valid {
		return result, errors.New("Unauthorized")
	}
	claims := token.Claims.(jwt.MapClaims)

	exp := claims["exp"].(float64)
	_exp := time.Unix(int64(exp), 0)
	if _exp.Before(time.Now()) {
		return result, errors.New("Unauthorized")
	}

	return AuthInfo{
		Phone: claims["phone"].(string),
		Role:  claims["role"].(string),
	}, nil
}

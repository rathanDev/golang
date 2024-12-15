package model

import (
	"github.com/dgrijalva/jwt-go"
	"svc/constant"
)

type Claims struct {
	UserId   string        `json:"userId"`
	Username string        `json:"username"`
	UserRole constant.Role `json:"userRole"`
	jwt.StandardClaims
}

package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}
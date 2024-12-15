package service

import (
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
	"svc/model"
)

// move to config
var appName = "APP_NAME"
var jwtSecretKey = []byte("secret-key")

func GenerateJwt(userId string, username string) (string, error) {
	log.Println("GenerateJwt ", userId, username)
	expirationTime := time.Now().Add(24 * time.Hour) // expires in 24 hours
	claims := &model.Claims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    appName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		log.Println("Error signing token")
		return "", err
	}

	return signedToken, nil
}

func ValidateJwt(tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		log.Println("Err parsing token", err)
		return nil, err
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

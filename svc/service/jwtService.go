package service

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"svc/config"
	"svc/model"
	"time"
)

// move to config
var appName = "APP_NAME"
var jwtSecretKey = []byte(config.GetConfig().Api.Key)

func GenerateJwt(user model.User) (string, error) {
	log.Println("GenerateJwt ", user)
	expirationTime := time.Now().Add(24 * time.Hour) // expires in 24 hours
	claims := &model.Claims{
		UserId:   user.ID,
		Username: user.Username,
		UserRole: user.Role,
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

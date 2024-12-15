package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"svc/config"
	"svc/model"
)

var jwtSecretKey = []byte(config.GetConfig().Api.Key)

func TokenAuthMiddleware() gin.HandlerFunc {
	log.Println("TokenAuthMiddleware")
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		log.Println("tokenStr", tokenString)

		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			c.Set("userId", claims.UserId)
			c.Set("username", claims.Username)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}

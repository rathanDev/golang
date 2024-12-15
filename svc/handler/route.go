package handler

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(route *gin.Engine) {
	route.GET("/users", GetUsers)
	route.POST("/signUp", SignUp)
	route.POST("/login", Login)
}

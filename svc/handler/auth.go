package handler

import (
	"svc/service"

	"github.com/gin-gonic/gin"
)

func SetRoutes(route *gin.Engine) {
	route.GET("/users", service.GetUsers)
	route.POST("/signUp", service.SignUp)
	route.POST("/login", service.Login)
}

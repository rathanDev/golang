package handler

import (
	"svc/middleware"
	"svc/service"

	"github.com/gin-gonic/gin"
)

func SetRoutes(route *gin.Engine) {
	route.POST("/signUp", service.SignUp)
	route.POST("/login", service.Login)

	authorized := route.Group("/")
	authorized.Use(middleware.TokenAuthMiddleware())
	authorized.GET("/users", service.GetUsers)
}

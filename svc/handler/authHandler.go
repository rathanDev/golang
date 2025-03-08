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
	{
		// authorized.Use(middleware.TokenAuthMiddleware())
		// authorized.GET("/users", service.GetUsers)

		// Admin-only route
		authorized.GET("/admin", middleware.TokenAuthMiddleware(secretKey, []constant.Role{constant.ADMIN}), handlers.AdminHandler)

		// Doctor and Admin route
		auth.GET("/doctor", middleware.AuthMiddleware(secretKey, []constant.Role{constant.DOC, constant.ADMIN}), handlers.DoctorHandler)

		// Patient and Admin route
		auth.GET("/patient", middleware.AuthMiddleware(secretKey, []constant.Role{constant.PAT, constant.ADMIN}), handlers.PatientHandler)


	}

}

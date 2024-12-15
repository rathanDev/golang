package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"svc/handler"
)

func main() {
	fmt.Println("Hello svc")

	route := gin.Default()

	route.GET("/users", handler.GetUsers)
	route.POST("/signUp", handler.SignUp)

	route.Run(":8080")
}

package main

import (
	"fmt"
	"svc/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello svc")

	route := gin.Default()

	route.GET("/users", handler.GetUsers)

	route.Run(":8080")
}

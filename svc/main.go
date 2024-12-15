package main

import (
	"github.com/gin-gonic/gin"
	"svc/handler"
)

func main() {
	route := gin.Default()
	handler.SetRoutes(route)
	route.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"svc/config"
	"svc/handler"
)

func main() {
	config.LoadConfig()
	route := gin.Default()
	handler.SetRoutes(route)
	route.Run(":8080")
}

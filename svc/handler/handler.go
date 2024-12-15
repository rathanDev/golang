package handler

import (
	"net/http"
	"svc/model"

	"github.com/gin-gonic/gin"
)

var users = []model.User{
	{ID: 1, Name: "Jana"},
	{ID: 2, Name: "Rathan"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

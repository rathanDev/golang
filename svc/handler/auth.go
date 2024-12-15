package handler

import (
	"fmt"
	"net/http"
	"svc/model"

	"github.com/gin-gonic/gin"
)

// var usersCredentials = []model.UserCredential{}

var users = []model.User{
	{ID: 1, Name: "Jana"},
	{ID: 2, Name: "Rathan"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func SignUp(c *gin.Context) {
	var newUser model.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Some err",
		})
		return
	}
	users = append(users, newUser)
	fmt.Println("users", users)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

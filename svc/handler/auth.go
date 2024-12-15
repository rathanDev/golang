package handler

import (
	"fmt"
	"log"
	"net/http"
	"svc/model"

	"github.com/gin-gonic/gin"
)

// var usersCredentials = []model.UserCredential{}

var users = []model.User{}
// {
// 	{ID: 1, Name: "Jana"},
// 	{ID: 2, Name: "Rathan"},
// }

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func SignUp(c *gin.Context) {
	log.Println("SignUp")
	var newUser model.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusConflict, gin.H{
			"error": "Some err",
		})
		return
	}
	log.Println("NewUser", newUser)
	users = append(users, newUser)
	fmt.Println("users", users)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

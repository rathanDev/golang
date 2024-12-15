package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"svc/model"
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

func Login(c *gin.Context) {
	var userCred model.UserCredential
	log.Println("Login")
	// In a real application, authenticate the user (this is just an example)
	//  username := c.PostForm("username")
	//  password := c.PostForm("password")
	err := c.BindJSON(&userCred)
	if err != nil {
		log.Println("Err at login", err)
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}
	log.Println("Login", userCred)

	for _, user := range users {
		if user.Username == userCred.Username && user.Password == userCred.Password {
			log.Println("User found", user)
			c.JSON(http.StatusCreated, gin.H{
				"message": "LoggedIn",
			})
			return
		}
	}
	log.Println("Login failed")
	c.JSON(http.StatusConflict, gin.H{
		"error": "Login failed",
	})
}

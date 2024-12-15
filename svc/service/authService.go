package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"svc/model"
)

// var usersCredentials = []model.UserCredential{}

var users = []model.User{{ID: "001", Name: "Jana", Username: "jana", Password: "Pwd1"}}

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

	var loggedInUser model.User

	for _, user := range users {
		if user.Username == userCred.Username && user.Password == userCred.Password {
			log.Println("User found", user)
			loggedInUser = user
			break
		}
	}

	if loggedInUser == (model.User{}) {
		log.Println("Login failed")
		c.JSON(http.StatusConflict, gin.H{
			"error": "Login failed",
		})
		return
	}

	tokenStr, err := GenerateJwt(loggedInUser.ID, loggedInUser.Username)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Unable to generate jwt",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": tokenStr,
	})
}

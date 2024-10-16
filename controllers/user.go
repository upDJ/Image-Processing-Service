package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/upDJ/Image-Processing-Service/models"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users := append(models.Users, newUser)
	c.JSON(200, users)
}

func LoginUser(c *gin.Context) {
	var userCredentials models.User

	if err := c.BindJSON(&userCredentials); err != nil {
		return
	}

	for _, user := range models.Users {
		isValidUsername := user.Username == userCredentials.Username
		isValidPassword := user.Password == userCredentials.Password

		if isValidUsername && isValidPassword {
			c.String(200, "Successful Login!")
			return
		}
	}

	c.String(404, "User Not Found")
}

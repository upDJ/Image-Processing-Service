package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/upDJ/Image-Processing-Service/models"
)

func createUser(c *gin.Context) []models.User {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return models.Users
	}

	users := append(models.Users, newUser)
	return users
}

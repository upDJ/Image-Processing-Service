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

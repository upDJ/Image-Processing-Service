package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/upDJ/Image-Processing-Service/controllers"
)

func createUser(r *gin.Engine) {
	r.POST("/register", controllers.CreateUser)
}

func SetupUserRoutes(r *gin.Engine) *gin.Engine {
	createUser(r)
  return r
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/upDJ/Image-Processing-Service/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.Ping)
	return r
}

package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/upDJ/Image-Processing-Service/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.PingController)
	return r
}

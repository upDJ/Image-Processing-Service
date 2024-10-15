package main

import (
	"github.com/upDJ/Image-Processing-Service/routers"
)

func main() {
	router := routers.SetupRouter()
	router = routers.SetupUserRoutes(router)
	// setup user routes
	router.Run("localhost:8080")
}

package main

import (
  "github.com/upDJ/Image-Processing-Service/routers"
)

func main() {
    router := routers.SetupRouter()
    // router.POST("/register")
  
    router.Run("localhost:8080")
}


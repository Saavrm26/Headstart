package main

import (
	"github.com/gin-gonic/gin"
  "hs/headstart/Routers"
)

func main() {
	router := gin.Default() 
  Routers.CatalougeRoutes(router)
	router.Run(":8080")
}

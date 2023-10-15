package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
  "hs/headstart/Routers"
)

func main() {
	router := gin.Default() 
  Routers.CatalougeRoutes(router)
	fmt.Printf("Starting server at port 8080")
	router.Run(":8080")
}

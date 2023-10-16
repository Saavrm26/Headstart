package main

import (
	"fmt"
	"hs/headstart/Routers"

	"github.com/gin-gonic/gin"
)



func main() {
  fmt.Println(nCatalougeValidatorFreeList)
  initCatalougeValidatorFreeList()

	router := gin.Default() 
  Routers.CatalougeRoutes(router)
	router.Run(":8080")
}
